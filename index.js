const robotjs = require('./build/Release/robotjs.node');

var keycode = {
  "`": 41,
  "1": 2,
  "2": 3,
  "3": 4,
  "4": 5,
  "5": 6,
  "6": 7,
  "7": 8,
  "8": 9,
  "9": 10,
  "0": 11,
  "-": 12,
  "+": 13,
  //
  "q":  16,
  "w":  17,
  "e":  18,
  "r":  19,
  "t":  20,
  "y":  21,
  "u":  22,
  "i":  23,
  "o":  24,
  "p":  25,
  "[":  26,
  "]":  27,
  "\\": 43,
  //
  "a": 30,
  "s": 31,
  "d": 32,
  "f": 33,
  "g": 34,
  "h": 35,
  "j": 36,
  "k": 37,
  "l": 38,
  ";": 39,
  "'": 40,
  //
  "z": 44,
  "x": 45,
  "c": 46,
  "v": 47,
  "b": 48,
  "n": 49,
  "m": 50,
  ",": 51,
  ".": 52,
  "/": 53,
  //
  "f1":  59,
  "f2":  60,
  "f3":  61,
  "f4":  62,
  "f5":  63,
  "f6":  64,
  "f7":  65,
  "f8":  66,
  "f9":  67,
  "f10": 68,
  "f11": 69,
  "f12": 70,
  // more
  "esc":     1,
  "delete":  14,
  "tab":     15,
  "ctrl":    29,
  "control": 29,
  "alt":     56,
  "space":   57,
  "shift":   42,
  "rshift":  54,
  "enter":   28,
  //
  "cmd":     3675,
  "command": 3675,
  "rcmd":    3676,
  "ralt":    3640,
  //
  "up":    57416,
  "down":  57424,
  "left":  57419,
  "right": 57421,
}

class EventListener {
  constructor() {
    this.pressed = new Array(256)
    for (var i=0; i<this.pressed.length; i++) {
      this.pressed[i] = false
    }
    this.listener = []
  }
  attach(when, cmds, cb) {
    var keys = []
    for (var i=0; i<cmds.length; i++) {
      var cmd = cmds[i]
      keys.push(keycode[cmd])
    }
    this.listener.push({
      when: when,
      keys: keys,
      cb: cb,
      prev: false,
    })
  }
  any(cb) {
    this.cb = cb
  }
  all_pressed(keys) {
    for (var i=0; i<keys.length; i++) {
      var key = keys[i]
      if (!this.pressed[key]) {
        return false
      }
    }
    return true
  }
  run() {
    robotjs.EventAll((e)=>{
      if (e.kind == 3 || e.kind == 4) {
        this.pressed[e.key_code] = true
      } else if (e.kind == 5) {
        this.pressed[e.key_code] = false
      }
      for (var i=0; i<this.listener.length; i++) {
        var listener = this.listener[i]
        var current = this.all_pressed(listener.keys)
        if (listener.when == 3 && !listener.prev && current) {
          listener.cb()
        } else if (listener.when == 5 && listener.prev && !current) {
          listener.cb()
        }
        listener.prev = current
      }
    })
  }
  end() {
    robotjs.EventEnd()
  }
}

module.exports = {
  "EventHook":    robotjs.EventHook,
  "EventProcess": robotjs.EventProcess,
  "EventEnd":     robotjs.EventEnd,

  "EventAll":     robotjs.EventAll,
  "EventListener":EventListener,

  // HookEnabled honk enable status
	HookEnabled  : 1,
	HookDisabled : 2,

	KeyDown : 3,
	KeyHold : 4,
	KeyUp   : 5,

	MouseUp    : 6,
	MouseHold  : 7,
	MouseDown  : 8,
	MouseMove  : 9,
	MouseDrag  : 10,
	MouseWheel : 11,

	FakeEvent : 12,
	// Keychar could be v
	CharUndefined : 0xFFFF,
	WheelUp       : -1,
	WheelDown     : 1,
}
