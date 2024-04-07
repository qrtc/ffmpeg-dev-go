package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

const (
	OUTPUT_BIT_RATE = 96000
	OUTPUT_CHANNELS = 2
)

// Open an input file and the required decoder.
func openInputFile(fileName string) (inputFormatContext *ffmpeg.AVFormatContext,
	inputCodecContext *ffmpeg.AVCodecContext, ret int32) {

	// Open the input file to read from it.
	if ret = ffmpeg.AvFormatOpenInput(&inputFormatContext, fileName, nil, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open input file '%s' (error '%s')\n", fileName, ffmpeg.AvErr2str(ret))
		return nil, nil, ret
	}

	// Get information on the input file (number of streams etc.).
	if ret = ffmpeg.AvFormatFindStreamInfo(inputFormatContext, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open find stream info (error '%s')\n", ffmpeg.AvErr2str(ret))
		ffmpeg.AvFormatCloseInput(&inputFormatContext)
		return nil, nil, ret
	}

	// Make sure that there is only one stream in the input file.
	if inputFormatContext.GetNbStreams() != 1 {
		fmt.Fprintf(os.Stderr, "Expected one audio input stream, but found %d\n", inputFormatContext.GetNbStreams())
		ffmpeg.AvFormatCloseInput(&inputFormatContext)
		return nil, nil, ffmpeg.AVERROR_EXIT
	}

	// Find a decoder for the audio stream.
	inputCodec := ffmpeg.AvCodecFindDecoder(inputFormatContext.GetStreams()[0].GetCodecpar().GetCodecId())
	if inputCodec == nil {
		fmt.Fprintf(os.Stderr, "Could not find input codec\n")
		ffmpeg.AvFormatCloseInput(&inputFormatContext)
		return nil, nil, ffmpeg.AVERROR_EXIT
	}

	// Allocate a new decoding context.
	avctx := ffmpeg.AvCodecAllocContext3(inputCodec)
	if avctx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate a decoding context\n")
		ffmpeg.AvFormatCloseInput(&inputFormatContext)
		return nil, nil, ffmpeg.AVERROR(syscall.ENOMEM)
	}

	// Initialize the stream parameters with demuxer information.
	if ret = ffmpeg.AvCodecParametersToContext(avctx, inputFormatContext.GetStreams()[0].GetCodecpar()); ret < 0 {
		ffmpeg.AvFormatCloseInput(&inputFormatContext)
		ffmpeg.AvCodecFreeContext(&avctx)
		return nil, nil, ret
	}

	// Open the decoder for the audio stream to use it later.
	if ret = ffmpeg.AvCodecOpen2(avctx, inputCodec, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open input codec (error '%s')\n", ffmpeg.AvErr2str(ret))
		ffmpeg.AvCodecFreeContext(&avctx)
		ffmpeg.AvFormatCloseInput(&inputFormatContext)
		return nil, nil, ret
	}

	// Save the decoder context for easier access later.
	inputCodecContext = avctx

	return inputFormatContext, inputCodecContext, ret
}

