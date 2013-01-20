// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// 3DL_stereo_control: http://www.opengl.org/registry/specs/3DL/stereo_control.txt
// 
// AMD_gpu_association: http://www.opengl.org/registry/specs/AMD/gpu_association.txt
// 
// ARB_buffer_region: http://www.opengl.org/registry/specs/ARB/buffer_region.txt
// 
// ARB_create_context: http://www.opengl.org/registry/specs/ARB/create_context.txt
// 
// ARB_extensions_string: http://www.opengl.org/registry/specs/ARB/extensions_string.txt
// 
// ARB_make_current_read: http://www.opengl.org/registry/specs/ARB/make_current_read.txt
// 
// ARB_pbuffer: http://www.opengl.org/registry/specs/ARB/pbuffer.txt
// 
// ARB_pixel_format: http://www.opengl.org/registry/specs/ARB/pixel_format.txt
// 
// ARB_render_texture: http://www.opengl.org/registry/specs/ARB/render_texture.txt
// 
// EXT_display_color_table: http://www.opengl.org/registry/specs/EXT/display_color_table.txt
// 
// EXT_extensions_string: http://www.opengl.org/registry/specs/EXT/extensions_string.txt
// 
// EXT_make_current_read: http://www.opengl.org/registry/specs/EXT/make_current_read.txt
// 
// EXT_pbuffer: http://www.opengl.org/registry/specs/EXT/pbuffer.txt
// 
// EXT_pixel_format: http://www.opengl.org/registry/specs/EXT/pixel_format.txt
// 
// EXT_swap_control: http://www.opengl.org/registry/specs/EXT/swap_control.txt
// 
// I3D_digital_video_control: http://www.opengl.org/registry/specs/I3D/digital_video_control.txt
// 
// I3D_gamma: http://www.opengl.org/registry/specs/I3D/gamma.txt
// 
// I3D_genlock: http://www.opengl.org/registry/specs/I3D/genlock.txt
// 
// I3D_image_buffer: http://www.opengl.org/registry/specs/I3D/image_buffer.txt
// 
// I3D_swap_frame_lock: http://www.opengl.org/registry/specs/I3D/swap_frame_lock.txt
// 
// I3D_swap_frame_usage: http://www.opengl.org/registry/specs/I3D/swap_frame_usage.txt
// 
// NV_DX_interop: http://www.opengl.org/registry/specs/NV/DX_interop.txt
// 
// NV_copy_image: http://www.opengl.org/registry/specs/NV/copy_image.txt
// 
// NV_gpu_affinity: http://www.opengl.org/registry/specs/NV/gpu_affinity.txt
// 
// NV_present_video: http://www.opengl.org/registry/specs/NV/present_video.txt
// 
// NV_swap_group: http://www.opengl.org/registry/specs/NV/swap_group.txt
// 
// NV_vertex_array_range: http://www.opengl.org/registry/specs/NV/vertex_array_range.txt
// 
// NV_video_capture: http://www.opengl.org/registry/specs/NV/video_capture.txt
// 
// NV_video_output: http://www.opengl.org/registry/specs/NV/video_output.txt
// 
// OML_sync_control: http://www.opengl.org/registry/specs/OML/sync_control.txt
// 
// wgl: http://www.opengl.org/registry/specs/wgl/.txt
// 
package wgl

