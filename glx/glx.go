// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// ARB_create_context: http://www.opengl.org/registry/specs/ARB/create_context.txt
// 
// ARB_create_context_profile: http://www.opengl.org/registry/specs/ARB/create_context_profile.txt
// 
// ARB_create_context_robustness: http://www.opengl.org/registry/specs/ARB/create_context_robustness.txt
// 
// ARB_fbconfig_float: http://www.opengl.org/registry/specs/ARB/fbconfig_float.txt
// 
// ARB_framebuffer_sRGB: http://www.opengl.org/registry/specs/ARB/framebuffer_sRGB.txt
// 
// ARB_get_proc_address: http://www.opengl.org/registry/specs/ARB/get_proc_address.txt
// 
// ARB_multisample: http://www.opengl.org/registry/specs/ARB/multisample.txt
// 
// EXT_fbconfig_packed_float: http://www.opengl.org/registry/specs/EXT/fbconfig_packed_float.txt
// 
// EXT_framebuffer_sRGB: http://www.opengl.org/registry/specs/EXT/framebuffer_sRGB.txt
// 
// EXT_import_context: http://www.opengl.org/registry/specs/EXT/import_context.txt
// 
// EXT_swap_control: http://www.opengl.org/registry/specs/EXT/swap_control.txt
// 
// EXT_swap_control_tear: http://www.opengl.org/registry/specs/EXT/swap_control_tear.txt
// 
// EXT_texture_from_pixmap: http://www.opengl.org/registry/specs/EXT/texture_from_pixmap.txt
// 
// EXT_visual_info: http://www.opengl.org/registry/specs/EXT/visual_info.txt
// 
// EXT_visual_rating: http://www.opengl.org/registry/specs/EXT/visual_rating.txt
// 
// INTEL_swap_event: http://www.opengl.org/registry/specs/INTEL/swap_event.txt
// 
// MESA_agp_offset: http://www.opengl.org/registry/specs/MESA/agp_offset.txt
// 
// MESA_copy_sub_buffer: http://www.opengl.org/registry/specs/MESA/copy_sub_buffer.txt
// 
// MESA_pixmap_colormap: http://www.opengl.org/registry/specs/MESA/pixmap_colormap.txt
// 
// MESA_release_buffers: http://www.opengl.org/registry/specs/MESA/release_buffers.txt
// 
// MESA_set_3dfx_mode: http://www.opengl.org/registry/specs/MESA/set_3dfx_mode.txt
// 
// NV_copy_image: http://www.opengl.org/registry/specs/NV/copy_image.txt
// 
// NV_float_buffer: http://www.opengl.org/registry/specs/NV/float_buffer.txt
// 
// NV_multisample_coverage: http://www.opengl.org/registry/specs/NV/multisample_coverage.txt
// 
// NV_present_video: http://www.opengl.org/registry/specs/NV/present_video.txt
// 
// NV_swap_group: http://www.opengl.org/registry/specs/NV/swap_group.txt
// 
// NV_video_capture: http://www.opengl.org/registry/specs/NV/video_capture.txt
// 
// NV_video_output: http://www.opengl.org/registry/specs/NV/video_output.txt
// 
// OML_swap_method: http://www.opengl.org/registry/specs/OML/swap_method.txt
// 
// OML_sync_control: http://www.opengl.org/registry/specs/OML/sync_control.txt
// 
// SGIS_multisample: http://www.opengl.org/registry/specs/SGIS/multisample.txt
// 
// SGIX_dmbuffer: http://www.opengl.org/registry/specs/SGIX/dmbuffer.txt
// 
// SGIX_fbconfig: http://www.opengl.org/registry/specs/SGIX/fbconfig.txt
// 
// SGIX_hyperpipe: http://www.opengl.org/registry/specs/SGIX/hyperpipe.txt
// 
// SGIX_pbuffer: http://www.opengl.org/registry/specs/SGIX/pbuffer.txt
// 
// SGIX_swap_barrier: http://www.opengl.org/registry/specs/SGIX/swap_barrier.txt
// 
// SGIX_swap_group: http://www.opengl.org/registry/specs/SGIX/swap_group.txt
// 
// SGIX_video_resize: http://www.opengl.org/registry/specs/SGIX/video_resize.txt
// 
// SGIX_video_source: http://www.opengl.org/registry/specs/SGIX/video_source.txt
// 
// SGIX_visual_select_group: http://www.opengl.org/registry/specs/SGIX/visual_select_group.txt
// 
// SGI_cushion: http://www.opengl.org/registry/specs/SGI/cushion.txt
// 
// SGI_make_current_read: http://www.opengl.org/registry/specs/SGI/make_current_read.txt
// 
// SGI_swap_control: http://www.opengl.org/registry/specs/SGI/swap_control.txt
// 
// SGI_video_sync: http://www.opengl.org/registry/specs/SGI/video_sync.txt
// 
// SUN_get_transparent_index: http://www.opengl.org/registry/specs/SUN/get_transparent_index.txt
// 
// VERSION_1_3
// 
// VERSION_1_4
// 
// glx: http://www.opengl.org/registry/specs/glx/.txt
// 
// http://www.opengl.org/sdk/docs/man
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
// //  ARB_create_context
// GLXContext (APIENTRYP ptrglCreateContextAttribsARB)(Display* dpy, GLXFBConfig config, GLXContext share_context, Bool direct, int* attrib_list);
// //  ARB_create_context_profile
// //  ARB_create_context_robustness
// //  ARB_fbconfig_float
// //  ARB_framebuffer_sRGB
// //  ARB_get_proc_address
// __GLXextFuncPtr (APIENTRYP ptrglGetProcAddressARB)(GLubyte* procName);
// //  ARB_multisample
// //  EXT_fbconfig_packed_float
// //  EXT_framebuffer_sRGB
// //  EXT_import_context
// Display * (APIENTRYP ptrglGetCurrentDisplayEXT)();
// int (APIENTRYP ptrglQueryContextInfoEXT)(Display* dpy, GLXContext context, int attribute, int* value);
// GLXContextID (APIENTRYP ptrglGetContextIDEXT)(const GLXContext context);
// GLXContext (APIENTRYP ptrglImportContextEXT)(Display* dpy, GLXContextID contextID);
// void (APIENTRYP ptrglFreeContextEXT)(Display* dpy, GLXContext context);
// //  EXT_swap_control
// void (APIENTRYP ptrglSwapIntervalEXT)(Display* dpy, GLXDrawable drawable, int interval);
// //  EXT_swap_control_tear
// //  EXT_texture_from_pixmap
// void (APIENTRYP ptrglBindTexImageEXT)(Display* dpy, GLXDrawable drawable, int buffer, int* attrib_list);
// void (APIENTRYP ptrglReleaseTexImageEXT)(Display* dpy, GLXDrawable drawable, int buffer);
// //  EXT_visual_info
// //  EXT_visual_rating
// //  INTEL_swap_event
// //  MESA_agp_offset
// unsigned int (APIENTRYP ptrglGetAGPOffsetMESA)(void* pointer);
// //  MESA_copy_sub_buffer
// void (APIENTRYP ptrglCopySubBufferMESA)(Display* dpy, GLXDrawable drawable, int x, int y, int width, int height);
// //  MESA_pixmap_colormap
// GLXPixmap (APIENTRYP ptrglCreateGLXPixmapMESA)(Display* dpy, XVisualInfo* visual, Pixmap pixmap, Colormap cmap);
// //  MESA_release_buffers
// Bool (APIENTRYP ptrglReleaseBuffersMESA)(Display* dpy, GLXDrawable drawable);
// //  MESA_set_3dfx_mode
// Bool (APIENTRYP ptrglSet3DfxModeMESA)(int mode);
// //  NV_copy_image
// void (APIENTRYP ptrglCopyImageSubDataNV)(Display* dpy, GLXContext srcCtx, GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, GLXContext dstCtx, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei width, GLsizei height, GLsizei depth);
// //  NV_float_buffer
// //  NV_multisample_coverage
// //  NV_present_video
// unsigned int * (APIENTRYP ptrglEnumerateVideoDevicesNV)(Display* dpy, int screen, int* nelements);
// int (APIENTRYP ptrglBindVideoDeviceNV)(Display* dpy, unsigned int video_slot, unsigned int video_device, int* attrib_list);
// //  NV_swap_group
// Bool (APIENTRYP ptrglJoinSwapGroupNV)(Display* dpy, GLXDrawable drawable, GLuint group);
// Bool (APIENTRYP ptrglBindSwapBarrierNV)(Display* dpy, GLuint group, GLuint barrier);
// Bool (APIENTRYP ptrglQuerySwapGroupNV)(Display* dpy, GLXDrawable drawable, GLuint* group, GLuint* barrier);
// Bool (APIENTRYP ptrglQueryMaxSwapGroupsNV)(Display* dpy, int screen, GLuint* maxGroups, GLuint* maxBarriers);
// Bool (APIENTRYP ptrglQueryFrameCountNV)(Display* dpy, int screen, GLuint* count);
// Bool (APIENTRYP ptrglResetFrameCountNV)(Display* dpy, int screen);
// //  NV_video_capture
// int (APIENTRYP ptrglBindVideoCaptureDeviceNV)(Display* dpy, unsigned int video_capture_slot, GLXVideoCaptureDeviceNV device);
// GLXVideoCaptureDeviceNV * (APIENTRYP ptrglEnumerateVideoCaptureDevicesNV)(Display* dpy, int screen, int* nelements);
// void (APIENTRYP ptrglLockVideoCaptureDeviceNV)(Display* dpy, GLXVideoCaptureDeviceNV device);
// int (APIENTRYP ptrglQueryVideoCaptureDeviceNV)(Display* dpy, GLXVideoCaptureDeviceNV device, int attribute, int* value);
// void (APIENTRYP ptrglReleaseVideoCaptureDeviceNV)(Display* dpy, GLXVideoCaptureDeviceNV device);
// //  NV_video_output
// int (APIENTRYP ptrglGetVideoDeviceNV)(Display* dpy, int screen, int numVideoDevices, GLXVideoDeviceNV* pVideoDevice);
// int (APIENTRYP ptrglReleaseVideoDeviceNV)(Display* dpy, int screen, GLXVideoDeviceNV VideoDevice);
// int (APIENTRYP ptrglBindVideoImageNV)(Display* dpy, GLXVideoDeviceNV VideoDevice, GLXPbuffer pbuf, int iVideoBuffer);
// int (APIENTRYP ptrglReleaseVideoImageNV)(Display* dpy, GLXPbuffer pbuf);
// int (APIENTRYP ptrglSendPbufferToVideoNV)(Display* dpy, GLXPbuffer pbuf, int iBufferType, unsigned long* pulCounterPbuffer, GLboolean bBlock);
// int (APIENTRYP ptrglGetVideoInfoNV)(Display* dpy, int screen, GLXVideoDeviceNV VideoDevice, unsigned long* pulCounterOutputPbuffer, unsigned long* pulCounterOutputVideo);
// //  OML_swap_method
// //  OML_sync_control
// Bool (APIENTRYP ptrglGetSyncValuesOML)(Display* dpy, GLXDrawable drawable, int64_t* ust, int64_t* msc, int64_t* sbc);
// Bool (APIENTRYP ptrglGetMscRateOML)(Display* dpy, GLXDrawable drawable, int32_t* numerator, int32_t* denominator);
// int64_t (APIENTRYP ptrglSwapBuffersMscOML)(Display* dpy, GLXDrawable drawable, int64_t target_msc, int64_t divisor, int64_t remainder);
// Bool (APIENTRYP ptrglWaitForMscOML)(Display* dpy, GLXDrawable drawable, int64_t target_msc, int64_t divisor, int64_t remainder, int64_t* ust, int64_t* msc, int64_t* sbc);
// Bool (APIENTRYP ptrglWaitForSbcOML)(Display* dpy, GLXDrawable drawable, int64_t target_sbc, int64_t* ust, int64_t* msc, int64_t* sbc);
// //  SGIS_multisample
// //  SGIX_dmbuffer
// Bool (APIENTRYP ptrglAssociateDMPbufferSGIX)(Display* dpy, GLXPbufferSGIX pbuffer, DMparams* params, DMbuffer dmbuffer);
// //  SGIX_fbconfig
// int (APIENTRYP ptrglGetFBConfigAttribSGIX)(Display* dpy, GLXFBConfigSGIX config, int attribute, int* value);
// GLXFBConfigSGIX * (APIENTRYP ptrglChooseFBConfigSGIX)(Display* dpy, int screen, int* attrib_list, int* nelements);
// GLXPixmap (APIENTRYP ptrglCreateGLXPixmapWithConfigSGIX)(Display* dpy, GLXFBConfigSGIX config, Pixmap pixmap);
// GLXContext (APIENTRYP ptrglCreateContextWithConfigSGIX)(Display* dpy, GLXFBConfigSGIX config, int render_type, GLXContext share_list, Bool direct);
// XVisualInfo * (APIENTRYP ptrglGetVisualFromFBConfigSGIX)(Display* dpy, GLXFBConfigSGIX config);
// GLXFBConfigSGIX (APIENTRYP ptrglGetFBConfigFromVisualSGIX)(Display* dpy, XVisualInfo* vis);
// //  SGIX_hyperpipe
// GLXHyperpipeNetworkSGIX * (APIENTRYP ptrglQueryHyperpipeNetworkSGIX)(Display* dpy, int* npipes);
// int (APIENTRYP ptrglHyperpipeConfigSGIX)(Display* dpy, int networkId, int npipes, GLXHyperpipeConfigSGIX* cfg, int* hpId);
// GLXHyperpipeConfigSGIX * (APIENTRYP ptrglQueryHyperpipeConfigSGIX)(Display* dpy, int hpId, int* npipes);
// int (APIENTRYP ptrglDestroyHyperpipeConfigSGIX)(Display* dpy, int hpId);
// int (APIENTRYP ptrglBindHyperpipeSGIX)(Display* dpy, int hpId);
// int (APIENTRYP ptrglQueryHyperpipeBestAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, void* attribList, void* returnAttribList);
// int (APIENTRYP ptrglHyperpipeAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, void* attribList);
// int (APIENTRYP ptrglQueryHyperpipeAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, void* returnAttribList);
// //  SGIX_pbuffer
// GLXPbufferSGIX (APIENTRYP ptrglCreateGLXPbufferSGIX)(Display* dpy, GLXFBConfigSGIX config, unsigned int width, unsigned int height, int* attrib_list);
// void (APIENTRYP ptrglDestroyGLXPbufferSGIX)(Display* dpy, GLXPbufferSGIX pbuf);
// int (APIENTRYP ptrglQueryGLXPbufferSGIX)(Display* dpy, GLXPbufferSGIX pbuf, int attribute, unsigned int* value);
// void (APIENTRYP ptrglSelectEventSGIX)(Display* dpy, GLXDrawable drawable, unsigned long mask);
// void (APIENTRYP ptrglGetSelectedEventSGIX)(Display* dpy, GLXDrawable drawable, unsigned long* mask);
// //  SGIX_swap_barrier
// void (APIENTRYP ptrglBindSwapBarrierSGIX)(Display* dpy, GLXDrawable drawable, int barrier);
// Bool (APIENTRYP ptrglQueryMaxSwapBarriersSGIX)(Display* dpy, int screen, int* max);
// //  SGIX_swap_group
// void (APIENTRYP ptrglJoinSwapGroupSGIX)(Display* dpy, GLXDrawable drawable, GLXDrawable member);
// //  SGIX_video_resize
// int (APIENTRYP ptrglBindChannelToWindowSGIX)(Display* display, int screen, int channel, Window window);
// int (APIENTRYP ptrglChannelRectSGIX)(Display* display, int screen, int channel, int x, int y, int w, int h);
// int (APIENTRYP ptrglQueryChannelRectSGIX)(Display* display, int screen, int channel, int* dx, int* dy, int* dw, int* dh);
// int (APIENTRYP ptrglQueryChannelDeltasSGIX)(Display* display, int screen, int channel, int* x, int* y, int* w, int* h);
// int (APIENTRYP ptrglChannelRectSyncSGIX)(Display* display, int screen, int channel, GLenum synctype);
// //  SGIX_video_source
// GLXVideoSourceSGIX (APIENTRYP ptrglCreateGLXVideoSourceSGIX)(Display* display, int screen, VLServer server, VLPath path, int nodeClass, VLNode drainNode);
// void (APIENTRYP ptrglDestroyGLXVideoSourceSGIX)(Display* dpy, GLXVideoSourceSGIX glxvideosource);
// //  SGIX_visual_select_group
// //  SGI_cushion
// void (APIENTRYP ptrglCushionSGI)(Display* dpy, Window window, float cushion);
// //  SGI_make_current_read
// Bool (APIENTRYP ptrglMakeCurrentReadSGI)(Display* dpy, GLXDrawable draw, GLXDrawable read, GLXContext ctx);
// GLXDrawable (APIENTRYP ptrglGetCurrentReadDrawableSGI)();
// //  SGI_swap_control
// int (APIENTRYP ptrglSwapIntervalSGI)(int interval);
// //  SGI_video_sync
// int (APIENTRYP ptrglGetVideoSyncSGI)(unsigned int* count);
// int (APIENTRYP ptrglWaitVideoSyncSGI)(int divisor, int remainder, unsigned int* count);
// //  SUN_get_transparent_index
// Status (APIENTRYP ptrglGetTransparentIndexSUN)(Display* dpy, Window overlay, Window underlay, long* pTransparentIndex);
// //  VERSION_1_3
// GLXFBConfig * (APIENTRYP ptrglGetFBConfigs)(Display* dpy, int screen, int* nelements);
// GLXFBConfig * (APIENTRYP ptrglChooseFBConfig)(Display* dpy, int screen, int* attrib_list, int* nelements);
// int (APIENTRYP ptrglGetFBConfigAttrib)(Display* dpy, GLXFBConfig config, int attribute, int* value);
// XVisualInfo * (APIENTRYP ptrglGetVisualFromFBConfig)(Display* dpy, GLXFBConfig config);
// GLXWindow (APIENTRYP ptrglCreateWindow)(Display* dpy, GLXFBConfig config, Window win, int* attrib_list);
// void (APIENTRYP ptrglDestroyWindow)(Display* dpy, GLXWindow win);
// GLXPixmap (APIENTRYP ptrglCreatePixmap)(Display* dpy, GLXFBConfig config, Pixmap pixmap, int* attrib_list);
// void (APIENTRYP ptrglDestroyPixmap)(Display* dpy, GLXPixmap pixmap);
// GLXPbuffer (APIENTRYP ptrglCreatePbuffer)(Display* dpy, GLXFBConfig config, int* attrib_list);
// void (APIENTRYP ptrglDestroyPbuffer)(Display* dpy, GLXPbuffer pbuf);
// void (APIENTRYP ptrglQueryDrawable)(Display* dpy, GLXDrawable draw, int attribute, unsigned int* value);
// GLXContext (APIENTRYP ptrglCreateNewContext)(Display* dpy, GLXFBConfig config, int render_type, GLXContext share_list, Bool direct);
// Bool (APIENTRYP ptrglMakeContextCurrent)(Display* dpy, GLXDrawable draw, GLXDrawable read, GLXContext ctx);
// GLXDrawable (APIENTRYP ptrglGetCurrentReadDrawable)();
// Display * (APIENTRYP ptrglGetCurrentDisplay)();
// int (APIENTRYP ptrglQueryContext)(Display* dpy, GLXContext ctx, int attribute, int* value);
// void (APIENTRYP ptrglSelectEvent)(Display* dpy, GLXDrawable draw, unsigned long event_mask);
// void (APIENTRYP ptrglGetSelectedEvent)(Display* dpy, GLXDrawable draw, unsigned long* event_mask);
// //  VERSION_1_4
// __GLXextFuncPtr (APIENTRYP ptrglGetProcAddress)(GLubyte* procName);
// //  glx
// void (APIENTRYP ptrglRender)();
// void (APIENTRYP ptrglRenderLarge)();
// void (APIENTRYP ptrglCreateContext)(GLint gc_id, GLint screen, GLint visual, GLint share_list);
// void (APIENTRYP ptrglDestroyContext)(GLint context);
// void (APIENTRYP ptrglMakeCurrent)(GLint drawable, GLint context);
// void (APIENTRYP ptrglIsDirect)(GLint dpy, GLint context);
// void (APIENTRYP ptrglQueryVersion)(GLint* major, GLint* minor);
// void (APIENTRYP ptrglWaitGL)(GLint context);
// void (APIENTRYP ptrglWaitX)();
// void (APIENTRYP ptrglCopyContext)(GLint source, GLint dest, GLint mask);
// void (APIENTRYP ptrglSwapBuffers)(GLint drawable);
// void (APIENTRYP ptrglUseXFont)(GLint font, GLint first, GLint count, GLint list_base);
// void (APIENTRYP ptrglCreateGLXPixmap)(GLint visual, GLint pixmap, GLint glxpixmap);
// void (APIENTRYP ptrglGetVisualConfigs)();
// void (APIENTRYP ptrglDestroyGLXPixmap)(GLint pixmap);
// void (APIENTRYP ptrglVendorPrivate)();
// void (APIENTRYP ptrglVendorPrivateWithReply)();
// void (APIENTRYP ptrglQueryExtensionsString)(GLint screen);
// void (APIENTRYP ptrglQueryServerString)(GLint screen, GLint name);
// void (APIENTRYP ptrglClientInfo)();
// void (APIENTRYP ptrglGetFBConfigs)();
// void (APIENTRYP ptrglCreatePixmap)(GLint config, GLint pixmap, GLint glxpixmap);
// void (APIENTRYP ptrglDestroyPixmap)(GLint glxpixmap);
// void (APIENTRYP ptrglCreateNewContext)(GLint config, GLint render_type, GLint share_list, GLint direct);
// void (APIENTRYP ptrglQueryContext)();
// void (APIENTRYP ptrglMakeContextCurrent)(GLint drawable, GLint readdrawable, GLint context);
// void (APIENTRYP ptrglCreatePbuffer)(GLint config, GLint pbuffer);
// void (APIENTRYP ptrglDestroyPbuffer)(GLint pbuffer);
// void (APIENTRYP ptrglGetDrawableAttributes)(GLint drawable);
// void (APIENTRYP ptrglChangeDrawableAttributes)(GLint drawable);
// void (APIENTRYP ptrglCreateWindow)(GLint config, GLint window, GLint glxwindow);
// void (APIENTRYP ptrglDestroyWindow)(GLint glxwindow);
// void (APIENTRYP ptrglSwapIntervalSGI)();
// void (APIENTRYP ptrglMakeCurrentReadSGI)(GLint drawable, GLint readdrawable, GLint context);
// void (APIENTRYP ptrglCreateGLXVideoSourceSGIX)(GLint dpy, GLint screen, GLint server, GLint path, GLint class, GLint node);
// void (APIENTRYP ptrglDestroyGLXVideoSourceSGIX)(GLint dpy, GLint glxvideosource);
// void (APIENTRYP ptrglQueryContextInfoEXT)();
// void (APIENTRYP ptrglGetFBConfigsSGIX)();
// void (APIENTRYP ptrglCreateContextWithConfigSGIX)(GLint gc_id, GLint screen, GLint config, GLint share_list);
// void (APIENTRYP ptrglCreateGLXPixmapWithConfigSGIX)(GLint config, GLint pixmap, GLint glxpixmap);
// void (APIENTRYP ptrglCreateGLXPbufferSGIX)(GLint config, GLint pbuffer);
// void (APIENTRYP ptrglDestroyGLXPbufferSGIX)(GLint pbuffer);
// void (APIENTRYP ptrglChangeDrawableAttributesSGIX)(GLint drawable);
// void (APIENTRYP ptrglGetDrawableAttributesSGIX)(GLint drawable);
// void (APIENTRYP ptrglJoinSwapGroupSGIX)(GLint window, GLint group);
// void (APIENTRYP ptrglBindSwapBarrierSGIX)(GLint window, GLint barrier);
// void (APIENTRYP ptrglQueryMaxSwapBarriersSGIX)();
// GLXHyperpipeNetworkSGIX * (APIENTRYP ptrglQueryHyperpipeNetworkSGIX)(Display* dpy, int* npipes);
// int (APIENTRYP ptrglHyperpipeConfigSGIX)(Display* dpy, int networkId, int npipes, GLXHyperpipeConfigSGIX* cfg, int* hpId);
// GLXHyperpipeConfigSGIX * (APIENTRYP ptrglQueryHyperpipeConfigSGIX)(Display* dpy, int hpId, int* npipes);
// int (APIENTRYP ptrglDestroyHyperpipeConfigSGIX)(Display* dpy, int hpId);
// int (APIENTRYP ptrglBindHyperpipeSGIX)(Display* dpy, int hpId);
// int (APIENTRYP ptrglQueryHyperpipeBestAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, GLvoid* attribList, GLvoid* returnAttribList);
// int (APIENTRYP ptrglHyperpipeAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, void* attribList);
// int (APIENTRYP ptrglQueryHyperpipeAttribSGIX)(Display* dpy, int timeSlice, int attrib, int size, void* returnAttribList);
// 
// //  ARB_create_context
// GLXContext goglCreateContextAttribsARB(Display* dpy, GLXFBConfig config, GLXContext share_context, Bool direct, int* attrib_list) {
// 	return (*ptrglCreateContextAttribsARB)(dpy, config, share_context, direct, attrib_list);
// }
// //  ARB_create_context_profile
// //  ARB_create_context_robustness
// //  ARB_fbconfig_float
// //  ARB_framebuffer_sRGB
// //  ARB_get_proc_address
// __GLXextFuncPtr goglGetProcAddressARB(GLubyte* procName) {
// 	return (*ptrglGetProcAddressARB)(procName);
// }
// //  ARB_multisample
// //  EXT_fbconfig_packed_float
// //  EXT_framebuffer_sRGB
// //  EXT_import_context
// Display * goglGetCurrentDisplayEXT() {
// 	return (*ptrglGetCurrentDisplayEXT)();
// }
// int goglQueryContextInfoEXT(Display* dpy, GLXContext context, int attribute, int* value) {
// 	return (*ptrglQueryContextInfoEXT)(dpy, context, attribute, value);
// }
// GLXContextID goglGetContextIDEXT(const GLXContext context) {
// 	return (*ptrglGetContextIDEXT)(context);
// }
// GLXContext goglImportContextEXT(Display* dpy, GLXContextID contextID) {
// 	return (*ptrglImportContextEXT)(dpy, contextID);
// }
// void goglFreeContextEXT(Display* dpy, GLXContext context) {
// 	(*ptrglFreeContextEXT)(dpy, context);
// }
// //  EXT_swap_control
// void goglSwapIntervalEXT(Display* dpy, GLXDrawable drawable, int interval) {
// 	(*ptrglSwapIntervalEXT)(dpy, drawable, interval);
// }
// //  EXT_swap_control_tear
// //  EXT_texture_from_pixmap
// void goglBindTexImageEXT(Display* dpy, GLXDrawable drawable, int buffer, int* attrib_list) {
// 	(*ptrglBindTexImageEXT)(dpy, drawable, buffer, attrib_list);
// }
// void goglReleaseTexImageEXT(Display* dpy, GLXDrawable drawable, int buffer) {
// 	(*ptrglReleaseTexImageEXT)(dpy, drawable, buffer);
// }
// //  EXT_visual_info
// //  EXT_visual_rating
// //  INTEL_swap_event
// //  MESA_agp_offset
// unsigned int goglGetAGPOffsetMESA(void* pointer) {
// 	return (*ptrglGetAGPOffsetMESA)(pointer);
// }
// //  MESA_copy_sub_buffer
// void goglCopySubBufferMESA(Display* dpy, GLXDrawable drawable, int x, int y, int width, int height) {
// 	(*ptrglCopySubBufferMESA)(dpy, drawable, x, y, width, height);
// }
// //  MESA_pixmap_colormap
// GLXPixmap goglCreateGLXPixmapMESA(Display* dpy, XVisualInfo* visual, Pixmap pixmap, Colormap cmap) {
// 	return (*ptrglCreateGLXPixmapMESA)(dpy, visual, pixmap, cmap);
// }
// //  MESA_release_buffers
// Bool goglReleaseBuffersMESA(Display* dpy, GLXDrawable drawable) {
// 	return (*ptrglReleaseBuffersMESA)(dpy, drawable);
// }
// //  MESA_set_3dfx_mode
// Bool goglSet3DfxModeMESA(int mode) {
// 	return (*ptrglSet3DfxModeMESA)(mode);
// }
// //  NV_copy_image
// void goglCopyImageSubDataNV(Display* dpy, GLXContext srcCtx, GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, GLXContext dstCtx, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei width, GLsizei height, GLsizei depth) {
// 	(*ptrglCopyImageSubDataNV)(dpy, srcCtx, srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstCtx, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, width, height, depth);
// }
// //  NV_float_buffer
// //  NV_multisample_coverage
// //  NV_present_video
// unsigned int * goglEnumerateVideoDevicesNV(Display* dpy, int screen, int* nelements) {
// 	return (*ptrglEnumerateVideoDevicesNV)(dpy, screen, nelements);
// }
// int goglBindVideoDeviceNV(Display* dpy, unsigned int video_slot, unsigned int video_device, int* attrib_list) {
// 	return (*ptrglBindVideoDeviceNV)(dpy, video_slot, video_device, attrib_list);
// }
// //  NV_swap_group
// Bool goglJoinSwapGroupNV(Display* dpy, GLXDrawable drawable, GLuint group) {
// 	return (*ptrglJoinSwapGroupNV)(dpy, drawable, group);
// }
// Bool goglBindSwapBarrierNV(Display* dpy, GLuint group, GLuint barrier) {
// 	return (*ptrglBindSwapBarrierNV)(dpy, group, barrier);
// }
// Bool goglQuerySwapGroupNV(Display* dpy, GLXDrawable drawable, GLuint* group, GLuint* barrier) {
// 	return (*ptrglQuerySwapGroupNV)(dpy, drawable, group, barrier);
// }
// Bool goglQueryMaxSwapGroupsNV(Display* dpy, int screen, GLuint* maxGroups, GLuint* maxBarriers) {
// 	return (*ptrglQueryMaxSwapGroupsNV)(dpy, screen, maxGroups, maxBarriers);
// }
// Bool goglQueryFrameCountNV(Display* dpy, int screen, GLuint* count) {
// 	return (*ptrglQueryFrameCountNV)(dpy, screen, count);
// }
// Bool goglResetFrameCountNV(Display* dpy, int screen) {
// 	return (*ptrglResetFrameCountNV)(dpy, screen);
// }
// //  NV_video_capture
// int goglBindVideoCaptureDeviceNV(Display* dpy, unsigned int video_capture_slot, GLXVideoCaptureDeviceNV device) {
// 	return (*ptrglBindVideoCaptureDeviceNV)(dpy, video_capture_slot, device);
// }
// GLXVideoCaptureDeviceNV * goglEnumerateVideoCaptureDevicesNV(Display* dpy, int screen, int* nelements) {
// 	return (*ptrglEnumerateVideoCaptureDevicesNV)(dpy, screen, nelements);
// }
// void goglLockVideoCaptureDeviceNV(Display* dpy, GLXVideoCaptureDeviceNV device) {
// 	(*ptrglLockVideoCaptureDeviceNV)(dpy, device);
// }
// int goglQueryVideoCaptureDeviceNV(Display* dpy, GLXVideoCaptureDeviceNV device, int attribute, int* value) {
// 	return (*ptrglQueryVideoCaptureDeviceNV)(dpy, device, attribute, value);
// }
// void goglReleaseVideoCaptureDeviceNV(Display* dpy, GLXVideoCaptureDeviceNV device) {
// 	(*ptrglReleaseVideoCaptureDeviceNV)(dpy, device);
// }
// //  NV_video_output
// int goglGetVideoDeviceNV(Display* dpy, int screen, int numVideoDevices, GLXVideoDeviceNV* pVideoDevice) {
// 	return (*ptrglGetVideoDeviceNV)(dpy, screen, numVideoDevices, pVideoDevice);
// }
// int goglReleaseVideoDeviceNV(Display* dpy, int screen, GLXVideoDeviceNV VideoDevice) {
// 	return (*ptrglReleaseVideoDeviceNV)(dpy, screen, VideoDevice);
// }
// int goglBindVideoImageNV(Display* dpy, GLXVideoDeviceNV VideoDevice, GLXPbuffer pbuf, int iVideoBuffer) {
// 	return (*ptrglBindVideoImageNV)(dpy, VideoDevice, pbuf, iVideoBuffer);
// }
// int goglReleaseVideoImageNV(Display* dpy, GLXPbuffer pbuf) {
// 	return (*ptrglReleaseVideoImageNV)(dpy, pbuf);
// }
// int goglSendPbufferToVideoNV(Display* dpy, GLXPbuffer pbuf, int iBufferType, unsigned long* pulCounterPbuffer, GLboolean bBlock) {
// 	return (*ptrglSendPbufferToVideoNV)(dpy, pbuf, iBufferType, pulCounterPbuffer, bBlock);
// }
// int goglGetVideoInfoNV(Display* dpy, int screen, GLXVideoDeviceNV VideoDevice, unsigned long* pulCounterOutputPbuffer, unsigned long* pulCounterOutputVideo) {
// 	return (*ptrglGetVideoInfoNV)(dpy, screen, VideoDevice, pulCounterOutputPbuffer, pulCounterOutputVideo);
// }
// //  OML_swap_method
// //  OML_sync_control
// Bool goglGetSyncValuesOML(Display* dpy, GLXDrawable drawable, int64_t* ust, int64_t* msc, int64_t* sbc) {
// 	return (*ptrglGetSyncValuesOML)(dpy, drawable, ust, msc, sbc);
// }
// Bool goglGetMscRateOML(Display* dpy, GLXDrawable drawable, int32_t* numerator, int32_t* denominator) {
// 	return (*ptrglGetMscRateOML)(dpy, drawable, numerator, denominator);
// }
// int64_t goglSwapBuffersMscOML(Display* dpy, GLXDrawable drawable, int64_t target_msc, int64_t divisor, int64_t remainder) {
// 	return (*ptrglSwapBuffersMscOML)(dpy, drawable, target_msc, divisor, remainder);
// }
// Bool goglWaitForMscOML(Display* dpy, GLXDrawable drawable, int64_t target_msc, int64_t divisor, int64_t remainder, int64_t* ust, int64_t* msc, int64_t* sbc) {
// 	return (*ptrglWaitForMscOML)(dpy, drawable, target_msc, divisor, remainder, ust, msc, sbc);
// }
// Bool goglWaitForSbcOML(Display* dpy, GLXDrawable drawable, int64_t target_sbc, int64_t* ust, int64_t* msc, int64_t* sbc) {
// 	return (*ptrglWaitForSbcOML)(dpy, drawable, target_sbc, ust, msc, sbc);
// }
// //  SGIS_multisample
// //  SGIX_dmbuffer
// Bool goglAssociateDMPbufferSGIX(Display* dpy, GLXPbufferSGIX pbuffer, DMparams* params, DMbuffer dmbuffer) {
// 	return (*ptrglAssociateDMPbufferSGIX)(dpy, pbuffer, params, dmbuffer);
// }
// //  SGIX_fbconfig
// int goglGetFBConfigAttribSGIX(Display* dpy, GLXFBConfigSGIX config, int attribute, int* value) {
// 	return (*ptrglGetFBConfigAttribSGIX)(dpy, config, attribute, value);
// }
// GLXFBConfigSGIX * goglChooseFBConfigSGIX(Display* dpy, int screen, int* attrib_list, int* nelements) {
// 	return (*ptrglChooseFBConfigSGIX)(dpy, screen, attrib_list, nelements);
// }
// GLXPixmap goglCreateGLXPixmapWithConfigSGIX(Display* dpy, GLXFBConfigSGIX config, Pixmap pixmap) {
// 	return (*ptrglCreateGLXPixmapWithConfigSGIX)(dpy, config, pixmap);
// }
// GLXContext goglCreateContextWithConfigSGIX(Display* dpy, GLXFBConfigSGIX config, int render_type, GLXContext share_list, Bool direct) {
// 	return (*ptrglCreateContextWithConfigSGIX)(dpy, config, render_type, share_list, direct);
// }
// XVisualInfo * goglGetVisualFromFBConfigSGIX(Display* dpy, GLXFBConfigSGIX config) {
// 	return (*ptrglGetVisualFromFBConfigSGIX)(dpy, config);
// }
// GLXFBConfigSGIX goglGetFBConfigFromVisualSGIX(Display* dpy, XVisualInfo* vis) {
// 	return (*ptrglGetFBConfigFromVisualSGIX)(dpy, vis);
// }
// //  SGIX_hyperpipe
// GLXHyperpipeNetworkSGIX * goglQueryHyperpipeNetworkSGIX(Display* dpy, int* npipes) {
// 	return (*ptrglQueryHyperpipeNetworkSGIX)(dpy, npipes);
// }
// int goglHyperpipeConfigSGIX(Display* dpy, int networkId, int npipes, GLXHyperpipeConfigSGIX* cfg, int* hpId) {
// 	return (*ptrglHyperpipeConfigSGIX)(dpy, networkId, npipes, cfg, hpId);
// }
// GLXHyperpipeConfigSGIX * goglQueryHyperpipeConfigSGIX(Display* dpy, int hpId, int* npipes) {
// 	return (*ptrglQueryHyperpipeConfigSGIX)(dpy, hpId, npipes);
// }
// int goglDestroyHyperpipeConfigSGIX(Display* dpy, int hpId) {
// 	return (*ptrglDestroyHyperpipeConfigSGIX)(dpy, hpId);
// }
// int goglBindHyperpipeSGIX(Display* dpy, int hpId) {
// 	return (*ptrglBindHyperpipeSGIX)(dpy, hpId);
// }
// int goglQueryHyperpipeBestAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, void* attribList, void* returnAttribList) {
// 	return (*ptrglQueryHyperpipeBestAttribSGIX)(dpy, timeSlice, attrib, size, attribList, returnAttribList);
// }
// int goglHyperpipeAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, void* attribList) {
// 	return (*ptrglHyperpipeAttribSGIX)(dpy, timeSlice, attrib, size, attribList);
// }
// int goglQueryHyperpipeAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, void* returnAttribList) {
// 	return (*ptrglQueryHyperpipeAttribSGIX)(dpy, timeSlice, attrib, size, returnAttribList);
// }
// //  SGIX_pbuffer
// GLXPbufferSGIX goglCreateGLXPbufferSGIX(Display* dpy, GLXFBConfigSGIX config, unsigned int width, unsigned int height, int* attrib_list) {
// 	return (*ptrglCreateGLXPbufferSGIX)(dpy, config, width, height, attrib_list);
// }
// void goglDestroyGLXPbufferSGIX(Display* dpy, GLXPbufferSGIX pbuf) {
// 	(*ptrglDestroyGLXPbufferSGIX)(dpy, pbuf);
// }
// int goglQueryGLXPbufferSGIX(Display* dpy, GLXPbufferSGIX pbuf, int attribute, unsigned int* value) {
// 	return (*ptrglQueryGLXPbufferSGIX)(dpy, pbuf, attribute, value);
// }
// void goglSelectEventSGIX(Display* dpy, GLXDrawable drawable, unsigned long mask) {
// 	(*ptrglSelectEventSGIX)(dpy, drawable, mask);
// }
// void goglGetSelectedEventSGIX(Display* dpy, GLXDrawable drawable, unsigned long* mask) {
// 	(*ptrglGetSelectedEventSGIX)(dpy, drawable, mask);
// }
// //  SGIX_swap_barrier
// void goglBindSwapBarrierSGIX(Display* dpy, GLXDrawable drawable, int barrier) {
// 	(*ptrglBindSwapBarrierSGIX)(dpy, drawable, barrier);
// }
// Bool goglQueryMaxSwapBarriersSGIX(Display* dpy, int screen, int* max) {
// 	return (*ptrglQueryMaxSwapBarriersSGIX)(dpy, screen, max);
// }
// //  SGIX_swap_group
// void goglJoinSwapGroupSGIX(Display* dpy, GLXDrawable drawable, GLXDrawable member) {
// 	(*ptrglJoinSwapGroupSGIX)(dpy, drawable, member);
// }
// //  SGIX_video_resize
// int goglBindChannelToWindowSGIX(Display* display, int screen, int channel, Window window) {
// 	return (*ptrglBindChannelToWindowSGIX)(display, screen, channel, window);
// }
// int goglChannelRectSGIX(Display* display, int screen, int channel, int x, int y, int w, int h) {
// 	return (*ptrglChannelRectSGIX)(display, screen, channel, x, y, w, h);
// }
// int goglQueryChannelRectSGIX(Display* display, int screen, int channel, int* dx, int* dy, int* dw, int* dh) {
// 	return (*ptrglQueryChannelRectSGIX)(display, screen, channel, dx, dy, dw, dh);
// }
// int goglQueryChannelDeltasSGIX(Display* display, int screen, int channel, int* x, int* y, int* w, int* h) {
// 	return (*ptrglQueryChannelDeltasSGIX)(display, screen, channel, x, y, w, h);
// }
// int goglChannelRectSyncSGIX(Display* display, int screen, int channel, GLenum synctype) {
// 	return (*ptrglChannelRectSyncSGIX)(display, screen, channel, synctype);
// }
// //  SGIX_video_source
// GLXVideoSourceSGIX goglCreateGLXVideoSourceSGIX(Display* display, int screen, VLServer server, VLPath path, int nodeClass, VLNode drainNode) {
// 	return (*ptrglCreateGLXVideoSourceSGIX)(display, screen, server, path, nodeClass, drainNode);
// }
// void goglDestroyGLXVideoSourceSGIX(Display* dpy, GLXVideoSourceSGIX glxvideosource) {
// 	(*ptrglDestroyGLXVideoSourceSGIX)(dpy, glxvideosource);
// }
// //  SGIX_visual_select_group
// //  SGI_cushion
// void goglCushionSGI(Display* dpy, Window window, float cushion) {
// 	(*ptrglCushionSGI)(dpy, window, cushion);
// }
// //  SGI_make_current_read
// Bool goglMakeCurrentReadSGI(Display* dpy, GLXDrawable draw, GLXDrawable read, GLXContext ctx) {
// 	return (*ptrglMakeCurrentReadSGI)(dpy, draw, read, ctx);
// }
// GLXDrawable goglGetCurrentReadDrawableSGI() {
// 	return (*ptrglGetCurrentReadDrawableSGI)();
// }
// //  SGI_swap_control
// int goglSwapIntervalSGI(int interval) {
// 	return (*ptrglSwapIntervalSGI)(interval);
// }
// //  SGI_video_sync
// int goglGetVideoSyncSGI(unsigned int* count) {
// 	return (*ptrglGetVideoSyncSGI)(count);
// }
// int goglWaitVideoSyncSGI(int divisor, int remainder, unsigned int* count) {
// 	return (*ptrglWaitVideoSyncSGI)(divisor, remainder, count);
// }
// //  SUN_get_transparent_index
// Status goglGetTransparentIndexSUN(Display* dpy, Window overlay, Window underlay, long* pTransparentIndex) {
// 	return (*ptrglGetTransparentIndexSUN)(dpy, overlay, underlay, pTransparentIndex);
// }
// //  VERSION_1_3
// GLXFBConfig * goglGetFBConfigs(Display* dpy, int screen, int* nelements) {
// 	return (*ptrglGetFBConfigs)(dpy, screen, nelements);
// }
// GLXFBConfig * goglChooseFBConfig(Display* dpy, int screen, int* attrib_list, int* nelements) {
// 	return (*ptrglChooseFBConfig)(dpy, screen, attrib_list, nelements);
// }
// int goglGetFBConfigAttrib(Display* dpy, GLXFBConfig config, int attribute, int* value) {
// 	return (*ptrglGetFBConfigAttrib)(dpy, config, attribute, value);
// }
// XVisualInfo * goglGetVisualFromFBConfig(Display* dpy, GLXFBConfig config) {
// 	return (*ptrglGetVisualFromFBConfig)(dpy, config);
// }
// GLXWindow goglCreateWindow(Display* dpy, GLXFBConfig config, Window win, int* attrib_list) {
// 	return (*ptrglCreateWindow)(dpy, config, win, attrib_list);
// }
// void goglDestroyWindow(Display* dpy, GLXWindow win) {
// 	(*ptrglDestroyWindow)(dpy, win);
// }
// GLXPixmap goglCreatePixmap(Display* dpy, GLXFBConfig config, Pixmap pixmap, int* attrib_list) {
// 	return (*ptrglCreatePixmap)(dpy, config, pixmap, attrib_list);
// }
// void goglDestroyPixmap(Display* dpy, GLXPixmap pixmap) {
// 	(*ptrglDestroyPixmap)(dpy, pixmap);
// }
// GLXPbuffer goglCreatePbuffer(Display* dpy, GLXFBConfig config, int* attrib_list) {
// 	return (*ptrglCreatePbuffer)(dpy, config, attrib_list);
// }
// void goglDestroyPbuffer(Display* dpy, GLXPbuffer pbuf) {
// 	(*ptrglDestroyPbuffer)(dpy, pbuf);
// }
// void goglQueryDrawable(Display* dpy, GLXDrawable draw, int attribute, unsigned int* value) {
// 	(*ptrglQueryDrawable)(dpy, draw, attribute, value);
// }
// GLXContext goglCreateNewContext(Display* dpy, GLXFBConfig config, int render_type, GLXContext share_list, Bool direct) {
// 	return (*ptrglCreateNewContext)(dpy, config, render_type, share_list, direct);
// }
// Bool goglMakeContextCurrent(Display* dpy, GLXDrawable draw, GLXDrawable read, GLXContext ctx) {
// 	return (*ptrglMakeContextCurrent)(dpy, draw, read, ctx);
// }
// GLXDrawable goglGetCurrentReadDrawable() {
// 	return (*ptrglGetCurrentReadDrawable)();
// }
// Display * goglGetCurrentDisplay() {
// 	return (*ptrglGetCurrentDisplay)();
// }
// int goglQueryContext(Display* dpy, GLXContext ctx, int attribute, int* value) {
// 	return (*ptrglQueryContext)(dpy, ctx, attribute, value);
// }
// void goglSelectEvent(Display* dpy, GLXDrawable draw, unsigned long event_mask) {
// 	(*ptrglSelectEvent)(dpy, draw, event_mask);
// }
// void goglGetSelectedEvent(Display* dpy, GLXDrawable draw, unsigned long* event_mask) {
// 	(*ptrglGetSelectedEvent)(dpy, draw, event_mask);
// }
// //  VERSION_1_4
// __GLXextFuncPtr goglGetProcAddress(GLubyte* procName) {
// 	return (*ptrglGetProcAddress)(procName);
// }
// //  glx
// void goglRender() {
// 	(*ptrglRender)();
// }
// void goglRenderLarge() {
// 	(*ptrglRenderLarge)();
// }
// void goglCreateContext(GLint gc_id, GLint screen, GLint visual, GLint share_list) {
// 	(*ptrglCreateContext)(gc_id, screen, visual, share_list);
// }
// void goglDestroyContext(GLint context) {
// 	(*ptrglDestroyContext)(context);
// }
// void goglMakeCurrent(GLint drawable, GLint context) {
// 	(*ptrglMakeCurrent)(drawable, context);
// }
// void goglIsDirect(GLint dpy, GLint context) {
// 	(*ptrglIsDirect)(dpy, context);
// }
// void goglQueryVersion(GLint* major, GLint* minor) {
// 	(*ptrglQueryVersion)(major, minor);
// }
// void goglWaitGL(GLint context) {
// 	(*ptrglWaitGL)(context);
// }
// void goglWaitX() {
// 	(*ptrglWaitX)();
// }
// void goglCopyContext(GLint source, GLint dest, GLint mask) {
// 	(*ptrglCopyContext)(source, dest, mask);
// }
// void goglSwapBuffers(GLint drawable) {
// 	(*ptrglSwapBuffers)(drawable);
// }
// void goglUseXFont(GLint font, GLint first, GLint count, GLint list_base) {
// 	(*ptrglUseXFont)(font, first, count, list_base);
// }
// void goglCreateGLXPixmap(GLint visual, GLint pixmap, GLint glxpixmap) {
// 	(*ptrglCreateGLXPixmap)(visual, pixmap, glxpixmap);
// }
// void goglGetVisualConfigs() {
// 	(*ptrglGetVisualConfigs)();
// }
// void goglDestroyGLXPixmap(GLint pixmap) {
// 	(*ptrglDestroyGLXPixmap)(pixmap);
// }
// void goglVendorPrivate() {
// 	(*ptrglVendorPrivate)();
// }
// void goglVendorPrivateWithReply() {
// 	(*ptrglVendorPrivateWithReply)();
// }
// void goglQueryExtensionsString(GLint screen) {
// 	(*ptrglQueryExtensionsString)(screen);
// }
// void goglQueryServerString(GLint screen, GLint name) {
// 	(*ptrglQueryServerString)(screen, name);
// }
// void goglClientInfo() {
// 	(*ptrglClientInfo)();
// }
// void goglGetFBConfigs() {
// 	(*ptrglGetFBConfigs)();
// }
// void goglCreatePixmap(GLint config, GLint pixmap, GLint glxpixmap) {
// 	(*ptrglCreatePixmap)(config, pixmap, glxpixmap);
// }
// void goglDestroyPixmap(GLint glxpixmap) {
// 	(*ptrglDestroyPixmap)(glxpixmap);
// }
// void goglCreateNewContext(GLint config, GLint render_type, GLint share_list, GLint direct) {
// 	(*ptrglCreateNewContext)(config, render_type, share_list, direct);
// }
// void goglQueryContext() {
// 	(*ptrglQueryContext)();
// }
// void goglMakeContextCurrent(GLint drawable, GLint readdrawable, GLint context) {
// 	(*ptrglMakeContextCurrent)(drawable, readdrawable, context);
// }
// void goglCreatePbuffer(GLint config, GLint pbuffer) {
// 	(*ptrglCreatePbuffer)(config, pbuffer);
// }
// void goglDestroyPbuffer(GLint pbuffer) {
// 	(*ptrglDestroyPbuffer)(pbuffer);
// }
// void goglGetDrawableAttributes(GLint drawable) {
// 	(*ptrglGetDrawableAttributes)(drawable);
// }
// void goglChangeDrawableAttributes(GLint drawable) {
// 	(*ptrglChangeDrawableAttributes)(drawable);
// }
// void goglCreateWindow(GLint config, GLint window, GLint glxwindow) {
// 	(*ptrglCreateWindow)(config, window, glxwindow);
// }
// void goglDestroyWindow(GLint glxwindow) {
// 	(*ptrglDestroyWindow)(glxwindow);
// }
// void goglSwapIntervalSGI() {
// 	(*ptrglSwapIntervalSGI)();
// }
// void goglMakeCurrentReadSGI(GLint drawable, GLint readdrawable, GLint context) {
// 	(*ptrglMakeCurrentReadSGI)(drawable, readdrawable, context);
// }
// void goglCreateGLXVideoSourceSGIX(GLint dpy, GLint screen, GLint server, GLint path, GLint class, GLint node) {
// 	(*ptrglCreateGLXVideoSourceSGIX)(dpy, screen, server, path, class, node);
// }
// void goglDestroyGLXVideoSourceSGIX(GLint dpy, GLint glxvideosource) {
// 	(*ptrglDestroyGLXVideoSourceSGIX)(dpy, glxvideosource);
// }
// void goglQueryContextInfoEXT() {
// 	(*ptrglQueryContextInfoEXT)();
// }
// void goglGetFBConfigsSGIX() {
// 	(*ptrglGetFBConfigsSGIX)();
// }
// void goglCreateContextWithConfigSGIX(GLint gc_id, GLint screen, GLint config, GLint share_list) {
// 	(*ptrglCreateContextWithConfigSGIX)(gc_id, screen, config, share_list);
// }
// void goglCreateGLXPixmapWithConfigSGIX(GLint config, GLint pixmap, GLint glxpixmap) {
// 	(*ptrglCreateGLXPixmapWithConfigSGIX)(config, pixmap, glxpixmap);
// }
// void goglCreateGLXPbufferSGIX(GLint config, GLint pbuffer) {
// 	(*ptrglCreateGLXPbufferSGIX)(config, pbuffer);
// }
// void goglDestroyGLXPbufferSGIX(GLint pbuffer) {
// 	(*ptrglDestroyGLXPbufferSGIX)(pbuffer);
// }
// void goglChangeDrawableAttributesSGIX(GLint drawable) {
// 	(*ptrglChangeDrawableAttributesSGIX)(drawable);
// }
// void goglGetDrawableAttributesSGIX(GLint drawable) {
// 	(*ptrglGetDrawableAttributesSGIX)(drawable);
// }
// void goglJoinSwapGroupSGIX(GLint window, GLint group) {
// 	(*ptrglJoinSwapGroupSGIX)(window, group);
// }
// void goglBindSwapBarrierSGIX(GLint window, GLint barrier) {
// 	(*ptrglBindSwapBarrierSGIX)(window, barrier);
// }
// void goglQueryMaxSwapBarriersSGIX() {
// 	(*ptrglQueryMaxSwapBarriersSGIX)();
// }
// GLXHyperpipeNetworkSGIX * goglQueryHyperpipeNetworkSGIX(Display* dpy, int* npipes) {
// 	return (*ptrglQueryHyperpipeNetworkSGIX)(dpy, npipes);
// }
// int goglHyperpipeConfigSGIX(Display* dpy, int networkId, int npipes, GLXHyperpipeConfigSGIX* cfg, int* hpId) {
// 	return (*ptrglHyperpipeConfigSGIX)(dpy, networkId, npipes, cfg, hpId);
// }
// GLXHyperpipeConfigSGIX * goglQueryHyperpipeConfigSGIX(Display* dpy, int hpId, int* npipes) {
// 	return (*ptrglQueryHyperpipeConfigSGIX)(dpy, hpId, npipes);
// }
// int goglDestroyHyperpipeConfigSGIX(Display* dpy, int hpId) {
// 	return (*ptrglDestroyHyperpipeConfigSGIX)(dpy, hpId);
// }
// int goglBindHyperpipeSGIX(Display* dpy, int hpId) {
// 	return (*ptrglBindHyperpipeSGIX)(dpy, hpId);
// }
// int goglQueryHyperpipeBestAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, GLvoid* attribList, GLvoid* returnAttribList) {
// 	return (*ptrglQueryHyperpipeBestAttribSGIX)(dpy, timeSlice, attrib, size, attribList, returnAttribList);
// }
// int goglHyperpipeAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, void* attribList) {
// 	return (*ptrglHyperpipeAttribSGIX)(dpy, timeSlice, attrib, size, attribList);
// }
// int goglQueryHyperpipeAttribSGIX(Display* dpy, int timeSlice, int attrib, int size, void* returnAttribList) {
// 	return (*ptrglQueryHyperpipeAttribSGIX)(dpy, timeSlice, attrib, size, returnAttribList);
// }
// 
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
// int init_ARB_fbconfig_float() {
// 	return 0;
// }
// int init_ARB_framebuffer_sRGB() {
// 	return 0;
// }
// int init_ARB_get_proc_address() {
// 	ptrglGetProcAddressARB = goglGetProcAddress("glGetProcAddressARB");
// 	if(ptrglGetProcAddressARB == NULL) return 1;
// 	return 0;
// }
// int init_ARB_multisample() {
// 	return 0;
// }
// int init_EXT_fbconfig_packed_float() {
// 	return 0;
// }
// int init_EXT_framebuffer_sRGB() {
// 	return 0;
// }
// int init_EXT_import_context() {
// 	ptrglGetCurrentDisplayEXT = goglGetProcAddress("glGetCurrentDisplayEXT");
// 	if(ptrglGetCurrentDisplayEXT == NULL) return 1;
// 	ptrglQueryContextInfoEXT = goglGetProcAddress("glQueryContextInfoEXT");
// 	if(ptrglQueryContextInfoEXT == NULL) return 1;
// 	ptrglGetContextIDEXT = goglGetProcAddress("glGetContextIDEXT");
// 	if(ptrglGetContextIDEXT == NULL) return 1;
// 	ptrglImportContextEXT = goglGetProcAddress("glImportContextEXT");
// 	if(ptrglImportContextEXT == NULL) return 1;
// 	ptrglFreeContextEXT = goglGetProcAddress("glFreeContextEXT");
// 	if(ptrglFreeContextEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_swap_control() {
// 	ptrglSwapIntervalEXT = goglGetProcAddress("glSwapIntervalEXT");
// 	if(ptrglSwapIntervalEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_swap_control_tear() {
// 	return 0;
// }
// int init_EXT_texture_from_pixmap() {
// 	ptrglBindTexImageEXT = goglGetProcAddress("glBindTexImageEXT");
// 	if(ptrglBindTexImageEXT == NULL) return 1;
// 	ptrglReleaseTexImageEXT = goglGetProcAddress("glReleaseTexImageEXT");
// 	if(ptrglReleaseTexImageEXT == NULL) return 1;
// 	return 0;
// }
// int init_EXT_visual_info() {
// 	return 0;
// }
// int init_EXT_visual_rating() {
// 	return 0;
// }
// int init_INTEL_swap_event() {
// 	return 0;
// }
// int init_MESA_agp_offset() {
// 	ptrglGetAGPOffsetMESA = goglGetProcAddress("glGetAGPOffsetMESA");
// 	if(ptrglGetAGPOffsetMESA == NULL) return 1;
// 	return 0;
// }
// int init_MESA_copy_sub_buffer() {
// 	ptrglCopySubBufferMESA = goglGetProcAddress("glCopySubBufferMESA");
// 	if(ptrglCopySubBufferMESA == NULL) return 1;
// 	return 0;
// }
// int init_MESA_pixmap_colormap() {
// 	ptrglCreateGLXPixmapMESA = goglGetProcAddress("glCreateGLXPixmapMESA");
// 	if(ptrglCreateGLXPixmapMESA == NULL) return 1;
// 	return 0;
// }
// int init_MESA_release_buffers() {
// 	ptrglReleaseBuffersMESA = goglGetProcAddress("glReleaseBuffersMESA");
// 	if(ptrglReleaseBuffersMESA == NULL) return 1;
// 	return 0;
// }
// int init_MESA_set_3dfx_mode() {
// 	ptrglSet3DfxModeMESA = goglGetProcAddress("glSet3DfxModeMESA");
// 	if(ptrglSet3DfxModeMESA == NULL) return 1;
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
// int init_NV_multisample_coverage() {
// 	return 0;
// }
// int init_NV_present_video() {
// 	ptrglEnumerateVideoDevicesNV = goglGetProcAddress("glEnumerateVideoDevicesNV");
// 	if(ptrglEnumerateVideoDevicesNV == NULL) return 1;
// 	ptrglBindVideoDeviceNV = goglGetProcAddress("glBindVideoDeviceNV");
// 	if(ptrglBindVideoDeviceNV == NULL) return 1;
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
// int init_OML_swap_method() {
// 	return 0;
// }
// int init_OML_sync_control() {
// 	ptrglGetSyncValuesOML = goglGetProcAddress("glGetSyncValuesOML");
// 	if(ptrglGetSyncValuesOML == NULL) return 1;
// 	ptrglGetMscRateOML = goglGetProcAddress("glGetMscRateOML");
// 	if(ptrglGetMscRateOML == NULL) return 1;
// 	ptrglSwapBuffersMscOML = goglGetProcAddress("glSwapBuffersMscOML");
// 	if(ptrglSwapBuffersMscOML == NULL) return 1;
// 	ptrglWaitForMscOML = goglGetProcAddress("glWaitForMscOML");
// 	if(ptrglWaitForMscOML == NULL) return 1;
// 	ptrglWaitForSbcOML = goglGetProcAddress("glWaitForSbcOML");
// 	if(ptrglWaitForSbcOML == NULL) return 1;
// 	return 0;
// }
// int init_SGIS_multisample() {
// 	return 0;
// }
// int init_SGIX_dmbuffer() {
// 	ptrglAssociateDMPbufferSGIX = goglGetProcAddress("glAssociateDMPbufferSGIX");
// 	if(ptrglAssociateDMPbufferSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_fbconfig() {
// 	ptrglGetFBConfigAttribSGIX = goglGetProcAddress("glGetFBConfigAttribSGIX");
// 	if(ptrglGetFBConfigAttribSGIX == NULL) return 1;
// 	ptrglChooseFBConfigSGIX = goglGetProcAddress("glChooseFBConfigSGIX");
// 	if(ptrglChooseFBConfigSGIX == NULL) return 1;
// 	ptrglCreateGLXPixmapWithConfigSGIX = goglGetProcAddress("glCreateGLXPixmapWithConfigSGIX");
// 	if(ptrglCreateGLXPixmapWithConfigSGIX == NULL) return 1;
// 	ptrglCreateContextWithConfigSGIX = goglGetProcAddress("glCreateContextWithConfigSGIX");
// 	if(ptrglCreateContextWithConfigSGIX == NULL) return 1;
// 	ptrglGetVisualFromFBConfigSGIX = goglGetProcAddress("glGetVisualFromFBConfigSGIX");
// 	if(ptrglGetVisualFromFBConfigSGIX == NULL) return 1;
// 	ptrglGetFBConfigFromVisualSGIX = goglGetProcAddress("glGetFBConfigFromVisualSGIX");
// 	if(ptrglGetFBConfigFromVisualSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_hyperpipe() {
// 	ptrglQueryHyperpipeNetworkSGIX = goglGetProcAddress("glQueryHyperpipeNetworkSGIX");
// 	if(ptrglQueryHyperpipeNetworkSGIX == NULL) return 1;
// 	ptrglHyperpipeConfigSGIX = goglGetProcAddress("glHyperpipeConfigSGIX");
// 	if(ptrglHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglQueryHyperpipeConfigSGIX = goglGetProcAddress("glQueryHyperpipeConfigSGIX");
// 	if(ptrglQueryHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglDestroyHyperpipeConfigSGIX = goglGetProcAddress("glDestroyHyperpipeConfigSGIX");
// 	if(ptrglDestroyHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglBindHyperpipeSGIX = goglGetProcAddress("glBindHyperpipeSGIX");
// 	if(ptrglBindHyperpipeSGIX == NULL) return 1;
// 	ptrglQueryHyperpipeBestAttribSGIX = goglGetProcAddress("glQueryHyperpipeBestAttribSGIX");
// 	if(ptrglQueryHyperpipeBestAttribSGIX == NULL) return 1;
// 	ptrglHyperpipeAttribSGIX = goglGetProcAddress("glHyperpipeAttribSGIX");
// 	if(ptrglHyperpipeAttribSGIX == NULL) return 1;
// 	ptrglQueryHyperpipeAttribSGIX = goglGetProcAddress("glQueryHyperpipeAttribSGIX");
// 	if(ptrglQueryHyperpipeAttribSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_pbuffer() {
// 	ptrglCreateGLXPbufferSGIX = goglGetProcAddress("glCreateGLXPbufferSGIX");
// 	if(ptrglCreateGLXPbufferSGIX == NULL) return 1;
// 	ptrglDestroyGLXPbufferSGIX = goglGetProcAddress("glDestroyGLXPbufferSGIX");
// 	if(ptrglDestroyGLXPbufferSGIX == NULL) return 1;
// 	ptrglQueryGLXPbufferSGIX = goglGetProcAddress("glQueryGLXPbufferSGIX");
// 	if(ptrglQueryGLXPbufferSGIX == NULL) return 1;
// 	ptrglSelectEventSGIX = goglGetProcAddress("glSelectEventSGIX");
// 	if(ptrglSelectEventSGIX == NULL) return 1;
// 	ptrglGetSelectedEventSGIX = goglGetProcAddress("glGetSelectedEventSGIX");
// 	if(ptrglGetSelectedEventSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_swap_barrier() {
// 	ptrglBindSwapBarrierSGIX = goglGetProcAddress("glBindSwapBarrierSGIX");
// 	if(ptrglBindSwapBarrierSGIX == NULL) return 1;
// 	ptrglQueryMaxSwapBarriersSGIX = goglGetProcAddress("glQueryMaxSwapBarriersSGIX");
// 	if(ptrglQueryMaxSwapBarriersSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_swap_group() {
// 	ptrglJoinSwapGroupSGIX = goglGetProcAddress("glJoinSwapGroupSGIX");
// 	if(ptrglJoinSwapGroupSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_video_resize() {
// 	ptrglBindChannelToWindowSGIX = goglGetProcAddress("glBindChannelToWindowSGIX");
// 	if(ptrglBindChannelToWindowSGIX == NULL) return 1;
// 	ptrglChannelRectSGIX = goglGetProcAddress("glChannelRectSGIX");
// 	if(ptrglChannelRectSGIX == NULL) return 1;
// 	ptrglQueryChannelRectSGIX = goglGetProcAddress("glQueryChannelRectSGIX");
// 	if(ptrglQueryChannelRectSGIX == NULL) return 1;
// 	ptrglQueryChannelDeltasSGIX = goglGetProcAddress("glQueryChannelDeltasSGIX");
// 	if(ptrglQueryChannelDeltasSGIX == NULL) return 1;
// 	ptrglChannelRectSyncSGIX = goglGetProcAddress("glChannelRectSyncSGIX");
// 	if(ptrglChannelRectSyncSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_video_source() {
// 	ptrglCreateGLXVideoSourceSGIX = goglGetProcAddress("glCreateGLXVideoSourceSGIX");
// 	if(ptrglCreateGLXVideoSourceSGIX == NULL) return 1;
// 	ptrglDestroyGLXVideoSourceSGIX = goglGetProcAddress("glDestroyGLXVideoSourceSGIX");
// 	if(ptrglDestroyGLXVideoSourceSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_visual_select_group() {
// 	return 0;
// }
// int init_SGI_cushion() {
// 	ptrglCushionSGI = goglGetProcAddress("glCushionSGI");
// 	if(ptrglCushionSGI == NULL) return 1;
// 	return 0;
// }
// int init_SGI_make_current_read() {
// 	ptrglMakeCurrentReadSGI = goglGetProcAddress("glMakeCurrentReadSGI");
// 	if(ptrglMakeCurrentReadSGI == NULL) return 1;
// 	ptrglGetCurrentReadDrawableSGI = goglGetProcAddress("glGetCurrentReadDrawableSGI");
// 	if(ptrglGetCurrentReadDrawableSGI == NULL) return 1;
// 	return 0;
// }
// int init_SGI_swap_control() {
// 	ptrglSwapIntervalSGI = goglGetProcAddress("glSwapIntervalSGI");
// 	if(ptrglSwapIntervalSGI == NULL) return 1;
// 	return 0;
// }
// int init_SGI_video_sync() {
// 	ptrglGetVideoSyncSGI = goglGetProcAddress("glGetVideoSyncSGI");
// 	if(ptrglGetVideoSyncSGI == NULL) return 1;
// 	ptrglWaitVideoSyncSGI = goglGetProcAddress("glWaitVideoSyncSGI");
// 	if(ptrglWaitVideoSyncSGI == NULL) return 1;
// 	return 0;
// }
// int init_SUN_get_transparent_index() {
// 	ptrglGetTransparentIndexSUN = goglGetProcAddress("glGetTransparentIndexSUN");
// 	if(ptrglGetTransparentIndexSUN == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_3() {
// 	ptrglGetFBConfigs = goglGetProcAddress("glGetFBConfigs");
// 	if(ptrglGetFBConfigs == NULL) return 1;
// 	ptrglChooseFBConfig = goglGetProcAddress("glChooseFBConfig");
// 	if(ptrglChooseFBConfig == NULL) return 1;
// 	ptrglGetFBConfigAttrib = goglGetProcAddress("glGetFBConfigAttrib");
// 	if(ptrglGetFBConfigAttrib == NULL) return 1;
// 	ptrglGetVisualFromFBConfig = goglGetProcAddress("glGetVisualFromFBConfig");
// 	if(ptrglGetVisualFromFBConfig == NULL) return 1;
// 	ptrglCreateWindow = goglGetProcAddress("glCreateWindow");
// 	if(ptrglCreateWindow == NULL) return 1;
// 	ptrglDestroyWindow = goglGetProcAddress("glDestroyWindow");
// 	if(ptrglDestroyWindow == NULL) return 1;
// 	ptrglCreatePixmap = goglGetProcAddress("glCreatePixmap");
// 	if(ptrglCreatePixmap == NULL) return 1;
// 	ptrglDestroyPixmap = goglGetProcAddress("glDestroyPixmap");
// 	if(ptrglDestroyPixmap == NULL) return 1;
// 	ptrglCreatePbuffer = goglGetProcAddress("glCreatePbuffer");
// 	if(ptrglCreatePbuffer == NULL) return 1;
// 	ptrglDestroyPbuffer = goglGetProcAddress("glDestroyPbuffer");
// 	if(ptrglDestroyPbuffer == NULL) return 1;
// 	ptrglQueryDrawable = goglGetProcAddress("glQueryDrawable");
// 	if(ptrglQueryDrawable == NULL) return 1;
// 	ptrglCreateNewContext = goglGetProcAddress("glCreateNewContext");
// 	if(ptrglCreateNewContext == NULL) return 1;
// 	ptrglMakeContextCurrent = goglGetProcAddress("glMakeContextCurrent");
// 	if(ptrglMakeContextCurrent == NULL) return 1;
// 	ptrglGetCurrentReadDrawable = goglGetProcAddress("glGetCurrentReadDrawable");
// 	if(ptrglGetCurrentReadDrawable == NULL) return 1;
// 	ptrglGetCurrentDisplay = goglGetProcAddress("glGetCurrentDisplay");
// 	if(ptrglGetCurrentDisplay == NULL) return 1;
// 	ptrglQueryContext = goglGetProcAddress("glQueryContext");
// 	if(ptrglQueryContext == NULL) return 1;
// 	ptrglSelectEvent = goglGetProcAddress("glSelectEvent");
// 	if(ptrglSelectEvent == NULL) return 1;
// 	ptrglGetSelectedEvent = goglGetProcAddress("glGetSelectedEvent");
// 	if(ptrglGetSelectedEvent == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_4() {
// 	ptrglGetProcAddress = goglGetProcAddress("glGetProcAddress");
// 	if(ptrglGetProcAddress == NULL) return 1;
// 	return 0;
// }
// int init_glx() {
// 	ptrglRender = goglGetProcAddress("glRender");
// 	if(ptrglRender == NULL) return 1;
// 	ptrglRenderLarge = goglGetProcAddress("glRenderLarge");
// 	if(ptrglRenderLarge == NULL) return 1;
// 	ptrglCreateContext = goglGetProcAddress("glCreateContext");
// 	if(ptrglCreateContext == NULL) return 1;
// 	ptrglDestroyContext = goglGetProcAddress("glDestroyContext");
// 	if(ptrglDestroyContext == NULL) return 1;
// 	ptrglMakeCurrent = goglGetProcAddress("glMakeCurrent");
// 	if(ptrglMakeCurrent == NULL) return 1;
// 	ptrglIsDirect = goglGetProcAddress("glIsDirect");
// 	if(ptrglIsDirect == NULL) return 1;
// 	ptrglQueryVersion = goglGetProcAddress("glQueryVersion");
// 	if(ptrglQueryVersion == NULL) return 1;
// 	ptrglWaitGL = goglGetProcAddress("glWaitGL");
// 	if(ptrglWaitGL == NULL) return 1;
// 	ptrglWaitX = goglGetProcAddress("glWaitX");
// 	if(ptrglWaitX == NULL) return 1;
// 	ptrglCopyContext = goglGetProcAddress("glCopyContext");
// 	if(ptrglCopyContext == NULL) return 1;
// 	ptrglSwapBuffers = goglGetProcAddress("glSwapBuffers");
// 	if(ptrglSwapBuffers == NULL) return 1;
// 	ptrglUseXFont = goglGetProcAddress("glUseXFont");
// 	if(ptrglUseXFont == NULL) return 1;
// 	ptrglCreateGLXPixmap = goglGetProcAddress("glCreateGLXPixmap");
// 	if(ptrglCreateGLXPixmap == NULL) return 1;
// 	ptrglGetVisualConfigs = goglGetProcAddress("glGetVisualConfigs");
// 	if(ptrglGetVisualConfigs == NULL) return 1;
// 	ptrglDestroyGLXPixmap = goglGetProcAddress("glDestroyGLXPixmap");
// 	if(ptrglDestroyGLXPixmap == NULL) return 1;
// 	ptrglVendorPrivate = goglGetProcAddress("glVendorPrivate");
// 	if(ptrglVendorPrivate == NULL) return 1;
// 	ptrglVendorPrivateWithReply = goglGetProcAddress("glVendorPrivateWithReply");
// 	if(ptrglVendorPrivateWithReply == NULL) return 1;
// 	ptrglQueryExtensionsString = goglGetProcAddress("glQueryExtensionsString");
// 	if(ptrglQueryExtensionsString == NULL) return 1;
// 	ptrglQueryServerString = goglGetProcAddress("glQueryServerString");
// 	if(ptrglQueryServerString == NULL) return 1;
// 	ptrglClientInfo = goglGetProcAddress("glClientInfo");
// 	if(ptrglClientInfo == NULL) return 1;
// 	ptrglGetFBConfigs = goglGetProcAddress("glGetFBConfigs");
// 	if(ptrglGetFBConfigs == NULL) return 1;
// 	ptrglCreatePixmap = goglGetProcAddress("glCreatePixmap");
// 	if(ptrglCreatePixmap == NULL) return 1;
// 	ptrglDestroyPixmap = goglGetProcAddress("glDestroyPixmap");
// 	if(ptrglDestroyPixmap == NULL) return 1;
// 	ptrglCreateNewContext = goglGetProcAddress("glCreateNewContext");
// 	if(ptrglCreateNewContext == NULL) return 1;
// 	ptrglQueryContext = goglGetProcAddress("glQueryContext");
// 	if(ptrglQueryContext == NULL) return 1;
// 	ptrglMakeContextCurrent = goglGetProcAddress("glMakeContextCurrent");
// 	if(ptrglMakeContextCurrent == NULL) return 1;
// 	ptrglCreatePbuffer = goglGetProcAddress("glCreatePbuffer");
// 	if(ptrglCreatePbuffer == NULL) return 1;
// 	ptrglDestroyPbuffer = goglGetProcAddress("glDestroyPbuffer");
// 	if(ptrglDestroyPbuffer == NULL) return 1;
// 	ptrglGetDrawableAttributes = goglGetProcAddress("glGetDrawableAttributes");
// 	if(ptrglGetDrawableAttributes == NULL) return 1;
// 	ptrglChangeDrawableAttributes = goglGetProcAddress("glChangeDrawableAttributes");
// 	if(ptrglChangeDrawableAttributes == NULL) return 1;
// 	ptrglCreateWindow = goglGetProcAddress("glCreateWindow");
// 	if(ptrglCreateWindow == NULL) return 1;
// 	ptrglDestroyWindow = goglGetProcAddress("glDestroyWindow");
// 	if(ptrglDestroyWindow == NULL) return 1;
// 	ptrglSwapIntervalSGI = goglGetProcAddress("glSwapIntervalSGI");
// 	if(ptrglSwapIntervalSGI == NULL) return 1;
// 	ptrglMakeCurrentReadSGI = goglGetProcAddress("glMakeCurrentReadSGI");
// 	if(ptrglMakeCurrentReadSGI == NULL) return 1;
// 	ptrglCreateGLXVideoSourceSGIX = goglGetProcAddress("glCreateGLXVideoSourceSGIX");
// 	if(ptrglCreateGLXVideoSourceSGIX == NULL) return 1;
// 	ptrglDestroyGLXVideoSourceSGIX = goglGetProcAddress("glDestroyGLXVideoSourceSGIX");
// 	if(ptrglDestroyGLXVideoSourceSGIX == NULL) return 1;
// 	ptrglQueryContextInfoEXT = goglGetProcAddress("glQueryContextInfoEXT");
// 	if(ptrglQueryContextInfoEXT == NULL) return 1;
// 	ptrglGetFBConfigsSGIX = goglGetProcAddress("glGetFBConfigsSGIX");
// 	if(ptrglGetFBConfigsSGIX == NULL) return 1;
// 	ptrglCreateContextWithConfigSGIX = goglGetProcAddress("glCreateContextWithConfigSGIX");
// 	if(ptrglCreateContextWithConfigSGIX == NULL) return 1;
// 	ptrglCreateGLXPixmapWithConfigSGIX = goglGetProcAddress("glCreateGLXPixmapWithConfigSGIX");
// 	if(ptrglCreateGLXPixmapWithConfigSGIX == NULL) return 1;
// 	ptrglCreateGLXPbufferSGIX = goglGetProcAddress("glCreateGLXPbufferSGIX");
// 	if(ptrglCreateGLXPbufferSGIX == NULL) return 1;
// 	ptrglDestroyGLXPbufferSGIX = goglGetProcAddress("glDestroyGLXPbufferSGIX");
// 	if(ptrglDestroyGLXPbufferSGIX == NULL) return 1;
// 	ptrglChangeDrawableAttributesSGIX = goglGetProcAddress("glChangeDrawableAttributesSGIX");
// 	if(ptrglChangeDrawableAttributesSGIX == NULL) return 1;
// 	ptrglGetDrawableAttributesSGIX = goglGetProcAddress("glGetDrawableAttributesSGIX");
// 	if(ptrglGetDrawableAttributesSGIX == NULL) return 1;
// 	ptrglJoinSwapGroupSGIX = goglGetProcAddress("glJoinSwapGroupSGIX");
// 	if(ptrglJoinSwapGroupSGIX == NULL) return 1;
// 	ptrglBindSwapBarrierSGIX = goglGetProcAddress("glBindSwapBarrierSGIX");
// 	if(ptrglBindSwapBarrierSGIX == NULL) return 1;
// 	ptrglQueryMaxSwapBarriersSGIX = goglGetProcAddress("glQueryMaxSwapBarriersSGIX");
// 	if(ptrglQueryMaxSwapBarriersSGIX == NULL) return 1;
// 	ptrglQueryHyperpipeNetworkSGIX = goglGetProcAddress("glQueryHyperpipeNetworkSGIX");
// 	if(ptrglQueryHyperpipeNetworkSGIX == NULL) return 1;
// 	ptrglHyperpipeConfigSGIX = goglGetProcAddress("glHyperpipeConfigSGIX");
// 	if(ptrglHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglQueryHyperpipeConfigSGIX = goglGetProcAddress("glQueryHyperpipeConfigSGIX");
// 	if(ptrglQueryHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglDestroyHyperpipeConfigSGIX = goglGetProcAddress("glDestroyHyperpipeConfigSGIX");
// 	if(ptrglDestroyHyperpipeConfigSGIX == NULL) return 1;
// 	ptrglBindHyperpipeSGIX = goglGetProcAddress("glBindHyperpipeSGIX");
// 	if(ptrglBindHyperpipeSGIX == NULL) return 1;
// 	ptrglQueryHyperpipeBestAttribSGIX = goglGetProcAddress("glQueryHyperpipeBestAttribSGIX");
// 	if(ptrglQueryHyperpipeBestAttribSGIX == NULL) return 1;
// 	ptrglHyperpipeAttribSGIX = goglGetProcAddress("glHyperpipeAttribSGIX");
// 	if(ptrglHyperpipeAttribSGIX == NULL) return 1;
// 	ptrglQueryHyperpipeAttribSGIX = goglGetProcAddress("glQueryHyperpipeAttribSGIX");
// 	if(ptrglQueryHyperpipeAttribSGIX == NULL) return 1;
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