// Open an output file and the required encoder.
func openOutputFile(filename string, inputCodecContext *ffmpeg.AVCodecContext) (outputFormatContext *ffmpeg.AVFormatContext,
	outputCodecContext *ffmpeg.AVCodecContext, ret int32) {

	var outputIOContext *ffmpeg.AVIOContext
	var outputCodec *ffmpeg.AVCodec
	var stream *ffmpeg.AVStream
	var avctx *ffmpeg.AVCodecContext

	// Open the output file to write to it.
	if ret = ffmpeg.AvIOOpen(&outputIOContext, filename, ffmpeg.AVIO_FLAG_WRITE); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open output file '%s' (error '%s')\n", filename, ffmpeg.AvErr2str(ret))
		return nil, nil, ret
	}

	// Create a new format context for the output container format.
	if outputFormatContext = ffmpeg.AvFormatAllocContext(); outputFormatContext == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate output format context\n")
		return nil, nil, ffmpeg.AVERROR(syscall.ENOMEM)
	}

	// Associate the output file (pointer) with the container format context.
	outputFormatContext.SetPb(outputIOContext)

	// Guess the desired container format based on the file extension.
	outputFormatContext.SetOformat(ffmpeg.AvGuessFormat(ffmpeg.NIL, filename, ffmpeg.NIL))
	if outputFormatContext.GetOformat() == nil {
		fmt.Fprintf(os.Stderr, "Could not find output file format\n")
		goto cleanup
	}

	outputFormatContext.SetUrl(filename)
	if len(outputFormatContext.GetUrl()) == 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate url.\n")
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto cleanup
	}

	// Find the encoder to be used by its name.
	if outputCodec = ffmpeg.AvCodecFindEncoder(ffmpeg.AV_CODEC_ID_AAC); outputCodec == nil {
		fmt.Fprintf(os.Stderr, "Could not find an AAC encoder.\n")
		goto cleanup
	}

	// Create a new audio stream in the output file container.
	if stream = ffmpeg.AvFormatNewStream(outputFormatContext, nil); stream == nil {
		fmt.Fprintf(os.Stderr, "Could not create new stream\n")
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto cleanup
	}

	if avctx = ffmpeg.AvCodecAllocContext3(outputCodec); avctx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate an encoding context\n")
		ret = ffmpeg.AVERROR(syscall.ENOMEM)
		goto cleanup
	}

	// Set the basic encoder parameters.
	// The input file's sample rate is used to avoid a sample rate conversion.
	ffmpeg.AvChannelLayoutDefault(avctx.GetChLayoutAddr(), OUTPUT_CHANNELS)
	avctx.SetSampleRate(inputCodecContext.GetSampleRate())
	avctx.SetSampleFmt(outputCodec.GetSampleFmts()[0])
	avctx.SetBitRate(OUTPUT_BIT_RATE)

	// Allow the use of the experimental AAC encoder.
	avctx.SetStrictStdCompliance(ffmpeg.FF_COMPLIANCE_EXPERIMENTAL)

	// Set the sample rate for the container.
	stream.SetTimeBase(ffmpeg.AvMakeQ(inputCodecContext.GetSampleRate(), 1))

	// Some container formats (like MP4) require global headers to be present.
	// Mark the encoder so that it behaves accordingly.
	if (outputFormatContext.GetOformat().GetFlags() & ffmpeg.AVFMT_GLOBALHEADER) != 0 {
		avctx.SetFlags(avctx.GetFlags() | ffmpeg.AV_CODEC_FLAG_GLOBAL_HEADER)
	}

	// Open the encoder for the audio stream to use it later.
	if ret = ffmpeg.AvCodecOpen2(avctx, outputCodec, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open output codec (error '%s')\n", ffmpeg.AvErr2str(ret))
		goto cleanup
	}

	ret = ffmpeg.AvCodecParametersFromContext(stream.GetCodecpar(), avctx)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not initialize stream parameters\n")
		goto cleanup
	}

	// Save the encoder context for easier access later.
	outputCodecContext = avctx

	return outputFormatContext, outputCodecContext, 0

cleanup:
	ffmpeg.AvCodecFreeContext(&avctx)
	ffmpeg.AvIOClosep(outputFormatContext.GetPbAddr())
	if ret > 0 {
		ret = ffmpeg.AVERROR_EXIT
	}
	return nil, nil, ret
}

// Initialize one data packet for reading or writing.
func initPacket() (packet *ffmpeg.AVPacket, ret int32) {
	if packet = ffmpeg.AvPacketAlloc(); packet == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate packet\n")
		return nil, ffmpeg.AVERROR(syscall.ENOMEM)
	}
	return packet, 0
}

// Initialize one audio frame for reading from the input file.
func initInputFrame() (frame *ffmpeg.AVFrame, ret int32) {
	if frame = ffmpeg.AvFrameAlloc(); frame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate input frame\n")
		return nil, ffmpeg.AVERROR(syscall.ENOMEM)
	}
	return frame, 0
}