// #cgo darwin  LDFLAGS: -framework OpenGL
// #cgo linux   LDFLAGS: -lGL
// #cgo windows LDFLAGS: -lopengl32
// 
// #include <stdlib.h>
// #if defined(__APPLE__)
// #include <dlfcn.h>
// #elif defined(_WIN32)
// #define WIN32_LEAN_AND_MEAN 1
// #include <windows.h>
// #else
// #include <X11/Xlib.h>
// #include <GL/glx.h>
// #endif
// 
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
// #ifndef GLAPI
// #define GLAPI extern
// #endif
// 
// typedef unsigned int GLenum;
// typedef unsigned char GLboolean;
// typedef unsigned int GLbitfield;
// typedef signed char GLbyte;
// typedef short GLshort;
// typedef int GLint;
// typedef int GLsizei;
// typedef unsigned char GLubyte;
// typedef unsigned short GLushort;
// typedef unsigned int GLuint;
// typedef unsigned short GLhalf;
// typedef float GLfloat;
// typedef float GLclampf;
// typedef double GLdouble;
// typedef double GLclampd;
// typedef void GLvoid;
// 
// #ifndef WGL_ARB_pbuffer
// DECLARE_HANDLE(HPBUFFERARB);
// #endif
// #ifndef WGL_EXT_pbuffer
// DECLARE_HANDLE(HPBUFFEREXT);
// #endif
// #ifndef WGL_NV_present_video
// DECLARE_HANDLE(HVIDEOOUTPUTDEVICENV);
// #endif
// #ifndef WGL_NV_video_output
// DECLARE_HANDLE(HPVIDEODEV);
// #endif
// #ifndef WGL_NV_gpu_affinity
// DECLARE_HANDLE(HPGPUNV);
// DECLARE_HANDLE(HGPUNV);
// 
// typedef struct _GPU_DEVICE {
// DWORD  cb;
// CHAR   DeviceName[32];
// CHAR   DeviceString[128];
// DWORD  Flags;
// RECT   rcVirtualScreen;
// } GPU_DEVICE, *PGPU_DEVICE;
// #endif
// #ifndef WGL_NV_video_capture
// DECLARE_HANDLE(HVIDEOINPUTDEVICENV);
// #endif
// 
// #ifdef _WIN32
// static HMODULE opengl32 = NULL;
// #endif
// 
// static void* gowglGetProcAddress(const char* name) { 
// #ifdef __APPLE__
// 	return dlsym(RTLD_DEFAULT, name);
// #elif _WIN32
// 	void* pf = wglGetProcAddress((LPCSTR)name);
// 	if(pf) {
// 		return pf;
// 	}
// 	if(opengl32 == NULL) {
// 		opengl32 = LoadLibraryA("opengl32.dll");
// 	}
// 	return GetProcAddress(opengl32, (LPCSTR)name);
// #else
// 	return glXGetProcAddress((const GLubyte*)name);
// #endif
// }
// 
// //  3DL_stereo_control
// BOOL (APIENTRYP ptrwglSetStereoEmitterState3DL)(HDC hDC, UINT uState);
// //  AMD_gpu_association
// UINT (APIENTRYP ptrwglGetGPUIDsAMD)(UINT maxCount, UINT* ids);
// INT (APIENTRYP ptrwglGetGPUInfoAMD)(UINT id, int property, GLenum dataType, UINT size, void* data);
// UINT (APIENTRYP ptrwglGetContextGPUIDAMD)(HGLRC hglrc);
// HGLRC (APIENTRYP ptrwglCreateAssociatedContextAMD)(UINT id);
// HGLRC (APIENTRYP ptrwglCreateAssociatedContextAttribsAMD)(UINT id, HGLRC hShareContext, int* attribList);
// BOOL (APIENTRYP ptrwglDeleteAssociatedContextAMD)(HGLRC hglrc);
// BOOL (APIENTRYP ptrwglMakeAssociatedContextCurrentAMD)(HGLRC hglrc);
// HGLRC (APIENTRYP ptrwglGetCurrentAssociatedContextAMD)();
// VOID (APIENTRYP ptrwglBlitContextFramebufferAMD)(HGLRC dstCtx, GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
// //  ARB_buffer_region
// HANDLE (APIENTRYP ptrwglCreateBufferRegionARB)(HDC hDC, int iLayerPlane, UINT uType);
// VOID (APIENTRYP ptrwglDeleteBufferRegionARB)(HANDLE hRegion);
// BOOL (APIENTRYP ptrwglSaveBufferRegionARB)(HANDLE hRegion, int x, int y, int width, int height);
// BOOL (APIENTRYP ptrwglRestoreBufferRegionARB)(HANDLE hRegion, int x, int y, int width, int height, int xSrc, int ySrc);
// //  ARB_create_context
// HGLRC (APIENTRYP ptrwglCreateContextAttribsARB)(HDC hDC, HGLRC hShareContext, int* attribList);
// //  ARB_extensions_string
// const char * (APIENTRYP ptrwglGetExtensionsStringARB)(HDC hdc);
// //  ARB_make_current_read
// BOOL (APIENTRYP ptrwglMakeContextCurrentARB)(HDC hDrawDC, HDC hReadDC, HGLRC hglrc);
// HDC (APIENTRYP ptrwglGetCurrentReadDCARB)();
// //  ARB_pbuffer
// HPBUFFERARB (APIENTRYP ptrwglCreatePbufferARB)(HDC hDC, int iPixelFormat, int iWidth, int iHeight, int* piAttribList);
// HDC (APIENTRYP ptrwglGetPbufferDCARB)(HPBUFFERARB hPbuffer);
// int (APIENTRYP ptrwglReleasePbufferDCARB)(HPBUFFERARB hPbuffer, HDC hDC);
// BOOL (APIENTRYP ptrwglDestroyPbufferARB)(HPBUFFERARB hPbuffer);
// BOOL (APIENTRYP ptrwglQueryPbufferARB)(HPBUFFERARB hPbuffer, int iAttribute, int* piValue);
// //  ARB_pixel_format
// BOOL (APIENTRYP ptrwglGetPixelFormatAttribivARB)(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, int* piValues);
// BOOL (APIENTRYP ptrwglGetPixelFormatAttribfvARB)(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, FLOAT* pfValues);
// BOOL (APIENTRYP ptrwglChoosePixelFormatARB)(HDC hdc, int* piAttribIList, FLOAT* pfAttribFList, UINT nMaxFormats, int* piFormats, UINT* nNumFormats);
// //  ARB_render_texture
// BOOL (APIENTRYP ptrwglBindTexImageARB)(HPBUFFERARB hPbuffer, int iBuffer);
// BOOL (APIENTRYP ptrwglReleaseTexImageARB)(HPBUFFERARB hPbuffer, int iBuffer);
// BOOL (APIENTRYP ptrwglSetPbufferAttribARB)(HPBUFFERARB hPbuffer, int* piAttribList);
// //  EXT_display_color_table
// GLboolean (APIENTRYP ptrwglCreateDisplayColorTableEXT)(GLushort id);
// GLboolean (APIENTRYP ptrwglLoadDisplayColorTableEXT)(GLushort* table, GLuint length);
// GLboolean (APIENTRYP ptrwglBindDisplayColorTableEXT)(GLushort id);
// VOID (APIENTRYP ptrwglDestroyDisplayColorTableEXT)(GLushort id);
// //  EXT_extensions_string
// const char * (APIENTRYP ptrwglGetExtensionsStringEXT)();
// //  EXT_make_current_read
// BOOL (APIENTRYP ptrwglMakeContextCurrentEXT)(HDC hDrawDC, HDC hReadDC, HGLRC hglrc);
// HDC (APIENTRYP ptrwglGetCurrentReadDCEXT)();
// //  EXT_pbuffer
// HPBUFFEREXT (APIENTRYP ptrwglCreatePbufferEXT)(HDC hDC, int iPixelFormat, int iWidth, int iHeight, int* piAttribList);
// HDC (APIENTRYP ptrwglGetPbufferDCEXT)(HPBUFFEREXT hPbuffer);
// int (APIENTRYP ptrwglReleasePbufferDCEXT)(HPBUFFEREXT hPbuffer, HDC hDC);
// BOOL (APIENTRYP ptrwglDestroyPbufferEXT)(HPBUFFEREXT hPbuffer);
// BOOL (APIENTRYP ptrwglQueryPbufferEXT)(HPBUFFEREXT hPbuffer, int iAttribute, int* piValue);
// //  EXT_pixel_format
// BOOL (APIENTRYP ptrwglGetPixelFormatAttribivEXT)(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, int* piValues);
// BOOL (APIENTRYP ptrwglGetPixelFormatAttribfvEXT)(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, FLOAT* pfValues);
// BOOL (APIENTRYP ptrwglChoosePixelFormatEXT)(HDC hdc, int* piAttribIList, FLOAT* pfAttribFList, UINT nMaxFormats, int* piFormats, UINT* nNumFormats);
// //  EXT_swap_control
// BOOL (APIENTRYP ptrwglSwapIntervalEXT)(int interval);
// int (APIENTRYP ptrwglGetSwapIntervalEXT)();
// //  I3D_digital_video_control
// BOOL (APIENTRYP ptrwglGetDigitalVideoParametersI3D)(HDC hDC, int iAttribute, int* piValue);
// BOOL (APIENTRYP ptrwglSetDigitalVideoParametersI3D)(HDC hDC, int iAttribute, int* piValue);
// //  I3D_gamma
// BOOL (APIENTRYP ptrwglGetGammaTableParametersI3D)(HDC hDC, int iAttribute, int* piValue);
// BOOL (APIENTRYP ptrwglSetGammaTableParametersI3D)(HDC hDC, int iAttribute, int* piValue);
// BOOL (APIENTRYP ptrwglGetGammaTableI3D)(HDC hDC, int iEntries, USHORT* puRed, USHORT* puGreen, USHORT* puBlue);
// BOOL (APIENTRYP ptrwglSetGammaTableI3D)(HDC hDC, int iEntries, USHORT* puRed, USHORT* puGreen, USHORT* puBlue);
// //  I3D_genlock
// BOOL (APIENTRYP ptrwglEnableGenlockI3D)(HDC hDC);
// BOOL (APIENTRYP ptrwglDisableGenlockI3D)(HDC hDC);
// BOOL (APIENTRYP ptrwglIsEnabledGenlockI3D)(HDC hDC, BOOL* pFlag);
// BOOL (APIENTRYP ptrwglGenlockSourceI3D)(HDC hDC, UINT uSource);
// BOOL (APIENTRYP ptrwglGetGenlockSourceI3D)(HDC hDC, UINT* uSource);
// BOOL (APIENTRYP ptrwglGenlockSourceEdgeI3D)(HDC hDC, UINT uEdge);
// BOOL (APIENTRYP ptrwglGetGenlockSourceEdgeI3D)(HDC hDC, UINT* uEdge);
// BOOL (APIENTRYP ptrwglGenlockSampleRateI3D)(HDC hDC, UINT uRate);
// BOOL (APIENTRYP ptrwglGetGenlockSampleRateI3D)(HDC hDC, UINT* uRate);
// BOOL (APIENTRYP ptrwglGenlockSourceDelayI3D)(HDC hDC, UINT uDelay);
// BOOL (APIENTRYP ptrwglGetGenlockSourceDelayI3D)(HDC hDC, UINT* uDelay);
// BOOL (APIENTRYP ptrwglQueryGenlockMaxSourceDelayI3D)(HDC hDC, UINT* uMaxLineDelay, UINT* uMaxPixelDelay);
// //  I3D_image_buffer
// LPVOID (APIENTRYP ptrwglCreateImageBufferI3D)(HDC hDC, DWORD dwSize, UINT uFlags);
// BOOL (APIENTRYP ptrwglDestroyImageBufferI3D)(HDC hDC, LPVOID pAddress);
// BOOL (APIENTRYP ptrwglAssociateImageBufferEventsI3D)(HDC hDC, HANDLE* pEvent, LPVOID* pAddress, DWORD* pSize, UINT count);
// BOOL (APIENTRYP ptrwglReleaseImageBufferEventsI3D)(HDC hDC, LPVOID* pAddress, UINT count);
// //  I3D_swap_frame_lock
// BOOL (APIENTRYP ptrwglEnableFrameLockI3D)();
// BOOL (APIENTRYP ptrwglDisableFrameLockI3D)();
// BOOL (APIENTRYP ptrwglIsEnabledFrameLockI3D)(BOOL* pFlag);
// BOOL (APIENTRYP ptrwglQueryFrameLockMasterI3D)(BOOL* pFlag);
// //  I3D_swap_frame_usage
// BOOL (APIENTRYP ptrwglGetFrameUsageI3D)(float* pUsage);
// BOOL (APIENTRYP ptrwglBeginFrameTrackingI3D)();
// BOOL (APIENTRYP ptrwglEndFrameTrackingI3D)();
// BOOL (APIENTRYP ptrwglQueryFrameTrackingI3D)(DWORD* pFrameCount, DWORD* pMissedFrames, float* pLastMissedUsage);
// //  NV_DX_interop
// BOOL (APIENTRYP ptrwglDXSetResourceShareHandleNV)(void* dxObject, HANDLE shareHandle);
// HANDLE (APIENTRYP ptrwglDXOpenDeviceNV)(void* dxDevice);
// BOOL (APIENTRYP ptrwglDXCloseDeviceNV)(HANDLE hDevice);
// HANDLE (APIENTRYP ptrwglDXRegisterObjectNV)(HANDLE hDevice, void* dxObject, GLuint name, GLenum type, GLenum access);
// BOOL (APIENTRYP ptrwglDXUnregisterObjectNV)(HANDLE hDevice, HANDLE hObject);
// BOOL (APIENTRYP ptrwglDXObjectAccessNV)(HANDLE hObject, GLenum access);
// BOOL (APIENTRYP ptrwglDXLockObjectsNV)(HANDLE hDevice, GLint count, HANDLE* hObjects);
// BOOL (APIENTRYP ptrwglDXUnlockObjectsNV)(HANDLE hDevice, GLint count, HANDLE* hObjects);
// //  NV_copy_image
// BOOL (APIENTRYP ptrwglCopyImageSubDataNV)(HGLRC hSrcRC, GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, HGLRC hDstRC, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei width, GLsizei height, GLsizei depth);
// //  NV_gpu_affinity
// BOOL (APIENTRYP ptrwglEnumGpusNV)(UINT iGpuIndex, HGPUNV* phGpu);
// BOOL (APIENTRYP ptrwglEnumGpuDevicesNV)(HGPUNV hGpu, UINT iDeviceIndex, PGPU_DEVICE lpGpuDevice);
// HDC (APIENTRYP ptrwglCreateAffinityDCNV)(HGPUNV* phGpuList);
// BOOL (APIENTRYP ptrwglEnumGpusFromAffinityDCNV)(HDC hAffinityDC, UINT iGpuIndex, HGPUNV* hGpu);
// BOOL (APIENTRYP ptrwglDeleteDCNV)(HDC hdc);
// //  NV_present_video
// int (APIENTRYP ptrwglEnumerateVideoDevicesNV)(HDC hDC, HVIDEOOUTPUTDEVICENV* phDeviceList);
// BOOL (APIENTRYP ptrwglBindVideoDeviceNV)(HDC hDC, unsigned int uVideoSlot, HVIDEOOUTPUTDEVICENV hVideoDevice, int* piAttribList);
// BOOL (APIENTRYP ptrwglQueryCurrentContextNV)(int iAttribute, int* piValue);
// //  NV_swap_group
// BOOL (APIENTRYP ptrwglJoinSwapGroupNV)(HDC hDC, GLuint group);
// BOOL (APIENTRYP ptrwglBindSwapBarrierNV)(GLuint group, GLuint barrier);
// BOOL (APIENTRYP ptrwglQuerySwapGroupNV)(HDC hDC, GLuint* group, GLuint* barrier);
// BOOL (APIENTRYP ptrwglQueryMaxSwapGroupsNV)(HDC hDC, GLuint* maxGroups, GLuint* maxBarriers);
// BOOL (APIENTRYP ptrwglQueryFrameCountNV)(HDC hDC, GLuint* count);
// BOOL (APIENTRYP ptrwglResetFrameCountNV)(HDC hDC);
// //  NV_vertex_array_range
// void* (APIENTRYP ptrwglAllocateMemoryNV)(GLsizei size, GLfloat readfreq, GLfloat writefreq, GLfloat priority);
// void (APIENTRYP ptrwglFreeMemoryNV)(void* pointer);
// //  NV_video_capture
// BOOL (APIENTRYP ptrwglBindVideoCaptureDeviceNV)(UINT uVideoSlot, HVIDEOINPUTDEVICENV hDevice);
// UINT (APIENTRYP ptrwglEnumerateVideoCaptureDevicesNV)(HDC hDc, HVIDEOINPUTDEVICENV* phDeviceList);
// BOOL (APIENTRYP ptrwglLockVideoCaptureDeviceNV)(HDC hDc, HVIDEOINPUTDEVICENV hDevice);
// BOOL (APIENTRYP ptrwglQueryVideoCaptureDeviceNV)(HDC hDc, HVIDEOINPUTDEVICENV hDevice, int iAttribute, int* piValue);
// BOOL (APIENTRYP ptrwglReleaseVideoCaptureDeviceNV)(HDC hDc, HVIDEOINPUTDEVICENV hDevice);
// //  NV_video_output
// BOOL (APIENTRYP ptrwglGetVideoDeviceNV)(HDC hDC, int numDevices, HPVIDEODEV* hVideoDevice);
// BOOL (APIENTRYP ptrwglReleaseVideoDeviceNV)(HPVIDEODEV hVideoDevice);
// BOOL (APIENTRYP ptrwglBindVideoImageNV)(HPVIDEODEV hVideoDevice, HPBUFFERARB hPbuffer, int iVideoBuffer);
// BOOL (APIENTRYP ptrwglReleaseVideoImageNV)(HPBUFFERARB hPbuffer, int iVideoBuffer);
// BOOL (APIENTRYP ptrwglSendPbufferToVideoNV)(HPBUFFERARB hPbuffer, int iBufferType, unsigned long* pulCounterPbuffer, BOOL bBlock);
// BOOL (APIENTRYP ptrwglGetVideoInfoNV)(HPVIDEODEV hpVideoDevice, unsigned long* pulCounterOutputPbuffer, unsigned long* pulCounterOutputVideo);
// //  OML_sync_control
// BOOL (APIENTRYP ptrwglGetSyncValuesOML)(HDC hdc, INT64* ust, INT64* msc, INT64* sbc);
// BOOL (APIENTRYP ptrwglGetMscRateOML)(HDC hdc, INT32* numerator, INT32* denominator);
// INT64 (APIENTRYP ptrwglSwapBuffersMscOML)(HDC hdc, INT64 target_msc, INT64 divisor, INT64 remainder);
// INT64 (APIENTRYP ptrwglSwapLayerBuffersMscOML)(HDC hdc, int fuPlanes, INT64 target_msc, INT64 divisor, INT64 remainder);
// BOOL (APIENTRYP ptrwglWaitForMscOML)(HDC hdc, INT64 target_msc, INT64 divisor, INT64 remainder, INT64* ust, INT64* msc, INT64* sbc);
// BOOL (APIENTRYP ptrwglWaitForSbcOML)(HDC hdc, INT64 target_sbc, INT64* ust, INT64* msc, INT64* sbc);
// //  wgl
// HGLRC (APIENTRYP ptrwglCreateContext)(HDC hDc);
// BOOL (APIENTRYP ptrwglDeleteContext)(HGLRC oldContext);
// HGLRC (APIENTRYP ptrwglGetCurrentContext)();
// BOOL (APIENTRYP ptrwglMakeCurrent)(HDC hDc, HGLRC newContext);
// BOOL (APIENTRYP ptrwglCopyContext)(HGLRC hglrcSrc, HGLRC hglrcDst, UINT mask);
// int (APIENTRYP ptrwglChoosePixelFormat)(HDC hDc, PIXELFORMATDESCRIPTOR* pPfd);
// HDC (APIENTRYP ptrwglGetCurrentDC)();
// PROC (APIENTRYP ptrwglGetDefaultProcAddress)(LPCSTR lpszProc);
// PROC (APIENTRYP ptrwglGetProcAddress)(LPCSTR lpszProc);
// int (APIENTRYP ptrwglGetPixelFormat)(HDC hdc);
// BOOL (APIENTRYP ptrwglSetPixelFormat)(HDC hdc, int ipfd, PIXELFORMATDESCRIPTOR* ppfd);
// BOOL (APIENTRYP ptrwglSwapBuffers)(HDC hdc);
// BOOL (APIENTRYP ptrwglShareLists)(HGLRC hrcSrvShare, HGLRC hrcSrvSource);
// HGLRC (APIENTRYP ptrwglCreateLayerContext)(HDC hDc, int level);
// BOOL (APIENTRYP ptrwglDescribeLayerPlane)(HDC hDc, int pixelFormat, int layerPlane, UINT nBytes, LAYERPLANEDESCRIPTOR* plpd);
// int (APIENTRYP ptrwglSetLayerPaletteEntries)(HDC hdc, int iLayerPlane, int iStart, int cEntries, COLORREF* pcr);
// int (APIENTRYP ptrwglGetLayerPaletteEntries)(HDC hdc, int iLayerPlane, int iStart, int cEntries, COLORREF* pcr);
// BOOL (APIENTRYP ptrwglRealizeLayerPalette)(HDC hdc, int iLayerPlane, BOOL bRealize);
// BOOL (APIENTRYP ptrwglSwapLayerBuffers)(HDC hdc, UINT fuFlags);
// BOOL (APIENTRYP ptrwglUseFontBitmapsA)(HDC hDC, DWORD first, DWORD count, DWORD listBase);
// BOOL (APIENTRYP ptrwglUseFontBitmapsW)(HDC hDC, DWORD first, DWORD count, DWORD listBase);
// 
// //  3DL_stereo_control
// BOOL gowglSetStereoEmitterState3DL(HDC hDC, UINT uState) {
// 	return (*ptrwglSetStereoEmitterState3DL)(hDC, uState);
// }
// //  AMD_gpu_association
// UINT gowglGetGPUIDsAMD(UINT maxCount, UINT* ids) {
// 	return (*ptrwglGetGPUIDsAMD)(maxCount, ids);
// }
// INT gowglGetGPUInfoAMD(UINT id, int property, GLenum dataType, UINT size, void* data) {
// 	return (*ptrwglGetGPUInfoAMD)(id, property, dataType, size, data);
// }
// UINT gowglGetContextGPUIDAMD(HGLRC hglrc) {
// 	return (*ptrwglGetContextGPUIDAMD)(hglrc);
// }
// HGLRC gowglCreateAssociatedContextAMD(UINT id) {
// 	return (*ptrwglCreateAssociatedContextAMD)(id);
// }
// HGLRC gowglCreateAssociatedContextAttribsAMD(UINT id, HGLRC hShareContext, int* attribList) {
// 	return (*ptrwglCreateAssociatedContextAttribsAMD)(id, hShareContext, attribList);
// }
// BOOL gowglDeleteAssociatedContextAMD(HGLRC hglrc) {
// 	return (*ptrwglDeleteAssociatedContextAMD)(hglrc);
// }
// BOOL gowglMakeAssociatedContextCurrentAMD(HGLRC hglrc) {
// 	return (*ptrwglMakeAssociatedContextCurrentAMD)(hglrc);
// }
// HGLRC gowglGetCurrentAssociatedContextAMD() {
// 	return (*ptrwglGetCurrentAssociatedContextAMD)();
// }
// VOID gowglBlitContextFramebufferAMD(HGLRC dstCtx, GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter) {
// 	(*ptrwglBlitContextFramebufferAMD)(dstCtx, srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// //  ARB_buffer_region
// HANDLE gowglCreateBufferRegionARB(HDC hDC, int iLayerPlane, UINT uType) {
// 	return (*ptrwglCreateBufferRegionARB)(hDC, iLayerPlane, uType);
// }
// VOID gowglDeleteBufferRegionARB(HANDLE hRegion) {
// 	(*ptrwglDeleteBufferRegionARB)(hRegion);
// }
// BOOL gowglSaveBufferRegionARB(HANDLE hRegion, int x, int y, int width, int height) {
// 	return (*ptrwglSaveBufferRegionARB)(hRegion, x, y, width, height);
// }
// BOOL gowglRestoreBufferRegionARB(HANDLE hRegion, int x, int y, int width, int height, int xSrc, int ySrc) {
// 	return (*ptrwglRestoreBufferRegionARB)(hRegion, x, y, width, height, xSrc, ySrc);
// }
// //  ARB_create_context
// HGLRC gowglCreateContextAttribsARB(HDC hDC, HGLRC hShareContext, int* attribList) {
// 	return (*ptrwglCreateContextAttribsARB)(hDC, hShareContext, attribList);
// }
// //  ARB_extensions_string
// const char * gowglGetExtensionsStringARB(HDC hdc) {
// 	return (*ptrwglGetExtensionsStringARB)(hdc);
// }
// //  ARB_make_current_read
// BOOL gowglMakeContextCurrentARB(HDC hDrawDC, HDC hReadDC, HGLRC hglrc) {
// 	return (*ptrwglMakeContextCurrentARB)(hDrawDC, hReadDC, hglrc);
// }
// HDC gowglGetCurrentReadDCARB() {
// 	return (*ptrwglGetCurrentReadDCARB)();
// }
// //  ARB_pbuffer
// HPBUFFERARB gowglCreatePbufferARB(HDC hDC, int iPixelFormat, int iWidth, int iHeight, int* piAttribList) {
// 	return (*ptrwglCreatePbufferARB)(hDC, iPixelFormat, iWidth, iHeight, piAttribList);
// }
// HDC gowglGetPbufferDCARB(HPBUFFERARB hPbuffer) {
// 	return (*ptrwglGetPbufferDCARB)(hPbuffer);
// }
// int gowglReleasePbufferDCARB(HPBUFFERARB hPbuffer, HDC hDC) {
// 	return (*ptrwglReleasePbufferDCARB)(hPbuffer, hDC);
// }
// BOOL gowglDestroyPbufferARB(HPBUFFERARB hPbuffer) {
// 	return (*ptrwglDestroyPbufferARB)(hPbuffer);
// }
// BOOL gowglQueryPbufferARB(HPBUFFERARB hPbuffer, int iAttribute, int* piValue) {
// 	return (*ptrwglQueryPbufferARB)(hPbuffer, iAttribute, piValue);
// }
// //  ARB_pixel_format
// BOOL gowglGetPixelFormatAttribivARB(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, int* piValues) {
// 	return (*ptrwglGetPixelFormatAttribivARB)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, piValues);
// }
// BOOL gowglGetPixelFormatAttribfvARB(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, FLOAT* pfValues) {
// 	return (*ptrwglGetPixelFormatAttribfvARB)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, pfValues);
// }
// BOOL gowglChoosePixelFormatARB(HDC hdc, int* piAttribIList, FLOAT* pfAttribFList, UINT nMaxFormats, int* piFormats, UINT* nNumFormats) {
// 	return (*ptrwglChoosePixelFormatARB)(hdc, piAttribIList, pfAttribFList, nMaxFormats, piFormats, nNumFormats);
// }
// //  ARB_render_texture
// BOOL gowglBindTexImageARB(HPBUFFERARB hPbuffer, int iBuffer) {
// 	return (*ptrwglBindTexImageARB)(hPbuffer, iBuffer);
// }
// BOOL gowglReleaseTexImageARB(HPBUFFERARB hPbuffer, int iBuffer) {
// 	return (*ptrwglReleaseTexImageARB)(hPbuffer, iBuffer);
// }
// BOOL gowglSetPbufferAttribARB(HPBUFFERARB hPbuffer, int* piAttribList) {
// 	return (*ptrwglSetPbufferAttribARB)(hPbuffer, piAttribList);
// }
// //  EXT_display_color_table
// GLboolean gowglCreateDisplayColorTableEXT(GLushort id) {
// 	return (*ptrwglCreateDisplayColorTableEXT)(id);
// }
// GLboolean gowglLoadDisplayColorTableEXT(GLushort* table, GLuint length) {
// 	return (*ptrwglLoadDisplayColorTableEXT)(table, length);
// }
// GLboolean gowglBindDisplayColorTableEXT(GLushort id) {
// 	return (*ptrwglBindDisplayColorTableEXT)(id);
// }
// VOID gowglDestroyDisplayColorTableEXT(GLushort id) {
// 	(*ptrwglDestroyDisplayColorTableEXT)(id);
// }
// //  EXT_extensions_string
// const char * gowglGetExtensionsStringEXT() {
// 	return (*ptrwglGetExtensionsStringEXT)();
// }
// //  EXT_make_current_read
// BOOL gowglMakeContextCurrentEXT(HDC hDrawDC, HDC hReadDC, HGLRC hglrc) {
// 	return (*ptrwglMakeContextCurrentEXT)(hDrawDC, hReadDC, hglrc);
// }
// HDC gowglGetCurrentReadDCEXT() {
// 	return (*ptrwglGetCurrentReadDCEXT)();
// }
// //  EXT_pbuffer
// HPBUFFEREXT gowglCreatePbufferEXT(HDC hDC, int iPixelFormat, int iWidth, int iHeight, int* piAttribList) {
// 	return (*ptrwglCreatePbufferEXT)(hDC, iPixelFormat, iWidth, iHeight, piAttribList);
// }
// HDC gowglGetPbufferDCEXT(HPBUFFEREXT hPbuffer) {
// 	return (*ptrwglGetPbufferDCEXT)(hPbuffer);
// }
// int gowglReleasePbufferDCEXT(HPBUFFEREXT hPbuffer, HDC hDC) {
// 	return (*ptrwglReleasePbufferDCEXT)(hPbuffer, hDC);
// }
// BOOL gowglDestroyPbufferEXT(HPBUFFEREXT hPbuffer) {
// 	return (*ptrwglDestroyPbufferEXT)(hPbuffer);
// }
// BOOL gowglQueryPbufferEXT(HPBUFFEREXT hPbuffer, int iAttribute, int* piValue) {
// 	return (*ptrwglQueryPbufferEXT)(hPbuffer, iAttribute, piValue);
// }
// //  EXT_pixel_format
// BOOL gowglGetPixelFormatAttribivEXT(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, int* piValues) {
// 	return (*ptrwglGetPixelFormatAttribivEXT)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, piValues);
// }
// BOOL gowglGetPixelFormatAttribfvEXT(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, FLOAT* pfValues) {
// 	return (*ptrwglGetPixelFormatAttribfvEXT)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, pfValues);
// }
// BOOL gowglChoosePixelFormatEXT(HDC hdc, int* piAttribIList, FLOAT* pfAttribFList, UINT nMaxFormats, int* piFormats, UINT* nNumFormats) {
// 	return (*ptrwglChoosePixelFormatEXT)(hdc, piAttribIList, pfAttribFList, nMaxFormats, piFormats, nNumFormats);
// }
// //  EXT_swap_control
// BOOL gowglSwapIntervalEXT(int interval) {
// 	return (*ptrwglSwapIntervalEXT)(interval);
// }
// int gowglGetSwapIntervalEXT() {
// 	return (*ptrwglGetSwapIntervalEXT)();
// }
// //  I3D_digital_video_control
// BOOL gowglGetDigitalVideoParametersI3D(HDC hDC, int iAttribute, int* piValue) {
// 	return (*ptrwglGetDigitalVideoParametersI3D)(hDC, iAttribute, piValue);
// }
// BOOL gowglSetDigitalVideoParametersI3D(HDC hDC, int iAttribute, int* piValue) {
// 	return (*ptrwglSetDigitalVideoParametersI3D)(hDC, iAttribute, piValue);
// }
// //  I3D_gamma
// BOOL gowglGetGammaTableParametersI3D(HDC hDC, int iAttribute, int* piValue) {
// 	return (*ptrwglGetGammaTableParametersI3D)(hDC, iAttribute, piValue);
// }
// BOOL gowglSetGammaTableParametersI3D(HDC hDC, int iAttribute, int* piValue) {
// 	return (*ptrwglSetGammaTableParametersI3D)(hDC, iAttribute, piValue);
// }
// BOOL gowglGetGammaTableI3D(HDC hDC, int iEntries, USHORT* puRed, USHORT* puGreen, USHORT* puBlue) {
// 	return (*ptrwglGetGammaTableI3D)(hDC, iEntries, puRed, puGreen, puBlue);
// }
// BOOL gowglSetGammaTableI3D(HDC hDC, int iEntries, USHORT* puRed, USHORT* puGreen, USHORT* puBlue) {
// 	return (*ptrwglSetGammaTableI3D)(hDC, iEntries, puRed, puGreen, puBlue);
// }
// //  I3D_genlock
// BOOL gowglEnableGenlockI3D(HDC hDC) {
// 	return (*ptrwglEnableGenlockI3D)(hDC);
// }
// BOOL gowglDisableGenlockI3D(HDC hDC) {
// 	return (*ptrwglDisableGenlockI3D)(hDC);
// }
// BOOL gowglIsEnabledGenlockI3D(HDC hDC, BOOL* pFlag) {
// 	return (*ptrwglIsEnabledGenlockI3D)(hDC, pFlag);
// }
// BOOL gowglGenlockSourceI3D(HDC hDC, UINT uSource) {
// 	return (*ptrwglGenlockSourceI3D)(hDC, uSource);
// }
// BOOL gowglGetGenlockSourceI3D(HDC hDC, UINT* uSource) {
// 	return (*ptrwglGetGenlockSourceI3D)(hDC, uSource);
// }
// BOOL gowglGenlockSourceEdgeI3D(HDC hDC, UINT uEdge) {
// 	return (*ptrwglGenlockSourceEdgeI3D)(hDC, uEdge);
// }
// BOOL gowglGetGenlockSourceEdgeI3D(HDC hDC, UINT* uEdge) {
// 	return (*ptrwglGetGenlockSourceEdgeI3D)(hDC, uEdge);
// }
// BOOL gowglGenlockSampleRateI3D(HDC hDC, UINT uRate) {
// 	return (*ptrwglGenlockSampleRateI3D)(hDC, uRate);
// }
// BOOL gowglGetGenlockSampleRateI3D(HDC hDC, UINT* uRate) {
// 	return (*ptrwglGetGenlockSampleRateI3D)(hDC, uRate);
// }
// BOOL gowglGenlockSourceDelayI3D(HDC hDC, UINT uDelay) {
// 	return (*ptrwglGenlockSourceDelayI3D)(hDC, uDelay);
// }
// BOOL gowglGetGenlockSourceDelayI3D(HDC hDC, UINT* uDelay) {
// 	return (*ptrwglGetGenlockSourceDelayI3D)(hDC, uDelay);
// }
// BOOL gowglQueryGenlockMaxSourceDelayI3D(HDC hDC, UINT* uMaxLineDelay, UINT* uMaxPixelDelay) {
// 	return (*ptrwglQueryGenlockMaxSourceDelayI3D)(hDC, uMaxLineDelay, uMaxPixelDelay);
// }
// //  I3D_image_buffer
// LPVOID gowglCreateImageBufferI3D(HDC hDC, DWORD dwSize, UINT uFlags) {
// 	return (*ptrwglCreateImageBufferI3D)(hDC, dwSize, uFlags);
// }
// BOOL gowglDestroyImageBufferI3D(HDC hDC, LPVOID pAddress) {
// 	return (*ptrwglDestroyImageBufferI3D)(hDC, pAddress);
// }
// BOOL gowglAssociateImageBufferEventsI3D(HDC hDC, HANDLE* pEvent, LPVOID* pAddress, DWORD* pSize, UINT count) {
// 	return (*ptrwglAssociateImageBufferEventsI3D)(hDC, pEvent, pAddress, pSize, count);
// }
// BOOL gowglReleaseImageBufferEventsI3D(HDC hDC, LPVOID* pAddress, UINT count) {
// 	return (*ptrwglReleaseImageBufferEventsI3D)(hDC, pAddress, count);
// }
// //  I3D_swap_frame_lock
// BOOL gowglEnableFrameLockI3D() {
// 	return (*ptrwglEnableFrameLockI3D)();
// }
// BOOL gowglDisableFrameLockI3D() {
// 	return (*ptrwglDisableFrameLockI3D)();
// }
// BOOL gowglIsEnabledFrameLockI3D(BOOL* pFlag) {
// 	return (*ptrwglIsEnabledFrameLockI3D)(pFlag);
// }
// BOOL gowglQueryFrameLockMasterI3D(BOOL* pFlag) {
// 	return (*ptrwglQueryFrameLockMasterI3D)(pFlag);
// }
// //  I3D_swap_frame_usage
// BOOL gowglGetFrameUsageI3D(float* pUsage) {
// 	return (*ptrwglGetFrameUsageI3D)(pUsage);
// }
// BOOL gowglBeginFrameTrackingI3D() {
// 	return (*ptrwglBeginFrameTrackingI3D)();
// }
// BOOL gowglEndFrameTrackingI3D() {
// 	return (*ptrwglEndFrameTrackingI3D)();
// }
// BOOL gowglQueryFrameTrackingI3D(DWORD* pFrameCount, DWORD* pMissedFrames, float* pLastMissedUsage) {
// 	return (*ptrwglQueryFrameTrackingI3D)(pFrameCount, pMissedFrames, pLastMissedUsage);
// }
// //  NV_DX_interop
// BOOL gowglDXSetResourceShareHandleNV(void* dxObject, HANDLE shareHandle) {
// 	return (*ptrwglDXSetResourceShareHandleNV)(dxObject, shareHandle);
// }
// HANDLE gowglDXOpenDeviceNV(void* dxDevice) {
// 	return (*ptrwglDXOpenDeviceNV)(dxDevice);
// }
// BOOL gowglDXCloseDeviceNV(HANDLE hDevice) {
// 	return (*ptrwglDXCloseDeviceNV)(hDevice);
// }
// HANDLE gowglDXRegisterObjectNV(HANDLE hDevice, void* dxObject, GLuint name, GLenum type_, GLenum access) {
// 	return (*ptrwglDXRegisterObjectNV)(hDevice, dxObject, name, type_, access);
// }
// BOOL gowglDXUnregisterObjectNV(HANDLE hDevice, HANDLE hObject) {
// 	return (*ptrwglDXUnregisterObjectNV)(hDevice, hObject);
// }
// BOOL gowglDXObjectAccessNV(HANDLE hObject, GLenum access) {
// 	return (*ptrwglDXObjectAccessNV)(hObject, access);
// }
// BOOL gowglDXLockObjectsNV(HANDLE hDevice, GLint count, HANDLE* hObjects) {
// 	return (*ptrwglDXLockObjectsNV)(hDevice, count, hObjects);
// }
// BOOL gowglDXUnlockObjectsNV(HANDLE hDevice, GLint count, HANDLE* hObjects) {
// 	return (*ptrwglDXUnlockObjectsNV)(hDevice, count, hObjects);
// }
// //  NV_copy_image
// BOOL gowglCopyImageSubDataNV(HGLRC hSrcRC, GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, HGLRC hDstRC, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei width, GLsizei height, GLsizei depth) {
// 	return (*ptrwglCopyImageSubDataNV)(hSrcRC, srcName, srcTarget, srcLevel, srcX, srcY, srcZ, hDstRC, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, width, height, depth);
// }
// //  NV_gpu_affinity
// BOOL gowglEnumGpusNV(UINT iGpuIndex, HGPUNV* phGpu) {
// 	return (*ptrwglEnumGpusNV)(iGpuIndex, phGpu);
// }
// BOOL gowglEnumGpuDevicesNV(HGPUNV hGpu, UINT iDeviceIndex, PGPU_DEVICE lpGpuDevice) {
// 	return (*ptrwglEnumGpuDevicesNV)(hGpu, iDeviceIndex, lpGpuDevice);
// }
// HDC gowglCreateAffinityDCNV(HGPUNV* phGpuList) {
// 	return (*ptrwglCreateAffinityDCNV)(phGpuList);
// }
// BOOL gowglEnumGpusFromAffinityDCNV(HDC hAffinityDC, UINT iGpuIndex, HGPUNV* hGpu) {
// 	return (*ptrwglEnumGpusFromAffinityDCNV)(hAffinityDC, iGpuIndex, hGpu);
// }
// BOOL gowglDeleteDCNV(HDC hdc) {
// 	return (*ptrwglDeleteDCNV)(hdc);
// }
// //  NV_present_video
// int gowglEnumerateVideoDevicesNV(HDC hDC, HVIDEOOUTPUTDEVICENV* phDeviceList) {
// 	return (*ptrwglEnumerateVideoDevicesNV)(hDC, phDeviceList);
// }
// BOOL gowglBindVideoDeviceNV(HDC hDC, unsigned int uVideoSlot, HVIDEOOUTPUTDEVICENV hVideoDevice, int* piAttribList) {
// 	return (*ptrwglBindVideoDeviceNV)(hDC, uVideoSlot, hVideoDevice, piAttribList);
// }
// BOOL gowglQueryCurrentContextNV(int iAttribute, int* piValue) {
// 	return (*ptrwglQueryCurrentContextNV)(iAttribute, piValue);
// }
// //  NV_swap_group
// BOOL gowglJoinSwapGroupNV(HDC hDC, GLuint group) {
// 	return (*ptrwglJoinSwapGroupNV)(hDC, group);
// }
// BOOL gowglBindSwapBarrierNV(GLuint group, GLuint barrier) {
// 	return (*ptrwglBindSwapBarrierNV)(group, barrier);
// }
// BOOL gowglQuerySwapGroupNV(HDC hDC, GLuint* group, GLuint* barrier) {
// 	return (*ptrwglQuerySwapGroupNV)(hDC, group, barrier);
// }
// BOOL gowglQueryMaxSwapGroupsNV(HDC hDC, GLuint* maxGroups, GLuint* maxBarriers) {
// 	return (*ptrwglQueryMaxSwapGroupsNV)(hDC, maxGroups, maxBarriers);
// }
// BOOL gowglQueryFrameCountNV(HDC hDC, GLuint* count) {
// 	return (*ptrwglQueryFrameCountNV)(hDC, count);
// }
// BOOL gowglResetFrameCountNV(HDC hDC) {
// 	return (*ptrwglResetFrameCountNV)(hDC);
// }
// //  NV_vertex_array_range
// void* gowglAllocateMemoryNV(GLsizei size, GLfloat readfreq, GLfloat writefreq, GLfloat priority) {
// 	return (*ptrwglAllocateMemoryNV)(size, readfreq, writefreq, priority);
// }
// void gowglFreeMemoryNV(void* pointer) {
// 	(*ptrwglFreeMemoryNV)(pointer);
// }
// //  NV_video_capture
// BOOL gowglBindVideoCaptureDeviceNV(UINT uVideoSlot, HVIDEOINPUTDEVICENV hDevice) {
// 	return (*ptrwglBindVideoCaptureDeviceNV)(uVideoSlot, hDevice);
// }
// UINT gowglEnumerateVideoCaptureDevicesNV(HDC hDc, HVIDEOINPUTDEVICENV* phDeviceList) {
// 	return (*ptrwglEnumerateVideoCaptureDevicesNV)(hDc, phDeviceList);
// }
// BOOL gowglLockVideoCaptureDeviceNV(HDC hDc, HVIDEOINPUTDEVICENV hDevice) {
// 	return (*ptrwglLockVideoCaptureDeviceNV)(hDc, hDevice);
// }
// BOOL gowglQueryVideoCaptureDeviceNV(HDC hDc, HVIDEOINPUTDEVICENV hDevice, int iAttribute, int* piValue) {
// 	return (*ptrwglQueryVideoCaptureDeviceNV)(hDc, hDevice, iAttribute, piValue);
// }
// BOOL gowglReleaseVideoCaptureDeviceNV(HDC hDc, HVIDEOINPUTDEVICENV hDevice) {
// 	return (*ptrwglReleaseVideoCaptureDeviceNV)(hDc, hDevice);
// }
// //  NV_video_output
// BOOL gowglGetVideoDeviceNV(HDC hDC, int numDevices, HPVIDEODEV* hVideoDevice) {
// 	return (*ptrwglGetVideoDeviceNV)(hDC, numDevices, hVideoDevice);
// }
// BOOL gowglReleaseVideoDeviceNV(HPVIDEODEV hVideoDevice) {
// 	return (*ptrwglReleaseVideoDeviceNV)(hVideoDevice);
// }
// BOOL gowglBindVideoImageNV(HPVIDEODEV hVideoDevice, HPBUFFERARB hPbuffer, int iVideoBuffer) {
// 	return (*ptrwglBindVideoImageNV)(hVideoDevice, hPbuffer, iVideoBuffer);
// }
// BOOL gowglReleaseVideoImageNV(HPBUFFERARB hPbuffer, int iVideoBuffer) {
// 	return (*ptrwglReleaseVideoImageNV)(hPbuffer, iVideoBuffer);
// }
// BOOL gowglSendPbufferToVideoNV(HPBUFFERARB hPbuffer, int iBufferType, unsigned long* pulCounterPbuffer, BOOL bBlock) {
// 	return (*ptrwglSendPbufferToVideoNV)(hPbuffer, iBufferType, pulCounterPbuffer, bBlock);
// }
// BOOL gowglGetVideoInfoNV(HPVIDEODEV hpVideoDevice, unsigned long* pulCounterOutputPbuffer, unsigned long* pulCounterOutputVideo) {
// 	return (*ptrwglGetVideoInfoNV)(hpVideoDevice, pulCounterOutputPbuffer, pulCounterOutputVideo);
// }
// //  OML_sync_control
// BOOL gowglGetSyncValuesOML(HDC hdc, INT64* ust, INT64* msc, INT64* sbc) {
// 	return (*ptrwglGetSyncValuesOML)(hdc, ust, msc, sbc);
// }
// BOOL gowglGetMscRateOML(HDC hdc, INT32* numerator, INT32* denominator) {
// 	return (*ptrwglGetMscRateOML)(hdc, numerator, denominator);
// }
// INT64 gowglSwapBuffersMscOML(HDC hdc, INT64 target_msc, INT64 divisor, INT64 remainder) {
// 	return (*ptrwglSwapBuffersMscOML)(hdc, target_msc, divisor, remainder);
// }
// INT64 gowglSwapLayerBuffersMscOML(HDC hdc, int fuPlanes, INT64 target_msc, INT64 divisor, INT64 remainder) {
// 	return (*ptrwglSwapLayerBuffersMscOML)(hdc, fuPlanes, target_msc, divisor, remainder);
// }
// BOOL gowglWaitForMscOML(HDC hdc, INT64 target_msc, INT64 divisor, INT64 remainder, INT64* ust, INT64* msc, INT64* sbc) {
// 	return (*ptrwglWaitForMscOML)(hdc, target_msc, divisor, remainder, ust, msc, sbc);
// }
// BOOL gowglWaitForSbcOML(HDC hdc, INT64 target_sbc, INT64* ust, INT64* msc, INT64* sbc) {
// 	return (*ptrwglWaitForSbcOML)(hdc, target_sbc, ust, msc, sbc);
// }
// //  wgl
// HGLRC gowglCreateContext(HDC hDc) {
// 	return (*ptrwglCreateContext)(hDc);
// }
// BOOL gowglDeleteContext(HGLRC oldContext) {
// 	return (*ptrwglDeleteContext)(oldContext);
// }
// HGLRC gowglGetCurrentContext() {
// 	return (*ptrwglGetCurrentContext)();
// }
// BOOL gowglMakeCurrent(HDC hDc, HGLRC newContext) {
// 	return (*ptrwglMakeCurrent)(hDc, newContext);
// }
// BOOL gowglCopyContext(HGLRC hglrcSrc, HGLRC hglrcDst, UINT mask) {
// 	return (*ptrwglCopyContext)(hglrcSrc, hglrcDst, mask);
// }
// int gowglChoosePixelFormat(HDC hDc, PIXELFORMATDESCRIPTOR* pPfd) {
// 	return (*ptrwglChoosePixelFormat)(hDc, pPfd);
// }
// HDC gowglGetCurrentDC() {
// 	return (*ptrwglGetCurrentDC)();
// }
// PROC gowglGetDefaultProcAddress(LPCSTR lpszProc) {
// 	return (*ptrwglGetDefaultProcAddress)(lpszProc);
// }
// PROC gowglGetProcAddress(LPCSTR lpszProc) {
// 	return (*ptrwglGetProcAddress)(lpszProc);
// }
// int gowglGetPixelFormat(HDC hdc) {
// 	return (*ptrwglGetPixelFormat)(hdc);
// }
// BOOL gowglSetPixelFormat(HDC hdc, int ipfd, PIXELFORMATDESCRIPTOR* ppfd) {
// 	return (*ptrwglSetPixelFormat)(hdc, ipfd, ppfd);
// }
// BOOL gowglSwapBuffers(HDC hdc) {
// 	return (*ptrwglSwapBuffers)(hdc);
// }
// BOOL gowglShareLists(HGLRC hrcSrvShare, HGLRC hrcSrvSource) {
// 	return (*ptrwglShareLists)(hrcSrvShare, hrcSrvSource);
// }
// HGLRC gowglCreateLayerContext(HDC hDc, int level) {
// 	return (*ptrwglCreateLayerContext)(hDc, level);
// }
// BOOL gowglDescribeLayerPlane(HDC hDc, int pixelFormat, int layerPlane, UINT nBytes, LAYERPLANEDESCRIPTOR* plpd) {
// 	return (*ptrwglDescribeLayerPlane)(hDc, pixelFormat, layerPlane, nBytes, plpd);
// }
// int gowglSetLayerPaletteEntries(HDC hdc, int iLayerPlane, int iStart, int cEntries, COLORREF* pcr) {
// 	return (*ptrwglSetLayerPaletteEntries)(hdc, iLayerPlane, iStart, cEntries, pcr);
// }
// int gowglGetLayerPaletteEntries(HDC hdc, int iLayerPlane, int iStart, int cEntries, COLORREF* pcr) {
// 	return (*ptrwglGetLayerPaletteEntries)(hdc, iLayerPlane, iStart, cEntries, pcr);
// }
// BOOL gowglRealizeLayerPalette(HDC hdc, int iLayerPlane, BOOL bRealize) {
// 	return (*ptrwglRealizeLayerPalette)(hdc, iLayerPlane, bRealize);
// }
// BOOL gowglSwapLayerBuffers(HDC hdc, UINT fuFlags) {
// 	return (*ptrwglSwapLayerBuffers)(hdc, fuFlags);
// }
// BOOL gowglUseFontBitmapsA(HDC hDC, DWORD first, DWORD count, DWORD listBase) {
// 	return (*ptrwglUseFontBitmapsA)(hDC, first, count, listBase);
// }
// BOOL gowglUseFontBitmapsW(HDC hDC, DWORD first, DWORD count, DWORD listBase) {
// 	return (*ptrwglUseFontBitmapsW)(hDC, first, count, listBase);
// }
// 
// int init_3DL_stereo_control() {
// 	ptrwglSetStereoEmitterState3DL = gowglGetProcAddress("wglSetStereoEmitterState3DL");
// 	if(ptrwglSetStereoEmitterState3DL == NULL) return 1;
// 	return 0;
// }
// int init_AMD_gpu_association() {
// 	ptrwglGetGPUIDsAMD = gowglGetProcAddress("wglGetGPUIDsAMD");
// 	if(ptrwglGetGPUIDsAMD == NULL) return 1;
// 	ptrwglGetGPUInfoAMD = gowglGetProcAddress("wglGetGPUInfoAMD");
// 	if(ptrwglGetGPUInfoAMD == NULL) return 1;
// 	ptrwglGetContextGPUIDAMD = gowglGetProcAddress("wglGetContextGPUIDAMD");
// 	if(ptrwglGetContextGPUIDAMD == NULL) return 1;
// 	ptrwglCreateAssociatedContextAMD = gowglGetProcAddress("wglCreateAssociatedContextAMD");
// 	if(ptrwglCreateAssociatedContextAMD == NULL) return 1;
// 	ptrwglCreateAssociatedContextAttribsAMD = gowglGetProcAddress("wglCreateAssociatedContextAttribsAMD");
// 	if(ptrwglCreateAssociatedContextAttribsAMD == NULL) return 1;
// 	ptrwglDeleteAssociatedContextAMD = gowglGetProcAddress("wglDeleteAssociatedContextAMD");
// 	if(ptrwglDeleteAssociatedContextAMD == NULL) return 1;
// 	ptrwglMakeAssociatedContextCurrentAMD = gowglGetProcAddress("wglMakeAssociatedContextCurrentAMD");
// 	if(ptrwglMakeAssociatedContextCurrentAMD == NULL) return 1;
// 	ptrwglGetCurrentAssociatedContextAMD = gowglGetProcAddress("wglGetCurrentAssociatedContextAMD");
// 	if(ptrwglGetCurrentAssociatedContextAMD == NULL) return 1;
// 	ptrwglBlitContextFramebufferAMD = gowglGetProcAddress("wglBlitContextFramebufferAMD");
// 	if(ptrwglBlitContextFramebufferAMD == NULL) return 1;
// 	return 0;
// }
// int init_ARB_buffer_region() {
// 	ptrwglCreateBufferRegionARB = gowglGetProcAddress("wglCreateBufferRegionARB");
// 	if(ptrwglCreateBufferRegionARB == NULL) return 1;
// 	ptrwglDeleteBufferRegionARB = gowglGetProcAddress("wglDeleteBufferRegionARB");
// 	if(ptrwglDeleteBufferRegionARB == NULL) return 1;
// 	ptrwglSaveBufferRegionARB = gowglGetProcAddress("wglSaveBufferRegionARB");
// 	if(ptrwglSaveBufferRegionARB == NULL) return 1;
// 	ptrwglRestoreBufferRegionARB = gowglGetProcAddress("wglRestoreBufferRegionARB");
// 	if(ptrwglRestoreBufferRegionARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_create_context() {
// 	ptrwglCreateContextAttribsARB = gowglGetProcAddress("wglCreateContextAttribsARB");
// 	if(ptrwglCreateContextAttribsARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_extensions_string() {
// 	ptrwglGetExtensionsStringARB = gowglGetProcAddress("wglGetExtensionsStringARB");
// 	if(ptrwglGetExtensionsStringARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_make_current_read() {
// 	ptrwglMakeContextCurrentARB = gowglGetProcAddress("wglMakeContextCurrentARB");
// 	if(ptrwglMakeContextCurrentARB == NULL) return 1;
// 	ptrwglGetCurrentReadDCARB = gowglGetProcAddress("wglGetCurrentReadDCARB");
// 	if(ptrwglGetCurrentReadDCARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_pbuffer() {
// 	ptrwglCreatePbufferARB = gowglGetProcAddress("wglCreatePbufferARB");
// 	if(ptrwglCreatePbufferARB == NULL) return 1;
// 	ptrwglGetPbufferDCARB = gowglGetProcAddress("wglGetPbufferDCARB");
// 	if(ptrwglGetPbufferDCARB == NULL) return 1;
// 	ptrwglReleasePbufferDCARB = gowglGetProcAddress("wglReleasePbufferDCARB");
// 	if(ptrwglReleasePbufferDCARB == NULL) return 1;
// 	ptrwglDestroyPbufferARB = gowglGetProcAddress("wglDestroyPbufferARB");
// 	if(ptrwglDestroyPbufferARB == NULL) return 1;
// 	ptrwglQueryPbufferARB = gowglGetProcAddress("wglQueryPbufferARB");
// 	if(ptrwglQueryPbufferARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_pixel_format() {
// 	ptrwglGetPixelFormatAttribivARB = gowglGetProcAddress("wglGetPixelFormatAttribivARB");
// 	if(ptrwglGetPixelFormatAttribivARB == NULL) return 1;
// 	ptrwglGetPixelFormatAttribfvARB = gowglGetProcAddress("wglGetPixelFormatAttribfvARB");
// 	if(ptrwglGetPixelFormatAttribfvARB == NULL) return 1;
// 	ptrwglChoosePixelFormatARB = gowglGetProcAddress("wglChoosePixelFormatARB");
// 	if(ptrwglChoosePixelFormatARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_render_texture() {
// 	ptrwglBindTexImageARB = gowglGetProcAddress("wglBindTexImageARB");
// 	if(ptrwglBindTexImageARB == NULL) return 1;
// 	ptrwglReleaseTexImageARB = gowglGetProcAddress("wglReleaseTexImageARB");
// 	if(ptrwglReleaseTexImageARB == NULL) return 1;
// 	ptrwglSetPbufferAttribARB = gowglGetProcAddress("wglSetPbufferAttribARB");
// 	if(ptrwglSetPbufferAttribARB == NULL) return 1;
// 	return 0;
// }
// int init_EXT_display_color_table() {
// 	ptrwglCreateDisplayColorTableEXT = gowglGetProcAddress("wglCreateDisplayColorTableEXT");
// 	if(ptrwglCreateDisplayColorTableEXT == NULL) return 1;
// 	ptrwglLoadDisplayColorTableEXT = gowglGetProcAddress("wglLoadDisplayColorTableEXT");
// 	if(ptrwglLoadDisplayColorTableEXT == NULL) return 1;
// 	ptrwglBindDisplayColorTableEXT = gowglGetProcAddress("wglBindDisplayColorTableEXT");
// 	if(ptrwglBindDisplayColorTableEXT == NULL) return 1;
// 	ptrwglDestroyDisplayColorTableEXT = gowglGetProcAddress("wglDestroyDisplayColorTableEXT");
// 	if(ptrwglDestroyDisplayColorTableEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_extensions_string() {
// 	ptrwglGetExtensionsStringEXT = gowglGetProcAddress("wglGetExtensionsStringEXT");
// 	if(ptrwglGetExtensionsStringEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_make_current_read() {
// 	ptrwglMakeContextCurrentEXT = gowglGetProcAddress("wglMakeContextCurrentEXT");
// 	if(ptrwglMakeContextCurrentEXT == NULL) return 1;
// 	ptrwglGetCurrentReadDCEXT = gowglGetProcAddress("wglGetCurrentReadDCEXT");
// 	if(ptrwglGetCurrentReadDCEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_pbuffer() {
// 	ptrwglCreatePbufferEXT = gowglGetProcAddress("wglCreatePbufferEXT");
// 	if(ptrwglCreatePbufferEXT == NULL) return 1;
// 	ptrwglGetPbufferDCEXT = gowglGetProcAddress("wglGetPbufferDCEXT");
// 	if(ptrwglGetPbufferDCEXT == NULL) return 1;
// 	ptrwglReleasePbufferDCEXT = gowglGetProcAddress("wglReleasePbufferDCEXT");
// 	if(ptrwglReleasePbufferDCEXT == NULL) return 1;
// 	ptrwglDestroyPbufferEXT = gowglGetProcAddress("wglDestroyPbufferEXT");
// 	if(ptrwglDestroyPbufferEXT == NULL) return 1;
// 	ptrwglQueryPbufferEXT = gowglGetProcAddress("wglQueryPbufferEXT");
// 	if(ptrwglQueryPbufferEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_pixel_format() {
// 	ptrwglGetPixelFormatAttribivEXT = gowglGetProcAddress("wglGetPixelFormatAttribivEXT");
// 	if(ptrwglGetPixelFormatAttribivEXT == NULL) return 1;
// 	ptrwglGetPixelFormatAttribfvEXT = gowglGetProcAddress("wglGetPixelFormatAttribfvEXT");
// 	if(ptrwglGetPixelFormatAttribfvEXT == NULL) return 1;
// 	ptrwglChoosePixelFormatEXT = gowglGetProcAddress("wglChoosePixelFormatEXT");
// 	if(ptrwglChoosePixelFormatEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_swap_control() {
// 	ptrwglSwapIntervalEXT = gowglGetProcAddress("wglSwapIntervalEXT");
// 	if(ptrwglSwapIntervalEXT == NULL) return 1;
// 	ptrwglGetSwapIntervalEXT = gowglGetProcAddress("wglGetSwapIntervalEXT");
// 	if(ptrwglGetSwapIntervalEXT == NULL) return 1;
// 	return 0;
// }
// int init_I3D_digital_video_control() {
// 	ptrwglGetDigitalVideoParametersI3D = gowglGetProcAddress("wglGetDigitalVideoParametersI3D");
// 	if(ptrwglGetDigitalVideoParametersI3D == NULL) return 1;
// 	ptrwglSetDigitalVideoParametersI3D = gowglGetProcAddress("wglSetDigitalVideoParametersI3D");
// 	if(ptrwglSetDigitalVideoParametersI3D == NULL) return 1;
// 	return 0;
// }
// int init_I3D_gamma() {
// 	ptrwglGetGammaTableParametersI3D = gowglGetProcAddress("wglGetGammaTableParametersI3D");
// 	if(ptrwglGetGammaTableParametersI3D == NULL) return 1;
// 	ptrwglSetGammaTableParametersI3D = gowglGetProcAddress("wglSetGammaTableParametersI3D");
// 	if(ptrwglSetGammaTableParametersI3D == NULL) return 1;
// 	ptrwglGetGammaTableI3D = gowglGetProcAddress("wglGetGammaTableI3D");
// 	if(ptrwglGetGammaTableI3D == NULL) return 1;
// 	ptrwglSetGammaTableI3D = gowglGetProcAddress("wglSetGammaTableI3D");
// 	if(ptrwglSetGammaTableI3D == NULL) return 1;
// 	return 0;
// }
// int init_I3D_genlock() {
// 	ptrwglEnableGenlockI3D = gowglGetProcAddress("wglEnableGenlockI3D");
// 	if(ptrwglEnableGenlockI3D == NULL) return 1;
// 	ptrwglDisableGenlockI3D = gowglGetProcAddress("wglDisableGenlockI3D");
// 	if(ptrwglDisableGenlockI3D == NULL) return 1;
// 	ptrwglIsEnabledGenlockI3D = gowglGetProcAddress("wglIsEnabledGenlockI3D");
// 	if(ptrwglIsEnabledGenlockI3D == NULL) return 1;
// 	ptrwglGenlockSourceI3D = gowglGetProcAddress("wglGenlockSourceI3D");
// 	if(ptrwglGenlockSourceI3D == NULL) return 1;
// 	ptrwglGetGenlockSourceI3D = gowglGetProcAddress("wglGetGenlockSourceI3D");
// 	if(ptrwglGetGenlockSourceI3D == NULL) return 1;
// 	ptrwglGenlockSourceEdgeI3D = gowglGetProcAddress("wglGenlockSourceEdgeI3D");
// 	if(ptrwglGenlockSourceEdgeI3D == NULL) return 1;
// 	ptrwglGetGenlockSourceEdgeI3D = gowglGetProcAddress("wglGetGenlockSourceEdgeI3D");
// 	if(ptrwglGetGenlockSourceEdgeI3D == NULL) return 1;
// 	ptrwglGenlockSampleRateI3D = gowglGetProcAddress("wglGenlockSampleRateI3D");
// 	if(ptrwglGenlockSampleRateI3D == NULL) return 1;
// 	ptrwglGetGenlockSampleRateI3D = gowglGetProcAddress("wglGetGenlockSampleRateI3D");
// 	if(ptrwglGetGenlockSampleRateI3D == NULL) return 1;
// 	ptrwglGenlockSourceDelayI3D = gowglGetProcAddress("wglGenlockSourceDelayI3D");
// 	if(ptrwglGenlockSourceDelayI3D == NULL) return 1;
// 	ptrwglGetGenlockSourceDelayI3D = gowglGetProcAddress("wglGetGenlockSourceDelayI3D");
// 	if(ptrwglGetGenlockSourceDelayI3D == NULL) return 1;
// 	ptrwglQueryGenlockMaxSourceDelayI3D = gowglGetProcAddress("wglQueryGenlockMaxSourceDelayI3D");
// 	if(ptrwglQueryGenlockMaxSourceDelayI3D == NULL) return 1;
// 	return 0;
// }
// int init_I3D_image_buffer() {
// 	ptrwglCreateImageBufferI3D = gowglGetProcAddress("wglCreateImageBufferI3D");
// 	if(ptrwglCreateImageBufferI3D == NULL) return 1;
// 	ptrwglDestroyImageBufferI3D = gowglGetProcAddress("wglDestroyImageBufferI3D");
// 	if(ptrwglDestroyImageBufferI3D == NULL) return 1;
// 	ptrwglAssociateImageBufferEventsI3D = gowglGetProcAddress("wglAssociateImageBufferEventsI3D");
// 	if(ptrwglAssociateImageBufferEventsI3D == NULL) return 1;
// 	ptrwglReleaseImageBufferEventsI3D = gowglGetProcAddress("wglReleaseImageBufferEventsI3D");
// 	if(ptrwglReleaseImageBufferEventsI3D == NULL) return 1;
// 	return 0;
// }
// int init_I3D_swap_frame_lock() {
// 	ptrwglEnableFrameLockI3D = gowglGetProcAddress("wglEnableFrameLockI3D");
// 	if(ptrwglEnableFrameLockI3D == NULL) return 1;
// 	ptrwglDisableFrameLockI3D = gowglGetProcAddress("wglDisableFrameLockI3D");
// 	if(ptrwglDisableFrameLockI3D == NULL) return 1;
// 	ptrwglIsEnabledFrameLockI3D = gowglGetProcAddress("wglIsEnabledFrameLockI3D");
// 	if(ptrwglIsEnabledFrameLockI3D == NULL) return 1;
// 	ptrwglQueryFrameLockMasterI3D = gowglGetProcAddress("wglQueryFrameLockMasterI3D");
// 	if(ptrwglQueryFrameLockMasterI3D == NULL) return 1;
// 	return 0;
// }
// int init_I3D_swap_frame_usage() {
// 	ptrwglGetFrameUsageI3D = gowglGetProcAddress("wglGetFrameUsageI3D");
// 	if(ptrwglGetFrameUsageI3D == NULL) return 1;
// 	ptrwglBeginFrameTrackingI3D = gowglGetProcAddress("wglBeginFrameTrackingI3D");
// 	if(ptrwglBeginFrameTrackingI3D == NULL) return 1;
// 	ptrwglEndFrameTrackingI3D = gowglGetProcAddress("wglEndFrameTrackingI3D");
// 	if(ptrwglEndFrameTrackingI3D == NULL) return 1;
// 	ptrwglQueryFrameTrackingI3D = gowglGetProcAddress("wglQueryFrameTrackingI3D");
// 	if(ptrwglQueryFrameTrackingI3D == NULL) return 1;
// 	return 0;
// }
// int init_NV_DX_interop() {
// 	ptrwglDXSetResourceShareHandleNV = gowglGetProcAddress("wglDXSetResourceShareHandleNV");
// 	if(ptrwglDXSetResourceShareHandleNV == NULL) return 1;
// 	ptrwglDXOpenDeviceNV = gowglGetProcAddress("wglDXOpenDeviceNV");
// 	if(ptrwglDXOpenDeviceNV == NULL) return 1;
// 	ptrwglDXCloseDeviceNV = gowglGetProcAddress("wglDXCloseDeviceNV");
// 	if(ptrwglDXCloseDeviceNV == NULL) return 1;
// 	ptrwglDXRegisterObjectNV = gowglGetProcAddress("wglDXRegisterObjectNV");
// 	if(ptrwglDXRegisterObjectNV == NULL) return 1;
// 	ptrwglDXUnregisterObjectNV = gowglGetProcAddress("wglDXUnregisterObjectNV");
// 	if(ptrwglDXUnregisterObjectNV == NULL) return 1;
// 	ptrwglDXObjectAccessNV = gowglGetProcAddress("wglDXObjectAccessNV");
// 	if(ptrwglDXObjectAccessNV == NULL) return 1;
// 	ptrwglDXLockObjectsNV = gowglGetProcAddress("wglDXLockObjectsNV");
// 	if(ptrwglDXLockObjectsNV == NULL) return 1;
// 	ptrwglDXUnlockObjectsNV = gowglGetProcAddress("wglDXUnlockObjectsNV");
// 	if(ptrwglDXUnlockObjectsNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_copy_image() {
// 	ptrwglCopyImageSubDataNV = gowglGetProcAddress("wglCopyImageSubDataNV");
// 	if(ptrwglCopyImageSubDataNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_gpu_affinity() {
// 	ptrwglEnumGpusNV = gowglGetProcAddress("wglEnumGpusNV");
// 	if(ptrwglEnumGpusNV == NULL) return 1;
// 	ptrwglEnumGpuDevicesNV = gowglGetProcAddress("wglEnumGpuDevicesNV");
// 	if(ptrwglEnumGpuDevicesNV == NULL) return 1;
// 	ptrwglCreateAffinityDCNV = gowglGetProcAddress("wglCreateAffinityDCNV");
// 	if(ptrwglCreateAffinityDCNV == NULL) return 1;
// 	ptrwglEnumGpusFromAffinityDCNV = gowglGetProcAddress("wglEnumGpusFromAffinityDCNV");
// 	if(ptrwglEnumGpusFromAffinityDCNV == NULL) return 1;
// 	ptrwglDeleteDCNV = gowglGetProcAddress("wglDeleteDCNV");
// 	if(ptrwglDeleteDCNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_present_video() {
// 	ptrwglEnumerateVideoDevicesNV = gowglGetProcAddress("wglEnumerateVideoDevicesNV");
// 	if(ptrwglEnumerateVideoDevicesNV == NULL) return 1;
// 	ptrwglBindVideoDeviceNV = gowglGetProcAddress("wglBindVideoDeviceNV");
// 	if(ptrwglBindVideoDeviceNV == NULL) return 1;
// 	ptrwglQueryCurrentContextNV = gowglGetProcAddress("wglQueryCurrentContextNV");
// 	if(ptrwglQueryCurrentContextNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_swap_group() {
// 	ptrwglJoinSwapGroupNV = gowglGetProcAddress("wglJoinSwapGroupNV");
// 	if(ptrwglJoinSwapGroupNV == NULL) return 1;
// 	ptrwglBindSwapBarrierNV = gowglGetProcAddress("wglBindSwapBarrierNV");
// 	if(ptrwglBindSwapBarrierNV == NULL) return 1;
// 	ptrwglQuerySwapGroupNV = gowglGetProcAddress("wglQuerySwapGroupNV");
// 	if(ptrwglQuerySwapGroupNV == NULL) return 1;
// 	ptrwglQueryMaxSwapGroupsNV = gowglGetProcAddress("wglQueryMaxSwapGroupsNV");
// 	if(ptrwglQueryMaxSwapGroupsNV == NULL) return 1;
// 	ptrwglQueryFrameCountNV = gowglGetProcAddress("wglQueryFrameCountNV");
// 	if(ptrwglQueryFrameCountNV == NULL) return 1;
// 	ptrwglResetFrameCountNV = gowglGetProcAddress("wglResetFrameCountNV");
// 	if(ptrwglResetFrameCountNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_vertex_array_range() {
// 	ptrwglAllocateMemoryNV = gowglGetProcAddress("wglAllocateMemoryNV");
// 	if(ptrwglAllocateMemoryNV == NULL) return 1;
// 	ptrwglFreeMemoryNV = gowglGetProcAddress("wglFreeMemoryNV");
// 	if(ptrwglFreeMemoryNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_video_capture() {
// 	ptrwglBindVideoCaptureDeviceNV = gowglGetProcAddress("wglBindVideoCaptureDeviceNV");
// 	if(ptrwglBindVideoCaptureDeviceNV == NULL) return 1;
// 	ptrwglEnumerateVideoCaptureDevicesNV = gowglGetProcAddress("wglEnumerateVideoCaptureDevicesNV");
// 	if(ptrwglEnumerateVideoCaptureDevicesNV == NULL) return 1;
// 	ptrwglLockVideoCaptureDeviceNV = gowglGetProcAddress("wglLockVideoCaptureDeviceNV");
// 	if(ptrwglLockVideoCaptureDeviceNV == NULL) return 1;
// 	ptrwglQueryVideoCaptureDeviceNV = gowglGetProcAddress("wglQueryVideoCaptureDeviceNV");
// 	if(ptrwglQueryVideoCaptureDeviceNV == NULL) return 1;
// 	ptrwglReleaseVideoCaptureDeviceNV = gowglGetProcAddress("wglReleaseVideoCaptureDeviceNV");
// 	if(ptrwglReleaseVideoCaptureDeviceNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_video_output() {
// 	ptrwglGetVideoDeviceNV = gowglGetProcAddress("wglGetVideoDeviceNV");
// 	if(ptrwglGetVideoDeviceNV == NULL) return 1;
// 	ptrwglReleaseVideoDeviceNV = gowglGetProcAddress("wglReleaseVideoDeviceNV");
// 	if(ptrwglReleaseVideoDeviceNV == NULL) return 1;
// 	ptrwglBindVideoImageNV = gowglGetProcAddress("wglBindVideoImageNV");
// 	if(ptrwglBindVideoImageNV == NULL) return 1;
// 	ptrwglReleaseVideoImageNV = gowglGetProcAddress("wglReleaseVideoImageNV");
// 	if(ptrwglReleaseVideoImageNV == NULL) return 1;
// 	ptrwglSendPbufferToVideoNV = gowglGetProcAddress("wglSendPbufferToVideoNV");
// 	if(ptrwglSendPbufferToVideoNV == NULL) return 1;
// 	ptrwglGetVideoInfoNV = gowglGetProcAddress("wglGetVideoInfoNV");
// 	if(ptrwglGetVideoInfoNV == NULL) return 1;
// 	return 0;
// }
// int init_OML_sync_control() {
// 	ptrwglGetSyncValuesOML = gowglGetProcAddress("wglGetSyncValuesOML");
// 	if(ptrwglGetSyncValuesOML == NULL) return 1;
// 	ptrwglGetMscRateOML = gowglGetProcAddress("wglGetMscRateOML");
// 	if(ptrwglGetMscRateOML == NULL) return 1;
// 	ptrwglSwapBuffersMscOML = gowglGetProcAddress("wglSwapBuffersMscOML");
// 	if(ptrwglSwapBuffersMscOML == NULL) return 1;
// 	ptrwglSwapLayerBuffersMscOML = gowglGetProcAddress("wglSwapLayerBuffersMscOML");
// 	if(ptrwglSwapLayerBuffersMscOML == NULL) return 1;
// 	ptrwglWaitForMscOML = gowglGetProcAddress("wglWaitForMscOML");
// 	if(ptrwglWaitForMscOML == NULL) return 1;
// 	ptrwglWaitForSbcOML = gowglGetProcAddress("wglWaitForSbcOML");
// 	if(ptrwglWaitForSbcOML == NULL) return 1;
// 	return 0;
// }
// int init_wgl() {
// 	ptrwglCreateContext = gowglGetProcAddress("wglCreateContext");
// 	if(ptrwglCreateContext == NULL) return 1;
// 	ptrwglDeleteContext = gowglGetProcAddress("wglDeleteContext");
// 	if(ptrwglDeleteContext == NULL) return 1;
// 	ptrwglGetCurrentContext = gowglGetProcAddress("wglGetCurrentContext");
// 	if(ptrwglGetCurrentContext == NULL) return 1;
// 	ptrwglMakeCurrent = gowglGetProcAddress("wglMakeCurrent");
// 	if(ptrwglMakeCurrent == NULL) return 1;
// 	ptrwglCopyContext = gowglGetProcAddress("wglCopyContext");
// 	if(ptrwglCopyContext == NULL) return 1;
// 	ptrwglChoosePixelFormat = gowglGetProcAddress("wglChoosePixelFormat");
// 	if(ptrwglChoosePixelFormat == NULL) return 1;
// 	ptrwglGetCurrentDC = gowglGetProcAddress("wglGetCurrentDC");
// 	if(ptrwglGetCurrentDC == NULL) return 1;
// 	ptrwglGetDefaultProcAddress = gowglGetProcAddress("wglGetDefaultProcAddress");
// 	if(ptrwglGetDefaultProcAddress == NULL) return 1;
// 	ptrwglGetProcAddress = gowglGetProcAddress("wglGetProcAddress");
// 	if(ptrwglGetProcAddress == NULL) return 1;
// 	ptrwglGetPixelFormat = gowglGetProcAddress("wglGetPixelFormat");
// 	if(ptrwglGetPixelFormat == NULL) return 1;
// 	ptrwglSetPixelFormat = gowglGetProcAddress("wglSetPixelFormat");
// 	if(ptrwglSetPixelFormat == NULL) return 1;
// 	ptrwglSwapBuffers = gowglGetProcAddress("wglSwapBuffers");
// 	if(ptrwglSwapBuffers == NULL) return 1;
// 	ptrwglShareLists = gowglGetProcAddress("wglShareLists");
// 	if(ptrwglShareLists == NULL) return 1;
// 	ptrwglCreateLayerContext = gowglGetProcAddress("wglCreateLayerContext");
// 	if(ptrwglCreateLayerContext == NULL) return 1;
// 	ptrwglDescribeLayerPlane = gowglGetProcAddress("wglDescribeLayerPlane");
// 	if(ptrwglDescribeLayerPlane == NULL) return 1;
// 	ptrwglSetLayerPaletteEntries = gowglGetProcAddress("wglSetLayerPaletteEntries");
// 	if(ptrwglSetLayerPaletteEntries == NULL) return 1;
// 	ptrwglGetLayerPaletteEntries = gowglGetProcAddress("wglGetLayerPaletteEntries");
// 	if(ptrwglGetLayerPaletteEntries == NULL) return 1;
// 	ptrwglRealizeLayerPalette = gowglGetProcAddress("wglRealizeLayerPalette");
// 	if(ptrwglRealizeLayerPalette == NULL) return 1;
// 	ptrwglSwapLayerBuffers = gowglGetProcAddress("wglSwapLayerBuffers");
// 	if(ptrwglSwapLayerBuffers == NULL) return 1;
// 	ptrwglUseFontBitmapsA = gowglGetProcAddress("wglUseFontBitmapsA");
// 	if(ptrwglUseFontBitmapsA == NULL) return 1;
// 	ptrwglUseFontBitmapsW = gowglGetProcAddress("wglUseFontBitmapsW");
// 	if(ptrwglUseFontBitmapsW == NULL) return 1;
// 	return 0;
// }
// 
import "C"
import "unsafe"
import "errors"

