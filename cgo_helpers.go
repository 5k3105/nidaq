// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 26 Sep 2017 20:00:22 MDT.
// By https://git.io/c-for-go. DO NOT EDIT.

package nidaqmx

/*
#include "NIDAQmx.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

func (x EveryNSamplesEventCallbackPtr) PassRef() (ref *C.DAQmxEveryNSamplesEventCallbackPtr, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if everyNSamplesEventCallbackPtrB317614EFunc == nil {
		everyNSamplesEventCallbackPtrB317614EFunc = x
	}
	return (*C.DAQmxEveryNSamplesEventCallbackPtr)(C.DAQmxEveryNSamplesEventCallbackPtr_b317614e), nil
}

func (x EveryNSamplesEventCallbackPtr) PassValue() (ref C.DAQmxEveryNSamplesEventCallbackPtr, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if everyNSamplesEventCallbackPtrB317614EFunc == nil {
		everyNSamplesEventCallbackPtrB317614EFunc = x
	}
	return (C.DAQmxEveryNSamplesEventCallbackPtr)(C.DAQmxEveryNSamplesEventCallbackPtr_b317614e), nil
}

func NewEveryNSamplesEventCallbackPtrRef(ref unsafe.Pointer) *EveryNSamplesEventCallbackPtr {
	return (*EveryNSamplesEventCallbackPtr)(ref)
}

//export everyNSamplesEventCallbackPtrB317614E
func everyNSamplesEventCallbackPtrB317614E(cTaskHandle C.TaskHandle, cEveryNsamplesEventType C.int32, cNSamples C.uInt32, cCallbackData unsafe.Pointer) C.int32 {
	if everyNSamplesEventCallbackPtrB317614EFunc != nil {
		TaskHandleb317614e := (unsafe.Pointer)(cTaskHandle)
		EveryNsamplesEventTypeb317614e := (int32)(cEveryNsamplesEventType)
		NSamplesb317614e := (uint32)(cNSamples)
		CallbackDatab317614e := (unsafe.Pointer)(unsafe.Pointer(cCallbackData))
		retb317614e := everyNSamplesEventCallbackPtrB317614EFunc(TaskHandleb317614e, EveryNsamplesEventTypeb317614e, NSamplesb317614e, CallbackDatab317614e)
		ret, _ := (C.int32)(retb317614e), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var everyNSamplesEventCallbackPtrB317614EFunc EveryNSamplesEventCallbackPtr

func (x DoneEventCallbackPtr) PassRef() (ref *C.DAQmxDoneEventCallbackPtr, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if doneEventCallbackPtr950801AFFunc == nil {
		doneEventCallbackPtr950801AFFunc = x
	}
	return (*C.DAQmxDoneEventCallbackPtr)(C.DAQmxDoneEventCallbackPtr_950801af), nil
}

func (x DoneEventCallbackPtr) PassValue() (ref C.DAQmxDoneEventCallbackPtr, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if doneEventCallbackPtr950801AFFunc == nil {
		doneEventCallbackPtr950801AFFunc = x
	}
	return (C.DAQmxDoneEventCallbackPtr)(C.DAQmxDoneEventCallbackPtr_950801af), nil
}

func NewDoneEventCallbackPtrRef(ref unsafe.Pointer) *DoneEventCallbackPtr {
	return (*DoneEventCallbackPtr)(ref)
}

//export doneEventCallbackPtr950801AF
func doneEventCallbackPtr950801AF(cTaskHandle C.TaskHandle, cStatus C.int32, cCallbackData unsafe.Pointer) C.int32 {
	if doneEventCallbackPtr950801AFFunc != nil {
		TaskHandle950801af := (unsafe.Pointer)(cTaskHandle)
		Status950801af := (int32)(cStatus)
		CallbackData950801af := (unsafe.Pointer)(unsafe.Pointer(cCallbackData))
		ret950801af := doneEventCallbackPtr950801AFFunc(TaskHandle950801af, Status950801af, CallbackData950801af)
		ret, _ := (C.int32)(ret950801af), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var doneEventCallbackPtr950801AFFunc DoneEventCallbackPtr

func (x SignalEventCallbackPtr) PassRef() (ref *C.DAQmxSignalEventCallbackPtr, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if signalEventCallbackPtr97D1D041Func == nil {
		signalEventCallbackPtr97D1D041Func = x
	}
	return (*C.DAQmxSignalEventCallbackPtr)(C.DAQmxSignalEventCallbackPtr_97d1d041), nil
}

func (x SignalEventCallbackPtr) PassValue() (ref C.DAQmxSignalEventCallbackPtr, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if signalEventCallbackPtr97D1D041Func == nil {
		signalEventCallbackPtr97D1D041Func = x
	}
	return (C.DAQmxSignalEventCallbackPtr)(C.DAQmxSignalEventCallbackPtr_97d1d041), nil
}

func NewSignalEventCallbackPtrRef(ref unsafe.Pointer) *SignalEventCallbackPtr {
	return (*SignalEventCallbackPtr)(ref)
}

//export signalEventCallbackPtr97D1D041
func signalEventCallbackPtr97D1D041(cTaskHandle C.TaskHandle, cSignalID C.int32, cCallbackData unsafe.Pointer) C.int32 {
	if signalEventCallbackPtr97D1D041Func != nil {
		TaskHandle97d1d041 := (unsafe.Pointer)(cTaskHandle)
		SignalID97d1d041 := (int32)(cSignalID)
		CallbackData97d1d041 := (unsafe.Pointer)(unsafe.Pointer(cCallbackData))
		ret97d1d041 := signalEventCallbackPtr97D1D041Func(TaskHandle97d1d041, SignalID97d1d041, CallbackData97d1d041)
		ret, _ := (C.int32)(ret97d1d041), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var signalEventCallbackPtr97D1D041Func SignalEventCallbackPtr

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// unpackPUInt8String represents the data from Go string as *C.uInt8 and avoids copying.
func unpackPUInt8String(str string) (*C.uInt8, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.uInt8)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}
