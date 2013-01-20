// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// ARB_create_context: https://www.opengl.org/registry/specs/ARB/create_context.txt
// 
// ARB_get_proc_address: https://www.opengl.org/registry/specs/ARB/get_proc_address.txt
// 
// EXT_import_context: https://www.opengl.org/registry/specs/EXT/import_context.txt
// 
// EXT_swap_control: https://www.opengl.org/registry/specs/EXT/swap_control.txt
// 
// EXT_texture_from_pixmap: https://www.opengl.org/registry/specs/EXT/texture_from_pixmap.txt
// 
// MESA_agp_offset: https://www.opengl.org/registry/specs/MESA/agp_offset.txt
// 
// MESA_copy_sub_buffer: https://www.opengl.org/registry/specs/MESA/copy_sub_buffer.txt
// 
// MESA_pixmap_colormap: https://www.opengl.org/registry/specs/MESA/pixmap_colormap.txt
// 
// MESA_release_buffers: https://www.opengl.org/registry/specs/MESA/release_buffers.txt
// 
// MESA_set_3dfx_mode: https://www.opengl.org/registry/specs/MESA/set_3dfx_mode.txt
// 
// NV_copy_image: https://www.opengl.org/registry/specs/NV/copy_image.txt
// 
// NV_present_video: https://www.opengl.org/registry/specs/NV/present_video.txt
// 
// NV_swap_group: https://www.opengl.org/registry/specs/NV/swap_group.txt
// 
// NV_video_capture: https://www.opengl.org/registry/specs/NV/video_capture.txt
// 
// NV_video_output: https://www.opengl.org/registry/specs/NV/video_output.txt
// 
// OML_sync_control: https://www.opengl.org/registry/specs/OML/sync_control.txt
// 
// SGIX_dmbuffer: https://www.opengl.org/registry/specs/SGIX/dmbuffer.txt
// 
// SGIX_fbconfig: https://www.opengl.org/registry/specs/SGIX/fbconfig.txt
// 
// SGIX_hyperpipe: https://www.opengl.org/registry/specs/SGIX/hyperpipe.txt
// 
// SGIX_pbuffer: https://www.opengl.org/registry/specs/SGIX/pbuffer.txt
// 
// SGIX_swap_barrier: https://www.opengl.org/registry/specs/SGIX/swap_barrier.txt
// 
// SGIX_swap_group: https://www.opengl.org/registry/specs/SGIX/swap_group.txt
// 
// SGIX_video_resize: https://www.opengl.org/registry/specs/SGIX/video_resize.txt
// 
// SGIX_video_source: https://www.opengl.org/registry/specs/SGIX/video_source.txt
// 
// SGI_cushion: https://www.opengl.org/registry/specs/SGI/cushion.txt
// 
// SGI_make_current_read: https://www.opengl.org/registry/specs/SGI/make_current_read.txt
// 
// SGI_swap_control: https://www.opengl.org/registry/specs/SGI/swap_control.txt
// 
// SGI_video_sync: https://www.opengl.org/registry/specs/SGI/video_sync.txt
// 
// SUN_get_transparent_index: https://www.opengl.org/registry/specs/SUN/get_transparent_index.txt
// 
// VERSION_1_3
// 
// VERSION_1_4
// 
// glx: https://www.opengl.org/registry/specs/glx/.txt
// 
// https://www.opengl.org/sdk/docs/man
// 
package glx

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
// #ifndef GLX_ARB_get_proc_address
// typedef void (*__GLXextFuncPtr)(void);
// #endif
// 
// #ifndef GLX_SGIX_video_source
// typedef XID GLXVideoSourceSGIX;
// #endif
// 
// #ifndef GLX_SGIX_fbconfig
// typedef XID GLXFBConfigIDSGIX;
// typedef struct __GLXFBConfigRec *GLXFBConfigSGIX;
// #endif
// 
// #ifndef GLX_SGIX_pbuffer
// typedef XID GLXPbufferSGIX;
// typedef struct {
// int type;
// unsigned long serial;	  /* # of last request processed by server */
// Bool send_event;		  /* true if this came for SendEvent request */
// Display *display;		  /* display the event was read from */
// GLXDrawable drawable;	  /* i.d. of Drawable */
// int event_type;		  /* GLX_DAMAGED_SGIX or GLX_SAVED_SGIX */
// int draw_type;		  /* GLX_WINDOW_SGIX or GLX_PBUFFER_SGIX */
// unsigned int mask;	  /* mask indicating which buffers are affected*/
// int x, y;
// int width, height;
// int count;		  /* if nonzero, at least this many more */
// } GLXBufferClobberEventSGIX;
// #endif
// 
// #ifndef GLX_NV_video_output
// typedef unsigned int GLXVideoDeviceNV;
// #endif
// 
// #ifndef GLX_NV_video_capture
// typedef XID GLXVideoCaptureDeviceNV;
// #endif
// 
// #ifndef GLEXT_64_TYPES_DEFINED
// /* This code block is duplicated in glext.h, so must be protected */
// #define GLEXT_64_TYPES_DEFINED
// /* Define int32_t, int64_t, and uint64_t types for UST/MSC */
// /* (as used in the GLX_OML_sync_control extension). */
// #if defined(__STDC_VERSION__) && __STDC_VERSION__ >= 199901L
// #include <inttypes.h>
// #elif defined(__sun__) || defined(__digital__)
// #include <inttypes.h>
// #if defined(__STDC__)
// #if defined(__arch64__) || defined(_LP64)
// typedef long int int64_t;
// typedef unsigned long int uint64_t;
// #else
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #endif /* __arch64__ */
// #endif /* __STDC__ */
// #elif defined( __VMS ) || defined(__sgi)
// #include <inttypes.h>
// #elif defined(__SCO__) || defined(__USLC__)
// #include <stdint.h>
// #elif defined(__UNIXOS2__) || defined(__SOL64__)
// typedef long int int32_t;
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #elif defined(_WIN32) && defined(__GNUC__)
// #include <stdint.h>
// #elif defined(_WIN32)
// typedef __int32 int32_t;
// typedef __int64 int64_t;
// typedef unsigned __int64 uint64_t;
// #else
// #include <inttypes.h>     /* Fallback option */
// #endif
// #endif
// 
// #ifdef _WIN32
// static HMODULE opengl32 = NULL;
// #endif
// 
// static void* goglxGetProcAddress(const char* name) { 
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
// //  ARB_create_context
// GLXContext (APIENTRYP ptrglxCreateContextAttribsARB)(Display* dpy, GLXFBConfig config, GLXContext share_context, Bool direct, int* attrib_list);
// //  ARB_get_proc_address
// __GLXextFuncPtr (APIENTRYP ptrglxGetProcAddressARB)(GLubyte* procName);
// //  EXT_import_context
// Display * (APIENTRYP ptrglxGetCurrentDisplayEXT)();
// int (APIENTRYP ptrglxQueryContextInfoEXT)(Display* dpy, GLXContext context, int attribute, int* value);
// GLXContextID (APIENTRYP ptrglxGetContextIDEXT)(const GLXContext context);
// GLXContext (APIENTRYP ptrglxImportContextEXT)(Display* dpy, GLXContextID contextID);
// void (APIENTRYP ptrglxFreeContextEXT)(Display* dpy, GLXContext context);
// //  EXT_swap_control
// void (APIENTRYP ptrglxSwapIntervalEXT)(Display* dpy, GLXDrawable drawable, int interval);
// //  EXT_texture_from_pixmap
// void (APIENTRYP ptrglxBindTexImageEXT)(Display* dpy, GLXDrawable drawable, int buffer, int* attrib_list);
// void (APIENTRYP ptrglxReleaseTexImageEXT)(Display* dpy, GLXDrawable drawable, int buffer);
// //  MESA_agp_offset
// unsigned int (APIENTRYP ptrglxGetAGPOffsetMESA)(void* pointer);
// //  MESA_copy_sub_buffer
// void (APIENTRYP ptrglxCopySubBufferMESA)(Display* dpy, GLXDrawable drawable, int x, int y, int width, int height);
// //  MESA_pixmap_colormap
// GLXPixmap (APIENTRYP ptrglxCreateGLXPixmapMESA)(Display* dpy, XVisualInfo* visual, Pixmap pixmap, Colormap cmap);
// //  MESA_release_buffers
// Bool (APIENTRYP ptrglxReleaseBuffersMESA)(Display* dpy, GLXDrawable drawable);
// //  MESA_set_3dfx_mode
// Bool (APIENTRYP ptrglxSet3DfxModeMESA)(int mode);
// //  NV_copy_image
// void (APIENTRYP ptrglxCopyImageSubDataNV)(Display* dpy, GLXContext srcCtx, GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, GLXContext dstCtx, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei width, GLsizei height, GLsizei depth);
// //  NV_present_video
// unsigned int * (APIENTRYP ptrglxEnumerateVideoDevicesNV)(Display* dpy, int screen, int* nelements);
// int (APIENTRYP ptrglxBindVideoDeviceNV)(Display* dpy, unsigned int video_slot, unsigned int video_device, int* attrib_list);
// //  NV_swap_group
// Bool (APIENTRYP ptrglxJoinSwapGroupNV)(Display* dpy, GLXDrawable drawable, GLuint group);
// Bool (APIENTRYP ptrglxBindSwapBarrierNV)(Display* dpy, GLuint group, GLuint barrier);
// Bool (APIENTRYP ptrglxQuerySwapGroupNV)(Display* dpy, GLXDrawable drawable, GLuint* group, GLuint* barrier);
// Bool (APIENTRYP ptrglxQueryMaxSwapGroupsNV)(Display* dpy, int screen, GLuint* maxGroups, GLuint* maxBarriers);
// Bool (APIENTRYP ptrglxQueryFrameCountNV)(Display* dpy, int screen, GLuint* count);
// Bool (APIENTRYP ptrglxResetFrameCountNV)(Display* dpy, int screen);
// //  NV_video_capture
// int (APIENTRYP ptrglxBindVideoCaptureDeviceNV)(Display* dpy, unsigned int video_capture_slot, GLXVideoCaptureDeviceNV device);
// GLXVideoCaptureDeviceNV * (APIENTRYP ptrglxEnumerateVideoCaptureDevicesNV)(Display* dpy, int screen, int* nelements);
// void (APIENTRYP ptrglxLockVideoCaptureDeviceNV)(Display* dpy, GLXVideoCaptureDeviceNV device);
// int (APIENTRYP ptrglxQueryVideoCaptureDeviceNV)(Display* dpy, GLXVideoCaptureDeviceNV device, int attribute, int* value);
// void (APIENTRYP ptrglxReleaseVideoCaptureDeviceNV)(Display* dpy, GLXVideoCaptureDeviceNV device);
// //  NV_video_output
// int (APIENTRYP ptrglxGetVideoDeviceNV)(Display* dpy, int screen, int numVideoDevices, GLXVideoDeviceNV* pVideoDevice);
// int (APIENTRYP ptrglxReleaseVideoDeviceNV)(Display* dpy, int screen, GLXVideoDeviceNV VideoDevice);
// int (APIENTRYP ptrglxBindVideoImageNV)(Display* dpy, GLXVideoDeviceNV VideoDevice, GLXPbuffer pbuf, int iVideoBuffer);
// int (APIENTRYP ptrglxReleaseVideoImageNV)(Display* dpy, GLXPbuffer pbuf);
// int (APIENTRYP ptrglxSendPbufferToVideoNV)(Display* dpy, GLXPbuffer pbuf, int iBufferType, unsigned long* pulCounterPbuffer, GLboolean bBlock);
// int (APIENTRYP ptrglxGetVideoInfoNV)(Display* dpy, int screen, GLXVideoDeviceNV VideoDevice, unsigned long* pulCounterOutputPbuffer, unsigned long* pulCounterOutputVideo);
// //  OML_sync_control
// Bool (APIENTRYP ptrglxGetSyncValuesOML)(Display* dpy, GLXDrawable drawable, int64_t* ust, int64_t* msc, int64_t* sbc);
// Bool (APIENTRYP ptrglxGetMscRateOML)(Display* dpy, GLXDrawable drawable, int32_t* numerator, int32_t* denominator);
// int64_t (APIENTRYP ptrglxSwapBuffersMscOML)(Display* dpy, GLXDrawable drawable, int64_t target_msc, int64_t divisor, int64_t remainder);
// Bool (APIENTRYP ptrglxWaitForMscOML)(Display* dpy, GLXDrawable drawable, int64_t target_msc, int64_t divisor, int64_t remainder, int64_t* ust, int64_t* msc, int64_t* sbc);
// Bool (APIENTRYP ptrglxWaitForSbcOML)(Display* dpy, GLXDrawable drawable, int64_t target_sbc, int64_t* ust, int64_t* msc, int64_t* sbc);
// //  SGIX_dmbuffer
// Bool (APIENTRYP ptrglxAssociateDMPbufferSGIX)(Display* dpy, GLXPbufferSGIX pbuffer, DMparams* params, DMbuffer dmbuffer);
// //  SGIX_fbconfig
// int (APIENTRYP ptrglxGetFBConfigAttribSGIX)(Display* dpy, GLXFBConfigSGIX config, int attribute, int* value);
// GLXFBConfigSGIX * (APIENTRYP ptrglxChooseFBConfigSGIX)(Display* dpy, int screen, int* attrib_list, int* nelements);
// GLXPixmap (APIENTRYP ptrglxCreateGLXPixmapWithConfigSGIX)(Display* dpy, GLXFBConfigSGIX config, Pixmap pixmap);
// GLXContext (APIENTRYP ptrglxCreateContextWithConfigSGIX)(Display* dpy, GLXFBConfigSGIX config, int render_type, GLXContext share_list, Bool direct);
// XVisualInfo * (APIENTRYP ptrglxGetVisualFromFBConfigSGIX)(Display* dpy, GLXFBConfigSGIX config);
// GLXFBConfigSGIX (APIENTRYP ptrglxGetFBConfigFromVisualSGIX)(Display* dpy, XVisualInfo* vis);
// //  SGIX_hyperpipe
// GLXHyperpipeNetworkSGIX * (APIENTRYP ptrglxQueryHyperpipeNetworkSGIX)(Display* dpy, int* npipes);
// int (APIENTRYP ptrglxHyperpipeConfigSGIX)(Display* dpy, int networkId, int npipes, GLXHyperpipeConfigSGIX* cfg, int* hpId);
// GLXHyperpipeConfigSGIX * (APIENTRYP ptrglxQueryHyperpipeConfigSGIX)(Display* dpy, int hpId, int* npipes);
// int (APIENTRYP ptrglxDestroyHyperpipeConfigSGIX)(Display* dpy, int hpId);
// int (APIENTRYP ptrglxBindHyperpipeSGIX)(Display* dpy, int hpId);
// int (APIENTRYP ptrglxQueryHyperpipeBestAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, void* attribList, void* returnAttribList);
// int (APIENTRYP ptrglxHyperpipeAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, void* attribList);
// int (APIENTRYP ptrglxQueryHyperpipeAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, void* returnAttribList);
// //  SGIX_pbuffer
// GLXPbufferSGIX (APIENTRYP ptrglxCreateGLXPbufferSGIX)(Display* dpy, GLXFBConfigSGIX config, unsigned int width, unsigned int height, int* attrib_list);
// void (APIENTRYP ptrglxDestroyGLXPbufferSGIX)(Display* dpy, GLXPbufferSGIX pbuf);
// int (APIENTRYP ptrglxQueryGLXPbufferSGIX)(Display* dpy, GLXPbufferSGIX pbuf, int attribute, unsigned int* value);
// void (APIENTRYP ptrglxSelectEventSGIX)(Display* dpy, GLXDrawable drawable, unsigned long mask);
// void (APIENTRYP ptrglxGetSelectedEventSGIX)(Display* dpy, GLXDrawable drawable, unsigned long* mask);
// //  SGIX_swap_barrier
// void (APIENTRYP ptrglxBindSwapBarrierSGIX)(Display* dpy, GLXDrawable drawable, int barrier);
// Bool (APIENTRYP ptrglxQueryMaxSwapBarriersSGIX)(Display* dpy, int screen, int* max);
// //  SGIX_swap_group
// void (APIENTRYP ptrglxJoinSwapGroupSGIX)(Display* dpy, GLXDrawable drawable, GLXDrawable member);
// //  SGIX_video_resize
// int (APIENTRYP ptrglxBindChannelToWindowSGIX)(Display* display, int screen, int channel, Window window);
// int (APIENTRYP ptrglxChannelRectSGIX)(Display* display, int screen, int channel, int x, int y, int w, int h);
// int (APIENTRYP ptrglxQueryChannelRectSGIX)(Display* display, int screen, int channel, int* dx, int* dy, int* dw, int* dh);
// int (APIENTRYP ptrglxQueryChannelDeltasSGIX)(Display* display, int screen, int channel, int* x, int* y, int* w, int* h);
// int (APIENTRYP ptrglxChannelRectSyncSGIX)(Display* display, int screen, int channel, GLenum synctype);
// //  SGIX_video_source
// GLXVideoSourceSGIX (APIENTRYP ptrglxCreateGLXVideoSourceSGIX)(Display* display, int screen, VLServer server, VLPath path, int nodeClass, VLNode drainNode);
// void (APIENTRYP ptrglxDestroyGLXVideoSourceSGIX)(Display* dpy, GLXVideoSourceSGIX glxvideosource);
// //  SGI_cushion
// void (APIENTRYP ptrglxCushionSGI)(Display* dpy, Window window, float cushion);
// //  SGI_make_current_read
// Bool (APIENTRYP ptrglxMakeCurrentReadSGI)(Display* dpy, GLXDrawable draw, GLXDrawable read, GLXContext ctx);
// GLXDrawable (APIENTRYP ptrglxGetCurrentReadDrawableSGI)();
// //  SGI_swap_control
// int (APIENTRYP ptrglxSwapIntervalSGI)(int interval);
// //  SGI_video_sync
// int (APIENTRYP ptrglxGetVideoSyncSGI)(unsigned int* count);
// int (APIENTRYP ptrglxWaitVideoSyncSGI)(int divisor, int remainder, unsigned int* count);
// //  SUN_get_transparent_index
// Status (APIENTRYP ptrglxGetTransparentIndexSUN)(Display* dpy, Window overlay, Window underlay, long* pTransparentIndex);
// //  VERSION_1_3
// GLXFBConfig * (APIENTRYP ptrglxGetFBConfigs)(Display* dpy, int screen, int* nelements);
// GLXFBConfig * (APIENTRYP ptrglxChooseFBConfig)(Display* dpy, int screen, int* attrib_list, int* nelements);
// int (APIENTRYP ptrglxGetFBConfigAttrib)(Display* dpy, GLXFBConfig config, int attribute, int* value);
// XVisualInfo * (APIENTRYP ptrglxGetVisualFromFBConfig)(Display* dpy, GLXFBConfig config);
// GLXWindow (APIENTRYP ptrglxCreateWindow)(Display* dpy, GLXFBConfig config, Window win, int* attrib_list);
// void (APIENTRYP ptrglxDestroyWindow)(Display* dpy, GLXWindow win);
// GLXPixmap (APIENTRYP ptrglxCreatePixmap)(Display* dpy, GLXFBConfig config, Pixmap pixmap, int* attrib_list);
// void (APIENTRYP ptrglxDestroyPixmap)(Display* dpy, GLXPixmap pixmap);
// GLXPbuffer (APIENTRYP ptrglxCreatePbuffer)(Display* dpy, GLXFBConfig config, int* attrib_list);
// void (APIENTRYP ptrglxDestroyPbuffer)(Display* dpy, GLXPbuffer pbuf);
// void (APIENTRYP ptrglxQueryDrawable)(Display* dpy, GLXDrawable draw, int attribute, unsigned int* value);
// GLXContext (APIENTRYP ptrglxCreateNewContext)(Display* dpy, GLXFBConfig config, int render_type, GLXContext share_list, Bool direct);
// Bool (APIENTRYP ptrglxMakeContextCurrent)(Display* dpy, GLXDrawable draw, GLXDrawable read, GLXContext ctx);
// GLXDrawable (APIENTRYP ptrglxGetCurrentReadDrawable)();
// Display * (APIENTRYP ptrglxGetCurrentDisplay)();
// int (APIENTRYP ptrglxQueryContext)(Display* dpy, GLXContext ctx, int attribute, int* value);
// void (APIENTRYP ptrglxSelectEvent)(Display* dpy, GLXDrawable draw, unsigned long event_mask);
// void (APIENTRYP ptrglxGetSelectedEvent)(Display* dpy, GLXDrawable draw, unsigned long* event_mask);
// //  VERSION_1_4
// __GLXextFuncPtr (APIENTRYP ptrglxGetProcAddress)(GLubyte* procName);
// //  glx
// void (APIENTRYP ptrglxRender)();
// void (APIENTRYP ptrglxRenderLarge)();
// void (APIENTRYP ptrglxCreateContext)(GLint gc_id, GLint screen, GLint visual, GLint share_list);
// void (APIENTRYP ptrglxDestroyContext)(GLint context);
// void (APIENTRYP ptrglxMakeCurrent)(GLint drawable, GLint context);
// void (APIENTRYP ptrglxIsDirect)(GLint dpy, GLint context);
// void (APIENTRYP ptrglxQueryVersion)(GLint* major, GLint* minor);
// void (APIENTRYP ptrglxWaitGL)(GLint context);
// void (APIENTRYP ptrglxWaitX)();
// void (APIENTRYP ptrglxCopyContext)(GLint source, GLint dest, GLint mask);
// void (APIENTRYP ptrglxSwapBuffers)(GLint drawable);
// void (APIENTRYP ptrglxUseXFont)(GLint font, GLint first, GLint count, GLint list_base);
// void (APIENTRYP ptrglxCreateGLXPixmap)(GLint visual, GLint pixmap, GLint glxpixmap);
// void (APIENTRYP ptrglxGetVisualConfigs)();
// void (APIENTRYP ptrglxDestroyGLXPixmap)(GLint pixmap);
// void (APIENTRYP ptrglxVendorPrivate)();
// void (APIENTRYP ptrglxVendorPrivateWithReply)();
// void (APIENTRYP ptrglxQueryExtensionsString)(GLint screen);
// void (APIENTRYP ptrglxQueryServerString)(GLint screen, GLint name);
// void (APIENTRYP ptrglxClientInfo)();
// void (APIENTRYP ptrglxGetFBConfigs)();
// void (APIENTRYP ptrglxCreatePixmap)(GLint config, GLint pixmap, GLint glxpixmap);
// void (APIENTRYP ptrglxDestroyPixmap)(GLint glxpixmap);
// void (APIENTRYP ptrglxCreateNewContext)(GLint config, GLint render_type, GLint share_list, GLint direct);
// void (APIENTRYP ptrglxQueryContext)();
// void (APIENTRYP ptrglxMakeContextCurrent)(GLint drawable, GLint readdrawable, GLint context);
// void (APIENTRYP ptrglxCreatePbuffer)(GLint config, GLint pbuffer);
// void (APIENTRYP ptrglxDestroyPbuffer)(GLint pbuffer);
// void (APIENTRYP ptrglxGetDrawableAttributes)(GLint drawable);
// void (APIENTRYP ptrglxChangeDrawableAttributes)(GLint drawable);
// void (APIENTRYP ptrglxCreateWindow)(GLint config, GLint window, GLint glxwindow);
// void (APIENTRYP ptrglxDestroyWindow)(GLint glxwindow);
// void (APIENTRYP ptrglxSwapIntervalSGI)();
// void (APIENTRYP ptrglxMakeCurrentReadSGI)(GLint drawable, GLint readdrawable, GLint context);
// void (APIENTRYP ptrglxCreateGLXVideoSourceSGIX)(GLint dpy, GLint screen, GLint server, GLint path, GLint class, GLint node);
// void (APIENTRYP ptrglxDestroyGLXVideoSourceSGIX)(GLint dpy, GLint glxvideosource);
// void (APIENTRYP ptrglxQueryContextInfoEXT)();
// void (APIENTRYP ptrglxGetFBConfigsSGIX)();
// void (APIENTRYP ptrglxCreateContextWithConfigSGIX)(GLint gc_id, GLint screen, GLint config, GLint share_list);
// void (APIENTRYP ptrglxCreateGLXPixmapWithConfigSGIX)(GLint config, GLint pixmap, GLint glxpixmap);
// void (APIENTRYP ptrglxCreateGLXPbufferSGIX)(GLint config, GLint pbuffer);
// void (APIENTRYP ptrglxDestroyGLXPbufferSGIX)(GLint pbuffer);
// void (APIENTRYP ptrglxChangeDrawableAttributesSGIX)(GLint drawable);
// void (APIENTRYP ptrglxGetDrawableAttributesSGIX)(GLint drawable);
// void (APIENTRYP ptrglxJoinSwapGroupSGIX)(GLint window, GLint group);
// void (APIENTRYP ptrglxBindSwapBarrierSGIX)(GLint window, GLint barrier);
// void (APIENTRYP ptrglxQueryMaxSwapBarriersSGIX)();
// GLXHyperpipeNetworkSGIX * (APIENTRYP ptrglxQueryHyperpipeNetworkSGIX)(Display* dpy, int* npipes);
// int (APIENTRYP ptrglxHyperpipeConfigSGIX)(Display* dpy, int networkId, int npipes, GLXHyperpipeConfigSGIX* cfg, int* hpId);
// GLXHyperpipeConfigSGIX * (APIENTRYP ptrglxQueryHyperpipeConfigSGIX)(Display* dpy, int hpId, int* npipes);
// int (APIENTRYP ptrglxDestroyHyperpipeConfigSGIX)(Display* dpy, int hpId);
// int (APIENTRYP ptrglxBindHyperpipeSGIX)(Display* dpy, int hpId);
// int (APIENTRYP ptrglxQueryHyperpipeBestAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, GLvoid* attribList, GLvoid* returnAttribList);
// int (APIENTRYP ptrglxHyperpipeAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, void* attribList);
// int (APIENTRYP ptrglxQueryHyperpipeAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, void* returnAttribList);
// 
// //  ARB_create_context
// GLXContext goglxCreateContextAttribsARB(Display* dpy, GLXFBConfig config, GLXContext share_context, Bool direct, int* attrib_list) {
// 	return (*ptrglxCreateContextAttribsARB)(dpy, config, share_context, direct, attrib_list);
// }
// //  ARB_get_proc_address
// __GLXextFuncPtr goglxGetProcAddressARB(GLubyte* procName) {
// 	return (*ptrglxGetProcAddressARB)(procName);
// }
// //  EXT_import_context
// Display * goglxGetCurrentDisplayEXT() {
// 	return (*ptrglxGetCurrentDisplayEXT)();
// }
// int goglxQueryContextInfoEXT(Display* dpy, GLXContext context, int attribute, int* value) {
// 	return (*ptrglxQueryContextInfoEXT)(dpy, context, attribute, value);
// }
// GLXContextID goglxGetContextIDEXT(const GLXContext context) {
// 	return (*ptrglxGetContextIDEXT)(context);
// }
// GLXContext goglxImportContextEXT(Display* dpy, GLXContextID contextID) {
// 	return (*ptrglxImportContextEXT)(dpy, contextID);
// }
// void goglxFreeContextEXT(Display* dpy, GLXContext context) {
// 	(*ptrglxFreeContextEXT)(dpy, context);
// }
// //  EXT_swap_control
// void goglxSwapIntervalEXT(Display* dpy, GLXDrawable drawable, int interval) {
// 	(*ptrglxSwapIntervalEXT)(dpy, drawable, interval);
// }
// //  EXT_texture_from_pixmap
// void goglxBindTexImageEXT(Display* dpy, GLXDrawable drawable, int buffer, int* attrib_list) {
// 	(*ptrglxBindTexImageEXT)(dpy, drawable, buffer, attrib_list);
// }
// void goglxReleaseTexImageEXT(Display* dpy, GLXDrawable drawable, int buffer) {
// 	(*ptrglxReleaseTexImageEXT)(dpy, drawable, buffer);
// }
// //  MESA_agp_offset
// unsigned int goglxGetAGPOffsetMESA(void* pointer) {
// 	return (*ptrglxGetAGPOffsetMESA)(pointer);
// }
// //  MESA_copy_sub_buffer
// void goglxCopySubBufferMESA(Display* dpy, GLXDrawable drawable, int x, int y, int width, int height) {
// 	(*ptrglxCopySubBufferMESA)(dpy, drawable, x, y, width, height);
// }
// //  MESA_pixmap_colormap
// GLXPixmap goglxCreateGLXPixmapMESA(Display* dpy, XVisualInfo* visual, Pixmap pixmap, Colormap cmap) {
// 	return (*ptrglxCreateGLXPixmapMESA)(dpy, visual, pixmap, cmap);
// }
// //  MESA_release_buffers
// Bool goglxReleaseBuffersMESA(Display* dpy, GLXDrawable drawable) {
// 	return (*ptrglxReleaseBuffersMESA)(dpy, drawable);
// }
// //  MESA_set_3dfx_mode
// Bool goglxSet3DfxModeMESA(int mode) {
// 	return (*ptrglxSet3DfxModeMESA)(mode);
// }
// //  NV_copy_image
// void goglxCopyImageSubDataNV(Display* dpy, GLXContext srcCtx, GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, GLXContext dstCtx, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei width, GLsizei height, GLsizei depth) {
// 	(*ptrglxCopyImageSubDataNV)(dpy, srcCtx, srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstCtx, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, width, height, depth);
// }
// //  NV_present_video
// unsigned int * goglxEnumerateVideoDevicesNV(Display* dpy, int screen, int* nelements) {
// 	return (*ptrglxEnumerateVideoDevicesNV)(dpy, screen, nelements);
// }
// int goglxBindVideoDeviceNV(Display* dpy, unsigned int video_slot, unsigned int video_device, int* attrib_list) {
// 	return (*ptrglxBindVideoDeviceNV)(dpy, video_slot, video_device, attrib_list);
// }
// //  NV_swap_group
// Bool goglxJoinSwapGroupNV(Display* dpy, GLXDrawable drawable, GLuint group) {
// 	return (*ptrglxJoinSwapGroupNV)(dpy, drawable, group);
// }
// Bool goglxBindSwapBarrierNV(Display* dpy, GLuint group, GLuint barrier) {
// 	return (*ptrglxBindSwapBarrierNV)(dpy, group, barrier);
// }
// Bool goglxQuerySwapGroupNV(Display* dpy, GLXDrawable drawable, GLuint* group, GLuint* barrier) {
// 	return (*ptrglxQuerySwapGroupNV)(dpy, drawable, group, barrier);
// }
// Bool goglxQueryMaxSwapGroupsNV(Display* dpy, int screen, GLuint* maxGroups, GLuint* maxBarriers) {
// 	return (*ptrglxQueryMaxSwapGroupsNV)(dpy, screen, maxGroups, maxBarriers);
// }
// Bool goglxQueryFrameCountNV(Display* dpy, int screen, GLuint* count) {
// 	return (*ptrglxQueryFrameCountNV)(dpy, screen, count);
// }
// Bool goglxResetFrameCountNV(Display* dpy, int screen) {
// 	return (*ptrglxResetFrameCountNV)(dpy, screen);
// }
// //  NV_video_capture
// int goglxBindVideoCaptureDeviceNV(Display* dpy, unsigned int video_capture_slot, GLXVideoCaptureDeviceNV device) {
// 	return (*ptrglxBindVideoCaptureDeviceNV)(dpy, video_capture_slot, device);
// }
// GLXVideoCaptureDeviceNV * goglxEnumerateVideoCaptureDevicesNV(Display* dpy, int screen, int* nelements) {
// 	return (*ptrglxEnumerateVideoCaptureDevicesNV)(dpy, screen, nelements);
// }
// void goglxLockVideoCaptureDeviceNV(Display* dpy, GLXVideoCaptureDeviceNV device) {
// 	(*ptrglxLockVideoCaptureDeviceNV)(dpy, device);
// }
// int goglxQueryVideoCaptureDeviceNV(Display* dpy, GLXVideoCaptureDeviceNV device, int attribute, int* value) {
// 	return (*ptrglxQueryVideoCaptureDeviceNV)(dpy, device, attribute, value);
// }
// void goglxReleaseVideoCaptureDeviceNV(Display* dpy, GLXVideoCaptureDeviceNV device) {
// 	(*ptrglxReleaseVideoCaptureDeviceNV)(dpy, device);
// }
// //  NV_video_output
// int goglxGetVideoDeviceNV(Display* dpy, int screen, int numVideoDevices, GLXVideoDeviceNV* pVideoDevice) {
// 	return (*ptrglxGetVideoDeviceNV)(dpy, screen, numVideoDevices, pVideoDevice);
// }
// int goglxReleaseVideoDeviceNV(Display* dpy, int screen, GLXVideoDeviceNV VideoDevice) {
// 	return (*ptrglxReleaseVideoDeviceNV)(dpy, screen, VideoDevice);
// }
// int goglxBindVideoImageNV(Display* dpy, GLXVideoDeviceNV VideoDevice, GLXPbuffer pbuf, int iVideoBuffer) {
// 	return (*ptrglxBindVideoImageNV)(dpy, VideoDevice, pbuf, iVideoBuffer);
// }
// int goglxReleaseVideoImageNV(Display* dpy, GLXPbuffer pbuf) {
// 	return (*ptrglxReleaseVideoImageNV)(dpy, pbuf);
// }
// int goglxSendPbufferToVideoNV(Display* dpy, GLXPbuffer pbuf, int iBufferType, unsigned long* pulCounterPbuffer, GLboolean bBlock) {
// 	return (*ptrglxSendPbufferToVideoNV)(dpy, pbuf, iBufferType, pulCounterPbuffer, bBlock);
// }
// int goglxGetVideoInfoNV(Display* dpy, int screen, GLXVideoDeviceNV VideoDevice, unsigned long* pulCounterOutputPbuffer, unsigned long* pulCounterOutputVideo) {
// 	return (*ptrglxGetVideoInfoNV)(dpy, screen, VideoDevice, pulCounterOutputPbuffer, pulCounterOutputVideo);
// }
// //  OML_sync_control
// Bool goglxGetSyncValuesOML(Display* dpy, GLXDrawable drawable, int64_t* ust, int64_t* msc, int64_t* sbc) {
// 	return (*ptrglxGetSyncValuesOML)(dpy, drawable, ust, msc, sbc);
// }
// Bool goglxGetMscRateOML(Display* dpy, GLXDrawable drawable, int32_t* numerator, int32_t* denominator) {
// 	return (*ptrglxGetMscRateOML)(dpy, drawable, numerator, denominator);
// }
// int64_t goglxSwapBuffersMscOML(Display* dpy, GLXDrawable drawable, int64_t target_msc, int64_t divisor, int64_t remainder) {
// 	return (*ptrglxSwapBuffersMscOML)(dpy, drawable, target_msc, divisor, remainder);
// }
// Bool goglxWaitForMscOML(Display* dpy, GLXDrawable drawable, int64_t target_msc, int64_t divisor, int64_t remainder, int64_t* ust, int64_t* msc, int64_t* sbc) {
// 	return (*ptrglxWaitForMscOML)(dpy, drawable, target_msc, divisor, remainder, ust, msc, sbc);
// }
// Bool goglxWaitForSbcOML(Display* dpy, GLXDrawable drawable, int64_t target_sbc, int64_t* ust, int64_t* msc, int64_t* sbc) {
// 	return (*ptrglxWaitForSbcOML)(dpy, drawable, target_sbc, ust, msc, sbc);
// }
// //  SGIX_dmbuffer
// Bool goglxAssociateDMPbufferSGIX(Display* dpy, GLXPbufferSGIX pbuffer, DMparams* params, DMbuffer dmbuffer) {
// 	return (*ptrglxAssociateDMPbufferSGIX)(dpy, pbuffer, params, dmbuffer);
// }
// //  SGIX_fbconfig
// int goglxGetFBConfigAttribSGIX(Display* dpy, GLXFBConfigSGIX config, int attribute, int* value) {
// 	return (*ptrglxGetFBConfigAttribSGIX)(dpy, config, attribute, value);
// }
// GLXFBConfigSGIX * goglxChooseFBConfigSGIX(Display* dpy, int screen, int* attrib_list, int* nelements) {
// 	return (*ptrglxChooseFBConfigSGIX)(dpy, screen, attrib_list, nelements);
// }
// GLXPixmap goglxCreateGLXPixmapWithConfigSGIX(Display* dpy, GLXFBConfigSGIX config, Pixmap pixmap) {
// 	return (*ptrglxCreateGLXPixmapWithConfigSGIX)(dpy, config, pixmap);
// }
// GLXContext goglxCreateContextWithConfigSGIX(Display* dpy, GLXFBConfigSGIX config, int render_type, GLXContext share_list, Bool direct) {
// 	return (*ptrglxCreateContextWithConfigSGIX)(dpy, config, render_type, share_list, direct);
// }
// XVisualInfo * goglxGetVisualFromFBConfigSGIX(Display* dpy, GLXFBConfigSGIX config) {
// 	return (*ptrglxGetVisualFromFBConfigSGIX)(dpy, config);
// }
// GLXFBConfigSGIX goglxGetFBConfigFromVisualSGIX(Display* dpy, XVisualInfo* vis) {
// 	return (*ptrglxGetFBConfigFromVisualSGIX)(dpy, vis);
// }
// //  SGIX_hyperpipe
// GLXHyperpipeNetworkSGIX * goglxQueryHyperpipeNetworkSGIX(Display* dpy, int* npipes) {
// 	return (*ptrglxQueryHyperpipeNetworkSGIX)(dpy, npipes);
// }
// int goglxHyperpipeConfigSGIX(Display* dpy, int networkId, int npipes, GLXHyperpipeConfigSGIX* cfg, int* hpId) {
// 	return (*ptrglxHyperpipeConfigSGIX)(dpy, networkId, npipes, cfg, hpId);
// }
// GLXHyperpipeConfigSGIX * goglxQueryHyperpipeConfigSGIX(Display* dpy, int hpId, int* npipes) {
// 	return (*ptrglxQueryHyperpipeConfigSGIX)(dpy, hpId, npipes);
// }
// int goglxDestroyHyperpipeConfigSGIX(Display* dpy, int hpId) {
// 	return (*ptrglxDestroyHyperpipeConfigSGIX)(dpy, hpId);
// }
// int goglxBindHyperpipeSGIX(Display* dpy, int hpId) {
// 	return (*ptrglxBindHyperpipeSGIX)(dpy, hpId);
// }
// int goglxQueryHyperpipeBestAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, void* attribList, void* returnAttribList) {
// 	return (*ptrglxQueryHyperpipeBestAttribSGIX)(dpy, timeSlice, attrib, size, attribList, returnAttribList);
// }
// int goglxHyperpipeAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, void* attribList) {
// 	return (*ptrglxHyperpipeAttribSGIX)(dpy, timeSlice, attrib, size, attribList);
// }
// int goglxQueryHyperpipeAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, void* returnAttribList) {
// 	return (*ptrglxQueryHyperpipeAttribSGIX)(dpy, timeSlice, attrib, size, returnAttribList);
// }
// //  SGIX_pbuffer
// GLXPbufferSGIX goglxCreateGLXPbufferSGIX(Display* dpy, GLXFBConfigSGIX config, unsigned int width, unsigned int height, int* attrib_list) {
// 	return (*ptrglxCreateGLXPbufferSGIX)(dpy, config, width, height, attrib_list);
// }
// void goglxDestroyGLXPbufferSGIX(Display* dpy, GLXPbufferSGIX pbuf) {
// 	(*ptrglxDestroyGLXPbufferSGIX)(dpy, pbuf);
// }
// int goglxQueryGLXPbufferSGIX(Display* dpy, GLXPbufferSGIX pbuf, int attribute, unsigned int* value) {
// 	return (*ptrglxQueryGLXPbufferSGIX)(dpy, pbuf, attribute, value);
// }
// void goglxSelectEventSGIX(Display* dpy, GLXDrawable drawable, unsigned long mask) {
// 	(*ptrglxSelectEventSGIX)(dpy, drawable, mask);
// }
// void goglxGetSelectedEventSGIX(Display* dpy, GLXDrawable drawable, unsigned long* mask) {
// 	(*ptrglxGetSelectedEventSGIX)(dpy, drawable, mask);
// }
// //  SGIX_swap_barrier
// void goglxBindSwapBarrierSGIX(Display* dpy, GLXDrawable drawable, int barrier) {
// 	(*ptrglxBindSwapBarrierSGIX)(dpy, drawable, barrier);
// }
// Bool goglxQueryMaxSwapBarriersSGIX(Display* dpy, int screen, int* max) {
// 	return (*ptrglxQueryMaxSwapBarriersSGIX)(dpy, screen, max);
// }
// //  SGIX_swap_group
// void goglxJoinSwapGroupSGIX(Display* dpy, GLXDrawable drawable, GLXDrawable member) {
// 	(*ptrglxJoinSwapGroupSGIX)(dpy, drawable, member);
// }
// //  SGIX_video_resize
// int goglxBindChannelToWindowSGIX(Display* display, int screen, int channel, Window window) {
// 	return (*ptrglxBindChannelToWindowSGIX)(display, screen, channel, window);
// }
// int goglxChannelRectSGIX(Display* display, int screen, int channel, int x, int y, int w, int h) {
// 	return (*ptrglxChannelRectSGIX)(display, screen, channel, x, y, w, h);
// }
// int goglxQueryChannelRectSGIX(Display* display, int screen, int channel, int* dx, int* dy, int* dw, int* dh) {
// 	return (*ptrglxQueryChannelRectSGIX)(display, screen, channel, dx, dy, dw, dh);
// }
// int goglxQueryChannelDeltasSGIX(Display* display, int screen, int channel, int* x, int* y, int* w, int* h) {
// 	return (*ptrglxQueryChannelDeltasSGIX)(display, screen, channel, x, y, w, h);
// }
// int goglxChannelRectSyncSGIX(Display* display, int screen, int channel, GLenum synctype) {
// 	return (*ptrglxChannelRectSyncSGIX)(display, screen, channel, synctype);
// }
// //  SGIX_video_source
// GLXVideoSourceSGIX goglxCreateGLXVideoSourceSGIX(Display* display, int screen, VLServer server, VLPath path, int nodeClass, VLNode drainNode) {
// 	return (*ptrglxCreateGLXVideoSourceSGIX)(display, screen, server, path, nodeClass, drainNode);
// }
// void goglxDestroyGLXVideoSourceSGIX(Display* dpy, GLXVideoSourceSGIX glxvideosource) {
// 	(*ptrglxDestroyGLXVideoSourceSGIX)(dpy, glxvideosource);
// }
// //  SGI_cushion
// void goglxCushionSGI(Display* dpy, Window window, float cushion) {
// 	(*ptrglxCushionSGI)(dpy, window, cushion);
// }
// //  SGI_make_current_read
// Bool goglxMakeCurrentReadSGI(Display* dpy, GLXDrawable draw, GLXDrawable read, GLXContext ctx) {
// 	return (*ptrglxMakeCurrentReadSGI)(dpy, draw, read, ctx);
// }
// GLXDrawable goglxGetCurrentReadDrawableSGI() {
// 	return (*ptrglxGetCurrentReadDrawableSGI)();
// }
// //  SGI_swap_control
// int goglxSwapIntervalSGI(int interval) {
// 	return (*ptrglxSwapIntervalSGI)(interval);
// }
// //  SGI_video_sync
// int goglxGetVideoSyncSGI(unsigned int* count) {
// 	return (*ptrglxGetVideoSyncSGI)(count);
// }
// int goglxWaitVideoSyncSGI(int divisor, int remainder, unsigned int* count) {
// 	return (*ptrglxWaitVideoSyncSGI)(divisor, remainder, count);
// }
// //  SUN_get_transparent_index
// Status goglxGetTransparentIndexSUN(Display* dpy, Window overlay, Window underlay, long* pTransparentIndex) {
// 	return (*ptrglxGetTransparentIndexSUN)(dpy, overlay, underlay, pTransparentIndex);
// }
// //  VERSION_1_3
// GLXFBConfig * goglxGetFBConfigs(Display* dpy, int screen, int* nelements) {
// 	return (*ptrglxGetFBConfigs)(dpy, screen, nelements);
// }
// GLXFBConfig * goglxChooseFBConfig(Display* dpy, int screen, int* attrib_list, int* nelements) {
// 	return (*ptrglxChooseFBConfig)(dpy, screen, attrib_list, nelements);
// }
// int goglxGetFBConfigAttrib(Display* dpy, GLXFBConfig config, int attribute, int* value) {
// 	return (*ptrglxGetFBConfigAttrib)(dpy, config, attribute, value);
// }
// XVisualInfo * goglxGetVisualFromFBConfig(Display* dpy, GLXFBConfig config) {
// 	return (*ptrglxGetVisualFromFBConfig)(dpy, config);
// }
// GLXWindow goglxCreateWindow(Display* dpy, GLXFBConfig config, Window win, int* attrib_list) {
// 	return (*ptrglxCreateWindow)(dpy, config, win, attrib_list);
// }
// void goglxDestroyWindow(Display* dpy, GLXWindow win) {
// 	(*ptrglxDestroyWindow)(dpy, win);
// }
// GLXPixmap goglxCreatePixmap(Display* dpy, GLXFBConfig config, Pixmap pixmap, int* attrib_list) {
// 	return (*ptrglxCreatePixmap)(dpy, config, pixmap, attrib_list);
// }
// void goglxDestroyPixmap(Display* dpy, GLXPixmap pixmap) {
// 	(*ptrglxDestroyPixmap)(dpy, pixmap);
// }
// GLXPbuffer goglxCreatePbuffer(Display* dpy, GLXFBConfig config, int* attrib_list) {
// 	return (*ptrglxCreatePbuffer)(dpy, config, attrib_list);
// }
// void goglxDestroyPbuffer(Display* dpy, GLXPbuffer pbuf) {
// 	(*ptrglxDestroyPbuffer)(dpy, pbuf);
// }
// void goglxQueryDrawable(Display* dpy, GLXDrawable draw, int attribute, unsigned int* value) {
// 	(*ptrglxQueryDrawable)(dpy, draw, attribute, value);
// }
// GLXContext goglxCreateNewContext(Display* dpy, GLXFBConfig config, int render_type, GLXContext share_list, Bool direct) {
// 	return (*ptrglxCreateNewContext)(dpy, config, render_type, share_list, direct);
// }
// Bool goglxMakeContextCurrent(Display* dpy, GLXDrawable draw, GLXDrawable read, GLXContext ctx) {
// 	return (*ptrglxMakeContextCurrent)(dpy, draw, read, ctx);
// }
// GLXDrawable goglxGetCurrentReadDrawable() {
// 	return (*ptrglxGetCurrentReadDrawable)();
// }
// Display * goglxGetCurrentDisplay() {
// 	return (*ptrglxGetCurrentDisplay)();
// }
// int goglxQueryContext(Display* dpy, GLXContext ctx, int attribute, int* value) {
// 	return (*ptrglxQueryContext)(dpy, ctx, attribute, value);
// }
// void goglxSelectEvent(Display* dpy, GLXDrawable draw, unsigned long event_mask) {
// 	(*ptrglxSelectEvent)(dpy, draw, event_mask);
// }
// void goglxGetSelectedEvent(Display* dpy, GLXDrawable draw, unsigned long* event_mask) {
// 	(*ptrglxGetSelectedEvent)(dpy, draw, event_mask);
// }
// //  VERSION_1_4
// __GLXextFuncPtr goglxGetProcAddress(GLubyte* procName) {
// 	return (*ptrglxGetProcAddress)(procName);
// }
// //  glx
// void goglxRender() {
// 	(*ptrglxRender)();
// }
// void goglxRenderLarge() {
// 	(*ptrglxRenderLarge)();
// }
// void goglxCreateContext(GLint gc_id, GLint screen, GLint visual, GLint share_list) {
// 	(*ptrglxCreateContext)(gc_id, screen, visual, share_list);
// }
// void goglxDestroyContext(GLint context) {
// 	(*ptrglxDestroyContext)(context);
// }
// void goglxMakeCurrent(GLint drawable, GLint context) {
// 	(*ptrglxMakeCurrent)(drawable, context);
// }
// void goglxIsDirect(GLint dpy, GLint context) {
// 	(*ptrglxIsDirect)(dpy, context);
// }
// void goglxQueryVersion(GLint* major, GLint* minor) {
// 	(*ptrglxQueryVersion)(major, minor);
// }
// void goglxWaitGL(GLint context) {
// 	(*ptrglxWaitGL)(context);
// }
// void goglxWaitX() {
// 	(*ptrglxWaitX)();
// }
// void goglxCopyContext(GLint source, GLint dest, GLint mask) {
// 	(*ptrglxCopyContext)(source, dest, mask);
// }
// void goglxSwapBuffers(GLint drawable) {
// 	(*ptrglxSwapBuffers)(drawable);
// }
// void goglxUseXFont(GLint font, GLint first, GLint count, GLint list_base) {
// 	(*ptrglxUseXFont)(font, first, count, list_base);
// }
// void goglxCreateGLXPixmap(GLint visual, GLint pixmap, GLint glxpixmap) {
// 	(*ptrglxCreateGLXPixmap)(visual, pixmap, glxpixmap);
// }
// void goglxGetVisualConfigs() {
// 	(*ptrglxGetVisualConfigs)();
// }
// void goglxDestroyGLXPixmap(GLint pixmap) {
// 	(*ptrglxDestroyGLXPixmap)(pixmap);
// }
// void goglxVendorPrivate() {
// 	(*ptrglxVendorPrivate)();
// }
// void goglxVendorPrivateWithReply() {
// 	(*ptrglxVendorPrivateWithReply)();
// }
// void goglxQueryExtensionsString(GLint screen) {
// 	(*ptrglxQueryExtensionsString)(screen);
// }
// void goglxQueryServerString(GLint screen, GLint name) {
// 	(*ptrglxQueryServerString)(screen, name);
// }
// void goglxClientInfo() {
// 	(*ptrglxClientInfo)();
// }
// void goglxGetFBConfigs() {
// 	(*ptrglxGetFBConfigs)();
// }
// void goglxCreatePixmap(GLint config, GLint pixmap, GLint glxpixmap) {
// 	(*ptrglxCreatePixmap)(config, pixmap, glxpixmap);
// }
// void goglxDestroyPixmap(GLint glxpixmap) {
// 	(*ptrglxDestroyPixmap)(glxpixmap);
// }
// void goglxCreateNewContext(GLint config, GLint render_type, GLint share_list, GLint direct) {
// 	(*ptrglxCreateNewContext)(config, render_type, share_list, direct);
// }
// void goglxQueryContext() {
// 	(*ptrglxQueryContext)();
// }
// void goglxMakeContextCurrent(GLint drawable, GLint readdrawable, GLint context) {
// 	(*ptrglxMakeContextCurrent)(drawable, readdrawable, context);
// }
// void goglxCreatePbuffer(GLint config, GLint pbuffer) {
// 	(*ptrglxCreatePbuffer)(config, pbuffer);
// }
// void goglxDestroyPbuffer(GLint pbuffer) {
// 	(*ptrglxDestroyPbuffer)(pbuffer);
// }
// void goglxGetDrawableAttributes(GLint drawable) {
// 	(*ptrglxGetDrawableAttributes)(drawable);
// }
// void goglxChangeDrawableAttributes(GLint drawable) {
// 	(*ptrglxChangeDrawableAttributes)(drawable);
// }
// void goglxCreateWindow(GLint config, GLint window, GLint glxwindow) {
// 	(*ptrglxCreateWindow)(config, window, glxwindow);
// }
// void goglxDestroyWindow(GLint glxwindow) {
// 	(*ptrglxDestroyWindow)(glxwindow);
// }
// void goglxSwapIntervalSGI() {
// 	(*ptrglxSwapIntervalSGI)();
// }
// void goglxMakeCurrentReadSGI(GLint drawable, GLint readdrawable, GLint context) {
// 	(*ptrglxMakeCurrentReadSGI)(drawable, readdrawable, context);
// }
// void goglxCreateGLXVideoSourceSGIX(GLint dpy, GLint screen, GLint server, GLint path, GLint class, GLint node) {
// 	(*ptrglxCreateGLXVideoSourceSGIX)(dpy, screen, server, path, class, node);
// }
// void goglxDestroyGLXVideoSourceSGIX(GLint dpy, GLint glxvideosource) {
// 	(*ptrglxDestroyGLXVideoSourceSGIX)(dpy, glxvideosource);
// }
// void goglxQueryContextInfoEXT() {
// 	(*ptrglxQueryContextInfoEXT)();
// }
// void goglxGetFBConfigsSGIX() {
// 	(*ptrglxGetFBConfigsSGIX)();
// }
// void goglxCreateContextWithConfigSGIX(GLint gc_id, GLint screen, GLint config, GLint share_list) {
// 	(*ptrglxCreateContextWithConfigSGIX)(gc_id, screen, config, share_list);
// }
// void goglxCreateGLXPixmapWithConfigSGIX(GLint config, GLint pixmap, GLint glxpixmap) {
// 	(*ptrglxCreateGLXPixmapWithConfigSGIX)(config, pixmap, glxpixmap);
// }
// void goglxCreateGLXPbufferSGIX(GLint config, GLint pbuffer) {
// 	(*ptrglxCreateGLXPbufferSGIX)(config, pbuffer);
// }
// void goglxDestroyGLXPbufferSGIX(GLint pbuffer) {
// 	(*ptrglxDestroyGLXPbufferSGIX)(pbuffer);
// }
// void goglxChangeDrawableAttributesSGIX(GLint drawable) {
// 	(*ptrglxChangeDrawableAttributesSGIX)(drawable);
// }
// void goglxGetDrawableAttributesSGIX(GLint drawable) {
// 	(*ptrglxGetDrawableAttributesSGIX)(drawable);
// }
// void goglxJoinSwapGroupSGIX(GLint window, GLint group) {
// 	(*ptrglxJoinSwapGroupSGIX)(window, group);
// }
// void goglxBindSwapBarrierSGIX(GLint window, GLint barrier) {
// 	(*ptrglxBindSwapBarrierSGIX)(window, barrier);
// }
// void goglxQueryMaxSwapBarriersSGIX() {
// 	(*ptrglxQueryMaxSwapBarriersSGIX)();
// }
// GLXHyperpipeNetworkSGIX * goglxQueryHyperpipeNetworkSGIX(Display* dpy, int* npipes) {
// 	return (*ptrglxQueryHyperpipeNetworkSGIX)(dpy, npipes);
// }
// int goglxHyperpipeConfigSGIX(Display* dpy, int networkId, int npipes, GLXHyperpipeConfigSGIX* cfg, int* hpId) {
// 	return (*ptrglxHyperpipeConfigSGIX)(dpy, networkId, npipes, cfg, hpId);
// }
// GLXHyperpipeConfigSGIX * goglxQueryHyperpipeConfigSGIX(Display* dpy, int hpId, int* npipes) {
// 	return (*ptrglxQueryHyperpipeConfigSGIX)(dpy, hpId, npipes);
// }
// int goglxDestroyHyperpipeConfigSGIX(Display* dpy, int hpId) {
// 	return (*ptrglxDestroyHyperpipeConfigSGIX)(dpy, hpId);
// }
// int goglxBindHyperpipeSGIX(Display* dpy, int hpId) {
// 	return (*ptrglxBindHyperpipeSGIX)(dpy, hpId);
// }
// int goglxQueryHyperpipeBestAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, GLvoid* attribList, GLvoid* returnAttribList) {
// 	return (*ptrglxQueryHyperpipeBestAttribSGIX)(dpy, timeSlice, attrib, size, attribList, returnAttribList);
// }
// int goglxHyperpipeAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, void* attribList) {
// 	return (*ptrglxHyperpipeAttribSGIX)(dpy, timeSlice, attrib, size, attribList);
// }
// int goglxQueryHyperpipeAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, void* returnAttribList) {
// 	return (*ptrglxQueryHyperpipeAttribSGIX)(dpy, timeSlice, attrib, size, returnAttribList);
// }
// 
// int init_ARB_create_context() {
// 	ptrglxCreateContextAttribsARB = goglxGetProcAddress("glxCreateContextAttribsARB");
// 	if(ptrglxCreateContextAttribsARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_get_proc_address() {
// 	ptrglxGetProcAddressARB = goglxGetProcAddress("glxGetProcAddressARB");
// 	if(ptrglxGetProcAddressARB == NULL) return 1;
// 	return 0;
// }
// int init_EXT_import_context() {
// 	ptrglxGetCurrentDisplayEXT = goglxGetProcAddress("glxGetCurrentDisplayEXT");
// 	if(ptrglxGetCurrentDisplayEXT == NULL) return 1;
// 	ptrglxQueryContextInfoEXT = goglxGetProcAddress("glxQueryContextInfoEXT");
// 	if(ptrglxQueryContextInfoEXT == NULL) return 1;
// 	ptrglxGetContextIDEXT = goglxGetProcAddress("glxGetContextIDEXT");
// 	if(ptrglxGetContextIDEXT == NULL) return 1;
// 	ptrglxImportContextEXT = goglxGetProcAddress("glxImportContextEXT");
// 	if(ptrglxImportContextEXT == NULL) return 1;
// 	ptrglxFreeContextEXT = goglxGetProcAddress("glxFreeContextEXT");
// 	if(ptrglxFreeContextEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_swap_control() {
// 	ptrglxSwapIntervalEXT = goglxGetProcAddress("glxSwapIntervalEXT");
// 	if(ptrglxSwapIntervalEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_texture_from_pixmap() {
// 	ptrglxBindTexImageEXT = goglxGetProcAddress("glxBindTexImageEXT");
// 	if(ptrglxBindTexImageEXT == NULL) return 1;
// 	ptrglxReleaseTexImageEXT = goglxGetProcAddress("glxReleaseTexImageEXT");
// 	if(ptrglxReleaseTexImageEXT == NULL) return 1;
// 	return 0;
// }
// int init_MESA_agp_offset() {
// 	ptrglxGetAGPOffsetMESA = goglxGetProcAddress("glxGetAGPOffsetMESA");
// 	if(ptrglxGetAGPOffsetMESA == NULL) return 1;
// 	return 0;
// }
// int init_MESA_copy_sub_buffer() {
// 	ptrglxCopySubBufferMESA = goglxGetProcAddress("glxCopySubBufferMESA");
// 	if(ptrglxCopySubBufferMESA == NULL) return 1;
// 	return 0;
// }
// int init_MESA_pixmap_colormap() {
// 	ptrglxCreateGLXPixmapMESA = goglxGetProcAddress("glxCreateGLXPixmapMESA");
// 	if(ptrglxCreateGLXPixmapMESA == NULL) return 1;
// 	return 0;
// }
// int init_MESA_release_buffers() {
// 	ptrglxReleaseBuffersMESA = goglxGetProcAddress("glxReleaseBuffersMESA");
// 	if(ptrglxReleaseBuffersMESA == NULL) return 1;
// 	return 0;
// }
// int init_MESA_set_3dfx_mode() {
// 	ptrglxSet3DfxModeMESA = goglxGetProcAddress("glxSet3DfxModeMESA");
// 	if(ptrglxSet3DfxModeMESA == NULL) return 1;
// 	return 0;
// }
// int init_NV_copy_image() {
// 	ptrglxCopyImageSubDataNV = goglxGetProcAddress("glxCopyImageSubDataNV");
// 	if(ptrglxCopyImageSubDataNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_present_video() {
// 	ptrglxEnumerateVideoDevicesNV = goglxGetProcAddress("glxEnumerateVideoDevicesNV");
// 	if(ptrglxEnumerateVideoDevicesNV == NULL) return 1;
// 	ptrglxBindVideoDeviceNV = goglxGetProcAddress("glxBindVideoDeviceNV");
// 	if(ptrglxBindVideoDeviceNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_swap_group() {
// 	ptrglxJoinSwapGroupNV = goglxGetProcAddress("glxJoinSwapGroupNV");
// 	if(ptrglxJoinSwapGroupNV == NULL) return 1;
// 	ptrglxBindSwapBarrierNV = goglxGetProcAddress("glxBindSwapBarrierNV");
// 	if(ptrglxBindSwapBarrierNV == NULL) return 1;
// 	ptrglxQuerySwapGroupNV = goglxGetProcAddress("glxQuerySwapGroupNV");
// 	if(ptrglxQuerySwapGroupNV == NULL) return 1;
// 	ptrglxQueryMaxSwapGroupsNV = goglxGetProcAddress("glxQueryMaxSwapGroupsNV");
// 	if(ptrglxQueryMaxSwapGroupsNV == NULL) return 1;
// 	ptrglxQueryFrameCountNV = goglxGetProcAddress("glxQueryFrameCountNV");
// 	if(ptrglxQueryFrameCountNV == NULL) return 1;
// 	ptrglxResetFrameCountNV = goglxGetProcAddress("glxResetFrameCountNV");
// 	if(ptrglxResetFrameCountNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_video_capture() {
// 	ptrglxBindVideoCaptureDeviceNV = goglxGetProcAddress("glxBindVideoCaptureDeviceNV");
// 	if(ptrglxBindVideoCaptureDeviceNV == NULL) return 1;
// 	ptrglxEnumerateVideoCaptureDevicesNV = goglxGetProcAddress("glxEnumerateVideoCaptureDevicesNV");
// 	if(ptrglxEnumerateVideoCaptureDevicesNV == NULL) return 1;
// 	ptrglxLockVideoCaptureDeviceNV = goglxGetProcAddress("glxLockVideoCaptureDeviceNV");
// 	if(ptrglxLockVideoCaptureDeviceNV == NULL) return 1;
// 	ptrglxQueryVideoCaptureDeviceNV = goglxGetProcAddress("glxQueryVideoCaptureDeviceNV");
// 	if(ptrglxQueryVideoCaptureDeviceNV == NULL) return 1;
// 	ptrglxReleaseVideoCaptureDeviceNV = goglxGetProcAddress("glxReleaseVideoCaptureDeviceNV");
// 	if(ptrglxReleaseVideoCaptureDeviceNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_video_output() {
// 	ptrglxGetVideoDeviceNV = goglxGetProcAddress("glxGetVideoDeviceNV");
// 	if(ptrglxGetVideoDeviceNV == NULL) return 1;
// 	ptrglxReleaseVideoDeviceNV = goglxGetProcAddress("glxReleaseVideoDeviceNV");
// 	if(ptrglxReleaseVideoDeviceNV == NULL) return 1;
// 	ptrglxBindVideoImageNV = goglxGetProcAddress("glxBindVideoImageNV");
// 	if(ptrglxBindVideoImageNV == NULL) return 1;
// 	ptrglxReleaseVideoImageNV = goglxGetProcAddress("glxReleaseVideoImageNV");
// 	if(ptrglxReleaseVideoImageNV == NULL) return 1;
// 	ptrglxSendPbufferToVideoNV = goglxGetProcAddress("glxSendPbufferToVideoNV");
// 	if(ptrglxSendPbufferToVideoNV == NULL) return 1;
// 	ptrglxGetVideoInfoNV = goglxGetProcAddress("glxGetVideoInfoNV");
// 	if(ptrglxGetVideoInfoNV == NULL) return 1;
// 	return 0;
// }
// int init_OML_sync_control() {
// 	ptrglxGetSyncValuesOML = goglxGetProcAddress("glxGetSyncValuesOML");
// 	if(ptrglxGetSyncValuesOML == NULL) return 1;
// 	ptrglxGetMscRateOML = goglxGetProcAddress("glxGetMscRateOML");
// 	if(ptrglxGetMscRateOML == NULL) return 1;
// 	ptrglxSwapBuffersMscOML = goglxGetProcAddress("glxSwapBuffersMscOML");
// 	if(ptrglxSwapBuffersMscOML == NULL) return 1;
// 	ptrglxWaitForMscOML = goglxGetProcAddress("glxWaitForMscOML");
// 	if(ptrglxWaitForMscOML == NULL) return 1;
// 	ptrglxWaitForSbcOML = goglxGetProcAddress("glxWaitForSbcOML");
// 	if(ptrglxWaitForSbcOML == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_dmbuffer() {
// 	ptrglxAssociateDMPbufferSGIX = goglxGetProcAddress("glxAssociateDMPbufferSGIX");
// 	if(ptrglxAssociateDMPbufferSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_fbconfig() {
// 	ptrglxGetFBConfigAttribSGIX = goglxGetProcAddress("glxGetFBConfigAttribSGIX");
// 	if(ptrglxGetFBConfigAttribSGIX == NULL) return 1;
// 	ptrglxChooseFBConfigSGIX = goglxGetProcAddress("glxChooseFBConfigSGIX");
// 	if(ptrglxChooseFBConfigSGIX == NULL) return 1;
// 	ptrglxCreateGLXPixmapWithConfigSGIX = goglxGetProcAddress("glxCreateGLXPixmapWithConfigSGIX");
// 	if(ptrglxCreateGLXPixmapWithConfigSGIX == NULL) return 1;
// 	ptrglxCreateContextWithConfigSGIX = goglxGetProcAddress("glxCreateContextWithConfigSGIX");
// 	if(ptrglxCreateContextWithConfigSGIX == NULL) return 1;
// 	ptrglxGetVisualFromFBConfigSGIX = goglxGetProcAddress("glxGetVisualFromFBConfigSGIX");
// 	if(ptrglxGetVisualFromFBConfigSGIX == NULL) return 1;
// 	ptrglxGetFBConfigFromVisualSGIX = goglxGetProcAddress("glxGetFBConfigFromVisualSGIX");
// 	if(ptrglxGetFBConfigFromVisualSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_hyperpipe() {
// 	ptrglxQueryHyperpipeNetworkSGIX = goglxGetProcAddress("glxQueryHyperpipeNetworkSGIX");
// 	if(ptrglxQueryHyperpipeNetworkSGIX == NULL) return 1;
// 	ptrglxHyperpipeConfigSGIX = goglxGetProcAddress("glxHyperpipeConfigSGIX");
// 	if(ptrglxHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglxQueryHyperpipeConfigSGIX = goglxGetProcAddress("glxQueryHyperpipeConfigSGIX");
// 	if(ptrglxQueryHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglxDestroyHyperpipeConfigSGIX = goglxGetProcAddress("glxDestroyHyperpipeConfigSGIX");
// 	if(ptrglxDestroyHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglxBindHyperpipeSGIX = goglxGetProcAddress("glxBindHyperpipeSGIX");
// 	if(ptrglxBindHyperpipeSGIX == NULL) return 1;
// 	ptrglxQueryHyperpipeBestAttribSGIX = goglxGetProcAddress("glxQueryHyperpipeBestAttribSGIX");
// 	if(ptrglxQueryHyperpipeBestAttribSGIX == NULL) return 1;
// 	ptrglxHyperpipeAttribSGIX = goglxGetProcAddress("glxHyperpipeAttribSGIX");
// 	if(ptrglxHyperpipeAttribSGIX == NULL) return 1;
// 	ptrglxQueryHyperpipeAttribSGIX = goglxGetProcAddress("glxQueryHyperpipeAttribSGIX");
// 	if(ptrglxQueryHyperpipeAttribSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_pbuffer() {
// 	ptrglxCreateGLXPbufferSGIX = goglxGetProcAddress("glxCreateGLXPbufferSGIX");
// 	if(ptrglxCreateGLXPbufferSGIX == NULL) return 1;
// 	ptrglxDestroyGLXPbufferSGIX = goglxGetProcAddress("glxDestroyGLXPbufferSGIX");
// 	if(ptrglxDestroyGLXPbufferSGIX == NULL) return 1;
// 	ptrglxQueryGLXPbufferSGIX = goglxGetProcAddress("glxQueryGLXPbufferSGIX");
// 	if(ptrglxQueryGLXPbufferSGIX == NULL) return 1;
// 	ptrglxSelectEventSGIX = goglxGetProcAddress("glxSelectEventSGIX");
// 	if(ptrglxSelectEventSGIX == NULL) return 1;
// 	ptrglxGetSelectedEventSGIX = goglxGetProcAddress("glxGetSelectedEventSGIX");
// 	if(ptrglxGetSelectedEventSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_swap_barrier() {
// 	ptrglxBindSwapBarrierSGIX = goglxGetProcAddress("glxBindSwapBarrierSGIX");
// 	if(ptrglxBindSwapBarrierSGIX == NULL) return 1;
// 	ptrglxQueryMaxSwapBarriersSGIX = goglxGetProcAddress("glxQueryMaxSwapBarriersSGIX");
// 	if(ptrglxQueryMaxSwapBarriersSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_swap_group() {
// 	ptrglxJoinSwapGroupSGIX = goglxGetProcAddress("glxJoinSwapGroupSGIX");
// 	if(ptrglxJoinSwapGroupSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_video_resize() {
// 	ptrglxBindChannelToWindowSGIX = goglxGetProcAddress("glxBindChannelToWindowSGIX");
// 	if(ptrglxBindChannelToWindowSGIX == NULL) return 1;
// 	ptrglxChannelRectSGIX = goglxGetProcAddress("glxChannelRectSGIX");
// 	if(ptrglxChannelRectSGIX == NULL) return 1;
// 	ptrglxQueryChannelRectSGIX = goglxGetProcAddress("glxQueryChannelRectSGIX");
// 	if(ptrglxQueryChannelRectSGIX == NULL) return 1;
// 	ptrglxQueryChannelDeltasSGIX = goglxGetProcAddress("glxQueryChannelDeltasSGIX");
// 	if(ptrglxQueryChannelDeltasSGIX == NULL) return 1;
// 	ptrglxChannelRectSyncSGIX = goglxGetProcAddress("glxChannelRectSyncSGIX");
// 	if(ptrglxChannelRectSyncSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_video_source() {
// 	ptrglxCreateGLXVideoSourceSGIX = goglxGetProcAddress("glxCreateGLXVideoSourceSGIX");
// 	if(ptrglxCreateGLXVideoSourceSGIX == NULL) return 1;
// 	ptrglxDestroyGLXVideoSourceSGIX = goglxGetProcAddress("glxDestroyGLXVideoSourceSGIX");
// 	if(ptrglxDestroyGLXVideoSourceSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGI_cushion() {
// 	ptrglxCushionSGI = goglxGetProcAddress("glxCushionSGI");
// 	if(ptrglxCushionSGI == NULL) return 1;
// 	return 0;
// }
// int init_SGI_make_current_read() {
// 	ptrglxMakeCurrentReadSGI = goglxGetProcAddress("glxMakeCurrentReadSGI");
// 	if(ptrglxMakeCurrentReadSGI == NULL) return 1;
// 	ptrglxGetCurrentReadDrawableSGI = goglxGetProcAddress("glxGetCurrentReadDrawableSGI");
// 	if(ptrglxGetCurrentReadDrawableSGI == NULL) return 1;
// 	return 0;
// }
// int init_SGI_swap_control() {
// 	ptrglxSwapIntervalSGI = goglxGetProcAddress("glxSwapIntervalSGI");
// 	if(ptrglxSwapIntervalSGI == NULL) return 1;
// 	return 0;
// }
// int init_SGI_video_sync() {
// 	ptrglxGetVideoSyncSGI = goglxGetProcAddress("glxGetVideoSyncSGI");
// 	if(ptrglxGetVideoSyncSGI == NULL) return 1;
// 	ptrglxWaitVideoSyncSGI = goglxGetProcAddress("glxWaitVideoSyncSGI");
// 	if(ptrglxWaitVideoSyncSGI == NULL) return 1;
// 	return 0;
// }
// int init_SUN_get_transparent_index() {
// 	ptrglxGetTransparentIndexSUN = goglxGetProcAddress("glxGetTransparentIndexSUN");
// 	if(ptrglxGetTransparentIndexSUN == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_3() {
// 	ptrglxGetFBConfigs = goglxGetProcAddress("glxGetFBConfigs");
// 	if(ptrglxGetFBConfigs == NULL) return 1;
// 	ptrglxChooseFBConfig = goglxGetProcAddress("glxChooseFBConfig");
// 	if(ptrglxChooseFBConfig == NULL) return 1;
// 	ptrglxGetFBConfigAttrib = goglxGetProcAddress("glxGetFBConfigAttrib");
// 	if(ptrglxGetFBConfigAttrib == NULL) return 1;
// 	ptrglxGetVisualFromFBConfig = goglxGetProcAddress("glxGetVisualFromFBConfig");
// 	if(ptrglxGetVisualFromFBConfig == NULL) return 1;
// 	ptrglxCreateWindow = goglxGetProcAddress("glxCreateWindow");
// 	if(ptrglxCreateWindow == NULL) return 1;
// 	ptrglxDestroyWindow = goglxGetProcAddress("glxDestroyWindow");
// 	if(ptrglxDestroyWindow == NULL) return 1;
// 	ptrglxCreatePixmap = goglxGetProcAddress("glxCreatePixmap");
// 	if(ptrglxCreatePixmap == NULL) return 1;
// 	ptrglxDestroyPixmap = goglxGetProcAddress("glxDestroyPixmap");
// 	if(ptrglxDestroyPixmap == NULL) return 1;
// 	ptrglxCreatePbuffer = goglxGetProcAddress("glxCreatePbuffer");
// 	if(ptrglxCreatePbuffer == NULL) return 1;
// 	ptrglxDestroyPbuffer = goglxGetProcAddress("glxDestroyPbuffer");
// 	if(ptrglxDestroyPbuffer == NULL) return 1;
// 	ptrglxQueryDrawable = goglxGetProcAddress("glxQueryDrawable");
// 	if(ptrglxQueryDrawable == NULL) return 1;
// 	ptrglxCreateNewContext = goglxGetProcAddress("glxCreateNewContext");
// 	if(ptrglxCreateNewContext == NULL) return 1;
// 	ptrglxMakeContextCurrent = goglxGetProcAddress("glxMakeContextCurrent");
// 	if(ptrglxMakeContextCurrent == NULL) return 1;
// 	ptrglxGetCurrentReadDrawable = goglxGetProcAddress("glxGetCurrentReadDrawable");
// 	if(ptrglxGetCurrentReadDrawable == NULL) return 1;
// 	ptrglxGetCurrentDisplay = goglxGetProcAddress("glxGetCurrentDisplay");
// 	if(ptrglxGetCurrentDisplay == NULL) return 1;
// 	ptrglxQueryContext = goglxGetProcAddress("glxQueryContext");
// 	if(ptrglxQueryContext == NULL) return 1;
// 	ptrglxSelectEvent = goglxGetProcAddress("glxSelectEvent");
// 	if(ptrglxSelectEvent == NULL) return 1;
// 	ptrglxGetSelectedEvent = goglxGetProcAddress("glxGetSelectedEvent");
// 	if(ptrglxGetSelectedEvent == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_4() {
// 	ptrglxGetProcAddress = goglxGetProcAddress("glxGetProcAddress");
// 	if(ptrglxGetProcAddress == NULL) return 1;
// 	return 0;
// }
// int init_glx() {
// 	ptrglxRender = goglxGetProcAddress("glxRender");
// 	if(ptrglxRender == NULL) return 1;
// 	ptrglxRenderLarge = goglxGetProcAddress("glxRenderLarge");
// 	if(ptrglxRenderLarge == NULL) return 1;
// 	ptrglxCreateContext = goglxGetProcAddress("glxCreateContext");
// 	if(ptrglxCreateContext == NULL) return 1;
// 	ptrglxDestroyContext = goglxGetProcAddress("glxDestroyContext");
// 	if(ptrglxDestroyContext == NULL) return 1;
// 	ptrglxMakeCurrent = goglxGetProcAddress("glxMakeCurrent");
// 	if(ptrglxMakeCurrent == NULL) return 1;
// 	ptrglxIsDirect = goglxGetProcAddress("glxIsDirect");
// 	if(ptrglxIsDirect == NULL) return 1;
// 	ptrglxQueryVersion = goglxGetProcAddress("glxQueryVersion");
// 	if(ptrglxQueryVersion == NULL) return 1;
// 	ptrglxWaitGL = goglxGetProcAddress("glxWaitGL");
// 	if(ptrglxWaitGL == NULL) return 1;
// 	ptrglxWaitX = goglxGetProcAddress("glxWaitX");
// 	if(ptrglxWaitX == NULL) return 1;
// 	ptrglxCopyContext = goglxGetProcAddress("glxCopyContext");
// 	if(ptrglxCopyContext == NULL) return 1;
// 	ptrglxSwapBuffers = goglxGetProcAddress("glxSwapBuffers");
// 	if(ptrglxSwapBuffers == NULL) return 1;
// 	ptrglxUseXFont = goglxGetProcAddress("glxUseXFont");
// 	if(ptrglxUseXFont == NULL) return 1;
// 	ptrglxCreateGLXPixmap = goglxGetProcAddress("glxCreateGLXPixmap");
// 	if(ptrglxCreateGLXPixmap == NULL) return 1;
// 	ptrglxGetVisualConfigs = goglxGetProcAddress("glxGetVisualConfigs");
// 	if(ptrglxGetVisualConfigs == NULL) return 1;
// 	ptrglxDestroyGLXPixmap = goglxGetProcAddress("glxDestroyGLXPixmap");
// 	if(ptrglxDestroyGLXPixmap == NULL) return 1;
// 	ptrglxVendorPrivate = goglxGetProcAddress("glxVendorPrivate");
// 	if(ptrglxVendorPrivate == NULL) return 1;
// 	ptrglxVendorPrivateWithReply = goglxGetProcAddress("glxVendorPrivateWithReply");
// 	if(ptrglxVendorPrivateWithReply == NULL) return 1;
// 	ptrglxQueryExtensionsString = goglxGetProcAddress("glxQueryExtensionsString");
// 	if(ptrglxQueryExtensionsString == NULL) return 1;
// 	ptrglxQueryServerString = goglxGetProcAddress("glxQueryServerString");
// 	if(ptrglxQueryServerString == NULL) return 1;
// 	ptrglxClientInfo = goglxGetProcAddress("glxClientInfo");
// 	if(ptrglxClientInfo == NULL) return 1;
// 	ptrglxGetFBConfigs = goglxGetProcAddress("glxGetFBConfigs");
// 	if(ptrglxGetFBConfigs == NULL) return 1;
// 	ptrglxCreatePixmap = goglxGetProcAddress("glxCreatePixmap");
// 	if(ptrglxCreatePixmap == NULL) return 1;
// 	ptrglxDestroyPixmap = goglxGetProcAddress("glxDestroyPixmap");
// 	if(ptrglxDestroyPixmap == NULL) return 1;
// 	ptrglxCreateNewContext = goglxGetProcAddress("glxCreateNewContext");
// 	if(ptrglxCreateNewContext == NULL) return 1;
// 	ptrglxQueryContext = goglxGetProcAddress("glxQueryContext");
// 	if(ptrglxQueryContext == NULL) return 1;
// 	ptrglxMakeContextCurrent = goglxGetProcAddress("glxMakeContextCurrent");
// 	if(ptrglxMakeContextCurrent == NULL) return 1;
// 	ptrglxCreatePbuffer = goglxGetProcAddress("glxCreatePbuffer");
// 	if(ptrglxCreatePbuffer == NULL) return 1;
// 	ptrglxDestroyPbuffer = goglxGetProcAddress("glxDestroyPbuffer");
// 	if(ptrglxDestroyPbuffer == NULL) return 1;
// 	ptrglxGetDrawableAttributes = goglxGetProcAddress("glxGetDrawableAttributes");
// 	if(ptrglxGetDrawableAttributes == NULL) return 1;
// 	ptrglxChangeDrawableAttributes = goglxGetProcAddress("glxChangeDrawableAttributes");
// 	if(ptrglxChangeDrawableAttributes == NULL) return 1;
// 	ptrglxCreateWindow = goglxGetProcAddress("glxCreateWindow");
// 	if(ptrglxCreateWindow == NULL) return 1;
// 	ptrglxDestroyWindow = goglxGetProcAddress("glxDestroyWindow");
// 	if(ptrglxDestroyWindow == NULL) return 1;
// 	ptrglxSwapIntervalSGI = goglxGetProcAddress("glxSwapIntervalSGI");
// 	if(ptrglxSwapIntervalSGI == NULL) return 1;
// 	ptrglxMakeCurrentReadSGI = goglxGetProcAddress("glxMakeCurrentReadSGI");
// 	if(ptrglxMakeCurrentReadSGI == NULL) return 1;
// 	ptrglxCreateGLXVideoSourceSGIX = goglxGetProcAddress("glxCreateGLXVideoSourceSGIX");
// 	if(ptrglxCreateGLXVideoSourceSGIX == NULL) return 1;
// 	ptrglxDestroyGLXVideoSourceSGIX = goglxGetProcAddress("glxDestroyGLXVideoSourceSGIX");
// 	if(ptrglxDestroyGLXVideoSourceSGIX == NULL) return 1;
// 	ptrglxQueryContextInfoEXT = goglxGetProcAddress("glxQueryContextInfoEXT");
// 	if(ptrglxQueryContextInfoEXT == NULL) return 1;
// 	ptrglxGetFBConfigsSGIX = goglxGetProcAddress("glxGetFBConfigsSGIX");
// 	if(ptrglxGetFBConfigsSGIX == NULL) return 1;
// 	ptrglxCreateContextWithConfigSGIX = goglxGetProcAddress("glxCreateContextWithConfigSGIX");
// 	if(ptrglxCreateContextWithConfigSGIX == NULL) return 1;
// 	ptrglxCreateGLXPixmapWithConfigSGIX = goglxGetProcAddress("glxCreateGLXPixmapWithConfigSGIX");
// 	if(ptrglxCreateGLXPixmapWithConfigSGIX == NULL) return 1;
// 	ptrglxCreateGLXPbufferSGIX = goglxGetProcAddress("glxCreateGLXPbufferSGIX");
// 	if(ptrglxCreateGLXPbufferSGIX == NULL) return 1;
// 	ptrglxDestroyGLXPbufferSGIX = goglxGetProcAddress("glxDestroyGLXPbufferSGIX");
// 	if(ptrglxDestroyGLXPbufferSGIX == NULL) return 1;
// 	ptrglxChangeDrawableAttributesSGIX = goglxGetProcAddress("glxChangeDrawableAttributesSGIX");
// 	if(ptrglxChangeDrawableAttributesSGIX == NULL) return 1;
// 	ptrglxGetDrawableAttributesSGIX = goglxGetProcAddress("glxGetDrawableAttributesSGIX");
// 	if(ptrglxGetDrawableAttributesSGIX == NULL) return 1;
// 	ptrglxJoinSwapGroupSGIX = goglxGetProcAddress("glxJoinSwapGroupSGIX");
// 	if(ptrglxJoinSwapGroupSGIX == NULL) return 1;
// 	ptrglxBindSwapBarrierSGIX = goglxGetProcAddress("glxBindSwapBarrierSGIX");
// 	if(ptrglxBindSwapBarrierSGIX == NULL) return 1;
// 	ptrglxQueryMaxSwapBarriersSGIX = goglxGetProcAddress("glxQueryMaxSwapBarriersSGIX");
// 	if(ptrglxQueryMaxSwapBarriersSGIX == NULL) return 1;
// 	ptrglxQueryHyperpipeNetworkSGIX = goglxGetProcAddress("glxQueryHyperpipeNetworkSGIX");
// 	if(ptrglxQueryHyperpipeNetworkSGIX == NULL) return 1;
// 	ptrglxHyperpipeConfigSGIX = goglxGetProcAddress("glxHyperpipeConfigSGIX");
// 	if(ptrglxHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglxQueryHyperpipeConfigSGIX = goglxGetProcAddress("glxQueryHyperpipeConfigSGIX");
// 	if(ptrglxQueryHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglxDestroyHyperpipeConfigSGIX = goglxGetProcAddress("glxDestroyHyperpipeConfigSGIX");
// 	if(ptrglxDestroyHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglxBindHyperpipeSGIX = goglxGetProcAddress("glxBindHyperpipeSGIX");
// 	if(ptrglxBindHyperpipeSGIX == NULL) return 1;
// 	ptrglxQueryHyperpipeBestAttribSGIX = goglxGetProcAddress("glxQueryHyperpipeBestAttribSGIX");
// 	if(ptrglxQueryHyperpipeBestAttribSGIX == NULL) return 1;
// 	ptrglxHyperpipeAttribSGIX = goglxGetProcAddress("glxHyperpipeAttribSGIX");
// 	if(ptrglxHyperpipeAttribSGIX == NULL) return 1;
// 	ptrglxQueryHyperpipeAttribSGIX = goglxGetProcAddress("glxQueryHyperpipeAttribSGIX");
// 	if(ptrglxQueryHyperpipeAttribSGIX == NULL) return 1;
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

