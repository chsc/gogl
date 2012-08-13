// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// 3DFX_multisample: http://www.opengl.org/registry/specs/3DFX/multisample.txt
// 
// 3DL_stereo_control: http://www.opengl.org/registry/specs/3DL/stereo_control.txt
// 
// AMD_gpu_association: http://www.opengl.org/registry/specs/AMD/gpu_association.txt
// 
// ARB_buffer_region: http://www.opengl.org/registry/specs/ARB/buffer_region.txt
// 
// ARB_create_context: http://www.opengl.org/registry/specs/ARB/create_context.txt
// 
// ARB_create_context_profile: http://www.opengl.org/registry/specs/ARB/create_context_profile.txt
// 
// ARB_create_context_robustness: http://www.opengl.org/registry/specs/ARB/create_context_robustness.txt
// 
// ARB_extensions_string: http://www.opengl.org/registry/specs/ARB/extensions_string.txt
// 
// ARB_framebuffer_sRGB: http://www.opengl.org/registry/specs/ARB/framebuffer_sRGB.txt
// 
// ARB_make_current_read: http://www.opengl.org/registry/specs/ARB/make_current_read.txt
// 
// ARB_multisample: http://www.opengl.org/registry/specs/ARB/multisample.txt
// 
// ARB_pbuffer: http://www.opengl.org/registry/specs/ARB/pbuffer.txt
// 
// ARB_pixel_format: http://www.opengl.org/registry/specs/ARB/pixel_format.txt
// 
// ARB_pixel_format_float: http://www.opengl.org/registry/specs/ARB/pixel_format_float.txt
// 
// ARB_render_texture: http://www.opengl.org/registry/specs/ARB/render_texture.txt
// 
// ATI_pixel_format_float: http://www.opengl.org/registry/specs/ATI/pixel_format_float.txt
// 
// EXT_depth_float: http://www.opengl.org/registry/specs/EXT/depth_float.txt
// 
// EXT_display_color_table: http://www.opengl.org/registry/specs/EXT/display_color_table.txt
// 
// EXT_extensions_string: http://www.opengl.org/registry/specs/EXT/extensions_string.txt
// 
// EXT_framebuffer_sRGB: http://www.opengl.org/registry/specs/EXT/framebuffer_sRGB.txt
// 
// EXT_make_current_read: http://www.opengl.org/registry/specs/EXT/make_current_read.txt
// 
// EXT_multisample: http://www.opengl.org/registry/specs/EXT/multisample.txt
// 
// EXT_pbuffer: http://www.opengl.org/registry/specs/EXT/pbuffer.txt
// 
// EXT_pixel_format: http://www.opengl.org/registry/specs/EXT/pixel_format.txt
// 
// EXT_pixel_format_packed_float: http://www.opengl.org/registry/specs/EXT/pixel_format_packed_float.txt
// 
// EXT_swap_control: http://www.opengl.org/registry/specs/EXT/swap_control.txt
// 
// EXT_swap_control_tear: http://www.opengl.org/registry/specs/EXT/swap_control_tear.txt
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
// NV_DX_interop2: http://www.opengl.org/registry/specs/NV/DX_interop2.txt
// 
// NV_copy_image: http://www.opengl.org/registry/specs/NV/copy_image.txt
// 
// NV_float_buffer: http://www.opengl.org/registry/specs/NV/float_buffer.txt
// 
// NV_gpu_affinity: http://www.opengl.org/registry/specs/NV/gpu_affinity.txt
// 
// NV_multisample_coverage: http://www.opengl.org/registry/specs/NV/multisample_coverage.txt
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
// static void* goglGetProcAddress(const char* name) { 
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
// //  3DFX_multisample
// //  3DL_stereo_control
// BOOL (APIENTRYP ptrglSetStereoEmitterState3DL)(HDC hDC, UINT uState);
// //  AMD_gpu_association
// UINT (APIENTRYP ptrglGetGPUIDsAMD)(UINT maxCount, UINT* ids);
// INT (APIENTRYP ptrglGetGPUInfoAMD)(UINT id, int property, GLenum dataType, UINT size, void* data);
// UINT (APIENTRYP ptrglGetContextGPUIDAMD)(HGLRC hglrc);
// HGLRC (APIENTRYP ptrglCreateAssociatedContextAMD)(UINT id);
// HGLRC (APIENTRYP ptrglCreateAssociatedContextAttribsAMD)(UINT id, HGLRC hShareContext, int* attribList);
// BOOL (APIENTRYP ptrglDeleteAssociatedContextAMD)(HGLRC hglrc);
// BOOL (APIENTRYP ptrglMakeAssociatedContextCurrentAMD)(HGLRC hglrc);
// HGLRC (APIENTRYP ptrglGetCurrentAssociatedContextAMD)();
// VOID (APIENTRYP ptrglBlitContextFramebufferAMD)(HGLRC dstCtx, GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
// //  ARB_buffer_region
// HANDLE (APIENTRYP ptrglCreateBufferRegionARB)(HDC hDC, int iLayerPlane, UINT uType);
// VOID (APIENTRYP ptrglDeleteBufferRegionARB)(HANDLE hRegion);
// BOOL (APIENTRYP ptrglSaveBufferRegionARB)(HANDLE hRegion, int x, int y, int width, int height);
// BOOL (APIENTRYP ptrglRestoreBufferRegionARB)(HANDLE hRegion, int x, int y, int width, int height, int xSrc, int ySrc);
// //  ARB_create_context
// HGLRC (APIENTRYP ptrglCreateContextAttribsARB)(HDC hDC, HGLRC hShareContext, int* attribList);
// //  ARB_create_context_profile
// //  ARB_create_context_robustness
// //  ARB_extensions_string
// const char * (APIENTRYP ptrglGetExtensionsStringARB)(HDC hdc);
// //  ARB_framebuffer_sRGB
// //  ARB_make_current_read
// BOOL (APIENTRYP ptrglMakeContextCurrentARB)(HDC hDrawDC, HDC hReadDC, HGLRC hglrc);
// HDC (APIENTRYP ptrglGetCurrentReadDCARB)();
// //  ARB_multisample
// //  ARB_pbuffer
// HPBUFFERARB (APIENTRYP ptrglCreatePbufferARB)(HDC hDC, int iPixelFormat, int iWidth, int iHeight, int* piAttribList);
// HDC (APIENTRYP ptrglGetPbufferDCARB)(HPBUFFERARB hPbuffer);
// int (APIENTRYP ptrglReleasePbufferDCARB)(HPBUFFERARB hPbuffer, HDC hDC);
// BOOL (APIENTRYP ptrglDestroyPbufferARB)(HPBUFFERARB hPbuffer);
// BOOL (APIENTRYP ptrglQueryPbufferARB)(HPBUFFERARB hPbuffer, int iAttribute, int* piValue);
// //  ARB_pixel_format
// BOOL (APIENTRYP ptrglGetPixelFormatAttribivARB)(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, int* piValues);
// BOOL (APIENTRYP ptrglGetPixelFormatAttribfvARB)(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, FLOAT* pfValues);
// BOOL (APIENTRYP ptrglChoosePixelFormatARB)(HDC hdc, int* piAttribIList, FLOAT* pfAttribFList, UINT nMaxFormats, int* piFormats, UINT* nNumFormats);
// //  ARB_pixel_format_float
// //  ARB_render_texture
// BOOL (APIENTRYP ptrglBindTexImageARB)(HPBUFFERARB hPbuffer, int iBuffer);
// BOOL (APIENTRYP ptrglReleaseTexImageARB)(HPBUFFERARB hPbuffer, int iBuffer);
// BOOL (APIENTRYP ptrglSetPbufferAttribARB)(HPBUFFERARB hPbuffer, int* piAttribList);
// //  ATI_pixel_format_float
// //  EXT_depth_float
// //  EXT_display_color_table
// GLboolean (APIENTRYP ptrglCreateDisplayColorTableEXT)(GLushort id);
// GLboolean (APIENTRYP ptrglLoadDisplayColorTableEXT)(GLushort* table, GLuint length);
// GLboolean (APIENTRYP ptrglBindDisplayColorTableEXT)(GLushort id);
// VOID (APIENTRYP ptrglDestroyDisplayColorTableEXT)(GLushort id);
// //  EXT_extensions_string
// const char * (APIENTRYP ptrglGetExtensionsStringEXT)();
// //  EXT_framebuffer_sRGB
// //  EXT_make_current_read
// BOOL (APIENTRYP ptrglMakeContextCurrentEXT)(HDC hDrawDC, HDC hReadDC, HGLRC hglrc);
// HDC (APIENTRYP ptrglGetCurrentReadDCEXT)();
// //  EXT_multisample
// //  EXT_pbuffer
// HPBUFFEREXT (APIENTRYP ptrglCreatePbufferEXT)(HDC hDC, int iPixelFormat, int iWidth, int iHeight, int* piAttribList);
// HDC (APIENTRYP ptrglGetPbufferDCEXT)(HPBUFFEREXT hPbuffer);
// int (APIENTRYP ptrglReleasePbufferDCEXT)(HPBUFFEREXT hPbuffer, HDC hDC);
// BOOL (APIENTRYP ptrglDestroyPbufferEXT)(HPBUFFEREXT hPbuffer);
// BOOL (APIENTRYP ptrglQueryPbufferEXT)(HPBUFFEREXT hPbuffer, int iAttribute, int* piValue);
// //  EXT_pixel_format
// BOOL (APIENTRYP ptrglGetPixelFormatAttribivEXT)(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, int* piValues);
// BOOL (APIENTRYP ptrglGetPixelFormatAttribfvEXT)(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, FLOAT* pfValues);
// BOOL (APIENTRYP ptrglChoosePixelFormatEXT)(HDC hdc, int* piAttribIList, FLOAT* pfAttribFList, UINT nMaxFormats, int* piFormats, UINT* nNumFormats);
// //  EXT_pixel_format_packed_float
// //  EXT_swap_control
// BOOL (APIENTRYP ptrglSwapIntervalEXT)(int interval);
// int (APIENTRYP ptrglGetSwapIntervalEXT)();
// //  EXT_swap_control_tear
// //  I3D_digital_video_control
// BOOL (APIENTRYP ptrglGetDigitalVideoParametersI3D)(HDC hDC, int iAttribute, int* piValue);
// BOOL (APIENTRYP ptrglSetDigitalVideoParametersI3D)(HDC hDC, int iAttribute, int* piValue);
// //  I3D_gamma
// BOOL (APIENTRYP ptrglGetGammaTableParametersI3D)(HDC hDC, int iAttribute, int* piValue);
// BOOL (APIENTRYP ptrglSetGammaTableParametersI3D)(HDC hDC, int iAttribute, int* piValue);
// BOOL (APIENTRYP ptrglGetGammaTableI3D)(HDC hDC, int iEntries, USHORT* puRed, USHORT* puGreen, USHORT* puBlue);
// BOOL (APIENTRYP ptrglSetGammaTableI3D)(HDC hDC, int iEntries, USHORT* puRed, USHORT* puGreen, USHORT* puBlue);
// //  I3D_genlock
// BOOL (APIENTRYP ptrglEnableGenlockI3D)(HDC hDC);
// BOOL (APIENTRYP ptrglDisableGenlockI3D)(HDC hDC);
// BOOL (APIENTRYP ptrglIsEnabledGenlockI3D)(HDC hDC, BOOL* pFlag);
// BOOL (APIENTRYP ptrglGenlockSourceI3D)(HDC hDC, UINT uSource);
// BOOL (APIENTRYP ptrglGetGenlockSourceI3D)(HDC hDC, UINT* uSource);
// BOOL (APIENTRYP ptrglGenlockSourceEdgeI3D)(HDC hDC, UINT uEdge);
// BOOL (APIENTRYP ptrglGetGenlockSourceEdgeI3D)(HDC hDC, UINT* uEdge);
// BOOL (APIENTRYP ptrglGenlockSampleRateI3D)(HDC hDC, UINT uRate);
// BOOL (APIENTRYP ptrglGetGenlockSampleRateI3D)(HDC hDC, UINT* uRate);
// BOOL (APIENTRYP ptrglGenlockSourceDelayI3D)(HDC hDC, UINT uDelay);
// BOOL (APIENTRYP ptrglGetGenlockSourceDelayI3D)(HDC hDC, UINT* uDelay);
// BOOL (APIENTRYP ptrglQueryGenlockMaxSourceDelayI3D)(HDC hDC, UINT* uMaxLineDelay, UINT* uMaxPixelDelay);
// //  I3D_image_buffer
// LPVOID (APIENTRYP ptrglCreateImageBufferI3D)(HDC hDC, DWORD dwSize, UINT uFlags);
// BOOL (APIENTRYP ptrglDestroyImageBufferI3D)(HDC hDC, LPVOID pAddress);
// BOOL (APIENTRYP ptrglAssociateImageBufferEventsI3D)(HDC hDC, HANDLE* pEvent, LPVOID* pAddress, DWORD* pSize, UINT count);
// BOOL (APIENTRYP ptrglReleaseImageBufferEventsI3D)(HDC hDC, LPVOID* pAddress, UINT count);
// //  I3D_swap_frame_lock
// BOOL (APIENTRYP ptrglEnableFrameLockI3D)();
// BOOL (APIENTRYP ptrglDisableFrameLockI3D)();
// BOOL (APIENTRYP ptrglIsEnabledFrameLockI3D)(BOOL* pFlag);
// BOOL (APIENTRYP ptrglQueryFrameLockMasterI3D)(BOOL* pFlag);
// //  I3D_swap_frame_usage
// BOOL (APIENTRYP ptrglGetFrameUsageI3D)(float* pUsage);
// BOOL (APIENTRYP ptrglBeginFrameTrackingI3D)();
// BOOL (APIENTRYP ptrglEndFrameTrackingI3D)();
// BOOL (APIENTRYP ptrglQueryFrameTrackingI3D)(DWORD* pFrameCount, DWORD* pMissedFrames, float* pLastMissedUsage);
// //  NV_DX_interop
// BOOL (APIENTRYP ptrglDXSetResourceShareHandleNV)(void* dxObject, HANDLE shareHandle);
// HANDLE (APIENTRYP ptrglDXOpenDeviceNV)(void* dxDevice);
// BOOL (APIENTRYP ptrglDXCloseDeviceNV)(HANDLE hDevice);
// HANDLE (APIENTRYP ptrglDXRegisterObjectNV)(HANDLE hDevice, void* dxObject, GLuint name, GLenum type, GLenum access);
// BOOL (APIENTRYP ptrglDXUnregisterObjectNV)(HANDLE hDevice, HANDLE hObject);
// BOOL (APIENTRYP ptrglDXObjectAccessNV)(HANDLE hObject, GLenum access);
// BOOL (APIENTRYP ptrglDXLockObjectsNV)(HANDLE hDevice, GLint count, HANDLE* hObjects);
// BOOL (APIENTRYP ptrglDXUnlockObjectsNV)(HANDLE hDevice, GLint count, HANDLE* hObjects);
// //  NV_DX_interop2
// //  NV_copy_image
// BOOL (APIENTRYP ptrglCopyImageSubDataNV)(HGLRC hSrcRC, GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, HGLRC hDstRC, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei width, GLsizei height, GLsizei depth);
// //  NV_float_buffer
// //  NV_gpu_affinity
// BOOL (APIENTRYP ptrglEnumGpusNV)(UINT iGpuIndex, HGPUNV* phGpu);
// BOOL (APIENTRYP ptrglEnumGpuDevicesNV)(HGPUNV hGpu, UINT iDeviceIndex, PGPU_DEVICE lpGpuDevice);
// HDC (APIENTRYP ptrglCreateAffinityDCNV)(HGPUNV* phGpuList);
// BOOL (APIENTRYP ptrglEnumGpusFromAffinityDCNV)(HDC hAffinityDC, UINT iGpuIndex, HGPUNV* hGpu);
// BOOL (APIENTRYP ptrglDeleteDCNV)(HDC hdc);
// //  NV_multisample_coverage
// //  NV_present_video
// int (APIENTRYP ptrglEnumerateVideoDevicesNV)(HDC hDC, HVIDEOOUTPUTDEVICENV* phDeviceList);
// BOOL (APIENTRYP ptrglBindVideoDeviceNV)(HDC hDC, unsigned int uVideoSlot, HVIDEOOUTPUTDEVICENV hVideoDevice, int* piAttribList);
// BOOL (APIENTRYP ptrglQueryCurrentContextNV)(int iAttribute, int* piValue);
// //  NV_swap_group
// BOOL (APIENTRYP ptrglJoinSwapGroupNV)(HDC hDC, GLuint group);
// BOOL (APIENTRYP ptrglBindSwapBarrierNV)(GLuint group, GLuint barrier);
// BOOL (APIENTRYP ptrglQuerySwapGroupNV)(HDC hDC, GLuint* group, GLuint* barrier);
// BOOL (APIENTRYP ptrglQueryMaxSwapGroupsNV)(HDC hDC, GLuint* maxGroups, GLuint* maxBarriers);
// BOOL (APIENTRYP ptrglQueryFrameCountNV)(HDC hDC, GLuint* count);
// BOOL (APIENTRYP ptrglResetFrameCountNV)(HDC hDC);
// //  NV_vertex_array_range
// void* (APIENTRYP ptrglAllocateMemoryNV)(GLsizei size, GLfloat readfreq, GLfloat writefreq, GLfloat priority);
// void (APIENTRYP ptrglFreeMemoryNV)(void* pointer);
// //  NV_video_capture
// BOOL (APIENTRYP ptrglBindVideoCaptureDeviceNV)(UINT uVideoSlot, HVIDEOINPUTDEVICENV hDevice);
// UINT (APIENTRYP ptrglEnumerateVideoCaptureDevicesNV)(HDC hDc, HVIDEOINPUTDEVICENV* phDeviceList);
// BOOL (APIENTRYP ptrglLockVideoCaptureDeviceNV)(HDC hDc, HVIDEOINPUTDEVICENV hDevice);
// BOOL (APIENTRYP ptrglQueryVideoCaptureDeviceNV)(HDC hDc, HVIDEOINPUTDEVICENV hDevice, int iAttribute, int* piValue);
// BOOL (APIENTRYP ptrglReleaseVideoCaptureDeviceNV)(HDC hDc, HVIDEOINPUTDEVICENV hDevice);
// //  NV_video_output
// BOOL (APIENTRYP ptrglGetVideoDeviceNV)(HDC hDC, int numDevices, HPVIDEODEV* hVideoDevice);
// BOOL (APIENTRYP ptrglReleaseVideoDeviceNV)(HPVIDEODEV hVideoDevice);
// BOOL (APIENTRYP ptrglBindVideoImageNV)(HPVIDEODEV hVideoDevice, HPBUFFERARB hPbuffer, int iVideoBuffer);
// BOOL (APIENTRYP ptrglReleaseVideoImageNV)(HPBUFFERARB hPbuffer, int iVideoBuffer);
// BOOL (APIENTRYP ptrglSendPbufferToVideoNV)(HPBUFFERARB hPbuffer, int iBufferType, unsigned long* pulCounterPbuffer, BOOL bBlock);
// BOOL (APIENTRYP ptrglGetVideoInfoNV)(HPVIDEODEV hpVideoDevice, unsigned long* pulCounterOutputPbuffer, unsigned long* pulCounterOutputVideo);
// //  OML_sync_control
// BOOL (APIENTRYP ptrglGetSyncValuesOML)(HDC hdc, INT64* ust, INT64* msc, INT64* sbc);
// BOOL (APIENTRYP ptrglGetMscRateOML)(HDC hdc, INT32* numerator, INT32* denominator);
// INT64 (APIENTRYP ptrglSwapBuffersMscOML)(HDC hdc, INT64 target_msc, INT64 divisor, INT64 remainder);
// INT64 (APIENTRYP ptrglSwapLayerBuffersMscOML)(HDC hdc, int fuPlanes, INT64 target_msc, INT64 divisor, INT64 remainder);
// BOOL (APIENTRYP ptrglWaitForMscOML)(HDC hdc, INT64 target_msc, INT64 divisor, INT64 remainder, INT64* ust, INT64* msc, INT64* sbc);
// BOOL (APIENTRYP ptrglWaitForSbcOML)(HDC hdc, INT64 target_sbc, INT64* ust, INT64* msc, INT64* sbc);
// //  wgl
// HGLRC (APIENTRYP ptrglCreateContext)(HDC hDc);
// BOOL (APIENTRYP ptrglDeleteContext)(HGLRC oldContext);
// HGLRC (APIENTRYP ptrglGetCurrentContext)();
// BOOL (APIENTRYP ptrglMakeCurrent)(HDC hDc, HGLRC newContext);
// BOOL (APIENTRYP ptrglCopyContext)(HGLRC hglrcSrc, HGLRC hglrcDst, UINT mask);
// int (APIENTRYP ptrglChoosePixelFormat)(HDC hDc, PIXELFORMATDESCRIPTOR* pPfd);
// HDC (APIENTRYP ptrglGetCurrentDC)();
// PROC (APIENTRYP ptrglGetDefaultProcAddress)(LPCSTR lpszProc);
// PROC (APIENTRYP ptrglGetProcAddress)(LPCSTR lpszProc);
// int (APIENTRYP ptrglGetPixelFormat)(HDC hdc);
// BOOL (APIENTRYP ptrglSetPixelFormat)(HDC hdc, int ipfd, PIXELFORMATDESCRIPTOR* ppfd);
// BOOL (APIENTRYP ptrglSwapBuffers)(HDC hdc);
// BOOL (APIENTRYP ptrglShareLists)(HGLRC hrcSrvShare, HGLRC hrcSrvSource);
// HGLRC (APIENTRYP ptrglCreateLayerContext)(HDC hDc, int level);
// BOOL (APIENTRYP ptrglDescribeLayerPlane)(HDC hDc, int pixelFormat, int layerPlane, UINT nBytes, LAYERPLANEDESCRIPTOR* plpd);
// int (APIENTRYP ptrglSetLayerPaletteEntries)(HDC hdc, int iLayerPlane, int iStart, int cEntries, COLORREF* pcr);
// int (APIENTRYP ptrglGetLayerPaletteEntries)(HDC hdc, int iLayerPlane, int iStart, int cEntries, COLORREF* pcr);
// BOOL (APIENTRYP ptrglRealizeLayerPalette)(HDC hdc, int iLayerPlane, BOOL bRealize);
// BOOL (APIENTRYP ptrglSwapLayerBuffers)(HDC hdc, UINT fuFlags);
// BOOL (APIENTRYP ptrglUseFontBitmapsA)(HDC hDC, DWORD first, DWORD count, DWORD listBase);
// BOOL (APIENTRYP ptrglUseFontBitmapsW)(HDC hDC, DWORD first, DWORD count, DWORD listBase);
// 
// //  3DFX_multisample
// //  3DL_stereo_control
// BOOL goglSetStereoEmitterState3DL(HDC hDC, UINT uState) {
// 	return (*ptrglSetStereoEmitterState3DL)(hDC, uState);
// }
// //  AMD_gpu_association
// UINT goglGetGPUIDsAMD(UINT maxCount, UINT* ids) {
// 	return (*ptrglGetGPUIDsAMD)(maxCount, ids);
// }
// INT goglGetGPUInfoAMD(UINT id, int property, GLenum dataType, UINT size, void* data) {
// 	return (*ptrglGetGPUInfoAMD)(id, property, dataType, size, data);
// }
// UINT goglGetContextGPUIDAMD(HGLRC hglrc) {
// 	return (*ptrglGetContextGPUIDAMD)(hglrc);
// }
// HGLRC goglCreateAssociatedContextAMD(UINT id) {
// 	return (*ptrglCreateAssociatedContextAMD)(id);
// }
// HGLRC goglCreateAssociatedContextAttribsAMD(UINT id, HGLRC hShareContext, int* attribList) {
// 	return (*ptrglCreateAssociatedContextAttribsAMD)(id, hShareContext, attribList);
// }
// BOOL goglDeleteAssociatedContextAMD(HGLRC hglrc) {
// 	return (*ptrglDeleteAssociatedContextAMD)(hglrc);
// }
// BOOL goglMakeAssociatedContextCurrentAMD(HGLRC hglrc) {
// 	return (*ptrglMakeAssociatedContextCurrentAMD)(hglrc);
// }
// HGLRC goglGetCurrentAssociatedContextAMD() {
// 	return (*ptrglGetCurrentAssociatedContextAMD)();
// }
// VOID goglBlitContextFramebufferAMD(HGLRC dstCtx, GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter) {
// 	return (*ptrglBlitContextFramebufferAMD)(dstCtx, srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// //  ARB_buffer_region
// HANDLE goglCreateBufferRegionARB(HDC hDC, int iLayerPlane, UINT uType) {
// 	return (*ptrglCreateBufferRegionARB)(hDC, iLayerPlane, uType);
// }
// VOID goglDeleteBufferRegionARB(HANDLE hRegion) {
// 	return (*ptrglDeleteBufferRegionARB)(hRegion);
// }
// BOOL goglSaveBufferRegionARB(HANDLE hRegion, int x, int y, int width, int height) {
// 	return (*ptrglSaveBufferRegionARB)(hRegion, x, y, width, height);
// }
// BOOL goglRestoreBufferRegionARB(HANDLE hRegion, int x, int y, int width, int height, int xSrc, int ySrc) {
// 	return (*ptrglRestoreBufferRegionARB)(hRegion, x, y, width, height, xSrc, ySrc);
// }
// //  ARB_create_context
// HGLRC goglCreateContextAttribsARB(HDC hDC, HGLRC hShareContext, int* attribList) {
// 	return (*ptrglCreateContextAttribsARB)(hDC, hShareContext, attribList);
// }
// //  ARB_create_context_profile
// //  ARB_create_context_robustness
// //  ARB_extensions_string
// const char * goglGetExtensionsStringARB(HDC hdc) {
// 	return (*ptrglGetExtensionsStringARB)(hdc);
// }
// //  ARB_framebuffer_sRGB
// //  ARB_make_current_read
// BOOL goglMakeContextCurrentARB(HDC hDrawDC, HDC hReadDC, HGLRC hglrc) {
// 	return (*ptrglMakeContextCurrentARB)(hDrawDC, hReadDC, hglrc);
// }
// HDC goglGetCurrentReadDCARB() {
// 	return (*ptrglGetCurrentReadDCARB)();
// }
// //  ARB_multisample
// //  ARB_pbuffer
// HPBUFFERARB goglCreatePbufferARB(HDC hDC, int iPixelFormat, int iWidth, int iHeight, int* piAttribList) {
// 	return (*ptrglCreatePbufferARB)(hDC, iPixelFormat, iWidth, iHeight, piAttribList);
// }
// HDC goglGetPbufferDCARB(HPBUFFERARB hPbuffer) {
// 	return (*ptrglGetPbufferDCARB)(hPbuffer);
// }
// int goglReleasePbufferDCARB(HPBUFFERARB hPbuffer, HDC hDC) {
// 	return (*ptrglReleasePbufferDCARB)(hPbuffer, hDC);
// }
// BOOL goglDestroyPbufferARB(HPBUFFERARB hPbuffer) {
// 	return (*ptrglDestroyPbufferARB)(hPbuffer);
// }
// BOOL goglQueryPbufferARB(HPBUFFERARB hPbuffer, int iAttribute, int* piValue) {
// 	return (*ptrglQueryPbufferARB)(hPbuffer, iAttribute, piValue);
// }
// //  ARB_pixel_format
// BOOL goglGetPixelFormatAttribivARB(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, int* piValues) {
// 	return (*ptrglGetPixelFormatAttribivARB)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, piValues);
// }
// BOOL goglGetPixelFormatAttribfvARB(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, FLOAT* pfValues) {
// 	return (*ptrglGetPixelFormatAttribfvARB)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, pfValues);
// }
// BOOL goglChoosePixelFormatARB(HDC hdc, int* piAttribIList, FLOAT* pfAttribFList, UINT nMaxFormats, int* piFormats, UINT* nNumFormats) {
// 	return (*ptrglChoosePixelFormatARB)(hdc, piAttribIList, pfAttribFList, nMaxFormats, piFormats, nNumFormats);
// }
// //  ARB_pixel_format_float
// //  ARB_render_texture
// BOOL goglBindTexImageARB(HPBUFFERARB hPbuffer, int iBuffer) {
// 	return (*ptrglBindTexImageARB)(hPbuffer, iBuffer);
// }
// BOOL goglReleaseTexImageARB(HPBUFFERARB hPbuffer, int iBuffer) {
// 	return (*ptrglReleaseTexImageARB)(hPbuffer, iBuffer);
// }
// BOOL goglSetPbufferAttribARB(HPBUFFERARB hPbuffer, int* piAttribList) {
// 	return (*ptrglSetPbufferAttribARB)(hPbuffer, piAttribList);
// }
// //  ATI_pixel_format_float
// //  EXT_depth_float
// //  EXT_display_color_table
// GLboolean goglCreateDisplayColorTableEXT(GLushort id) {
// 	return (*ptrglCreateDisplayColorTableEXT)(id);
// }
// GLboolean goglLoadDisplayColorTableEXT(GLushort* table, GLuint length) {
// 	return (*ptrglLoadDisplayColorTableEXT)(table, length);
// }
// GLboolean goglBindDisplayColorTableEXT(GLushort id) {
// 	return (*ptrglBindDisplayColorTableEXT)(id);
// }
// VOID goglDestroyDisplayColorTableEXT(GLushort id) {
// 	return (*ptrglDestroyDisplayColorTableEXT)(id);
// }
// //  EXT_extensions_string
// const char * goglGetExtensionsStringEXT() {
// 	return (*ptrglGetExtensionsStringEXT)();
// }
// //  EXT_framebuffer_sRGB
// //  EXT_make_current_read
// BOOL goglMakeContextCurrentEXT(HDC hDrawDC, HDC hReadDC, HGLRC hglrc) {
// 	return (*ptrglMakeContextCurrentEXT)(hDrawDC, hReadDC, hglrc);
// }
// HDC goglGetCurrentReadDCEXT() {
// 	return (*ptrglGetCurrentReadDCEXT)();
// }
// //  EXT_multisample
// //  EXT_pbuffer
// HPBUFFEREXT goglCreatePbufferEXT(HDC hDC, int iPixelFormat, int iWidth, int iHeight, int* piAttribList) {
// 	return (*ptrglCreatePbufferEXT)(hDC, iPixelFormat, iWidth, iHeight, piAttribList);
// }
// HDC goglGetPbufferDCEXT(HPBUFFEREXT hPbuffer) {
// 	return (*ptrglGetPbufferDCEXT)(hPbuffer);
// }
// int goglReleasePbufferDCEXT(HPBUFFEREXT hPbuffer, HDC hDC) {
// 	return (*ptrglReleasePbufferDCEXT)(hPbuffer, hDC);
// }
// BOOL goglDestroyPbufferEXT(HPBUFFEREXT hPbuffer) {
// 	return (*ptrglDestroyPbufferEXT)(hPbuffer);
// }
// BOOL goglQueryPbufferEXT(HPBUFFEREXT hPbuffer, int iAttribute, int* piValue) {
// 	return (*ptrglQueryPbufferEXT)(hPbuffer, iAttribute, piValue);
// }
// //  EXT_pixel_format
// BOOL goglGetPixelFormatAttribivEXT(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, int* piValues) {
// 	return (*ptrglGetPixelFormatAttribivEXT)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, piValues);
// }
// BOOL goglGetPixelFormatAttribfvEXT(HDC hdc, int iPixelFormat, int iLayerPlane, UINT nAttributes, int* piAttributes, FLOAT* pfValues) {
// 	return (*ptrglGetPixelFormatAttribfvEXT)(hdc, iPixelFormat, iLayerPlane, nAttributes, piAttributes, pfValues);
// }
// BOOL goglChoosePixelFormatEXT(HDC hdc, int* piAttribIList, FLOAT* pfAttribFList, UINT nMaxFormats, int* piFormats, UINT* nNumFormats) {
// 	return (*ptrglChoosePixelFormatEXT)(hdc, piAttribIList, pfAttribFList, nMaxFormats, piFormats, nNumFormats);
// }
// //  EXT_pixel_format_packed_float
// //  EXT_swap_control
// BOOL goglSwapIntervalEXT(int interval) {
// 	return (*ptrglSwapIntervalEXT)(interval);
// }
// int goglGetSwapIntervalEXT() {
// 	return (*ptrglGetSwapIntervalEXT)();
// }
// //  EXT_swap_control_tear
// //  I3D_digital_video_control
// BOOL goglGetDigitalVideoParametersI3D(HDC hDC, int iAttribute, int* piValue) {
// 	return (*ptrglGetDigitalVideoParametersI3D)(hDC, iAttribute, piValue);
// }
// BOOL goglSetDigitalVideoParametersI3D(HDC hDC, int iAttribute, int* piValue) {
// 	return (*ptrglSetDigitalVideoParametersI3D)(hDC, iAttribute, piValue);
// }
// //  I3D_gamma
// BOOL goglGetGammaTableParametersI3D(HDC hDC, int iAttribute, int* piValue) {
// 	return (*ptrglGetGammaTableParametersI3D)(hDC, iAttribute, piValue);
// }
// BOOL goglSetGammaTableParametersI3D(HDC hDC, int iAttribute, int* piValue) {
// 	return (*ptrglSetGammaTableParametersI3D)(hDC, iAttribute, piValue);
// }
// BOOL goglGetGammaTableI3D(HDC hDC, int iEntries, USHORT* puRed, USHORT* puGreen, USHORT* puBlue) {
// 	return (*ptrglGetGammaTableI3D)(hDC, iEntries, puRed, puGreen, puBlue);
// }
// BOOL goglSetGammaTableI3D(HDC hDC, int iEntries, USHORT* puRed, USHORT* puGreen, USHORT* puBlue) {
// 	return (*ptrglSetGammaTableI3D)(hDC, iEntries, puRed, puGreen, puBlue);
// }
// //  I3D_genlock
// BOOL goglEnableGenlockI3D(HDC hDC) {
// 	return (*ptrglEnableGenlockI3D)(hDC);
// }
// BOOL goglDisableGenlockI3D(HDC hDC) {
// 	return (*ptrglDisableGenlockI3D)(hDC);
// }
// BOOL goglIsEnabledGenlockI3D(HDC hDC, BOOL* pFlag) {
// 	return (*ptrglIsEnabledGenlockI3D)(hDC, pFlag);
// }
// BOOL goglGenlockSourceI3D(HDC hDC, UINT uSource) {
// 	return (*ptrglGenlockSourceI3D)(hDC, uSource);
// }
// BOOL goglGetGenlockSourceI3D(HDC hDC, UINT* uSource) {
// 	return (*ptrglGetGenlockSourceI3D)(hDC, uSource);
// }
// BOOL goglGenlockSourceEdgeI3D(HDC hDC, UINT uEdge) {
// 	return (*ptrglGenlockSourceEdgeI3D)(hDC, uEdge);
// }
// BOOL goglGetGenlockSourceEdgeI3D(HDC hDC, UINT* uEdge) {
// 	return (*ptrglGetGenlockSourceEdgeI3D)(hDC, uEdge);
// }
// BOOL goglGenlockSampleRateI3D(HDC hDC, UINT uRate) {
// 	return (*ptrglGenlockSampleRateI3D)(hDC, uRate);
// }
// BOOL goglGetGenlockSampleRateI3D(HDC hDC, UINT* uRate) {
// 	return (*ptrglGetGenlockSampleRateI3D)(hDC, uRate);
// }
// BOOL goglGenlockSourceDelayI3D(HDC hDC, UINT uDelay) {
// 	return (*ptrglGenlockSourceDelayI3D)(hDC, uDelay);
// }
// BOOL goglGetGenlockSourceDelayI3D(HDC hDC, UINT* uDelay) {
// 	return (*ptrglGetGenlockSourceDelayI3D)(hDC, uDelay);
// }
// BOOL goglQueryGenlockMaxSourceDelayI3D(HDC hDC, UINT* uMaxLineDelay, UINT* uMaxPixelDelay) {
// 	return (*ptrglQueryGenlockMaxSourceDelayI3D)(hDC, uMaxLineDelay, uMaxPixelDelay);
// }
// //  I3D_image_buffer
// LPVOID goglCreateImageBufferI3D(HDC hDC, DWORD dwSize, UINT uFlags) {
// 	return (*ptrglCreateImageBufferI3D)(hDC, dwSize, uFlags);
// }
// BOOL goglDestroyImageBufferI3D(HDC hDC, LPVOID pAddress) {
// 	return (*ptrglDestroyImageBufferI3D)(hDC, pAddress);
// }
// BOOL goglAssociateImageBufferEventsI3D(HDC hDC, HANDLE* pEvent, LPVOID* pAddress, DWORD* pSize, UINT count) {
// 	return (*ptrglAssociateImageBufferEventsI3D)(hDC, pEvent, pAddress, pSize, count);
// }
// BOOL goglReleaseImageBufferEventsI3D(HDC hDC, LPVOID* pAddress, UINT count) {
// 	return (*ptrglReleaseImageBufferEventsI3D)(hDC, pAddress, count);
// }
// //  I3D_swap_frame_lock
// BOOL goglEnableFrameLockI3D() {
// 	return (*ptrglEnableFrameLockI3D)();
// }
// BOOL goglDisableFrameLockI3D() {
// 	return (*ptrglDisableFrameLockI3D)();
// }
// BOOL goglIsEnabledFrameLockI3D(BOOL* pFlag) {
// 	return (*ptrglIsEnabledFrameLockI3D)(pFlag);
// }
// BOOL goglQueryFrameLockMasterI3D(BOOL* pFlag) {
// 	return (*ptrglQueryFrameLockMasterI3D)(pFlag);
// }
// //  I3D_swap_frame_usage
// BOOL goglGetFrameUsageI3D(float* pUsage) {
// 	return (*ptrglGetFrameUsageI3D)(pUsage);
// }
// BOOL goglBeginFrameTrackingI3D() {
// 	return (*ptrglBeginFrameTrackingI3D)();
// }
// BOOL goglEndFrameTrackingI3D() {
// 	return (*ptrglEndFrameTrackingI3D)();
// }
// BOOL goglQueryFrameTrackingI3D(DWORD* pFrameCount, DWORD* pMissedFrames, float* pLastMissedUsage) {
// 	return (*ptrglQueryFrameTrackingI3D)(pFrameCount, pMissedFrames, pLastMissedUsage);
// }
// //  NV_DX_interop
// BOOL goglDXSetResourceShareHandleNV(void* dxObject, HANDLE shareHandle) {
// 	return (*ptrglDXSetResourceShareHandleNV)(dxObject, shareHandle);
// }
// HANDLE goglDXOpenDeviceNV(void* dxDevice) {
// 	return (*ptrglDXOpenDeviceNV)(dxDevice);
// }
// BOOL goglDXCloseDeviceNV(HANDLE hDevice) {
// 	return (*ptrglDXCloseDeviceNV)(hDevice);
// }
// HANDLE goglDXRegisterObjectNV(HANDLE hDevice, void* dxObject, GLuint name, GLenum type_, GLenum access) {
// 	return (*ptrglDXRegisterObjectNV)(hDevice, dxObject, name, type_, access);
// }
// BOOL goglDXUnregisterObjectNV(HANDLE hDevice, HANDLE hObject) {
// 	return (*ptrglDXUnregisterObjectNV)(hDevice, hObject);
// }
// BOOL goglDXObjectAccessNV(HANDLE hObject, GLenum access) {
// 	return (*ptrglDXObjectAccessNV)(hObject, access);
// }
// BOOL goglDXLockObjectsNV(HANDLE hDevice, GLint count, HANDLE* hObjects) {
// 	return (*ptrglDXLockObjectsNV)(hDevice, count, hObjects);
// }
// BOOL goglDXUnlockObjectsNV(HANDLE hDevice, GLint count, HANDLE* hObjects) {
// 	return (*ptrglDXUnlockObjectsNV)(hDevice, count, hObjects);
// }
// //  NV_DX_interop2
// //  NV_copy_image
// BOOL goglCopyImageSubDataNV(HGLRC hSrcRC, GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, HGLRC hDstRC, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei width, GLsizei height, GLsizei depth) {
// 	return (*ptrglCopyImageSubDataNV)(hSrcRC, srcName, srcTarget, srcLevel, srcX, srcY, srcZ, hDstRC, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, width, height, depth);
// }
// //  NV_float_buffer
// //  NV_gpu_affinity
// BOOL goglEnumGpusNV(UINT iGpuIndex, HGPUNV* phGpu) {
// 	return (*ptrglEnumGpusNV)(iGpuIndex, phGpu);
// }
// BOOL goglEnumGpuDevicesNV(HGPUNV hGpu, UINT iDeviceIndex, PGPU_DEVICE lpGpuDevice) {
// 	return (*ptrglEnumGpuDevicesNV)(hGpu, iDeviceIndex, lpGpuDevice);
// }
// HDC goglCreateAffinityDCNV(HGPUNV* phGpuList) {
// 	return (*ptrglCreateAffinityDCNV)(phGpuList);
// }
// BOOL goglEnumGpusFromAffinityDCNV(HDC hAffinityDC, UINT iGpuIndex, HGPUNV* hGpu) {
// 	return (*ptrglEnumGpusFromAffinityDCNV)(hAffinityDC, iGpuIndex, hGpu);
// }
// BOOL goglDeleteDCNV(HDC hdc) {
// 	return (*ptrglDeleteDCNV)(hdc);
// }
// //  NV_multisample_coverage
// //  NV_present_video
// int goglEnumerateVideoDevicesNV(HDC hDC, HVIDEOOUTPUTDEVICENV* phDeviceList) {
// 	return (*ptrglEnumerateVideoDevicesNV)(hDC, phDeviceList);
// }
// BOOL goglBindVideoDeviceNV(HDC hDC, unsigned int uVideoSlot, HVIDEOOUTPUTDEVICENV hVideoDevice, int* piAttribList) {
// 	return (*ptrglBindVideoDeviceNV)(hDC, uVideoSlot, hVideoDevice, piAttribList);
// }
// BOOL goglQueryCurrentContextNV(int iAttribute, int* piValue) {
// 	return (*ptrglQueryCurrentContextNV)(iAttribute, piValue);
// }
// //  NV_swap_group
// BOOL goglJoinSwapGroupNV(HDC hDC, GLuint group) {
// 	return (*ptrglJoinSwapGroupNV)(hDC, group);
// }
// BOOL goglBindSwapBarrierNV(GLuint group, GLuint barrier) {
// 	return (*ptrglBindSwapBarrierNV)(group, barrier);
// }
// BOOL goglQuerySwapGroupNV(HDC hDC, GLuint* group, GLuint* barrier) {
// 	return (*ptrglQuerySwapGroupNV)(hDC, group, barrier);
// }
// BOOL goglQueryMaxSwapGroupsNV(HDC hDC, GLuint* maxGroups, GLuint* maxBarriers) {
// 	return (*ptrglQueryMaxSwapGroupsNV)(hDC, maxGroups, maxBarriers);
// }
// BOOL goglQueryFrameCountNV(HDC hDC, GLuint* count) {
// 	return (*ptrglQueryFrameCountNV)(hDC, count);
// }
// BOOL goglResetFrameCountNV(HDC hDC) {
// 	return (*ptrglResetFrameCountNV)(hDC);
// }
// //  NV_vertex_array_range
// void* goglAllocateMemoryNV(GLsizei size, GLfloat readfreq, GLfloat writefreq, GLfloat priority) {
// 	return (*ptrglAllocateMemoryNV)(size, readfreq, writefreq, priority);
// }
// void goglFreeMemoryNV(void* pointer) {
// 	(*ptrglFreeMemoryNV)(pointer);
// }
// //  NV_video_capture
// BOOL goglBindVideoCaptureDeviceNV(UINT uVideoSlot, HVIDEOINPUTDEVICENV hDevice) {
// 	return (*ptrglBindVideoCaptureDeviceNV)(uVideoSlot, hDevice);
// }
// UINT goglEnumerateVideoCaptureDevicesNV(HDC hDc, HVIDEOINPUTDEVICENV* phDeviceList) {
// 	return (*ptrglEnumerateVideoCaptureDevicesNV)(hDc, phDeviceList);
// }
// BOOL goglLockVideoCaptureDeviceNV(HDC hDc, HVIDEOINPUTDEVICENV hDevice) {
// 	return (*ptrglLockVideoCaptureDeviceNV)(hDc, hDevice);
// }
// BOOL goglQueryVideoCaptureDeviceNV(HDC hDc, HVIDEOINPUTDEVICENV hDevice, int iAttribute, int* piValue) {
// 	return (*ptrglQueryVideoCaptureDeviceNV)(hDc, hDevice, iAttribute, piValue);
// }
// BOOL goglReleaseVideoCaptureDeviceNV(HDC hDc, HVIDEOINPUTDEVICENV hDevice) {
// 	return (*ptrglReleaseVideoCaptureDeviceNV)(hDc, hDevice);
// }
// //  NV_video_output
// BOOL goglGetVideoDeviceNV(HDC hDC, int numDevices, HPVIDEODEV* hVideoDevice) {
// 	return (*ptrglGetVideoDeviceNV)(hDC, numDevices, hVideoDevice);
// }
// BOOL goglReleaseVideoDeviceNV(HPVIDEODEV hVideoDevice) {
// 	return (*ptrglReleaseVideoDeviceNV)(hVideoDevice);
// }
// BOOL goglBindVideoImageNV(HPVIDEODEV hVideoDevice, HPBUFFERARB hPbuffer, int iVideoBuffer) {
// 	return (*ptrglBindVideoImageNV)(hVideoDevice, hPbuffer, iVideoBuffer);
// }
// BOOL goglReleaseVideoImageNV(HPBUFFERARB hPbuffer, int iVideoBuffer) {
// 	return (*ptrglReleaseVideoImageNV)(hPbuffer, iVideoBuffer);
// }
// BOOL goglSendPbufferToVideoNV(HPBUFFERARB hPbuffer, int iBufferType, unsigned long* pulCounterPbuffer, BOOL bBlock) {
// 	return (*ptrglSendPbufferToVideoNV)(hPbuffer, iBufferType, pulCounterPbuffer, bBlock);
// }
// BOOL goglGetVideoInfoNV(HPVIDEODEV hpVideoDevice, unsigned long* pulCounterOutputPbuffer, unsigned long* pulCounterOutputVideo) {
// 	return (*ptrglGetVideoInfoNV)(hpVideoDevice, pulCounterOutputPbuffer, pulCounterOutputVideo);
// }
// //  OML_sync_control
// BOOL goglGetSyncValuesOML(HDC hdc, INT64* ust, INT64* msc, INT64* sbc) {
// 	return (*ptrglGetSyncValuesOML)(hdc, ust, msc, sbc);
// }
// BOOL goglGetMscRateOML(HDC hdc, INT32* numerator, INT32* denominator) {
// 	return (*ptrglGetMscRateOML)(hdc, numerator, denominator);
// }
// INT64 goglSwapBuffersMscOML(HDC hdc, INT64 target_msc, INT64 divisor, INT64 remainder) {
// 	return (*ptrglSwapBuffersMscOML)(hdc, target_msc, divisor, remainder);
// }
// INT64 goglSwapLayerBuffersMscOML(HDC hdc, int fuPlanes, INT64 target_msc, INT64 divisor, INT64 remainder) {
// 	return (*ptrglSwapLayerBuffersMscOML)(hdc, fuPlanes, target_msc, divisor, remainder);
// }
// BOOL goglWaitForMscOML(HDC hdc, INT64 target_msc, INT64 divisor, INT64 remainder, INT64* ust, INT64* msc, INT64* sbc) {
// 	return (*ptrglWaitForMscOML)(hdc, target_msc, divisor, remainder, ust, msc, sbc);
// }
// BOOL goglWaitForSbcOML(HDC hdc, INT64 target_sbc, INT64* ust, INT64* msc, INT64* sbc) {
// 	return (*ptrglWaitForSbcOML)(hdc, target_sbc, ust, msc, sbc);
// }
// //  wgl
// HGLRC goglCreateContext(HDC hDc) {
// 	return (*ptrglCreateContext)(hDc);
// }
// BOOL goglDeleteContext(HGLRC oldContext) {
// 	return (*ptrglDeleteContext)(oldContext);
// }
// HGLRC goglGetCurrentContext() {
// 	return (*ptrglGetCurrentContext)();
// }
// BOOL goglMakeCurrent(HDC hDc, HGLRC newContext) {
// 	return (*ptrglMakeCurrent)(hDc, newContext);
// }
// BOOL goglCopyContext(HGLRC hglrcSrc, HGLRC hglrcDst, UINT mask) {
// 	return (*ptrglCopyContext)(hglrcSrc, hglrcDst, mask);
// }
// int goglChoosePixelFormat(HDC hDc, PIXELFORMATDESCRIPTOR* pPfd) {
// 	return (*ptrglChoosePixelFormat)(hDc, pPfd);
// }
// HDC goglGetCurrentDC() {
// 	return (*ptrglGetCurrentDC)();
// }
// PROC goglGetDefaultProcAddress(LPCSTR lpszProc) {
// 	return (*ptrglGetDefaultProcAddress)(lpszProc);
// }
// PROC goglGetProcAddress(LPCSTR lpszProc) {
// 	return (*ptrglGetProcAddress)(lpszProc);
// }
// int goglGetPixelFormat(HDC hdc) {
// 	return (*ptrglGetPixelFormat)(hdc);
// }
// BOOL goglSetPixelFormat(HDC hdc, int ipfd, PIXELFORMATDESCRIPTOR* ppfd) {
// 	return (*ptrglSetPixelFormat)(hdc, ipfd, ppfd);
// }
// BOOL goglSwapBuffers(HDC hdc) {
// 	return (*ptrglSwapBuffers)(hdc);
// }
// BOOL goglShareLists(HGLRC hrcSrvShare, HGLRC hrcSrvSource) {
// 	return (*ptrglShareLists)(hrcSrvShare, hrcSrvSource);
// }
// HGLRC goglCreateLayerContext(HDC hDc, int level) {
// 	return (*ptrglCreateLayerContext)(hDc, level);
// }
// BOOL goglDescribeLayerPlane(HDC hDc, int pixelFormat, int layerPlane, UINT nBytes, LAYERPLANEDESCRIPTOR* plpd) {
// 	return (*ptrglDescribeLayerPlane)(hDc, pixelFormat, layerPlane, nBytes, plpd);
// }
// int goglSetLayerPaletteEntries(HDC hdc, int iLayerPlane, int iStart, int cEntries, COLORREF* pcr) {
// 	return (*ptrglSetLayerPaletteEntries)(hdc, iLayerPlane, iStart, cEntries, pcr);
// }
// int goglGetLayerPaletteEntries(HDC hdc, int iLayerPlane, int iStart, int cEntries, COLORREF* pcr) {
// 	return (*ptrglGetLayerPaletteEntries)(hdc, iLayerPlane, iStart, cEntries, pcr);
// }
// BOOL goglRealizeLayerPalette(HDC hdc, int iLayerPlane, BOOL bRealize) {
// 	return (*ptrglRealizeLayerPalette)(hdc, iLayerPlane, bRealize);
// }
// BOOL goglSwapLayerBuffers(HDC hdc, UINT fuFlags) {
// 	return (*ptrglSwapLayerBuffers)(hdc, fuFlags);
// }
// BOOL goglUseFontBitmapsA(HDC hDC, DWORD first, DWORD count, DWORD listBase) {
// 	return (*ptrglUseFontBitmapsA)(hDC, first, count, listBase);
// }
// BOOL goglUseFontBitmapsW(HDC hDC, DWORD first, DWORD count, DWORD listBase) {
// 	return (*ptrglUseFontBitmapsW)(hDC, first, count, listBase);
// }
// 
// int init_3DFX_multisample() {
// 	return 0;
// }
// int init_3DL_stereo_control() {
// 	ptrglSetStereoEmitterState3DL = goglGetProcAddress("glSetStereoEmitterState3DL");
// 	if(ptrglSetStereoEmitterState3DL == NULL) return 1;
// 	return 0;
// }
// int init_AMD_gpu_association() {
// 	ptrglGetGPUIDsAMD = goglGetProcAddress("glGetGPUIDsAMD");
// 	if(ptrglGetGPUIDsAMD == NULL) return 1;
// 	ptrglGetGPUInfoAMD = goglGetProcAddress("glGetGPUInfoAMD");
// 	if(ptrglGetGPUInfoAMD == NULL) return 1;
// 	ptrglGetContextGPUIDAMD = goglGetProcAddress("glGetContextGPUIDAMD");
// 	if(ptrglGetContextGPUIDAMD == NULL) return 1;
// 	ptrglCreateAssociatedContextAMD = goglGetProcAddress("glCreateAssociatedContextAMD");
// 	if(ptrglCreateAssociatedContextAMD == NULL) return 1;
// 	ptrglCreateAssociatedContextAttribsAMD = goglGetProcAddress("glCreateAssociatedContextAttribsAMD");
// 	if(ptrglCreateAssociatedContextAttribsAMD == NULL) return 1;
// 	ptrglDeleteAssociatedContextAMD = goglGetProcAddress("glDeleteAssociatedContextAMD");
// 	if(ptrglDeleteAssociatedContextAMD == NULL) return 1;
// 	ptrglMakeAssociatedContextCurrentAMD = goglGetProcAddress("glMakeAssociatedContextCurrentAMD");
// 	if(ptrglMakeAssociatedContextCurrentAMD == NULL) return 1;
// 	ptrglGetCurrentAssociatedContextAMD = goglGetProcAddress("glGetCurrentAssociatedContextAMD");
// 	if(ptrglGetCurrentAssociatedContextAMD == NULL) return 1;
// 	ptrglBlitContextFramebufferAMD = goglGetProcAddress("glBlitContextFramebufferAMD");
// 	if(ptrglBlitContextFramebufferAMD == NULL) return 1;
// 	return 0;
// }
// int init_ARB_buffer_region() {
// 	ptrglCreateBufferRegionARB = goglGetProcAddress("glCreateBufferRegionARB");
// 	if(ptrglCreateBufferRegionARB == NULL) return 1;
// 	ptrglDeleteBufferRegionARB = goglGetProcAddress("glDeleteBufferRegionARB");
// 	if(ptrglDeleteBufferRegionARB == NULL) return 1;
// 	ptrglSaveBufferRegionARB = goglGetProcAddress("glSaveBufferRegionARB");
// 	if(ptrglSaveBufferRegionARB == NULL) return 1;
// 	ptrglRestoreBufferRegionARB = goglGetProcAddress("glRestoreBufferRegionARB");
// 	if(ptrglRestoreBufferRegionARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_create_context() {
// 	ptrglCreateContextAttribsARB = goglGetProcAddress("glCreateContextAttribsARB");
// 	if(ptrglCreateContextAttribsARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_create_context_profile() {
// 	return 0;
// }
// int init_ARB_create_context_robustness() {
// 	return 0;
// }
// int init_ARB_extensions_string() {
// 	ptrglGetExtensionsStringARB = goglGetProcAddress("glGetExtensionsStringARB");
// 	if(ptrglGetExtensionsStringARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_framebuffer_sRGB() {
// 	return 0;
// }
// int init_ARB_make_current_read() {
// 	ptrglMakeContextCurrentARB = goglGetProcAddress("glMakeContextCurrentARB");
// 	if(ptrglMakeContextCurrentARB == NULL) return 1;
// 	ptrglGetCurrentReadDCARB = goglGetProcAddress("glGetCurrentReadDCARB");
// 	if(ptrglGetCurrentReadDCARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_multisample() {
// 	return 0;
// }
// int init_ARB_pbuffer() {
// 	ptrglCreatePbufferARB = goglGetProcAddress("glCreatePbufferARB");
// 	if(ptrglCreatePbufferARB == NULL) return 1;
// 	ptrglGetPbufferDCARB = goglGetProcAddress("glGetPbufferDCARB");
// 	if(ptrglGetPbufferDCARB == NULL) return 1;
// 	ptrglReleasePbufferDCARB = goglGetProcAddress("glReleasePbufferDCARB");
// 	if(ptrglReleasePbufferDCARB == NULL) return 1;
// 	ptrglDestroyPbufferARB = goglGetProcAddress("glDestroyPbufferARB");
// 	if(ptrglDestroyPbufferARB == NULL) return 1;
// 	ptrglQueryPbufferARB = goglGetProcAddress("glQueryPbufferARB");
// 	if(ptrglQueryPbufferARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_pixel_format() {
// 	ptrglGetPixelFormatAttribivARB = goglGetProcAddress("glGetPixelFormatAttribivARB");
// 	if(ptrglGetPixelFormatAttribivARB == NULL) return 1;
// 	ptrglGetPixelFormatAttribfvARB = goglGetProcAddress("glGetPixelFormatAttribfvARB");
// 	if(ptrglGetPixelFormatAttribfvARB == NULL) return 1;
// 	ptrglChoosePixelFormatARB = goglGetProcAddress("glChoosePixelFormatARB");
// 	if(ptrglChoosePixelFormatARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_pixel_format_float() {
// 	return 0;
// }
// int init_ARB_render_texture() {
// 	ptrglBindTexImageARB = goglGetProcAddress("glBindTexImageARB");
// 	if(ptrglBindTexImageARB == NULL) return 1;
// 	ptrglReleaseTexImageARB = goglGetProcAddress("glReleaseTexImageARB");
// 	if(ptrglReleaseTexImageARB == NULL) return 1;
// 	ptrglSetPbufferAttribARB = goglGetProcAddress("glSetPbufferAttribARB");
// 	if(ptrglSetPbufferAttribARB == NULL) return 1;
// 	return 0;
// }
// int init_ATI_pixel_format_float() {
// 	return 0;
// }
// int init_EXT_depth_float() {
// 	return 0;
// }
// int init_EXT_display_color_table() {
// 	ptrglCreateDisplayColorTableEXT = goglGetProcAddress("glCreateDisplayColorTableEXT");
// 	if(ptrglCreateDisplayColorTableEXT == NULL) return 1;
// 	ptrglLoadDisplayColorTableEXT = goglGetProcAddress("glLoadDisplayColorTableEXT");
// 	if(ptrglLoadDisplayColorTableEXT == NULL) return 1;
// 	ptrglBindDisplayColorTableEXT = goglGetProcAddress("glBindDisplayColorTableEXT");
// 	if(ptrglBindDisplayColorTableEXT == NULL) return 1;
// 	ptrglDestroyDisplayColorTableEXT = goglGetProcAddress("glDestroyDisplayColorTableEXT");
// 	if(ptrglDestroyDisplayColorTableEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_extensions_string() {
// 	ptrglGetExtensionsStringEXT = goglGetProcAddress("glGetExtensionsStringEXT");
// 	if(ptrglGetExtensionsStringEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_framebuffer_sRGB() {
// 	return 0;
// }
// int init_EXT_make_current_read() {
// 	ptrglMakeContextCurrentEXT = goglGetProcAddress("glMakeContextCurrentEXT");
// 	if(ptrglMakeContextCurrentEXT == NULL) return 1;
// 	ptrglGetCurrentReadDCEXT = goglGetProcAddress("glGetCurrentReadDCEXT");
// 	if(ptrglGetCurrentReadDCEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_multisample() {
// 	return 0;
// }
// int init_EXT_pbuffer() {
// 	ptrglCreatePbufferEXT = goglGetProcAddress("glCreatePbufferEXT");
// 	if(ptrglCreatePbufferEXT == NULL) return 1;
// 	ptrglGetPbufferDCEXT = goglGetProcAddress("glGetPbufferDCEXT");
// 	if(ptrglGetPbufferDCEXT == NULL) return 1;
// 	ptrglReleasePbufferDCEXT = goglGetProcAddress("glReleasePbufferDCEXT");
// 	if(ptrglReleasePbufferDCEXT == NULL) return 1;
// 	ptrglDestroyPbufferEXT = goglGetProcAddress("glDestroyPbufferEXT");
// 	if(ptrglDestroyPbufferEXT == NULL) return 1;
// 	ptrglQueryPbufferEXT = goglGetProcAddress("glQueryPbufferEXT");
// 	if(ptrglQueryPbufferEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_pixel_format() {
// 	ptrglGetPixelFormatAttribivEXT = goglGetProcAddress("glGetPixelFormatAttribivEXT");
// 	if(ptrglGetPixelFormatAttribivEXT == NULL) return 1;
// 	ptrglGetPixelFormatAttribfvEXT = goglGetProcAddress("glGetPixelFormatAttribfvEXT");
// 	if(ptrglGetPixelFormatAttribfvEXT == NULL) return 1;
// 	ptrglChoosePixelFormatEXT = goglGetProcAddress("glChoosePixelFormatEXT");
// 	if(ptrglChoosePixelFormatEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_pixel_format_packed_float() {
// 	return 0;
// }
// int init_EXT_swap_control() {
// 	ptrglSwapIntervalEXT = goglGetProcAddress("glSwapIntervalEXT");
// 	if(ptrglSwapIntervalEXT == NULL) return 1;
// 	ptrglGetSwapIntervalEXT = goglGetProcAddress("glGetSwapIntervalEXT");
// 	if(ptrglGetSwapIntervalEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_swap_control_tear() {
// 	return 0;
// }
// int init_I3D_digital_video_control() {
// 	ptrglGetDigitalVideoParametersI3D = goglGetProcAddress("glGetDigitalVideoParametersI3D");
// 	if(ptrglGetDigitalVideoParametersI3D == NULL) return 1;
// 	ptrglSetDigitalVideoParametersI3D = goglGetProcAddress("glSetDigitalVideoParametersI3D");
// 	if(ptrglSetDigitalVideoParametersI3D == NULL) return 1;
// 	return 0;
// }
// int init_I3D_gamma() {
// 	ptrglGetGammaTableParametersI3D = goglGetProcAddress("glGetGammaTableParametersI3D");
// 	if(ptrglGetGammaTableParametersI3D == NULL) return 1;
// 	ptrglSetGammaTableParametersI3D = goglGetProcAddress("glSetGammaTableParametersI3D");
// 	if(ptrglSetGammaTableParametersI3D == NULL) return 1;
// 	ptrglGetGammaTableI3D = goglGetProcAddress("glGetGammaTableI3D");
// 	if(ptrglGetGammaTableI3D == NULL) return 1;
// 	ptrglSetGammaTableI3D = goglGetProcAddress("glSetGammaTableI3D");
// 	if(ptrglSetGammaTableI3D == NULL) return 1;
// 	return 0;
// }
// int init_I3D_genlock() {
// 	ptrglEnableGenlockI3D = goglGetProcAddress("glEnableGenlockI3D");
// 	if(ptrglEnableGenlockI3D == NULL) return 1;
// 	ptrglDisableGenlockI3D = goglGetProcAddress("glDisableGenlockI3D");
// 	if(ptrglDisableGenlockI3D == NULL) return 1;
// 	ptrglIsEnabledGenlockI3D = goglGetProcAddress("glIsEnabledGenlockI3D");
// 	if(ptrglIsEnabledGenlockI3D == NULL) return 1;
// 	ptrglGenlockSourceI3D = goglGetProcAddress("glGenlockSourceI3D");
// 	if(ptrglGenlockSourceI3D == NULL) return 1;
// 	ptrglGetGenlockSourceI3D = goglGetProcAddress("glGetGenlockSourceI3D");
// 	if(ptrglGetGenlockSourceI3D == NULL) return 1;
// 	ptrglGenlockSourceEdgeI3D = goglGetProcAddress("glGenlockSourceEdgeI3D");
// 	if(ptrglGenlockSourceEdgeI3D == NULL) return 1;
// 	ptrglGetGenlockSourceEdgeI3D = goglGetProcAddress("glGetGenlockSourceEdgeI3D");
// 	if(ptrglGetGenlockSourceEdgeI3D == NULL) return 1;
// 	ptrglGenlockSampleRateI3D = goglGetProcAddress("glGenlockSampleRateI3D");
// 	if(ptrglGenlockSampleRateI3D == NULL) return 1;
// 	ptrglGetGenlockSampleRateI3D = goglGetProcAddress("glGetGenlockSampleRateI3D");
// 	if(ptrglGetGenlockSampleRateI3D == NULL) return 1;
// 	ptrglGenlockSourceDelayI3D = goglGetProcAddress("glGenlockSourceDelayI3D");
// 	if(ptrglGenlockSourceDelayI3D == NULL) return 1;
// 	ptrglGetGenlockSourceDelayI3D = goglGetProcAddress("glGetGenlockSourceDelayI3D");
// 	if(ptrglGetGenlockSourceDelayI3D == NULL) return 1;
// 	ptrglQueryGenlockMaxSourceDelayI3D = goglGetProcAddress("glQueryGenlockMaxSourceDelayI3D");
// 	if(ptrglQueryGenlockMaxSourceDelayI3D == NULL) return 1;
// 	return 0;
// }
// int init_I3D_image_buffer() {
// 	ptrglCreateImageBufferI3D = goglGetProcAddress("glCreateImageBufferI3D");
// 	if(ptrglCreateImageBufferI3D == NULL) return 1;
// 	ptrglDestroyImageBufferI3D = goglGetProcAddress("glDestroyImageBufferI3D");
// 	if(ptrglDestroyImageBufferI3D == NULL) return 1;
// 	ptrglAssociateImageBufferEventsI3D = goglGetProcAddress("glAssociateImageBufferEventsI3D");
// 	if(ptrglAssociateImageBufferEventsI3D == NULL) return 1;
// 	ptrglReleaseImageBufferEventsI3D = goglGetProcAddress("glReleaseImageBufferEventsI3D");
// 	if(ptrglReleaseImageBufferEventsI3D == NULL) return 1;
// 	return 0;
// }
// int init_I3D_swap_frame_lock() {
// 	ptrglEnableFrameLockI3D = goglGetProcAddress("glEnableFrameLockI3D");
// 	if(ptrglEnableFrameLockI3D == NULL) return 1;
// 	ptrglDisableFrameLockI3D = goglGetProcAddress("glDisableFrameLockI3D");
// 	if(ptrglDisableFrameLockI3D == NULL) return 1;
// 	ptrglIsEnabledFrameLockI3D = goglGetProcAddress("glIsEnabledFrameLockI3D");
// 	if(ptrglIsEnabledFrameLockI3D == NULL) return 1;
// 	ptrglQueryFrameLockMasterI3D = goglGetProcAddress("glQueryFrameLockMasterI3D");
// 	if(ptrglQueryFrameLockMasterI3D == NULL) return 1;
// 	return 0;
// }
// int init_I3D_swap_frame_usage() {
// 	ptrglGetFrameUsageI3D = goglGetProcAddress("glGetFrameUsageI3D");
// 	if(ptrglGetFrameUsageI3D == NULL) return 1;
// 	ptrglBeginFrameTrackingI3D = goglGetProcAddress("glBeginFrameTrackingI3D");
// 	if(ptrglBeginFrameTrackingI3D == NULL) return 1;
// 	ptrglEndFrameTrackingI3D = goglGetProcAddress("glEndFrameTrackingI3D");
// 	if(ptrglEndFrameTrackingI3D == NULL) return 1;
// 	ptrglQueryFrameTrackingI3D = goglGetProcAddress("glQueryFrameTrackingI3D");
// 	if(ptrglQueryFrameTrackingI3D == NULL) return 1;
// 	return 0;
// }
// int init_NV_DX_interop() {
// 	ptrglDXSetResourceShareHandleNV = goglGetProcAddress("glDXSetResourceShareHandleNV");
// 	if(ptrglDXSetResourceShareHandleNV == NULL) return 1;
// 	ptrglDXOpenDeviceNV = goglGetProcAddress("glDXOpenDeviceNV");
// 	if(ptrglDXOpenDeviceNV == NULL) return 1;
// 	ptrglDXCloseDeviceNV = goglGetProcAddress("glDXCloseDeviceNV");
// 	if(ptrglDXCloseDeviceNV == NULL) return 1;
// 	ptrglDXRegisterObjectNV = goglGetProcAddress("glDXRegisterObjectNV");
// 	if(ptrglDXRegisterObjectNV == NULL) return 1;
// 	ptrglDXUnregisterObjectNV = goglGetProcAddress("glDXUnregisterObjectNV");
// 	if(ptrglDXUnregisterObjectNV == NULL) return 1;
// 	ptrglDXObjectAccessNV = goglGetProcAddress("glDXObjectAccessNV");
// 	if(ptrglDXObjectAccessNV == NULL) return 1;
// 	ptrglDXLockObjectsNV = goglGetProcAddress("glDXLockObjectsNV");
// 	if(ptrglDXLockObjectsNV == NULL) return 1;
// 	ptrglDXUnlockObjectsNV = goglGetProcAddress("glDXUnlockObjectsNV");
// 	if(ptrglDXUnlockObjectsNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_DX_interop2() {
// 	return 0;
// }
// int init_NV_copy_image() {
// 	ptrglCopyImageSubDataNV = goglGetProcAddress("glCopyImageSubDataNV");
// 	if(ptrglCopyImageSubDataNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_float_buffer() {
// 	return 0;
// }
// int init_NV_gpu_affinity() {
// 	ptrglEnumGpusNV = goglGetProcAddress("glEnumGpusNV");
// 	if(ptrglEnumGpusNV == NULL) return 1;
// 	ptrglEnumGpuDevicesNV = goglGetProcAddress("glEnumGpuDevicesNV");
// 	if(ptrglEnumGpuDevicesNV == NULL) return 1;
// 	ptrglCreateAffinityDCNV = goglGetProcAddress("glCreateAffinityDCNV");
// 	if(ptrglCreateAffinityDCNV == NULL) return 1;
// 	ptrglEnumGpusFromAffinityDCNV = goglGetProcAddress("glEnumGpusFromAffinityDCNV");
// 	if(ptrglEnumGpusFromAffinityDCNV == NULL) return 1;
// 	ptrglDeleteDCNV = goglGetProcAddress("glDeleteDCNV");
// 	if(ptrglDeleteDCNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_multisample_coverage() {
// 	return 0;
// }
// int init_NV_present_video() {
// 	ptrglEnumerateVideoDevicesNV = goglGetProcAddress("glEnumerateVideoDevicesNV");
// 	if(ptrglEnumerateVideoDevicesNV == NULL) return 1;
// 	ptrglBindVideoDeviceNV = goglGetProcAddress("glBindVideoDeviceNV");
// 	if(ptrglBindVideoDeviceNV == NULL) return 1;
// 	ptrglQueryCurrentContextNV = goglGetProcAddress("glQueryCurrentContextNV");
// 	if(ptrglQueryCurrentContextNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_swap_group() {
// 	ptrglJoinSwapGroupNV = goglGetProcAddress("glJoinSwapGroupNV");
// 	if(ptrglJoinSwapGroupNV == NULL) return 1;
// 	ptrglBindSwapBarrierNV = goglGetProcAddress("glBindSwapBarrierNV");
// 	if(ptrglBindSwapBarrierNV == NULL) return 1;
// 	ptrglQuerySwapGroupNV = goglGetProcAddress("glQuerySwapGroupNV");
// 	if(ptrglQuerySwapGroupNV == NULL) return 1;
// 	ptrglQueryMaxSwapGroupsNV = goglGetProcAddress("glQueryMaxSwapGroupsNV");
// 	if(ptrglQueryMaxSwapGroupsNV == NULL) return 1;
// 	ptrglQueryFrameCountNV = goglGetProcAddress("glQueryFrameCountNV");
// 	if(ptrglQueryFrameCountNV == NULL) return 1;
// 	ptrglResetFrameCountNV = goglGetProcAddress("glResetFrameCountNV");
// 	if(ptrglResetFrameCountNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_vertex_array_range() {
// 	ptrglAllocateMemoryNV = goglGetProcAddress("glAllocateMemoryNV");
// 	if(ptrglAllocateMemoryNV == NULL) return 1;
// 	ptrglFreeMemoryNV = goglGetProcAddress("glFreeMemoryNV");
// 	if(ptrglFreeMemoryNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_video_capture() {
// 	ptrglBindVideoCaptureDeviceNV = goglGetProcAddress("glBindVideoCaptureDeviceNV");
// 	if(ptrglBindVideoCaptureDeviceNV == NULL) return 1;
// 	ptrglEnumerateVideoCaptureDevicesNV = goglGetProcAddress("glEnumerateVideoCaptureDevicesNV");
// 	if(ptrglEnumerateVideoCaptureDevicesNV == NULL) return 1;
// 	ptrglLockVideoCaptureDeviceNV = goglGetProcAddress("glLockVideoCaptureDeviceNV");
// 	if(ptrglLockVideoCaptureDeviceNV == NULL) return 1;
// 	ptrglQueryVideoCaptureDeviceNV = goglGetProcAddress("glQueryVideoCaptureDeviceNV");
// 	if(ptrglQueryVideoCaptureDeviceNV == NULL) return 1;
// 	ptrglReleaseVideoCaptureDeviceNV = goglGetProcAddress("glReleaseVideoCaptureDeviceNV");
// 	if(ptrglReleaseVideoCaptureDeviceNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_video_output() {
// 	ptrglGetVideoDeviceNV = goglGetProcAddress("glGetVideoDeviceNV");
// 	if(ptrglGetVideoDeviceNV == NULL) return 1;
// 	ptrglReleaseVideoDeviceNV = goglGetProcAddress("glReleaseVideoDeviceNV");
// 	if(ptrglReleaseVideoDeviceNV == NULL) return 1;
// 	ptrglBindVideoImageNV = goglGetProcAddress("glBindVideoImageNV");
// 	if(ptrglBindVideoImageNV == NULL) return 1;
// 	ptrglReleaseVideoImageNV = goglGetProcAddress("glReleaseVideoImageNV");
// 	if(ptrglReleaseVideoImageNV == NULL) return 1;
// 	ptrglSendPbufferToVideoNV = goglGetProcAddress("glSendPbufferToVideoNV");
// 	if(ptrglSendPbufferToVideoNV == NULL) return 1;
// 	ptrglGetVideoInfoNV = goglGetProcAddress("glGetVideoInfoNV");
// 	if(ptrglGetVideoInfoNV == NULL) return 1;
// 	return 0;
// }
// int init_OML_sync_control() {
// 	ptrglGetSyncValuesOML = goglGetProcAddress("glGetSyncValuesOML");
// 	if(ptrglGetSyncValuesOML == NULL) return 1;
// 	ptrglGetMscRateOML = goglGetProcAddress("glGetMscRateOML");
// 	if(ptrglGetMscRateOML == NULL) return 1;
// 	ptrglSwapBuffersMscOML = goglGetProcAddress("glSwapBuffersMscOML");
// 	if(ptrglSwapBuffersMscOML == NULL) return 1;
// 	ptrglSwapLayerBuffersMscOML = goglGetProcAddress("glSwapLayerBuffersMscOML");
// 	if(ptrglSwapLayerBuffersMscOML == NULL) return 1;
// 	ptrglWaitForMscOML = goglGetProcAddress("glWaitForMscOML");
// 	if(ptrglWaitForMscOML == NULL) return 1;
// 	ptrglWaitForSbcOML = goglGetProcAddress("glWaitForSbcOML");
// 	if(ptrglWaitForSbcOML == NULL) return 1;
// 	return 0;
// }
// int init_wgl() {
// 	ptrglCreateContext = goglGetProcAddress("glCreateContext");
// 	if(ptrglCreateContext == NULL) return 1;
// 	ptrglDeleteContext = goglGetProcAddress("glDeleteContext");
// 	if(ptrglDeleteContext == NULL) return 1;
// 	ptrglGetCurrentContext = goglGetProcAddress("glGetCurrentContext");
// 	if(ptrglGetCurrentContext == NULL) return 1;
// 	ptrglMakeCurrent = goglGetProcAddress("glMakeCurrent");
// 	if(ptrglMakeCurrent == NULL) return 1;
// 	ptrglCopyContext = goglGetProcAddress("glCopyContext");
// 	if(ptrglCopyContext == NULL) return 1;
// 	ptrglChoosePixelFormat = goglGetProcAddress("glChoosePixelFormat");
// 	if(ptrglChoosePixelFormat == NULL) return 1;
// 	ptrglGetCurrentDC = goglGetProcAddress("glGetCurrentDC");
// 	if(ptrglGetCurrentDC == NULL) return 1;
// 	ptrglGetDefaultProcAddress = goglGetProcAddress("glGetDefaultProcAddress");
// 	if(ptrglGetDefaultProcAddress == NULL) return 1;
// 	ptrglGetProcAddress = goglGetProcAddress("glGetProcAddress");
// 	if(ptrglGetProcAddress == NULL) return 1;
// 	ptrglGetPixelFormat = goglGetProcAddress("glGetPixelFormat");
// 	if(ptrglGetPixelFormat == NULL) return 1;
// 	ptrglSetPixelFormat = goglGetProcAddress("glSetPixelFormat");
// 	if(ptrglSetPixelFormat == NULL) return 1;
// 	ptrglSwapBuffers = goglGetProcAddress("glSwapBuffers");
// 	if(ptrglSwapBuffers == NULL) return 1;
// 	ptrglShareLists = goglGetProcAddress("glShareLists");
// 	if(ptrglShareLists == NULL) return 1;
// 	ptrglCreateLayerContext = goglGetProcAddress("glCreateLayerContext");
// 	if(ptrglCreateLayerContext == NULL) return 1;
// 	ptrglDescribeLayerPlane = goglGetProcAddress("glDescribeLayerPlane");
// 	if(ptrglDescribeLayerPlane == NULL) return 1;
// 	ptrglSetLayerPaletteEntries = goglGetProcAddress("glSetLayerPaletteEntries");
// 	if(ptrglSetLayerPaletteEntries == NULL) return 1;
// 	ptrglGetLayerPaletteEntries = goglGetProcAddress("glGetLayerPaletteEntries");
// 	if(ptrglGetLayerPaletteEntries == NULL) return 1;
// 	ptrglRealizeLayerPalette = goglGetProcAddress("glRealizeLayerPalette");
// 	if(ptrglRealizeLayerPalette == NULL) return 1;
// 	ptrglSwapLayerBuffers = goglGetProcAddress("glSwapLayerBuffers");
// 	if(ptrglSwapLayerBuffers == NULL) return 1;
// 	ptrglUseFontBitmapsA = goglGetProcAddress("glUseFontBitmapsA");
// 	if(ptrglUseFontBitmapsA == NULL) return 1;
// 	ptrglUseFontBitmapsW = goglGetProcAddress("glUseFontBitmapsW");
// 	if(ptrglUseFontBitmapsW == NULL) return 1;
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
// 3DFX_multisample

