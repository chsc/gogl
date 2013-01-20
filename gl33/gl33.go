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
// http://www.opengl.org/sdk/docs/man3
// 
package gl33

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
// #ifndef GL_KHR_debug
// typedef void (APIENTRY *GLDEBUGPROC)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,GLvoid *userParam);
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
// void (APIENTRYP ptrglIndexub)(GLubyte c);
// void (APIENTRYP ptrglIndexubv)(GLubyte* c);
// //  VERSION_1_2
// void (APIENTRYP ptrglBlendColor)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglBlendEquation)(GLenum mode);
// void (APIENTRYP ptrglDrawRangeElements)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices);
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
// void (APIENTRYP ptrglMultiDrawArrays)(GLenum mode, GLint* first, GLsizei* count, GLsizei drawcount);
// void (APIENTRYP ptrglMultiDrawElements)(GLenum mode, GLsizei* count, GLenum type, GLvoid* const* indices, GLsizei drawcount);
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
// void (APIENTRYP ptrglDrawArraysInstanced)(GLenum mode, GLint first, GLsizei count, GLsizei instancecount);
// void (APIENTRYP ptrglDrawElementsInstanced)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei instancecount);
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
// void (APIENTRYP ptrglDrawElementsInstancedBaseVertex)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei instancecount, GLint basevertex);
// void (APIENTRYP ptrglMultiDrawElementsBaseVertex)(GLenum mode, GLsizei* count, GLenum type, GLvoid* const* indices, GLsizei drawcount, GLint* basevertex);
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
// void goglIndexub(GLubyte c) {
// 	(*ptrglIndexub)(c);
// }
// void goglIndexubv(GLubyte* c) {
// 	(*ptrglIndexubv)(c);
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
// void goglMultiDrawArrays(GLenum mode, GLint* first, GLsizei* count, GLsizei drawcount) {
// 	(*ptrglMultiDrawArrays)(mode, first, count, drawcount);
// }
// void goglMultiDrawElements(GLenum mode, GLsizei* count, GLenum type_, GLvoid* const* indices, GLsizei drawcount) {
// 	(*ptrglMultiDrawElements)(mode, count, type_, indices, drawcount);
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
// void goglDrawArraysInstanced(GLenum mode, GLint first, GLsizei count, GLsizei instancecount) {
// 	(*ptrglDrawArraysInstanced)(mode, first, count, instancecount);
// }
// void goglDrawElementsInstanced(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLsizei instancecount) {
// 	(*ptrglDrawElementsInstanced)(mode, count, type_, indices, instancecount);
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
// void goglDrawElementsInstancedBaseVertex(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLsizei instancecount, GLint basevertex) {
// 	(*ptrglDrawElementsInstancedBaseVertex)(mode, count, type_, indices, instancecount, basevertex);
// }
// void goglMultiDrawElementsBaseVertex(GLenum mode, GLsizei* count, GLenum type_, GLvoid* const* indices, GLsizei drawcount, GLint* basevertex) {
// 	(*ptrglMultiDrawElementsBaseVertex)(mode, count, type_, indices, drawcount, basevertex);
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
// 	ptrglIndexub = goglGetProcAddress("glIndexub");
// 	if(ptrglIndexub == NULL) return 1;
// 	ptrglIndexubv = goglGetProcAddress("glIndexubv");
// 	if(ptrglIndexubv == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_2() {
// 	ptrglBlendColor = goglGetProcAddress("glBlendColor");
// 	if(ptrglBlendColor == NULL) return 1;
// 	ptrglBlendEquation = goglGetProcAddress("glBlendEquation");
// 	if(ptrglBlendEquation == NULL) return 1;
// 	ptrglDrawRangeElements = goglGetProcAddress("glDrawRangeElements");
// 	if(ptrglDrawRangeElements == NULL) return 1;
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
	ALPHA = 0x1906
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
	ALWAYS = 0x0207
	AMBIENT = 0x1200
	AMBIENT_AND_DIFFUSE = 0x1602
	AND = 0x1501
	AND_INVERTED = 0x1504
	AND_REVERSE = 0x1502
	ATTRIB_STACK_DEPTH = 0x0BB0
	AUTO_NORMAL = 0x0D80
	AUX0 = 0x0409
	AUX1 = 0x040A
	AUX2 = 0x040B
	AUX3 = 0x040C
	AUX_BUFFERS = 0x0C00
	BACK = 0x0405
	BACK_LEFT = 0x0402
	BACK_RIGHT = 0x0403
	BITMAP = 0x1A00
	BITMAP_TOKEN = 0x0704
	BLEND = 0x0BE2
	BLEND_DST = 0x0BE0
	BLEND_SRC = 0x0BE1
	BLUE = 0x1905
	BLUE_BIAS = 0x0D1B
	BLUE_BITS = 0x0D54
	BLUE_SCALE = 0x0D1A
	BYTE = 0x1400
	C3F_V3F = 0x2A24
	C4F_N3F_V3F = 0x2A26
	C4UB_V2F = 0x2A22
	C4UB_V3F = 0x2A23
	CCW = 0x0901
	CLAMP = 0x2900
	CLEAR = 0x1500
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
	COLOR = 0x1800
	COLOR_ARRAY = 0x8076
	COLOR_ARRAY_POINTER = 0x8090
	COLOR_ARRAY_SIZE = 0x8081
	COLOR_ARRAY_STRIDE = 0x8083
	COLOR_ARRAY_TYPE = 0x8082
	COLOR_BUFFER_BIT = 0x00004000
	COLOR_CLEAR_VALUE = 0x0C22
	COLOR_INDEX = 0x1900
	COLOR_INDEXES = 0x1603
	COLOR_LOGIC_OP = 0x0BF2
	COLOR_MATERIAL = 0x0B57
	COLOR_MATERIAL_FACE = 0x0B55
	COLOR_MATERIAL_PARAMETER = 0x0B56
	COLOR_WRITEMASK = 0x0C23
	COMPILE = 0x1300
	COMPILE_AND_EXECUTE = 0x1301
	CONSTANT_ATTENUATION = 0x1207
	COPY = 0x1503
	COPY_INVERTED = 0x150C
	COPY_PIXEL_TOKEN = 0x0706
	CULL_FACE = 0x0B44
	CULL_FACE_MODE = 0x0B45
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
	CW = 0x0900
	DECAL = 0x2101
	DECR = 0x1E03
	DEPTH = 0x1801
	DEPTH_BIAS = 0x0D1F
	DEPTH_BITS = 0x0D56
	DEPTH_BUFFER_BIT = 0x00000100
	DEPTH_CLEAR_VALUE = 0x0B73
	DEPTH_COMPONENT = 0x1902
	DEPTH_FUNC = 0x0B74
	DEPTH_RANGE = 0x0B70
	DEPTH_SCALE = 0x0D1E
	DEPTH_TEST = 0x0B71
	DEPTH_WRITEMASK = 0x0B72
	DIFFUSE = 0x1201
	DITHER = 0x0BD0
	DOMAIN = 0x0A02
	DONT_CARE = 0x1100
	DOUBLE = 0x140A
	DOUBLEBUFFER = 0x0C32
	DRAW_BUFFER = 0x0C01
	DRAW_PIXEL_TOKEN = 0x0705
	DST_ALPHA = 0x0304
	DST_COLOR = 0x0306
	EDGE_FLAG = 0x0B43
	EDGE_FLAG_ARRAY = 0x8079
	EDGE_FLAG_ARRAY_POINTER = 0x8093
	EDGE_FLAG_ARRAY_STRIDE = 0x808C
	EMISSION = 0x1600
	ENABLE_BIT = 0x00002000
	EQUAL = 0x0202
	EQUIV = 0x1509
	EVAL_BIT = 0x00010000
	EXP = 0x0800
	EXP2 = 0x0801
	EXTENSIONS = 0x1F03
	EYE_LINEAR = 0x2400
	EYE_PLANE = 0x2502
	FALSE = 0
	FASTEST = 0x1101
	FEEDBACK = 0x1C01
	FEEDBACK_BUFFER_POINTER = 0x0DF0
	FEEDBACK_BUFFER_SIZE = 0x0DF1
	FEEDBACK_BUFFER_TYPE = 0x0DF2
	FILL = 0x1B02
	FLAT = 0x1D00
	FLOAT = 0x1406
	FOG = 0x0B60
	FOG_BIT = 0x00000080
	FOG_COLOR = 0x0B66
	FOG_DENSITY = 0x0B62
	FOG_END = 0x0B64
	FOG_HINT = 0x0C54
	FOG_INDEX = 0x0B61
	FOG_MODE = 0x0B65
	FOG_START = 0x0B63
	FRONT = 0x0404
	FRONT_AND_BACK = 0x0408
	FRONT_FACE = 0x0B46
	FRONT_LEFT = 0x0400
	FRONT_RIGHT = 0x0401
	GEQUAL = 0x0206
	GREATER = 0x0204
	GREEN = 0x1904
	GREEN_BIAS = 0x0D19
	GREEN_BITS = 0x0D53
	GREEN_SCALE = 0x0D18
	HINT_BIT = 0x00008000
	INCR = 0x1E02
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
	INT = 0x1404
	INTENSITY = 0x8049
	INTENSITY12 = 0x804C
	INTENSITY16 = 0x804D
	INTENSITY4 = 0x804A
	INTENSITY8 = 0x804B
	INVALID_ENUM = 0x0500
	INVALID_OPERATION = 0x0502
	INVALID_VALUE = 0x0501
	INVERT = 0x150A
	KEEP = 0x1E00
	LEFT = 0x0406
	LEQUAL = 0x0203
	LESS = 0x0201
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
	LINE = 0x1B01
	LINEAR = 0x2601
	LINEAR_ATTENUATION = 0x1208
	LINEAR_MIPMAP_LINEAR = 0x2703
	LINEAR_MIPMAP_NEAREST = 0x2701
	LINES = 0x0001
	LINE_BIT = 0x00000004
	LINE_LOOP = 0x0002
	LINE_RESET_TOKEN = 0x0707
	LINE_SMOOTH = 0x0B20
	LINE_SMOOTH_HINT = 0x0C52
	LINE_STIPPLE = 0x0B24
	LINE_STIPPLE_PATTERN = 0x0B25
	LINE_STIPPLE_REPEAT = 0x0B26
	LINE_STRIP = 0x0003
	LINE_TOKEN = 0x0702
	LINE_WIDTH = 0x0B21
	LINE_WIDTH_GRANULARITY = 0x0B23
	LINE_WIDTH_RANGE = 0x0B22
	LIST_BASE = 0x0B32
	LIST_BIT = 0x00020000
	LIST_INDEX = 0x0B33
	LIST_MODE = 0x0B30
	LOAD = 0x0101
	LOGIC_OP = 0x0BF1
	LOGIC_OP_MODE = 0x0BF0
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
	MAX_TEXTURE_SIZE = 0x0D33
	MAX_TEXTURE_STACK_DEPTH = 0x0D39
	MAX_VIEWPORT_DIMS = 0x0D3A
	MODELVIEW = 0x1700
	MODELVIEW_MATRIX = 0x0BA6
	MODELVIEW_STACK_DEPTH = 0x0BA3
	MODULATE = 0x2100
	MULT = 0x0103
	N3F_V3F = 0x2A25
	NAME_STACK_DEPTH = 0x0D70
	NAND = 0x150E
	NEAREST = 0x2600
	NEAREST_MIPMAP_LINEAR = 0x2702
	NEAREST_MIPMAP_NEAREST = 0x2700
	NEVER = 0x0200
	NICEST = 0x1102
	NONE = 0
	NOOP = 0x1505
	NOR = 0x1508
	NORMALIZE = 0x0BA1
	NORMAL_ARRAY = 0x8075
	NORMAL_ARRAY_POINTER = 0x808F
	NORMAL_ARRAY_STRIDE = 0x807F
	NORMAL_ARRAY_TYPE = 0x807E
	NOTEQUAL = 0x0205
	NO_ERROR = 0
	OBJECT_LINEAR = 0x2401
	OBJECT_PLANE = 0x2501
	ONE = 1
	ONE_MINUS_DST_ALPHA = 0x0305
	ONE_MINUS_DST_COLOR = 0x0307
	ONE_MINUS_SRC_ALPHA = 0x0303
	ONE_MINUS_SRC_COLOR = 0x0301
	OR = 0x1507
	ORDER = 0x0A01
	OR_INVERTED = 0x150D
	OR_REVERSE = 0x150B
	OUT_OF_MEMORY = 0x0505
	PACK_ALIGNMENT = 0x0D05
	PACK_LSB_FIRST = 0x0D01
	PACK_ROW_LENGTH = 0x0D02
	PACK_SKIP_PIXELS = 0x0D04
	PACK_SKIP_ROWS = 0x0D03
	PACK_SWAP_BYTES = 0x0D00
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
	POINT = 0x1B00
	POINTS = 0x0000
	POINT_BIT = 0x00000002
	POINT_SIZE = 0x0B11
	POINT_SIZE_GRANULARITY = 0x0B13
	POINT_SIZE_RANGE = 0x0B12
	POINT_SMOOTH = 0x0B10
	POINT_SMOOTH_HINT = 0x0C51
	POINT_TOKEN = 0x0701
	POLYGON = 0x0009
	POLYGON_BIT = 0x00000008
	POLYGON_MODE = 0x0B40
	POLYGON_OFFSET_FACTOR = 0x8038
	POLYGON_OFFSET_FILL = 0x8037
	POLYGON_OFFSET_LINE = 0x2A02
	POLYGON_OFFSET_POINT = 0x2A01
	POLYGON_OFFSET_UNITS = 0x2A00
	POLYGON_SMOOTH = 0x0B41
	POLYGON_SMOOTH_HINT = 0x0C53
	POLYGON_STIPPLE = 0x0B42
	POLYGON_STIPPLE_BIT = 0x00000010
	POLYGON_TOKEN = 0x0703
	POSITION = 0x1203
	PROJECTION = 0x1701
	PROJECTION_MATRIX = 0x0BA7
	PROJECTION_STACK_DEPTH = 0x0BA4
	PROXY_TEXTURE_1D = 0x8063
	PROXY_TEXTURE_2D = 0x8064
	Q = 0x2003
	QUADRATIC_ATTENUATION = 0x1209
	QUADS = 0x0007
	QUAD_STRIP = 0x0008
	R = 0x2002
	R3_G3_B2 = 0x2A10
	READ_BUFFER = 0x0C02
	RED = 0x1903
	RED_BIAS = 0x0D15
	RED_BITS = 0x0D52
	RED_SCALE = 0x0D14
	RENDER = 0x1C00
	RENDERER = 0x1F01
	RENDER_MODE = 0x0C40
	REPEAT = 0x2901
	REPLACE = 0x1E01
	RETURN = 0x0102
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
	RGBA_MODE = 0x0C31
	RIGHT = 0x0407
	S = 0x2000
	SCISSOR_BIT = 0x00080000
	SCISSOR_BOX = 0x0C10
	SCISSOR_TEST = 0x0C11
	SELECT = 0x1C02
	SELECTION_BUFFER_POINTER = 0x0DF3
	SELECTION_BUFFER_SIZE = 0x0DF4
	SET = 0x150F
	SHADE_MODEL = 0x0B54
	SHININESS = 0x1601
	SHORT = 0x1402
	SMOOTH = 0x1D01
	SPECULAR = 0x1202
	SPHERE_MAP = 0x2402
	SPOT_CUTOFF = 0x1206
	SPOT_DIRECTION = 0x1204
	SPOT_EXPONENT = 0x1205
	SRC_ALPHA = 0x0302
	SRC_ALPHA_SATURATE = 0x0308
	SRC_COLOR = 0x0300
	STACK_OVERFLOW = 0x0503
	STACK_UNDERFLOW = 0x0504
	STENCIL = 0x1802
	STENCIL_BITS = 0x0D57
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
	T = 0x2001
	T2F_C3F_V3F = 0x2A2A
	T2F_C4F_N3F_V3F = 0x2A2C
	T2F_C4UB_V3F = 0x2A29
	T2F_N3F_V3F = 0x2A2B
	T2F_V3F = 0x2A27
	T4F_C4F_N3F_V4F = 0x2A2D
	T4F_V4F = 0x2A28
	TEXTURE = 0x1702
	TEXTURE_1D = 0x0DE0
	TEXTURE_2D = 0x0DE1
	TEXTURE_ALPHA_SIZE = 0x805F
	TEXTURE_BINDING_1D = 0x8068
	TEXTURE_BINDING_2D = 0x8069
	TEXTURE_BIT = 0x00040000
	TEXTURE_BLUE_SIZE = 0x805E
	TEXTURE_BORDER = 0x1005
	TEXTURE_BORDER_COLOR = 0x1004
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
	TEXTURE_GREEN_SIZE = 0x805D
	TEXTURE_HEIGHT = 0x1001
	TEXTURE_INTENSITY_SIZE = 0x8061
	TEXTURE_INTERNAL_FORMAT = 0x1003
	TEXTURE_LUMINANCE_SIZE = 0x8060
	TEXTURE_MAG_FILTER = 0x2800
	TEXTURE_MATRIX = 0x0BA8
	TEXTURE_MIN_FILTER = 0x2801
	TEXTURE_PRIORITY = 0x8066
	TEXTURE_RED_SIZE = 0x805C
	TEXTURE_RESIDENT = 0x8067
	TEXTURE_STACK_DEPTH = 0x0BA5
	TEXTURE_WIDTH = 0x1000
	TEXTURE_WRAP_S = 0x2802
	TEXTURE_WRAP_T = 0x2803
	TRANSFORM_BIT = 0x00001000
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
	V2F = 0x2A20
	V3F = 0x2A21
	VENDOR = 0x1F00
	VERSION = 0x1F02
	VERTEX_ARRAY = 0x8074
	VERTEX_ARRAY_POINTER = 0x808E
	VERTEX_ARRAY_SIZE = 0x807A
	VERTEX_ARRAY_STRIDE = 0x807C
	VERTEX_ARRAY_TYPE = 0x807B
	VIEWPORT = 0x0BA2
	VIEWPORT_BIT = 0x00000800
	XOR = 0x1506
	ZERO = 0
	ZOOM_X = 0x0D16
	ZOOM_Y = 0x0D17
)
// VERSION_1_2
const (
	ALIASED_LINE_WIDTH_RANGE = 0x846E
	ALIASED_POINT_SIZE_RANGE = 0x846D
	BGR = 0x80E0
	BGRA = 0x80E1
	CLAMP_TO_EDGE = 0x812F
	LIGHT_MODEL_COLOR_CONTROL = 0x81F8
	MAX_3D_TEXTURE_SIZE = 0x8073
	MAX_ELEMENTS_INDICES = 0x80E9
	MAX_ELEMENTS_VERTICES = 0x80E8
	PACK_IMAGE_HEIGHT = 0x806C
	PACK_SKIP_IMAGES = 0x806B
	PROXY_TEXTURE_3D = 0x8070
	RESCALE_NORMAL = 0x803A
	SEPARATE_SPECULAR_COLOR = 0x81FA
	SINGLE_COLOR = 0x81F9
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
// VERSION_1_3
const (
	ACTIVE_TEXTURE = 0x84E0
	ADD_SIGNED = 0x8574
	CLAMP_TO_BORDER = 0x812D
	CLIENT_ACTIVE_TEXTURE = 0x84E1
	COMBINE = 0x8570
	COMBINE_ALPHA = 0x8572
	COMBINE_RGB = 0x8571
	COMPRESSED_ALPHA = 0x84E9
	COMPRESSED_INTENSITY = 0x84EC
	COMPRESSED_LUMINANCE = 0x84EA
	COMPRESSED_LUMINANCE_ALPHA = 0x84EB
	COMPRESSED_RGB = 0x84ED
	COMPRESSED_RGBA = 0x84EE
	COMPRESSED_TEXTURE_FORMATS = 0x86A3
	CONSTANT = 0x8576
	DOT3_RGB = 0x86AE
	DOT3_RGBA = 0x86AF
	INTERPOLATE = 0x8575
	MAX_CUBE_MAP_TEXTURE_SIZE = 0x851C
	MAX_TEXTURE_UNITS = 0x84E2
	MULTISAMPLE = 0x809D
	MULTISAMPLE_BIT = 0x20000000
	NORMAL_MAP = 0x8511
	NUM_COMPRESSED_TEXTURE_FORMATS = 0x86A2
	OPERAND0_ALPHA = 0x8598
	OPERAND0_RGB = 0x8590
	OPERAND1_ALPHA = 0x8599
	OPERAND1_RGB = 0x8591
	OPERAND2_ALPHA = 0x859A
	OPERAND2_RGB = 0x8592
	PREVIOUS = 0x8578
	PRIMARY_COLOR = 0x8577
	PROXY_TEXTURE_CUBE_MAP = 0x851B
	REFLECTION_MAP = 0x8512
	RGB_SCALE = 0x8573
	SAMPLES = 0x80A9
	SAMPLE_ALPHA_TO_COVERAGE = 0x809E
	SAMPLE_ALPHA_TO_ONE = 0x809F
	SAMPLE_BUFFERS = 0x80A8
	SAMPLE_COVERAGE = 0x80A0
	SAMPLE_COVERAGE_INVERT = 0x80AB
	SAMPLE_COVERAGE_VALUE = 0x80AA
	SOURCE0_ALPHA = 0x8588
	SOURCE0_RGB = 0x8580
	SOURCE1_ALPHA = 0x8589
	SOURCE1_RGB = 0x8581
	SOURCE2_ALPHA = 0x858A
	SOURCE2_RGB = 0x8582
	SUBTRACT = 0x84E7
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
	COLOR_SUM = 0x8458
	COMPARE_R_TO_TEXTURE = 0x884E
	CURRENT_FOG_COORDINATE = 0x8453
	CURRENT_SECONDARY_COLOR = 0x8459
	DECR_WRAP = 0x8508
	DEPTH_COMPONENT16 = 0x81A5
	DEPTH_COMPONENT24 = 0x81A6
	DEPTH_COMPONENT32 = 0x81A7
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
	INCR_WRAP = 0x8507
	MAX_TEXTURE_LOD_BIAS = 0x84FD
	MIRRORED_REPEAT = 0x8370
	POINT_DISTANCE_ATTENUATION = 0x8129
	POINT_FADE_THRESHOLD_SIZE = 0x8128
	POINT_SIZE_MAX = 0x8127
	POINT_SIZE_MIN = 0x8126
	SECONDARY_COLOR_ARRAY = 0x845E
	SECONDARY_COLOR_ARRAY_POINTER = 0x845D
	SECONDARY_COLOR_ARRAY_SIZE = 0x845A
	SECONDARY_COLOR_ARRAY_STRIDE = 0x845C
	SECONDARY_COLOR_ARRAY_TYPE = 0x845B
	TEXTURE_COMPARE_FUNC = 0x884D
	TEXTURE_COMPARE_MODE = 0x884C
	TEXTURE_DEPTH_SIZE = 0x884A
	TEXTURE_FILTER_CONTROL = 0x8500
	TEXTURE_LOD_BIAS = 0x8501
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
	COLOR_ARRAY_BUFFER_BINDING = 0x8898
	CURRENT_FOG_COORD = 0x8453
	CURRENT_QUERY = 0x8865
	DYNAMIC_COPY = 0x88EA
	DYNAMIC_DRAW = 0x88E8
	DYNAMIC_READ = 0x88E9
	EDGE_FLAG_ARRAY_BUFFER_BINDING = 0x889B
	ELEMENT_ARRAY_BUFFER = 0x8893
	ELEMENT_ARRAY_BUFFER_BINDING = 0x8895
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
	QUERY_COUNTER_BITS = 0x8864
	QUERY_RESULT = 0x8866
	QUERY_RESULT_AVAILABLE = 0x8867
	READ_ONLY = 0x88B8
	READ_WRITE = 0x88BA
	SAMPLES_PASSED = 0x8914
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING = 0x889C
	SRC0_ALPHA = 0x8588
	SRC0_RGB = 0x8580
	SRC1_ALPHA = 0x8589
	SRC1_RGB = 0x8581
	SRC2_ALPHA = 0x858A
	SRC2_RGB = 0x8582
	STATIC_COPY = 0x88E6
	STATIC_DRAW = 0x88E4
	STATIC_READ = 0x88E5
	STREAM_COPY = 0x88E2
	STREAM_DRAW = 0x88E0
	STREAM_READ = 0x88E1
	TEXTURE_COORD_ARRAY_BUFFER_BINDING = 0x889A
	VERTEX_ARRAY_BUFFER_BINDING = 0x8896
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = 0x889F
	WEIGHT_ARRAY_BUFFER_BINDING = 0x889E
	WRITE_ONLY = 0x88B9
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
	COORD_REPLACE = 0x8862
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
	MAX_TEXTURE_COORDS = 0x8871
	MAX_TEXTURE_IMAGE_UNITS = 0x8872
	MAX_VARYING_FLOATS = 0x8B4B
	MAX_VERTEX_ATTRIBS = 0x8869
	MAX_VERTEX_TEXTURE_IMAGE_UNITS = 0x8B4C
	MAX_VERTEX_UNIFORM_COMPONENTS = 0x8B4A
	POINT_SPRITE = 0x8861
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
	VERTEX_PROGRAM_TWO_SIDE = 0x8643
	VERTEX_SHADER = 0x8B31
)
// VERSION_2_1
const (
	COMPRESSED_SLUMINANCE = 0x8C4A
	COMPRESSED_SLUMINANCE_ALPHA = 0x8C4B
	COMPRESSED_SRGB = 0x8C48
	COMPRESSED_SRGB_ALPHA = 0x8C49
	CURRENT_RASTER_SECONDARY_COLOR = 0x845F
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
	SLUMINANCE = 0x8C46
	SLUMINANCE8 = 0x8C47
	SLUMINANCE8_ALPHA8 = 0x8C45
	SLUMINANCE_ALPHA = 0x8C44
	SRGB = 0x8C40
	SRGB8 = 0x8C41
	SRGB8_ALPHA8 = 0x8C43
	SRGB_ALPHA = 0x8C42
)
// VERSION_3_0
const (
	ALPHA_INTEGER = 0x8D97
	BGRA_INTEGER = 0x8D9B
	BGR_INTEGER = 0x8D9A
	BLUE_INTEGER = 0x8D96
	BUFFER_ACCESS_FLAGS = 0x911F
	BUFFER_MAP_LENGTH = 0x9120
	BUFFER_MAP_OFFSET = 0x9121
	CLAMP_FRAGMENT_COLOR = 0x891B
	CLAMP_READ_COLOR = 0x891C
	CLAMP_VERTEX_COLOR = 0x891A
	CLIP_DISTANCE0 = 0x3000
	CLIP_DISTANCE1 = 0x3001
	CLIP_DISTANCE2 = 0x3002
	CLIP_DISTANCE3 = 0x3003
	CLIP_DISTANCE4 = 0x3004
	CLIP_DISTANCE5 = 0x3005
	CLIP_DISTANCE6 = 0x3006
	CLIP_DISTANCE7 = 0x3007
	COLOR_ATTACHMENT0 = 0x8CE0
	COLOR_ATTACHMENT1 = 0x8CE1
	COLOR_ATTACHMENT10 = 0x8CEA
	COLOR_ATTACHMENT11 = 0x8CEB
	COLOR_ATTACHMENT12 = 0x8CEC
	COLOR_ATTACHMENT13 = 0x8CED
	COLOR_ATTACHMENT14 = 0x8CEE
	COLOR_ATTACHMENT15 = 0x8CEF
	COLOR_ATTACHMENT2 = 0x8CE2
	COLOR_ATTACHMENT3 = 0x8CE3
	COLOR_ATTACHMENT4 = 0x8CE4
	COLOR_ATTACHMENT5 = 0x8CE5
	COLOR_ATTACHMENT6 = 0x8CE6
	COLOR_ATTACHMENT7 = 0x8CE7
	COLOR_ATTACHMENT8 = 0x8CE8
	COLOR_ATTACHMENT9 = 0x8CE9
	COMPARE_REF_TO_TEXTURE = 0x884E
	COMPRESSED_RED = 0x8225
	COMPRESSED_RED_RGTC1 = 0x8DBB
	COMPRESSED_RG = 0x8226
	COMPRESSED_RG_RGTC2 = 0x8DBD
	COMPRESSED_SIGNED_RED_RGTC1 = 0x8DBC
	COMPRESSED_SIGNED_RG_RGTC2 = 0x8DBE
	CONTEXT_FLAGS = 0x821E
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT = 0x0001
	DEPTH24_STENCIL8 = 0x88F0
	DEPTH32F_STENCIL8 = 0x8CAD
	DEPTH_ATTACHMENT = 0x8D00
	DEPTH_COMPONENT32F = 0x8CAC
	DEPTH_STENCIL = 0x84F9
	DEPTH_STENCIL_ATTACHMENT = 0x821A
	DRAW_FRAMEBUFFER = 0x8CA9
	DRAW_FRAMEBUFFER_BINDING = 0x8CA6
	FIXED_ONLY = 0x891D
	FLOAT_32_UNSIGNED_INT_24_8_REV = 0x8DAD
	FRAMEBUFFER = 0x8D40
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE = 0x8215
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE = 0x8214
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING = 0x8210
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE = 0x8211
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE = 0x8216
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE = 0x8213
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME = 0x8CD1
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE = 0x8CD0
	FRAMEBUFFER_ATTACHMENT_RED_SIZE = 0x8212
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE = 0x8217
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 0x8CD3
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = 0x8CD2
	FRAMEBUFFER_BINDING = 0x8CA6
	FRAMEBUFFER_COMPLETE = 0x8CD5
	FRAMEBUFFER_DEFAULT = 0x8218
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT = 0x8CD6
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER = 0x8CDB
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = 0x8CD7
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE = 0x8D56
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER = 0x8CDC
	FRAMEBUFFER_SRGB = 0x8DB9
	FRAMEBUFFER_UNDEFINED = 0x8219
	FRAMEBUFFER_UNSUPPORTED = 0x8CDD
	GREEN_INTEGER = 0x8D95
	HALF_FLOAT = 0x140B
	INDEX = 0x8222
	INTERLEAVED_ATTRIBS = 0x8C8C
	INT_SAMPLER_1D = 0x8DC9
	INT_SAMPLER_1D_ARRAY = 0x8DCE
	INT_SAMPLER_2D = 0x8DCA
	INT_SAMPLER_2D_ARRAY = 0x8DCF
	INT_SAMPLER_3D = 0x8DCB
	INT_SAMPLER_CUBE = 0x8DCC
	INVALID_FRAMEBUFFER_OPERATION = 0x0506
	MAJOR_VERSION = 0x821B
	MAP_FLUSH_EXPLICIT_BIT = 0x0010
	MAP_INVALIDATE_BUFFER_BIT = 0x0008
	MAP_INVALIDATE_RANGE_BIT = 0x0004
	MAP_READ_BIT = 0x0001
	MAP_UNSYNCHRONIZED_BIT = 0x0020
	MAP_WRITE_BIT = 0x0002
	MAX_ARRAY_TEXTURE_LAYERS = 0x88FF
	MAX_CLIP_DISTANCES = 0x0D32
	MAX_COLOR_ATTACHMENTS = 0x8CDF
	MAX_PROGRAM_TEXEL_OFFSET = 0x8905
	MAX_RENDERBUFFER_SIZE = 0x84E8
	MAX_SAMPLES = 0x8D57
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS = 0x8C8A
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS = 0x8C8B
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS = 0x8C80
	MINOR_VERSION = 0x821C
	MIN_PROGRAM_TEXEL_OFFSET = 0x8904
	NUM_EXTENSIONS = 0x821D
	PRIMITIVES_GENERATED = 0x8C87
	PROXY_TEXTURE_1D_ARRAY = 0x8C19
	PROXY_TEXTURE_2D_ARRAY = 0x8C1B
	QUERY_BY_REGION_NO_WAIT = 0x8E16
	QUERY_BY_REGION_WAIT = 0x8E15
	QUERY_NO_WAIT = 0x8E14
	QUERY_WAIT = 0x8E13
	R11F_G11F_B10F = 0x8C3A
	R16 = 0x822A
	R16F = 0x822D
	R16I = 0x8233
	R16UI = 0x8234
	R32F = 0x822E
	R32I = 0x8235
	R32UI = 0x8236
	R8 = 0x8229
	R8I = 0x8231
	R8UI = 0x8232
	RASTERIZER_DISCARD = 0x8C89
	READ_FRAMEBUFFER = 0x8CA8
	READ_FRAMEBUFFER_BINDING = 0x8CAA
	RED_INTEGER = 0x8D94
	RENDERBUFFER = 0x8D41
	RENDERBUFFER_ALPHA_SIZE = 0x8D53
	RENDERBUFFER_BINDING = 0x8CA7
	RENDERBUFFER_BLUE_SIZE = 0x8D52
	RENDERBUFFER_DEPTH_SIZE = 0x8D54
	RENDERBUFFER_GREEN_SIZE = 0x8D51
	RENDERBUFFER_HEIGHT = 0x8D43
	RENDERBUFFER_INTERNAL_FORMAT = 0x8D44
	RENDERBUFFER_RED_SIZE = 0x8D50
	RENDERBUFFER_SAMPLES = 0x8CAB
	RENDERBUFFER_STENCIL_SIZE = 0x8D55
	RENDERBUFFER_WIDTH = 0x8D42
	RG = 0x8227
	RG16 = 0x822C
	RG16F = 0x822F
	RG16I = 0x8239
	RG16UI = 0x823A
	RG32F = 0x8230
	RG32I = 0x823B
	RG32UI = 0x823C
	RG8 = 0x822B
	RG8I = 0x8237
	RG8UI = 0x8238
	RGB16F = 0x881B
	RGB16I = 0x8D89
	RGB16UI = 0x8D77
	RGB32F = 0x8815
	RGB32I = 0x8D83
	RGB32UI = 0x8D71
	RGB8I = 0x8D8F
	RGB8UI = 0x8D7D
	RGB9_E5 = 0x8C3D
	RGBA16F = 0x881A
	RGBA16I = 0x8D88
	RGBA16UI = 0x8D76
	RGBA32F = 0x8814
	RGBA32I = 0x8D82
	RGBA32UI = 0x8D70
	RGBA8I = 0x8D8E
	RGBA8UI = 0x8D7C
	RGBA_INTEGER = 0x8D99
	RGB_INTEGER = 0x8D98
	RG_INTEGER = 0x8228
	SAMPLER_1D_ARRAY = 0x8DC0
	SAMPLER_1D_ARRAY_SHADOW = 0x8DC3
	SAMPLER_2D_ARRAY = 0x8DC1
	SAMPLER_2D_ARRAY_SHADOW = 0x8DC4
	SAMPLER_CUBE_SHADOW = 0x8DC5
	SEPARATE_ATTRIBS = 0x8C8D
	STENCIL_ATTACHMENT = 0x8D20
	STENCIL_INDEX1 = 0x8D46
	STENCIL_INDEX16 = 0x8D49
	STENCIL_INDEX4 = 0x8D47
	STENCIL_INDEX8 = 0x8D48
	TEXTURE_1D_ARRAY = 0x8C18
	TEXTURE_2D_ARRAY = 0x8C1A
	TEXTURE_ALPHA_TYPE = 0x8C13
	TEXTURE_BINDING_1D_ARRAY = 0x8C1C
	TEXTURE_BINDING_2D_ARRAY = 0x8C1D
	TEXTURE_BLUE_TYPE = 0x8C12
	TEXTURE_DEPTH_TYPE = 0x8C16
	TEXTURE_GREEN_TYPE = 0x8C11
	TEXTURE_INTENSITY_TYPE = 0x8C15
	TEXTURE_LUMINANCE_TYPE = 0x8C14
	TEXTURE_RED_TYPE = 0x8C10
	TEXTURE_SHARED_SIZE = 0x8C3F
	TEXTURE_STENCIL_SIZE = 0x88F1
	TRANSFORM_FEEDBACK_BUFFER = 0x8C8E
	TRANSFORM_FEEDBACK_BUFFER_BINDING = 0x8C8F
	TRANSFORM_FEEDBACK_BUFFER_MODE = 0x8C7F
	TRANSFORM_FEEDBACK_BUFFER_SIZE = 0x8C85
	TRANSFORM_FEEDBACK_BUFFER_START = 0x8C84
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = 0x8C88
	TRANSFORM_FEEDBACK_VARYINGS = 0x8C83
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH = 0x8C76
	UNSIGNED_INT_10F_11F_11F_REV = 0x8C3B
	UNSIGNED_INT_24_8 = 0x84FA
	UNSIGNED_INT_5_9_9_9_REV = 0x8C3E
	UNSIGNED_INT_SAMPLER_1D = 0x8DD1
	UNSIGNED_INT_SAMPLER_1D_ARRAY = 0x8DD6
	UNSIGNED_INT_SAMPLER_2D = 0x8DD2
	UNSIGNED_INT_SAMPLER_2D_ARRAY = 0x8DD7
	UNSIGNED_INT_SAMPLER_3D = 0x8DD3
	UNSIGNED_INT_SAMPLER_CUBE = 0x8DD4
	UNSIGNED_INT_VEC2 = 0x8DC6
	UNSIGNED_INT_VEC3 = 0x8DC7
	UNSIGNED_INT_VEC4 = 0x8DC8
	UNSIGNED_NORMALIZED = 0x8C17
	VERTEX_ARRAY_BINDING = 0x85B5
	VERTEX_ATTRIB_ARRAY_INTEGER = 0x88FD
)
// VERSION_3_1
const (
	ACTIVE_UNIFORM_BLOCKS = 0x8A36
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH = 0x8A35
	COPY_READ_BUFFER = 0x8F36
	COPY_WRITE_BUFFER = 0x8F37
	INT_SAMPLER_2D_RECT = 0x8DCD
	INT_SAMPLER_BUFFER = 0x8DD0
	INVALID_INDEX = 0xFFFFFFFF
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS = 0x8A33
	MAX_COMBINED_UNIFORM_BLOCKS = 0x8A2E
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS = 0x8A31
	MAX_FRAGMENT_UNIFORM_BLOCKS = 0x8A2D
	MAX_RECTANGLE_TEXTURE_SIZE = 0x84F8
	MAX_TEXTURE_BUFFER_SIZE = 0x8C2B
	MAX_UNIFORM_BLOCK_SIZE = 0x8A30
	MAX_UNIFORM_BUFFER_BINDINGS = 0x8A2F
	MAX_VERTEX_UNIFORM_BLOCKS = 0x8A2B
	PRIMITIVE_RESTART = 0x8F9D
	PRIMITIVE_RESTART_INDEX = 0x8F9E
	PROXY_TEXTURE_RECTANGLE = 0x84F7
	R16_SNORM = 0x8F98
	R8_SNORM = 0x8F94
	RED_SNORM = 0x8F90
	RG16_SNORM = 0x8F99
	RG8_SNORM = 0x8F95
	RGB16_SNORM = 0x8F9A
	RGB8_SNORM = 0x8F96
	RGBA16_SNORM = 0x8F9B
	RGBA8_SNORM = 0x8F97
	RGBA_SNORM = 0x8F93
	RGB_SNORM = 0x8F92
	RG_SNORM = 0x8F91
	SAMPLER_2D_RECT = 0x8B63
	SAMPLER_2D_RECT_SHADOW = 0x8B64
	SAMPLER_BUFFER = 0x8DC2
	SIGNED_NORMALIZED = 0x8F9C
	TEXTURE_BINDING_BUFFER = 0x8C2C
	TEXTURE_BINDING_RECTANGLE = 0x84F6
	TEXTURE_BUFFER = 0x8C2A
	TEXTURE_BUFFER_DATA_STORE_BINDING = 0x8C2D
	TEXTURE_RECTANGLE = 0x84F5
	UNIFORM_ARRAY_STRIDE = 0x8A3C
	UNIFORM_BLOCK_ACTIVE_UNIFORMS = 0x8A42
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES = 0x8A43
	UNIFORM_BLOCK_BINDING = 0x8A3F
	UNIFORM_BLOCK_DATA_SIZE = 0x8A40
	UNIFORM_BLOCK_INDEX = 0x8A3A
	UNIFORM_BLOCK_NAME_LENGTH = 0x8A41
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER = 0x8A46
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER = 0x8A44
	UNIFORM_BUFFER = 0x8A11
	UNIFORM_BUFFER_BINDING = 0x8A28
	UNIFORM_BUFFER_OFFSET_ALIGNMENT = 0x8A34
	UNIFORM_BUFFER_SIZE = 0x8A2A
	UNIFORM_BUFFER_START = 0x8A29
	UNIFORM_IS_ROW_MAJOR = 0x8A3E
	UNIFORM_MATRIX_STRIDE = 0x8A3D
	UNIFORM_NAME_LENGTH = 0x8A39
	UNIFORM_OFFSET = 0x8A3B
	UNIFORM_SIZE = 0x8A38
	UNIFORM_TYPE = 0x8A37
	UNSIGNED_INT_SAMPLER_2D_RECT = 0x8DD5
	UNSIGNED_INT_SAMPLER_BUFFER = 0x8DD8
)
// VERSION_3_2
const (
	ALREADY_SIGNALED = 0x911A
	CONDITION_SATISFIED = 0x911C
	CONTEXT_COMPATIBILITY_PROFILE_BIT = 0x00000002
	CONTEXT_CORE_PROFILE_BIT = 0x00000001
	CONTEXT_PROFILE_MASK = 0x9126
	DEPTH_CLAMP = 0x864F
	FIRST_VERTEX_CONVENTION = 0x8E4D
	FRAMEBUFFER_ATTACHMENT_LAYERED = 0x8DA7
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS = 0x8DA8
	GEOMETRY_INPUT_TYPE = 0x8917
	GEOMETRY_OUTPUT_TYPE = 0x8918
	GEOMETRY_SHADER = 0x8DD9
	GEOMETRY_VERTICES_OUT = 0x8916
	INT_SAMPLER_2D_MULTISAMPLE = 0x9109
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY = 0x910C
	LAST_VERTEX_CONVENTION = 0x8E4E
	LINES_ADJACENCY = 0x000A
	LINE_STRIP_ADJACENCY = 0x000B
	MAX_COLOR_TEXTURE_SAMPLES = 0x910E
	MAX_DEPTH_TEXTURE_SAMPLES = 0x910F
	MAX_FRAGMENT_INPUT_COMPONENTS = 0x9125
	MAX_GEOMETRY_INPUT_COMPONENTS = 0x9123
	MAX_GEOMETRY_OUTPUT_COMPONENTS = 0x9124
	MAX_GEOMETRY_OUTPUT_VERTICES = 0x8DE0
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS = 0x8C29
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS = 0x8DE1
	MAX_GEOMETRY_UNIFORM_COMPONENTS = 0x8DDF
	MAX_INTEGER_SAMPLES = 0x9110
	MAX_SAMPLE_MASK_WORDS = 0x8E59
	MAX_SERVER_WAIT_TIMEOUT = 0x9111
	MAX_VERTEX_OUTPUT_COMPONENTS = 0x9122
	OBJECT_TYPE = 0x9112
	PROGRAM_POINT_SIZE = 0x8642
	PROVOKING_VERTEX = 0x8E4F
	PROXY_TEXTURE_2D_MULTISAMPLE = 0x9101
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY = 0x9103
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION = 0x8E4C
	SAMPLER_2D_MULTISAMPLE = 0x9108
	SAMPLER_2D_MULTISAMPLE_ARRAY = 0x910B
	SAMPLE_MASK = 0x8E51
	SAMPLE_MASK_VALUE = 0x8E52
	SAMPLE_POSITION = 0x8E50
	SIGNALED = 0x9119
	SYNC_CONDITION = 0x9113
	SYNC_FENCE = 0x9116
	SYNC_FLAGS = 0x9115
	SYNC_FLUSH_COMMANDS_BIT = 0x00000001
	SYNC_GPU_COMMANDS_COMPLETE = 0x9117
	SYNC_STATUS = 0x9114
	TEXTURE_2D_MULTISAMPLE = 0x9100
	TEXTURE_2D_MULTISAMPLE_ARRAY = 0x9102
	TEXTURE_BINDING_2D_MULTISAMPLE = 0x9104
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY = 0x9105
	TEXTURE_CUBE_MAP_SEAMLESS = 0x884F
	TEXTURE_FIXED_SAMPLE_LOCATIONS = 0x9107
	TEXTURE_SAMPLES = 0x9106
	TIMEOUT_EXPIRED = 0x911B
	TIMEOUT_IGNORED = 0xFFFFFFFFFFFFFFFF
	TRIANGLES_ADJACENCY = 0x000C
	TRIANGLE_STRIP_ADJACENCY = 0x000D
	UNSIGNALED = 0x9118
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE = 0x910A
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY = 0x910D
	WAIT_FAILED = 0x911D
)
// VERSION_3_3
const (
	ANY_SAMPLES_PASSED = 0x8C2F
	INT_2_10_10_10_REV = 0x8D9F
	MAX_DUAL_SOURCE_DRAW_BUFFERS = 0x88FC
	ONE_MINUS_SRC1_ALPHA = 0x88FB
	ONE_MINUS_SRC1_COLOR = 0x88FA
	RGB10_A2UI = 0x906F
	SAMPLER_BINDING = 0x8919
	SRC1_COLOR = 0x88F9
	TEXTURE_SWIZZLE_A = 0x8E45
	TEXTURE_SWIZZLE_B = 0x8E44
	TEXTURE_SWIZZLE_G = 0x8E43
	TEXTURE_SWIZZLE_R = 0x8E42
	TEXTURE_SWIZZLE_RGBA = 0x8E46
	TIMESTAMP = 0x8E28
	TIME_ELAPSED = 0x88BF
	VERTEX_ATTRIB_ARRAY_DIVISOR = 0x88FE
)
// VERSION_1_0

// http://www.opengl.org/sdk/docs/man3/xhtml/glCullFace.xml
func CullFace(mode Enum)  {
	C.goglCullFace((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFrontFace.xml
func FrontFace(mode Enum)  {
	C.goglFrontFace((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glHint.xml
func Hint(target Enum, mode Enum)  {
	C.goglHint((C.GLenum)(target), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glLineWidth.xml
func LineWidth(width Float)  {
	C.goglLineWidth((C.GLfloat)(width))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glPointSize.xml
func PointSize(size Float)  {
	C.goglPointSize((C.GLfloat)(size))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glPolygonMode.xml
func PolygonMode(face Enum, mode Enum)  {
	C.goglPolygonMode((C.GLenum)(face), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glScissor.xml
func Scissor(x Int, y Int, width Sizei, height Sizei)  {
	C.goglScissor((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterf.xml
func TexParameterf(target Enum, pname Enum, param Float)  {
	C.goglTexParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterfv.xml
func TexParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexParameteri.xml
func TexParameteri(target Enum, pname Enum, param Int)  {
	C.goglTexParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexParameteriv.xml
func TexParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexImage1D.xml
func TexImage1D(target Enum, level Int, internalformat Int, width Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexImage2D.xml
func TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawBuffer.xml
func DrawBuffer(mode Enum)  {
	C.goglDrawBuffer((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClear.xml
func Clear(mask Bitfield)  {
	C.goglClear((C.GLbitfield)(mask))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClearColor.xml
func ClearColor(red Float, green Float, blue Float, alpha Float)  {
	C.goglClearColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClearStencil.xml
func ClearStencil(s Int)  {
	C.goglClearStencil((C.GLint)(s))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClearDepth.xml
func ClearDepth(depth Double)  {
	C.goglClearDepth((C.GLdouble)(depth))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glStencilMask.xml
func StencilMask(mask Uint)  {
	C.goglStencilMask((C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glColorMask.xml
func ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean)  {
	C.goglColorMask((C.GLboolean)(red), (C.GLboolean)(green), (C.GLboolean)(blue), (C.GLboolean)(alpha))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDepthMask.xml
func DepthMask(flag Boolean)  {
	C.goglDepthMask((C.GLboolean)(flag))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDisable.xml
func Disable(cap Enum)  {
	C.goglDisable((C.GLenum)(cap))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glEnable.xml
func Enable(cap Enum)  {
	C.goglEnable((C.GLenum)(cap))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFinish.xml
func Finish()  {
	C.goglFinish()
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFlush.xml
func Flush()  {
	C.goglFlush()
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBlendFunc.xml
func BlendFunc(sfactor Enum, dfactor Enum)  {
	C.goglBlendFunc((C.GLenum)(sfactor), (C.GLenum)(dfactor))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glLogicOp.xml
func LogicOp(opcode Enum)  {
	C.goglLogicOp((C.GLenum)(opcode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glStencilFunc.xml
func StencilFunc(func_ Enum, ref Int, mask Uint)  {
	C.goglStencilFunc((C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glStencilOp.xml
func StencilOp(fail Enum, zfail Enum, zpass Enum)  {
	C.goglStencilOp((C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDepthFunc.xml
func DepthFunc(func_ Enum)  {
	C.goglDepthFunc((C.GLenum)(func_))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glPixelStoref.xml
func PixelStoref(pname Enum, param Float)  {
	C.goglPixelStoref((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glPixelStorei.xml
func PixelStorei(pname Enum, param Int)  {
	C.goglPixelStorei((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glReadBuffer.xml
func ReadBuffer(mode Enum)  {
	C.goglReadBuffer((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glReadPixels.xml
func ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglReadPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetBooleanv.xml
func GetBooleanv(pname Enum, params *Boolean)  {
	C.goglGetBooleanv((C.GLenum)(pname), (*C.GLboolean)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetDoublev.xml
func GetDoublev(pname Enum, params *Double)  {
	C.goglGetDoublev((C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetError.xml
func GetError() Enum {
	return (Enum)(C.goglGetError())
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetFloatv.xml
func GetFloatv(pname Enum, params *Float)  {
	C.goglGetFloatv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetIntegerv.xml
func GetIntegerv(pname Enum, params *Int)  {
	C.goglGetIntegerv((C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetString.xml
func GetString(name Enum) *Ubyte {
	return (*Ubyte)(C.goglGetString((C.GLenum)(name)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetTexImage.xml
func GetTexImage(target Enum, level Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglGetTexImage((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameterfv.xml
func GetTexParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameteriv.xml
func GetTexParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetTexLevelParameterfv.xml
func GetTexLevelParameterfv(target Enum, level Int, pname Enum, params *Float)  {
	C.goglGetTexLevelParameterfv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetTexLevelParameteriv.xml
func GetTexLevelParameteriv(target Enum, level Int, pname Enum, params *Int)  {
	C.goglGetTexLevelParameteriv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsEnabled.xml
func IsEnabled(cap Enum) Boolean {
	return (Boolean)(C.goglIsEnabled((C.GLenum)(cap)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDepthRange.xml
func DepthRange(near_ Double, far_ Double)  {
	C.goglDepthRange((C.GLdouble)(near_), (C.GLdouble)(far_))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glViewport.xml
func Viewport(x Int, y Int, width Sizei, height Sizei)  {
	C.goglViewport((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_1

// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawArrays.xml
func DrawArrays(mode Enum, first Int, count Sizei)  {
	C.goglDrawArrays((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawElements.xml
func DrawElements(mode Enum, count Sizei, type_ Enum, indices Pointer)  {
	C.goglDrawElements((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetPointerv.xml
func GetPointerv(pname Enum, params *Pointer)  {
	C.goglGetPointerv((C.GLenum)(pname), (*unsafe.Pointer)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glPolygonOffset.xml
func PolygonOffset(factor Float, units Float)  {
	C.goglPolygonOffset((C.GLfloat)(factor), (C.GLfloat)(units))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexImage1D.xml
func CopyTexImage1D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, border Int)  {
	C.goglCopyTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexImage2D.xml
func CopyTexImage2D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, height Sizei, border Int)  {
	C.goglCopyTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexSubImage1D.xml
func CopyTexSubImage1D(target Enum, level Int, xoffset Int, x Int, y Int, width Sizei)  {
	C.goglCopyTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexSubImage2D.xml
func CopyTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, x Int, y Int, width Sizei, height Sizei)  {
	C.goglCopyTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexSubImage1D.xml
func TexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexSubImage2D.xml
func TexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindTexture.xml
func BindTexture(target Enum, texture Uint)  {
	C.goglBindTexture((C.GLenum)(target), (C.GLuint)(texture))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteTextures.xml
func DeleteTextures(n Sizei, textures *Uint)  {
	C.goglDeleteTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGenTextures.xml
func GenTextures(n Sizei, textures *Uint)  {
	C.goglGenTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsTexture.xml
func IsTexture(texture Uint) Boolean {
	return (Boolean)(C.goglIsTexture((C.GLuint)(texture)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIndexub.xml
func Indexub(c Ubyte)  {
	C.goglIndexub((C.GLubyte)(c))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIndexubv.xml
func Indexubv(c *Ubyte)  {
	C.goglIndexubv((*C.GLubyte)(c))
}
// VERSION_1_2

// http://www.opengl.org/sdk/docs/man3/xhtml/glBlendColor.xml
func BlendColor(red Float, green Float, blue Float, alpha Float)  {
	C.goglBlendColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBlendEquation.xml
func BlendEquation(mode Enum)  {
	C.goglBlendEquation((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawRangeElements.xml
func DrawRangeElements(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Pointer)  {
	C.goglDrawRangeElements((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexSubImage3D.xml
func TexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCopyTexSubImage3D.xml
func CopyTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, x Int, y Int, width Sizei, height Sizei)  {
	C.goglCopyTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_3

// http://www.opengl.org/sdk/docs/man3/xhtml/glActiveTexture.xml
func ActiveTexture(texture Enum)  {
	C.goglActiveTexture((C.GLenum)(texture))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glSampleCoverage.xml
func SampleCoverage(value Float, invert Boolean)  {
	C.goglSampleCoverage((C.GLfloat)(value), (C.GLboolean)(invert))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexImage3D.xml
func CompressedTexImage3D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, depth Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexImage2D.xml
func CompressedTexImage2D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexImage1D.xml
func CompressedTexImage1D(target Enum, level Int, internalformat Enum, width Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexSubImage3D.xml
func CompressedTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexSubImage2D.xml
func CompressedTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCompressedTexSubImage1D.xml
func CompressedTexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetCompressedTexImage.xml
func GetCompressedTexImage(target Enum, level Int, img Pointer)  {
	C.goglGetCompressedTexImage((C.GLenum)(target), (C.GLint)(level), (unsafe.Pointer)(img))
}
// VERSION_1_4

// http://www.opengl.org/sdk/docs/man3/xhtml/glBlendFuncSeparate.xml
func BlendFuncSeparate(sfactorRGB Enum, dfactorRGB Enum, sfactorAlpha Enum, dfactorAlpha Enum)  {
	C.goglBlendFuncSeparate((C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiDrawArrays.xml
func MultiDrawArrays(mode Enum, first *Int, count *Sizei, drawcount Sizei)  {
	C.goglMultiDrawArrays((C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(drawcount))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiDrawElements.xml
func MultiDrawElements(mode Enum, count *Sizei, type_ Enum, indices *Pointer, drawcount Sizei)  {
	C.goglMultiDrawElements((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(drawcount))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glPointParameterf.xml
func PointParameterf(pname Enum, param Float)  {
	C.goglPointParameterf((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glPointParameterfv.xml
func PointParameterfv(pname Enum, params *Float)  {
	C.goglPointParameterfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glPointParameteri.xml
func PointParameteri(pname Enum, param Int)  {
	C.goglPointParameteri((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glPointParameteriv.xml
func PointParameteriv(pname Enum, params *Int)  {
	C.goglPointParameteriv((C.GLenum)(pname), (*C.GLint)(params))
}
// VERSION_1_5

// http://www.opengl.org/sdk/docs/man3/xhtml/glGenQueries.xml
func GenQueries(n Sizei, ids *Uint)  {
	C.goglGenQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteQueries.xml
func DeleteQueries(n Sizei, ids *Uint)  {
	C.goglDeleteQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsQuery.xml
func IsQuery(id Uint) Boolean {
	return (Boolean)(C.goglIsQuery((C.GLuint)(id)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBeginQuery.xml
func BeginQuery(target Enum, id Uint)  {
	C.goglBeginQuery((C.GLenum)(target), (C.GLuint)(id))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glEndQuery.xml
func EndQuery(target Enum)  {
	C.goglEndQuery((C.GLenum)(target))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetQueryiv.xml
func GetQueryiv(target Enum, pname Enum, params *Int)  {
	C.goglGetQueryiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetQueryObjectiv.xml
func GetQueryObjectiv(id Uint, pname Enum, params *Int)  {
	C.goglGetQueryObjectiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetQueryObjectuiv.xml
func GetQueryObjectuiv(id Uint, pname Enum, params *Uint)  {
	C.goglGetQueryObjectuiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindBuffer.xml
func BindBuffer(target Enum, buffer Uint)  {
	C.goglBindBuffer((C.GLenum)(target), (C.GLuint)(buffer))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteBuffers.xml
func DeleteBuffers(n Sizei, buffers *Uint)  {
	C.goglDeleteBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGenBuffers.xml
func GenBuffers(n Sizei, buffers *Uint)  {
	C.goglGenBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsBuffer.xml
func IsBuffer(buffer Uint) Boolean {
	return (Boolean)(C.goglIsBuffer((C.GLuint)(buffer)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBufferData.xml
func BufferData(target Enum, size Sizeiptr, data Pointer, usage Enum)  {
	C.goglBufferData((C.GLenum)(target), (C.GLsizeiptr)(size), (unsafe.Pointer)(data), (C.GLenum)(usage))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBufferSubData.xml
func BufferSubData(target Enum, offset Intptr, size Sizeiptr, data Pointer)  {
	C.goglBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetBufferSubData.xml
func GetBufferSubData(target Enum, offset Intptr, size Sizeiptr, data Pointer)  {
	C.goglGetBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMapBuffer.xml
func MapBuffer(target Enum, access Enum) Pointer {
	return (Pointer)(C.goglMapBuffer((C.GLenum)(target), (C.GLenum)(access)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUnmapBuffer.xml
func UnmapBuffer(target Enum) Boolean {
	return (Boolean)(C.goglUnmapBuffer((C.GLenum)(target)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetBufferParameteriv.xml
func GetBufferParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetBufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetBufferPointerv.xml
func GetBufferPointerv(target Enum, pname Enum, params *Pointer)  {
	C.goglGetBufferPointerv((C.GLenum)(target), (C.GLenum)(pname), (*unsafe.Pointer)(params))
}
// VERSION_2_0

// http://www.opengl.org/sdk/docs/man3/xhtml/glBlendEquationSeparate.xml
func BlendEquationSeparate(modeRGB Enum, modeAlpha Enum)  {
	C.goglBlendEquationSeparate((C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawBuffers.xml
func DrawBuffers(n Sizei, bufs *Enum)  {
	C.goglDrawBuffers((C.GLsizei)(n), (*C.GLenum)(bufs))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glStencilOpSeparate.xml
func StencilOpSeparate(face Enum, sfail Enum, dpfail Enum, dppass Enum)  {
	C.goglStencilOpSeparate((C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glStencilFuncSeparate.xml
func StencilFuncSeparate(face Enum, func_ Enum, ref Int, mask Uint)  {
	C.goglStencilFuncSeparate((C.GLenum)(face), (C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glStencilMaskSeparate.xml
func StencilMaskSeparate(face Enum, mask Uint)  {
	C.goglStencilMaskSeparate((C.GLenum)(face), (C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glAttachShader.xml
func AttachShader(program Uint, shader Uint)  {
	C.goglAttachShader((C.GLuint)(program), (C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindAttribLocation.xml
func BindAttribLocation(program Uint, index Uint, name *Char)  {
	C.goglBindAttribLocation((C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCompileShader.xml
func CompileShader(shader Uint)  {
	C.goglCompileShader((C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCreateProgram.xml
func CreateProgram() Uint {
	return (Uint)(C.goglCreateProgram())
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCreateShader.xml
func CreateShader(type_ Enum) Uint {
	return (Uint)(C.goglCreateShader((C.GLenum)(type_)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteProgram.xml
func DeleteProgram(program Uint)  {
	C.goglDeleteProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteShader.xml
func DeleteShader(shader Uint)  {
	C.goglDeleteShader((C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDetachShader.xml
func DetachShader(program Uint, shader Uint)  {
	C.goglDetachShader((C.GLuint)(program), (C.GLuint)(shader))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDisableVertexAttribArray.xml
func DisableVertexAttribArray(index Uint)  {
	C.goglDisableVertexAttribArray((C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glEnableVertexAttribArray.xml
func EnableVertexAttribArray(index Uint)  {
	C.goglEnableVertexAttribArray((C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveAttrib.xml
func GetActiveAttrib(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)  {
	C.goglGetActiveAttrib((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveUniform.xml
func GetActiveUniform(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)  {
	C.goglGetActiveUniform((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetAttachedShaders.xml
func GetAttachedShaders(program Uint, maxCount Sizei, count *Sizei, obj *Uint)  {
	C.goglGetAttachedShaders((C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(count), (*C.GLuint)(obj))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetAttribLocation.xml
func GetAttribLocation(program Uint, name *Char) Int {
	return (Int)(C.goglGetAttribLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetProgramiv.xml
func GetProgramiv(program Uint, pname Enum, params *Int)  {
	C.goglGetProgramiv((C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetProgramInfoLog.xml
func GetProgramInfoLog(program Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.goglGetProgramInfoLog((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderiv.xml
func GetShaderiv(shader Uint, pname Enum, params *Int)  {
	C.goglGetShaderiv((C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderInfoLog.xml
func GetShaderInfoLog(shader Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.goglGetShaderInfoLog((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderSource.xml
func GetShaderSource(shader Uint, bufSize Sizei, length *Sizei, source *Char)  {
	C.goglGetShaderSource((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(source))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformLocation.xml
func GetUniformLocation(program Uint, name *Char) Int {
	return (Int)(C.goglGetUniformLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformfv.xml
func GetUniformfv(program Uint, location Int, params *Float)  {
	C.goglGetUniformfv((C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformiv.xml
func GetUniformiv(program Uint, location Int, params *Int)  {
	C.goglGetUniformiv((C.GLuint)(program), (C.GLint)(location), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribdv.xml
func GetVertexAttribdv(index Uint, pname Enum, params *Double)  {
	C.goglGetVertexAttribdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribfv.xml
func GetVertexAttribfv(index Uint, pname Enum, params *Float)  {
	C.goglGetVertexAttribfv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribiv.xml
func GetVertexAttribiv(index Uint, pname Enum, params *Int)  {
	C.goglGetVertexAttribiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribPointerv.xml
func GetVertexAttribPointerv(index Uint, pname Enum, pointer *Pointer)  {
	C.goglGetVertexAttribPointerv((C.GLuint)(index), (C.GLenum)(pname), (*unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsProgram.xml
func IsProgram(program Uint) Boolean {
	return (Boolean)(C.goglIsProgram((C.GLuint)(program)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsShader.xml
func IsShader(shader Uint) Boolean {
	return (Boolean)(C.goglIsShader((C.GLuint)(shader)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glLinkProgram.xml
func LinkProgram(program Uint)  {
	C.goglLinkProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glShaderSource.xml
func ShaderSource(shader Uint, count Sizei, string_ **Char, length *Int)  {
	C.goglShaderSource((C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(string_)), (*C.GLint)(length))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUseProgram.xml
func UseProgram(program Uint)  {
	C.goglUseProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform1f.xml
func Uniform1f(location Int, v0 Float)  {
	C.goglUniform1f((C.GLint)(location), (C.GLfloat)(v0))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform2f.xml
func Uniform2f(location Int, v0 Float, v1 Float)  {
	C.goglUniform2f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform3f.xml
func Uniform3f(location Int, v0 Float, v1 Float, v2 Float)  {
	C.goglUniform3f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform4f.xml
func Uniform4f(location Int, v0 Float, v1 Float, v2 Float, v3 Float)  {
	C.goglUniform4f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform1i.xml
func Uniform1i(location Int, v0 Int)  {
	C.goglUniform1i((C.GLint)(location), (C.GLint)(v0))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform2i.xml
func Uniform2i(location Int, v0 Int, v1 Int)  {
	C.goglUniform2i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform3i.xml
func Uniform3i(location Int, v0 Int, v1 Int, v2 Int)  {
	C.goglUniform3i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform4i.xml
func Uniform4i(location Int, v0 Int, v1 Int, v2 Int, v3 Int)  {
	C.goglUniform4i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform1fv.xml
func Uniform1fv(location Int, count Sizei, value *Float)  {
	C.goglUniform1fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform2fv.xml
func Uniform2fv(location Int, count Sizei, value *Float)  {
	C.goglUniform2fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform3fv.xml
func Uniform3fv(location Int, count Sizei, value *Float)  {
	C.goglUniform3fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform4fv.xml
func Uniform4fv(location Int, count Sizei, value *Float)  {
	C.goglUniform4fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform1iv.xml
func Uniform1iv(location Int, count Sizei, value *Int)  {
	C.goglUniform1iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform2iv.xml
func Uniform2iv(location Int, count Sizei, value *Int)  {
	C.goglUniform2iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform3iv.xml
func Uniform3iv(location Int, count Sizei, value *Int)  {
	C.goglUniform3iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform4iv.xml
func Uniform4iv(location Int, count Sizei, value *Int)  {
	C.goglUniform4iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix2fv.xml
func UniformMatrix2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix3fv.xml
func UniformMatrix3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix4fv.xml
func UniformMatrix4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glValidateProgram.xml
func ValidateProgram(program Uint)  {
	C.goglValidateProgram((C.GLuint)(program))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribPointer.xml
func VertexAttribPointer(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, pointer Pointer)  {
	C.goglVertexAttribPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// VERSION_2_1

// http://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix2x3fv.xml
func UniformMatrix2x3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix3x2fv.xml
func UniformMatrix3x2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix2x4fv.xml
func UniformMatrix2x4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix4x2fv.xml
func UniformMatrix4x2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix3x4fv.xml
func UniformMatrix3x4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniformMatrix4x3fv.xml
func UniformMatrix4x3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// VERSION_3_0

// http://www.opengl.org/sdk/docs/man3/xhtml/glColorMaski.xml
func ColorMaski(index Uint, r Boolean, g Boolean, b Boolean, a Boolean)  {
	C.goglColorMaski((C.GLuint)(index), (C.GLboolean)(r), (C.GLboolean)(g), (C.GLboolean)(b), (C.GLboolean)(a))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetBooleani_v.xml
func GetBooleani_v(target Enum, index Uint, data *Boolean)  {
	C.goglGetBooleani_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLboolean)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetIntegeri_v.xml
func GetIntegeri_v(target Enum, index Uint, data *Int)  {
	C.goglGetIntegeri_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glEnablei.xml
func Enablei(target Enum, index Uint)  {
	C.goglEnablei((C.GLenum)(target), (C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDisablei.xml
func Disablei(target Enum, index Uint)  {
	C.goglDisablei((C.GLenum)(target), (C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsEnabledi.xml
func IsEnabledi(target Enum, index Uint) Boolean {
	return (Boolean)(C.goglIsEnabledi((C.GLenum)(target), (C.GLuint)(index)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBeginTransformFeedback.xml
func BeginTransformFeedback(primitiveMode Enum)  {
	C.goglBeginTransformFeedback((C.GLenum)(primitiveMode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glEndTransformFeedback.xml
func EndTransformFeedback()  {
	C.goglEndTransformFeedback()
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindBufferRange.xml
func BindBufferRange(target Enum, index Uint, buffer Uint, offset Intptr, size Sizeiptr)  {
	C.goglBindBufferRange((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindBufferBase.xml
func BindBufferBase(target Enum, index Uint, buffer Uint)  {
	C.goglBindBufferBase((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTransformFeedbackVaryings.xml
func TransformFeedbackVaryings(program Uint, count Sizei, varyings **Char, bufferMode Enum)  {
	C.goglTransformFeedbackVaryings((C.GLuint)(program), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(varyings)), (C.GLenum)(bufferMode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetTransformFeedbackVarying.xml
func GetTransformFeedbackVarying(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Sizei, type_ *Enum, name *Char)  {
	C.goglGetTransformFeedbackVarying((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLsizei)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClampColor.xml
func ClampColor(target Enum, clamp Enum)  {
	C.goglClampColor((C.GLenum)(target), (C.GLenum)(clamp))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBeginConditionalRender.xml
func BeginConditionalRender(id Uint, mode Enum)  {
	C.goglBeginConditionalRender((C.GLuint)(id), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glEndConditionalRender.xml
func EndConditionalRender()  {
	C.goglEndConditionalRender()
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribIPointer.xml
func VertexAttribIPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglVertexAttribIPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribIiv.xml
func GetVertexAttribIiv(index Uint, pname Enum, params *Int)  {
	C.goglGetVertexAttribIiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetVertexAttribIuiv.xml
func GetVertexAttribIuiv(index Uint, pname Enum, params *Uint)  {
	C.goglGetVertexAttribIuiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI1i.xml
func VertexAttribI1i(index Uint, x Int)  {
	C.goglVertexAttribI1i((C.GLuint)(index), (C.GLint)(x))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI2i.xml
func VertexAttribI2i(index Uint, x Int, y Int)  {
	C.goglVertexAttribI2i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI3i.xml
func VertexAttribI3i(index Uint, x Int, y Int, z Int)  {
	C.goglVertexAttribI3i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4i.xml
func VertexAttribI4i(index Uint, x Int, y Int, z Int, w Int)  {
	C.goglVertexAttribI4i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI1ui.xml
func VertexAttribI1ui(index Uint, x Uint)  {
	C.goglVertexAttribI1ui((C.GLuint)(index), (C.GLuint)(x))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI2ui.xml
func VertexAttribI2ui(index Uint, x Uint, y Uint)  {
	C.goglVertexAttribI2ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI3ui.xml
func VertexAttribI3ui(index Uint, x Uint, y Uint, z Uint)  {
	C.goglVertexAttribI3ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4ui.xml
func VertexAttribI4ui(index Uint, x Uint, y Uint, z Uint, w Uint)  {
	C.goglVertexAttribI4ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI1iv.xml
func VertexAttribI1iv(index Uint, v *Int)  {
	C.goglVertexAttribI1iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI2iv.xml
func VertexAttribI2iv(index Uint, v *Int)  {
	C.goglVertexAttribI2iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI3iv.xml
func VertexAttribI3iv(index Uint, v *Int)  {
	C.goglVertexAttribI3iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4iv.xml
func VertexAttribI4iv(index Uint, v *Int)  {
	C.goglVertexAttribI4iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI1uiv.xml
func VertexAttribI1uiv(index Uint, v *Uint)  {
	C.goglVertexAttribI1uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI2uiv.xml
func VertexAttribI2uiv(index Uint, v *Uint)  {
	C.goglVertexAttribI2uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI3uiv.xml
func VertexAttribI3uiv(index Uint, v *Uint)  {
	C.goglVertexAttribI3uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4uiv.xml
func VertexAttribI4uiv(index Uint, v *Uint)  {
	C.goglVertexAttribI4uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4bv.xml
func VertexAttribI4bv(index Uint, v *Byte)  {
	C.goglVertexAttribI4bv((C.GLuint)(index), (*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4sv.xml
func VertexAttribI4sv(index Uint, v *Short)  {
	C.goglVertexAttribI4sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4ubv.xml
func VertexAttribI4ubv(index Uint, v *Ubyte)  {
	C.goglVertexAttribI4ubv((C.GLuint)(index), (*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribI4usv.xml
func VertexAttribI4usv(index Uint, v *Ushort)  {
	C.goglVertexAttribI4usv((C.GLuint)(index), (*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformuiv.xml
func GetUniformuiv(program Uint, location Int, params *Uint)  {
	C.goglGetUniformuiv((C.GLuint)(program), (C.GLint)(location), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindFragDataLocation.xml
func BindFragDataLocation(program Uint, color Uint, name *Char)  {
	C.goglBindFragDataLocation((C.GLuint)(program), (C.GLuint)(color), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetFragDataLocation.xml
func GetFragDataLocation(program Uint, name *Char) Int {
	return (Int)(C.goglGetFragDataLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform1ui.xml
func Uniform1ui(location Int, v0 Uint)  {
	C.goglUniform1ui((C.GLint)(location), (C.GLuint)(v0))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform2ui.xml
func Uniform2ui(location Int, v0 Uint, v1 Uint)  {
	C.goglUniform2ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform3ui.xml
func Uniform3ui(location Int, v0 Uint, v1 Uint, v2 Uint)  {
	C.goglUniform3ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform4ui.xml
func Uniform4ui(location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint)  {
	C.goglUniform4ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform1uiv.xml
func Uniform1uiv(location Int, count Sizei, value *Uint)  {
	C.goglUniform1uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform2uiv.xml
func Uniform2uiv(location Int, count Sizei, value *Uint)  {
	C.goglUniform2uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform3uiv.xml
func Uniform3uiv(location Int, count Sizei, value *Uint)  {
	C.goglUniform3uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniform4uiv.xml
func Uniform4uiv(location Int, count Sizei, value *Uint)  {
	C.goglUniform4uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterIiv.xml
func TexParameterIiv(target Enum, pname Enum, params *Int)  {
	C.goglTexParameterIiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexParameterIuiv.xml
func TexParameterIuiv(target Enum, pname Enum, params *Uint)  {
	C.goglTexParameterIuiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameterIiv.xml
func GetTexParameterIiv(target Enum, pname Enum, params *Int)  {
	C.goglGetTexParameterIiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetTexParameterIuiv.xml
func GetTexParameterIuiv(target Enum, pname Enum, params *Uint)  {
	C.goglGetTexParameterIuiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClearBufferiv.xml
func ClearBufferiv(buffer Enum, drawbuffer Int, value *Int)  {
	C.goglClearBufferiv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClearBufferuiv.xml
func ClearBufferuiv(buffer Enum, drawbuffer Int, value *Uint)  {
	C.goglClearBufferuiv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClearBufferfv.xml
func ClearBufferfv(buffer Enum, drawbuffer Int, value *Float)  {
	C.goglClearBufferfv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClearBufferfi.xml
func ClearBufferfi(buffer Enum, drawbuffer Int, depth Float, stencil Int)  {
	C.goglClearBufferfi((C.GLenum)(buffer), (C.GLint)(drawbuffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetStringi.xml
func GetStringi(name Enum, index Uint) *Ubyte {
	return (*Ubyte)(C.goglGetStringi((C.GLenum)(name), (C.GLuint)(index)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsRenderbuffer.xml
func IsRenderbuffer(renderbuffer Uint) Boolean {
	return (Boolean)(C.goglIsRenderbuffer((C.GLuint)(renderbuffer)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindRenderbuffer.xml
func BindRenderbuffer(target Enum, renderbuffer Uint)  {
	C.goglBindRenderbuffer((C.GLenum)(target), (C.GLuint)(renderbuffer))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteRenderbuffers.xml
func DeleteRenderbuffers(n Sizei, renderbuffers *Uint)  {
	C.goglDeleteRenderbuffers((C.GLsizei)(n), (*C.GLuint)(renderbuffers))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGenRenderbuffers.xml
func GenRenderbuffers(n Sizei, renderbuffers *Uint)  {
	C.goglGenRenderbuffers((C.GLsizei)(n), (*C.GLuint)(renderbuffers))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glRenderbufferStorage.xml
func RenderbufferStorage(target Enum, internalformat Enum, width Sizei, height Sizei)  {
	C.goglRenderbufferStorage((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetRenderbufferParameteriv.xml
func GetRenderbufferParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetRenderbufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsFramebuffer.xml
func IsFramebuffer(framebuffer Uint) Boolean {
	return (Boolean)(C.goglIsFramebuffer((C.GLuint)(framebuffer)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindFramebuffer.xml
func BindFramebuffer(target Enum, framebuffer Uint)  {
	C.goglBindFramebuffer((C.GLenum)(target), (C.GLuint)(framebuffer))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteFramebuffers.xml
func DeleteFramebuffers(n Sizei, framebuffers *Uint)  {
	C.goglDeleteFramebuffers((C.GLsizei)(n), (*C.GLuint)(framebuffers))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGenFramebuffers.xml
func GenFramebuffers(n Sizei, framebuffers *Uint)  {
	C.goglGenFramebuffers((C.GLsizei)(n), (*C.GLuint)(framebuffers))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCheckFramebufferStatus.xml
func CheckFramebufferStatus(target Enum) Enum {
	return (Enum)(C.goglCheckFramebufferStatus((C.GLenum)(target)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTexture1D.xml
func FramebufferTexture1D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int)  {
	C.goglFramebufferTexture1D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTexture2D.xml
func FramebufferTexture2D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int)  {
	C.goglFramebufferTexture2D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTexture3D.xml
func FramebufferTexture3D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int, zoffset Int)  {
	C.goglFramebufferTexture3D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(zoffset))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferRenderbuffer.xml
func FramebufferRenderbuffer(target Enum, attachment Enum, renderbuffertarget Enum, renderbuffer Uint)  {
	C.goglFramebufferRenderbuffer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetFramebufferAttachmentParameteriv.xml
func GetFramebufferAttachmentParameteriv(target Enum, attachment Enum, pname Enum, params *Int)  {
	C.goglGetFramebufferAttachmentParameteriv((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGenerateMipmap.xml
func GenerateMipmap(target Enum)  {
	C.goglGenerateMipmap((C.GLenum)(target))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBlitFramebuffer.xml
func BlitFramebuffer(srcX0 Int, srcY0 Int, srcX1 Int, srcY1 Int, dstX0 Int, dstY0 Int, dstX1 Int, dstY1 Int, mask Bitfield, filter Enum)  {
	C.goglBlitFramebuffer((C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glRenderbufferStorageMultisample.xml
func RenderbufferStorageMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei)  {
	C.goglRenderbufferStorageMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTextureLayer.xml
func FramebufferTextureLayer(target Enum, attachment Enum, texture Uint, level Int, layer Int)  {
	C.goglFramebufferTextureLayer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMapBufferRange.xml
func MapBufferRange(target Enum, offset Intptr, length Sizeiptr, access Bitfield) Pointer {
	return (Pointer)(C.goglMapBufferRange((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length), (C.GLbitfield)(access)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFlushMappedBufferRange.xml
func FlushMappedBufferRange(target Enum, offset Intptr, length Sizeiptr)  {
	C.goglFlushMappedBufferRange((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindVertexArray.xml
func BindVertexArray(array Uint)  {
	C.goglBindVertexArray((C.GLuint)(array))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteVertexArrays.xml
func DeleteVertexArrays(n Sizei, arrays *Uint)  {
	C.goglDeleteVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGenVertexArrays.xml
func GenVertexArrays(n Sizei, arrays *Uint)  {
	C.goglGenVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsVertexArray.xml
func IsVertexArray(array Uint) Boolean {
	return (Boolean)(C.goglIsVertexArray((C.GLuint)(array)))
}
// VERSION_3_1

// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawArraysInstanced.xml
func DrawArraysInstanced(mode Enum, first Int, count Sizei, instancecount Sizei)  {
	C.goglDrawArraysInstanced((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawElementsInstanced.xml
func DrawElementsInstanced(mode Enum, count Sizei, type_ Enum, indices Pointer, instancecount Sizei)  {
	C.goglDrawElementsInstanced((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexBuffer.xml
func TexBuffer(target Enum, internalformat Enum, buffer Uint)  {
	C.goglTexBuffer((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glPrimitiveRestartIndex.xml
func PrimitiveRestartIndex(index Uint)  {
	C.goglPrimitiveRestartIndex((C.GLuint)(index))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glCopyBufferSubData.xml
func CopyBufferSubData(readTarget Enum, writeTarget Enum, readOffset Intptr, writeOffset Intptr, size Sizeiptr)  {
	C.goglCopyBufferSubData((C.GLenum)(readTarget), (C.GLenum)(writeTarget), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizeiptr)(size))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformIndices.xml
func GetUniformIndices(program Uint, uniformCount Sizei, uniformNames **Char, uniformIndices *Uint)  {
	C.goglGetUniformIndices((C.GLuint)(program), (C.GLsizei)(uniformCount), (**C.GLchar)(unsafe.Pointer(uniformNames)), (*C.GLuint)(uniformIndices))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveUniformsiv.xml
func GetActiveUniformsiv(program Uint, uniformCount Sizei, uniformIndices *Uint, pname Enum, params *Int)  {
	C.goglGetActiveUniformsiv((C.GLuint)(program), (C.GLsizei)(uniformCount), (*C.GLuint)(uniformIndices), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveUniformName.xml
func GetActiveUniformName(program Uint, uniformIndex Uint, bufSize Sizei, length *Sizei, uniformName *Char)  {
	C.goglGetActiveUniformName((C.GLuint)(program), (C.GLuint)(uniformIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(uniformName))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformBlockIndex.xml
func GetUniformBlockIndex(program Uint, uniformBlockName *Char) Uint {
	return (Uint)(C.goglGetUniformBlockIndex((C.GLuint)(program), (*C.GLchar)(uniformBlockName)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveUniformBlockiv.xml
func GetActiveUniformBlockiv(program Uint, uniformBlockIndex Uint, pname Enum, params *Int)  {
	C.goglGetActiveUniformBlockiv((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetActiveUniformBlockName.xml
func GetActiveUniformBlockName(program Uint, uniformBlockIndex Uint, bufSize Sizei, length *Sizei, uniformBlockName *Char)  {
	C.goglGetActiveUniformBlockName((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(uniformBlockName))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glUniformBlockBinding.xml
func UniformBlockBinding(program Uint, uniformBlockIndex Uint, uniformBlockBinding Uint)  {
	C.goglUniformBlockBinding((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLuint)(uniformBlockBinding))
}
// VERSION_3_2

// http://www.opengl.org/sdk/docs/man3/xhtml/glGetInteger64i_v.xml
func GetInteger64i_v(target Enum, index Uint, data *Int64)  {
	C.goglGetInteger64i_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLint64)(data))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetBufferParameteri64v.xml
func GetBufferParameteri64v(target Enum, pname Enum, params *Int64)  {
	C.goglGetBufferParameteri64v((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint64)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFramebufferTexture.xml
func FramebufferTexture(target Enum, attachment Enum, texture Uint, level Int)  {
	C.goglFramebufferTexture((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawElementsBaseVertex.xml
func DrawElementsBaseVertex(mode Enum, count Sizei, type_ Enum, indices Pointer, basevertex Int)  {
	C.goglDrawElementsBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLint)(basevertex))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawRangeElementsBaseVertex.xml
func DrawRangeElementsBaseVertex(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Pointer, basevertex Int)  {
	C.goglDrawRangeElementsBaseVertex((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLint)(basevertex))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawElementsInstancedBaseVertex.xml
func DrawElementsInstancedBaseVertex(mode Enum, count Sizei, type_ Enum, indices Pointer, instancecount Sizei, basevertex Int)  {
	C.goglDrawElementsInstancedBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount), (C.GLint)(basevertex))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiDrawElementsBaseVertex.xml
func MultiDrawElementsBaseVertex(mode Enum, count *Sizei, type_ Enum, indices *Pointer, drawcount Sizei, basevertex *Int)  {
	C.goglMultiDrawElementsBaseVertex((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(drawcount), (*C.GLint)(basevertex))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glProvokingVertex.xml
func ProvokingVertex(mode Enum)  {
	C.goglProvokingVertex((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glFenceSync.xml
func FenceSync(condition Enum, flags Bitfield) Sync {
	return (Sync)(C.goglFenceSync((C.GLenum)(condition), (C.GLbitfield)(flags)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsSync.xml
func IsSync(sync Sync) Boolean {
	return (Boolean)(C.goglIsSync((C.GLsync)(sync)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteSync.xml
func DeleteSync(sync Sync)  {
	C.goglDeleteSync((C.GLsync)(sync))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClientWaitSync.xml
func ClientWaitSync(sync Sync, flags Bitfield, timeout Uint64) Enum {
	return (Enum)(C.goglClientWaitSync((C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glWaitSync.xml
func WaitSync(sync Sync, flags Bitfield, timeout Uint64)  {
	C.goglWaitSync((C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetInteger64v.xml
func GetInteger64v(pname Enum, params *Int64)  {
	C.goglGetInteger64v((C.GLenum)(pname), (*C.GLint64)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetSynciv.xml
func GetSynciv(sync Sync, pname Enum, bufSize Sizei, length *Sizei, values *Int)  {
	C.goglGetSynciv((C.GLsync)(sync), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(values))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexImage2DMultisample.xml
func TexImage2DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, fixedsamplelocations Boolean)  {
	C.goglTexImage2DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(fixedsamplelocations))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexImage3DMultisample.xml
func TexImage3DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean)  {
	C.goglTexImage3DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(fixedsamplelocations))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetMultisamplefv.xml
func GetMultisamplefv(pname Enum, index Uint, val *Float)  {
	C.goglGetMultisamplefv((C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(val))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glSampleMaski.xml
func SampleMaski(index Uint, mask Bitfield)  {
	C.goglSampleMaski((C.GLuint)(index), (C.GLbitfield)(mask))
}
// VERSION_3_3

// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribDivisor.xml
func VertexAttribDivisor(index Uint, divisor Uint)  {
	C.goglVertexAttribDivisor((C.GLuint)(index), (C.GLuint)(divisor))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindFragDataLocationIndexed.xml
func BindFragDataLocationIndexed(program Uint, colorNumber Uint, index Uint, name *Char)  {
	C.goglBindFragDataLocationIndexed((C.GLuint)(program), (C.GLuint)(colorNumber), (C.GLuint)(index), (*C.GLchar)(name))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetFragDataIndex.xml
func GetFragDataIndex(program Uint, name *Char) Int {
	return (Int)(C.goglGetFragDataIndex((C.GLuint)(program), (*C.GLchar)(name)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGenSamplers.xml
func GenSamplers(count Sizei, samplers *Uint)  {
	C.goglGenSamplers((C.GLsizei)(count), (*C.GLuint)(samplers))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteSamplers.xml
func DeleteSamplers(count Sizei, samplers *Uint)  {
	C.goglDeleteSamplers((C.GLsizei)(count), (*C.GLuint)(samplers))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glIsSampler.xml
func IsSampler(sampler Uint) Boolean {
	return (Boolean)(C.goglIsSampler((C.GLuint)(sampler)))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindSampler.xml
func BindSampler(unit Uint, sampler Uint)  {
	C.goglBindSampler((C.GLuint)(unit), (C.GLuint)(sampler))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glSamplerParameteri.xml
func SamplerParameteri(sampler Uint, pname Enum, param Int)  {
	C.goglSamplerParameteri((C.GLuint)(sampler), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glSamplerParameteriv.xml
func SamplerParameteriv(sampler Uint, pname Enum, param *Int)  {
	C.goglSamplerParameteriv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glSamplerParameterf.xml
func SamplerParameterf(sampler Uint, pname Enum, param Float)  {
	C.goglSamplerParameterf((C.GLuint)(sampler), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glSamplerParameterfv.xml
func SamplerParameterfv(sampler Uint, pname Enum, param *Float)  {
	C.goglSamplerParameterfv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glSamplerParameterIiv.xml
func SamplerParameterIiv(sampler Uint, pname Enum, param *Int)  {
	C.goglSamplerParameterIiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glSamplerParameterIuiv.xml
func SamplerParameterIuiv(sampler Uint, pname Enum, param *Uint)  {
	C.goglSamplerParameterIuiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(param))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetSamplerParameteriv.xml
func GetSamplerParameteriv(sampler Uint, pname Enum, params *Int)  {
	C.goglGetSamplerParameteriv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetSamplerParameterIiv.xml
func GetSamplerParameterIiv(sampler Uint, pname Enum, params *Int)  {
	C.goglGetSamplerParameterIiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetSamplerParameterfv.xml
func GetSamplerParameterfv(sampler Uint, pname Enum, params *Float)  {
	C.goglGetSamplerParameterfv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetSamplerParameterIuiv.xml
func GetSamplerParameterIuiv(sampler Uint, pname Enum, params *Uint)  {
	C.goglGetSamplerParameterIuiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glQueryCounter.xml
func QueryCounter(id Uint, target Enum)  {
	C.goglQueryCounter((C.GLuint)(id), (C.GLenum)(target))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetQueryObjecti64v.xml
func GetQueryObjecti64v(id Uint, pname Enum, params *Int64)  {
	C.goglGetQueryObjecti64v((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint64)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glGetQueryObjectui64v.xml
func GetQueryObjectui64v(id Uint, pname Enum, params *Uint64)  {
	C.goglGetQueryObjectui64v((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint64)(params))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexP2ui.xml
func VertexP2ui(type_ Enum, value Uint)  {
	C.goglVertexP2ui((C.GLenum)(type_), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexP2uiv.xml
func VertexP2uiv(type_ Enum, value *Uint)  {
	C.goglVertexP2uiv((C.GLenum)(type_), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexP3ui.xml
func VertexP3ui(type_ Enum, value Uint)  {
	C.goglVertexP3ui((C.GLenum)(type_), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexP3uiv.xml
func VertexP3uiv(type_ Enum, value *Uint)  {
	C.goglVertexP3uiv((C.GLenum)(type_), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexP4ui.xml
func VertexP4ui(type_ Enum, value Uint)  {
	C.goglVertexP4ui((C.GLenum)(type_), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexP4uiv.xml
func VertexP4uiv(type_ Enum, value *Uint)  {
	C.goglVertexP4uiv((C.GLenum)(type_), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexCoordP1ui.xml
func TexCoordP1ui(type_ Enum, coords Uint)  {
	C.goglTexCoordP1ui((C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexCoordP1uiv.xml
func TexCoordP1uiv(type_ Enum, coords *Uint)  {
	C.goglTexCoordP1uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexCoordP2ui.xml
func TexCoordP2ui(type_ Enum, coords Uint)  {
	C.goglTexCoordP2ui((C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexCoordP2uiv.xml
func TexCoordP2uiv(type_ Enum, coords *Uint)  {
	C.goglTexCoordP2uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexCoordP3ui.xml
func TexCoordP3ui(type_ Enum, coords Uint)  {
	C.goglTexCoordP3ui((C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexCoordP3uiv.xml
func TexCoordP3uiv(type_ Enum, coords *Uint)  {
	C.goglTexCoordP3uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexCoordP4ui.xml
func TexCoordP4ui(type_ Enum, coords Uint)  {
	C.goglTexCoordP4ui((C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexCoordP4uiv.xml
func TexCoordP4uiv(type_ Enum, coords *Uint)  {
	C.goglTexCoordP4uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiTexCoordP1ui.xml
func MultiTexCoordP1ui(texture Enum, type_ Enum, coords Uint)  {
	C.goglMultiTexCoordP1ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiTexCoordP1uiv.xml
func MultiTexCoordP1uiv(texture Enum, type_ Enum, coords *Uint)  {
	C.goglMultiTexCoordP1uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiTexCoordP2ui.xml
func MultiTexCoordP2ui(texture Enum, type_ Enum, coords Uint)  {
	C.goglMultiTexCoordP2ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiTexCoordP2uiv.xml
func MultiTexCoordP2uiv(texture Enum, type_ Enum, coords *Uint)  {
	C.goglMultiTexCoordP2uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiTexCoordP3ui.xml
func MultiTexCoordP3ui(texture Enum, type_ Enum, coords Uint)  {
	C.goglMultiTexCoordP3ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiTexCoordP3uiv.xml
func MultiTexCoordP3uiv(texture Enum, type_ Enum, coords *Uint)  {
	C.goglMultiTexCoordP3uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiTexCoordP4ui.xml
func MultiTexCoordP4ui(texture Enum, type_ Enum, coords Uint)  {
	C.goglMultiTexCoordP4ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiTexCoordP4uiv.xml
func MultiTexCoordP4uiv(texture Enum, type_ Enum, coords *Uint)  {
	C.goglMultiTexCoordP4uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glNormalP3ui.xml
func NormalP3ui(type_ Enum, coords Uint)  {
	C.goglNormalP3ui((C.GLenum)(type_), (C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glNormalP3uiv.xml
func NormalP3uiv(type_ Enum, coords *Uint)  {
	C.goglNormalP3uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glColorP3ui.xml
func ColorP3ui(type_ Enum, color Uint)  {
	C.goglColorP3ui((C.GLenum)(type_), (C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glColorP3uiv.xml
func ColorP3uiv(type_ Enum, color *Uint)  {
	C.goglColorP3uiv((C.GLenum)(type_), (*C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glColorP4ui.xml
func ColorP4ui(type_ Enum, color Uint)  {
	C.goglColorP4ui((C.GLenum)(type_), (C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glColorP4uiv.xml
func ColorP4uiv(type_ Enum, color *Uint)  {
	C.goglColorP4uiv((C.GLenum)(type_), (*C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glSecondaryColorP3ui.xml
func SecondaryColorP3ui(type_ Enum, color Uint)  {
	C.goglSecondaryColorP3ui((C.GLenum)(type_), (C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glSecondaryColorP3uiv.xml
func SecondaryColorP3uiv(type_ Enum, color *Uint)  {
	C.goglSecondaryColorP3uiv((C.GLenum)(type_), (*C.GLuint)(color))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribP1ui.xml
func VertexAttribP1ui(index Uint, type_ Enum, normalized Boolean, value Uint)  {
	C.goglVertexAttribP1ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribP1uiv.xml
func VertexAttribP1uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)  {
	C.goglVertexAttribP1uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribP2ui.xml
func VertexAttribP2ui(index Uint, type_ Enum, normalized Boolean, value Uint)  {
	C.goglVertexAttribP2ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribP2uiv.xml
func VertexAttribP2uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)  {
	C.goglVertexAttribP2uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribP3ui.xml
func VertexAttribP3ui(index Uint, type_ Enum, normalized Boolean, value Uint)  {
	C.goglVertexAttribP3ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribP3uiv.xml
func VertexAttribP3uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)  {
	C.goglVertexAttribP3uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribP4ui.xml
func VertexAttribP4ui(index Uint, type_ Enum, normalized Boolean, value Uint)  {
	C.goglVertexAttribP4ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribP4uiv.xml
func VertexAttribP4uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)  {
	C.goglVertexAttribP4uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
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