// 3DFX_multisample
const (
	SAMPLES_3DFX = 0x8051
	SAMPLE_BUFFERS_3DFX = 0x8050
)
// AMD_gpu_association
const (
	GPU_CLOCK_AMD = 0x21A4
	GPU_FASTEST_TARGET_GPUS_AMD = 0x21A2
	GPU_NUM_PIPES_AMD = 0x21A5
	GPU_NUM_RB_AMD = 0x21A7
	GPU_NUM_SIMD_AMD = 0x21A6
	GPU_NUM_SPI_AMD = 0x21A8
	GPU_OPENGL_VERSION_STRING_AMD = 0x1F02
	GPU_RAM_AMD = 0x21A3
	GPU_RENDERER_STRING_AMD = 0x1F01
	GPU_VENDOR_AMD = 0x1F00
)
// ARB_create_context
const (
	CONTEXT_DEBUG_BIT_ARB = 0x00000001
	CONTEXT_FLAGS_ARB = 0x2094
	CONTEXT_FORWARD_COMPATIBLE_BIT_ARB = 0x00000002
	CONTEXT_MAJOR_VERSION_ARB = 0x2091
	CONTEXT_MINOR_VERSION_ARB = 0x2092
)
// ARB_create_context_profile
const (
	CONTEXT_COMPATIBILITY_PROFILE_BIT_ARB = 0x00000002
	CONTEXT_CORE_PROFILE_BIT_ARB = 0x00000001
	CONTEXT_PROFILE_MASK_ARB = 0x9126
)
// ARB_create_context_robustness
const (
	CONTEXT_RESET_NOTIFICATION_STRATEGY_ARB = 0x8256
	CONTEXT_ROBUST_ACCESS_BIT_ARB = 0x00000004
	LOSE_CONTEXT_ON_RESET_ARB = 0x8252
	NO_RESET_NOTIFICATION_ARB = 0x8261
)
// ARB_fbconfig_float
const (
	RGBA_FLOAT_BIT_ARB = 0x00000004
	RGBA_FLOAT_TYPE_ARB = 0x20B9
)
// ARB_framebuffer_sRGB
const (
	FRAMEBUFFER_SRGB_CAPABLE_ARB = 0x20B2
)
// ARB_get_proc_address
const (
)
// ARB_multisample
const (
	SAMPLES_ARB = 100001
	SAMPLE_BUFFERS_ARB = 100000
)
// ARB_vertex_buffer_object
const (
	CONTEXT_ALLOW_BUFFER_BYTE_ORDER_MISMATCH_ARB = 0x2095
)
// EXT_buffer_age
const (
	BACK_BUFFER_AGE_EXT = 0x20F4
)
// EXT_create_context_es2_profile
const (
	CONTEXT_ES2_PROFILE_BIT_EXT = 0x00000004
)
// EXT_fbconfig_packed_float
const (
	RGBA_UNSIGNED_FLOAT_BIT_EXT = 0x00000008
	RGBA_UNSIGNED_FLOAT_TYPE_EXT = 0x20B1
)
// EXT_framebuffer_sRGB
const (
	FRAMEBUFFER_SRGB_CAPABLE_EXT = 0x20B2
)
// EXT_import_context
const (
	SHARE_CONTEXT_EXT = 0x800A
	VISUAL_ID_EXT = 0x800B
)
// EXT_swap_control
const (
	MAX_SWAP_INTERVAL_EXT = 0x20F2
	SWAP_INTERVAL_EXT = 0x20F1
)
// EXT_swap_control_tear
const (
	LATE_SWAPS_TEAR_EXT = 0x20F3
)
// EXT_texture_from_pixmap
const (
	AUX0_EXT = 0x20E2
	AUX1_EXT = 0x20E3
	AUX2_EXT = 0x20E4
	AUX3_EXT = 0x20E5
	AUX4_EXT = 0x20E6
	AUX5_EXT = 0x20E7
	AUX6_EXT = 0x20E8
	AUX7_EXT = 0x20E9
	AUX8_EXT = 0x20EA
	AUX9_EXT = 0x20EB
	BACK_EXT = 0x20E0
	BACK_LEFT_EXT = 0x20E0
	BACK_RIGHT_EXT = 0x20E1
	BIND_TO_MIPMAP_TEXTURE_EXT = 0x20D2
	BIND_TO_TEXTURE_RGBA_EXT = 0x20D1
	BIND_TO_TEXTURE_RGB_EXT = 0x20D0
	BIND_TO_TEXTURE_TARGETS_EXT = 0x20D3
	FRONT_EXT = 0x20DE
	FRONT_LEFT_EXT = 0x20DE
	FRONT_RIGHT_EXT = 0x20DF
	MIPMAP_TEXTURE_EXT = 0x20D7
	TEXTURE_1D_BIT_EXT = 0x00000001
	TEXTURE_1D_EXT = 0x20DB
	TEXTURE_2D_BIT_EXT = 0x00000002
	TEXTURE_2D_EXT = 0x20DC
	TEXTURE_FORMAT_EXT = 0x20D5
	TEXTURE_FORMAT_NONE_EXT = 0x20D8
	TEXTURE_FORMAT_RGBA_EXT = 0x20DA
	TEXTURE_FORMAT_RGB_EXT = 0x20D9
	TEXTURE_RECTANGLE_BIT_EXT = 0x00000004
	TEXTURE_RECTANGLE_EXT = 0x20DD
	TEXTURE_TARGET_EXT = 0x20D6
	Y_INVERTED_EXT = 0x20D4
)
// EXT_visual_info
const (
	DIRECT_COLOR_EXT = 0x8003
	GRAY_SCALE_EXT = 0x8006
	PSEUDO_COLOR_EXT = 0x8004
	STATIC_COLOR_EXT = 0x8005
	STATIC_GRAY_EXT = 0x8007
	TRANSPARENT_ALPHA_VALUE_EXT = 0x28
	TRANSPARENT_BLUE_VALUE_EXT = 0x27
	TRANSPARENT_GREEN_VALUE_EXT = 0x26
	TRANSPARENT_INDEX_EXT = 0x8009
	TRANSPARENT_INDEX_VALUE_EXT = 0x24
	TRANSPARENT_RED_VALUE_EXT = 0x25
	TRANSPARENT_RGB_EXT = 0x8008
	TRANSPARENT_TYPE_EXT = 0x23
	TRUE_COLOR_EXT = 0x8002
	X_VISUAL_TYPE_EXT = 0x22
)
// EXT_visual_rating
const (
	NON_CONFORMANT_VISUAL_EXT = 0x800D
	SLOW_VISUAL_EXT = 0x8001
	VISUAL_CAVEAT_EXT = 0x20
)
// INTEL_swap_event
const (
	BUFFER_SWAP_COMPLETE_INTEL_MASK = 0x04000000
	COPY_COMPLETE_INTEL = 0x8181
	EXCHANGE_COMPLETE_INTEL = 0x8180
	FLIP_COMPLETE_INTEL = 0x8182
)
// MESA_agp_offset
const (
)
// MESA_copy_sub_buffer
const (
)
// MESA_pixmap_colormap
const (
)
// MESA_release_buffers
const (
)
// MESA_set_3dfx_mode
const (
	X3DFX_FULLSCREEN_MODE_MESA = 0x2
	X3DFX_WINDOW_MODE_MESA = 0x1
)
// NV_copy_image
const (
)
// NV_float_buffer
const (
	FLOAT_COMPONENTS_NV = 0x20B0
)
// NV_multisample_coverage
const (
	COLOR_SAMPLES_NV = 0x20B3
	COVERAGE_SAMPLES_NV = 100001
)
// NV_present_video
const (
	NUM_VIDEO_SLOTS_NV = 0x20F0
)
// NV_swap_group
const (
)
// NV_video_capture
const (
	DEVICE_ID_NV = 0x20CD
	NUM_VIDEO_CAPTURE_SLOTS_NV = 0x20CF
	UNIQUE_ID_NV = 0x20CE
)
// NV_video_out
const (
	VIDEO_OUT_ALPHA_NV = 0x20C4
	VIDEO_OUT_COLOR_AND_ALPHA_NV = 0x20C6
	VIDEO_OUT_COLOR_AND_DEPTH_NV = 0x20C7
	VIDEO_OUT_COLOR_NV = 0x20C3
	VIDEO_OUT_DEPTH_NV = 0x20C5
	VIDEO_OUT_FIELD_1_NV = 0x20C9
	VIDEO_OUT_FIELD_2_NV = 0x20CA
	VIDEO_OUT_FRAME_NV = 0x20C8
	VIDEO_OUT_STACKED_FIELDS_1_2_NV = 0x20CB
	VIDEO_OUT_STACKED_FIELDS_2_1_NV = 0x20CC
)
// OML_swap_method
const (
	SWAP_COPY_OML = 0x8062
	SWAP_EXCHANGE_OML = 0x8061
	SWAP_METHOD_OML = 0x8060
	SWAP_UNDEFINED_OML = 0x8063
)
// OML_sync_control
const (
)
// SGIS_blended_overlay
const (
	BLENDED_RGBA_SGIS = 0x8025
)
// SGIS_multisample
const (
	SAMPLES_SGIS = 100001
	SAMPLE_BUFFERS_SGIS = 100000
)
// SGIS_shared_multisample
const (
	MULTISAMPLE_SUB_RECT_HEIGHT_SGIS = 0x8027
	MULTISAMPLE_SUB_RECT_WIDTH_SGIS = 0x8026
)
// SGIX_dmbuffer
const (
	DIGITAL_MEDIA_PBUFFER_SGIX = 0x8024
)
// SGIX_fbconfig
const (
	COLOR_INDEX_BIT_SGIX = 0x00000002
	COLOR_INDEX_TYPE_SGIX = 0x8015
	DRAWABLE_TYPE_SGIX = 0x8010
	FBCONFIG_ID_SGIX = 0x8013
	PIXMAP_BIT_SGIX = 0x00000002
	RENDER_TYPE_SGIX = 0x8011
	RGBA_BIT_SGIX = 0x00000001
	RGBA_TYPE_SGIX = 0x8014
	WINDOW_BIT_SGIX = 0x00000001
	X_RENDERABLE_SGIX = 0x8012
)
// SGIX_hyperpipe
const (
	BAD_HYPERPIPE_CONFIG_SGIX = 91
	BAD_HYPERPIPE_SGIX = 92
	HYPERPIPE_DISPLAY_PIPE_SGIX = 0x00000001
	HYPERPIPE_ID_SGIX = 0x8030
	HYPERPIPE_PIPE_NAME_LENGTH_SGIX = 80
	HYPERPIPE_PIXEL_AVERAGE_SGIX = 0x00000004
	HYPERPIPE_RENDER_PIPE_SGIX = 0x00000002
	HYPERPIPE_STEREO_SGIX = 0x00000003
	PIPE_RECT_LIMITS_SGIX = 0x00000002
	PIPE_RECT_SGIX = 0x00000001
)
// SGIX_pbuffer
const (
	ACCUM_BUFFER_BIT_SGIX = 0x00000080
	AUX_BUFFERS_BIT_SGIX = 0x00000010
	BACK_LEFT_BUFFER_BIT_SGIX = 0x00000004
	BACK_RIGHT_BUFFER_BIT_SGIX = 0x00000008
	BUFFER_CLOBBER_MASK_SGIX = 0x08000000
	DAMAGED_SGIX = 0x8020
	DEPTH_BUFFER_BIT_SGIX = 0x00000020
	EVENT_MASK_SGIX = 0x801F
	FRONT_LEFT_BUFFER_BIT_SGIX = 0x00000001
	FRONT_RIGHT_BUFFER_BIT_SGIX = 0x00000002
	HEIGHT_SGIX = 0x801E
	LARGEST_PBUFFER_SGIX = 0x801C
	MAX_PBUFFER_HEIGHT_SGIX = 0x8017
	MAX_PBUFFER_PIXELS_SGIX = 0x8018
	MAX_PBUFFER_WIDTH_SGIX = 0x8016
	OPTIMAL_PBUFFER_HEIGHT_SGIX = 0x801A
	OPTIMAL_PBUFFER_WIDTH_SGIX = 0x8019
	PBUFFER_BIT_SGIX = 0x00000004
	PBUFFER_SGIX = 0x8023
	PRESERVED_CONTENTS_SGIX = 0x801B
	SAMPLE_BUFFERS_BIT_SGIX = 0x00000100
	SAVED_SGIX = 0x8021
	STENCIL_BUFFER_BIT_SGIX = 0x00000040
	WIDTH_SGIX = 0x801D
	WINDOW_SGIX = 0x8022
)
// SGIX_swap_barrier
const (
)
// SGIX_swap_group
const (
)
// SGIX_video_resize
const (
	SYNC_FRAME_SGIX = 0x00000000
	SYNC_SWAP_SGIX = 0x00000001
)
// SGIX_video_source
const (
)
// SGIX_visual_select_group
const (
	VISUAL_SELECT_GROUP_SGIX = 0x8028
)
// SGI_cushion
const (
)
// SGI_make_current_read
const (
)
// SGI_swap_control
const (
)
// SGI_video_sync
const (
)
// SUN_get_transparent_index
const (
)
// VERSION_1_3
const (
	ACCUM_BUFFER_BIT = 0x00000080
	AUX_BUFFERS_BIT = 0x00000010
	BACK_LEFT_BUFFER_BIT = 0x00000004
	BACK_RIGHT_BUFFER_BIT = 0x00000008
	COLOR_INDEX_BIT = 0x00000002
	COLOR_INDEX_TYPE = 0x8015
	CONFIG_CAVEAT = 0x20
	DAMAGED = 0x8020
	DEPTH_BUFFER_BIT = 0x00000020
	DIRECT_COLOR = 0x8003
	DONT_CARE = 0xFFFFFFFF
	DRAWABLE_TYPE = 0x8010
	EVENT_MASK = 0x801F
	FBCONFIG_ID = 0x8013
	FRONT_LEFT_BUFFER_BIT = 0x00000001
	FRONT_RIGHT_BUFFER_BIT = 0x00000002
	GRAY_SCALE = 0x8006
	HEIGHT = 0x801E
	LARGEST_PBUFFER = 0x801C
	MAX_PBUFFER_HEIGHT = 0x8017
	MAX_PBUFFER_PIXELS = 0x8018
	MAX_PBUFFER_WIDTH = 0x8016
	NONE = 0x8000
	NON_CONFORMANT_CONFIG = 0x800D
	PBUFFER = 0x8023
	PBUFFER_BIT = 0x00000004
	PBUFFER_CLOBBER_MASK = 0x08000000
	PBUFFER_HEIGHT = 0x8040
	PBUFFER_WIDTH = 0x8041
	PIXMAP_BIT = 0x00000002
	PRESERVED_CONTENTS = 0x801B
	PSEUDO_COLOR = 0x8004
	RENDER_TYPE = 0x8011
	RGBA_BIT = 0x00000001
	RGBA_TYPE = 0x8014
	SAVED = 0x8021
	SCREEN = 0x800C
	SLOW_CONFIG = 0x8001
	STATIC_COLOR = 0x8005
	STATIC_GRAY = 0x8007
	STENCIL_BUFFER_BIT = 0x00000040
	TRANSPARENT_ALPHA_VALUE = 0x28
	TRANSPARENT_BLUE_VALUE = 0x27
	TRANSPARENT_GREEN_VALUE = 0x26
	TRANSPARENT_INDEX = 0x8009
	TRANSPARENT_INDEX_VALUE = 0x24
	TRANSPARENT_RED_VALUE = 0x25
	TRANSPARENT_RGB = 0x8008
	TRANSPARENT_TYPE = 0x23
	TRUE_COLOR = 0x8002
	VISUAL_ID = 0x800B
	WIDTH = 0x801D
	WINDOW = 0x8022
	WINDOW_BIT = 0x00000001
	X_RENDERABLE = 0x8012
	X_VISUAL_TYPE = 0x22
)
// VERSION_1_4
const (
	SAMPLES = 100001
	SAMPLE_BUFFERS = 100000
)
// ARB_create_context