// 3DL_stereo_control

func SetStereoEmitterState3DL(hDC Pointer, uState uint32) int32 {
	return (int32)(C.goglSetStereoEmitterState3DL((C.HDC)(hDC), (C.UINT)(uState)))
}
// AMD_gpu_association

func GetGPUIDsAMD(maxCount uint32, ids *uint32) uint32 {
	return (uint32)(C.goglGetGPUIDsAMD((C.UINT)(maxCount), (*C.UINT)(ids)))
}
func GetGPUInfoAMD(id uint32, property Int, dataType Enum, size uint32, data Pointer) int32 {
	return (int32)(C.goglGetGPUInfoAMD((C.UINT)(id), (C.int)(property), (C.GLenum)(dataType), (C.UINT)(size), (unsafe.Pointer)(data)))
}
func GetContextGPUIDAMD(hglrc Pointer) uint32 {
	return (uint32)(C.goglGetContextGPUIDAMD((C.HGLRC)(hglrc)))
}
func CreateAssociatedContextAMD(id uint32) Pointer {
	return (Pointer)(C.goglCreateAssociatedContextAMD((C.UINT)(id)))
}
func CreateAssociatedContextAttribsAMD(id uint32, hShareContext Pointer, attribList *Int) Pointer {
	return (Pointer)(C.goglCreateAssociatedContextAttribsAMD((C.UINT)(id), (C.HGLRC)(hShareContext), (*C.int)(attribList)))
}
func DeleteAssociatedContextAMD(hglrc Pointer) int32 {
	return (int32)(C.goglDeleteAssociatedContextAMD((C.HGLRC)(hglrc)))
}
func MakeAssociatedContextCurrentAMD(hglrc Pointer) int32 {
	return (int32)(C.goglMakeAssociatedContextCurrentAMD((C.HGLRC)(hglrc)))
}
func GetCurrentAssociatedContextAMD() Pointer {
	return (Pointer)(C.goglGetCurrentAssociatedContextAMD())
}
func BlitContextFramebufferAMD(dstCtx Pointer, srcX0 Int, srcY0 Int, srcX1 Int, srcY1 Int, dstX0 Int, dstY0 Int, dstX1 Int, dstY1 Int, mask Bitfield, filter Enum)  {
	return ()(C.goglBlitContextFramebufferAMD((C.HGLRC)(dstCtx), (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter)))
}
// ARB_buffer_region

