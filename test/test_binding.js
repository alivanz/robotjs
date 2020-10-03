const robotjs = require("../lib/binding.js");
const assert = require("assert");

assert(robotjs.HelloWorld, "The expected function is undefined");

function testBasic() {
    const result =  robotjs.HelloWorld("hello");
    robotjs.GoPrint();
    assert.strictEqual(result, "world", "Unexpected value returned");
}

assert.doesNotThrow(testBasic, undefined, "testBasic threw an expection");

console.log("Tests passed- everything looks OK!");
