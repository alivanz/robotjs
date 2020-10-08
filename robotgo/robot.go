package main

import (
	"reflect"
	"unsafe"

	hook "github.com/robotn/gohook"
)

/*
#include <stdint.h>
#include "types.h"
*/
import "C"

var (
	poll chan unsafe.Pointer
)

//export eventHook
func eventHook(when C.uint8_t, n C.int, s **C.char, cb unsafe.Pointer) {
	keys := make([]string, n)
	var arr []*C.char
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&arr))
	hdr.Data = uintptr(unsafe.Pointer(s))
	hdr.Len = int(n)
	hdr.Cap = int(n)
	for i, cstr := range arr {
		keys[i] = C.GoString(cstr)
	}
	hook.Register(uint8(when), keys, func(e hook.Event) {
		if cb == nil {
			return
		}
		poll <- cb
		// var event C.Event_t
		// event.Kind = C.uint8_t(e.Kind)
		// event.When = C.uint64_t(e.When.Unix())
		// event.Mask = C.uint16_t(e.Mask)
		// event.Keycode = C.uint16_t(e.Keycode)
		// event.Rawcode = C.uint16_t(e.Rawcode)
		// event.Keychar = C.uint8_t(e.Keychar)
		// event.Button = C.uint16_t(e.Button)
		// event.Clicks = C.uint16_t(e.Clicks)
		// event.X = C.int16_t(e.X)
		// event.Y = C.int16_t(e.Y)
		// event.Amount = C.uint16_t(e.Amount)
		// event.Rotation = C.int32_t(e.Rotation)
		// event.Direction = C.uint8_t(e.Direction)
		// log.Print("go call back")
		// C.call_cb(helper, cb, &event)
	})
}

//export pollEventCallback
func pollEventCallback() unsafe.Pointer {
	select {
	case ptr := <-poll:
		return ptr
	}
}

//export eventProcess
func eventProcess() {
	poll = make(chan unsafe.Pointer)
	go func() {
		c := hook.Start()
		<-hook.Process(c)
		close(poll)
	}()
}

//export eventEnd
func eventEnd() {
	hook.End()
}

func main() {
}