func CreateBufferRegionARB(hDC Pointer, iLayerPlane Int, uType uint32) Pointer {
	return (Pointer)(C.goglCreateBufferRegionARB((C.HDC)(hDC), (C.int)(iLayerPlane), (C.UINT)(uType)))
}
func DeleteBufferRegionARB(hRegion Pointer)  {
	return ()(C.goglDeleteBufferRegionARB((C.HANDLE)(hRegion)))
}
func SaveBufferRegionARB(hRegion Pointer, x Int, y Int, width Int, height Int) int32 {
	return (int32)(C.goglSaveBufferRegionARB((C.HANDLE)(hRegion), (C.int)(x), (C.int)(y), (C.int)(width), (C.int)(height)))
}
func RestoreBufferRegionARB(hRegion Pointer, x Int, y Int, width Int, height Int, xSrc Int, ySrc Int) int32 {
	return (int32)(C.goglRestoreBufferRegionARB((C.HANDLE)(hRegion), (C.int)(x), (C.int)(y), (C.int)(width), (C.int)(height), (C.int)(xSrc), (C.int)(ySrc)))
}
// ARB_create_context

func CreateContextAttribsARB(hDC Pointer, hShareContext Pointer, attribList *Int) Pointer {
	return (Pointer)(C.goglCreateContextAttribsARB((C.HDC)(hDC), (C.HGLRC)(hShareContext), (*C.int)(attribList)))
}
// ARB_create_context_profile

