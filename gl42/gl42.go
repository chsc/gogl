// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// VERSION_1_0
// 
// VERSION_1_1
// 
// VERSION_1_2
// 
// VERSION_1_3
// 
// VERSION_1_4
// 
// VERSION_1_5
// 
// VERSION_2_0
// 
// VERSION_2_1
// 
// VERSION_3_0
// 
// VERSION_3_1
// 
// VERSION_3_2
// 
// VERSION_3_3
// 
// VERSION_4_0
// 
// VERSION_4_1
// 
// VERSION_4_2
// 
// http://www.opengl.org/sdk/docs/man4
// 
package gl42

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
// void* goglGetProcAddress(const char* name) { 
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
// void (APIENTRYP ptrglClearColor)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglClearStencil)(GLint s);
// void (APIENTRYP ptrglClearDepth)(GLdouble depth);
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
// void (APIENTRYP ptrglDepthRange)(GLdouble near, GLdouble far);
// void (APIENTRYP ptrglViewport)(GLint x, GLint y, GLsizei width, GLsizei height);
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
// //  VERSION_1_2
// void (APIENTRYP ptrglBlendColor)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglBlendEquation)(GLenum mode);
// void (APIENTRYP ptrglDrawRangeElements)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices);
// void (APIENTRYP ptrglTexImage3D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglCopyTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height);
// //  VERSION_1_3
// void (APIENTRYP ptrglActiveTexture)(GLenum texture);
// void (APIENTRYP ptrglSampleCoverage)(GLfloat value, GLboolean invert);
// void (APIENTRYP ptrglCompressedTexImage3D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglGetCompressedTexImage)(GLenum target, GLint level, GLvoid* img);
// //  VERSION_1_4
// void (APIENTRYP ptrglBlendFuncSeparate)(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha);
// void (APIENTRYP ptrglMultiDrawArrays)(GLenum mode, GLint* first, GLsizei* count, GLsizei primcount);
// void (APIENTRYP ptrglMultiDrawElements)(GLenum mode, GLsizei* count, GLenum type, GLvoid* const* indices, GLsizei primcount);
// void (APIENTRYP ptrglPointParameterf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglPointParameterfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglPointParameteri)(GLenum pname, GLint param);
// void (APIENTRYP ptrglPointParameteriv)(GLenum pname, GLint* params);
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
// void (APIENTRYP ptrglShaderSource)(GLuint shader, GLsizei count, GLchar* const* string, GLint* length);
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
// //  VERSION_3_0
// void (APIENTRYP ptrglColorMaski)(GLuint index, GLboolean r, GLboolean g, GLboolean b, GLboolean a);
// void (APIENTRYP ptrglGetBooleani_v)(GLenum target, GLuint index, GLboolean* data);
// void (APIENTRYP ptrglGetIntegeri_v)(GLenum target, GLuint index, GLint* data);
// void (APIENTRYP ptrglEnablei)(GLenum target, GLuint index);
// void (APIENTRYP ptrglDisablei)(GLenum target, GLuint index);
// GLboolean (APIENTRYP ptrglIsEnabledi)(GLenum target, GLuint index);
// void (APIENTRYP ptrglBeginTransformFeedback)(GLenum primitiveMode);
// void (APIENTRYP ptrglEndTransformFeedback)();
// void (APIENTRYP ptrglBindBufferRange)(GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size);
// void (APIENTRYP ptrglBindBufferBase)(GLenum target, GLuint index, GLuint buffer);
// void (APIENTRYP ptrglTransformFeedbackVaryings)(GLuint program, GLsizei count, GLchar* const* varyings, GLenum bufferMode);
// void (APIENTRYP ptrglGetTransformFeedbackVarying)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLsizei* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrglClampColor)(GLenum target, GLenum clamp);
// void (APIENTRYP ptrglBeginConditionalRender)(GLuint id, GLenum mode);
// void (APIENTRYP ptrglEndConditionalRender)();
// void (APIENTRYP ptrglVertexAttribIPointer)(GLuint index, GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglGetVertexAttribIiv)(GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetVertexAttribIuiv)(GLuint index, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglVertexAttribI1i)(GLuint index, GLint x);
// void (APIENTRYP ptrglVertexAttribI2i)(GLuint index, GLint x, GLint y);
// void (APIENTRYP ptrglVertexAttribI3i)(GLuint index, GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglVertexAttribI4i)(GLuint index, GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglVertexAttribI1ui)(GLuint index, GLuint x);
// void (APIENTRYP ptrglVertexAttribI2ui)(GLuint index, GLuint x, GLuint y);
// void (APIENTRYP ptrglVertexAttribI3ui)(GLuint index, GLuint x, GLuint y, GLuint z);
// void (APIENTRYP ptrglVertexAttribI4ui)(GLuint index, GLuint x, GLuint y, GLuint z, GLuint w);
// void (APIENTRYP ptrglVertexAttribI1iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglVertexAttribI2iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglVertexAttribI3iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglVertexAttribI4iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglVertexAttribI1uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttribI2uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttribI3uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttribI4uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttribI4bv)(GLuint index, GLbyte* v);
// void (APIENTRYP ptrglVertexAttribI4sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttribI4ubv)(GLuint index, GLubyte* v);
// void (APIENTRYP ptrglVertexAttribI4usv)(GLuint index, GLushort* v);
// void (APIENTRYP ptrglGetUniformuiv)(GLuint program, GLint location, GLuint* params);
// void (APIENTRYP ptrglBindFragDataLocation)(GLuint program, GLuint color, GLchar* name);
// GLint (APIENTRYP ptrglGetFragDataLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglUniform1ui)(GLint location, GLuint v0);
// void (APIENTRYP ptrglUniform2ui)(GLint location, GLuint v0, GLuint v1);
// void (APIENTRYP ptrglUniform3ui)(GLint location, GLuint v0, GLuint v1, GLuint v2);
// void (APIENTRYP ptrglUniform4ui)(GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3);
// void (APIENTRYP ptrglUniform1uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglUniform2uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglUniform3uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglUniform4uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglTexParameterIiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglTexParameterIuiv)(GLenum target, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglGetTexParameterIiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetTexParameterIuiv)(GLenum target, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglClearBufferiv)(GLenum buffer, GLint drawbuffer, GLint* value);
// void (APIENTRYP ptrglClearBufferuiv)(GLenum buffer, GLint drawbuffer, GLuint* value);
// void (APIENTRYP ptrglClearBufferfv)(GLenum buffer, GLint drawbuffer, GLfloat* value);
// void (APIENTRYP ptrglClearBufferfi)(GLenum buffer, GLint drawbuffer, GLfloat depth, GLint stencil);
// const GLubyte * (APIENTRYP ptrglGetStringi)(GLenum name, GLuint index);
// GLboolean (APIENTRYP ptrglIsRenderbuffer)(GLuint renderbuffer);
// void (APIENTRYP ptrglBindRenderbuffer)(GLenum target, GLuint renderbuffer);
// void (APIENTRYP ptrglDeleteRenderbuffers)(GLsizei n, GLuint* renderbuffers);
// void (APIENTRYP ptrglGenRenderbuffers)(GLsizei n, GLuint* renderbuffers);
// void (APIENTRYP ptrglRenderbufferStorage)(GLenum target, GLenum internalformat, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglGetRenderbufferParameteriv)(GLenum target, GLenum pname, GLint* params);
// GLboolean (APIENTRYP ptrglIsFramebuffer)(GLuint framebuffer);
// void (APIENTRYP ptrglBindFramebuffer)(GLenum target, GLuint framebuffer);
// void (APIENTRYP ptrglDeleteFramebuffers)(GLsizei n, GLuint* framebuffers);
// void (APIENTRYP ptrglGenFramebuffers)(GLsizei n, GLuint* framebuffers);
// GLenum (APIENTRYP ptrglCheckFramebufferStatus)(GLenum target);
// void (APIENTRYP ptrglFramebufferTexture1D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
// void (APIENTRYP ptrglFramebufferTexture2D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
// void (APIENTRYP ptrglFramebufferTexture3D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint zoffset);
// void (APIENTRYP ptrglFramebufferRenderbuffer)(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer);
// void (APIENTRYP ptrglGetFramebufferAttachmentParameteriv)(GLenum target, GLenum attachment, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGenerateMipmap)(GLenum target);
// void (APIENTRYP ptrglBlitFramebuffer)(GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
// void (APIENTRYP ptrglRenderbufferStorageMultisample)(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglFramebufferTextureLayer)(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer);
// GLvoid* (APIENTRYP ptrglMapBufferRange)(GLenum target, GLintptr offset, GLsizeiptr length, GLbitfield access);
// void (APIENTRYP ptrglFlushMappedBufferRange)(GLenum target, GLintptr offset, GLsizeiptr length);
// void (APIENTRYP ptrglBindVertexArray)(GLuint array);
// void (APIENTRYP ptrglDeleteVertexArrays)(GLsizei n, GLuint* arrays);
// void (APIENTRYP ptrglGenVertexArrays)(GLsizei n, GLuint* arrays);
// GLboolean (APIENTRYP ptrglIsVertexArray)(GLuint array);
// //  VERSION_3_1
// void (APIENTRYP ptrglDrawArraysInstanced)(GLenum mode, GLint first, GLsizei count, GLsizei primcount);
// void (APIENTRYP ptrglDrawElementsInstanced)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei primcount);
// void (APIENTRYP ptrglTexBuffer)(GLenum target, GLenum internalformat, GLuint buffer);
// void (APIENTRYP ptrglPrimitiveRestartIndex)(GLuint index);
// void (APIENTRYP ptrglCopyBufferSubData)(GLenum readTarget, GLenum writeTarget, GLintptr readOffset, GLintptr writeOffset, GLsizeiptr size);
// void (APIENTRYP ptrglGetUniformIndices)(GLuint program, GLsizei uniformCount, GLchar* const* uniformNames, GLuint* uniformIndices);
// void (APIENTRYP ptrglGetActiveUniformsiv)(GLuint program, GLsizei uniformCount, GLuint* uniformIndices, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetActiveUniformName)(GLuint program, GLuint uniformIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformName);
// GLuint (APIENTRYP ptrglGetUniformBlockIndex)(GLuint program, GLchar* uniformBlockName);
// void (APIENTRYP ptrglGetActiveUniformBlockiv)(GLuint program, GLuint uniformBlockIndex, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetActiveUniformBlockName)(GLuint program, GLuint uniformBlockIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformBlockName);
// void (APIENTRYP ptrglUniformBlockBinding)(GLuint program, GLuint uniformBlockIndex, GLuint uniformBlockBinding);
// //  VERSION_3_2
// void (APIENTRYP ptrglGetInteger64i_v)(GLenum target, GLuint index, GLint64* data);
// void (APIENTRYP ptrglGetBufferParameteri64v)(GLenum target, GLenum pname, GLint64* params);
// void (APIENTRYP ptrglFramebufferTexture)(GLenum target, GLenum attachment, GLuint texture, GLint level);
// void (APIENTRYP ptrglDrawElementsBaseVertex)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLint basevertex);
// void (APIENTRYP ptrglDrawRangeElementsBaseVertex)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices, GLint basevertex);
// void (APIENTRYP ptrglDrawElementsInstancedBaseVertex)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei primcount, GLint basevertex);
// void (APIENTRYP ptrglMultiDrawElementsBaseVertex)(GLenum mode, GLsizei* count, GLenum type, GLvoid* const* indices, GLsizei primcount, GLint* basevertex);
// void (APIENTRYP ptrglProvokingVertex)(GLenum mode);
// GLsync (APIENTRYP ptrglFenceSync)(GLenum condition, GLbitfield flags);
// GLboolean (APIENTRYP ptrglIsSync)(GLsync sync);
// void (APIENTRYP ptrglDeleteSync)(GLsync sync);
// GLenum (APIENTRYP ptrglClientWaitSync)(GLsync sync, GLbitfield flags, GLuint64 timeout);
// void (APIENTRYP ptrglWaitSync)(GLsync sync, GLbitfield flags, GLuint64 timeout);
// void (APIENTRYP ptrglGetInteger64v)(GLenum pname, GLint64* params);
// void (APIENTRYP ptrglGetSynciv)(GLsync sync, GLenum pname, GLsizei bufSize, GLsizei* length, GLint* values);
// void (APIENTRYP ptrglTexImage2DMultisample)(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrglTexImage3DMultisample)(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrglGetMultisamplefv)(GLenum pname, GLuint index, GLfloat* val);
// void (APIENTRYP ptrglSampleMaski)(GLuint index, GLbitfield mask);
// //  VERSION_3_3
// void (APIENTRYP ptrglVertexAttribDivisor)(GLuint index, GLuint divisor);
// void (APIENTRYP ptrglBindFragDataLocationIndexed)(GLuint program, GLuint colorNumber, GLuint index, GLchar* name);
// GLint (APIENTRYP ptrglGetFragDataIndex)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglGenSamplers)(GLsizei count, GLuint* samplers);
// void (APIENTRYP ptrglDeleteSamplers)(GLsizei count, GLuint* samplers);
// GLboolean (APIENTRYP ptrglIsSampler)(GLuint sampler);
// void (APIENTRYP ptrglBindSampler)(GLuint unit, GLuint sampler);
// void (APIENTRYP ptrglSamplerParameteri)(GLuint sampler, GLenum pname, GLint param);
// void (APIENTRYP ptrglSamplerParameteriv)(GLuint sampler, GLenum pname, GLint* param);
// void (APIENTRYP ptrglSamplerParameterf)(GLuint sampler, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglSamplerParameterfv)(GLuint sampler, GLenum pname, GLfloat* param);
// void (APIENTRYP ptrglSamplerParameterIiv)(GLuint sampler, GLenum pname, GLint* param);
// void (APIENTRYP ptrglSamplerParameterIuiv)(GLuint sampler, GLenum pname, GLuint* param);
// void (APIENTRYP ptrglGetSamplerParameteriv)(GLuint sampler, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetSamplerParameterIiv)(GLuint sampler, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetSamplerParameterfv)(GLuint sampler, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetSamplerParameterIuiv)(GLuint sampler, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglQueryCounter)(GLuint id, GLenum target);
// void (APIENTRYP ptrglGetQueryObjecti64v)(GLuint id, GLenum pname, GLint64* params);
// void (APIENTRYP ptrglGetQueryObjectui64v)(GLuint id, GLenum pname, GLuint64* params);
// void (APIENTRYP ptrglVertexP2ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrglVertexP2uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrglVertexP3ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrglVertexP3uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrglVertexP4ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrglVertexP4uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrglTexCoordP1ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglTexCoordP1uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglTexCoordP2ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglTexCoordP2uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglTexCoordP3ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglTexCoordP3uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglTexCoordP4ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglTexCoordP4uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglMultiTexCoordP1ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglMultiTexCoordP1uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglMultiTexCoordP2ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglMultiTexCoordP2uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglMultiTexCoordP3ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglMultiTexCoordP3uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglMultiTexCoordP4ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglMultiTexCoordP4uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglNormalP3ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglNormalP3uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglColorP3ui)(GLenum type, GLuint color);
// void (APIENTRYP ptrglColorP3uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrglColorP4ui)(GLenum type, GLuint color);
// void (APIENTRYP ptrglColorP4uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrglSecondaryColorP3ui)(GLenum type, GLuint color);
// void (APIENTRYP ptrglSecondaryColorP3uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrglVertexAttribP1ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglVertexAttribP1uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrglVertexAttribP2ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglVertexAttribP2uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrglVertexAttribP3ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglVertexAttribP3uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrglVertexAttribP4ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglVertexAttribP4uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// //  VERSION_4_0
// void (APIENTRYP ptrglMinSampleShading)(GLfloat value);
// void (APIENTRYP ptrglBlendEquationi)(GLuint buf, GLenum mode);
// void (APIENTRYP ptrglBlendEquationSeparatei)(GLuint buf, GLenum modeRGB, GLenum modeAlpha);
// void (APIENTRYP ptrglBlendFunci)(GLuint buf, GLenum src, GLenum dst);
// void (APIENTRYP ptrglBlendFuncSeparatei)(GLuint buf, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha);
// void (APIENTRYP ptrglDrawArraysIndirect)(GLenum mode, GLvoid* indirect);
// void (APIENTRYP ptrglDrawElementsIndirect)(GLenum mode, GLenum type, GLvoid* indirect);
// void (APIENTRYP ptrglUniform1d)(GLint location, GLdouble x);
// void (APIENTRYP ptrglUniform2d)(GLint location, GLdouble x, GLdouble y);
// void (APIENTRYP ptrglUniform3d)(GLint location, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglUniform4d)(GLint location, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglUniform1dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglUniform2dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglUniform3dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglUniform4dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix2dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix3dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix4dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix2x3dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix2x4dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix3x2dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix3x4dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix4x2dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix4x3dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglGetUniformdv)(GLuint program, GLint location, GLdouble* params);
// GLint (APIENTRYP ptrglGetSubroutineUniformLocation)(GLuint program, GLenum shadertype, GLchar* name);
// GLuint (APIENTRYP ptrglGetSubroutineIndex)(GLuint program, GLenum shadertype, GLchar* name);
// void (APIENTRYP ptrglGetActiveSubroutineUniformiv)(GLuint program, GLenum shadertype, GLuint index, GLenum pname, GLint* values);
// void (APIENTRYP ptrglGetActiveSubroutineUniformName)(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name);
// void (APIENTRYP ptrglGetActiveSubroutineName)(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name);
// void (APIENTRYP ptrglUniformSubroutinesuiv)(GLenum shadertype, GLsizei count, GLuint* indices);
// void (APIENTRYP ptrglGetUniformSubroutineuiv)(GLenum shadertype, GLint location, GLuint* params);
// void (APIENTRYP ptrglGetProgramStageiv)(GLuint program, GLenum shadertype, GLenum pname, GLint* values);
// void (APIENTRYP ptrglPatchParameteri)(GLenum pname, GLint value);
// void (APIENTRYP ptrglPatchParameterfv)(GLenum pname, GLfloat* values);
// void (APIENTRYP ptrglBindTransformFeedback)(GLenum target, GLuint id);
// void (APIENTRYP ptrglDeleteTransformFeedbacks)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglGenTransformFeedbacks)(GLsizei n, GLuint* ids);
// GLboolean (APIENTRYP ptrglIsTransformFeedback)(GLuint id);
// void (APIENTRYP ptrglPauseTransformFeedback)();
// void (APIENTRYP ptrglResumeTransformFeedback)();
// void (APIENTRYP ptrglDrawTransformFeedback)(GLenum mode, GLuint id);
// void (APIENTRYP ptrglDrawTransformFeedbackStream)(GLenum mode, GLuint id, GLuint stream);
// void (APIENTRYP ptrglBeginQueryIndexed)(GLenum target, GLuint index, GLuint id);
// void (APIENTRYP ptrglEndQueryIndexed)(GLenum target, GLuint index);
// void (APIENTRYP ptrglGetQueryIndexediv)(GLenum target, GLuint index, GLenum pname, GLint* params);
// //  VERSION_4_1
// void (APIENTRYP ptrglReleaseShaderCompiler)();
// void (APIENTRYP ptrglShaderBinary)(GLsizei count, GLuint* shaders, GLenum binaryformat, GLvoid* binary, GLsizei length);
// void (APIENTRYP ptrglGetShaderPrecisionFormat)(GLenum shadertype, GLenum precisiontype, GLint* range, GLint* precision);
// void (APIENTRYP ptrglDepthRangef)(GLfloat n, GLfloat f);
// void (APIENTRYP ptrglClearDepthf)(GLfloat d);
// void (APIENTRYP ptrglGetProgramBinary)(GLuint program, GLsizei bufSize, GLsizei* length, GLenum* binaryFormat, GLvoid* binary);
// void (APIENTRYP ptrglProgramBinary)(GLuint program, GLenum binaryFormat, GLvoid* binary, GLsizei length);
// void (APIENTRYP ptrglProgramParameteri)(GLuint program, GLenum pname, GLint value);
// void (APIENTRYP ptrglUseProgramStages)(GLuint pipeline, GLbitfield stages, GLuint program);
// void (APIENTRYP ptrglActiveShaderProgram)(GLuint pipeline, GLuint program);
// GLuint (APIENTRYP ptrglCreateShaderProgramv)(GLenum type, GLsizei count, GLchar* const* strings);
// void (APIENTRYP ptrglBindProgramPipeline)(GLuint pipeline);
// void (APIENTRYP ptrglDeleteProgramPipelines)(GLsizei n, GLuint* pipelines);
// void (APIENTRYP ptrglGenProgramPipelines)(GLsizei n, GLuint* pipelines);
// GLboolean (APIENTRYP ptrglIsProgramPipeline)(GLuint pipeline);
// void (APIENTRYP ptrglGetProgramPipelineiv)(GLuint pipeline, GLenum pname, GLint* params);
// void (APIENTRYP ptrglProgramUniform1i)(GLuint program, GLint location, GLint v0);
// void (APIENTRYP ptrglProgramUniform1iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglProgramUniform1f)(GLuint program, GLint location, GLfloat v0);
// void (APIENTRYP ptrglProgramUniform1fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglProgramUniform1d)(GLuint program, GLint location, GLdouble v0);
// void (APIENTRYP ptrglProgramUniform1dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglProgramUniform1ui)(GLuint program, GLint location, GLuint v0);
// void (APIENTRYP ptrglProgramUniform1uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglProgramUniform2i)(GLuint program, GLint location, GLint v0, GLint v1);
// void (APIENTRYP ptrglProgramUniform2iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglProgramUniform2f)(GLuint program, GLint location, GLfloat v0, GLfloat v1);
// void (APIENTRYP ptrglProgramUniform2fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglProgramUniform2d)(GLuint program, GLint location, GLdouble v0, GLdouble v1);
// void (APIENTRYP ptrglProgramUniform2dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglProgramUniform2ui)(GLuint program, GLint location, GLuint v0, GLuint v1);
// void (APIENTRYP ptrglProgramUniform2uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglProgramUniform3i)(GLuint program, GLint location, GLint v0, GLint v1, GLint v2);
// void (APIENTRYP ptrglProgramUniform3iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglProgramUniform3f)(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2);
// void (APIENTRYP ptrglProgramUniform3fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglProgramUniform3d)(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2);
// void (APIENTRYP ptrglProgramUniform3dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglProgramUniform3ui)(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2);
// void (APIENTRYP ptrglProgramUniform3uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglProgramUniform4i)(GLuint program, GLint location, GLint v0, GLint v1, GLint v2, GLint v3);
// void (APIENTRYP ptrglProgramUniform4iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglProgramUniform4f)(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3);
// void (APIENTRYP ptrglProgramUniform4fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglProgramUniform4d)(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3);
// void (APIENTRYP ptrglProgramUniform4dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglProgramUniform4ui)(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3);
// void (APIENTRYP ptrglProgramUniform4uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglProgramUniformMatrix2fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix3fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix4fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix2dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniformMatrix3dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniformMatrix4dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniformMatrix2x3fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix3x2fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix2x4fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix4x2fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix3x4fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix4x3fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix2x3dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniformMatrix3x2dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniformMatrix2x4dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniformMatrix4x2dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniformMatrix3x4dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniformMatrix4x3dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglValidateProgramPipeline)(GLuint pipeline);
// void (APIENTRYP ptrglGetProgramPipelineInfoLog)(GLuint pipeline, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrglVertexAttribL1d)(GLuint index, GLdouble x);
// void (APIENTRYP ptrglVertexAttribL2d)(GLuint index, GLdouble x, GLdouble y);
// void (APIENTRYP ptrglVertexAttribL3d)(GLuint index, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglVertexAttribL4d)(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglVertexAttribL1dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttribL2dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttribL3dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttribL4dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttribLPointer)(GLuint index, GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglGetVertexAttribLdv)(GLuint index, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglViewportArrayv)(GLuint first, GLsizei count, GLfloat* v);
// void (APIENTRYP ptrglViewportIndexedf)(GLuint index, GLfloat x, GLfloat y, GLfloat w, GLfloat h);
// void (APIENTRYP ptrglViewportIndexedfv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglScissorArrayv)(GLuint first, GLsizei count, GLint* v);
// void (APIENTRYP ptrglScissorIndexed)(GLuint index, GLint left, GLint bottom, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglScissorIndexedv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglDepthRangeArrayv)(GLuint first, GLsizei count, GLdouble* v);
// void (APIENTRYP ptrglDepthRangeIndexed)(GLuint index, GLdouble n, GLdouble f);
// void (APIENTRYP ptrglGetFloati_v)(GLenum target, GLuint index, GLfloat* data);
// void (APIENTRYP ptrglGetDoublei_v)(GLenum target, GLuint index, GLdouble* data);
// //  VERSION_4_2
// void (APIENTRYP ptrglDrawArraysInstancedBaseInstance)(GLenum mode, GLint first, GLsizei count, GLsizei primcount, GLuint baseinstance);
// void (APIENTRYP ptrglDrawElementsInstancedBaseInstance)(GLenum mode, GLsizei count, GLenum type, void* indices, GLsizei primcount, GLuint baseinstance);
// void (APIENTRYP ptrglDrawElementsInstancedBaseVertexBaseInstance)(GLenum mode, GLsizei count, GLenum type, void* indices, GLsizei primcount, GLint basevertex, GLuint baseinstance);
// void (APIENTRYP ptrglDrawTransformFeedbackInstanced)(GLenum mode, GLuint id, GLsizei primcount);
// void (APIENTRYP ptrglDrawTransformFeedbackStreamInstanced)(GLenum mode, GLuint id, GLuint stream, GLsizei primcount);
// void (APIENTRYP ptrglGetInternalformativ)(GLenum target, GLenum internalformat, GLenum pname, GLsizei bufSize, GLint* params);
// void (APIENTRYP ptrglGetActiveAtomicCounterBufferiv)(GLuint program, GLuint bufferIndex, GLenum pname, GLint* params);
// void (APIENTRYP ptrglBindImageTexture)(GLuint unit, GLuint texture, GLint level, GLboolean layered, GLint layer, GLenum access, GLenum format);
// void (APIENTRYP ptrglMemoryBarrier)(GLbitfield barriers);
// void (APIENTRYP ptrglTexStorage1D)(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width);
// void (APIENTRYP ptrglTexStorage2D)(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglTexStorage3D)(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth);
// void (APIENTRYP ptrglTextureStorage1DEXT)(GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width);
// void (APIENTRYP ptrglTextureStorage2DEXT)(GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglTextureStorage3DEXT)(GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth);
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
// void goglClearColor(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
// 	(*ptrglClearColor)(red, green, blue, alpha);
// }
// void goglClearStencil(GLint s) {
// 	(*ptrglClearStencil)(s);
// }
// void goglClearDepth(GLdouble depth) {
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
// void goglDepthRange(GLdouble near_, GLdouble far_) {
// 	(*ptrglDepthRange)(near_, far_);
// }
// void goglViewport(GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrglViewport)(x, y, width, height);
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
// //  VERSION_1_2
// void goglBlendColor(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
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
// //  VERSION_1_3
// void goglActiveTexture(GLenum texture) {
// 	(*ptrglActiveTexture)(texture);
// }
// void goglSampleCoverage(GLfloat value, GLboolean invert) {
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
// //  VERSION_1_4
// void goglBlendFuncSeparate(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha) {
// 	(*ptrglBlendFuncSeparate)(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha);
// }
// void goglMultiDrawArrays(GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
// 	(*ptrglMultiDrawArrays)(mode, first, count, primcount);
// }
// void goglMultiDrawElements(GLenum mode, GLsizei* count, GLenum type_, GLvoid* const* indices, GLsizei primcount) {
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
// void goglShaderSource(GLuint shader, GLsizei count, GLchar* const* string_, GLint* length) {
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
// //  VERSION_3_0
// void goglColorMaski(GLuint index, GLboolean r, GLboolean g, GLboolean b, GLboolean a) {
// 	(*ptrglColorMaski)(index, r, g, b, a);
// }
// void goglGetBooleani_v(GLenum target, GLuint index, GLboolean* data) {
// 	(*ptrglGetBooleani_v)(target, index, data);
// }
// void goglGetIntegeri_v(GLenum target, GLuint index, GLint* data) {
// 	(*ptrglGetIntegeri_v)(target, index, data);
// }
// void goglEnablei(GLenum target, GLuint index) {
// 	(*ptrglEnablei)(target, index);
// }
// void goglDisablei(GLenum target, GLuint index) {
// 	(*ptrglDisablei)(target, index);
// }
// GLboolean goglIsEnabledi(GLenum target, GLuint index) {
// 	return (*ptrglIsEnabledi)(target, index);
// }
// void goglBeginTransformFeedback(GLenum primitiveMode) {
// 	(*ptrglBeginTransformFeedback)(primitiveMode);
// }
// void goglEndTransformFeedback() {
// 	(*ptrglEndTransformFeedback)();
// }
// void goglBindBufferRange(GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size) {
// 	(*ptrglBindBufferRange)(target, index, buffer, offset, size);
// }
// void goglBindBufferBase(GLenum target, GLuint index, GLuint buffer) {
// 	(*ptrglBindBufferBase)(target, index, buffer);
// }
// void goglTransformFeedbackVaryings(GLuint program, GLsizei count, GLchar* const* varyings, GLenum bufferMode) {
// 	(*ptrglTransformFeedbackVaryings)(program, count, varyings, bufferMode);
// }
// void goglGetTransformFeedbackVarying(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLsizei* size, GLenum* type_, GLchar* name) {
// 	(*ptrglGetTransformFeedbackVarying)(program, index, bufSize, length, size, type_, name);
// }
// void goglClampColor(GLenum target, GLenum clamp) {
// 	(*ptrglClampColor)(target, clamp);
// }
// void goglBeginConditionalRender(GLuint id, GLenum mode) {
// 	(*ptrglBeginConditionalRender)(id, mode);
// }
// void goglEndConditionalRender() {
// 	(*ptrglEndConditionalRender)();
// }
// void goglVertexAttribIPointer(GLuint index, GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglVertexAttribIPointer)(index, size, type_, stride, pointer);
// }
// void goglGetVertexAttribIiv(GLuint index, GLenum pname, GLint* params) {
// 	(*ptrglGetVertexAttribIiv)(index, pname, params);
// }
// void goglGetVertexAttribIuiv(GLuint index, GLenum pname, GLuint* params) {
// 	(*ptrglGetVertexAttribIuiv)(index, pname, params);
// }
// void goglVertexAttribI1i(GLuint index, GLint x) {
// 	(*ptrglVertexAttribI1i)(index, x);
// }
// void goglVertexAttribI2i(GLuint index, GLint x, GLint y) {
// 	(*ptrglVertexAttribI2i)(index, x, y);
// }
// void goglVertexAttribI3i(GLuint index, GLint x, GLint y, GLint z) {
// 	(*ptrglVertexAttribI3i)(index, x, y, z);
// }
// void goglVertexAttribI4i(GLuint index, GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrglVertexAttribI4i)(index, x, y, z, w);
// }
// void goglVertexAttribI1ui(GLuint index, GLuint x) {
// 	(*ptrglVertexAttribI1ui)(index, x);
// }
// void goglVertexAttribI2ui(GLuint index, GLuint x, GLuint y) {
// 	(*ptrglVertexAttribI2ui)(index, x, y);
// }
// void goglVertexAttribI3ui(GLuint index, GLuint x, GLuint y, GLuint z) {
// 	(*ptrglVertexAttribI3ui)(index, x, y, z);
// }
// void goglVertexAttribI4ui(GLuint index, GLuint x, GLuint y, GLuint z, GLuint w) {
// 	(*ptrglVertexAttribI4ui)(index, x, y, z, w);
// }
// void goglVertexAttribI1iv(GLuint index, GLint* v) {
// 	(*ptrglVertexAttribI1iv)(index, v);
// }
// void goglVertexAttribI2iv(GLuint index, GLint* v) {
// 	(*ptrglVertexAttribI2iv)(index, v);
// }
// void goglVertexAttribI3iv(GLuint index, GLint* v) {
// 	(*ptrglVertexAttribI3iv)(index, v);
// }
// void goglVertexAttribI4iv(GLuint index, GLint* v) {
// 	(*ptrglVertexAttribI4iv)(index, v);
// }
// void goglVertexAttribI1uiv(GLuint index, GLuint* v) {
// 	(*ptrglVertexAttribI1uiv)(index, v);
// }
// void goglVertexAttribI2uiv(GLuint index, GLuint* v) {
// 	(*ptrglVertexAttribI2uiv)(index, v);
// }
// void goglVertexAttribI3uiv(GLuint index, GLuint* v) {
// 	(*ptrglVertexAttribI3uiv)(index, v);
// }
// void goglVertexAttribI4uiv(GLuint index, GLuint* v) {
// 	(*ptrglVertexAttribI4uiv)(index, v);
// }
// void goglVertexAttribI4bv(GLuint index, GLbyte* v) {
// 	(*ptrglVertexAttribI4bv)(index, v);
// }
// void goglVertexAttribI4sv(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttribI4sv)(index, v);
// }
// void goglVertexAttribI4ubv(GLuint index, GLubyte* v) {
// 	(*ptrglVertexAttribI4ubv)(index, v);
// }
// void goglVertexAttribI4usv(GLuint index, GLushort* v) {
// 	(*ptrglVertexAttribI4usv)(index, v);
// }
// void goglGetUniformuiv(GLuint program, GLint location, GLuint* params) {
// 	(*ptrglGetUniformuiv)(program, location, params);
// }
// void goglBindFragDataLocation(GLuint program, GLuint color, GLchar* name) {
// 	(*ptrglBindFragDataLocation)(program, color, name);
// }
// GLint goglGetFragDataLocation(GLuint program, GLchar* name) {
// 	return (*ptrglGetFragDataLocation)(program, name);
// }
// void goglUniform1ui(GLint location, GLuint v0) {
// 	(*ptrglUniform1ui)(location, v0);
// }
// void goglUniform2ui(GLint location, GLuint v0, GLuint v1) {
// 	(*ptrglUniform2ui)(location, v0, v1);
// }
// void goglUniform3ui(GLint location, GLuint v0, GLuint v1, GLuint v2) {
// 	(*ptrglUniform3ui)(location, v0, v1, v2);
// }
// void goglUniform4ui(GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3) {
// 	(*ptrglUniform4ui)(location, v0, v1, v2, v3);
// }
// void goglUniform1uiv(GLint location, GLsizei count, GLuint* value) {
// 	(*ptrglUniform1uiv)(location, count, value);
// }
// void goglUniform2uiv(GLint location, GLsizei count, GLuint* value) {
// 	(*ptrglUniform2uiv)(location, count, value);
// }
// void goglUniform3uiv(GLint location, GLsizei count, GLuint* value) {
// 	(*ptrglUniform3uiv)(location, count, value);
// }
// void goglUniform4uiv(GLint location, GLsizei count, GLuint* value) {
// 	(*ptrglUniform4uiv)(location, count, value);
// }
// void goglTexParameterIiv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglTexParameterIiv)(target, pname, params);
// }
// void goglTexParameterIuiv(GLenum target, GLenum pname, GLuint* params) {
// 	(*ptrglTexParameterIuiv)(target, pname, params);
// }
// void goglGetTexParameterIiv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetTexParameterIiv)(target, pname, params);
// }
// void goglGetTexParameterIuiv(GLenum target, GLenum pname, GLuint* params) {
// 	(*ptrglGetTexParameterIuiv)(target, pname, params);
// }
// void goglClearBufferiv(GLenum buffer, GLint drawbuffer, GLint* value) {
// 	(*ptrglClearBufferiv)(buffer, drawbuffer, value);
// }
// void goglClearBufferuiv(GLenum buffer, GLint drawbuffer, GLuint* value) {
// 	(*ptrglClearBufferuiv)(buffer, drawbuffer, value);
// }
// void goglClearBufferfv(GLenum buffer, GLint drawbuffer, GLfloat* value) {
// 	(*ptrglClearBufferfv)(buffer, drawbuffer, value);
// }
// void goglClearBufferfi(GLenum buffer, GLint drawbuffer, GLfloat depth, GLint stencil) {
// 	(*ptrglClearBufferfi)(buffer, drawbuffer, depth, stencil);
// }
// const GLubyte * goglGetStringi(GLenum name, GLuint index) {
// 	return (*ptrglGetStringi)(name, index);
// }
// GLboolean goglIsRenderbuffer(GLuint renderbuffer) {
// 	return (*ptrglIsRenderbuffer)(renderbuffer);
// }
// void goglBindRenderbuffer(GLenum target, GLuint renderbuffer) {
// 	(*ptrglBindRenderbuffer)(target, renderbuffer);
// }
// void goglDeleteRenderbuffers(GLsizei n, GLuint* renderbuffers) {
// 	(*ptrglDeleteRenderbuffers)(n, renderbuffers);
// }
// void goglGenRenderbuffers(GLsizei n, GLuint* renderbuffers) {
// 	(*ptrglGenRenderbuffers)(n, renderbuffers);
// }
// void goglRenderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height) {
// 	(*ptrglRenderbufferStorage)(target, internalformat, width, height);
// }
// void goglGetRenderbufferParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetRenderbufferParameteriv)(target, pname, params);
// }
// GLboolean goglIsFramebuffer(GLuint framebuffer) {
// 	return (*ptrglIsFramebuffer)(framebuffer);
// }
// void goglBindFramebuffer(GLenum target, GLuint framebuffer) {
// 	(*ptrglBindFramebuffer)(target, framebuffer);
// }
// void goglDeleteFramebuffers(GLsizei n, GLuint* framebuffers) {
// 	(*ptrglDeleteFramebuffers)(n, framebuffers);
// }
// void goglGenFramebuffers(GLsizei n, GLuint* framebuffers) {
// 	(*ptrglGenFramebuffers)(n, framebuffers);
// }
// GLenum goglCheckFramebufferStatus(GLenum target) {
// 	return (*ptrglCheckFramebufferStatus)(target);
// }
// void goglFramebufferTexture1D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level) {
// 	(*ptrglFramebufferTexture1D)(target, attachment, textarget, texture, level);
// }
// void goglFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level) {
// 	(*ptrglFramebufferTexture2D)(target, attachment, textarget, texture, level);
// }
// void goglFramebufferTexture3D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint zoffset) {
// 	(*ptrglFramebufferTexture3D)(target, attachment, textarget, texture, level, zoffset);
// }
// void goglFramebufferRenderbuffer(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer) {
// 	(*ptrglFramebufferRenderbuffer)(target, attachment, renderbuffertarget, renderbuffer);
// }
// void goglGetFramebufferAttachmentParameteriv(GLenum target, GLenum attachment, GLenum pname, GLint* params) {
// 	(*ptrglGetFramebufferAttachmentParameteriv)(target, attachment, pname, params);
// }
// void goglGenerateMipmap(GLenum target) {
// 	(*ptrglGenerateMipmap)(target);
// }
// void goglBlitFramebuffer(GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter) {
// 	(*ptrglBlitFramebuffer)(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// void goglRenderbufferStorageMultisample(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height) {
// 	(*ptrglRenderbufferStorageMultisample)(target, samples, internalformat, width, height);
// }
// void goglFramebufferTextureLayer(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer) {
// 	(*ptrglFramebufferTextureLayer)(target, attachment, texture, level, layer);
// }
// GLvoid* goglMapBufferRange(GLenum target, GLintptr offset, GLsizeiptr length, GLbitfield access) {
// 	return (*ptrglMapBufferRange)(target, offset, length, access);
// }
// void goglFlushMappedBufferRange(GLenum target, GLintptr offset, GLsizeiptr length) {
// 	(*ptrglFlushMappedBufferRange)(target, offset, length);
// }
// void goglBindVertexArray(GLuint array) {
// 	(*ptrglBindVertexArray)(array);
// }
// void goglDeleteVertexArrays(GLsizei n, GLuint* arrays) {
// 	(*ptrglDeleteVertexArrays)(n, arrays);
// }
// void goglGenVertexArrays(GLsizei n, GLuint* arrays) {
// 	(*ptrglGenVertexArrays)(n, arrays);
// }
// GLboolean goglIsVertexArray(GLuint array) {
// 	return (*ptrglIsVertexArray)(array);
// }
// //  VERSION_3_1
// void goglDrawArraysInstanced(GLenum mode, GLint first, GLsizei count, GLsizei primcount) {
// 	(*ptrglDrawArraysInstanced)(mode, first, count, primcount);
// }
// void goglDrawElementsInstanced(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLsizei primcount) {
// 	(*ptrglDrawElementsInstanced)(mode, count, type_, indices, primcount);
// }
// void goglTexBuffer(GLenum target, GLenum internalformat, GLuint buffer) {
// 	(*ptrglTexBuffer)(target, internalformat, buffer);
// }
// void goglPrimitiveRestartIndex(GLuint index) {
// 	(*ptrglPrimitiveRestartIndex)(index);
// }
// void goglCopyBufferSubData(GLenum readTarget, GLenum writeTarget, GLintptr readOffset, GLintptr writeOffset, GLsizeiptr size) {
// 	(*ptrglCopyBufferSubData)(readTarget, writeTarget, readOffset, writeOffset, size);
// }
// void goglGetUniformIndices(GLuint program, GLsizei uniformCount, GLchar* const* uniformNames, GLuint* uniformIndices) {
// 	(*ptrglGetUniformIndices)(program, uniformCount, uniformNames, uniformIndices);
// }
// void goglGetActiveUniformsiv(GLuint program, GLsizei uniformCount, GLuint* uniformIndices, GLenum pname, GLint* params) {
// 	(*ptrglGetActiveUniformsiv)(program, uniformCount, uniformIndices, pname, params);
// }
// void goglGetActiveUniformName(GLuint program, GLuint uniformIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformName) {
// 	(*ptrglGetActiveUniformName)(program, uniformIndex, bufSize, length, uniformName);
// }
// GLuint goglGetUniformBlockIndex(GLuint program, GLchar* uniformBlockName) {
// 	return (*ptrglGetUniformBlockIndex)(program, uniformBlockName);
// }
// void goglGetActiveUniformBlockiv(GLuint program, GLuint uniformBlockIndex, GLenum pname, GLint* params) {
// 	(*ptrglGetActiveUniformBlockiv)(program, uniformBlockIndex, pname, params);
// }
// void goglGetActiveUniformBlockName(GLuint program, GLuint uniformBlockIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformBlockName) {
// 	(*ptrglGetActiveUniformBlockName)(program, uniformBlockIndex, bufSize, length, uniformBlockName);
// }
// void goglUniformBlockBinding(GLuint program, GLuint uniformBlockIndex, GLuint uniformBlockBinding) {
// 	(*ptrglUniformBlockBinding)(program, uniformBlockIndex, uniformBlockBinding);
// }
// //  VERSION_3_2
// void goglGetInteger64i_v(GLenum target, GLuint index, GLint64* data) {
// 	(*ptrglGetInteger64i_v)(target, index, data);
// }
// void goglGetBufferParameteri64v(GLenum target, GLenum pname, GLint64* params) {
// 	(*ptrglGetBufferParameteri64v)(target, pname, params);
// }
// void goglFramebufferTexture(GLenum target, GLenum attachment, GLuint texture, GLint level) {
// 	(*ptrglFramebufferTexture)(target, attachment, texture, level);
// }
// void goglDrawElementsBaseVertex(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLint basevertex) {
// 	(*ptrglDrawElementsBaseVertex)(mode, count, type_, indices, basevertex);
// }
// void goglDrawRangeElementsBaseVertex(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type_, GLvoid* indices, GLint basevertex) {
// 	(*ptrglDrawRangeElementsBaseVertex)(mode, start, end, count, type_, indices, basevertex);
// }
// void goglDrawElementsInstancedBaseVertex(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLsizei primcount, GLint basevertex) {
// 	(*ptrglDrawElementsInstancedBaseVertex)(mode, count, type_, indices, primcount, basevertex);
// }
// void goglMultiDrawElementsBaseVertex(GLenum mode, GLsizei* count, GLenum type_, GLvoid* const* indices, GLsizei primcount, GLint* basevertex) {
// 	(*ptrglMultiDrawElementsBaseVertex)(mode, count, type_, indices, primcount, basevertex);
// }
// void goglProvokingVertex(GLenum mode) {
// 	(*ptrglProvokingVertex)(mode);
// }
// GLsync goglFenceSync(GLenum condition, GLbitfield flags) {
// 	return (*ptrglFenceSync)(condition, flags);
// }
// GLboolean goglIsSync(GLsync sync) {
// 	return (*ptrglIsSync)(sync);
// }
// void goglDeleteSync(GLsync sync) {
// 	(*ptrglDeleteSync)(sync);
// }
// GLenum goglClientWaitSync(GLsync sync, GLbitfield flags, GLuint64 timeout) {
// 	return (*ptrglClientWaitSync)(sync, flags, timeout);
// }
// void goglWaitSync(GLsync sync, GLbitfield flags, GLuint64 timeout) {
// 	(*ptrglWaitSync)(sync, flags, timeout);
// }
// void goglGetInteger64v(GLenum pname, GLint64* params) {
// 	(*ptrglGetInteger64v)(pname, params);
// }
// void goglGetSynciv(GLsync sync, GLenum pname, GLsizei bufSize, GLsizei* length, GLint* values) {
// 	(*ptrglGetSynciv)(sync, pname, bufSize, length, values);
// }
// void goglTexImage2DMultisample(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations) {
// 	(*ptrglTexImage2DMultisample)(target, samples, internalformat, width, height, fixedsamplelocations);
// }
// void goglTexImage3DMultisample(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations) {
// 	(*ptrglTexImage3DMultisample)(target, samples, internalformat, width, height, depth, fixedsamplelocations);
// }
// void goglGetMultisamplefv(GLenum pname, GLuint index, GLfloat* val) {
// 	(*ptrglGetMultisamplefv)(pname, index, val);
// }
// void goglSampleMaski(GLuint index, GLbitfield mask) {
// 	(*ptrglSampleMaski)(index, mask);
// }
// //  VERSION_3_3
// void goglVertexAttribDivisor(GLuint index, GLuint divisor) {
// 	(*ptrglVertexAttribDivisor)(index, divisor);
// }
// void goglBindFragDataLocationIndexed(GLuint program, GLuint colorNumber, GLuint index, GLchar* name) {
// 	(*ptrglBindFragDataLocationIndexed)(program, colorNumber, index, name);
// }
// GLint goglGetFragDataIndex(GLuint program, GLchar* name) {
// 	return (*ptrglGetFragDataIndex)(program, name);
// }
// void goglGenSamplers(GLsizei count, GLuint* samplers) {
// 	(*ptrglGenSamplers)(count, samplers);
// }
// void goglDeleteSamplers(GLsizei count, GLuint* samplers) {
// 	(*ptrglDeleteSamplers)(count, samplers);
// }
// GLboolean goglIsSampler(GLuint sampler) {
// 	return (*ptrglIsSampler)(sampler);
// }
// void goglBindSampler(GLuint unit, GLuint sampler) {
// 	(*ptrglBindSampler)(unit, sampler);
// }
// void goglSamplerParameteri(GLuint sampler, GLenum pname, GLint param) {
// 	(*ptrglSamplerParameteri)(sampler, pname, param);
// }
// void goglSamplerParameteriv(GLuint sampler, GLenum pname, GLint* param) {
// 	(*ptrglSamplerParameteriv)(sampler, pname, param);
// }
// void goglSamplerParameterf(GLuint sampler, GLenum pname, GLfloat param) {
// 	(*ptrglSamplerParameterf)(sampler, pname, param);
// }
// void goglSamplerParameterfv(GLuint sampler, GLenum pname, GLfloat* param) {
// 	(*ptrglSamplerParameterfv)(sampler, pname, param);
// }
// void goglSamplerParameterIiv(GLuint sampler, GLenum pname, GLint* param) {
// 	(*ptrglSamplerParameterIiv)(sampler, pname, param);
// }
// void goglSamplerParameterIuiv(GLuint sampler, GLenum pname, GLuint* param) {
// 	(*ptrglSamplerParameterIuiv)(sampler, pname, param);
// }
// void goglGetSamplerParameteriv(GLuint sampler, GLenum pname, GLint* params) {
// 	(*ptrglGetSamplerParameteriv)(sampler, pname, params);
// }
// void goglGetSamplerParameterIiv(GLuint sampler, GLenum pname, GLint* params) {
// 	(*ptrglGetSamplerParameterIiv)(sampler, pname, params);
// }
// void goglGetSamplerParameterfv(GLuint sampler, GLenum pname, GLfloat* params) {
// 	(*ptrglGetSamplerParameterfv)(sampler, pname, params);
// }
// void goglGetSamplerParameterIuiv(GLuint sampler, GLenum pname, GLuint* params) {
// 	(*ptrglGetSamplerParameterIuiv)(sampler, pname, params);
// }
// void goglQueryCounter(GLuint id, GLenum target) {
// 	(*ptrglQueryCounter)(id, target);
// }
// void goglGetQueryObjecti64v(GLuint id, GLenum pname, GLint64* params) {
// 	(*ptrglGetQueryObjecti64v)(id, pname, params);
// }
// void goglGetQueryObjectui64v(GLuint id, GLenum pname, GLuint64* params) {
// 	(*ptrglGetQueryObjectui64v)(id, pname, params);
// }
// void goglVertexP2ui(GLenum type_, GLuint value) {
// 	(*ptrglVertexP2ui)(type_, value);
// }
// void goglVertexP2uiv(GLenum type_, GLuint* value) {
// 	(*ptrglVertexP2uiv)(type_, value);
// }
// void goglVertexP3ui(GLenum type_, GLuint value) {
// 	(*ptrglVertexP3ui)(type_, value);
// }
// void goglVertexP3uiv(GLenum type_, GLuint* value) {
// 	(*ptrglVertexP3uiv)(type_, value);
// }
// void goglVertexP4ui(GLenum type_, GLuint value) {
// 	(*ptrglVertexP4ui)(type_, value);
// }
// void goglVertexP4uiv(GLenum type_, GLuint* value) {
// 	(*ptrglVertexP4uiv)(type_, value);
// }
// void goglTexCoordP1ui(GLenum type_, GLuint coords) {
// 	(*ptrglTexCoordP1ui)(type_, coords);
// }
// void goglTexCoordP1uiv(GLenum type_, GLuint* coords) {
// 	(*ptrglTexCoordP1uiv)(type_, coords);
// }
// void goglTexCoordP2ui(GLenum type_, GLuint coords) {
// 	(*ptrglTexCoordP2ui)(type_, coords);
// }
// void goglTexCoordP2uiv(GLenum type_, GLuint* coords) {
// 	(*ptrglTexCoordP2uiv)(type_, coords);
// }
// void goglTexCoordP3ui(GLenum type_, GLuint coords) {
// 	(*ptrglTexCoordP3ui)(type_, coords);
// }
// void goglTexCoordP3uiv(GLenum type_, GLuint* coords) {
// 	(*ptrglTexCoordP3uiv)(type_, coords);
// }
// void goglTexCoordP4ui(GLenum type_, GLuint coords) {
// 	(*ptrglTexCoordP4ui)(type_, coords);
// }
// void goglTexCoordP4uiv(GLenum type_, GLuint* coords) {
// 	(*ptrglTexCoordP4uiv)(type_, coords);
// }
// void goglMultiTexCoordP1ui(GLenum texture, GLenum type_, GLuint coords) {
// 	(*ptrglMultiTexCoordP1ui)(texture, type_, coords);
// }
// void goglMultiTexCoordP1uiv(GLenum texture, GLenum type_, GLuint* coords) {
// 	(*ptrglMultiTexCoordP1uiv)(texture, type_, coords);
// }
// void goglMultiTexCoordP2ui(GLenum texture, GLenum type_, GLuint coords) {
// 	(*ptrglMultiTexCoordP2ui)(texture, type_, coords);
// }
// void goglMultiTexCoordP2uiv(GLenum texture, GLenum type_, GLuint* coords) {
// 	(*ptrglMultiTexCoordP2uiv)(texture, type_, coords);
// }
// void goglMultiTexCoordP3ui(GLenum texture, GLenum type_, GLuint coords) {
// 	(*ptrglMultiTexCoordP3ui)(texture, type_, coords);
// }
// void goglMultiTexCoordP3uiv(GLenum texture, GLenum type_, GLuint* coords) {
// 	(*ptrglMultiTexCoordP3uiv)(texture, type_, coords);
// }
// void goglMultiTexCoordP4ui(GLenum texture, GLenum type_, GLuint coords) {
// 	(*ptrglMultiTexCoordP4ui)(texture, type_, coords);
// }
// void goglMultiTexCoordP4uiv(GLenum texture, GLenum type_, GLuint* coords) {
// 	(*ptrglMultiTexCoordP4uiv)(texture, type_, coords);
// }
// void goglNormalP3ui(GLenum type_, GLuint coords) {
// 	(*ptrglNormalP3ui)(type_, coords);
// }
// void goglNormalP3uiv(GLenum type_, GLuint* coords) {
// 	(*ptrglNormalP3uiv)(type_, coords);
// }
// void goglColorP3ui(GLenum type_, GLuint color) {
// 	(*ptrglColorP3ui)(type_, color);
// }
// void goglColorP3uiv(GLenum type_, GLuint* color) {
// 	(*ptrglColorP3uiv)(type_, color);
// }
// void goglColorP4ui(GLenum type_, GLuint color) {
// 	(*ptrglColorP4ui)(type_, color);
// }
// void goglColorP4uiv(GLenum type_, GLuint* color) {
// 	(*ptrglColorP4uiv)(type_, color);
// }
// void goglSecondaryColorP3ui(GLenum type_, GLuint color) {
// 	(*ptrglSecondaryColorP3ui)(type_, color);
// }
// void goglSecondaryColorP3uiv(GLenum type_, GLuint* color) {
// 	(*ptrglSecondaryColorP3uiv)(type_, color);
// }
// void goglVertexAttribP1ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {
// 	(*ptrglVertexAttribP1ui)(index, type_, normalized, value);
// }
// void goglVertexAttribP1uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {
// 	(*ptrglVertexAttribP1uiv)(index, type_, normalized, value);
// }
// void goglVertexAttribP2ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {
// 	(*ptrglVertexAttribP2ui)(index, type_, normalized, value);
// }
// void goglVertexAttribP2uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {
// 	(*ptrglVertexAttribP2uiv)(index, type_, normalized, value);
// }
// void goglVertexAttribP3ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {
// 	(*ptrglVertexAttribP3ui)(index, type_, normalized, value);
// }
// void goglVertexAttribP3uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {
// 	(*ptrglVertexAttribP3uiv)(index, type_, normalized, value);
// }
// void goglVertexAttribP4ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {
// 	(*ptrglVertexAttribP4ui)(index, type_, normalized, value);
// }
// void goglVertexAttribP4uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {
// 	(*ptrglVertexAttribP4uiv)(index, type_, normalized, value);
// }
// //  VERSION_4_0
// void goglMinSampleShading(GLfloat value) {
// 	(*ptrglMinSampleShading)(value);
// }
// void goglBlendEquationi(GLuint buf, GLenum mode) {
// 	(*ptrglBlendEquationi)(buf, mode);
// }
// void goglBlendEquationSeparatei(GLuint buf, GLenum modeRGB, GLenum modeAlpha) {
// 	(*ptrglBlendEquationSeparatei)(buf, modeRGB, modeAlpha);
// }
// void goglBlendFunci(GLuint buf, GLenum src, GLenum dst) {
// 	(*ptrglBlendFunci)(buf, src, dst);
// }
// void goglBlendFuncSeparatei(GLuint buf, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
// 	(*ptrglBlendFuncSeparatei)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
// }
// void goglDrawArraysIndirect(GLenum mode, GLvoid* indirect) {
// 	(*ptrglDrawArraysIndirect)(mode, indirect);
// }
// void goglDrawElementsIndirect(GLenum mode, GLenum type_, GLvoid* indirect) {
// 	(*ptrglDrawElementsIndirect)(mode, type_, indirect);
// }
// void goglUniform1d(GLint location, GLdouble x) {
// 	(*ptrglUniform1d)(location, x);
// }
// void goglUniform2d(GLint location, GLdouble x, GLdouble y) {
// 	(*ptrglUniform2d)(location, x, y);
// }
// void goglUniform3d(GLint location, GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglUniform3d)(location, x, y, z);
// }
// void goglUniform4d(GLint location, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglUniform4d)(location, x, y, z, w);
// }
// void goglUniform1dv(GLint location, GLsizei count, GLdouble* value) {
// 	(*ptrglUniform1dv)(location, count, value);
// }
// void goglUniform2dv(GLint location, GLsizei count, GLdouble* value) {
// 	(*ptrglUniform2dv)(location, count, value);
// }
// void goglUniform3dv(GLint location, GLsizei count, GLdouble* value) {
// 	(*ptrglUniform3dv)(location, count, value);
// }
// void goglUniform4dv(GLint location, GLsizei count, GLdouble* value) {
// 	(*ptrglUniform4dv)(location, count, value);
// }
// void goglUniformMatrix2dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglUniformMatrix2dv)(location, count, transpose, value);
// }
// void goglUniformMatrix3dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglUniformMatrix3dv)(location, count, transpose, value);
// }
// void goglUniformMatrix4dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglUniformMatrix4dv)(location, count, transpose, value);
// }
// void goglUniformMatrix2x3dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglUniformMatrix2x3dv)(location, count, transpose, value);
// }
// void goglUniformMatrix2x4dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglUniformMatrix2x4dv)(location, count, transpose, value);
// }
// void goglUniformMatrix3x2dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglUniformMatrix3x2dv)(location, count, transpose, value);
// }
// void goglUniformMatrix3x4dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglUniformMatrix3x4dv)(location, count, transpose, value);
// }
// void goglUniformMatrix4x2dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglUniformMatrix4x2dv)(location, count, transpose, value);
// }
// void goglUniformMatrix4x3dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglUniformMatrix4x3dv)(location, count, transpose, value);
// }
// void goglGetUniformdv(GLuint program, GLint location, GLdouble* params) {
// 	(*ptrglGetUniformdv)(program, location, params);
// }
// GLint goglGetSubroutineUniformLocation(GLuint program, GLenum shadertype, GLchar* name) {
// 	return (*ptrglGetSubroutineUniformLocation)(program, shadertype, name);
// }
// GLuint goglGetSubroutineIndex(GLuint program, GLenum shadertype, GLchar* name) {
// 	return (*ptrglGetSubroutineIndex)(program, shadertype, name);
// }
// void goglGetActiveSubroutineUniformiv(GLuint program, GLenum shadertype, GLuint index, GLenum pname, GLint* values) {
// 	(*ptrglGetActiveSubroutineUniformiv)(program, shadertype, index, pname, values);
// }
// void goglGetActiveSubroutineUniformName(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name) {
// 	(*ptrglGetActiveSubroutineUniformName)(program, shadertype, index, bufsize, length, name);
// }
// void goglGetActiveSubroutineName(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name) {
// 	(*ptrglGetActiveSubroutineName)(program, shadertype, index, bufsize, length, name);
// }
// void goglUniformSubroutinesuiv(GLenum shadertype, GLsizei count, GLuint* indices) {
// 	(*ptrglUniformSubroutinesuiv)(shadertype, count, indices);
// }
// void goglGetUniformSubroutineuiv(GLenum shadertype, GLint location, GLuint* params) {
// 	(*ptrglGetUniformSubroutineuiv)(shadertype, location, params);
// }
// void goglGetProgramStageiv(GLuint program, GLenum shadertype, GLenum pname, GLint* values) {
// 	(*ptrglGetProgramStageiv)(program, shadertype, pname, values);
// }
// void goglPatchParameteri(GLenum pname, GLint value) {
// 	(*ptrglPatchParameteri)(pname, value);
// }
// void goglPatchParameterfv(GLenum pname, GLfloat* values) {
// 	(*ptrglPatchParameterfv)(pname, values);
// }
// void goglBindTransformFeedback(GLenum target, GLuint id) {
// 	(*ptrglBindTransformFeedback)(target, id);
// }
// void goglDeleteTransformFeedbacks(GLsizei n, GLuint* ids) {
// 	(*ptrglDeleteTransformFeedbacks)(n, ids);
// }
// void goglGenTransformFeedbacks(GLsizei n, GLuint* ids) {
// 	(*ptrglGenTransformFeedbacks)(n, ids);
// }
// GLboolean goglIsTransformFeedback(GLuint id) {
// 	return (*ptrglIsTransformFeedback)(id);
// }
// void goglPauseTransformFeedback() {
// 	(*ptrglPauseTransformFeedback)();
// }
// void goglResumeTransformFeedback() {
// 	(*ptrglResumeTransformFeedback)();
// }
// void goglDrawTransformFeedback(GLenum mode, GLuint id) {
// 	(*ptrglDrawTransformFeedback)(mode, id);
// }
// void goglDrawTransformFeedbackStream(GLenum mode, GLuint id, GLuint stream) {
// 	(*ptrglDrawTransformFeedbackStream)(mode, id, stream);
// }
// void goglBeginQueryIndexed(GLenum target, GLuint index, GLuint id) {
// 	(*ptrglBeginQueryIndexed)(target, index, id);
// }
// void goglEndQueryIndexed(GLenum target, GLuint index) {
// 	(*ptrglEndQueryIndexed)(target, index);
// }
// void goglGetQueryIndexediv(GLenum target, GLuint index, GLenum pname, GLint* params) {
// 	(*ptrglGetQueryIndexediv)(target, index, pname, params);
// }
// //  VERSION_4_1
// void goglReleaseShaderCompiler() {
// 	(*ptrglReleaseShaderCompiler)();
// }
// void goglShaderBinary(GLsizei count, GLuint* shaders, GLenum binaryformat, GLvoid* binary, GLsizei length) {
// 	(*ptrglShaderBinary)(count, shaders, binaryformat, binary, length);
// }
// void goglGetShaderPrecisionFormat(GLenum shadertype, GLenum precisiontype, GLint* range_, GLint* precision) {
// 	(*ptrglGetShaderPrecisionFormat)(shadertype, precisiontype, range_, precision);
// }
// void goglDepthRangef(GLfloat n, GLfloat f) {
// 	(*ptrglDepthRangef)(n, f);
// }
// void goglClearDepthf(GLfloat d) {
// 	(*ptrglClearDepthf)(d);
// }
// void goglGetProgramBinary(GLuint program, GLsizei bufSize, GLsizei* length, GLenum* binaryFormat, GLvoid* binary) {
// 	(*ptrglGetProgramBinary)(program, bufSize, length, binaryFormat, binary);
// }
// void goglProgramBinary(GLuint program, GLenum binaryFormat, GLvoid* binary, GLsizei length) {
// 	(*ptrglProgramBinary)(program, binaryFormat, binary, length);
// }
// void goglProgramParameteri(GLuint program, GLenum pname, GLint value) {
// 	(*ptrglProgramParameteri)(program, pname, value);
// }
// void goglUseProgramStages(GLuint pipeline, GLbitfield stages, GLuint program) {
// 	(*ptrglUseProgramStages)(pipeline, stages, program);
// }
// void goglActiveShaderProgram(GLuint pipeline, GLuint program) {
// 	(*ptrglActiveShaderProgram)(pipeline, program);
// }
// GLuint goglCreateShaderProgramv(GLenum type_, GLsizei count, GLchar* const* strings) {
// 	return (*ptrglCreateShaderProgramv)(type_, count, strings);
// }
// void goglBindProgramPipeline(GLuint pipeline) {
// 	(*ptrglBindProgramPipeline)(pipeline);
// }
// void goglDeleteProgramPipelines(GLsizei n, GLuint* pipelines) {
// 	(*ptrglDeleteProgramPipelines)(n, pipelines);
// }
// void goglGenProgramPipelines(GLsizei n, GLuint* pipelines) {
// 	(*ptrglGenProgramPipelines)(n, pipelines);
// }
// GLboolean goglIsProgramPipeline(GLuint pipeline) {
// 	return (*ptrglIsProgramPipeline)(pipeline);
// }
// void goglGetProgramPipelineiv(GLuint pipeline, GLenum pname, GLint* params) {
// 	(*ptrglGetProgramPipelineiv)(pipeline, pname, params);
// }
// void goglProgramUniform1i(GLuint program, GLint location, GLint v0) {
// 	(*ptrglProgramUniform1i)(program, location, v0);
// }
// void goglProgramUniform1iv(GLuint program, GLint location, GLsizei count, GLint* value) {
// 	(*ptrglProgramUniform1iv)(program, location, count, value);
// }
// void goglProgramUniform1f(GLuint program, GLint location, GLfloat v0) {
// 	(*ptrglProgramUniform1f)(program, location, v0);
// }
// void goglProgramUniform1fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrglProgramUniform1fv)(program, location, count, value);
// }
// void goglProgramUniform1d(GLuint program, GLint location, GLdouble v0) {
// 	(*ptrglProgramUniform1d)(program, location, v0);
// }
// void goglProgramUniform1dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {
// 	(*ptrglProgramUniform1dv)(program, location, count, value);
// }
// void goglProgramUniform1ui(GLuint program, GLint location, GLuint v0) {
// 	(*ptrglProgramUniform1ui)(program, location, v0);
// }
// void goglProgramUniform1uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {
// 	(*ptrglProgramUniform1uiv)(program, location, count, value);
// }
// void goglProgramUniform2i(GLuint program, GLint location, GLint v0, GLint v1) {
// 	(*ptrglProgramUniform2i)(program, location, v0, v1);
// }
// void goglProgramUniform2iv(GLuint program, GLint location, GLsizei count, GLint* value) {
// 	(*ptrglProgramUniform2iv)(program, location, count, value);
// }
// void goglProgramUniform2f(GLuint program, GLint location, GLfloat v0, GLfloat v1) {
// 	(*ptrglProgramUniform2f)(program, location, v0, v1);
// }
// void goglProgramUniform2fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrglProgramUniform2fv)(program, location, count, value);
// }
// void goglProgramUniform2d(GLuint program, GLint location, GLdouble v0, GLdouble v1) {
// 	(*ptrglProgramUniform2d)(program, location, v0, v1);
// }
// void goglProgramUniform2dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {
// 	(*ptrglProgramUniform2dv)(program, location, count, value);
// }
// void goglProgramUniform2ui(GLuint program, GLint location, GLuint v0, GLuint v1) {
// 	(*ptrglProgramUniform2ui)(program, location, v0, v1);
// }
// void goglProgramUniform2uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {
// 	(*ptrglProgramUniform2uiv)(program, location, count, value);
// }
// void goglProgramUniform3i(GLuint program, GLint location, GLint v0, GLint v1, GLint v2) {
// 	(*ptrglProgramUniform3i)(program, location, v0, v1, v2);
// }
// void goglProgramUniform3iv(GLuint program, GLint location, GLsizei count, GLint* value) {
// 	(*ptrglProgramUniform3iv)(program, location, count, value);
// }
// void goglProgramUniform3f(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
// 	(*ptrglProgramUniform3f)(program, location, v0, v1, v2);
// }
// void goglProgramUniform3fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrglProgramUniform3fv)(program, location, count, value);
// }
// void goglProgramUniform3d(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2) {
// 	(*ptrglProgramUniform3d)(program, location, v0, v1, v2);
// }
// void goglProgramUniform3dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {
// 	(*ptrglProgramUniform3dv)(program, location, count, value);
// }
// void goglProgramUniform3ui(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2) {
// 	(*ptrglProgramUniform3ui)(program, location, v0, v1, v2);
// }
// void goglProgramUniform3uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {
// 	(*ptrglProgramUniform3uiv)(program, location, count, value);
// }
// void goglProgramUniform4i(GLuint program, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
// 	(*ptrglProgramUniform4i)(program, location, v0, v1, v2, v3);
// }
// void goglProgramUniform4iv(GLuint program, GLint location, GLsizei count, GLint* value) {
// 	(*ptrglProgramUniform4iv)(program, location, count, value);
// }
// void goglProgramUniform4f(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
// 	(*ptrglProgramUniform4f)(program, location, v0, v1, v2, v3);
// }
// void goglProgramUniform4fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrglProgramUniform4fv)(program, location, count, value);
// }
// void goglProgramUniform4d(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {
// 	(*ptrglProgramUniform4d)(program, location, v0, v1, v2, v3);
// }
// void goglProgramUniform4dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {
// 	(*ptrglProgramUniform4dv)(program, location, count, value);
// }
// void goglProgramUniform4ui(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3) {
// 	(*ptrglProgramUniform4ui)(program, location, v0, v1, v2, v3);
// }
// void goglProgramUniform4uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {
// 	(*ptrglProgramUniform4uiv)(program, location, count, value);
// }
// void goglProgramUniformMatrix2fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglProgramUniformMatrix2fv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix3fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglProgramUniformMatrix3fv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix4fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglProgramUniformMatrix4fv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix2dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglProgramUniformMatrix2dv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix3dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglProgramUniformMatrix3dv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix4dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglProgramUniformMatrix4dv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix2x3fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglProgramUniformMatrix2x3fv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix3x2fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglProgramUniformMatrix3x2fv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix2x4fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglProgramUniformMatrix2x4fv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix4x2fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglProgramUniformMatrix4x2fv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix3x4fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglProgramUniformMatrix3x4fv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix4x3fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrglProgramUniformMatrix4x3fv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix2x3dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglProgramUniformMatrix2x3dv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix3x2dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglProgramUniformMatrix3x2dv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix2x4dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglProgramUniformMatrix2x4dv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix4x2dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglProgramUniformMatrix4x2dv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix3x4dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglProgramUniformMatrix3x4dv)(program, location, count, transpose, value);
// }
// void goglProgramUniformMatrix4x3dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {
// 	(*ptrglProgramUniformMatrix4x3dv)(program, location, count, transpose, value);
// }
// void goglValidateProgramPipeline(GLuint pipeline) {
// 	(*ptrglValidateProgramPipeline)(pipeline);
// }
// void goglGetProgramPipelineInfoLog(GLuint pipeline, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {
// 	(*ptrglGetProgramPipelineInfoLog)(pipeline, bufSize, length, infoLog);
// }
// void goglVertexAttribL1d(GLuint index, GLdouble x) {
// 	(*ptrglVertexAttribL1d)(index, x);
// }
// void goglVertexAttribL2d(GLuint index, GLdouble x, GLdouble y) {
// 	(*ptrglVertexAttribL2d)(index, x, y);
// }
// void goglVertexAttribL3d(GLuint index, GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglVertexAttribL3d)(index, x, y, z);
// }
// void goglVertexAttribL4d(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglVertexAttribL4d)(index, x, y, z, w);
// }
// void goglVertexAttribL1dv(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttribL1dv)(index, v);
// }
// void goglVertexAttribL2dv(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttribL2dv)(index, v);
// }
// void goglVertexAttribL3dv(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttribL3dv)(index, v);
// }
// void goglVertexAttribL4dv(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttribL4dv)(index, v);
// }
// void goglVertexAttribLPointer(GLuint index, GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglVertexAttribLPointer)(index, size, type_, stride, pointer);
// }
// void goglGetVertexAttribLdv(GLuint index, GLenum pname, GLdouble* params) {
// 	(*ptrglGetVertexAttribLdv)(index, pname, params);
// }
// void goglViewportArrayv(GLuint first, GLsizei count, GLfloat* v) {
// 	(*ptrglViewportArrayv)(first, count, v);
// }
// void goglViewportIndexedf(GLuint index, GLfloat x, GLfloat y, GLfloat w, GLfloat h) {
// 	(*ptrglViewportIndexedf)(index, x, y, w, h);
// }
// void goglViewportIndexedfv(GLuint index, GLfloat* v) {
// 	(*ptrglViewportIndexedfv)(index, v);
// }
// void goglScissorArrayv(GLuint first, GLsizei count, GLint* v) {
// 	(*ptrglScissorArrayv)(first, count, v);
// }
// void goglScissorIndexed(GLuint index, GLint left, GLint bottom, GLsizei width, GLsizei height) {
// 	(*ptrglScissorIndexed)(index, left, bottom, width, height);
// }
// void goglScissorIndexedv(GLuint index, GLint* v) {
// 	(*ptrglScissorIndexedv)(index, v);
// }
// void goglDepthRangeArrayv(GLuint first, GLsizei count, GLdouble* v) {
// 	(*ptrglDepthRangeArrayv)(first, count, v);
// }
// void goglDepthRangeIndexed(GLuint index, GLdouble n, GLdouble f) {
// 	(*ptrglDepthRangeIndexed)(index, n, f);
// }
// void goglGetFloati_v(GLenum target, GLuint index, GLfloat* data) {
// 	(*ptrglGetFloati_v)(target, index, data);
// }
// void goglGetDoublei_v(GLenum target, GLuint index, GLdouble* data) {
// 	(*ptrglGetDoublei_v)(target, index, data);
// }
// //  VERSION_4_2
// void goglDrawArraysInstancedBaseInstance(GLenum mode, GLint first, GLsizei count, GLsizei primcount, GLuint baseinstance) {
// 	(*ptrglDrawArraysInstancedBaseInstance)(mode, first, count, primcount, baseinstance);
// }
// void goglDrawElementsInstancedBaseInstance(GLenum mode, GLsizei count, GLenum type_, void* indices, GLsizei primcount, GLuint baseinstance) {
// 	(*ptrglDrawElementsInstancedBaseInstance)(mode, count, type_, indices, primcount, baseinstance);
// }
// void goglDrawElementsInstancedBaseVertexBaseInstance(GLenum mode, GLsizei count, GLenum type_, void* indices, GLsizei primcount, GLint basevertex, GLuint baseinstance) {
// 	(*ptrglDrawElementsInstancedBaseVertexBaseInstance)(mode, count, type_, indices, primcount, basevertex, baseinstance);
// }
// void goglDrawTransformFeedbackInstanced(GLenum mode, GLuint id, GLsizei primcount) {
// 	(*ptrglDrawTransformFeedbackInstanced)(mode, id, primcount);
// }
// void goglDrawTransformFeedbackStreamInstanced(GLenum mode, GLuint id, GLuint stream, GLsizei primcount) {
// 	(*ptrglDrawTransformFeedbackStreamInstanced)(mode, id, stream, primcount);
// }
// void goglGetInternalformativ(GLenum target, GLenum internalformat, GLenum pname, GLsizei bufSize, GLint* params) {
// 	(*ptrglGetInternalformativ)(target, internalformat, pname, bufSize, params);
// }
// void goglGetActiveAtomicCounterBufferiv(GLuint program, GLuint bufferIndex, GLenum pname, GLint* params) {
// 	(*ptrglGetActiveAtomicCounterBufferiv)(program, bufferIndex, pname, params);
// }
// void goglBindImageTexture(GLuint unit, GLuint texture, GLint level, GLboolean layered, GLint layer, GLenum access, GLenum format) {
// 	(*ptrglBindImageTexture)(unit, texture, level, layered, layer, access, format);
// }
// void goglMemoryBarrier(GLbitfield barriers) {
// 	(*ptrglMemoryBarrier)(barriers);
// }
// void goglTexStorage1D(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width) {
// 	(*ptrglTexStorage1D)(target, levels, internalformat, width);
// }
// void goglTexStorage2D(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height) {
// 	(*ptrglTexStorage2D)(target, levels, internalformat, width, height);
// }
// void goglTexStorage3D(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth) {
// 	(*ptrglTexStorage3D)(target, levels, internalformat, width, height, depth);
// }
// void goglTextureStorage1DEXT(GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width) {
// 	(*ptrglTextureStorage1DEXT)(texture, target, levels, internalformat, width);
// }
// void goglTextureStorage2DEXT(GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height) {
// 	(*ptrglTextureStorage2DEXT)(texture, target, levels, internalformat, width, height);
// }
// void goglTextureStorage3DEXT(GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth) {
// 	(*ptrglTextureStorage3DEXT)(texture, target, levels, internalformat, width, height, depth);
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
// int init_VERSION_3_0() {
// 	ptrglColorMaski = goglGetProcAddress("glColorMaski");
// 	if(ptrglColorMaski == NULL) return 1;
// 	ptrglGetBooleani_v = goglGetProcAddress("glGetBooleani_v");
// 	if(ptrglGetBooleani_v == NULL) return 1;
// 	ptrglGetIntegeri_v = goglGetProcAddress("glGetIntegeri_v");
// 	if(ptrglGetIntegeri_v == NULL) return 1;
// 	ptrglEnablei = goglGetProcAddress("glEnablei");
// 	if(ptrglEnablei == NULL) return 1;
// 	ptrglDisablei = goglGetProcAddress("glDisablei");
// 	if(ptrglDisablei == NULL) return 1;
// 	ptrglIsEnabledi = goglGetProcAddress("glIsEnabledi");
// 	if(ptrglIsEnabledi == NULL) return 1;
// 	ptrglBeginTransformFeedback = goglGetProcAddress("glBeginTransformFeedback");
// 	if(ptrglBeginTransformFeedback == NULL) return 1;
// 	ptrglEndTransformFeedback = goglGetProcAddress("glEndTransformFeedback");
// 	if(ptrglEndTransformFeedback == NULL) return 1;
// 	ptrglBindBufferRange = goglGetProcAddress("glBindBufferRange");
// 	if(ptrglBindBufferRange == NULL) return 1;
// 	ptrglBindBufferBase = goglGetProcAddress("glBindBufferBase");
// 	if(ptrglBindBufferBase == NULL) return 1;
// 	ptrglTransformFeedbackVaryings = goglGetProcAddress("glTransformFeedbackVaryings");
// 	if(ptrglTransformFeedbackVaryings == NULL) return 1;
// 	ptrglGetTransformFeedbackVarying = goglGetProcAddress("glGetTransformFeedbackVarying");
// 	if(ptrglGetTransformFeedbackVarying == NULL) return 1;
// 	ptrglClampColor = goglGetProcAddress("glClampColor");
// 	if(ptrglClampColor == NULL) return 1;
// 	ptrglBeginConditionalRender = goglGetProcAddress("glBeginConditionalRender");
// 	if(ptrglBeginConditionalRender == NULL) return 1;
// 	ptrglEndConditionalRender = goglGetProcAddress("glEndConditionalRender");
// 	if(ptrglEndConditionalRender == NULL) return 1;
// 	ptrglVertexAttribIPointer = goglGetProcAddress("glVertexAttribIPointer");
// 	if(ptrglVertexAttribIPointer == NULL) return 1;
// 	ptrglGetVertexAttribIiv = goglGetProcAddress("glGetVertexAttribIiv");
// 	if(ptrglGetVertexAttribIiv == NULL) return 1;
// 	ptrglGetVertexAttribIuiv = goglGetProcAddress("glGetVertexAttribIuiv");
// 	if(ptrglGetVertexAttribIuiv == NULL) return 1;
// 	ptrglVertexAttribI1i = goglGetProcAddress("glVertexAttribI1i");
// 	if(ptrglVertexAttribI1i == NULL) return 1;
// 	ptrglVertexAttribI2i = goglGetProcAddress("glVertexAttribI2i");
// 	if(ptrglVertexAttribI2i == NULL) return 1;
// 	ptrglVertexAttribI3i = goglGetProcAddress("glVertexAttribI3i");
// 	if(ptrglVertexAttribI3i == NULL) return 1;
// 	ptrglVertexAttribI4i = goglGetProcAddress("glVertexAttribI4i");
// 	if(ptrglVertexAttribI4i == NULL) return 1;
// 	ptrglVertexAttribI1ui = goglGetProcAddress("glVertexAttribI1ui");
// 	if(ptrglVertexAttribI1ui == NULL) return 1;
// 	ptrglVertexAttribI2ui = goglGetProcAddress("glVertexAttribI2ui");
// 	if(ptrglVertexAttribI2ui == NULL) return 1;
// 	ptrglVertexAttribI3ui = goglGetProcAddress("glVertexAttribI3ui");
// 	if(ptrglVertexAttribI3ui == NULL) return 1;
// 	ptrglVertexAttribI4ui = goglGetProcAddress("glVertexAttribI4ui");
// 	if(ptrglVertexAttribI4ui == NULL) return 1;
// 	ptrglVertexAttribI1iv = goglGetProcAddress("glVertexAttribI1iv");
// 	if(ptrglVertexAttribI1iv == NULL) return 1;
// 	ptrglVertexAttribI2iv = goglGetProcAddress("glVertexAttribI2iv");
// 	if(ptrglVertexAttribI2iv == NULL) return 1;
// 	ptrglVertexAttribI3iv = goglGetProcAddress("glVertexAttribI3iv");
// 	if(ptrglVertexAttribI3iv == NULL) return 1;
// 	ptrglVertexAttribI4iv = goglGetProcAddress("glVertexAttribI4iv");
// 	if(ptrglVertexAttribI4iv == NULL) return 1;
// 	ptrglVertexAttribI1uiv = goglGetProcAddress("glVertexAttribI1uiv");
// 	if(ptrglVertexAttribI1uiv == NULL) return 1;
// 	ptrglVertexAttribI2uiv = goglGetProcAddress("glVertexAttribI2uiv");
// 	if(ptrglVertexAttribI2uiv == NULL) return 1;
// 	ptrglVertexAttribI3uiv = goglGetProcAddress("glVertexAttribI3uiv");
// 	if(ptrglVertexAttribI3uiv == NULL) return 1;
// 	ptrglVertexAttribI4uiv = goglGetProcAddress("glVertexAttribI4uiv");
// 	if(ptrglVertexAttribI4uiv == NULL) return 1;
// 	ptrglVertexAttribI4bv = goglGetProcAddress("glVertexAttribI4bv");
// 	if(ptrglVertexAttribI4bv == NULL) return 1;
// 	ptrglVertexAttribI4sv = goglGetProcAddress("glVertexAttribI4sv");
// 	if(ptrglVertexAttribI4sv == NULL) return 1;
// 	ptrglVertexAttribI4ubv = goglGetProcAddress("glVertexAttribI4ubv");
// 	if(ptrglVertexAttribI4ubv == NULL) return 1;
// 	ptrglVertexAttribI4usv = goglGetProcAddress("glVertexAttribI4usv");
// 	if(ptrglVertexAttribI4usv == NULL) return 1;
// 	ptrglGetUniformuiv = goglGetProcAddress("glGetUniformuiv");
// 	if(ptrglGetUniformuiv == NULL) return 1;
// 	ptrglBindFragDataLocation = goglGetProcAddress("glBindFragDataLocation");
// 	if(ptrglBindFragDataLocation == NULL) return 1;
// 	ptrglGetFragDataLocation = goglGetProcAddress("glGetFragDataLocation");
// 	if(ptrglGetFragDataLocation == NULL) return 1;
// 	ptrglUniform1ui = goglGetProcAddress("glUniform1ui");
// 	if(ptrglUniform1ui == NULL) return 1;
// 	ptrglUniform2ui = goglGetProcAddress("glUniform2ui");
// 	if(ptrglUniform2ui == NULL) return 1;
// 	ptrglUniform3ui = goglGetProcAddress("glUniform3ui");
// 	if(ptrglUniform3ui == NULL) return 1;
// 	ptrglUniform4ui = goglGetProcAddress("glUniform4ui");
// 	if(ptrglUniform4ui == NULL) return 1;
// 	ptrglUniform1uiv = goglGetProcAddress("glUniform1uiv");
// 	if(ptrglUniform1uiv == NULL) return 1;
// 	ptrglUniform2uiv = goglGetProcAddress("glUniform2uiv");
// 	if(ptrglUniform2uiv == NULL) return 1;
// 	ptrglUniform3uiv = goglGetProcAddress("glUniform3uiv");
// 	if(ptrglUniform3uiv == NULL) return 1;
// 	ptrglUniform4uiv = goglGetProcAddress("glUniform4uiv");
// 	if(ptrglUniform4uiv == NULL) return 1;
// 	ptrglTexParameterIiv = goglGetProcAddress("glTexParameterIiv");
// 	if(ptrglTexParameterIiv == NULL) return 1;
// 	ptrglTexParameterIuiv = goglGetProcAddress("glTexParameterIuiv");
// 	if(ptrglTexParameterIuiv == NULL) return 1;
// 	ptrglGetTexParameterIiv = goglGetProcAddress("glGetTexParameterIiv");
// 	if(ptrglGetTexParameterIiv == NULL) return 1;
// 	ptrglGetTexParameterIuiv = goglGetProcAddress("glGetTexParameterIuiv");
// 	if(ptrglGetTexParameterIuiv == NULL) return 1;
// 	ptrglClearBufferiv = goglGetProcAddress("glClearBufferiv");
// 	if(ptrglClearBufferiv == NULL) return 1;
// 	ptrglClearBufferuiv = goglGetProcAddress("glClearBufferuiv");
// 	if(ptrglClearBufferuiv == NULL) return 1;
// 	ptrglClearBufferfv = goglGetProcAddress("glClearBufferfv");
// 	if(ptrglClearBufferfv == NULL) return 1;
// 	ptrglClearBufferfi = goglGetProcAddress("glClearBufferfi");
// 	if(ptrglClearBufferfi == NULL) return 1;
// 	ptrglGetStringi = goglGetProcAddress("glGetStringi");
// 	if(ptrglGetStringi == NULL) return 1;
// 	ptrglIsRenderbuffer = goglGetProcAddress("glIsRenderbuffer");
// 	if(ptrglIsRenderbuffer == NULL) return 1;
// 	ptrglBindRenderbuffer = goglGetProcAddress("glBindRenderbuffer");
// 	if(ptrglBindRenderbuffer == NULL) return 1;
// 	ptrglDeleteRenderbuffers = goglGetProcAddress("glDeleteRenderbuffers");
// 	if(ptrglDeleteRenderbuffers == NULL) return 1;
// 	ptrglGenRenderbuffers = goglGetProcAddress("glGenRenderbuffers");
// 	if(ptrglGenRenderbuffers == NULL) return 1;
// 	ptrglRenderbufferStorage = goglGetProcAddress("glRenderbufferStorage");
// 	if(ptrglRenderbufferStorage == NULL) return 1;
// 	ptrglGetRenderbufferParameteriv = goglGetProcAddress("glGetRenderbufferParameteriv");
// 	if(ptrglGetRenderbufferParameteriv == NULL) return 1;
// 	ptrglIsFramebuffer = goglGetProcAddress("glIsFramebuffer");
// 	if(ptrglIsFramebuffer == NULL) return 1;
// 	ptrglBindFramebuffer = goglGetProcAddress("glBindFramebuffer");
// 	if(ptrglBindFramebuffer == NULL) return 1;
// 	ptrglDeleteFramebuffers = goglGetProcAddress("glDeleteFramebuffers");
// 	if(ptrglDeleteFramebuffers == NULL) return 1;
// 	ptrglGenFramebuffers = goglGetProcAddress("glGenFramebuffers");
// 	if(ptrglGenFramebuffers == NULL) return 1;
// 	ptrglCheckFramebufferStatus = goglGetProcAddress("glCheckFramebufferStatus");
// 	if(ptrglCheckFramebufferStatus == NULL) return 1;
// 	ptrglFramebufferTexture1D = goglGetProcAddress("glFramebufferTexture1D");
// 	if(ptrglFramebufferTexture1D == NULL) return 1;
// 	ptrglFramebufferTexture2D = goglGetProcAddress("glFramebufferTexture2D");
// 	if(ptrglFramebufferTexture2D == NULL) return 1;
// 	ptrglFramebufferTexture3D = goglGetProcAddress("glFramebufferTexture3D");
// 	if(ptrglFramebufferTexture3D == NULL) return 1;
// 	ptrglFramebufferRenderbuffer = goglGetProcAddress("glFramebufferRenderbuffer");
// 	if(ptrglFramebufferRenderbuffer == NULL) return 1;
// 	ptrglGetFramebufferAttachmentParameteriv = goglGetProcAddress("glGetFramebufferAttachmentParameteriv");
// 	if(ptrglGetFramebufferAttachmentParameteriv == NULL) return 1;
// 	ptrglGenerateMipmap = goglGetProcAddress("glGenerateMipmap");
// 	if(ptrglGenerateMipmap == NULL) return 1;
// 	ptrglBlitFramebuffer = goglGetProcAddress("glBlitFramebuffer");
// 	if(ptrglBlitFramebuffer == NULL) return 1;
// 	ptrglRenderbufferStorageMultisample = goglGetProcAddress("glRenderbufferStorageMultisample");
// 	if(ptrglRenderbufferStorageMultisample == NULL) return 1;
// 	ptrglFramebufferTextureLayer = goglGetProcAddress("glFramebufferTextureLayer");
// 	if(ptrglFramebufferTextureLayer == NULL) return 1;
// 	ptrglMapBufferRange = goglGetProcAddress("glMapBufferRange");
// 	if(ptrglMapBufferRange == NULL) return 1;
// 	ptrglFlushMappedBufferRange = goglGetProcAddress("glFlushMappedBufferRange");
// 	if(ptrglFlushMappedBufferRange == NULL) return 1;
// 	ptrglBindVertexArray = goglGetProcAddress("glBindVertexArray");
// 	if(ptrglBindVertexArray == NULL) return 1;
// 	ptrglDeleteVertexArrays = goglGetProcAddress("glDeleteVertexArrays");
// 	if(ptrglDeleteVertexArrays == NULL) return 1;
// 	ptrglGenVertexArrays = goglGetProcAddress("glGenVertexArrays");
// 	if(ptrglGenVertexArrays == NULL) return 1;
// 	ptrglIsVertexArray = goglGetProcAddress("glIsVertexArray");
// 	if(ptrglIsVertexArray == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_3_1() {
// 	ptrglDrawArraysInstanced = goglGetProcAddress("glDrawArraysInstanced");
// 	if(ptrglDrawArraysInstanced == NULL) return 1;
// 	ptrglDrawElementsInstanced = goglGetProcAddress("glDrawElementsInstanced");
// 	if(ptrglDrawElementsInstanced == NULL) return 1;
// 	ptrglTexBuffer = goglGetProcAddress("glTexBuffer");
// 	if(ptrglTexBuffer == NULL) return 1;
// 	ptrglPrimitiveRestartIndex = goglGetProcAddress("glPrimitiveRestartIndex");
// 	if(ptrglPrimitiveRestartIndex == NULL) return 1;
// 	ptrglCopyBufferSubData = goglGetProcAddress("glCopyBufferSubData");
// 	if(ptrglCopyBufferSubData == NULL) return 1;
// 	ptrglGetUniformIndices = goglGetProcAddress("glGetUniformIndices");
// 	if(ptrglGetUniformIndices == NULL) return 1;
// 	ptrglGetActiveUniformsiv = goglGetProcAddress("glGetActiveUniformsiv");
// 	if(ptrglGetActiveUniformsiv == NULL) return 1;
// 	ptrglGetActiveUniformName = goglGetProcAddress("glGetActiveUniformName");
// 	if(ptrglGetActiveUniformName == NULL) return 1;
// 	ptrglGetUniformBlockIndex = goglGetProcAddress("glGetUniformBlockIndex");
// 	if(ptrglGetUniformBlockIndex == NULL) return 1;
// 	ptrglGetActiveUniformBlockiv = goglGetProcAddress("glGetActiveUniformBlockiv");
// 	if(ptrglGetActiveUniformBlockiv == NULL) return 1;
// 	ptrglGetActiveUniformBlockName = goglGetProcAddress("glGetActiveUniformBlockName");
// 	if(ptrglGetActiveUniformBlockName == NULL) return 1;
// 	ptrglUniformBlockBinding = goglGetProcAddress("glUniformBlockBinding");
// 	if(ptrglUniformBlockBinding == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_3_2() {
// 	ptrglGetInteger64i_v = goglGetProcAddress("glGetInteger64i_v");
// 	if(ptrglGetInteger64i_v == NULL) return 1;
// 	ptrglGetBufferParameteri64v = goglGetProcAddress("glGetBufferParameteri64v");
// 	if(ptrglGetBufferParameteri64v == NULL) return 1;
// 	ptrglFramebufferTexture = goglGetProcAddress("glFramebufferTexture");
// 	if(ptrglFramebufferTexture == NULL) return 1;
// 	ptrglDrawElementsBaseVertex = goglGetProcAddress("glDrawElementsBaseVertex");
// 	if(ptrglDrawElementsBaseVertex == NULL) return 1;
// 	ptrglDrawRangeElementsBaseVertex = goglGetProcAddress("glDrawRangeElementsBaseVertex");
// 	if(ptrglDrawRangeElementsBaseVertex == NULL) return 1;
// 	ptrglDrawElementsInstancedBaseVertex = goglGetProcAddress("glDrawElementsInstancedBaseVertex");
// 	if(ptrglDrawElementsInstancedBaseVertex == NULL) return 1;
// 	ptrglMultiDrawElementsBaseVertex = goglGetProcAddress("glMultiDrawElementsBaseVertex");
// 	if(ptrglMultiDrawElementsBaseVertex == NULL) return 1;
// 	ptrglProvokingVertex = goglGetProcAddress("glProvokingVertex");
// 	if(ptrglProvokingVertex == NULL) return 1;
// 	ptrglFenceSync = goglGetProcAddress("glFenceSync");
// 	if(ptrglFenceSync == NULL) return 1;
// 	ptrglIsSync = goglGetProcAddress("glIsSync");
// 	if(ptrglIsSync == NULL) return 1;
// 	ptrglDeleteSync = goglGetProcAddress("glDeleteSync");
// 	if(ptrglDeleteSync == NULL) return 1;
// 	ptrglClientWaitSync = goglGetProcAddress("glClientWaitSync");
// 	if(ptrglClientWaitSync == NULL) return 1;
// 	ptrglWaitSync = goglGetProcAddress("glWaitSync");
// 	if(ptrglWaitSync == NULL) return 1;
// 	ptrglGetInteger64v = goglGetProcAddress("glGetInteger64v");
// 	if(ptrglGetInteger64v == NULL) return 1;
// 	ptrglGetSynciv = goglGetProcAddress("glGetSynciv");
// 	if(ptrglGetSynciv == NULL) return 1;
// 	ptrglTexImage2DMultisample = goglGetProcAddress("glTexImage2DMultisample");
// 	if(ptrglTexImage2DMultisample == NULL) return 1;
// 	ptrglTexImage3DMultisample = goglGetProcAddress("glTexImage3DMultisample");
// 	if(ptrglTexImage3DMultisample == NULL) return 1;
// 	ptrglGetMultisamplefv = goglGetProcAddress("glGetMultisamplefv");
// 	if(ptrglGetMultisamplefv == NULL) return 1;
// 	ptrglSampleMaski = goglGetProcAddress("glSampleMaski");
// 	if(ptrglSampleMaski == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_3_3() {
// 	ptrglVertexAttribDivisor = goglGetProcAddress("glVertexAttribDivisor");
// 	if(ptrglVertexAttribDivisor == NULL) return 1;
// 	ptrglBindFragDataLocationIndexed = goglGetProcAddress("glBindFragDataLocationIndexed");
// 	if(ptrglBindFragDataLocationIndexed == NULL) return 1;
// 	ptrglGetFragDataIndex = goglGetProcAddress("glGetFragDataIndex");
// 	if(ptrglGetFragDataIndex == NULL) return 1;
// 	ptrglGenSamplers = goglGetProcAddress("glGenSamplers");
// 	if(ptrglGenSamplers == NULL) return 1;
// 	ptrglDeleteSamplers = goglGetProcAddress("glDeleteSamplers");
// 	if(ptrglDeleteSamplers == NULL) return 1;
// 	ptrglIsSampler = goglGetProcAddress("glIsSampler");
// 	if(ptrglIsSampler == NULL) return 1;
// 	ptrglBindSampler = goglGetProcAddress("glBindSampler");
// 	if(ptrglBindSampler == NULL) return 1;
// 	ptrglSamplerParameteri = goglGetProcAddress("glSamplerParameteri");
// 	if(ptrglSamplerParameteri == NULL) return 1;
// 	ptrglSamplerParameteriv = goglGetProcAddress("glSamplerParameteriv");
// 	if(ptrglSamplerParameteriv == NULL) return 1;
// 	ptrglSamplerParameterf = goglGetProcAddress("glSamplerParameterf");
// 	if(ptrglSamplerParameterf == NULL) return 1;
// 	ptrglSamplerParameterfv = goglGetProcAddress("glSamplerParameterfv");
// 	if(ptrglSamplerParameterfv == NULL) return 1;
// 	ptrglSamplerParameterIiv = goglGetProcAddress("glSamplerParameterIiv");
// 	if(ptrglSamplerParameterIiv == NULL) return 1;
// 	ptrglSamplerParameterIuiv = goglGetProcAddress("glSamplerParameterIuiv");
// 	if(ptrglSamplerParameterIuiv == NULL) return 1;
// 	ptrglGetSamplerParameteriv = goglGetProcAddress("glGetSamplerParameteriv");
// 	if(ptrglGetSamplerParameteriv == NULL) return 1;
// 	ptrglGetSamplerParameterIiv = goglGetProcAddress("glGetSamplerParameterIiv");
// 	if(ptrglGetSamplerParameterIiv == NULL) return 1;
// 	ptrglGetSamplerParameterfv = goglGetProcAddress("glGetSamplerParameterfv");
// 	if(ptrglGetSamplerParameterfv == NULL) return 1;
// 	ptrglGetSamplerParameterIuiv = goglGetProcAddress("glGetSamplerParameterIuiv");
// 	if(ptrglGetSamplerParameterIuiv == NULL) return 1;
// 	ptrglQueryCounter = goglGetProcAddress("glQueryCounter");
// 	if(ptrglQueryCounter == NULL) return 1;
// 	ptrglGetQueryObjecti64v = goglGetProcAddress("glGetQueryObjecti64v");
// 	if(ptrglGetQueryObjecti64v == NULL) return 1;
// 	ptrglGetQueryObjectui64v = goglGetProcAddress("glGetQueryObjectui64v");
// 	if(ptrglGetQueryObjectui64v == NULL) return 1;
// 	ptrglVertexP2ui = goglGetProcAddress("glVertexP2ui");
// 	if(ptrglVertexP2ui == NULL) return 1;
// 	ptrglVertexP2uiv = goglGetProcAddress("glVertexP2uiv");
// 	if(ptrglVertexP2uiv == NULL) return 1;
// 	ptrglVertexP3ui = goglGetProcAddress("glVertexP3ui");
// 	if(ptrglVertexP3ui == NULL) return 1;
// 	ptrglVertexP3uiv = goglGetProcAddress("glVertexP3uiv");
// 	if(ptrglVertexP3uiv == NULL) return 1;
// 	ptrglVertexP4ui = goglGetProcAddress("glVertexP4ui");
// 	if(ptrglVertexP4ui == NULL) return 1;
// 	ptrglVertexP4uiv = goglGetProcAddress("glVertexP4uiv");
// 	if(ptrglVertexP4uiv == NULL) return 1;
// 	ptrglTexCoordP1ui = goglGetProcAddress("glTexCoordP1ui");
// 	if(ptrglTexCoordP1ui == NULL) return 1;
// 	ptrglTexCoordP1uiv = goglGetProcAddress("glTexCoordP1uiv");
// 	if(ptrglTexCoordP1uiv == NULL) return 1;
// 	ptrglTexCoordP2ui = goglGetProcAddress("glTexCoordP2ui");
// 	if(ptrglTexCoordP2ui == NULL) return 1;
// 	ptrglTexCoordP2uiv = goglGetProcAddress("glTexCoordP2uiv");
// 	if(ptrglTexCoordP2uiv == NULL) return 1;
// 	ptrglTexCoordP3ui = goglGetProcAddress("glTexCoordP3ui");
// 	if(ptrglTexCoordP3ui == NULL) return 1;
// 	ptrglTexCoordP3uiv = goglGetProcAddress("glTexCoordP3uiv");
// 	if(ptrglTexCoordP3uiv == NULL) return 1;
// 	ptrglTexCoordP4ui = goglGetProcAddress("glTexCoordP4ui");
// 	if(ptrglTexCoordP4ui == NULL) return 1;
// 	ptrglTexCoordP4uiv = goglGetProcAddress("glTexCoordP4uiv");
// 	if(ptrglTexCoordP4uiv == NULL) return 1;
// 	ptrglMultiTexCoordP1ui = goglGetProcAddress("glMultiTexCoordP1ui");
// 	if(ptrglMultiTexCoordP1ui == NULL) return 1;
// 	ptrglMultiTexCoordP1uiv = goglGetProcAddress("glMultiTexCoordP1uiv");
// 	if(ptrglMultiTexCoordP1uiv == NULL) return 1;
// 	ptrglMultiTexCoordP2ui = goglGetProcAddress("glMultiTexCoordP2ui");
// 	if(ptrglMultiTexCoordP2ui == NULL) return 1;
// 	ptrglMultiTexCoordP2uiv = goglGetProcAddress("glMultiTexCoordP2uiv");
// 	if(ptrglMultiTexCoordP2uiv == NULL) return 1;
// 	ptrglMultiTexCoordP3ui = goglGetProcAddress("glMultiTexCoordP3ui");
// 	if(ptrglMultiTexCoordP3ui == NULL) return 1;
// 	ptrglMultiTexCoordP3uiv = goglGetProcAddress("glMultiTexCoordP3uiv");
// 	if(ptrglMultiTexCoordP3uiv == NULL) return 1;
// 	ptrglMultiTexCoordP4ui = goglGetProcAddress("glMultiTexCoordP4ui");
// 	if(ptrglMultiTexCoordP4ui == NULL) return 1;
// 	ptrglMultiTexCoordP4uiv = goglGetProcAddress("glMultiTexCoordP4uiv");
// 	if(ptrglMultiTexCoordP4uiv == NULL) return 1;
// 	ptrglNormalP3ui = goglGetProcAddress("glNormalP3ui");
// 	if(ptrglNormalP3ui == NULL) return 1;
// 	ptrglNormalP3uiv = goglGetProcAddress("glNormalP3uiv");
// 	if(ptrglNormalP3uiv == NULL) return 1;
// 	ptrglColorP3ui = goglGetProcAddress("glColorP3ui");
// 	if(ptrglColorP3ui == NULL) return 1;
// 	ptrglColorP3uiv = goglGetProcAddress("glColorP3uiv");
// 	if(ptrglColorP3uiv == NULL) return 1;
// 	ptrglColorP4ui = goglGetProcAddress("glColorP4ui");
// 	if(ptrglColorP4ui == NULL) return 1;
// 	ptrglColorP4uiv = goglGetProcAddress("glColorP4uiv");
// 	if(ptrglColorP4uiv == NULL) return 1;
// 	ptrglSecondaryColorP3ui = goglGetProcAddress("glSecondaryColorP3ui");
// 	if(ptrglSecondaryColorP3ui == NULL) return 1;
// 	ptrglSecondaryColorP3uiv = goglGetProcAddress("glSecondaryColorP3uiv");
// 	if(ptrglSecondaryColorP3uiv == NULL) return 1;
// 	ptrglVertexAttribP1ui = goglGetProcAddress("glVertexAttribP1ui");
// 	if(ptrglVertexAttribP1ui == NULL) return 1;
// 	ptrglVertexAttribP1uiv = goglGetProcAddress("glVertexAttribP1uiv");
// 	if(ptrglVertexAttribP1uiv == NULL) return 1;
// 	ptrglVertexAttribP2ui = goglGetProcAddress("glVertexAttribP2ui");
// 	if(ptrglVertexAttribP2ui == NULL) return 1;
// 	ptrglVertexAttribP2uiv = goglGetProcAddress("glVertexAttribP2uiv");
// 	if(ptrglVertexAttribP2uiv == NULL) return 1;
// 	ptrglVertexAttribP3ui = goglGetProcAddress("glVertexAttribP3ui");
// 	if(ptrglVertexAttribP3ui == NULL) return 1;
// 	ptrglVertexAttribP3uiv = goglGetProcAddress("glVertexAttribP3uiv");
// 	if(ptrglVertexAttribP3uiv == NULL) return 1;
// 	ptrglVertexAttribP4ui = goglGetProcAddress("glVertexAttribP4ui");
// 	if(ptrglVertexAttribP4ui == NULL) return 1;
// 	ptrglVertexAttribP4uiv = goglGetProcAddress("glVertexAttribP4uiv");
// 	if(ptrglVertexAttribP4uiv == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_4_0() {
// 	ptrglMinSampleShading = goglGetProcAddress("glMinSampleShading");
// 	if(ptrglMinSampleShading == NULL) return 1;
// 	ptrglBlendEquationi = goglGetProcAddress("glBlendEquationi");
// 	if(ptrglBlendEquationi == NULL) return 1;
// 	ptrglBlendEquationSeparatei = goglGetProcAddress("glBlendEquationSeparatei");
// 	if(ptrglBlendEquationSeparatei == NULL) return 1;
// 	ptrglBlendFunci = goglGetProcAddress("glBlendFunci");
// 	if(ptrglBlendFunci == NULL) return 1;
// 	ptrglBlendFuncSeparatei = goglGetProcAddress("glBlendFuncSeparatei");
// 	if(ptrglBlendFuncSeparatei == NULL) return 1;
// 	ptrglDrawArraysIndirect = goglGetProcAddress("glDrawArraysIndirect");
// 	if(ptrglDrawArraysIndirect == NULL) return 1;
// 	ptrglDrawElementsIndirect = goglGetProcAddress("glDrawElementsIndirect");
// 	if(ptrglDrawElementsIndirect == NULL) return 1;
// 	ptrglUniform1d = goglGetProcAddress("glUniform1d");
// 	if(ptrglUniform1d == NULL) return 1;
// 	ptrglUniform2d = goglGetProcAddress("glUniform2d");
// 	if(ptrglUniform2d == NULL) return 1;
// 	ptrglUniform3d = goglGetProcAddress("glUniform3d");
// 	if(ptrglUniform3d == NULL) return 1;
// 	ptrglUniform4d = goglGetProcAddress("glUniform4d");
// 	if(ptrglUniform4d == NULL) return 1;
// 	ptrglUniform1dv = goglGetProcAddress("glUniform1dv");
// 	if(ptrglUniform1dv == NULL) return 1;
// 	ptrglUniform2dv = goglGetProcAddress("glUniform2dv");
// 	if(ptrglUniform2dv == NULL) return 1;
// 	ptrglUniform3dv = goglGetProcAddress("glUniform3dv");
// 	if(ptrglUniform3dv == NULL) return 1;
// 	ptrglUniform4dv = goglGetProcAddress("glUniform4dv");
// 	if(ptrglUniform4dv == NULL) return 1;
// 	ptrglUniformMatrix2dv = goglGetProcAddress("glUniformMatrix2dv");
// 	if(ptrglUniformMatrix2dv == NULL) return 1;
// 	ptrglUniformMatrix3dv = goglGetProcAddress("glUniformMatrix3dv");
// 	if(ptrglUniformMatrix3dv == NULL) return 1;
// 	ptrglUniformMatrix4dv = goglGetProcAddress("glUniformMatrix4dv");
// 	if(ptrglUniformMatrix4dv == NULL) return 1;
// 	ptrglUniformMatrix2x3dv = goglGetProcAddress("glUniformMatrix2x3dv");
// 	if(ptrglUniformMatrix2x3dv == NULL) return 1;
// 	ptrglUniformMatrix2x4dv = goglGetProcAddress("glUniformMatrix2x4dv");
// 	if(ptrglUniformMatrix2x4dv == NULL) return 1;
// 	ptrglUniformMatrix3x2dv = goglGetProcAddress("glUniformMatrix3x2dv");
// 	if(ptrglUniformMatrix3x2dv == NULL) return 1;
// 	ptrglUniformMatrix3x4dv = goglGetProcAddress("glUniformMatrix3x4dv");
// 	if(ptrglUniformMatrix3x4dv == NULL) return 1;
// 	ptrglUniformMatrix4x2dv = goglGetProcAddress("glUniformMatrix4x2dv");
// 	if(ptrglUniformMatrix4x2dv == NULL) return 1;
// 	ptrglUniformMatrix4x3dv = goglGetProcAddress("glUniformMatrix4x3dv");
// 	if(ptrglUniformMatrix4x3dv == NULL) return 1;
// 	ptrglGetUniformdv = goglGetProcAddress("glGetUniformdv");
// 	if(ptrglGetUniformdv == NULL) return 1;
// 	ptrglGetSubroutineUniformLocation = goglGetProcAddress("glGetSubroutineUniformLocation");
// 	if(ptrglGetSubroutineUniformLocation == NULL) return 1;
// 	ptrglGetSubroutineIndex = goglGetProcAddress("glGetSubroutineIndex");
// 	if(ptrglGetSubroutineIndex == NULL) return 1;
// 	ptrglGetActiveSubroutineUniformiv = goglGetProcAddress("glGetActiveSubroutineUniformiv");
// 	if(ptrglGetActiveSubroutineUniformiv == NULL) return 1;
// 	ptrglGetActiveSubroutineUniformName = goglGetProcAddress("glGetActiveSubroutineUniformName");
// 	if(ptrglGetActiveSubroutineUniformName == NULL) return 1;
// 	ptrglGetActiveSubroutineName = goglGetProcAddress("glGetActiveSubroutineName");
// 	if(ptrglGetActiveSubroutineName == NULL) return 1;
// 	ptrglUniformSubroutinesuiv = goglGetProcAddress("glUniformSubroutinesuiv");
// 	if(ptrglUniformSubroutinesuiv == NULL) return 1;
// 	ptrglGetUniformSubroutineuiv = goglGetProcAddress("glGetUniformSubroutineuiv");
// 	if(ptrglGetUniformSubroutineuiv == NULL) return 1;
// 	ptrglGetProgramStageiv = goglGetProcAddress("glGetProgramStageiv");
// 	if(ptrglGetProgramStageiv == NULL) return 1;
// 	ptrglPatchParameteri = goglGetProcAddress("glPatchParameteri");
// 	if(ptrglPatchParameteri == NULL) return 1;
// 	ptrglPatchParameterfv = goglGetProcAddress("glPatchParameterfv");
// 	if(ptrglPatchParameterfv == NULL) return 1;
// 	ptrglBindTransformFeedback = goglGetProcAddress("glBindTransformFeedback");
// 	if(ptrglBindTransformFeedback == NULL) return 1;
// 	ptrglDeleteTransformFeedbacks = goglGetProcAddress("glDeleteTransformFeedbacks");
// 	if(ptrglDeleteTransformFeedbacks == NULL) return 1;
// 	ptrglGenTransformFeedbacks = goglGetProcAddress("glGenTransformFeedbacks");
// 	if(ptrglGenTransformFeedbacks == NULL) return 1;
// 	ptrglIsTransformFeedback = goglGetProcAddress("glIsTransformFeedback");
// 	if(ptrglIsTransformFeedback == NULL) return 1;
// 	ptrglPauseTransformFeedback = goglGetProcAddress("glPauseTransformFeedback");
// 	if(ptrglPauseTransformFeedback == NULL) return 1;
// 	ptrglResumeTransformFeedback = goglGetProcAddress("glResumeTransformFeedback");
// 	if(ptrglResumeTransformFeedback == NULL) return 1;
// 	ptrglDrawTransformFeedback = goglGetProcAddress("glDrawTransformFeedback");
// 	if(ptrglDrawTransformFeedback == NULL) return 1;
// 	ptrglDrawTransformFeedbackStream = goglGetProcAddress("glDrawTransformFeedbackStream");
// 	if(ptrglDrawTransformFeedbackStream == NULL) return 1;
// 	ptrglBeginQueryIndexed = goglGetProcAddress("glBeginQueryIndexed");
// 	if(ptrglBeginQueryIndexed == NULL) return 1;
// 	ptrglEndQueryIndexed = goglGetProcAddress("glEndQueryIndexed");
// 	if(ptrglEndQueryIndexed == NULL) return 1;
// 	ptrglGetQueryIndexediv = goglGetProcAddress("glGetQueryIndexediv");
// 	if(ptrglGetQueryIndexediv == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_4_1() {
// 	ptrglReleaseShaderCompiler = goglGetProcAddress("glReleaseShaderCompiler");
// 	if(ptrglReleaseShaderCompiler == NULL) return 1;
// 	ptrglShaderBinary = goglGetProcAddress("glShaderBinary");
// 	if(ptrglShaderBinary == NULL) return 1;
// 	ptrglGetShaderPrecisionFormat = goglGetProcAddress("glGetShaderPrecisionFormat");
// 	if(ptrglGetShaderPrecisionFormat == NULL) return 1;
// 	ptrglDepthRangef = goglGetProcAddress("glDepthRangef");
// 	if(ptrglDepthRangef == NULL) return 1;
// 	ptrglClearDepthf = goglGetProcAddress("glClearDepthf");
// 	if(ptrglClearDepthf == NULL) return 1;
// 	ptrglGetProgramBinary = goglGetProcAddress("glGetProgramBinary");
// 	if(ptrglGetProgramBinary == NULL) return 1;
// 	ptrglProgramBinary = goglGetProcAddress("glProgramBinary");
// 	if(ptrglProgramBinary == NULL) return 1;
// 	ptrglProgramParameteri = goglGetProcAddress("glProgramParameteri");
// 	if(ptrglProgramParameteri == NULL) return 1;
// 	ptrglUseProgramStages = goglGetProcAddress("glUseProgramStages");
// 	if(ptrglUseProgramStages == NULL) return 1;
// 	ptrglActiveShaderProgram = goglGetProcAddress("glActiveShaderProgram");
// 	if(ptrglActiveShaderProgram == NULL) return 1;
// 	ptrglCreateShaderProgramv = goglGetProcAddress("glCreateShaderProgramv");
// 	if(ptrglCreateShaderProgramv == NULL) return 1;
// 	ptrglBindProgramPipeline = goglGetProcAddress("glBindProgramPipeline");
// 	if(ptrglBindProgramPipeline == NULL) return 1;
// 	ptrglDeleteProgramPipelines = goglGetProcAddress("glDeleteProgramPipelines");
// 	if(ptrglDeleteProgramPipelines == NULL) return 1;
// 	ptrglGenProgramPipelines = goglGetProcAddress("glGenProgramPipelines");
// 	if(ptrglGenProgramPipelines == NULL) return 1;
// 	ptrglIsProgramPipeline = goglGetProcAddress("glIsProgramPipeline");
// 	if(ptrglIsProgramPipeline == NULL) return 1;
// 	ptrglGetProgramPipelineiv = goglGetProcAddress("glGetProgramPipelineiv");
// 	if(ptrglGetProgramPipelineiv == NULL) return 1;
// 	ptrglProgramUniform1i = goglGetProcAddress("glProgramUniform1i");
// 	if(ptrglProgramUniform1i == NULL) return 1;
// 	ptrglProgramUniform1iv = goglGetProcAddress("glProgramUniform1iv");
// 	if(ptrglProgramUniform1iv == NULL) return 1;
// 	ptrglProgramUniform1f = goglGetProcAddress("glProgramUniform1f");
// 	if(ptrglProgramUniform1f == NULL) return 1;
// 	ptrglProgramUniform1fv = goglGetProcAddress("glProgramUniform1fv");
// 	if(ptrglProgramUniform1fv == NULL) return 1;
// 	ptrglProgramUniform1d = goglGetProcAddress("glProgramUniform1d");
// 	if(ptrglProgramUniform1d == NULL) return 1;
// 	ptrglProgramUniform1dv = goglGetProcAddress("glProgramUniform1dv");
// 	if(ptrglProgramUniform1dv == NULL) return 1;
// 	ptrglProgramUniform1ui = goglGetProcAddress("glProgramUniform1ui");
// 	if(ptrglProgramUniform1ui == NULL) return 1;
// 	ptrglProgramUniform1uiv = goglGetProcAddress("glProgramUniform1uiv");
// 	if(ptrglProgramUniform1uiv == NULL) return 1;
// 	ptrglProgramUniform2i = goglGetProcAddress("glProgramUniform2i");
// 	if(ptrglProgramUniform2i == NULL) return 1;
// 	ptrglProgramUniform2iv = goglGetProcAddress("glProgramUniform2iv");
// 	if(ptrglProgramUniform2iv == NULL) return 1;
// 	ptrglProgramUniform2f = goglGetProcAddress("glProgramUniform2f");
// 	if(ptrglProgramUniform2f == NULL) return 1;
// 	ptrglProgramUniform2fv = goglGetProcAddress("glProgramUniform2fv");
// 	if(ptrglProgramUniform2fv == NULL) return 1;
// 	ptrglProgramUniform2d = goglGetProcAddress("glProgramUniform2d");
// 	if(ptrglProgramUniform2d == NULL) return 1;
// 	ptrglProgramUniform2dv = goglGetProcAddress("glProgramUniform2dv");
// 	if(ptrglProgramUniform2dv == NULL) return 1;
// 	ptrglProgramUniform2ui = goglGetProcAddress("glProgramUniform2ui");
// 	if(ptrglProgramUniform2ui == NULL) return 1;
// 	ptrglProgramUniform2uiv = goglGetProcAddress("glProgramUniform2uiv");
// 	if(ptrglProgramUniform2uiv == NULL) return 1;
// 	ptrglProgramUniform3i = goglGetProcAddress("glProgramUniform3i");
// 	if(ptrglProgramUniform3i == NULL) return 1;
// 	ptrglProgramUniform3iv = goglGetProcAddress("glProgramUniform3iv");
// 	if(ptrglProgramUniform3iv == NULL) return 1;
// 	ptrglProgramUniform3f = goglGetProcAddress("glProgramUniform3f");
// 	if(ptrglProgramUniform3f == NULL) return 1;
// 	ptrglProgramUniform3fv = goglGetProcAddress("glProgramUniform3fv");
// 	if(ptrglProgramUniform3fv == NULL) return 1;
// 	ptrglProgramUniform3d = goglGetProcAddress("glProgramUniform3d");
// 	if(ptrglProgramUniform3d == NULL) return 1;
// 	ptrglProgramUniform3dv = goglGetProcAddress("glProgramUniform3dv");
// 	if(ptrglProgramUniform3dv == NULL) return 1;
// 	ptrglProgramUniform3ui = goglGetProcAddress("glProgramUniform3ui");
// 	if(ptrglProgramUniform3ui == NULL) return 1;
// 	ptrglProgramUniform3uiv = goglGetProcAddress("glProgramUniform3uiv");
// 	if(ptrglProgramUniform3uiv == NULL) return 1;
// 	ptrglProgramUniform4i = goglGetProcAddress("glProgramUniform4i");
// 	if(ptrglProgramUniform4i == NULL) return 1;
// 	ptrglProgramUniform4iv = goglGetProcAddress("glProgramUniform4iv");
// 	if(ptrglProgramUniform4iv == NULL) return 1;
// 	ptrglProgramUniform4f = goglGetProcAddress("glProgramUniform4f");
// 	if(ptrglProgramUniform4f == NULL) return 1;
// 	ptrglProgramUniform4fv = goglGetProcAddress("glProgramUniform4fv");
// 	if(ptrglProgramUniform4fv == NULL) return 1;
// 	ptrglProgramUniform4d = goglGetProcAddress("glProgramUniform4d");
// 	if(ptrglProgramUniform4d == NULL) return 1;
// 	ptrglProgramUniform4dv = goglGetProcAddress("glProgramUniform4dv");
// 	if(ptrglProgramUniform4dv == NULL) return 1;
// 	ptrglProgramUniform4ui = goglGetProcAddress("glProgramUniform4ui");
// 	if(ptrglProgramUniform4ui == NULL) return 1;
// 	ptrglProgramUniform4uiv = goglGetProcAddress("glProgramUniform4uiv");
// 	if(ptrglProgramUniform4uiv == NULL) return 1;
// 	ptrglProgramUniformMatrix2fv = goglGetProcAddress("glProgramUniformMatrix2fv");
// 	if(ptrglProgramUniformMatrix2fv == NULL) return 1;
// 	ptrglProgramUniformMatrix3fv = goglGetProcAddress("glProgramUniformMatrix3fv");
// 	if(ptrglProgramUniformMatrix3fv == NULL) return 1;
// 	ptrglProgramUniformMatrix4fv = goglGetProcAddress("glProgramUniformMatrix4fv");
// 	if(ptrglProgramUniformMatrix4fv == NULL) return 1;
// 	ptrglProgramUniformMatrix2dv = goglGetProcAddress("glProgramUniformMatrix2dv");
// 	if(ptrglProgramUniformMatrix2dv == NULL) return 1;
// 	ptrglProgramUniformMatrix3dv = goglGetProcAddress("glProgramUniformMatrix3dv");
// 	if(ptrglProgramUniformMatrix3dv == NULL) return 1;
// 	ptrglProgramUniformMatrix4dv = goglGetProcAddress("glProgramUniformMatrix4dv");
// 	if(ptrglProgramUniformMatrix4dv == NULL) return 1;
// 	ptrglProgramUniformMatrix2x3fv = goglGetProcAddress("glProgramUniformMatrix2x3fv");
// 	if(ptrglProgramUniformMatrix2x3fv == NULL) return 1;
// 	ptrglProgramUniformMatrix3x2fv = goglGetProcAddress("glProgramUniformMatrix3x2fv");
// 	if(ptrglProgramUniformMatrix3x2fv == NULL) return 1;
// 	ptrglProgramUniformMatrix2x4fv = goglGetProcAddress("glProgramUniformMatrix2x4fv");
// 	if(ptrglProgramUniformMatrix2x4fv == NULL) return 1;
// 	ptrglProgramUniformMatrix4x2fv = goglGetProcAddress("glProgramUniformMatrix4x2fv");
// 	if(ptrglProgramUniformMatrix4x2fv == NULL) return 1;
// 	ptrglProgramUniformMatrix3x4fv = goglGetProcAddress("glProgramUniformMatrix3x4fv");
// 	if(ptrglProgramUniformMatrix3x4fv == NULL) return 1;
// 	ptrglProgramUniformMatrix4x3fv = goglGetProcAddress("glProgramUniformMatrix4x3fv");
// 	if(ptrglProgramUniformMatrix4x3fv == NULL) return 1;
// 	ptrglProgramUniformMatrix2x3dv = goglGetProcAddress("glProgramUniformMatrix2x3dv");
// 	if(ptrglProgramUniformMatrix2x3dv == NULL) return 1;
// 	ptrglProgramUniformMatrix3x2dv = goglGetProcAddress("glProgramUniformMatrix3x2dv");
// 	if(ptrglProgramUniformMatrix3x2dv == NULL) return 1;
// 	ptrglProgramUniformMatrix2x4dv = goglGetProcAddress("glProgramUniformMatrix2x4dv");
// 	if(ptrglProgramUniformMatrix2x4dv == NULL) return 1;
// 	ptrglProgramUniformMatrix4x2dv = goglGetProcAddress("glProgramUniformMatrix4x2dv");
// 	if(ptrglProgramUniformMatrix4x2dv == NULL) return 1;
// 	ptrglProgramUniformMatrix3x4dv = goglGetProcAddress("glProgramUniformMatrix3x4dv");
// 	if(ptrglProgramUniformMatrix3x4dv == NULL) return 1;
// 	ptrglProgramUniformMatrix4x3dv = goglGetProcAddress("glProgramUniformMatrix4x3dv");
// 	if(ptrglProgramUniformMatrix4x3dv == NULL) return 1;
// 	ptrglValidateProgramPipeline = goglGetProcAddress("glValidateProgramPipeline");
// 	if(ptrglValidateProgramPipeline == NULL) return 1;
// 	ptrglGetProgramPipelineInfoLog = goglGetProcAddress("glGetProgramPipelineInfoLog");
// 	if(ptrglGetProgramPipelineInfoLog == NULL) return 1;
// 	ptrglVertexAttribL1d = goglGetProcAddress("glVertexAttribL1d");
// 	if(ptrglVertexAttribL1d == NULL) return 1;
// 	ptrglVertexAttribL2d = goglGetProcAddress("glVertexAttribL2d");
// 	if(ptrglVertexAttribL2d == NULL) return 1;
// 	ptrglVertexAttribL3d = goglGetProcAddress("glVertexAttribL3d");
// 	if(ptrglVertexAttribL3d == NULL) return 1;
// 	ptrglVertexAttribL4d = goglGetProcAddress("glVertexAttribL4d");
// 	if(ptrglVertexAttribL4d == NULL) return 1;
// 	ptrglVertexAttribL1dv = goglGetProcAddress("glVertexAttribL1dv");
// 	if(ptrglVertexAttribL1dv == NULL) return 1;
// 	ptrglVertexAttribL2dv = goglGetProcAddress("glVertexAttribL2dv");
// 	if(ptrglVertexAttribL2dv == NULL) return 1;
// 	ptrglVertexAttribL3dv = goglGetProcAddress("glVertexAttribL3dv");
// 	if(ptrglVertexAttribL3dv == NULL) return 1;
// 	ptrglVertexAttribL4dv = goglGetProcAddress("glVertexAttribL4dv");
// 	if(ptrglVertexAttribL4dv == NULL) return 1;
// 	ptrglVertexAttribLPointer = goglGetProcAddress("glVertexAttribLPointer");
// 	if(ptrglVertexAttribLPointer == NULL) return 1;
// 	ptrglGetVertexAttribLdv = goglGetProcAddress("glGetVertexAttribLdv");
// 	if(ptrglGetVertexAttribLdv == NULL) return 1;
// 	ptrglViewportArrayv = goglGetProcAddress("glViewportArrayv");
// 	if(ptrglViewportArrayv == NULL) return 1;
// 	ptrglViewportIndexedf = goglGetProcAddress("glViewportIndexedf");
// 	if(ptrglViewportIndexedf == NULL) return 1;
// 	ptrglViewportIndexedfv = goglGetProcAddress("glViewportIndexedfv");
// 	if(ptrglViewportIndexedfv == NULL) return 1;
// 	ptrglScissorArrayv = goglGetProcAddress("glScissorArrayv");
// 	if(ptrglScissorArrayv == NULL) return 1;
// 	ptrglScissorIndexed = goglGetProcAddress("glScissorIndexed");
// 	if(ptrglScissorIndexed == NULL) return 1;
// 	ptrglScissorIndexedv = goglGetProcAddress("glScissorIndexedv");
// 	if(ptrglScissorIndexedv == NULL) return 1;
// 	ptrglDepthRangeArrayv = goglGetProcAddress("glDepthRangeArrayv");
// 	if(ptrglDepthRangeArrayv == NULL) return 1;
// 	ptrglDepthRangeIndexed = goglGetProcAddress("glDepthRangeIndexed");
// 	if(ptrglDepthRangeIndexed == NULL) return 1;
// 	ptrglGetFloati_v = goglGetProcAddress("glGetFloati_v");
// 	if(ptrglGetFloati_v == NULL) return 1;
// 	ptrglGetDoublei_v = goglGetProcAddress("glGetDoublei_v");
// 	if(ptrglGetDoublei_v == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_4_2() {
// 	ptrglDrawArraysInstancedBaseInstance = goglGetProcAddress("glDrawArraysInstancedBaseInstance");
// 	if(ptrglDrawArraysInstancedBaseInstance == NULL) return 1;
// 	ptrglDrawElementsInstancedBaseInstance = goglGetProcAddress("glDrawElementsInstancedBaseInstance");
// 	if(ptrglDrawElementsInstancedBaseInstance == NULL) return 1;
// 	ptrglDrawElementsInstancedBaseVertexBaseInstance = goglGetProcAddress("glDrawElementsInstancedBaseVertexBaseInstance");
// 	if(ptrglDrawElementsInstancedBaseVertexBaseInstance == NULL) return 1;
// 	ptrglDrawTransformFeedbackInstanced = goglGetProcAddress("glDrawTransformFeedbackInstanced");
// 	if(ptrglDrawTransformFeedbackInstanced == NULL) return 1;
// 	ptrglDrawTransformFeedbackStreamInstanced = goglGetProcAddress("glDrawTransformFeedbackStreamInstanced");
// 	if(ptrglDrawTransformFeedbackStreamInstanced == NULL) return 1;
// 	ptrglGetInternalformativ = goglGetProcAddress("glGetInternalformativ");
// 	if(ptrglGetInternalformativ == NULL) return 1;
// 	ptrglGetActiveAtomicCounterBufferiv = goglGetProcAddress("glGetActiveAtomicCounterBufferiv");
// 	if(ptrglGetActiveAtomicCounterBufferiv == NULL) return 1;
// 	ptrglBindImageTexture = goglGetProcAddress("glBindImageTexture");
// 	if(ptrglBindImageTexture == NULL) return 1;
// 	ptrglMemoryBarrier = goglGetProcAddress("glMemoryBarrier");
// 	if(ptrglMemoryBarrier == NULL) return 1;
// 	ptrglTexStorage1D = goglGetProcAddress("glTexStorage1D");
// 	if(ptrglTexStorage1D == NULL) return 1;
// 	ptrglTexStorage2D = goglGetProcAddress("glTexStorage2D");
// 	if(ptrglTexStorage2D == NULL) return 1;
// 	ptrglTexStorage3D = goglGetProcAddress("glTexStorage3D");
// 	if(ptrglTexStorage3D == NULL) return 1;
// 	ptrglTextureStorage1DEXT = goglGetProcAddress("glTextureStorage1DEXT");
// 	if(ptrglTextureStorage1DEXT == NULL) return 1;
// 	ptrglTextureStorage2DEXT = goglGetProcAddress("glTextureStorage2DEXT");
// 	if(ptrglTextureStorage2DEXT == NULL) return 1;
// 	ptrglTextureStorage3DEXT = goglGetProcAddress("glTextureStorage3DEXT");
// 	if(ptrglTextureStorage3DEXT == NULL) return 1;
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
	TEXTURE = 0x1702
	LINE_LOOP = 0x0002
	CW = 0x0900
	OR_REVERSE = 0x150B
	ONE_MINUS_DST_COLOR = 0x0307
	BLUE = 0x1905
	COPY_INVERTED = 0x150C
	PACK_ALIGNMENT = 0x0D05
	CCW = 0x0901
	TEXTURE_BORDER_COLOR = 0x1004
	TEXTURE_MIN_FILTER = 0x2801
	RIGHT = 0x0407
	VENDOR = 0x1F00
	COLOR = 0x1800
	ALWAYS = 0x0207
	FRONT_AND_BACK = 0x0408
	UNSIGNED_BYTE = 0x1401
	STENCIL_INDEX = 0x1901
	FASTEST = 0x1101
	MAX_VIEWPORT_DIMS = 0x0D3A
	INVALID_OPERATION = 0x0502
	INVERT = 0x150A
	UNPACK_SWAP_BYTES = 0x0CF0
	INVALID_VALUE = 0x0501
	TRIANGLE_FAN = 0x0006
	UNPACK_SKIP_PIXELS = 0x0CF4
	POINT = 0x1B00
	PROXY_TEXTURE_1D = 0x8063
	UNSIGNED_INT = 0x1405
	SET = 0x150F
	FRONT_RIGHT = 0x0401
	STENCIL_TEST = 0x0B90
	DRAW_BUFFER = 0x0C01
	ONE_MINUS_SRC_ALPHA = 0x0303
	SCISSOR_TEST = 0x0C11
	DEPTH_COMPONENT = 0x1902
	PACK_SKIP_ROWS = 0x0D03
	EXTENSIONS = 0x1F03
	DONT_CARE = 0x1100
	AND_INVERTED = 0x1504
	RGB5 = 0x8050
	SRC_ALPHA_SATURATE = 0x0308
	POINT_SIZE_RANGE = 0x0B12
	NEAREST = 0x2600
	TEXTURE_BLUE_SIZE = 0x805E
	NOTEQUAL = 0x0205
	FLOAT = 0x1406
	DECR = 0x1E03
	STENCIL_PASS_DEPTH_FAIL = 0x0B95
	SUBPIXEL_BITS = 0x0D50
	NONE = 0
	LINES = 0x0001
	LINE_SMOOTH = 0x0B20
	SCISSOR_BOX = 0x0C10
	PACK_ROW_LENGTH = 0x0D02
	NO_ERROR = 0
	LINE_STRIP = 0x0003
	GREATER = 0x0204
	STENCIL_CLEAR_VALUE = 0x0B91
	EQUAL = 0x0202
	TRUE = 1
	DST_ALPHA = 0x0304
	LINEAR = 0x2601
	FRONT = 0x0404
	POLYGON_OFFSET_FILL = 0x8037
	POLYGON_OFFSET_LINE = 0x2A02
	RGB8 = 0x8051
	ONE_MINUS_DST_ALPHA = 0x0305
	VIEWPORT = 0x0BA2
	CLEAR = 0x1500
	NEVER = 0x0200
	ALPHA = 0x1906
	TEXTURE_BINDING_1D = 0x8068
	CULL_FACE_MODE = 0x0B45
	BYTE = 0x1400
	UNPACK_LSB_FIRST = 0x0CF1
	DOUBLE = 0x140A
	OR = 0x1507
	AND_REVERSE = 0x1502
	TEXTURE_WRAP_S = 0x2802
	AND = 0x1501
	DEPTH_CLEAR_VALUE = 0x0B73
	STEREO = 0x0C33
	FILL = 0x1B02
	NEAREST_MIPMAP_LINEAR = 0x2702
	ZERO = 0
	PACK_SKIP_PIXELS = 0x0D04
	PACK_SWAP_BYTES = 0x0D00
	TEXTURE_RED_SIZE = 0x805C
	SRC_ALPHA = 0x0302
	LINE_WIDTH = 0x0B21
	RGBA12 = 0x805A
	FALSE = 0
	DEPTH_TEST = 0x0B71
	BLEND_DST = 0x0BE0
	STENCIL_WRITEMASK = 0x0B98
	RGB16 = 0x8054
	TEXTURE_1D = 0x0DE0
	RGB10_A2 = 0x8059
	REPEAT = 0x2901
	KEEP = 0x1E00
	DITHER = 0x0BD0
	COLOR_LOGIC_OP = 0x0BF2
	UNPACK_SKIP_ROWS = 0x0CF3
	MAX_TEXTURE_SIZE = 0x0D33
	STENCIL_FUNC = 0x0B92
	R3_G3_B2 = 0x2A10
	TEXTURE_INTERNAL_FORMAT = 0x1003
	SHORT = 0x1402
	DST_COLOR = 0x0306
	OUT_OF_MEMORY = 0x0505
	DEPTH_RANGE = 0x0B70
	LEFT = 0x0406
	TEXTURE_WRAP_T = 0x2803
	DEPTH = 0x1801
	FRONT_LEFT = 0x0400
	TEXTURE_2D = 0x0DE1
	POINT_SIZE = 0x0B11
	LINEAR_MIPMAP_NEAREST = 0x2701
	LINEAR_MIPMAP_LINEAR = 0x2703
	CULL_FACE = 0x0B44
	VERSION = 0x1F02
	TEXTURE_ALPHA_SIZE = 0x805F
	NICEST = 0x1102
	RGB5_A1 = 0x8057
	POINTS = 0x0000
	TRIANGLE_STRIP = 0x0005
	RGB = 0x1907
	POLYGON_SMOOTH = 0x0B41
	REPLACE = 0x1E01
	RED = 0x1903
	GEQUAL = 0x0206
	BACK = 0x0405
	PACK_LSB_FIRST = 0x0D01
	NOOP = 0x1505
	RGBA8 = 0x8058
	DEPTH_FUNC = 0x0B74
	LINE_SMOOTH_HINT = 0x0C52
	EQUIV = 0x1509
	LINE = 0x1B01
	STENCIL = 0x1802
	POLYGON_OFFSET_POINT = 0x2A01
	TEXTURE_GREEN_SIZE = 0x805D
	LINE_WIDTH_GRANULARITY = 0x0B23
	COLOR_WRITEMASK = 0x0C23
	TEXTURE_WIDTH = 0x1000
	PROXY_TEXTURE_2D = 0x8064
	TEXTURE_MAG_FILTER = 0x2800
	NEAREST_MIPMAP_NEAREST = 0x2700
	LINE_WIDTH_RANGE = 0x0B22
	TRIANGLES = 0x0004
	POLYGON_OFFSET_FACTOR = 0x8038
	RGBA2 = 0x8055
	LESS = 0x0201
	UNPACK_ROW_LENGTH = 0x0CF2
	STENCIL_FAIL = 0x0B94
	DEPTH_WRITEMASK = 0x0B72
	BLEND = 0x0BE2
	RGB10 = 0x8052
	STENCIL_PASS_DEPTH_PASS = 0x0B96
	RENDERER = 0x1F01
	LEQUAL = 0x0203
	UNPACK_ALIGNMENT = 0x0CF5
	DEPTH_BUFFER_BIT = 0x00000100
	BLEND_SRC = 0x0BE1
	ONE_MINUS_SRC_COLOR = 0x0301
	RGBA = 0x1908
	BACK_LEFT = 0x0402
	FRONT_FACE = 0x0B46
	SRC_COLOR = 0x0300
	COLOR_CLEAR_VALUE = 0x0C22
	UNSIGNED_SHORT = 0x1403
	STENCIL_REF = 0x0B97
	RGBA16 = 0x805B
	GREEN = 0x1904
	INT = 0x1404
	STENCIL_BUFFER_BIT = 0x00000400
	BACK_RIGHT = 0x0403
	READ_BUFFER = 0x0C02
	OR_INVERTED = 0x150D
	NAND = 0x150E
	INCR = 0x1E02
	RGB12 = 0x8053
	LOGIC_OP_MODE = 0x0BF0
	POLYGON_OFFSET_UNITS = 0x2A00
	RGBA4 = 0x8056
	NOR = 0x1508
	POINT_SIZE_GRANULARITY = 0x0B13
	INVALID_ENUM = 0x0500
	ONE = 1
	STENCIL_VALUE_MASK = 0x0B93
	COLOR_BUFFER_BIT = 0x00004000
	XOR = 0x1506
	DOUBLEBUFFER = 0x0C32
	POLYGON_SMOOTH_HINT = 0x0C53
	TEXTURE_HEIGHT = 0x1001
	TEXTURE_BINDING_2D = 0x8069
	COPY = 0x1503
	RGB4 = 0x804F
)
// VERSION_1_2
const (
	SMOOTH_LINE_WIDTH_RANGE = 0x0B22
	SMOOTH_LINE_WIDTH_GRANULARITY = 0x0B23
	PACK_SKIP_IMAGES = 0x806B
	UNSIGNED_BYTE_2_3_3_REV = 0x8362
	UNSIGNED_INT_8_8_8_8 = 0x8035
	UNSIGNED_SHORT_4_4_4_4_REV = 0x8365
	UNSIGNED_SHORT_5_6_5_REV = 0x8364
	TEXTURE_3D = 0x806F
	PROXY_TEXTURE_3D = 0x8070
	CLAMP_TO_EDGE = 0x812F
	MAX_ELEMENTS_INDICES = 0x80E9
	SMOOTH_POINT_SIZE_RANGE = 0x0B12
	UNSIGNED_SHORT_4_4_4_4 = 0x8033
	UNSIGNED_INT_10_10_10_2 = 0x8036
	TEXTURE_DEPTH = 0x8071
	SMOOTH_POINT_SIZE_GRANULARITY = 0x0B13
	BGR = 0x80E0
	MAX_3D_TEXTURE_SIZE = 0x8073
	UNPACK_SKIP_IMAGES = 0x806D
	UNSIGNED_BYTE_3_3_2 = 0x8032
	BGRA = 0x80E1
	TEXTURE_BINDING_3D = 0x806A
	UNSIGNED_SHORT_5_5_5_1 = 0x8034
	UNPACK_IMAGE_HEIGHT = 0x806E
	TEXTURE_BASE_LEVEL = 0x813C
	TEXTURE_WRAP_R = 0x8072
	TEXTURE_MIN_LOD = 0x813A
	ALIASED_LINE_WIDTH_RANGE = 0x846E
	UNSIGNED_SHORT_1_5_5_5_REV = 0x8366
	TEXTURE_MAX_LOD = 0x813B
	PACK_IMAGE_HEIGHT = 0x806C
	TEXTURE_MAX_LEVEL = 0x813D
	UNSIGNED_INT_2_10_10_10_REV = 0x8368
	UNSIGNED_INT_8_8_8_8_REV = 0x8367
	MAX_ELEMENTS_VERTICES = 0x80E8
	UNSIGNED_SHORT_5_6_5 = 0x8363
)
// VERSION_1_3
const (
	TEXTURE_COMPRESSED = 0x86A1
	COMPRESSED_RGBA = 0x84EE
	TEXTURE_BINDING_CUBE_MAP = 0x8514
	TEXTURE17 = 0x84D1
	TEXTURE30 = 0x84DE
	TEXTURE29 = 0x84DD
	TEXTURE22 = 0x84D6
	TEXTURE18 = 0x84D2
	NUM_COMPRESSED_TEXTURE_FORMATS = 0x86A2
	TEXTURE_COMPRESSED_IMAGE_SIZE = 0x86A0
	TEXTURE13 = 0x84CD
	SAMPLE_COVERAGE = 0x80A0
	TEXTURE12 = 0x84CC
	SAMPLES = 0x80A9
	TEXTURE4 = 0x84C4
	TEXTURE7 = 0x84C7
	SAMPLE_COVERAGE_VALUE = 0x80AA
	TEXTURE_CUBE_MAP_NEGATIVE_X = 0x8516
	TEXTURE0 = 0x84C0
	TEXTURE16 = 0x84D0
	TEXTURE20 = 0x84D4
	TEXTURE28 = 0x84DC
	CLAMP_TO_BORDER = 0x812D
	MAX_CUBE_MAP_TEXTURE_SIZE = 0x851C
	TEXTURE_COMPRESSION_HINT = 0x84EF
	TEXTURE9 = 0x84C9
	TEXTURE25 = 0x84D9
	MULTISAMPLE = 0x809D
	TEXTURE21 = 0x84D5
	TEXTURE23 = 0x84D7
	ACTIVE_TEXTURE = 0x84E0
	TEXTURE8 = 0x84C8
	PROXY_TEXTURE_CUBE_MAP = 0x851B
	TEXTURE_CUBE_MAP_POSITIVE_Y = 0x8517
	SAMPLE_BUFFERS = 0x80A8
	TEXTURE10 = 0x84CA
	TEXTURE14 = 0x84CE
	TEXTURE3 = 0x84C3
	TEXTURE15 = 0x84CF
	SAMPLE_ALPHA_TO_ONE = 0x809F
	SAMPLE_ALPHA_TO_COVERAGE = 0x809E
	TEXTURE24 = 0x84D8
	TEXTURE19 = 0x84D3
	TEXTURE11 = 0x84CB
	TEXTURE31 = 0x84DF
	TEXTURE27 = 0x84DB
	TEXTURE_CUBE_MAP_NEGATIVE_Y = 0x8518
	TEXTURE1 = 0x84C1
	TEXTURE2 = 0x84C2
	SAMPLE_COVERAGE_INVERT = 0x80AB
	TEXTURE6 = 0x84C6
	COMPRESSED_TEXTURE_FORMATS = 0x86A3
	TEXTURE_CUBE_MAP = 0x8513
	TEXTURE_CUBE_MAP_NEGATIVE_Z = 0x851A
	TEXTURE_CUBE_MAP_POSITIVE_X = 0x8515
	TEXTURE26 = 0x84DA
	TEXTURE_CUBE_MAP_POSITIVE_Z = 0x8519
	TEXTURE5 = 0x84C5
	COMPRESSED_RGB = 0x84ED
)
// VERSION_1_4
const (
	MIRRORED_REPEAT = 0x8370
	TEXTURE_DEPTH_SIZE = 0x884A
	DECR_WRAP = 0x8508
	MAX_TEXTURE_LOD_BIAS = 0x84FD
	BLEND_SRC_ALPHA = 0x80CB
	BLEND_DST_ALPHA = 0x80CA
	POINT_FADE_THRESHOLD_SIZE = 0x8128
	DEPTH_COMPONENT16 = 0x81A5
	TEXTURE_LOD_BIAS = 0x8501
	DEPTH_COMPONENT32 = 0x81A7
	TEXTURE_COMPARE_FUNC = 0x884D
	DEPTH_COMPONENT24 = 0x81A6
	INCR_WRAP = 0x8507
	TEXTURE_COMPARE_MODE = 0x884C
	BLEND_DST_RGB = 0x80C8
	BLEND_SRC_RGB = 0x80C9
)
// VERSION_1_5
const (
	QUERY_COUNTER_BITS = 0x8864
	READ_ONLY = 0x88B8
	BUFFER_SIZE = 0x8764
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = 0x889F
	STREAM_READ = 0x88E1
	ELEMENT_ARRAY_BUFFER = 0x8893
	DYNAMIC_READ = 0x88E9
	STREAM_DRAW = 0x88E0
	ARRAY_BUFFER = 0x8892
	BUFFER_USAGE = 0x8765
	STATIC_COPY = 0x88E6
	STREAM_COPY = 0x88E2
	WRITE_ONLY = 0x88B9
	QUERY_RESULT_AVAILABLE = 0x8867
	DYNAMIC_DRAW = 0x88E8
	READ_WRITE = 0x88BA
	BUFFER_MAPPED = 0x88BC
	STATIC_READ = 0x88E5
	SAMPLES_PASSED = 0x8914
	DYNAMIC_COPY = 0x88EA
	ELEMENT_ARRAY_BUFFER_BINDING = 0x8895
	CURRENT_QUERY = 0x8865
	BUFFER_ACCESS = 0x88BB
	STATIC_DRAW = 0x88E4
	ARRAY_BUFFER_BINDING = 0x8894
	QUERY_RESULT = 0x8866
	BUFFER_MAP_POINTER = 0x88BD
)
// VERSION_2_0
const (
	SAMPLER_2D_SHADOW = 0x8B62
	MAX_FRAGMENT_UNIFORM_COMPONENTS = 0x8B49
	STENCIL_BACK_FAIL = 0x8801
	DRAW_BUFFER6 = 0x882B
	BLEND_EQUATION_ALPHA = 0x883D
	SAMPLER_CUBE = 0x8B60
	FLOAT_MAT4 = 0x8B5C
	VERTEX_ATTRIB_ARRAY_POINTER = 0x8645
	DRAW_BUFFER8 = 0x882D
	MAX_VARYING_FLOATS = 0x8B4B
	SHADER_SOURCE_LENGTH = 0x8B88
	FLOAT_VEC2 = 0x8B50
	LINK_STATUS = 0x8B82
	FLOAT_VEC4 = 0x8B52
	DRAW_BUFFER4 = 0x8829
	MAX_VERTEX_ATTRIBS = 0x8869
	FLOAT_MAT3 = 0x8B5B
	SAMPLER_1D = 0x8B5D
	COMPILE_STATUS = 0x8B81
	MAX_TEXTURE_IMAGE_UNITS = 0x8872
	VERTEX_ATTRIB_ARRAY_TYPE = 0x8625
	SAMPLER_1D_SHADOW = 0x8B61
	DRAW_BUFFER5 = 0x882A
	DRAW_BUFFER12 = 0x8831
	VERTEX_ATTRIB_ARRAY_NORMALIZED = 0x886A
	DRAW_BUFFER1 = 0x8826
	VERTEX_SHADER = 0x8B31
	INT_VEC3 = 0x8B54
	BOOL_VEC3 = 0x8B58
	ACTIVE_UNIFORM_MAX_LENGTH = 0x8B87
	DRAW_BUFFER7 = 0x882C
	FRAGMENT_SHADER = 0x8B30
	VERTEX_ATTRIB_ARRAY_SIZE = 0x8623
	VALIDATE_STATUS = 0x8B83
	LOWER_LEFT = 0x8CA1
	POINT_SPRITE_COORD_ORIGIN = 0x8CA0
	CURRENT_PROGRAM = 0x8B8D
	SHADING_LANGUAGE_VERSION = 0x8B8C
	DRAW_BUFFER9 = 0x882E
	BLEND_EQUATION_RGB = 0x8009
	BOOL_VEC4 = 0x8B59
	DRAW_BUFFER10 = 0x882F
	ACTIVE_ATTRIBUTE_MAX_LENGTH = 0x8B8A
	FLOAT_VEC3 = 0x8B51
	SAMPLER_2D = 0x8B5E
	STENCIL_BACK_WRITEMASK = 0x8CA5
	MAX_DRAW_BUFFERS = 0x8824
	BOOL = 0x8B56
	CURRENT_VERTEX_ATTRIB = 0x8626
	INT_VEC4 = 0x8B55
	DRAW_BUFFER14 = 0x8833
	STENCIL_BACK_FUNC = 0x8800
	ACTIVE_UNIFORMS = 0x8B86
	BOOL_VEC2 = 0x8B57
	FRAGMENT_SHADER_DERIVATIVE_HINT = 0x8B8B
	VERTEX_PROGRAM_POINT_SIZE = 0x8642
	DRAW_BUFFER0 = 0x8825
	INT_VEC2 = 0x8B53
	STENCIL_BACK_PASS_DEPTH_PASS = 0x8803
	STENCIL_BACK_REF = 0x8CA3
	DRAW_BUFFER3 = 0x8828
	INFO_LOG_LENGTH = 0x8B84
	SAMPLER_3D = 0x8B5F
	ACTIVE_ATTRIBUTES = 0x8B89
	VERTEX_ATTRIB_ARRAY_ENABLED = 0x8622
	MAX_VERTEX_UNIFORM_COMPONENTS = 0x8B4A
	DRAW_BUFFER11 = 0x8830
	DRAW_BUFFER13 = 0x8832
	MAX_COMBINED_TEXTURE_IMAGE_UNITS = 0x8B4D
	ATTACHED_SHADERS = 0x8B85
	STENCIL_BACK_PASS_DEPTH_FAIL = 0x8802
	DRAW_BUFFER15 = 0x8834
	SHADER_TYPE = 0x8B4F
	UPPER_LEFT = 0x8CA2
	FLOAT_MAT2 = 0x8B5A
	DELETE_STATUS = 0x8B80
	STENCIL_BACK_VALUE_MASK = 0x8CA4
	DRAW_BUFFER2 = 0x8827
	VERTEX_ATTRIB_ARRAY_STRIDE = 0x8624
	MAX_VERTEX_TEXTURE_IMAGE_UNITS = 0x8B4C
)
// VERSION_2_1
const (
	FLOAT_MAT3x2 = 0x8B67
	PIXEL_PACK_BUFFER_BINDING = 0x88ED
	PIXEL_PACK_BUFFER = 0x88EB
	SRGB = 0x8C40
	COMPRESSED_SRGB_ALPHA = 0x8C49
	FLOAT_MAT3x4 = 0x8B68
	PIXEL_UNPACK_BUFFER_BINDING = 0x88EF
	FLOAT_MAT2x4 = 0x8B66
	SRGB_ALPHA = 0x8C42
	SRGB8_ALPHA8 = 0x8C43
	SRGB8 = 0x8C41
	FLOAT_MAT2x3 = 0x8B65
	FLOAT_MAT4x2 = 0x8B69
	COMPRESSED_SRGB = 0x8C48
	FLOAT_MAT4x3 = 0x8B6A
	PIXEL_UNPACK_BUFFER = 0x88EC
)
// VERSION_3_0
const (
	INT_SAMPLER_CUBE = 0x8DCC
	MAX_ARRAY_TEXTURE_LAYERS = 0x88FF
	COLOR_ATTACHMENT3 = 0x8CE3
	BLUE_INTEGER = 0x8D96
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS = 0x8C80
	R32UI = 0x8236
	TEXTURE_BINDING_1D_ARRAY = 0x8C1C
	DEPTH_STENCIL = 0x84F9
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME = 0x8CD1
	COLOR_ATTACHMENT2 = 0x8CE2
	MAX_COLOR_ATTACHMENTS = 0x8CDF
	RGBA16I = 0x8D88
	RG8UI = 0x8238
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = 0x8CD7
	RGBA_INTEGER = 0x8D99
	RGBA16UI = 0x8D76
	UNSIGNED_INT_SAMPLER_1D_ARRAY = 0x8DD6
	TEXTURE_GREEN_TYPE = 0x8C11
	RG16 = 0x822C
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING = 0x8210
	FRAMEBUFFER_DEFAULT = 0x8218
	QUERY_NO_WAIT = 0x8E14
	RG32F = 0x8230
	RENDERBUFFER_DEPTH_SIZE = 0x8D54
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS = 0x8C8B
	UNSIGNED_INT_10F_11F_11F_REV = 0x8C3B
	UNSIGNED_INT_SAMPLER_1D = 0x8DD1
	RGB32F = 0x8815
	STENCIL_ATTACHMENT = 0x8D20
	COMPARE_REF_TO_TEXTURE = 0x884E
	INT_SAMPLER_1D_ARRAY = 0x8DCE
	MAP_INVALIDATE_BUFFER_BIT = 0x0008
	INT_SAMPLER_2D = 0x8DCA
	COLOR_ATTACHMENT5 = 0x8CE5
	RG_INTEGER = 0x8228
	READ_FRAMEBUFFER_BINDING = 0x8CAA
	CLAMP_READ_COLOR = 0x891C
	R16F = 0x822D
	FIXED_ONLY = 0x891D
	RGBA32I = 0x8D82
	MAX_PROGRAM_TEXEL_OFFSET = 0x8905
	BGRA_INTEGER = 0x8D9B
	TRANSFORM_FEEDBACK_VARYINGS = 0x8C83
	R8UI = 0x8232
	COLOR_ATTACHMENT15 = 0x8CEF
	RENDERBUFFER_SAMPLES = 0x8CAB
	SAMPLER_1D_ARRAY = 0x8DC0
	R32I = 0x8235
	FRAMEBUFFER_COMPLETE = 0x8CD5
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE = 0x8217
	SAMPLER_2D_ARRAY = 0x8DC1
	RG16F = 0x822F
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS = 0x8C8A
	UNSIGNED_INT_SAMPLER_2D_ARRAY = 0x8DD7
	COLOR_ATTACHMENT11 = 0x8CEB
	TEXTURE_BLUE_TYPE = 0x8C12
	FRAMEBUFFER_ATTACHMENT_RED_SIZE = 0x8212
	RGBA8I = 0x8D8E
	R32F = 0x822E
	CLIP_DISTANCE1 = 0x3001
	DRAW_FRAMEBUFFER = 0x8CA9
	TRANSFORM_FEEDBACK_BUFFER = 0x8C8E
	RGBA16F = 0x881A
	RENDERBUFFER_BINDING = 0x8CA7
	DEPTH_STENCIL_ATTACHMENT = 0x821A
	READ_FRAMEBUFFER = 0x8CA8
	R11F_G11F_B10F = 0x8C3A
	TRANSFORM_FEEDBACK_BUFFER_SIZE = 0x8C85
	DEPTH32F_STENCIL8 = 0x8CAD
	MAP_UNSYNCHRONIZED_BIT = 0x0020
	UNSIGNED_INT_VEC4 = 0x8DC8
	MAP_WRITE_BIT = 0x0002
	CLIP_DISTANCE2 = 0x3002
	PROXY_TEXTURE_2D_ARRAY = 0x8C1B
	R16I = 0x8233
	RG = 0x8227
	COMPRESSED_RED = 0x8225
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = 0x8C88
	COLOR_ATTACHMENT9 = 0x8CE9
	TEXTURE_SHARED_SIZE = 0x8C3F
	STENCIL_INDEX4 = 0x8D47
	GREEN_INTEGER = 0x8D95
	UNSIGNED_INT_24_8 = 0x84FA
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE = 0x8D56
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT = 0x8CD6
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE = 0x8213
	QUERY_BY_REGION_NO_WAIT = 0x8E16
	COLOR_ATTACHMENT13 = 0x8CED
	RENDERBUFFER = 0x8D41
	RED_INTEGER = 0x8D94
	CONTEXT_FLAGS = 0x821E
	CLIP_DISTANCE0 = 0x3000
	RENDERBUFFER_INTERNAL_FORMAT = 0x8D44
	RGB16I = 0x8D89
	RENDERBUFFER_ALPHA_SIZE = 0x8D53
	COLOR_ATTACHMENT6 = 0x8CE6
	MIN_PROGRAM_TEXEL_OFFSET = 0x8904
	BUFFER_ACCESS_FLAGS = 0x911F
	SAMPLER_1D_ARRAY_SHADOW = 0x8DC3
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT = 0x0001
	FRAMEBUFFER_UNDEFINED = 0x8219
	INT_SAMPLER_3D = 0x8DCB
	MAP_READ_BIT = 0x0001
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE = 0x8216
	RGB_INTEGER = 0x8D98
	MAX_RENDERBUFFER_SIZE = 0x84E8
	INVALID_FRAMEBUFFER_OPERATION = 0x0506
	RG16I = 0x8239
	FLOAT_32_UNSIGNED_INT_24_8_REV = 0x8DAD
	RENDERBUFFER_WIDTH = 0x8D42
	SEPARATE_ATTRIBS = 0x8C8D
	RENDERBUFFER_BLUE_SIZE = 0x8D52
	TEXTURE_DEPTH_TYPE = 0x8C16
	TEXTURE_STENCIL_SIZE = 0x88F1
	RGBA8UI = 0x8D7C
	QUERY_WAIT = 0x8E13
	DRAW_FRAMEBUFFER_BINDING = FRAMEBUFFER_BINDING
	INTERLEAVED_ATTRIBS = 0x8C8C
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH = 0x8C76
	DEPTH_COMPONENT32F = 0x8CAC
	SAMPLER_2D_ARRAY_SHADOW = 0x8DC4
	RENDERBUFFER_GREEN_SIZE = 0x8D51
	RASTERIZER_DISCARD = 0x8C89
	TEXTURE_1D_ARRAY = 0x8C18
	UNSIGNED_INT_VEC3 = 0x8DC7
	STENCIL_INDEX16 = 0x8D49
	UNSIGNED_INT_5_9_9_9_REV = 0x8C3E
	STENCIL_INDEX1 = 0x8D46
	TRANSFORM_FEEDBACK_BUFFER_START = 0x8C84
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER = 0x8CDB
	TRANSFORM_FEEDBACK_BUFFER_BINDING = 0x8C8F
	HALF_FLOAT = 0x140B
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER = 0x8CDC
	TEXTURE_2D_ARRAY = 0x8C1A
	MAX_SAMPLES = 0x8D57
	STENCIL_INDEX8 = 0x8D48
	MAP_FLUSH_EXPLICIT_BIT = 0x0010
	COMPRESSED_RG_RGTC2 = 0x8DBD
	COLOR_ATTACHMENT0 = 0x8CE0
	FRAMEBUFFER_BINDING = 0x8CA6
	MAX_CLIP_DISTANCES = 0x0D32
	R16 = 0x822A
	RGB16F = 0x881B
	RENDERBUFFER_STENCIL_SIZE = 0x8D55
	UNSIGNED_INT_SAMPLER_CUBE = 0x8DD4
	PRIMITIVES_GENERATED = 0x8C87
	FRAMEBUFFER_UNSUPPORTED = 0x8CDD
	INT_SAMPLER_2D_ARRAY = 0x8DCF
	RG16UI = 0x823A
	CLIP_DISTANCE3 = 0x3003
	CLIP_DISTANCE4 = 0x3004
	RGB32I = 0x8D83
	TRANSFORM_FEEDBACK_BUFFER_MODE = 0x8C7F
	RG32UI = 0x823C
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE = 0x8CD0
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 0x8CD3
	RGB32UI = 0x8D71
	UNSIGNED_NORMALIZED = 0x8C17
	RG32I = 0x823B
	TEXTURE_RED_TYPE = 0x8C10
	DEPTH24_STENCIL8 = 0x88F0
	QUERY_BY_REGION_WAIT = 0x8E15
	COLOR_ATTACHMENT1 = 0x8CE1
	VERTEX_ATTRIB_ARRAY_INTEGER = 0x88FD
	RGB8UI = 0x8D7D
	PROXY_TEXTURE_1D_ARRAY = 0x8C19
	TEXTURE_ALPHA_TYPE = 0x8C13
	MAP_INVALIDATE_RANGE_BIT = 0x0004
	R8I = 0x8231
	COMPRESSED_SIGNED_RG_RGTC2 = 0x8DBE
	UNSIGNED_INT_SAMPLER_2D = 0x8DD2
	BGR_INTEGER = 0x8D9A
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE = 0x8214
	CLIP_DISTANCE6 = 0x3006
	COLOR_ATTACHMENT10 = 0x8CEA
	COLOR_ATTACHMENT8 = 0x8CE8
	FRAMEBUFFER = 0x8D40
	COLOR_ATTACHMENT14 = 0x8CEE
	CLIP_DISTANCE5 = 0x3005
	COLOR_ATTACHMENT4 = 0x8CE4
	DEPTH_ATTACHMENT = 0x8D00
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE = 0x8211
	COLOR_ATTACHMENT12 = 0x8CEC
	BUFFER_MAP_LENGTH = 0x9120
	UNSIGNED_INT_VEC2 = 0x8DC6
	CLIP_DISTANCE7 = 0x3007
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE = 0x8215
	MINOR_VERSION = 0x821C
	COLOR_ATTACHMENT7 = 0x8CE7
	SAMPLER_CUBE_SHADOW = 0x8DC5
	INDEX = 0x8222
	TEXTURE_BINDING_2D_ARRAY = 0x8C1D
	RENDERBUFFER_RED_SIZE = 0x8D50
	NUM_EXTENSIONS = 0x821D
	COMPRESSED_RED_RGTC1 = 0x8DBB
	RGB16UI = 0x8D77
	BUFFER_MAP_OFFSET = 0x9121
	RGB9_E5 = 0x8C3D
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = 0x8CD2
	FRAMEBUFFER_SRGB = 0x8DB9
	R8 = 0x8229
	RGB8I = 0x8D8F
	INT_SAMPLER_1D = 0x8DC9
	UNSIGNED_INT_SAMPLER_3D = 0x8DD3
	RENDERBUFFER_HEIGHT = 0x8D43
	COMPRESSED_SIGNED_RED_RGTC1 = 0x8DBC
	VERTEX_ARRAY_BINDING = 0x85B5
	RG8 = 0x822B
	RG8I = 0x8237
	RGBA32F = 0x8814
	MAJOR_VERSION = 0x821B
	RGBA32UI = 0x8D70
	R16UI = 0x8234
	COMPRESSED_RG = 0x8226
)
// VERSION_3_1
const (
	TEXTURE_BINDING_BUFFER = 0x8C2C
	UNIFORM_BLOCK_INDEX = 0x8A3A
	PRIMITIVE_RESTART_INDEX = 0x8F9E
	RGBA_SNORM = 0x8F93
	RGB8_SNORM = 0x8F96
	MAX_UNIFORM_BUFFER_BINDINGS = 0x8A2F
	COPY_READ_BUFFER = COPY_READ_BUFFER_BINDING
	UNIFORM_TYPE = 0x8A37
	SAMPLER_2D_RECT = 0x8B63
	RG16_SNORM = 0x8F99
	UNIFORM_BUFFER_SIZE = 0x8A2A
	UNIFORM_ARRAY_STRIDE = 0x8A3C
	MAX_RECTANGLE_TEXTURE_SIZE = 0x84F8
	RED_SNORM = 0x8F90
	RGB_SNORM = 0x8F92
	TEXTURE_RECTANGLE = 0x84F5
	UNIFORM_BLOCK_BINDING = 0x8A3F
	UNIFORM_SIZE = 0x8A38
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER = 0x8A46
	COPY_WRITE_BUFFER = COPY_WRITE_BUFFER_BINDING
	TEXTURE_BINDING_RECTANGLE = 0x84F6
	UNIFORM_BLOCK_ACTIVE_UNIFORMS = 0x8A42
	UNIFORM_NAME_LENGTH = 0x8A39
	RGBA16_SNORM = 0x8F9B
	INVALID_INDEX = 0xFFFFFFFF
	UNIFORM_BUFFER = 0x8A11
	R8_SNORM = 0x8F94
	MAX_COMBINED_UNIFORM_BLOCKS = 0x8A2E
	RGB16_SNORM = 0x8F9A
	UNIFORM_BUFFER_OFFSET_ALIGNMENT = 0x8A34
	UNSIGNED_INT_SAMPLER_2D_RECT = 0x8DD5
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS = 0x8A31
	UNSIGNED_INT_SAMPLER_BUFFER = 0x8DD8
	UNIFORM_OFFSET = 0x8A3B
	TEXTURE_BUFFER_FORMAT = 0x8C2E
	UNIFORM_MATRIX_STRIDE = 0x8A3D
	INT_SAMPLER_BUFFER = 0x8DD0
	MAX_VERTEX_UNIFORM_BLOCKS = 0x8A2B
	UNIFORM_BUFFER_START = 0x8A29
	TEXTURE_BUFFER_DATA_STORE_BINDING = 0x8C2D
	R16_SNORM = 0x8F98
	INT_SAMPLER_2D_RECT = 0x8DCD
	ACTIVE_UNIFORM_BLOCKS = 0x8A36
	SAMPLER_2D_RECT_SHADOW = 0x8B64
	SAMPLER_BUFFER = 0x8DC2
	UNIFORM_BLOCK_NAME_LENGTH = 0x8A41
	RG_SNORM = 0x8F91
	MAX_TEXTURE_BUFFER_SIZE = 0x8C2B
	PROXY_TEXTURE_RECTANGLE = 0x84F7
	SIGNED_NORMALIZED = 0x8F9C
	PRIMITIVE_RESTART = 0x8F9D
	UNIFORM_IS_ROW_MAJOR = 0x8A3E
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH = 0x8A35
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES = 0x8A43
	MAX_FRAGMENT_UNIFORM_BLOCKS = 0x8A2D
	TEXTURE_BUFFER = 0x8C2A
	UNIFORM_BUFFER_BINDING = 0x8A28
	UNIFORM_BLOCK_DATA_SIZE = 0x8A40
	RG8_SNORM = 0x8F95
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS = 0x8A33
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER = 0x8A44
	MAX_UNIFORM_BLOCK_SIZE = 0x8A30
	RGBA8_SNORM = 0x8F97
)
// VERSION_3_2
const (
	MAX_SERVER_WAIT_TIMEOUT = 0x9111
	MAX_GEOMETRY_INPUT_COMPONENTS = 0x9123
	SAMPLE_MASK_VALUE = 0x8E52
	PROGRAM_POINT_SIZE = 0x8642
	TEXTURE_CUBE_MAP_SEAMLESS = 0x884F
	MAX_GEOMETRY_OUTPUT_COMPONENTS = 0x9124
	DEPTH_CLAMP = 0x864F
	SAMPLE_POSITION = 0x8E50
	MAX_GEOMETRY_UNIFORM_COMPONENTS = 0x8DDF
	CONTEXT_CORE_PROFILE_BIT = 0x00000001
	LINES_ADJACENCY = 0x000A
	UNSIGNALED = 0x9118
	ALREADY_SIGNALED = 0x911A
	LINE_STRIP_ADJACENCY = 0x000B
	SAMPLE_MASK = 0x8E51
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY = 0x9103
	MAX_VERTEX_OUTPUT_COMPONENTS = 0x9122
	SIGNALED = 0x9119
	GEOMETRY_VERTICES_OUT = 0x8916
	TEXTURE_FIXED_SAMPLE_LOCATIONS = 0x9107
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION = 0x8E4C
	MAX_COLOR_TEXTURE_SAMPLES = 0x910E
	TIMEOUT_EXPIRED = 0x911B
	SAMPLER_2D_MULTISAMPLE = 0x9108
	PROXY_TEXTURE_2D_MULTISAMPLE = 0x9101
	FRAMEBUFFER_ATTACHMENT_LAYERED = 0x8DA7
	MAX_DEPTH_TEXTURE_SAMPLES = 0x910F
	TEXTURE_BINDING_2D_MULTISAMPLE = 0x9104
	CONTEXT_COMPATIBILITY_PROFILE_BIT = 0x00000002
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS = 0x8C29
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY = 0x910D
	PROVOKING_VERTEX = 0x8E4F
	TEXTURE_2D_MULTISAMPLE_ARRAY = 0x9102
	MAX_INTEGER_SAMPLES = 0x9110
	SYNC_GPU_COMMANDS_COMPLETE = 0x9117
	MAX_SAMPLE_MASK_WORDS = 0x8E59
	GEOMETRY_OUTPUT_TYPE = 0x8918
	SYNC_FLAGS = 0x9115
	SAMPLER_2D_MULTISAMPLE_ARRAY = 0x910B
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE = 0x910A
	CONTEXT_PROFILE_MASK = 0x9126
	CONDITION_SATISFIED = 0x911C
	MAX_FRAGMENT_INPUT_COMPONENTS = 0x9125
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY = 0x910C
	TIMEOUT_IGNORED = 0xFFFFFFFFFFFFFFFF
	GEOMETRY_INPUT_TYPE = 0x8917
	TRIANGLES_ADJACENCY = 0x000C
	GEOMETRY_SHADER = 0x8DD9
	LAST_VERTEX_CONVENTION = 0x8E4E
	SYNC_FLUSH_COMMANDS_BIT = 0x00000001
	INT_SAMPLER_2D_MULTISAMPLE = 0x9109
	TRIANGLE_STRIP_ADJACENCY = 0x000D
	WAIT_FAILED = 0x911D
	SYNC_FENCE = 0x9116
	FIRST_VERTEX_CONVENTION = 0x8E4D
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS = 0x8DA8
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS = 0x8DE1
	TEXTURE_SAMPLES = 0x9106
	SYNC_STATUS = 0x9114
	SYNC_CONDITION = 0x9113
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY = 0x9105
	OBJECT_TYPE = 0x9112
	MAX_GEOMETRY_OUTPUT_VERTICES = 0x8DE0
	TEXTURE_2D_MULTISAMPLE = 0x9100
)
// VERSION_3_3
const (
	TEXTURE_SWIZZLE_B = 0x8E44
	ONE_MINUS_SRC1_COLOR = 0x88FA
	TEXTURE_SWIZZLE_G = 0x8E43
	TIMESTAMP = 0x8E28
	SAMPLER_BINDING = 0x8919
	TEXTURE_SWIZZLE_R = 0x8E42
	INT_2_10_10_10_REV = 0x8D9F
	TEXTURE_SWIZZLE_RGBA = 0x8E46
	TIME_ELAPSED = 0x88BF
	SRC1_COLOR = 0x88F9
	ANY_SAMPLES_PASSED = 0x8C2F
	VERTEX_ATTRIB_ARRAY_DIVISOR = 0x88FE
	RGB10_A2UI = 0x906F
	ONE_MINUS_SRC1_ALPHA = 0x88FB
	MAX_DUAL_SOURCE_DRAW_BUFFERS = 0x88FC
	TEXTURE_SWIZZLE_A = 0x8E45
)
// VERSION_4_0
const (
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS = 0x8E47
	DOUBLE_MAT3x2 = 0x8F4B
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS = 0x8E7F
	PATCH_DEFAULT_OUTER_LEVEL = 0x8E74
	DOUBLE_MAT3x4 = 0x8F4C
	TESS_GEN_VERTEX_ORDER = 0x8E78
	ISOLINES = 0x8E7A
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER = 0x84F0
	PATCH_VERTICES = 0x8E72
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS = 0x8E81
	PATCH_DEFAULT_INNER_LEVEL = 0x8E73
	MIN_SAMPLE_SHADING_VALUE = 0x8C37
	TESS_GEN_MODE = 0x8E76
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS = 0x8E82
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS = 0x8E8A
	SAMPLER_CUBE_MAP_ARRAY = 0x900C
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS = 0x8E1E
	TRANSFORM_FEEDBACK_BUFFER_PAUSED = TRANSFORM_FEEDBACK_PAUSED
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS = 0x8E83
	MAX_FRAGMENT_INTERPOLATION_OFFSET = 0x8E5C
	COMPATIBLE_SUBROUTINES = 0x8E4B
	MAX_SUBROUTINE_UNIFORM_LOCATIONS = 0x8DE8
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY = 0x900F
	GEOMETRY_SHADER_INVOCATIONS = 0x887F
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x84F1
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET = 0x8E5F
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS = 0x8E85
	MAX_TESS_EVALUATION_INPUT_COMPONENTS = 0x886D
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET = 0x8E5E
	TESS_GEN_POINT_MODE = 0x8E79
	DOUBLE_MAT4x3 = 0x8F4E
	MAX_GEOMETRY_SHADER_INVOCATIONS = 0x8E5A
	DOUBLE_MAT2x3 = 0x8F49
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE = TRANSFORM_FEEDBACK_ACTIVE
	ACTIVE_SUBROUTINE_UNIFORMS = 0x8DE6
	DOUBLE_VEC2 = 0x8FFC
	SAMPLE_SHADING = 0x8C36
	SAMPLER_CUBE_MAP_ARRAY_SHADOW = 0x900D
	MAX_TESS_CONTROL_UNIFORM_BLOCKS = 0x8E89
	MAX_TESS_GEN_LEVEL = 0x8E7E
	FRACTIONAL_ODD = 0x8E7B
	TRANSFORM_FEEDBACK = 0x8E22
	DRAW_INDIRECT_BUFFER_BINDING = 0x8F43
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS = 0x8E80
	TEXTURE_BINDING_CUBE_MAP_ARRAY = 0x900A
	TESS_GEN_SPACING = 0x8E77
	MAX_TESS_CONTROL_INPUT_COMPONENTS = 0x886C
	DOUBLE_VEC4 = 0x8FFE
	DOUBLE_MAT4 = 0x8F48
	INT_SAMPLER_CUBE_MAP_ARRAY = 0x900E
	TEXTURE_CUBE_MAP_ARRAY = 0x9009
	FRACTIONAL_EVEN = 0x8E7C
	PROXY_TEXTURE_CUBE_MAP_ARRAY = 0x900B
	ACTIVE_SUBROUTINES = 0x8DE5
	DOUBLE_MAT3 = 0x8F47
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH = 0x8E49
	TESS_CONTROL_OUTPUT_VERTICES = 0x8E75
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS = 0x8E1F
	DOUBLE_MAT4x2 = 0x8F4D
	MAX_TESS_PATCH_COMPONENTS = 0x8E84
	MAX_VERTEX_STREAMS = 0x8E71
	TESS_CONTROL_SHADER = 0x8E88
	DOUBLE_VEC3 = 0x8FFD
	PATCHES = 0x000E
	MAX_TRANSFORM_FEEDBACK_BUFFERS = 0x8E70
	TESS_EVALUATION_SHADER = 0x8E87
	MAX_PATCH_VERTICES = 0x8E7D
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS = 0x8E86
	FRAGMENT_INTERPOLATION_OFFSET_BITS = 0x8E5D
	NUM_COMPATIBLE_SUBROUTINES = 0x8E4A
	DOUBLE_MAT2x4 = 0x8F4A
	TRANSFORM_FEEDBACK_BINDING = 0x8E25
	MIN_FRAGMENT_INTERPOLATION_OFFSET = 0x8E5B
	DOUBLE_MAT2 = 0x8F46
	ACTIVE_SUBROUTINE_MAX_LENGTH = 0x8E48
	DRAW_INDIRECT_BUFFER = 0x8F3F
	MAX_SUBROUTINES = 0x8DE7
)
// VERSION_4_1
const (
	LOW_INT = 0x8DF3
	VERTEX_SHADER_BIT = 0x00000001
	FIXED = 0x140C
	MEDIUM_INT = 0x8DF4
	IMPLEMENTATION_COLOR_READ_FORMAT = 0x8B9B
	PROGRAM_SEPARABLE = 0x8258
	GEOMETRY_SHADER_BIT = 0x00000004
	LOW_FLOAT = 0x8DF0
	VIEWPORT_BOUNDS_RANGE = 0x825D
	MEDIUM_FLOAT = 0x8DF1
	MAX_VARYING_VECTORS = 0x8DFC
	SHADER_COMPILER = 0x8DFA
	PROGRAM_BINARY_FORMATS = 0x87FF
	PROGRAM_BINARY_LENGTH = 0x8741
	HIGH_FLOAT = 0x8DF2
	TESS_EVALUATION_SHADER_BIT = 0x00000010
	ALL_SHADER_BITS = 0xFFFFFFFF
	VIEWPORT_SUBPIXEL_BITS = 0x825C
	IMPLEMENTATION_COLOR_READ_TYPE = 0x8B9A
	RGB565 = 0x8D62
	MAX_FRAGMENT_UNIFORM_VECTORS = 0x8DFD
	ACTIVE_PROGRAM = 0x8259
	UNDEFINED_VERTEX = 0x8260
	MAX_VIEWPORTS = 0x825B
	HIGH_INT = 0x8DF5
	MAX_VERTEX_UNIFORM_VECTORS = 0x8DFB
	FRAGMENT_SHADER_BIT = 0x00000002
	NUM_PROGRAM_BINARY_FORMATS = 0x87FE
	PROGRAM_PIPELINE_BINDING = 0x825A
	VIEWPORT_INDEX_PROVOKING_VERTEX = 0x825F
	TESS_CONTROL_SHADER_BIT = 0x00000008
	PROGRAM_BINARY_RETRIEVABLE_HINT = 0x8257
	LAYER_PROVOKING_VERTEX = 0x825E
	NUM_SHADER_BINARY_FORMATS = 0x8DF9
)
// VERSION_4_2
const (
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS = 0x92D0
	IMAGE_BINDING_LEVEL = 0x8F3B
	IMAGE_FORMAT_COMPATIBILITY_TYPE = 0x90C7
	ATOMIC_COUNTER_BUFFER = 0x92C0
	TEXTURE_FETCH_BARRIER_BIT = 0x00000008
	INT_IMAGE_2D_ARRAY = 0x905E
	PACK_COMPRESSED_BLOCK_DEPTH = 0x912D
	MAX_ATOMIC_COUNTER_BUFFER_SIZE = 0x92D8
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER = 0x92CB
	UNPACK_COMPRESSED_BLOCK_WIDTH = 0x9127
	ACTIVE_ATOMIC_COUNTER_BUFFERS = 0x92D9
	MAX_FRAGMENT_ATOMIC_COUNTERS = 0x92D6
	TEXTURE_UPDATE_BARRIER_BIT = 0x00000100
	UNPACK_COMPRESSED_BLOCK_SIZE = 0x912A
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY = 0x906C
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS = 0x90CC
	IMAGE_BINDING_LAYERED = 0x8F3C
	ELEMENT_ARRAY_BARRIER_BIT = 0x00000002
	UNSIGNED_INT_IMAGE_CUBE = 0x9066
	BUFFER_UPDATE_BARRIER_BIT = 0x00000200
	IMAGE_1D = 0x904C
	IMAGE_2D_RECT = 0x904F
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES = 0x92C6
	ATOMIC_COUNTER_BUFFER_SIZE = 0x92C3
	PACK_COMPRESSED_BLOCK_SIZE = 0x912E
	MAX_TESS_CONTROL_ATOMIC_COUNTERS = 0x92D3
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS = 0x92C5
	ALL_BARRIER_BITS = 0xFFFFFFFF
	IMAGE_BINDING_ACCESS = 0x8F3E
	IMAGE_BINDING_NAME = 0x8F3A
	MAX_VERTEX_IMAGE_UNIFORMS = 0x90CA
	IMAGE_3D = 0x904E
	ATOMIC_COUNTER_BARRIER_BIT = 0x00001000
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS = 0x92CD
	MAX_COMBINED_ATOMIC_COUNTERS = 0x92D7
	INT_IMAGE_2D = 0x9058
	UNSIGNED_INT_IMAGE_2D_RECT = 0x9065
	IMAGE_BUFFER = 0x9051
	PACK_COMPRESSED_BLOCK_HEIGHT = 0x912C
	TEXTURE_IMMUTABLE_FORMAT = 0x912F
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS = 0x90C9
	IMAGE_2D_MULTISAMPLE_ARRAY = 0x9056
	UNSIGNED_INT_IMAGE_2D = 0x9063
	ATOMIC_COUNTER_BUFFER_START = 0x92C2
	INT_IMAGE_CUBE_MAP_ARRAY = 0x905F
	COMMAND_BARRIER_BIT = 0x00000040
	IMAGE_1D_ARRAY = 0x9052
	UNSIGNED_INT_ATOMIC_COUNTER = 0x92DB
	ATOMIC_COUNTER_BUFFER_BINDING = 0x92C1
	UNSIGNED_INT_IMAGE_3D = 0x9064
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS = 0x92CE
	IMAGE_CUBE_MAP_ARRAY = 0x9054
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS = 0x92D4
	INT_IMAGE_1D_ARRAY = 0x905D
	IMAGE_BINDING_FORMAT = 0x906E
	IMAGE_BINDING_LAYER = 0x8F3D
	UNPACK_COMPRESSED_BLOCK_HEIGHT = 0x9128
	UNIFORM_BARRIER_BIT = 0x00000004
	INT_IMAGE_CUBE = 0x905B
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER = 0x92C8
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS = 0x8F39
	INT_IMAGE_BUFFER = 0x905C
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER = 0x92CA
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT = 0x00000001
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY = 0x906A
	MIN_MAP_BUFFER_ALIGNMENT = 0x90BC
	MAX_COMBINED_IMAGE_UNIFORMS = 0x90CF
	MAX_GEOMETRY_IMAGE_UNIFORMS = 0x90CD
	SHADER_IMAGE_ACCESS_BARRIER_BIT = 0x00000020
	MAX_VERTEX_ATOMIC_COUNTERS = 0x92D2
	IMAGE_2D_MULTISAMPLE = 0x9055
	TRANSFORM_FEEDBACK_BARRIER_BIT = 0x00000800
	IMAGE_2D_ARRAY = 0x9053
	IMAGE_2D = 0x904D
	UNSIGNED_INT_IMAGE_2D_ARRAY = 0x9069
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX = 0x92DA
	INT_IMAGE_3D = 0x9059
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	INT_IMAGE_1D = 0x9057
	MAX_GEOMETRY_ATOMIC_COUNTERS = 0x92D5
	UNSIGNED_INT_IMAGE_1D_ARRAY = 0x9068
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS = 0x92D1
	INT_IMAGE_2D_MULTISAMPLE = 0x9060
	MAX_FRAGMENT_IMAGE_UNIFORMS = 0x90CE
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER = 0x92C7
	IMAGE_CUBE = 0x9050
	MAX_TESS_CONTROL_IMAGE_UNIFORMS = 0x90CB
	NUM_SAMPLE_COUNTS = 0x9380
	UNSIGNED_INT_IMAGE_1D = 0x9062
	UNSIGNED_INT_IMAGE_BUFFER = 0x9067
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE = 0x906B
	PIXEL_BUFFER_BARRIER_BIT = 0x00000080
	FRAMEBUFFER_BARRIER_BIT = 0x00000400
	INT_IMAGE_2D_MULTISAMPLE_ARRAY = 0x9061
	INT_IMAGE_2D_RECT = 0x905A
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE = 0x90C8
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS = 0x92DC
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS = 0x92CF
	UNPACK_COMPRESSED_BLOCK_DEPTH = 0x9129
	PACK_COMPRESSED_BLOCK_WIDTH = 0x912B
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS = 0x92CC
	ATOMIC_COUNTER_BUFFER_DATA_SIZE = 0x92C4
	MAX_IMAGE_UNITS = 0x8F38
	MAX_IMAGE_SAMPLES = 0x906D
)
// VERSION_1_0

