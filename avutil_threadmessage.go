// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/threadmessage.h>

typedef void (*av_thread_message_free_func)(void *msg);
*/
import "C"
import "unsafe"

// AVThreadMessageQueue
type AVThreadMessageQueue C.struct_AVThreadMessageQueue

// AVThreadMessageFlags
type AVThreadMessageFlags = C.enum_AVThreadMessageFlags

const (
	AV_THREAD_MESSAGE_NONBLOCK = AVThreadMessageFlags(C.AV_THREAD_MESSAGE_NONBLOCK)
)

// AvThreadMessageQueueAlloc allocates a new message queue.
func AvThreadMessageQueueAlloc[T Integer](mq **AVThreadMessageQueue, nelem, elsize T) int32 {
	return (int32)(C.av_thread_message_queue_alloc(
		(**C.struct_AVThreadMessageQueue)(unsafe.Pointer(mq)), (C.uint)(nelem), (C.uint)(elsize)))
}

// AvThreadMessageQueueFree frees a message queue.
func AvThreadMessageQueueFree(mq **AVThreadMessageQueue) {
	C.av_thread_message_queue_free((**C.struct_AVThreadMessageQueue)(unsafe.Pointer(mq)))
}

// AvThreadMessageQueueSend sends a message on the queue.
func AvThreadMessageQueueSend(mq *AVThreadMessageQueue, msg CVoidPointer, flags uint32) int32 {
	return (int32)(C.av_thread_message_queue_send((*C.struct_AVThreadMessageQueue)(mq),
		VoidPointer(msg), (C.uint)(flags)))
}

// AvThreadMessageQueueRecv receives a message from the queue.
func AvThreadMessageQueueRecv(mq *AVThreadMessageQueue, msg CVoidPointer, flags uint32) int32 {
	return (int32)(C.av_thread_message_queue_recv((*C.struct_AVThreadMessageQueue)(mq),
		VoidPointer(msg), (C.uint)(flags)))
}

// AvThreadMessageQueueSetErrSend sets the sending error code.
func AvThreadMessageQueueSetErrSend(mq *AVThreadMessageQueue, err int32) {
	C.av_thread_message_queue_set_err_send((*C.struct_AVThreadMessageQueue)(mq), (C.int)(err))
}

// AvThreadMessageQueueSetErrRecv set the receiving error code.
func AvThreadMessageQueueSetErrRecv(mq *AVThreadMessageQueue, err int32) {
	C.av_thread_message_queue_set_err_recv((*C.struct_AVThreadMessageQueue)(mq), (C.int)(err))
}

// typedef void (*av_thread_message_free_func)(void *msg);
type AvThreadMessageFreeFunc = C.av_thread_message_free_func

// AvThreadMessageQueueSetFreeFunc sets the optional free message callback function which will be called if an
// operation is removing messages from the queue.
func AvThreadMessageQueueSetFreeFunc(mq *AVThreadMessageQueue, freeFunc AvThreadMessageFreeFunc) {
	C.av_thread_message_queue_set_free_func((*C.struct_AVThreadMessageQueue)(mq),
		(C.av_thread_message_free_func)(freeFunc))
}

// AvThreadMessageFlush flushes the message queue.
func AvThreadMessageFlush(mq *AVThreadMessageQueue) {
	C.av_thread_message_flush((*C.struct_AVThreadMessageQueue)(mq))
}