// ARB_create_context_robustness

// ARB_extensions_string

func GetExtensionsStringARB(hdc Pointer) *byte {
	return (*byte)(C.goglGetExtensionsStringARB((C.HDC)(hdc)))
}
// ARB_framebuffer_sRGB

// ARB_make_current_read

func MakeContextCurrentARB(hDrawDC Pointer, hReadDC Pointer, hglrc Pointer) int32 {
	return (int32)(C.goglMakeContextCurrentARB((C.HDC)(hDrawDC), (C.HDC)(hReadDC), (C.HGLRC)(hglrc)))
}
func GetCurrentReadDCARB() Pointer {
	return (Pointer)(C.goglGetCurrentReadDCARB())
}
// ARB_multisample

// ARB_pbuffer

func CreatePbufferARB(hDC Pointer, iPixelFormat Int, iWidth Int, iHeight Int, piAttribList *Int) Pointer {
	return (Pointer)(C.goglCreatePbufferARB((C.HDC)(hDC), (C.int)(iPixelFormat), (C.int)(iWidth), (C.int)(iHeight), (*C.int)(piAttribList)))
}
func GetPbufferDCARB(hPbuffer Pointer) Pointer {
	return (Pointer)(C.goglGetPbufferDCARB((C.HPBUFFERARB)(hPbuffer)))
}
func ReleasePbufferDCARB(hPbuffer Pointer, hDC Pointer) Int {
	return (Int)(C.goglReleasePbufferDCARB((C.HPBUFFERARB)(hPbuffer), (C.HDC)(hDC)))
}
func DestroyPbufferARB(hPbuffer Pointer) int32 {
	return (int32)(C.goglDestroyPbufferARB((C.HPBUFFERARB)(hPbuffer)))
}
func QueryPbufferARB(hPbuffer Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.goglQueryPbufferARB((C.HPBUFFERARB)(hPbuffer), (C.int)(iAttribute), (*C.int)(piValue)))
}
// ARB_pixel_format