// http://www.opengl.org/sdk/docs/man/xhtml/glCreateContextAttribsARB.xml
func CreateContextAttribsARB(dpy Pointer, config Pointer, share_context Pointer, direct int, attrib_list *Int) Pointer {
	return (Pointer)(C.goglCreateContextAttribsARB((*C.Display)(dpy), (C.GLXFBConfig)(config), (C.GLXContext)(share_context), (C.int)(direct), (*C.int)(attrib_list)))
}
// ARB_create_context_profile

// ARB_create_context_robustness

// ARB_fbconfig_float

// ARB_framebuffer_sRGB

// ARB_get_proc_address

// http://www.opengl.org/sdk/docs/man/xhtml/glGetProcAddressARB.xml
func GetProcAddressARB(procName *Ubyte) Pointer {
	return (Pointer)(C.goglGetProcAddressARB((*C.GLubyte)(procName)))
}
// ARB_multisample

// EXT_fbconfig_packed_float

// EXT_framebuffer_sRGB

// EXT_import_context

// http://www.opengl.org/sdk/docs/man/xhtml/glGetCurrentDisplayEXT.xml
func GetCurrentDisplayEXT() Pointer {
	return (Pointer)(C.goglGetCurrentDisplayEXT())
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryContextInfoEXT.xml
func QueryContextInfoEXT(dpy Pointer, context Pointer, attribute Int, value *Int) Int {
	return (Int)(C.goglQueryContextInfoEXT((*C.Display)(dpy), (C.GLXContext)(context), (C.int)(attribute), (*C.int)(value)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetContextIDEXT.xml
func GetContextIDEXT(context Pointer) uint32 {
	return (uint32)(C.goglGetContextIDEXT((C.GLXContext)(context)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glImportContextEXT.xml
func ImportContextEXT(dpy Pointer, contextID uint32) Pointer {
	return (Pointer)(C.goglImportContextEXT((*C.Display)(dpy), (C.GLXContextID)(contextID)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFreeContextEXT.xml
func FreeContextEXT(dpy Pointer, context Pointer)  {
	C.goglFreeContextEXT((*C.Display)(dpy), (C.GLXContext)(context))
}
// EXT_swap_control

// http://www.opengl.org/sdk/docs/man/xhtml/glSwapIntervalEXT.xml
func SwapIntervalEXT(dpy Pointer, drawable Pointer, interval Int)  {
	C.goglSwapIntervalEXT((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int)(interval))
}
// EXT_swap_control_tear

// EXT_texture_from_pixmap

// http://www.opengl.org/sdk/docs/man/xhtml/glBindTexImageEXT.xml
func BindTexImageEXT(dpy Pointer, drawable Pointer, buffer Int, attrib_list *Int)  {
	C.goglBindTexImageEXT((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int)(buffer), (*C.int)(attrib_list))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glReleaseTexImageEXT.xml
func ReleaseTexImageEXT(dpy Pointer, drawable Pointer, buffer Int)  {
	C.goglReleaseTexImageEXT((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int)(buffer))
}
// EXT_visual_info

// EXT_visual_rating

// INTEL_swap_event

// MESA_agp_offset

// http://www.opengl.org/sdk/docs/man/xhtml/glGetAGPOffsetMESA.xml
func GetAGPOffsetMESA(pointer Pointer) uint32 {
	return (uint32)(C.goglGetAGPOffsetMESA((unsafe.Pointer)(pointer)))
}
// MESA_copy_sub_buffer

// http://www.opengl.org/sdk/docs/man/xhtml/glCopySubBufferMESA.xml
func CopySubBufferMESA(dpy Pointer, drawable Pointer, x Int, y Int, width Int, height Int)  {
	C.goglCopySubBufferMESA((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int)(x), (C.int)(y), (C.int)(width), (C.int)(height))
}
// MESA_pixmap_colormap

// http://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPixmapMESA.xml
func CreateGLXPixmapMESA(dpy Pointer, visual Pointer, pixmap Pointer, cmap Pointer) Pointer {
	return (Pointer)(C.goglCreateGLXPixmapMESA((*C.Display)(dpy), (C.XVisualInfo)(visual), (C.Pixmap)(pixmap), (C.Colormap)(cmap)))
}
// MESA_release_buffers

// http://www.opengl.org/sdk/docs/man/xhtml/glReleaseBuffersMESA.xml
func ReleaseBuffersMESA(dpy Pointer, drawable Pointer) int {
	return (int)(C.goglReleaseBuffersMESA((*C.Display)(dpy), (C.GLXDrawable)(drawable)))
}
// MESA_set_3dfx_mode

// http://www.opengl.org/sdk/docs/man/xhtml/glSet3DfxModeMESA.xml
func Set3DfxModeMESA(mode Int) int {
	return (int)(C.goglSet3DfxModeMESA((C.int)(mode)))
}
// NV_copy_image

// http://www.opengl.org/sdk/docs/man/xhtml/glCopyImageSubDataNV.xml
func CopyImageSubDataNV(dpy Pointer, srcCtx Pointer, srcName Uint, srcTarget Enum, srcLevel Int, srcX Int, srcY Int, srcZ Int, dstCtx Pointer, dstName Uint, dstTarget Enum, dstLevel Int, dstX Int, dstY Int, dstZ Int, width Sizei, height Sizei, depth Sizei)  {
	C.goglCopyImageSubDataNV((*C.Display)(dpy), (C.GLXContext)(srcCtx), (C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.GLXContext)(dstCtx), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
// NV_float_buffer

// NV_multisample_coverage

// NV_present_video

// http://www.opengl.org/sdk/docs/man/xhtml/glEnumerateVideoDevicesNV.xml
func EnumerateVideoDevicesNV(dpy Pointer, screen Int, nelements *Int) *uint32 {
	return (*uint32)(C.goglEnumerateVideoDevicesNV((*C.Display)(dpy), (C.int)(screen), (*C.int)(nelements)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBindVideoDeviceNV.xml
func BindVideoDeviceNV(dpy Pointer, video_slot uint32, video_device uint32, attrib_list *Int) Int {
	return (Int)(C.goglBindVideoDeviceNV((*C.Display)(dpy), (C.uint32)(video_slot), (C.uint32)(video_device), (*C.int)(attrib_list)))
}
// NV_swap_group

// http://www.opengl.org/sdk/docs/man/xhtml/glJoinSwapGroupNV.xml
func JoinSwapGroupNV(dpy Pointer, drawable Pointer, group Uint) int {
	return (int)(C.goglJoinSwapGroupNV((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.GLuint)(group)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBindSwapBarrierNV.xml
func BindSwapBarrierNV(dpy Pointer, group Uint, barrier Uint) int {
	return (int)(C.goglBindSwapBarrierNV((*C.Display)(dpy), (C.GLuint)(group), (C.GLuint)(barrier)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQuerySwapGroupNV.xml
func QuerySwapGroupNV(dpy Pointer, drawable Pointer, group *Uint, barrier *Uint) int {
	return (int)(C.goglQuerySwapGroupNV((*C.Display)(dpy), (C.GLXDrawable)(drawable), (*C.GLuint)(group), (*C.GLuint)(barrier)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryMaxSwapGroupsNV.xml
func QueryMaxSwapGroupsNV(dpy Pointer, screen Int, maxGroups *Uint, maxBarriers *Uint) int {
	return (int)(C.goglQueryMaxSwapGroupsNV((*C.Display)(dpy), (C.int)(screen), (*C.GLuint)(maxGroups), (*C.GLuint)(maxBarriers)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryFrameCountNV.xml
func QueryFrameCountNV(dpy Pointer, screen Int, count *Uint) int {
	return (int)(C.goglQueryFrameCountNV((*C.Display)(dpy), (C.int)(screen), (*C.GLuint)(count)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glResetFrameCountNV.xml
func ResetFrameCountNV(dpy Pointer, screen Int) int {
	return (int)(C.goglResetFrameCountNV((*C.Display)(dpy), (C.int)(screen)))
}
// NV_video_capture

// http://www.opengl.org/sdk/docs/man/xhtml/glBindVideoCaptureDeviceNV.xml
func BindVideoCaptureDeviceNV(dpy Pointer, video_capture_slot uint32, device Pointer) Int {
	return (Int)(C.goglBindVideoCaptureDeviceNV((*C.Display)(dpy), (C.uint32)(video_capture_slot), (C.GLXVideoCaptureDeviceNV)(device)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEnumerateVideoCaptureDevicesNV.xml
func EnumerateVideoCaptureDevicesNV(dpy Pointer, screen Int, nelements *Int) Pointer {
	return (Pointer)(C.goglEnumerateVideoCaptureDevicesNV((*C.Display)(dpy), (C.int)(screen), (*C.int)(nelements)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLockVideoCaptureDeviceNV.xml
func LockVideoCaptureDeviceNV(dpy Pointer, device Pointer)  {
	C.goglLockVideoCaptureDeviceNV((*C.Display)(dpy), (C.GLXVideoCaptureDeviceNV)(device))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryVideoCaptureDeviceNV.xml
func QueryVideoCaptureDeviceNV(dpy Pointer, device Pointer, attribute Int, value *Int) Int {
	return (Int)(C.goglQueryVideoCaptureDeviceNV((*C.Display)(dpy), (C.GLXVideoCaptureDeviceNV)(device), (C.int)(attribute), (*C.int)(value)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glReleaseVideoCaptureDeviceNV.xml
func ReleaseVideoCaptureDeviceNV(dpy Pointer, device Pointer)  {
	C.goglReleaseVideoCaptureDeviceNV((*C.Display)(dpy), (C.GLXVideoCaptureDeviceNV)(device))
}
// NV_video_output

// http://www.opengl.org/sdk/docs/man/xhtml/glGetVideoDeviceNV.xml
func GetVideoDeviceNV(dpy Pointer, screen Int, numVideoDevices Int, pVideoDevice Pointer) Int {
	return (Int)(C.goglGetVideoDeviceNV((*C.Display)(dpy), (C.int)(screen), (C.int)(numVideoDevices), (C.GLXVideoDeviceNV)(pVideoDevice)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glReleaseVideoDeviceNV.xml
func ReleaseVideoDeviceNV(dpy Pointer, screen Int, VideoDevice Pointer) Int {
	return (Int)(C.goglReleaseVideoDeviceNV((*C.Display)(dpy), (C.int)(screen), (C.GLXVideoDeviceNV)(VideoDevice)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBindVideoImageNV.xml
func BindVideoImageNV(dpy Pointer, VideoDevice Pointer, pbuf Pointer, iVideoBuffer Int) Int {
	return (Int)(C.goglBindVideoImageNV((*C.Display)(dpy), (C.GLXVideoDeviceNV)(VideoDevice), (C.GLXPbuffer)(pbuf), (C.int)(iVideoBuffer)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glReleaseVideoImageNV.xml
func ReleaseVideoImageNV(dpy Pointer, pbuf Pointer) Int {
	return (Int)(C.goglReleaseVideoImageNV((*C.Display)(dpy), (C.GLXPbuffer)(pbuf)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSendPbufferToVideoNV.xml
func SendPbufferToVideoNV(dpy Pointer, pbuf Pointer, iBufferType Int, pulCounterPbuffer *uint32, bBlock Boolean) Int {
	return (Int)(C.goglSendPbufferToVideoNV((*C.Display)(dpy), (C.GLXPbuffer)(pbuf), (C.int)(iBufferType), (*C.unsigned_long)(pulCounterPbuffer), (C.GLboolean)(bBlock)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetVideoInfoNV.xml
func GetVideoInfoNV(dpy Pointer, screen Int, VideoDevice Pointer, pulCounterOutputPbuffer *uint32, pulCounterOutputVideo *uint32) Int {
	return (Int)(C.goglGetVideoInfoNV((*C.Display)(dpy), (C.int)(screen), (C.GLXVideoDeviceNV)(VideoDevice), (*C.unsigned_long)(pulCounterOutputPbuffer), (*C.unsigned_long)(pulCounterOutputVideo)))
}
// OML_swap_method

// OML_sync_control

// http://www.opengl.org/sdk/docs/man/xhtml/glGetSyncValuesOML.xml
func GetSyncValuesOML(dpy Pointer, drawable Pointer, ust *int64, msc *int64, sbc *int64) int {
	return (int)(C.goglGetSyncValuesOML((*C.Display)(dpy), (C.GLXDrawable)(drawable), (*C.int64_t)(ust), (*C.int64_t)(msc), (*C.int64_t)(sbc)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMscRateOML.xml
func GetMscRateOML(dpy Pointer, drawable Pointer, numerator *int32, denominator *int32) int {
	return (int)(C.goglGetMscRateOML((*C.Display)(dpy), (C.GLXDrawable)(drawable), (*C.int32_t)(numerator), (*C.int32_t)(denominator)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSwapBuffersMscOML.xml
func SwapBuffersMscOML(dpy Pointer, drawable Pointer, target_msc int64, divisor int64, remainder int64) int64 {
	return (int64)(C.goglSwapBuffersMscOML((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int64_t)(target_msc), (C.int64_t)(divisor), (C.int64_t)(remainder)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWaitForMscOML.xml
func WaitForMscOML(dpy Pointer, drawable Pointer, target_msc int64, divisor int64, remainder int64, ust *int64, msc *int64, sbc *int64) int {
	return (int)(C.goglWaitForMscOML((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int64_t)(target_msc), (C.int64_t)(divisor), (C.int64_t)(remainder), (*C.int64_t)(ust), (*C.int64_t)(msc), (*C.int64_t)(sbc)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWaitForSbcOML.xml
func WaitForSbcOML(dpy Pointer, drawable Pointer, target_sbc int64, ust *int64, msc *int64, sbc *int64) int {
	return (int)(C.goglWaitForSbcOML((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int64_t)(target_sbc), (*C.int64_t)(ust), (*C.int64_t)(msc), (*C.int64_t)(sbc)))
}
// SGIS_multisample

// SGIX_dmbuffer

// http://www.opengl.org/sdk/docs/man/xhtml/glAssociateDMPbufferSGIX.xml
func AssociateDMPbufferSGIX(dpy Pointer, pbuffer Pointer, params Pointer, dmbuffer Pointer) int {
	return (int)(C.goglAssociateDMPbufferSGIX((*C.Display)(dpy), (C.GLXPbufferSGIX)(pbuffer), (C.DMparams)(params), (C.DMbuffer)(dmbuffer)))
}
// SGIX_fbconfig

// http://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigAttribSGIX.xml
func GetFBConfigAttribSGIX(dpy Pointer, config Pointer, attribute Int, value *Int) Int {
	return (Int)(C.goglGetFBConfigAttribSGIX((*C.Display)(dpy), (C.GLXFBConfigSGIX)(config), (C.int)(attribute), (*C.int)(value)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glChooseFBConfigSGIX.xml
func ChooseFBConfigSGIX(dpy Pointer, screen Int, attrib_list *Int, nelements *Int) Pointer {
	return (Pointer)(C.goglChooseFBConfigSGIX((*C.Display)(dpy), (C.int)(screen), (*C.int)(attrib_list), (*C.int)(nelements)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPixmapWithConfigSGIX.xml
func CreateGLXPixmapWithConfigSGIX(dpy Pointer, config Pointer, pixmap Pointer) Pointer {
	return (Pointer)(C.goglCreateGLXPixmapWithConfigSGIX((*C.Display)(dpy), (C.GLXFBConfigSGIX)(config), (C.Pixmap)(pixmap)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateContextWithConfigSGIX.xml
func CreateContextWithConfigSGIX(dpy Pointer, config Pointer, render_type Int, share_list Pointer, direct int) Pointer {
	return (Pointer)(C.goglCreateContextWithConfigSGIX((*C.Display)(dpy), (C.GLXFBConfigSGIX)(config), (C.int)(render_type), (C.GLXContext)(share_list), (C.int)(direct)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetVisualFromFBConfigSGIX.xml
func GetVisualFromFBConfigSGIX(dpy Pointer, config Pointer) Pointer {
	return (Pointer)(C.goglGetVisualFromFBConfigSGIX((*C.Display)(dpy), (C.GLXFBConfigSGIX)(config)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigFromVisualSGIX.xml
func GetFBConfigFromVisualSGIX(dpy Pointer, vis Pointer) Pointer {
	return (Pointer)(C.goglGetFBConfigFromVisualSGIX((*C.Display)(dpy), (C.XVisualInfo)(vis)))
}
// SGIX_hyperpipe

// http://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeNetworkSGIX.xml
func QueryHyperpipeNetworkSGIX(dpy Pointer, npipes *Int) Pointer {
	return (Pointer)(C.goglQueryHyperpipeNetworkSGIX((*C.Display)(dpy), (*C.int)(npipes)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glHyperpipeConfigSGIX.xml
func HyperpipeConfigSGIX(dpy Pointer, networkId Int, npipes Int, cfg Pointer, hpId *Int) Int {
	return (Int)(C.goglHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(networkId), (C.int)(npipes), (*C.GLXHyperpipeConfigSGIX)(cfg), (*C.int)(hpId)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeConfigSGIX.xml
func QueryHyperpipeConfigSGIX(dpy Pointer, hpId Int, npipes *Int) Pointer {
	return (Pointer)(C.goglQueryHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(hpId), (*C.int)(npipes)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyHyperpipeConfigSGIX.xml
func DestroyHyperpipeConfigSGIX(dpy Pointer, hpId Int) Int {
	return (Int)(C.goglDestroyHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(hpId)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBindHyperpipeSGIX.xml
func BindHyperpipeSGIX(dpy Pointer, hpId Int) Int {
	return (Int)(C.goglBindHyperpipeSGIX((*C.Display)(dpy), (C.int)(hpId)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeBestAttribSGIX.xml
func QueryHyperpipeBestAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, attribList Pointer, returnAttribList Pointer) Int {
	return (Int)(C.goglQueryHyperpipeBestAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(attribList), (unsafe.Pointer)(returnAttribList)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glHyperpipeAttribSGIX.xml
func HyperpipeAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, attribList Pointer) Int {
	return (Int)(C.goglHyperpipeAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(attribList)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeAttribSGIX.xml
func QueryHyperpipeAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, returnAttribList Pointer) Int {
	return (Int)(C.goglQueryHyperpipeAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(returnAttribList)))
}
// SGIX_pbuffer

// http://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPbufferSGIX.xml
func CreateGLXPbufferSGIX(dpy Pointer, config Pointer, width uint32, height uint32, attrib_list *Int) Pointer {
	return (Pointer)(C.goglCreateGLXPbufferSGIX((*C.Display)(dpy), (C.GLXFBConfigSGIX)(config), (C.uint32)(width), (C.uint32)(height), (*C.int)(attrib_list)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyGLXPbufferSGIX.xml
func DestroyGLXPbufferSGIX(dpy Pointer, pbuf Pointer)  {
	C.goglDestroyGLXPbufferSGIX((*C.Display)(dpy), (C.GLXPbufferSGIX)(pbuf))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryGLXPbufferSGIX.xml
func QueryGLXPbufferSGIX(dpy Pointer, pbuf Pointer, attribute Int, value *uint32) Int {
	return (Int)(C.goglQueryGLXPbufferSGIX((*C.Display)(dpy), (C.GLXPbufferSGIX)(pbuf), (C.int)(attribute), (*C.uint32)(value)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSelectEventSGIX.xml
func SelectEventSGIX(dpy Pointer, drawable Pointer, mask uint32)  {
	C.goglSelectEventSGIX((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.unsigned_long)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetSelectedEventSGIX.xml
func GetSelectedEventSGIX(dpy Pointer, drawable Pointer, mask *uint32)  {
	C.goglGetSelectedEventSGIX((*C.Display)(dpy), (C.GLXDrawable)(drawable), (*C.unsigned_long)(mask))
}
// SGIX_swap_barrier

// http://www.opengl.org/sdk/docs/man/xhtml/glBindSwapBarrierSGIX.xml
func BindSwapBarrierSGIX(dpy Pointer, drawable Pointer, barrier Int)  {
	C.goglBindSwapBarrierSGIX((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.int)(barrier))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryMaxSwapBarriersSGIX.xml
func QueryMaxSwapBarriersSGIX(dpy Pointer, screen Int, max *Int) int {
	return (int)(C.goglQueryMaxSwapBarriersSGIX((*C.Display)(dpy), (C.int)(screen), (*C.int)(max)))
}
// SGIX_swap_group

// http://www.opengl.org/sdk/docs/man/xhtml/glJoinSwapGroupSGIX.xml
func JoinSwapGroupSGIX(dpy Pointer, drawable Pointer, member Pointer)  {
	C.goglJoinSwapGroupSGIX((*C.Display)(dpy), (C.GLXDrawable)(drawable), (C.GLXDrawable)(member))
}
// SGIX_video_resize

// http://www.opengl.org/sdk/docs/man/xhtml/glBindChannelToWindowSGIX.xml
func BindChannelToWindowSGIX(display Pointer, screen Int, channel Int, window Pointer) Int {
	return (Int)(C.goglBindChannelToWindowSGIX((*C.Display)(display), (C.int)(screen), (C.int)(channel), (C.Window)(window)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glChannelRectSGIX.xml
func ChannelRectSGIX(display Pointer, screen Int, channel Int, x Int, y Int, w Int, h Int) Int {
	return (Int)(C.goglChannelRectSGIX((*C.Display)(display), (C.int)(screen), (C.int)(channel), (C.int)(x), (C.int)(y), (C.int)(w), (C.int)(h)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryChannelRectSGIX.xml
func QueryChannelRectSGIX(display Pointer, screen Int, channel Int, dx *Int, dy *Int, dw *Int, dh *Int) Int {
	return (Int)(C.goglQueryChannelRectSGIX((*C.Display)(display), (C.int)(screen), (C.int)(channel), (*C.int)(dx), (*C.int)(dy), (*C.int)(dw), (*C.int)(dh)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryChannelDeltasSGIX.xml
func QueryChannelDeltasSGIX(display Pointer, screen Int, channel Int, x *Int, y *Int, w *Int, h *Int) Int {
	return (Int)(C.goglQueryChannelDeltasSGIX((*C.Display)(display), (C.int)(screen), (C.int)(channel), (*C.int)(x), (*C.int)(y), (*C.int)(w), (*C.int)(h)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glChannelRectSyncSGIX.xml
func ChannelRectSyncSGIX(display Pointer, screen Int, channel Int, synctype Enum) Int {
	return (Int)(C.goglChannelRectSyncSGIX((*C.Display)(display), (C.int)(screen), (C.int)(channel), (C.GLenum)(synctype)))
}
// SGIX_video_source

// http://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXVideoSourceSGIX.xml
func CreateGLXVideoSourceSGIX(display Pointer, screen Int, server Pointer, path Pointer, nodeClass Int, drainNode Pointer) Pointer {
	return (Pointer)(C.goglCreateGLXVideoSourceSGIX((*C.Display)(display), (C.int)(screen), (C.VLServer)(server), (C.VLPath)(path), (C.int)(nodeClass), (C.VLNode)(drainNode)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyGLXVideoSourceSGIX.xml
func DestroyGLXVideoSourceSGIX(dpy Pointer, glxvideosource Pointer)  {
	C.goglDestroyGLXVideoSourceSGIX((*C.Display)(dpy), (*C.GGLXVideoSourceSGIX)(glxvideosource))
}
// SGIX_visual_select_group

// SGI_cushion

// http://www.opengl.org/sdk/docs/man/xhtml/glCushionSGI.xml
func CushionSGI(dpy Pointer, window Pointer, cushion float32)  {
	C.goglCushionSGI((*C.Display)(dpy), (C.Window)(window), (C.float32)(cushion))
}
// SGI_make_current_read

// http://www.opengl.org/sdk/docs/man/xhtml/glMakeCurrentReadSGI.xml
func MakeCurrentReadSGI(dpy Pointer, draw Pointer, read Pointer, ctx Pointer) int {
	return (int)(C.goglMakeCurrentReadSGI((*C.Display)(dpy), (C.GLXDrawable)(draw), (C.GLXDrawable)(read), (C.GLXContext)(ctx)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetCurrentReadDrawableSGI.xml
func GetCurrentReadDrawableSGI() Pointer {
	return (Pointer)(C.goglGetCurrentReadDrawableSGI())
}
// SGI_swap_control

// http://www.opengl.org/sdk/docs/man/xhtml/glSwapIntervalSGI.xml
func SwapIntervalSGI(interval Int) Int {
	return (Int)(C.goglSwapIntervalSGI((C.int)(interval)))
}
// SGI_video_sync

// http://www.opengl.org/sdk/docs/man/xhtml/glGetVideoSyncSGI.xml
func GetVideoSyncSGI(count *uint32) Int {
	return (Int)(C.goglGetVideoSyncSGI((*C.uint32)(count)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWaitVideoSyncSGI.xml
func WaitVideoSyncSGI(divisor Int, remainder Int, count *uint32) Int {
	return (Int)(C.goglWaitVideoSyncSGI((C.int)(divisor), (C.int)(remainder), (*C.uint32)(count)))
}
// SUN_get_transparent_index

// http://www.opengl.org/sdk/docs/man/xhtml/glGetTransparentIndexSUN.xml
func GetTransparentIndexSUN(dpy Pointer, overlay Pointer, underlay Pointer, pTransparentIndex *int32) int {
	return (int)(C.goglGetTransparentIndexSUN((*C.Display)(dpy), (C.Window)(overlay), (C.Window)(underlay), (*C.long)(pTransparentIndex)))
}
// VERSION_1_3

// http://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigs.xml
func GetFBConfigs(dpy Pointer, screen Int, nelements *Int) Pointer {
	return (Pointer)(C.goglGetFBConfigs((*C.Display)(dpy), (C.int)(screen), (*C.int)(nelements)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glChooseFBConfig.xml
func ChooseFBConfig(dpy Pointer, screen Int, attrib_list *Int, nelements *Int) Pointer {
	return (Pointer)(C.goglChooseFBConfig((*C.Display)(dpy), (C.int)(screen), (*C.int)(attrib_list), (*C.int)(nelements)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigAttrib.xml
func GetFBConfigAttrib(dpy Pointer, config Pointer, attribute Int, value *Int) Int {
	return (Int)(C.goglGetFBConfigAttrib((*C.Display)(dpy), (C.GLXFBConfig)(config), (C.int)(attribute), (*C.int)(value)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetVisualFromFBConfig.xml
func GetVisualFromFBConfig(dpy Pointer, config Pointer) Pointer {
	return (Pointer)(C.goglGetVisualFromFBConfig((*C.Display)(dpy), (C.GLXFBConfig)(config)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateWindow.xml
func CreateWindow(dpy Pointer, config Pointer, win Pointer, attrib_list *Int) Pointer {
	return (Pointer)(C.goglCreateWindow((*C.Display)(dpy), (C.GLXFBConfig)(config), (C.Window)(win), (*C.int)(attrib_list)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyWindow.xml
func DestroyWindow(dpy Pointer, win Pointer)  {
	C.goglDestroyWindow((*C.Display)(dpy), (C.GLXWindow)(win))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreatePixmap.xml
func CreatePixmap(dpy Pointer, config Pointer, pixmap Pointer, attrib_list *Int) Pointer {
	return (Pointer)(C.goglCreatePixmap((*C.Display)(dpy), (C.GLXFBConfig)(config), (C.Pixmap)(pixmap), (*C.int)(attrib_list)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyPixmap.xml
func DestroyPixmap(dpy Pointer, pixmap Pointer)  {
	C.goglDestroyPixmap((*C.Display)(dpy), (C.GLXPixmap)(pixmap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreatePbuffer.xml
func CreatePbuffer(dpy Pointer, config Pointer, attrib_list *Int) Pointer {
	return (Pointer)(C.goglCreatePbuffer((*C.Display)(dpy), (C.GLXFBConfig)(config), (*C.int)(attrib_list)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyPbuffer.xml
func DestroyPbuffer(dpy Pointer, pbuf Pointer)  {
	C.goglDestroyPbuffer((*C.Display)(dpy), (C.GLXPbuffer)(pbuf))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryDrawable.xml
func QueryDrawable(dpy Pointer, draw Pointer, attribute Int, value *uint32)  {
	C.goglQueryDrawable((*C.Display)(dpy), (C.GLXDrawable)(draw), (C.int)(attribute), (*C.uint32)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateNewContext.xml
func CreateNewContext(dpy Pointer, config Pointer, render_type Int, share_list Pointer, direct int) Pointer {
	return (Pointer)(C.goglCreateNewContext((*C.Display)(dpy), (C.GLXFBConfig)(config), (C.int)(render_type), (C.GLXContext)(share_list), (C.int)(direct)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMakeContextCurrent.xml
func MakeContextCurrent(dpy Pointer, draw Pointer, read Pointer, ctx Pointer) int {
	return (int)(C.goglMakeContextCurrent((*C.Display)(dpy), (C.GLXDrawable)(draw), (C.GLXDrawable)(read), (C.GLXContext)(ctx)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetCurrentReadDrawable.xml
func GetCurrentReadDrawable() Pointer {
	return (Pointer)(C.goglGetCurrentReadDrawable())
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetCurrentDisplay.xml
func GetCurrentDisplay() Pointer {
	return (Pointer)(C.goglGetCurrentDisplay())
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryContext.xml
func QueryContext(dpy Pointer, ctx Pointer, attribute Int, value *Int) Int {
	return (Int)(C.goglQueryContext((*C.Display)(dpy), (C.GLXContext)(ctx), (C.int)(attribute), (*C.int)(value)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSelectEvent.xml
func SelectEvent(dpy Pointer, draw Pointer, event_mask uint32)  {
	C.goglSelectEvent((*C.Display)(dpy), (C.GLXDrawable)(draw), (C.unsigned_long)(event_mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetSelectedEvent.xml
func GetSelectedEvent(dpy Pointer, draw Pointer, event_mask *uint32)  {
	C.goglGetSelectedEvent((*C.Display)(dpy), (C.GLXDrawable)(draw), (*C.unsigned_long)(event_mask))
}
// VERSION_1_4

// http://www.opengl.org/sdk/docs/man/xhtml/glGetProcAddress.xml
func GetProcAddress(procName *Ubyte) Pointer {
	return (Pointer)(C.goglGetProcAddress((*C.GLubyte)(procName)))
}
// glx

// http://www.opengl.org/sdk/docs/man/xhtml/glRender.xml
func Render()  {
	C.goglRender()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRenderLarge.xml
func RenderLarge()  {
	C.goglRenderLarge()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateContext.xml
func CreateContext(gc_id Int, screen Int, visual Int, share_list Int)  {
	C.goglCreateContext((C.GLint)(gc_id), (C.GLint)(screen), (C.GLint)(visual), (C.GLint)(share_list))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyContext.xml
func DestroyContext(context Int)  {
	C.goglDestroyContext((C.GLint)(context))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMakeCurrent.xml
func MakeCurrent(drawable Int, context Int)  {
	C.goglMakeCurrent((C.GLint)(drawable), (C.GLint)(context))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIsDirect.xml
func IsDirect(dpy Int, context Int)  {
	C.goglIsDirect((C.GLint)(dpy), (C.GLint)(context))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryVersion.xml
func QueryVersion(major *Int, minor *Int)  {
	C.goglQueryVersion((*C.GLint)(major), (*C.GLint)(minor))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWaitGL.xml
func WaitGL(context Int)  {
	C.goglWaitGL((C.GLint)(context))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWaitX.xml
func WaitX()  {
	C.goglWaitX()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyContext.xml
func CopyContext(source Int, dest Int, mask Int)  {
	C.goglCopyContext((C.GLint)(source), (C.GLint)(dest), (C.GLint)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSwapBuffers.xml
func SwapBuffers(drawable Int)  {
	C.goglSwapBuffers((C.GLint)(drawable))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUseXFont.xml
func UseXFont(font Int, first Int, count Int, list_base Int)  {
	C.goglUseXFont((C.GLint)(font), (C.GLint)(first), (C.GLint)(count), (C.GLint)(list_base))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPixmap.xml
func CreateGLXPixmap(visual Int, pixmap Int, glxpixmap Int)  {
	C.goglCreateGLXPixmap((C.GLint)(visual), (C.GLint)(pixmap), (C.GLint)(glxpixmap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetVisualConfigs.xml
func GetVisualConfigs()  {
	C.goglGetVisualConfigs()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyGLXPixmap.xml
func DestroyGLXPixmap(pixmap Int)  {
	C.goglDestroyGLXPixmap((C.GLint)(pixmap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVendorPrivate.xml
func VendorPrivate()  {
	C.goglVendorPrivate()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVendorPrivateWithReply.xml
func VendorPrivateWithReply()  {
	C.goglVendorPrivateWithReply()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryExtensionsString.xml
func QueryExtensionsString(screen Int)  {
	C.goglQueryExtensionsString((C.GLint)(screen))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryServerString.xml
func QueryServerString(screen Int, name Int)  {
	C.goglQueryServerString((C.GLint)(screen), (C.GLint)(name))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClientInfo.xml
func ClientInfo()  {
	C.goglClientInfo()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigs.xml
func GetFBConfigs()  {
	C.goglGetFBConfigs()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreatePixmap.xml
func CreatePixmap(config Int, pixmap Int, glxpixmap Int)  {
	C.goglCreatePixmap((C.GLint)(config), (C.GLint)(pixmap), (C.GLint)(glxpixmap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyPixmap.xml
func DestroyPixmap(glxpixmap Int)  {
	C.goglDestroyPixmap((C.GLint)(glxpixmap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateNewContext.xml
func CreateNewContext(config Int, render_type Int, share_list Int, direct Int)  {
	C.goglCreateNewContext((C.GLint)(config), (C.GLint)(render_type), (C.GLint)(share_list), (C.GLint)(direct))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryContext.xml
func QueryContext()  {
	C.goglQueryContext()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMakeContextCurrent.xml
func MakeContextCurrent(drawable Int, readdrawable Int, context Int)  {
	C.goglMakeContextCurrent((C.GLint)(drawable), (C.GLint)(readdrawable), (C.GLint)(context))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreatePbuffer.xml
func CreatePbuffer(config Int, pbuffer Int)  {
	C.goglCreatePbuffer((C.GLint)(config), (C.GLint)(pbuffer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyPbuffer.xml
func DestroyPbuffer(pbuffer Int)  {
	C.goglDestroyPbuffer((C.GLint)(pbuffer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetDrawableAttributes.xml
func GetDrawableAttributes(drawable Int)  {
	C.goglGetDrawableAttributes((C.GLint)(drawable))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glChangeDrawableAttributes.xml
func ChangeDrawableAttributes(drawable Int)  {
	C.goglChangeDrawableAttributes((C.GLint)(drawable))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateWindow.xml
func CreateWindow(config Int, window Int, glxwindow Int)  {
	C.goglCreateWindow((C.GLint)(config), (C.GLint)(window), (C.GLint)(glxwindow))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyWindow.xml
func DestroyWindow(glxwindow Int)  {
	C.goglDestroyWindow((C.GLint)(glxwindow))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSwapIntervalSGI.xml
func SwapIntervalSGI()  {
	C.goglSwapIntervalSGI()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMakeCurrentReadSGI.xml
func MakeCurrentReadSGI(drawable Int, readdrawable Int, context Int)  {
	C.goglMakeCurrentReadSGI((C.GLint)(drawable), (C.GLint)(readdrawable), (C.GLint)(context))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXVideoSourceSGIX.xml
func CreateGLXVideoSourceSGIX(dpy Int, screen Int, server Int, path Int, class Int, node Int)  {
	C.goglCreateGLXVideoSourceSGIX((C.GLint)(dpy), (C.GLint)(screen), (C.GLint)(server), (C.GLint)(path), (C.GLint)(class), (C.GLint)(node))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyGLXVideoSourceSGIX.xml
func DestroyGLXVideoSourceSGIX(dpy Int, glxvideosource Int)  {
	C.goglDestroyGLXVideoSourceSGIX((C.GLint)(dpy), (C.GLint)(glxvideosource))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryContextInfoEXT.xml
func QueryContextInfoEXT()  {
	C.goglQueryContextInfoEXT()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetFBConfigsSGIX.xml
func GetFBConfigsSGIX()  {
	C.goglGetFBConfigsSGIX()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateContextWithConfigSGIX.xml
func CreateContextWithConfigSGIX(gc_id Int, screen Int, config Int, share_list Int)  {
	C.goglCreateContextWithConfigSGIX((C.GLint)(gc_id), (C.GLint)(screen), (C.GLint)(config), (C.GLint)(share_list))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPixmapWithConfigSGIX.xml
func CreateGLXPixmapWithConfigSGIX(config Int, pixmap Int, glxpixmap Int)  {
	C.goglCreateGLXPixmapWithConfigSGIX((C.GLint)(config), (C.GLint)(pixmap), (C.GLint)(glxpixmap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateGLXPbufferSGIX.xml
func CreateGLXPbufferSGIX(config Int, pbuffer Int)  {
	C.goglCreateGLXPbufferSGIX((C.GLint)(config), (C.GLint)(pbuffer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyGLXPbufferSGIX.xml
func DestroyGLXPbufferSGIX(pbuffer Int)  {
	C.goglDestroyGLXPbufferSGIX((C.GLint)(pbuffer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glChangeDrawableAttributesSGIX.xml
func ChangeDrawableAttributesSGIX(drawable Int)  {
	C.goglChangeDrawableAttributesSGIX((C.GLint)(drawable))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetDrawableAttributesSGIX.xml
func GetDrawableAttributesSGIX(drawable Int)  {
	C.goglGetDrawableAttributesSGIX((C.GLint)(drawable))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glJoinSwapGroupSGIX.xml
func JoinSwapGroupSGIX(window Int, group Int)  {
	C.goglJoinSwapGroupSGIX((C.GLint)(window), (C.GLint)(group))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBindSwapBarrierSGIX.xml
func BindSwapBarrierSGIX(window Int, barrier Int)  {
	C.goglBindSwapBarrierSGIX((C.GLint)(window), (C.GLint)(barrier))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryMaxSwapBarriersSGIX.xml
func QueryMaxSwapBarriersSGIX()  {
	C.goglQueryMaxSwapBarriersSGIX()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeNetworkSGIX.xml
func QueryHyperpipeNetworkSGIX(dpy Pointer, npipes *Int) Pointer {
	return (Pointer)(C.goglQueryHyperpipeNetworkSGIX((*C.Display)(dpy), (*C.int)(npipes)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glHyperpipeConfigSGIX.xml
func HyperpipeConfigSGIX(dpy Pointer, networkId Int, npipes Int, cfg Pointer, hpId *Int) Int {
	return (Int)(C.goglHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(networkId), (C.int)(npipes), (*C.GLXHyperpipeConfigSGIX)(cfg), (*C.int)(hpId)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeConfigSGIX.xml
func QueryHyperpipeConfigSGIX(dpy Pointer, hpId Int, npipes *Int) Pointer {
	return (Pointer)(C.goglQueryHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(hpId), (*C.int)(npipes)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDestroyHyperpipeConfigSGIX.xml
func DestroyHyperpipeConfigSGIX(dpy Pointer, hpId Int) Int {
	return (Int)(C.goglDestroyHyperpipeConfigSGIX((*C.Display)(dpy), (C.int)(hpId)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBindHyperpipeSGIX.xml
func BindHyperpipeSGIX(dpy Pointer, hpId Int) Int {
	return (Int)(C.goglBindHyperpipeSGIX((*C.Display)(dpy), (C.int)(hpId)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeBestAttribSGIX.xml
func QueryHyperpipeBestAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, attribList Pointer, returnAttribList Pointer) Int {
	return (Int)(C.goglQueryHyperpipeBestAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(attribList), (unsafe.Pointer)(returnAttribList)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glHyperpipeAttribSGIX.xml
func HyperpipeAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, attribList Pointer) Int {
	return (Int)(C.goglHyperpipeAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(attribList)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glQueryHyperpipeAttribSGIX.xml
func QueryHyperpipeAttribSGIX(dpy Pointer, timeSlice Int, attrib Int, size Int, returnAttribList Pointer) Int {
	return (Int)(C.goglQueryHyperpipeAttribSGIX((*C.Display)(dpy), (C.int)(timeSlice), (C.int)(attrib), (C.int)(size), (unsafe.Pointer)(returnAttribList)))
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
func InitArbFbconfigFloat() error {
	var ret C.int
	if ret = C.init_ARB_fbconfig_float(); ret != 0 {
		return errors.New("unable to initialize ARB_fbconfig_float")
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
func InitArbGetProcAddress() error {
	var ret C.int
	if ret = C.init_ARB_get_proc_address(); ret != 0 {
		return errors.New("unable to initialize ARB_get_proc_address")
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
func InitExtFbconfigPackedFloat() error {
	var ret C.int
	if ret = C.init_EXT_fbconfig_packed_float(); ret != 0 {
		return errors.New("unable to initialize EXT_fbconfig_packed_float")
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
func InitExtSwapControlTear() error {
	var ret C.int
	if ret = C.init_EXT_swap_control_tear(); ret != 0 {
		return errors.New("unable to initialize EXT_swap_control_tear")
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
func InitExtVisualInfo() error {
	var ret C.int
	if ret = C.init_EXT_visual_info(); ret != 0 {
		return errors.New("unable to initialize EXT_visual_info")
	}
	return nil
}
func InitExtVisualRating() error {
	var ret C.int
	if ret = C.init_EXT_visual_rating(); ret != 0 {
		return errors.New("unable to initialize EXT_visual_rating")
	}
	return nil
}
func InitIntelSwapEvent() error {
	var ret C.int
	if ret = C.init_INTEL_swap_event(); ret != 0 {
		return errors.New("unable to initialize INTEL_swap_event")
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
func InitNvFloatBuffer() error {
	var ret C.int
	if ret = C.init_NV_float_buffer(); ret != 0 {
		return errors.New("unable to initialize NV_float_buffer")
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
func InitOmlSwapMethod() error {
	var ret C.int
	if ret = C.init_OML_swap_method(); ret != 0 {
		return errors.New("unable to initialize OML_swap_method")
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
func InitSgisMultisample() error {
	var ret C.int
	if ret = C.init_SGIS_multisample(); ret != 0 {
		return errors.New("unable to initialize SGIS_multisample")
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
func InitSgixVisualSelectGroup() error {
	var ret C.int
	if ret = C.init_SGIX_visual_select_group(); ret != 0 {
		return errors.New("unable to initialize SGIX_visual_select_group")
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