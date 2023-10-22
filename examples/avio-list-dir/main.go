package main

import (
	"fmt"
	"os"

	ffmpeg "github.com/qrtc/ffmpeg-dev-go"
)

func typeString(_type int32) string {
	switch _type {
	case ffmpeg.AVIO_ENTRY_DIRECTORY:
		return "<DIR>"
	case ffmpeg.AVIO_ENTRY_FILE:
		return "<FILE>"
	case ffmpeg.AVIO_ENTRY_BLOCK_DEVICE:
		return "<BLOCK DEVICE>"
	case ffmpeg.AVIO_ENTRY_CHARACTER_DEVICE:
		return "<CHARACTER DEVICE>"
	case ffmpeg.AVIO_ENTRY_NAMED_PIPE:
		return "<PIPE>"
	case ffmpeg.AVIO_ENTRY_SYMBOLIC_LINK:
		return "<LINK>"
	case ffmpeg.AVIO_ENTRY_SOCKET:
		return "<SOCKET>"
	case ffmpeg.AVIO_ENTRY_SERVER:
		return "<SERVER>"
	case ffmpeg.AVIO_ENTRY_SHARE:
		return "<SHARE>"
	case ffmpeg.AVIO_ENTRY_WORKGROUP:
		return "<WORKGROUP>"
	case ffmpeg.AVIO_ENTRY_UNKNOWN:
	default:
		break
	}
	return "<UNKNOWN>"
}

func listOp(inputDir string) int32 {
	var entry *ffmpeg.AVIODirEntry
	var ctx *ffmpeg.AVIODirContext
	var cnt, ret int32
	var filemode, uidAndGid string

	if ret = ffmpeg.AvIOOpenDir(&ctx, inputDir, nil); ret < 0 {
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot open directory: %s.\n", ffmpeg.AvErr2str(ret))
		goto fail
	}

	for {
		if ret = ffmpeg.AvIOReadDir(ctx, &entry); ret < 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_ERROR, "Cannot list directory: %s.\n", ffmpeg.AvErr2str(ret))
			goto fail
		}
		if entry == nil {
			break
		}
		if entry.GetFilemode() == -1 {
			filemode = "???"
		} else {
			filemode = fmt.Sprintf("%3d", entry.GetFilemode())
		}
		uidAndGid = fmt.Sprintf("%d(%d)", entry.GetUserId(), entry.GetGroupId())
		if cnt == 0 {
			ffmpeg.AvLog(nil, ffmpeg.AV_LOG_INFO, "%-9s %12s %30s %10s %s %16s %16s %16s\n",
				"TYPE", "SIZE", "NAME", "UID(GID)", "UGO", "MODIFIED",
				"ACCESSED", "STATUS_CHANGED")
		}
		ffmpeg.AvLog(nil, ffmpeg.AV_LOG_INFO, "%-9s %12d %30s %10s %s %d %d %d\n",
			typeString(entry.GetType()),
			entry.GetSize(),
			entry.GetName(),
			uidAndGid,
			filemode,
			entry.GetModificationTimestamp(),
			entry.GetAccessTimestamp(),
			entry.GetStatusChangeTimestamp())
		ffmpeg.AvIOFreeDirectoryEntry(&entry)
		cnt++
	}

fail:
	ffmpeg.AvIOCloseDir(&ctx)
	return ret
}

func usage(programName string) {
	fmt.Fprintf(os.Stderr, "usage: %s input_dir\n"+
		"API example program to show how to list files in directory "+
		"accessed through AVIOContext.\n", programName)
	os.Exit(1)
}

func main() {
	var ret int32

	ffmpeg.AvLogSetLevel(ffmpeg.AV_LOG_DEBUG)

	if len(os.Args) < 2 {
		usage(os.Args[0])
		os.Exit(1)
	}

	ffmpeg.AvFormatNetworkInit()

	ret = listOp(os.Args[1])

	ffmpeg.AvFormatNetworkDeinit()

	if ret < 0 {
		os.Exit(1)
	}
}