func GetPixelFormatAttribivARB(hdc Pointer, iPixelFormat Int, iLayerPlane Int, nAttributes uint32, piAttributes *Int, piValues *Int) int32 {
	return (int32)(C.goglGetPixelFormatAttribivARB((C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(piAttributes), (*C.int)(piValues)))
}
func GetPixelFormatAttribfvARB(hdc Pointer, iPixelFormat Int, iLayerPlane Int, nAttributes uint32, piAttributes *Int, pfValues *float32) int32 {
	return (int32)(C.goglGetPixelFormatAttribfvARB((C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(piAttributes), (*C.FLOAT)(pfValues)))
}
func ChoosePixelFormatARB(hdc Pointer, piAttribIList *Int, pfAttribFList *float32, nMaxFormats uint32, piFormats *Int, nNumFormats *uint32) int32 {
	return (int32)(C.goglChoosePixelFormatARB((C.HDC)(hdc), (*C.int)(piAttribIList), (*C.FLOAT)(pfAttribFList), (C.UINT)(nMaxFormats), (*C.int)(piFormats), (*C.UINT)(nNumFormats)))
}
// ARB_pixel_format_float

// ARB_render_texture

func BindTexImageARB(hPbuffer Pointer, iBuffer Int) int32 {
	return (int32)(C.goglBindTexImageARB((C.HPBUFFERARB)(hPbuffer), (C.int)(iBuffer)))
}
func ReleaseTexImageARB(hPbuffer Pointer, iBuffer Int) int32 {
	return (int32)(C.goglReleaseTexImageARB((C.HPBUFFERARB)(hPbuffer), (C.int)(iBuffer)))
}
func SetPbufferAttribARB(hPbuffer Pointer, piAttribList *Int) int32 {
	return (int32)(C.goglSetPbufferAttribARB((C.HPBUFFERARB)(hPbuffer), (*C.int)(piAttribList)))
}
// ATI_pixel_format_float

// EXT_depth_float

// EXT_display_color_table

func CreateDisplayColorTableEXT(id Ushort) Boolean {
	return (Boolean)(C.goglCreateDisplayColorTableEXT((C.GLushort)(id)))
}
func LoadDisplayColorTableEXT(table *Ushort, length Uint) Boolean {
	return (Boolean)(C.goglLoadDisplayColorTableEXT((*C.GLushort)(table), (C.GLuint)(length)))
}
func BindDisplayColorTableEXT(id Ushort) Boolean {
	return (Boolean)(C.goglBindDisplayColorTableEXT((C.GLushort)(id)))
}
func DestroyDisplayColorTableEXT(id Ushort)  {
	return ()(C.goglDestroyDisplayColorTableEXT((C.GLushort)(id)))
}
// EXT_extensions_string

func GetExtensionsStringEXT() *byte {
	return (*byte)(C.goglGetExtensionsStringEXT())
}
// EXT_framebuffer_sRGB

// EXT_make_current_read

func MakeContextCurrentEXT(hDrawDC Pointer, hReadDC Pointer, hglrc Pointer) int32 {
	return (int32)(C.goglMakeContextCurrentEXT((C.HDC)(hDrawDC), (C.HDC)(hReadDC), (C.HGLRC)(hglrc)))
}
func GetCurrentReadDCEXT() Pointer {
	return (Pointer)(C.goglGetCurrentReadDCEXT())
}
// EXT_multisample

// EXT_pbuffer

func CreatePbufferEXT(hDC Pointer, iPixelFormat Int, iWidth Int, iHeight Int, piAttribList *Int) Pointer {
	return (Pointer)(C.goglCreatePbufferEXT((C.HDC)(hDC), (C.int)(iPixelFormat), (C.int)(iWidth), (C.int)(iHeight), (*C.int)(piAttribList)))
}
func GetPbufferDCEXT(hPbuffer Pointer) Pointer {
	return (Pointer)(C.goglGetPbufferDCEXT((C.HPBUFFEREXT)(hPbuffer)))
}
func ReleasePbufferDCEXT(hPbuffer Pointer, hDC Pointer) Int {
	return (Int)(C.goglReleasePbufferDCEXT((C.HPBUFFEREXT)(hPbuffer), (C.HDC)(hDC)))
}
func DestroyPbufferEXT(hPbuffer Pointer) int32 {
	return (int32)(C.goglDestroyPbufferEXT((C.HPBUFFEREXT)(hPbuffer)))
}
func QueryPbufferEXT(hPbuffer Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.goglQueryPbufferEXT((C.HPBUFFEREXT)(hPbuffer), (C.int)(iAttribute), (*C.int)(piValue)))
}
// EXT_pixel_format

func GetPixelFormatAttribivEXT(hdc Pointer, iPixelFormat Int, iLayerPlane Int, nAttributes uint32, piAttributes *Int, piValues *Int) int32 {
	return (int32)(C.goglGetPixelFormatAttribivEXT((C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(piAttributes), (*C.int)(piValues)))
}
func GetPixelFormatAttribfvEXT(hdc Pointer, iPixelFormat Int, iLayerPlane Int, nAttributes uint32, piAttributes *Int, pfValues *float32) int32 {
	return (int32)(C.goglGetPixelFormatAttribfvEXT((C.HDC)(hdc), (C.int)(iPixelFormat), (C.int)(iLayerPlane), (C.UINT)(nAttributes), (*C.int)(piAttributes), (*C.FLOAT)(pfValues)))
}
func ChoosePixelFormatEXT(hdc Pointer, piAttribIList *Int, pfAttribFList *float32, nMaxFormats uint32, piFormats *Int, nNumFormats *uint32) int32 {
	return (int32)(C.goglChoosePixelFormatEXT((C.HDC)(hdc), (*C.int)(piAttribIList), (*C.FLOAT)(pfAttribFList), (C.UINT)(nMaxFormats), (*C.int)(piFormats), (*C.UINT)(nNumFormats)))
}
// EXT_pixel_format_packed_float

// EXT_swap_control

func SwapIntervalEXT(interval Int) int32 {
	return (int32)(C.goglSwapIntervalEXT((C.int)(interval)))
}
func GetSwapIntervalEXT() Int {
	return (Int)(C.goglGetSwapIntervalEXT())
}
// EXT_swap_control_tear

// I3D_digital_video_control

func GetDigitalVideoParametersI3D(hDC Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.goglGetDigitalVideoParametersI3D((C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(piValue)))
}
func SetDigitalVideoParametersI3D(hDC Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.goglSetDigitalVideoParametersI3D((C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(piValue)))
}
// I3D_gamma

func GetGammaTableParametersI3D(hDC Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.goglGetGammaTableParametersI3D((C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(piValue)))
}
func SetGammaTableParametersI3D(hDC Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.goglSetGammaTableParametersI3D((C.HDC)(hDC), (C.int)(iAttribute), (*C.int)(piValue)))
}
func GetGammaTableI3D(hDC Pointer, iEntries Int, puRed *uint16, puGreen *uint16, puBlue *uint16) int32 {
	return (int32)(C.goglGetGammaTableI3D((C.HDC)(hDC), (C.int)(iEntries), (*C.USHORT)(puRed), (*C.USHORT)(puGreen), (*C.USHORT)(puBlue)))
}
func SetGammaTableI3D(hDC Pointer, iEntries Int, puRed *uint16, puGreen *uint16, puBlue *uint16) int32 {
	return (int32)(C.goglSetGammaTableI3D((C.HDC)(hDC), (C.int)(iEntries), (*C.USHORT)(puRed), (*C.USHORT)(puGreen), (*C.USHORT)(puBlue)))
}
// I3D_genlock

func EnableGenlockI3D(hDC Pointer) int32 {
	return (int32)(C.goglEnableGenlockI3D((C.HDC)(hDC)))
}
func DisableGenlockI3D(hDC Pointer) int32 {
	return (int32)(C.goglDisableGenlockI3D((C.HDC)(hDC)))
}
func IsEnabledGenlockI3D(hDC Pointer, pFlag *int32) int32 {
	return (int32)(C.goglIsEnabledGenlockI3D((C.HDC)(hDC), (*C.BOOL)(pFlag)))
}
func GenlockSourceI3D(hDC Pointer, uSource uint32) int32 {
	return (int32)(C.goglGenlockSourceI3D((C.HDC)(hDC), (C.UINT)(uSource)))
}
func GetGenlockSourceI3D(hDC Pointer, uSource *uint32) int32 {
	return (int32)(C.goglGetGenlockSourceI3D((C.HDC)(hDC), (*C.UINT)(uSource)))
}
func GenlockSourceEdgeI3D(hDC Pointer, uEdge uint32) int32 {
	return (int32)(C.goglGenlockSourceEdgeI3D((C.HDC)(hDC), (C.UINT)(uEdge)))
}
func GetGenlockSourceEdgeI3D(hDC Pointer, uEdge *uint32) int32 {
	return (int32)(C.goglGetGenlockSourceEdgeI3D((C.HDC)(hDC), (*C.UINT)(uEdge)))
}
func GenlockSampleRateI3D(hDC Pointer, uRate uint32) int32 {
	return (int32)(C.goglGenlockSampleRateI3D((C.HDC)(hDC), (C.UINT)(uRate)))
}
func GetGenlockSampleRateI3D(hDC Pointer, uRate *uint32) int32 {
	return (int32)(C.goglGetGenlockSampleRateI3D((C.HDC)(hDC), (*C.UINT)(uRate)))
}
func GenlockSourceDelayI3D(hDC Pointer, uDelay uint32) int32 {
	return (int32)(C.goglGenlockSourceDelayI3D((C.HDC)(hDC), (C.UINT)(uDelay)))
}
func GetGenlockSourceDelayI3D(hDC Pointer, uDelay *uint32) int32 {
	return (int32)(C.goglGetGenlockSourceDelayI3D((C.HDC)(hDC), (*C.UINT)(uDelay)))
}
func QueryGenlockMaxSourceDelayI3D(hDC Pointer, uMaxLineDelay *uint32, uMaxPixelDelay *uint32) int32 {
	return (int32)(C.goglQueryGenlockMaxSourceDelayI3D((C.HDC)(hDC), (*C.UINT)(uMaxLineDelay), (*C.UINT)(uMaxPixelDelay)))
}
// I3D_image_buffer

func CreateImageBufferI3D(hDC Pointer, dwSize uint32, uFlags uint32) Pointer {
	return (Pointer)(C.goglCreateImageBufferI3D((C.HDC)(hDC), (C.DWORD)(dwSize), (C.UINT)(uFlags)))
}
func DestroyImageBufferI3D(hDC Pointer, pAddress Pointer) int32 {
	return (int32)(C.goglDestroyImageBufferI3D((C.HDC)(hDC), (C.LPVOID)(pAddress)))
}
func AssociateImageBufferEventsI3D(hDC Pointer, pEvent Pointer, pAddress Pointer, pSize *uint32, count uint32) int32 {
	return (int32)(C.goglAssociateImageBufferEventsI3D((C.HDC)(hDC), (C.HANDLE)(pEvent), (C.LPVOID)(pAddress), (*C.DWORD)(pSize), (C.UINT)(count)))
}
func ReleaseImageBufferEventsI3D(hDC Pointer, pAddress Pointer, count uint32) int32 {
	return (int32)(C.goglReleaseImageBufferEventsI3D((C.HDC)(hDC), (C.LPVOID)(pAddress), (C.UINT)(count)))
}
// I3D_swap_frame_lock

func EnableFrameLockI3D() int32 {
	return (int32)(C.goglEnableFrameLockI3D())
}
func DisableFrameLockI3D() int32 {
	return (int32)(C.goglDisableFrameLockI3D())
}
func IsEnabledFrameLockI3D(pFlag *int32) int32 {
	return (int32)(C.goglIsEnabledFrameLockI3D((*C.BOOL)(pFlag)))
}
func QueryFrameLockMasterI3D(pFlag *int32) int32 {
	return (int32)(C.goglQueryFrameLockMasterI3D((*C.BOOL)(pFlag)))
}
// I3D_swap_frame_usage

func GetFrameUsageI3D(pUsage *float32) int32 {
	return (int32)(C.goglGetFrameUsageI3D((*C.float32)(pUsage)))
}
func BeginFrameTrackingI3D() int32 {
	return (int32)(C.goglBeginFrameTrackingI3D())
}
func EndFrameTrackingI3D() int32 {
	return (int32)(C.goglEndFrameTrackingI3D())
}
func QueryFrameTrackingI3D(pFrameCount *uint32, pMissedFrames *uint32, pLastMissedUsage *float32) int32 {
	return (int32)(C.goglQueryFrameTrackingI3D((*C.DWORD)(pFrameCount), (*C.DWORD)(pMissedFrames), (*C.float32)(pLastMissedUsage)))
}
// NV_DX_interop

func DXSetResourceShareHandleNV(dxObject Pointer, shareHandle Pointer) int32 {
	return (int32)(C.goglDXSetResourceShareHandleNV((unsafe.Pointer)(dxObject), (C.HANDLE)(shareHandle)))
}
func DXOpenDeviceNV(dxDevice Pointer) Pointer {
	return (Pointer)(C.goglDXOpenDeviceNV((unsafe.Pointer)(dxDevice)))
}
func DXCloseDeviceNV(hDevice Pointer) int32 {
	return (int32)(C.goglDXCloseDeviceNV((C.HANDLE)(hDevice)))
}
func DXRegisterObjectNV(hDevice Pointer, dxObject Pointer, name Uint, type_ Enum, access Enum) Pointer {
	return (Pointer)(C.goglDXRegisterObjectNV((C.HANDLE)(hDevice), (unsafe.Pointer)(dxObject), (C.GLuint)(name), (C.GLenum)(type_), (C.GLenum)(access)))
}
func DXUnregisterObjectNV(hDevice Pointer, hObject Pointer) int32 {
	return (int32)(C.goglDXUnregisterObjectNV((C.HANDLE)(hDevice), (C.HANDLE)(hObject)))
}
func DXObjectAccessNV(hObject Pointer, access Enum) int32 {
	return (int32)(C.goglDXObjectAccessNV((C.HANDLE)(hObject), (C.GLenum)(access)))
}
func DXLockObjectsNV(hDevice Pointer, count Int, hObjects Pointer) int32 {
	return (int32)(C.goglDXLockObjectsNV((C.HANDLE)(hDevice), (C.GLint)(count), (C.HANDLE)(hObjects)))
}
func DXUnlockObjectsNV(hDevice Pointer, count Int, hObjects Pointer) int32 {
	return (int32)(C.goglDXUnlockObjectsNV((C.HANDLE)(hDevice), (C.GLint)(count), (C.HANDLE)(hObjects)))
}
// NV_DX_interop2

// NV_copy_image

func CopyImageSubDataNV(hSrcRC Pointer, srcName Uint, srcTarget Enum, srcLevel Int, srcX Int, srcY Int, srcZ Int, hDstRC Pointer, dstName Uint, dstTarget Enum, dstLevel Int, dstX Int, dstY Int, dstZ Int, width Sizei, height Sizei, depth Sizei) int32 {
	return (int32)(C.goglCopyImageSubDataNV((C.HGLRC)(hSrcRC), (C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.HGLRC)(hDstRC), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth)))
}
// NV_float_buffer