// Initialize the audio resampler based on the input and output codec settings.
// If the input and output sample formats differ, a conversion is required
// libswresample takes care of this, but requires initialization.
func initResampler(inputCodecContext, outputCodecContext *ffmpeg.AVCodecContext) (
	resampleContext *ffmpeg.SwrContext, ret int32) {

	// Create a resampler context for the conversion.
	// Set the conversion parameters.
	// Default channel layouts based on the number of channels
	// are assumed for simplicity (they are sometimes not detected
	// properly by the demuxer and/or decoder).
	if ret = ffmpeg.SwrAllocSetOpts2(&resampleContext,
		outputCodecContext.GetChLayoutAddr(),
		outputCodecContext.GetSampleFmt(),
		outputCodecContext.GetSampleRate(),
		inputCodecContext.GetChLayoutAddr(),
		inputCodecContext.GetSampleFmt(),
		inputCodecContext.GetSampleRate(),
		0, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate resample context\n")
		return nil, ffmpeg.AVERROR(syscall.ENOMEM)
	}

	if outputCodecContext.GetSampleRate() != inputCodecContext.GetSampleRate() {
		panic("resample has to be handled differently")
	}

	// Open the resampler with the specified parameters.
	if ret = ffmpeg.SwrInit(resampleContext); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open resample context\n")
		ffmpeg.SwrFree(&resampleContext)
		return nil, ret
	}
	return resampleContext, 0
}

// Initialize a FIFO buffer for the audio samples to be encoded.
func initFifo(outputCodecContext *ffmpeg.AVCodecContext) (fifo *ffmpeg.AVAudioFifo, ret int32) {
	// Create the FIFO buffer based on the specified output sample format
	if fifo = ffmpeg.AvAudioFifoAlloc(outputCodecContext.GetSampleFmt(),
		outputCodecContext.GetChLayoutAddr().GetNbChannels(), 1); fifo == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate FIFO\n")
		return nil, ffmpeg.AVERROR(syscall.ENOMEM)
	}
	return fifo, 0
}

// Write the header of the output file container.
func writeOutputFileHeader(outputFormatContext *ffmpeg.AVFormatContext) int32 {
	if ret := ffmpeg.AvFormatWriteHeader(outputFormatContext, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not write output file header (error '%s')\n", ffmpeg.AvErr2str(ret))
		return ret
	}
	return 0
}

// Decode one audio frame from the input file.
func decodeAudioFrame(frame *ffmpeg.AVFrame,
	inputFormatContext *ffmpeg.AVFormatContext,
	inputCodecContext *ffmpeg.AVCodecContext) (dataPresent, finished, ret int32) {
	// Packet used for temporary storage.
	inputPacket, ret := initPacket()
	if ret < 0 {
		return 0, 0, ret
	}

	// Read one audio frame from the input file into a temporary packet.
	if ret = ffmpeg.AvReadFrame(inputFormatContext, inputPacket); ret < 0 {
		// If we are at the end of the file, flush the decoder below.
		if ret == ffmpeg.AVERROR_EOF {
			finished = 1
		} else {
			fmt.Fprintf(os.Stderr, "Could not read frame (error '%s')\n", ffmpeg.AvErr2str(ret))
			goto cleanup
		}
	}

	// Send the audio frame stored in the temporary packet to the decoder.
	// The input audio stream decoder is used to do this.
	if ret = ffmpeg.AvCodecSendPacket(inputCodecContext, inputPacket); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not send packet for decoding (error '%s')\n", ffmpeg.AvErr2str(ret))
		goto cleanup
	}

	// Receive one frame from the decoder.
	ret = ffmpeg.AvCodecReceiveFrame(inputCodecContext, frame)
	switch {
	// If the decoder asks for more data to be able to decode a frame,
	// return indicating that no data is present.
	case ret == ffmpeg.AVERROR(syscall.EAGAIN):
		ret = 0
		goto cleanup
	// If the end of the input file is reached, stop decoding.
	case ret == ffmpeg.AVERROR_EOF:
		finished = 1
		ret = 0
		goto cleanup
	case ret < 0:
		fmt.Fprintf(os.Stderr, "Could not decode frame (error '%s')\n", ffmpeg.AvErr2str(ret))
		goto cleanup
	// Default case: Return decoded data.
	default:
		dataPresent = 1
		goto cleanup
	}

