const robotjs = require("../index.js");
const assert = require("assert");

function testBasic() {
    // robotjs.EventHook(robotjs.KeyDown, ["w"], function(){
    //   console.log("w pressed")
    // });
    // robotjs.EventHook(robotjs.KeyDown, ["ctrl", "w"], function(){
    //   console.log("ctrl+w pressed")
    // });
    // robotjs.EventHook(robotjs.KeyDown, ["ctrl", "q"], function(){
    //   console.log("quit")
    //   robotjs.EventEnd()
    // });
    // robotjs.EventProcess();
    // assert.strictEqual(result, "world", "Unexpected value returned");
    robotjs.EventAll(function(e) {
      console.log(e)
      if (e.key_code == 16) {
        robotjs.EventEnd()
      }
    });
}

testBasic();

assert.doesNotThrow(testBasic, undefined, "testBasic threw an expection");

console.log("Tests passed- everything looks OK!");