// NV_gpu_affinity

func EnumGpusNV(iGpuIndex uint32, phGpu Pointer) int32 {
	return (int32)(C.goglEnumGpusNV((C.UINT)(iGpuIndex), (C.HGPUNV)(phGpu)))
}
func EnumGpuDevicesNV(hGpu Pointer, iDeviceIndex uint32, lpGpuDevice Pointer) int32 {
	return (int32)(C.goglEnumGpuDevicesNV((C.HGPUNV)(hGpu), (C.UINT)(iDeviceIndex), (C.PGPU_DEVICE)(lpGpuDevice)))
}
func CreateAffinityDCNV(phGpuList Pointer) Pointer {
	return (Pointer)(C.goglCreateAffinityDCNV((C.HGPUNV)(phGpuList)))
}
func EnumGpusFromAffinityDCNV(hAffinityDC Pointer, iGpuIndex uint32, hGpu Pointer) int32 {
	return (int32)(C.goglEnumGpusFromAffinityDCNV((C.HDC)(hAffinityDC), (C.UINT)(iGpuIndex), (C.HGPUNV)(hGpu)))
}
func DeleteDCNV(hdc Pointer) int32 {
	return (int32)(C.goglDeleteDCNV((C.HDC)(hdc)))
}
// NV_multisample_coverage