cleanup:
	ffmpeg.AvPacketFree(&inputPacket)
	return dataPresent, finished, ret
}

// Initialize a temporary storage for the specified number of audio samples.
// The conversion requires temporary storage due to the different format.
// The number of audio samples to be allocated is specified in frame_size.
func initConvertedSamples(outputCodecContext *ffmpeg.AVCodecContext,
	frameSize int32) (convertedInputSamples **uint8, ret int32) {

	// Allocate as many pointers as there are audio channels.
	// Each pointer will later point to the audio samples of the corresponding
	// channels (although it may be NULL for interleaved formats).
	if convertedInputSamples = (**uint8)(ffmpeg.AvCalloc(outputCodecContext.GetChLayoutAddr().GetNbChannels(),
		unsafe.Sizeof(*convertedInputSamples))); convertedInputSamples == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate converted input sample pointers\n")
		return nil, ffmpeg.AVERROR(syscall.ENOMEM)
	}

	// Allocate memory for the samples of all channels in one consecutive
	// block for convenience.
	if ret = ffmpeg.AvSamplesAlloc(convertedInputSamples, nil,
		outputCodecContext.GetChLayoutAddr().GetNbChannels(),
		frameSize,
		outputCodecContext.GetSampleFmt(), 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate converted input samples (error '%s')\n", ffmpeg.AvErr2str(ret))
		ffmpeg.AvFreep(&convertedInputSamples)
		return nil, ret
	}
	return convertedInputSamples, 0
}

// Convert the input audio samples into the output sample format.
// The conversion happens on a per-frame basis, the size of which is
// specified by frame_size.
func convertSamples(inputData, convertedData **uint8,
	frameSize int32, resampleContext *ffmpeg.SwrContext) (ret int32) {
	if ret = ffmpeg.SwrConvert(resampleContext, convertedData, frameSize, inputData, frameSize); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not convert input samples (error '%s')\n", ffmpeg.AvErr2str(ret))
		return ret
	}
	return 0
}

// Add converted input audio samples to the FIFO buffer for later processing.
func addSamplesToFifo(fifo *ffmpeg.AVAudioFifo, convertedInputSamples **uint8, frameSize int32) int32 {
	// Make the FIFO as large as it needs to be to hold both,
	// the old and the new samples.
	if ret := ffmpeg.AvAudioFifoRealloc(fifo, ffmpeg.AvAudioFifoSize(fifo)+frameSize); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not reallocate FIFO\n")
		return ret
	}

	// Store the new samples in the FIFO buffer.
	if ret := ffmpeg.AvAudioFifoWrite(fifo, convertedInputSamples, frameSize); ret < frameSize {
		fmt.Fprintf(os.Stderr, "Could not write data to FIFO\n")
		return ffmpeg.AVERROR_EXIT
	}
	return 0
}