// https://www.opengl.org/sdk/docs/man/xhtml/glCreateContextAttribsARB.xml
func CreateContextAttribsARB(dpy Pointer, config Pointer, share_context Pointer, direct int, attrib_list *Int) Pointer {
	return (Pointer)(C.goglxCreateContextAttribsARB((*C.Display)(dpy), (C.GLXFBConfig)(config), (C.GLXContext)(share_context), (C.int)(direct), (*C.int)(attrib_list)))
}
// ARB_get_proc_address

// https://www.opengl.org/sdk/docs/man/xhtml/glGetProcAddressARB.xml
func GetProcAddressARB(procName *Ubyte) Pointer {
	return (Pointer)(C.goglxGetProcAddressARB((*C.GLubyte)(procName)))
}
// EXT_import_context

// https://www.opengl.org/sdk/docs/man/xhtml/glGetCurrentDisplayEXT.xml
func GetCurrentDisplayEXT() Pointer {
	return (Pointer)(C.goglxGetCurrentDisplayEXT())
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryContextInfoEXT.xml
func QueryContextInfoEXT(dpy Pointer, context Pointer, attribute Int, value *Int) Int {
	return (Int)(C.goglxQueryContextInfoEXT((*C.Display)(dpy), (C.GLXContext)(context), (C.int)(attribute), (*C.int)(value)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetContextIDEXT.xml
func GetContextIDEXT(context Pointer) uint32 {
	return (uint32)(C.goglxGetContextIDEXT((C.GLXContext)(context)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glImportContextEXT.xml
func ImportContextEXT(dpy Pointer, contextID uint32) Pointer {
	return (Pointer)(C.goglxImportContextEXT((*C.Display)(dpy), (C.GLXContextID)(contextID)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glFreeContextEXT.xml
func FreeContextEXT(dpy Pointer, context Pointer)  {
	C.goglxFreeContextEXT((*C.Display)(dpy), (C.GLXContext)(context))
}
// EXT_swap_control

// https://www.opengl.org/sdk/docs/man/xhtml/glSwapIntervalEXT.xml
func SwapIntervalEXT(dpy Pointer, drawable Pointer, interval Int)  {
	C.goglxSwapIntervalEXT((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int)(interval))
}
// EXT_texture_from_pixmap

// https://www.opengl.org/sdk/docs/man/xhtml/glBindTexImageEXT.xml
func BindTexImageEXT(dpy Pointer, drawable Pointer, buffer Int, attrib_list *Int)  {
	C.goglxBindTexImageEXT((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int)(buffer), (*C.int)(attrib_list))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glReleaseTexImageEXT.xml
func ReleaseTexImageEXT(dpy Pointer, drawable Pointer, buffer Int)  {
	C.goglxReleaseTexImageEXT((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int)(buffer))
}
// MESA_agp_offset

// https://www.opengl.org/sdk/docs/man/xhtml/glGetAGPOffsetMESA.xml
func GetAGPOffsetMESA(pointer Pointer) uint32 {
	return (uint32)(C.goglxGetAGPOffsetMESA((unsafe.Pointer)(pointer)))
}
// MESA_copy_sub_buffer

// https://www.opengl.org/sdk/docs/man/xhtml/glCopySubBufferMESA.xml
func CopySubBufferMESA(dpy Pointer, drawable Pointer, x Int, y Int, width Int, height Int)  {
	C.goglxCopySubBufferMESA((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int)(x), (C.int)(y), (C.int)(width), (C.int)(height))
}
// MESA_pixmap_colormap

// https://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPixmapMESA.xml
func CreateGLXPixmapMESA(dpy Pointer, visual Pointer, pixmap Pointer, cmap Pointer) Pointer {
	return (Pointer)(C.goglxCreateGLXPixmapMESA((*C.Display)(dpy), (C.XVisualInfo)(visual), (C.Pixmap)(pixmap), (C.Colormap)(cmap)))
}
// MESA_release_buffers

// https://www.opengl.org/sdk/docs/man/xhtml/glReleaseBuffersMESA.xml
func ReleaseBuffersMESA(dpy Pointer, drawable Pointer) int {
	return (int)(C.goglxReleaseBuffersMESA((*C.Display)(dpy), (C.GLXDrawable)(drawable)))
}
// MESA_set_3dfx_mode

// https://www.opengl.org/sdk/docs/man/xhtml/glSet3DfxModeMESA.xml
func Set3DfxModeMESA(mode Int) int {
	return (int)(C.goglxSet3DfxModeMESA((C.int)(mode)))
}
// NV_copy_image

// https://www.opengl.org/sdk/docs/man/xhtml/glCopyImageSubDataNV.xml
func CopyImageSubDataNV(dpy Pointer, srcCtx Pointer, srcName Uint, srcTarget Enum, srcLevel Int, srcX Int, srcY Int, srcZ Int, dstCtx Pointer, dstName Uint, dstTarget Enum, dstLevel Int, dstX Int, dstY Int, dstZ Int, width Sizei, height Sizei, depth Sizei)  {
	C.goglxCopyImageSubDataNV((*C.Display)(dpy), (C.GLXContext)(srcCtx), (C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.GLXContext)(dstCtx), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
// NV_present_video

// https://www.opengl.org/sdk/docs/man/xhtml/glEnumerateVideoDevicesNV.xml
func EnumerateVideoDevicesNV(dpy Pointer, screen Int, nelements *Int) *uint32 {
	return (*uint32)(C.goglxEnumerateVideoDevicesNV((*C.Display)(dpy), (C.int)(screen), (*C.int)(nelements)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glBindVideoDeviceNV.xml
func BindVideoDeviceNV(dpy Pointer, video_slot uint32, video_device uint32, attrib_list *Int) Int {
	return (Int)(C.goglxBindVideoDeviceNV((*C.Display)(dpy), (C.uint32)(video_slot), (C.uint32)(video_device), (*C.int)(attrib_list)))
}
// NV_swap_group

// https://www.opengl.org/sdk/docs/man/xhtml/glJoinSwapGroupNV.xml
func JoinSwapGroupNV(dpy Pointer, drawable Pointer, group Uint) int {
	return (int)(C.goglxJoinSwapGroupNV((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.GLuint)(group)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glBindSwapBarrierNV.xml
func BindSwapBarrierNV(dpy Pointer, group Uint, barrier Uint) int {
	return (int)(C.goglxBindSwapBarrierNV((*C.Display)(dpy), (C.GLuint)(group), (C.GLuint)(barrier)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQuerySwapGroupNV.xml
func QuerySwapGroupNV(dpy Pointer, drawable Pointer, group *Uint, barrier *Uint) int {
	return (int)(C.goglxQuerySwapGroupNV((*C.Display)(dpy), (C.GLXDrawable)(drawable), (*C.GLuint)(group), (*C.GLuint)(barrier)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryMaxSwapGroupsNV.xml
func QueryMaxSwapGroupsNV(dpy Pointer, screen Int, maxGroups *Uint, maxBarriers *Uint) int {
	return (int)(C.goglxQueryMaxSwapGroupsNV((*C.Display)(dpy), (C.int)(screen), (*C.GLuint)(maxGroups), (*C.GLuint)(maxBarriers)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryFrameCountNV.xml
func QueryFrameCountNV(dpy Pointer, screen Int, count *Uint) int {
	return (int)(C.goglxQueryFrameCountNV((*C.Display)(dpy), (C.int)(screen), (*C.GLuint)(count)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glResetFrameCountNV.xml
func ResetFrameCountNV(dpy Pointer, screen Int) int {
	return (int)(C.goglxResetFrameCountNV((*C.Display)(dpy), (C.int)(screen)))
}
// NV_video_capture

// https://www.opengl.org/sdk/docs/man/xhtml/glBindVideoCaptureDeviceNV.xml
func BindVideoCaptureDeviceNV(dpy Pointer, video_capture_slot uint32, device Pointer) Int {
	return (Int)(C.goglxBindVideoCaptureDeviceNV((*C.Display)(dpy), (C.uint32)(video_capture_slot), (C.GLXVideoCaptureDeviceNV)(device)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glEnumerateVideoCaptureDevicesNV.xml
func EnumerateVideoCaptureDevicesNV(dpy Pointer, screen Int, nelements *Int) Pointer {
	return (Pointer)(C.goglxEnumerateVideoCaptureDevicesNV((*C.Display)(dpy), (C.int)(screen), (*C.int)(nelements)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glLockVideoCaptureDeviceNV.xml
func LockVideoCaptureDeviceNV(dpy Pointer, device Pointer)  {
	C.goglxLockVideoCaptureDeviceNV((*C.Display)(dpy), (C.GLXVideoCaptureDeviceNV)(device))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryVideoCaptureDeviceNV.xml
func QueryVideoCaptureDeviceNV(dpy Pointer, device Pointer, attribute Int, value *Int) Int {
	return (Int)(C.goglxQueryVideoCaptureDeviceNV((*C.Display)(dpy), (C.GLXVideoCaptureDeviceNV)(device), (C.int)(attribute), (*C.int)(value)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glReleaseVideoCaptureDeviceNV.xml
func ReleaseVideoCaptureDeviceNV(dpy Pointer, device Pointer)  {
	C.goglxReleaseVideoCaptureDeviceNV((*C.Display)(dpy), (C.GLXVideoCaptureDeviceNV)(device))
}
// NV_video_output

// https://www.opengl.org/sdk/docs/man/xhtml/glGetVideoDeviceNV.xml
func GetVideoDeviceNV(dpy Pointer, screen Int, numVideoDevices Int, pVideoDevice Pointer) Int {
	return (Int)(C.goglxGetVideoDeviceNV((*C.Display)(dpy), (C.int)(screen), (C.int)(numVideoDevices), (C.GLXVideoDeviceNV)(pVideoDevice)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glReleaseVideoDeviceNV.xml
func ReleaseVideoDeviceNV(dpy Pointer, screen Int, VideoDevice Pointer) Int {
	return (Int)(C.goglxReleaseVideoDeviceNV((*C.Display)(dpy), (C.int)(screen), (C.GLXVideoDeviceNV)(VideoDevice)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glBindVideoImageNV.xml
func BindVideoImageNV(dpy Pointer, VideoDevice Pointer, pbuf Pointer, iVideoBuffer Int) Int {
	return (Int)(C.goglxBindVideoImageNV((*C.Display)(dpy), (C.GLXVideoDeviceNV)(VideoDevice), (C.GLXPbuffer)(pbuf), (C.int)(iVideoBuffer)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glReleaseVideoImageNV.xml
func ReleaseVideoImageNV(dpy Pointer, pbuf Pointer) Int {
	return (Int)(C.goglxReleaseVideoImageNV((*C.Display)(dpy), (C.GLXPbuffer)(pbuf)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glSendPbufferToVideoNV.xml
func SendPbufferToVideoNV(dpy Pointer, pbuf Pointer, iBufferType Int, pulCounterPbuffer *uint32, bBlock Boolean) Int {
	return (Int)(C.goglxSendPbufferToVideoNV((*C.Display)(dpy), (C.GLXPbuffer)(pbuf), (C.int)(iBufferType), (*C.unsigned_long)(pulCounterPbuffer), (C.GLboolean)(bBlock)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetVideoInfoNV.xml
func GetVideoInfoNV(dpy Pointer, screen Int, VideoDevice Pointer, pulCounterOutputPbuffer *uint32, pulCounterOutputVideo *uint32) Int {
	return (Int)(C.goglxGetVideoInfoNV((*C.Display)(dpy), (C.int)(screen), (C.GLXVideoDeviceNV)(VideoDevice), (*C.unsigned_long)(pulCounterOutputPbuffer), (*C.unsigned_long)(pulCounterOutputVideo)))
}
// OML_sync_control

// https://www.opengl.org/sdk/docs/man/xhtml/glGetSyncValuesOML.xml
func GetSyncValuesOML(dpy Pointer, drawable Pointer, ust *int64, msc *int64, sbc *int64) int {
	return (int)(C.goglxGetSyncValuesOML((*C.Display)(dpy), (C.GLXDrawable)(drawable), (*C.int64_t)(ust), (*C.int64_t)(msc), (*C.int64_t)(sbc)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetMscRateOML.xml
func GetMscRateOML(dpy Pointer, drawable Pointer, numerator *int32, denominator *int32) int {
	return (int)(C.goglxGetMscRateOML((*C.Display)(dpy), (C.GLXDrawable)(drawable), (*C.int32_t)(numerator), (*C.int32_t)(denominator)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glSwapBuffersMscOML.xml
func SwapBuffersMscOML(dpy Pointer, drawable Pointer, target_msc int64, divisor int64, remainder int64) int64 {
	return (int64)(C.goglxSwapBuffersMscOML((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int64_t)(target_msc), (C.int64_t)(divisor), (C.int64_t)(remainder)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glWaitForMscOML.xml
func WaitForMscOML(dpy Pointer, drawable Pointer, target_msc int64, divisor int64, remainder int64, ust *int64, msc *int64, sbc *int64) int {
	return (int)(C.goglxWaitForMscOML((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int64_t)(target_msc), (C.int64_t)(divisor), (C.int64_t)(remainder), (*C.int64_t)(ust), (*C.int64_t)(msc), (*C.int64_t)(sbc)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glWaitForSbcOML.xml
func WaitForSbcOML(dpy Pointer, drawable Pointer, target_sbc int64, ust *int64, msc *int64, sbc *int64) int {
	return (int)(C.goglxWaitForSbcOML((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int64_t)(target_sbc), (*C.int64_t)(ust), (*C.int64_t)(msc), (*C.int64_t)(sbc)))
}
// SGIX_dmbuffer

// https://www.opengl.org/sdk/docs/man/xhtml/glAssociateDMPbufferSGIX.xml
func AssociateDMPbufferSGIX(dpy Pointer, pbuffer Pointer, params Pointer, dmbuffer Pointer) int {
	return (int)(C.goglxAssociateDMPbufferSGIX((*C.Display)(dpy), (C.GLXPbufferSGIX)(pbuffer), (C.DMparams)(params), (C.DMbuffer)(dmbuffer)))
}
// SGIX_fbconfig

// https://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigAttribSGIX.xml
func GetFBConfigAttribSGIX(dpy Pointer, config Pointer, attribute Int, value *Int) Int {
	return (Int)(C.goglxGetFBConfigAttribSGIX((*C.Display)(dpy), (C.GLXFBConfigSGIX)(config), (C.int)(attribute), (*C.int)(value)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glChooseFBConfigSGIX.xml
func ChooseFBConfigSGIX(dpy Pointer, screen Int, attrib_list *Int, nelements *Int) Pointer {
	return (Pointer)(C.goglxChooseFBConfigSGIX((*C.Display)(dpy), (C.int)(screen), (*C.int)(attrib_list), (*C.int)(nelements)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPixmapWithConfigSGIX.xml
func CreateGLXPixmapWithConfigSGIX(dpy Pointer, config Pointer, pixmap Pointer) Pointer {
	return (Pointer)(C.goglxCreateGLXPixmapWithConfigSGIX((*C.Display)(dpy), (C.GLXFBConfigSGIX)(config), (C.Pixmap)(pixmap)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateContextWithConfigSGIX.xml
func CreateContextWithConfigSGIX(dpy Pointer, config Pointer, render_type Int, share_list Pointer, direct int) Pointer {
	return (Pointer)(C.goglxCreateContextWithConfigSGIX((*C.Display)(dpy), (C.GLXFBConfigSGIX)(config), (C.int)(render_type), (C.GLXContext)(share_list), (C.int)(direct)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetVisualFromFBConfigSGIX.xml
func GetVisualFromFBConfigSGIX(dpy Pointer, config Pointer) Pointer {
	return (Pointer)(C.goglxGetVisualFromFBConfigSGIX((*C.Display)(dpy), (C.GLXFBConfigSGIX)(config)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigFromVisualSGIX.xml
func GetFBConfigFromVisualSGIX(dpy Pointer, vis Pointer) Pointer {
	return (Pointer)(C.goglxGetFBConfigFromVisualSGIX((*C.Display)(dpy), (C.XVisualInfo)(vis)))
}
// SGIX_hyperpipe

// https://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeNetworkSGIX.xml
func QueryHyperpipeNetworkSGIX(dpy Pointer, npipes *Int) Pointer {
	return (Pointer)(C.goglxQueryHyperpipeNetworkSGIX((*C.Display)(dpy), (*C.int)(npipes)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glHyperpipeConfigSGIX.xml
func HyperpipeConfigSGIX(dpy Pointer, networkId Int, npipes Int, cfg Pointer, hpId *Int) Int {
	return (Int)(C.goglxHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(networkId), (C.int)(npipes), (*C.GLXHyperpipeConfigSGIX)(cfg), (*C.int)(hpId)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeConfigSGIX.xml
func QueryHyperpipeConfigSGIX(dpy Pointer, hpId Int, npipes *Int) Pointer {
	return (Pointer)(C.goglxQueryHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(hpId), (*C.int)(npipes)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyHyperpipeConfigSGIX.xml
func DestroyHyperpipeConfigSGIX(dpy Pointer, hpId Int) Int {
	return (Int)(C.goglxDestroyHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(hpId)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glBindHyperpipeSGIX.xml
func BindHyperpipeSGIX(dpy Pointer, hpId Int) Int {
	return (Int)(C.goglxBindHyperpipeSGIX((*C.Display)(dpy), (C.int)(hpId)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeBestAttribSGIX.xml
func QueryHyperpipeBestAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, attribList Pointer, returnAttribList Pointer) Int {
	return (Int)(C.goglxQueryHyperpipeBestAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(attribList), (unsafe.Pointer)(returnAttribList)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glHyperpipeAttribSGIX.xml
func HyperpipeAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, attribList Pointer) Int {
	return (Int)(C.goglxHyperpipeAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(attribList)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeAttribSGIX.xml
func QueryHyperpipeAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, returnAttribList Pointer) Int {
	return (Int)(C.goglxQueryHyperpipeAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(returnAttribList)))
}
// SGIX_pbuffer

// https://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPbufferSGIX.xml
func CreateGLXPbufferSGIX(dpy Pointer, config Pointer, width uint32, height uint32, attrib_list *Int) Pointer {
	return (Pointer)(C.goglxCreateGLXPbufferSGIX((*C.Display)(dpy), (C.GLXFBConfigSGIX)(config), (C.uint32)(width), (C.uint32)(height), (*C.int)(attrib_list)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyGLXPbufferSGIX.xml
func DestroyGLXPbufferSGIX(dpy Pointer, pbuf Pointer)  {
	C.goglxDestroyGLXPbufferSGIX((*C.Display)(dpy), (C.GLXPbufferSGIX)(pbuf))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryGLXPbufferSGIX.xml
func QueryGLXPbufferSGIX(dpy Pointer, pbuf Pointer, attribute Int, value *uint32) Int {
	return (Int)(C.goglxQueryGLXPbufferSGIX((*C.Display)(dpy), (C.GLXPbufferSGIX)(pbuf), (C.int)(attribute), (*C.uint32)(value)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glSelectEventSGIX.xml
func SelectEventSGIX(dpy Pointer, drawable Pointer, mask uint32)  {
	C.goglxSelectEventSGIX((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.unsigned_long)(mask))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetSelectedEventSGIX.xml
func GetSelectedEventSGIX(dpy Pointer, drawable Pointer, mask *uint32)  {
	C.goglxGetSelectedEventSGIX((*C.Display)(dpy), (C.GLXDrawable)(drawable), (*C.unsigned_long)(mask))
}
// SGIX_swap_barrier

// https://www.opengl.org/sdk/docs/man/xhtml/glBindSwapBarrierSGIX.xml
func BindSwapBarrierSGIX(dpy Pointer, drawable Pointer, barrier Int)  {
	C.goglxBindSwapBarrierSGIX((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int)(barrier))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryMaxSwapBarriersSGIX.xml
func QueryMaxSwapBarriersSGIX(dpy Pointer, screen Int, max *Int) int {
	return (int)(C.goglxQueryMaxSwapBarriersSGIX((*C.Display)(dpy), (C.int)(screen), (*C.int)(max)))
}
// SGIX_swap_group

// https://www.opengl.org/sdk/docs/man/xhtml/glJoinSwapGroupSGIX.xml
func JoinSwapGroupSGIX(dpy Pointer, drawable Pointer, member Pointer)  {
	C.goglxJoinSwapGroupSGIX((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.GLXDrawable)(member))
}
// SGIX_video_resize

// https://www.opengl.org/sdk/docs/man/xhtml/glBindChannelToWindowSGIX.xml
func BindChannelToWindowSGIX(display Pointer, screen Int, channel Int, window Pointer) Int {
	return (Int)(C.goglxBindChannelToWindowSGIX((*C.Display)(display), (C.int)(screen), (C.int)(channel), (C.Window)(window)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glChannelRectSGIX.xml
func ChannelRectSGIX(display Pointer, screen Int, channel Int, x Int, y Int, w Int, h Int) Int {
	return (Int)(C.goglxChannelRectSGIX((*C.Display)(display), (C.int)(screen), (C.int)(channel), (C.int)(x), (C.int)(y), (C.int)(w), (C.int)(h)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryChannelRectSGIX.xml
func QueryChannelRectSGIX(display Pointer, screen Int, channel Int, dx *Int, dy *Int, dw *Int, dh *Int) Int {
	return (Int)(C.goglxQueryChannelRectSGIX((*C.Display)(display), (C.int)(screen), (C.int)(channel), (*C.int)(dx), (*C.int)(dy), (*C.int)(dw), (*C.int)(dh)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryChannelDeltasSGIX.xml
func QueryChannelDeltasSGIX(display Pointer, screen Int, channel Int, x *Int, y *Int, w *Int, h *Int) Int {
	return (Int)(C.goglxQueryChannelDeltasSGIX((*C.Display)(display), (C.int)(screen), (C.int)(channel), (*C.int)(x), (*C.int)(y), (*C.int)(w), (*C.int)(h)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glChannelRectSyncSGIX.xml
func ChannelRectSyncSGIX(display Pointer, screen Int, channel Int, synctype Enum) Int {
	return (Int)(C.goglxChannelRectSyncSGIX((*C.Display)(display), (C.int)(screen), (C.int)(channel), (C.GLenum)(synctype)))
}
// SGIX_video_source

// https://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXVideoSourceSGIX.xml
func CreateGLXVideoSourceSGIX(display Pointer, screen Int, server Pointer, path Pointer, nodeClass Int, drainNode Pointer) Pointer {
	return (Pointer)(C.goglxCreateGLXVideoSourceSGIX((*C.Display)(display), (C.int)(screen), (C.VLServer)(server), (C.VLPath)(path), (C.int)(nodeClass), (C.VLNode)(drainNode)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyGLXVideoSourceSGIX.xml
func DestroyGLXVideoSourceSGIX(dpy Pointer, glxvideosource Pointer)  {
	C.goglxDestroyGLXVideoSourceSGIX((*C.Display)(dpy), (*C.GGLXVideoSourceSGIX)(glxvideosource))
}
// SGI_cushion

// https://www.opengl.org/sdk/docs/man/xhtml/glCushionSGI.xml
func CushionSGI(dpy Pointer, window Pointer, cushion float32)  {
	C.goglxCushionSGI((*C.Display)(dpy), (C.Window)(window), (C.float32)(cushion))
}
// SGI_make_current_read

// https://www.opengl.org/sdk/docs/man/xhtml/glMakeCurrentReadSGI.xml
func MakeCurrentReadSGI(dpy Pointer, draw Pointer, read Pointer, ctx Pointer) int {
	return (int)(C.goglxMakeCurrentReadSGI((*C.Display)(dpy), (C.GLXDrawable)(draw), (C.GLXDrawable)(read), (C.GLXContext)(ctx)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetCurrentReadDrawableSGI.xml
func GetCurrentReadDrawableSGI() Pointer {
	return (Pointer)(C.goglxGetCurrentReadDrawableSGI())
}
// SGI_swap_control

// https://www.opengl.org/sdk/docs/man/xhtml/glSwapIntervalSGI.xml
func SwapIntervalSGI(interval Int) Int {
	return (Int)(C.goglxSwapIntervalSGI((C.int)(interval)))
}
// SGI_video_sync

// https://www.opengl.org/sdk/docs/man/xhtml/glGetVideoSyncSGI.xml
func GetVideoSyncSGI(count *uint32) Int {
	return (Int)(C.goglxGetVideoSyncSGI((*C.uint32)(count)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glWaitVideoSyncSGI.xml
func WaitVideoSyncSGI(divisor Int, remainder Int, count *uint32) Int {
	return (Int)(C.goglxWaitVideoSyncSGI((C.int)(divisor), (C.int)(remainder), (*C.uint32)(count)))
}
// SUN_get_transparent_index

// https://www.opengl.org/sdk/docs/man/xhtml/glGetTransparentIndexSUN.xml
func GetTransparentIndexSUN(dpy Pointer, overlay Pointer, underlay Pointer, pTransparentIndex *int32) int {
	return (int)(C.goglxGetTransparentIndexSUN((*C.Display)(dpy), (C.Window)(overlay), (C.Window)(underlay), (*C.long)(pTransparentIndex)))
}
// VERSION_1_3

// https://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigs.xml
func GetFBConfigs(dpy Pointer, screen Int, nelements *Int) Pointer {
	return (Pointer)(C.goglxGetFBConfigs((*C.Display)(dpy), (C.int)(screen), (*C.int)(nelements)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glChooseFBConfig.xml
func ChooseFBConfig(dpy Pointer, screen Int, attrib_list *Int, nelements *Int) Pointer {
	return (Pointer)(C.goglxChooseFBConfig((*C.Display)(dpy), (C.int)(screen), (*C.int)(attrib_list), (*C.int)(nelements)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigAttrib.xml
func GetFBConfigAttrib(dpy Pointer, config Pointer, attribute Int, value *Int) Int {
	return (Int)(C.goglxGetFBConfigAttrib((*C.Display)(dpy), (C.GLXFBConfig)(config), (C.int)(attribute), (*C.int)(value)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetVisualFromFBConfig.xml
func GetVisualFromFBConfig(dpy Pointer, config Pointer) Pointer {
	return (Pointer)(C.goglxGetVisualFromFBConfig((*C.Display)(dpy), (C.GLXFBConfig)(config)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateWindow.xml
func CreateWindow(dpy Pointer, config Pointer, win Pointer, attrib_list *Int) Pointer {
	return (Pointer)(C.goglxCreateWindow((*C.Display)(dpy), (C.GLXFBConfig)(config), (C.Window)(win), (*C.int)(attrib_list)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyWindow.xml
func DestroyWindow(dpy Pointer, win Pointer)  {
	C.goglxDestroyWindow((*C.Display)(dpy), (C.GLXWindow)(win))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreatePixmap.xml
func CreatePixmap(dpy Pointer, config Pointer, pixmap Pointer, attrib_list *Int) Pointer {
	return (Pointer)(C.goglxCreatePixmap((*C.Display)(dpy), (C.GLXFBConfig)(config), (C.Pixmap)(pixmap), (*C.int)(attrib_list)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyPixmap.xml
func DestroyPixmap(dpy Pointer, pixmap Pointer)  {
	C.goglxDestroyPixmap((*C.Display)(dpy), (C.GLXPixmap)(pixmap))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreatePbuffer.xml
func CreatePbuffer(dpy Pointer, config Pointer, attrib_list *Int) Pointer {
	return (Pointer)(C.goglxCreatePbuffer((*C.Display)(dpy), (C.GLXFBConfig)(config), (*C.int)(attrib_list)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyPbuffer.xml
func DestroyPbuffer(dpy Pointer, pbuf Pointer)  {
	C.goglxDestroyPbuffer((*C.Display)(dpy), (C.GLXPbuffer)(pbuf))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryDrawable.xml
func QueryDrawable(dpy Pointer, draw Pointer, attribute Int, value *uint32)  {
	C.goglxQueryDrawable((*C.Display)(dpy), (C.GLXDrawable)(draw), (C.int)(attribute), (*C.uint32)(value))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateNewContext.xml
func CreateNewContext(dpy Pointer, config Pointer, render_type Int, share_list Pointer, direct int) Pointer {
	return (Pointer)(C.goglxCreateNewContext((*C.Display)(dpy), (C.GLXFBConfig)(config), (C.int)(render_type), (C.GLXContext)(share_list), (C.int)(direct)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glMakeContextCurrent.xml
func MakeContextCurrent(dpy Pointer, draw Pointer, read Pointer, ctx Pointer) int {
	return (int)(C.goglxMakeContextCurrent((*C.Display)(dpy), (C.GLXDrawable)(draw), (C.GLXDrawable)(read), (C.GLXContext)(ctx)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetCurrentReadDrawable.xml
func GetCurrentReadDrawable() Pointer {
	return (Pointer)(C.goglxGetCurrentReadDrawable())
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetCurrentDisplay.xml
func GetCurrentDisplay() Pointer {
	return (Pointer)(C.goglxGetCurrentDisplay())
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryContext.xml
func QueryContext(dpy Pointer, ctx Pointer, attribute Int, value *Int) Int {
	return (Int)(C.goglxQueryContext((*C.Display)(dpy), (C.GLXContext)(ctx), (C.int)(attribute), (*C.int)(value)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glSelectEvent.xml
func SelectEvent(dpy Pointer, draw Pointer, event_mask uint32)  {
	C.goglxSelectEvent((*C.Display)(dpy), (C.GLXDrawable)(draw), (C.unsigned_long)(event_mask))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetSelectedEvent.xml
func GetSelectedEvent(dpy Pointer, draw Pointer, event_mask *uint32)  {
	C.goglxGetSelectedEvent((*C.Display)(dpy), (C.GLXDrawable)(draw), (*C.unsigned_long)(event_mask))
}
// VERSION_1_4

// https://www.opengl.org/sdk/docs/man/xhtml/glGetProcAddress.xml
func GetProcAddress(procName *Ubyte) Pointer {
	return (Pointer)(C.goglxGetProcAddress((*C.GLubyte)(procName)))
}
// glx

// https://www.opengl.org/sdk/docs/man/xhtml/glRender.xml
func Render()  {
	C.goglxRender()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glRenderLarge.xml
func RenderLarge()  {
	C.goglxRenderLarge()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateContext.xml
func CreateContext(gc_id Int, screen Int, visual Int, share_list Int)  {
	C.goglxCreateContext((C.GLint)(gc_id), (C.GLint)(screen), (C.GLint)(visual), (C.GLint)(share_list))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyContext.xml
func DestroyContext(context Int)  {
	C.goglxDestroyContext((C.GLint)(context))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glMakeCurrent.xml
func MakeCurrent(drawable Int, context Int)  {
	C.goglxMakeCurrent((C.GLint)(drawable), (C.GLint)(context))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glIsDirect.xml
func IsDirect(dpy Int, context Int)  {
	C.goglxIsDirect((C.GLint)(dpy), (C.GLint)(context))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryVersion.xml
func QueryVersion(major *Int, minor *Int)  {
	C.goglxQueryVersion((*C.GLint)(major), (*C.GLint)(minor))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glWaitGL.xml
func WaitGL(context Int)  {
	C.goglxWaitGL((C.GLint)(context))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glWaitX.xml
func WaitX()  {
	C.goglxWaitX()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCopyContext.xml
func CopyContext(source Int, dest Int, mask Int)  {
	C.goglxCopyContext((C.GLint)(source), (C.GLint)(dest), (C.GLint)(mask))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glSwapBuffers.xml
func SwapBuffers(drawable Int)  {
	C.goglxSwapBuffers((C.GLint)(drawable))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glUseXFont.xml
func UseXFont(font Int, first Int, count Int, list_base Int)  {
	C.goglxUseXFont((C.GLint)(font), (C.GLint)(first), (C.GLint)(count), (C.GLint)(list_base))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPixmap.xml
func CreateGLXPixmap(visual Int, pixmap Int, glxpixmap Int)  {
	C.goglxCreateGLXPixmap((C.GLint)(visual), (C.GLint)(pixmap), (C.GLint)(glxpixmap))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetVisualConfigs.xml
func GetVisualConfigs()  {
	C.goglxGetVisualConfigs()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyGLXPixmap.xml
func DestroyGLXPixmap(pixmap Int)  {
	C.goglxDestroyGLXPixmap((C.GLint)(pixmap))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glVendorPrivate.xml
func VendorPrivate()  {
	C.goglxVendorPrivate()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glVendorPrivateWithReply.xml
func VendorPrivateWithReply()  {
	C.goglxVendorPrivateWithReply()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryExtensionsString.xml
func QueryExtensionsString(screen Int)  {
	C.goglxQueryExtensionsString((C.GLint)(screen))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryServerString.xml
func QueryServerString(screen Int, name Int)  {
	C.goglxQueryServerString((C.GLint)(screen), (C.GLint)(name))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glClientInfo.xml
func ClientInfo()  {
	C.goglxClientInfo()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigs.xml
func GetFBConfigs()  {
	C.goglxGetFBConfigs()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreatePixmap.xml
func CreatePixmap(config Int, pixmap Int, glxpixmap Int)  {
	C.goglxCreatePixmap((C.GLint)(config), (C.GLint)(pixmap), (C.GLint)(glxpixmap))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyPixmap.xml
func DestroyPixmap(glxpixmap Int)  {
	C.goglxDestroyPixmap((C.GLint)(glxpixmap))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateNewContext.xml
func CreateNewContext(config Int, render_type Int, share_list Int, direct Int)  {
	C.goglxCreateNewContext((C.GLint)(config), (C.GLint)(render_type), (C.GLint)(share_list), (C.GLint)(direct))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryContext.xml
func QueryContext()  {
	C.goglxQueryContext()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glMakeContextCurrent.xml
func MakeContextCurrent(drawable Int, readdrawable Int, context Int)  {
	C.goglxMakeContextCurrent((C.GLint)(drawable), (C.GLint)(readdrawable), (C.GLint)(context))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreatePbuffer.xml
func CreatePbuffer(config Int, pbuffer Int)  {
	C.goglxCreatePbuffer((C.GLint)(config), (C.GLint)(pbuffer))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyPbuffer.xml
func DestroyPbuffer(pbuffer Int)  {
	C.goglxDestroyPbuffer((C.GLint)(pbuffer))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetDrawableAttributes.xml
func GetDrawableAttributes(drawable Int)  {
	C.goglxGetDrawableAttributes((C.GLint)(drawable))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glChangeDrawableAttributes.xml
func ChangeDrawableAttributes(drawable Int)  {
	C.goglxChangeDrawableAttributes((C.GLint)(drawable))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateWindow.xml
func CreateWindow(config Int, window Int, glxwindow Int)  {
	C.goglxCreateWindow((C.GLint)(config), (C.GLint)(window), (C.GLint)(glxwindow))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyWindow.xml
func DestroyWindow(glxwindow Int)  {
	C.goglxDestroyWindow((C.GLint)(glxwindow))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glSwapIntervalSGI.xml
func SwapIntervalSGI()  {
	C.goglxSwapIntervalSGI()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glMakeCurrentReadSGI.xml
func MakeCurrentReadSGI(drawable Int, readdrawable Int, context Int)  {
	C.goglxMakeCurrentReadSGI((C.GLint)(drawable), (C.GLint)(readdrawable), (C.GLint)(context))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXVideoSourceSGIX.xml
func CreateGLXVideoSourceSGIX(dpy Int, screen Int, server Int, path Int, class Int, node Int)  {
	C.goglxCreateGLXVideoSourceSGIX((C.GLint)(dpy), (C.GLint)(screen), (C.GLint)(server), (C.GLint)(path), (C.GLint)(class), (C.GLint)(node))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyGLXVideoSourceSGIX.xml
func DestroyGLXVideoSourceSGIX(dpy Int, glxvideosource Int)  {
	C.goglxDestroyGLXVideoSourceSGIX((C.GLint)(dpy), (C.GLint)(glxvideosource))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryContextInfoEXT.xml
func QueryContextInfoEXT()  {
	C.goglxQueryContextInfoEXT()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigsSGIX.xml
func GetFBConfigsSGIX()  {
	C.goglxGetFBConfigsSGIX()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateContextWithConfigSGIX.xml
func CreateContextWithConfigSGIX(gc_id Int, screen Int, config Int, share_list Int)  {
	C.goglxCreateContextWithConfigSGIX((C.GLint)(gc_id), (C.GLint)(screen), (C.GLint)(config), (C.GLint)(share_list))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPixmapWithConfigSGIX.xml
func CreateGLXPixmapWithConfigSGIX(config Int, pixmap Int, glxpixmap Int)  {
	C.goglxCreateGLXPixmapWithConfigSGIX((C.GLint)(config), (C.GLint)(pixmap), (C.GLint)(glxpixmap))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPbufferSGIX.xml
func CreateGLXPbufferSGIX(config Int, pbuffer Int)  {
	C.goglxCreateGLXPbufferSGIX((C.GLint)(config), (C.GLint)(pbuffer))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyGLXPbufferSGIX.xml
func DestroyGLXPbufferSGIX(pbuffer Int)  {
	C.goglxDestroyGLXPbufferSGIX((C.GLint)(pbuffer))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glChangeDrawableAttributesSGIX.xml
func ChangeDrawableAttributesSGIX(drawable Int)  {
	C.goglxChangeDrawableAttributesSGIX((C.GLint)(drawable))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glGetDrawableAttributesSGIX.xml
func GetDrawableAttributesSGIX(drawable Int)  {
	C.goglxGetDrawableAttributesSGIX((C.GLint)(drawable))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glJoinSwapGroupSGIX.xml
func JoinSwapGroupSGIX(window Int, group Int)  {
	C.goglxJoinSwapGroupSGIX((C.GLint)(window), (C.GLint)(group))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glBindSwapBarrierSGIX.xml
func BindSwapBarrierSGIX(window Int, barrier Int)  {
	C.goglxBindSwapBarrierSGIX((C.GLint)(window), (C.GLint)(barrier))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryMaxSwapBarriersSGIX.xml
func QueryMaxSwapBarriersSGIX()  {
	C.goglxQueryMaxSwapBarriersSGIX()
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeNetworkSGIX.xml
func QueryHyperpipeNetworkSGIX(dpy Pointer, npipes *Int) Pointer {
	return (Pointer)(C.goglxQueryHyperpipeNetworkSGIX((*C.Display)(dpy), (*C.int)(npipes)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glHyperpipeConfigSGIX.xml
func HyperpipeConfigSGIX(dpy Pointer, networkId Int, npipes Int, cfg Pointer, hpId *Int) Int {
	return (Int)(C.goglxHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(networkId), (C.int)(npipes), (*C.GLXHyperpipeConfigSGIX)(cfg), (*C.int)(hpId)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeConfigSGIX.xml
func QueryHyperpipeConfigSGIX(dpy Pointer, hpId Int, npipes *Int) Pointer {
	return (Pointer)(C.goglxQueryHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(hpId), (*C.int)(npipes)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glDestroyHyperpipeConfigSGIX.xml
func DestroyHyperpipeConfigSGIX(dpy Pointer, hpId Int) Int {
	return (Int)(C.goglxDestroyHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(hpId)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glBindHyperpipeSGIX.xml
func BindHyperpipeSGIX(dpy Pointer, hpId Int) Int {
	return (Int)(C.goglxBindHyperpipeSGIX((*C.Display)(dpy), (C.int)(hpId)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeBestAttribSGIX.xml
func QueryHyperpipeBestAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, attribList Pointer, returnAttribList Pointer) Int {
	return (Int)(C.goglxQueryHyperpipeBestAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(attribList), (unsafe.Pointer)(returnAttribList)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glHyperpipeAttribSGIX.xml
func HyperpipeAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, attribList Pointer) Int {
	return (Int)(C.goglxHyperpipeAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(attribList)))
}
// https://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeAttribSGIX.xml
func QueryHyperpipeAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, returnAttribList Pointer) Int {
	return (Int)(C.goglxQueryHyperpipeAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(returnAttribList)))
}
func InitArbCreateContext() error {
	var ret C.int
	if ret = C.init_ARB_create_context(); ret != 0 {
		return errors.New("unable to initialize ARB_create_context")
	}
	return nil
}
func InitArbGetProcAddress() error {
	var ret C.int
	if ret = C.init_ARB_get_proc_address(); ret != 0 {
		return errors.New("unable to initialize ARB_get_proc_address")
	}
	return nil
}
func InitExtImportContext() error {
	var ret C.int
	if ret = C.init_EXT_import_context(); ret != 0 {
		return errors.New("unable to initialize EXT_import_context")
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
func InitExtTextureFromPixmap() error {
	var ret C.int
	if ret = C.init_EXT_texture_from_pixmap(); ret != 0 {
		return errors.New("unable to initialize EXT_texture_from_pixmap")
	}
	return nil
}
func InitMesaAgpOffset() error {
	var ret C.int
	if ret = C.init_MESA_agp_offset(); ret != 0 {
		return errors.New("unable to initialize MESA_agp_offset")
	}
	return nil
}
func InitMesaCopySubBuffer() error {
	var ret C.int
	if ret = C.init_MESA_copy_sub_buffer(); ret != 0 {
		return errors.New("unable to initialize MESA_copy_sub_buffer")
	}
	return nil
}
func InitMesaPixmapColormap() error {
	var ret C.int
	if ret = C.init_MESA_pixmap_colormap(); ret != 0 {
		return errors.New("unable to initialize MESA_pixmap_colormap")
	}
	return nil
}
func InitMesaReleaseBuffers() error {
	var ret C.int
	if ret = C.init_MESA_release_buffers(); ret != 0 {
		return errors.New("unable to initialize MESA_release_buffers")
	}
	return nil
}
func InitMesaSet3dfxMode() error {
	var ret C.int
	if ret = C.init_MESA_set_3dfx_mode(); ret != 0 {
		return errors.New("unable to initialize MESA_set_3dfx_mode")
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
func InitSgixDmbuffer() error {
	var ret C.int
	if ret = C.init_SGIX_dmbuffer(); ret != 0 {
		return errors.New("unable to initialize SGIX_dmbuffer")
	}
	return nil
}
func InitSgixFbconfig() error {
	var ret C.int
	if ret = C.init_SGIX_fbconfig(); ret != 0 {
		return errors.New("unable to initialize SGIX_fbconfig")
	}
	return nil
}
func InitSgixHyperpipe() error {
	var ret C.int
	if ret = C.init_SGIX_hyperpipe(); ret != 0 {
		return errors.New("unable to initialize SGIX_hyperpipe")
	}
	return nil
}
func InitSgixPbuffer() error {
	var ret C.int
	if ret = C.init_SGIX_pbuffer(); ret != 0 {
		return errors.New("unable to initialize SGIX_pbuffer")
	}
	return nil
}
func InitSgixSwapBarrier() error {
	var ret C.int
	if ret = C.init_SGIX_swap_barrier(); ret != 0 {
		return errors.New("unable to initialize SGIX_swap_barrier")
	}
	return nil
}
func InitSgixSwapGroup() error {
	var ret C.int
	if ret = C.init_SGIX_swap_group(); ret != 0 {
		return errors.New("unable to initialize SGIX_swap_group")
	}
	return nil
}
func InitSgixVideoResize() error {
	var ret C.int
	if ret = C.init_SGIX_video_resize(); ret != 0 {
		return errors.New("unable to initialize SGIX_video_resize")
	}
	return nil
}
func InitSgixVideoSource() error {
	var ret C.int
	if ret = C.init_SGIX_video_source(); ret != 0 {
		return errors.New("unable to initialize SGIX_video_source")
	}
	return nil
}
func InitSgiCushion() error {
	var ret C.int
	if ret = C.init_SGI_cushion(); ret != 0 {
		return errors.New("unable to initialize SGI_cushion")
	}
	return nil
}
func InitSgiMakeCurrentRead() error {
	var ret C.int
	if ret = C.init_SGI_make_current_read(); ret != 0 {
		return errors.New("unable to initialize SGI_make_current_read")
	}
	return nil
}
func InitSgiSwapControl() error {
	var ret C.int
	if ret = C.init_SGI_swap_control(); ret != 0 {
		return errors.New("unable to initialize SGI_swap_control")
	}
	return nil
}
func InitSgiVideoSync() error {
	var ret C.int
	if ret = C.init_SGI_video_sync(); ret != 0 {
		return errors.New("unable to initialize SGI_video_sync")
	}
	return nil
}
func InitSunGetTransparentIndex() error {
	var ret C.int
	if ret = C.init_SUN_get_transparent_index(); ret != 0 {
		return errors.New("unable to initialize SUN_get_transparent_index")
	}
	return nil
}
func InitVersion13() error {
	var ret C.int
	if ret = C.init_VERSION_1_3(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_3")
	}
	return nil
}
func InitVersion14() error {
	var ret C.int
	if ret = C.init_VERSION_1_4(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_4")
	}
	return nil
}
func InitGlx() error {
	var ret C.int
	if ret = C.init_glx(); ret != 0 {
		return errors.New("unable to initialize glx")
	}
	return nil
}
// EOF