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

Napi::Boolean EventHook(const Napi::CallbackInfo& info) {
  Napi::Env env = info.Env();
  Napi::Array arr = info[1].As<Napi::Array>();
  int n = arr.Length();
  const char **list = (const char **)malloc(sizeof(char*) * n);
  for (int i=0; i<n; i++) {
    Napi::Value elem = arr[i];
    std::string s = elem.ToString();
    const char *cstr = strdup(s.c_str());
    list[i] = cstr;
  }
  eventHook(
    info[0].ToNumber().Uint32Value(),
    n,
    (char **)list,
    0
  );
  for (int i=0; i<n; i++) {
    free((void *)list[i]);
  }
  free((void *)list);
  return Napi::Boolean::New(env, true);
}

Napi::Boolean EventProcess(const Napi::CallbackInfo& info) {
  Napi::Env env = info.Env();
  eventProcess();
  return Napi::Boolean::New(env, true);
}
Napi::Boolean EventEnd(const Napi::CallbackInfo& info) {
  Napi::Env env = info.Env();
  eventEnd();
  return Napi::Boolean::New(env, true);
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
  exports.Set(
    Napi::String::New(env, "EventHook"),
    Napi::Function::New(env, EventHook)
  );
  exports.Set(
    Napi::String::New(env, "EventProcess"),
    Napi::Function::New(env, EventProcess)
  );
  exports.Set(
    Napi::String::New(env, "EventEnd"),
    Napi::Function::New(env, EventEnd)
  );
  return exports;
}

NODE_API_MODULE(robotjs, Init)