// Read one audio frame from the input file, decode, convert and store it in the FIFO buffer.
func readDecodeConvertAndStore(fifo *ffmpeg.AVAudioFifo,
	inputFormatContext *ffmpeg.AVFormatContext,
	inputCodecContext *ffmpeg.AVCodecContext,
	outputCodecContext *ffmpeg.AVCodecContext,
	resamplerContext *ffmpeg.SwrContext) (finished, ret int32) {

	// Temporary storage of the input samples of the frame read from the file.
	var inputFrame *ffmpeg.AVFrame
	// Temporary storage for the converted input samples.
	var convertedInputSamples **uint8
	var dataPresent int32
	ret = ffmpeg.AVERROR_EXIT

	// Initialize temporary storage for one input frame.
	if inputFrame, ret = initInputFrame(); ret != 0 {
		goto cleanup
	}
	// Decode one frame worth of audio samples.
	if dataPresent, finished, ret = decodeAudioFrame(inputFrame,
		inputFormatContext, inputCodecContext); ret != 0 {
		goto cleanup
	}
	// If we are at the end of the file and there are no more samples
	// in the decoder which are delayed, we are actually finished.
	// This must not be treated as an error.
	if finished != 0 {
		ret = 0
		goto cleanup
	}
	// If there is decoded data, convert and store it.
	if dataPresent != 0 {
		// Initialize the temporary storage for the converted input samples.
		if convertedInputSamples, ret = initConvertedSamples(outputCodecContext,
			inputFrame.GetNbSamples()); ret != 0 {
			goto cleanup
		}

		// Convert the input samples to the desired output sample format.
		// This requires a temporary storage provided by converted_input_samples.
		if ret = convertSamples(inputFrame.GetExtendedData(),
			convertedInputSamples, inputFrame.GetNbSamples(), resamplerContext); ret != 0 {
			goto cleanup
		}

		// Add the converted input samples to the FIFO buffer for later processing.
		if ret = addSamplesToFifo(fifo, convertedInputSamples,
			inputFrame.GetNbSamples()); ret != 0 {
			goto cleanup
		}
		ret = 0
	}
	ret = 0

cleanup:
	if convertedInputSamples != nil {
		ffmpeg.AvFreep(&convertedInputSamples)
	}
	ffmpeg.AvFrameFree(&inputFrame)

	return finished, ret
}

// Initialize one input frame for writing to the output file.
func initOutputFrame(outputcodecContext *ffmpeg.AVCodecContext,
	frameSize int32) (frame *ffmpeg.AVFrame, ret int32) {

	// Create a new frame to store the audio samples.
	if frame = ffmpeg.AvFrameAlloc(); frame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate output frame\n")
		return nil, ffmpeg.AVERROR_EXIT
	}

	// Set the frame's parameters, especially its size and format.
	// av_frame_get_buffer needs this to allocate memory for the
	// audio samples of the frame.
	// Default channel layouts based on the number of channels
	// are assumed for simplicity.
	frame.SetNbSamples(frameSize)
	ffmpeg.AvChannelLayoutCopy(frame.GetChLayoutAddr(), outputcodecContext.GetChLayoutAddr())
	frame.SetFormat(outputcodecContext.GetSampleFmt())
	frame.SetSampleRate(outputcodecContext.GetSampleRate())

	// Allocate the samples of the created frame. This call will make
	// sure that the audio frame can hold as many samples as specified.
	if ret = ffmpeg.AvFrameGetBuffer(frame, 0); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate output frame samples (error '%s')\n", ffmpeg.AvErr2str(ret))
		ffmpeg.AvFrameFree(&frame)
		return nil, ret
	}

	return frame, 0
}

// Global timestamp for the audio frames.
var pts int64