// http://www.opengl.org/sdk/docs/man4/xhtml/glCullFace.xml
func CullFace(mode Enum)  {
	C.goglCullFace((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFrontFace.xml
func FrontFace(mode Enum)  {
	C.goglFrontFace((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glHint.xml
func Hint(target Enum, mode Enum)  {
	C.goglHint((C.GLenum)(target), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glLineWidth.xml
func LineWidth(width Float)  {
	C.goglLineWidth((C.GLfloat)(width))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPointSize.xml
func PointSize(size Float)  {
	C.goglPointSize((C.GLfloat)(size))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPolygonMode.xml
func PolygonMode(face Enum, mode Enum)  {
	C.goglPolygonMode((C.GLenum)(face), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glScissor.xml
func Scissor(x Int, y Int, width Sizei, height Sizei)  {
	C.goglScissor((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexParameterf.xml
func TexParameterf(target Enum, pname Enum, param Float)  {
	C.goglTexParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexParameterfv.xml
func TexParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexParameteri.xml
func TexParameteri(target Enum, pname Enum, param Int)  {
	C.goglTexParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexParameteriv.xml
func TexParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexImage1D.xml
func TexImage1D(target Enum, level Int, internalformat Int, width Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexImage2D.xml
func TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawBuffer.xml
func DrawBuffer(mode Enum)  {
	C.goglDrawBuffer((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClear.xml
func Clear(mask Bitfield)  {
	C.goglClear((C.GLbitfield)(mask))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClearColor.xml
func ClearColor(red Float, green Float, blue Float, alpha Float)  {
	C.goglClearColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClearStencil.xml
func ClearStencil(s Int)  {
	C.goglClearStencil((C.GLint)(s))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClearDepth.xml
func ClearDepth(depth Double)  {
	C.goglClearDepth((C.GLdouble)(depth))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glStencilMask.xml
func StencilMask(mask Uint)  {
	C.goglStencilMask((C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glColorMask.xml
func ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean)  {
	C.goglColorMask((C.GLboolean)(red), (C.GLboolean)(green), (C.GLboolean)(blue), (C.GLboolean)(alpha))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDepthMask.xml
func DepthMask(flag Boolean)  {
	C.goglDepthMask((C.GLboolean)(flag))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDisable.xml
func Disable(cap Enum)  {
	C.goglDisable((C.GLenum)(cap))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glEnable.xml
func Enable(cap Enum)  {
	C.goglEnable((C.GLenum)(cap))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFinish.xml
func Finish()  {
	C.goglFinish()
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFlush.xml
func Flush()  {
	C.goglFlush()
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBlendFunc.xml
func BlendFunc(sfactor Enum, dfactor Enum)  {
	C.goglBlendFunc((C.GLenum)(sfactor), (C.GLenum)(dfactor))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glLogicOp.xml
func LogicOp(opcode Enum)  {
	C.goglLogicOp((C.GLenum)(opcode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glStencilFunc.xml
func StencilFunc(func_ Enum, ref Int, mask Uint)  {
	C.goglStencilFunc((C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glStencilOp.xml
func StencilOp(fail Enum, zfail Enum, zpass Enum)  {
	C.goglStencilOp((C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDepthFunc.xml
func DepthFunc(func_ Enum)  {
	C.goglDepthFunc((C.GLenum)(func_))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPixelStoref.xml
func PixelStoref(pname Enum, param Float)  {
	C.goglPixelStoref((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPixelStorei.xml
func PixelStorei(pname Enum, param Int)  {
	C.goglPixelStorei((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glReadBuffer.xml
func ReadBuffer(mode Enum)  {
	C.goglReadBuffer((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glReadPixels.xml
func ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglReadPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetBooleanv.xml
func GetBooleanv(pname Enum, params *Boolean)  {
	C.goglGetBooleanv((C.GLenum)(pname), (*C.GLboolean)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetDoublev.xml
func GetDoublev(pname Enum, params *Double)  {
	C.goglGetDoublev((C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetError.xml
func GetError() Enum {
	return (Enum)(C.goglGetError())
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetFloatv.xml
func GetFloatv(pname Enum, params *Float)  {
	C.goglGetFloatv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetIntegerv.xml
func GetIntegerv(pname Enum, params *Int)  {
	C.goglGetIntegerv((C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetString.xml
func GetString(name Enum) *Ubyte {
	return (*Ubyte)(C.goglGetString((C.GLenum)(name)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetTexImage.xml
func GetTexImage(target Enum, level Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglGetTexImage((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetTexParameterfv.xml
func GetTexParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetTexParameteriv.xml
func GetTexParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetTexLevelParameterfv.xml
func GetTexLevelParameterfv(target Enum, level Int, pname Enum, params *Float)  {
	C.goglGetTexLevelParameterfv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetTexLevelParameteriv.xml
func GetTexLevelParameteriv(target Enum, level Int, pname Enum, params *Int)  {
	C.goglGetTexLevelParameteriv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsEnabled.xml
func IsEnabled(cap Enum) Boolean {
	return (Boolean)(C.goglIsEnabled((C.GLenum)(cap)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDepthRange.xml
func DepthRange(near_ Double, far_ Double)  {
	C.goglDepthRange((C.GLdouble)(near_), (C.GLdouble)(far_))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glViewport.xml
func Viewport(x Int, y Int, width Sizei, height Sizei)  {
	C.goglViewport((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_1

// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawArrays.xml
func DrawArrays(mode Enum, first Int, count Sizei)  {
	C.goglDrawArrays((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawElements.xml
func DrawElements(mode Enum, count Sizei, type_ Enum, indices Pointer)  {
	C.goglDrawElements((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetPointerv.xml
func GetPointerv(pname Enum, params *Pointer)  {
	C.goglGetPointerv((C.GLenum)(pname), (*unsafe.Pointer)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPolygonOffset.xml
func PolygonOffset(factor Float, units Float)  {
	C.goglPolygonOffset((C.GLfloat)(factor), (C.GLfloat)(units))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCopyTexImage1D.xml
func CopyTexImage1D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, border Int)  {
	C.goglCopyTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCopyTexImage2D.xml
func CopyTexImage2D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, height Sizei, border Int)  {
	C.goglCopyTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCopyTexSubImage1D.xml
func CopyTexSubImage1D(target Enum, level Int, xoffset Int, x Int, y Int, width Sizei)  {
	C.goglCopyTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCopyTexSubImage2D.xml
func CopyTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, x Int, y Int, width Sizei, height Sizei)  {
	C.goglCopyTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexSubImage1D.xml
func TexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexSubImage2D.xml
func TexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindTexture.xml
func BindTexture(target Enum, texture Uint)  {
	C.goglBindTexture((C.GLenum)(target), (C.GLuint)(texture))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteTextures.xml
func DeleteTextures(n Sizei, textures *Uint)  {
	C.goglDeleteTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGenTextures.xml
func GenTextures(n Sizei, textures *Uint)  {
	C.goglGenTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsTexture.xml
func IsTexture(texture Uint) Boolean {
	return (Boolean)(C.goglIsTexture((C.GLuint)(texture)))
}
// VERSION_1_2

// http://www.opengl.org/sdk/docs/man4/xhtml/glBlendColor.xml
func BlendColor(red Float, green Float, blue Float, alpha Float)  {
	C.goglBlendColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBlendEquation.xml
func BlendEquation(mode Enum)  {
	C.goglBlendEquation((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawRangeElements.xml
func DrawRangeElements(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Pointer)  {
	C.goglDrawRangeElements((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexImage3D.xml
func TexImage3D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, depth Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexSubImage3D.xml
func TexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCopyTexSubImage3D.xml
func CopyTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, x Int, y Int, width Sizei, height Sizei)  {
	C.goglCopyTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_3

// http://www.opengl.org/sdk/docs/man4/xhtml/glActiveTexture.xml
func ActiveTexture(texture Enum)  {
	C.goglActiveTexture((C.GLenum)(texture))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glSampleCoverage.xml
func SampleCoverage(value Float, invert Boolean)  {
	C.goglSampleCoverage((C.GLfloat)(value), (C.GLboolean)(invert))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCompressedTexImage3D.xml
func CompressedTexImage3D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, depth Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCompressedTexImage2D.xml
func CompressedTexImage2D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCompressedTexImage1D.xml
func CompressedTexImage1D(target Enum, level Int, internalformat Enum, width Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCompressedTexSubImage3D.xml
func CompressedTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCompressedTexSubImage2D.xml
func CompressedTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCompressedTexSubImage1D.xml
func CompressedTexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetCompressedTexImage.xml
func GetCompressedTexImage(target Enum, level Int, img Pointer)  {
	C.goglGetCompressedTexImage((C.GLenum)(target), (C.GLint)(level), (unsafe.Pointer)(img))
}
// VERSION_1_4

// http://www.opengl.org/sdk/docs/man4/xhtml/glBlendFuncSeparate.xml
func BlendFuncSeparate(sfactorRGB Enum, dfactorRGB Enum, sfactorAlpha Enum, dfactorAlpha Enum)  {
	C.goglBlendFuncSeparate((C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiDrawArrays.xml
func MultiDrawArrays(mode Enum, first *Int, count *Sizei, primcount Sizei)  {
	C.goglMultiDrawArrays((C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(primcount))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiDrawElements.xml
func MultiDrawElements(mode Enum, count *Sizei, type_ Enum, indices *Pointer, primcount Sizei)  {
	C.goglMultiDrawElements((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(primcount))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPointParameterf.xml
func PointParameterf(pname Enum, param Float)  {
	C.goglPointParameterf((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPointParameterfv.xml
func PointParameterfv(pname Enum, params *Float)  {
	C.goglPointParameterfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPointParameteri.xml
func PointParameteri(pname Enum, param Int)  {
	C.goglPointParameteri((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPointParameteriv.xml
func PointParameteriv(pname Enum, params *Int)  {
	C.goglPointParameteriv((C.GLenum)(pname), (*C.GLint)(params))
}
// VERSION_1_5

// http://www.opengl.org/sdk/docs/man4/xhtml/glGenQueries.xml
func GenQueries(n Sizei, ids *Uint)  {
	C.goglGenQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteQueries.xml
func DeleteQueries(n Sizei, ids *Uint)  {
	C.goglDeleteQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsQuery.xml
func IsQuery(id Uint) Boolean {
	return (Boolean)(C.goglIsQuery((C.GLuint)(id)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBeginQuery.xml
func BeginQuery(target Enum, id Uint)  {
	C.goglBeginQuery((C.GLenum)(target), (C.GLuint)(id))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glEndQuery.xml
func EndQuery(target Enum)  {
	C.goglEndQuery((C.GLenum)(target))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetQueryiv.xml
func GetQueryiv(target Enum, pname Enum, params *Int)  {
	C.goglGetQueryiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetQueryObjectiv.xml
func GetQueryObjectiv(id Uint, pname Enum, params *Int)  {
	C.goglGetQueryObjectiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetQueryObjectuiv.xml
func GetQueryObjectuiv(id Uint, pname Enum, params *Uint)  {
	C.goglGetQueryObjectuiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindBuffer.xml
func BindBuffer(target Enum, buffer Uint)  {
	C.goglBindBuffer((C.GLenum)(target), (C.GLuint)(buffer))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteBuffers.xml
func DeleteBuffers(n Sizei, buffers *Uint)  {
	C.goglDeleteBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGenBuffers.xml
func GenBuffers(n Sizei, buffers *Uint)  {
	C.goglGenBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsBuffer.xml
func IsBuffer(buffer Uint) Boolean {
	return (Boolean)(C.goglIsBuffer((C.GLuint)(buffer)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBufferData.xml
func BufferData(target Enum, size Sizeiptr, data Pointer, usage Enum)  {
	C.goglBufferData((C.GLenum)(target), (C.GLsizeiptr)(size), (unsafe.Pointer)(data), (C.GLenum)(usage))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBufferSubData.xml
func BufferSubData(target Enum, offset Intptr, size Sizeiptr, data Pointer)  {
	C.goglBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetBufferSubData.xml
func GetBufferSubData(target Enum, offset Intptr, size Sizeiptr, data Pointer)  {
	C.goglGetBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMapBuffer.xml
func MapBuffer(target Enum, access Enum) Pointer {
	return (Pointer)(C.goglMapBuffer((C.GLenum)(target), (C.GLenum)(access)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUnmapBuffer.xml
func UnmapBuffer(target Enum) Boolean {
	return (Boolean)(C.goglUnmapBuffer((C.GLenum)(target)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetBufferParameteriv.xml
func GetBufferParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetBufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetBufferPointerv.xml
func GetBufferPointerv(target Enum, pname Enum, params *Pointer)  {
	C.goglGetBufferPointerv((C.GLenum)(target), (C.GLenum)(pname), (*unsafe.Pointer)(params))
}
// VERSION_2_0

// http://www.opengl.org/sdk/docs/man4/xhtml/glBlendEquationSeparate.xml
func BlendEquationSeparate(modeRGB Enum, modeAlpha Enum)  {
	C.goglBlendEquationSeparate((C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawBuffers.xml
func DrawBuffers(n Sizei, bufs *Enum)  {
	C.goglDrawBuffers((C.GLsizei)(n), (*C.GLenum)(bufs))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glStencilOpSeparate.xml
func StencilOpSeparate(face Enum, sfail Enum, dpfail Enum, dppass Enum)  {
	C.goglStencilOpSeparate((C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glStencilFuncSeparate.xml
func StencilFuncSeparate(face Enum, func_ Enum, ref Int, mask Uint)  {
	C.goglStencilFuncSeparate((C.GLenum)(face), (C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glStencilMaskSeparate.xml
func StencilMaskSeparate(face Enum, mask Uint)  {
	C.goglStencilMaskSeparate((C.GLenum)(face), (C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glAttachShader.xml
func AttachShader(program Uint, shader Uint)  {
	C.goglAttachShader((C.GLuint)(program), (C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindAttribLocation.xml
func BindAttribLocation(program Uint, index Uint, name *Char)  {
	C.goglBindAttribLocation((C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCompileShader.xml
func CompileShader(shader Uint)  {
	C.goglCompileShader((C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCreateProgram.xml
func CreateProgram() Uint {
	return (Uint)(C.goglCreateProgram())
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCreateShader.xml
func CreateShader(type_ Enum) Uint {
	return (Uint)(C.goglCreateShader((C.GLenum)(type_)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteProgram.xml
func DeleteProgram(program Uint)  {
	C.goglDeleteProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteShader.xml
func DeleteShader(shader Uint)  {
	C.goglDeleteShader((C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDetachShader.xml
func DetachShader(program Uint, shader Uint)  {
	C.goglDetachShader((C.GLuint)(program), (C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDisableVertexAttribArray.xml
func DisableVertexAttribArray(index Uint)  {
	C.goglDisableVertexAttribArray((C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glEnableVertexAttribArray.xml
func EnableVertexAttribArray(index Uint)  {
	C.goglEnableVertexAttribArray((C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetActiveAttrib.xml
func GetActiveAttrib(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)  {
	C.goglGetActiveAttrib((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetActiveUniform.xml
func GetActiveUniform(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)  {
	C.goglGetActiveUniform((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetAttachedShaders.xml
func GetAttachedShaders(program Uint, maxCount Sizei, count *Sizei, obj *Uint)  {
	C.goglGetAttachedShaders((C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(count), (*C.GLuint)(obj))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetAttribLocation.xml
func GetAttribLocation(program Uint, name *Char) Int {
	return (Int)(C.goglGetAttribLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetProgramiv.xml
func GetProgramiv(program Uint, pname Enum, params *Int)  {
	C.goglGetProgramiv((C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetProgramInfoLog.xml
func GetProgramInfoLog(program Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.goglGetProgramInfoLog((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetShaderiv.xml
func GetShaderiv(shader Uint, pname Enum, params *Int)  {
	C.goglGetShaderiv((C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetShaderInfoLog.xml
func GetShaderInfoLog(shader Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.goglGetShaderInfoLog((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetShaderSource.xml
func GetShaderSource(shader Uint, bufSize Sizei, length *Sizei, source *Char)  {
	C.goglGetShaderSource((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(source))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetUniformLocation.xml
func GetUniformLocation(program Uint, name *Char) Int {
	return (Int)(C.goglGetUniformLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetUniformfv.xml
func GetUniformfv(program Uint, location Int, params *Float)  {
	C.goglGetUniformfv((C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetUniformiv.xml
func GetUniformiv(program Uint, location Int, params *Int)  {
	C.goglGetUniformiv((C.GLuint)(program), (C.GLint)(location), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetVertexAttribdv.xml
func GetVertexAttribdv(index Uint, pname Enum, params *Double)  {
	C.goglGetVertexAttribdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetVertexAttribfv.xml
func GetVertexAttribfv(index Uint, pname Enum, params *Float)  {
	C.goglGetVertexAttribfv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetVertexAttribiv.xml
func GetVertexAttribiv(index Uint, pname Enum, params *Int)  {
	C.goglGetVertexAttribiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetVertexAttribPointerv.xml
func GetVertexAttribPointerv(index Uint, pname Enum, pointer *Pointer)  {
	C.goglGetVertexAttribPointerv((C.GLuint)(index), (C.GLenum)(pname), (*unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsProgram.xml
func IsProgram(program Uint) Boolean {
	return (Boolean)(C.goglIsProgram((C.GLuint)(program)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsShader.xml
func IsShader(shader Uint) Boolean {
	return (Boolean)(C.goglIsShader((C.GLuint)(shader)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glLinkProgram.xml
func LinkProgram(program Uint)  {
	C.goglLinkProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glShaderSource.xml
func ShaderSource(shader Uint, count Sizei, string_ **Char, length *Int)  {
	C.goglShaderSource((C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(string_)), (*C.GLint)(length))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUseProgram.xml
func UseProgram(program Uint)  {
	C.goglUseProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform1f.xml
func Uniform1f(location Int, v0 Float)  {
	C.goglUniform1f((C.GLint)(location), (C.GLfloat)(v0))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform2f.xml
func Uniform2f(location Int, v0 Float, v1 Float)  {
	C.goglUniform2f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform3f.xml
func Uniform3f(location Int, v0 Float, v1 Float, v2 Float)  {
	C.goglUniform3f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform4f.xml
func Uniform4f(location Int, v0 Float, v1 Float, v2 Float, v3 Float)  {
	C.goglUniform4f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform1i.xml
func Uniform1i(location Int, v0 Int)  {
	C.goglUniform1i((C.GLint)(location), (C.GLint)(v0))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform2i.xml
func Uniform2i(location Int, v0 Int, v1 Int)  {
	C.goglUniform2i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform3i.xml
func Uniform3i(location Int, v0 Int, v1 Int, v2 Int)  {
	C.goglUniform3i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform4i.xml
func Uniform4i(location Int, v0 Int, v1 Int, v2 Int, v3 Int)  {
	C.goglUniform4i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform1fv.xml
func Uniform1fv(location Int, count Sizei, value *Float)  {
	C.goglUniform1fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform2fv.xml
func Uniform2fv(location Int, count Sizei, value *Float)  {
	C.goglUniform2fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform3fv.xml
func Uniform3fv(location Int, count Sizei, value *Float)  {
	C.goglUniform3fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform4fv.xml
func Uniform4fv(location Int, count Sizei, value *Float)  {
	C.goglUniform4fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform1iv.xml
func Uniform1iv(location Int, count Sizei, value *Int)  {
	C.goglUniform1iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform2iv.xml
func Uniform2iv(location Int, count Sizei, value *Int)  {
	C.goglUniform2iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform3iv.xml
func Uniform3iv(location Int, count Sizei, value *Int)  {
	C.goglUniform3iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform4iv.xml
func Uniform4iv(location Int, count Sizei, value *Int)  {
	C.goglUniform4iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix2fv.xml
func UniformMatrix2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix3fv.xml
func UniformMatrix3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix4fv.xml
func UniformMatrix4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glValidateProgram.xml
func ValidateProgram(program Uint)  {
	C.goglValidateProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib1d.xml
func VertexAttrib1d(index Uint, x Double)  {
	C.goglVertexAttrib1d((C.GLuint)(index), (C.GLdouble)(x))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib1dv.xml
func VertexAttrib1dv(index Uint, v *Double)  {
	C.goglVertexAttrib1dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib1f.xml
func VertexAttrib1f(index Uint, x Float)  {
	C.goglVertexAttrib1f((C.GLuint)(index), (C.GLfloat)(x))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib1fv.xml
func VertexAttrib1fv(index Uint, v *Float)  {
	C.goglVertexAttrib1fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib1s.xml
func VertexAttrib1s(index Uint, x Short)  {
	C.goglVertexAttrib1s((C.GLuint)(index), (C.GLshort)(x))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib1sv.xml
func VertexAttrib1sv(index Uint, v *Short)  {
	C.goglVertexAttrib1sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib2d.xml
func VertexAttrib2d(index Uint, x Double, y Double)  {
	C.goglVertexAttrib2d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib2dv.xml
func VertexAttrib2dv(index Uint, v *Double)  {
	C.goglVertexAttrib2dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib2f.xml
func VertexAttrib2f(index Uint, x Float, y Float)  {
	C.goglVertexAttrib2f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib2fv.xml
func VertexAttrib2fv(index Uint, v *Float)  {
	C.goglVertexAttrib2fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib2s.xml
func VertexAttrib2s(index Uint, x Short, y Short)  {
	C.goglVertexAttrib2s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib2sv.xml
func VertexAttrib2sv(index Uint, v *Short)  {
	C.goglVertexAttrib2sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib3d.xml
func VertexAttrib3d(index Uint, x Double, y Double, z Double)  {
	C.goglVertexAttrib3d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib3dv.xml
func VertexAttrib3dv(index Uint, v *Double)  {
	C.goglVertexAttrib3dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib3f.xml
func VertexAttrib3f(index Uint, x Float, y Float, z Float)  {
	C.goglVertexAttrib3f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib3fv.xml
func VertexAttrib3fv(index Uint, v *Float)  {
	C.goglVertexAttrib3fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib3s.xml
func VertexAttrib3s(index Uint, x Short, y Short, z Short)  {
	C.goglVertexAttrib3s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib3sv.xml
func VertexAttrib3sv(index Uint, v *Short)  {
	C.goglVertexAttrib3sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4Nbv.xml
func VertexAttrib4Nbv(index Uint, v *Byte)  {
	C.goglVertexAttrib4Nbv((C.GLuint)(index), (*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4Niv.xml
func VertexAttrib4Niv(index Uint, v *Int)  {
	C.goglVertexAttrib4Niv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4Nsv.xml
func VertexAttrib4Nsv(index Uint, v *Short)  {
	C.goglVertexAttrib4Nsv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4Nub.xml
func VertexAttrib4Nub(index Uint, x Ubyte, y Ubyte, z Ubyte, w Ubyte)  {
	C.goglVertexAttrib4Nub((C.GLuint)(index), (C.GLubyte)(x), (C.GLubyte)(y), (C.GLubyte)(z), (C.GLubyte)(w))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4Nubv.xml
func VertexAttrib4Nubv(index Uint, v *Ubyte)  {
	C.goglVertexAttrib4Nubv((C.GLuint)(index), (*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4Nuiv.xml
func VertexAttrib4Nuiv(index Uint, v *Uint)  {
	C.goglVertexAttrib4Nuiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4Nusv.xml
func VertexAttrib4Nusv(index Uint, v *Ushort)  {
	C.goglVertexAttrib4Nusv((C.GLuint)(index), (*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4bv.xml
func VertexAttrib4bv(index Uint, v *Byte)  {
	C.goglVertexAttrib4bv((C.GLuint)(index), (*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4d.xml
func VertexAttrib4d(index Uint, x Double, y Double, z Double, w Double)  {
	C.goglVertexAttrib4d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4dv.xml
func VertexAttrib4dv(index Uint, v *Double)  {
	C.goglVertexAttrib4dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4f.xml
func VertexAttrib4f(index Uint, x Float, y Float, z Float, w Float)  {
	C.goglVertexAttrib4f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4fv.xml
func VertexAttrib4fv(index Uint, v *Float)  {
	C.goglVertexAttrib4fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4iv.xml
func VertexAttrib4iv(index Uint, v *Int)  {
	C.goglVertexAttrib4iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4s.xml
func VertexAttrib4s(index Uint, x Short, y Short, z Short, w Short)  {
	C.goglVertexAttrib4s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4sv.xml
func VertexAttrib4sv(index Uint, v *Short)  {
	C.goglVertexAttrib4sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4ubv.xml
func VertexAttrib4ubv(index Uint, v *Ubyte)  {
	C.goglVertexAttrib4ubv((C.GLuint)(index), (*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4uiv.xml
func VertexAttrib4uiv(index Uint, v *Uint)  {
	C.goglVertexAttrib4uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttrib4usv.xml
func VertexAttrib4usv(index Uint, v *Ushort)  {
	C.goglVertexAttrib4usv((C.GLuint)(index), (*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribPointer.xml
func VertexAttribPointer(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, pointer Pointer)  {
	C.goglVertexAttribPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// VERSION_2_1

// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix2x3fv.xml
func UniformMatrix2x3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix3x2fv.xml
func UniformMatrix3x2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix2x4fv.xml
func UniformMatrix2x4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix4x2fv.xml
func UniformMatrix4x2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix3x4fv.xml
func UniformMatrix3x4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix4x3fv.xml
func UniformMatrix4x3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// VERSION_3_0

// http://www.opengl.org/sdk/docs/man4/xhtml/glColorMaski.xml
func ColorMaski(index Uint, r Boolean, g Boolean, b Boolean, a Boolean)  {
	C.goglColorMaski((C.GLuint)(index), (C.GLboolean)(r), (C.GLboolean)(g), (C.GLboolean)(b), (C.GLboolean)(a))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetBooleani_v.xml
func GetBooleani_v(target Enum, index Uint, data *Boolean)  {
	C.goglGetBooleani_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLboolean)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetIntegeri_v.xml
func GetIntegeri_v(target Enum, index Uint, data *Int)  {
	C.goglGetIntegeri_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glEnablei.xml
func Enablei(target Enum, index Uint)  {
	C.goglEnablei((C.GLenum)(target), (C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDisablei.xml
func Disablei(target Enum, index Uint)  {
	C.goglDisablei((C.GLenum)(target), (C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsEnabledi.xml
func IsEnabledi(target Enum, index Uint) Boolean {
	return (Boolean)(C.goglIsEnabledi((C.GLenum)(target), (C.GLuint)(index)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBeginTransformFeedback.xml
func BeginTransformFeedback(primitiveMode Enum)  {
	C.goglBeginTransformFeedback((C.GLenum)(primitiveMode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glEndTransformFeedback.xml
func EndTransformFeedback()  {
	C.goglEndTransformFeedback()
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindBufferRange.xml
func BindBufferRange(target Enum, index Uint, buffer Uint, offset Intptr, size Sizeiptr)  {
	C.goglBindBufferRange((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindBufferBase.xml
func BindBufferBase(target Enum, index Uint, buffer Uint)  {
	C.goglBindBufferBase((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTransformFeedbackVaryings.xml
func TransformFeedbackVaryings(program Uint, count Sizei, varyings **Char, bufferMode Enum)  {
	C.goglTransformFeedbackVaryings((C.GLuint)(program), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(varyings)), (C.GLenum)(bufferMode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetTransformFeedbackVarying.xml
func GetTransformFeedbackVarying(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Sizei, type_ *Enum, name *Char)  {
	C.goglGetTransformFeedbackVarying((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLsizei)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClampColor.xml
func ClampColor(target Enum, clamp Enum)  {
	C.goglClampColor((C.GLenum)(target), (C.GLenum)(clamp))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBeginConditionalRender.xml
func BeginConditionalRender(id Uint, mode Enum)  {
	C.goglBeginConditionalRender((C.GLuint)(id), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glEndConditionalRender.xml
func EndConditionalRender()  {
	C.goglEndConditionalRender()
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribIPointer.xml
func VertexAttribIPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglVertexAttribIPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetVertexAttribIiv.xml
func GetVertexAttribIiv(index Uint, pname Enum, params *Int)  {
	C.goglGetVertexAttribIiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetVertexAttribIuiv.xml
func GetVertexAttribIuiv(index Uint, pname Enum, params *Uint)  {
	C.goglGetVertexAttribIuiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI1i.xml
func VertexAttribI1i(index Uint, x Int)  {
	C.goglVertexAttribI1i((C.GLuint)(index), (C.GLint)(x))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI2i.xml
func VertexAttribI2i(index Uint, x Int, y Int)  {
	C.goglVertexAttribI2i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI3i.xml
func VertexAttribI3i(index Uint, x Int, y Int, z Int)  {
	C.goglVertexAttribI3i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI4i.xml
func VertexAttribI4i(index Uint, x Int, y Int, z Int, w Int)  {
	C.goglVertexAttribI4i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI1ui.xml
func VertexAttribI1ui(index Uint, x Uint)  {
	C.goglVertexAttribI1ui((C.GLuint)(index), (C.GLuint)(x))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI2ui.xml
func VertexAttribI2ui(index Uint, x Uint, y Uint)  {
	C.goglVertexAttribI2ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI3ui.xml
func VertexAttribI3ui(index Uint, x Uint, y Uint, z Uint)  {
	C.goglVertexAttribI3ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI4ui.xml
func VertexAttribI4ui(index Uint, x Uint, y Uint, z Uint, w Uint)  {
	C.goglVertexAttribI4ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI1iv.xml
func VertexAttribI1iv(index Uint, v *Int)  {
	C.goglVertexAttribI1iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI2iv.xml
func VertexAttribI2iv(index Uint, v *Int)  {
	C.goglVertexAttribI2iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI3iv.xml
func VertexAttribI3iv(index Uint, v *Int)  {
	C.goglVertexAttribI3iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI4iv.xml
func VertexAttribI4iv(index Uint, v *Int)  {
	C.goglVertexAttribI4iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI1uiv.xml
func VertexAttribI1uiv(index Uint, v *Uint)  {
	C.goglVertexAttribI1uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI2uiv.xml
func VertexAttribI2uiv(index Uint, v *Uint)  {
	C.goglVertexAttribI2uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI3uiv.xml
func VertexAttribI3uiv(index Uint, v *Uint)  {
	C.goglVertexAttribI3uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI4uiv.xml
func VertexAttribI4uiv(index Uint, v *Uint)  {
	C.goglVertexAttribI4uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI4bv.xml
func VertexAttribI4bv(index Uint, v *Byte)  {
	C.goglVertexAttribI4bv((C.GLuint)(index), (*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI4sv.xml
func VertexAttribI4sv(index Uint, v *Short)  {
	C.goglVertexAttribI4sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI4ubv.xml
func VertexAttribI4ubv(index Uint, v *Ubyte)  {
	C.goglVertexAttribI4ubv((C.GLuint)(index), (*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribI4usv.xml
func VertexAttribI4usv(index Uint, v *Ushort)  {
	C.goglVertexAttribI4usv((C.GLuint)(index), (*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetUniformuiv.xml
func GetUniformuiv(program Uint, location Int, params *Uint)  {
	C.goglGetUniformuiv((C.GLuint)(program), (C.GLint)(location), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindFragDataLocation.xml
func BindFragDataLocation(program Uint, color Uint, name *Char)  {
	C.goglBindFragDataLocation((C.GLuint)(program), (C.GLuint)(color), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetFragDataLocation.xml
func GetFragDataLocation(program Uint, name *Char) Int {
	return (Int)(C.goglGetFragDataLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform1ui.xml
func Uniform1ui(location Int, v0 Uint)  {
	C.goglUniform1ui((C.GLint)(location), (C.GLuint)(v0))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform2ui.xml
func Uniform2ui(location Int, v0 Uint, v1 Uint)  {
	C.goglUniform2ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform3ui.xml
func Uniform3ui(location Int, v0 Uint, v1 Uint, v2 Uint)  {
	C.goglUniform3ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform4ui.xml
func Uniform4ui(location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint)  {
	C.goglUniform4ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform1uiv.xml
func Uniform1uiv(location Int, count Sizei, value *Uint)  {
	C.goglUniform1uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform2uiv.xml
func Uniform2uiv(location Int, count Sizei, value *Uint)  {
	C.goglUniform2uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform3uiv.xml
func Uniform3uiv(location Int, count Sizei, value *Uint)  {
	C.goglUniform3uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform4uiv.xml
func Uniform4uiv(location Int, count Sizei, value *Uint)  {
	C.goglUniform4uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexParameterIiv.xml
func TexParameterIiv(target Enum, pname Enum, params *Int)  {
	C.goglTexParameterIiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexParameterIuiv.xml
func TexParameterIuiv(target Enum, pname Enum, params *Uint)  {
	C.goglTexParameterIuiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetTexParameterIiv.xml
func GetTexParameterIiv(target Enum, pname Enum, params *Int)  {
	C.goglGetTexParameterIiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetTexParameterIuiv.xml
func GetTexParameterIuiv(target Enum, pname Enum, params *Uint)  {
	C.goglGetTexParameterIuiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClearBufferiv.xml
func ClearBufferiv(buffer Enum, drawbuffer Int, value *Int)  {
	C.goglClearBufferiv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClearBufferuiv.xml
func ClearBufferuiv(buffer Enum, drawbuffer Int, value *Uint)  {
	C.goglClearBufferuiv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClearBufferfv.xml
func ClearBufferfv(buffer Enum, drawbuffer Int, value *Float)  {
	C.goglClearBufferfv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClearBufferfi.xml
func ClearBufferfi(buffer Enum, drawbuffer Int, depth Float, stencil Int)  {
	C.goglClearBufferfi((C.GLenum)(buffer), (C.GLint)(drawbuffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetStringi.xml
func GetStringi(name Enum, index Uint) *Ubyte {
	return (*Ubyte)(C.goglGetStringi((C.GLenum)(name), (C.GLuint)(index)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsRenderbuffer.xml
func IsRenderbuffer(renderbuffer Uint) Boolean {
	return (Boolean)(C.goglIsRenderbuffer((C.GLuint)(renderbuffer)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindRenderbuffer.xml
func BindRenderbuffer(target Enum, renderbuffer Uint)  {
	C.goglBindRenderbuffer((C.GLenum)(target), (C.GLuint)(renderbuffer))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteRenderbuffers.xml
func DeleteRenderbuffers(n Sizei, renderbuffers *Uint)  {
	C.goglDeleteRenderbuffers((C.GLsizei)(n), (*C.GLuint)(renderbuffers))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGenRenderbuffers.xml
func GenRenderbuffers(n Sizei, renderbuffers *Uint)  {
	C.goglGenRenderbuffers((C.GLsizei)(n), (*C.GLuint)(renderbuffers))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glRenderbufferStorage.xml
func RenderbufferStorage(target Enum, internalformat Enum, width Sizei, height Sizei)  {
	C.goglRenderbufferStorage((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetRenderbufferParameteriv.xml
func GetRenderbufferParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetRenderbufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsFramebuffer.xml
func IsFramebuffer(framebuffer Uint) Boolean {
	return (Boolean)(C.goglIsFramebuffer((C.GLuint)(framebuffer)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindFramebuffer.xml
func BindFramebuffer(target Enum, framebuffer Uint)  {
	C.goglBindFramebuffer((C.GLenum)(target), (C.GLuint)(framebuffer))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteFramebuffers.xml
func DeleteFramebuffers(n Sizei, framebuffers *Uint)  {
	C.goglDeleteFramebuffers((C.GLsizei)(n), (*C.GLuint)(framebuffers))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGenFramebuffers.xml
func GenFramebuffers(n Sizei, framebuffers *Uint)  {
	C.goglGenFramebuffers((C.GLsizei)(n), (*C.GLuint)(framebuffers))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCheckFramebufferStatus.xml
func CheckFramebufferStatus(target Enum) Enum {
	return (Enum)(C.goglCheckFramebufferStatus((C.GLenum)(target)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFramebufferTexture1D.xml
func FramebufferTexture1D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int)  {
	C.goglFramebufferTexture1D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFramebufferTexture2D.xml
func FramebufferTexture2D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int)  {
	C.goglFramebufferTexture2D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFramebufferTexture3D.xml
func FramebufferTexture3D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int, zoffset Int)  {
	C.goglFramebufferTexture3D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(zoffset))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFramebufferRenderbuffer.xml
func FramebufferRenderbuffer(target Enum, attachment Enum, renderbuffertarget Enum, renderbuffer Uint)  {
	C.goglFramebufferRenderbuffer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetFramebufferAttachmentParameteriv.xml
func GetFramebufferAttachmentParameteriv(target Enum, attachment Enum, pname Enum, params *Int)  {
	C.goglGetFramebufferAttachmentParameteriv((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGenerateMipmap.xml
func GenerateMipmap(target Enum)  {
	C.goglGenerateMipmap((C.GLenum)(target))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBlitFramebuffer.xml
func BlitFramebuffer(srcX0 Int, srcY0 Int, srcX1 Int, srcY1 Int, dstX0 Int, dstY0 Int, dstX1 Int, dstY1 Int, mask Bitfield, filter Enum)  {
	C.goglBlitFramebuffer((C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glRenderbufferStorageMultisample.xml
func RenderbufferStorageMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei)  {
	C.goglRenderbufferStorageMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFramebufferTextureLayer.xml
func FramebufferTextureLayer(target Enum, attachment Enum, texture Uint, level Int, layer Int)  {
	C.goglFramebufferTextureLayer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMapBufferRange.xml
func MapBufferRange(target Enum, offset Intptr, length Sizeiptr, access Bitfield) Pointer {
	return (Pointer)(C.goglMapBufferRange((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length), (C.GLbitfield)(access)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFlushMappedBufferRange.xml
func FlushMappedBufferRange(target Enum, offset Intptr, length Sizeiptr)  {
	C.goglFlushMappedBufferRange((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindVertexArray.xml
func BindVertexArray(array Uint)  {
	C.goglBindVertexArray((C.GLuint)(array))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteVertexArrays.xml
func DeleteVertexArrays(n Sizei, arrays *Uint)  {
	C.goglDeleteVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGenVertexArrays.xml
func GenVertexArrays(n Sizei, arrays *Uint)  {
	C.goglGenVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsVertexArray.xml
func IsVertexArray(array Uint) Boolean {
	return (Boolean)(C.goglIsVertexArray((C.GLuint)(array)))
}
// VERSION_3_1

// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawArraysInstanced.xml
func DrawArraysInstanced(mode Enum, first Int, count Sizei, primcount Sizei)  {
	C.goglDrawArraysInstanced((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(primcount))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawElementsInstanced.xml
func DrawElementsInstanced(mode Enum, count Sizei, type_ Enum, indices Pointer, primcount Sizei)  {
	C.goglDrawElementsInstanced((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(primcount))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexBuffer.xml
func TexBuffer(target Enum, internalformat Enum, buffer Uint)  {
	C.goglTexBuffer((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPrimitiveRestartIndex.xml
func PrimitiveRestartIndex(index Uint)  {
	C.goglPrimitiveRestartIndex((C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCopyBufferSubData.xml
func CopyBufferSubData(readTarget Enum, writeTarget Enum, readOffset Intptr, writeOffset Intptr, size Sizeiptr)  {
	C.goglCopyBufferSubData((C.GLenum)(readTarget), (C.GLenum)(writeTarget), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizeiptr)(size))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetUniformIndices.xml
func GetUniformIndices(program Uint, uniformCount Sizei, uniformNames **Char, uniformIndices *Uint)  {
	C.goglGetUniformIndices((C.GLuint)(program), (C.GLsizei)(uniformCount), (**C.GLchar)(unsafe.Pointer(uniformNames)), (*C.GLuint)(uniformIndices))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetActiveUniformsiv.xml
func GetActiveUniformsiv(program Uint, uniformCount Sizei, uniformIndices *Uint, pname Enum, params *Int)  {
	C.goglGetActiveUniformsiv((C.GLuint)(program), (C.GLsizei)(uniformCount), (*C.GLuint)(uniformIndices), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetActiveUniformName.xml
func GetActiveUniformName(program Uint, uniformIndex Uint, bufSize Sizei, length *Sizei, uniformName *Char)  {
	C.goglGetActiveUniformName((C.GLuint)(program), (C.GLuint)(uniformIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(uniformName))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetUniformBlockIndex.xml
func GetUniformBlockIndex(program Uint, uniformBlockName *Char) Uint {
	return (Uint)(C.goglGetUniformBlockIndex((C.GLuint)(program), (*C.GLchar)(uniformBlockName)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetActiveUniformBlockiv.xml
func GetActiveUniformBlockiv(program Uint, uniformBlockIndex Uint, pname Enum, params *Int)  {
	C.goglGetActiveUniformBlockiv((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetActiveUniformBlockName.xml
func GetActiveUniformBlockName(program Uint, uniformBlockIndex Uint, bufSize Sizei, length *Sizei, uniformBlockName *Char)  {
	C.goglGetActiveUniformBlockName((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(uniformBlockName))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformBlockBinding.xml
func UniformBlockBinding(program Uint, uniformBlockIndex Uint, uniformBlockBinding Uint)  {
	C.goglUniformBlockBinding((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLuint)(uniformBlockBinding))
}
// VERSION_3_2

// http://www.opengl.org/sdk/docs/man4/xhtml/glGetInteger64i_v.xml
func GetInteger64i_v(target Enum, index Uint, data *Int64)  {
	C.goglGetInteger64i_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLint64)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetBufferParameteri64v.xml
func GetBufferParameteri64v(target Enum, pname Enum, params *Int64)  {
	C.goglGetBufferParameteri64v((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint64)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFramebufferTexture.xml
func FramebufferTexture(target Enum, attachment Enum, texture Uint, level Int)  {
	C.goglFramebufferTexture((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawElementsBaseVertex.xml
func DrawElementsBaseVertex(mode Enum, count Sizei, type_ Enum, indices Pointer, basevertex Int)  {
	C.goglDrawElementsBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLint)(basevertex))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawRangeElementsBaseVertex.xml
func DrawRangeElementsBaseVertex(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Pointer, basevertex Int)  {
	C.goglDrawRangeElementsBaseVertex((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLint)(basevertex))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawElementsInstancedBaseVertex.xml
func DrawElementsInstancedBaseVertex(mode Enum, count Sizei, type_ Enum, indices Pointer, primcount Sizei, basevertex Int)  {
	C.goglDrawElementsInstancedBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(primcount), (C.GLint)(basevertex))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiDrawElementsBaseVertex.xml
func MultiDrawElementsBaseVertex(mode Enum, count *Sizei, type_ Enum, indices *Pointer, primcount Sizei, basevertex *Int)  {
	C.goglMultiDrawElementsBaseVertex((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(primcount), (*C.GLint)(basevertex))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProvokingVertex.xml
func ProvokingVertex(mode Enum)  {
	C.goglProvokingVertex((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glFenceSync.xml
func FenceSync(condition Enum, flags Bitfield) Sync {
	return (Sync)(C.goglFenceSync((C.GLenum)(condition), (C.GLbitfield)(flags)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsSync.xml
func IsSync(sync Sync) Boolean {
	return (Boolean)(C.goglIsSync((C.GLsync)(sync)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteSync.xml
func DeleteSync(sync Sync)  {
	C.goglDeleteSync((C.GLsync)(sync))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClientWaitSync.xml
func ClientWaitSync(sync Sync, flags Bitfield, timeout Uint64) Enum {
	return (Enum)(C.goglClientWaitSync((C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glWaitSync.xml
func WaitSync(sync Sync, flags Bitfield, timeout Uint64)  {
	C.goglWaitSync((C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetInteger64v.xml
func GetInteger64v(pname Enum, params *Int64)  {
	C.goglGetInteger64v((C.GLenum)(pname), (*C.GLint64)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetSynciv.xml
func GetSynciv(sync Sync, pname Enum, bufSize Sizei, length *Sizei, values *Int)  {
	C.goglGetSynciv((C.GLsync)(sync), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(values))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexImage2DMultisample.xml
func TexImage2DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, fixedsamplelocations Boolean)  {
	C.goglTexImage2DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(fixedsamplelocations))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexImage3DMultisample.xml
func TexImage3DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean)  {
	C.goglTexImage3DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(fixedsamplelocations))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetMultisamplefv.xml
func GetMultisamplefv(pname Enum, index Uint, val *Float)  {
	C.goglGetMultisamplefv((C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(val))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glSampleMaski.xml
func SampleMaski(index Uint, mask Bitfield)  {
	C.goglSampleMaski((C.GLuint)(index), (C.GLbitfield)(mask))
}
// VERSION_3_3

// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribDivisor.xml
func VertexAttribDivisor(index Uint, divisor Uint)  {
	C.goglVertexAttribDivisor((C.GLuint)(index), (C.GLuint)(divisor))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindFragDataLocationIndexed.xml
func BindFragDataLocationIndexed(program Uint, colorNumber Uint, index Uint, name *Char)  {
	C.goglBindFragDataLocationIndexed((C.GLuint)(program), (C.GLuint)(colorNumber), (C.GLuint)(index), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetFragDataIndex.xml
func GetFragDataIndex(program Uint, name *Char) Int {
	return (Int)(C.goglGetFragDataIndex((C.GLuint)(program), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGenSamplers.xml
func GenSamplers(count Sizei, samplers *Uint)  {
	C.goglGenSamplers((C.GLsizei)(count), (*C.GLuint)(samplers))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteSamplers.xml
func DeleteSamplers(count Sizei, samplers *Uint)  {
	C.goglDeleteSamplers((C.GLsizei)(count), (*C.GLuint)(samplers))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsSampler.xml
func IsSampler(sampler Uint) Boolean {
	return (Boolean)(C.goglIsSampler((C.GLuint)(sampler)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindSampler.xml
func BindSampler(unit Uint, sampler Uint)  {
	C.goglBindSampler((C.GLuint)(unit), (C.GLuint)(sampler))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glSamplerParameteri.xml
func SamplerParameteri(sampler Uint, pname Enum, param Int)  {
	C.goglSamplerParameteri((C.GLuint)(sampler), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glSamplerParameteriv.xml
func SamplerParameteriv(sampler Uint, pname Enum, param *Int)  {
	C.goglSamplerParameteriv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glSamplerParameterf.xml
func SamplerParameterf(sampler Uint, pname Enum, param Float)  {
	C.goglSamplerParameterf((C.GLuint)(sampler), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glSamplerParameterfv.xml
func SamplerParameterfv(sampler Uint, pname Enum, param *Float)  {
	C.goglSamplerParameterfv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glSamplerParameterIiv.xml
func SamplerParameterIiv(sampler Uint, pname Enum, param *Int)  {
	C.goglSamplerParameterIiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glSamplerParameterIuiv.xml
func SamplerParameterIuiv(sampler Uint, pname Enum, param *Uint)  {
	C.goglSamplerParameterIuiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(param))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetSamplerParameteriv.xml
func GetSamplerParameteriv(sampler Uint, pname Enum, params *Int)  {
	C.goglGetSamplerParameteriv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetSamplerParameterIiv.xml
func GetSamplerParameterIiv(sampler Uint, pname Enum, params *Int)  {
	C.goglGetSamplerParameterIiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetSamplerParameterfv.xml
func GetSamplerParameterfv(sampler Uint, pname Enum, params *Float)  {
	C.goglGetSamplerParameterfv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetSamplerParameterIuiv.xml
func GetSamplerParameterIuiv(sampler Uint, pname Enum, params *Uint)  {
	C.goglGetSamplerParameterIuiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glQueryCounter.xml
func QueryCounter(id Uint, target Enum)  {
	C.goglQueryCounter((C.GLuint)(id), (C.GLenum)(target))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetQueryObjecti64v.xml
func GetQueryObjecti64v(id Uint, pname Enum, params *Int64)  {
	C.goglGetQueryObjecti64v((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint64)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetQueryObjectui64v.xml
func GetQueryObjectui64v(id Uint, pname Enum, params *Uint64)  {
	C.goglGetQueryObjectui64v((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint64)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexP2ui.xml
func VertexP2ui(type_ Enum, value Uint)  {
	C.goglVertexP2ui((C.GLenum)(type_), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexP2uiv.xml
func VertexP2uiv(type_ Enum, value *Uint)  {
	C.goglVertexP2uiv((C.GLenum)(type_), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexP3ui.xml
func VertexP3ui(type_ Enum, value Uint)  {
	C.goglVertexP3ui((C.GLenum)(type_), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexP3uiv.xml
func VertexP3uiv(type_ Enum, value *Uint)  {
	C.goglVertexP3uiv((C.GLenum)(type_), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexP4ui.xml
func VertexP4ui(type_ Enum, value Uint)  {
	C.goglVertexP4ui((C.GLenum)(type_), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexP4uiv.xml
func VertexP4uiv(type_ Enum, value *Uint)  {
	C.goglVertexP4uiv((C.GLenum)(type_), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexCoordP1ui.xml
func TexCoordP1ui(type_ Enum, coords Uint)  {
	C.goglTexCoordP1ui((C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexCoordP1uiv.xml
func TexCoordP1uiv(type_ Enum, coords *Uint)  {
	C.goglTexCoordP1uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexCoordP2ui.xml
func TexCoordP2ui(type_ Enum, coords Uint)  {
	C.goglTexCoordP2ui((C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexCoordP2uiv.xml
func TexCoordP2uiv(type_ Enum, coords *Uint)  {
	C.goglTexCoordP2uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexCoordP3ui.xml
func TexCoordP3ui(type_ Enum, coords Uint)  {
	C.goglTexCoordP3ui((C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexCoordP3uiv.xml
func TexCoordP3uiv(type_ Enum, coords *Uint)  {
	C.goglTexCoordP3uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexCoordP4ui.xml
func TexCoordP4ui(type_ Enum, coords Uint)  {
	C.goglTexCoordP4ui((C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexCoordP4uiv.xml
func TexCoordP4uiv(type_ Enum, coords *Uint)  {
	C.goglTexCoordP4uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiTexCoordP1ui.xml
func MultiTexCoordP1ui(texture Enum, type_ Enum, coords Uint)  {
	C.goglMultiTexCoordP1ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiTexCoordP1uiv.xml
func MultiTexCoordP1uiv(texture Enum, type_ Enum, coords *Uint)  {
	C.goglMultiTexCoordP1uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiTexCoordP2ui.xml
func MultiTexCoordP2ui(texture Enum, type_ Enum, coords Uint)  {
	C.goglMultiTexCoordP2ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiTexCoordP2uiv.xml
func MultiTexCoordP2uiv(texture Enum, type_ Enum, coords *Uint)  {
	C.goglMultiTexCoordP2uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiTexCoordP3ui.xml
func MultiTexCoordP3ui(texture Enum, type_ Enum, coords Uint)  {
	C.goglMultiTexCoordP3ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiTexCoordP3uiv.xml
func MultiTexCoordP3uiv(texture Enum, type_ Enum, coords *Uint)  {
	C.goglMultiTexCoordP3uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiTexCoordP4ui.xml
func MultiTexCoordP4ui(texture Enum, type_ Enum, coords Uint)  {
	C.goglMultiTexCoordP4ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMultiTexCoordP4uiv.xml
func MultiTexCoordP4uiv(texture Enum, type_ Enum, coords *Uint)  {
	C.goglMultiTexCoordP4uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glNormalP3ui.xml
func NormalP3ui(type_ Enum, coords Uint)  {
	C.goglNormalP3ui((C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glNormalP3uiv.xml
func NormalP3uiv(type_ Enum, coords *Uint)  {
	C.goglNormalP3uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glColorP3ui.xml
func ColorP3ui(type_ Enum, color Uint)  {
	C.goglColorP3ui((C.GLenum)(type_), (C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glColorP3uiv.xml
func ColorP3uiv(type_ Enum, color *Uint)  {
	C.goglColorP3uiv((C.GLenum)(type_), (*C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glColorP4ui.xml
func ColorP4ui(type_ Enum, color Uint)  {
	C.goglColorP4ui((C.GLenum)(type_), (C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glColorP4uiv.xml
func ColorP4uiv(type_ Enum, color *Uint)  {
	C.goglColorP4uiv((C.GLenum)(type_), (*C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glSecondaryColorP3ui.xml
func SecondaryColorP3ui(type_ Enum, color Uint)  {
	C.goglSecondaryColorP3ui((C.GLenum)(type_), (C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glSecondaryColorP3uiv.xml
func SecondaryColorP3uiv(type_ Enum, color *Uint)  {
	C.goglSecondaryColorP3uiv((C.GLenum)(type_), (*C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribP1ui.xml
func VertexAttribP1ui(index Uint, type_ Enum, normalized Boolean, value Uint)  {
	C.goglVertexAttribP1ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribP1uiv.xml
func VertexAttribP1uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)  {
	C.goglVertexAttribP1uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribP2ui.xml
func VertexAttribP2ui(index Uint, type_ Enum, normalized Boolean, value Uint)  {
	C.goglVertexAttribP2ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribP2uiv.xml
func VertexAttribP2uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)  {
	C.goglVertexAttribP2uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribP3ui.xml
func VertexAttribP3ui(index Uint, type_ Enum, normalized Boolean, value Uint)  {
	C.goglVertexAttribP3ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribP3uiv.xml
func VertexAttribP3uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)  {
	C.goglVertexAttribP3uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribP4ui.xml
func VertexAttribP4ui(index Uint, type_ Enum, normalized Boolean, value Uint)  {
	C.goglVertexAttribP4ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribP4uiv.xml
func VertexAttribP4uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)  {
	C.goglVertexAttribP4uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}
// VERSION_4_0

// http://www.opengl.org/sdk/docs/man4/xhtml/glMinSampleShading.xml
func MinSampleShading(value Float)  {
	C.goglMinSampleShading((C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBlendEquationi.xml
func BlendEquationi(buf Uint, mode Enum)  {
	C.goglBlendEquationi((C.GLuint)(buf), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBlendEquationSeparatei.xml
func BlendEquationSeparatei(buf Uint, modeRGB Enum, modeAlpha Enum)  {
	C.goglBlendEquationSeparatei((C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBlendFunci.xml
func BlendFunci(buf Uint, src Enum, dst Enum)  {
	C.goglBlendFunci((C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBlendFuncSeparatei.xml
func BlendFuncSeparatei(buf Uint, srcRGB Enum, dstRGB Enum, srcAlpha Enum, dstAlpha Enum)  {
	C.goglBlendFuncSeparatei((C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawArraysIndirect.xml
func DrawArraysIndirect(mode Enum, indirect Pointer)  {
	C.goglDrawArraysIndirect((C.GLenum)(mode), (unsafe.Pointer)(indirect))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawElementsIndirect.xml
func DrawElementsIndirect(mode Enum, type_ Enum, indirect Pointer)  {
	C.goglDrawElementsIndirect((C.GLenum)(mode), (C.GLenum)(type_), (unsafe.Pointer)(indirect))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform1d.xml
func Uniform1d(location Int, x Double)  {
	C.goglUniform1d((C.GLint)(location), (C.GLdouble)(x))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform2d.xml
func Uniform2d(location Int, x Double, y Double)  {
	C.goglUniform2d((C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform3d.xml
func Uniform3d(location Int, x Double, y Double, z Double)  {
	C.goglUniform3d((C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform4d.xml
func Uniform4d(location Int, x Double, y Double, z Double, w Double)  {
	C.goglUniform4d((C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform1dv.xml
func Uniform1dv(location Int, count Sizei, value *Double)  {
	C.goglUniform1dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform2dv.xml
func Uniform2dv(location Int, count Sizei, value *Double)  {
	C.goglUniform2dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform3dv.xml
func Uniform3dv(location Int, count Sizei, value *Double)  {
	C.goglUniform3dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniform4dv.xml
func Uniform4dv(location Int, count Sizei, value *Double)  {
	C.goglUniform4dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix2dv.xml
func UniformMatrix2dv(location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglUniformMatrix2dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix3dv.xml
func UniformMatrix3dv(location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglUniformMatrix3dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix4dv.xml
func UniformMatrix4dv(location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglUniformMatrix4dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix2x3dv.xml
func UniformMatrix2x3dv(location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglUniformMatrix2x3dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix2x4dv.xml
func UniformMatrix2x4dv(location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglUniformMatrix2x4dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix3x2dv.xml
func UniformMatrix3x2dv(location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglUniformMatrix3x2dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix3x4dv.xml
func UniformMatrix3x4dv(location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglUniformMatrix3x4dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix4x2dv.xml
func UniformMatrix4x2dv(location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglUniformMatrix4x2dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformMatrix4x3dv.xml
func UniformMatrix4x3dv(location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglUniformMatrix4x3dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetUniformdv.xml
func GetUniformdv(program Uint, location Int, params *Double)  {
	C.goglGetUniformdv((C.GLuint)(program), (C.GLint)(location), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetSubroutineUniformLocation.xml
func GetSubroutineUniformLocation(program Uint, shadertype Enum, name *Char) Int {
	return (Int)(C.goglGetSubroutineUniformLocation((C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetSubroutineIndex.xml
func GetSubroutineIndex(program Uint, shadertype Enum, name *Char) Uint {
	return (Uint)(C.goglGetSubroutineIndex((C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetActiveSubroutineUniformiv.xml
func GetActiveSubroutineUniformiv(program Uint, shadertype Enum, index Uint, pname Enum, values *Int)  {
	C.goglGetActiveSubroutineUniformiv((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(values))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetActiveSubroutineUniformName.xml
func GetActiveSubroutineUniformName(program Uint, shadertype Enum, index Uint, bufsize Sizei, length *Sizei, name *Char)  {
	C.goglGetActiveSubroutineUniformName((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(length), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetActiveSubroutineName.xml
func GetActiveSubroutineName(program Uint, shadertype Enum, index Uint, bufsize Sizei, length *Sizei, name *Char)  {
	C.goglGetActiveSubroutineName((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(length), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUniformSubroutinesuiv.xml
func UniformSubroutinesuiv(shadertype Enum, count Sizei, indices *Uint)  {
	C.goglUniformSubroutinesuiv((C.GLenum)(shadertype), (C.GLsizei)(count), (*C.GLuint)(indices))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetUniformSubroutineuiv.xml
func GetUniformSubroutineuiv(shadertype Enum, location Int, params *Uint)  {
	C.goglGetUniformSubroutineuiv((C.GLenum)(shadertype), (C.GLint)(location), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetProgramStageiv.xml
func GetProgramStageiv(program Uint, shadertype Enum, pname Enum, values *Int)  {
	C.goglGetProgramStageiv((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLenum)(pname), (*C.GLint)(values))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPatchParameteri.xml
func PatchParameteri(pname Enum, value Int)  {
	C.goglPatchParameteri((C.GLenum)(pname), (C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPatchParameterfv.xml
func PatchParameterfv(pname Enum, values *Float)  {
	C.goglPatchParameterfv((C.GLenum)(pname), (*C.GLfloat)(values))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindTransformFeedback.xml
func BindTransformFeedback(target Enum, id Uint)  {
	C.goglBindTransformFeedback((C.GLenum)(target), (C.GLuint)(id))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteTransformFeedbacks.xml
func DeleteTransformFeedbacks(n Sizei, ids *Uint)  {
	C.goglDeleteTransformFeedbacks((C.GLsizei)(n), (*C.GLuint)(ids))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGenTransformFeedbacks.xml
func GenTransformFeedbacks(n Sizei, ids *Uint)  {
	C.goglGenTransformFeedbacks((C.GLsizei)(n), (*C.GLuint)(ids))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsTransformFeedback.xml
func IsTransformFeedback(id Uint) Boolean {
	return (Boolean)(C.goglIsTransformFeedback((C.GLuint)(id)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glPauseTransformFeedback.xml
func PauseTransformFeedback()  {
	C.goglPauseTransformFeedback()
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glResumeTransformFeedback.xml
func ResumeTransformFeedback()  {
	C.goglResumeTransformFeedback()
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawTransformFeedback.xml
func DrawTransformFeedback(mode Enum, id Uint)  {
	C.goglDrawTransformFeedback((C.GLenum)(mode), (C.GLuint)(id))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawTransformFeedbackStream.xml
func DrawTransformFeedbackStream(mode Enum, id Uint, stream Uint)  {
	C.goglDrawTransformFeedbackStream((C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBeginQueryIndexed.xml
func BeginQueryIndexed(target Enum, index Uint, id Uint)  {
	C.goglBeginQueryIndexed((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(id))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glEndQueryIndexed.xml
func EndQueryIndexed(target Enum, index Uint)  {
	C.goglEndQueryIndexed((C.GLenum)(target), (C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetQueryIndexediv.xml
func GetQueryIndexediv(target Enum, index Uint, pname Enum, params *Int)  {
	C.goglGetQueryIndexediv((C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
// VERSION_4_1

// http://www.opengl.org/sdk/docs/man4/xhtml/glReleaseShaderCompiler.xml
func ReleaseShaderCompiler()  {
	C.goglReleaseShaderCompiler()
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glShaderBinary.xml
func ShaderBinary(count Sizei, shaders *Uint, binaryformat Enum, binary Pointer, length Sizei)  {
	C.goglShaderBinary((C.GLsizei)(count), (*C.GLuint)(shaders), (C.GLenum)(binaryformat), (unsafe.Pointer)(binary), (C.GLsizei)(length))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetShaderPrecisionFormat.xml
func GetShaderPrecisionFormat(shadertype Enum, precisiontype Enum, range_ *Int, precision *Int)  {
	C.goglGetShaderPrecisionFormat((C.GLenum)(shadertype), (C.GLenum)(precisiontype), (*C.GLint)(range_), (*C.GLint)(precision))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDepthRangef.xml
func DepthRangef(n Float, f Float)  {
	C.goglDepthRangef((C.GLfloat)(n), (C.GLfloat)(f))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glClearDepthf.xml
func ClearDepthf(d Float)  {
	C.goglClearDepthf((C.GLfloat)(d))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetProgramBinary.xml
func GetProgramBinary(program Uint, bufSize Sizei, length *Sizei, binaryFormat *Enum, binary Pointer)  {
	C.goglGetProgramBinary((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLenum)(binaryFormat), (unsafe.Pointer)(binary))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramBinary.xml
func ProgramBinary(program Uint, binaryFormat Enum, binary Pointer, length Sizei)  {
	C.goglProgramBinary((C.GLuint)(program), (C.GLenum)(binaryFormat), (unsafe.Pointer)(binary), (C.GLsizei)(length))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramParameteri.xml
func ProgramParameteri(program Uint, pname Enum, value Int)  {
	C.goglProgramParameteri((C.GLuint)(program), (C.GLenum)(pname), (C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glUseProgramStages.xml
func UseProgramStages(pipeline Uint, stages Bitfield, program Uint)  {
	C.goglUseProgramStages((C.GLuint)(pipeline), (C.GLbitfield)(stages), (C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glActiveShaderProgram.xml
func ActiveShaderProgram(pipeline Uint, program Uint)  {
	C.goglActiveShaderProgram((C.GLuint)(pipeline), (C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glCreateShaderProgramv.xml
func CreateShaderProgramv(type_ Enum, count Sizei, strings **Char) Uint {
	return (Uint)(C.goglCreateShaderProgramv((C.GLenum)(type_), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(strings))))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindProgramPipeline.xml
func BindProgramPipeline(pipeline Uint)  {
	C.goglBindProgramPipeline((C.GLuint)(pipeline))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDeleteProgramPipelines.xml
func DeleteProgramPipelines(n Sizei, pipelines *Uint)  {
	C.goglDeleteProgramPipelines((C.GLsizei)(n), (*C.GLuint)(pipelines))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGenProgramPipelines.xml
func GenProgramPipelines(n Sizei, pipelines *Uint)  {
	C.goglGenProgramPipelines((C.GLsizei)(n), (*C.GLuint)(pipelines))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glIsProgramPipeline.xml
func IsProgramPipeline(pipeline Uint) Boolean {
	return (Boolean)(C.goglIsProgramPipeline((C.GLuint)(pipeline)))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetProgramPipelineiv.xml
func GetProgramPipelineiv(pipeline Uint, pname Enum, params *Int)  {
	C.goglGetProgramPipelineiv((C.GLuint)(pipeline), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform1i.xml
func ProgramUniform1i(program Uint, location Int, v0 Int)  {
	C.goglProgramUniform1i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform1iv.xml
func ProgramUniform1iv(program Uint, location Int, count Sizei, value *Int)  {
	C.goglProgramUniform1iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform1f.xml
func ProgramUniform1f(program Uint, location Int, v0 Float)  {
	C.goglProgramUniform1f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform1fv.xml
func ProgramUniform1fv(program Uint, location Int, count Sizei, value *Float)  {
	C.goglProgramUniform1fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform1d.xml
func ProgramUniform1d(program Uint, location Int, v0 Double)  {
	C.goglProgramUniform1d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform1dv.xml
func ProgramUniform1dv(program Uint, location Int, count Sizei, value *Double)  {
	C.goglProgramUniform1dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform1ui.xml
func ProgramUniform1ui(program Uint, location Int, v0 Uint)  {
	C.goglProgramUniform1ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform1uiv.xml
func ProgramUniform1uiv(program Uint, location Int, count Sizei, value *Uint)  {
	C.goglProgramUniform1uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform2i.xml
func ProgramUniform2i(program Uint, location Int, v0 Int, v1 Int)  {
	C.goglProgramUniform2i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform2iv.xml
func ProgramUniform2iv(program Uint, location Int, count Sizei, value *Int)  {
	C.goglProgramUniform2iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform2f.xml
func ProgramUniform2f(program Uint, location Int, v0 Float, v1 Float)  {
	C.goglProgramUniform2f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform2fv.xml
func ProgramUniform2fv(program Uint, location Int, count Sizei, value *Float)  {
	C.goglProgramUniform2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform2d.xml
func ProgramUniform2d(program Uint, location Int, v0 Double, v1 Double)  {
	C.goglProgramUniform2d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform2dv.xml
func ProgramUniform2dv(program Uint, location Int, count Sizei, value *Double)  {
	C.goglProgramUniform2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform2ui.xml
func ProgramUniform2ui(program Uint, location Int, v0 Uint, v1 Uint)  {
	C.goglProgramUniform2ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform2uiv.xml
func ProgramUniform2uiv(program Uint, location Int, count Sizei, value *Uint)  {
	C.goglProgramUniform2uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform3i.xml
func ProgramUniform3i(program Uint, location Int, v0 Int, v1 Int, v2 Int)  {
	C.goglProgramUniform3i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform3iv.xml
func ProgramUniform3iv(program Uint, location Int, count Sizei, value *Int)  {
	C.goglProgramUniform3iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform3f.xml
func ProgramUniform3f(program Uint, location Int, v0 Float, v1 Float, v2 Float)  {
	C.goglProgramUniform3f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform3fv.xml
func ProgramUniform3fv(program Uint, location Int, count Sizei, value *Float)  {
	C.goglProgramUniform3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform3d.xml
func ProgramUniform3d(program Uint, location Int, v0 Double, v1 Double, v2 Double)  {
	C.goglProgramUniform3d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform3dv.xml
func ProgramUniform3dv(program Uint, location Int, count Sizei, value *Double)  {
	C.goglProgramUniform3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform3ui.xml
func ProgramUniform3ui(program Uint, location Int, v0 Uint, v1 Uint, v2 Uint)  {
	C.goglProgramUniform3ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform3uiv.xml
func ProgramUniform3uiv(program Uint, location Int, count Sizei, value *Uint)  {
	C.goglProgramUniform3uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform4i.xml
func ProgramUniform4i(program Uint, location Int, v0 Int, v1 Int, v2 Int, v3 Int)  {
	C.goglProgramUniform4i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform4iv.xml
func ProgramUniform4iv(program Uint, location Int, count Sizei, value *Int)  {
	C.goglProgramUniform4iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform4f.xml
func ProgramUniform4f(program Uint, location Int, v0 Float, v1 Float, v2 Float, v3 Float)  {
	C.goglProgramUniform4f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform4fv.xml
func ProgramUniform4fv(program Uint, location Int, count Sizei, value *Float)  {
	C.goglProgramUniform4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform4d.xml
func ProgramUniform4d(program Uint, location Int, v0 Double, v1 Double, v2 Double, v3 Double)  {
	C.goglProgramUniform4d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLdouble)(v3))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform4dv.xml
func ProgramUniform4dv(program Uint, location Int, count Sizei, value *Double)  {
	C.goglProgramUniform4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform4ui.xml
func ProgramUniform4ui(program Uint, location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint)  {
	C.goglProgramUniform4ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniform4uiv.xml
func ProgramUniform4uiv(program Uint, location Int, count Sizei, value *Uint)  {
	C.goglProgramUniform4uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix2fv.xml
func ProgramUniformMatrix2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglProgramUniformMatrix2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix3fv.xml
func ProgramUniformMatrix3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglProgramUniformMatrix3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix4fv.xml
func ProgramUniformMatrix4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglProgramUniformMatrix4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix2dv.xml
func ProgramUniformMatrix2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglProgramUniformMatrix2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix3dv.xml
func ProgramUniformMatrix3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglProgramUniformMatrix3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix4dv.xml
func ProgramUniformMatrix4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglProgramUniformMatrix4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix2x3fv.xml
func ProgramUniformMatrix2x3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglProgramUniformMatrix2x3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix3x2fv.xml
func ProgramUniformMatrix3x2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglProgramUniformMatrix3x2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix2x4fv.xml
func ProgramUniformMatrix2x4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglProgramUniformMatrix2x4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix4x2fv.xml
func ProgramUniformMatrix4x2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglProgramUniformMatrix4x2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix3x4fv.xml
func ProgramUniformMatrix3x4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglProgramUniformMatrix3x4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix4x3fv.xml
func ProgramUniformMatrix4x3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglProgramUniformMatrix4x3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix2x3dv.xml
func ProgramUniformMatrix2x3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglProgramUniformMatrix2x3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix3x2dv.xml
func ProgramUniformMatrix3x2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglProgramUniformMatrix3x2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix2x4dv.xml
func ProgramUniformMatrix2x4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglProgramUniformMatrix2x4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix4x2dv.xml
func ProgramUniformMatrix4x2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglProgramUniformMatrix4x2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix3x4dv.xml
func ProgramUniformMatrix3x4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglProgramUniformMatrix3x4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glProgramUniformMatrix4x3dv.xml
func ProgramUniformMatrix4x3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)  {
	C.goglProgramUniformMatrix4x3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glValidateProgramPipeline.xml
func ValidateProgramPipeline(pipeline Uint)  {
	C.goglValidateProgramPipeline((C.GLuint)(pipeline))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetProgramPipelineInfoLog.xml
func GetProgramPipelineInfoLog(pipeline Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.goglGetProgramPipelineInfoLog((C.GLuint)(pipeline), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribL1d.xml
func VertexAttribL1d(index Uint, x Double)  {
	C.goglVertexAttribL1d((C.GLuint)(index), (C.GLdouble)(x))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribL2d.xml
func VertexAttribL2d(index Uint, x Double, y Double)  {
	C.goglVertexAttribL2d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribL3d.xml
func VertexAttribL3d(index Uint, x Double, y Double, z Double)  {
	C.goglVertexAttribL3d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribL4d.xml
func VertexAttribL4d(index Uint, x Double, y Double, z Double, w Double)  {
	C.goglVertexAttribL4d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribL1dv.xml
func VertexAttribL1dv(index Uint, v *Double)  {
	C.goglVertexAttribL1dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribL2dv.xml
func VertexAttribL2dv(index Uint, v *Double)  {
	C.goglVertexAttribL2dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribL3dv.xml
func VertexAttribL3dv(index Uint, v *Double)  {
	C.goglVertexAttribL3dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribL4dv.xml
func VertexAttribL4dv(index Uint, v *Double)  {
	C.goglVertexAttribL4dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glVertexAttribLPointer.xml
func VertexAttribLPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglVertexAttribLPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetVertexAttribLdv.xml
func GetVertexAttribLdv(index Uint, pname Enum, params *Double)  {
	C.goglGetVertexAttribLdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glViewportArrayv.xml
func ViewportArrayv(first Uint, count Sizei, v *Float)  {
	C.goglViewportArrayv((C.GLuint)(first), (C.GLsizei)(count), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glViewportIndexedf.xml
func ViewportIndexedf(index Uint, x Float, y Float, w Float, h Float)  {
	C.goglViewportIndexedf((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(w), (C.GLfloat)(h))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glViewportIndexedfv.xml
func ViewportIndexedfv(index Uint, v *Float)  {
	C.goglViewportIndexedfv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glScissorArrayv.xml
func ScissorArrayv(first Uint, count Sizei, v *Int)  {
	C.goglScissorArrayv((C.GLuint)(first), (C.GLsizei)(count), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glScissorIndexed.xml
func ScissorIndexed(index Uint, left Int, bottom Int, width Sizei, height Sizei)  {
	C.goglScissorIndexed((C.GLuint)(index), (C.GLint)(left), (C.GLint)(bottom), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glScissorIndexedv.xml
func ScissorIndexedv(index Uint, v *Int)  {
	C.goglScissorIndexedv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDepthRangeArrayv.xml
func DepthRangeArrayv(first Uint, count Sizei, v *Double)  {
	C.goglDepthRangeArrayv((C.GLuint)(first), (C.GLsizei)(count), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDepthRangeIndexed.xml
func DepthRangeIndexed(index Uint, n Double, f Double)  {
	C.goglDepthRangeIndexed((C.GLuint)(index), (C.GLdouble)(n), (C.GLdouble)(f))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetFloati_v.xml
func GetFloati_v(target Enum, index Uint, data *Float)  {
	C.goglGetFloati_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLfloat)(data))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetDoublei_v.xml
func GetDoublei_v(target Enum, index Uint, data *Double)  {
	C.goglGetDoublei_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLdouble)(data))
}
// VERSION_4_2

// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawArraysInstancedBaseInstance.xml
func DrawArraysInstancedBaseInstance(mode Enum, first Int, count Sizei, primcount Sizei, baseinstance Uint)  {
	C.goglDrawArraysInstancedBaseInstance((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(primcount), (C.GLuint)(baseinstance))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawElementsInstancedBaseInstance.xml
func DrawElementsInstancedBaseInstance(mode Enum, count Sizei, type_ Enum, indices Pointer, primcount Sizei, baseinstance Uint)  {
	C.goglDrawElementsInstancedBaseInstance((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(primcount), (C.GLuint)(baseinstance))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawElementsInstancedBaseVertexBaseInstance.xml
func DrawElementsInstancedBaseVertexBaseInstance(mode Enum, count Sizei, type_ Enum, indices Pointer, primcount Sizei, basevertex Int, baseinstance Uint)  {
	C.goglDrawElementsInstancedBaseVertexBaseInstance((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(primcount), (C.GLint)(basevertex), (C.GLuint)(baseinstance))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawTransformFeedbackInstanced.xml
func DrawTransformFeedbackInstanced(mode Enum, id Uint, primcount Sizei)  {
	C.goglDrawTransformFeedbackInstanced((C.GLenum)(mode), (C.GLuint)(id), (C.GLsizei)(primcount))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glDrawTransformFeedbackStreamInstanced.xml
func DrawTransformFeedbackStreamInstanced(mode Enum, id Uint, stream Uint, primcount Sizei)  {
	C.goglDrawTransformFeedbackStreamInstanced((C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream), (C.GLsizei)(primcount))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetInternalformativ.xml
func GetInternalformativ(target Enum, internalformat Enum, pname Enum, bufSize Sizei, params *Int)  {
	C.goglGetInternalformativ((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glGetActiveAtomicCounterBufferiv.xml
func GetActiveAtomicCounterBufferiv(program Uint, bufferIndex Uint, pname Enum, params *Int)  {
	C.goglGetActiveAtomicCounterBufferiv((C.GLuint)(program), (C.GLuint)(bufferIndex), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glBindImageTexture.xml
func BindImageTexture(unit Uint, texture Uint, level Int, layered Boolean, layer Int, access Enum, format Enum)  {
	C.goglBindImageTexture((C.GLuint)(unit), (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(layered), (C.GLint)(layer), (C.GLenum)(access), (C.GLenum)(format))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glMemoryBarrier.xml
func MemoryBarrier(barriers Bitfield)  {
	C.goglMemoryBarrier((C.GLbitfield)(barriers))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexStorage1D.xml
func TexStorage1D(target Enum, levels Sizei, internalformat Enum, width Sizei)  {
	C.goglTexStorage1D((C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexStorage2D.xml
func TexStorage2D(target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei)  {
	C.goglTexStorage2D((C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTexStorage3D.xml
func TexStorage3D(target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei, depth Sizei)  {
	C.goglTexStorage3D((C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTextureStorage1DEXT.xml
func TextureStorage1DEXT(texture Uint, target Enum, levels Sizei, internalformat Enum, width Sizei)  {
	C.goglTextureStorage1DEXT((C.GLuint)(texture), (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTextureStorage2DEXT.xml
func TextureStorage2DEXT(texture Uint, target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei)  {
	C.goglTextureStorage2DEXT((C.GLuint)(texture), (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man4/xhtml/glTextureStorage3DEXT.xml
func TextureStorage3DEXT(texture Uint, target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei, depth Sizei)  {
	C.goglTextureStorage3DEXT((C.GLuint)(texture), (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
func InitVersion10() error {
	var ret C.int
	if ret = C.init_VERSION_1_0(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_0")
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
func InitVersion12() error {
	var ret C.int
	if ret = C.init_VERSION_1_2(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_2")
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
func InitVersion30() error {
	var ret C.int
	if ret = C.init_VERSION_3_0(); ret != 0 {
		return errors.New("unable to initialize VERSION_3_0")
	}
	return nil
}
func InitVersion31() error {
	var ret C.int
	if ret = C.init_VERSION_3_1(); ret != 0 {
		return errors.New("unable to initialize VERSION_3_1")
	}
	return nil
}
func InitVersion32() error {
	var ret C.int
	if ret = C.init_VERSION_3_2(); ret != 0 {
		return errors.New("unable to initialize VERSION_3_2")
	}
	return nil
}
func InitVersion33() error {
	var ret C.int
	if ret = C.init_VERSION_3_3(); ret != 0 {
		return errors.New("unable to initialize VERSION_3_3")
	}
	return nil
}
func InitVersion40() error {
	var ret C.int
	if ret = C.init_VERSION_4_0(); ret != 0 {
		return errors.New("unable to initialize VERSION_4_0")
	}
	return nil
}
func InitVersion41() error {
	var ret C.int
	if ret = C.init_VERSION_4_1(); ret != 0 {
		return errors.New("unable to initialize VERSION_4_1")
	}
	return nil
}
func InitVersion42() error {
	var ret C.int
	if ret = C.init_VERSION_4_2(); ret != 0 {
		return errors.New("unable to initialize VERSION_4_2")
	}
	return nil
}
func Init() error {
	var err error
	if err = InitVersion10(); err != nil {
		return err
	}
	if err = InitVersion11(); err != nil {
		return err
	}
	if err = InitVersion12(); err != nil {
		return err
	}
	if err = InitVersion13(); err != nil {
		return err
	}
	if err = InitVersion14(); err != nil {
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
	if err = InitVersion30(); err != nil {
		return err
	}
	if err = InitVersion31(); err != nil {
		return err
	}
	if err = InitVersion32(); err != nil {
		return err
	}
	if err = InitVersion33(); err != nil {
		return err
	}
	if err = InitVersion40(); err != nil {
		return err
	}
	if err = InitVersion41(); err != nil {
		return err
	}
	if err = InitVersion42(); err != nil {
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