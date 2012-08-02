// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// VERSION_1_0
// 
// VERSION_1_0_DEPRECATED
// 
// VERSION_1_1
// 
// VERSION_1_1_DEPRECATED
// 
// VERSION_1_2
// 
// VERSION_1_2_DEPRECATED
// 
// VERSION_1_3
// 
// VERSION_1_3_DEPRECATED
// 
// VERSION_1_4
// 
// VERSION_1_4_DEPRECATED
// 
// VERSION_1_5
// 
// VERSION_2_0
// 
// VERSION_2_1
// 
// http://www.opengl.org/sdk/docs/man
// 
package gl21

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
// #include <stddef.h>
// #ifndef GL_VERSION_2_0
// /* GL type for program/shader text */
// typedef char GLchar;
// #endif
// 
// #ifndef GL_VERSION_1_5
// /* GL types for handling large vertex buffer objects */
// typedef ptrdiff_t GLintptr;
// typedef ptrdiff_t GLsizeiptr;
// #endif
// 
// #ifndef GL_ARB_vertex_buffer_object
// /* GL types for handling large vertex buffer objects */
// typedef ptrdiff_t GLintptrARB;
// typedef ptrdiff_t GLsizeiptrARB;
// #endif
// 
// #ifndef GL_ARB_shader_objects
// /* GL types for program/shader text and shader object handles */
// typedef char GLcharARB;
// typedef unsigned int GLhandleARB;
// #endif
// 
// /* GL type for "half" precision (s10e5) float data in host memory */
// #ifndef GL_ARB_half_float_pixel
// typedef unsigned short GLhalfARB;
// #endif
// 
// #ifndef GL_NV_half_float
// typedef unsigned short GLhalfNV;
// #endif
// 
// #ifndef GLEXT_64_TYPES_DEFINED
// /* This code block is duplicated in glxext.h, so must be protected */
// #define GLEXT_64_TYPES_DEFINED
// /* Define int32_t, int64_t, and uint64_t types for UST/MSC */
// /* (as used in the GL_EXT_timer_query extension). */
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
// /* Fallback if nothing above works */
// #include <inttypes.h>
// #endif
// #endif
// 
// #ifndef GL_EXT_timer_query
// typedef int64_t GLint64EXT;
// typedef uint64_t GLuint64EXT;
// #endif
// 
// #ifndef GL_ARB_sync
// typedef int64_t GLint64;
// typedef uint64_t GLuint64;
// typedef struct __GLsync *GLsync;
// #endif
// 
// #ifndef GL_ARB_cl_event
// /* These incomplete types let us declare types compatible with OpenCL's cl_context and cl_event */
// struct _cl_context;
// struct _cl_event;
// #endif
// 
// #ifndef GL_ARB_debug_output
// typedef void (APIENTRY *GLDEBUGPROCARB)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,GLvoid *userParam);
// #endif
// 
// #ifndef GL_AMD_debug_output
// typedef void (APIENTRY *GLDEBUGPROCAMD)(GLuint id,GLenum category,GLenum severity,GLsizei length,const GLchar *message,GLvoid *userParam);
// #endif
// 
// #ifndef GL_NV_vdpau_interop
// typedef GLintptr GLvdpauSurfaceNV;
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
// //  VERSION_1_0
// void (APIENTRYP ptrglCullFace)(GLenum mode);
// void (APIENTRYP ptrglFrontFace)(GLenum mode);
// void (APIENTRYP ptrglHint)(GLenum target, GLenum mode);
// void (APIENTRYP ptrglLineWidth)(GLfloat width);
// void (APIENTRYP ptrglPointSize)(GLfloat size);
// void (APIENTRYP ptrglPolygonMode)(GLenum face, GLenum mode);
// void (APIENTRYP ptrglScissor)(GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglTexParameterf)(GLenum target, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglTexParameteri)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrglTexParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglTexImage1D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglTexImage2D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglDrawBuffer)(GLenum mode);
// void (APIENTRYP ptrglClear)(GLbitfield mask);
// void (APIENTRYP ptrglClearColor)(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
// void (APIENTRYP ptrglClearStencil)(GLint s);
// void (APIENTRYP ptrglClearDepth)(GLclampd depth);
// void (APIENTRYP ptrglStencilMask)(GLuint mask);
// void (APIENTRYP ptrglColorMask)(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
// void (APIENTRYP ptrglDepthMask)(GLboolean flag);
// void (APIENTRYP ptrglDisable)(GLenum cap);
// void (APIENTRYP ptrglEnable)(GLenum cap);
// void (APIENTRYP ptrglFinish)();
// void (APIENTRYP ptrglFlush)();
// void (APIENTRYP ptrglBlendFunc)(GLenum sfactor, GLenum dfactor);
// void (APIENTRYP ptrglLogicOp)(GLenum opcode);
// void (APIENTRYP ptrglStencilFunc)(GLenum func, GLint ref, GLuint mask);
// void (APIENTRYP ptrglStencilOp)(GLenum fail, GLenum zfail, GLenum zpass);
// void (APIENTRYP ptrglDepthFunc)(GLenum func);
// void (APIENTRYP ptrglPixelStoref)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglPixelStorei)(GLenum pname, GLint param);
// void (APIENTRYP ptrglReadBuffer)(GLenum mode);
// void (APIENTRYP ptrglReadPixels)(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglGetBooleanv)(GLenum pname, GLboolean* params);
// void (APIENTRYP ptrglGetDoublev)(GLenum pname, GLdouble* params);
// GLenum (APIENTRYP ptrglGetError)();
// void (APIENTRYP ptrglGetFloatv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetIntegerv)(GLenum pname, GLint* params);
// const GLubyte * (APIENTRYP ptrglGetString)(GLenum name);
// void (APIENTRYP ptrglGetTexImage)(GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglGetTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetTexParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetTexLevelParameterfv)(GLenum target, GLint level, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetTexLevelParameteriv)(GLenum target, GLint level, GLenum pname, GLint* params);
// GLboolean (APIENTRYP ptrglIsEnabled)(GLenum cap);
// void (APIENTRYP ptrglDepthRange)(GLclampd near, GLclampd far);
// void (APIENTRYP ptrglViewport)(GLint x, GLint y, GLsizei width, GLsizei height);
// //  VERSION_1_0_DEPRECATED
// void (APIENTRYP ptrglNewList)(GLuint list, GLenum mode);
// void (APIENTRYP ptrglEndList)();
// void (APIENTRYP ptrglCallList)(GLuint list);
// void (APIENTRYP ptrglCallLists)(GLsizei n, GLenum type, GLvoid* lists);
// void (APIENTRYP ptrglDeleteLists)(GLuint list, GLsizei range);
// GLuint (APIENTRYP ptrglGenLists)(GLsizei range);
// void (APIENTRYP ptrglListBase)(GLuint base);
// void (APIENTRYP ptrglBegin)(GLenum mode);
// void (APIENTRYP ptrglBitmap)(GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap);
// void (APIENTRYP ptrglColor3b)(GLbyte red, GLbyte green, GLbyte blue);
// void (APIENTRYP ptrglColor3bv)(GLbyte* v);
// void (APIENTRYP ptrglColor3d)(GLdouble red, GLdouble green, GLdouble blue);
// void (APIENTRYP ptrglColor3dv)(GLdouble* v);
// void (APIENTRYP ptrglColor3f)(GLfloat red, GLfloat green, GLfloat blue);
// void (APIENTRYP ptrglColor3fv)(GLfloat* v);
// void (APIENTRYP ptrglColor3i)(GLint red, GLint green, GLint blue);
// void (APIENTRYP ptrglColor3iv)(GLint* v);
// void (APIENTRYP ptrglColor3s)(GLshort red, GLshort green, GLshort blue);
// void (APIENTRYP ptrglColor3sv)(GLshort* v);
// void (APIENTRYP ptrglColor3ub)(GLubyte red, GLubyte green, GLubyte blue);
// void (APIENTRYP ptrglColor3ubv)(GLubyte* v);
// void (APIENTRYP ptrglColor3ui)(GLuint red, GLuint green, GLuint blue);
// void (APIENTRYP ptrglColor3uiv)(GLuint* v);
// void (APIENTRYP ptrglColor3us)(GLushort red, GLushort green, GLushort blue);
// void (APIENTRYP ptrglColor3usv)(GLushort* v);
// void (APIENTRYP ptrglColor4b)(GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha);
// void (APIENTRYP ptrglColor4bv)(GLbyte* v);
// void (APIENTRYP ptrglColor4d)(GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha);
// void (APIENTRYP ptrglColor4dv)(GLdouble* v);
// void (APIENTRYP ptrglColor4f)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglColor4fv)(GLfloat* v);
// void (APIENTRYP ptrglColor4i)(GLint red, GLint green, GLint blue, GLint alpha);
// void (APIENTRYP ptrglColor4iv)(GLint* v);
// void (APIENTRYP ptrglColor4s)(GLshort red, GLshort green, GLshort blue, GLshort alpha);
// void (APIENTRYP ptrglColor4sv)(GLshort* v);
// void (APIENTRYP ptrglColor4ub)(GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha);
// void (APIENTRYP ptrglColor4ubv)(GLubyte* v);
// void (APIENTRYP ptrglColor4ui)(GLuint red, GLuint green, GLuint blue, GLuint alpha);
// void (APIENTRYP ptrglColor4uiv)(GLuint* v);
// void (APIENTRYP ptrglColor4us)(GLushort red, GLushort green, GLushort blue, GLushort alpha);
// void (APIENTRYP ptrglColor4usv)(GLushort* v);
// void (APIENTRYP ptrglEdgeFlag)(GLboolean flag);
// void (APIENTRYP ptrglEdgeFlagv)(GLboolean* flag);
// void (APIENTRYP ptrglEnd)();
// void (APIENTRYP ptrglIndexd)(GLdouble c);
// void (APIENTRYP ptrglIndexdv)(GLdouble* c);
// void (APIENTRYP ptrglIndexf)(GLfloat c);
// void (APIENTRYP ptrglIndexfv)(GLfloat* c);
// void (APIENTRYP ptrglIndexi)(GLint c);
// void (APIENTRYP ptrglIndexiv)(GLint* c);
// void (APIENTRYP ptrglIndexs)(GLshort c);
// void (APIENTRYP ptrglIndexsv)(GLshort* c);
// void (APIENTRYP ptrglNormal3b)(GLbyte nx, GLbyte ny, GLbyte nz);
// void (APIENTRYP ptrglNormal3bv)(GLbyte* v);
// void (APIENTRYP ptrglNormal3d)(GLdouble nx, GLdouble ny, GLdouble nz);
// void (APIENTRYP ptrglNormal3dv)(GLdouble* v);
// void (APIENTRYP ptrglNormal3f)(GLfloat nx, GLfloat ny, GLfloat nz);
// void (APIENTRYP ptrglNormal3fv)(GLfloat* v);
// void (APIENTRYP ptrglNormal3i)(GLint nx, GLint ny, GLint nz);
// void (APIENTRYP ptrglNormal3iv)(GLint* v);
// void (APIENTRYP ptrglNormal3s)(GLshort nx, GLshort ny, GLshort nz);
// void (APIENTRYP ptrglNormal3sv)(GLshort* v);
// void (APIENTRYP ptrglRasterPos2d)(GLdouble x, GLdouble y);
// void (APIENTRYP ptrglRasterPos2dv)(GLdouble* v);
// void (APIENTRYP ptrglRasterPos2f)(GLfloat x, GLfloat y);
// void (APIENTRYP ptrglRasterPos2fv)(GLfloat* v);
// void (APIENTRYP ptrglRasterPos2i)(GLint x, GLint y);
// void (APIENTRYP ptrglRasterPos2iv)(GLint* v);
// void (APIENTRYP ptrglRasterPos2s)(GLshort x, GLshort y);
// void (APIENTRYP ptrglRasterPos2sv)(GLshort* v);
// void (APIENTRYP ptrglRasterPos3d)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglRasterPos3dv)(GLdouble* v);
// void (APIENTRYP ptrglRasterPos3f)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglRasterPos3fv)(GLfloat* v);
// void (APIENTRYP ptrglRasterPos3i)(GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglRasterPos3iv)(GLint* v);
// void (APIENTRYP ptrglRasterPos3s)(GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrglRasterPos3sv)(GLshort* v);
// void (APIENTRYP ptrglRasterPos4d)(GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglRasterPos4dv)(GLdouble* v);
// void (APIENTRYP ptrglRasterPos4f)(GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglRasterPos4fv)(GLfloat* v);
// void (APIENTRYP ptrglRasterPos4i)(GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglRasterPos4iv)(GLint* v);
// void (APIENTRYP ptrglRasterPos4s)(GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrglRasterPos4sv)(GLshort* v);
// void (APIENTRYP ptrglRectd)(GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2);
// void (APIENTRYP ptrglRectdv)(GLdouble* v1, GLdouble* v2);
// void (APIENTRYP ptrglRectf)(GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2);
// void (APIENTRYP ptrglRectfv)(GLfloat* v1, GLfloat* v2);
// void (APIENTRYP ptrglRecti)(GLint x1, GLint y1, GLint x2, GLint y2);
// void (APIENTRYP ptrglRectiv)(GLint* v1, GLint* v2);
// void (APIENTRYP ptrglRects)(GLshort x1, GLshort y1, GLshort x2, GLshort y2);
// void (APIENTRYP ptrglRectsv)(GLshort* v1, GLshort* v2);
// void (APIENTRYP ptrglTexCoord1d)(GLdouble s);
// void (APIENTRYP ptrglTexCoord1dv)(GLdouble* v);
// void (APIENTRYP ptrglTexCoord1f)(GLfloat s);
// void (APIENTRYP ptrglTexCoord1fv)(GLfloat* v);
// void (APIENTRYP ptrglTexCoord1i)(GLint s);
// void (APIENTRYP ptrglTexCoord1iv)(GLint* v);
// void (APIENTRYP ptrglTexCoord1s)(GLshort s);
// void (APIENTRYP ptrglTexCoord1sv)(GLshort* v);
// void (APIENTRYP ptrglTexCoord2d)(GLdouble s, GLdouble t);
// void (APIENTRYP ptrglTexCoord2dv)(GLdouble* v);
// void (APIENTRYP ptrglTexCoord2f)(GLfloat s, GLfloat t);
// void (APIENTRYP ptrglTexCoord2fv)(GLfloat* v);
// void (APIENTRYP ptrglTexCoord2i)(GLint s, GLint t);
// void (APIENTRYP ptrglTexCoord2iv)(GLint* v);
// void (APIENTRYP ptrglTexCoord2s)(GLshort s, GLshort t);
// void (APIENTRYP ptrglTexCoord2sv)(GLshort* v);
// void (APIENTRYP ptrglTexCoord3d)(GLdouble s, GLdouble t, GLdouble r);
// void (APIENTRYP ptrglTexCoord3dv)(GLdouble* v);
// void (APIENTRYP ptrglTexCoord3f)(GLfloat s, GLfloat t, GLfloat r);
// void (APIENTRYP ptrglTexCoord3fv)(GLfloat* v);
// void (APIENTRYP ptrglTexCoord3i)(GLint s, GLint t, GLint r);
// void (APIENTRYP ptrglTexCoord3iv)(GLint* v);
// void (APIENTRYP ptrglTexCoord3s)(GLshort s, GLshort t, GLshort r);
// void (APIENTRYP ptrglTexCoord3sv)(GLshort* v);
// void (APIENTRYP ptrglTexCoord4d)(GLdouble s, GLdouble t, GLdouble r, GLdouble q);
// void (APIENTRYP ptrglTexCoord4dv)(GLdouble* v);
// void (APIENTRYP ptrglTexCoord4f)(GLfloat s, GLfloat t, GLfloat r, GLfloat q);
// void (APIENTRYP ptrglTexCoord4fv)(GLfloat* v);
// void (APIENTRYP ptrglTexCoord4i)(GLint s, GLint t, GLint r, GLint q);
// void (APIENTRYP ptrglTexCoord4iv)(GLint* v);
// void (APIENTRYP ptrglTexCoord4s)(GLshort s, GLshort t, GLshort r, GLshort q);
// void (APIENTRYP ptrglTexCoord4sv)(GLshort* v);
// void (APIENTRYP ptrglVertex2d)(GLdouble x, GLdouble y);
// void (APIENTRYP ptrglVertex2dv)(GLdouble* v);
// void (APIENTRYP ptrglVertex2f)(GLfloat x, GLfloat y);
// void (APIENTRYP ptrglVertex2fv)(GLfloat* v);
// void (APIENTRYP ptrglVertex2i)(GLint x, GLint y);
// void (APIENTRYP ptrglVertex2iv)(GLint* v);
// void (APIENTRYP ptrglVertex2s)(GLshort x, GLshort y);
// void (APIENTRYP ptrglVertex2sv)(GLshort* v);
// void (APIENTRYP ptrglVertex3d)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglVertex3dv)(GLdouble* v);
// void (APIENTRYP ptrglVertex3f)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglVertex3fv)(GLfloat* v);
// void (APIENTRYP ptrglVertex3i)(GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglVertex3iv)(GLint* v);
// void (APIENTRYP ptrglVertex3s)(GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrglVertex3sv)(GLshort* v);
// void (APIENTRYP ptrglVertex4d)(GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglVertex4dv)(GLdouble* v);
// void (APIENTRYP ptrglVertex4f)(GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglVertex4fv)(GLfloat* v);
// void (APIENTRYP ptrglVertex4i)(GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglVertex4iv)(GLint* v);
// void (APIENTRYP ptrglVertex4s)(GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrglVertex4sv)(GLshort* v);
// void (APIENTRYP ptrglClipPlane)(GLenum plane, GLdouble* equation);
// void (APIENTRYP ptrglColorMaterial)(GLenum face, GLenum mode);
// void (APIENTRYP ptrglFogf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglFogfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglFogi)(GLenum pname, GLint param);
// void (APIENTRYP ptrglFogiv)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglLightf)(GLenum light, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglLightfv)(GLenum light, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglLighti)(GLenum light, GLenum pname, GLint param);
// void (APIENTRYP ptrglLightiv)(GLenum light, GLenum pname, GLint* params);
// void (APIENTRYP ptrglLightModelf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglLightModelfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglLightModeli)(GLenum pname, GLint param);
// void (APIENTRYP ptrglLightModeliv)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglLineStipple)(GLint factor, GLushort pattern);
// void (APIENTRYP ptrglMaterialf)(GLenum face, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglMaterialfv)(GLenum face, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglMateriali)(GLenum face, GLenum pname, GLint param);
// void (APIENTRYP ptrglMaterialiv)(GLenum face, GLenum pname, GLint* params);
// void (APIENTRYP ptrglPolygonStipple)(GLubyte* mask);
// void (APIENTRYP ptrglShadeModel)(GLenum mode);
// void (APIENTRYP ptrglTexEnvf)(GLenum target, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglTexEnvfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglTexEnvi)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrglTexEnviv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglTexGend)(GLenum coord, GLenum pname, GLdouble param);
// void (APIENTRYP ptrglTexGendv)(GLenum coord, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglTexGenf)(GLenum coord, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglTexGenfv)(GLenum coord, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglTexGeni)(GLenum coord, GLenum pname, GLint param);
// void (APIENTRYP ptrglTexGeniv)(GLenum coord, GLenum pname, GLint* params);
// void (APIENTRYP ptrglFeedbackBuffer)(GLsizei size, GLenum type, GLfloat* buffer);
// void (APIENTRYP ptrglSelectBuffer)(GLsizei size, GLuint* buffer);
// GLint (APIENTRYP ptrglRenderMode)(GLenum mode);
// void (APIENTRYP ptrglInitNames)();
// void (APIENTRYP ptrglLoadName)(GLuint name);
// void (APIENTRYP ptrglPassThrough)(GLfloat token);
// void (APIENTRYP ptrglPopName)();
// void (APIENTRYP ptrglPushName)(GLuint name);
// void (APIENTRYP ptrglClearAccum)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglClearIndex)(GLfloat c);
// void (APIENTRYP ptrglIndexMask)(GLuint mask);
// void (APIENTRYP ptrglAccum)(GLenum op, GLfloat value);
// void (APIENTRYP ptrglPopAttrib)();
// void (APIENTRYP ptrglPushAttrib)(GLbitfield mask);
// void (APIENTRYP ptrglMap1d)(GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points);
// void (APIENTRYP ptrglMap1f)(GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points);
// void (APIENTRYP ptrglMap2d)(GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points);
// void (APIENTRYP ptrglMap2f)(GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points);
// void (APIENTRYP ptrglMapGrid1d)(GLint un, GLdouble u1, GLdouble u2);
// void (APIENTRYP ptrglMapGrid1f)(GLint un, GLfloat u1, GLfloat u2);
// void (APIENTRYP ptrglMapGrid2d)(GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2);
// void (APIENTRYP ptrglMapGrid2f)(GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2);
// void (APIENTRYP ptrglEvalCoord1d)(GLdouble u);
// void (APIENTRYP ptrglEvalCoord1dv)(GLdouble* u);
// void (APIENTRYP ptrglEvalCoord1f)(GLfloat u);
// void (APIENTRYP ptrglEvalCoord1fv)(GLfloat* u);
// void (APIENTRYP ptrglEvalCoord2d)(GLdouble u, GLdouble v);
// void (APIENTRYP ptrglEvalCoord2dv)(GLdouble* u);
// void (APIENTRYP ptrglEvalCoord2f)(GLfloat u, GLfloat v);
// void (APIENTRYP ptrglEvalCoord2fv)(GLfloat* u);
// void (APIENTRYP ptrglEvalMesh1)(GLenum mode, GLint i1, GLint i2);
// void (APIENTRYP ptrglEvalPoint1)(GLint i);
// void (APIENTRYP ptrglEvalMesh2)(GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2);
// void (APIENTRYP ptrglEvalPoint2)(GLint i, GLint j);
// void (APIENTRYP ptrglAlphaFunc)(GLenum func, GLclampf ref);
// void (APIENTRYP ptrglPixelZoom)(GLfloat xfactor, GLfloat yfactor);
// void (APIENTRYP ptrglPixelTransferf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglPixelTransferi)(GLenum pname, GLint param);
// void (APIENTRYP ptrglPixelMapfv)(GLenum map, GLsizei mapsize, GLfloat* values);
// void (APIENTRYP ptrglPixelMapuiv)(GLenum map, GLsizei mapsize, GLuint* values);
// void (APIENTRYP ptrglPixelMapusv)(GLenum map, GLsizei mapsize, GLushort* values);
// void (APIENTRYP ptrglCopyPixels)(GLint x, GLint y, GLsizei width, GLsizei height, GLenum type);
// void (APIENTRYP ptrglDrawPixels)(GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglGetClipPlane)(GLenum plane, GLdouble* equation);
// void (APIENTRYP ptrglGetLightfv)(GLenum light, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetLightiv)(GLenum light, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetMapdv)(GLenum target, GLenum query, GLdouble* v);
// void (APIENTRYP ptrglGetMapfv)(GLenum target, GLenum query, GLfloat* v);
// void (APIENTRYP ptrglGetMapiv)(GLenum target, GLenum query, GLint* v);
// void (APIENTRYP ptrglGetMaterialfv)(GLenum face, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetMaterialiv)(GLenum face, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetPixelMapfv)(GLenum map, GLfloat* values);
// void (APIENTRYP ptrglGetPixelMapuiv)(GLenum map, GLuint* values);
// void (APIENTRYP ptrglGetPixelMapusv)(GLenum map, GLushort* values);
// void (APIENTRYP ptrglGetPolygonStipple)(GLubyte* mask);
// void (APIENTRYP ptrglGetTexEnvfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetTexEnviv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetTexGendv)(GLenum coord, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglGetTexGenfv)(GLenum coord, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetTexGeniv)(GLenum coord, GLenum pname, GLint* params);
// GLboolean (APIENTRYP ptrglIsList)(GLuint list);
// void (APIENTRYP ptrglFrustum)(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar);
// void (APIENTRYP ptrglLoadIdentity)();
// void (APIENTRYP ptrglLoadMatrixf)(GLfloat* m);
// void (APIENTRYP ptrglLoadMatrixd)(GLdouble* m);
// void (APIENTRYP ptrglMatrixMode)(GLenum mode);
// void (APIENTRYP ptrglMultMatrixf)(GLfloat* m);
// void (APIENTRYP ptrglMultMatrixd)(GLdouble* m);
// void (APIENTRYP ptrglOrtho)(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar);
// void (APIENTRYP ptrglPopMatrix)();
// void (APIENTRYP ptrglPushMatrix)();
// void (APIENTRYP ptrglRotated)(GLdouble angle, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglRotatef)(GLfloat angle, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglScaled)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglScalef)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglTranslated)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglTranslatef)(GLfloat x, GLfloat y, GLfloat z);
// //  VERSION_1_1
// void (APIENTRYP ptrglDrawArrays)(GLenum mode, GLint first, GLsizei count);
// void (APIENTRYP ptrglDrawElements)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices);
// void (APIENTRYP ptrglGetPointerv)(GLenum pname, GLvoid** params);
// void (APIENTRYP ptrglPolygonOffset)(GLfloat factor, GLfloat units);
// void (APIENTRYP ptrglCopyTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border);
// void (APIENTRYP ptrglCopyTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border);
// void (APIENTRYP ptrglCopyTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrglCopyTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglBindTexture)(GLenum target, GLuint texture);
// void (APIENTRYP ptrglDeleteTextures)(GLsizei n, GLuint* textures);
// void (APIENTRYP ptrglGenTextures)(GLsizei n, GLuint* textures);
// GLboolean (APIENTRYP ptrglIsTexture)(GLuint texture);
// //  VERSION_1_1_DEPRECATED
// void (APIENTRYP ptrglArrayElement)(GLint i);
// void (APIENTRYP ptrglColorPointer)(GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglDisableClientState)(GLenum array);
// void (APIENTRYP ptrglEdgeFlagPointer)(GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglEnableClientState)(GLenum array);
// void (APIENTRYP ptrglIndexPointer)(GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglInterleavedArrays)(GLenum format, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglNormalPointer)(GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglTexCoordPointer)(GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglVertexPointer)(GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// GLboolean (APIENTRYP ptrglAreTexturesResident)(GLsizei n, GLuint* textures, GLboolean* residences);
// void (APIENTRYP ptrglPrioritizeTextures)(GLsizei n, GLuint* textures, GLclampf* priorities);
// void (APIENTRYP ptrglIndexub)(GLubyte c);
// void (APIENTRYP ptrglIndexubv)(GLubyte* c);
// void (APIENTRYP ptrglPopClientAttrib)();
// void (APIENTRYP ptrglPushClientAttrib)(GLbitfield mask);
// //  VERSION_1_2
// void (APIENTRYP ptrglBlendColor)(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
// void (APIENTRYP ptrglBlendEquation)(GLenum mode);
// void (APIENTRYP ptrglDrawRangeElements)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices);
// void (APIENTRYP ptrglTexImage3D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglCopyTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height);
// //  VERSION_1_2_DEPRECATED
// void (APIENTRYP ptrglColorTable)(GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* table);
// void (APIENTRYP ptrglColorTableParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglColorTableParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglCopyColorTable)(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrglGetColorTable)(GLenum target, GLenum format, GLenum type, GLvoid* table);
// void (APIENTRYP ptrglGetColorTableParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetColorTableParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglColorSubTable)(GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data);
// void (APIENTRYP ptrglCopyColorSubTable)(GLenum target, GLsizei start, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrglConvolutionFilter1D)(GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* image);
// void (APIENTRYP ptrglConvolutionFilter2D)(GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* image);
// void (APIENTRYP ptrglConvolutionParameterf)(GLenum target, GLenum pname, GLfloat params);
// void (APIENTRYP ptrglConvolutionParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglConvolutionParameteri)(GLenum target, GLenum pname, GLint params);
// void (APIENTRYP ptrglConvolutionParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglCopyConvolutionFilter1D)(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrglCopyConvolutionFilter2D)(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglGetConvolutionFilter)(GLenum target, GLenum format, GLenum type, GLvoid* image);
// void (APIENTRYP ptrglGetConvolutionParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetConvolutionParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetSeparableFilter)(GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span);
// void (APIENTRYP ptrglSeparableFilter2D)(GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column);
// void (APIENTRYP ptrglGetHistogram)(GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values);
// void (APIENTRYP ptrglGetHistogramParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetHistogramParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetMinmax)(GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values);
// void (APIENTRYP ptrglGetMinmaxParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetMinmaxParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglHistogram)(GLenum target, GLsizei width, GLenum internalformat, GLboolean sink);
// void (APIENTRYP ptrglMinmax)(GLenum target, GLenum internalformat, GLboolean sink);
// void (APIENTRYP ptrglResetHistogram)(GLenum target);
// void (APIENTRYP ptrglResetMinmax)(GLenum target);
// //  VERSION_1_3
// void (APIENTRYP ptrglActiveTexture)(GLenum texture);
// void (APIENTRYP ptrglSampleCoverage)(GLclampf value, GLboolean invert);
// void (APIENTRYP ptrglCompressedTexImage3D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglGetCompressedTexImage)(GLenum target, GLint level, GLvoid* img);
// //  VERSION_1_3_DEPRECATED
// void (APIENTRYP ptrglClientActiveTexture)(GLenum texture);
// void (APIENTRYP ptrglMultiTexCoord1d)(GLenum target, GLdouble s);
// void (APIENTRYP ptrglMultiTexCoord1dv)(GLenum target, GLdouble* v);
// void (APIENTRYP ptrglMultiTexCoord1f)(GLenum target, GLfloat s);
// void (APIENTRYP ptrglMultiTexCoord1fv)(GLenum target, GLfloat* v);
// void (APIENTRYP ptrglMultiTexCoord1i)(GLenum target, GLint s);
// void (APIENTRYP ptrglMultiTexCoord1iv)(GLenum target, GLint* v);
// void (APIENTRYP ptrglMultiTexCoord1s)(GLenum target, GLshort s);
// void (APIENTRYP ptrglMultiTexCoord1sv)(GLenum target, GLshort* v);
// void (APIENTRYP ptrglMultiTexCoord2d)(GLenum target, GLdouble s, GLdouble t);
// void (APIENTRYP ptrglMultiTexCoord2dv)(GLenum target, GLdouble* v);
// void (APIENTRYP ptrglMultiTexCoord2f)(GLenum target, GLfloat s, GLfloat t);
// void (APIENTRYP ptrglMultiTexCoord2fv)(GLenum target, GLfloat* v);
// void (APIENTRYP ptrglMultiTexCoord2i)(GLenum target, GLint s, GLint t);
// void (APIENTRYP ptrglMultiTexCoord2iv)(GLenum target, GLint* v);
// void (APIENTRYP ptrglMultiTexCoord2s)(GLenum target, GLshort s, GLshort t);
// void (APIENTRYP ptrglMultiTexCoord2sv)(GLenum target, GLshort* v);
// void (APIENTRYP ptrglMultiTexCoord3d)(GLenum target, GLdouble s, GLdouble t, GLdouble r);
// void (APIENTRYP ptrglMultiTexCoord3dv)(GLenum target, GLdouble* v);
// void (APIENTRYP ptrglMultiTexCoord3f)(GLenum target, GLfloat s, GLfloat t, GLfloat r);
// void (APIENTRYP ptrglMultiTexCoord3fv)(GLenum target, GLfloat* v);
// void (APIENTRYP ptrglMultiTexCoord3i)(GLenum target, GLint s, GLint t, GLint r);
// void (APIENTRYP ptrglMultiTexCoord3iv)(GLenum target, GLint* v);
// void (APIENTRYP ptrglMultiTexCoord3s)(GLenum target, GLshort s, GLshort t, GLshort r);
// void (APIENTRYP ptrglMultiTexCoord3sv)(GLenum target, GLshort* v);
// void (APIENTRYP ptrglMultiTexCoord4d)(GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q);
// void (APIENTRYP ptrglMultiTexCoord4dv)(GLenum target, GLdouble* v);
// void (APIENTRYP ptrglMultiTexCoord4f)(GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q);
// void (APIENTRYP ptrglMultiTexCoord4fv)(GLenum target, GLfloat* v);
// void (APIENTRYP ptrglMultiTexCoord4i)(GLenum target, GLint s, GLint t, GLint r, GLint q);
// void (APIENTRYP ptrglMultiTexCoord4iv)(GLenum target, GLint* v);
// void (APIENTRYP ptrglMultiTexCoord4s)(GLenum target, GLshort s, GLshort t, GLshort r, GLshort q);
// void (APIENTRYP ptrglMultiTexCoord4sv)(GLenum target, GLshort* v);
// void (APIENTRYP ptrglLoadTransposeMatrixf)(GLfloat* m);
// void (APIENTRYP ptrglLoadTransposeMatrixd)(GLdouble* m);
// void (APIENTRYP ptrglMultTransposeMatrixf)(GLfloat* m);
// void (APIENTRYP ptrglMultTransposeMatrixd)(GLdouble* m);
// //  VERSION_1_4
// void (APIENTRYP ptrglBlendFuncSeparate)(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha);
// void (APIENTRYP ptrglMultiDrawArrays)(GLenum mode, GLint* first, GLsizei* count, GLsizei primcount);
// void (APIENTRYP ptrglMultiDrawElements)(GLenum mode, GLsizei* count, GLenum type, GLvoid** indices, GLsizei primcount);
// void (APIENTRYP ptrglPointParameterf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglPointParameterfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglPointParameteri)(GLenum pname, GLint param);
// void (APIENTRYP ptrglPointParameteriv)(GLenum pname, GLint* params);
// //  VERSION_1_4_DEPRECATED
// void (APIENTRYP ptrglFogCoordf)(GLfloat coord);
// void (APIENTRYP ptrglFogCoordfv)(GLfloat* coord);
// void (APIENTRYP ptrglFogCoordd)(GLdouble coord);
// void (APIENTRYP ptrglFogCoorddv)(GLdouble* coord);
// void (APIENTRYP ptrglFogCoordPointer)(GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglSecondaryColor3b)(GLbyte red, GLbyte green, GLbyte blue);
// void (APIENTRYP ptrglSecondaryColor3bv)(GLbyte* v);
// void (APIENTRYP ptrglSecondaryColor3d)(GLdouble red, GLdouble green, GLdouble blue);
// void (APIENTRYP ptrglSecondaryColor3dv)(GLdouble* v);
// void (APIENTRYP ptrglSecondaryColor3f)(GLfloat red, GLfloat green, GLfloat blue);
// void (APIENTRYP ptrglSecondaryColor3fv)(GLfloat* v);
// void (APIENTRYP ptrglSecondaryColor3i)(GLint red, GLint green, GLint blue);
// void (APIENTRYP ptrglSecondaryColor3iv)(GLint* v);
// void (APIENTRYP ptrglSecondaryColor3s)(GLshort red, GLshort green, GLshort blue);
// void (APIENTRYP ptrglSecondaryColor3sv)(GLshort* v);
// void (APIENTRYP ptrglSecondaryColor3ub)(GLubyte red, GLubyte green, GLubyte blue);
// void (APIENTRYP ptrglSecondaryColor3ubv)(GLubyte* v);
// void (APIENTRYP ptrglSecondaryColor3ui)(GLuint red, GLuint green, GLuint blue);
// void (APIENTRYP ptrglSecondaryColor3uiv)(GLuint* v);
// void (APIENTRYP ptrglSecondaryColor3us)(GLushort red, GLushort green, GLushort blue);
// void (APIENTRYP ptrglSecondaryColor3usv)(GLushort* v);
// void (APIENTRYP ptrglSecondaryColorPointer)(GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglWindowPos2d)(GLdouble x, GLdouble y);
// void (APIENTRYP ptrglWindowPos2dv)(GLdouble* v);
// void (APIENTRYP ptrglWindowPos2f)(GLfloat x, GLfloat y);
// void (APIENTRYP ptrglWindowPos2fv)(GLfloat* v);
// void (APIENTRYP ptrglWindowPos2i)(GLint x, GLint y);
// void (APIENTRYP ptrglWindowPos2iv)(GLint* v);
// void (APIENTRYP ptrglWindowPos2s)(GLshort x, GLshort y);
// void (APIENTRYP ptrglWindowPos2sv)(GLshort* v);
// void (APIENTRYP ptrglWindowPos3d)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglWindowPos3dv)(GLdouble* v);
// void (APIENTRYP ptrglWindowPos3f)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglWindowPos3fv)(GLfloat* v);
// void (APIENTRYP ptrglWindowPos3i)(GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglWindowPos3iv)(GLint* v);
// void (APIENTRYP ptrglWindowPos3s)(GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrglWindowPos3sv)(GLshort* v);
// //  VERSION_1_5
// void (APIENTRYP ptrglGenQueries)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglDeleteQueries)(GLsizei n, GLuint* ids);
// GLboolean (APIENTRYP ptrglIsQuery)(GLuint id);
// void (APIENTRYP ptrglBeginQuery)(GLenum target, GLuint id);
// void (APIENTRYP ptrglEndQuery)(GLenum target);
// void (APIENTRYP ptrglGetQueryiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetQueryObjectiv)(GLuint id, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetQueryObjectuiv)(GLuint id, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglBindBuffer)(GLenum target, GLuint buffer);
// void (APIENTRYP ptrglDeleteBuffers)(GLsizei n, GLuint* buffers);
// void (APIENTRYP ptrglGenBuffers)(GLsizei n, GLuint* buffers);
// GLboolean (APIENTRYP ptrglIsBuffer)(GLuint buffer);
// void (APIENTRYP ptrglBufferData)(GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage);
// void (APIENTRYP ptrglBufferSubData)(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data);
// void (APIENTRYP ptrglGetBufferSubData)(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data);
// GLvoid* (APIENTRYP ptrglMapBuffer)(GLenum target, GLenum access);
// GLboolean (APIENTRYP ptrglUnmapBuffer)(GLenum target);
// void (APIENTRYP ptrglGetBufferParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetBufferPointerv)(GLenum target, GLenum pname, GLvoid** params);
// //  VERSION_2_0
// void (APIENTRYP ptrglBlendEquationSeparate)(GLenum modeRGB, GLenum modeAlpha);
// void (APIENTRYP ptrglDrawBuffers)(GLsizei n, GLenum* bufs);
// void (APIENTRYP ptrglStencilOpSeparate)(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass);
// void (APIENTRYP ptrglStencilFuncSeparate)(GLenum face, GLenum func, GLint ref, GLuint mask);
// void (APIENTRYP ptrglStencilMaskSeparate)(GLenum face, GLuint mask);
// void (APIENTRYP ptrglAttachShader)(GLuint program, GLuint shader);
// void (APIENTRYP ptrglBindAttribLocation)(GLuint program, GLuint index, GLchar* name);
// void (APIENTRYP ptrglCompileShader)(GLuint shader);
// GLuint (APIENTRYP ptrglCreateProgram)();
// GLuint (APIENTRYP ptrglCreateShader)(GLenum type);
// void (APIENTRYP ptrglDeleteProgram)(GLuint program);
// void (APIENTRYP ptrglDeleteShader)(GLuint shader);
// void (APIENTRYP ptrglDetachShader)(GLuint program, GLuint shader);
// void (APIENTRYP ptrglDisableVertexAttribArray)(GLuint index);
// void (APIENTRYP ptrglEnableVertexAttribArray)(GLuint index);
// void (APIENTRYP ptrglGetActiveAttrib)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrglGetActiveUniform)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrglGetAttachedShaders)(GLuint program, GLsizei maxCount, GLsizei* count, GLuint* obj);
// GLint (APIENTRYP ptrglGetAttribLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglGetProgramiv)(GLuint program, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetProgramInfoLog)(GLuint program, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrglGetShaderiv)(GLuint shader, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetShaderInfoLog)(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrglGetShaderSource)(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source);
// GLint (APIENTRYP ptrglGetUniformLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglGetUniformfv)(GLuint program, GLint location, GLfloat* params);
// void (APIENTRYP ptrglGetUniformiv)(GLuint program, GLint location, GLint* params);
// void (APIENTRYP ptrglGetVertexAttribdv)(GLuint index, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglGetVertexAttribfv)(GLuint index, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetVertexAttribiv)(GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetVertexAttribPointerv)(GLuint index, GLenum pname, GLvoid** pointer);
// GLboolean (APIENTRYP ptrglIsProgram)(GLuint program);
// GLboolean (APIENTRYP ptrglIsShader)(GLuint shader);
// void (APIENTRYP ptrglLinkProgram)(GLuint program);
// void (APIENTRYP ptrglShaderSource)(GLuint shader, GLsizei count, GLchar** string, GLint* length);
// void (APIENTRYP ptrglUseProgram)(GLuint program);
// void (APIENTRYP ptrglUniform1f)(GLint location, GLfloat v0);
// void (APIENTRYP ptrglUniform2f)(GLint location, GLfloat v0, GLfloat v1);
// void (APIENTRYP ptrglUniform3f)(GLint location, GLfloat v0, GLfloat v1, GLfloat v2);
// void (APIENTRYP ptrglUniform4f)(GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3);
// void (APIENTRYP ptrglUniform1i)(GLint location, GLint v0);
// void (APIENTRYP ptrglUniform2i)(GLint location, GLint v0, GLint v1);
// void (APIENTRYP ptrglUniform3i)(GLint location, GLint v0, GLint v1, GLint v2);
// void (APIENTRYP ptrglUniform4i)(GLint location, GLint v0, GLint v1, GLint v2, GLint v3);
// void (APIENTRYP ptrglUniform1fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglUniform2fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglUniform3fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglUniform4fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglUniform1iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglUniform2iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglUniform3iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglUniform4iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglUniformMatrix2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglUniformMatrix3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglUniformMatrix4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglValidateProgram)(GLuint program);
// void (APIENTRYP ptrglVertexAttrib1d)(GLuint index, GLdouble x);
// void (APIENTRYP ptrglVertexAttrib1dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttrib1f)(GLuint index, GLfloat x);
// void (APIENTRYP ptrglVertexAttrib1fv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglVertexAttrib1s)(GLuint index, GLshort x);
// void (APIENTRYP ptrglVertexAttrib1sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttrib2d)(GLuint index, GLdouble x, GLdouble y);
// void (APIENTRYP ptrglVertexAttrib2dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttrib2f)(GLuint index, GLfloat x, GLfloat y);
// void (APIENTRYP ptrglVertexAttrib2fv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglVertexAttrib2s)(GLuint index, GLshort x, GLshort y);
// void (APIENTRYP ptrglVertexAttrib2sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttrib3d)(GLuint index, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglVertexAttrib3dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttrib3f)(GLuint index, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglVertexAttrib3fv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglVertexAttrib3s)(GLuint index, GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrglVertexAttrib3sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttrib4Nbv)(GLuint index, GLbyte* v);
// void (APIENTRYP ptrglVertexAttrib4Niv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglVertexAttrib4Nsv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttrib4Nub)(GLuint index, GLubyte x, GLubyte y, GLubyte z, GLubyte w);
// void (APIENTRYP ptrglVertexAttrib4Nubv)(GLuint index, GLubyte* v);
// void (APIENTRYP ptrglVertexAttrib4Nuiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttrib4Nusv)(GLuint index, GLushort* v);
// void (APIENTRYP ptrglVertexAttrib4bv)(GLuint index, GLbyte* v);
// void (APIENTRYP ptrglVertexAttrib4d)(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglVertexAttrib4dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttrib4f)(GLuint index, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglVertexAttrib4fv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglVertexAttrib4iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglVertexAttrib4s)(GLuint index, GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrglVertexAttrib4sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttrib4ubv)(GLuint index, GLubyte* v);
// void (APIENTRYP ptrglVertexAttrib4uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttrib4usv)(GLuint index, GLushort* v);
// void (APIENTRYP ptrglVertexAttribPointer)(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer);
// //  VERSION_2_1
// void (APIENTRYP ptrglUniformMatrix2x3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglUniformMatrix3x2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglUniformMatrix2x4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglUniformMatrix4x2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglUniformMatrix3x4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglUniformMatrix4x3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// 
// //  VERSION_1_0
// void goglCullFace(GLenum mode) {
// 	(*ptrglCullFace)(mode);
// }
// void goglFrontFace(GLenum mode) {
// 	(*ptrglFrontFace)(mode);
// }
// void goglHint(GLenum target, GLenum mode) {
// 	(*ptrglHint)(target, mode);
// }
// void goglLineWidth(GLfloat width) {
// 	(*ptrglLineWidth)(width);
// }
// void goglPointSize(GLfloat size) {
// 	(*ptrglPointSize)(size);
// }
// void goglPolygonMode(GLenum face, GLenum mode) {
// 	(*ptrglPolygonMode)(face, mode);
// }
// void goglScissor(GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrglScissor)(x, y, width, height);
// }
// void goglTexParameterf(GLenum target, GLenum pname, GLfloat param) {
// 	(*ptrglTexParameterf)(target, pname, param);
// }
// void goglTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglTexParameterfv)(target, pname, params);
// }
// void goglTexParameteri(GLenum target, GLenum pname, GLint param) {
// 	(*ptrglTexParameteri)(target, pname, param);
// }
// void goglTexParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglTexParameteriv)(target, pname, params);
// }
// void goglTexImage1D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglTexImage1D)(target, level, internalformat, width, border, format, type_, pixels);
// }
// void goglTexImage2D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglTexImage2D)(target, level, internalformat, width, height, border, format, type_, pixels);
// }
// void goglDrawBuffer(GLenum mode) {
// 	(*ptrglDrawBuffer)(mode);
// }
// void goglClear(GLbitfield mask) {
// 	(*ptrglClear)(mask);
// }
// void goglClearColor(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
// 	(*ptrglClearColor)(red, green, blue, alpha);
// }
// void goglClearStencil(GLint s) {
// 	(*ptrglClearStencil)(s);
// }
// void goglClearDepth(GLclampd depth) {
// 	(*ptrglClearDepth)(depth);
// }
// void goglStencilMask(GLuint mask) {
// 	(*ptrglStencilMask)(mask);
// }
// void goglColorMask(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
// 	(*ptrglColorMask)(red, green, blue, alpha);
// }
// void goglDepthMask(GLboolean flag) {
// 	(*ptrglDepthMask)(flag);
// }
// void goglDisable(GLenum cap) {
// 	(*ptrglDisable)(cap);
// }
// void goglEnable(GLenum cap) {
// 	(*ptrglEnable)(cap);
// }
// void goglFinish() {
// 	(*ptrglFinish)();
// }
// void goglFlush() {
// 	(*ptrglFlush)();
// }
// void goglBlendFunc(GLenum sfactor, GLenum dfactor) {
// 	(*ptrglBlendFunc)(sfactor, dfactor);
// }
// void goglLogicOp(GLenum opcode) {
// 	(*ptrglLogicOp)(opcode);
// }
// void goglStencilFunc(GLenum func_, GLint ref, GLuint mask) {
// 	(*ptrglStencilFunc)(func_, ref, mask);
// }
// void goglStencilOp(GLenum fail, GLenum zfail, GLenum zpass) {
// 	(*ptrglStencilOp)(fail, zfail, zpass);
// }
// void goglDepthFunc(GLenum func_) {
// 	(*ptrglDepthFunc)(func_);
// }
// void goglPixelStoref(GLenum pname, GLfloat param) {
// 	(*ptrglPixelStoref)(pname, param);
// }
// void goglPixelStorei(GLenum pname, GLint param) {
// 	(*ptrglPixelStorei)(pname, param);
// }
// void goglReadBuffer(GLenum mode) {
// 	(*ptrglReadBuffer)(mode);
// }
// void goglReadPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglReadPixels)(x, y, width, height, format, type_, pixels);
// }
// void goglGetBooleanv(GLenum pname, GLboolean* params) {
// 	(*ptrglGetBooleanv)(pname, params);
// }
// void goglGetDoublev(GLenum pname, GLdouble* params) {
// 	(*ptrglGetDoublev)(pname, params);
// }
// GLenum goglGetError() {
// 	return (*ptrglGetError)();
// }
// void goglGetFloatv(GLenum pname, GLfloat* params) {
// 	(*ptrglGetFloatv)(pname, params);
// }
// void goglGetIntegerv(GLenum pname, GLint* params) {
// 	(*ptrglGetIntegerv)(pname, params);
// }
// const GLubyte * goglGetString(GLenum name) {
// 	return (*ptrglGetString)(name);
// }
// void goglGetTexImage(GLenum target, GLint level, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglGetTexImage)(target, level, format, type_, pixels);
// }
// void goglGetTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetTexParameterfv)(target, pname, params);
// }
// void goglGetTexParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetTexParameteriv)(target, pname, params);
// }
// void goglGetTexLevelParameterfv(GLenum target, GLint level, GLenum pname, GLfloat* params) {
// 	(*ptrglGetTexLevelParameterfv)(target, level, pname, params);
// }
// void goglGetTexLevelParameteriv(GLenum target, GLint level, GLenum pname, GLint* params) {
// 	(*ptrglGetTexLevelParameteriv)(target, level, pname, params);
// }
// GLboolean goglIsEnabled(GLenum cap) {
// 	return (*ptrglIsEnabled)(cap);
// }
// void goglDepthRange(GLclampd near_, GLclampd far_) {
// 	(*ptrglDepthRange)(near_, far_);
// }
// void goglViewport(GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrglViewport)(x, y, width, height);
// }
// //  VERSION_1_0_DEPRECATED
// void goglNewList(GLuint list, GLenum mode) {
// 	(*ptrglNewList)(list, mode);
// }
// void goglEndList() {
// 	(*ptrglEndList)();
// }
// void goglCallList(GLuint list) {
// 	(*ptrglCallList)(list);
// }
// void goglCallLists(GLsizei n, GLenum type_, GLvoid* lists) {
// 	(*ptrglCallLists)(n, type_, lists);
// }
// void goglDeleteLists(GLuint list, GLsizei range_) {
// 	(*ptrglDeleteLists)(list, range_);
// }
// GLuint goglGenLists(GLsizei range_) {
// 	return (*ptrglGenLists)(range_);
// }
// void goglListBase(GLuint base) {
// 	(*ptrglListBase)(base);
// }
// void goglBegin(GLenum mode) {
// 	(*ptrglBegin)(mode);
// }
// void goglBitmap(GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
// 	(*ptrglBitmap)(width, height, xorig, yorig, xmove, ymove, bitmap);
// }
// void goglColor3b(GLbyte red, GLbyte green, GLbyte blue) {
// 	(*ptrglColor3b)(red, green, blue);
// }
// void goglColor3bv(GLbyte* v) {
// 	(*ptrglColor3bv)(v);
// }
// void goglColor3d(GLdouble red, GLdouble green, GLdouble blue) {
// 	(*ptrglColor3d)(red, green, blue);
// }
// void goglColor3dv(GLdouble* v) {
// 	(*ptrglColor3dv)(v);
// }
// void goglColor3f(GLfloat red, GLfloat green, GLfloat blue) {
// 	(*ptrglColor3f)(red, green, blue);
// }
// void goglColor3fv(GLfloat* v) {
// 	(*ptrglColor3fv)(v);
// }
// void goglColor3i(GLint red, GLint green, GLint blue) {
// 	(*ptrglColor3i)(red, green, blue);
// }
// void goglColor3iv(GLint* v) {
// 	(*ptrglColor3iv)(v);
// }
// void goglColor3s(GLshort red, GLshort green, GLshort blue) {
// 	(*ptrglColor3s)(red, green, blue);
// }
// void goglColor3sv(GLshort* v) {
// 	(*ptrglColor3sv)(v);
// }
// void goglColor3ub(GLubyte red, GLubyte green, GLubyte blue) {
// 	(*ptrglColor3ub)(red, green, blue);
// }
// void goglColor3ubv(GLubyte* v) {
// 	(*ptrglColor3ubv)(v);
// }
// void goglColor3ui(GLuint red, GLuint green, GLuint blue) {
// 	(*ptrglColor3ui)(red, green, blue);
// }
// void goglColor3uiv(GLuint* v) {
// 	(*ptrglColor3uiv)(v);
// }
// void goglColor3us(GLushort red, GLushort green, GLushort blue) {
// 	(*ptrglColor3us)(red, green, blue);
// }
// void goglColor3usv(GLushort* v) {
// 	(*ptrglColor3usv)(v);
// }
// void goglColor4b(GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
// 	(*ptrglColor4b)(red, green, blue, alpha);
// }
// void goglColor4bv(GLbyte* v) {
// 	(*ptrglColor4bv)(v);
// }
// void goglColor4d(GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
// 	(*ptrglColor4d)(red, green, blue, alpha);
// }
// void goglColor4dv(GLdouble* v) {
// 	(*ptrglColor4dv)(v);
// }
// void goglColor4f(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
// 	(*ptrglColor4f)(red, green, blue, alpha);
// }
// void goglColor4fv(GLfloat* v) {
// 	(*ptrglColor4fv)(v);
// }
// void goglColor4i(GLint red, GLint green, GLint blue, GLint alpha) {
// 	(*ptrglColor4i)(red, green, blue, alpha);
// }
// void goglColor4iv(GLint* v) {
// 	(*ptrglColor4iv)(v);
// }
// void goglColor4s(GLshort red, GLshort green, GLshort blue, GLshort alpha) {
// 	(*ptrglColor4s)(red, green, blue, alpha);
// }
// void goglColor4sv(GLshort* v) {
// 	(*ptrglColor4sv)(v);
// }
// void goglColor4ub(GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
// 	(*ptrglColor4ub)(red, green, blue, alpha);
// }
// void goglColor4ubv(GLubyte* v) {
// 	(*ptrglColor4ubv)(v);
// }
// void goglColor4ui(GLuint red, GLuint green, GLuint blue, GLuint alpha) {
// 	(*ptrglColor4ui)(red, green, blue, alpha);
// }
// void goglColor4uiv(GLuint* v) {
// 	(*ptrglColor4uiv)(v);
// }
// void goglColor4us(GLushort red, GLushort green, GLushort blue, GLushort alpha) {
// 	(*ptrglColor4us)(red, green, blue, alpha);
// }
// void goglColor4usv(GLushort* v) {
// 	(*ptrglColor4usv)(v);
// }
// void goglEdgeFlag(GLboolean flag) {
// 	(*ptrglEdgeFlag)(flag);
// }
// void goglEdgeFlagv(GLboolean* flag) {
// 	(*ptrglEdgeFlagv)(flag);
// }
// void goglEnd() {
// 	(*ptrglEnd)();
// }
// void goglIndexd(GLdouble c) {
// 	(*ptrglIndexd)(c);
// }
// void goglIndexdv(GLdouble* c) {
// 	(*ptrglIndexdv)(c);
// }
// void goglIndexf(GLfloat c) {
// 	(*ptrglIndexf)(c);
// }
// void goglIndexfv(GLfloat* c) {
// 	(*ptrglIndexfv)(c);
// }
// void goglIndexi(GLint c) {
// 	(*ptrglIndexi)(c);
// }
// void goglIndexiv(GLint* c) {
// 	(*ptrglIndexiv)(c);
// }
// void goglIndexs(GLshort c) {
// 	(*ptrglIndexs)(c);
// }
// void goglIndexsv(GLshort* c) {
// 	(*ptrglIndexsv)(c);
// }
// void goglNormal3b(GLbyte nx, GLbyte ny, GLbyte nz) {
// 	(*ptrglNormal3b)(nx, ny, nz);
// }
// void goglNormal3bv(GLbyte* v) {
// 	(*ptrglNormal3bv)(v);
// }
// void goglNormal3d(GLdouble nx, GLdouble ny, GLdouble nz) {
// 	(*ptrglNormal3d)(nx, ny, nz);
// }
// void goglNormal3dv(GLdouble* v) {
// 	(*ptrglNormal3dv)(v);
// }
// void goglNormal3f(GLfloat nx, GLfloat ny, GLfloat nz) {
// 	(*ptrglNormal3f)(nx, ny, nz);
// }
// void goglNormal3fv(GLfloat* v) {
// 	(*ptrglNormal3fv)(v);
// }
// void goglNormal3i(GLint nx, GLint ny, GLint nz) {
// 	(*ptrglNormal3i)(nx, ny, nz);
// }
// void goglNormal3iv(GLint* v) {
// 	(*ptrglNormal3iv)(v);
// }
// void goglNormal3s(GLshort nx, GLshort ny, GLshort nz) {
// 	(*ptrglNormal3s)(nx, ny, nz);
// }
// void goglNormal3sv(GLshort* v) {
// 	(*ptrglNormal3sv)(v);
// }
// void goglRasterPos2d(GLdouble x, GLdouble y) {
// 	(*ptrglRasterPos2d)(x, y);
// }
// void goglRasterPos2dv(GLdouble* v) {
// 	(*ptrglRasterPos2dv)(v);
// }
// void goglRasterPos2f(GLfloat x, GLfloat y) {
// 	(*ptrglRasterPos2f)(x, y);
// }
// void goglRasterPos2fv(GLfloat* v) {
// 	(*ptrglRasterPos2fv)(v);
// }
// void goglRasterPos2i(GLint x, GLint y) {
// 	(*ptrglRasterPos2i)(x, y);
// }
// void goglRasterPos2iv(GLint* v) {
// 	(*ptrglRasterPos2iv)(v);
// }
// void goglRasterPos2s(GLshort x, GLshort y) {
// 	(*ptrglRasterPos2s)(x, y);
// }
// void goglRasterPos2sv(GLshort* v) {
// 	(*ptrglRasterPos2sv)(v);
// }
// void goglRasterPos3d(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglRasterPos3d)(x, y, z);
// }
// void goglRasterPos3dv(GLdouble* v) {
// 	(*ptrglRasterPos3dv)(v);
// }
// void goglRasterPos3f(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglRasterPos3f)(x, y, z);
// }
// void goglRasterPos3fv(GLfloat* v) {
// 	(*ptrglRasterPos3fv)(v);
// }
// void goglRasterPos3i(GLint x, GLint y, GLint z) {
// 	(*ptrglRasterPos3i)(x, y, z);
// }
// void goglRasterPos3iv(GLint* v) {
// 	(*ptrglRasterPos3iv)(v);
// }
// void goglRasterPos3s(GLshort x, GLshort y, GLshort z) {
// 	(*ptrglRasterPos3s)(x, y, z);
// }
// void goglRasterPos3sv(GLshort* v) {
// 	(*ptrglRasterPos3sv)(v);
// }
// void goglRasterPos4d(GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglRasterPos4d)(x, y, z, w);
// }
// void goglRasterPos4dv(GLdouble* v) {
// 	(*ptrglRasterPos4dv)(v);
// }
// void goglRasterPos4f(GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglRasterPos4f)(x, y, z, w);
// }
// void goglRasterPos4fv(GLfloat* v) {
// 	(*ptrglRasterPos4fv)(v);
// }
// void goglRasterPos4i(GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrglRasterPos4i)(x, y, z, w);
// }
// void goglRasterPos4iv(GLint* v) {
// 	(*ptrglRasterPos4iv)(v);
// }
// void goglRasterPos4s(GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrglRasterPos4s)(x, y, z, w);
// }
// void goglRasterPos4sv(GLshort* v) {
// 	(*ptrglRasterPos4sv)(v);
// }
// void goglRectd(GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
// 	(*ptrglRectd)(x1, y1, x2, y2);
// }
// void goglRectdv(GLdouble* v1, GLdouble* v2) {
// 	(*ptrglRectdv)(v1, v2);
// }
// void goglRectf(GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
// 	(*ptrglRectf)(x1, y1, x2, y2);
// }
// void goglRectfv(GLfloat* v1, GLfloat* v2) {
// 	(*ptrglRectfv)(v1, v2);
// }
// void goglRecti(GLint x1, GLint y1, GLint x2, GLint y2) {
// 	(*ptrglRecti)(x1, y1, x2, y2);
// }
// void goglRectiv(GLint* v1, GLint* v2) {
// 	(*ptrglRectiv)(v1, v2);
// }
// void goglRects(GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
// 	(*ptrglRects)(x1, y1, x2, y2);
// }
// void goglRectsv(GLshort* v1, GLshort* v2) {
// 	(*ptrglRectsv)(v1, v2);
// }
// void goglTexCoord1d(GLdouble s) {
// 	(*ptrglTexCoord1d)(s);
// }
// void goglTexCoord1dv(GLdouble* v) {
// 	(*ptrglTexCoord1dv)(v);
// }
// void goglTexCoord1f(GLfloat s) {
// 	(*ptrglTexCoord1f)(s);
// }
// void goglTexCoord1fv(GLfloat* v) {
// 	(*ptrglTexCoord1fv)(v);
// }
// void goglTexCoord1i(GLint s) {
// 	(*ptrglTexCoord1i)(s);
// }
// void goglTexCoord1iv(GLint* v) {
// 	(*ptrglTexCoord1iv)(v);
// }
// void goglTexCoord1s(GLshort s) {
// 	(*ptrglTexCoord1s)(s);
// }
// void goglTexCoord1sv(GLshort* v) {
// 	(*ptrglTexCoord1sv)(v);
// }
// void goglTexCoord2d(GLdouble s, GLdouble t) {
// 	(*ptrglTexCoord2d)(s, t);
// }
// void goglTexCoord2dv(GLdouble* v) {
// 	(*ptrglTexCoord2dv)(v);
// }
// void goglTexCoord2f(GLfloat s, GLfloat t) {
// 	(*ptrglTexCoord2f)(s, t);
// }
// void goglTexCoord2fv(GLfloat* v) {
// 	(*ptrglTexCoord2fv)(v);
// }
// void goglTexCoord2i(GLint s, GLint t) {
// 	(*ptrglTexCoord2i)(s, t);
// }
// void goglTexCoord2iv(GLint* v) {
// 	(*ptrglTexCoord2iv)(v);
// }
// void goglTexCoord2s(GLshort s, GLshort t) {
// 	(*ptrglTexCoord2s)(s, t);
// }
// void goglTexCoord2sv(GLshort* v) {
// 	(*ptrglTexCoord2sv)(v);
// }
// void goglTexCoord3d(GLdouble s, GLdouble t, GLdouble r) {
// 	(*ptrglTexCoord3d)(s, t, r);
// }
// void goglTexCoord3dv(GLdouble* v) {
// 	(*ptrglTexCoord3dv)(v);
// }
// void goglTexCoord3f(GLfloat s, GLfloat t, GLfloat r) {
// 	(*ptrglTexCoord3f)(s, t, r);
// }
// void goglTexCoord3fv(GLfloat* v) {
// 	(*ptrglTexCoord3fv)(v);
// }
// void goglTexCoord3i(GLint s, GLint t, GLint r) {
// 	(*ptrglTexCoord3i)(s, t, r);
// }
// void goglTexCoord3iv(GLint* v) {
// 	(*ptrglTexCoord3iv)(v);
// }
// void goglTexCoord3s(GLshort s, GLshort t, GLshort r) {
// 	(*ptrglTexCoord3s)(s, t, r);
// }
// void goglTexCoord3sv(GLshort* v) {
// 	(*ptrglTexCoord3sv)(v);
// }
// void goglTexCoord4d(GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
// 	(*ptrglTexCoord4d)(s, t, r, q);
// }
// void goglTexCoord4dv(GLdouble* v) {
// 	(*ptrglTexCoord4dv)(v);
// }
// void goglTexCoord4f(GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
// 	(*ptrglTexCoord4f)(s, t, r, q);
// }
// void goglTexCoord4fv(GLfloat* v) {
// 	(*ptrglTexCoord4fv)(v);
// }
// void goglTexCoord4i(GLint s, GLint t, GLint r, GLint q) {
// 	(*ptrglTexCoord4i)(s, t, r, q);
// }
// void goglTexCoord4iv(GLint* v) {
// 	(*ptrglTexCoord4iv)(v);
// }
// void goglTexCoord4s(GLshort s, GLshort t, GLshort r, GLshort q) {
// 	(*ptrglTexCoord4s)(s, t, r, q);
// }
// void goglTexCoord4sv(GLshort* v) {
// 	(*ptrglTexCoord4sv)(v);
// }
// void goglVertex2d(GLdouble x, GLdouble y) {
// 	(*ptrglVertex2d)(x, y);
// }
// void goglVertex2dv(GLdouble* v) {
// 	(*ptrglVertex2dv)(v);
// }
// void goglVertex2f(GLfloat x, GLfloat y) {
// 	(*ptrglVertex2f)(x, y);
// }
// void goglVertex2fv(GLfloat* v) {
// 	(*ptrglVertex2fv)(v);
// }
// void goglVertex2i(GLint x, GLint y) {
// 	(*ptrglVertex2i)(x, y);
// }
// void goglVertex2iv(GLint* v) {
// 	(*ptrglVertex2iv)(v);
// }
// void goglVertex2s(GLshort x, GLshort y) {
// 	(*ptrglVertex2s)(x, y);
// }
// void goglVertex2sv(GLshort* v) {
// 	(*ptrglVertex2sv)(v);
// }
// void goglVertex3d(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglVertex3d)(x, y, z);
// }
// void goglVertex3dv(GLdouble* v) {
// 	(*ptrglVertex3dv)(v);
// }
// void goglVertex3f(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglVertex3f)(x, y, z);
// }
// void goglVertex3fv(GLfloat* v) {
// 	(*ptrglVertex3fv)(v);
// }
// void goglVertex3i(GLint x, GLint y, GLint z) {
// 	(*ptrglVertex3i)(x, y, z);
// }
// void goglVertex3iv(GLint* v) {
// 	(*ptrglVertex3iv)(v);
// }
// void goglVertex3s(GLshort x, GLshort y, GLshort z) {
// 	(*ptrglVertex3s)(x, y, z);
// }
// void goglVertex3sv(GLshort* v) {
// 	(*ptrglVertex3sv)(v);
// }
// void goglVertex4d(GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglVertex4d)(x, y, z, w);
// }
// void goglVertex4dv(GLdouble* v) {
// 	(*ptrglVertex4dv)(v);
// }
// void goglVertex4f(GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglVertex4f)(x, y, z, w);
// }
// void goglVertex4fv(GLfloat* v) {
// 	(*ptrglVertex4fv)(v);
// }
// void goglVertex4i(GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrglVertex4i)(x, y, z, w);
// }
// void goglVertex4iv(GLint* v) {
// 	(*ptrglVertex4iv)(v);
// }
// void goglVertex4s(GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrglVertex4s)(x, y, z, w);
// }
// void goglVertex4sv(GLshort* v) {
// 	(*ptrglVertex4sv)(v);
// }
// void goglClipPlane(GLenum plane, GLdouble* equation) {
// 	(*ptrglClipPlane)(plane, equation);
// }
// void goglColorMaterial(GLenum face, GLenum mode) {
// 	(*ptrglColorMaterial)(face, mode);
// }
// void goglFogf(GLenum pname, GLfloat param) {
// 	(*ptrglFogf)(pname, param);
// }
// void goglFogfv(GLenum pname, GLfloat* params) {
// 	(*ptrglFogfv)(pname, params);
// }
// void goglFogi(GLenum pname, GLint param) {
// 	(*ptrglFogi)(pname, param);
// }
// void goglFogiv(GLenum pname, GLint* params) {
// 	(*ptrglFogiv)(pname, params);
// }
// void goglLightf(GLenum light, GLenum pname, GLfloat param) {
// 	(*ptrglLightf)(light, pname, param);
// }
// void goglLightfv(GLenum light, GLenum pname, GLfloat* params) {
// 	(*ptrglLightfv)(light, pname, params);
// }
// void goglLighti(GLenum light, GLenum pname, GLint param) {
// 	(*ptrglLighti)(light, pname, param);
// }
// void goglLightiv(GLenum light, GLenum pname, GLint* params) {
// 	(*ptrglLightiv)(light, pname, params);
// }
// void goglLightModelf(GLenum pname, GLfloat param) {
// 	(*ptrglLightModelf)(pname, param);
// }
// void goglLightModelfv(GLenum pname, GLfloat* params) {
// 	(*ptrglLightModelfv)(pname, params);
// }
// void goglLightModeli(GLenum pname, GLint param) {
// 	(*ptrglLightModeli)(pname, param);
// }
// void goglLightModeliv(GLenum pname, GLint* params) {
// 	(*ptrglLightModeliv)(pname, params);
// }
// void goglLineStipple(GLint factor, GLushort pattern) {
// 	(*ptrglLineStipple)(factor, pattern);
// }
// void goglMaterialf(GLenum face, GLenum pname, GLfloat param) {
// 	(*ptrglMaterialf)(face, pname, param);
// }
// void goglMaterialfv(GLenum face, GLenum pname, GLfloat* params) {
// 	(*ptrglMaterialfv)(face, pname, params);
// }
// void goglMateriali(GLenum face, GLenum pname, GLint param) {
// 	(*ptrglMateriali)(face, pname, param);
// }
// void goglMaterialiv(GLenum face, GLenum pname, GLint* params) {
// 	(*ptrglMaterialiv)(face, pname, params);
// }
// void goglPolygonStipple(GLubyte* mask) {
// 	(*ptrglPolygonStipple)(mask);
// }
// void goglShadeModel(GLenum mode) {
// 	(*ptrglShadeModel)(mode);
// }
// void goglTexEnvf(GLenum target, GLenum pname, GLfloat param) {
// 	(*ptrglTexEnvf)(target, pname, param);
// }
// void goglTexEnvfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglTexEnvfv)(target, pname, params);
// }
// void goglTexEnvi(GLenum target, GLenum pname, GLint param) {
// 	(*ptrglTexEnvi)(target, pname, param);
// }
// void goglTexEnviv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglTexEnviv)(target, pname, params);
// }
// void goglTexGend(GLenum coord, GLenum pname, GLdouble param) {
// 	(*ptrglTexGend)(coord, pname, param);
// }
// void goglTexGendv(GLenum coord, GLenum pname, GLdouble* params) {
// 	(*ptrglTexGendv)(coord, pname, params);
// }
// void goglTexGenf(GLenum coord, GLenum pname, GLfloat param) {
// 	(*ptrglTexGenf)(coord, pname, param);
// }
// void goglTexGenfv(GLenum coord, GLenum pname, GLfloat* params) {
// 	(*ptrglTexGenfv)(coord, pname, params);
// }
// void goglTexGeni(GLenum coord, GLenum pname, GLint param) {
// 	(*ptrglTexGeni)(coord, pname, param);
// }
// void goglTexGeniv(GLenum coord, GLenum pname, GLint* params) {
// 	(*ptrglTexGeniv)(coord, pname, params);
// }
// void goglFeedbackBuffer(GLsizei size, GLenum type_, GLfloat* buffer) {
// 	(*ptrglFeedbackBuffer)(size, type_, buffer);
// }
// void goglSelectBuffer(GLsizei size, GLuint* buffer) {
// 	(*ptrglSelectBuffer)(size, buffer);
// }
// GLint goglRenderMode(GLenum mode) {
// 	return (*ptrglRenderMode)(mode);
// }
// void goglInitNames() {
// 	(*ptrglInitNames)();
// }
// void goglLoadName(GLuint name) {
// 	(*ptrglLoadName)(name);
// }
// void goglPassThrough(GLfloat token) {
// 	(*ptrglPassThrough)(token);
// }
// void goglPopName() {
// 	(*ptrglPopName)();
// }
// void goglPushName(GLuint name) {
// 	(*ptrglPushName)(name);
// }
// void goglClearAccum(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
// 	(*ptrglClearAccum)(red, green, blue, alpha);
// }
// void goglClearIndex(GLfloat c) {
// 	(*ptrglClearIndex)(c);
// }
// void goglIndexMask(GLuint mask) {
// 	(*ptrglIndexMask)(mask);
// }
// void goglAccum(GLenum op, GLfloat value) {
// 	(*ptrglAccum)(op, value);
// }
// void goglPopAttrib() {
// 	(*ptrglPopAttrib)();
// }
// void goglPushAttrib(GLbitfield mask) {
// 	(*ptrglPushAttrib)(mask);
// }
// void goglMap1d(GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
// 	(*ptrglMap1d)(target, u1, u2, stride, order, points);
// }
// void goglMap1f(GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
// 	(*ptrglMap1f)(target, u1, u2, stride, order, points);
// }
// void goglMap2d(GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
// 	(*ptrglMap2d)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// void goglMap2f(GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
// 	(*ptrglMap2f)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// void goglMapGrid1d(GLint un, GLdouble u1, GLdouble u2) {
// 	(*ptrglMapGrid1d)(un, u1, u2);
// }
// void goglMapGrid1f(GLint un, GLfloat u1, GLfloat u2) {
// 	(*ptrglMapGrid1f)(un, u1, u2);
// }
// void goglMapGrid2d(GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
// 	(*ptrglMapGrid2d)(un, u1, u2, vn, v1, v2);
// }
// void goglMapGrid2f(GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
// 	(*ptrglMapGrid2f)(un, u1, u2, vn, v1, v2);
// }
// void goglEvalCoord1d(GLdouble u) {
// 	(*ptrglEvalCoord1d)(u);
// }
// void goglEvalCoord1dv(GLdouble* u) {
// 	(*ptrglEvalCoord1dv)(u);
// }
// void goglEvalCoord1f(GLfloat u) {
// 	(*ptrglEvalCoord1f)(u);
// }
// void goglEvalCoord1fv(GLfloat* u) {
// 	(*ptrglEvalCoord1fv)(u);
// }
// void goglEvalCoord2d(GLdouble u, GLdouble v) {
// 	(*ptrglEvalCoord2d)(u, v);
// }
// void goglEvalCoord2dv(GLdouble* u) {
// 	(*ptrglEvalCoord2dv)(u);
// }
// void goglEvalCoord2f(GLfloat u, GLfloat v) {
// 	(*ptrglEvalCoord2f)(u, v);
// }
// void goglEvalCoord2fv(GLfloat* u) {
// 	(*ptrglEvalCoord2fv)(u);
// }
// void goglEvalMesh1(GLenum mode, GLint i1, GLint i2) {
// 	(*ptrglEvalMesh1)(mode, i1, i2);
// }
// void goglEvalPoint1(GLint i) {
// 	(*ptrglEvalPoint1)(i);
// }
// void goglEvalMesh2(GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
// 	(*ptrglEvalMesh2)(mode, i1, i2, j1, j2);
// }
// void goglEvalPoint2(GLint i, GLint j) {
// 	(*ptrglEvalPoint2)(i, j);
// }
// void goglAlphaFunc(GLenum func_, GLclampf ref) {
// 	(*ptrglAlphaFunc)(func_, ref);
// }
// void goglPixelZoom(GLfloat xfactor, GLfloat yfactor) {
// 	(*ptrglPixelZoom)(xfactor, yfactor);
// }
// void goglPixelTransferf(GLenum pname, GLfloat param) {
// 	(*ptrglPixelTransferf)(pname, param);
// }
// void goglPixelTransferi(GLenum pname, GLint param) {
// 	(*ptrglPixelTransferi)(pname, param);
// }
// void goglPixelMapfv(GLenum map_, GLsizei mapsize, GLfloat* values) {
// 	(*ptrglPixelMapfv)(map_, mapsize, values);
// }
// void goglPixelMapuiv(GLenum map_, GLsizei mapsize, GLuint* values) {
// 	(*ptrglPixelMapuiv)(map_, mapsize, values);
// }
// void goglPixelMapusv(GLenum map_, GLsizei mapsize, GLushort* values) {
// 	(*ptrglPixelMapusv)(map_, mapsize, values);
// }
// void goglCopyPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum type_) {
// 	(*ptrglCopyPixels)(x, y, width, height, type_);
// }
// void goglDrawPixels(GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglDrawPixels)(width, height, format, type_, pixels);
// }
// void goglGetClipPlane(GLenum plane, GLdouble* equation) {
// 	(*ptrglGetClipPlane)(plane, equation);
// }
// void goglGetLightfv(GLenum light, GLenum pname, GLfloat* params) {
// 	(*ptrglGetLightfv)(light, pname, params);
// }
// void goglGetLightiv(GLenum light, GLenum pname, GLint* params) {
// 	(*ptrglGetLightiv)(light, pname, params);
// }
// void goglGetMapdv(GLenum target, GLenum query, GLdouble* v) {
// 	(*ptrglGetMapdv)(target, query, v);
// }
// void goglGetMapfv(GLenum target, GLenum query, GLfloat* v) {
// 	(*ptrglGetMapfv)(target, query, v);
// }
// void goglGetMapiv(GLenum target, GLenum query, GLint* v) {
// 	(*ptrglGetMapiv)(target, query, v);
// }
// void goglGetMaterialfv(GLenum face, GLenum pname, GLfloat* params) {
// 	(*ptrglGetMaterialfv)(face, pname, params);
// }
// void goglGetMaterialiv(GLenum face, GLenum pname, GLint* params) {
// 	(*ptrglGetMaterialiv)(face, pname, params);
// }
// void goglGetPixelMapfv(GLenum map_, GLfloat* values) {
// 	(*ptrglGetPixelMapfv)(map_, values);
// }
// void goglGetPixelMapuiv(GLenum map_, GLuint* values) {
// 	(*ptrglGetPixelMapuiv)(map_, values);
// }
// void goglGetPixelMapusv(GLenum map_, GLushort* values) {
// 	(*ptrglGetPixelMapusv)(map_, values);
// }
// void goglGetPolygonStipple(GLubyte* mask) {
// 	(*ptrglGetPolygonStipple)(mask);
// }
// void goglGetTexEnvfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetTexEnvfv)(target, pname, params);
// }
// void goglGetTexEnviv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetTexEnviv)(target, pname, params);
// }
// void goglGetTexGendv(GLenum coord, GLenum pname, GLdouble* params) {
// 	(*ptrglGetTexGendv)(coord, pname, params);
// }
// void goglGetTexGenfv(GLenum coord, GLenum pname, GLfloat* params) {
// 	(*ptrglGetTexGenfv)(coord, pname, params);
// }
// void goglGetTexGeniv(GLenum coord, GLenum pname, GLint* params) {
// 	(*ptrglGetTexGeniv)(coord, pname, params);
// }
// GLboolean goglIsList(GLuint list) {
// 	return (*ptrglIsList)(list);
// }
// void goglFrustum(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
// 	(*ptrglFrustum)(left, right, bottom, top, zNear, zFar);
// }
// void goglLoadIdentity() {
// 	(*ptrglLoadIdentity)();
// }
// void goglLoadMatrixf(GLfloat* m) {
// 	(*ptrglLoadMatrixf)(m);
// }
// void goglLoadMatrixd(GLdouble* m) {
// 	(*ptrglLoadMatrixd)(m);
// }
// void goglMatrixMode(GLenum mode) {
// 	(*ptrglMatrixMode)(mode);
// }
// void goglMultMatrixf(GLfloat* m) {
// 	(*ptrglMultMatrixf)(m);
// }
// void goglMultMatrixd(GLdouble* m) {
// 	(*ptrglMultMatrixd)(m);
// }
// void goglOrtho(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
// 	(*ptrglOrtho)(left, right, bottom, top, zNear, zFar);
// }
// void goglPopMatrix() {
// 	(*ptrglPopMatrix)();
// }
// void goglPushMatrix() {
// 	(*ptrglPushMatrix)();
// }
// void goglRotated(GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglRotated)(angle, x, y, z);
// }
// void goglRotatef(GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglRotatef)(angle, x, y, z);
// }
// void goglScaled(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglScaled)(x, y, z);
// }
// void goglScalef(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglScalef)(x, y, z);
// }
// void goglTranslated(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglTranslated)(x, y, z);
// }
// void goglTranslatef(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglTranslatef)(x, y, z);
// }
// //  VERSION_1_1
// void goglDrawArrays(GLenum mode, GLint first, GLsizei count) {
// 	(*ptrglDrawArrays)(mode, first, count);
// }
// void goglDrawElements(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices) {
// 	(*ptrglDrawElements)(mode, count, type_, indices);
// }
// void goglGetPointerv(GLenum pname, GLvoid** params) {
// 	(*ptrglGetPointerv)(pname, params);
// }
// void goglPolygonOffset(GLfloat factor, GLfloat units) {
// 	(*ptrglPolygonOffset)(factor, units);
// }
// void goglCopyTexImage1D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border) {
// 	(*ptrglCopyTexImage1D)(target, level, internalformat, x, y, width, border);
// }
// void goglCopyTexImage2D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
// 	(*ptrglCopyTexImage2D)(target, level, internalformat, x, y, width, height, border);
// }
// void goglCopyTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
// 	(*ptrglCopyTexSubImage1D)(target, level, xoffset, x, y, width);
// }
// void goglCopyTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrglCopyTexSubImage2D)(target, level, xoffset, yoffset, x, y, width, height);
// }
// void goglTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglTexSubImage1D)(target, level, xoffset, width, format, type_, pixels);
// }
// void goglTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglTexSubImage2D)(target, level, xoffset, yoffset, width, height, format, type_, pixels);
// }
// void goglBindTexture(GLenum target, GLuint texture) {
// 	(*ptrglBindTexture)(target, texture);
// }
// void goglDeleteTextures(GLsizei n, GLuint* textures) {
// 	(*ptrglDeleteTextures)(n, textures);
// }
// void goglGenTextures(GLsizei n, GLuint* textures) {
// 	(*ptrglGenTextures)(n, textures);
// }
// GLboolean goglIsTexture(GLuint texture) {
// 	return (*ptrglIsTexture)(texture);
// }
// //  VERSION_1_1_DEPRECATED
// void goglArrayElement(GLint i) {
// 	(*ptrglArrayElement)(i);
// }
// void goglColorPointer(GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglColorPointer)(size, type_, stride, pointer);
// }
// void goglDisableClientState(GLenum array) {
// 	(*ptrglDisableClientState)(array);
// }
// void goglEdgeFlagPointer(GLsizei stride, GLvoid* pointer) {
// 	(*ptrglEdgeFlagPointer)(stride, pointer);
// }
// void goglEnableClientState(GLenum array) {
// 	(*ptrglEnableClientState)(array);
// }
// void goglIndexPointer(GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglIndexPointer)(type_, stride, pointer);
// }
// void goglInterleavedArrays(GLenum format, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglInterleavedArrays)(format, stride, pointer);
// }
// void goglNormalPointer(GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglNormalPointer)(type_, stride, pointer);
// }
// void goglTexCoordPointer(GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglTexCoordPointer)(size, type_, stride, pointer);
// }
// void goglVertexPointer(GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglVertexPointer)(size, type_, stride, pointer);
// }
// GLboolean goglAreTexturesResident(GLsizei n, GLuint* textures, GLboolean* residences) {
// 	return (*ptrglAreTexturesResident)(n, textures, residences);
// }
// void goglPrioritizeTextures(GLsizei n, GLuint* textures, GLclampf* priorities) {
// 	(*ptrglPrioritizeTextures)(n, textures, priorities);
// }
// void goglIndexub(GLubyte c) {
// 	(*ptrglIndexub)(c);
// }
// void goglIndexubv(GLubyte* c) {
// 	(*ptrglIndexubv)(c);
// }
// void goglPopClientAttrib() {
// 	(*ptrglPopClientAttrib)();
// }
// void goglPushClientAttrib(GLbitfield mask) {
// 	(*ptrglPushClientAttrib)(mask);
// }
// //  VERSION_1_2
// void goglBlendColor(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
// 	(*ptrglBlendColor)(red, green, blue, alpha);
// }
// void goglBlendEquation(GLenum mode) {
// 	(*ptrglBlendEquation)(mode);
// }
// void goglDrawRangeElements(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type_, GLvoid* indices) {
// 	(*ptrglDrawRangeElements)(mode, start, end, count, type_, indices);
// }
// void goglTexImage3D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglTexImage3D)(target, level, internalformat, width, height, depth, border, format, type_, pixels);
// }
// void goglTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglTexSubImage3D)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type_, pixels);
// }
// void goglCopyTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrglCopyTexSubImage3D)(target, level, xoffset, yoffset, zoffset, x, y, width, height);
// }
// //  VERSION_1_2_DEPRECATED
// void goglColorTable(GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type_, GLvoid* table) {
// 	(*ptrglColorTable)(target, internalformat, width, format, type_, table);
// }
// void goglColorTableParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglColorTableParameterfv)(target, pname, params);
// }
// void goglColorTableParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglColorTableParameteriv)(target, pname, params);
// }
// void goglCopyColorTable(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
// 	(*ptrglCopyColorTable)(target, internalformat, x, y, width);
// }
// void goglGetColorTable(GLenum target, GLenum format, GLenum type_, GLvoid* table) {
// 	(*ptrglGetColorTable)(target, format, type_, table);
// }
// void goglGetColorTableParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetColorTableParameterfv)(target, pname, params);
// }
// void goglGetColorTableParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetColorTableParameteriv)(target, pname, params);
// }
// void goglColorSubTable(GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type_, GLvoid* data) {
// 	(*ptrglColorSubTable)(target, start, count, format, type_, data);
// }
// void goglCopyColorSubTable(GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
// 	(*ptrglCopyColorSubTable)(target, start, x, y, width);
// }
// void goglConvolutionFilter1D(GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type_, GLvoid* image) {
// 	(*ptrglConvolutionFilter1D)(target, internalformat, width, format, type_, image);
// }
// void goglConvolutionFilter2D(GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* image) {
// 	(*ptrglConvolutionFilter2D)(target, internalformat, width, height, format, type_, image);
// }
// void goglConvolutionParameterf(GLenum target, GLenum pname, GLfloat params) {
// 	(*ptrglConvolutionParameterf)(target, pname, params);
// }
// void goglConvolutionParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglConvolutionParameterfv)(target, pname, params);
// }
// void goglConvolutionParameteri(GLenum target, GLenum pname, GLint params) {
// 	(*ptrglConvolutionParameteri)(target, pname, params);
// }
// void goglConvolutionParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglConvolutionParameteriv)(target, pname, params);
// }
// void goglCopyConvolutionFilter1D(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
// 	(*ptrglCopyConvolutionFilter1D)(target, internalformat, x, y, width);
// }
// void goglCopyConvolutionFilter2D(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrglCopyConvolutionFilter2D)(target, internalformat, x, y, width, height);
// }
// void goglGetConvolutionFilter(GLenum target, GLenum format, GLenum type_, GLvoid* image) {
// 	(*ptrglGetConvolutionFilter)(target, format, type_, image);
// }
// void goglGetConvolutionParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetConvolutionParameterfv)(target, pname, params);
// }
// void goglGetConvolutionParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetConvolutionParameteriv)(target, pname, params);
// }
// void goglGetSeparableFilter(GLenum target, GLenum format, GLenum type_, GLvoid* row, GLvoid* column, GLvoid* span) {
// 	(*ptrglGetSeparableFilter)(target, format, type_, row, column, span);
// }
// void goglSeparableFilter2D(GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* row, GLvoid* column) {
// 	(*ptrglSeparableFilter2D)(target, internalformat, width, height, format, type_, row, column);
// }
// void goglGetHistogram(GLenum target, GLboolean reset, GLenum format, GLenum type_, GLvoid* values) {
// 	(*ptrglGetHistogram)(target, reset, format, type_, values);
// }
// void goglGetHistogramParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetHistogramParameterfv)(target, pname, params);
// }
// void goglGetHistogramParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetHistogramParameteriv)(target, pname, params);
// }
// void goglGetMinmax(GLenum target, GLboolean reset, GLenum format, GLenum type_, GLvoid* values) {
// 	(*ptrglGetMinmax)(target, reset, format, type_, values);
// }
// void goglGetMinmaxParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetMinmaxParameterfv)(target, pname, params);
// }
// void goglGetMinmaxParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetMinmaxParameteriv)(target, pname, params);
// }
// void goglHistogram(GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
// 	(*ptrglHistogram)(target, width, internalformat, sink);
// }
// void goglMinmax(GLenum target, GLenum internalformat, GLboolean sink) {
// 	(*ptrglMinmax)(target, internalformat, sink);
// }
// void goglResetHistogram(GLenum target) {
// 	(*ptrglResetHistogram)(target);
// }
// void goglResetMinmax(GLenum target) {
// 	(*ptrglResetMinmax)(target);
// }
// //  VERSION_1_3
// void goglActiveTexture(GLenum texture) {
// 	(*ptrglActiveTexture)(texture);
// }
// void goglSampleCoverage(GLclampf value, GLboolean invert) {
// 	(*ptrglSampleCoverage)(value, invert);
// }
// void goglCompressedTexImage3D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
// 	(*ptrglCompressedTexImage3D)(target, level, internalformat, width, height, depth, border, imageSize, data);
// }
// void goglCompressedTexImage2D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
// 	(*ptrglCompressedTexImage2D)(target, level, internalformat, width, height, border, imageSize, data);
// }
// void goglCompressedTexImage1D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
// 	(*ptrglCompressedTexImage1D)(target, level, internalformat, width, border, imageSize, data);
// }
// void goglCompressedTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
// 	(*ptrglCompressedTexSubImage3D)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
// }
// void goglCompressedTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
// 	(*ptrglCompressedTexSubImage2D)(target, level, xoffset, yoffset, width, height, format, imageSize, data);
// }
// void goglCompressedTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
// 	(*ptrglCompressedTexSubImage1D)(target, level, xoffset, width, format, imageSize, data);
// }
// void goglGetCompressedTexImage(GLenum target, GLint level, GLvoid* img) {
// 	(*ptrglGetCompressedTexImage)(target, level, img);
// }
// //  VERSION_1_3_DEPRECATED
// void goglClientActiveTexture(GLenum texture) {
// 	(*ptrglClientActiveTexture)(texture);
// }
// void goglMultiTexCoord1d(GLenum target, GLdouble s) {
// 	(*ptrglMultiTexCoord1d)(target, s);
// }
// void goglMultiTexCoord1dv(GLenum target, GLdouble* v) {
// 	(*ptrglMultiTexCoord1dv)(target, v);
// }
// void goglMultiTexCoord1f(GLenum target, GLfloat s) {
// 	(*ptrglMultiTexCoord1f)(target, s);
// }
// void goglMultiTexCoord1fv(GLenum target, GLfloat* v) {
// 	(*ptrglMultiTexCoord1fv)(target, v);
// }
// void goglMultiTexCoord1i(GLenum target, GLint s) {
// 	(*ptrglMultiTexCoord1i)(target, s);
// }
// void goglMultiTexCoord1iv(GLenum target, GLint* v) {
// 	(*ptrglMultiTexCoord1iv)(target, v);
// }
// void goglMultiTexCoord1s(GLenum target, GLshort s) {
// 	(*ptrglMultiTexCoord1s)(target, s);
// }
// void goglMultiTexCoord1sv(GLenum target, GLshort* v) {
// 	(*ptrglMultiTexCoord1sv)(target, v);
// }
// void goglMultiTexCoord2d(GLenum target, GLdouble s, GLdouble t) {
// 	(*ptrglMultiTexCoord2d)(target, s, t);
// }
// void goglMultiTexCoord2dv(GLenum target, GLdouble* v) {
// 	(*ptrglMultiTexCoord2dv)(target, v);
// }
// void goglMultiTexCoord2f(GLenum target, GLfloat s, GLfloat t) {
// 	(*ptrglMultiTexCoord2f)(target, s, t);
// }
// void goglMultiTexCoord2fv(GLenum target, GLfloat* v) {
// 	(*ptrglMultiTexCoord2fv)(target, v);
// }
// void goglMultiTexCoord2i(GLenum target, GLint s, GLint t) {
// 	(*ptrglMultiTexCoord2i)(target, s, t);
// }
// void goglMultiTexCoord2iv(GLenum target, GLint* v) {
// 	(*ptrglMultiTexCoord2iv)(target, v);
// }
// void goglMultiTexCoord2s(GLenum target, GLshort s, GLshort t) {
// 	(*ptrglMultiTexCoord2s)(target, s, t);
// }
// void goglMultiTexCoord2sv(GLenum target, GLshort* v) {
// 	(*ptrglMultiTexCoord2sv)(target, v);
// }
// void goglMultiTexCoord3d(GLenum target, GLdouble s, GLdouble t, GLdouble r) {
// 	(*ptrglMultiTexCoord3d)(target, s, t, r);
// }
// void goglMultiTexCoord3dv(GLenum target, GLdouble* v) {
// 	(*ptrglMultiTexCoord3dv)(target, v);
// }
// void goglMultiTexCoord3f(GLenum target, GLfloat s, GLfloat t, GLfloat r) {
// 	(*ptrglMultiTexCoord3f)(target, s, t, r);
// }
// void goglMultiTexCoord3fv(GLenum target, GLfloat* v) {
// 	(*ptrglMultiTexCoord3fv)(target, v);
// }
// void goglMultiTexCoord3i(GLenum target, GLint s, GLint t, GLint r) {
// 	(*ptrglMultiTexCoord3i)(target, s, t, r);
// }
// void goglMultiTexCoord3iv(GLenum target, GLint* v) {
// 	(*ptrglMultiTexCoord3iv)(target, v);
// }
// void goglMultiTexCoord3s(GLenum target, GLshort s, GLshort t, GLshort r) {
// 	(*ptrglMultiTexCoord3s)(target, s, t, r);
// }
// void goglMultiTexCoord3sv(GLenum target, GLshort* v) {
// 	(*ptrglMultiTexCoord3sv)(target, v);
// }
// void goglMultiTexCoord4d(GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
// 	(*ptrglMultiTexCoord4d)(target, s, t, r, q);
// }
// void goglMultiTexCoord4dv(GLenum target, GLdouble* v) {
// 	(*ptrglMultiTexCoord4dv)(target, v);
// }
// void goglMultiTexCoord4f(GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
// 	(*ptrglMultiTexCoord4f)(target, s, t, r, q);
// }
// void goglMultiTexCoord4fv(GLenum target, GLfloat* v) {
// 	(*ptrglMultiTexCoord4fv)(target, v);
// }
// void goglMultiTexCoord4i(GLenum target, GLint s, GLint t, GLint r, GLint q) {
// 	(*ptrglMultiTexCoord4i)(target, s, t, r, q);
// }
// void goglMultiTexCoord4iv(GLenum target, GLint* v) {
// 	(*ptrglMultiTexCoord4iv)(target, v);
// }
// void goglMultiTexCoord4s(GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
// 	(*ptrglMultiTexCoord4s)(target, s, t, r, q);
// }
// void goglMultiTexCoord4sv(GLenum target, GLshort* v) {
// 	(*ptrglMultiTexCoord4sv)(target, v);
// }
// void goglLoadTransposeMatrixf(GLfloat* m) {
// 	(*ptrglLoadTransposeMatrixf)(m);
// }
// void goglLoadTransposeMatrixd(GLdouble* m) {
// 	(*ptrglLoadTransposeMatrixd)(m);
// }
// void goglMultTransposeMatrixf(GLfloat* m) {
// 	(*ptrglMultTransposeMatrixf)(m);
// }
// void goglMultTransposeMatrixd(GLdouble* m) {
// 	(*ptrglMultTransposeMatrixd)(m);
// }
// //  VERSION_1_4
// void goglBlendFuncSeparate(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha) {
// 	(*ptrglBlendFuncSeparate)(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha);
// }
// void goglMultiDrawArrays(GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
// 	(*ptrglMultiDrawArrays)(mode, first, count, primcount);
// }
// void goglMultiDrawElements(GLenum mode, GLsizei* count, GLenum type_, GLvoid** indices, GLsizei primcount) {
// 	(*ptrglMultiDrawElements)(mode, count, type_, indices, primcount);
// }
// void goglPointParameterf(GLenum pname, GLfloat param) {
// 	(*ptrglPointParameterf)(pname, param);
// }
// void goglPointParameterfv(GLenum pname, GLfloat* params) {
// 	(*ptrglPointParameterfv)(pname, params);
// }
// void goglPointParameteri(GLenum pname, GLint param) {
// 	(*ptrglPointParameteri)(pname, param);
// }
// void goglPointParameteriv(GLenum pname, GLint* params) {
// 	(*ptrglPointParameteriv)(pname, params);
// }
// //  VERSION_1_4_DEPRECATED
// void goglFogCoordf(GLfloat coord) {
// 	(*ptrglFogCoordf)(coord);
// }
// void goglFogCoordfv(GLfloat* coord) {
// 	(*ptrglFogCoordfv)(coord);
// }
// void goglFogCoordd(GLdouble coord) {
// 	(*ptrglFogCoordd)(coord);
// }
// void goglFogCoorddv(GLdouble* coord) {
// 	(*ptrglFogCoorddv)(coord);
// }
// void goglFogCoordPointer(GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglFogCoordPointer)(type_, stride, pointer);
// }
// void goglSecondaryColor3b(GLbyte red, GLbyte green, GLbyte blue) {
// 	(*ptrglSecondaryColor3b)(red, green, blue);
// }
// void goglSecondaryColor3bv(GLbyte* v) {
// 	(*ptrglSecondaryColor3bv)(v);
// }
// void goglSecondaryColor3d(GLdouble red, GLdouble green, GLdouble blue) {
// 	(*ptrglSecondaryColor3d)(red, green, blue);
// }
// void goglSecondaryColor3dv(GLdouble* v) {
// 	(*ptrglSecondaryColor3dv)(v);
// }
// void goglSecondaryColor3f(GLfloat red, GLfloat green, GLfloat blue) {
// 	(*ptrglSecondaryColor3f)(red, green, blue);
// }
// void goglSecondaryColor3fv(GLfloat* v) {
// 	(*ptrglSecondaryColor3fv)(v);
// }
// void goglSecondaryColor3i(GLint red, GLint green, GLint blue) {
// 	(*ptrglSecondaryColor3i)(red, green, blue);
// }
// void goglSecondaryColor3iv(GLint* v) {
// 	(*ptrglSecondaryColor3iv)(v);
// }
// void goglSecondaryColor3s(GLshort red, GLshort green, GLshort blue) {
// 	(*ptrglSecondaryColor3s)(red, green, blue);
// }
// void goglSecondaryColor3sv(GLshort* v) {
// 	(*ptrglSecondaryColor3sv)(v);
// }
// void goglSecondaryColor3ub(GLubyte red, GLubyte green, GLubyte blue) {
// 	(*ptrglSecondaryColor3ub)(red, green, blue);
// }
// void goglSecondaryColor3ubv(GLubyte* v) {
// 	(*ptrglSecondaryColor3ubv)(v);
// }
// void goglSecondaryColor3ui(GLuint red, GLuint green, GLuint blue) {
// 	(*ptrglSecondaryColor3ui)(red, green, blue);
// }
// void goglSecondaryColor3uiv(GLuint* v) {
// 	(*ptrglSecondaryColor3uiv)(v);
// }
// void goglSecondaryColor3us(GLushort red, GLushort green, GLushort blue) {
// 	(*ptrglSecondaryColor3us)(red, green, blue);
// }
// void goglSecondaryColor3usv(GLushort* v) {
// 	(*ptrglSecondaryColor3usv)(v);
// }
// void goglSecondaryColorPointer(GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglSecondaryColorPointer)(size, type_, stride, pointer);
// }
// void goglWindowPos2d(GLdouble x, GLdouble y) {
// 	(*ptrglWindowPos2d)(x, y);
// }
// void goglWindowPos2dv(GLdouble* v) {
// 	(*ptrglWindowPos2dv)(v);
// }
// void goglWindowPos2f(GLfloat x, GLfloat y) {
// 	(*ptrglWindowPos2f)(x, y);
// }
// void goglWindowPos2fv(GLfloat* v) {
// 	(*ptrglWindowPos2fv)(v);
// }
// void goglWindowPos2i(GLint x, GLint y) {
// 	(*ptrglWindowPos2i)(x, y);
// }
// void goglWindowPos2iv(GLint* v) {
// 	(*ptrglWindowPos2iv)(v);
// }
// void goglWindowPos2s(GLshort x, GLshort y) {
// 	(*ptrglWindowPos2s)(x, y);
// }
// void goglWindowPos2sv(GLshort* v) {
// 	(*ptrglWindowPos2sv)(v);
// }
// void goglWindowPos3d(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglWindowPos3d)(x, y, z);
// }
// void goglWindowPos3dv(GLdouble* v) {
// 	(*ptrglWindowPos3dv)(v);
// }
// void goglWindowPos3f(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglWindowPos3f)(x, y, z);
// }
// void goglWindowPos3fv(GLfloat* v) {
// 	(*ptrglWindowPos3fv)(v);
// }
// void goglWindowPos3i(GLint x, GLint y, GLint z) {
// 	(*ptrglWindowPos3i)(x, y, z);
// }
// void goglWindowPos3iv(GLint* v) {
// 	(*ptrglWindowPos3iv)(v);
// }
// void goglWindowPos3s(GLshort x, GLshort y, GLshort z) {
// 	(*ptrglWindowPos3s)(x, y, z);
// }
// void goglWindowPos3sv(GLshort* v) {
// 	(*ptrglWindowPos3sv)(v);
// }
// //  VERSION_1_5
// void goglGenQueries(GLsizei n, GLuint* ids) {
// 	(*ptrglGenQueries)(n, ids);
// }
// void goglDeleteQueries(GLsizei n, GLuint* ids) {
// 	(*ptrglDeleteQueries)(n, ids);
// }
// GLboolean goglIsQuery(GLuint id) {
// 	return (*ptrglIsQuery)(id);
// }
// void goglBeginQuery(GLenum target, GLuint id) {
// 	(*ptrglBeginQuery)(target, id);
// }
// void goglEndQuery(GLenum target) {
// 	(*ptrglEndQuery)(target);
// }
// void goglGetQueryiv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetQueryiv)(target, pname, params);
// }
// void goglGetQueryObjectiv(GLuint id, GLenum pname, GLint* params) {
// 	(*ptrglGetQueryObjectiv)(id, pname, params);
// }
// void goglGetQueryObjectuiv(GLuint id, GLenum pname, GLuint* params) {
// 	(*ptrglGetQueryObjectuiv)(id, pname, params);
// }
// void goglBindBuffer(GLenum target, GLuint buffer) {
// 	(*ptrglBindBuffer)(target, buffer);
// }
// void goglDeleteBuffers(GLsizei n, GLuint* buffers) {
// 	(*ptrglDeleteBuffers)(n, buffers);
// }
// void goglGenBuffers(GLsizei n, GLuint* buffers) {
// 	(*ptrglGenBuffers)(n, buffers);
// }
// GLboolean goglIsBuffer(GLuint buffer) {
// 	return (*ptrglIsBuffer)(buffer);
// }
// void goglBufferData(GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
// 	(*ptrglBufferData)(target, size, data, usage);
// }
// void goglBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
// 	(*ptrglBufferSubData)(target, offset, size, data);
// }
// void goglGetBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
// 	(*ptrglGetBufferSubData)(target, offset, size, data);
// }
// GLvoid* goglMapBuffer(GLenum target, GLenum access) {
// 	return (*ptrglMapBuffer)(target, access);
// }
// GLboolean goglUnmapBuffer(GLenum target) {
// 	return (*ptrglUnmapBuffer)(target);
// }
// void goglGetBufferParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetBufferParameteriv)(target, pname, params);
// }
// void goglGetBufferPointerv(GLenum target, GLenum pname, GLvoid** params) {
// 	(*ptrglGetBufferPointerv)(target, pname, params);
// }
// //  VERSION_2_0
// void goglBlendEquationSeparate(GLenum modeRGB, GLenum modeAlpha) {
// 	(*ptrglBlendEquationSeparate)(modeRGB, modeAlpha);
// }
// void goglDrawBuffers(GLsizei n, GLenum* bufs) {
// 	(*ptrglDrawBuffers)(n, bufs);
// }
// void goglStencilOpSeparate(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
// 	(*ptrglStencilOpSeparate)(face, sfail, dpfail, dppass);
// }
// void goglStencilFuncSeparate(GLenum face, GLenum func_, GLint ref, GLuint mask) {
// 	(*ptrglStencilFuncSeparate)(face, func_, ref, mask);
// }
// void goglStencilMaskSeparate(GLenum face, GLuint mask) {
// 	(*ptrglStencilMaskSeparate)(face, mask);
// }
// void goglAttachShader(GLuint program, GLuint shader) {
// 	(*ptrglAttachShader)(program, shader);
// }
// void goglBindAttribLocation(GLuint program, GLuint index, GLchar* name) {
// 	(*ptrglBindAttribLocation)(program, index, name);
// }
// void goglCompileShader(GLuint shader) {
// 	(*ptrglCompileShader)(shader);
// }
// GLuint goglCreateProgram() {
// 	return (*ptrglCreateProgram)();
// }
// GLuint goglCreateShader(GLenum type_) {
// 	return (*ptrglCreateShader)(type_);
// }
// void goglDeleteProgram(GLuint program) {
// 	(*ptrglDeleteProgram)(program);
// }
// void goglDeleteShader(GLuint shader) {
// 	(*ptrglDeleteShader)(shader);
// }
// void goglDetachShader(GLuint program, GLuint shader) {
// 	(*ptrglDetachShader)(program, shader);
// }
// void goglDisableVertexAttribArray(GLuint index) {
// 	(*ptrglDisableVertexAttribArray)(index);
// }
// void goglEnableVertexAttribArray(GLuint index) {
// 	(*ptrglEnableVertexAttribArray)(index);
// }
// void goglGetActiveAttrib(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type_, GLchar* name) {
// 	(*ptrglGetActiveAttrib)(program, index, bufSize, length, size, type_, name);
// }
// void goglGetActiveUniform(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type_, GLchar* name) {
// 	(*ptrglGetActiveUniform)(program, index, bufSize, length, size, type_, name);
// }
// void goglGetAttachedShaders(GLuint program, GLsizei maxCount, GLsizei* count, GLuint* obj) {
// 	(*ptrglGetAttachedShaders)(program, maxCount, count, obj);
// }
// GLint goglGetAttribLocation(GLuint program, GLchar* name) {
// 	return (*ptrglGetAttribLocation)(program, name);
// }
// void goglGetProgramiv(GLuint program, GLenum pname, GLint* params) {
// 	(*ptrglGetProgramiv)(program, pname, params);
// }
// void goglGetProgramInfoLog(GLuint program, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {
// 	(*ptrglGetProgramInfoLog)(program, bufSize, length, infoLog);
// }
// void goglGetShaderiv(GLuint shader, GLenum pname, GLint* params) {
// 	(*ptrglGetShaderiv)(shader, pname, params);
// }
// void goglGetShaderInfoLog(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {
// 	(*ptrglGetShaderInfoLog)(shader, bufSize, length, infoLog);
// }
// void goglGetShaderSource(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
// 	(*ptrglGetShaderSource)(shader, bufSize, length, source);
// }
// GLint goglGetUniformLocation(GLuint program, GLchar* name) {
// 	return (*ptrglGetUniformLocation)(program, name);
// }
// void goglGetUniformfv(GLuint program, GLint location, GLfloat* params) {
// 	(*ptrglGetUniformfv)(program, location, params);
// }
// void goglGetUniformiv(GLuint program, GLint location, GLint* params) {
// 	(*ptrglGetUniformiv)(program, location, params);
// }
// void goglGetVertexAttribdv(GLuint index, GLenum pname, GLdouble* params) {
// 	(*ptrglGetVertexAttribdv)(index, pname, params);
// }
// void goglGetVertexAttribfv(GLuint index, GLenum pname, GLfloat* params) {
// 	(*ptrglGetVertexAttribfv)(index, pname, params);
// }
// void goglGetVertexAttribiv(GLuint index, GLenum pname, GLint* params) {
// 	(*ptrglGetVertexAttribiv)(index, pname, params);
// }
// void goglGetVertexAttribPointerv(GLuint index, GLenum pname, GLvoid** pointer) {
// 	(*ptrglGetVertexAttribPointerv)(index, pname, pointer);
// }
// GLboolean goglIsProgram(GLuint program) {
// 	return (*ptrglIsProgram)(program);
// }
// GLboolean goglIsShader(GLuint shader) {
// 	return (*ptrglIsShader)(shader);
// }
// void goglLinkProgram(GLuint program) {
// 	(*ptrglLinkProgram)(program);
// }
// void goglShaderSource(GLuint shader, GLsizei count, GLchar** string_, GLint* length) {
// 	(*ptrglShaderSource)(shader, count, string_, length);
// }
// void goglUseProgram(GLuint program) {
// 	(*ptrglUseProgram)(program);
// }
// void goglUniform1f(GLint location, GLfloat v0) {
// 	(*ptrglUniform1f)(location, v0);
// }
// void goglUniform2f(GLint location, GLfloat v0, GLfloat v1) {
// 	(*ptrglUniform2f)(location, v0, v1);
// }
// void goglUniform3f(GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
// 	(*ptrglUniform3f)(location, v0, v1, v2);
// }
// void goglUniform4f(GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
// 	(*ptrglUniform4f)(location, v0, v1, v2, v3);
// }
// void goglUniform1i(GLint location, GLint v0) {
// 	(*ptrglUniform1i)(location, v0);
// }
// void goglUniform2i(GLint location, GLint v0, GLint v1) {
// 	(*ptrglUniform2i)(location, v0, v1);
// }
// void goglUniform3i(GLint location, GLint v0, GLint v1, GLint v2) {
// 	(*ptrglUniform3i)(location, v0, v1, v2);
// }
// void goglUniform4i(GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
// 	(*ptrglUniform4i)(location, v0, v1, v2, v3);
// }
// void goglUniform1fv(GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrglUniform1fv)(location, count, value);
// }
// void goglUniform2fv(GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrglUniform2fv)(location, count, value);
// }
// void goglUniform3fv(GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrglUniform3fv)(location, count, value);
// }
// void goglUniform4fv(GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrglUniform4fv)(location, count, value);
// }
// void goglUniform1iv(GLint location, GLsizei count, GLint* value) {
// 	(*ptrglUniform1iv)(location, count, value);
// }
// void goglUniform2iv(GLint location, GLsizei count, GLint* value) {
// 	(*ptrglUniform2iv)(location, count, value);
// }
// void goglUniform3iv(GLint location, GLsizei count, GLint* value) {
// 	(*ptrglUniform3iv)(location, count, value);
// }
// void goglUniform4iv(GLint location, GLsizei count, GLint* value) {
// 	(*ptrglUniform4iv)(location, count, value);
// }
// void goglUniformMatrix2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglUniformMatrix2fv)(location, count, transpose, value);
// }
// void goglUniformMatrix3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglUniformMatrix3fv)(location, count, transpose, value);
// }
// void goglUniformMatrix4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglUniformMatrix4fv)(location, count, transpose, value);
// }
// void goglValidateProgram(GLuint program) {
// 	(*ptrglValidateProgram)(program);
// }
// void goglVertexAttrib1d(GLuint index, GLdouble x) {
// 	(*ptrglVertexAttrib1d)(index, x);
// }
// void goglVertexAttrib1dv(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttrib1dv)(index, v);
// }
// void goglVertexAttrib1f(GLuint index, GLfloat x) {
// 	(*ptrglVertexAttrib1f)(index, x);
// }
// void goglVertexAttrib1fv(GLuint index, GLfloat* v) {
// 	(*ptrglVertexAttrib1fv)(index, v);
// }
// void goglVertexAttrib1s(GLuint index, GLshort x) {
// 	(*ptrglVertexAttrib1s)(index, x);
// }
// void goglVertexAttrib1sv(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttrib1sv)(index, v);
// }
// void goglVertexAttrib2d(GLuint index, GLdouble x, GLdouble y) {
// 	(*ptrglVertexAttrib2d)(index, x, y);
// }
// void goglVertexAttrib2dv(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttrib2dv)(index, v);
// }
// void goglVertexAttrib2f(GLuint index, GLfloat x, GLfloat y) {
// 	(*ptrglVertexAttrib2f)(index, x, y);
// }
// void goglVertexAttrib2fv(GLuint index, GLfloat* v) {
// 	(*ptrglVertexAttrib2fv)(index, v);
// }
// void goglVertexAttrib2s(GLuint index, GLshort x, GLshort y) {
// 	(*ptrglVertexAttrib2s)(index, x, y);
// }
// void goglVertexAttrib2sv(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttrib2sv)(index, v);
// }
// void goglVertexAttrib3d(GLuint index, GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglVertexAttrib3d)(index, x, y, z);
// }
// void goglVertexAttrib3dv(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttrib3dv)(index, v);
// }
// void goglVertexAttrib3f(GLuint index, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglVertexAttrib3f)(index, x, y, z);
// }
// void goglVertexAttrib3fv(GLuint index, GLfloat* v) {
// 	(*ptrglVertexAttrib3fv)(index, v);
// }
// void goglVertexAttrib3s(GLuint index, GLshort x, GLshort y, GLshort z) {
// 	(*ptrglVertexAttrib3s)(index, x, y, z);
// }
// void goglVertexAttrib3sv(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttrib3sv)(index, v);
// }
// void goglVertexAttrib4Nbv(GLuint index, GLbyte* v) {
// 	(*ptrglVertexAttrib4Nbv)(index, v);
// }
// void goglVertexAttrib4Niv(GLuint index, GLint* v) {
// 	(*ptrglVertexAttrib4Niv)(index, v);
// }
// void goglVertexAttrib4Nsv(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttrib4Nsv)(index, v);
// }
// void goglVertexAttrib4Nub(GLuint index, GLubyte x, GLubyte y, GLubyte z, GLubyte w) {
// 	(*ptrglVertexAttrib4Nub)(index, x, y, z, w);
// }
// void goglVertexAttrib4Nubv(GLuint index, GLubyte* v) {
// 	(*ptrglVertexAttrib4Nubv)(index, v);
// }
// void goglVertexAttrib4Nuiv(GLuint index, GLuint* v) {
// 	(*ptrglVertexAttrib4Nuiv)(index, v);
// }
// void goglVertexAttrib4Nusv(GLuint index, GLushort* v) {
// 	(*ptrglVertexAttrib4Nusv)(index, v);
// }
// void goglVertexAttrib4bv(GLuint index, GLbyte* v) {
// 	(*ptrglVertexAttrib4bv)(index, v);
// }
// void goglVertexAttrib4d(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglVertexAttrib4d)(index, x, y, z, w);
// }
// void goglVertexAttrib4dv(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttrib4dv)(index, v);
// }
// void goglVertexAttrib4f(GLuint index, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglVertexAttrib4f)(index, x, y, z, w);
// }
// void goglVertexAttrib4fv(GLuint index, GLfloat* v) {
// 	(*ptrglVertexAttrib4fv)(index, v);
// }
// void goglVertexAttrib4iv(GLuint index, GLint* v) {
// 	(*ptrglVertexAttrib4iv)(index, v);
// }
// void goglVertexAttrib4s(GLuint index, GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrglVertexAttrib4s)(index, x, y, z, w);
// }
// void goglVertexAttrib4sv(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttrib4sv)(index, v);
// }
// void goglVertexAttrib4ubv(GLuint index, GLubyte* v) {
// 	(*ptrglVertexAttrib4ubv)(index, v);
// }
// void goglVertexAttrib4uiv(GLuint index, GLuint* v) {
// 	(*ptrglVertexAttrib4uiv)(index, v);
// }
// void goglVertexAttrib4usv(GLuint index, GLushort* v) {
// 	(*ptrglVertexAttrib4usv)(index, v);
// }
// void goglVertexAttribPointer(GLuint index, GLint size, GLenum type_, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglVertexAttribPointer)(index, size, type_, normalized, stride, pointer);
// }
// //  VERSION_2_1
// void goglUniformMatrix2x3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglUniformMatrix2x3fv)(location, count, transpose, value);
// }
// void goglUniformMatrix3x2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglUniformMatrix3x2fv)(location, count, transpose, value);
// }
// void goglUniformMatrix2x4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglUniformMatrix2x4fv)(location, count, transpose, value);
// }
// void goglUniformMatrix4x2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglUniformMatrix4x2fv)(location, count, transpose, value);
// }
// void goglUniformMatrix3x4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglUniformMatrix3x4fv)(location, count, transpose, value);
// }
// void goglUniformMatrix4x3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglUniformMatrix4x3fv)(location, count, transpose, value);
// }
// 
// int init_VERSION_1_0() {
// 	ptrglCullFace = goglGetProcAddress("glCullFace");
// 	if(ptrglCullFace == NULL) return 1;
// 	ptrglFrontFace = goglGetProcAddress("glFrontFace");
// 	if(ptrglFrontFace == NULL) return 1;
// 	ptrglHint = goglGetProcAddress("glHint");
// 	if(ptrglHint == NULL) return 1;
// 	ptrglLineWidth = goglGetProcAddress("glLineWidth");
// 	if(ptrglLineWidth == NULL) return 1;
// 	ptrglPointSize = goglGetProcAddress("glPointSize");
// 	if(ptrglPointSize == NULL) return 1;
// 	ptrglPolygonMode = goglGetProcAddress("glPolygonMode");
// 	if(ptrglPolygonMode == NULL) return 1;
// 	ptrglScissor = goglGetProcAddress("glScissor");
// 	if(ptrglScissor == NULL) return 1;
// 	ptrglTexParameterf = goglGetProcAddress("glTexParameterf");
// 	if(ptrglTexParameterf == NULL) return 1;
// 	ptrglTexParameterfv = goglGetProcAddress("glTexParameterfv");
// 	if(ptrglTexParameterfv == NULL) return 1;
// 	ptrglTexParameteri = goglGetProcAddress("glTexParameteri");
// 	if(ptrglTexParameteri == NULL) return 1;
// 	ptrglTexParameteriv = goglGetProcAddress("glTexParameteriv");
// 	if(ptrglTexParameteriv == NULL) return 1;
// 	ptrglTexImage1D = goglGetProcAddress("glTexImage1D");
// 	if(ptrglTexImage1D == NULL) return 1;
// 	ptrglTexImage2D = goglGetProcAddress("glTexImage2D");
// 	if(ptrglTexImage2D == NULL) return 1;
// 	ptrglDrawBuffer = goglGetProcAddress("glDrawBuffer");
// 	if(ptrglDrawBuffer == NULL) return 1;
// 	ptrglClear = goglGetProcAddress("glClear");
// 	if(ptrglClear == NULL) return 1;
// 	ptrglClearColor = goglGetProcAddress("glClearColor");
// 	if(ptrglClearColor == NULL) return 1;
// 	ptrglClearStencil = goglGetProcAddress("glClearStencil");
// 	if(ptrglClearStencil == NULL) return 1;
// 	ptrglClearDepth = goglGetProcAddress("glClearDepth");
// 	if(ptrglClearDepth == NULL) return 1;
// 	ptrglStencilMask = goglGetProcAddress("glStencilMask");
// 	if(ptrglStencilMask == NULL) return 1;
// 	ptrglColorMask = goglGetProcAddress("glColorMask");
// 	if(ptrglColorMask == NULL) return 1;
// 	ptrglDepthMask = goglGetProcAddress("glDepthMask");
// 	if(ptrglDepthMask == NULL) return 1;
// 	ptrglDisable = goglGetProcAddress("glDisable");
// 	if(ptrglDisable == NULL) return 1;
// 	ptrglEnable = goglGetProcAddress("glEnable");
// 	if(ptrglEnable == NULL) return 1;
// 	ptrglFinish = goglGetProcAddress("glFinish");
// 	if(ptrglFinish == NULL) return 1;
// 	ptrglFlush = goglGetProcAddress("glFlush");
// 	if(ptrglFlush == NULL) return 1;
// 	ptrglBlendFunc = goglGetProcAddress("glBlendFunc");
// 	if(ptrglBlendFunc == NULL) return 1;
// 	ptrglLogicOp = goglGetProcAddress("glLogicOp");
// 	if(ptrglLogicOp == NULL) return 1;
// 	ptrglStencilFunc = goglGetProcAddress("glStencilFunc");
// 	if(ptrglStencilFunc == NULL) return 1;
// 	ptrglStencilOp = goglGetProcAddress("glStencilOp");
// 	if(ptrglStencilOp == NULL) return 1;
// 	ptrglDepthFunc = goglGetProcAddress("glDepthFunc");
// 	if(ptrglDepthFunc == NULL) return 1;
// 	ptrglPixelStoref = goglGetProcAddress("glPixelStoref");
// 	if(ptrglPixelStoref == NULL) return 1;
// 	ptrglPixelStorei = goglGetProcAddress("glPixelStorei");
// 	if(ptrglPixelStorei == NULL) return 1;
// 	ptrglReadBuffer = goglGetProcAddress("glReadBuffer");
// 	if(ptrglReadBuffer == NULL) return 1;
// 	ptrglReadPixels = goglGetProcAddress("glReadPixels");
// 	if(ptrglReadPixels == NULL) return 1;
// 	ptrglGetBooleanv = goglGetProcAddress("glGetBooleanv");
// 	if(ptrglGetBooleanv == NULL) return 1;
// 	ptrglGetDoublev = goglGetProcAddress("glGetDoublev");
// 	if(ptrglGetDoublev == NULL) return 1;
// 	ptrglGetError = goglGetProcAddress("glGetError");
// 	if(ptrglGetError == NULL) return 1;
// 	ptrglGetFloatv = goglGetProcAddress("glGetFloatv");
// 	if(ptrglGetFloatv == NULL) return 1;
// 	ptrglGetIntegerv = goglGetProcAddress("glGetIntegerv");
// 	if(ptrglGetIntegerv == NULL) return 1;
// 	ptrglGetString = goglGetProcAddress("glGetString");
// 	if(ptrglGetString == NULL) return 1;
// 	ptrglGetTexImage = goglGetProcAddress("glGetTexImage");
// 	if(ptrglGetTexImage == NULL) return 1;
// 	ptrglGetTexParameterfv = goglGetProcAddress("glGetTexParameterfv");
// 	if(ptrglGetTexParameterfv == NULL) return 1;
// 	ptrglGetTexParameteriv = goglGetProcAddress("glGetTexParameteriv");
// 	if(ptrglGetTexParameteriv == NULL) return 1;
// 	ptrglGetTexLevelParameterfv = goglGetProcAddress("glGetTexLevelParameterfv");
// 	if(ptrglGetTexLevelParameterfv == NULL) return 1;
// 	ptrglGetTexLevelParameteriv = goglGetProcAddress("glGetTexLevelParameteriv");
// 	if(ptrglGetTexLevelParameteriv == NULL) return 1;
// 	ptrglIsEnabled = goglGetProcAddress("glIsEnabled");
// 	if(ptrglIsEnabled == NULL) return 1;
// 	ptrglDepthRange = goglGetProcAddress("glDepthRange");
// 	if(ptrglDepthRange == NULL) return 1;
// 	ptrglViewport = goglGetProcAddress("glViewport");
// 	if(ptrglViewport == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_0_DEPRECATED() {
// 	ptrglNewList = goglGetProcAddress("glNewList");
// 	if(ptrglNewList == NULL) return 1;
// 	ptrglEndList = goglGetProcAddress("glEndList");
// 	if(ptrglEndList == NULL) return 1;
// 	ptrglCallList = goglGetProcAddress("glCallList");
// 	if(ptrglCallList == NULL) return 1;
// 	ptrglCallLists = goglGetProcAddress("glCallLists");
// 	if(ptrglCallLists == NULL) return 1;
// 	ptrglDeleteLists = goglGetProcAddress("glDeleteLists");
// 	if(ptrglDeleteLists == NULL) return 1;
// 	ptrglGenLists = goglGetProcAddress("glGenLists");
// 	if(ptrglGenLists == NULL) return 1;
// 	ptrglListBase = goglGetProcAddress("glListBase");
// 	if(ptrglListBase == NULL) return 1;
// 	ptrglBegin = goglGetProcAddress("glBegin");
// 	if(ptrglBegin == NULL) return 1;
// 	ptrglBitmap = goglGetProcAddress("glBitmap");
// 	if(ptrglBitmap == NULL) return 1;
// 	ptrglColor3b = goglGetProcAddress("glColor3b");
// 	if(ptrglColor3b == NULL) return 1;
// 	ptrglColor3bv = goglGetProcAddress("glColor3bv");
// 	if(ptrglColor3bv == NULL) return 1;
// 	ptrglColor3d = goglGetProcAddress("glColor3d");
// 	if(ptrglColor3d == NULL) return 1;
// 	ptrglColor3dv = goglGetProcAddress("glColor3dv");
// 	if(ptrglColor3dv == NULL) return 1;
// 	ptrglColor3f = goglGetProcAddress("glColor3f");
// 	if(ptrglColor3f == NULL) return 1;
// 	ptrglColor3fv = goglGetProcAddress("glColor3fv");
// 	if(ptrglColor3fv == NULL) return 1;
// 	ptrglColor3i = goglGetProcAddress("glColor3i");
// 	if(ptrglColor3i == NULL) return 1;
// 	ptrglColor3iv = goglGetProcAddress("glColor3iv");
// 	if(ptrglColor3iv == NULL) return 1;
// 	ptrglColor3s = goglGetProcAddress("glColor3s");
// 	if(ptrglColor3s == NULL) return 1;
// 	ptrglColor3sv = goglGetProcAddress("glColor3sv");
// 	if(ptrglColor3sv == NULL) return 1;
// 	ptrglColor3ub = goglGetProcAddress("glColor3ub");
// 	if(ptrglColor3ub == NULL) return 1;
// 	ptrglColor3ubv = goglGetProcAddress("glColor3ubv");
// 	if(ptrglColor3ubv == NULL) return 1;
// 	ptrglColor3ui = goglGetProcAddress("glColor3ui");
// 	if(ptrglColor3ui == NULL) return 1;
// 	ptrglColor3uiv = goglGetProcAddress("glColor3uiv");
// 	if(ptrglColor3uiv == NULL) return 1;
// 	ptrglColor3us = goglGetProcAddress("glColor3us");
// 	if(ptrglColor3us == NULL) return 1;
// 	ptrglColor3usv = goglGetProcAddress("glColor3usv");
// 	if(ptrglColor3usv == NULL) return 1;
// 	ptrglColor4b = goglGetProcAddress("glColor4b");
// 	if(ptrglColor4b == NULL) return 1;
// 	ptrglColor4bv = goglGetProcAddress("glColor4bv");
// 	if(ptrglColor4bv == NULL) return 1;
// 	ptrglColor4d = goglGetProcAddress("glColor4d");
// 	if(ptrglColor4d == NULL) return 1;
// 	ptrglColor4dv = goglGetProcAddress("glColor4dv");
// 	if(ptrglColor4dv == NULL) return 1;
// 	ptrglColor4f = goglGetProcAddress("glColor4f");
// 	if(ptrglColor4f == NULL) return 1;
// 	ptrglColor4fv = goglGetProcAddress("glColor4fv");
// 	if(ptrglColor4fv == NULL) return 1;
// 	ptrglColor4i = goglGetProcAddress("glColor4i");
// 	if(ptrglColor4i == NULL) return 1;
// 	ptrglColor4iv = goglGetProcAddress("glColor4iv");
// 	if(ptrglColor4iv == NULL) return 1;
// 	ptrglColor4s = goglGetProcAddress("glColor4s");
// 	if(ptrglColor4s == NULL) return 1;
// 	ptrglColor4sv = goglGetProcAddress("glColor4sv");
// 	if(ptrglColor4sv == NULL) return 1;
// 	ptrglColor4ub = goglGetProcAddress("glColor4ub");
// 	if(ptrglColor4ub == NULL) return 1;
// 	ptrglColor4ubv = goglGetProcAddress("glColor4ubv");
// 	if(ptrglColor4ubv == NULL) return 1;
// 	ptrglColor4ui = goglGetProcAddress("glColor4ui");
// 	if(ptrglColor4ui == NULL) return 1;
// 	ptrglColor4uiv = goglGetProcAddress("glColor4uiv");
// 	if(ptrglColor4uiv == NULL) return 1;
// 	ptrglColor4us = goglGetProcAddress("glColor4us");
// 	if(ptrglColor4us == NULL) return 1;
// 	ptrglColor4usv = goglGetProcAddress("glColor4usv");
// 	if(ptrglColor4usv == NULL) return 1;
// 	ptrglEdgeFlag = goglGetProcAddress("glEdgeFlag");
// 	if(ptrglEdgeFlag == NULL) return 1;
// 	ptrglEdgeFlagv = goglGetProcAddress("glEdgeFlagv");
// 	if(ptrglEdgeFlagv == NULL) return 1;
// 	ptrglEnd = goglGetProcAddress("glEnd");
// 	if(ptrglEnd == NULL) return 1;
// 	ptrglIndexd = goglGetProcAddress("glIndexd");
// 	if(ptrglIndexd == NULL) return 1;
// 	ptrglIndexdv = goglGetProcAddress("glIndexdv");
// 	if(ptrglIndexdv == NULL) return 1;
// 	ptrglIndexf = goglGetProcAddress("glIndexf");
// 	if(ptrglIndexf == NULL) return 1;
// 	ptrglIndexfv = goglGetProcAddress("glIndexfv");
// 	if(ptrglIndexfv == NULL) return 1;
// 	ptrglIndexi = goglGetProcAddress("glIndexi");
// 	if(ptrglIndexi == NULL) return 1;
// 	ptrglIndexiv = goglGetProcAddress("glIndexiv");
// 	if(ptrglIndexiv == NULL) return 1;
// 	ptrglIndexs = goglGetProcAddress("glIndexs");
// 	if(ptrglIndexs == NULL) return 1;
// 	ptrglIndexsv = goglGetProcAddress("glIndexsv");
// 	if(ptrglIndexsv == NULL) return 1;
// 	ptrglNormal3b = goglGetProcAddress("glNormal3b");
// 	if(ptrglNormal3b == NULL) return 1;
// 	ptrglNormal3bv = goglGetProcAddress("glNormal3bv");
// 	if(ptrglNormal3bv == NULL) return 1;
// 	ptrglNormal3d = goglGetProcAddress("glNormal3d");
// 	if(ptrglNormal3d == NULL) return 1;
// 	ptrglNormal3dv = goglGetProcAddress("glNormal3dv");
// 	if(ptrglNormal3dv == NULL) return 1;
// 	ptrglNormal3f = goglGetProcAddress("glNormal3f");
// 	if(ptrglNormal3f == NULL) return 1;
// 	ptrglNormal3fv = goglGetProcAddress("glNormal3fv");
// 	if(ptrglNormal3fv == NULL) return 1;
// 	ptrglNormal3i = goglGetProcAddress("glNormal3i");
// 	if(ptrglNormal3i == NULL) return 1;
// 	ptrglNormal3iv = goglGetProcAddress("glNormal3iv");
// 	if(ptrglNormal3iv == NULL) return 1;
// 	ptrglNormal3s = goglGetProcAddress("glNormal3s");
// 	if(ptrglNormal3s == NULL) return 1;
// 	ptrglNormal3sv = goglGetProcAddress("glNormal3sv");
// 	if(ptrglNormal3sv == NULL) return 1;
// 	ptrglRasterPos2d = goglGetProcAddress("glRasterPos2d");
// 	if(ptrglRasterPos2d == NULL) return 1;
// 	ptrglRasterPos2dv = goglGetProcAddress("glRasterPos2dv");
// 	if(ptrglRasterPos2dv == NULL) return 1;
// 	ptrglRasterPos2f = goglGetProcAddress("glRasterPos2f");
// 	if(ptrglRasterPos2f == NULL) return 1;
// 	ptrglRasterPos2fv = goglGetProcAddress("glRasterPos2fv");
// 	if(ptrglRasterPos2fv == NULL) return 1;
// 	ptrglRasterPos2i = goglGetProcAddress("glRasterPos2i");
// 	if(ptrglRasterPos2i == NULL) return 1;
// 	ptrglRasterPos2iv = goglGetProcAddress("glRasterPos2iv");
// 	if(ptrglRasterPos2iv == NULL) return 1;
// 	ptrglRasterPos2s = goglGetProcAddress("glRasterPos2s");
// 	if(ptrglRasterPos2s == NULL) return 1;
// 	ptrglRasterPos2sv = goglGetProcAddress("glRasterPos2sv");
// 	if(ptrglRasterPos2sv == NULL) return 1;
// 	ptrglRasterPos3d = goglGetProcAddress("glRasterPos3d");
// 	if(ptrglRasterPos3d == NULL) return 1;
// 	ptrglRasterPos3dv = goglGetProcAddress("glRasterPos3dv");
// 	if(ptrglRasterPos3dv == NULL) return 1;
// 	ptrglRasterPos3f = goglGetProcAddress("glRasterPos3f");
// 	if(ptrglRasterPos3f == NULL) return 1;
// 	ptrglRasterPos3fv = goglGetProcAddress("glRasterPos3fv");
// 	if(ptrglRasterPos3fv == NULL) return 1;
// 	ptrglRasterPos3i = goglGetProcAddress("glRasterPos3i");
// 	if(ptrglRasterPos3i == NULL) return 1;
// 	ptrglRasterPos3iv = goglGetProcAddress("glRasterPos3iv");
// 	if(ptrglRasterPos3iv == NULL) return 1;
// 	ptrglRasterPos3s = goglGetProcAddress("glRasterPos3s");
// 	if(ptrglRasterPos3s == NULL) return 1;
// 	ptrglRasterPos3sv = goglGetProcAddress("glRasterPos3sv");
// 	if(ptrglRasterPos3sv == NULL) return 1;
// 	ptrglRasterPos4d = goglGetProcAddress("glRasterPos4d");
// 	if(ptrglRasterPos4d == NULL) return 1;
// 	ptrglRasterPos4dv = goglGetProcAddress("glRasterPos4dv");
// 	if(ptrglRasterPos4dv == NULL) return 1;
// 	ptrglRasterPos4f = goglGetProcAddress("glRasterPos4f");
// 	if(ptrglRasterPos4f == NULL) return 1;
// 	ptrglRasterPos4fv = goglGetProcAddress("glRasterPos4fv");
// 	if(ptrglRasterPos4fv == NULL) return 1;
// 	ptrglRasterPos4i = goglGetProcAddress("glRasterPos4i");
// 	if(ptrglRasterPos4i == NULL) return 1;
// 	ptrglRasterPos4iv = goglGetProcAddress("glRasterPos4iv");
// 	if(ptrglRasterPos4iv == NULL) return 1;
// 	ptrglRasterPos4s = goglGetProcAddress("glRasterPos4s");
// 	if(ptrglRasterPos4s == NULL) return 1;
// 	ptrglRasterPos4sv = goglGetProcAddress("glRasterPos4sv");
// 	if(ptrglRasterPos4sv == NULL) return 1;
// 	ptrglRectd = goglGetProcAddress("glRectd");
// 	if(ptrglRectd == NULL) return 1;
// 	ptrglRectdv = goglGetProcAddress("glRectdv");
// 	if(ptrglRectdv == NULL) return 1;
// 	ptrglRectf = goglGetProcAddress("glRectf");
// 	if(ptrglRectf == NULL) return 1;
// 	ptrglRectfv = goglGetProcAddress("glRectfv");
// 	if(ptrglRectfv == NULL) return 1;
// 	ptrglRecti = goglGetProcAddress("glRecti");
// 	if(ptrglRecti == NULL) return 1;
// 	ptrglRectiv = goglGetProcAddress("glRectiv");
// 	if(ptrglRectiv == NULL) return 1;
// 	ptrglRects = goglGetProcAddress("glRects");
// 	if(ptrglRects == NULL) return 1;
// 	ptrglRectsv = goglGetProcAddress("glRectsv");
// 	if(ptrglRectsv == NULL) return 1;
// 	ptrglTexCoord1d = goglGetProcAddress("glTexCoord1d");
// 	if(ptrglTexCoord1d == NULL) return 1;
// 	ptrglTexCoord1dv = goglGetProcAddress("glTexCoord1dv");
// 	if(ptrglTexCoord1dv == NULL) return 1;
// 	ptrglTexCoord1f = goglGetProcAddress("glTexCoord1f");
// 	if(ptrglTexCoord1f == NULL) return 1;
// 	ptrglTexCoord1fv = goglGetProcAddress("glTexCoord1fv");
// 	if(ptrglTexCoord1fv == NULL) return 1;
// 	ptrglTexCoord1i = goglGetProcAddress("glTexCoord1i");
// 	if(ptrglTexCoord1i == NULL) return 1;
// 	ptrglTexCoord1iv = goglGetProcAddress("glTexCoord1iv");
// 	if(ptrglTexCoord1iv == NULL) return 1;
// 	ptrglTexCoord1s = goglGetProcAddress("glTexCoord1s");
// 	if(ptrglTexCoord1s == NULL) return 1;
// 	ptrglTexCoord1sv = goglGetProcAddress("glTexCoord1sv");
// 	if(ptrglTexCoord1sv == NULL) return 1;
// 	ptrglTexCoord2d = goglGetProcAddress("glTexCoord2d");
// 	if(ptrglTexCoord2d == NULL) return 1;
// 	ptrglTexCoord2dv = goglGetProcAddress("glTexCoord2dv");
// 	if(ptrglTexCoord2dv == NULL) return 1;
// 	ptrglTexCoord2f = goglGetProcAddress("glTexCoord2f");
// 	if(ptrglTexCoord2f == NULL) return 1;
// 	ptrglTexCoord2fv = goglGetProcAddress("glTexCoord2fv");
// 	if(ptrglTexCoord2fv == NULL) return 1;
// 	ptrglTexCoord2i = goglGetProcAddress("glTexCoord2i");
// 	if(ptrglTexCoord2i == NULL) return 1;
// 	ptrglTexCoord2iv = goglGetProcAddress("glTexCoord2iv");
// 	if(ptrglTexCoord2iv == NULL) return 1;
// 	ptrglTexCoord2s = goglGetProcAddress("glTexCoord2s");
// 	if(ptrglTexCoord2s == NULL) return 1;
// 	ptrglTexCoord2sv = goglGetProcAddress("glTexCoord2sv");
// 	if(ptrglTexCoord2sv == NULL) return 1;
// 	ptrglTexCoord3d = goglGetProcAddress("glTexCoord3d");
// 	if(ptrglTexCoord3d == NULL) return 1;
// 	ptrglTexCoord3dv = goglGetProcAddress("glTexCoord3dv");
// 	if(ptrglTexCoord3dv == NULL) return 1;
// 	ptrglTexCoord3f = goglGetProcAddress("glTexCoord3f");
// 	if(ptrglTexCoord3f == NULL) return 1;
// 	ptrglTexCoord3fv = goglGetProcAddress("glTexCoord3fv");
// 	if(ptrglTexCoord3fv == NULL) return 1;
// 	ptrglTexCoord3i = goglGetProcAddress("glTexCoord3i");
// 	if(ptrglTexCoord3i == NULL) return 1;
// 	ptrglTexCoord3iv = goglGetProcAddress("glTexCoord3iv");
// 	if(ptrglTexCoord3iv == NULL) return 1;
// 	ptrglTexCoord3s = goglGetProcAddress("glTexCoord3s");
// 	if(ptrglTexCoord3s == NULL) return 1;
// 	ptrglTexCoord3sv = goglGetProcAddress("glTexCoord3sv");
// 	if(ptrglTexCoord3sv == NULL) return 1;
// 	ptrglTexCoord4d = goglGetProcAddress("glTexCoord4d");
// 	if(ptrglTexCoord4d == NULL) return 1;
// 	ptrglTexCoord4dv = goglGetProcAddress("glTexCoord4dv");
// 	if(ptrglTexCoord4dv == NULL) return 1;
// 	ptrglTexCoord4f = goglGetProcAddress("glTexCoord4f");
// 	if(ptrglTexCoord4f == NULL) return 1;
// 	ptrglTexCoord4fv = goglGetProcAddress("glTexCoord4fv");
// 	if(ptrglTexCoord4fv == NULL) return 1;
// 	ptrglTexCoord4i = goglGetProcAddress("glTexCoord4i");
// 	if(ptrglTexCoord4i == NULL) return 1;
// 	ptrglTexCoord4iv = goglGetProcAddress("glTexCoord4iv");
// 	if(ptrglTexCoord4iv == NULL) return 1;
// 	ptrglTexCoord4s = goglGetProcAddress("glTexCoord4s");
// 	if(ptrglTexCoord4s == NULL) return 1;
// 	ptrglTexCoord4sv = goglGetProcAddress("glTexCoord4sv");
// 	if(ptrglTexCoord4sv == NULL) return 1;
// 	ptrglVertex2d = goglGetProcAddress("glVertex2d");
// 	if(ptrglVertex2d == NULL) return 1;
// 	ptrglVertex2dv = goglGetProcAddress("glVertex2dv");
// 	if(ptrglVertex2dv == NULL) return 1;
// 	ptrglVertex2f = goglGetProcAddress("glVertex2f");
// 	if(ptrglVertex2f == NULL) return 1;
// 	ptrglVertex2fv = goglGetProcAddress("glVertex2fv");
// 	if(ptrglVertex2fv == NULL) return 1;
// 	ptrglVertex2i = goglGetProcAddress("glVertex2i");
// 	if(ptrglVertex2i == NULL) return 1;
// 	ptrglVertex2iv = goglGetProcAddress("glVertex2iv");
// 	if(ptrglVertex2iv == NULL) return 1;
// 	ptrglVertex2s = goglGetProcAddress("glVertex2s");
// 	if(ptrglVertex2s == NULL) return 1;
// 	ptrglVertex2sv = goglGetProcAddress("glVertex2sv");
// 	if(ptrglVertex2sv == NULL) return 1;
// 	ptrglVertex3d = goglGetProcAddress("glVertex3d");
// 	if(ptrglVertex3d == NULL) return 1;
// 	ptrglVertex3dv = goglGetProcAddress("glVertex3dv");
// 	if(ptrglVertex3dv == NULL) return 1;
// 	ptrglVertex3f = goglGetProcAddress("glVertex3f");
// 	if(ptrglVertex3f == NULL) return 1;
// 	ptrglVertex3fv = goglGetProcAddress("glVertex3fv");
// 	if(ptrglVertex3fv == NULL) return 1;
// 	ptrglVertex3i = goglGetProcAddress("glVertex3i");
// 	if(ptrglVertex3i == NULL) return 1;
// 	ptrglVertex3iv = goglGetProcAddress("glVertex3iv");
// 	if(ptrglVertex3iv == NULL) return 1;
// 	ptrglVertex3s = goglGetProcAddress("glVertex3s");
// 	if(ptrglVertex3s == NULL) return 1;
// 	ptrglVertex3sv = goglGetProcAddress("glVertex3sv");
// 	if(ptrglVertex3sv == NULL) return 1;
// 	ptrglVertex4d = goglGetProcAddress("glVertex4d");
// 	if(ptrglVertex4d == NULL) return 1;
// 	ptrglVertex4dv = goglGetProcAddress("glVertex4dv");
// 	if(ptrglVertex4dv == NULL) return 1;
// 	ptrglVertex4f = goglGetProcAddress("glVertex4f");
// 	if(ptrglVertex4f == NULL) return 1;
// 	ptrglVertex4fv = goglGetProcAddress("glVertex4fv");
// 	if(ptrglVertex4fv == NULL) return 1;
// 	ptrglVertex4i = goglGetProcAddress("glVertex4i");
// 	if(ptrglVertex4i == NULL) return 1;
// 	ptrglVertex4iv = goglGetProcAddress("glVertex4iv");
// 	if(ptrglVertex4iv == NULL) return 1;
// 	ptrglVertex4s = goglGetProcAddress("glVertex4s");
// 	if(ptrglVertex4s == NULL) return 1;
// 	ptrglVertex4sv = goglGetProcAddress("glVertex4sv");
// 	if(ptrglVertex4sv == NULL) return 1;
// 	ptrglClipPlane = goglGetProcAddress("glClipPlane");
// 	if(ptrglClipPlane == NULL) return 1;
// 	ptrglColorMaterial = goglGetProcAddress("glColorMaterial");
// 	if(ptrglColorMaterial == NULL) return 1;
// 	ptrglFogf = goglGetProcAddress("glFogf");
// 	if(ptrglFogf == NULL) return 1;
// 	ptrglFogfv = goglGetProcAddress("glFogfv");
// 	if(ptrglFogfv == NULL) return 1;
// 	ptrglFogi = goglGetProcAddress("glFogi");
// 	if(ptrglFogi == NULL) return 1;
// 	ptrglFogiv = goglGetProcAddress("glFogiv");
// 	if(ptrglFogiv == NULL) return 1;
// 	ptrglLightf = goglGetProcAddress("glLightf");
// 	if(ptrglLightf == NULL) return 1;
// 	ptrglLightfv = goglGetProcAddress("glLightfv");
// 	if(ptrglLightfv == NULL) return 1;
// 	ptrglLighti = goglGetProcAddress("glLighti");
// 	if(ptrglLighti == NULL) return 1;
// 	ptrglLightiv = goglGetProcAddress("glLightiv");
// 	if(ptrglLightiv == NULL) return 1;
// 	ptrglLightModelf = goglGetProcAddress("glLightModelf");
// 	if(ptrglLightModelf == NULL) return 1;
// 	ptrglLightModelfv = goglGetProcAddress("glLightModelfv");
// 	if(ptrglLightModelfv == NULL) return 1;
// 	ptrglLightModeli = goglGetProcAddress("glLightModeli");
// 	if(ptrglLightModeli == NULL) return 1;
// 	ptrglLightModeliv = goglGetProcAddress("glLightModeliv");
// 	if(ptrglLightModeliv == NULL) return 1;
// 	ptrglLineStipple = goglGetProcAddress("glLineStipple");
// 	if(ptrglLineStipple == NULL) return 1;
// 	ptrglMaterialf = goglGetProcAddress("glMaterialf");
// 	if(ptrglMaterialf == NULL) return 1;
// 	ptrglMaterialfv = goglGetProcAddress("glMaterialfv");
// 	if(ptrglMaterialfv == NULL) return 1;
// 	ptrglMateriali = goglGetProcAddress("glMateriali");
// 	if(ptrglMateriali == NULL) return 1;
// 	ptrglMaterialiv = goglGetProcAddress("glMaterialiv");
// 	if(ptrglMaterialiv == NULL) return 1;
// 	ptrglPolygonStipple = goglGetProcAddress("glPolygonStipple");
// 	if(ptrglPolygonStipple == NULL) return 1;
// 	ptrglShadeModel = goglGetProcAddress("glShadeModel");
// 	if(ptrglShadeModel == NULL) return 1;
// 	ptrglTexEnvf = goglGetProcAddress("glTexEnvf");
// 	if(ptrglTexEnvf == NULL) return 1;
// 	ptrglTexEnvfv = goglGetProcAddress("glTexEnvfv");
// 	if(ptrglTexEnvfv == NULL) return 1;
// 	ptrglTexEnvi = goglGetProcAddress("glTexEnvi");
// 	if(ptrglTexEnvi == NULL) return 1;
// 	ptrglTexEnviv = goglGetProcAddress("glTexEnviv");
// 	if(ptrglTexEnviv == NULL) return 1;
// 	ptrglTexGend = goglGetProcAddress("glTexGend");
// 	if(ptrglTexGend == NULL) return 1;
// 	ptrglTexGendv = goglGetProcAddress("glTexGendv");
// 	if(ptrglTexGendv == NULL) return 1;
// 	ptrglTexGenf = goglGetProcAddress("glTexGenf");
// 	if(ptrglTexGenf == NULL) return 1;
// 	ptrglTexGenfv = goglGetProcAddress("glTexGenfv");
// 	if(ptrglTexGenfv == NULL) return 1;
// 	ptrglTexGeni = goglGetProcAddress("glTexGeni");
// 	if(ptrglTexGeni == NULL) return 1;
// 	ptrglTexGeniv = goglGetProcAddress("glTexGeniv");
// 	if(ptrglTexGeniv == NULL) return 1;
// 	ptrglFeedbackBuffer = goglGetProcAddress("glFeedbackBuffer");
// 	if(ptrglFeedbackBuffer == NULL) return 1;
// 	ptrglSelectBuffer = goglGetProcAddress("glSelectBuffer");
// 	if(ptrglSelectBuffer == NULL) return 1;
// 	ptrglRenderMode = goglGetProcAddress("glRenderMode");
// 	if(ptrglRenderMode == NULL) return 1;
// 	ptrglInitNames = goglGetProcAddress("glInitNames");
// 	if(ptrglInitNames == NULL) return 1;
// 	ptrglLoadName = goglGetProcAddress("glLoadName");
// 	if(ptrglLoadName == NULL) return 1;
// 	ptrglPassThrough = goglGetProcAddress("glPassThrough");
// 	if(ptrglPassThrough == NULL) return 1;
// 	ptrglPopName = goglGetProcAddress("glPopName");
// 	if(ptrglPopName == NULL) return 1;
// 	ptrglPushName = goglGetProcAddress("glPushName");
// 	if(ptrglPushName == NULL) return 1;
// 	ptrglClearAccum = goglGetProcAddress("glClearAccum");
// 	if(ptrglClearAccum == NULL) return 1;
// 	ptrglClearIndex = goglGetProcAddress("glClearIndex");
// 	if(ptrglClearIndex == NULL) return 1;
// 	ptrglIndexMask = goglGetProcAddress("glIndexMask");
// 	if(ptrglIndexMask == NULL) return 1;
// 	ptrglAccum = goglGetProcAddress("glAccum");
// 	if(ptrglAccum == NULL) return 1;
// 	ptrglPopAttrib = goglGetProcAddress("glPopAttrib");
// 	if(ptrglPopAttrib == NULL) return 1;
// 	ptrglPushAttrib = goglGetProcAddress("glPushAttrib");
// 	if(ptrglPushAttrib == NULL) return 1;
// 	ptrglMap1d = goglGetProcAddress("glMap1d");
// 	if(ptrglMap1d == NULL) return 1;
// 	ptrglMap1f = goglGetProcAddress("glMap1f");
// 	if(ptrglMap1f == NULL) return 1;
// 	ptrglMap2d = goglGetProcAddress("glMap2d");
// 	if(ptrglMap2d == NULL) return 1;
// 	ptrglMap2f = goglGetProcAddress("glMap2f");
// 	if(ptrglMap2f == NULL) return 1;
// 	ptrglMapGrid1d = goglGetProcAddress("glMapGrid1d");
// 	if(ptrglMapGrid1d == NULL) return 1;
// 	ptrglMapGrid1f = goglGetProcAddress("glMapGrid1f");
// 	if(ptrglMapGrid1f == NULL) return 1;
// 	ptrglMapGrid2d = goglGetProcAddress("glMapGrid2d");
// 	if(ptrglMapGrid2d == NULL) return 1;
// 	ptrglMapGrid2f = goglGetProcAddress("glMapGrid2f");
// 	if(ptrglMapGrid2f == NULL) return 1;
// 	ptrglEvalCoord1d = goglGetProcAddress("glEvalCoord1d");
// 	if(ptrglEvalCoord1d == NULL) return 1;
// 	ptrglEvalCoord1dv = goglGetProcAddress("glEvalCoord1dv");
// 	if(ptrglEvalCoord1dv == NULL) return 1;
// 	ptrglEvalCoord1f = goglGetProcAddress("glEvalCoord1f");
// 	if(ptrglEvalCoord1f == NULL) return 1;
// 	ptrglEvalCoord1fv = goglGetProcAddress("glEvalCoord1fv");
// 	if(ptrglEvalCoord1fv == NULL) return 1;
// 	ptrglEvalCoord2d = goglGetProcAddress("glEvalCoord2d");
// 	if(ptrglEvalCoord2d == NULL) return 1;
// 	ptrglEvalCoord2dv = goglGetProcAddress("glEvalCoord2dv");
// 	if(ptrglEvalCoord2dv == NULL) return 1;
// 	ptrglEvalCoord2f = goglGetProcAddress("glEvalCoord2f");
// 	if(ptrglEvalCoord2f == NULL) return 1;
// 	ptrglEvalCoord2fv = goglGetProcAddress("glEvalCoord2fv");
// 	if(ptrglEvalCoord2fv == NULL) return 1;
// 	ptrglEvalMesh1 = goglGetProcAddress("glEvalMesh1");
// 	if(ptrglEvalMesh1 == NULL) return 1;
// 	ptrglEvalPoint1 = goglGetProcAddress("glEvalPoint1");
// 	if(ptrglEvalPoint1 == NULL) return 1;
// 	ptrglEvalMesh2 = goglGetProcAddress("glEvalMesh2");
// 	if(ptrglEvalMesh2 == NULL) return 1;
// 	ptrglEvalPoint2 = goglGetProcAddress("glEvalPoint2");
// 	if(ptrglEvalPoint2 == NULL) return 1;
// 	ptrglAlphaFunc = goglGetProcAddress("glAlphaFunc");
// 	if(ptrglAlphaFunc == NULL) return 1;
// 	ptrglPixelZoom = goglGetProcAddress("glPixelZoom");
// 	if(ptrglPixelZoom == NULL) return 1;
// 	ptrglPixelTransferf = goglGetProcAddress("glPixelTransferf");
// 	if(ptrglPixelTransferf == NULL) return 1;
// 	ptrglPixelTransferi = goglGetProcAddress("glPixelTransferi");
// 	if(ptrglPixelTransferi == NULL) return 1;
// 	ptrglPixelMapfv = goglGetProcAddress("glPixelMapfv");
// 	if(ptrglPixelMapfv == NULL) return 1;
// 	ptrglPixelMapuiv = goglGetProcAddress("glPixelMapuiv");
// 	if(ptrglPixelMapuiv == NULL) return 1;
// 	ptrglPixelMapusv = goglGetProcAddress("glPixelMapusv");
// 	if(ptrglPixelMapusv == NULL) return 1;
// 	ptrglCopyPixels = goglGetProcAddress("glCopyPixels");
// 	if(ptrglCopyPixels == NULL) return 1;
// 	ptrglDrawPixels = goglGetProcAddress("glDrawPixels");
// 	if(ptrglDrawPixels == NULL) return 1;
// 	ptrglGetClipPlane = goglGetProcAddress("glGetClipPlane");
// 	if(ptrglGetClipPlane == NULL) return 1;
// 	ptrglGetLightfv = goglGetProcAddress("glGetLightfv");
// 	if(ptrglGetLightfv == NULL) return 1;
// 	ptrglGetLightiv = goglGetProcAddress("glGetLightiv");
// 	if(ptrglGetLightiv == NULL) return 1;
// 	ptrglGetMapdv = goglGetProcAddress("glGetMapdv");
// 	if(ptrglGetMapdv == NULL) return 1;
// 	ptrglGetMapfv = goglGetProcAddress("glGetMapfv");
// 	if(ptrglGetMapfv == NULL) return 1;
// 	ptrglGetMapiv = goglGetProcAddress("glGetMapiv");
// 	if(ptrglGetMapiv == NULL) return 1;
// 	ptrglGetMaterialfv = goglGetProcAddress("glGetMaterialfv");
// 	if(ptrglGetMaterialfv == NULL) return 1;
// 	ptrglGetMaterialiv = goglGetProcAddress("glGetMaterialiv");
// 	if(ptrglGetMaterialiv == NULL) return 1;
// 	ptrglGetPixelMapfv = goglGetProcAddress("glGetPixelMapfv");
// 	if(ptrglGetPixelMapfv == NULL) return 1;
// 	ptrglGetPixelMapuiv = goglGetProcAddress("glGetPixelMapuiv");
// 	if(ptrglGetPixelMapuiv == NULL) return 1;
// 	ptrglGetPixelMapusv = goglGetProcAddress("glGetPixelMapusv");
// 	if(ptrglGetPixelMapusv == NULL) return 1;
// 	ptrglGetPolygonStipple = goglGetProcAddress("glGetPolygonStipple");
// 	if(ptrglGetPolygonStipple == NULL) return 1;
// 	ptrglGetTexEnvfv = goglGetProcAddress("glGetTexEnvfv");
// 	if(ptrglGetTexEnvfv == NULL) return 1;
// 	ptrglGetTexEnviv = goglGetProcAddress("glGetTexEnviv");
// 	if(ptrglGetTexEnviv == NULL) return 1;
// 	ptrglGetTexGendv = goglGetProcAddress("glGetTexGendv");
// 	if(ptrglGetTexGendv == NULL) return 1;
// 	ptrglGetTexGenfv = goglGetProcAddress("glGetTexGenfv");
// 	if(ptrglGetTexGenfv == NULL) return 1;
// 	ptrglGetTexGeniv = goglGetProcAddress("glGetTexGeniv");
// 	if(ptrglGetTexGeniv == NULL) return 1;
// 	ptrglIsList = goglGetProcAddress("glIsList");
// 	if(ptrglIsList == NULL) return 1;
// 	ptrglFrustum = goglGetProcAddress("glFrustum");
// 	if(ptrglFrustum == NULL) return 1;
// 	ptrglLoadIdentity = goglGetProcAddress("glLoadIdentity");
// 	if(ptrglLoadIdentity == NULL) return 1;
// 	ptrglLoadMatrixf = goglGetProcAddress("glLoadMatrixf");
// 	if(ptrglLoadMatrixf == NULL) return 1;
// 	ptrglLoadMatrixd = goglGetProcAddress("glLoadMatrixd");
// 	if(ptrglLoadMatrixd == NULL) return 1;
// 	ptrglMatrixMode = goglGetProcAddress("glMatrixMode");
// 	if(ptrglMatrixMode == NULL) return 1;
// 	ptrglMultMatrixf = goglGetProcAddress("glMultMatrixf");
// 	if(ptrglMultMatrixf == NULL) return 1;
// 	ptrglMultMatrixd = goglGetProcAddress("glMultMatrixd");
// 	if(ptrglMultMatrixd == NULL) return 1;
// 	ptrglOrtho = goglGetProcAddress("glOrtho");
// 	if(ptrglOrtho == NULL) return 1;
// 	ptrglPopMatrix = goglGetProcAddress("glPopMatrix");
// 	if(ptrglPopMatrix == NULL) return 1;
// 	ptrglPushMatrix = goglGetProcAddress("glPushMatrix");
// 	if(ptrglPushMatrix == NULL) return 1;
// 	ptrglRotated = goglGetProcAddress("glRotated");
// 	if(ptrglRotated == NULL) return 1;
// 	ptrglRotatef = goglGetProcAddress("glRotatef");
// 	if(ptrglRotatef == NULL) return 1;
// 	ptrglScaled = goglGetProcAddress("glScaled");
// 	if(ptrglScaled == NULL) return 1;
// 	ptrglScalef = goglGetProcAddress("glScalef");
// 	if(ptrglScalef == NULL) return 1;
// 	ptrglTranslated = goglGetProcAddress("glTranslated");
// 	if(ptrglTranslated == NULL) return 1;
// 	ptrglTranslatef = goglGetProcAddress("glTranslatef");
// 	if(ptrglTranslatef == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_1() {
// 	ptrglDrawArrays = goglGetProcAddress("glDrawArrays");
// 	if(ptrglDrawArrays == NULL) return 1;
// 	ptrglDrawElements = goglGetProcAddress("glDrawElements");
// 	if(ptrglDrawElements == NULL) return 1;
// 	ptrglGetPointerv = goglGetProcAddress("glGetPointerv");
// 	if(ptrglGetPointerv == NULL) return 1;
// 	ptrglPolygonOffset = goglGetProcAddress("glPolygonOffset");
// 	if(ptrglPolygonOffset == NULL) return 1;
// 	ptrglCopyTexImage1D = goglGetProcAddress("glCopyTexImage1D");
// 	if(ptrglCopyTexImage1D == NULL) return 1;
// 	ptrglCopyTexImage2D = goglGetProcAddress("glCopyTexImage2D");
// 	if(ptrglCopyTexImage2D == NULL) return 1;
// 	ptrglCopyTexSubImage1D = goglGetProcAddress("glCopyTexSubImage1D");
// 	if(ptrglCopyTexSubImage1D == NULL) return 1;
// 	ptrglCopyTexSubImage2D = goglGetProcAddress("glCopyTexSubImage2D");
// 	if(ptrglCopyTexSubImage2D == NULL) return 1;
// 	ptrglTexSubImage1D = goglGetProcAddress("glTexSubImage1D");
// 	if(ptrglTexSubImage1D == NULL) return 1;
// 	ptrglTexSubImage2D = goglGetProcAddress("glTexSubImage2D");
// 	if(ptrglTexSubImage2D == NULL) return 1;
// 	ptrglBindTexture = goglGetProcAddress("glBindTexture");
// 	if(ptrglBindTexture == NULL) return 1;
// 	ptrglDeleteTextures = goglGetProcAddress("glDeleteTextures");
// 	if(ptrglDeleteTextures == NULL) return 1;
// 	ptrglGenTextures = goglGetProcAddress("glGenTextures");
// 	if(ptrglGenTextures == NULL) return 1;
// 	ptrglIsTexture = goglGetProcAddress("glIsTexture");
// 	if(ptrglIsTexture == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_1_DEPRECATED() {
// 	ptrglArrayElement = goglGetProcAddress("glArrayElement");
// 	if(ptrglArrayElement == NULL) return 1;
// 	ptrglColorPointer = goglGetProcAddress("glColorPointer");
// 	if(ptrglColorPointer == NULL) return 1;
// 	ptrglDisableClientState = goglGetProcAddress("glDisableClientState");
// 	if(ptrglDisableClientState == NULL) return 1;
// 	ptrglEdgeFlagPointer = goglGetProcAddress("glEdgeFlagPointer");
// 	if(ptrglEdgeFlagPointer == NULL) return 1;
// 	ptrglEnableClientState = goglGetProcAddress("glEnableClientState");
// 	if(ptrglEnableClientState == NULL) return 1;
// 	ptrglIndexPointer = goglGetProcAddress("glIndexPointer");
// 	if(ptrglIndexPointer == NULL) return 1;
// 	ptrglInterleavedArrays = goglGetProcAddress("glInterleavedArrays");
// 	if(ptrglInterleavedArrays == NULL) return 1;
// 	ptrglNormalPointer = goglGetProcAddress("glNormalPointer");
// 	if(ptrglNormalPointer == NULL) return 1;
// 	ptrglTexCoordPointer = goglGetProcAddress("glTexCoordPointer");
// 	if(ptrglTexCoordPointer == NULL) return 1;
// 	ptrglVertexPointer = goglGetProcAddress("glVertexPointer");
// 	if(ptrglVertexPointer == NULL) return 1;
// 	ptrglAreTexturesResident = goglGetProcAddress("glAreTexturesResident");
// 	if(ptrglAreTexturesResident == NULL) return 1;
// 	ptrglPrioritizeTextures = goglGetProcAddress("glPrioritizeTextures");
// 	if(ptrglPrioritizeTextures == NULL) return 1;
// 	ptrglIndexub = goglGetProcAddress("glIndexub");
// 	if(ptrglIndexub == NULL) return 1;
// 	ptrglIndexubv = goglGetProcAddress("glIndexubv");
// 	if(ptrglIndexubv == NULL) return 1;
// 	ptrglPopClientAttrib = goglGetProcAddress("glPopClientAttrib");
// 	if(ptrglPopClientAttrib == NULL) return 1;
// 	ptrglPushClientAttrib = goglGetProcAddress("glPushClientAttrib");
// 	if(ptrglPushClientAttrib == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_2() {
// 	ptrglBlendColor = goglGetProcAddress("glBlendColor");
// 	if(ptrglBlendColor == NULL) return 1;
// 	ptrglBlendEquation = goglGetProcAddress("glBlendEquation");
// 	if(ptrglBlendEquation == NULL) return 1;
// 	ptrglDrawRangeElements = goglGetProcAddress("glDrawRangeElements");
// 	if(ptrglDrawRangeElements == NULL) return 1;
// 	ptrglTexImage3D = goglGetProcAddress("glTexImage3D");
// 	if(ptrglTexImage3D == NULL) return 1;
// 	ptrglTexSubImage3D = goglGetProcAddress("glTexSubImage3D");
// 	if(ptrglTexSubImage3D == NULL) return 1;
// 	ptrglCopyTexSubImage3D = goglGetProcAddress("glCopyTexSubImage3D");
// 	if(ptrglCopyTexSubImage3D == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_2_DEPRECATED() {
// 	ptrglColorTable = goglGetProcAddress("glColorTable");
// 	if(ptrglColorTable == NULL) return 1;
// 	ptrglColorTableParameterfv = goglGetProcAddress("glColorTableParameterfv");
// 	if(ptrglColorTableParameterfv == NULL) return 1;
// 	ptrglColorTableParameteriv = goglGetProcAddress("glColorTableParameteriv");
// 	if(ptrglColorTableParameteriv == NULL) return 1;
// 	ptrglCopyColorTable = goglGetProcAddress("glCopyColorTable");
// 	if(ptrglCopyColorTable == NULL) return 1;
// 	ptrglGetColorTable = goglGetProcAddress("glGetColorTable");
// 	if(ptrglGetColorTable == NULL) return 1;
// 	ptrglGetColorTableParameterfv = goglGetProcAddress("glGetColorTableParameterfv");
// 	if(ptrglGetColorTableParameterfv == NULL) return 1;
// 	ptrglGetColorTableParameteriv = goglGetProcAddress("glGetColorTableParameteriv");
// 	if(ptrglGetColorTableParameteriv == NULL) return 1;
// 	ptrglColorSubTable = goglGetProcAddress("glColorSubTable");
// 	if(ptrglColorSubTable == NULL) return 1;
// 	ptrglCopyColorSubTable = goglGetProcAddress("glCopyColorSubTable");
// 	if(ptrglCopyColorSubTable == NULL) return 1;
// 	ptrglConvolutionFilter1D = goglGetProcAddress("glConvolutionFilter1D");
// 	if(ptrglConvolutionFilter1D == NULL) return 1;
// 	ptrglConvolutionFilter2D = goglGetProcAddress("glConvolutionFilter2D");
// 	if(ptrglConvolutionFilter2D == NULL) return 1;
// 	ptrglConvolutionParameterf = goglGetProcAddress("glConvolutionParameterf");
// 	if(ptrglConvolutionParameterf == NULL) return 1;
// 	ptrglConvolutionParameterfv = goglGetProcAddress("glConvolutionParameterfv");
// 	if(ptrglConvolutionParameterfv == NULL) return 1;
// 	ptrglConvolutionParameteri = goglGetProcAddress("glConvolutionParameteri");
// 	if(ptrglConvolutionParameteri == NULL) return 1;
// 	ptrglConvolutionParameteriv = goglGetProcAddress("glConvolutionParameteriv");
// 	if(ptrglConvolutionParameteriv == NULL) return 1;
// 	ptrglCopyConvolutionFilter1D = goglGetProcAddress("glCopyConvolutionFilter1D");
// 	if(ptrglCopyConvolutionFilter1D == NULL) return 1;
// 	ptrglCopyConvolutionFilter2D = goglGetProcAddress("glCopyConvolutionFilter2D");
// 	if(ptrglCopyConvolutionFilter2D == NULL) return 1;
// 	ptrglGetConvolutionFilter = goglGetProcAddress("glGetConvolutionFilter");
// 	if(ptrglGetConvolutionFilter == NULL) return 1;
// 	ptrglGetConvolutionParameterfv = goglGetProcAddress("glGetConvolutionParameterfv");
// 	if(ptrglGetConvolutionParameterfv == NULL) return 1;
// 	ptrglGetConvolutionParameteriv = goglGetProcAddress("glGetConvolutionParameteriv");
// 	if(ptrglGetConvolutionParameteriv == NULL) return 1;
// 	ptrglGetSeparableFilter = goglGetProcAddress("glGetSeparableFilter");
// 	if(ptrglGetSeparableFilter == NULL) return 1;
// 	ptrglSeparableFilter2D = goglGetProcAddress("glSeparableFilter2D");
// 	if(ptrglSeparableFilter2D == NULL) return 1;
// 	ptrglGetHistogram = goglGetProcAddress("glGetHistogram");
// 	if(ptrglGetHistogram == NULL) return 1;
// 	ptrglGetHistogramParameterfv = goglGetProcAddress("glGetHistogramParameterfv");
// 	if(ptrglGetHistogramParameterfv == NULL) return 1;
// 	ptrglGetHistogramParameteriv = goglGetProcAddress("glGetHistogramParameteriv");
// 	if(ptrglGetHistogramParameteriv == NULL) return 1;
// 	ptrglGetMinmax = goglGetProcAddress("glGetMinmax");
// 	if(ptrglGetMinmax == NULL) return 1;
// 	ptrglGetMinmaxParameterfv = goglGetProcAddress("glGetMinmaxParameterfv");
// 	if(ptrglGetMinmaxParameterfv == NULL) return 1;
// 	ptrglGetMinmaxParameteriv = goglGetProcAddress("glGetMinmaxParameteriv");
// 	if(ptrglGetMinmaxParameteriv == NULL) return 1;
// 	ptrglHistogram = goglGetProcAddress("glHistogram");
// 	if(ptrglHistogram == NULL) return 1;
// 	ptrglMinmax = goglGetProcAddress("glMinmax");
// 	if(ptrglMinmax == NULL) return 1;
// 	ptrglResetHistogram = goglGetProcAddress("glResetHistogram");
// 	if(ptrglResetHistogram == NULL) return 1;
// 	ptrglResetMinmax = goglGetProcAddress("glResetMinmax");
// 	if(ptrglResetMinmax == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_3() {
// 	ptrglActiveTexture = goglGetProcAddress("glActiveTexture");
// 	if(ptrglActiveTexture == NULL) return 1;
// 	ptrglSampleCoverage = goglGetProcAddress("glSampleCoverage");
// 	if(ptrglSampleCoverage == NULL) return 1;
// 	ptrglCompressedTexImage3D = goglGetProcAddress("glCompressedTexImage3D");
// 	if(ptrglCompressedTexImage3D == NULL) return 1;
// 	ptrglCompressedTexImage2D = goglGetProcAddress("glCompressedTexImage2D");
// 	if(ptrglCompressedTexImage2D == NULL) return 1;
// 	ptrglCompressedTexImage1D = goglGetProcAddress("glCompressedTexImage1D");
// 	if(ptrglCompressedTexImage1D == NULL) return 1;
// 	ptrglCompressedTexSubImage3D = goglGetProcAddress("glCompressedTexSubImage3D");
// 	if(ptrglCompressedTexSubImage3D == NULL) return 1;
// 	ptrglCompressedTexSubImage2D = goglGetProcAddress("glCompressedTexSubImage2D");
// 	if(ptrglCompressedTexSubImage2D == NULL) return 1;
// 	ptrglCompressedTexSubImage1D = goglGetProcAddress("glCompressedTexSubImage1D");
// 	if(ptrglCompressedTexSubImage1D == NULL) return 1;
// 	ptrglGetCompressedTexImage = goglGetProcAddress("glGetCompressedTexImage");
// 	if(ptrglGetCompressedTexImage == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_3_DEPRECATED() {
// 	ptrglClientActiveTexture = goglGetProcAddress("glClientActiveTexture");
// 	if(ptrglClientActiveTexture == NULL) return 1;
// 	ptrglMultiTexCoord1d = goglGetProcAddress("glMultiTexCoord1d");
// 	if(ptrglMultiTexCoord1d == NULL) return 1;
// 	ptrglMultiTexCoord1dv = goglGetProcAddress("glMultiTexCoord1dv");
// 	if(ptrglMultiTexCoord1dv == NULL) return 1;
// 	ptrglMultiTexCoord1f = goglGetProcAddress("glMultiTexCoord1f");
// 	if(ptrglMultiTexCoord1f == NULL) return 1;
// 	ptrglMultiTexCoord1fv = goglGetProcAddress("glMultiTexCoord1fv");
// 	if(ptrglMultiTexCoord1fv == NULL) return 1;
// 	ptrglMultiTexCoord1i = goglGetProcAddress("glMultiTexCoord1i");
// 	if(ptrglMultiTexCoord1i == NULL) return 1;
// 	ptrglMultiTexCoord1iv = goglGetProcAddress("glMultiTexCoord1iv");
// 	if(ptrglMultiTexCoord1iv == NULL) return 1;
// 	ptrglMultiTexCoord1s = goglGetProcAddress("glMultiTexCoord1s");
// 	if(ptrglMultiTexCoord1s == NULL) return 1;
// 	ptrglMultiTexCoord1sv = goglGetProcAddress("glMultiTexCoord1sv");
// 	if(ptrglMultiTexCoord1sv == NULL) return 1;
// 	ptrglMultiTexCoord2d = goglGetProcAddress("glMultiTexCoord2d");
// 	if(ptrglMultiTexCoord2d == NULL) return 1;
// 	ptrglMultiTexCoord2dv = goglGetProcAddress("glMultiTexCoord2dv");
// 	if(ptrglMultiTexCoord2dv == NULL) return 1;
// 	ptrglMultiTexCoord2f = goglGetProcAddress("glMultiTexCoord2f");
// 	if(ptrglMultiTexCoord2f == NULL) return 1;
// 	ptrglMultiTexCoord2fv = goglGetProcAddress("glMultiTexCoord2fv");
// 	if(ptrglMultiTexCoord2fv == NULL) return 1;
// 	ptrglMultiTexCoord2i = goglGetProcAddress("glMultiTexCoord2i");
// 	if(ptrglMultiTexCoord2i == NULL) return 1;
// 	ptrglMultiTexCoord2iv = goglGetProcAddress("glMultiTexCoord2iv");
// 	if(ptrglMultiTexCoord2iv == NULL) return 1;
// 	ptrglMultiTexCoord2s = goglGetProcAddress("glMultiTexCoord2s");
// 	if(ptrglMultiTexCoord2s == NULL) return 1;
// 	ptrglMultiTexCoord2sv = goglGetProcAddress("glMultiTexCoord2sv");
// 	if(ptrglMultiTexCoord2sv == NULL) return 1;
// 	ptrglMultiTexCoord3d = goglGetProcAddress("glMultiTexCoord3d");
// 	if(ptrglMultiTexCoord3d == NULL) return 1;
// 	ptrglMultiTexCoord3dv = goglGetProcAddress("glMultiTexCoord3dv");
// 	if(ptrglMultiTexCoord3dv == NULL) return 1;
// 	ptrglMultiTexCoord3f = goglGetProcAddress("glMultiTexCoord3f");
// 	if(ptrglMultiTexCoord3f == NULL) return 1;
// 	ptrglMultiTexCoord3fv = goglGetProcAddress("glMultiTexCoord3fv");
// 	if(ptrglMultiTexCoord3fv == NULL) return 1;
// 	ptrglMultiTexCoord3i = goglGetProcAddress("glMultiTexCoord3i");
// 	if(ptrglMultiTexCoord3i == NULL) return 1;
// 	ptrglMultiTexCoord3iv = goglGetProcAddress("glMultiTexCoord3iv");
// 	if(ptrglMultiTexCoord3iv == NULL) return 1;
// 	ptrglMultiTexCoord3s = goglGetProcAddress("glMultiTexCoord3s");
// 	if(ptrglMultiTexCoord3s == NULL) return 1;
// 	ptrglMultiTexCoord3sv = goglGetProcAddress("glMultiTexCoord3sv");
// 	if(ptrglMultiTexCoord3sv == NULL) return 1;
// 	ptrglMultiTexCoord4d = goglGetProcAddress("glMultiTexCoord4d");
// 	if(ptrglMultiTexCoord4d == NULL) return 1;
// 	ptrglMultiTexCoord4dv = goglGetProcAddress("glMultiTexCoord4dv");
// 	if(ptrglMultiTexCoord4dv == NULL) return 1;
// 	ptrglMultiTexCoord4f = goglGetProcAddress("glMultiTexCoord4f");
// 	if(ptrglMultiTexCoord4f == NULL) return 1;
// 	ptrglMultiTexCoord4fv = goglGetProcAddress("glMultiTexCoord4fv");
// 	if(ptrglMultiTexCoord4fv == NULL) return 1;
// 	ptrglMultiTexCoord4i = goglGetProcAddress("glMultiTexCoord4i");
// 	if(ptrglMultiTexCoord4i == NULL) return 1;
// 	ptrglMultiTexCoord4iv = goglGetProcAddress("glMultiTexCoord4iv");
// 	if(ptrglMultiTexCoord4iv == NULL) return 1;
// 	ptrglMultiTexCoord4s = goglGetProcAddress("glMultiTexCoord4s");
// 	if(ptrglMultiTexCoord4s == NULL) return 1;
// 	ptrglMultiTexCoord4sv = goglGetProcAddress("glMultiTexCoord4sv");
// 	if(ptrglMultiTexCoord4sv == NULL) return 1;
// 	ptrglLoadTransposeMatrixf = goglGetProcAddress("glLoadTransposeMatrixf");
// 	if(ptrglLoadTransposeMatrixf == NULL) return 1;
// 	ptrglLoadTransposeMatrixd = goglGetProcAddress("glLoadTransposeMatrixd");
// 	if(ptrglLoadTransposeMatrixd == NULL) return 1;
// 	ptrglMultTransposeMatrixf = goglGetProcAddress("glMultTransposeMatrixf");
// 	if(ptrglMultTransposeMatrixf == NULL) return 1;
// 	ptrglMultTransposeMatrixd = goglGetProcAddress("glMultTransposeMatrixd");
// 	if(ptrglMultTransposeMatrixd == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_4() {
// 	ptrglBlendFuncSeparate = goglGetProcAddress("glBlendFuncSeparate");
// 	if(ptrglBlendFuncSeparate == NULL) return 1;
// 	ptrglMultiDrawArrays = goglGetProcAddress("glMultiDrawArrays");
// 	if(ptrglMultiDrawArrays == NULL) return 1;
// 	ptrglMultiDrawElements = goglGetProcAddress("glMultiDrawElements");
// 	if(ptrglMultiDrawElements == NULL) return 1;
// 	ptrglPointParameterf = goglGetProcAddress("glPointParameterf");
// 	if(ptrglPointParameterf == NULL) return 1;
// 	ptrglPointParameterfv = goglGetProcAddress("glPointParameterfv");
// 	if(ptrglPointParameterfv == NULL) return 1;
// 	ptrglPointParameteri = goglGetProcAddress("glPointParameteri");
// 	if(ptrglPointParameteri == NULL) return 1;
// 	ptrglPointParameteriv = goglGetProcAddress("glPointParameteriv");
// 	if(ptrglPointParameteriv == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_4_DEPRECATED() {
// 	ptrglFogCoordf = goglGetProcAddress("glFogCoordf");
// 	if(ptrglFogCoordf == NULL) return 1;
// 	ptrglFogCoordfv = goglGetProcAddress("glFogCoordfv");
// 	if(ptrglFogCoordfv == NULL) return 1;
// 	ptrglFogCoordd = goglGetProcAddress("glFogCoordd");
// 	if(ptrglFogCoordd == NULL) return 1;
// 	ptrglFogCoorddv = goglGetProcAddress("glFogCoorddv");
// 	if(ptrglFogCoorddv == NULL) return 1;
// 	ptrglFogCoordPointer = goglGetProcAddress("glFogCoordPointer");
// 	if(ptrglFogCoordPointer == NULL) return 1;
// 	ptrglSecondaryColor3b = goglGetProcAddress("glSecondaryColor3b");
// 	if(ptrglSecondaryColor3b == NULL) return 1;
// 	ptrglSecondaryColor3bv = goglGetProcAddress("glSecondaryColor3bv");
// 	if(ptrglSecondaryColor3bv == NULL) return 1;
// 	ptrglSecondaryColor3d = goglGetProcAddress("glSecondaryColor3d");
// 	if(ptrglSecondaryColor3d == NULL) return 1;
// 	ptrglSecondaryColor3dv = goglGetProcAddress("glSecondaryColor3dv");
// 	if(ptrglSecondaryColor3dv == NULL) return 1;
// 	ptrglSecondaryColor3f = goglGetProcAddress("glSecondaryColor3f");
// 	if(ptrglSecondaryColor3f == NULL) return 1;
// 	ptrglSecondaryColor3fv = goglGetProcAddress("glSecondaryColor3fv");
// 	if(ptrglSecondaryColor3fv == NULL) return 1;
// 	ptrglSecondaryColor3i = goglGetProcAddress("glSecondaryColor3i");
// 	if(ptrglSecondaryColor3i == NULL) return 1;
// 	ptrglSecondaryColor3iv = goglGetProcAddress("glSecondaryColor3iv");
// 	if(ptrglSecondaryColor3iv == NULL) return 1;
// 	ptrglSecondaryColor3s = goglGetProcAddress("glSecondaryColor3s");
// 	if(ptrglSecondaryColor3s == NULL) return 1;
// 	ptrglSecondaryColor3sv = goglGetProcAddress("glSecondaryColor3sv");
// 	if(ptrglSecondaryColor3sv == NULL) return 1;
// 	ptrglSecondaryColor3ub = goglGetProcAddress("glSecondaryColor3ub");
// 	if(ptrglSecondaryColor3ub == NULL) return 1;
// 	ptrglSecondaryColor3ubv = goglGetProcAddress("glSecondaryColor3ubv");
// 	if(ptrglSecondaryColor3ubv == NULL) return 1;
// 	ptrglSecondaryColor3ui = goglGetProcAddress("glSecondaryColor3ui");
// 	if(ptrglSecondaryColor3ui == NULL) return 1;
// 	ptrglSecondaryColor3uiv = goglGetProcAddress("glSecondaryColor3uiv");
// 	if(ptrglSecondaryColor3uiv == NULL) return 1;
// 	ptrglSecondaryColor3us = goglGetProcAddress("glSecondaryColor3us");
// 	if(ptrglSecondaryColor3us == NULL) return 1;
// 	ptrglSecondaryColor3usv = goglGetProcAddress("glSecondaryColor3usv");
// 	if(ptrglSecondaryColor3usv == NULL) return 1;
// 	ptrglSecondaryColorPointer = goglGetProcAddress("glSecondaryColorPointer");
// 	if(ptrglSecondaryColorPointer == NULL) return 1;
// 	ptrglWindowPos2d = goglGetProcAddress("glWindowPos2d");
// 	if(ptrglWindowPos2d == NULL) return 1;
// 	ptrglWindowPos2dv = goglGetProcAddress("glWindowPos2dv");
// 	if(ptrglWindowPos2dv == NULL) return 1;
// 	ptrglWindowPos2f = goglGetProcAddress("glWindowPos2f");
// 	if(ptrglWindowPos2f == NULL) return 1;
// 	ptrglWindowPos2fv = goglGetProcAddress("glWindowPos2fv");
// 	if(ptrglWindowPos2fv == NULL) return 1;
// 	ptrglWindowPos2i = goglGetProcAddress("glWindowPos2i");
// 	if(ptrglWindowPos2i == NULL) return 1;
// 	ptrglWindowPos2iv = goglGetProcAddress("glWindowPos2iv");
// 	if(ptrglWindowPos2iv == NULL) return 1;
// 	ptrglWindowPos2s = goglGetProcAddress("glWindowPos2s");
// 	if(ptrglWindowPos2s == NULL) return 1;
// 	ptrglWindowPos2sv = goglGetProcAddress("glWindowPos2sv");
// 	if(ptrglWindowPos2sv == NULL) return 1;
// 	ptrglWindowPos3d = goglGetProcAddress("glWindowPos3d");
// 	if(ptrglWindowPos3d == NULL) return 1;
// 	ptrglWindowPos3dv = goglGetProcAddress("glWindowPos3dv");
// 	if(ptrglWindowPos3dv == NULL) return 1;
// 	ptrglWindowPos3f = goglGetProcAddress("glWindowPos3f");
// 	if(ptrglWindowPos3f == NULL) return 1;
// 	ptrglWindowPos3fv = goglGetProcAddress("glWindowPos3fv");
// 	if(ptrglWindowPos3fv == NULL) return 1;
// 	ptrglWindowPos3i = goglGetProcAddress("glWindowPos3i");
// 	if(ptrglWindowPos3i == NULL) return 1;
// 	ptrglWindowPos3iv = goglGetProcAddress("glWindowPos3iv");
// 	if(ptrglWindowPos3iv == NULL) return 1;
// 	ptrglWindowPos3s = goglGetProcAddress("glWindowPos3s");
// 	if(ptrglWindowPos3s == NULL) return 1;
// 	ptrglWindowPos3sv = goglGetProcAddress("glWindowPos3sv");
// 	if(ptrglWindowPos3sv == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_5() {
// 	ptrglGenQueries = goglGetProcAddress("glGenQueries");
// 	if(ptrglGenQueries == NULL) return 1;
// 	ptrglDeleteQueries = goglGetProcAddress("glDeleteQueries");
// 	if(ptrglDeleteQueries == NULL) return 1;
// 	ptrglIsQuery = goglGetProcAddress("glIsQuery");
// 	if(ptrglIsQuery == NULL) return 1;
// 	ptrglBeginQuery = goglGetProcAddress("glBeginQuery");
// 	if(ptrglBeginQuery == NULL) return 1;
// 	ptrglEndQuery = goglGetProcAddress("glEndQuery");
// 	if(ptrglEndQuery == NULL) return 1;
// 	ptrglGetQueryiv = goglGetProcAddress("glGetQueryiv");
// 	if(ptrglGetQueryiv == NULL) return 1;
// 	ptrglGetQueryObjectiv = goglGetProcAddress("glGetQueryObjectiv");
// 	if(ptrglGetQueryObjectiv == NULL) return 1;
// 	ptrglGetQueryObjectuiv = goglGetProcAddress("glGetQueryObjectuiv");
// 	if(ptrglGetQueryObjectuiv == NULL) return 1;
// 	ptrglBindBuffer = goglGetProcAddress("glBindBuffer");
// 	if(ptrglBindBuffer == NULL) return 1;
// 	ptrglDeleteBuffers = goglGetProcAddress("glDeleteBuffers");
// 	if(ptrglDeleteBuffers == NULL) return 1;
// 	ptrglGenBuffers = goglGetProcAddress("glGenBuffers");
// 	if(ptrglGenBuffers == NULL) return 1;
// 	ptrglIsBuffer = goglGetProcAddress("glIsBuffer");
// 	if(ptrglIsBuffer == NULL) return 1;
// 	ptrglBufferData = goglGetProcAddress("glBufferData");
// 	if(ptrglBufferData == NULL) return 1;
// 	ptrglBufferSubData = goglGetProcAddress("glBufferSubData");
// 	if(ptrglBufferSubData == NULL) return 1;
// 	ptrglGetBufferSubData = goglGetProcAddress("glGetBufferSubData");
// 	if(ptrglGetBufferSubData == NULL) return 1;
// 	ptrglMapBuffer = goglGetProcAddress("glMapBuffer");
// 	if(ptrglMapBuffer == NULL) return 1;
// 	ptrglUnmapBuffer = goglGetProcAddress("glUnmapBuffer");
// 	if(ptrglUnmapBuffer == NULL) return 1;
// 	ptrglGetBufferParameteriv = goglGetProcAddress("glGetBufferParameteriv");
// 	if(ptrglGetBufferParameteriv == NULL) return 1;
// 	ptrglGetBufferPointerv = goglGetProcAddress("glGetBufferPointerv");
// 	if(ptrglGetBufferPointerv == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_2_0() {
// 	ptrglBlendEquationSeparate = goglGetProcAddress("glBlendEquationSeparate");
// 	if(ptrglBlendEquationSeparate == NULL) return 1;
// 	ptrglDrawBuffers = goglGetProcAddress("glDrawBuffers");
// 	if(ptrglDrawBuffers == NULL) return 1;
// 	ptrglStencilOpSeparate = goglGetProcAddress("glStencilOpSeparate");
// 	if(ptrglStencilOpSeparate == NULL) return 1;
// 	ptrglStencilFuncSeparate = goglGetProcAddress("glStencilFuncSeparate");
// 	if(ptrglStencilFuncSeparate == NULL) return 1;
// 	ptrglStencilMaskSeparate = goglGetProcAddress("glStencilMaskSeparate");
// 	if(ptrglStencilMaskSeparate == NULL) return 1;
// 	ptrglAttachShader = goglGetProcAddress("glAttachShader");
// 	if(ptrglAttachShader == NULL) return 1;
// 	ptrglBindAttribLocation = goglGetProcAddress("glBindAttribLocation");
// 	if(ptrglBindAttribLocation == NULL) return 1;
// 	ptrglCompileShader = goglGetProcAddress("glCompileShader");
// 	if(ptrglCompileShader == NULL) return 1;
// 	ptrglCreateProgram = goglGetProcAddress("glCreateProgram");
// 	if(ptrglCreateProgram == NULL) return 1;
// 	ptrglCreateShader = goglGetProcAddress("glCreateShader");
// 	if(ptrglCreateShader == NULL) return 1;
// 	ptrglDeleteProgram = goglGetProcAddress("glDeleteProgram");
// 	if(ptrglDeleteProgram == NULL) return 1;
// 	ptrglDeleteShader = goglGetProcAddress("glDeleteShader");
// 	if(ptrglDeleteShader == NULL) return 1;
// 	ptrglDetachShader = goglGetProcAddress("glDetachShader");
// 	if(ptrglDetachShader == NULL) return 1;
// 	ptrglDisableVertexAttribArray = goglGetProcAddress("glDisableVertexAttribArray");
// 	if(ptrglDisableVertexAttribArray == NULL) return 1;
// 	ptrglEnableVertexAttribArray = goglGetProcAddress("glEnableVertexAttribArray");
// 	if(ptrglEnableVertexAttribArray == NULL) return 1;
// 	ptrglGetActiveAttrib = goglGetProcAddress("glGetActiveAttrib");
// 	if(ptrglGetActiveAttrib == NULL) return 1;
// 	ptrglGetActiveUniform = goglGetProcAddress("glGetActiveUniform");
// 	if(ptrglGetActiveUniform == NULL) return 1;
// 	ptrglGetAttachedShaders = goglGetProcAddress("glGetAttachedShaders");
// 	if(ptrglGetAttachedShaders == NULL) return 1;
// 	ptrglGetAttribLocation = goglGetProcAddress("glGetAttribLocation");
// 	if(ptrglGetAttribLocation == NULL) return 1;
// 	ptrglGetProgramiv = goglGetProcAddress("glGetProgramiv");
// 	if(ptrglGetProgramiv == NULL) return 1;
// 	ptrglGetProgramInfoLog = goglGetProcAddress("glGetProgramInfoLog");
// 	if(ptrglGetProgramInfoLog == NULL) return 1;
// 	ptrglGetShaderiv = goglGetProcAddress("glGetShaderiv");
// 	if(ptrglGetShaderiv == NULL) return 1;
// 	ptrglGetShaderInfoLog = goglGetProcAddress("glGetShaderInfoLog");
// 	if(ptrglGetShaderInfoLog == NULL) return 1;
// 	ptrglGetShaderSource = goglGetProcAddress("glGetShaderSource");
// 	if(ptrglGetShaderSource == NULL) return 1;
// 	ptrglGetUniformLocation = goglGetProcAddress("glGetUniformLocation");
// 	if(ptrglGetUniformLocation == NULL) return 1;
// 	ptrglGetUniformfv = goglGetProcAddress("glGetUniformfv");
// 	if(ptrglGetUniformfv == NULL) return 1;
// 	ptrglGetUniformiv = goglGetProcAddress("glGetUniformiv");
// 	if(ptrglGetUniformiv == NULL) return 1;
// 	ptrglGetVertexAttribdv = goglGetProcAddress("glGetVertexAttribdv");
// 	if(ptrglGetVertexAttribdv == NULL) return 1;
// 	ptrglGetVertexAttribfv = goglGetProcAddress("glGetVertexAttribfv");
// 	if(ptrglGetVertexAttribfv == NULL) return 1;
// 	ptrglGetVertexAttribiv = goglGetProcAddress("glGetVertexAttribiv");
// 	if(ptrglGetVertexAttribiv == NULL) return 1;
// 	ptrglGetVertexAttribPointerv = goglGetProcAddress("glGetVertexAttribPointerv");
// 	if(ptrglGetVertexAttribPointerv == NULL) return 1;
// 	ptrglIsProgram = goglGetProcAddress("glIsProgram");
// 	if(ptrglIsProgram == NULL) return 1;
// 	ptrglIsShader = goglGetProcAddress("glIsShader");
// 	if(ptrglIsShader == NULL) return 1;
// 	ptrglLinkProgram = goglGetProcAddress("glLinkProgram");
// 	if(ptrglLinkProgram == NULL) return 1;
// 	ptrglShaderSource = goglGetProcAddress("glShaderSource");
// 	if(ptrglShaderSource == NULL) return 1;
// 	ptrglUseProgram = goglGetProcAddress("glUseProgram");
// 	if(ptrglUseProgram == NULL) return 1;
// 	ptrglUniform1f = goglGetProcAddress("glUniform1f");
// 	if(ptrglUniform1f == NULL) return 1;
// 	ptrglUniform2f = goglGetProcAddress("glUniform2f");
// 	if(ptrglUniform2f == NULL) return 1;
// 	ptrglUniform3f = goglGetProcAddress("glUniform3f");
// 	if(ptrglUniform3f == NULL) return 1;
// 	ptrglUniform4f = goglGetProcAddress("glUniform4f");
// 	if(ptrglUniform4f == NULL) return 1;
// 	ptrglUniform1i = goglGetProcAddress("glUniform1i");
// 	if(ptrglUniform1i == NULL) return 1;
// 	ptrglUniform2i = goglGetProcAddress("glUniform2i");
// 	if(ptrglUniform2i == NULL) return 1;
// 	ptrglUniform3i = goglGetProcAddress("glUniform3i");
// 	if(ptrglUniform3i == NULL) return 1;
// 	ptrglUniform4i = goglGetProcAddress("glUniform4i");
// 	if(ptrglUniform4i == NULL) return 1;
// 	ptrglUniform1fv = goglGetProcAddress("glUniform1fv");
// 	if(ptrglUniform1fv == NULL) return 1;
// 	ptrglUniform2fv = goglGetProcAddress("glUniform2fv");
// 	if(ptrglUniform2fv == NULL) return 1;
// 	ptrglUniform3fv = goglGetProcAddress("glUniform3fv");
// 	if(ptrglUniform3fv == NULL) return 1;
// 	ptrglUniform4fv = goglGetProcAddress("glUniform4fv");
// 	if(ptrglUniform4fv == NULL) return 1;
// 	ptrglUniform1iv = goglGetProcAddress("glUniform1iv");
// 	if(ptrglUniform1iv == NULL) return 1;
// 	ptrglUniform2iv = goglGetProcAddress("glUniform2iv");
// 	if(ptrglUniform2iv == NULL) return 1;
// 	ptrglUniform3iv = goglGetProcAddress("glUniform3iv");
// 	if(ptrglUniform3iv == NULL) return 1;
// 	ptrglUniform4iv = goglGetProcAddress("glUniform4iv");
// 	if(ptrglUniform4iv == NULL) return 1;
// 	ptrglUniformMatrix2fv = goglGetProcAddress("glUniformMatrix2fv");
// 	if(ptrglUniformMatrix2fv == NULL) return 1;
// 	ptrglUniformMatrix3fv = goglGetProcAddress("glUniformMatrix3fv");
// 	if(ptrglUniformMatrix3fv == NULL) return 1;
// 	ptrglUniformMatrix4fv = goglGetProcAddress("glUniformMatrix4fv");
// 	if(ptrglUniformMatrix4fv == NULL) return 1;
// 	ptrglValidateProgram = goglGetProcAddress("glValidateProgram");
// 	if(ptrglValidateProgram == NULL) return 1;
// 	ptrglVertexAttrib1d = goglGetProcAddress("glVertexAttrib1d");
// 	if(ptrglVertexAttrib1d == NULL) return 1;
// 	ptrglVertexAttrib1dv = goglGetProcAddress("glVertexAttrib1dv");
// 	if(ptrglVertexAttrib1dv == NULL) return 1;
// 	ptrglVertexAttrib1f = goglGetProcAddress("glVertexAttrib1f");
// 	if(ptrglVertexAttrib1f == NULL) return 1;
// 	ptrglVertexAttrib1fv = goglGetProcAddress("glVertexAttrib1fv");
// 	if(ptrglVertexAttrib1fv == NULL) return 1;
// 	ptrglVertexAttrib1s = goglGetProcAddress("glVertexAttrib1s");
// 	if(ptrglVertexAttrib1s == NULL) return 1;
// 	ptrglVertexAttrib1sv = goglGetProcAddress("glVertexAttrib1sv");
// 	if(ptrglVertexAttrib1sv == NULL) return 1;
// 	ptrglVertexAttrib2d = goglGetProcAddress("glVertexAttrib2d");
// 	if(ptrglVertexAttrib2d == NULL) return 1;
// 	ptrglVertexAttrib2dv = goglGetProcAddress("glVertexAttrib2dv");
// 	if(ptrglVertexAttrib2dv == NULL) return 1;
// 	ptrglVertexAttrib2f = goglGetProcAddress("glVertexAttrib2f");
// 	if(ptrglVertexAttrib2f == NULL) return 1;
// 	ptrglVertexAttrib2fv = goglGetProcAddress("glVertexAttrib2fv");
// 	if(ptrglVertexAttrib2fv == NULL) return 1;
// 	ptrglVertexAttrib2s = goglGetProcAddress("glVertexAttrib2s");
// 	if(ptrglVertexAttrib2s == NULL) return 1;
// 	ptrglVertexAttrib2sv = goglGetProcAddress("glVertexAttrib2sv");
// 	if(ptrglVertexAttrib2sv == NULL) return 1;
// 	ptrglVertexAttrib3d = goglGetProcAddress("glVertexAttrib3d");
// 	if(ptrglVertexAttrib3d == NULL) return 1;
// 	ptrglVertexAttrib3dv = goglGetProcAddress("glVertexAttrib3dv");
// 	if(ptrglVertexAttrib3dv == NULL) return 1;
// 	ptrglVertexAttrib3f = goglGetProcAddress("glVertexAttrib3f");
// 	if(ptrglVertexAttrib3f == NULL) return 1;
// 	ptrglVertexAttrib3fv = goglGetProcAddress("glVertexAttrib3fv");
// 	if(ptrglVertexAttrib3fv == NULL) return 1;
// 	ptrglVertexAttrib3s = goglGetProcAddress("glVertexAttrib3s");
// 	if(ptrglVertexAttrib3s == NULL) return 1;
// 	ptrglVertexAttrib3sv = goglGetProcAddress("glVertexAttrib3sv");
// 	if(ptrglVertexAttrib3sv == NULL) return 1;
// 	ptrglVertexAttrib4Nbv = goglGetProcAddress("glVertexAttrib4Nbv");
// 	if(ptrglVertexAttrib4Nbv == NULL) return 1;
// 	ptrglVertexAttrib4Niv = goglGetProcAddress("glVertexAttrib4Niv");
// 	if(ptrglVertexAttrib4Niv == NULL) return 1;
// 	ptrglVertexAttrib4Nsv = goglGetProcAddress("glVertexAttrib4Nsv");
// 	if(ptrglVertexAttrib4Nsv == NULL) return 1;
// 	ptrglVertexAttrib4Nub = goglGetProcAddress("glVertexAttrib4Nub");
// 	if(ptrglVertexAttrib4Nub == NULL) return 1;
// 	ptrglVertexAttrib4Nubv = goglGetProcAddress("glVertexAttrib4Nubv");
// 	if(ptrglVertexAttrib4Nubv == NULL) return 1;
// 	ptrglVertexAttrib4Nuiv = goglGetProcAddress("glVertexAttrib4Nuiv");
// 	if(ptrglVertexAttrib4Nuiv == NULL) return 1;
// 	ptrglVertexAttrib4Nusv = goglGetProcAddress("glVertexAttrib4Nusv");
// 	if(ptrglVertexAttrib4Nusv == NULL) return 1;
// 	ptrglVertexAttrib4bv = goglGetProcAddress("glVertexAttrib4bv");
// 	if(ptrglVertexAttrib4bv == NULL) return 1;
// 	ptrglVertexAttrib4d = goglGetProcAddress("glVertexAttrib4d");
// 	if(ptrglVertexAttrib4d == NULL) return 1;
// 	ptrglVertexAttrib4dv = goglGetProcAddress("glVertexAttrib4dv");
// 	if(ptrglVertexAttrib4dv == NULL) return 1;
// 	ptrglVertexAttrib4f = goglGetProcAddress("glVertexAttrib4f");
// 	if(ptrglVertexAttrib4f == NULL) return 1;
// 	ptrglVertexAttrib4fv = goglGetProcAddress("glVertexAttrib4fv");
// 	if(ptrglVertexAttrib4fv == NULL) return 1;
// 	ptrglVertexAttrib4iv = goglGetProcAddress("glVertexAttrib4iv");
// 	if(ptrglVertexAttrib4iv == NULL) return 1;
// 	ptrglVertexAttrib4s = goglGetProcAddress("glVertexAttrib4s");
// 	if(ptrglVertexAttrib4s == NULL) return 1;
// 	ptrglVertexAttrib4sv = goglGetProcAddress("glVertexAttrib4sv");
// 	if(ptrglVertexAttrib4sv == NULL) return 1;
// 	ptrglVertexAttrib4ubv = goglGetProcAddress("glVertexAttrib4ubv");
// 	if(ptrglVertexAttrib4ubv == NULL) return 1;
// 	ptrglVertexAttrib4uiv = goglGetProcAddress("glVertexAttrib4uiv");
// 	if(ptrglVertexAttrib4uiv == NULL) return 1;
// 	ptrglVertexAttrib4usv = goglGetProcAddress("glVertexAttrib4usv");
// 	if(ptrglVertexAttrib4usv == NULL) return 1;
// 	ptrglVertexAttribPointer = goglGetProcAddress("glVertexAttribPointer");
// 	if(ptrglVertexAttribPointer == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_2_1() {
// 	ptrglUniformMatrix2x3fv = goglGetProcAddress("glUniformMatrix2x3fv");
// 	if(ptrglUniformMatrix2x3fv == NULL) return 1;
// 	ptrglUniformMatrix3x2fv = goglGetProcAddress("glUniformMatrix3x2fv");
// 	if(ptrglUniformMatrix3x2fv == NULL) return 1;
// 	ptrglUniformMatrix2x4fv = goglGetProcAddress("glUniformMatrix2x4fv");
// 	if(ptrglUniformMatrix2x4fv == NULL) return 1;
// 	ptrglUniformMatrix4x2fv = goglGetProcAddress("glUniformMatrix4x2fv");
// 	if(ptrglUniformMatrix4x2fv == NULL) return 1;
// 	ptrglUniformMatrix3x4fv = goglGetProcAddress("glUniformMatrix3x4fv");
// 	if(ptrglUniformMatrix3x4fv == NULL) return 1;
// 	ptrglUniformMatrix4x3fv = goglGetProcAddress("glUniformMatrix4x3fv");
// 	if(ptrglUniformMatrix4x3fv == NULL) return 1;
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

// VERSION_1_1
const (
	ALPHA = 0x1906
	ALWAYS = 0x0207
	AND = 0x1501
	AND_INVERTED = 0x1504
	AND_REVERSE = 0x1502
	BACK = 0x0405
	BACK_LEFT = 0x0402
	BACK_RIGHT = 0x0403
	BLEND = 0x0BE2
	BLEND_DST = 0x0BE0
	BLEND_SRC = 0x0BE1
	BLUE = 0x1905
	BYTE = 0x1400
	CCW = 0x0901
	CLEAR = 0x1500
	COLOR = 0x1800
	COLOR_BUFFER_BIT = 0x00004000
	COLOR_CLEAR_VALUE = 0x0C22
	COLOR_LOGIC_OP = 0x0BF2
	COLOR_WRITEMASK = 0x0C23
	COPY = 0x1503
	COPY_INVERTED = 0x150C
	CULL_FACE = 0x0B44
	CULL_FACE_MODE = 0x0B45
	CW = 0x0900
	DECR = 0x1E03
	DEPTH = 0x1801
	DEPTH_BUFFER_BIT = 0x00000100
	DEPTH_CLEAR_VALUE = 0x0B73
	DEPTH_COMPONENT = 0x1902
	DEPTH_FUNC = 0x0B74
	DEPTH_RANGE = 0x0B70
	DEPTH_TEST = 0x0B71
	DEPTH_WRITEMASK = 0x0B72
	DITHER = 0x0BD0
	DONT_CARE = 0x1100
	DOUBLE = 0x140A
	DOUBLEBUFFER = 0x0C32
	DRAW_BUFFER = 0x0C01
	DST_ALPHA = 0x0304
	DST_COLOR = 0x0306
	EQUAL = 0x0202
	EQUIV = 0x1509
	EXTENSIONS = 0x1F03
	FALSE = 0
	FASTEST = 0x1101
	FILL = 0x1B02
	FLOAT = 0x1406
	FRONT = 0x0404
	FRONT_AND_BACK = 0x0408
	FRONT_FACE = 0x0B46
	FRONT_LEFT = 0x0400
	FRONT_RIGHT = 0x0401
	GEQUAL = 0x0206
	GREATER = 0x0204
	GREEN = 0x1904
	INCR = 0x1E02
	INT = 0x1404
	INVALID_ENUM = 0x0500
	INVALID_OPERATION = 0x0502
	INVALID_VALUE = 0x0501
	INVERT = 0x150A
	KEEP = 0x1E00
	LEFT = 0x0406
	LEQUAL = 0x0203
	LESS = 0x0201
	LINE = 0x1B01
	LINEAR = 0x2601
	LINEAR_MIPMAP_LINEAR = 0x2703
	LINEAR_MIPMAP_NEAREST = 0x2701
	LINES = 0x0001
	LINE_LOOP = 0x0002
	LINE_SMOOTH = 0x0B20
	LINE_SMOOTH_HINT = 0x0C52
	LINE_STRIP = 0x0003
	LINE_WIDTH = 0x0B21
	LINE_WIDTH_GRANULARITY = 0x0B23
	LINE_WIDTH_RANGE = 0x0B22
	LOGIC_OP_MODE = 0x0BF0
	MAX_TEXTURE_SIZE = 0x0D33
	MAX_VIEWPORT_DIMS = 0x0D3A
	NAND = 0x150E
	NEAREST = 0x2600
	NEAREST_MIPMAP_LINEAR = 0x2702
	NEAREST_MIPMAP_NEAREST = 0x2700
	NEVER = 0x0200
	NICEST = 0x1102
	NONE = 0
	NOOP = 0x1505
	NOR = 0x1508
	NOTEQUAL = 0x0205
	NO_ERROR = 0
	ONE = 1
	ONE_MINUS_DST_ALPHA = 0x0305
	ONE_MINUS_DST_COLOR = 0x0307
	ONE_MINUS_SRC_ALPHA = 0x0303
	ONE_MINUS_SRC_COLOR = 0x0301
	OR = 0x1507
	OR_INVERTED = 0x150D
	OR_REVERSE = 0x150B
	OUT_OF_MEMORY = 0x0505
	PACK_ALIGNMENT = 0x0D05
	PACK_LSB_FIRST = 0x0D01
	PACK_ROW_LENGTH = 0x0D02
	PACK_SKIP_PIXELS = 0x0D04
	PACK_SKIP_ROWS = 0x0D03
	PACK_SWAP_BYTES = 0x0D00
	POINT = 0x1B00
	POINTS = 0x0000
	POINT_SIZE = 0x0B11
	POINT_SIZE_GRANULARITY = 0x0B13
	POINT_SIZE_RANGE = 0x0B12
	POLYGON_OFFSET_FACTOR = 0x8038
	POLYGON_OFFSET_FILL = 0x8037
	POLYGON_OFFSET_LINE = 0x2A02
	POLYGON_OFFSET_POINT = 0x2A01
	POLYGON_OFFSET_UNITS = 0x2A00
	POLYGON_SMOOTH = 0x0B41
	POLYGON_SMOOTH_HINT = 0x0C53
	PROXY_TEXTURE_1D = 0x8063
	PROXY_TEXTURE_2D = 0x8064
	R3_G3_B2 = 0x2A10
	READ_BUFFER = 0x0C02
	RED = 0x1903
	RENDERER = 0x1F01
	REPEAT = 0x2901
	REPLACE = 0x1E01
	RGB = 0x1907
	RGB10 = 0x8052
	RGB10_A2 = 0x8059
	RGB12 = 0x8053
	RGB16 = 0x8054
	RGB4 = 0x804F
	RGB5 = 0x8050
	RGB5_A1 = 0x8057
	RGB8 = 0x8051
	RGBA = 0x1908
	RGBA12 = 0x805A
	RGBA16 = 0x805B
	RGBA2 = 0x8055
	RGBA4 = 0x8056
	RGBA8 = 0x8058
	RIGHT = 0x0407
	SCISSOR_BOX = 0x0C10
	SCISSOR_TEST = 0x0C11
	SET = 0x150F
	SHORT = 0x1402
	SRC_ALPHA = 0x0302
	SRC_ALPHA_SATURATE = 0x0308
	SRC_COLOR = 0x0300
	STENCIL = 0x1802
	STENCIL_BUFFER_BIT = 0x00000400
	STENCIL_CLEAR_VALUE = 0x0B91
	STENCIL_FAIL = 0x0B94
	STENCIL_FUNC = 0x0B92
	STENCIL_INDEX = 0x1901
	STENCIL_PASS_DEPTH_FAIL = 0x0B95
	STENCIL_PASS_DEPTH_PASS = 0x0B96
	STENCIL_REF = 0x0B97
	STENCIL_TEST = 0x0B90
	STENCIL_VALUE_MASK = 0x0B93
	STENCIL_WRITEMASK = 0x0B98
	STEREO = 0x0C33
	SUBPIXEL_BITS = 0x0D50
	TEXTURE = 0x1702
	TEXTURE_1D = 0x0DE0
	TEXTURE_2D = 0x0DE1
	TEXTURE_ALPHA_SIZE = 0x805F
	TEXTURE_BINDING_1D = 0x8068
	TEXTURE_BINDING_2D = 0x8069
	TEXTURE_BLUE_SIZE = 0x805E
	TEXTURE_BORDER_COLOR = 0x1004
	TEXTURE_GREEN_SIZE = 0x805D
	TEXTURE_HEIGHT = 0x1001
	TEXTURE_INTERNAL_FORMAT = 0x1003
	TEXTURE_MAG_FILTER = 0x2800
	TEXTURE_MIN_FILTER = 0x2801
	TEXTURE_RED_SIZE = 0x805C
	TEXTURE_WIDTH = 0x1000
	TEXTURE_WRAP_S = 0x2802
	TEXTURE_WRAP_T = 0x2803
	TRIANGLES = 0x0004
	TRIANGLE_FAN = 0x0006
	TRIANGLE_STRIP = 0x0005
	TRUE = 1
	UNPACK_ALIGNMENT = 0x0CF5
	UNPACK_LSB_FIRST = 0x0CF1
	UNPACK_ROW_LENGTH = 0x0CF2
	UNPACK_SKIP_PIXELS = 0x0CF4
	UNPACK_SKIP_ROWS = 0x0CF3
	UNPACK_SWAP_BYTES = 0x0CF0
	UNSIGNED_BYTE = 0x1401
	UNSIGNED_INT = 0x1405
	UNSIGNED_SHORT = 0x1403
	VENDOR = 0x1F00
	VERSION = 0x1F02
	VIEWPORT = 0x0BA2
	XOR = 0x1506
	ZERO = 0
)
// VERSION_1_1_DEPRECATED
const (
	X2D = 0x0600
	X2_BYTES = 0x1407
	X3D = 0x0601
	X3D_COLOR = 0x0602
	X3D_COLOR_TEXTURE = 0x0603
	X3_BYTES = 0x1408
	X4D_COLOR_TEXTURE = 0x0604
	X4_BYTES = 0x1409
	ACCUM = 0x0100
	ACCUM_ALPHA_BITS = 0x0D5B
	ACCUM_BLUE_BITS = 0x0D5A
	ACCUM_BUFFER_BIT = 0x00000200
	ACCUM_CLEAR_VALUE = 0x0B80
	ACCUM_GREEN_BITS = 0x0D59
	ACCUM_RED_BITS = 0x0D58
	ADD = 0x0104
	ALL_ATTRIB_BITS = 0xFFFFFFFF
	ALPHA12 = 0x803D
	ALPHA16 = 0x803E
	ALPHA4 = 0x803B
	ALPHA8 = 0x803C
	ALPHA_BIAS = 0x0D1D
	ALPHA_BITS = 0x0D55
	ALPHA_SCALE = 0x0D1C
	ALPHA_TEST = 0x0BC0
	ALPHA_TEST_FUNC = 0x0BC1
	ALPHA_TEST_REF = 0x0BC2
	AMBIENT = 0x1200
	AMBIENT_AND_DIFFUSE = 0x1602
	ATTRIB_STACK_DEPTH = 0x0BB0
	AUTO_NORMAL = 0x0D80
	AUX0 = 0x0409
	AUX1 = 0x040A
	AUX2 = 0x040B
	AUX3 = 0x040C
	AUX_BUFFERS = 0x0C00
	BITMAP = 0x1A00
	BITMAP_TOKEN = 0x0704
	BLUE_BIAS = 0x0D1B
	BLUE_BITS = 0x0D54
	BLUE_SCALE = 0x0D1A
	C3F_V3F = 0x2A24
	C4F_N3F_V3F = 0x2A26
	C4UB_V2F = 0x2A22
	C4UB_V3F = 0x2A23
	CLAMP = 0x2900
	CLIENT_ALL_ATTRIB_BITS = 0xFFFFFFFF
	CLIENT_ATTRIB_STACK_DEPTH = 0x0BB1
	CLIENT_PIXEL_STORE_BIT = 0x00000001
	CLIENT_VERTEX_ARRAY_BIT = 0x00000002
	CLIP_PLANE0 = 0x3000
	CLIP_PLANE1 = 0x3001
	CLIP_PLANE2 = 0x3002
	CLIP_PLANE3 = 0x3003
	CLIP_PLANE4 = 0x3004
	CLIP_PLANE5 = 0x3005
	COEFF = 0x0A00
	COLOR_ARRAY = 0x8076
	COLOR_ARRAY_POINTER = 0x8090
	COLOR_ARRAY_SIZE = 0x8081
	COLOR_ARRAY_STRIDE = 0x8083
	COLOR_ARRAY_TYPE = 0x8082
	COLOR_INDEX = 0x1900
	COLOR_INDEXES = 0x1603
	COLOR_MATERIAL = 0x0B57
	COLOR_MATERIAL_FACE = 0x0B55
	COLOR_MATERIAL_PARAMETER = 0x0B56
	COMPILE = 0x1300
	COMPILE_AND_EXECUTE = 0x1301
	CONSTANT_ATTENUATION = 0x1207
	COPY_PIXEL_TOKEN = 0x0706
	CURRENT_BIT = 0x00000001
	CURRENT_COLOR = 0x0B00
	CURRENT_INDEX = 0x0B01
	CURRENT_NORMAL = 0x0B02
	CURRENT_RASTER_COLOR = 0x0B04
	CURRENT_RASTER_DISTANCE = 0x0B09
	CURRENT_RASTER_INDEX = 0x0B05
	CURRENT_RASTER_POSITION = 0x0B07
	CURRENT_RASTER_POSITION_VALID = 0x0B08
	CURRENT_RASTER_TEXTURE_COORDS = 0x0B06
	CURRENT_TEXTURE_COORDS = 0x0B03
	DECAL = 0x2101
	DEPTH_BIAS = 0x0D1F
	DEPTH_BITS = 0x0D56
	DEPTH_SCALE = 0x0D1E
	DIFFUSE = 0x1201
	DOMAIN = 0x0A02
	DRAW_PIXEL_TOKEN = 0x0705
	EDGE_FLAG = 0x0B43
	EDGE_FLAG_ARRAY = 0x8079
	EDGE_FLAG_ARRAY_POINTER = 0x8093
	EDGE_FLAG_ARRAY_STRIDE = 0x808C
	EMISSION = 0x1600
	ENABLE_BIT = 0x00002000
	EVAL_BIT = 0x00010000
	EXP = 0x0800
	EXP2 = 0x0801
	EYE_LINEAR = 0x2400
	EYE_PLANE = 0x2502
	FEEDBACK = 0x1C01
	FEEDBACK_BUFFER_POINTER = 0x0DF0
	FEEDBACK_BUFFER_SIZE = 0x0DF1
	FEEDBACK_BUFFER_TYPE = 0x0DF2
	FLAT = 0x1D00
	FOG = 0x0B60
	FOG_BIT = 0x00000080
	FOG_COLOR = 0x0B66
	FOG_DENSITY = 0x0B62
	FOG_END = 0x0B64
	FOG_HINT = 0x0C54
	FOG_INDEX = 0x0B61
	FOG_MODE = 0x0B65
	FOG_START = 0x0B63
	GREEN_BIAS = 0x0D19
	GREEN_BITS = 0x0D53
	GREEN_SCALE = 0x0D18
	HINT_BIT = 0x00008000
	INDEX_ARRAY = 0x8077
	INDEX_ARRAY_POINTER = 0x8091
	INDEX_ARRAY_STRIDE = 0x8086
	INDEX_ARRAY_TYPE = 0x8085
	INDEX_BITS = 0x0D51
	INDEX_CLEAR_VALUE = 0x0C20
	INDEX_LOGIC_OP = 0x0BF1
	INDEX_MODE = 0x0C30
	INDEX_OFFSET = 0x0D13
	INDEX_SHIFT = 0x0D12
	INDEX_WRITEMASK = 0x0C21
	INTENSITY = 0x8049
	INTENSITY12 = 0x804C
	INTENSITY16 = 0x804D
	INTENSITY4 = 0x804A
	INTENSITY8 = 0x804B
	LIGHT0 = 0x4000
	LIGHT1 = 0x4001
	LIGHT2 = 0x4002
	LIGHT3 = 0x4003
	LIGHT4 = 0x4004
	LIGHT5 = 0x4005
	LIGHT6 = 0x4006
	LIGHT7 = 0x4007
	LIGHTING = 0x0B50
	LIGHTING_BIT = 0x00000040
	LIGHT_MODEL_AMBIENT = 0x0B53
	LIGHT_MODEL_LOCAL_VIEWER = 0x0B51
	LIGHT_MODEL_TWO_SIDE = 0x0B52
	LINEAR_ATTENUATION = 0x1208
	LINE_BIT = 0x00000004
	LINE_RESET_TOKEN = 0x0707
	LINE_STIPPLE = 0x0B24
	LINE_STIPPLE_PATTERN = 0x0B25
	LINE_STIPPLE_REPEAT = 0x0B26
	LINE_TOKEN = 0x0702
	LIST_BASE = 0x0B32
	LIST_BIT = 0x00020000
	LIST_INDEX = 0x0B33
	LIST_MODE = 0x0B30
	LOAD = 0x0101
	LOGIC_OP = 0x0BF1
	LUMINANCE = 0x1909
	LUMINANCE12 = 0x8041
	LUMINANCE12_ALPHA12 = 0x8047
	LUMINANCE12_ALPHA4 = 0x8046
	LUMINANCE16 = 0x8042
	LUMINANCE16_ALPHA16 = 0x8048
	LUMINANCE4 = 0x803F
	LUMINANCE4_ALPHA4 = 0x8043
	LUMINANCE6_ALPHA2 = 0x8044
	LUMINANCE8 = 0x8040
	LUMINANCE8_ALPHA8 = 0x8045
	LUMINANCE_ALPHA = 0x190A
	MAP1_COLOR_4 = 0x0D90
	MAP1_GRID_DOMAIN = 0x0DD0
	MAP1_GRID_SEGMENTS = 0x0DD1
	MAP1_INDEX = 0x0D91
	MAP1_NORMAL = 0x0D92
	MAP1_TEXTURE_COORD_1 = 0x0D93
	MAP1_TEXTURE_COORD_2 = 0x0D94
	MAP1_TEXTURE_COORD_3 = 0x0D95
	MAP1_TEXTURE_COORD_4 = 0x0D96
	MAP1_VERTEX_3 = 0x0D97
	MAP1_VERTEX_4 = 0x0D98
	MAP2_COLOR_4 = 0x0DB0
	MAP2_GRID_DOMAIN = 0x0DD2
	MAP2_GRID_SEGMENTS = 0x0DD3
	MAP2_INDEX = 0x0DB1
	MAP2_NORMAL = 0x0DB2
	MAP2_TEXTURE_COORD_1 = 0x0DB3
	MAP2_TEXTURE_COORD_2 = 0x0DB4
	MAP2_TEXTURE_COORD_3 = 0x0DB5
	MAP2_TEXTURE_COORD_4 = 0x0DB6
	MAP2_VERTEX_3 = 0x0DB7
	MAP2_VERTEX_4 = 0x0DB8
	MAP_COLOR = 0x0D10
	MAP_STENCIL = 0x0D11
	MATRIX_MODE = 0x0BA0
	MAX_ATTRIB_STACK_DEPTH = 0x0D35
	MAX_CLIENT_ATTRIB_STACK_DEPTH = 0x0D3B
	MAX_CLIP_PLANES = 0x0D32
	MAX_EVAL_ORDER = 0x0D30
	MAX_LIGHTS = 0x0D31
	MAX_LIST_NESTING = 0x0B31
	MAX_MODELVIEW_STACK_DEPTH = 0x0D36
	MAX_NAME_STACK_DEPTH = 0x0D37
	MAX_PIXEL_MAP_TABLE = 0x0D34
	MAX_PROJECTION_STACK_DEPTH = 0x0D38
	MAX_TEXTURE_STACK_DEPTH = 0x0D39
	MODELVIEW = 0x1700
	MODELVIEW_MATRIX = 0x0BA6
	MODELVIEW_STACK_DEPTH = 0x0BA3
	MODULATE = 0x2100
	MULT = 0x0103
	N3F_V3F = 0x2A25
	NAME_STACK_DEPTH = 0x0D70
	NORMALIZE = 0x0BA1
	NORMAL_ARRAY = 0x8075
	NORMAL_ARRAY_POINTER = 0x808F
	NORMAL_ARRAY_STRIDE = 0x807F
	NORMAL_ARRAY_TYPE = 0x807E
	OBJECT_LINEAR = 0x2401
	OBJECT_PLANE = 0x2501
	ORDER = 0x0A01
	PASS_THROUGH_TOKEN = 0x0700
	PERSPECTIVE_CORRECTION_HINT = 0x0C50
	PIXEL_MAP_A_TO_A = 0x0C79
	PIXEL_MAP_A_TO_A_SIZE = 0x0CB9
	PIXEL_MAP_B_TO_B = 0x0C78
	PIXEL_MAP_B_TO_B_SIZE = 0x0CB8
	PIXEL_MAP_G_TO_G = 0x0C77
	PIXEL_MAP_G_TO_G_SIZE = 0x0CB7
	PIXEL_MAP_I_TO_A = 0x0C75
	PIXEL_MAP_I_TO_A_SIZE = 0x0CB5
	PIXEL_MAP_I_TO_B = 0x0C74
	PIXEL_MAP_I_TO_B_SIZE = 0x0CB4
	PIXEL_MAP_I_TO_G = 0x0C73
	PIXEL_MAP_I_TO_G_SIZE = 0x0CB3
	PIXEL_MAP_I_TO_I = 0x0C70
	PIXEL_MAP_I_TO_I_SIZE = 0x0CB0
	PIXEL_MAP_I_TO_R = 0x0C72
	PIXEL_MAP_I_TO_R_SIZE = 0x0CB2
	PIXEL_MAP_R_TO_R = 0x0C76
	PIXEL_MAP_R_TO_R_SIZE = 0x0CB6
	PIXEL_MAP_S_TO_S = 0x0C71
	PIXEL_MAP_S_TO_S_SIZE = 0x0CB1
	PIXEL_MODE_BIT = 0x00000020
	POINT_BIT = 0x00000002
	POINT_SMOOTH = 0x0B10
	POINT_SMOOTH_HINT = 0x0C51
	POINT_TOKEN = 0x0701
	POLYGON = 0x0009
	POLYGON_BIT = 0x00000008
	POLYGON_MODE = 0x0B40
	POLYGON_STIPPLE = 0x0B42
	POLYGON_STIPPLE_BIT = 0x00000010
	POLYGON_TOKEN = 0x0703
	POSITION = 0x1203
	PROJECTION = 0x1701
	PROJECTION_MATRIX = 0x0BA7
	PROJECTION_STACK_DEPTH = 0x0BA4
	Q = 0x2003
	QUADRATIC_ATTENUATION = 0x1209
	QUADS = 0x0007
	QUAD_STRIP = 0x0008
	R = 0x2002
	RED_BIAS = 0x0D15
	RED_BITS = 0x0D52
	RED_SCALE = 0x0D14
	RENDER = 0x1C00
	RENDER_MODE = 0x0C40
	RETURN = 0x0102
	RGBA_MODE = 0x0C31
	S = 0x2000
	SCISSOR_BIT = 0x00080000
	SELECT = 0x1C02
	SELECTION_BUFFER_POINTER = 0x0DF3
	SELECTION_BUFFER_SIZE = 0x0DF4
	SHADE_MODEL = 0x0B54
	SHININESS = 0x1601
	SMOOTH = 0x1D01
	SPECULAR = 0x1202
	SPHERE_MAP = 0x2402
	SPOT_CUTOFF = 0x1206
	SPOT_DIRECTION = 0x1204
	SPOT_EXPONENT = 0x1205
	STACK_OVERFLOW = 0x0503
	STACK_UNDERFLOW = 0x0504
	STENCIL_BITS = 0x0D57
	T = 0x2001
	T2F_C3F_V3F = 0x2A2A
	T2F_C4F_N3F_V3F = 0x2A2C
	T2F_C4UB_V3F = 0x2A29
	T2F_N3F_V3F = 0x2A2B
	T2F_V3F = 0x2A27
	T4F_C4F_N3F_V4F = 0x2A2D
	T4F_V4F = 0x2A28
	TEXTURE_BIT = 0x00040000
	TEXTURE_BORDER = 0x1005
	TEXTURE_COMPONENTS = 0x1003
	TEXTURE_COORD_ARRAY = 0x8078
	TEXTURE_COORD_ARRAY_POINTER = 0x8092
	TEXTURE_COORD_ARRAY_SIZE = 0x8088
	TEXTURE_COORD_ARRAY_STRIDE = 0x808A
	TEXTURE_COORD_ARRAY_TYPE = 0x8089
	TEXTURE_ENV = 0x2300
	TEXTURE_ENV_COLOR = 0x2201
	TEXTURE_ENV_MODE = 0x2200
	TEXTURE_GEN_MODE = 0x2500
	TEXTURE_GEN_Q = 0x0C63
	TEXTURE_GEN_R = 0x0C62
	TEXTURE_GEN_S = 0x0C60
	TEXTURE_GEN_T = 0x0C61
	TEXTURE_INTENSITY_SIZE = 0x8061
	TEXTURE_LUMINANCE_SIZE = 0x8060
	TEXTURE_MATRIX = 0x0BA8
	TEXTURE_PRIORITY = 0x8066
	TEXTURE_RESIDENT = 0x8067
	TEXTURE_STACK_DEPTH = 0x0BA5
	TRANSFORM_BIT = 0x00001000
	V2F = 0x2A20
	V3F = 0x2A21
	VERTEX_ARRAY = 0x8074
	VERTEX_ARRAY_POINTER = 0x808E
	VERTEX_ARRAY_SIZE = 0x807A
	VERTEX_ARRAY_STRIDE = 0x807C
	VERTEX_ARRAY_TYPE = 0x807B
	VIEWPORT_BIT = 0x00000800
	ZOOM_X = 0x0D16
	ZOOM_Y = 0x0D17
)
// VERSION_1_2
const (
	ALIASED_LINE_WIDTH_RANGE = 0x846E
	BGR = 0x80E0
	BGRA = 0x80E1
	CLAMP_TO_EDGE = 0x812F
	MAX_3D_TEXTURE_SIZE = 0x8073
	MAX_ELEMENTS_INDICES = 0x80E9
	MAX_ELEMENTS_VERTICES = 0x80E8
	PACK_IMAGE_HEIGHT = 0x806C
	PACK_SKIP_IMAGES = 0x806B
	PROXY_TEXTURE_3D = 0x8070
	SMOOTH_LINE_WIDTH_GRANULARITY = 0x0B23
	SMOOTH_LINE_WIDTH_RANGE = 0x0B22
	SMOOTH_POINT_SIZE_GRANULARITY = 0x0B13
	SMOOTH_POINT_SIZE_RANGE = 0x0B12
	TEXTURE_3D = 0x806F
	TEXTURE_BASE_LEVEL = 0x813C
	TEXTURE_BINDING_3D = 0x806A
	TEXTURE_DEPTH = 0x8071
	TEXTURE_MAX_LEVEL = 0x813D
	TEXTURE_MAX_LOD = 0x813B
	TEXTURE_MIN_LOD = 0x813A
	TEXTURE_WRAP_R = 0x8072
	UNPACK_IMAGE_HEIGHT = 0x806E
	UNPACK_SKIP_IMAGES = 0x806D
	UNSIGNED_BYTE_2_3_3_REV = 0x8362
	UNSIGNED_BYTE_3_3_2 = 0x8032
	UNSIGNED_INT_10_10_10_2 = 0x8036
	UNSIGNED_INT_2_10_10_10_REV = 0x8368
	UNSIGNED_INT_8_8_8_8 = 0x8035
	UNSIGNED_INT_8_8_8_8_REV = 0x8367
	UNSIGNED_SHORT_1_5_5_5_REV = 0x8366
	UNSIGNED_SHORT_4_4_4_4 = 0x8033
	UNSIGNED_SHORT_4_4_4_4_REV = 0x8365
	UNSIGNED_SHORT_5_5_5_1 = 0x8034
	UNSIGNED_SHORT_5_6_5 = 0x8363
	UNSIGNED_SHORT_5_6_5_REV = 0x8364
)
// VERSION_1_2_DEPRECATED
const (
	ALIASED_POINT_SIZE_RANGE = 0x846D
	LIGHT_MODEL_COLOR_CONTROL = 0x81F8
	RESCALE_NORMAL = 0x803A
	SEPARATE_SPECULAR_COLOR = 0x81FA
	SINGLE_COLOR = 0x81F9
)
// VERSION_1_3
const (
	ACTIVE_TEXTURE = 0x84E0
	CLAMP_TO_BORDER = 0x812D
	COMPRESSED_RGB = 0x84ED
	COMPRESSED_RGBA = 0x84EE
	COMPRESSED_TEXTURE_FORMATS = 0x86A3
	MAX_CUBE_MAP_TEXTURE_SIZE = 0x851C
	MULTISAMPLE = 0x809D
	NUM_COMPRESSED_TEXTURE_FORMATS = 0x86A2
	PROXY_TEXTURE_CUBE_MAP = 0x851B
	SAMPLES = 0x80A9
	SAMPLE_ALPHA_TO_COVERAGE = 0x809E
	SAMPLE_ALPHA_TO_ONE = 0x809F
	SAMPLE_BUFFERS = 0x80A8
	SAMPLE_COVERAGE = 0x80A0
	SAMPLE_COVERAGE_INVERT = 0x80AB
	SAMPLE_COVERAGE_VALUE = 0x80AA
	TEXTURE0 = 0x84C0
	TEXTURE1 = 0x84C1
	TEXTURE10 = 0x84CA
	TEXTURE11 = 0x84CB
	TEXTURE12 = 0x84CC
	TEXTURE13 = 0x84CD
	TEXTURE14 = 0x84CE
	TEXTURE15 = 0x84CF
	TEXTURE16 = 0x84D0
	TEXTURE17 = 0x84D1
	TEXTURE18 = 0x84D2
	TEXTURE19 = 0x84D3
	TEXTURE2 = 0x84C2
	TEXTURE20 = 0x84D4
	TEXTURE21 = 0x84D5
	TEXTURE22 = 0x84D6
	TEXTURE23 = 0x84D7
	TEXTURE24 = 0x84D8
	TEXTURE25 = 0x84D9
	TEXTURE26 = 0x84DA
	TEXTURE27 = 0x84DB
	TEXTURE28 = 0x84DC
	TEXTURE29 = 0x84DD
	TEXTURE3 = 0x84C3
	TEXTURE30 = 0x84DE
	TEXTURE31 = 0x84DF
	TEXTURE4 = 0x84C4
	TEXTURE5 = 0x84C5
	TEXTURE6 = 0x84C6
	TEXTURE7 = 0x84C7
	TEXTURE8 = 0x84C8
	TEXTURE9 = 0x84C9
	TEXTURE_BINDING_CUBE_MAP = 0x8514
	TEXTURE_COMPRESSED = 0x86A1
	TEXTURE_COMPRESSED_IMAGE_SIZE = 0x86A0
	TEXTURE_COMPRESSION_HINT = 0x84EF
	TEXTURE_CUBE_MAP = 0x8513
	TEXTURE_CUBE_MAP_NEGATIVE_X = 0x8516
	TEXTURE_CUBE_MAP_NEGATIVE_Y = 0x8518
	TEXTURE_CUBE_MAP_NEGATIVE_Z = 0x851A
	TEXTURE_CUBE_MAP_POSITIVE_X = 0x8515
	TEXTURE_CUBE_MAP_POSITIVE_Y = 0x8517
	TEXTURE_CUBE_MAP_POSITIVE_Z = 0x8519
)
// VERSION_1_3_DEPRECATED
const (
	ADD_SIGNED = 0x8574
	CLIENT_ACTIVE_TEXTURE = 0x84E1
	COMBINE = 0x8570
	COMBINE_ALPHA = 0x8572
	COMBINE_RGB = 0x8571
	COMPRESSED_ALPHA = 0x84E9
	COMPRESSED_INTENSITY = 0x84EC
	COMPRESSED_LUMINANCE = 0x84EA
	COMPRESSED_LUMINANCE_ALPHA = 0x84EB
	CONSTANT = 0x8576
	DOT3_RGB = 0x86AE
	DOT3_RGBA = 0x86AF
	INTERPOLATE = 0x8575
	MAX_TEXTURE_UNITS = 0x84E2
	MULTISAMPLE_BIT = 0x20000000
	NORMAL_MAP = 0x8511
	OPERAND0_ALPHA = 0x8598
	OPERAND0_RGB = 0x8590
	OPERAND1_ALPHA = 0x8599
	OPERAND1_RGB = 0x8591
	OPERAND2_ALPHA = 0x859A
	OPERAND2_RGB = 0x8592
	PREVIOUS = 0x8578
	PRIMARY_COLOR = 0x8577
	REFLECTION_MAP = 0x8512
	RGB_SCALE = 0x8573
	SOURCE0_ALPHA = 0x8588
	SOURCE0_RGB = 0x8580
	SOURCE1_ALPHA = 0x8589
	SOURCE1_RGB = 0x8581
	SOURCE2_ALPHA = 0x858A
	SOURCE2_RGB = 0x8582
	SUBTRACT = 0x84E7
	TRANSPOSE_COLOR_MATRIX = 0x84E6
	TRANSPOSE_MODELVIEW_MATRIX = 0x84E3
	TRANSPOSE_PROJECTION_MATRIX = 0x84E4
	TRANSPOSE_TEXTURE_MATRIX = 0x84E5
)
// VERSION_1_4
const (
	BLEND_DST_ALPHA = 0x80CA
	BLEND_DST_RGB = 0x80C8
	BLEND_SRC_ALPHA = 0x80CB
	BLEND_SRC_RGB = 0x80C9
	DECR_WRAP = 0x8508
	DEPTH_COMPONENT16 = 0x81A5
	DEPTH_COMPONENT24 = 0x81A6
	DEPTH_COMPONENT32 = 0x81A7
	INCR_WRAP = 0x8507
	MAX_TEXTURE_LOD_BIAS = 0x84FD
	MIRRORED_REPEAT = 0x8370
	POINT_FADE_THRESHOLD_SIZE = 0x8128
	TEXTURE_COMPARE_FUNC = 0x884D
	TEXTURE_COMPARE_MODE = 0x884C
	TEXTURE_DEPTH_SIZE = 0x884A
	TEXTURE_LOD_BIAS = 0x8501
)
// VERSION_1_4_DEPRECATED
const (
	COLOR_SUM = 0x8458
	COMPARE_R_TO_TEXTURE = 0x884E
	CURRENT_FOG_COORDINATE = 0x8453
	CURRENT_SECONDARY_COLOR = 0x8459
	DEPTH_TEXTURE_MODE = 0x884B
	FOG_COORDINATE = 0x8451
	FOG_COORDINATE_ARRAY = 0x8457
	FOG_COORDINATE_ARRAY_POINTER = 0x8456
	FOG_COORDINATE_ARRAY_STRIDE = 0x8455
	FOG_COORDINATE_ARRAY_TYPE = 0x8454
	FOG_COORDINATE_SOURCE = 0x8450
	FRAGMENT_DEPTH = 0x8452
	GENERATE_MIPMAP = 0x8191
	GENERATE_MIPMAP_HINT = 0x8192
	POINT_DISTANCE_ATTENUATION = 0x8129
	POINT_SIZE_MAX = 0x8127
	POINT_SIZE_MIN = 0x8126
	SECONDARY_COLOR_ARRAY = 0x845E
	SECONDARY_COLOR_ARRAY_POINTER = 0x845D
	SECONDARY_COLOR_ARRAY_SIZE = 0x845A
	SECONDARY_COLOR_ARRAY_STRIDE = 0x845C
	SECONDARY_COLOR_ARRAY_TYPE = 0x845B
	TEXTURE_FILTER_CONTROL = 0x8500
)
// VERSION_1_5
const (
	ARRAY_BUFFER = 0x8892
	ARRAY_BUFFER_BINDING = 0x8894
	BUFFER_ACCESS = 0x88BB
	BUFFER_MAPPED = 0x88BC
	BUFFER_MAP_POINTER = 0x88BD
	BUFFER_SIZE = 0x8764
	BUFFER_USAGE = 0x8765
	CURRENT_QUERY = 0x8865
	DYNAMIC_COPY = 0x88EA
	DYNAMIC_DRAW = 0x88E8
	DYNAMIC_READ = 0x88E9
	ELEMENT_ARRAY_BUFFER = 0x8893
	ELEMENT_ARRAY_BUFFER_BINDING = 0x8895
	QUERY_COUNTER_BITS = 0x8864
	QUERY_RESULT = 0x8866
	QUERY_RESULT_AVAILABLE = 0x8867
	READ_ONLY = 0x88B8
	READ_WRITE = 0x88BA
	SAMPLES_PASSED = 0x8914
	STATIC_COPY = 0x88E6
	STATIC_DRAW = 0x88E4
	STATIC_READ = 0x88E5
	STREAM_COPY = 0x88E2
	STREAM_DRAW = 0x88E0
	STREAM_READ = 0x88E1
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = 0x889F
	WRITE_ONLY = 0x88B9
)
// VERSION_1_5_DEPRECATED
const (
	COLOR_ARRAY_BUFFER_BINDING = 0x8898
	CURRENT_FOG_COORD = 0x8453
	EDGE_FLAG_ARRAY_BUFFER_BINDING = 0x889B
	FOG_COORD = 0x8451
	FOG_COORDINATE_ARRAY_BUFFER_BINDING = 0x889D
	FOG_COORD_ARRAY = 0x8457
	FOG_COORD_ARRAY_BUFFER_BINDING = 0x889D
	FOG_COORD_ARRAY_POINTER = 0x8456
	FOG_COORD_ARRAY_STRIDE = 0x8455
	FOG_COORD_ARRAY_TYPE = 0x8454
	FOG_COORD_SRC = 0x8450
	INDEX_ARRAY_BUFFER_BINDING = 0x8899
	NORMAL_ARRAY_BUFFER_BINDING = 0x8897
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING = 0x889C
	SRC0_ALPHA = 0x8588
	SRC0_RGB = 0x8580
	SRC1_ALPHA = 0x8589
	SRC1_RGB = 0x8581
	SRC2_ALPHA = 0x858A
	SRC2_RGB = 0x8582
	TEXTURE_COORD_ARRAY_BUFFER_BINDING = 0x889A
	VERTEX_ARRAY_BUFFER_BINDING = 0x8896
	WEIGHT_ARRAY_BUFFER_BINDING = 0x889E
)
// VERSION_2_0
const (
	ACTIVE_ATTRIBUTES = 0x8B89
	ACTIVE_ATTRIBUTE_MAX_LENGTH = 0x8B8A
	ACTIVE_UNIFORMS = 0x8B86
	ACTIVE_UNIFORM_MAX_LENGTH = 0x8B87
	ATTACHED_SHADERS = 0x8B85
	BLEND_EQUATION_ALPHA = 0x883D
	BLEND_EQUATION_RGB = 0x8009
	BOOL = 0x8B56
	BOOL_VEC2 = 0x8B57
	BOOL_VEC3 = 0x8B58
	BOOL_VEC4 = 0x8B59
	COMPILE_STATUS = 0x8B81
	CURRENT_PROGRAM = 0x8B8D
	CURRENT_VERTEX_ATTRIB = 0x8626
	DELETE_STATUS = 0x8B80
	DRAW_BUFFER0 = 0x8825
	DRAW_BUFFER1 = 0x8826
	DRAW_BUFFER10 = 0x882F
	DRAW_BUFFER11 = 0x8830
	DRAW_BUFFER12 = 0x8831
	DRAW_BUFFER13 = 0x8832
	DRAW_BUFFER14 = 0x8833
	DRAW_BUFFER15 = 0x8834
	DRAW_BUFFER2 = 0x8827
	DRAW_BUFFER3 = 0x8828
	DRAW_BUFFER4 = 0x8829
	DRAW_BUFFER5 = 0x882A
	DRAW_BUFFER6 = 0x882B
	DRAW_BUFFER7 = 0x882C
	DRAW_BUFFER8 = 0x882D
	DRAW_BUFFER9 = 0x882E
	FLOAT_MAT2 = 0x8B5A
	FLOAT_MAT3 = 0x8B5B
	FLOAT_MAT4 = 0x8B5C
	FLOAT_VEC2 = 0x8B50
	FLOAT_VEC3 = 0x8B51
	FLOAT_VEC4 = 0x8B52
	FRAGMENT_SHADER = 0x8B30
	FRAGMENT_SHADER_DERIVATIVE_HINT = 0x8B8B
	INFO_LOG_LENGTH = 0x8B84
	INT_VEC2 = 0x8B53
	INT_VEC3 = 0x8B54
	INT_VEC4 = 0x8B55
	LINK_STATUS = 0x8B82
	LOWER_LEFT = 0x8CA1
	MAX_COMBINED_TEXTURE_IMAGE_UNITS = 0x8B4D
	MAX_DRAW_BUFFERS = 0x8824
	MAX_FRAGMENT_UNIFORM_COMPONENTS = 0x8B49
	MAX_TEXTURE_IMAGE_UNITS = 0x8872
	MAX_VARYING_FLOATS = 0x8B4B
	MAX_VERTEX_ATTRIBS = 0x8869
	MAX_VERTEX_TEXTURE_IMAGE_UNITS = 0x8B4C
	MAX_VERTEX_UNIFORM_COMPONENTS = 0x8B4A
	POINT_SPRITE_COORD_ORIGIN = 0x8CA0
	SAMPLER_1D = 0x8B5D
	SAMPLER_1D_SHADOW = 0x8B61
	SAMPLER_2D = 0x8B5E
	SAMPLER_2D_SHADOW = 0x8B62
	SAMPLER_3D = 0x8B5F
	SAMPLER_CUBE = 0x8B60
	SHADER_SOURCE_LENGTH = 0x8B88
	SHADER_TYPE = 0x8B4F
	SHADING_LANGUAGE_VERSION = 0x8B8C
	STENCIL_BACK_FAIL = 0x8801
	STENCIL_BACK_FUNC = 0x8800
	STENCIL_BACK_PASS_DEPTH_FAIL = 0x8802
	STENCIL_BACK_PASS_DEPTH_PASS = 0x8803
	STENCIL_BACK_REF = 0x8CA3
	STENCIL_BACK_VALUE_MASK = 0x8CA4
	STENCIL_BACK_WRITEMASK = 0x8CA5
	UPPER_LEFT = 0x8CA2
	VALIDATE_STATUS = 0x8B83
	VERTEX_ATTRIB_ARRAY_ENABLED = 0x8622
	VERTEX_ATTRIB_ARRAY_NORMALIZED = 0x886A
	VERTEX_ATTRIB_ARRAY_POINTER = 0x8645
	VERTEX_ATTRIB_ARRAY_SIZE = 0x8623
	VERTEX_ATTRIB_ARRAY_STRIDE = 0x8624
	VERTEX_ATTRIB_ARRAY_TYPE = 0x8625
	VERTEX_PROGRAM_POINT_SIZE = 0x8642
	VERTEX_SHADER = 0x8B31
)
// VERSION_2_0_DEPRECATED
const (
	COORD_REPLACE = 0x8862
	MAX_TEXTURE_COORDS = 0x8871
	POINT_SPRITE = 0x8861
	VERTEX_PROGRAM_TWO_SIDE = 0x8643
)
// VERSION_2_1
const (
	COMPRESSED_SRGB = 0x8C48
	COMPRESSED_SRGB_ALPHA = 0x8C49
	FLOAT_MAT2x3 = 0x8B65
	FLOAT_MAT2x4 = 0x8B66
	FLOAT_MAT3x2 = 0x8B67
	FLOAT_MAT3x4 = 0x8B68
	FLOAT_MAT4x2 = 0x8B69
	FLOAT_MAT4x3 = 0x8B6A
	PIXEL_PACK_BUFFER = 0x88EB
	PIXEL_PACK_BUFFER_BINDING = 0x88ED
	PIXEL_UNPACK_BUFFER = 0x88EC
	PIXEL_UNPACK_BUFFER_BINDING = 0x88EF
	SRGB = 0x8C40
	SRGB8 = 0x8C41
	SRGB8_ALPHA8 = 0x8C43
	SRGB_ALPHA = 0x8C42
)
// VERSION_2_1_DEPRECATED
const (
	COMPRESSED_SLUMINANCE = 0x8C4A
	COMPRESSED_SLUMINANCE_ALPHA = 0x8C4B
	CURRENT_RASTER_SECONDARY_COLOR = 0x845F
	SLUMINANCE = 0x8C46
	SLUMINANCE8 = 0x8C47
	SLUMINANCE8_ALPHA8 = 0x8C45
	SLUMINANCE_ALPHA = 0x8C44
)
// VERSION_1_0

// http://www.opengl.org/sdk/docs/man/xhtml/glCullFace.xml
func CullFace(mode Enum)  {
	C.goglCullFace((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFrontFace.xml
func FrontFace(mode Enum)  {
	C.goglFrontFace((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glHint.xml
func Hint(target Enum, mode Enum)  {
	C.goglHint((C.GLenum)(target), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLineWidth.xml
func LineWidth(width Float)  {
	C.goglLineWidth((C.GLfloat)(width))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPointSize.xml
func PointSize(size Float)  {
	C.goglPointSize((C.GLfloat)(size))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPolygonMode.xml
func PolygonMode(face Enum, mode Enum)  {
	C.goglPolygonMode((C.GLenum)(face), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glScissor.xml
func Scissor(x Int, y Int, width Sizei, height Sizei)  {
	C.goglScissor((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterf.xml
func TexParameterf(target Enum, pname Enum, param Float)  {
	C.goglTexParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterfv.xml
func TexParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexParameteri.xml
func TexParameteri(target Enum, pname Enum, param Int)  {
	C.goglTexParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexParameteriv.xml
func TexParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexImage1D.xml
func TexImage1D(target Enum, level Int, internalformat Int, width Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexImage2D.xml
func TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDrawBuffer.xml
func DrawBuffer(mode Enum)  {
	C.goglDrawBuffer((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClear.xml
func Clear(mask Bitfield)  {
	C.goglClear((C.GLbitfield)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClearColor.xml
func ClearColor(red Clampf, green Clampf, blue Clampf, alpha Clampf)  {
	C.goglClearColor((C.GLclampf)(red), (C.GLclampf)(green), (C.GLclampf)(blue), (C.GLclampf)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClearStencil.xml
func ClearStencil(s Int)  {
	C.goglClearStencil((C.GLint)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClearDepth.xml
func ClearDepth(depth Clampd)  {
	C.goglClearDepth((C.GLclampd)(depth))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glStencilMask.xml
func StencilMask(mask Uint)  {
	C.goglStencilMask((C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColorMask.xml
func ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean)  {
	C.goglColorMask((C.GLboolean)(red), (C.GLboolean)(green), (C.GLboolean)(blue), (C.GLboolean)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDepthMask.xml
func DepthMask(flag Boolean)  {
	C.goglDepthMask((C.GLboolean)(flag))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDisable.xml
func Disable(cap Enum)  {
	C.goglDisable((C.GLenum)(cap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEnable.xml
func Enable(cap Enum)  {
	C.goglEnable((C.GLenum)(cap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFinish.xml
func Finish()  {
	C.goglFinish()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFlush.xml
func Flush()  {
	C.goglFlush()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBlendFunc.xml
func BlendFunc(sfactor Enum, dfactor Enum)  {
	C.goglBlendFunc((C.GLenum)(sfactor), (C.GLenum)(dfactor))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLogicOp.xml
func LogicOp(opcode Enum)  {
	C.goglLogicOp((C.GLenum)(opcode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glStencilFunc.xml
func StencilFunc(func_ Enum, ref Int, mask Uint)  {
	C.goglStencilFunc((C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glStencilOp.xml
func StencilOp(fail Enum, zfail Enum, zpass Enum)  {
	C.goglStencilOp((C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDepthFunc.xml
func DepthFunc(func_ Enum)  {
	C.goglDepthFunc((C.GLenum)(func_))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelStoref.xml
func PixelStoref(pname Enum, param Float)  {
	C.goglPixelStoref((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelStorei.xml
func PixelStorei(pname Enum, param Int)  {
	C.goglPixelStorei((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glReadBuffer.xml
func ReadBuffer(mode Enum)  {
	C.goglReadBuffer((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glReadPixels.xml
func ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglReadPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetBooleanv.xml
func GetBooleanv(pname Enum, params *Boolean)  {
	C.goglGetBooleanv((C.GLenum)(pname), (*C.GLboolean)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetDoublev.xml
func GetDoublev(pname Enum, params *Double)  {
	C.goglGetDoublev((C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetError.xml
func GetError() Enum {
	return (Enum)(C.goglGetError())
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetFloatv.xml
func GetFloatv(pname Enum, params *Float)  {
	C.goglGetFloatv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetIntegerv.xml
func GetIntegerv(pname Enum, params *Int)  {
	C.goglGetIntegerv((C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetString.xml
func GetString(name Enum) *Ubyte {
	return (*Ubyte)(C.goglGetString((C.GLenum)(name)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexImage.xml
func GetTexImage(target Enum, level Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglGetTexImage((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterfv.xml
func GetTexParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameteriv.xml
func GetTexParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexLevelParameterfv.xml
func GetTexLevelParameterfv(target Enum, level Int, pname Enum, params *Float)  {
	C.goglGetTexLevelParameterfv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexLevelParameteriv.xml
func GetTexLevelParameteriv(target Enum, level Int, pname Enum, params *Int)  {
	C.goglGetTexLevelParameteriv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIsEnabled.xml
func IsEnabled(cap Enum) Boolean {
	return (Boolean)(C.goglIsEnabled((C.GLenum)(cap)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDepthRange.xml
func DepthRange(near_ Clampd, far_ Clampd)  {
	C.goglDepthRange((C.GLclampd)(near_), (C.GLclampd)(far_))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glViewport.xml
func Viewport(x Int, y Int, width Sizei, height Sizei)  {
	C.goglViewport((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_0_DEPRECATED

// http://www.opengl.org/sdk/docs/man/xhtml/glNewList.xml
func NewList(list Uint, mode Enum)  {
	C.goglNewList((C.GLuint)(list), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEndList.xml
func EndList()  {
	C.goglEndList()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCallList.xml
func CallList(list Uint)  {
	C.goglCallList((C.GLuint)(list))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCallLists.xml
func CallLists(n Sizei, type_ Enum, lists Pointer)  {
	C.goglCallLists((C.GLsizei)(n), (C.GLenum)(type_), (unsafe.Pointer)(lists))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDeleteLists.xml
func DeleteLists(list Uint, range_ Sizei)  {
	C.goglDeleteLists((C.GLuint)(list), (C.GLsizei)(range_))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGenLists.xml
func GenLists(range_ Sizei) Uint {
	return (Uint)(C.goglGenLists((C.GLsizei)(range_)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glListBase.xml
func ListBase(base Uint)  {
	C.goglListBase((C.GLuint)(base))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBegin.xml
func Begin(mode Enum)  {
	C.goglBegin((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBitmap.xml
func Bitmap(width Sizei, height Sizei, xorig Float, yorig Float, xmove Float, ymove Float, bitmap *Ubyte)  {
	C.goglBitmap((C.GLsizei)(width), (C.GLsizei)(height), (C.GLfloat)(xorig), (C.GLfloat)(yorig), (C.GLfloat)(xmove), (C.GLfloat)(ymove), (*C.GLubyte)(bitmap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3b.xml
func Color3b(red Byte, green Byte, blue Byte)  {
	C.goglColor3b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3bv.xml
func Color3bv(v *Byte)  {
	C.goglColor3bv((*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3d.xml
func Color3d(red Double, green Double, blue Double)  {
	C.goglColor3d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3dv.xml
func Color3dv(v *Double)  {
	C.goglColor3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3f.xml
func Color3f(red Float, green Float, blue Float)  {
	C.goglColor3f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3fv.xml
func Color3fv(v *Float)  {
	C.goglColor3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3i.xml
func Color3i(red Int, green Int, blue Int)  {
	C.goglColor3i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3iv.xml
func Color3iv(v *Int)  {
	C.goglColor3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3s.xml
func Color3s(red Short, green Short, blue Short)  {
	C.goglColor3s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3sv.xml
func Color3sv(v *Short)  {
	C.goglColor3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3ub.xml
func Color3ub(red Ubyte, green Ubyte, blue Ubyte)  {
	C.goglColor3ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3ubv.xml
func Color3ubv(v *Ubyte)  {
	C.goglColor3ubv((*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3ui.xml
func Color3ui(red Uint, green Uint, blue Uint)  {
	C.goglColor3ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3uiv.xml
func Color3uiv(v *Uint)  {
	C.goglColor3uiv((*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3us.xml
func Color3us(red Ushort, green Ushort, blue Ushort)  {
	C.goglColor3us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3usv.xml
func Color3usv(v *Ushort)  {
	C.goglColor3usv((*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4b.xml
func Color4b(red Byte, green Byte, blue Byte, alpha Byte)  {
	C.goglColor4b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue), (C.GLbyte)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4bv.xml
func Color4bv(v *Byte)  {
	C.goglColor4bv((*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4d.xml
func Color4d(red Double, green Double, blue Double, alpha Double)  {
	C.goglColor4d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue), (C.GLdouble)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4dv.xml
func Color4dv(v *Double)  {
	C.goglColor4dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4f.xml
func Color4f(red Float, green Float, blue Float, alpha Float)  {
	C.goglColor4f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4fv.xml
func Color4fv(v *Float)  {
	C.goglColor4fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4i.xml
func Color4i(red Int, green Int, blue Int, alpha Int)  {
	C.goglColor4i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue), (C.GLint)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4iv.xml
func Color4iv(v *Int)  {
	C.goglColor4iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4s.xml
func Color4s(red Short, green Short, blue Short, alpha Short)  {
	C.goglColor4s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue), (C.GLshort)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4sv.xml
func Color4sv(v *Short)  {
	C.goglColor4sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4ub.xml
func Color4ub(red Ubyte, green Ubyte, blue Ubyte, alpha Ubyte)  {
	C.goglColor4ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue), (C.GLubyte)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4ubv.xml
func Color4ubv(v *Ubyte)  {
	C.goglColor4ubv((*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4ui.xml
func Color4ui(red Uint, green Uint, blue Uint, alpha Uint)  {
	C.goglColor4ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue), (C.GLuint)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4uiv.xml
func Color4uiv(v *Uint)  {
	C.goglColor4uiv((*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4us.xml
func Color4us(red Ushort, green Ushort, blue Ushort, alpha Ushort)  {
	C.goglColor4us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue), (C.GLushort)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4usv.xml
func Color4usv(v *Ushort)  {
	C.goglColor4usv((*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEdgeFlag.xml
func EdgeFlag(flag Boolean)  {
	C.goglEdgeFlag((C.GLboolean)(flag))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEdgeFlagv.xml
func EdgeFlagv(flag *Boolean)  {
	C.goglEdgeFlagv((*C.GLboolean)(flag))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEnd.xml
func End()  {
	C.goglEnd()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexd.xml
func Indexd(c Double)  {
	C.goglIndexd((C.GLdouble)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexdv.xml
func Indexdv(c *Double)  {
	C.goglIndexdv((*C.GLdouble)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexf.xml
func Indexf(c Float)  {
	C.goglIndexf((C.GLfloat)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexfv.xml
func Indexfv(c *Float)  {
	C.goglIndexfv((*C.GLfloat)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexi.xml
func Indexi(c Int)  {
	C.goglIndexi((C.GLint)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexiv.xml
func Indexiv(c *Int)  {
	C.goglIndexiv((*C.GLint)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexs.xml
func Indexs(c Short)  {
	C.goglIndexs((C.GLshort)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexsv.xml
func Indexsv(c *Short)  {
	C.goglIndexsv((*C.GLshort)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3b.xml
func Normal3b(nx Byte, ny Byte, nz Byte)  {
	C.goglNormal3b((C.GLbyte)(nx), (C.GLbyte)(ny), (C.GLbyte)(nz))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3bv.xml
func Normal3bv(v *Byte)  {
	C.goglNormal3bv((*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3d.xml
func Normal3d(nx Double, ny Double, nz Double)  {
	C.goglNormal3d((C.GLdouble)(nx), (C.GLdouble)(ny), (C.GLdouble)(nz))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3dv.xml
func Normal3dv(v *Double)  {
	C.goglNormal3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3f.xml
func Normal3f(nx Float, ny Float, nz Float)  {
	C.goglNormal3f((C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3fv.xml
func Normal3fv(v *Float)  {
	C.goglNormal3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3i.xml
func Normal3i(nx Int, ny Int, nz Int)  {
	C.goglNormal3i((C.GLint)(nx), (C.GLint)(ny), (C.GLint)(nz))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3iv.xml
func Normal3iv(v *Int)  {
	C.goglNormal3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3s.xml
func Normal3s(nx Short, ny Short, nz Short)  {
	C.goglNormal3s((C.GLshort)(nx), (C.GLshort)(ny), (C.GLshort)(nz))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3sv.xml
func Normal3sv(v *Short)  {
	C.goglNormal3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2d.xml
func RasterPos2d(x Double, y Double)  {
	C.goglRasterPos2d((C.GLdouble)(x), (C.GLdouble)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2dv.xml
func RasterPos2dv(v *Double)  {
	C.goglRasterPos2dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2f.xml
func RasterPos2f(x Float, y Float)  {
	C.goglRasterPos2f((C.GLfloat)(x), (C.GLfloat)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2fv.xml
func RasterPos2fv(v *Float)  {
	C.goglRasterPos2fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2i.xml
func RasterPos2i(x Int, y Int)  {
	C.goglRasterPos2i((C.GLint)(x), (C.GLint)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2iv.xml
func RasterPos2iv(v *Int)  {
	C.goglRasterPos2iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2s.xml
func RasterPos2s(x Short, y Short)  {
	C.goglRasterPos2s((C.GLshort)(x), (C.GLshort)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2sv.xml
func RasterPos2sv(v *Short)  {
	C.goglRasterPos2sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3d.xml
func RasterPos3d(x Double, y Double, z Double)  {
	C.goglRasterPos3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3dv.xml
func RasterPos3dv(v *Double)  {
	C.goglRasterPos3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3f.xml
func RasterPos3f(x Float, y Float, z Float)  {
	C.goglRasterPos3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3fv.xml
func RasterPos3fv(v *Float)  {
	C.goglRasterPos3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3i.xml
func RasterPos3i(x Int, y Int, z Int)  {
	C.goglRasterPos3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3iv.xml
func RasterPos3iv(v *Int)  {
	C.goglRasterPos3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3s.xml
func RasterPos3s(x Short, y Short, z Short)  {
	C.goglRasterPos3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3sv.xml
func RasterPos3sv(v *Short)  {
	C.goglRasterPos3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4d.xml
func RasterPos4d(x Double, y Double, z Double, w Double)  {
	C.goglRasterPos4d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4dv.xml
func RasterPos4dv(v *Double)  {
	C.goglRasterPos4dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4f.xml
func RasterPos4f(x Float, y Float, z Float, w Float)  {
	C.goglRasterPos4f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4fv.xml
func RasterPos4fv(v *Float)  {
	C.goglRasterPos4fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4i.xml
func RasterPos4i(x Int, y Int, z Int, w Int)  {
	C.goglRasterPos4i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4iv.xml
func RasterPos4iv(v *Int)  {
	C.goglRasterPos4iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4s.xml
func RasterPos4s(x Short, y Short, z Short, w Short)  {
	C.goglRasterPos4s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4sv.xml
func RasterPos4sv(v *Short)  {
	C.goglRasterPos4sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectd.xml
func Rectd(x1 Double, y1 Double, x2 Double, y2 Double)  {
	C.goglRectd((C.GLdouble)(x1), (C.GLdouble)(y1), (C.GLdouble)(x2), (C.GLdouble)(y2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectdv.xml
func Rectdv(v1 *Double, v2 *Double)  {
	C.goglRectdv((*C.GLdouble)(v1), (*C.GLdouble)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectf.xml
func Rectf(x1 Float, y1 Float, x2 Float, y2 Float)  {
	C.goglRectf((C.GLfloat)(x1), (C.GLfloat)(y1), (C.GLfloat)(x2), (C.GLfloat)(y2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectfv.xml
func Rectfv(v1 *Float, v2 *Float)  {
	C.goglRectfv((*C.GLfloat)(v1), (*C.GLfloat)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRecti.xml
func Recti(x1 Int, y1 Int, x2 Int, y2 Int)  {
	C.goglRecti((C.GLint)(x1), (C.GLint)(y1), (C.GLint)(x2), (C.GLint)(y2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectiv.xml
func Rectiv(v1 *Int, v2 *Int)  {
	C.goglRectiv((*C.GLint)(v1), (*C.GLint)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRects.xml
func Rects(x1 Short, y1 Short, x2 Short, y2 Short)  {
	C.goglRects((C.GLshort)(x1), (C.GLshort)(y1), (C.GLshort)(x2), (C.GLshort)(y2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectsv.xml
func Rectsv(v1 *Short, v2 *Short)  {
	C.goglRectsv((*C.GLshort)(v1), (*C.GLshort)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1d.xml
func TexCoord1d(s Double)  {
	C.goglTexCoord1d((C.GLdouble)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1dv.xml
func TexCoord1dv(v *Double)  {
	C.goglTexCoord1dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1f.xml
func TexCoord1f(s Float)  {
	C.goglTexCoord1f((C.GLfloat)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1fv.xml
func TexCoord1fv(v *Float)  {
	C.goglTexCoord1fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1i.xml
func TexCoord1i(s Int)  {
	C.goglTexCoord1i((C.GLint)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1iv.xml
func TexCoord1iv(v *Int)  {
	C.goglTexCoord1iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1s.xml
func TexCoord1s(s Short)  {
	C.goglTexCoord1s((C.GLshort)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1sv.xml
func TexCoord1sv(v *Short)  {
	C.goglTexCoord1sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2d.xml
func TexCoord2d(s Double, t Double)  {
	C.goglTexCoord2d((C.GLdouble)(s), (C.GLdouble)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2dv.xml
func TexCoord2dv(v *Double)  {
	C.goglTexCoord2dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2f.xml
func TexCoord2f(s Float, t Float)  {
	C.goglTexCoord2f((C.GLfloat)(s), (C.GLfloat)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2fv.xml
func TexCoord2fv(v *Float)  {
	C.goglTexCoord2fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2i.xml
func TexCoord2i(s Int, t Int)  {
	C.goglTexCoord2i((C.GLint)(s), (C.GLint)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2iv.xml
func TexCoord2iv(v *Int)  {
	C.goglTexCoord2iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2s.xml
func TexCoord2s(s Short, t Short)  {
	C.goglTexCoord2s((C.GLshort)(s), (C.GLshort)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2sv.xml
func TexCoord2sv(v *Short)  {
	C.goglTexCoord2sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3d.xml
func TexCoord3d(s Double, t Double, r Double)  {
	C.goglTexCoord3d((C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3dv.xml
func TexCoord3dv(v *Double)  {
	C.goglTexCoord3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3f.xml
func TexCoord3f(s Float, t Float, r Float)  {
	C.goglTexCoord3f((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3fv.xml
func TexCoord3fv(v *Float)  {
	C.goglTexCoord3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3i.xml
func TexCoord3i(s Int, t Int, r Int)  {
	C.goglTexCoord3i((C.GLint)(s), (C.GLint)(t), (C.GLint)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3iv.xml
func TexCoord3iv(v *Int)  {
	C.goglTexCoord3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3s.xml
func TexCoord3s(s Short, t Short, r Short)  {
	C.goglTexCoord3s((C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3sv.xml
func TexCoord3sv(v *Short)  {
	C.goglTexCoord3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4d.xml
func TexCoord4d(s Double, t Double, r Double, q Double)  {
	C.goglTexCoord4d((C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r), (C.GLdouble)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4dv.xml
func TexCoord4dv(v *Double)  {
	C.goglTexCoord4dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4f.xml
func TexCoord4f(s Float, t Float, r Float, q Float)  {
	C.goglTexCoord4f((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4fv.xml
func TexCoord4fv(v *Float)  {
	C.goglTexCoord4fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4i.xml
func TexCoord4i(s Int, t Int, r Int, q Int)  {
	C.goglTexCoord4i((C.GLint)(s), (C.GLint)(t), (C.GLint)(r), (C.GLint)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4iv.xml
func TexCoord4iv(v *Int)  {
	C.goglTexCoord4iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4s.xml
func TexCoord4s(s Short, t Short, r Short, q Short)  {
	C.goglTexCoord4s((C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r), (C.GLshort)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4sv.xml
func TexCoord4sv(v *Short)  {
	C.goglTexCoord4sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2d.xml
func Vertex2d(x Double, y Double)  {
	C.goglVertex2d((C.GLdouble)(x), (C.GLdouble)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2dv.xml
func Vertex2dv(v *Double)  {
	C.goglVertex2dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2f.xml
func Vertex2f(x Float, y Float)  {
	C.goglVertex2f((C.GLfloat)(x), (C.GLfloat)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2fv.xml
func Vertex2fv(v *Float)  {
	C.goglVertex2fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2i.xml
func Vertex2i(x Int, y Int)  {
	C.goglVertex2i((C.GLint)(x), (C.GLint)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2iv.xml
func Vertex2iv(v *Int)  {
	C.goglVertex2iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2s.xml
func Vertex2s(x Short, y Short)  {
	C.goglVertex2s((C.GLshort)(x), (C.GLshort)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2sv.xml
func Vertex2sv(v *Short)  {
	C.goglVertex2sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3d.xml
func Vertex3d(x Double, y Double, z Double)  {
	C.goglVertex3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3dv.xml
func Vertex3dv(v *Double)  {
	C.goglVertex3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3f.xml
func Vertex3f(x Float, y Float, z Float)  {
	C.goglVertex3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3fv.xml
func Vertex3fv(v *Float)  {
	C.goglVertex3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3i.xml
func Vertex3i(x Int, y Int, z Int)  {
	C.goglVertex3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3iv.xml
func Vertex3iv(v *Int)  {
	C.goglVertex3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3s.xml
func Vertex3s(x Short, y Short, z Short)  {
	C.goglVertex3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3sv.xml
func Vertex3sv(v *Short)  {
	C.goglVertex3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4d.xml
func Vertex4d(x Double, y Double, z Double, w Double)  {
	C.goglVertex4d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4dv.xml
func Vertex4dv(v *Double)  {
	C.goglVertex4dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4f.xml
func Vertex4f(x Float, y Float, z Float, w Float)  {
	C.goglVertex4f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4fv.xml
func Vertex4fv(v *Float)  {
	C.goglVertex4fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4i.xml
func Vertex4i(x Int, y Int, z Int, w Int)  {
	C.goglVertex4i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4iv.xml
func Vertex4iv(v *Int)  {
	C.goglVertex4iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4s.xml
func Vertex4s(x Short, y Short, z Short, w Short)  {
	C.goglVertex4s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4sv.xml
func Vertex4sv(v *Short)  {
	C.goglVertex4sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClipPlane.xml
func ClipPlane(plane Enum, equation *Double)  {
	C.goglClipPlane((C.GLenum)(plane), (*C.GLdouble)(equation))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColorMaterial.xml
func ColorMaterial(face Enum, mode Enum)  {
	C.goglColorMaterial((C.GLenum)(face), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogf.xml
func Fogf(pname Enum, param Float)  {
	C.goglFogf((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogfv.xml
func Fogfv(pname Enum, params *Float)  {
	C.goglFogfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogi.xml
func Fogi(pname Enum, param Int)  {
	C.goglFogi((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogiv.xml
func Fogiv(pname Enum, params *Int)  {
	C.goglFogiv((C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightf.xml
func Lightf(light Enum, pname Enum, param Float)  {
	C.goglLightf((C.GLenum)(light), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightfv.xml
func Lightfv(light Enum, pname Enum, params *Float)  {
	C.goglLightfv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLighti.xml
func Lighti(light Enum, pname Enum, param Int)  {
	C.goglLighti((C.GLenum)(light), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightiv.xml
func Lightiv(light Enum, pname Enum, params *Int)  {
	C.goglLightiv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightModelf.xml
func LightModelf(pname Enum, param Float)  {
	C.goglLightModelf((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightModelfv.xml
func LightModelfv(pname Enum, params *Float)  {
	C.goglLightModelfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightModeli.xml
func LightModeli(pname Enum, param Int)  {
	C.goglLightModeli((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightModeliv.xml
func LightModeliv(pname Enum, params *Int)  {
	C.goglLightModeliv((C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLineStipple.xml
func LineStipple(factor Int, pattern Ushort)  {
	C.goglLineStipple((C.GLint)(factor), (C.GLushort)(pattern))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMaterialf.xml
func Materialf(face Enum, pname Enum, param Float)  {
	C.goglMaterialf((C.GLenum)(face), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMaterialfv.xml
func Materialfv(face Enum, pname Enum, params *Float)  {
	C.goglMaterialfv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMateriali.xml
func Materiali(face Enum, pname Enum, param Int)  {
	C.goglMateriali((C.GLenum)(face), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMaterialiv.xml
func Materialiv(face Enum, pname Enum, params *Int)  {
	C.goglMaterialiv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPolygonStipple.xml
func PolygonStipple(mask *Ubyte)  {
	C.goglPolygonStipple((*C.GLubyte)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glShadeModel.xml
func ShadeModel(mode Enum)  {
	C.goglShadeModel((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexEnvf.xml
func TexEnvf(target Enum, pname Enum, param Float)  {
	C.goglTexEnvf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexEnvfv.xml
func TexEnvfv(target Enum, pname Enum, params *Float)  {
	C.goglTexEnvfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexEnvi.xml
func TexEnvi(target Enum, pname Enum, param Int)  {
	C.goglTexEnvi((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexEnviv.xml
func TexEnviv(target Enum, pname Enum, params *Int)  {
	C.goglTexEnviv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGend.xml
func TexGend(coord Enum, pname Enum, param Double)  {
	C.goglTexGend((C.GLenum)(coord), (C.GLenum)(pname), (C.GLdouble)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGendv.xml
func TexGendv(coord Enum, pname Enum, params *Double)  {
	C.goglTexGendv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGenf.xml
func TexGenf(coord Enum, pname Enum, param Float)  {
	C.goglTexGenf((C.GLenum)(coord), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGenfv.xml
func TexGenfv(coord Enum, pname Enum, params *Float)  {
	C.goglTexGenfv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGeni.xml
func TexGeni(coord Enum, pname Enum, param Int)  {
	C.goglTexGeni((C.GLenum)(coord), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGeniv.xml
func TexGeniv(coord Enum, pname Enum, params *Int)  {
	C.goglTexGeniv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFeedbackBuffer.xml
func FeedbackBuffer(size Sizei, type_ Enum, buffer *Float)  {
	C.goglFeedbackBuffer((C.GLsizei)(size), (C.GLenum)(type_), (*C.GLfloat)(buffer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSelectBuffer.xml
func SelectBuffer(size Sizei, buffer *Uint)  {
	C.goglSelectBuffer((C.GLsizei)(size), (*C.GLuint)(buffer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRenderMode.xml
func RenderMode(mode Enum) Int {
	return (Int)(C.goglRenderMode((C.GLenum)(mode)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glInitNames.xml
func InitNames()  {
	C.goglInitNames()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLoadName.xml
func LoadName(name Uint)  {
	C.goglLoadName((C.GLuint)(name))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPassThrough.xml
func PassThrough(token Float)  {
	C.goglPassThrough((C.GLfloat)(token))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPopName.xml
func PopName()  {
	C.goglPopName()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPushName.xml
func PushName(name Uint)  {
	C.goglPushName((C.GLuint)(name))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClearAccum.xml
func ClearAccum(red Float, green Float, blue Float, alpha Float)  {
	C.goglClearAccum((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClearIndex.xml
func ClearIndex(c Float)  {
	C.goglClearIndex((C.GLfloat)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexMask.xml
func IndexMask(mask Uint)  {
	C.goglIndexMask((C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glAccum.xml
func Accum(op Enum, value Float)  {
	C.goglAccum((C.GLenum)(op), (C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPopAttrib.xml
func PopAttrib()  {
	C.goglPopAttrib()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPushAttrib.xml
func PushAttrib(mask Bitfield)  {
	C.goglPushAttrib((C.GLbitfield)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMap1d.xml
func Map1d(target Enum, u1 Double, u2 Double, stride Int, order Int, points *Double)  {
	C.goglMap1d((C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLdouble)(points))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMap1f.xml
func Map1f(target Enum, u1 Float, u2 Float, stride Int, order Int, points *Float)  {
	C.goglMap1f((C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLfloat)(points))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMap2d.xml
func Map2d(target Enum, u1 Double, u2 Double, ustride Int, uorder Int, v1 Double, v2 Double, vstride Int, vorder Int, points *Double)  {
	C.goglMap2d((C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLdouble)(points))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMap2f.xml
func Map2f(target Enum, u1 Float, u2 Float, ustride Int, uorder Int, v1 Float, v2 Float, vstride Int, vorder Int, points *Float)  {
	C.goglMap2f((C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLfloat)(points))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMapGrid1d.xml
func MapGrid1d(un Int, u1 Double, u2 Double)  {
	C.goglMapGrid1d((C.GLint)(un), (C.GLdouble)(u1), (C.GLdouble)(u2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMapGrid1f.xml
func MapGrid1f(un Int, u1 Float, u2 Float)  {
	C.goglMapGrid1f((C.GLint)(un), (C.GLfloat)(u1), (C.GLfloat)(u2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMapGrid2d.xml
func MapGrid2d(un Int, u1 Double, u2 Double, vn Int, v1 Double, v2 Double)  {
	C.goglMapGrid2d((C.GLint)(un), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(vn), (C.GLdouble)(v1), (C.GLdouble)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMapGrid2f.xml
func MapGrid2f(un Int, u1 Float, u2 Float, vn Int, v1 Float, v2 Float)  {
	C.goglMapGrid2f((C.GLint)(un), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(vn), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord1d.xml
func EvalCoord1d(u Double)  {
	C.goglEvalCoord1d((C.GLdouble)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord1dv.xml
func EvalCoord1dv(u *Double)  {
	C.goglEvalCoord1dv((*C.GLdouble)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord1f.xml
func EvalCoord1f(u Float)  {
	C.goglEvalCoord1f((C.GLfloat)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord1fv.xml
func EvalCoord1fv(u *Float)  {
	C.goglEvalCoord1fv((*C.GLfloat)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord2d.xml
func EvalCoord2d(u Double, v Double)  {
	C.goglEvalCoord2d((C.GLdouble)(u), (C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord2dv.xml
func EvalCoord2dv(u *Double)  {
	C.goglEvalCoord2dv((*C.GLdouble)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord2f.xml
func EvalCoord2f(u Float, v Float)  {
	C.goglEvalCoord2f((C.GLfloat)(u), (C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord2fv.xml
func EvalCoord2fv(u *Float)  {
	C.goglEvalCoord2fv((*C.GLfloat)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalMesh1.xml
func EvalMesh1(mode Enum, i1 Int, i2 Int)  {
	C.goglEvalMesh1((C.GLenum)(mode), (C.GLint)(i1), (C.GLint)(i2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalPoint1.xml
func EvalPoint1(i Int)  {
	C.goglEvalPoint1((C.GLint)(i))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalMesh2.xml
func EvalMesh2(mode Enum, i1 Int, i2 Int, j1 Int, j2 Int)  {
	C.goglEvalMesh2((C.GLenum)(mode), (C.GLint)(i1), (C.GLint)(i2), (C.GLint)(j1), (C.GLint)(j2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalPoint2.xml
func EvalPoint2(i Int, j Int)  {
	C.goglEvalPoint2((C.GLint)(i), (C.GLint)(j))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glAlphaFunc.xml
func AlphaFunc(func_ Enum, ref Clampf)  {
	C.goglAlphaFunc((C.GLenum)(func_), (C.GLclampf)(ref))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelZoom.xml
func PixelZoom(xfactor Float, yfactor Float)  {
	C.goglPixelZoom((C.GLfloat)(xfactor), (C.GLfloat)(yfactor))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelTransferf.xml
func PixelTransferf(pname Enum, param Float)  {
	C.goglPixelTransferf((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelTransferi.xml
func PixelTransferi(pname Enum, param Int)  {
	C.goglPixelTransferi((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelMapfv.xml
func PixelMapfv(map_ Enum, mapsize Sizei, values *Float)  {
	C.goglPixelMapfv((C.GLenum)(map_), (C.GLsizei)(mapsize), (*C.GLfloat)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelMapuiv.xml
func PixelMapuiv(map_ Enum, mapsize Sizei, values *Uint)  {
	C.goglPixelMapuiv((C.GLenum)(map_), (C.GLsizei)(mapsize), (*C.GLuint)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelMapusv.xml
func PixelMapusv(map_ Enum, mapsize Sizei, values *Ushort)  {
	C.goglPixelMapusv((C.GLenum)(map_), (C.GLsizei)(mapsize), (*C.GLushort)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyPixels.xml
func CopyPixels(x Int, y Int, width Sizei, height Sizei, type_ Enum)  {
	C.goglCopyPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(type_))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDrawPixels.xml
func DrawPixels(width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglDrawPixels((C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetClipPlane.xml
func GetClipPlane(plane Enum, equation *Double)  {
	C.goglGetClipPlane((C.GLenum)(plane), (*C.GLdouble)(equation))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetLightfv.xml
func GetLightfv(light Enum, pname Enum, params *Float)  {
	C.goglGetLightfv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetLightiv.xml
func GetLightiv(light Enum, pname Enum, params *Int)  {
	C.goglGetLightiv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMapdv.xml
func GetMapdv(target Enum, query Enum, v *Double)  {
	C.goglGetMapdv((C.GLenum)(target), (C.GLenum)(query), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMapfv.xml
func GetMapfv(target Enum, query Enum, v *Float)  {
	C.goglGetMapfv((C.GLenum)(target), (C.GLenum)(query), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMapiv.xml
func GetMapiv(target Enum, query Enum, v *Int)  {
	C.goglGetMapiv((C.GLenum)(target), (C.GLenum)(query), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMaterialfv.xml
func GetMaterialfv(face Enum, pname Enum, params *Float)  {
	C.goglGetMaterialfv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMaterialiv.xml
func GetMaterialiv(face Enum, pname Enum, params *Int)  {
	C.goglGetMaterialiv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetPixelMapfv.xml
func GetPixelMapfv(map_ Enum, values *Float)  {
	C.goglGetPixelMapfv((C.GLenum)(map_), (*C.GLfloat)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetPixelMapuiv.xml
func GetPixelMapuiv(map_ Enum, values *Uint)  {
	C.goglGetPixelMapuiv((C.GLenum)(map_), (*C.GLuint)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetPixelMapusv.xml
func GetPixelMapusv(map_ Enum, values *Ushort)  {
	C.goglGetPixelMapusv((C.GLenum)(map_), (*C.GLushort)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetPolygonStipple.xml
func GetPolygonStipple(mask *Ubyte)  {
	C.goglGetPolygonStipple((*C.GLubyte)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexEnvfv.xml
func GetTexEnvfv(target Enum, pname Enum, params *Float)  {
	C.goglGetTexEnvfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexEnviv.xml
func GetTexEnviv(target Enum, pname Enum, params *Int)  {
	C.goglGetTexEnviv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexGendv.xml
func GetTexGendv(coord Enum, pname Enum, params *Double)  {
	C.goglGetTexGendv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexGenfv.xml
func GetTexGenfv(coord Enum, pname Enum, params *Float)  {
	C.goglGetTexGenfv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexGeniv.xml
func GetTexGeniv(coord Enum, pname Enum, params *Int)  {
	C.goglGetTexGeniv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIsList.xml
func IsList(list Uint) Boolean {
	return (Boolean)(C.goglIsList((C.GLuint)(list)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFrustum.xml
func Frustum(left Double, right Double, bottom Double, top Double, zNear Double, zFar Double)  {
	C.goglFrustum((C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLoadIdentity.xml
func LoadIdentity()  {
	C.goglLoadIdentity()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLoadMatrixf.xml
func LoadMatrixf(m *Float)  {
	C.goglLoadMatrixf((*C.GLfloat)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLoadMatrixd.xml
func LoadMatrixd(m *Double)  {
	C.goglLoadMatrixd((*C.GLdouble)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMatrixMode.xml
func MatrixMode(mode Enum)  {
	C.goglMatrixMode((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultMatrixf.xml
func MultMatrixf(m *Float)  {
	C.goglMultMatrixf((*C.GLfloat)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultMatrixd.xml
func MultMatrixd(m *Double)  {
	C.goglMultMatrixd((*C.GLdouble)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glOrtho.xml
func Ortho(left Double, right Double, bottom Double, top Double, zNear Double, zFar Double)  {
	C.goglOrtho((C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPopMatrix.xml
func PopMatrix()  {
	C.goglPopMatrix()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPushMatrix.xml
func PushMatrix()  {
	C.goglPushMatrix()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRotated.xml
func Rotated(angle Double, x Double, y Double, z Double)  {
	C.goglRotated((C.GLdouble)(angle), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRotatef.xml
func Rotatef(angle Float, x Float, y Float, z Float)  {
	C.goglRotatef((C.GLfloat)(angle), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glScaled.xml
func Scaled(x Double, y Double, z Double)  {
	C.goglScaled((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glScalef.xml
func Scalef(x Float, y Float, z Float)  {
	C.goglScalef((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTranslated.xml
func Translated(x Double, y Double, z Double)  {
	C.goglTranslated((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTranslatef.xml
func Translatef(x Float, y Float, z Float)  {
	C.goglTranslatef((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// VERSION_1_1

// http://www.opengl.org/sdk/docs/man/xhtml/glDrawArrays.xml
func DrawArrays(mode Enum, first Int, count Sizei)  {
	C.goglDrawArrays((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDrawElements.xml
func DrawElements(mode Enum, count Sizei, type_ Enum, indices Pointer)  {
	C.goglDrawElements((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetPointerv.xml
func GetPointerv(pname Enum, params *Pointer)  {
	C.goglGetPointerv((C.GLenum)(pname), (*unsafe.Pointer)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPolygonOffset.xml
func PolygonOffset(factor Float, units Float)  {
	C.goglPolygonOffset((C.GLfloat)(factor), (C.GLfloat)(units))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexImage1D.xml
func CopyTexImage1D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, border Int)  {
	C.goglCopyTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexImage2D.xml
func CopyTexImage2D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, height Sizei, border Int)  {
	C.goglCopyTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage1D.xml
func CopyTexSubImage1D(target Enum, level Int, xoffset Int, x Int, y Int, width Sizei)  {
	C.goglCopyTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage2D.xml
func CopyTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, x Int, y Int, width Sizei, height Sizei)  {
	C.goglCopyTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage1D.xml
func TexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage2D.xml
func TexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBindTexture.xml
func BindTexture(target Enum, texture Uint)  {
	C.goglBindTexture((C.GLenum)(target), (C.GLuint)(texture))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDeleteTextures.xml
func DeleteTextures(n Sizei, textures *Uint)  {
	C.goglDeleteTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGenTextures.xml
func GenTextures(n Sizei, textures *Uint)  {
	C.goglGenTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIsTexture.xml
func IsTexture(texture Uint) Boolean {
	return (Boolean)(C.goglIsTexture((C.GLuint)(texture)))
}
// VERSION_1_1_DEPRECATED

// http://www.opengl.org/sdk/docs/man/xhtml/glArrayElement.xml
func ArrayElement(i Int)  {
	C.goglArrayElement((C.GLint)(i))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColorPointer.xml
func ColorPointer(size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglColorPointer((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDisableClientState.xml
func DisableClientState(array Enum)  {
	C.goglDisableClientState((C.GLenum)(array))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEdgeFlagPointer.xml
func EdgeFlagPointer(stride Sizei, pointer Pointer)  {
	C.goglEdgeFlagPointer((C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEnableClientState.xml
func EnableClientState(array Enum)  {
	C.goglEnableClientState((C.GLenum)(array))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexPointer.xml
func IndexPointer(type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglIndexPointer((C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glInterleavedArrays.xml
func InterleavedArrays(format Enum, stride Sizei, pointer Pointer)  {
	C.goglInterleavedArrays((C.GLenum)(format), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormalPointer.xml
func NormalPointer(type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglNormalPointer((C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordPointer.xml
func TexCoordPointer(size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglTexCoordPointer((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexPointer.xml
func VertexPointer(size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglVertexPointer((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glAreTexturesResident.xml
func AreTexturesResident(n Sizei, textures *Uint, residences *Boolean) Boolean {
	return (Boolean)(C.goglAreTexturesResident((C.GLsizei)(n), (*C.GLuint)(textures), (*C.GLboolean)(residences)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPrioritizeTextures.xml
func PrioritizeTextures(n Sizei, textures *Uint, priorities *Clampf)  {
	C.goglPrioritizeTextures((C.GLsizei)(n), (*C.GLuint)(textures), (*C.GLclampf)(priorities))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexub.xml
func Indexub(c Ubyte)  {
	C.goglIndexub((C.GLubyte)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexubv.xml
func Indexubv(c *Ubyte)  {
	C.goglIndexubv((*C.GLubyte)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPopClientAttrib.xml
func PopClientAttrib()  {
	C.goglPopClientAttrib()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPushClientAttrib.xml
func PushClientAttrib(mask Bitfield)  {
	C.goglPushClientAttrib((C.GLbitfield)(mask))
}
// VERSION_1_2

// http://www.opengl.org/sdk/docs/man/xhtml/glBlendColor.xml
func BlendColor(red Clampf, green Clampf, blue Clampf, alpha Clampf)  {
	C.goglBlendColor((C.GLclampf)(red), (C.GLclampf)(green), (C.GLclampf)(blue), (C.GLclampf)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquation.xml
func BlendEquation(mode Enum)  {
	C.goglBlendEquation((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDrawRangeElements.xml
func DrawRangeElements(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Pointer)  {
	C.goglDrawRangeElements((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexImage3D.xml
func TexImage3D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, depth Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage3D.xml
func TexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage3D.xml
func CopyTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, x Int, y Int, width Sizei, height Sizei)  {
	C.goglCopyTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_2_DEPRECATED

// http://www.opengl.org/sdk/docs/man/xhtml/glColorTable.xml
func ColorTable(target Enum, internalformat Enum, width Sizei, format Enum, type_ Enum, table Pointer)  {
	C.goglColorTable((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(table))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColorTableParameterfv.xml
func ColorTableParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglColorTableParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColorTableParameteriv.xml
func ColorTableParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglColorTableParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyColorTable.xml
func CopyColorTable(target Enum, internalformat Enum, x Int, y Int, width Sizei)  {
	C.goglCopyColorTable((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetColorTable.xml
func GetColorTable(target Enum, format Enum, type_ Enum, table Pointer)  {
	C.goglGetColorTable((C.GLenum)(target), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(table))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetColorTableParameterfv.xml
func GetColorTableParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetColorTableParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetColorTableParameteriv.xml
func GetColorTableParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetColorTableParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColorSubTable.xml
func ColorSubTable(target Enum, start Sizei, count Sizei, format Enum, type_ Enum, data Pointer)  {
	C.goglColorSubTable((C.GLenum)(target), (C.GLsizei)(start), (C.GLsizei)(count), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyColorSubTable.xml
func CopyColorSubTable(target Enum, start Sizei, x Int, y Int, width Sizei)  {
	C.goglCopyColorSubTable((C.GLenum)(target), (C.GLsizei)(start), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glConvolutionFilter1D.xml
func ConvolutionFilter1D(target Enum, internalformat Enum, width Sizei, format Enum, type_ Enum, image Pointer)  {
	C.goglConvolutionFilter1D((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(image))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glConvolutionFilter2D.xml
func ConvolutionFilter2D(target Enum, internalformat Enum, width Sizei, height Sizei, format Enum, type_ Enum, image Pointer)  {
	C.goglConvolutionFilter2D((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(image))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glConvolutionParameterf.xml
func ConvolutionParameterf(target Enum, pname Enum, params Float)  {
	C.goglConvolutionParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glConvolutionParameterfv.xml
func ConvolutionParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglConvolutionParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glConvolutionParameteri.xml
func ConvolutionParameteri(target Enum, pname Enum, params Int)  {
	C.goglConvolutionParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glConvolutionParameteriv.xml
func ConvolutionParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglConvolutionParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyConvolutionFilter1D.xml
func CopyConvolutionFilter1D(target Enum, internalformat Enum, x Int, y Int, width Sizei)  {
	C.goglCopyConvolutionFilter1D((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyConvolutionFilter2D.xml
func CopyConvolutionFilter2D(target Enum, internalformat Enum, x Int, y Int, width Sizei, height Sizei)  {
	C.goglCopyConvolutionFilter2D((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetConvolutionFilter.xml
func GetConvolutionFilter(target Enum, format Enum, type_ Enum, image Pointer)  {
	C.goglGetConvolutionFilter((C.GLenum)(target), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(image))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetConvolutionParameterfv.xml
func GetConvolutionParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetConvolutionParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetConvolutionParameteriv.xml
func GetConvolutionParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetConvolutionParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetSeparableFilter.xml
func GetSeparableFilter(target Enum, format Enum, type_ Enum, row Pointer, column Pointer, span Pointer)  {
	C.goglGetSeparableFilter((C.GLenum)(target), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(row), (unsafe.Pointer)(column), (unsafe.Pointer)(span))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSeparableFilter2D.xml
func SeparableFilter2D(target Enum, internalformat Enum, width Sizei, height Sizei, format Enum, type_ Enum, row Pointer, column Pointer)  {
	C.goglSeparableFilter2D((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(row), (unsafe.Pointer)(column))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetHistogram.xml
func GetHistogram(target Enum, reset Boolean, format Enum, type_ Enum, values Pointer)  {
	C.goglGetHistogram((C.GLenum)(target), (C.GLboolean)(reset), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetHistogramParameterfv.xml
func GetHistogramParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetHistogramParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetHistogramParameteriv.xml
func GetHistogramParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetHistogramParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMinmax.xml
func GetMinmax(target Enum, reset Boolean, format Enum, type_ Enum, values Pointer)  {
	C.goglGetMinmax((C.GLenum)(target), (C.GLboolean)(reset), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMinmaxParameterfv.xml
func GetMinmaxParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetMinmaxParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMinmaxParameteriv.xml
func GetMinmaxParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetMinmaxParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glHistogram.xml
func Histogram(target Enum, width Sizei, internalformat Enum, sink Boolean)  {
	C.goglHistogram((C.GLenum)(target), (C.GLsizei)(width), (C.GLenum)(internalformat), (C.GLboolean)(sink))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMinmax.xml
func Minmax(target Enum, internalformat Enum, sink Boolean)  {
	C.goglMinmax((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLboolean)(sink))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glResetHistogram.xml
func ResetHistogram(target Enum)  {
	C.goglResetHistogram((C.GLenum)(target))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glResetMinmax.xml
func ResetMinmax(target Enum)  {
	C.goglResetMinmax((C.GLenum)(target))
}
// VERSION_1_3

// http://www.opengl.org/sdk/docs/man/xhtml/glActiveTexture.xml
func ActiveTexture(texture Enum)  {
	C.goglActiveTexture((C.GLenum)(texture))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSampleCoverage.xml
func SampleCoverage(value Clampf, invert Boolean)  {
	C.goglSampleCoverage((C.GLclampf)(value), (C.GLboolean)(invert))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage3D.xml
func CompressedTexImage3D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, depth Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage2D.xml
func CompressedTexImage2D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage1D.xml
func CompressedTexImage1D(target Enum, level Int, internalformat Enum, width Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage3D.xml
func CompressedTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage2D.xml
func CompressedTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage1D.xml
func CompressedTexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetCompressedTexImage.xml
func GetCompressedTexImage(target Enum, level Int, img Pointer)  {
	C.goglGetCompressedTexImage((C.GLenum)(target), (C.GLint)(level), (unsafe.Pointer)(img))
}
// VERSION_1_3_DEPRECATED

// http://www.opengl.org/sdk/docs/man/xhtml/glClientActiveTexture.xml
func ClientActiveTexture(texture Enum)  {
	C.goglClientActiveTexture((C.GLenum)(texture))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord1d.xml
func MultiTexCoord1d(target Enum, s Double)  {
	C.goglMultiTexCoord1d((C.GLenum)(target), (C.GLdouble)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord1dv.xml
func MultiTexCoord1dv(target Enum, v *Double)  {
	C.goglMultiTexCoord1dv((C.GLenum)(target), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord1f.xml
func MultiTexCoord1f(target Enum, s Float)  {
	C.goglMultiTexCoord1f((C.GLenum)(target), (C.GLfloat)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord1fv.xml
func MultiTexCoord1fv(target Enum, v *Float)  {
	C.goglMultiTexCoord1fv((C.GLenum)(target), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord1i.xml
func MultiTexCoord1i(target Enum, s Int)  {
	C.goglMultiTexCoord1i((C.GLenum)(target), (C.GLint)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord1iv.xml
func MultiTexCoord1iv(target Enum, v *Int)  {
	C.goglMultiTexCoord1iv((C.GLenum)(target), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord1s.xml
func MultiTexCoord1s(target Enum, s Short)  {
	C.goglMultiTexCoord1s((C.GLenum)(target), (C.GLshort)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord1sv.xml
func MultiTexCoord1sv(target Enum, v *Short)  {
	C.goglMultiTexCoord1sv((C.GLenum)(target), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord2d.xml
func MultiTexCoord2d(target Enum, s Double, t Double)  {
	C.goglMultiTexCoord2d((C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord2dv.xml
func MultiTexCoord2dv(target Enum, v *Double)  {
	C.goglMultiTexCoord2dv((C.GLenum)(target), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord2f.xml
func MultiTexCoord2f(target Enum, s Float, t Float)  {
	C.goglMultiTexCoord2f((C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord2fv.xml
func MultiTexCoord2fv(target Enum, v *Float)  {
	C.goglMultiTexCoord2fv((C.GLenum)(target), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord2i.xml
func MultiTexCoord2i(target Enum, s Int, t Int)  {
	C.goglMultiTexCoord2i((C.GLenum)(target), (C.GLint)(s), (C.GLint)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord2iv.xml
func MultiTexCoord2iv(target Enum, v *Int)  {
	C.goglMultiTexCoord2iv((C.GLenum)(target), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord2s.xml
func MultiTexCoord2s(target Enum, s Short, t Short)  {
	C.goglMultiTexCoord2s((C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord2sv.xml
func MultiTexCoord2sv(target Enum, v *Short)  {
	C.goglMultiTexCoord2sv((C.GLenum)(target), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord3d.xml
func MultiTexCoord3d(target Enum, s Double, t Double, r Double)  {
	C.goglMultiTexCoord3d((C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord3dv.xml
func MultiTexCoord3dv(target Enum, v *Double)  {
	C.goglMultiTexCoord3dv((C.GLenum)(target), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord3f.xml
func MultiTexCoord3f(target Enum, s Float, t Float, r Float)  {
	C.goglMultiTexCoord3f((C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord3fv.xml
func MultiTexCoord3fv(target Enum, v *Float)  {
	C.goglMultiTexCoord3fv((C.GLenum)(target), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord3i.xml
func MultiTexCoord3i(target Enum, s Int, t Int, r Int)  {
	C.goglMultiTexCoord3i((C.GLenum)(target), (C.GLint)(s), (C.GLint)(t), (C.GLint)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord3iv.xml
func MultiTexCoord3iv(target Enum, v *Int)  {
	C.goglMultiTexCoord3iv((C.GLenum)(target), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord3s.xml
func MultiTexCoord3s(target Enum, s Short, t Short, r Short)  {
	C.goglMultiTexCoord3s((C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord3sv.xml
func MultiTexCoord3sv(target Enum, v *Short)  {
	C.goglMultiTexCoord3sv((C.GLenum)(target), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord4d.xml
func MultiTexCoord4d(target Enum, s Double, t Double, r Double, q Double)  {
	C.goglMultiTexCoord4d((C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r), (C.GLdouble)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord4dv.xml
func MultiTexCoord4dv(target Enum, v *Double)  {
	C.goglMultiTexCoord4dv((C.GLenum)(target), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord4f.xml
func MultiTexCoord4f(target Enum, s Float, t Float, r Float, q Float)  {
	C.goglMultiTexCoord4f((C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord4fv.xml
func MultiTexCoord4fv(target Enum, v *Float)  {
	C.goglMultiTexCoord4fv((C.GLenum)(target), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord4i.xml
func MultiTexCoord4i(target Enum, s Int, t Int, r Int, q Int)  {
	C.goglMultiTexCoord4i((C.GLenum)(target), (C.GLint)(s), (C.GLint)(t), (C.GLint)(r), (C.GLint)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord4iv.xml
func MultiTexCoord4iv(target Enum, v *Int)  {
	C.goglMultiTexCoord4iv((C.GLenum)(target), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord4s.xml
func MultiTexCoord4s(target Enum, s Short, t Short, r Short, q Short)  {
	C.goglMultiTexCoord4s((C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r), (C.GLshort)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoord4sv.xml
func MultiTexCoord4sv(target Enum, v *Short)  {
	C.goglMultiTexCoord4sv((C.GLenum)(target), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLoadTransposeMatrixf.xml
func LoadTransposeMatrixf(m *Float)  {
	C.goglLoadTransposeMatrixf((*C.GLfloat)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLoadTransposeMatrixd.xml
func LoadTransposeMatrixd(m *Double)  {
	C.goglLoadTransposeMatrixd((*C.GLdouble)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultTransposeMatrixf.xml
func MultTransposeMatrixf(m *Float)  {
	C.goglMultTransposeMatrixf((*C.GLfloat)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultTransposeMatrixd.xml
func MultTransposeMatrixd(m *Double)  {
	C.goglMultTransposeMatrixd((*C.GLdouble)(m))
}
// VERSION_1_4

// http://www.opengl.org/sdk/docs/man/xhtml/glBlendFuncSeparate.xml
func BlendFuncSeparate(sfactorRGB Enum, dfactorRGB Enum, sfactorAlpha Enum, dfactorAlpha Enum)  {
	C.goglBlendFuncSeparate((C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawArrays.xml
func MultiDrawArrays(mode Enum, first *Int, count *Sizei, primcount Sizei)  {
	C.goglMultiDrawArrays((C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(primcount))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawElements.xml
func MultiDrawElements(mode Enum, count *Sizei, type_ Enum, indices *Pointer, primcount Sizei)  {
	C.goglMultiDrawElements((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(primcount))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPointParameterf.xml
func PointParameterf(pname Enum, param Float)  {
	C.goglPointParameterf((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPointParameterfv.xml
func PointParameterfv(pname Enum, params *Float)  {
	C.goglPointParameterfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPointParameteri.xml
func PointParameteri(pname Enum, param Int)  {
	C.goglPointParameteri((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPointParameteriv.xml
func PointParameteriv(pname Enum, params *Int)  {
	C.goglPointParameteriv((C.GLenum)(pname), (*C.GLint)(params))
}
// VERSION_1_4_DEPRECATED

// http://www.opengl.org/sdk/docs/man/xhtml/glFogCoordf.xml
func FogCoordf(coord Float)  {
	C.goglFogCoordf((C.GLfloat)(coord))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogCoordfv.xml
func FogCoordfv(coord *Float)  {
	C.goglFogCoordfv((*C.GLfloat)(coord))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogCoordd.xml
func FogCoordd(coord Double)  {
	C.goglFogCoordd((C.GLdouble)(coord))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogCoorddv.xml
func FogCoorddv(coord *Double)  {
	C.goglFogCoorddv((*C.GLdouble)(coord))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogCoordPointer.xml
func FogCoordPointer(type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglFogCoordPointer((C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3b.xml
func SecondaryColor3b(red Byte, green Byte, blue Byte)  {
	C.goglSecondaryColor3b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3bv.xml
func SecondaryColor3bv(v *Byte)  {
	C.goglSecondaryColor3bv((*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3d.xml
func SecondaryColor3d(red Double, green Double, blue Double)  {
	C.goglSecondaryColor3d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3dv.xml
func SecondaryColor3dv(v *Double)  {
	C.goglSecondaryColor3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3f.xml
func SecondaryColor3f(red Float, green Float, blue Float)  {
	C.goglSecondaryColor3f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3fv.xml
func SecondaryColor3fv(v *Float)  {
	C.goglSecondaryColor3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3i.xml
func SecondaryColor3i(red Int, green Int, blue Int)  {
	C.goglSecondaryColor3i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3iv.xml
func SecondaryColor3iv(v *Int)  {
	C.goglSecondaryColor3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3s.xml
func SecondaryColor3s(red Short, green Short, blue Short)  {
	C.goglSecondaryColor3s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3sv.xml
func SecondaryColor3sv(v *Short)  {
	C.goglSecondaryColor3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3ub.xml
func SecondaryColor3ub(red Ubyte, green Ubyte, blue Ubyte)  {
	C.goglSecondaryColor3ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3ubv.xml
func SecondaryColor3ubv(v *Ubyte)  {
	C.goglSecondaryColor3ubv((*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3ui.xml
func SecondaryColor3ui(red Uint, green Uint, blue Uint)  {
	C.goglSecondaryColor3ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3uiv.xml
func SecondaryColor3uiv(v *Uint)  {
	C.goglSecondaryColor3uiv((*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3us.xml
func SecondaryColor3us(red Ushort, green Ushort, blue Ushort)  {
	C.goglSecondaryColor3us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColor3usv.xml
func SecondaryColor3usv(v *Ushort)  {
	C.goglSecondaryColor3usv((*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColorPointer.xml
func SecondaryColorPointer(size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglSecondaryColorPointer((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos2d.xml
func WindowPos2d(x Double, y Double)  {
	C.goglWindowPos2d((C.GLdouble)(x), (C.GLdouble)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos2dv.xml
func WindowPos2dv(v *Double)  {
	C.goglWindowPos2dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos2f.xml
func WindowPos2f(x Float, y Float)  {
	C.goglWindowPos2f((C.GLfloat)(x), (C.GLfloat)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos2fv.xml
func WindowPos2fv(v *Float)  {
	C.goglWindowPos2fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos2i.xml
func WindowPos2i(x Int, y Int)  {
	C.goglWindowPos2i((C.GLint)(x), (C.GLint)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos2iv.xml
func WindowPos2iv(v *Int)  {
	C.goglWindowPos2iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos2s.xml
func WindowPos2s(x Short, y Short)  {
	C.goglWindowPos2s((C.GLshort)(x), (C.GLshort)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos2sv.xml
func WindowPos2sv(v *Short)  {
	C.goglWindowPos2sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos3d.xml
func WindowPos3d(x Double, y Double, z Double)  {
	C.goglWindowPos3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos3dv.xml
func WindowPos3dv(v *Double)  {
	C.goglWindowPos3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos3f.xml
func WindowPos3f(x Float, y Float, z Float)  {
	C.goglWindowPos3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos3fv.xml
func WindowPos3fv(v *Float)  {
	C.goglWindowPos3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos3i.xml
func WindowPos3i(x Int, y Int, z Int)  {
	C.goglWindowPos3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos3iv.xml
func WindowPos3iv(v *Int)  {
	C.goglWindowPos3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos3s.xml
func WindowPos3s(x Short, y Short, z Short)  {
	C.goglWindowPos3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glWindowPos3sv.xml
func WindowPos3sv(v *Short)  {
	C.goglWindowPos3sv((*C.GLshort)(v))
}
// VERSION_1_5

// http://www.opengl.org/sdk/docs/man/xhtml/glGenQueries.xml
func GenQueries(n Sizei, ids *Uint)  {
	C.goglGenQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDeleteQueries.xml
func DeleteQueries(n Sizei, ids *Uint)  {
	C.goglDeleteQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIsQuery.xml
func IsQuery(id Uint) Boolean {
	return (Boolean)(C.goglIsQuery((C.GLuint)(id)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBeginQuery.xml
func BeginQuery(target Enum, id Uint)  {
	C.goglBeginQuery((C.GLenum)(target), (C.GLuint)(id))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEndQuery.xml
func EndQuery(target Enum)  {
	C.goglEndQuery((C.GLenum)(target))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryiv.xml
func GetQueryiv(target Enum, pname Enum, params *Int)  {
	C.goglGetQueryiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectiv.xml
func GetQueryObjectiv(id Uint, pname Enum, params *Int)  {
	C.goglGetQueryObjectiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectuiv.xml
func GetQueryObjectuiv(id Uint, pname Enum, params *Uint)  {
	C.goglGetQueryObjectuiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBindBuffer.xml
func BindBuffer(target Enum, buffer Uint)  {
	C.goglBindBuffer((C.GLenum)(target), (C.GLuint)(buffer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDeleteBuffers.xml
func DeleteBuffers(n Sizei, buffers *Uint)  {
	C.goglDeleteBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGenBuffers.xml
func GenBuffers(n Sizei, buffers *Uint)  {
	C.goglGenBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIsBuffer.xml
func IsBuffer(buffer Uint) Boolean {
	return (Boolean)(C.goglIsBuffer((C.GLuint)(buffer)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBufferData.xml
func BufferData(target Enum, size Sizeiptr, data Pointer, usage Enum)  {
	C.goglBufferData((C.GLenum)(target), (C.GLsizeiptr)(size), (unsafe.Pointer)(data), (C.GLenum)(usage))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBufferSubData.xml
func BufferSubData(target Enum, offset Intptr, size Sizeiptr, data Pointer)  {
	C.goglBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferSubData.xml
func GetBufferSubData(target Enum, offset Intptr, size Sizeiptr, data Pointer)  {
	C.goglGetBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMapBuffer.xml
func MapBuffer(target Enum, access Enum) Pointer {
	return (Pointer)(C.goglMapBuffer((C.GLenum)(target), (C.GLenum)(access)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUnmapBuffer.xml
func UnmapBuffer(target Enum) Boolean {
	return (Boolean)(C.goglUnmapBuffer((C.GLenum)(target)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferParameteriv.xml
func GetBufferParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetBufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferPointerv.xml
func GetBufferPointerv(target Enum, pname Enum, params *Pointer)  {
	C.goglGetBufferPointerv((C.GLenum)(target), (C.GLenum)(pname), (*unsafe.Pointer)(params))
}
// VERSION_2_0

// http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquationSeparate.xml
func BlendEquationSeparate(modeRGB Enum, modeAlpha Enum)  {
	C.goglBlendEquationSeparate((C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDrawBuffers.xml
func DrawBuffers(n Sizei, bufs *Enum)  {
	C.goglDrawBuffers((C.GLsizei)(n), (*C.GLenum)(bufs))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glStencilOpSeparate.xml
func StencilOpSeparate(face Enum, sfail Enum, dpfail Enum, dppass Enum)  {
	C.goglStencilOpSeparate((C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glStencilFuncSeparate.xml
func StencilFuncSeparate(face Enum, func_ Enum, ref Int, mask Uint)  {
	C.goglStencilFuncSeparate((C.GLenum)(face), (C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glStencilMaskSeparate.xml
func StencilMaskSeparate(face Enum, mask Uint)  {
	C.goglStencilMaskSeparate((C.GLenum)(face), (C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glAttachShader.xml
func AttachShader(program Uint, shader Uint)  {
	C.goglAttachShader((C.GLuint)(program), (C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBindAttribLocation.xml
func BindAttribLocation(program Uint, index Uint, name *Char)  {
	C.goglBindAttribLocation((C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCompileShader.xml
func CompileShader(shader Uint)  {
	C.goglCompileShader((C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateProgram.xml
func CreateProgram() Uint {
	return (Uint)(C.goglCreateProgram())
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCreateShader.xml
func CreateShader(type_ Enum) Uint {
	return (Uint)(C.goglCreateShader((C.GLenum)(type_)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDeleteProgram.xml
func DeleteProgram(program Uint)  {
	C.goglDeleteProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDeleteShader.xml
func DeleteShader(shader Uint)  {
	C.goglDeleteShader((C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDetachShader.xml
func DetachShader(program Uint, shader Uint)  {
	C.goglDetachShader((C.GLuint)(program), (C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDisableVertexAttribArray.xml
func DisableVertexAttribArray(index Uint)  {
	C.goglDisableVertexAttribArray((C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEnableVertexAttribArray.xml
func EnableVertexAttribArray(index Uint)  {
	C.goglEnableVertexAttribArray((C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveAttrib.xml
func GetActiveAttrib(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)  {
	C.goglGetActiveAttrib((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniform.xml
func GetActiveUniform(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)  {
	C.goglGetActiveUniform((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetAttachedShaders.xml
func GetAttachedShaders(program Uint, maxCount Sizei, count *Sizei, obj *Uint)  {
	C.goglGetAttachedShaders((C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(count), (*C.GLuint)(obj))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetAttribLocation.xml
func GetAttribLocation(program Uint, name *Char) Int {
	return (Int)(C.goglGetAttribLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramiv.xml
func GetProgramiv(program Uint, pname Enum, params *Int)  {
	C.goglGetProgramiv((C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramInfoLog.xml
func GetProgramInfoLog(program Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.goglGetProgramInfoLog((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderiv.xml
func GetShaderiv(shader Uint, pname Enum, params *Int)  {
	C.goglGetShaderiv((C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderInfoLog.xml
func GetShaderInfoLog(shader Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.goglGetShaderInfoLog((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderSource.xml
func GetShaderSource(shader Uint, bufSize Sizei, length *Sizei, source *Char)  {
	C.goglGetShaderSource((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(source))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformLocation.xml
func GetUniformLocation(program Uint, name *Char) Int {
	return (Int)(C.goglGetUniformLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformfv.xml
func GetUniformfv(program Uint, location Int, params *Float)  {
	C.goglGetUniformfv((C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformiv.xml
func GetUniformiv(program Uint, location Int, params *Int)  {
	C.goglGetUniformiv((C.GLuint)(program), (C.GLint)(location), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribdv.xml
func GetVertexAttribdv(index Uint, pname Enum, params *Double)  {
	C.goglGetVertexAttribdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribfv.xml
func GetVertexAttribfv(index Uint, pname Enum, params *Float)  {
	C.goglGetVertexAttribfv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribiv.xml
func GetVertexAttribiv(index Uint, pname Enum, params *Int)  {
	C.goglGetVertexAttribiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribPointerv.xml
func GetVertexAttribPointerv(index Uint, pname Enum, pointer *Pointer)  {
	C.goglGetVertexAttribPointerv((C.GLuint)(index), (C.GLenum)(pname), (*unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIsProgram.xml
func IsProgram(program Uint) Boolean {
	return (Boolean)(C.goglIsProgram((C.GLuint)(program)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIsShader.xml
func IsShader(shader Uint) Boolean {
	return (Boolean)(C.goglIsShader((C.GLuint)(shader)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLinkProgram.xml
func LinkProgram(program Uint)  {
	C.goglLinkProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glShaderSource.xml
func ShaderSource(shader Uint, count Sizei, string_ **Char, length *Int)  {
	C.goglShaderSource((C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(string_)), (*C.GLint)(length))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUseProgram.xml
func UseProgram(program Uint)  {
	C.goglUseProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform1f.xml
func Uniform1f(location Int, v0 Float)  {
	C.goglUniform1f((C.GLint)(location), (C.GLfloat)(v0))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform2f.xml
func Uniform2f(location Int, v0 Float, v1 Float)  {
	C.goglUniform2f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform3f.xml
func Uniform3f(location Int, v0 Float, v1 Float, v2 Float)  {
	C.goglUniform3f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform4f.xml
func Uniform4f(location Int, v0 Float, v1 Float, v2 Float, v3 Float)  {
	C.goglUniform4f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform1i.xml
func Uniform1i(location Int, v0 Int)  {
	C.goglUniform1i((C.GLint)(location), (C.GLint)(v0))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform2i.xml
func Uniform2i(location Int, v0 Int, v1 Int)  {
	C.goglUniform2i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform3i.xml
func Uniform3i(location Int, v0 Int, v1 Int, v2 Int)  {
	C.goglUniform3i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform4i.xml
func Uniform4i(location Int, v0 Int, v1 Int, v2 Int, v3 Int)  {
	C.goglUniform4i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform1fv.xml
func Uniform1fv(location Int, count Sizei, value *Float)  {
	C.goglUniform1fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform2fv.xml
func Uniform2fv(location Int, count Sizei, value *Float)  {
	C.goglUniform2fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform3fv.xml
func Uniform3fv(location Int, count Sizei, value *Float)  {
	C.goglUniform3fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform4fv.xml
func Uniform4fv(location Int, count Sizei, value *Float)  {
	C.goglUniform4fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform1iv.xml
func Uniform1iv(location Int, count Sizei, value *Int)  {
	C.goglUniform1iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform2iv.xml
func Uniform2iv(location Int, count Sizei, value *Int)  {
	C.goglUniform2iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform3iv.xml
func Uniform3iv(location Int, count Sizei, value *Int)  {
	C.goglUniform3iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniform4iv.xml
func Uniform4iv(location Int, count Sizei, value *Int)  {
	C.goglUniform4iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2fv.xml
func UniformMatrix2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3fv.xml
func UniformMatrix3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4fv.xml
func UniformMatrix4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glValidateProgram.xml
func ValidateProgram(program Uint)  {
	C.goglValidateProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib1d.xml
func VertexAttrib1d(index Uint, x Double)  {
	C.goglVertexAttrib1d((C.GLuint)(index), (C.GLdouble)(x))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib1dv.xml
func VertexAttrib1dv(index Uint, v *Double)  {
	C.goglVertexAttrib1dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib1f.xml
func VertexAttrib1f(index Uint, x Float)  {
	C.goglVertexAttrib1f((C.GLuint)(index), (C.GLfloat)(x))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib1fv.xml
func VertexAttrib1fv(index Uint, v *Float)  {
	C.goglVertexAttrib1fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib1s.xml
func VertexAttrib1s(index Uint, x Short)  {
	C.goglVertexAttrib1s((C.GLuint)(index), (C.GLshort)(x))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib1sv.xml
func VertexAttrib1sv(index Uint, v *Short)  {
	C.goglVertexAttrib1sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib2d.xml
func VertexAttrib2d(index Uint, x Double, y Double)  {
	C.goglVertexAttrib2d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib2dv.xml
func VertexAttrib2dv(index Uint, v *Double)  {
	C.goglVertexAttrib2dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib2f.xml
func VertexAttrib2f(index Uint, x Float, y Float)  {
	C.goglVertexAttrib2f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib2fv.xml
func VertexAttrib2fv(index Uint, v *Float)  {
	C.goglVertexAttrib2fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib2s.xml
func VertexAttrib2s(index Uint, x Short, y Short)  {
	C.goglVertexAttrib2s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib2sv.xml
func VertexAttrib2sv(index Uint, v *Short)  {
	C.goglVertexAttrib2sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib3d.xml
func VertexAttrib3d(index Uint, x Double, y Double, z Double)  {
	C.goglVertexAttrib3d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib3dv.xml
func VertexAttrib3dv(index Uint, v *Double)  {
	C.goglVertexAttrib3dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib3f.xml
func VertexAttrib3f(index Uint, x Float, y Float, z Float)  {
	C.goglVertexAttrib3f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib3fv.xml
func VertexAttrib3fv(index Uint, v *Float)  {
	C.goglVertexAttrib3fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib3s.xml
func VertexAttrib3s(index Uint, x Short, y Short, z Short)  {
	C.goglVertexAttrib3s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib3sv.xml
func VertexAttrib3sv(index Uint, v *Short)  {
	C.goglVertexAttrib3sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4Nbv.xml
func VertexAttrib4Nbv(index Uint, v *Byte)  {
	C.goglVertexAttrib4Nbv((C.GLuint)(index), (*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4Niv.xml
func VertexAttrib4Niv(index Uint, v *Int)  {
	C.goglVertexAttrib4Niv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4Nsv.xml
func VertexAttrib4Nsv(index Uint, v *Short)  {
	C.goglVertexAttrib4Nsv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4Nub.xml
func VertexAttrib4Nub(index Uint, x Ubyte, y Ubyte, z Ubyte, w Ubyte)  {
	C.goglVertexAttrib4Nub((C.GLuint)(index), (C.GLubyte)(x), (C.GLubyte)(y), (C.GLubyte)(z), (C.GLubyte)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4Nubv.xml
func VertexAttrib4Nubv(index Uint, v *Ubyte)  {
	C.goglVertexAttrib4Nubv((C.GLuint)(index), (*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4Nuiv.xml
func VertexAttrib4Nuiv(index Uint, v *Uint)  {
	C.goglVertexAttrib4Nuiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4Nusv.xml
func VertexAttrib4Nusv(index Uint, v *Ushort)  {
	C.goglVertexAttrib4Nusv((C.GLuint)(index), (*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4bv.xml
func VertexAttrib4bv(index Uint, v *Byte)  {
	C.goglVertexAttrib4bv((C.GLuint)(index), (*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4d.xml
func VertexAttrib4d(index Uint, x Double, y Double, z Double, w Double)  {
	C.goglVertexAttrib4d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4dv.xml
func VertexAttrib4dv(index Uint, v *Double)  {
	C.goglVertexAttrib4dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4f.xml
func VertexAttrib4f(index Uint, x Float, y Float, z Float, w Float)  {
	C.goglVertexAttrib4f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4fv.xml
func VertexAttrib4fv(index Uint, v *Float)  {
	C.goglVertexAttrib4fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4iv.xml
func VertexAttrib4iv(index Uint, v *Int)  {
	C.goglVertexAttrib4iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4s.xml
func VertexAttrib4s(index Uint, x Short, y Short, z Short, w Short)  {
	C.goglVertexAttrib4s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4sv.xml
func VertexAttrib4sv(index Uint, v *Short)  {
	C.goglVertexAttrib4sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4ubv.xml
func VertexAttrib4ubv(index Uint, v *Ubyte)  {
	C.goglVertexAttrib4ubv((C.GLuint)(index), (*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4uiv.xml
func VertexAttrib4uiv(index Uint, v *Uint)  {
	C.goglVertexAttrib4uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttrib4usv.xml
func VertexAttrib4usv(index Uint, v *Ushort)  {
	C.goglVertexAttrib4usv((C.GLuint)(index), (*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribPointer.xml
func VertexAttribPointer(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, pointer Pointer)  {
	C.goglVertexAttribPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// VERSION_2_1

// http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x3fv.xml
func UniformMatrix2x3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x2fv.xml
func UniformMatrix3x2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x4fv.xml
func UniformMatrix2x4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x2fv.xml
func UniformMatrix4x2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x4fv.xml
func UniformMatrix3x4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x3fv.xml
func UniformMatrix4x3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
func InitVersion10() error {
	var ret C.int
	if ret = C.init_VERSION_1_0(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_0")
	}
	return nil
}
func InitVersion10Deprecated() error {
	var ret C.int
	if ret = C.init_VERSION_1_0_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_0_DEPRECATED")
	}
	return nil
}
func InitVersion11() error {
	var ret C.int
	if ret = C.init_VERSION_1_1(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_1")
	}
	return nil
}
func InitVersion11Deprecated() error {
	var ret C.int
	if ret = C.init_VERSION_1_1_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_1_DEPRECATED")
	}
	return nil
}
func InitVersion12() error {
	var ret C.int
	if ret = C.init_VERSION_1_2(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_2")
	}
	return nil
}
func InitVersion12Deprecated() error {
	var ret C.int
	if ret = C.init_VERSION_1_2_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_2_DEPRECATED")
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
func InitVersion13Deprecated() error {
	var ret C.int
	if ret = C.init_VERSION_1_3_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_3_DEPRECATED")
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
func InitVersion14Deprecated() error {
	var ret C.int
	if ret = C.init_VERSION_1_4_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_4_DEPRECATED")
	}
	return nil
}
func InitVersion15() error {
	var ret C.int
	if ret = C.init_VERSION_1_5(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_5")
	}
	return nil
}
func InitVersion20() error {
	var ret C.int
	if ret = C.init_VERSION_2_0(); ret != 0 {
		return errors.New("unable to initialize VERSION_2_0")
	}
	return nil
}
func InitVersion21() error {
	var ret C.int
	if ret = C.init_VERSION_2_1(); ret != 0 {
		return errors.New("unable to initialize VERSION_2_1")
	}
	return nil
}
func Init() error {
	var err error
	if err = InitVersion10(); err != nil {
		return err
	}
	if err = InitVersion10Deprecated(); err != nil {
		return err
	}
	if err = InitVersion11(); err != nil {
		return err
	}
	if err = InitVersion11Deprecated(); err != nil {
		return err
	}
	if err = InitVersion12(); err != nil {
		return err
	}
	if err = InitVersion12Deprecated(); err != nil {
		return err
	}
	if err = InitVersion13(); err != nil {
		return err
	}
	if err = InitVersion13Deprecated(); err != nil {
		return err
	}
	if err = InitVersion14(); err != nil {
		return err
	}
	if err = InitVersion14Deprecated(); err != nil {
		return err
	}
	if err = InitVersion15(); err != nil {
		return err
	}
	if err = InitVersion20(); err != nil {
		return err
	}
	if err = InitVersion21(); err != nil {
		return err
	}
	return nil
}
//Go bool to GL boolean.
func GLBool(b bool) Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

// GL boolean to Go bool.
func GoBool(b Boolean) bool {
	return b == TRUE
}

// Go string to GL string.
func GLString(str string) *Char {
	return (*Char)(C.CString(str))
}

// Allocates a GL string.
func GLStringAlloc(length Sizei) *Char {
	return (*Char)(C.malloc(C.size_t(length)))
}

// Frees GL string.
func GLStringFree(str *Char) {
	C.free(unsafe.Pointer(str))
}

// GL string (GLchar*) to Go string.
func GoString(str *Char) string {
	return C.GoString((*C.char)(str))
}

// GL string (GLubyte*) to Go string.
func GoStringUb(str *Ubyte) string {
	return C.GoString((*C.char)(unsafe.Pointer(str)))
}

// GL string (GLchar*) with length to Go string.
func GoStringN(str *Char, length Sizei) string {
	return C.GoStringN((*C.char)(str), C.int(length))
}

// Converts a list of Go strings to a slice of GL strings.
// Usefull for ShaderSource().
func GLStringArray(strs ...string) []*Char {
	strSlice := make([]*Char, len(strs))
	for i, s := range strs {
		strSlice[i] = (*Char)(C.CString(s))
	}
	return strSlice
}

// Free GL string slice allocated by GLStringArray().
func GLStringArrayFree(strs []*Char) {
	for _, s := range strs {
		C.free(unsafe.Pointer(s))
	}
}

// Add offset to a pointer. Usefull for VertexAttribPointer, TexCoordPointer, NormalPointer, ... 
func Offset(p Pointer, o uintptr) Pointer {
	return Pointer(uintptr(p) + o)
}

// EOF