// Encode one frame worth of audio to the output file.
func encodeAudioFrame(frame *ffmpeg.AVFrame,
	outputFormatContext *ffmpeg.AVFormatContext,
	outputCodecContext *ffmpeg.AVCodecContext) (dataPresent, ret int32) {
	// Packet used for temporary storage.
	var outputPacket *ffmpeg.AVPacket

	if outputPacket, ret = initPacket(); ret < 0 {
		return dataPresent, ret
	}

	// Set a timestamp based on the sample rate for the container.
	if frame != nil {
		frame.SetPts(pts)
		pts += int64(frame.GetNbSamples())
	}

	// Send the audio frame stored in the temporary packet to the encoder.
	// The output audio stream encoder is used to do this.
	ret = ffmpeg.AvCodecSendFrame(outputCodecContext, frame)
	if ret == ffmpeg.AVERROR_EOF {
		ret = 0
		goto cleanup
	} else if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not send packet for encoding (error '%s')\n", ffmpeg.AvErr2str(ret))
		goto cleanup
	}

	// Receive one encoded frame from the encoder.
	ret = ffmpeg.AvCodecReceivePacket(outputCodecContext, outputPacket)
	// If the encoder asks for more data to be able to provide an
	// encoded frame, return indicating that no data is present.
	if ret == ffmpeg.AVERROR(syscall.EAGAIN) {
		ret = 0
		goto cleanup
	} else if ret == ffmpeg.AVERROR_EOF {
		// If the last frame has been encoded, stop encoding.
		ret = 0
		goto cleanup
	} else if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not encode frame (error '%s')\n", ffmpeg.AvErr2str(ret))
		goto cleanup
	} else {
		// Default case: Return encoded data.
		dataPresent = 1
	}

	// Write one audio frame from the temporary packet to the output file.
	if dataPresent != 0 {
		if ret = ffmpeg.AvWriteFrame(outputFormatContext, outputPacket); ret < 0 {
			fmt.Fprintf(os.Stderr, "Could not write frame (error '%s')\n", ffmpeg.AvErr2str(ret))
			goto cleanup
		}
	}

cleanup:
	ffmpeg.AvPacketFree(&outputPacket)
	return dataPresent, ret
}

// Load one audio frame from the FIFO buffer, encode and write it to the output file.
func loadEncodeAndWrite(fifo *ffmpeg.AVAudioFifo,
	outputFormatContext *ffmpeg.AVFormatContext,
	outputCodecContext *ffmpeg.AVCodecContext) (ret int32) {
	// Temporary storage of the output samples of the frame written to the file.
	var outputFrame *ffmpeg.AVFrame
	// Use the maximum number of possible samples per frame.
	// If there is less than the maximum possible frame size in the FIFO
	// buffer use this number. Otherwise, use the maximum possible frame size.
	frameSize := ffmpeg.FFMIN(ffmpeg.AvAudioFifoSize(fifo), outputCodecContext.GetFrameSize())

	// Initialize temporary storage for one output frame.
	if outputFrame, ret = initOutputFrame(outputCodecContext, frameSize); ret != 0 {
		return ffmpeg.AVERROR_EXIT
	}

	// Read as many samples from the FIFO buffer as required to fill the frame.
	// The samples are stored in the frame temporarily.
	if ret = ffmpeg.AvAudioFifoRead(fifo, outputFrame.GetData(), frameSize); ret < frameSize {
		fmt.Fprintf(os.Stderr, "Could not read data from FIFO\n")
		ffmpeg.AvFrameFree(&outputFrame)
		return ffmpeg.AVERROR_EXIT
	}

	// Encode one frame worth of audio samples.
	if _, ret = encodeAudioFrame(outputFrame,
		outputFormatContext, outputCodecContext); ret != 0 {
		ffmpeg.AvFrameFree(&outputFrame)
		return ffmpeg.AVERROR_EXIT
	}
	ffmpeg.AvFrameFree(&outputFrame)
	return 0
}

// Write the trailer of the output file container.
func writeOutputFileTrailer(outputFormatContext *ffmpeg.AVFormatContext) int32 {
	if ret := ffmpeg.AvWriteTrailer(outputFormatContext); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not write output file trailer (error '%s')\n", ffmpeg.AvErr2str(ret))
		return ret
	}
	return 0
}

