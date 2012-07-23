// +build windows

package gl42

// #include <stdlib.h>
// #define WIN32_LEAN_AND_MEAN 1
// #include <windows.h>
//
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
//
// typedef unsigned char GLboolean;
// typedef int GLint;
//
// void* goglGetProcAddress(const char* name);
//
// GLboolean (APIENTRYP ptrwglSwapIntervalEXT)(GLint interval);
//
// GLboolean gowglSwapIntervalEXT(GLint interval) {
// 	return (*ptrwglSwapIntervalEXT)(interval);
// }
//
// int init_WGL() {
//  ptrwglSwapIntervalEXT = goglGetProcAddress("wglSwapIntervalEXT");
// 	if(ptrwglSwapIntervalEXT == NULL) return 1;
// 	return 0;
// }
import "C"
import "errors"

func InitWgl() error {
	var ret C.int
	if ret = C.init_WGL(); ret != 0 {
		return errors.New("unable to initialize WGL")
	}
	return nil
}

func SwapInterval(interval Int) Boolean {
	return (Boolean)(C.gowglSwapIntervalEXT((C.GLint)(interval)))
}