type (
	Enum     C.GLenum
	Boolean  C.GLboolean
	Bitfield C.GLbitfield
	Byte     C.GLbyte
	Short    C.GLshort
	Int      C.GLint
	Sizei    C.GLsizei
	Ubyte    C.GLubyte
	Ushort   C.GLushort
	Uint     C.GLuint
	Half     C.GLhalf
	Float    C.GLfloat
	Clampf   C.GLclampf
	Double   C.GLdouble
	Clampd   C.GLclampd
	Char     C.GLchar
	Pointer  unsafe.Pointer
	Sync     C.GLsync
	Int64    C.GLint64
	Uint64   C.GLuint64
	Intptr   C.GLintptr
	Sizeiptr C.GLsizeiptr
)

// WGL_3DFX_multisample
const (
	WGL_SAMPLES_3DFX = 0x2061
	WGL_SAMPLE_BUFFERS_3DFX = 0x2060
)
// WGL_3DL_stereo_control
const (
	WGL_STEREO_EMITTER_DISABLE_3DL = 0x2056
	WGL_STEREO_EMITTER_ENABLE_3DL = 0x2055
	WGL_STEREO_POLARITY_INVERT_3DL = 0x2058
	WGL_STEREO_POLARITY_NORMAL_3DL = 0x2057
)
// WGL_AMD_gpu_association
const (
	WGL_GPU_CLOCK_AMD = 0x21A4
	WGL_GPU_FASTEST_TARGET_GPUS_AMD = 0x21A2
	WGL_GPU_NUM_PIPES_AMD = 0x21A5
	WGL_GPU_NUM_RB_AMD = 0x21A7
	WGL_GPU_NUM_SIMD_AMD = 0x21A6
	WGL_GPU_NUM_SPI_AMD = 0x21A8
	WGL_GPU_OPENGL_VERSION_STRING_AMD = 0x1F02
	WGL_GPU_RAM_AMD = 0x21A3
	WGL_GPU_RENDERER_STRING_AMD = 0x1F01
	WGL_GPU_VENDOR_AMD = 0x1F00
)
// WGL_ARB_buffer_region
const (
	WGL_BACK_COLOR_BUFFER_BIT_ARB = 0x00000002
	WGL_DEPTH_BUFFER_BIT_ARB = 0x00000004
	WGL_FRONT_COLOR_BUFFER_BIT_ARB = 0x00000001
	WGL_STENCIL_BUFFER_BIT_ARB = 0x00000008
)
// WGL_ARB_create_context
const (
	ERROR_INVALID_VERSION_ARB = 0x2095
	WGL_CONTEXT_DEBUG_BIT_ARB = 0x00000001
	WGL_CONTEXT_FLAGS_ARB = 0x2094
	WGL_CONTEXT_FORWARD_COMPATIBLE_BIT_ARB = 0x00000002
	WGL_CONTEXT_LAYER_PLANE_ARB = 0x2093
	WGL_CONTEXT_MAJOR_VERSION_ARB = 0x2091
	WGL_CONTEXT_MINOR_VERSION_ARB = 0x2092
)
// WGL_ARB_create_context_profile
const (
	ERROR_INVALID_PROFILE_ARB = 0x2096
	WGL_CONTEXT_COMPATIBILITY_PROFILE_BIT_ARB = 0x00000002
	WGL_CONTEXT_CORE_PROFILE_BIT_ARB = 0x00000001
	WGL_CONTEXT_PROFILE_MASK_ARB = 0x9126
)
// WGL_ARB_create_context_robustness
const (
	WGL_CONTEXT_RESET_NOTIFICATION_STRATEGY_ARB = 0x8256
	WGL_CONTEXT_ROBUST_ACCESS_BIT_ARB = 0x00000004
	WGL_LOSE_CONTEXT_ON_RESET_ARB = 0x8252
	WGL_NO_RESET_NOTIFICATION_ARB = 0x8261
)
// WGL_ARB_extensions_string
const (
)
// WGL_ARB_framebuffer_sRGB
const (
	WGL_FRAMEBUFFER_SRGB_CAPABLE_ARB = 0x20A9
)
// WGL_ARB_make_current_read
const (
	ERROR_INCOMPATIBLE_DEVICE_CONTEXTS_ARB = 0x2054
	ERROR_INVALID_PIXEL_TYPE_ARB = 0x2043
)
// WGL_ARB_multisample
const (
	WGL_SAMPLES_ARB = 0x2042
	WGL_SAMPLE_BUFFERS_ARB = 0x2041
)
// WGL_ARB_pbuffer
const (
	WGL_DRAW_TO_PBUFFER_ARB = 0x202D
	WGL_MAX_PBUFFER_HEIGHT_ARB = 0x2030
	WGL_MAX_PBUFFER_PIXELS_ARB = 0x202E
	WGL_MAX_PBUFFER_WIDTH_ARB = 0x202F
	WGL_PBUFFER_HEIGHT_ARB = 0x2035
	WGL_PBUFFER_LARGEST_ARB = 0x2033
	WGL_PBUFFER_LOST_ARB = 0x2036
	WGL_PBUFFER_WIDTH_ARB = 0x2034
)
// WGL_ARB_pixel_format
const (
	WGL_ACCELERATION_ARB = 0x2003
	WGL_ACCUM_ALPHA_BITS_ARB = 0x2021
	WGL_ACCUM_BITS_ARB = 0x201D
	WGL_ACCUM_BLUE_BITS_ARB = 0x2020
	WGL_ACCUM_GREEN_BITS_ARB = 0x201F
	WGL_ACCUM_RED_BITS_ARB = 0x201E
	WGL_ALPHA_BITS_ARB = 0x201B
	WGL_ALPHA_SHIFT_ARB = 0x201C
	WGL_AUX_BUFFERS_ARB = 0x2024
	WGL_BLUE_BITS_ARB = 0x2019
	WGL_BLUE_SHIFT_ARB = 0x201A
	WGL_COLOR_BITS_ARB = 0x2014
	WGL_DEPTH_BITS_ARB = 0x2022
	WGL_DOUBLE_BUFFER_ARB = 0x2011
	WGL_DRAW_TO_BITMAP_ARB = 0x2002
	WGL_DRAW_TO_WINDOW_ARB = 0x2001
	WGL_FULL_ACCELERATION_ARB = 0x2027
	WGL_GENERIC_ACCELERATION_ARB = 0x2026
	WGL_GREEN_BITS_ARB = 0x2017
	WGL_GREEN_SHIFT_ARB = 0x2018
	WGL_NEED_PALETTE_ARB = 0x2004
	WGL_NEED_SYSTEM_PALETTE_ARB = 0x2005
	WGL_NO_ACCELERATION_ARB = 0x2025
	WGL_NUMBER_OVERLAYS_ARB = 0x2008
	WGL_NUMBER_PIXEL_FORMATS_ARB = 0x2000
	WGL_NUMBER_UNDERLAYS_ARB = 0x2009
	WGL_PIXEL_TYPE_ARB = 0x2013
	WGL_RED_BITS_ARB = 0x2015
	WGL_RED_SHIFT_ARB = 0x2016
	WGL_SHARE_ACCUM_ARB = 0x200E
	WGL_SHARE_DEPTH_ARB = 0x200C
	WGL_SHARE_STENCIL_ARB = 0x200D
	WGL_STENCIL_BITS_ARB = 0x2023
	WGL_STEREO_ARB = 0x2012
	WGL_SUPPORT_GDI_ARB = 0x200F
	WGL_SUPPORT_OPENGL_ARB = 0x2010
	WGL_SWAP_COPY_ARB = 0x2029
	WGL_SWAP_EXCHANGE_ARB = 0x2028
	WGL_SWAP_LAYER_BUFFERS_ARB = 0x2006
	WGL_SWAP_METHOD_ARB = 0x2007
	WGL_SWAP_UNDEFINED_ARB = 0x202A
	WGL_TRANSPARENT_ARB = 0x200A
	WGL_TYPE_COLORINDEX_ARB = 0x202C
	WGL_TYPE_RGBA_ARB = 0x202B
)
// WGL_ARB_pixel_format_float
const (
	WGL_TYPE_RGBA_FLOAT_ARB = 0x21A0
)
// WGL_ARB_render_texture
const (
	WGL_AUX0_ARB = 0x2087
	WGL_AUX1_ARB = 0x2088
	WGL_AUX2_ARB = 0x2089
	WGL_AUX3_ARB = 0x208A
	WGL_AUX4_ARB = 0x208B
	WGL_AUX5_ARB = 0x208C
	WGL_AUX6_ARB = 0x208D
	WGL_AUX7_ARB = 0x208E
	WGL_AUX8_ARB = 0x208F
	WGL_AUX9_ARB = 0x2090
	WGL_BACK_LEFT_ARB = 0x2085
	WGL_BACK_RIGHT_ARB = 0x2086
	WGL_BIND_TO_TEXTURE_RGBA_ARB = 0x2071
	WGL_BIND_TO_TEXTURE_RGB_ARB = 0x2070
	WGL_CUBE_MAP_FACE_ARB = 0x207C
	WGL_FRONT_LEFT_ARB = 0x2083
	WGL_FRONT_RIGHT_ARB = 0x2084
	WGL_MIPMAP_LEVEL_ARB = 0x207B
	WGL_MIPMAP_TEXTURE_ARB = 0x2074
	WGL_NO_TEXTURE_ARB = 0x2077
	WGL_TEXTURE_1D_ARB = 0x2079
	WGL_TEXTURE_2D_ARB = 0x207A
	WGL_TEXTURE_CUBE_MAP_ARB = 0x2078
	WGL_TEXTURE_CUBE_MAP_NEGATIVE_X_ARB = 0x207E
	WGL_TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB = 0x2080
	WGL_TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB = 0x2082
	WGL_TEXTURE_CUBE_MAP_POSITIVE_X_ARB = 0x207D
	WGL_TEXTURE_CUBE_MAP_POSITIVE_Y_ARB = 0x207F
	WGL_TEXTURE_CUBE_MAP_POSITIVE_Z_ARB = 0x2081
	WGL_TEXTURE_FORMAT_ARB = 0x2072
	WGL_TEXTURE_RGBA_ARB = 0x2076
	WGL_TEXTURE_RGB_ARB = 0x2075
	WGL_TEXTURE_TARGET_ARB = 0x2073
)
// WGL_ATI_pixel_format_float
const (
	WGL_TYPE_RGBA_FLOAT_ATI = 0x21A0
)
// WGL_EXT_create_context_es2_profile
const (
	WGL_CONTEXT_ES2_PROFILE_BIT_EXT = 0x00000004
)
// WGL_EXT_depth_float
const (
	WGL_DEPTH_FLOAT_EXT = 0x2040
)
// WGL_EXT_framebuffer_sRGB
const (
	WGL_FRAMEBUFFER_SRGB_CAPABLE_EXT = 0x20A9
)
// WGL_EXT_make_current_read
const (
	ERROR_INVALID_PIXEL_TYPE_EXT = 0x2043
)
// WGL_EXT_multisample
const (
	WGL_SAMPLES_EXT = 0x2042
	WGL_SAMPLE_BUFFERS_EXT = 0x2041
)
// WGL_EXT_pbuffer
const (
	WGL_DRAW_TO_PBUFFER_EXT = 0x202D
	WGL_MAX_PBUFFER_HEIGHT_EXT = 0x2030
	WGL_MAX_PBUFFER_PIXELS_EXT = 0x202E
	WGL_MAX_PBUFFER_WIDTH_EXT = 0x202F
	WGL_OPTIMAL_PBUFFER_HEIGHT_EXT = 0x2032
	WGL_OPTIMAL_PBUFFER_WIDTH_EXT = 0x2031
	WGL_PBUFFER_HEIGHT_EXT = 0x2035
	WGL_PBUFFER_LARGEST_EXT = 0x2033
	WGL_PBUFFER_WIDTH_EXT = 0x2034
)
// WGL_EXT_pixel_format
const (
	WGL_ACCELERATION_EXT = 0x2003
	WGL_ACCUM_ALPHA_BITS_EXT = 0x2021
	WGL_ACCUM_BITS_EXT = 0x201D
	WGL_ACCUM_BLUE_BITS_EXT = 0x2020
	WGL_ACCUM_GREEN_BITS_EXT = 0x201F
	WGL_ACCUM_RED_BITS_EXT = 0x201E
	WGL_ALPHA_BITS_EXT = 0x201B
	WGL_ALPHA_SHIFT_EXT = 0x201C
	WGL_AUX_BUFFERS_EXT = 0x2024
	WGL_BLUE_BITS_EXT = 0x2019
	WGL_BLUE_SHIFT_EXT = 0x201A
	WGL_COLOR_BITS_EXT = 0x2014
	WGL_DEPTH_BITS_EXT = 0x2022
	WGL_DOUBLE_BUFFER_EXT = 0x2011
	WGL_DRAW_TO_BITMAP_EXT = 0x2002
	WGL_DRAW_TO_WINDOW_EXT = 0x2001
	WGL_FULL_ACCELERATION_EXT = 0x2027
	WGL_GENERIC_ACCELERATION_EXT = 0x2026
	WGL_GREEN_BITS_EXT = 0x2017
	WGL_GREEN_SHIFT_EXT = 0x2018
	WGL_NEED_PALETTE_EXT = 0x2004
	WGL_NEED_SYSTEM_PALETTE_EXT = 0x2005
	WGL_NO_ACCELERATION_EXT = 0x2025
	WGL_NUMBER_OVERLAYS_EXT = 0x2008
	WGL_NUMBER_PIXEL_FORMATS_EXT = 0x2000
	WGL_NUMBER_UNDERLAYS_EXT = 0x2009
	WGL_PIXEL_TYPE_EXT = 0x2013
	WGL_RED_BITS_EXT = 0x2015
	WGL_RED_SHIFT_EXT = 0x2016
	WGL_SHARE_ACCUM_EXT = 0x200E
	WGL_SHARE_DEPTH_EXT = 0x200C
	WGL_SHARE_STENCIL_EXT = 0x200D
	WGL_STENCIL_BITS_EXT = 0x2023
	WGL_STEREO_EXT = 0x2012
	WGL_SUPPORT_GDI_EXT = 0x200F
	WGL_SUPPORT_OPENGL_EXT = 0x2010
	WGL_SWAP_COPY_EXT = 0x2029
	WGL_SWAP_EXCHANGE_EXT = 0x2028
	WGL_SWAP_LAYER_BUFFERS_EXT = 0x2006
	WGL_SWAP_METHOD_EXT = 0x2007
	WGL_SWAP_UNDEFINED_EXT = 0x202A
	WGL_TRANSPARENT_EXT = 0x200A
	WGL_TRANSPARENT_VALUE_EXT = 0x200B
	WGL_TYPE_COLORINDEX_EXT = 0x202C
	WGL_TYPE_RGBA_EXT = 0x202B
)
// WGL_EXT_pixel_format_packed_float
const (
	WGL_TYPE_RGBA_UNSIGNED_FLOAT_EXT = 0x20A8
)
// WGL_EXT_swap_control
const (
)
// WGL_EXT_swap_control_tear
const (
)
// WGL_I3D_digital_video_control
const (
	WGL_DIGITAL_VIDEO_CURSOR_ALPHA_FRAMEBUFFER_I3D = 0x2050
	WGL_DIGITAL_VIDEO_CURSOR_ALPHA_VALUE_I3D = 0x2051
	WGL_DIGITAL_VIDEO_CURSOR_INCLUDED_I3D = 0x2052
	WGL_DIGITAL_VIDEO_GAMMA_CORRECTED_I3D = 0x2053
)
// WGL_I3D_gamma
const (
	WGL_GAMMA_EXCLUDE_DESKTOP_I3D = 0x204F
	WGL_GAMMA_TABLE_SIZE_I3D = 0x204E
)
// WGL_I3D_genlock
const (
	WGL_GENLOCK_SOURCE_DIGITAL_FIELD_I3D = 0x2049
	WGL_GENLOCK_SOURCE_DIGITAL_SYNC_I3D = 0x2048
	WGL_GENLOCK_SOURCE_EDGE_BOTH_I3D = 0x204C
	WGL_GENLOCK_SOURCE_EDGE_FALLING_I3D = 0x204A
	WGL_GENLOCK_SOURCE_EDGE_RISING_I3D = 0x204B
	WGL_GENLOCK_SOURCE_EXTERNAL_FIELD_I3D = 0x2046
	WGL_GENLOCK_SOURCE_EXTERNAL_SYNC_I3D = 0x2045
	WGL_GENLOCK_SOURCE_EXTERNAL_TTL_I3D = 0x2047
	WGL_GENLOCK_SOURCE_MULTIVIEW_I3D = 0x2044
)
// WGL_I3D_image_buffer
const (
	WGL_IMAGE_BUFFER_LOCK_I3D = 0x00000002
	WGL_IMAGE_BUFFER_MIN_ACCESS_I3D = 0x00000001
)
// WGL_I3D_swap_frame_lock
const (
)
// WGL_NV_DX_interop
const (
	WGL_ACCESS_READ_ONLY_NV = 0x00000000
	WGL_ACCESS_READ_WRITE_NV = 0x00000001
	WGL_ACCESS_WRITE_DISCARD_NV = 0x00000002
)
// WGL_NV_DX_interop2
const (
)
// WGL_NV_copy_image
const (
)
// WGL_NV_float_buffer
const (
	WGL_BIND_TO_TEXTURE_RECTANGLE_FLOAT_RGBA_NV = 0x20B4
	WGL_BIND_TO_TEXTURE_RECTANGLE_FLOAT_RGB_NV = 0x20B3
	WGL_BIND_TO_TEXTURE_RECTANGLE_FLOAT_RG_NV = 0x20B2
	WGL_BIND_TO_TEXTURE_RECTANGLE_FLOAT_R_NV = 0x20B1
	WGL_FLOAT_COMPONENTS_NV = 0x20B0
	WGL_TEXTURE_FLOAT_RGBA_NV = 0x20B8
	WGL_TEXTURE_FLOAT_RGB_NV = 0x20B7
	WGL_TEXTURE_FLOAT_RG_NV = 0x20B6
	WGL_TEXTURE_FLOAT_R_NV = 0x20B5
)
// WGL_NV_gpu_affinity
const (
	WGL_ERROR_INCOMPATIBLE_AFFINITY_MASKS_NV = 0x20D0
	WGL_ERROR_MISSING_AFFINITY_MASK_NV = 0x20D1
)
// WGL_NV_multisample_coverage
const (
	WGL_COLOR_SAMPLES_NV = 0x20B9
	WGL_COVERAGE_SAMPLES_NV = 0x2042
)
// WGL_NV_present_video
const (
	WGL_NUM_VIDEO_SLOTS_NV = 0x20F0
)
// WGL_NV_render_depth_texture
const (
	WGL_BIND_TO_TEXTURE_DEPTH_NV = 0x20A3
	WGL_BIND_TO_TEXTURE_RECTANGLE_DEPTH_NV = 0x20A4
	WGL_DEPTH_COMPONENT_NV = 0x20A7
	WGL_DEPTH_TEXTURE_FORMAT_NV = 0x20A5
	WGL_TEXTURE_DEPTH_COMPONENT_NV = 0x20A6
)
// WGL_NV_render_texture_rectangle
const (
	WGL_BIND_TO_TEXTURE_RECTANGLE_RGBA_NV = 0x20A1
	WGL_BIND_TO_TEXTURE_RECTANGLE_RGB_NV = 0x20A0
	WGL_TEXTURE_RECTANGLE_NV = 0x20A2
)
// WGL_NV_swap_group
const (
)
// WGL_NV_video_capture
const (
	WGL_NUM_VIDEO_CAPTURE_SLOTS_NV = 0x20CF
	WGL_UNIQUE_ID_NV = 0x20CE
)
// WGL_NV_video_output
const (
	WGL_BIND_TO_VIDEO_RGBA_NV = 0x20C1
	WGL_BIND_TO_VIDEO_RGB_AND_DEPTH_NV = 0x20C2
	WGL_BIND_TO_VIDEO_RGB_NV = 0x20C0
	WGL_VIDEO_OUT_ALPHA_NV = 0x20C4
	WGL_VIDEO_OUT_COLOR_AND_ALPHA_NV = 0x20C6
	WGL_VIDEO_OUT_COLOR_AND_DEPTH_NV = 0x20C7
	WGL_VIDEO_OUT_COLOR_NV = 0x20C3
	WGL_VIDEO_OUT_DEPTH_NV = 0x20C5
	WGL_VIDEO_OUT_FIELD_1 = 0x20C9
	WGL_VIDEO_OUT_FIELD_2 = 0x20CA
	WGL_VIDEO_OUT_FRAME = 0x20C8
	WGL_VIDEO_OUT_STACKED_FIELDS_1_2 = 0x20CB
	WGL_VIDEO_OUT_STACKED_FIELDS_2_1 = 0x20CC
)
// 3DL_stereo_control