func main() {
	var ret int32 = ffmpeg.AVERROR_EXIT
	var inputFormatContext *ffmpeg.AVFormatContext
	var inputCodecContext *ffmpeg.AVCodecContext
	var outputFormatContext *ffmpeg.AVFormatContext
	var outputCodecContext *ffmpeg.AVCodecContext
	var resampleContext *ffmpeg.SwrContext
	var fifo *ffmpeg.AVAudioFifo

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stdout, "Usage: %s <input file> <output file>\n", os.Args[0])
		os.Exit(1)
	}

	// Open the input file for reading.
	if inputFormatContext, inputCodecContext, ret = openInputFile(os.Args[1]); ret != 0 {
		goto cleanup
	}
	// Open the output file for writing.
	if outputFormatContext, outputCodecContext, ret = openOutputFile(os.Args[2],
		inputCodecContext); ret != 0 {
		goto cleanup
	}
	// Initialize the resampler to be able to convert audio sample formats.
	if resampleContext, ret = initResampler(inputCodecContext, outputCodecContext); ret != 0 {
		goto cleanup
	}
	// Initialize the FIFO buffer to store audio samples to be encoded.
	if fifo, ret = initFifo(outputCodecContext); ret != 0 {
		goto cleanup
	}
	// Write the header of the output file container.
	if ret = writeOutputFileHeader(outputFormatContext); ret != 0 {
		goto cleanup
	}

	// Loop as long as we have input samples to read or output samples
	// to write; abort as soon as we have neither.
	for {
		// Use the encoder's desired frame size for processing.
		var outputFrameSize int32 = outputCodecContext.GetFrameSize()
		var finished int32

		// Make sure that there is one frame worth of samples in the FIFO
		// buffer so that the encoder can do its work.
		// Since the decoder's and the encoder's frame size may differ, we
		// need to FIFO buffer to store as many frames worth of input samples
		// that they make up at least one frame worth of output samples.
		for ffmpeg.AvAudioFifoSize(fifo) < outputFrameSize {
			// Decode one frame worth of audio samples, convert it to the
			// output sample format and put it into the FIFO buffer.
			if finished, ret = readDecodeConvertAndStore(fifo, inputFormatContext,
				inputCodecContext, outputCodecContext, resampleContext); ret != 0 {
				goto cleanup
			}

			// If we are at the end of the input file, we continue
			// encoding the remaining audio samples to the output file.
			if finished != 0 {
				break
			}
		}

		// If we have enough samples for the encoder, we encode them.
		// At the end of the file, we pass the remaining samples to
		// the encoder.
		for ffmpeg.AvAudioFifoSize(fifo) >= outputFrameSize ||
			(finished != 0 && ffmpeg.AvAudioFifoSize(fifo) > 0) {
			// Take one frame worth of audio samples from the FIFO buffer,
			// encode it and write it to the output file.
			if ret = loadEncodeAndWrite(fifo, outputFormatContext, outputCodecContext); ret != 0 {
				goto cleanup
			}
		}

		// If we are at the end of the input file and have encoded
		// all remaining samples, we can exit this loop and finish.
		if finished != 0 {
			var dataWritten int32
			// Flush the encoder as it may have delayed frames.
			for ok := true; ok; ok = dataWritten != 0 {
				if dataWritten, ret = encodeAudioFrame(nil,
					outputFormatContext, outputCodecContext); ret != 0 {
					goto cleanup
				}
			}
			break
		}
	}

	// Write the trailer of the output file container.
	if ret = writeOutputFileTrailer(outputFormatContext); ret != 0 {
		goto cleanup
	}
	ret = 0

cleanup:
	if fifo != nil {
		ffmpeg.AvAudioFifoFree(fifo)
	}
	ffmpeg.SwrFree(&resampleContext)
	if outputCodecContext != nil {
		ffmpeg.AvCodecFreeContext(&outputCodecContext)
	}
	if outputFormatContext != nil {
		ffmpeg.AvIOClosep(outputFormatContext.GetPbAddr())
		ffmpeg.AvFormatFreeContext(outputFormatContext)
	}
	if inputCodecContext != nil {
		ffmpeg.AvCodecFreeContext(&inputCodecContext)
	}
	if inputFormatContext != nil {
		ffmpeg.AvFormatCloseInput(&inputFormatContext)
	}

	os.Exit(int(ret))
}
