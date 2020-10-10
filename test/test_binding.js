const robotjs = require("../index.js");

var listener = new robotjs.EventListener()
listener.attach(robotjs.KeyDown, ["ctrl"], function(){
  console.log("ctrl pressed")
})
listener.attach(robotjs.KeyUp, ["ctrl"], function(){
  console.log("ctrl released")
})
listener.attach(robotjs.KeyDown, ["alt"], function(){
  console.log("quit")
  listener.end()
})
listener.any(function(e){
  console.log(e)
})
listener.run()

console.log("Tests passed- everything looks OK!");
