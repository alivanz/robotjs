#include <napi.h>
#include "../robotgo/robotgo.h"

using namespace Napi;

void GoPrint(const Napi::CallbackInfo& info) {
  Print("hehehhehe");
}

Napi::String Method(const Napi::CallbackInfo& info) {
  Napi::Env env = info.Env();
  return Napi::String::New(env, "world");
}

Napi::Object Init(Napi::Env env, Napi::Object exports) {
  exports.Set(
    Napi::String::New(env, "HelloWorld"),
    Napi::Function::New(env, Method)
  );
  exports.Set(
    Napi::String::New(env, "GoPrint"),
    Napi::Function::New(env, GoPrint)
  );
  return exports;
}

NODE_API_MODULE(robotjs, Init)