func SetStereoEmitterState3DL(hDC Pointer, uState uint32) int32 {
	return (int32)(C.gowglSetStereoEmitterState3DL((C.HDC)(hDC), (C.UINT)(uState)))
}
// AMD_gpu_association

func GetGPUIDsAMD(maxCount uint32, ids *uint32) uint32 {
	return (uint32)(C.gowglGetGPUIDsAMD((C.UINT)(maxCount), (*C.UINT)(ids)))
}
func GetGPUInfoAMD(id uint32, property Int, dataType Enum, size uint32, data Pointer) int32 {
	return (int32)(C.gowglGetGPUInfoAMD((C.UINT)(id), (C.int)(property), (C.GLenum)(dataType), (C.UINT)(size), (unsafe.Pointer)(data)))
}
func GetContextGPUIDAMD(hglrc Pointer) uint32 {
	return (uint32)(C.gowglGetContextGPUIDAMD((C.HGLRC)(hglrc)))
}
func CreateAssociatedContextAMD(id uint32) Pointer {
	return (Pointer)(C.gowglCreateAssociatedContextAMD((C.UINT)(id)))
}
func CreateAssociatedContextAttribsAMD(id uint32, hShareContext Pointer, attribList *Int) Pointer {
	return (Pointer)(C.gowglCreateAssociatedContextAttribsAMD((C.UINT)(id), (C.HGLRC)(hShareContext), (*C.int)(attribList)))
}
func DeleteAssociatedContextAMD(hglrc Pointer) int32 {
	return (int32)(C.gowglDeleteAssociatedContextAMD((C.HGLRC)(hglrc)))
}
func MakeAssociatedContextCurrentAMD(hglrc Pointer) int32 {
	return (int32)(C.gowglMakeAssociatedContextCurrentAMD((C.HGLRC)(hglrc)))
}
func GetCurrentAssociatedContextAMD() Pointer {
	return (Pointer)(C.gowglGetCurrentAssociatedContextAMD())
}
func BlitContextFramebufferAMD(dstCtx Pointer, srcX0 Int, srcY0 Int, srcX1 Int, srcY1 Int, dstX0 Int, dstY0 Int, dstX1 Int, dstY1 Int, mask Bitfield, filter Enum)  {
	C.gowglBlitContextFramebufferAMD((C.HGLRC)(dstCtx), (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}
// ARB_buffer_region

func CreateBufferRegionARB(hDC Pointer, iLayerPlane Int, uType uint32) Pointer {
	return (Pointer)(C.gowglCreateBufferRegionARB((C.HDC)(hDC), (C.int)(iLayerPlane), (C.UINT)(uType)))
}
func DeleteBufferRegionARB(hRegion Pointer)  {
	C.gowglDeleteBufferRegionARB((C.HANDLE)(hRegion))
}
func SaveBufferRegionARB(hRegion Pointer, x Int, y Int, width Int, height Int) int32 {
	return (int32)(C.gowglSaveBufferRegionARB((C.HANDLE)(hRegion), (C.int)(x), (C.int)(y), (C.int)(width), (C.int)(height)))
}
func RestoreBufferRegionARB(hRegion Pointer, x Int, y Int, width Int, height Int, xSrc Int, ySrc Int) int32 {
	return (int32)(C.gowglRestoreBufferRegionARB((C.HANDLE)(hRegion), (C.int)(x), (C.int)(y), (C.int)(width), (C.int)(height), (C.int)(xSrc), (C.int)(ySrc)))
}
// ARB_create_context

func CreateContextAttribsARB(hDC Pointer, hShareContext Pointer, attribList *Int) Pointer {
	return (Pointer)(C.gowglCreateContextAttribsARB((C.HDC)(hDC), (C.HGLRC)(hShareContext), (*C.int)(attribList)))
}
// ARB_extensions_string

func GetExtensionsStringARB(hdc Pointer) *byte {
	return (*byte)(C.gowglGetExtensionsStringARB((C.HDC)(hdc)))
}
// ARB_make_current_read

func MakeContextCurrentARB(hDrawDC Pointer, hReadDC Pointer, hglrc Pointer) int32 {
	return (int32)(C.gowglMakeContextCurrentARB((C.HDC)(hDrawDC), (C.HDC)(hReadDC), (C.HGLRC)(hglrc)))
}
func GetCurrentReadDCARB() Pointer {
	return (Pointer)(C.gowglGetCurrentReadDCARB())
}
// ARB_pbuffer

func CreatePbufferARB(hDC Pointer, iPixelFormat Int, iWidth Int, iHeight Int, piAttribList *Int) Pointer {
	return (Pointer)(C.gowglCreatePbufferARB((C.HDC)(hDC), (C.int)(iPixelFormat), (C.int)(iWidth), (C.int)(iHeight), (*C.int)(piAttribList)))
}
func GetPbufferDCARB(hPbuffer Pointer) Pointer {
	return (Pointer)(C.gowglGetPbufferDCARB((C.HPBUFFERARB)(hPbuffer)))
}
func ReleasePbufferDCARB(hPbuffer Pointer, hDC Pointer) Int {
	return (Int)(C.gowglReleasePbufferDCARB((C.HPBUFFERARB)(hPbuffer), (C.HDC)(hDC)))
}
func DestroyPbufferARB(hPbuffer Pointer) int32 {
	return (int32)(C.gowglDestroyPbufferARB((C.HPBUFFERARB)(hPbuffer)))
}
func QueryPbufferARB(hPbuffer Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.gowglQueryPbufferARB((C.HPBUFFERARB)(hPbuffer), (C.int)(iAttribute), (*C.int)(piValue)))
}
// ARB_pixel_format

func GetPixelFormatAttribivARB(hdc Pointer, iPixelFormat Int, iLayerPlane Int, nAttributes uint32, piAttributes *Int, piValues *Int) int32 {
	return (int32)(C.gowglGetPixelFormatAttribivARB((C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(piAttributes), (*C.int)(piValues)))
}
func GetPixelFormatAttribfvARB(hdc Pointer, iPixelFormat Int, iLayerPlane Int, nAttributes uint32, piAttributes *Int, pfValues *float32) int32 {
	return (int32)(C.gowglGetPixelFormatAttribfvARB((C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(piAttributes), (*C.FLOAT)(pfValues)))
}
func ChoosePixelFormatARB(hdc Pointer, piAttribIList *Int, pfAttribFList *float32, nMaxFormats uint32, piFormats *Int, nNumFormats *uint32) int32 {
	return (int32)(C.gowglChoosePixelFormatARB((C.HDC)(hdc), (*C.int)(piAttribIList), (*C.FLOAT)(pfAttribFList), (C.UINT)(nMaxFormats), (*C.int)(piFormats), (*C.UINT)(nNumFormats)))
}
// ARB_render_texture

func BindTexImageARB(hPbuffer Pointer, iBuffer Int) int32 {
	return (int32)(C.gowglBindTexImageARB((C.HPBUFFERARB)(hPbuffer), (C.int)(iBuffer)))
}
func ReleaseTexImageARB(hPbuffer Pointer, iBuffer Int) int32 {
	return (int32)(C.gowglReleaseTexImageARB((C.HPBUFFERARB)(hPbuffer), (C.int)(iBuffer)))
}
func SetPbufferAttribARB(hPbuffer Pointer, piAttribList *Int) int32 {
	return (int32)(C.gowglSetPbufferAttribARB((C.HPBUFFERARB)(hPbuffer), (*C.int)(piAttribList)))
}
// EXT_display_color_table

func CreateDisplayColorTableEXT(id Ushort) Boolean {
	return (Boolean)(C.gowglCreateDisplayColorTableEXT((C.GLushort)(id)))
}
func LoadDisplayColorTableEXT(table *Ushort, length Uint) Boolean {
	return (Boolean)(C.gowglLoadDisplayColorTableEXT((*C.GLushort)(table), (C.GLuint)(length)))
}
func BindDisplayColorTableEXT(id Ushort) Boolean {
	return (Boolean)(C.gowglBindDisplayColorTableEXT((C.GLushort)(id)))
}
func DestroyDisplayColorTableEXT(id Ushort)  {
	C.gowglDestroyDisplayColorTableEXT((C.GLushort)(id))
}
// EXT_extensions_string

func GetExtensionsStringEXT() *byte {
	return (*byte)(C.gowglGetExtensionsStringEXT())
}
// EXT_make_current_read

func MakeContextCurrentEXT(hDrawDC Pointer, hReadDC Pointer, hglrc Pointer) int32 {
	return (int32)(C.gowglMakeContextCurrentEXT((C.HDC)(hDrawDC), (C.HDC)(hReadDC), (C.HGLRC)(hglrc)))
}
func GetCurrentReadDCEXT() Pointer {
	return (Pointer)(C.gowglGetCurrentReadDCEXT())
}
// EXT_pbuffer

func CreatePbufferEXT(hDC Pointer, iPixelFormat Int, iWidth Int, iHeight Int, piAttribList *Int) Pointer {
	return (Pointer)(C.gowglCreatePbufferEXT((C.HDC)(hDC), (C.int)(iPixelFormat), (C.int)(iWidth), (C.int)(iHeight), (*C.int)(piAttribList)))
}
func GetPbufferDCEXT(hPbuffer Pointer) Pointer {
	return (Pointer)(C.gowglGetPbufferDCEXT((C.HPBUFFEREXT)(hPbuffer)))
}
func ReleasePbufferDCEXT(hPbuffer Pointer, hDC Pointer) Int {
	return (Int)(C.gowglReleasePbufferDCEXT((C.HPBUFFEREXT)(hPbuffer), (C.HDC)(hDC)))
}
func DestroyPbufferEXT(hPbuffer Pointer) int32 {
	return (int32)(C.gowglDestroyPbufferEXT((C.HPBUFFEREXT)(hPbuffer)))
}
func QueryPbufferEXT(hPbuffer Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.gowglQueryPbufferEXT((C.HPBUFFEREXT)(hPbuffer), (C.int)(iAttribute), (*C.int)(piValue)))
}
// EXT_pixel_format

func GetPixelFormatAttribivEXT(hdc Pointer, iPixelFormat Int, iLayerPlane Int, nAttributes uint32, piAttributes *Int, piValues *Int) int32 {
	return (int32)(C.gowglGetPixelFormatAttribivEXT((C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(piAttributes), (*C.int)(piValues)))
}
func GetPixelFormatAttribfvEXT(hdc Pointer, iPixelFormat Int, iLayerPlane Int, nAttributes uint32, piAttributes *Int, pfValues *float32) int32 {
	return (int32)(C.gowglGetPixelFormatAttribfvEXT((C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(piAttributes), (*C.FLOAT)(pfValues)))
}
func ChoosePixelFormatEXT(hdc Pointer, piAttribIList *Int, pfAttribFList *float32, nMaxFormats uint32, piFormats *Int, nNumFormats *uint32) int32 {
	return (int32)(C.gowglChoosePixelFormatEXT((C.HDC)(hdc), (*C.int)(piAttribIList), (*C.FLOAT)(pfAttribFList), (C.UINT)(nMaxFormats), (*C.int)(piFormats), (*C.UINT)(nNumFormats)))
}
// EXT_swap_control

func SwapIntervalEXT(interval Int) int32 {
	return (int32)(C.gowglSwapIntervalEXT((C.int)(interval)))
}
func GetSwapIntervalEXT() Int {
	return (Int)(C.gowglGetSwapIntervalEXT())
}
// I3D_digital_video_control

func GetDigitalVideoParametersI3D(hDC Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.gowglGetDigitalVideoParametersI3D((C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(piValue)))
}
func SetDigitalVideoParametersI3D(hDC Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.gowglSetDigitalVideoParametersI3D((C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(piValue)))
}
// I3D_gamma

func GetGammaTableParametersI3D(hDC Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.gowglGetGammaTableParametersI3D((C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(piValue)))
}
func SetGammaTableParametersI3D(hDC Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.gowglSetGammaTableParametersI3D((C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(piValue)))
}
func GetGammaTableI3D(hDC Pointer, iEntries Int, puRed *uint16, puGreen *uint16, puBlue *uint16) int32 {
	return (int32)(C.gowglGetGammaTableI3D((C.HDC)(hDC), (C.int)(iEntries), (*C.USHORT)(puRed), (*C.USHORT)(puGreen), (*C.USHORT)(puBlue)))
}
func SetGammaTableI3D(hDC Pointer, iEntries Int, puRed *uint16, puGreen *uint16, puBlue *uint16) int32 {
	return (int32)(C.gowglSetGammaTableI3D((C.HDC)(hDC), (C.int)(iEntries), (*C.USHORT)(puRed), (*C.USHORT)(puGreen), (*C.USHORT)(puBlue)))
}
// I3D_genlock

func EnableGenlockI3D(hDC Pointer) int32 {
	return (int32)(C.gowglEnableGenlockI3D((C.HDC)(hDC)))
}
func DisableGenlockI3D(hDC Pointer) int32 {
	return (int32)(C.gowglDisableGenlockI3D((C.HDC)(hDC)))
}
func IsEnabledGenlockI3D(hDC Pointer, pFlag *int32) int32 {
	return (int32)(C.gowglIsEnabledGenlockI3D((C.HDC)(hDC), (*C.BOOL)(pFlag)))
}
func GenlockSourceI3D(hDC Pointer, uSource uint32) int32 {
	return (int32)(C.gowglGenlockSourceI3D((C.HDC)(hDC), (C.UINT)(uSource)))
}
func GetGenlockSourceI3D(hDC Pointer, uSource *uint32) int32 {
	return (int32)(C.gowglGetGenlockSourceI3D((C.HDC)(hDC), (*C.UINT)(uSource)))
}
func GenlockSourceEdgeI3D(hDC Pointer, uEdge uint32) int32 {
	return (int32)(C.gowglGenlockSourceEdgeI3D((C.HDC)(hDC), (C.UINT)(uEdge)))
}
func GetGenlockSourceEdgeI3D(hDC Pointer, uEdge *uint32) int32 {
	return (int32)(C.gowglGetGenlockSourceEdgeI3D((C.HDC)(hDC), (*C.UINT)(uEdge)))
}
func GenlockSampleRateI3D(hDC Pointer, uRate uint32) int32 {
	return (int32)(C.gowglGenlockSampleRateI3D((C.HDC)(hDC), (C.UINT)(uRate)))
}
func GetGenlockSampleRateI3D(hDC Pointer, uRate *uint32) int32 {
	return (int32)(C.gowglGetGenlockSampleRateI3D((C.HDC)(hDC), (*C.UINT)(uRate)))
}
func GenlockSourceDelayI3D(hDC Pointer, uDelay uint32) int32 {
	return (int32)(C.gowglGenlockSourceDelayI3D((C.HDC)(hDC), (C.UINT)(uDelay)))
}
func GetGenlockSourceDelayI3D(hDC Pointer, uDelay *uint32) int32 {
	return (int32)(C.gowglGetGenlockSourceDelayI3D((C.HDC)(hDC), (*C.UINT)(uDelay)))
}
func QueryGenlockMaxSourceDelayI3D(hDC Pointer, uMaxLineDelay *uint32, uMaxPixelDelay *uint32) int32 {
	return (int32)(C.gowglQueryGenlockMaxSourceDelayI3D((C.HDC)(hDC), (*C.UINT)(uMaxLineDelay), (*C.UINT)(uMaxPixelDelay)))
}
// I3D_image_buffer

func CreateImageBufferI3D(hDC Pointer, dwSize uint32, uFlags uint32) Pointer {
	return (Pointer)(C.gowglCreateImageBufferI3D((C.HDC)(hDC), (C.DWORD)(dwSize), (C.UINT)(uFlags)))
}
func DestroyImageBufferI3D(hDC Pointer, pAddress Pointer) int32 {
	return (int32)(C.gowglDestroyImageBufferI3D((C.HDC)(hDC), (C.LPVOID)(pAddress)))
}
func AssociateImageBufferEventsI3D(hDC Pointer, pEvent Pointer, pAddress Pointer, pSize *uint32, count uint32) int32 {
	return (int32)(C.gowglAssociateImageBufferEventsI3D((C.HDC)(hDC), (C.HANDLE)(pEvent), (C.LPVOID)(pAddress), (*C.DWORD)(pSize), (C.UINT)(count)))
}
func ReleaseImageBufferEventsI3D(hDC Pointer, pAddress Pointer, count uint32) int32 {
	return (int32)(C.gowglReleaseImageBufferEventsI3D((C.HDC)(hDC), (C.LPVOID)(pAddress), (C.UINT)(count)))
}
// I3D_swap_frame_lock

func EnableFrameLockI3D() int32 {
	return (int32)(C.gowglEnableFrameLockI3D())
}
func DisableFrameLockI3D() int32 {
	return (int32)(C.gowglDisableFrameLockI3D())
}
func IsEnabledFrameLockI3D(pFlag *int32) int32 {
	return (int32)(C.gowglIsEnabledFrameLockI3D((*C.BOOL)(pFlag)))
}
func QueryFrameLockMasterI3D(pFlag *int32) int32 {
	return (int32)(C.gowglQueryFrameLockMasterI3D((*C.BOOL)(pFlag)))
}
// I3D_swap_frame_usage

func GetFrameUsageI3D(pUsage *float32) int32 {
	return (int32)(C.gowglGetFrameUsageI3D((*C.float32)(pUsage)))
}
func BeginFrameTrackingI3D() int32 {
	return (int32)(C.gowglBeginFrameTrackingI3D())
}
func EndFrameTrackingI3D() int32 {
	return (int32)(C.gowglEndFrameTrackingI3D())
}
func QueryFrameTrackingI3D(pFrameCount *uint32, pMissedFrames *uint32, pLastMissedUsage *float32) int32 {
	return (int32)(C.gowglQueryFrameTrackingI3D((*C.DWORD)(pFrameCount), (*C.DWORD)(pMissedFrames), (*C.float32)(pLastMissedUsage)))
}
// NV_DX_interop

func DXSetResourceShareHandleNV(dxObject Pointer, shareHandle Pointer) int32 {
	return (int32)(C.gowglDXSetResourceShareHandleNV((unsafe.Pointer)(dxObject), (C.HANDLE)(shareHandle)))
}
func DXOpenDeviceNV(dxDevice Pointer) Pointer {
	return (Pointer)(C.gowglDXOpenDeviceNV((unsafe.Pointer)(dxDevice)))
}
func DXCloseDeviceNV(hDevice Pointer) int32 {
	return (int32)(C.gowglDXCloseDeviceNV((C.HANDLE)(hDevice)))
}
func DXRegisterObjectNV(hDevice Pointer, dxObject Pointer, name Uint, type_ Enum, access Enum) Pointer {
	return (Pointer)(C.gowglDXRegisterObjectNV((C.HANDLE)(hDevice), (unsafe.Pointer)(dxObject), (C.GLuint)(name), (C.GLenum)(type_), (C.GLenum)(access)))
}
func DXUnregisterObjectNV(hDevice Pointer, hObject Pointer) int32 {
	return (int32)(C.gowglDXUnregisterObjectNV((C.HANDLE)(hDevice), (C.HANDLE)(hObject)))
}
func DXObjectAccessNV(hObject Pointer, access Enum) int32 {
	return (int32)(C.gowglDXObjectAccessNV((C.HANDLE)(hObject), (C.GLenum)(access)))
}
func DXLockObjectsNV(hDevice Pointer, count Int, hObjects Pointer) int32 {
	return (int32)(C.gowglDXLockObjectsNV((C.HANDLE)(hDevice), (C.GLint)(count), (C.HANDLE)(hObjects)))
}
func DXUnlockObjectsNV(hDevice Pointer, count Int, hObjects Pointer) int32 {
	return (int32)(C.gowglDXUnlockObjectsNV((C.HANDLE)(hDevice), (C.GLint)(count), (C.HANDLE)(hObjects)))
}
// NV_copy_image

func CopyImageSubDataNV(hSrcRC Pointer, srcName Uint, srcTarget Enum, srcLevel Int, srcX Int, srcY Int, srcZ Int, hDstRC Pointer, dstName Uint, dstTarget Enum, dstLevel Int, dstX Int, dstY Int, dstZ Int, width Sizei, height Sizei, depth Sizei) int32 {
	return (int32)(C.gowglCopyImageSubDataNV((C.HGLRC)(hSrcRC), (C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.HGLRC)(hDstRC), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth)))
}
// NV_gpu_affinity

func EnumGpusNV(iGpuIndex uint32, phGpu Pointer) int32 {
	return (int32)(C.gowglEnumGpusNV((C.UINT)(iGpuIndex), (C.HGPUNV)(phGpu)))
}
func EnumGpuDevicesNV(hGpu Pointer, iDeviceIndex uint32, lpGpuDevice Pointer) int32 {
	return (int32)(C.gowglEnumGpuDevicesNV((C.HGPUNV)(hGpu), (C.UINT)(iDeviceIndex), (C.PGPU_DEVICE)(lpGpuDevice)))
}
func CreateAffinityDCNV(phGpuList Pointer) Pointer {
	return (Pointer)(C.gowglCreateAffinityDCNV((C.HGPUNV)(phGpuList)))
}
func EnumGpusFromAffinityDCNV(hAffinityDC Pointer, iGpuIndex uint32, hGpu Pointer) int32 {
	return (int32)(C.gowglEnumGpusFromAffinityDCNV((C.HDC)(hAffinityDC), (C.UINT)(iGpuIndex), (C.HGPUNV)(hGpu)))
}
func DeleteDCNV(hdc Pointer) int32 {
	return (int32)(C.gowglDeleteDCNV((C.HDC)(hdc)))
}
// NV_present_video

func EnumerateVideoDevicesNV(hDC Pointer, phDeviceList Pointer) Int {
	return (Int)(C.gowglEnumerateVideoDevicesNV((C.HDC)(hDC), (C.HVIDEOOUTPUTDEVICENV)(phDeviceList)))
}
func BindVideoDeviceNV(hDC Pointer, uVideoSlot uint32, hVideoDevice Pointer, piAttribList *Int) int32 {
	return (int32)(C.gowglBindVideoDeviceNV((C.HDC)(hDC), (C.uint32)(uVideoSlot), (C.HVIDEOOUTPUTDEVICENV)(hVideoDevice), (*C.int)(piAttribList)))
}
func QueryCurrentContextNV(iAttribute Int, piValue *Int) int32 {
	return (int32)(C.gowglQueryCurrentContextNV((C.int)(iAttribute), (*C.int)(piValue)))
}
// NV_swap_group

func JoinSwapGroupNV(hDC Pointer, group Uint) int32 {
	return (int32)(C.gowglJoinSwapGroupNV((C.HDC)(hDC), (C.GLuint)(group)))
}
func BindSwapBarrierNV(group Uint, barrier Uint) int32 {
	return (int32)(C.gowglBindSwapBarrierNV((C.GLuint)(group), (C.GLuint)(barrier)))
}
func QuerySwapGroupNV(hDC Pointer, group *Uint, barrier *Uint) int32 {
	return (int32)(C.gowglQuerySwapGroupNV((C.HDC)(hDC), (*C.GLuint)(group), (*C.GLuint)(barrier)))
}
func QueryMaxSwapGroupsNV(hDC Pointer, maxGroups *Uint, maxBarriers *Uint) int32 {
	return (int32)(C.gowglQueryMaxSwapGroupsNV((C.HDC)(hDC), (*C.GLuint)(maxGroups), (*C.GLuint)(maxBarriers)))
}
func QueryFrameCountNV(hDC Pointer, count *Uint) int32 {
	return (int32)(C.gowglQueryFrameCountNV((C.HDC)(hDC), (*C.GLuint)(count)))
}
func ResetFrameCountNV(hDC Pointer) int32 {
	return (int32)(C.gowglResetFrameCountNV((C.HDC)(hDC)))
}
// NV_vertex_array_range

func AllocateMemoryNV(size Sizei, readfreq Float, writefreq Float, priority Float) Pointer {
	return (Pointer)(C.gowglAllocateMemoryNV((C.GLsizei)(size), (C.GLfloat)(readfreq), (C.GLfloat)(writefreq), (C.GLfloat)(priority)))
}
func FreeMemoryNV(pointer Pointer)  {
	C.gowglFreeMemoryNV((unsafe.Pointer)(pointer))
}
// NV_video_capture

func BindVideoCaptureDeviceNV(uVideoSlot uint32, hDevice Pointer) int32 {
	return (int32)(C.gowglBindVideoCaptureDeviceNV((C.UINT)(uVideoSlot), (C.HVIDEOINPUTDEVICENV)(hDevice)))
}
func EnumerateVideoCaptureDevicesNV(hDc Pointer, phDeviceList Pointer) uint32 {
	return (uint32)(C.gowglEnumerateVideoCaptureDevicesNV((C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(phDeviceList)))
}
func LockVideoCaptureDeviceNV(hDc Pointer, hDevice Pointer) int32 {
	return (int32)(C.gowglLockVideoCaptureDeviceNV((C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(hDevice)))
}
func QueryVideoCaptureDeviceNV(hDc Pointer, hDevice Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.gowglQueryVideoCaptureDeviceNV((C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(hDevice), (C.int)(iAttribute), (*C.int)(piValue)))
}
func ReleaseVideoCaptureDeviceNV(hDc Pointer, hDevice Pointer) int32 {
	return (int32)(C.gowglReleaseVideoCaptureDeviceNV((C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(hDevice)))
}
// NV_video_output

func GetVideoDeviceNV(hDC Pointer, numDevices Int, hVideoDevice Pointer) int32 {
	return (int32)(C.gowglGetVideoDeviceNV((C.HDC)(hDC), (C.int)(numDevices), (C.HPVIDEODEV)(hVideoDevice)))
}
func ReleaseVideoDeviceNV(hVideoDevice Pointer) int32 {
	return (int32)(C.gowglReleaseVideoDeviceNV((C.HPVIDEODEV)(hVideoDevice)))
}
func BindVideoImageNV(hVideoDevice Pointer, hPbuffer Pointer, iVideoBuffer Int) int32 {
	return (int32)(C.gowglBindVideoImageNV((C.HPVIDEODEV)(hVideoDevice), (C.HPBUFFERARB)(hPbuffer), (C.int)(iVideoBuffer)))
}
func ReleaseVideoImageNV(hPbuffer Pointer, iVideoBuffer Int) int32 {
	return (int32)(C.gowglReleaseVideoImageNV((C.HPBUFFERARB)(hPbuffer), (C.int)(iVideoBuffer)))
}
func SendPbufferToVideoNV(hPbuffer Pointer, iBufferType Int, pulCounterPbuffer *uint32, bBlock int32) int32 {
	return (int32)(C.gowglSendPbufferToVideoNV((C.HPBUFFERARB)(hPbuffer), (C.int)(iBufferType), (*C.unsigned_long)(pulCounterPbuffer), (C.BOOL)(bBlock)))
}
func GetVideoInfoNV(hpVideoDevice Pointer, pulCounterOutputPbuffer *uint32, pulCounterOutputVideo *uint32) int32 {
	return (int32)(C.gowglGetVideoInfoNV((C.HPVIDEODEV)(hpVideoDevice), (*C.unsigned_long)(pulCounterOutputPbuffer), (*C.unsigned_long)(pulCounterOutputVideo)))
}
// OML_sync_control

func GetSyncValuesOML(hdc Pointer, ust *int64, msc *int64, sbc *int64) int32 {
	return (int32)(C.gowglGetSyncValuesOML((C.HDC)(hdc), (*C.INT64)(ust), (*C.INT64)(msc), (*C.INT64)(sbc)))
}
func GetMscRateOML(hdc Pointer, numerator *int32, denominator *int32) int32 {
	return (int32)(C.gowglGetMscRateOML((C.HDC)(hdc), (*C.INT32)(numerator), (*C.INT32)(denominator)))
}
func SwapBuffersMscOML(hdc Pointer, target_msc int64, divisor int64, remainder int64) int64 {
	return (int64)(C.gowglSwapBuffersMscOML((C.HDC)(hdc), (C.INT64)(target_msc), (C.INT64)(divisor), (C.INT64)(remainder)))
}
func SwapLayerBuffersMscOML(hdc Pointer, fuPlanes Int, target_msc int64, divisor int64, remainder int64) int64 {
	return (int64)(C.gowglSwapLayerBuffersMscOML((C.HDC)(hdc), (C.int)(fuPlanes), (C.INT64)(target_msc), (C.INT64)(divisor), (C.INT64)(remainder)))
}
func WaitForMscOML(hdc Pointer, target_msc int64, divisor int64, remainder int64, ust *int64, msc *int64, sbc *int64) int32 {
	return (int32)(C.gowglWaitForMscOML((C.HDC)(hdc), (C.INT64)(target_msc), (C.INT64)(divisor), (C.INT64)(remainder), (*C.INT64)(ust), (*C.INT64)(msc), (*C.INT64)(sbc)))
}
func WaitForSbcOML(hdc Pointer, target_sbc int64, ust *int64, msc *int64, sbc *int64) int32 {
	return (int32)(C.gowglWaitForSbcOML((C.HDC)(hdc), (C.INT64)(target_sbc), (*C.INT64)(ust), (*C.INT64)(msc), (*C.INT64)(sbc)))
}
// wgl

func CreateContext(hDc Pointer) Pointer {
	return (Pointer)(C.gowglCreateContext((C.HDC)(hDc)))
}
func DeleteContext(oldContext Pointer) int32 {
	return (int32)(C.gowglDeleteContext((C.HGLRC)(oldContext)))
}
func GetCurrentContext() Pointer {
	return (Pointer)(C.gowglGetCurrentContext())
}
func MakeCurrent(hDc Pointer, newContext Pointer) int32 {
	return (int32)(C.gowglMakeCurrent((C.HDC)(hDc), (C.HGLRC)(newContext)))
}
func CopyContext(hglrcSrc Pointer, hglrcDst Pointer, mask uint32) int32 {
	return (int32)(C.gowglCopyContext((C.HGLRC)(hglrcSrc), (C.HGLRC)(hglrcDst), (C.UINT)(mask)))
}
func ChoosePixelFormat(hDc Pointer, pPfd []byte) Int {
	return (Int)(C.gowglChoosePixelFormat((C.HDC)(hDc), (C.PIXELFORMATDESCRIPTOR)(pPfd)))
}
func GetCurrentDC() Pointer {
	return (Pointer)(C.gowglGetCurrentDC())
}
func GetDefaultProcAddress(lpszProc *byte) Pointer {
	return (Pointer)(C.gowglGetDefaultProcAddress((C.LPCSTR)(lpszProc)))
}
func GetProcAddress(lpszProc *byte) Pointer {
	return (Pointer)(C.gowglGetProcAddress((C.LPCSTR)(lpszProc)))
}
func GetPixelFormat(hdc Pointer) Int {
	return (Int)(C.gowglGetPixelFormat((C.HDC)(hdc)))
}
func SetPixelFormat(hdc Pointer, ipfd Int, ppfd []byte) int32 {
	return (int32)(C.gowglSetPixelFormat((C.HDC)(hdc), (C.int)(ipfd), (C.PIXELFORMATDESCRIPTOR)(ppfd)))
}
func SwapBuffers(hdc Pointer) int32 {
	return (int32)(C.gowglSwapBuffers((C.HDC)(hdc)))
}
func ShareLists(hrcSrvShare Pointer, hrcSrvSource Pointer) int32 {
	return (int32)(C.gowglShareLists((C.HGLRC)(hrcSrvShare), (C.HGLRC)(hrcSrvSource)))
}
func CreateLayerContext(hDc Pointer, level Int) Pointer {
	return (Pointer)(C.gowglCreateLayerContext((C.HDC)(hDc), (C.int)(level)))
}
func DescribeLayerPlane(hDc Pointer, pixelFormat Int, layerPlane Int, nBytes uint32, plpd []byte) int32 {
	return (int32)(C.gowglDescribeLayerPlane((C.HDC)(hDc), (C.int)(pixelFormat), (C.int)(layerPlane), (C.UINT)(nBytes), (C.LAYERPLANEDESCRIPTOR)(plpd)))
}
func SetLayerPaletteEntries(hdc Pointer, iLayerPlane Int, iStart Int, cEntries Int, pcr Pointer) Int {
	return (Int)(C.gowglSetLayerPaletteEntries((C.HDC)(hdc), (C.int)(iLayerPlane), (C.int)(iStart), (C.int)(cEntries), (C.COLORREF)(pcr)))
}
func GetLayerPaletteEntries(hdc Pointer, iLayerPlane Int, iStart Int, cEntries Int, pcr Pointer) Int {
	return (Int)(C.gowglGetLayerPaletteEntries((C.HDC)(hdc), (C.int)(iLayerPlane), (C.int)(iStart), (C.int)(cEntries), (C.COLORREF)(pcr)))
}
func RealizeLayerPalette(hdc Pointer, iLayerPlane Int, bRealize int32) int32 {
	return (int32)(C.gowglRealizeLayerPalette((C.HDC)(hdc), (C.int)(iLayerPlane), (C.BOOL)(bRealize)))
}
func SwapLayerBuffers(hdc Pointer, fuFlags uint32) int32 {
	return (int32)(C.gowglSwapLayerBuffers((C.HDC)(hdc), (C.UINT)(fuFlags)))
}
func UseFontBitmapsA(hDC Pointer, first uint32, count uint32, listBase uint32) int32 {
	return (int32)(C.gowglUseFontBitmapsA((C.HDC)(hDC), (C.DWORD)(first), (C.DWORD)(count), (C.DWORD)(listBase)))
}
func UseFontBitmapsW(hDC Pointer, first uint32, count uint32, listBase uint32) int32 {
	return (int32)(C.gowglUseFontBitmapsW((C.HDC)(hDC), (C.DWORD)(first), (C.DWORD)(count), (C.DWORD)(listBase)))
}
func Init3dlStereoControl() error {
	var ret C.int
	if ret = C.init_3DL_stereo_control(); ret != 0 {
		return errors.New("unable to initialize 3DL_stereo_control")
	}
	return nil
}
func InitAmdGpuAssociation() error {
	var ret C.int
	if ret = C.init_AMD_gpu_association(); ret != 0 {
		return errors.New("unable to initialize AMD_gpu_association")
	}
	return nil
}
func InitArbBufferRegion() error {
	var ret C.int
	if ret = C.init_ARB_buffer_region(); ret != 0 {
		return errors.New("unable to initialize ARB_buffer_region")
	}
	return nil
}
func InitArbCreateContext() error {
	var ret C.int
	if ret = C.init_ARB_create_context(); ret != 0 {
		return errors.New("unable to initialize ARB_create_context")
	}
	return nil
}
func InitArbExtensionsString() error {
	var ret C.int
	if ret = C.init_ARB_extensions_string(); ret != 0 {
		return errors.New("unable to initialize ARB_extensions_string")
	}
	return nil
}
func InitArbMakeCurrentRead() error {
	var ret C.int
	if ret = C.init_ARB_make_current_read(); ret != 0 {
		return errors.New("unable to initialize ARB_make_current_read")
	}
	return nil
}
func InitArbPbuffer() error {
	var ret C.int
	if ret = C.init_ARB_pbuffer(); ret != 0 {
		return errors.New("unable to initialize ARB_pbuffer")
	}
	return nil
}
func InitArbPixelFormat() error {
	var ret C.int
	if ret = C.init_ARB_pixel_format(); ret != 0 {
		return errors.New("unable to initialize ARB_pixel_format")
	}
	return nil
}
func InitArbRenderTexture() error {
	var ret C.int
	if ret = C.init_ARB_render_texture(); ret != 0 {
		return errors.New("unable to initialize ARB_render_texture")
	}
	return nil
}
func InitExtDisplayColorTable() error {
	var ret C.int
	if ret = C.init_EXT_display_color_table(); ret != 0 {
		return errors.New("unable to initialize EXT_display_color_table")
	}
	return nil
}
func InitExtExtensionsString() error {
	var ret C.int
	if ret = C.init_EXT_extensions_string(); ret != 0 {
		return errors.New("unable to initialize EXT_extensions_string")
	}
	return nil
}
func InitExtMakeCurrentRead() error {
	var ret C.int
	if ret = C.init_EXT_make_current_read(); ret != 0 {
		return errors.New("unable to initialize EXT_make_current_read")
	}
	return nil
}
func InitExtPbuffer() error {
	var ret C.int
	if ret = C.init_EXT_pbuffer(); ret != 0 {
		return errors.New("unable to initialize EXT_pbuffer")
	}
	return nil
}
func InitExtPixelFormat() error {
	var ret C.int
	if ret = C.init_EXT_pixel_format(); ret != 0 {
		return errors.New("unable to initialize EXT_pixel_format")
	}
	return nil
}
func InitExtSwapControl() error {
	var ret C.int
	if ret = C.init_EXT_swap_control(); ret != 0 {
		return errors.New("unable to initialize EXT_swap_control")
	}
	return nil
}
func InitI3dDigitalVideoControl() error {
	var ret C.int
	if ret = C.init_I3D_digital_video_control(); ret != 0 {
		return errors.New("unable to initialize I3D_digital_video_control")
	}
	return nil
}
func InitI3dGamma() error {
	var ret C.int
	if ret = C.init_I3D_gamma(); ret != 0 {
		return errors.New("unable to initialize I3D_gamma")
	}
	return nil
}
func InitI3dGenlock() error {
	var ret C.int
	if ret = C.init_I3D_genlock(); ret != 0 {
		return errors.New("unable to initialize I3D_genlock")
	}
	return nil
}
func InitI3dImageBuffer() error {
	var ret C.int
	if ret = C.init_I3D_image_buffer(); ret != 0 {
		return errors.New("unable to initialize I3D_image_buffer")
	}
	return nil
}
func InitI3dSwapFrameLock() error {
	var ret C.int
	if ret = C.init_I3D_swap_frame_lock(); ret != 0 {
		return errors.New("unable to initialize I3D_swap_frame_lock")
	}
	return nil
}
func InitI3dSwapFrameUsage() error {
	var ret C.int
	if ret = C.init_I3D_swap_frame_usage(); ret != 0 {
		return errors.New("unable to initialize I3D_swap_frame_usage")
	}
	return nil
}
func InitNvDxInterop() error {
	var ret C.int
	if ret = C.init_NV_DX_interop(); ret != 0 {
		return errors.New("unable to initialize NV_DX_interop")
	}
	return nil
}
func InitNvCopyImage() error {
	var ret C.int
	if ret = C.init_NV_copy_image(); ret != 0 {
		return errors.New("unable to initialize NV_copy_image")
	}
	return nil
}
func InitNvGpuAffinity() error {
	var ret C.int
	if ret = C.init_NV_gpu_affinity(); ret != 0 {
		return errors.New("unable to initialize NV_gpu_affinity")
	}
	return nil
}
func InitNvPresentVideo() error {
	var ret C.int
	if ret = C.init_NV_present_video(); ret != 0 {
		return errors.New("unable to initialize NV_present_video")
	}
	return nil
}
func InitNvSwapGroup() error {
	var ret C.int
	if ret = C.init_NV_swap_group(); ret != 0 {
		return errors.New("unable to initialize NV_swap_group")
	}
	return nil
}
func InitNvVertexArrayRange() error {
	var ret C.int
	if ret = C.init_NV_vertex_array_range(); ret != 0 {
		return errors.New("unable to initialize NV_vertex_array_range")
	}
	return nil
}
func InitNvVideoCapture() error {
	var ret C.int
	if ret = C.init_NV_video_capture(); ret != 0 {
		return errors.New("unable to initialize NV_video_capture")
	}
	return nil
}
func InitNvVideoOutput() error {
	var ret C.int
	if ret = C.init_NV_video_output(); ret != 0 {
		return errors.New("unable to initialize NV_video_output")
	}
	return nil
}
func InitOmlSyncControl() error {
	var ret C.int
	if ret = C.init_OML_sync_control(); ret != 0 {
		return errors.New("unable to initialize OML_sync_control")
	}
	return nil
}
func InitWgl() error {
	var ret C.int
	if ret = C.init_wgl(); ret != 0 {
		return errors.New("unable to initialize wgl")
	}
	return nil
}
// EOF