package main

import (
	"reflect"
	"unsafe"

	"github.com/mattn/go-pointer"
	hook "github.com/robotn/gohook"
)

/*
#include <stdint.h>
#include <stdlib.h>
#include "types.h"

static Event_t *NewEvent() {
	return (Event_t *)malloc(sizeof(Event_t));
}
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

//export eventStartListen
func eventStartListen() unsafe.Pointer {
	c := hook.Start()
	return pointer.Save(c)
}

//export eventRead
func eventRead(p unsafe.Pointer) unsafe.Pointer {
	c := pointer.Restore(p).(chan hook.Event)
	for e := range c {
		event := C.NewEvent()
		event.Kind = C.uint8_t(e.Kind)
		event.When = C.uint64_t(e.When.Unix())
		event.Mask = C.uint16_t(e.Mask)
		event.Keycode = C.uint16_t(e.Keycode)
		event.Rawcode = C.uint16_t(e.Rawcode)
		event.Keychar = C.uint8_t(e.Keychar)
		event.Button = C.uint16_t(e.Button)
		event.Clicks = C.uint16_t(e.Clicks)
		event.X = C.int16_t(e.X)
		event.Y = C.int16_t(e.Y)
		event.Amount = C.uint16_t(e.Amount)
		event.Rotation = C.int32_t(e.Rotation)
		event.Direction = C.uint8_t(e.Direction)
		return unsafe.Pointer(event)
	}
	pointer.Unref(p)
	return nil
}

func main() {
}
