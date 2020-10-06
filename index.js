const robotjs = require('./build/Release/robotjs.node');
module.exports = {
  "HelloWorld": robotjs.HelloWorld,
  "GoPrint": robotjs.GoPrint,
  "EventHook":    robotjs.EventHook,
  "EventProcess": robotjs.EventProcess,
  "EventEnd":     robotjs.EventEnd,

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