// NV_present_video

func EnumerateVideoDevicesNV(hDC Pointer, phDeviceList Pointer) Int {
	return (Int)(C.goglEnumerateVideoDevicesNV((C.HDC)(hDC), (C.HVIDEOOUTPUTDEVICENV)(phDeviceList)))
}
func BindVideoDeviceNV(hDC Pointer, uVideoSlot uint32, hVideoDevice Pointer, piAttribList *Int) int32 {
	return (int32)(C.goglBindVideoDeviceNV((C.HDC)(hDC), (C.uint32)(uVideoSlot), (C.HVIDEOOUTPUTDEVICENV)(hVideoDevice), (*C.int)(piAttribList)))
}
func QueryCurrentContextNV(iAttribute Int, piValue *Int) int32 {
	return (int32)(C.goglQueryCurrentContextNV((C.int)(iAttribute), (*C.int)(piValue)))
}
// NV_swap_group

func JoinSwapGroupNV(hDC Pointer, group Uint) int32 {
	return (int32)(C.goglJoinSwapGroupNV((C.HDC)(hDC), (C.GLuint)(group)))
}
func BindSwapBarrierNV(group Uint, barrier Uint) int32 {
	return (int32)(C.goglBindSwapBarrierNV((C.GLuint)(group), (C.GLuint)(barrier)))
}
func QuerySwapGroupNV(hDC Pointer, group *Uint, barrier *Uint) int32 {
	return (int32)(C.goglQuerySwapGroupNV((C.HDC)(hDC), (*C.GLuint)(group), (*C.GLuint)(barrier)))
}
func QueryMaxSwapGroupsNV(hDC Pointer, maxGroups *Uint, maxBarriers *Uint) int32 {
	return (int32)(C.goglQueryMaxSwapGroupsNV((C.HDC)(hDC), (*C.GLuint)(maxGroups), (*C.GLuint)(maxBarriers)))
}
func QueryFrameCountNV(hDC Pointer, count *Uint) int32 {
	return (int32)(C.goglQueryFrameCountNV((C.HDC)(hDC), (*C.GLuint)(count)))
}
func ResetFrameCountNV(hDC Pointer) int32 {
	return (int32)(C.goglResetFrameCountNV((C.HDC)(hDC)))
}
// NV_vertex_array_range

