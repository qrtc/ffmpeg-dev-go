// Copyright (c) 2023 QRTC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ffmpeg

/*
#include <libavutil/executor.h>

typedef int (*av_executor_priority_higher_func)(const AVTask *a, const AVTask *b);

typedef int (*av_executor_ready_func)(const AVTask *t, void *user_data);

typedef int (*av_executor_run_func)(AVTask *t, void *local_context, void *user_data);

*/
import "C"
import "unsafe"

// AVExecutor
type AVExecutor C.struct_AVExecutor

// AVTask
type AVTask C.struct_AVTask

// GetNext gets `AVTask.next` value.
func (t *AVTask) GetNext() *AVTask {
	return (*AVTask)(t.next)
}

// SetNext sets `AVTask.next` value.
func (t *AVTask) SetNext(v *AVTask) {
	t.next = (*C.struct_AVTask)(v)
}

// GetNextAddr gets `AVTask.next` address.
func (t *AVTask) GetNextAddr() **AVTask {
	return (**AVTask)(unsafe.Pointer(&t.next))
}

// AVTaskCallbacks
type AVTaskCallbacks C.struct_AVTaskCallbacks

// typedef int (*av_executor_priority_higher_func)(const AVTask *a, const AVTask *b);
type AvExecutorPriorityHigherFunc = C.av_executor_priority_higher_func

// typedef int (*av_executor_ready_func)(const AVTask *t, void *user_data);
type AvExecutorReadyFunc = C.av_executor_ready_func

// typedef int (*av_executor_run_func)(AVTask *t, void *local_context, void *user_data);
type AvExecutorRunFunc = C.av_executor_run_func

// GetUserData gets `AVTaskCallbacks.user_data` value.
func (cbs *AVTaskCallbacks) GetUserData() unsafe.Pointer {
	return cbs.user_data
}

// SetUserData sets `AVTaskCallbacks.user_data` value.
func (cbs *AVTaskCallbacks) SetUserData(v CVoidPointer) {
	cbs.user_data = VoidPointer(v)
}

// GetUserDataAddr gets `AVTaskCallbacks.user_data` address.
func (cbs *AVTaskCallbacks) GetUserDataAddr() *unsafe.Pointer {
	return (*unsafe.Pointer)(&cbs.user_data)
}

// GetLocalContextSize gets `AVTaskCallbacks.local_context_size` value.
func (cbs *AVTaskCallbacks) GetLocalContextSize() int32 {
	return (int32)(cbs.local_context_size)
}

// SetLocalContextSize sets `AVTaskCallbacks.local_context_size` value.
func (cbs *AVTaskCallbacks) SetLocalContextSize(v int32) {
	cbs.local_context_size = (C.int)(v)
}

// GetLocalContextSizeAddr gets `AVTaskCallbacks.local_context_size` address.
func (cbs *AVTaskCallbacks) GetLocalContextSizeAddr() *int32 {
	return (*int32)(&cbs.local_context_size)
}

// GetPriorityHigher gets `AVTaskCallbacks.priority_higher` value.
func (cbs *AVTaskCallbacks) GetPriorityHigher() AvExecutorPriorityHigherFunc {
	return (AvExecutorPriorityHigherFunc)(cbs.priority_higher)
}

// SetPriorityHigher sets `AVTaskCallbacks.priority_higher` value.
func (cbs *AVTaskCallbacks) SetPriorityHigher(v AvExecutorPriorityHigherFunc) {
	cbs.priority_higher = (C.av_executor_priority_higher_func)(v)
}

// GetPriorityHigherAddr gets `AVTaskCallbacks.priority_higher` address.
func (cbs *AVTaskCallbacks) GetPriorityHigherAddr() *AvExecutorPriorityHigherFunc {
	return (*AvExecutorPriorityHigherFunc)(&cbs.priority_higher)
}

// GetReady gets `AVTaskCallbacks.ready` value.
func (cbs *AVTaskCallbacks) GetReady() AvExecutorReadyFunc {
	return (AvExecutorReadyFunc)(cbs.ready)
}

// SetReady sets `AVTaskCallbacks.ready` value.
func (cbs *AVTaskCallbacks) SetReady(v AvExecutorReadyFunc) {
	cbs.ready = (C.av_executor_ready_func)(v)
}

// GetReadyAddr gets `AVTaskCallbacks.ready` address.
func (cbs *AVTaskCallbacks) GetReadyAddr() *AvExecutorReadyFunc {
	return (*AvExecutorReadyFunc)(&cbs.ready)
}

// GetRun gets `AVTaskCallbacks.run` value.
func (cbs *AVTaskCallbacks) GetRun() AvExecutorRunFunc {
	return (AvExecutorRunFunc)(cbs.run)
}

// SetRun sets `AVTaskCallbacks.run` value.
func (cbs *AVTaskCallbacks) SetRun(v AvExecutorRunFunc) {
	cbs.run = (C.av_executor_run_func)(v)
}

// GetRunAddr gets `AVTaskCallbacks.run` address.
func (cbs *AVTaskCallbacks) GetRunAddr() *AvExecutorRunFunc {
	return (*AvExecutorRunFunc)(&cbs.run)
}

// AvExecutorAlloc allocates executor.
func AvExecutorAlloc(callbacks *AVTaskCallbacks, threadCount int32) *AVExecutor {
	return (*AVExecutor)(C.av_executor_alloc((*C.struct_AVTaskCallbacks)(callbacks), (C.int)(threadCount)))
}

// AvExecutorFree frees executor.
func AvExecutorFree(e **AVExecutor) {
	C.av_executor_free((**C.struct_AVExecutor)(unsafe.Pointer(e)))
}

// AvExecutorExecute adds task to executor.
func AvExecutorExecute(e *AVExecutor, t *AVTask) {
	C.av_executor_execute((*C.struct_AVExecutor)(e), (*C.struct_AVTask)(t))
}