func AllocateMemoryNV(size Sizei, readfreq Float, writefreq Float, priority Float) Pointer {
	return (Pointer)(C.goglAllocateMemoryNV((C.GLsizei)(size), (C.GLfloat)(readfreq), (C.GLfloat)(writefreq), (C.GLfloat)(priority)))
}
func FreeMemoryNV(pointer Pointer)  {
	C.goglFreeMemoryNV((unsafe.Pointer)(pointer))
}
// NV_video_capture

func BindVideoCaptureDeviceNV(uVideoSlot uint32, hDevice Pointer) int32 {
	return (int32)(C.goglBindVideoCaptureDeviceNV((C.UINT)(uVideoSlot), (C.HVIDEOINPUTDEVICENV)(hDevice)))
}
func EnumerateVideoCaptureDevicesNV(hDc Pointer, phDeviceList Pointer) uint32 {
	return (uint32)(C.goglEnumerateVideoCaptureDevicesNV((C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(phDeviceList)))
}
func LockVideoCaptureDeviceNV(hDc Pointer, hDevice Pointer) int32 {
	return (int32)(C.goglLockVideoCaptureDeviceNV((C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(hDevice)))
}
func QueryVideoCaptureDeviceNV(hDc Pointer, hDevice Pointer, iAttribute Int, piValue *Int) int32 {
	return (int32)(C.goglQueryVideoCaptureDeviceNV((C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(hDevice), (C.int)(iAttribute), (*C.int)(piValue)))
}
func ReleaseVideoCaptureDeviceNV(hDc Pointer, hDevice Pointer) int32 {
	return (int32)(C.goglReleaseVideoCaptureDeviceNV((C.HDC)(hDc), (C.HVIDEOINPUTDEVICENV)(hDevice)))
}
// NV_video_output

func GetVideoDeviceNV(hDC Pointer, numDevices Int, hVideoDevice Pointer) int32 {
	return (int32)(C.goglGetVideoDeviceNV((C.HDC)(hDC), (C.int)(numDevices), (C.HPVIDEODEV)(hVideoDevice)))
}
func ReleaseVideoDeviceNV(hVideoDevice Pointer) int32 {
	return (int32)(C.goglReleaseVideoDeviceNV((C.HPVIDEODEV)(hVideoDevice)))
}
func BindVideoImageNV(hVideoDevice Pointer, hPbuffer Pointer, iVideoBuffer Int) int32 {
	return (int32)(C.goglBindVideoImageNV((C.HPVIDEODEV)(hVideoDevice), (C.HPBUFFERARB)(hPbuffer), (C.int)(iVideoBuffer)))
}
func ReleaseVideoImageNV(hPbuffer Pointer, iVideoBuffer Int) int32 {
	return (int32)(C.goglReleaseVideoImageNV((C.HPBUFFERARB)(hPbuffer), (C.int)(iVideoBuffer)))
}
func SendPbufferToVideoNV(hPbuffer Pointer, iBufferType Int, pulCounterPbuffer *uint32, bBlock int32) int32 {
	return (int32)(C.goglSendPbufferToVideoNV((C.HPBUFFERARB)(hPbuffer), (C.int)(iBufferType), (*C.unsigned_long)(pulCounterPbuffer), (C.BOOL)(bBlock)))
}
func GetVideoInfoNV(hpVideoDevice Pointer, pulCounterOutputPbuffer *uint32, pulCounterOutputVideo *uint32) int32 {
	return (int32)(C.goglGetVideoInfoNV((C.HPVIDEODEV)(hpVideoDevice), (*C.unsigned_long)(pulCounterOutputPbuffer), (*C.unsigned_long)(pulCounterOutputVideo)))
}
// OML_sync_control

func GetSyncValuesOML(hdc Pointer, ust *int64, msc *int64, sbc *int64) int32 {
	return (int32)(C.goglGetSyncValuesOML((C.HDC)(hdc), (*C.INT64)(ust), (*C.INT64)(msc), (*C.INT64)(sbc)))
}
func GetMscRateOML(hdc Pointer, numerator *int32, denominator *int32) int32 {
	return (int32)(C.goglGetMscRateOML((C.HDC)(hdc), (*C.INT32)(numerator), (*C.INT32)(denominator)))
}
func SwapBuffersMscOML(hdc Pointer, target_msc int64, divisor int64, remainder int64) int64 {
	return (int64)(C.goglSwapBuffersMscOML((C.HDC)(hdc), (C.INT64)(target_msc), (C.INT64)(divisor), (C.INT64)(remainder)))
}
func SwapLayerBuffersMscOML(hdc Pointer, fuPlanes Int, target_msc int64, divisor int64, remainder int64) int64 {
	return (int64)(C.goglSwapLayerBuffersMscOML((C.HDC)(hdc), (C.int)(fuPlanes), (C.INT64)(target_msc), (C.INT64)(divisor), (C.INT64)(remainder)))
}
func WaitForMscOML(hdc Pointer, target_msc int64, divisor int64, remainder int64, ust *int64, msc *int64, sbc *int64) int32 {
	return (int32)(C.goglWaitForMscOML((C.HDC)(hdc), (C.INT64)(target_msc), (C.INT64)(divisor), (C.INT64)(remainder), (*C.INT64)(ust), (*C.INT64)(msc), (*C.INT64)(sbc)))
}
func WaitForSbcOML(hdc Pointer, target_sbc int64, ust *int64, msc *int64, sbc *int64) int32 {
	return (int32)(C.goglWaitForSbcOML((C.HDC)(hdc), (C.INT64)(target_sbc), (*C.INT64)(ust), (*C.INT64)(msc), (*C.INT64)(sbc)))
}
// wgl

func CreateContext(hDc Pointer) Pointer {
	return (Pointer)(C.goglCreateContext((C.HDC)(hDc)))
}
func DeleteContext(oldContext Pointer) int32 {
	return (int32)(C.goglDeleteContext((C.HGLRC)(oldContext)))
}
func GetCurrentContext() Pointer {
	return (Pointer)(C.goglGetCurrentContext())
}
func MakeCurrent(hDc Pointer, newContext Pointer) int32 {
	return (int32)(C.goglMakeCurrent((C.HDC)(hDc), (C.HGLRC)(newContext)))
}
func CopyContext(hglrcSrc Pointer, hglrcDst Pointer, mask uint32) int32 {
	return (int32)(C.goglCopyContext((C.HGLRC)(hglrcSrc), (C.HGLRC)(hglrcDst), (C.UINT)(mask)))
}
func ChoosePixelFormat(hDc Pointer, pPfd []byte) Int {
	return (Int)(C.goglChoosePixelFormat((C.HDC)(hDc), (C.PIXELFORMATDESCRIPTOR)(pPfd)))
}
func GetCurrentDC() Pointer {
	return (Pointer)(C.goglGetCurrentDC())
}
func GetDefaultProcAddress(lpszProc *byte) Pointer {
	return (Pointer)(C.goglGetDefaultProcAddress((C.LPCSTR)(lpszProc)))
}
func GetProcAddress(lpszProc *byte) Pointer {
	return (Pointer)(C.goglGetProcAddress((C.LPCSTR)(lpszProc)))
}
func GetPixelFormat(hdc Pointer) Int {
	return (Int)(C.goglGetPixelFormat((C.HDC)(hdc)))
}
func SetPixelFormat(hdc Pointer, ipfd Int, ppfd []byte) int32 {
	return (int32)(C.goglSetPixelFormat((C.HDC)(hdc), (C.int)(ipfd), (C.PIXELFORMATDESCRIPTOR)(ppfd)))
}
func SwapBuffers(hdc Pointer) int32 {
	return (int32)(C.goglSwapBuffers((C.HDC)(hdc)))
}
func ShareLists(hrcSrvShare Pointer, hrcSrvSource Pointer) int32 {
	return (int32)(C.goglShareLists((C.HGLRC)(hrcSrvShare), (C.HGLRC)(hrcSrvSource)))
}
func CreateLayerContext(hDc Pointer, level Int) Pointer {
	return (Pointer)(C.goglCreateLayerContext((C.HDC)(hDc), (C.int)(level)))
}
func DescribeLayerPlane(hDc Pointer, pixelFormat Int, layerPlane Int, nBytes uint32, plpd []byte) int32 {
	return (int32)(C.goglDescribeLayerPlane((C.HDC)(hDc), (C.int)(pixelFormat), (C.int)(layerPlane), (C.UINT)(nBytes), (C.LAYERPLANEDESCRIPTOR)(plpd)))
}
func SetLayerPaletteEntries(hdc Pointer, iLayerPlane Int, iStart Int, cEntries Int, pcr Pointer) Int {
	return (Int)(C.goglSetLayerPaletteEntries((C.HDC)(hdc), (C.int)(iLayerPlane), (C.int)(iStart), (C.int)(cEntries), (C.COLORREF)(pcr)))
}
func GetLayerPaletteEntries(hdc Pointer, iLayerPlane Int, iStart Int, cEntries Int, pcr Pointer) Int {
	return (Int)(C.goglGetLayerPaletteEntries((C.HDC)(hdc), (C.int)(iLayerPlane), (C.int)(iStart), (C.int)(cEntries), (C.COLORREF)(pcr)))
}
func RealizeLayerPalette(hdc Pointer, iLayerPlane Int, bRealize int32) int32 {
	return (int32)(C.goglRealizeLayerPalette((C.HDC)(hdc), (C.int)(iLayerPlane), (C.BOOL)(bRealize)))
}
func SwapLayerBuffers(hdc Pointer, fuFlags uint32) int32 {
	return (int32)(C.goglSwapLayerBuffers((C.HDC)(hdc), (C.UINT)(fuFlags)))
}
func UseFontBitmapsA(hDC Pointer, first uint32, count uint32, listBase uint32) int32 {
	return (int32)(C.goglUseFontBitmapsA((C.HDC)(hDC), (C.DWORD)(first), (C.DWORD)(count), (C.DWORD)(listBase)))
}
func UseFontBitmapsW(hDC Pointer, first uint32, count uint32, listBase uint32) int32 {
	return (int32)(C.goglUseFontBitmapsW((C.HDC)(hDC), (C.DWORD)(first), (C.DWORD)(count), (C.DWORD)(listBase)))
}
func Init3dfxMultisample() error {
	var ret C.int
	if ret = C.init_3DFX_multisample(); ret != 0 {
		return errors.New("unable to initialize 3DFX_multisample")
	}
	return nil
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
func InitArbCreateContextProfile() error {
	var ret C.int
	if ret = C.init_ARB_create_context_profile(); ret != 0 {
		return errors.New("unable to initialize ARB_create_context_profile")
	}
	return nil
}
func InitArbCreateContextRobustness() error {
	var ret C.int
	if ret = C.init_ARB_create_context_robustness(); ret != 0 {
		return errors.New("unable to initialize ARB_create_context_robustness")
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
func InitArbFramebufferSrgb() error {
	var ret C.int
	if ret = C.init_ARB_framebuffer_sRGB(); ret != 0 {
		return errors.New("unable to initialize ARB_framebuffer_sRGB")
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
func InitArbMultisample() error {
	var ret C.int
	if ret = C.init_ARB_multisample(); ret != 0 {
		return errors.New("unable to initialize ARB_multisample")
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
func InitArbPixelFormatFloat() error {
	var ret C.int
	if ret = C.init_ARB_pixel_format_float(); ret != 0 {
		return errors.New("unable to initialize ARB_pixel_format_float")
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
func InitAtiPixelFormatFloat() error {
	var ret C.int
	if ret = C.init_ATI_pixel_format_float(); ret != 0 {
		return errors.New("unable to initialize ATI_pixel_format_float")
	}
	return nil
}
func InitExtDepthFloat() error {
	var ret C.int
	if ret = C.init_EXT_depth_float(); ret != 0 {
		return errors.New("unable to initialize EXT_depth_float")
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
func InitExtFramebufferSrgb() error {
	var ret C.int
	if ret = C.init_EXT_framebuffer_sRGB(); ret != 0 {
		return errors.New("unable to initialize EXT_framebuffer_sRGB")
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
func InitExtMultisample() error {
	var ret C.int
	if ret = C.init_EXT_multisample(); ret != 0 {
		return errors.New("unable to initialize EXT_multisample")
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
func InitExtPixelFormatPackedFloat() error {
	var ret C.int
	if ret = C.init_EXT_pixel_format_packed_float(); ret != 0 {
		return errors.New("unable to initialize EXT_pixel_format_packed_float")
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
func InitExtSwapControlTear() error {
	var ret C.int
	if ret = C.init_EXT_swap_control_tear(); ret != 0 {
		return errors.New("unable to initialize EXT_swap_control_tear")
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
func InitNvDxInterop2() error {
	var ret C.int
	if ret = C.init_NV_DX_interop2(); ret != 0 {
		return errors.New("unable to initialize NV_DX_interop2")
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
func InitNvFloatBuffer() error {
	var ret C.int
	if ret = C.init_NV_float_buffer(); ret != 0 {
		return errors.New("unable to initialize NV_float_buffer")
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
func InitNvMultisampleCoverage() error {
	var ret C.int
	if ret = C.init_NV_multisample_coverage(); ret != 0 {
		return errors.New("unable to initialize NV_multisample_coverage")
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