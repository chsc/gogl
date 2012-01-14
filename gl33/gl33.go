// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// VERSION_3_2
// 
// VERSION_3_3
// 
// VERSION_3_0
// 
// VERSION_3_1
// 
// VERSION_2_1
// 
// VERSION_2_0
// 
// VERSION_1_4
// 
// VERSION_1_5
// 
// VERSION_1_0
// 
// VERSION_1_1
// 
// VERSION_1_2
// 
// VERSION_1_3
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
// // TODO: add context handling header
// #elif defined(_WIN32)
// #define WIN32_LEAN_AND_MEAN 1
// #include <windows.h> // for wglGetProcAddress
// #else
// #include <X11/Xlib.h> // for glXGetProcAddress
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
// void* goglGetProcAddress(const char* name) { 
// #ifdef __APPLE__
// 	return NULL; // TODO: Add get proc addr for Mac.
// #elif _WIN32
// 	return wglGetProcAddress((LPCSTR)name); // TODO: Not tested
// #else
// 	return glXGetProcAddress((const GLubyte*)name);
// #endif
// }
// 
// //  VERSION_3_2
// void (APIENTRYP ptrgoglGetInteger64i_v)(GLenum target, GLuint index, GLint64* data);
// void (APIENTRYP ptrgoglGetBufferParameteri64v)(GLenum target, GLenum pname, GLint64* params);
// void (APIENTRYP ptrgoglFramebufferTexture)(GLenum target, GLenum attachment, GLuint texture, GLint level);
// void (APIENTRYP ptrgoglDrawElementsBaseVertex)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLint basevertex);
// void (APIENTRYP ptrgoglDrawRangeElementsBaseVertex)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices, GLint basevertex);
// void (APIENTRYP ptrgoglDrawElementsInstancedBaseVertex)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei primcount, GLint basevertex);
// void (APIENTRYP ptrgoglMultiDrawElementsBaseVertex)(GLenum mode, GLsizei* count, GLenum type, GLvoid** indices, GLsizei primcount, GLint* basevertex);
// void (APIENTRYP ptrgoglProvokingVertex)(GLenum mode);
// GLsync (APIENTRYP ptrgoglFenceSync)(GLenum condition, GLbitfield flags);
// GLboolean (APIENTRYP ptrgoglIsSync)(GLsync sync);
// void (APIENTRYP ptrgoglDeleteSync)(GLsync sync);
// GLenum (APIENTRYP ptrgoglClientWaitSync)(GLsync sync, GLbitfield flags, GLuint64 timeout);
// void (APIENTRYP ptrgoglWaitSync)(GLsync sync, GLbitfield flags, GLuint64 timeout);
// void (APIENTRYP ptrgoglGetInteger64v)(GLenum pname, GLint64* params);
// void (APIENTRYP ptrgoglGetSynciv)(GLsync sync, GLenum pname, GLsizei bufSize, GLsizei* length, GLint* values);
// void (APIENTRYP ptrgoglTexImage2DMultisample)(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrgoglTexImage3DMultisample)(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrgoglGetMultisamplefv)(GLenum pname, GLuint index, GLfloat* val);
// void (APIENTRYP ptrgoglSampleMaski)(GLuint index, GLbitfield mask);
// //  VERSION_3_3
// void (APIENTRYP ptrgoglVertexAttribDivisor)(GLuint index, GLuint divisor);
// void (APIENTRYP ptrgoglBindFragDataLocationIndexed)(GLuint program, GLuint colorNumber, GLuint index, GLchar* name);
// GLint (APIENTRYP ptrgoglGetFragDataIndex)(GLuint program, GLchar* name);
// void (APIENTRYP ptrgoglGenSamplers)(GLsizei count, GLuint* samplers);
// void (APIENTRYP ptrgoglDeleteSamplers)(GLsizei count, GLuint* samplers);
// GLboolean (APIENTRYP ptrgoglIsSampler)(GLuint sampler);
// void (APIENTRYP ptrgoglBindSampler)(GLuint unit, GLuint sampler);
// void (APIENTRYP ptrgoglSamplerParameteri)(GLuint sampler, GLenum pname, GLint param);
// void (APIENTRYP ptrgoglSamplerParameteriv)(GLuint sampler, GLenum pname, GLint* param);
// void (APIENTRYP ptrgoglSamplerParameterf)(GLuint sampler, GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglSamplerParameterfv)(GLuint sampler, GLenum pname, GLfloat* param);
// void (APIENTRYP ptrgoglSamplerParameterIiv)(GLuint sampler, GLenum pname, GLint* param);
// void (APIENTRYP ptrgoglSamplerParameterIuiv)(GLuint sampler, GLenum pname, GLuint* param);
// void (APIENTRYP ptrgoglGetSamplerParameteriv)(GLuint sampler, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetSamplerParameterIiv)(GLuint sampler, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetSamplerParameterfv)(GLuint sampler, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetSamplerParameterIuiv)(GLuint sampler, GLenum pname, GLuint* params);
// void (APIENTRYP ptrgoglQueryCounter)(GLuint id, GLenum target);
// void (APIENTRYP ptrgoglGetQueryObjecti64v)(GLuint id, GLenum pname, GLint64* params);
// void (APIENTRYP ptrgoglGetQueryObjectui64v)(GLuint id, GLenum pname, GLuint64* params);
// void (APIENTRYP ptrgoglVertexP2ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrgoglVertexP2uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrgoglVertexP3ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrgoglVertexP3uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrgoglVertexP4ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrgoglVertexP4uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrgoglTexCoordP1ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrgoglTexCoordP1uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrgoglTexCoordP2ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrgoglTexCoordP2uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrgoglTexCoordP3ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrgoglTexCoordP3uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrgoglTexCoordP4ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrgoglTexCoordP4uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrgoglMultiTexCoordP1ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrgoglMultiTexCoordP1uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrgoglMultiTexCoordP2ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrgoglMultiTexCoordP2uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrgoglMultiTexCoordP3ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrgoglMultiTexCoordP3uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrgoglMultiTexCoordP4ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrgoglMultiTexCoordP4uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrgoglNormalP3ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrgoglNormalP3uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrgoglColorP3ui)(GLenum type, GLuint color);
// void (APIENTRYP ptrgoglColorP3uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrgoglColorP4ui)(GLenum type, GLuint color);
// void (APIENTRYP ptrgoglColorP4uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrgoglSecondaryColorP3ui)(GLenum type, GLuint color);
// void (APIENTRYP ptrgoglSecondaryColorP3uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrgoglVertexAttribP1ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrgoglVertexAttribP1uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrgoglVertexAttribP2ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrgoglVertexAttribP2uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrgoglVertexAttribP3ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrgoglVertexAttribP3uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrgoglVertexAttribP4ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrgoglVertexAttribP4uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// //  VERSION_3_0
// void (APIENTRYP ptrgoglColorMaski)(GLuint index, GLboolean r, GLboolean g, GLboolean b, GLboolean a);
// void (APIENTRYP ptrgoglGetBooleani_v)(GLenum target, GLuint index, GLboolean* data);
// void (APIENTRYP ptrgoglGetIntegeri_v)(GLenum target, GLuint index, GLint* data);
// void (APIENTRYP ptrgoglEnablei)(GLenum target, GLuint index);
// void (APIENTRYP ptrgoglDisablei)(GLenum target, GLuint index);
// GLboolean (APIENTRYP ptrgoglIsEnabledi)(GLenum target, GLuint index);
// void (APIENTRYP ptrgoglBeginTransformFeedback)(GLenum primitiveMode);
// void (APIENTRYP ptrgoglEndTransformFeedback)();
// void (APIENTRYP ptrgoglBindBufferRange)(GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size);
// void (APIENTRYP ptrgoglBindBufferBase)(GLenum target, GLuint index, GLuint buffer);
// void (APIENTRYP ptrgoglTransformFeedbackVaryings)(GLuint program, GLsizei count, GLchar** varyings, GLenum bufferMode);
// void (APIENTRYP ptrgoglGetTransformFeedbackVarying)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLsizei* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrgoglClampColor)(GLenum target, GLenum clamp);
// void (APIENTRYP ptrgoglBeginConditionalRender)(GLuint id, GLenum mode);
// void (APIENTRYP ptrgoglEndConditionalRender)();
// void (APIENTRYP ptrgoglVertexAttribIPointer)(GLuint index, GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrgoglGetVertexAttribIiv)(GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetVertexAttribIuiv)(GLuint index, GLenum pname, GLuint* params);
// void (APIENTRYP ptrgoglVertexAttribI1i)(GLuint index, GLint x);
// void (APIENTRYP ptrgoglVertexAttribI2i)(GLuint index, GLint x, GLint y);
// void (APIENTRYP ptrgoglVertexAttribI3i)(GLuint index, GLint x, GLint y, GLint z);
// void (APIENTRYP ptrgoglVertexAttribI4i)(GLuint index, GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrgoglVertexAttribI1ui)(GLuint index, GLuint x);
// void (APIENTRYP ptrgoglVertexAttribI2ui)(GLuint index, GLuint x, GLuint y);
// void (APIENTRYP ptrgoglVertexAttribI3ui)(GLuint index, GLuint x, GLuint y, GLuint z);
// void (APIENTRYP ptrgoglVertexAttribI4ui)(GLuint index, GLuint x, GLuint y, GLuint z, GLuint w);
// void (APIENTRYP ptrgoglVertexAttribI1iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrgoglVertexAttribI2iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrgoglVertexAttribI3iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrgoglVertexAttribI4iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrgoglVertexAttribI1uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrgoglVertexAttribI2uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrgoglVertexAttribI3uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrgoglVertexAttribI4uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrgoglVertexAttribI4bv)(GLuint index, GLbyte* v);
// void (APIENTRYP ptrgoglVertexAttribI4sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrgoglVertexAttribI4ubv)(GLuint index, GLubyte* v);
// void (APIENTRYP ptrgoglVertexAttribI4usv)(GLuint index, GLushort* v);
// void (APIENTRYP ptrgoglGetUniformuiv)(GLuint program, GLint location, GLuint* params);
// void (APIENTRYP ptrgoglBindFragDataLocation)(GLuint program, GLuint color, GLchar* name);
// GLint (APIENTRYP ptrgoglGetFragDataLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrgoglUniform1ui)(GLint location, GLuint v0);
// void (APIENTRYP ptrgoglUniform2ui)(GLint location, GLuint v0, GLuint v1);
// void (APIENTRYP ptrgoglUniform3ui)(GLint location, GLuint v0, GLuint v1, GLuint v2);
// void (APIENTRYP ptrgoglUniform4ui)(GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3);
// void (APIENTRYP ptrgoglUniform1uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrgoglUniform2uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrgoglUniform3uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrgoglUniform4uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrgoglTexParameterIiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglTexParameterIuiv)(GLenum target, GLenum pname, GLuint* params);
// void (APIENTRYP ptrgoglGetTexParameterIiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetTexParameterIuiv)(GLenum target, GLenum pname, GLuint* params);
// void (APIENTRYP ptrgoglClearBufferiv)(GLenum buffer, GLint drawbuffer, GLint* value);
// void (APIENTRYP ptrgoglClearBufferuiv)(GLenum buffer, GLint drawbuffer, GLuint* value);
// void (APIENTRYP ptrgoglClearBufferfv)(GLenum buffer, GLint drawbuffer, GLfloat* value);
// void (APIENTRYP ptrgoglClearBufferfi)(GLenum buffer, GLint drawbuffer, GLfloat depth, GLint stencil);
// const GLubyte * (APIENTRYP ptrgoglGetStringi)(GLenum name, GLuint index);
// GLboolean (APIENTRYP ptrgoglIsRenderbuffer)(GLuint renderbuffer);
// void (APIENTRYP ptrgoglBindRenderbuffer)(GLenum target, GLuint renderbuffer);
// void (APIENTRYP ptrgoglDeleteRenderbuffers)(GLsizei n, GLuint* renderbuffers);
// void (APIENTRYP ptrgoglGenRenderbuffers)(GLsizei n, GLuint* renderbuffers);
// void (APIENTRYP ptrgoglRenderbufferStorage)(GLenum target, GLenum internalformat, GLsizei width, GLsizei height);
// void (APIENTRYP ptrgoglGetRenderbufferParameteriv)(GLenum target, GLenum pname, GLint* params);
// GLboolean (APIENTRYP ptrgoglIsFramebuffer)(GLuint framebuffer);
// void (APIENTRYP ptrgoglBindFramebuffer)(GLenum target, GLuint framebuffer);
// void (APIENTRYP ptrgoglDeleteFramebuffers)(GLsizei n, GLuint* framebuffers);
// void (APIENTRYP ptrgoglGenFramebuffers)(GLsizei n, GLuint* framebuffers);
// GLenum (APIENTRYP ptrgoglCheckFramebufferStatus)(GLenum target);
// void (APIENTRYP ptrgoglFramebufferTexture1D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
// void (APIENTRYP ptrgoglFramebufferTexture2D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
// void (APIENTRYP ptrgoglFramebufferTexture3D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint zoffset);
// void (APIENTRYP ptrgoglFramebufferRenderbuffer)(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer);
// void (APIENTRYP ptrgoglGetFramebufferAttachmentParameteriv)(GLenum target, GLenum attachment, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGenerateMipmap)(GLenum target);
// void (APIENTRYP ptrgoglBlitFramebuffer)(GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
// void (APIENTRYP ptrgoglRenderbufferStorageMultisample)(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
// void (APIENTRYP ptrgoglFramebufferTextureLayer)(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer);
// GLvoid* (APIENTRYP ptrgoglMapBufferRange)(GLenum target, GLintptr offset, GLsizeiptr length, GLbitfield access);
// void (APIENTRYP ptrgoglFlushMappedBufferRange)(GLenum target, GLintptr offset, GLsizeiptr length);
// void (APIENTRYP ptrgoglBindVertexArray)(GLuint array);
// void (APIENTRYP ptrgoglDeleteVertexArrays)(GLsizei n, GLuint* arrays);
// void (APIENTRYP ptrgoglGenVertexArrays)(GLsizei n, GLuint* arrays);
// GLboolean (APIENTRYP ptrgoglIsVertexArray)(GLuint array);
// //  VERSION_3_1
// void (APIENTRYP ptrgoglDrawArraysInstanced)(GLenum mode, GLint first, GLsizei count, GLsizei primcount);
// void (APIENTRYP ptrgoglDrawElementsInstanced)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei primcount);
// void (APIENTRYP ptrgoglTexBuffer)(GLenum target, GLenum internalformat, GLuint buffer);
// void (APIENTRYP ptrgoglPrimitiveRestartIndex)(GLuint index);
// void (APIENTRYP ptrgoglCopyBufferSubData)(GLenum readTarget, GLenum writeTarget, GLintptr readOffset, GLintptr writeOffset, GLsizeiptr size);
// void (APIENTRYP ptrgoglGetUniformIndices)(GLuint program, GLsizei uniformCount, GLchar** uniformNames, GLuint* uniformIndices);
// void (APIENTRYP ptrgoglGetActiveUniformsiv)(GLuint program, GLsizei uniformCount, GLuint* uniformIndices, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetActiveUniformName)(GLuint program, GLuint uniformIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformName);
// GLuint (APIENTRYP ptrgoglGetUniformBlockIndex)(GLuint program, GLchar* uniformBlockName);
// void (APIENTRYP ptrgoglGetActiveUniformBlockiv)(GLuint program, GLuint uniformBlockIndex, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetActiveUniformBlockName)(GLuint program, GLuint uniformBlockIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformBlockName);
// void (APIENTRYP ptrgoglUniformBlockBinding)(GLuint program, GLuint uniformBlockIndex, GLuint uniformBlockBinding);
// //  VERSION_2_1
// void (APIENTRYP ptrgoglUniformMatrix2x3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrgoglUniformMatrix3x2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrgoglUniformMatrix2x4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrgoglUniformMatrix4x2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrgoglUniformMatrix3x4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrgoglUniformMatrix4x3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// //  VERSION_2_0
// void (APIENTRYP ptrgoglBlendEquationSeparate)(GLenum modeRGB, GLenum modeAlpha);
// void (APIENTRYP ptrgoglDrawBuffers)(GLsizei n, GLenum* bufs);
// void (APIENTRYP ptrgoglStencilOpSeparate)(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass);
// void (APIENTRYP ptrgoglStencilFuncSeparate)(GLenum face, GLenum func, GLint ref, GLuint mask);
// void (APIENTRYP ptrgoglStencilMaskSeparate)(GLenum face, GLuint mask);
// void (APIENTRYP ptrgoglAttachShader)(GLuint program, GLuint shader);
// void (APIENTRYP ptrgoglBindAttribLocation)(GLuint program, GLuint index, GLchar* name);
// void (APIENTRYP ptrgoglCompileShader)(GLuint shader);
// GLuint (APIENTRYP ptrgoglCreateProgram)();
// GLuint (APIENTRYP ptrgoglCreateShader)(GLenum type);
// void (APIENTRYP ptrgoglDeleteProgram)(GLuint program);
// void (APIENTRYP ptrgoglDeleteShader)(GLuint shader);
// void (APIENTRYP ptrgoglDetachShader)(GLuint program, GLuint shader);
// void (APIENTRYP ptrgoglDisableVertexAttribArray)(GLuint index);
// void (APIENTRYP ptrgoglEnableVertexAttribArray)(GLuint index);
// void (APIENTRYP ptrgoglGetActiveAttrib)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrgoglGetActiveUniform)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrgoglGetAttachedShaders)(GLuint program, GLsizei maxCount, GLsizei* count, GLuint* obj);
// GLint (APIENTRYP ptrgoglGetAttribLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrgoglGetProgramiv)(GLuint program, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetProgramInfoLog)(GLuint program, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrgoglGetShaderiv)(GLuint shader, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetShaderInfoLog)(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrgoglGetShaderSource)(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source);
// GLint (APIENTRYP ptrgoglGetUniformLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrgoglGetUniformfv)(GLuint program, GLint location, GLfloat* params);
// void (APIENTRYP ptrgoglGetUniformiv)(GLuint program, GLint location, GLint* params);
// void (APIENTRYP ptrgoglGetVertexAttribdv)(GLuint index, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrgoglGetVertexAttribfv)(GLuint index, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetVertexAttribiv)(GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetVertexAttribPointerv)(GLuint index, GLenum pname, GLvoid** pointer);
// GLboolean (APIENTRYP ptrgoglIsProgram)(GLuint program);
// GLboolean (APIENTRYP ptrgoglIsShader)(GLuint shader);
// void (APIENTRYP ptrgoglLinkProgram)(GLuint program);
// void (APIENTRYP ptrgoglShaderSource)(GLuint shader, GLsizei count, GLchar** string, GLint* length);
// void (APIENTRYP ptrgoglUseProgram)(GLuint program);
// void (APIENTRYP ptrgoglUniform1f)(GLint location, GLfloat v0);
// void (APIENTRYP ptrgoglUniform2f)(GLint location, GLfloat v0, GLfloat v1);
// void (APIENTRYP ptrgoglUniform3f)(GLint location, GLfloat v0, GLfloat v1, GLfloat v2);
// void (APIENTRYP ptrgoglUniform4f)(GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3);
// void (APIENTRYP ptrgoglUniform1i)(GLint location, GLint v0);
// void (APIENTRYP ptrgoglUniform2i)(GLint location, GLint v0, GLint v1);
// void (APIENTRYP ptrgoglUniform3i)(GLint location, GLint v0, GLint v1, GLint v2);
// void (APIENTRYP ptrgoglUniform4i)(GLint location, GLint v0, GLint v1, GLint v2, GLint v3);
// void (APIENTRYP ptrgoglUniform1fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrgoglUniform2fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrgoglUniform3fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrgoglUniform4fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrgoglUniform1iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrgoglUniform2iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrgoglUniform3iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrgoglUniform4iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrgoglUniformMatrix2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrgoglUniformMatrix3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrgoglUniformMatrix4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrgoglValidateProgram)(GLuint program);
// void (APIENTRYP ptrgoglVertexAttrib1d)(GLuint index, GLdouble x);
// void (APIENTRYP ptrgoglVertexAttrib1dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrgoglVertexAttrib1f)(GLuint index, GLfloat x);
// void (APIENTRYP ptrgoglVertexAttrib1fv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrgoglVertexAttrib1s)(GLuint index, GLshort x);
// void (APIENTRYP ptrgoglVertexAttrib1sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrgoglVertexAttrib2d)(GLuint index, GLdouble x, GLdouble y);
// void (APIENTRYP ptrgoglVertexAttrib2dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrgoglVertexAttrib2f)(GLuint index, GLfloat x, GLfloat y);
// void (APIENTRYP ptrgoglVertexAttrib2fv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrgoglVertexAttrib2s)(GLuint index, GLshort x, GLshort y);
// void (APIENTRYP ptrgoglVertexAttrib2sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrgoglVertexAttrib3d)(GLuint index, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrgoglVertexAttrib3dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrgoglVertexAttrib3f)(GLuint index, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrgoglVertexAttrib3fv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrgoglVertexAttrib3s)(GLuint index, GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrgoglVertexAttrib3sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrgoglVertexAttrib4Nbv)(GLuint index, GLbyte* v);
// void (APIENTRYP ptrgoglVertexAttrib4Niv)(GLuint index, GLint* v);
// void (APIENTRYP ptrgoglVertexAttrib4Nsv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrgoglVertexAttrib4Nub)(GLuint index, GLubyte x, GLubyte y, GLubyte z, GLubyte w);
// void (APIENTRYP ptrgoglVertexAttrib4Nubv)(GLuint index, GLubyte* v);
// void (APIENTRYP ptrgoglVertexAttrib4Nuiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrgoglVertexAttrib4Nusv)(GLuint index, GLushort* v);
// void (APIENTRYP ptrgoglVertexAttrib4bv)(GLuint index, GLbyte* v);
// void (APIENTRYP ptrgoglVertexAttrib4d)(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrgoglVertexAttrib4dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrgoglVertexAttrib4f)(GLuint index, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrgoglVertexAttrib4fv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrgoglVertexAttrib4iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrgoglVertexAttrib4s)(GLuint index, GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrgoglVertexAttrib4sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrgoglVertexAttrib4ubv)(GLuint index, GLubyte* v);
// void (APIENTRYP ptrgoglVertexAttrib4uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrgoglVertexAttrib4usv)(GLuint index, GLushort* v);
// void (APIENTRYP ptrgoglVertexAttribPointer)(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer);
// //  VERSION_1_4
// void (APIENTRYP ptrgoglBlendFuncSeparate)(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha);
// void (APIENTRYP ptrgoglMultiDrawArrays)(GLenum mode, GLint* first, GLsizei* count, GLsizei primcount);
// void (APIENTRYP ptrgoglMultiDrawElements)(GLenum mode, GLsizei* count, GLenum type, GLvoid** indices, GLsizei primcount);
// void (APIENTRYP ptrgoglPointParameterf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglPointParameterfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglPointParameteri)(GLenum pname, GLint param);
// void (APIENTRYP ptrgoglPointParameteriv)(GLenum pname, GLint* params);
// //  VERSION_1_5
// void (APIENTRYP ptrgoglGenQueries)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrgoglDeleteQueries)(GLsizei n, GLuint* ids);
// GLboolean (APIENTRYP ptrgoglIsQuery)(GLuint id);
// void (APIENTRYP ptrgoglBeginQuery)(GLenum target, GLuint id);
// void (APIENTRYP ptrgoglEndQuery)(GLenum target);
// void (APIENTRYP ptrgoglGetQueryiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetQueryObjectiv)(GLuint id, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetQueryObjectuiv)(GLuint id, GLenum pname, GLuint* params);
// void (APIENTRYP ptrgoglBindBuffer)(GLenum target, GLuint buffer);
// void (APIENTRYP ptrgoglDeleteBuffers)(GLsizei n, GLuint* buffers);
// void (APIENTRYP ptrgoglGenBuffers)(GLsizei n, GLuint* buffers);
// GLboolean (APIENTRYP ptrgoglIsBuffer)(GLuint buffer);
// void (APIENTRYP ptrgoglBufferData)(GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage);
// void (APIENTRYP ptrgoglBufferSubData)(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data);
// void (APIENTRYP ptrgoglGetBufferSubData)(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data);
// GLvoid* (APIENTRYP ptrgoglMapBuffer)(GLenum target, GLenum access);
// GLboolean (APIENTRYP ptrgoglUnmapBuffer)(GLenum target);
// void (APIENTRYP ptrgoglGetBufferParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetBufferPointerv)(GLenum target, GLenum pname, GLvoid** params);
// //  VERSION_1_0
// void (APIENTRYP ptrgoglCullFace)(GLenum mode);
// void (APIENTRYP ptrgoglFrontFace)(GLenum mode);
// void (APIENTRYP ptrgoglHint)(GLenum target, GLenum mode);
// void (APIENTRYP ptrgoglLineWidth)(GLfloat width);
// void (APIENTRYP ptrgoglPointSize)(GLfloat size);
// void (APIENTRYP ptrgoglPolygonMode)(GLenum face, GLenum mode);
// void (APIENTRYP ptrgoglScissor)(GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrgoglTexParameterf)(GLenum target, GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglTexParameteri)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrgoglTexParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglTexImage1D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrgoglTexImage2D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrgoglDrawBuffer)(GLenum mode);
// void (APIENTRYP ptrgoglClear)(GLbitfield mask);
// void (APIENTRYP ptrgoglClearColor)(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
// void (APIENTRYP ptrgoglClearStencil)(GLint s);
// void (APIENTRYP ptrgoglClearDepth)(GLclampd depth);
// void (APIENTRYP ptrgoglStencilMask)(GLuint mask);
// void (APIENTRYP ptrgoglColorMask)(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
// void (APIENTRYP ptrgoglDepthMask)(GLboolean flag);
// void (APIENTRYP ptrgoglDisable)(GLenum cap);
// void (APIENTRYP ptrgoglEnable)(GLenum cap);
// void (APIENTRYP ptrgoglFinish)();
// void (APIENTRYP ptrgoglFlush)();
// void (APIENTRYP ptrgoglBlendFunc)(GLenum sfactor, GLenum dfactor);
// void (APIENTRYP ptrgoglLogicOp)(GLenum opcode);
// void (APIENTRYP ptrgoglStencilFunc)(GLenum func, GLint ref, GLuint mask);
// void (APIENTRYP ptrgoglStencilOp)(GLenum fail, GLenum zfail, GLenum zpass);
// void (APIENTRYP ptrgoglDepthFunc)(GLenum func);
// void (APIENTRYP ptrgoglPixelStoref)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglPixelStorei)(GLenum pname, GLint param);
// void (APIENTRYP ptrgoglReadBuffer)(GLenum mode);
// void (APIENTRYP ptrgoglReadPixels)(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrgoglGetBooleanv)(GLenum pname, GLboolean* params);
// void (APIENTRYP ptrgoglGetDoublev)(GLenum pname, GLdouble* params);
// GLenum (APIENTRYP ptrgoglGetError)();
// void (APIENTRYP ptrgoglGetFloatv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetIntegerv)(GLenum pname, GLint* params);
// const GLubyte * (APIENTRYP ptrgoglGetString)(GLenum name);
// void (APIENTRYP ptrgoglGetTexImage)(GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrgoglGetTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetTexParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetTexLevelParameterfv)(GLenum target, GLint level, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetTexLevelParameteriv)(GLenum target, GLint level, GLenum pname, GLint* params);
// GLboolean (APIENTRYP ptrgoglIsEnabled)(GLenum cap);
// void (APIENTRYP ptrgoglDepthRange)(GLclampd near, GLclampd far);
// void (APIENTRYP ptrgoglViewport)(GLint x, GLint y, GLsizei width, GLsizei height);
// //  VERSION_1_1
// void (APIENTRYP ptrgoglDrawArrays)(GLenum mode, GLint first, GLsizei count);
// void (APIENTRYP ptrgoglDrawElements)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices);
// void (APIENTRYP ptrgoglGetPointerv)(GLenum pname, GLvoid** params);
// void (APIENTRYP ptrgoglPolygonOffset)(GLfloat factor, GLfloat units);
// void (APIENTRYP ptrgoglCopyTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border);
// void (APIENTRYP ptrgoglCopyTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border);
// void (APIENTRYP ptrgoglCopyTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrgoglCopyTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrgoglTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrgoglTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrgoglBindTexture)(GLenum target, GLuint texture);
// void (APIENTRYP ptrgoglDeleteTextures)(GLsizei n, GLuint* textures);
// void (APIENTRYP ptrgoglGenTextures)(GLsizei n, GLuint* textures);
// GLboolean (APIENTRYP ptrgoglIsTexture)(GLuint texture);
// //  VERSION_1_2
// void (APIENTRYP ptrgoglBlendColor)(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
// void (APIENTRYP ptrgoglBlendEquation)(GLenum mode);
// void (APIENTRYP ptrgoglDrawRangeElements)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices);
// void (APIENTRYP ptrgoglTexImage3D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrgoglTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrgoglCopyTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height);
// //  VERSION_1_3
// void (APIENTRYP ptrgoglActiveTexture)(GLenum texture);
// void (APIENTRYP ptrgoglSampleCoverage)(GLclampf value, GLboolean invert);
// void (APIENTRYP ptrgoglCompressedTexImage3D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrgoglCompressedTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrgoglCompressedTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrgoglCompressedTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrgoglCompressedTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrgoglCompressedTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrgoglGetCompressedTexImage)(GLenum target, GLint level, GLvoid* img);
// 
// //  VERSION_2_1
// void goglUniformMatrix2x3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrgoglUniformMatrix2x3fv)(location, count, transpose, value);
// }
// void goglUniformMatrix3x2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrgoglUniformMatrix3x2fv)(location, count, transpose, value);
// }
// void goglUniformMatrix2x4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrgoglUniformMatrix2x4fv)(location, count, transpose, value);
// }
// void goglUniformMatrix4x2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrgoglUniformMatrix4x2fv)(location, count, transpose, value);
// }
// void goglUniformMatrix3x4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrgoglUniformMatrix3x4fv)(location, count, transpose, value);
// }
// void goglUniformMatrix4x3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrgoglUniformMatrix4x3fv)(location, count, transpose, value);
// }
// //  VERSION_2_0
// void goglBlendEquationSeparate(GLenum modeRGB, GLenum modeAlpha) {
// 	(*ptrgoglBlendEquationSeparate)(modeRGB, modeAlpha);
// }
// void goglDrawBuffers(GLsizei n, GLenum* bufs) {
// 	(*ptrgoglDrawBuffers)(n, bufs);
// }
// void goglStencilOpSeparate(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
// 	(*ptrgoglStencilOpSeparate)(face, sfail, dpfail, dppass);
// }
// void goglStencilFuncSeparate(GLenum face, GLenum func, GLint ref, GLuint mask) {
// 	(*ptrgoglStencilFuncSeparate)(face, func, ref, mask);
// }
// void goglStencilMaskSeparate(GLenum face, GLuint mask) {
// 	(*ptrgoglStencilMaskSeparate)(face, mask);
// }
// void goglAttachShader(GLuint program, GLuint shader) {
// 	(*ptrgoglAttachShader)(program, shader);
// }
// void goglBindAttribLocation(GLuint program, GLuint index, GLchar* name) {
// 	(*ptrgoglBindAttribLocation)(program, index, name);
// }
// void goglCompileShader(GLuint shader) {
// 	(*ptrgoglCompileShader)(shader);
// }
// GLuint goglCreateProgram() {
// 	return (*ptrgoglCreateProgram)();
// }
// GLuint goglCreateShader(GLenum type) {
// 	return (*ptrgoglCreateShader)(type);
// }
// void goglDeleteProgram(GLuint program) {
// 	(*ptrgoglDeleteProgram)(program);
// }
// void goglDeleteShader(GLuint shader) {
// 	(*ptrgoglDeleteShader)(shader);
// }
// void goglDetachShader(GLuint program, GLuint shader) {
// 	(*ptrgoglDetachShader)(program, shader);
// }
// void goglDisableVertexAttribArray(GLuint index) {
// 	(*ptrgoglDisableVertexAttribArray)(index);
// }
// void goglEnableVertexAttribArray(GLuint index) {
// 	(*ptrgoglEnableVertexAttribArray)(index);
// }
// void goglGetActiveAttrib(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
// 	(*ptrgoglGetActiveAttrib)(program, index, bufSize, length, size, type, name);
// }
// void goglGetActiveUniform(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
// 	(*ptrgoglGetActiveUniform)(program, index, bufSize, length, size, type, name);
// }
// void goglGetAttachedShaders(GLuint program, GLsizei maxCount, GLsizei* count, GLuint* obj) {
// 	(*ptrgoglGetAttachedShaders)(program, maxCount, count, obj);
// }
// GLint goglGetAttribLocation(GLuint program, GLchar* name) {
// 	return (*ptrgoglGetAttribLocation)(program, name);
// }
// void goglGetProgramiv(GLuint program, GLenum pname, GLint* params) {
// 	(*ptrgoglGetProgramiv)(program, pname, params);
// }
// void goglGetProgramInfoLog(GLuint program, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {
// 	(*ptrgoglGetProgramInfoLog)(program, bufSize, length, infoLog);
// }
// void goglGetShaderiv(GLuint shader, GLenum pname, GLint* params) {
// 	(*ptrgoglGetShaderiv)(shader, pname, params);
// }
// void goglGetShaderInfoLog(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {
// 	(*ptrgoglGetShaderInfoLog)(shader, bufSize, length, infoLog);
// }
// void goglGetShaderSource(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
// 	(*ptrgoglGetShaderSource)(shader, bufSize, length, source);
// }
// GLint goglGetUniformLocation(GLuint program, GLchar* name) {
// 	return (*ptrgoglGetUniformLocation)(program, name);
// }
// void goglGetUniformfv(GLuint program, GLint location, GLfloat* params) {
// 	(*ptrgoglGetUniformfv)(program, location, params);
// }
// void goglGetUniformiv(GLuint program, GLint location, GLint* params) {
// 	(*ptrgoglGetUniformiv)(program, location, params);
// }
// void goglGetVertexAttribdv(GLuint index, GLenum pname, GLdouble* params) {
// 	(*ptrgoglGetVertexAttribdv)(index, pname, params);
// }
// void goglGetVertexAttribfv(GLuint index, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetVertexAttribfv)(index, pname, params);
// }
// void goglGetVertexAttribiv(GLuint index, GLenum pname, GLint* params) {
// 	(*ptrgoglGetVertexAttribiv)(index, pname, params);
// }
// void goglGetVertexAttribPointerv(GLuint index, GLenum pname, GLvoid** pointer) {
// 	(*ptrgoglGetVertexAttribPointerv)(index, pname, pointer);
// }
// GLboolean goglIsProgram(GLuint program) {
// 	return (*ptrgoglIsProgram)(program);
// }
// GLboolean goglIsShader(GLuint shader) {
// 	return (*ptrgoglIsShader)(shader);
// }
// void goglLinkProgram(GLuint program) {
// 	(*ptrgoglLinkProgram)(program);
// }
// void goglShaderSource(GLuint shader, GLsizei count, GLchar** string, GLint* length) {
// 	(*ptrgoglShaderSource)(shader, count, string, length);
// }
// void goglUseProgram(GLuint program) {
// 	(*ptrgoglUseProgram)(program);
// }
// void goglUniform1f(GLint location, GLfloat v0) {
// 	(*ptrgoglUniform1f)(location, v0);
// }
// void goglUniform2f(GLint location, GLfloat v0, GLfloat v1) {
// 	(*ptrgoglUniform2f)(location, v0, v1);
// }
// void goglUniform3f(GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
// 	(*ptrgoglUniform3f)(location, v0, v1, v2);
// }
// void goglUniform4f(GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
// 	(*ptrgoglUniform4f)(location, v0, v1, v2, v3);
// }
// void goglUniform1i(GLint location, GLint v0) {
// 	(*ptrgoglUniform1i)(location, v0);
// }
// void goglUniform2i(GLint location, GLint v0, GLint v1) {
// 	(*ptrgoglUniform2i)(location, v0, v1);
// }
// void goglUniform3i(GLint location, GLint v0, GLint v1, GLint v2) {
// 	(*ptrgoglUniform3i)(location, v0, v1, v2);
// }
// void goglUniform4i(GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
// 	(*ptrgoglUniform4i)(location, v0, v1, v2, v3);
// }
// void goglUniform1fv(GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrgoglUniform1fv)(location, count, value);
// }
// void goglUniform2fv(GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrgoglUniform2fv)(location, count, value);
// }
// void goglUniform3fv(GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrgoglUniform3fv)(location, count, value);
// }
// void goglUniform4fv(GLint location, GLsizei count, GLfloat* value) {
// 	(*ptrgoglUniform4fv)(location, count, value);
// }
// void goglUniform1iv(GLint location, GLsizei count, GLint* value) {
// 	(*ptrgoglUniform1iv)(location, count, value);
// }
// void goglUniform2iv(GLint location, GLsizei count, GLint* value) {
// 	(*ptrgoglUniform2iv)(location, count, value);
// }
// void goglUniform3iv(GLint location, GLsizei count, GLint* value) {
// 	(*ptrgoglUniform3iv)(location, count, value);
// }
// void goglUniform4iv(GLint location, GLsizei count, GLint* value) {
// 	(*ptrgoglUniform4iv)(location, count, value);
// }
// void goglUniformMatrix2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrgoglUniformMatrix2fv)(location, count, transpose, value);
// }
// void goglUniformMatrix3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrgoglUniformMatrix3fv)(location, count, transpose, value);
// }
// void goglUniformMatrix4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
// 	(*ptrgoglUniformMatrix4fv)(location, count, transpose, value);
// }
// void goglValidateProgram(GLuint program) {
// 	(*ptrgoglValidateProgram)(program);
// }
// void goglVertexAttrib1d(GLuint index, GLdouble x) {
// 	(*ptrgoglVertexAttrib1d)(index, x);
// }
// void goglVertexAttrib1dv(GLuint index, GLdouble* v) {
// 	(*ptrgoglVertexAttrib1dv)(index, v);
// }
// void goglVertexAttrib1f(GLuint index, GLfloat x) {
// 	(*ptrgoglVertexAttrib1f)(index, x);
// }
// void goglVertexAttrib1fv(GLuint index, GLfloat* v) {
// 	(*ptrgoglVertexAttrib1fv)(index, v);
// }
// void goglVertexAttrib1s(GLuint index, GLshort x) {
// 	(*ptrgoglVertexAttrib1s)(index, x);
// }
// void goglVertexAttrib1sv(GLuint index, GLshort* v) {
// 	(*ptrgoglVertexAttrib1sv)(index, v);
// }
// void goglVertexAttrib2d(GLuint index, GLdouble x, GLdouble y) {
// 	(*ptrgoglVertexAttrib2d)(index, x, y);
// }
// void goglVertexAttrib2dv(GLuint index, GLdouble* v) {
// 	(*ptrgoglVertexAttrib2dv)(index, v);
// }
// void goglVertexAttrib2f(GLuint index, GLfloat x, GLfloat y) {
// 	(*ptrgoglVertexAttrib2f)(index, x, y);
// }
// void goglVertexAttrib2fv(GLuint index, GLfloat* v) {
// 	(*ptrgoglVertexAttrib2fv)(index, v);
// }
// void goglVertexAttrib2s(GLuint index, GLshort x, GLshort y) {
// 	(*ptrgoglVertexAttrib2s)(index, x, y);
// }
// void goglVertexAttrib2sv(GLuint index, GLshort* v) {
// 	(*ptrgoglVertexAttrib2sv)(index, v);
// }
// void goglVertexAttrib3d(GLuint index, GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrgoglVertexAttrib3d)(index, x, y, z);
// }
// void goglVertexAttrib3dv(GLuint index, GLdouble* v) {
// 	(*ptrgoglVertexAttrib3dv)(index, v);
// }
// void goglVertexAttrib3f(GLuint index, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrgoglVertexAttrib3f)(index, x, y, z);
// }
// void goglVertexAttrib3fv(GLuint index, GLfloat* v) {
// 	(*ptrgoglVertexAttrib3fv)(index, v);
// }
// void goglVertexAttrib3s(GLuint index, GLshort x, GLshort y, GLshort z) {
// 	(*ptrgoglVertexAttrib3s)(index, x, y, z);
// }
// void goglVertexAttrib3sv(GLuint index, GLshort* v) {
// 	(*ptrgoglVertexAttrib3sv)(index, v);
// }
// void goglVertexAttrib4Nbv(GLuint index, GLbyte* v) {
// 	(*ptrgoglVertexAttrib4Nbv)(index, v);
// }
// void goglVertexAttrib4Niv(GLuint index, GLint* v) {
// 	(*ptrgoglVertexAttrib4Niv)(index, v);
// }
// void goglVertexAttrib4Nsv(GLuint index, GLshort* v) {
// 	(*ptrgoglVertexAttrib4Nsv)(index, v);
// }
// void goglVertexAttrib4Nub(GLuint index, GLubyte x, GLubyte y, GLubyte z, GLubyte w) {
// 	(*ptrgoglVertexAttrib4Nub)(index, x, y, z, w);
// }
// void goglVertexAttrib4Nubv(GLuint index, GLubyte* v) {
// 	(*ptrgoglVertexAttrib4Nubv)(index, v);
// }
// void goglVertexAttrib4Nuiv(GLuint index, GLuint* v) {
// 	(*ptrgoglVertexAttrib4Nuiv)(index, v);
// }
// void goglVertexAttrib4Nusv(GLuint index, GLushort* v) {
// 	(*ptrgoglVertexAttrib4Nusv)(index, v);
// }
// void goglVertexAttrib4bv(GLuint index, GLbyte* v) {
// 	(*ptrgoglVertexAttrib4bv)(index, v);
// }
// void goglVertexAttrib4d(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrgoglVertexAttrib4d)(index, x, y, z, w);
// }
// void goglVertexAttrib4dv(GLuint index, GLdouble* v) {
// 	(*ptrgoglVertexAttrib4dv)(index, v);
// }
// void goglVertexAttrib4f(GLuint index, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrgoglVertexAttrib4f)(index, x, y, z, w);
// }
// void goglVertexAttrib4fv(GLuint index, GLfloat* v) {
// 	(*ptrgoglVertexAttrib4fv)(index, v);
// }
// void goglVertexAttrib4iv(GLuint index, GLint* v) {
// 	(*ptrgoglVertexAttrib4iv)(index, v);
// }
// void goglVertexAttrib4s(GLuint index, GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrgoglVertexAttrib4s)(index, x, y, z, w);
// }
// void goglVertexAttrib4sv(GLuint index, GLshort* v) {
// 	(*ptrgoglVertexAttrib4sv)(index, v);
// }
// void goglVertexAttrib4ubv(GLuint index, GLubyte* v) {
// 	(*ptrgoglVertexAttrib4ubv)(index, v);
// }
// void goglVertexAttrib4uiv(GLuint index, GLuint* v) {
// 	(*ptrgoglVertexAttrib4uiv)(index, v);
// }
// void goglVertexAttrib4usv(GLuint index, GLushort* v) {
// 	(*ptrgoglVertexAttrib4usv)(index, v);
// }
// void goglVertexAttribPointer(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglVertexAttribPointer)(index, size, type, normalized, stride, pointer);
// }
// //  VERSION_1_4
// void goglBlendFuncSeparate(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha) {
// 	(*ptrgoglBlendFuncSeparate)(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha);
// }
// void goglMultiDrawArrays(GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
// 	(*ptrgoglMultiDrawArrays)(mode, first, count, primcount);
// }
// void goglMultiDrawElements(GLenum mode, GLsizei* count, GLenum type, GLvoid** indices, GLsizei primcount) {
// 	(*ptrgoglMultiDrawElements)(mode, count, type, indices, primcount);
// }
// void goglPointParameterf(GLenum pname, GLfloat param) {
// 	(*ptrgoglPointParameterf)(pname, param);
// }
// void goglPointParameterfv(GLenum pname, GLfloat* params) {
// 	(*ptrgoglPointParameterfv)(pname, params);
// }
// void goglPointParameteri(GLenum pname, GLint param) {
// 	(*ptrgoglPointParameteri)(pname, param);
// }
// void goglPointParameteriv(GLenum pname, GLint* params) {
// 	(*ptrgoglPointParameteriv)(pname, params);
// }
// //  VERSION_1_5
// void goglGenQueries(GLsizei n, GLuint* ids) {
// 	(*ptrgoglGenQueries)(n, ids);
// }
// void goglDeleteQueries(GLsizei n, GLuint* ids) {
// 	(*ptrgoglDeleteQueries)(n, ids);
// }
// GLboolean goglIsQuery(GLuint id) {
// 	return (*ptrgoglIsQuery)(id);
// }
// void goglBeginQuery(GLenum target, GLuint id) {
// 	(*ptrgoglBeginQuery)(target, id);
// }
// void goglEndQuery(GLenum target) {
// 	(*ptrgoglEndQuery)(target);
// }
// void goglGetQueryiv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglGetQueryiv)(target, pname, params);
// }
// void goglGetQueryObjectiv(GLuint id, GLenum pname, GLint* params) {
// 	(*ptrgoglGetQueryObjectiv)(id, pname, params);
// }
// void goglGetQueryObjectuiv(GLuint id, GLenum pname, GLuint* params) {
// 	(*ptrgoglGetQueryObjectuiv)(id, pname, params);
// }
// void goglBindBuffer(GLenum target, GLuint buffer) {
// 	(*ptrgoglBindBuffer)(target, buffer);
// }
// void goglDeleteBuffers(GLsizei n, GLuint* buffers) {
// 	(*ptrgoglDeleteBuffers)(n, buffers);
// }
// void goglGenBuffers(GLsizei n, GLuint* buffers) {
// 	(*ptrgoglGenBuffers)(n, buffers);
// }
// GLboolean goglIsBuffer(GLuint buffer) {
// 	return (*ptrgoglIsBuffer)(buffer);
// }
// void goglBufferData(GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
// 	(*ptrgoglBufferData)(target, size, data, usage);
// }
// void goglBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
// 	(*ptrgoglBufferSubData)(target, offset, size, data);
// }
// void goglGetBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
// 	(*ptrgoglGetBufferSubData)(target, offset, size, data);
// }
// GLvoid* goglMapBuffer(GLenum target, GLenum access) {
// 	return (*ptrgoglMapBuffer)(target, access);
// }
// GLboolean goglUnmapBuffer(GLenum target) {
// 	return (*ptrgoglUnmapBuffer)(target);
// }
// void goglGetBufferParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglGetBufferParameteriv)(target, pname, params);
// }
// void goglGetBufferPointerv(GLenum target, GLenum pname, GLvoid** params) {
// 	(*ptrgoglGetBufferPointerv)(target, pname, params);
// }
// //  VERSION_1_0
// void goglCullFace(GLenum mode) {
// 	(*ptrgoglCullFace)(mode);
// }
// void goglFrontFace(GLenum mode) {
// 	(*ptrgoglFrontFace)(mode);
// }
// void goglHint(GLenum target, GLenum mode) {
// 	(*ptrgoglHint)(target, mode);
// }
// void goglLineWidth(GLfloat width) {
// 	(*ptrgoglLineWidth)(width);
// }
// void goglPointSize(GLfloat size) {
// 	(*ptrgoglPointSize)(size);
// }
// void goglPolygonMode(GLenum face, GLenum mode) {
// 	(*ptrgoglPolygonMode)(face, mode);
// }
// void goglScissor(GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrgoglScissor)(x, y, width, height);
// }
// void goglTexParameterf(GLenum target, GLenum pname, GLfloat param) {
// 	(*ptrgoglTexParameterf)(target, pname, param);
// }
// void goglTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrgoglTexParameterfv)(target, pname, params);
// }
// void goglTexParameteri(GLenum target, GLenum pname, GLint param) {
// 	(*ptrgoglTexParameteri)(target, pname, param);
// }
// void goglTexParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglTexParameteriv)(target, pname, params);
// }
// void goglTexImage1D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
// 	(*ptrgoglTexImage1D)(target, level, internalformat, width, border, format, type, pixels);
// }
// void goglTexImage2D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
// 	(*ptrgoglTexImage2D)(target, level, internalformat, width, height, border, format, type, pixels);
// }
// void goglDrawBuffer(GLenum mode) {
// 	(*ptrgoglDrawBuffer)(mode);
// }
// void goglClear(GLbitfield mask) {
// 	(*ptrgoglClear)(mask);
// }
// void goglClearColor(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
// 	(*ptrgoglClearColor)(red, green, blue, alpha);
// }
// void goglClearStencil(GLint s) {
// 	(*ptrgoglClearStencil)(s);
// }
// void goglClearDepth(GLclampd depth) {
// 	(*ptrgoglClearDepth)(depth);
// }
// void goglStencilMask(GLuint mask) {
// 	(*ptrgoglStencilMask)(mask);
// }
// void goglColorMask(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
// 	(*ptrgoglColorMask)(red, green, blue, alpha);
// }
// void goglDepthMask(GLboolean flag) {
// 	(*ptrgoglDepthMask)(flag);
// }
// void goglDisable(GLenum cap) {
// 	(*ptrgoglDisable)(cap);
// }
// void goglEnable(GLenum cap) {
// 	(*ptrgoglEnable)(cap);
// }
// void goglFinish() {
// 	(*ptrgoglFinish)();
// }
// void goglFlush() {
// 	(*ptrgoglFlush)();
// }
// void goglBlendFunc(GLenum sfactor, GLenum dfactor) {
// 	(*ptrgoglBlendFunc)(sfactor, dfactor);
// }
// void goglLogicOp(GLenum opcode) {
// 	(*ptrgoglLogicOp)(opcode);
// }
// void goglStencilFunc(GLenum func, GLint ref, GLuint mask) {
// 	(*ptrgoglStencilFunc)(func, ref, mask);
// }
// void goglStencilOp(GLenum fail, GLenum zfail, GLenum zpass) {
// 	(*ptrgoglStencilOp)(fail, zfail, zpass);
// }
// void goglDepthFunc(GLenum func) {
// 	(*ptrgoglDepthFunc)(func);
// }
// void goglPixelStoref(GLenum pname, GLfloat param) {
// 	(*ptrgoglPixelStoref)(pname, param);
// }
// void goglPixelStorei(GLenum pname, GLint param) {
// 	(*ptrgoglPixelStorei)(pname, param);
// }
// void goglReadBuffer(GLenum mode) {
// 	(*ptrgoglReadBuffer)(mode);
// }
// void goglReadPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
// 	(*ptrgoglReadPixels)(x, y, width, height, format, type, pixels);
// }
// void goglGetBooleanv(GLenum pname, GLboolean* params) {
// 	(*ptrgoglGetBooleanv)(pname, params);
// }
// void goglGetDoublev(GLenum pname, GLdouble* params) {
// 	(*ptrgoglGetDoublev)(pname, params);
// }
// GLenum goglGetError() {
// 	return (*ptrgoglGetError)();
// }
// void goglGetFloatv(GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetFloatv)(pname, params);
// }
// void goglGetIntegerv(GLenum pname, GLint* params) {
// 	(*ptrgoglGetIntegerv)(pname, params);
// }
// const GLubyte * goglGetString(GLenum name) {
// 	return (*ptrgoglGetString)(name);
// }
// void goglGetTexImage(GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
// 	(*ptrgoglGetTexImage)(target, level, format, type, pixels);
// }
// void goglGetTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetTexParameterfv)(target, pname, params);
// }
// void goglGetTexParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglGetTexParameteriv)(target, pname, params);
// }
// void goglGetTexLevelParameterfv(GLenum target, GLint level, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetTexLevelParameterfv)(target, level, pname, params);
// }
// void goglGetTexLevelParameteriv(GLenum target, GLint level, GLenum pname, GLint* params) {
// 	(*ptrgoglGetTexLevelParameteriv)(target, level, pname, params);
// }
// GLboolean goglIsEnabled(GLenum cap) {
// 	return (*ptrgoglIsEnabled)(cap);
// }
// void goglDepthRange(GLclampd near, GLclampd far) {
// 	(*ptrgoglDepthRange)(near, far);
// }
// void goglViewport(GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrgoglViewport)(x, y, width, height);
// }
// //  VERSION_1_1
// void goglDrawArrays(GLenum mode, GLint first, GLsizei count) {
// 	(*ptrgoglDrawArrays)(mode, first, count);
// }
// void goglDrawElements(GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
// 	(*ptrgoglDrawElements)(mode, count, type, indices);
// }
// void goglGetPointerv(GLenum pname, GLvoid** params) {
// 	(*ptrgoglGetPointerv)(pname, params);
// }
// void goglPolygonOffset(GLfloat factor, GLfloat units) {
// 	(*ptrgoglPolygonOffset)(factor, units);
// }
// void goglCopyTexImage1D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border) {
// 	(*ptrgoglCopyTexImage1D)(target, level, internalformat, x, y, width, border);
// }
// void goglCopyTexImage2D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
// 	(*ptrgoglCopyTexImage2D)(target, level, internalformat, x, y, width, height, border);
// }
// void goglCopyTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
// 	(*ptrgoglCopyTexSubImage1D)(target, level, xoffset, x, y, width);
// }
// void goglCopyTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrgoglCopyTexSubImage2D)(target, level, xoffset, yoffset, x, y, width, height);
// }
// void goglTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
// 	(*ptrgoglTexSubImage1D)(target, level, xoffset, width, format, type, pixels);
// }
// void goglTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
// 	(*ptrgoglTexSubImage2D)(target, level, xoffset, yoffset, width, height, format, type, pixels);
// }
// void goglBindTexture(GLenum target, GLuint texture) {
// 	(*ptrgoglBindTexture)(target, texture);
// }
// void goglDeleteTextures(GLsizei n, GLuint* textures) {
// 	(*ptrgoglDeleteTextures)(n, textures);
// }
// void goglGenTextures(GLsizei n, GLuint* textures) {
// 	(*ptrgoglGenTextures)(n, textures);
// }
// GLboolean goglIsTexture(GLuint texture) {
// 	return (*ptrgoglIsTexture)(texture);
// }
// //  VERSION_1_2
// void goglBlendColor(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
// 	(*ptrgoglBlendColor)(red, green, blue, alpha);
// }
// void goglBlendEquation(GLenum mode) {
// 	(*ptrgoglBlendEquation)(mode);
// }
// void goglDrawRangeElements(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
// 	(*ptrgoglDrawRangeElements)(mode, start, end, count, type, indices);
// }
// void goglTexImage3D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
// 	(*ptrgoglTexImage3D)(target, level, internalformat, width, height, depth, border, format, type, pixels);
// }
// void goglTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
// 	(*ptrgoglTexSubImage3D)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
// }
// void goglCopyTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrgoglCopyTexSubImage3D)(target, level, xoffset, yoffset, zoffset, x, y, width, height);
// }
// //  VERSION_1_3
// void goglActiveTexture(GLenum texture) {
// 	(*ptrgoglActiveTexture)(texture);
// }
// void goglSampleCoverage(GLclampf value, GLboolean invert) {
// 	(*ptrgoglSampleCoverage)(value, invert);
// }
// void goglCompressedTexImage3D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
// 	(*ptrgoglCompressedTexImage3D)(target, level, internalformat, width, height, depth, border, imageSize, data);
// }
// void goglCompressedTexImage2D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
// 	(*ptrgoglCompressedTexImage2D)(target, level, internalformat, width, height, border, imageSize, data);
// }
// void goglCompressedTexImage1D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
// 	(*ptrgoglCompressedTexImage1D)(target, level, internalformat, width, border, imageSize, data);
// }
// void goglCompressedTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
// 	(*ptrgoglCompressedTexSubImage3D)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
// }
// void goglCompressedTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
// 	(*ptrgoglCompressedTexSubImage2D)(target, level, xoffset, yoffset, width, height, format, imageSize, data);
// }
// void goglCompressedTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
// 	(*ptrgoglCompressedTexSubImage1D)(target, level, xoffset, width, format, imageSize, data);
// }
// void goglGetCompressedTexImage(GLenum target, GLint level, GLvoid* img) {
// 	(*ptrgoglGetCompressedTexImage)(target, level, img);
// }
// //  VERSION_3_2
// void goglGetInteger64i_v(GLenum target, GLuint index, GLint64* data) {
// 	(*ptrgoglGetInteger64i_v)(target, index, data);
// }
// void goglGetBufferParameteri64v(GLenum target, GLenum pname, GLint64* params) {
// 	(*ptrgoglGetBufferParameteri64v)(target, pname, params);
// }
// void goglFramebufferTexture(GLenum target, GLenum attachment, GLuint texture, GLint level) {
// 	(*ptrgoglFramebufferTexture)(target, attachment, texture, level);
// }
// void goglDrawElementsBaseVertex(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLint basevertex) {
// 	(*ptrgoglDrawElementsBaseVertex)(mode, count, type, indices, basevertex);
// }
// void goglDrawRangeElementsBaseVertex(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices, GLint basevertex) {
// 	(*ptrgoglDrawRangeElementsBaseVertex)(mode, start, end, count, type, indices, basevertex);
// }
// void goglDrawElementsInstancedBaseVertex(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei primcount, GLint basevertex) {
// 	(*ptrgoglDrawElementsInstancedBaseVertex)(mode, count, type, indices, primcount, basevertex);
// }
// void goglMultiDrawElementsBaseVertex(GLenum mode, GLsizei* count, GLenum type, GLvoid** indices, GLsizei primcount, GLint* basevertex) {
// 	(*ptrgoglMultiDrawElementsBaseVertex)(mode, count, type, indices, primcount, basevertex);
// }
// void goglProvokingVertex(GLenum mode) {
// 	(*ptrgoglProvokingVertex)(mode);
// }
// GLsync goglFenceSync(GLenum condition, GLbitfield flags) {
// 	return (*ptrgoglFenceSync)(condition, flags);
// }
// GLboolean goglIsSync(GLsync sync) {
// 	return (*ptrgoglIsSync)(sync);
// }
// void goglDeleteSync(GLsync sync) {
// 	(*ptrgoglDeleteSync)(sync);
// }
// GLenum goglClientWaitSync(GLsync sync, GLbitfield flags, GLuint64 timeout) {
// 	return (*ptrgoglClientWaitSync)(sync, flags, timeout);
// }
// void goglWaitSync(GLsync sync, GLbitfield flags, GLuint64 timeout) {
// 	(*ptrgoglWaitSync)(sync, flags, timeout);
// }
// void goglGetInteger64v(GLenum pname, GLint64* params) {
// 	(*ptrgoglGetInteger64v)(pname, params);
// }
// void goglGetSynciv(GLsync sync, GLenum pname, GLsizei bufSize, GLsizei* length, GLint* values) {
// 	(*ptrgoglGetSynciv)(sync, pname, bufSize, length, values);
// }
// void goglTexImage2DMultisample(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations) {
// 	(*ptrgoglTexImage2DMultisample)(target, samples, internalformat, width, height, fixedsamplelocations);
// }
// void goglTexImage3DMultisample(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations) {
// 	(*ptrgoglTexImage3DMultisample)(target, samples, internalformat, width, height, depth, fixedsamplelocations);
// }
// void goglGetMultisamplefv(GLenum pname, GLuint index, GLfloat* val) {
// 	(*ptrgoglGetMultisamplefv)(pname, index, val);
// }
// void goglSampleMaski(GLuint index, GLbitfield mask) {
// 	(*ptrgoglSampleMaski)(index, mask);
// }
// //  VERSION_3_3
// void goglVertexAttribDivisor(GLuint index, GLuint divisor) {
// 	(*ptrgoglVertexAttribDivisor)(index, divisor);
// }
// void goglBindFragDataLocationIndexed(GLuint program, GLuint colorNumber, GLuint index, GLchar* name) {
// 	(*ptrgoglBindFragDataLocationIndexed)(program, colorNumber, index, name);
// }
// GLint goglGetFragDataIndex(GLuint program, GLchar* name) {
// 	return (*ptrgoglGetFragDataIndex)(program, name);
// }
// void goglGenSamplers(GLsizei count, GLuint* samplers) {
// 	(*ptrgoglGenSamplers)(count, samplers);
// }
// void goglDeleteSamplers(GLsizei count, GLuint* samplers) {
// 	(*ptrgoglDeleteSamplers)(count, samplers);
// }
// GLboolean goglIsSampler(GLuint sampler) {
// 	return (*ptrgoglIsSampler)(sampler);
// }
// void goglBindSampler(GLuint unit, GLuint sampler) {
// 	(*ptrgoglBindSampler)(unit, sampler);
// }
// void goglSamplerParameteri(GLuint sampler, GLenum pname, GLint param) {
// 	(*ptrgoglSamplerParameteri)(sampler, pname, param);
// }
// void goglSamplerParameteriv(GLuint sampler, GLenum pname, GLint* param) {
// 	(*ptrgoglSamplerParameteriv)(sampler, pname, param);
// }
// void goglSamplerParameterf(GLuint sampler, GLenum pname, GLfloat param) {
// 	(*ptrgoglSamplerParameterf)(sampler, pname, param);
// }
// void goglSamplerParameterfv(GLuint sampler, GLenum pname, GLfloat* param) {
// 	(*ptrgoglSamplerParameterfv)(sampler, pname, param);
// }
// void goglSamplerParameterIiv(GLuint sampler, GLenum pname, GLint* param) {
// 	(*ptrgoglSamplerParameterIiv)(sampler, pname, param);
// }
// void goglSamplerParameterIuiv(GLuint sampler, GLenum pname, GLuint* param) {
// 	(*ptrgoglSamplerParameterIuiv)(sampler, pname, param);
// }
// void goglGetSamplerParameteriv(GLuint sampler, GLenum pname, GLint* params) {
// 	(*ptrgoglGetSamplerParameteriv)(sampler, pname, params);
// }
// void goglGetSamplerParameterIiv(GLuint sampler, GLenum pname, GLint* params) {
// 	(*ptrgoglGetSamplerParameterIiv)(sampler, pname, params);
// }
// void goglGetSamplerParameterfv(GLuint sampler, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetSamplerParameterfv)(sampler, pname, params);
// }
// void goglGetSamplerParameterIuiv(GLuint sampler, GLenum pname, GLuint* params) {
// 	(*ptrgoglGetSamplerParameterIuiv)(sampler, pname, params);
// }
// void goglQueryCounter(GLuint id, GLenum target) {
// 	(*ptrgoglQueryCounter)(id, target);
// }
// void goglGetQueryObjecti64v(GLuint id, GLenum pname, GLint64* params) {
// 	(*ptrgoglGetQueryObjecti64v)(id, pname, params);
// }
// void goglGetQueryObjectui64v(GLuint id, GLenum pname, GLuint64* params) {
// 	(*ptrgoglGetQueryObjectui64v)(id, pname, params);
// }
// void goglVertexP2ui(GLenum type, GLuint value) {
// 	(*ptrgoglVertexP2ui)(type, value);
// }
// void goglVertexP2uiv(GLenum type, GLuint* value) {
// 	(*ptrgoglVertexP2uiv)(type, value);
// }
// void goglVertexP3ui(GLenum type, GLuint value) {
// 	(*ptrgoglVertexP3ui)(type, value);
// }
// void goglVertexP3uiv(GLenum type, GLuint* value) {
// 	(*ptrgoglVertexP3uiv)(type, value);
// }
// void goglVertexP4ui(GLenum type, GLuint value) {
// 	(*ptrgoglVertexP4ui)(type, value);
// }
// void goglVertexP4uiv(GLenum type, GLuint* value) {
// 	(*ptrgoglVertexP4uiv)(type, value);
// }
// void goglTexCoordP1ui(GLenum type, GLuint coords) {
// 	(*ptrgoglTexCoordP1ui)(type, coords);
// }
// void goglTexCoordP1uiv(GLenum type, GLuint* coords) {
// 	(*ptrgoglTexCoordP1uiv)(type, coords);
// }
// void goglTexCoordP2ui(GLenum type, GLuint coords) {
// 	(*ptrgoglTexCoordP2ui)(type, coords);
// }
// void goglTexCoordP2uiv(GLenum type, GLuint* coords) {
// 	(*ptrgoglTexCoordP2uiv)(type, coords);
// }
// void goglTexCoordP3ui(GLenum type, GLuint coords) {
// 	(*ptrgoglTexCoordP3ui)(type, coords);
// }
// void goglTexCoordP3uiv(GLenum type, GLuint* coords) {
// 	(*ptrgoglTexCoordP3uiv)(type, coords);
// }
// void goglTexCoordP4ui(GLenum type, GLuint coords) {
// 	(*ptrgoglTexCoordP4ui)(type, coords);
// }
// void goglTexCoordP4uiv(GLenum type, GLuint* coords) {
// 	(*ptrgoglTexCoordP4uiv)(type, coords);
// }
// void goglMultiTexCoordP1ui(GLenum texture, GLenum type, GLuint coords) {
// 	(*ptrgoglMultiTexCoordP1ui)(texture, type, coords);
// }
// void goglMultiTexCoordP1uiv(GLenum texture, GLenum type, GLuint* coords) {
// 	(*ptrgoglMultiTexCoordP1uiv)(texture, type, coords);
// }
// void goglMultiTexCoordP2ui(GLenum texture, GLenum type, GLuint coords) {
// 	(*ptrgoglMultiTexCoordP2ui)(texture, type, coords);
// }
// void goglMultiTexCoordP2uiv(GLenum texture, GLenum type, GLuint* coords) {
// 	(*ptrgoglMultiTexCoordP2uiv)(texture, type, coords);
// }
// void goglMultiTexCoordP3ui(GLenum texture, GLenum type, GLuint coords) {
// 	(*ptrgoglMultiTexCoordP3ui)(texture, type, coords);
// }
// void goglMultiTexCoordP3uiv(GLenum texture, GLenum type, GLuint* coords) {
// 	(*ptrgoglMultiTexCoordP3uiv)(texture, type, coords);
// }
// void goglMultiTexCoordP4ui(GLenum texture, GLenum type, GLuint coords) {
// 	(*ptrgoglMultiTexCoordP4ui)(texture, type, coords);
// }
// void goglMultiTexCoordP4uiv(GLenum texture, GLenum type, GLuint* coords) {
// 	(*ptrgoglMultiTexCoordP4uiv)(texture, type, coords);
// }
// void goglNormalP3ui(GLenum type, GLuint coords) {
// 	(*ptrgoglNormalP3ui)(type, coords);
// }
// void goglNormalP3uiv(GLenum type, GLuint* coords) {
// 	(*ptrgoglNormalP3uiv)(type, coords);
// }
// void goglColorP3ui(GLenum type, GLuint color) {
// 	(*ptrgoglColorP3ui)(type, color);
// }
// void goglColorP3uiv(GLenum type, GLuint* color) {
// 	(*ptrgoglColorP3uiv)(type, color);
// }
// void goglColorP4ui(GLenum type, GLuint color) {
// 	(*ptrgoglColorP4ui)(type, color);
// }
// void goglColorP4uiv(GLenum type, GLuint* color) {
// 	(*ptrgoglColorP4uiv)(type, color);
// }
// void goglSecondaryColorP3ui(GLenum type, GLuint color) {
// 	(*ptrgoglSecondaryColorP3ui)(type, color);
// }
// void goglSecondaryColorP3uiv(GLenum type, GLuint* color) {
// 	(*ptrgoglSecondaryColorP3uiv)(type, color);
// }
// void goglVertexAttribP1ui(GLuint index, GLenum type, GLboolean normalized, GLuint value) {
// 	(*ptrgoglVertexAttribP1ui)(index, type, normalized, value);
// }
// void goglVertexAttribP1uiv(GLuint index, GLenum type, GLboolean normalized, GLuint* value) {
// 	(*ptrgoglVertexAttribP1uiv)(index, type, normalized, value);
// }
// void goglVertexAttribP2ui(GLuint index, GLenum type, GLboolean normalized, GLuint value) {
// 	(*ptrgoglVertexAttribP2ui)(index, type, normalized, value);
// }
// void goglVertexAttribP2uiv(GLuint index, GLenum type, GLboolean normalized, GLuint* value) {
// 	(*ptrgoglVertexAttribP2uiv)(index, type, normalized, value);
// }
// void goglVertexAttribP3ui(GLuint index, GLenum type, GLboolean normalized, GLuint value) {
// 	(*ptrgoglVertexAttribP3ui)(index, type, normalized, value);
// }
// void goglVertexAttribP3uiv(GLuint index, GLenum type, GLboolean normalized, GLuint* value) {
// 	(*ptrgoglVertexAttribP3uiv)(index, type, normalized, value);
// }
// void goglVertexAttribP4ui(GLuint index, GLenum type, GLboolean normalized, GLuint value) {
// 	(*ptrgoglVertexAttribP4ui)(index, type, normalized, value);
// }
// void goglVertexAttribP4uiv(GLuint index, GLenum type, GLboolean normalized, GLuint* value) {
// 	(*ptrgoglVertexAttribP4uiv)(index, type, normalized, value);
// }
// //  VERSION_3_0
// void goglColorMaski(GLuint index, GLboolean r, GLboolean g, GLboolean b, GLboolean a) {
// 	(*ptrgoglColorMaski)(index, r, g, b, a);
// }
// void goglGetBooleani_v(GLenum target, GLuint index, GLboolean* data) {
// 	(*ptrgoglGetBooleani_v)(target, index, data);
// }
// void goglGetIntegeri_v(GLenum target, GLuint index, GLint* data) {
// 	(*ptrgoglGetIntegeri_v)(target, index, data);
// }
// void goglEnablei(GLenum target, GLuint index) {
// 	(*ptrgoglEnablei)(target, index);
// }
// void goglDisablei(GLenum target, GLuint index) {
// 	(*ptrgoglDisablei)(target, index);
// }
// GLboolean goglIsEnabledi(GLenum target, GLuint index) {
// 	return (*ptrgoglIsEnabledi)(target, index);
// }
// void goglBeginTransformFeedback(GLenum primitiveMode) {
// 	(*ptrgoglBeginTransformFeedback)(primitiveMode);
// }
// void goglEndTransformFeedback() {
// 	(*ptrgoglEndTransformFeedback)();
// }
// void goglBindBufferRange(GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size) {
// 	(*ptrgoglBindBufferRange)(target, index, buffer, offset, size);
// }
// void goglBindBufferBase(GLenum target, GLuint index, GLuint buffer) {
// 	(*ptrgoglBindBufferBase)(target, index, buffer);
// }
// void goglTransformFeedbackVaryings(GLuint program, GLsizei count, GLchar** varyings, GLenum bufferMode) {
// 	(*ptrgoglTransformFeedbackVaryings)(program, count, varyings, bufferMode);
// }
// void goglGetTransformFeedbackVarying(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLsizei* size, GLenum* type, GLchar* name) {
// 	(*ptrgoglGetTransformFeedbackVarying)(program, index, bufSize, length, size, type, name);
// }
// void goglClampColor(GLenum target, GLenum clamp) {
// 	(*ptrgoglClampColor)(target, clamp);
// }
// void goglBeginConditionalRender(GLuint id, GLenum mode) {
// 	(*ptrgoglBeginConditionalRender)(id, mode);
// }
// void goglEndConditionalRender() {
// 	(*ptrgoglEndConditionalRender)();
// }
// void goglVertexAttribIPointer(GLuint index, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglVertexAttribIPointer)(index, size, type, stride, pointer);
// }
// void goglGetVertexAttribIiv(GLuint index, GLenum pname, GLint* params) {
// 	(*ptrgoglGetVertexAttribIiv)(index, pname, params);
// }
// void goglGetVertexAttribIuiv(GLuint index, GLenum pname, GLuint* params) {
// 	(*ptrgoglGetVertexAttribIuiv)(index, pname, params);
// }
// void goglVertexAttribI1i(GLuint index, GLint x) {
// 	(*ptrgoglVertexAttribI1i)(index, x);
// }
// void goglVertexAttribI2i(GLuint index, GLint x, GLint y) {
// 	(*ptrgoglVertexAttribI2i)(index, x, y);
// }
// void goglVertexAttribI3i(GLuint index, GLint x, GLint y, GLint z) {
// 	(*ptrgoglVertexAttribI3i)(index, x, y, z);
// }
// void goglVertexAttribI4i(GLuint index, GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrgoglVertexAttribI4i)(index, x, y, z, w);
// }
// void goglVertexAttribI1ui(GLuint index, GLuint x) {
// 	(*ptrgoglVertexAttribI1ui)(index, x);
// }
// void goglVertexAttribI2ui(GLuint index, GLuint x, GLuint y) {
// 	(*ptrgoglVertexAttribI2ui)(index, x, y);
// }
// void goglVertexAttribI3ui(GLuint index, GLuint x, GLuint y, GLuint z) {
// 	(*ptrgoglVertexAttribI3ui)(index, x, y, z);
// }
// void goglVertexAttribI4ui(GLuint index, GLuint x, GLuint y, GLuint z, GLuint w) {
// 	(*ptrgoglVertexAttribI4ui)(index, x, y, z, w);
// }
// void goglVertexAttribI1iv(GLuint index, GLint* v) {
// 	(*ptrgoglVertexAttribI1iv)(index, v);
// }
// void goglVertexAttribI2iv(GLuint index, GLint* v) {
// 	(*ptrgoglVertexAttribI2iv)(index, v);
// }
// void goglVertexAttribI3iv(GLuint index, GLint* v) {
// 	(*ptrgoglVertexAttribI3iv)(index, v);
// }
// void goglVertexAttribI4iv(GLuint index, GLint* v) {
// 	(*ptrgoglVertexAttribI4iv)(index, v);
// }
// void goglVertexAttribI1uiv(GLuint index, GLuint* v) {
// 	(*ptrgoglVertexAttribI1uiv)(index, v);
// }
// void goglVertexAttribI2uiv(GLuint index, GLuint* v) {
// 	(*ptrgoglVertexAttribI2uiv)(index, v);
// }
// void goglVertexAttribI3uiv(GLuint index, GLuint* v) {
// 	(*ptrgoglVertexAttribI3uiv)(index, v);
// }
// void goglVertexAttribI4uiv(GLuint index, GLuint* v) {
// 	(*ptrgoglVertexAttribI4uiv)(index, v);
// }
// void goglVertexAttribI4bv(GLuint index, GLbyte* v) {
// 	(*ptrgoglVertexAttribI4bv)(index, v);
// }
// void goglVertexAttribI4sv(GLuint index, GLshort* v) {
// 	(*ptrgoglVertexAttribI4sv)(index, v);
// }
// void goglVertexAttribI4ubv(GLuint index, GLubyte* v) {
// 	(*ptrgoglVertexAttribI4ubv)(index, v);
// }
// void goglVertexAttribI4usv(GLuint index, GLushort* v) {
// 	(*ptrgoglVertexAttribI4usv)(index, v);
// }
// void goglGetUniformuiv(GLuint program, GLint location, GLuint* params) {
// 	(*ptrgoglGetUniformuiv)(program, location, params);
// }
// void goglBindFragDataLocation(GLuint program, GLuint color, GLchar* name) {
// 	(*ptrgoglBindFragDataLocation)(program, color, name);
// }
// GLint goglGetFragDataLocation(GLuint program, GLchar* name) {
// 	return (*ptrgoglGetFragDataLocation)(program, name);
// }
// void goglUniform1ui(GLint location, GLuint v0) {
// 	(*ptrgoglUniform1ui)(location, v0);
// }
// void goglUniform2ui(GLint location, GLuint v0, GLuint v1) {
// 	(*ptrgoglUniform2ui)(location, v0, v1);
// }
// void goglUniform3ui(GLint location, GLuint v0, GLuint v1, GLuint v2) {
// 	(*ptrgoglUniform3ui)(location, v0, v1, v2);
// }
// void goglUniform4ui(GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3) {
// 	(*ptrgoglUniform4ui)(location, v0, v1, v2, v3);
// }
// void goglUniform1uiv(GLint location, GLsizei count, GLuint* value) {
// 	(*ptrgoglUniform1uiv)(location, count, value);
// }
// void goglUniform2uiv(GLint location, GLsizei count, GLuint* value) {
// 	(*ptrgoglUniform2uiv)(location, count, value);
// }
// void goglUniform3uiv(GLint location, GLsizei count, GLuint* value) {
// 	(*ptrgoglUniform3uiv)(location, count, value);
// }
// void goglUniform4uiv(GLint location, GLsizei count, GLuint* value) {
// 	(*ptrgoglUniform4uiv)(location, count, value);
// }
// void goglTexParameterIiv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglTexParameterIiv)(target, pname, params);
// }
// void goglTexParameterIuiv(GLenum target, GLenum pname, GLuint* params) {
// 	(*ptrgoglTexParameterIuiv)(target, pname, params);
// }
// void goglGetTexParameterIiv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglGetTexParameterIiv)(target, pname, params);
// }
// void goglGetTexParameterIuiv(GLenum target, GLenum pname, GLuint* params) {
// 	(*ptrgoglGetTexParameterIuiv)(target, pname, params);
// }
// void goglClearBufferiv(GLenum buffer, GLint drawbuffer, GLint* value) {
// 	(*ptrgoglClearBufferiv)(buffer, drawbuffer, value);
// }
// void goglClearBufferuiv(GLenum buffer, GLint drawbuffer, GLuint* value) {
// 	(*ptrgoglClearBufferuiv)(buffer, drawbuffer, value);
// }
// void goglClearBufferfv(GLenum buffer, GLint drawbuffer, GLfloat* value) {
// 	(*ptrgoglClearBufferfv)(buffer, drawbuffer, value);
// }
// void goglClearBufferfi(GLenum buffer, GLint drawbuffer, GLfloat depth, GLint stencil) {
// 	(*ptrgoglClearBufferfi)(buffer, drawbuffer, depth, stencil);
// }
// const GLubyte * goglGetStringi(GLenum name, GLuint index) {
// 	return (*ptrgoglGetStringi)(name, index);
// }
// GLboolean goglIsRenderbuffer(GLuint renderbuffer) {
// 	return (*ptrgoglIsRenderbuffer)(renderbuffer);
// }
// void goglBindRenderbuffer(GLenum target, GLuint renderbuffer) {
// 	(*ptrgoglBindRenderbuffer)(target, renderbuffer);
// }
// void goglDeleteRenderbuffers(GLsizei n, GLuint* renderbuffers) {
// 	(*ptrgoglDeleteRenderbuffers)(n, renderbuffers);
// }
// void goglGenRenderbuffers(GLsizei n, GLuint* renderbuffers) {
// 	(*ptrgoglGenRenderbuffers)(n, renderbuffers);
// }
// void goglRenderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height) {
// 	(*ptrgoglRenderbufferStorage)(target, internalformat, width, height);
// }
// void goglGetRenderbufferParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglGetRenderbufferParameteriv)(target, pname, params);
// }
// GLboolean goglIsFramebuffer(GLuint framebuffer) {
// 	return (*ptrgoglIsFramebuffer)(framebuffer);
// }
// void goglBindFramebuffer(GLenum target, GLuint framebuffer) {
// 	(*ptrgoglBindFramebuffer)(target, framebuffer);
// }
// void goglDeleteFramebuffers(GLsizei n, GLuint* framebuffers) {
// 	(*ptrgoglDeleteFramebuffers)(n, framebuffers);
// }
// void goglGenFramebuffers(GLsizei n, GLuint* framebuffers) {
// 	(*ptrgoglGenFramebuffers)(n, framebuffers);
// }
// GLenum goglCheckFramebufferStatus(GLenum target) {
// 	return (*ptrgoglCheckFramebufferStatus)(target);
// }
// void goglFramebufferTexture1D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level) {
// 	(*ptrgoglFramebufferTexture1D)(target, attachment, textarget, texture, level);
// }
// void goglFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level) {
// 	(*ptrgoglFramebufferTexture2D)(target, attachment, textarget, texture, level);
// }
// void goglFramebufferTexture3D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint zoffset) {
// 	(*ptrgoglFramebufferTexture3D)(target, attachment, textarget, texture, level, zoffset);
// }
// void goglFramebufferRenderbuffer(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer) {
// 	(*ptrgoglFramebufferRenderbuffer)(target, attachment, renderbuffertarget, renderbuffer);
// }
// void goglGetFramebufferAttachmentParameteriv(GLenum target, GLenum attachment, GLenum pname, GLint* params) {
// 	(*ptrgoglGetFramebufferAttachmentParameteriv)(target, attachment, pname, params);
// }
// void goglGenerateMipmap(GLenum target) {
// 	(*ptrgoglGenerateMipmap)(target);
// }
// void goglBlitFramebuffer(GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter) {
// 	(*ptrgoglBlitFramebuffer)(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// void goglRenderbufferStorageMultisample(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height) {
// 	(*ptrgoglRenderbufferStorageMultisample)(target, samples, internalformat, width, height);
// }
// void goglFramebufferTextureLayer(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer) {
// 	(*ptrgoglFramebufferTextureLayer)(target, attachment, texture, level, layer);
// }
// GLvoid* goglMapBufferRange(GLenum target, GLintptr offset, GLsizeiptr length, GLbitfield access) {
// 	return (*ptrgoglMapBufferRange)(target, offset, length, access);
// }
// void goglFlushMappedBufferRange(GLenum target, GLintptr offset, GLsizeiptr length) {
// 	(*ptrgoglFlushMappedBufferRange)(target, offset, length);
// }
// void goglBindVertexArray(GLuint array) {
// 	(*ptrgoglBindVertexArray)(array);
// }
// void goglDeleteVertexArrays(GLsizei n, GLuint* arrays) {
// 	(*ptrgoglDeleteVertexArrays)(n, arrays);
// }
// void goglGenVertexArrays(GLsizei n, GLuint* arrays) {
// 	(*ptrgoglGenVertexArrays)(n, arrays);
// }
// GLboolean goglIsVertexArray(GLuint array) {
// 	return (*ptrgoglIsVertexArray)(array);
// }
// //  VERSION_3_1
// void goglDrawArraysInstanced(GLenum mode, GLint first, GLsizei count, GLsizei primcount) {
// 	(*ptrgoglDrawArraysInstanced)(mode, first, count, primcount);
// }
// void goglDrawElementsInstanced(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei primcount) {
// 	(*ptrgoglDrawElementsInstanced)(mode, count, type, indices, primcount);
// }
// void goglTexBuffer(GLenum target, GLenum internalformat, GLuint buffer) {
// 	(*ptrgoglTexBuffer)(target, internalformat, buffer);
// }
// void goglPrimitiveRestartIndex(GLuint index) {
// 	(*ptrgoglPrimitiveRestartIndex)(index);
// }
// void goglCopyBufferSubData(GLenum readTarget, GLenum writeTarget, GLintptr readOffset, GLintptr writeOffset, GLsizeiptr size) {
// 	(*ptrgoglCopyBufferSubData)(readTarget, writeTarget, readOffset, writeOffset, size);
// }
// void goglGetUniformIndices(GLuint program, GLsizei uniformCount, GLchar** uniformNames, GLuint* uniformIndices) {
// 	(*ptrgoglGetUniformIndices)(program, uniformCount, uniformNames, uniformIndices);
// }
// void goglGetActiveUniformsiv(GLuint program, GLsizei uniformCount, GLuint* uniformIndices, GLenum pname, GLint* params) {
// 	(*ptrgoglGetActiveUniformsiv)(program, uniformCount, uniformIndices, pname, params);
// }
// void goglGetActiveUniformName(GLuint program, GLuint uniformIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformName) {
// 	(*ptrgoglGetActiveUniformName)(program, uniformIndex, bufSize, length, uniformName);
// }
// GLuint goglGetUniformBlockIndex(GLuint program, GLchar* uniformBlockName) {
// 	return (*ptrgoglGetUniformBlockIndex)(program, uniformBlockName);
// }
// void goglGetActiveUniformBlockiv(GLuint program, GLuint uniformBlockIndex, GLenum pname, GLint* params) {
// 	(*ptrgoglGetActiveUniformBlockiv)(program, uniformBlockIndex, pname, params);
// }
// void goglGetActiveUniformBlockName(GLuint program, GLuint uniformBlockIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformBlockName) {
// 	(*ptrgoglGetActiveUniformBlockName)(program, uniformBlockIndex, bufSize, length, uniformBlockName);
// }
// void goglUniformBlockBinding(GLuint program, GLuint uniformBlockIndex, GLuint uniformBlockBinding) {
// 	(*ptrgoglUniformBlockBinding)(program, uniformBlockIndex, uniformBlockBinding);
// }
// 
// int init_VERSION_2_1() {
// 	ptrgoglUniformMatrix2x3fv = goglGetProcAddress("glUniformMatrix2x3fv");
// 	if(ptrgoglUniformMatrix2x3fv == NULL) return 1;
// 	ptrgoglUniformMatrix3x2fv = goglGetProcAddress("glUniformMatrix3x2fv");
// 	if(ptrgoglUniformMatrix3x2fv == NULL) return 1;
// 	ptrgoglUniformMatrix2x4fv = goglGetProcAddress("glUniformMatrix2x4fv");
// 	if(ptrgoglUniformMatrix2x4fv == NULL) return 1;
// 	ptrgoglUniformMatrix4x2fv = goglGetProcAddress("glUniformMatrix4x2fv");
// 	if(ptrgoglUniformMatrix4x2fv == NULL) return 1;
// 	ptrgoglUniformMatrix3x4fv = goglGetProcAddress("glUniformMatrix3x4fv");
// 	if(ptrgoglUniformMatrix3x4fv == NULL) return 1;
// 	ptrgoglUniformMatrix4x3fv = goglGetProcAddress("glUniformMatrix4x3fv");
// 	if(ptrgoglUniformMatrix4x3fv == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_2_0() {
// 	ptrgoglBlendEquationSeparate = goglGetProcAddress("glBlendEquationSeparate");
// 	if(ptrgoglBlendEquationSeparate == NULL) return 1;
// 	ptrgoglDrawBuffers = goglGetProcAddress("glDrawBuffers");
// 	if(ptrgoglDrawBuffers == NULL) return 1;
// 	ptrgoglStencilOpSeparate = goglGetProcAddress("glStencilOpSeparate");
// 	if(ptrgoglStencilOpSeparate == NULL) return 1;
// 	ptrgoglStencilFuncSeparate = goglGetProcAddress("glStencilFuncSeparate");
// 	if(ptrgoglStencilFuncSeparate == NULL) return 1;
// 	ptrgoglStencilMaskSeparate = goglGetProcAddress("glStencilMaskSeparate");
// 	if(ptrgoglStencilMaskSeparate == NULL) return 1;
// 	ptrgoglAttachShader = goglGetProcAddress("glAttachShader");
// 	if(ptrgoglAttachShader == NULL) return 1;
// 	ptrgoglBindAttribLocation = goglGetProcAddress("glBindAttribLocation");
// 	if(ptrgoglBindAttribLocation == NULL) return 1;
// 	ptrgoglCompileShader = goglGetProcAddress("glCompileShader");
// 	if(ptrgoglCompileShader == NULL) return 1;
// 	ptrgoglCreateProgram = goglGetProcAddress("glCreateProgram");
// 	if(ptrgoglCreateProgram == NULL) return 1;
// 	ptrgoglCreateShader = goglGetProcAddress("glCreateShader");
// 	if(ptrgoglCreateShader == NULL) return 1;
// 	ptrgoglDeleteProgram = goglGetProcAddress("glDeleteProgram");
// 	if(ptrgoglDeleteProgram == NULL) return 1;
// 	ptrgoglDeleteShader = goglGetProcAddress("glDeleteShader");
// 	if(ptrgoglDeleteShader == NULL) return 1;
// 	ptrgoglDetachShader = goglGetProcAddress("glDetachShader");
// 	if(ptrgoglDetachShader == NULL) return 1;
// 	ptrgoglDisableVertexAttribArray = goglGetProcAddress("glDisableVertexAttribArray");
// 	if(ptrgoglDisableVertexAttribArray == NULL) return 1;
// 	ptrgoglEnableVertexAttribArray = goglGetProcAddress("glEnableVertexAttribArray");
// 	if(ptrgoglEnableVertexAttribArray == NULL) return 1;
// 	ptrgoglGetActiveAttrib = goglGetProcAddress("glGetActiveAttrib");
// 	if(ptrgoglGetActiveAttrib == NULL) return 1;
// 	ptrgoglGetActiveUniform = goglGetProcAddress("glGetActiveUniform");
// 	if(ptrgoglGetActiveUniform == NULL) return 1;
// 	ptrgoglGetAttachedShaders = goglGetProcAddress("glGetAttachedShaders");
// 	if(ptrgoglGetAttachedShaders == NULL) return 1;
// 	ptrgoglGetAttribLocation = goglGetProcAddress("glGetAttribLocation");
// 	if(ptrgoglGetAttribLocation == NULL) return 1;
// 	ptrgoglGetProgramiv = goglGetProcAddress("glGetProgramiv");
// 	if(ptrgoglGetProgramiv == NULL) return 1;
// 	ptrgoglGetProgramInfoLog = goglGetProcAddress("glGetProgramInfoLog");
// 	if(ptrgoglGetProgramInfoLog == NULL) return 1;
// 	ptrgoglGetShaderiv = goglGetProcAddress("glGetShaderiv");
// 	if(ptrgoglGetShaderiv == NULL) return 1;
// 	ptrgoglGetShaderInfoLog = goglGetProcAddress("glGetShaderInfoLog");
// 	if(ptrgoglGetShaderInfoLog == NULL) return 1;
// 	ptrgoglGetShaderSource = goglGetProcAddress("glGetShaderSource");
// 	if(ptrgoglGetShaderSource == NULL) return 1;
// 	ptrgoglGetUniformLocation = goglGetProcAddress("glGetUniformLocation");
// 	if(ptrgoglGetUniformLocation == NULL) return 1;
// 	ptrgoglGetUniformfv = goglGetProcAddress("glGetUniformfv");
// 	if(ptrgoglGetUniformfv == NULL) return 1;
// 	ptrgoglGetUniformiv = goglGetProcAddress("glGetUniformiv");
// 	if(ptrgoglGetUniformiv == NULL) return 1;
// 	ptrgoglGetVertexAttribdv = goglGetProcAddress("glGetVertexAttribdv");
// 	if(ptrgoglGetVertexAttribdv == NULL) return 1;
// 	ptrgoglGetVertexAttribfv = goglGetProcAddress("glGetVertexAttribfv");
// 	if(ptrgoglGetVertexAttribfv == NULL) return 1;
// 	ptrgoglGetVertexAttribiv = goglGetProcAddress("glGetVertexAttribiv");
// 	if(ptrgoglGetVertexAttribiv == NULL) return 1;
// 	ptrgoglGetVertexAttribPointerv = goglGetProcAddress("glGetVertexAttribPointerv");
// 	if(ptrgoglGetVertexAttribPointerv == NULL) return 1;
// 	ptrgoglIsProgram = goglGetProcAddress("glIsProgram");
// 	if(ptrgoglIsProgram == NULL) return 1;
// 	ptrgoglIsShader = goglGetProcAddress("glIsShader");
// 	if(ptrgoglIsShader == NULL) return 1;
// 	ptrgoglLinkProgram = goglGetProcAddress("glLinkProgram");
// 	if(ptrgoglLinkProgram == NULL) return 1;
// 	ptrgoglShaderSource = goglGetProcAddress("glShaderSource");
// 	if(ptrgoglShaderSource == NULL) return 1;
// 	ptrgoglUseProgram = goglGetProcAddress("glUseProgram");
// 	if(ptrgoglUseProgram == NULL) return 1;
// 	ptrgoglUniform1f = goglGetProcAddress("glUniform1f");
// 	if(ptrgoglUniform1f == NULL) return 1;
// 	ptrgoglUniform2f = goglGetProcAddress("glUniform2f");
// 	if(ptrgoglUniform2f == NULL) return 1;
// 	ptrgoglUniform3f = goglGetProcAddress("glUniform3f");
// 	if(ptrgoglUniform3f == NULL) return 1;
// 	ptrgoglUniform4f = goglGetProcAddress("glUniform4f");
// 	if(ptrgoglUniform4f == NULL) return 1;
// 	ptrgoglUniform1i = goglGetProcAddress("glUniform1i");
// 	if(ptrgoglUniform1i == NULL) return 1;
// 	ptrgoglUniform2i = goglGetProcAddress("glUniform2i");
// 	if(ptrgoglUniform2i == NULL) return 1;
// 	ptrgoglUniform3i = goglGetProcAddress("glUniform3i");
// 	if(ptrgoglUniform3i == NULL) return 1;
// 	ptrgoglUniform4i = goglGetProcAddress("glUniform4i");
// 	if(ptrgoglUniform4i == NULL) return 1;
// 	ptrgoglUniform1fv = goglGetProcAddress("glUniform1fv");
// 	if(ptrgoglUniform1fv == NULL) return 1;
// 	ptrgoglUniform2fv = goglGetProcAddress("glUniform2fv");
// 	if(ptrgoglUniform2fv == NULL) return 1;
// 	ptrgoglUniform3fv = goglGetProcAddress("glUniform3fv");
// 	if(ptrgoglUniform3fv == NULL) return 1;
// 	ptrgoglUniform4fv = goglGetProcAddress("glUniform4fv");
// 	if(ptrgoglUniform4fv == NULL) return 1;
// 	ptrgoglUniform1iv = goglGetProcAddress("glUniform1iv");
// 	if(ptrgoglUniform1iv == NULL) return 1;
// 	ptrgoglUniform2iv = goglGetProcAddress("glUniform2iv");
// 	if(ptrgoglUniform2iv == NULL) return 1;
// 	ptrgoglUniform3iv = goglGetProcAddress("glUniform3iv");
// 	if(ptrgoglUniform3iv == NULL) return 1;
// 	ptrgoglUniform4iv = goglGetProcAddress("glUniform4iv");
// 	if(ptrgoglUniform4iv == NULL) return 1;
// 	ptrgoglUniformMatrix2fv = goglGetProcAddress("glUniformMatrix2fv");
// 	if(ptrgoglUniformMatrix2fv == NULL) return 1;
// 	ptrgoglUniformMatrix3fv = goglGetProcAddress("glUniformMatrix3fv");
// 	if(ptrgoglUniformMatrix3fv == NULL) return 1;
// 	ptrgoglUniformMatrix4fv = goglGetProcAddress("glUniformMatrix4fv");
// 	if(ptrgoglUniformMatrix4fv == NULL) return 1;
// 	ptrgoglValidateProgram = goglGetProcAddress("glValidateProgram");
// 	if(ptrgoglValidateProgram == NULL) return 1;
// 	ptrgoglVertexAttrib1d = goglGetProcAddress("glVertexAttrib1d");
// 	if(ptrgoglVertexAttrib1d == NULL) return 1;
// 	ptrgoglVertexAttrib1dv = goglGetProcAddress("glVertexAttrib1dv");
// 	if(ptrgoglVertexAttrib1dv == NULL) return 1;
// 	ptrgoglVertexAttrib1f = goglGetProcAddress("glVertexAttrib1f");
// 	if(ptrgoglVertexAttrib1f == NULL) return 1;
// 	ptrgoglVertexAttrib1fv = goglGetProcAddress("glVertexAttrib1fv");
// 	if(ptrgoglVertexAttrib1fv == NULL) return 1;
// 	ptrgoglVertexAttrib1s = goglGetProcAddress("glVertexAttrib1s");
// 	if(ptrgoglVertexAttrib1s == NULL) return 1;
// 	ptrgoglVertexAttrib1sv = goglGetProcAddress("glVertexAttrib1sv");
// 	if(ptrgoglVertexAttrib1sv == NULL) return 1;
// 	ptrgoglVertexAttrib2d = goglGetProcAddress("glVertexAttrib2d");
// 	if(ptrgoglVertexAttrib2d == NULL) return 1;
// 	ptrgoglVertexAttrib2dv = goglGetProcAddress("glVertexAttrib2dv");
// 	if(ptrgoglVertexAttrib2dv == NULL) return 1;
// 	ptrgoglVertexAttrib2f = goglGetProcAddress("glVertexAttrib2f");
// 	if(ptrgoglVertexAttrib2f == NULL) return 1;
// 	ptrgoglVertexAttrib2fv = goglGetProcAddress("glVertexAttrib2fv");
// 	if(ptrgoglVertexAttrib2fv == NULL) return 1;
// 	ptrgoglVertexAttrib2s = goglGetProcAddress("glVertexAttrib2s");
// 	if(ptrgoglVertexAttrib2s == NULL) return 1;
// 	ptrgoglVertexAttrib2sv = goglGetProcAddress("glVertexAttrib2sv");
// 	if(ptrgoglVertexAttrib2sv == NULL) return 1;
// 	ptrgoglVertexAttrib3d = goglGetProcAddress("glVertexAttrib3d");
// 	if(ptrgoglVertexAttrib3d == NULL) return 1;
// 	ptrgoglVertexAttrib3dv = goglGetProcAddress("glVertexAttrib3dv");
// 	if(ptrgoglVertexAttrib3dv == NULL) return 1;
// 	ptrgoglVertexAttrib3f = goglGetProcAddress("glVertexAttrib3f");
// 	if(ptrgoglVertexAttrib3f == NULL) return 1;
// 	ptrgoglVertexAttrib3fv = goglGetProcAddress("glVertexAttrib3fv");
// 	if(ptrgoglVertexAttrib3fv == NULL) return 1;
// 	ptrgoglVertexAttrib3s = goglGetProcAddress("glVertexAttrib3s");
// 	if(ptrgoglVertexAttrib3s == NULL) return 1;
// 	ptrgoglVertexAttrib3sv = goglGetProcAddress("glVertexAttrib3sv");
// 	if(ptrgoglVertexAttrib3sv == NULL) return 1;
// 	ptrgoglVertexAttrib4Nbv = goglGetProcAddress("glVertexAttrib4Nbv");
// 	if(ptrgoglVertexAttrib4Nbv == NULL) return 1;
// 	ptrgoglVertexAttrib4Niv = goglGetProcAddress("glVertexAttrib4Niv");
// 	if(ptrgoglVertexAttrib4Niv == NULL) return 1;
// 	ptrgoglVertexAttrib4Nsv = goglGetProcAddress("glVertexAttrib4Nsv");
// 	if(ptrgoglVertexAttrib4Nsv == NULL) return 1;
// 	ptrgoglVertexAttrib4Nub = goglGetProcAddress("glVertexAttrib4Nub");
// 	if(ptrgoglVertexAttrib4Nub == NULL) return 1;
// 	ptrgoglVertexAttrib4Nubv = goglGetProcAddress("glVertexAttrib4Nubv");
// 	if(ptrgoglVertexAttrib4Nubv == NULL) return 1;
// 	ptrgoglVertexAttrib4Nuiv = goglGetProcAddress("glVertexAttrib4Nuiv");
// 	if(ptrgoglVertexAttrib4Nuiv == NULL) return 1;
// 	ptrgoglVertexAttrib4Nusv = goglGetProcAddress("glVertexAttrib4Nusv");
// 	if(ptrgoglVertexAttrib4Nusv == NULL) return 1;
// 	ptrgoglVertexAttrib4bv = goglGetProcAddress("glVertexAttrib4bv");
// 	if(ptrgoglVertexAttrib4bv == NULL) return 1;
// 	ptrgoglVertexAttrib4d = goglGetProcAddress("glVertexAttrib4d");
// 	if(ptrgoglVertexAttrib4d == NULL) return 1;
// 	ptrgoglVertexAttrib4dv = goglGetProcAddress("glVertexAttrib4dv");
// 	if(ptrgoglVertexAttrib4dv == NULL) return 1;
// 	ptrgoglVertexAttrib4f = goglGetProcAddress("glVertexAttrib4f");
// 	if(ptrgoglVertexAttrib4f == NULL) return 1;
// 	ptrgoglVertexAttrib4fv = goglGetProcAddress("glVertexAttrib4fv");
// 	if(ptrgoglVertexAttrib4fv == NULL) return 1;
// 	ptrgoglVertexAttrib4iv = goglGetProcAddress("glVertexAttrib4iv");
// 	if(ptrgoglVertexAttrib4iv == NULL) return 1;
// 	ptrgoglVertexAttrib4s = goglGetProcAddress("glVertexAttrib4s");
// 	if(ptrgoglVertexAttrib4s == NULL) return 1;
// 	ptrgoglVertexAttrib4sv = goglGetProcAddress("glVertexAttrib4sv");
// 	if(ptrgoglVertexAttrib4sv == NULL) return 1;
// 	ptrgoglVertexAttrib4ubv = goglGetProcAddress("glVertexAttrib4ubv");
// 	if(ptrgoglVertexAttrib4ubv == NULL) return 1;
// 	ptrgoglVertexAttrib4uiv = goglGetProcAddress("glVertexAttrib4uiv");
// 	if(ptrgoglVertexAttrib4uiv == NULL) return 1;
// 	ptrgoglVertexAttrib4usv = goglGetProcAddress("glVertexAttrib4usv");
// 	if(ptrgoglVertexAttrib4usv == NULL) return 1;
// 	ptrgoglVertexAttribPointer = goglGetProcAddress("glVertexAttribPointer");
// 	if(ptrgoglVertexAttribPointer == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_4() {
// 	ptrgoglBlendFuncSeparate = goglGetProcAddress("glBlendFuncSeparate");
// 	if(ptrgoglBlendFuncSeparate == NULL) return 1;
// 	ptrgoglMultiDrawArrays = goglGetProcAddress("glMultiDrawArrays");
// 	if(ptrgoglMultiDrawArrays == NULL) return 1;
// 	ptrgoglMultiDrawElements = goglGetProcAddress("glMultiDrawElements");
// 	if(ptrgoglMultiDrawElements == NULL) return 1;
// 	ptrgoglPointParameterf = goglGetProcAddress("glPointParameterf");
// 	if(ptrgoglPointParameterf == NULL) return 1;
// 	ptrgoglPointParameterfv = goglGetProcAddress("glPointParameterfv");
// 	if(ptrgoglPointParameterfv == NULL) return 1;
// 	ptrgoglPointParameteri = goglGetProcAddress("glPointParameteri");
// 	if(ptrgoglPointParameteri == NULL) return 1;
// 	ptrgoglPointParameteriv = goglGetProcAddress("glPointParameteriv");
// 	if(ptrgoglPointParameteriv == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_5() {
// 	ptrgoglGenQueries = goglGetProcAddress("glGenQueries");
// 	if(ptrgoglGenQueries == NULL) return 1;
// 	ptrgoglDeleteQueries = goglGetProcAddress("glDeleteQueries");
// 	if(ptrgoglDeleteQueries == NULL) return 1;
// 	ptrgoglIsQuery = goglGetProcAddress("glIsQuery");
// 	if(ptrgoglIsQuery == NULL) return 1;
// 	ptrgoglBeginQuery = goglGetProcAddress("glBeginQuery");
// 	if(ptrgoglBeginQuery == NULL) return 1;
// 	ptrgoglEndQuery = goglGetProcAddress("glEndQuery");
// 	if(ptrgoglEndQuery == NULL) return 1;
// 	ptrgoglGetQueryiv = goglGetProcAddress("glGetQueryiv");
// 	if(ptrgoglGetQueryiv == NULL) return 1;
// 	ptrgoglGetQueryObjectiv = goglGetProcAddress("glGetQueryObjectiv");
// 	if(ptrgoglGetQueryObjectiv == NULL) return 1;
// 	ptrgoglGetQueryObjectuiv = goglGetProcAddress("glGetQueryObjectuiv");
// 	if(ptrgoglGetQueryObjectuiv == NULL) return 1;
// 	ptrgoglBindBuffer = goglGetProcAddress("glBindBuffer");
// 	if(ptrgoglBindBuffer == NULL) return 1;
// 	ptrgoglDeleteBuffers = goglGetProcAddress("glDeleteBuffers");
// 	if(ptrgoglDeleteBuffers == NULL) return 1;
// 	ptrgoglGenBuffers = goglGetProcAddress("glGenBuffers");
// 	if(ptrgoglGenBuffers == NULL) return 1;
// 	ptrgoglIsBuffer = goglGetProcAddress("glIsBuffer");
// 	if(ptrgoglIsBuffer == NULL) return 1;
// 	ptrgoglBufferData = goglGetProcAddress("glBufferData");
// 	if(ptrgoglBufferData == NULL) return 1;
// 	ptrgoglBufferSubData = goglGetProcAddress("glBufferSubData");
// 	if(ptrgoglBufferSubData == NULL) return 1;
// 	ptrgoglGetBufferSubData = goglGetProcAddress("glGetBufferSubData");
// 	if(ptrgoglGetBufferSubData == NULL) return 1;
// 	ptrgoglMapBuffer = goglGetProcAddress("glMapBuffer");
// 	if(ptrgoglMapBuffer == NULL) return 1;
// 	ptrgoglUnmapBuffer = goglGetProcAddress("glUnmapBuffer");
// 	if(ptrgoglUnmapBuffer == NULL) return 1;
// 	ptrgoglGetBufferParameteriv = goglGetProcAddress("glGetBufferParameteriv");
// 	if(ptrgoglGetBufferParameteriv == NULL) return 1;
// 	ptrgoglGetBufferPointerv = goglGetProcAddress("glGetBufferPointerv");
// 	if(ptrgoglGetBufferPointerv == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_0() {
// 	ptrgoglCullFace = goglGetProcAddress("glCullFace");
// 	if(ptrgoglCullFace == NULL) return 1;
// 	ptrgoglFrontFace = goglGetProcAddress("glFrontFace");
// 	if(ptrgoglFrontFace == NULL) return 1;
// 	ptrgoglHint = goglGetProcAddress("glHint");
// 	if(ptrgoglHint == NULL) return 1;
// 	ptrgoglLineWidth = goglGetProcAddress("glLineWidth");
// 	if(ptrgoglLineWidth == NULL) return 1;
// 	ptrgoglPointSize = goglGetProcAddress("glPointSize");
// 	if(ptrgoglPointSize == NULL) return 1;
// 	ptrgoglPolygonMode = goglGetProcAddress("glPolygonMode");
// 	if(ptrgoglPolygonMode == NULL) return 1;
// 	ptrgoglScissor = goglGetProcAddress("glScissor");
// 	if(ptrgoglScissor == NULL) return 1;
// 	ptrgoglTexParameterf = goglGetProcAddress("glTexParameterf");
// 	if(ptrgoglTexParameterf == NULL) return 1;
// 	ptrgoglTexParameterfv = goglGetProcAddress("glTexParameterfv");
// 	if(ptrgoglTexParameterfv == NULL) return 1;
// 	ptrgoglTexParameteri = goglGetProcAddress("glTexParameteri");
// 	if(ptrgoglTexParameteri == NULL) return 1;
// 	ptrgoglTexParameteriv = goglGetProcAddress("glTexParameteriv");
// 	if(ptrgoglTexParameteriv == NULL) return 1;
// 	ptrgoglTexImage1D = goglGetProcAddress("glTexImage1D");
// 	if(ptrgoglTexImage1D == NULL) return 1;
// 	ptrgoglTexImage2D = goglGetProcAddress("glTexImage2D");
// 	if(ptrgoglTexImage2D == NULL) return 1;
// 	ptrgoglDrawBuffer = goglGetProcAddress("glDrawBuffer");
// 	if(ptrgoglDrawBuffer == NULL) return 1;
// 	ptrgoglClear = goglGetProcAddress("glClear");
// 	if(ptrgoglClear == NULL) return 1;
// 	ptrgoglClearColor = goglGetProcAddress("glClearColor");
// 	if(ptrgoglClearColor == NULL) return 1;
// 	ptrgoglClearStencil = goglGetProcAddress("glClearStencil");
// 	if(ptrgoglClearStencil == NULL) return 1;
// 	ptrgoglClearDepth = goglGetProcAddress("glClearDepth");
// 	if(ptrgoglClearDepth == NULL) return 1;
// 	ptrgoglStencilMask = goglGetProcAddress("glStencilMask");
// 	if(ptrgoglStencilMask == NULL) return 1;
// 	ptrgoglColorMask = goglGetProcAddress("glColorMask");
// 	if(ptrgoglColorMask == NULL) return 1;
// 	ptrgoglDepthMask = goglGetProcAddress("glDepthMask");
// 	if(ptrgoglDepthMask == NULL) return 1;
// 	ptrgoglDisable = goglGetProcAddress("glDisable");
// 	if(ptrgoglDisable == NULL) return 1;
// 	ptrgoglEnable = goglGetProcAddress("glEnable");
// 	if(ptrgoglEnable == NULL) return 1;
// 	ptrgoglFinish = goglGetProcAddress("glFinish");
// 	if(ptrgoglFinish == NULL) return 1;
// 	ptrgoglFlush = goglGetProcAddress("glFlush");
// 	if(ptrgoglFlush == NULL) return 1;
// 	ptrgoglBlendFunc = goglGetProcAddress("glBlendFunc");
// 	if(ptrgoglBlendFunc == NULL) return 1;
// 	ptrgoglLogicOp = goglGetProcAddress("glLogicOp");
// 	if(ptrgoglLogicOp == NULL) return 1;
// 	ptrgoglStencilFunc = goglGetProcAddress("glStencilFunc");
// 	if(ptrgoglStencilFunc == NULL) return 1;
// 	ptrgoglStencilOp = goglGetProcAddress("glStencilOp");
// 	if(ptrgoglStencilOp == NULL) return 1;
// 	ptrgoglDepthFunc = goglGetProcAddress("glDepthFunc");
// 	if(ptrgoglDepthFunc == NULL) return 1;
// 	ptrgoglPixelStoref = goglGetProcAddress("glPixelStoref");
// 	if(ptrgoglPixelStoref == NULL) return 1;
// 	ptrgoglPixelStorei = goglGetProcAddress("glPixelStorei");
// 	if(ptrgoglPixelStorei == NULL) return 1;
// 	ptrgoglReadBuffer = goglGetProcAddress("glReadBuffer");
// 	if(ptrgoglReadBuffer == NULL) return 1;
// 	ptrgoglReadPixels = goglGetProcAddress("glReadPixels");
// 	if(ptrgoglReadPixels == NULL) return 1;
// 	ptrgoglGetBooleanv = goglGetProcAddress("glGetBooleanv");
// 	if(ptrgoglGetBooleanv == NULL) return 1;
// 	ptrgoglGetDoublev = goglGetProcAddress("glGetDoublev");
// 	if(ptrgoglGetDoublev == NULL) return 1;
// 	ptrgoglGetError = goglGetProcAddress("glGetError");
// 	if(ptrgoglGetError == NULL) return 1;
// 	ptrgoglGetFloatv = goglGetProcAddress("glGetFloatv");
// 	if(ptrgoglGetFloatv == NULL) return 1;
// 	ptrgoglGetIntegerv = goglGetProcAddress("glGetIntegerv");
// 	if(ptrgoglGetIntegerv == NULL) return 1;
// 	ptrgoglGetString = goglGetProcAddress("glGetString");
// 	if(ptrgoglGetString == NULL) return 1;
// 	ptrgoglGetTexImage = goglGetProcAddress("glGetTexImage");
// 	if(ptrgoglGetTexImage == NULL) return 1;
// 	ptrgoglGetTexParameterfv = goglGetProcAddress("glGetTexParameterfv");
// 	if(ptrgoglGetTexParameterfv == NULL) return 1;
// 	ptrgoglGetTexParameteriv = goglGetProcAddress("glGetTexParameteriv");
// 	if(ptrgoglGetTexParameteriv == NULL) return 1;
// 	ptrgoglGetTexLevelParameterfv = goglGetProcAddress("glGetTexLevelParameterfv");
// 	if(ptrgoglGetTexLevelParameterfv == NULL) return 1;
// 	ptrgoglGetTexLevelParameteriv = goglGetProcAddress("glGetTexLevelParameteriv");
// 	if(ptrgoglGetTexLevelParameteriv == NULL) return 1;
// 	ptrgoglIsEnabled = goglGetProcAddress("glIsEnabled");
// 	if(ptrgoglIsEnabled == NULL) return 1;
// 	ptrgoglDepthRange = goglGetProcAddress("glDepthRange");
// 	if(ptrgoglDepthRange == NULL) return 1;
// 	ptrgoglViewport = goglGetProcAddress("glViewport");
// 	if(ptrgoglViewport == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_1() {
// 	ptrgoglDrawArrays = goglGetProcAddress("glDrawArrays");
// 	if(ptrgoglDrawArrays == NULL) return 1;
// 	ptrgoglDrawElements = goglGetProcAddress("glDrawElements");
// 	if(ptrgoglDrawElements == NULL) return 1;
// 	ptrgoglGetPointerv = goglGetProcAddress("glGetPointerv");
// 	if(ptrgoglGetPointerv == NULL) return 1;
// 	ptrgoglPolygonOffset = goglGetProcAddress("glPolygonOffset");
// 	if(ptrgoglPolygonOffset == NULL) return 1;
// 	ptrgoglCopyTexImage1D = goglGetProcAddress("glCopyTexImage1D");
// 	if(ptrgoglCopyTexImage1D == NULL) return 1;
// 	ptrgoglCopyTexImage2D = goglGetProcAddress("glCopyTexImage2D");
// 	if(ptrgoglCopyTexImage2D == NULL) return 1;
// 	ptrgoglCopyTexSubImage1D = goglGetProcAddress("glCopyTexSubImage1D");
// 	if(ptrgoglCopyTexSubImage1D == NULL) return 1;
// 	ptrgoglCopyTexSubImage2D = goglGetProcAddress("glCopyTexSubImage2D");
// 	if(ptrgoglCopyTexSubImage2D == NULL) return 1;
// 	ptrgoglTexSubImage1D = goglGetProcAddress("glTexSubImage1D");
// 	if(ptrgoglTexSubImage1D == NULL) return 1;
// 	ptrgoglTexSubImage2D = goglGetProcAddress("glTexSubImage2D");
// 	if(ptrgoglTexSubImage2D == NULL) return 1;
// 	ptrgoglBindTexture = goglGetProcAddress("glBindTexture");
// 	if(ptrgoglBindTexture == NULL) return 1;
// 	ptrgoglDeleteTextures = goglGetProcAddress("glDeleteTextures");
// 	if(ptrgoglDeleteTextures == NULL) return 1;
// 	ptrgoglGenTextures = goglGetProcAddress("glGenTextures");
// 	if(ptrgoglGenTextures == NULL) return 1;
// 	ptrgoglIsTexture = goglGetProcAddress("glIsTexture");
// 	if(ptrgoglIsTexture == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_2() {
// 	ptrgoglBlendColor = goglGetProcAddress("glBlendColor");
// 	if(ptrgoglBlendColor == NULL) return 1;
// 	ptrgoglBlendEquation = goglGetProcAddress("glBlendEquation");
// 	if(ptrgoglBlendEquation == NULL) return 1;
// 	ptrgoglDrawRangeElements = goglGetProcAddress("glDrawRangeElements");
// 	if(ptrgoglDrawRangeElements == NULL) return 1;
// 	ptrgoglTexImage3D = goglGetProcAddress("glTexImage3D");
// 	if(ptrgoglTexImage3D == NULL) return 1;
// 	ptrgoglTexSubImage3D = goglGetProcAddress("glTexSubImage3D");
// 	if(ptrgoglTexSubImage3D == NULL) return 1;
// 	ptrgoglCopyTexSubImage3D = goglGetProcAddress("glCopyTexSubImage3D");
// 	if(ptrgoglCopyTexSubImage3D == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_3() {
// 	ptrgoglActiveTexture = goglGetProcAddress("glActiveTexture");
// 	if(ptrgoglActiveTexture == NULL) return 1;
// 	ptrgoglSampleCoverage = goglGetProcAddress("glSampleCoverage");
// 	if(ptrgoglSampleCoverage == NULL) return 1;
// 	ptrgoglCompressedTexImage3D = goglGetProcAddress("glCompressedTexImage3D");
// 	if(ptrgoglCompressedTexImage3D == NULL) return 1;
// 	ptrgoglCompressedTexImage2D = goglGetProcAddress("glCompressedTexImage2D");
// 	if(ptrgoglCompressedTexImage2D == NULL) return 1;
// 	ptrgoglCompressedTexImage1D = goglGetProcAddress("glCompressedTexImage1D");
// 	if(ptrgoglCompressedTexImage1D == NULL) return 1;
// 	ptrgoglCompressedTexSubImage3D = goglGetProcAddress("glCompressedTexSubImage3D");
// 	if(ptrgoglCompressedTexSubImage3D == NULL) return 1;
// 	ptrgoglCompressedTexSubImage2D = goglGetProcAddress("glCompressedTexSubImage2D");
// 	if(ptrgoglCompressedTexSubImage2D == NULL) return 1;
// 	ptrgoglCompressedTexSubImage1D = goglGetProcAddress("glCompressedTexSubImage1D");
// 	if(ptrgoglCompressedTexSubImage1D == NULL) return 1;
// 	ptrgoglGetCompressedTexImage = goglGetProcAddress("glGetCompressedTexImage");
// 	if(ptrgoglGetCompressedTexImage == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_3_2() {
// 	ptrgoglGetInteger64i_v = goglGetProcAddress("glGetInteger64i_v");
// 	if(ptrgoglGetInteger64i_v == NULL) return 1;
// 	ptrgoglGetBufferParameteri64v = goglGetProcAddress("glGetBufferParameteri64v");
// 	if(ptrgoglGetBufferParameteri64v == NULL) return 1;
// 	ptrgoglFramebufferTexture = goglGetProcAddress("glFramebufferTexture");
// 	if(ptrgoglFramebufferTexture == NULL) return 1;
// 	ptrgoglDrawElementsBaseVertex = goglGetProcAddress("glDrawElementsBaseVertex");
// 	if(ptrgoglDrawElementsBaseVertex == NULL) return 1;
// 	ptrgoglDrawRangeElementsBaseVertex = goglGetProcAddress("glDrawRangeElementsBaseVertex");
// 	if(ptrgoglDrawRangeElementsBaseVertex == NULL) return 1;
// 	ptrgoglDrawElementsInstancedBaseVertex = goglGetProcAddress("glDrawElementsInstancedBaseVertex");
// 	if(ptrgoglDrawElementsInstancedBaseVertex == NULL) return 1;
// 	ptrgoglMultiDrawElementsBaseVertex = goglGetProcAddress("glMultiDrawElementsBaseVertex");
// 	if(ptrgoglMultiDrawElementsBaseVertex == NULL) return 1;
// 	ptrgoglProvokingVertex = goglGetProcAddress("glProvokingVertex");
// 	if(ptrgoglProvokingVertex == NULL) return 1;
// 	ptrgoglFenceSync = goglGetProcAddress("glFenceSync");
// 	if(ptrgoglFenceSync == NULL) return 1;
// 	ptrgoglIsSync = goglGetProcAddress("glIsSync");
// 	if(ptrgoglIsSync == NULL) return 1;
// 	ptrgoglDeleteSync = goglGetProcAddress("glDeleteSync");
// 	if(ptrgoglDeleteSync == NULL) return 1;
// 	ptrgoglClientWaitSync = goglGetProcAddress("glClientWaitSync");
// 	if(ptrgoglClientWaitSync == NULL) return 1;
// 	ptrgoglWaitSync = goglGetProcAddress("glWaitSync");
// 	if(ptrgoglWaitSync == NULL) return 1;
// 	ptrgoglGetInteger64v = goglGetProcAddress("glGetInteger64v");
// 	if(ptrgoglGetInteger64v == NULL) return 1;
// 	ptrgoglGetSynciv = goglGetProcAddress("glGetSynciv");
// 	if(ptrgoglGetSynciv == NULL) return 1;
// 	ptrgoglTexImage2DMultisample = goglGetProcAddress("glTexImage2DMultisample");
// 	if(ptrgoglTexImage2DMultisample == NULL) return 1;
// 	ptrgoglTexImage3DMultisample = goglGetProcAddress("glTexImage3DMultisample");
// 	if(ptrgoglTexImage3DMultisample == NULL) return 1;
// 	ptrgoglGetMultisamplefv = goglGetProcAddress("glGetMultisamplefv");
// 	if(ptrgoglGetMultisamplefv == NULL) return 1;
// 	ptrgoglSampleMaski = goglGetProcAddress("glSampleMaski");
// 	if(ptrgoglSampleMaski == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_3_3() {
// 	ptrgoglVertexAttribDivisor = goglGetProcAddress("glVertexAttribDivisor");
// 	if(ptrgoglVertexAttribDivisor == NULL) return 1;
// 	ptrgoglBindFragDataLocationIndexed = goglGetProcAddress("glBindFragDataLocationIndexed");
// 	if(ptrgoglBindFragDataLocationIndexed == NULL) return 1;
// 	ptrgoglGetFragDataIndex = goglGetProcAddress("glGetFragDataIndex");
// 	if(ptrgoglGetFragDataIndex == NULL) return 1;
// 	ptrgoglGenSamplers = goglGetProcAddress("glGenSamplers");
// 	if(ptrgoglGenSamplers == NULL) return 1;
// 	ptrgoglDeleteSamplers = goglGetProcAddress("glDeleteSamplers");
// 	if(ptrgoglDeleteSamplers == NULL) return 1;
// 	ptrgoglIsSampler = goglGetProcAddress("glIsSampler");
// 	if(ptrgoglIsSampler == NULL) return 1;
// 	ptrgoglBindSampler = goglGetProcAddress("glBindSampler");
// 	if(ptrgoglBindSampler == NULL) return 1;
// 	ptrgoglSamplerParameteri = goglGetProcAddress("glSamplerParameteri");
// 	if(ptrgoglSamplerParameteri == NULL) return 1;
// 	ptrgoglSamplerParameteriv = goglGetProcAddress("glSamplerParameteriv");
// 	if(ptrgoglSamplerParameteriv == NULL) return 1;
// 	ptrgoglSamplerParameterf = goglGetProcAddress("glSamplerParameterf");
// 	if(ptrgoglSamplerParameterf == NULL) return 1;
// 	ptrgoglSamplerParameterfv = goglGetProcAddress("glSamplerParameterfv");
// 	if(ptrgoglSamplerParameterfv == NULL) return 1;
// 	ptrgoglSamplerParameterIiv = goglGetProcAddress("glSamplerParameterIiv");
// 	if(ptrgoglSamplerParameterIiv == NULL) return 1;
// 	ptrgoglSamplerParameterIuiv = goglGetProcAddress("glSamplerParameterIuiv");
// 	if(ptrgoglSamplerParameterIuiv == NULL) return 1;
// 	ptrgoglGetSamplerParameteriv = goglGetProcAddress("glGetSamplerParameteriv");
// 	if(ptrgoglGetSamplerParameteriv == NULL) return 1;
// 	ptrgoglGetSamplerParameterIiv = goglGetProcAddress("glGetSamplerParameterIiv");
// 	if(ptrgoglGetSamplerParameterIiv == NULL) return 1;
// 	ptrgoglGetSamplerParameterfv = goglGetProcAddress("glGetSamplerParameterfv");
// 	if(ptrgoglGetSamplerParameterfv == NULL) return 1;
// 	ptrgoglGetSamplerParameterIuiv = goglGetProcAddress("glGetSamplerParameterIuiv");
// 	if(ptrgoglGetSamplerParameterIuiv == NULL) return 1;
// 	ptrgoglQueryCounter = goglGetProcAddress("glQueryCounter");
// 	if(ptrgoglQueryCounter == NULL) return 1;
// 	ptrgoglGetQueryObjecti64v = goglGetProcAddress("glGetQueryObjecti64v");
// 	if(ptrgoglGetQueryObjecti64v == NULL) return 1;
// 	ptrgoglGetQueryObjectui64v = goglGetProcAddress("glGetQueryObjectui64v");
// 	if(ptrgoglGetQueryObjectui64v == NULL) return 1;
// 	ptrgoglVertexP2ui = goglGetProcAddress("glVertexP2ui");
// 	if(ptrgoglVertexP2ui == NULL) return 1;
// 	ptrgoglVertexP2uiv = goglGetProcAddress("glVertexP2uiv");
// 	if(ptrgoglVertexP2uiv == NULL) return 1;
// 	ptrgoglVertexP3ui = goglGetProcAddress("glVertexP3ui");
// 	if(ptrgoglVertexP3ui == NULL) return 1;
// 	ptrgoglVertexP3uiv = goglGetProcAddress("glVertexP3uiv");
// 	if(ptrgoglVertexP3uiv == NULL) return 1;
// 	ptrgoglVertexP4ui = goglGetProcAddress("glVertexP4ui");
// 	if(ptrgoglVertexP4ui == NULL) return 1;
// 	ptrgoglVertexP4uiv = goglGetProcAddress("glVertexP4uiv");
// 	if(ptrgoglVertexP4uiv == NULL) return 1;
// 	ptrgoglTexCoordP1ui = goglGetProcAddress("glTexCoordP1ui");
// 	if(ptrgoglTexCoordP1ui == NULL) return 1;
// 	ptrgoglTexCoordP1uiv = goglGetProcAddress("glTexCoordP1uiv");
// 	if(ptrgoglTexCoordP1uiv == NULL) return 1;
// 	ptrgoglTexCoordP2ui = goglGetProcAddress("glTexCoordP2ui");
// 	if(ptrgoglTexCoordP2ui == NULL) return 1;
// 	ptrgoglTexCoordP2uiv = goglGetProcAddress("glTexCoordP2uiv");
// 	if(ptrgoglTexCoordP2uiv == NULL) return 1;
// 	ptrgoglTexCoordP3ui = goglGetProcAddress("glTexCoordP3ui");
// 	if(ptrgoglTexCoordP3ui == NULL) return 1;
// 	ptrgoglTexCoordP3uiv = goglGetProcAddress("glTexCoordP3uiv");
// 	if(ptrgoglTexCoordP3uiv == NULL) return 1;
// 	ptrgoglTexCoordP4ui = goglGetProcAddress("glTexCoordP4ui");
// 	if(ptrgoglTexCoordP4ui == NULL) return 1;
// 	ptrgoglTexCoordP4uiv = goglGetProcAddress("glTexCoordP4uiv");
// 	if(ptrgoglTexCoordP4uiv == NULL) return 1;
// 	ptrgoglMultiTexCoordP1ui = goglGetProcAddress("glMultiTexCoordP1ui");
// 	if(ptrgoglMultiTexCoordP1ui == NULL) return 1;
// 	ptrgoglMultiTexCoordP1uiv = goglGetProcAddress("glMultiTexCoordP1uiv");
// 	if(ptrgoglMultiTexCoordP1uiv == NULL) return 1;
// 	ptrgoglMultiTexCoordP2ui = goglGetProcAddress("glMultiTexCoordP2ui");
// 	if(ptrgoglMultiTexCoordP2ui == NULL) return 1;
// 	ptrgoglMultiTexCoordP2uiv = goglGetProcAddress("glMultiTexCoordP2uiv");
// 	if(ptrgoglMultiTexCoordP2uiv == NULL) return 1;
// 	ptrgoglMultiTexCoordP3ui = goglGetProcAddress("glMultiTexCoordP3ui");
// 	if(ptrgoglMultiTexCoordP3ui == NULL) return 1;
// 	ptrgoglMultiTexCoordP3uiv = goglGetProcAddress("glMultiTexCoordP3uiv");
// 	if(ptrgoglMultiTexCoordP3uiv == NULL) return 1;
// 	ptrgoglMultiTexCoordP4ui = goglGetProcAddress("glMultiTexCoordP4ui");
// 	if(ptrgoglMultiTexCoordP4ui == NULL) return 1;
// 	ptrgoglMultiTexCoordP4uiv = goglGetProcAddress("glMultiTexCoordP4uiv");
// 	if(ptrgoglMultiTexCoordP4uiv == NULL) return 1;
// 	ptrgoglNormalP3ui = goglGetProcAddress("glNormalP3ui");
// 	if(ptrgoglNormalP3ui == NULL) return 1;
// 	ptrgoglNormalP3uiv = goglGetProcAddress("glNormalP3uiv");
// 	if(ptrgoglNormalP3uiv == NULL) return 1;
// 	ptrgoglColorP3ui = goglGetProcAddress("glColorP3ui");
// 	if(ptrgoglColorP3ui == NULL) return 1;
// 	ptrgoglColorP3uiv = goglGetProcAddress("glColorP3uiv");
// 	if(ptrgoglColorP3uiv == NULL) return 1;
// 	ptrgoglColorP4ui = goglGetProcAddress("glColorP4ui");
// 	if(ptrgoglColorP4ui == NULL) return 1;
// 	ptrgoglColorP4uiv = goglGetProcAddress("glColorP4uiv");
// 	if(ptrgoglColorP4uiv == NULL) return 1;
// 	ptrgoglSecondaryColorP3ui = goglGetProcAddress("glSecondaryColorP3ui");
// 	if(ptrgoglSecondaryColorP3ui == NULL) return 1;
// 	ptrgoglSecondaryColorP3uiv = goglGetProcAddress("glSecondaryColorP3uiv");
// 	if(ptrgoglSecondaryColorP3uiv == NULL) return 1;
// 	ptrgoglVertexAttribP1ui = goglGetProcAddress("glVertexAttribP1ui");
// 	if(ptrgoglVertexAttribP1ui == NULL) return 1;
// 	ptrgoglVertexAttribP1uiv = goglGetProcAddress("glVertexAttribP1uiv");
// 	if(ptrgoglVertexAttribP1uiv == NULL) return 1;
// 	ptrgoglVertexAttribP2ui = goglGetProcAddress("glVertexAttribP2ui");
// 	if(ptrgoglVertexAttribP2ui == NULL) return 1;
// 	ptrgoglVertexAttribP2uiv = goglGetProcAddress("glVertexAttribP2uiv");
// 	if(ptrgoglVertexAttribP2uiv == NULL) return 1;
// 	ptrgoglVertexAttribP3ui = goglGetProcAddress("glVertexAttribP3ui");
// 	if(ptrgoglVertexAttribP3ui == NULL) return 1;
// 	ptrgoglVertexAttribP3uiv = goglGetProcAddress("glVertexAttribP3uiv");
// 	if(ptrgoglVertexAttribP3uiv == NULL) return 1;
// 	ptrgoglVertexAttribP4ui = goglGetProcAddress("glVertexAttribP4ui");
// 	if(ptrgoglVertexAttribP4ui == NULL) return 1;
// 	ptrgoglVertexAttribP4uiv = goglGetProcAddress("glVertexAttribP4uiv");
// 	if(ptrgoglVertexAttribP4uiv == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_3_0() {
// 	ptrgoglColorMaski = goglGetProcAddress("glColorMaski");
// 	if(ptrgoglColorMaski == NULL) return 1;
// 	ptrgoglGetBooleani_v = goglGetProcAddress("glGetBooleani_v");
// 	if(ptrgoglGetBooleani_v == NULL) return 1;
// 	ptrgoglGetIntegeri_v = goglGetProcAddress("glGetIntegeri_v");
// 	if(ptrgoglGetIntegeri_v == NULL) return 1;
// 	ptrgoglEnablei = goglGetProcAddress("glEnablei");
// 	if(ptrgoglEnablei == NULL) return 1;
// 	ptrgoglDisablei = goglGetProcAddress("glDisablei");
// 	if(ptrgoglDisablei == NULL) return 1;
// 	ptrgoglIsEnabledi = goglGetProcAddress("glIsEnabledi");
// 	if(ptrgoglIsEnabledi == NULL) return 1;
// 	ptrgoglBeginTransformFeedback = goglGetProcAddress("glBeginTransformFeedback");
// 	if(ptrgoglBeginTransformFeedback == NULL) return 1;
// 	ptrgoglEndTransformFeedback = goglGetProcAddress("glEndTransformFeedback");
// 	if(ptrgoglEndTransformFeedback == NULL) return 1;
// 	ptrgoglBindBufferRange = goglGetProcAddress("glBindBufferRange");
// 	if(ptrgoglBindBufferRange == NULL) return 1;
// 	ptrgoglBindBufferBase = goglGetProcAddress("glBindBufferBase");
// 	if(ptrgoglBindBufferBase == NULL) return 1;
// 	ptrgoglTransformFeedbackVaryings = goglGetProcAddress("glTransformFeedbackVaryings");
// 	if(ptrgoglTransformFeedbackVaryings == NULL) return 1;
// 	ptrgoglGetTransformFeedbackVarying = goglGetProcAddress("glGetTransformFeedbackVarying");
// 	if(ptrgoglGetTransformFeedbackVarying == NULL) return 1;
// 	ptrgoglClampColor = goglGetProcAddress("glClampColor");
// 	if(ptrgoglClampColor == NULL) return 1;
// 	ptrgoglBeginConditionalRender = goglGetProcAddress("glBeginConditionalRender");
// 	if(ptrgoglBeginConditionalRender == NULL) return 1;
// 	ptrgoglEndConditionalRender = goglGetProcAddress("glEndConditionalRender");
// 	if(ptrgoglEndConditionalRender == NULL) return 1;
// 	ptrgoglVertexAttribIPointer = goglGetProcAddress("glVertexAttribIPointer");
// 	if(ptrgoglVertexAttribIPointer == NULL) return 1;
// 	ptrgoglGetVertexAttribIiv = goglGetProcAddress("glGetVertexAttribIiv");
// 	if(ptrgoglGetVertexAttribIiv == NULL) return 1;
// 	ptrgoglGetVertexAttribIuiv = goglGetProcAddress("glGetVertexAttribIuiv");
// 	if(ptrgoglGetVertexAttribIuiv == NULL) return 1;
// 	ptrgoglVertexAttribI1i = goglGetProcAddress("glVertexAttribI1i");
// 	if(ptrgoglVertexAttribI1i == NULL) return 1;
// 	ptrgoglVertexAttribI2i = goglGetProcAddress("glVertexAttribI2i");
// 	if(ptrgoglVertexAttribI2i == NULL) return 1;
// 	ptrgoglVertexAttribI3i = goglGetProcAddress("glVertexAttribI3i");
// 	if(ptrgoglVertexAttribI3i == NULL) return 1;
// 	ptrgoglVertexAttribI4i = goglGetProcAddress("glVertexAttribI4i");
// 	if(ptrgoglVertexAttribI4i == NULL) return 1;
// 	ptrgoglVertexAttribI1ui = goglGetProcAddress("glVertexAttribI1ui");
// 	if(ptrgoglVertexAttribI1ui == NULL) return 1;
// 	ptrgoglVertexAttribI2ui = goglGetProcAddress("glVertexAttribI2ui");
// 	if(ptrgoglVertexAttribI2ui == NULL) return 1;
// 	ptrgoglVertexAttribI3ui = goglGetProcAddress("glVertexAttribI3ui");
// 	if(ptrgoglVertexAttribI3ui == NULL) return 1;
// 	ptrgoglVertexAttribI4ui = goglGetProcAddress("glVertexAttribI4ui");
// 	if(ptrgoglVertexAttribI4ui == NULL) return 1;
// 	ptrgoglVertexAttribI1iv = goglGetProcAddress("glVertexAttribI1iv");
// 	if(ptrgoglVertexAttribI1iv == NULL) return 1;
// 	ptrgoglVertexAttribI2iv = goglGetProcAddress("glVertexAttribI2iv");
// 	if(ptrgoglVertexAttribI2iv == NULL) return 1;
// 	ptrgoglVertexAttribI3iv = goglGetProcAddress("glVertexAttribI3iv");
// 	if(ptrgoglVertexAttribI3iv == NULL) return 1;
// 	ptrgoglVertexAttribI4iv = goglGetProcAddress("glVertexAttribI4iv");
// 	if(ptrgoglVertexAttribI4iv == NULL) return 1;
// 	ptrgoglVertexAttribI1uiv = goglGetProcAddress("glVertexAttribI1uiv");
// 	if(ptrgoglVertexAttribI1uiv == NULL) return 1;
// 	ptrgoglVertexAttribI2uiv = goglGetProcAddress("glVertexAttribI2uiv");
// 	if(ptrgoglVertexAttribI2uiv == NULL) return 1;
// 	ptrgoglVertexAttribI3uiv = goglGetProcAddress("glVertexAttribI3uiv");
// 	if(ptrgoglVertexAttribI3uiv == NULL) return 1;
// 	ptrgoglVertexAttribI4uiv = goglGetProcAddress("glVertexAttribI4uiv");
// 	if(ptrgoglVertexAttribI4uiv == NULL) return 1;
// 	ptrgoglVertexAttribI4bv = goglGetProcAddress("glVertexAttribI4bv");
// 	if(ptrgoglVertexAttribI4bv == NULL) return 1;
// 	ptrgoglVertexAttribI4sv = goglGetProcAddress("glVertexAttribI4sv");
// 	if(ptrgoglVertexAttribI4sv == NULL) return 1;
// 	ptrgoglVertexAttribI4ubv = goglGetProcAddress("glVertexAttribI4ubv");
// 	if(ptrgoglVertexAttribI4ubv == NULL) return 1;
// 	ptrgoglVertexAttribI4usv = goglGetProcAddress("glVertexAttribI4usv");
// 	if(ptrgoglVertexAttribI4usv == NULL) return 1;
// 	ptrgoglGetUniformuiv = goglGetProcAddress("glGetUniformuiv");
// 	if(ptrgoglGetUniformuiv == NULL) return 1;
// 	ptrgoglBindFragDataLocation = goglGetProcAddress("glBindFragDataLocation");
// 	if(ptrgoglBindFragDataLocation == NULL) return 1;
// 	ptrgoglGetFragDataLocation = goglGetProcAddress("glGetFragDataLocation");
// 	if(ptrgoglGetFragDataLocation == NULL) return 1;
// 	ptrgoglUniform1ui = goglGetProcAddress("glUniform1ui");
// 	if(ptrgoglUniform1ui == NULL) return 1;
// 	ptrgoglUniform2ui = goglGetProcAddress("glUniform2ui");
// 	if(ptrgoglUniform2ui == NULL) return 1;
// 	ptrgoglUniform3ui = goglGetProcAddress("glUniform3ui");
// 	if(ptrgoglUniform3ui == NULL) return 1;
// 	ptrgoglUniform4ui = goglGetProcAddress("glUniform4ui");
// 	if(ptrgoglUniform4ui == NULL) return 1;
// 	ptrgoglUniform1uiv = goglGetProcAddress("glUniform1uiv");
// 	if(ptrgoglUniform1uiv == NULL) return 1;
// 	ptrgoglUniform2uiv = goglGetProcAddress("glUniform2uiv");
// 	if(ptrgoglUniform2uiv == NULL) return 1;
// 	ptrgoglUniform3uiv = goglGetProcAddress("glUniform3uiv");
// 	if(ptrgoglUniform3uiv == NULL) return 1;
// 	ptrgoglUniform4uiv = goglGetProcAddress("glUniform4uiv");
// 	if(ptrgoglUniform4uiv == NULL) return 1;
// 	ptrgoglTexParameterIiv = goglGetProcAddress("glTexParameterIiv");
// 	if(ptrgoglTexParameterIiv == NULL) return 1;
// 	ptrgoglTexParameterIuiv = goglGetProcAddress("glTexParameterIuiv");
// 	if(ptrgoglTexParameterIuiv == NULL) return 1;
// 	ptrgoglGetTexParameterIiv = goglGetProcAddress("glGetTexParameterIiv");
// 	if(ptrgoglGetTexParameterIiv == NULL) return 1;
// 	ptrgoglGetTexParameterIuiv = goglGetProcAddress("glGetTexParameterIuiv");
// 	if(ptrgoglGetTexParameterIuiv == NULL) return 1;
// 	ptrgoglClearBufferiv = goglGetProcAddress("glClearBufferiv");
// 	if(ptrgoglClearBufferiv == NULL) return 1;
// 	ptrgoglClearBufferuiv = goglGetProcAddress("glClearBufferuiv");
// 	if(ptrgoglClearBufferuiv == NULL) return 1;
// 	ptrgoglClearBufferfv = goglGetProcAddress("glClearBufferfv");
// 	if(ptrgoglClearBufferfv == NULL) return 1;
// 	ptrgoglClearBufferfi = goglGetProcAddress("glClearBufferfi");
// 	if(ptrgoglClearBufferfi == NULL) return 1;
// 	ptrgoglGetStringi = goglGetProcAddress("glGetStringi");
// 	if(ptrgoglGetStringi == NULL) return 1;
// 	ptrgoglIsRenderbuffer = goglGetProcAddress("glIsRenderbuffer");
// 	if(ptrgoglIsRenderbuffer == NULL) return 1;
// 	ptrgoglBindRenderbuffer = goglGetProcAddress("glBindRenderbuffer");
// 	if(ptrgoglBindRenderbuffer == NULL) return 1;
// 	ptrgoglDeleteRenderbuffers = goglGetProcAddress("glDeleteRenderbuffers");
// 	if(ptrgoglDeleteRenderbuffers == NULL) return 1;
// 	ptrgoglGenRenderbuffers = goglGetProcAddress("glGenRenderbuffers");
// 	if(ptrgoglGenRenderbuffers == NULL) return 1;
// 	ptrgoglRenderbufferStorage = goglGetProcAddress("glRenderbufferStorage");
// 	if(ptrgoglRenderbufferStorage == NULL) return 1;
// 	ptrgoglGetRenderbufferParameteriv = goglGetProcAddress("glGetRenderbufferParameteriv");
// 	if(ptrgoglGetRenderbufferParameteriv == NULL) return 1;
// 	ptrgoglIsFramebuffer = goglGetProcAddress("glIsFramebuffer");
// 	if(ptrgoglIsFramebuffer == NULL) return 1;
// 	ptrgoglBindFramebuffer = goglGetProcAddress("glBindFramebuffer");
// 	if(ptrgoglBindFramebuffer == NULL) return 1;
// 	ptrgoglDeleteFramebuffers = goglGetProcAddress("glDeleteFramebuffers");
// 	if(ptrgoglDeleteFramebuffers == NULL) return 1;
// 	ptrgoglGenFramebuffers = goglGetProcAddress("glGenFramebuffers");
// 	if(ptrgoglGenFramebuffers == NULL) return 1;
// 	ptrgoglCheckFramebufferStatus = goglGetProcAddress("glCheckFramebufferStatus");
// 	if(ptrgoglCheckFramebufferStatus == NULL) return 1;
// 	ptrgoglFramebufferTexture1D = goglGetProcAddress("glFramebufferTexture1D");
// 	if(ptrgoglFramebufferTexture1D == NULL) return 1;
// 	ptrgoglFramebufferTexture2D = goglGetProcAddress("glFramebufferTexture2D");
// 	if(ptrgoglFramebufferTexture2D == NULL) return 1;
// 	ptrgoglFramebufferTexture3D = goglGetProcAddress("glFramebufferTexture3D");
// 	if(ptrgoglFramebufferTexture3D == NULL) return 1;
// 	ptrgoglFramebufferRenderbuffer = goglGetProcAddress("glFramebufferRenderbuffer");
// 	if(ptrgoglFramebufferRenderbuffer == NULL) return 1;
// 	ptrgoglGetFramebufferAttachmentParameteriv = goglGetProcAddress("glGetFramebufferAttachmentParameteriv");
// 	if(ptrgoglGetFramebufferAttachmentParameteriv == NULL) return 1;
// 	ptrgoglGenerateMipmap = goglGetProcAddress("glGenerateMipmap");
// 	if(ptrgoglGenerateMipmap == NULL) return 1;
// 	ptrgoglBlitFramebuffer = goglGetProcAddress("glBlitFramebuffer");
// 	if(ptrgoglBlitFramebuffer == NULL) return 1;
// 	ptrgoglRenderbufferStorageMultisample = goglGetProcAddress("glRenderbufferStorageMultisample");
// 	if(ptrgoglRenderbufferStorageMultisample == NULL) return 1;
// 	ptrgoglFramebufferTextureLayer = goglGetProcAddress("glFramebufferTextureLayer");
// 	if(ptrgoglFramebufferTextureLayer == NULL) return 1;
// 	ptrgoglMapBufferRange = goglGetProcAddress("glMapBufferRange");
// 	if(ptrgoglMapBufferRange == NULL) return 1;
// 	ptrgoglFlushMappedBufferRange = goglGetProcAddress("glFlushMappedBufferRange");
// 	if(ptrgoglFlushMappedBufferRange == NULL) return 1;
// 	ptrgoglBindVertexArray = goglGetProcAddress("glBindVertexArray");
// 	if(ptrgoglBindVertexArray == NULL) return 1;
// 	ptrgoglDeleteVertexArrays = goglGetProcAddress("glDeleteVertexArrays");
// 	if(ptrgoglDeleteVertexArrays == NULL) return 1;
// 	ptrgoglGenVertexArrays = goglGetProcAddress("glGenVertexArrays");
// 	if(ptrgoglGenVertexArrays == NULL) return 1;
// 	ptrgoglIsVertexArray = goglGetProcAddress("glIsVertexArray");
// 	if(ptrgoglIsVertexArray == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_3_1() {
// 	ptrgoglDrawArraysInstanced = goglGetProcAddress("glDrawArraysInstanced");
// 	if(ptrgoglDrawArraysInstanced == NULL) return 1;
// 	ptrgoglDrawElementsInstanced = goglGetProcAddress("glDrawElementsInstanced");
// 	if(ptrgoglDrawElementsInstanced == NULL) return 1;
// 	ptrgoglTexBuffer = goglGetProcAddress("glTexBuffer");
// 	if(ptrgoglTexBuffer == NULL) return 1;
// 	ptrgoglPrimitiveRestartIndex = goglGetProcAddress("glPrimitiveRestartIndex");
// 	if(ptrgoglPrimitiveRestartIndex == NULL) return 1;
// 	ptrgoglCopyBufferSubData = goglGetProcAddress("glCopyBufferSubData");
// 	if(ptrgoglCopyBufferSubData == NULL) return 1;
// 	ptrgoglGetUniformIndices = goglGetProcAddress("glGetUniformIndices");
// 	if(ptrgoglGetUniformIndices == NULL) return 1;
// 	ptrgoglGetActiveUniformsiv = goglGetProcAddress("glGetActiveUniformsiv");
// 	if(ptrgoglGetActiveUniformsiv == NULL) return 1;
// 	ptrgoglGetActiveUniformName = goglGetProcAddress("glGetActiveUniformName");
// 	if(ptrgoglGetActiveUniformName == NULL) return 1;
// 	ptrgoglGetUniformBlockIndex = goglGetProcAddress("glGetUniformBlockIndex");
// 	if(ptrgoglGetUniformBlockIndex == NULL) return 1;
// 	ptrgoglGetActiveUniformBlockiv = goglGetProcAddress("glGetActiveUniformBlockiv");
// 	if(ptrgoglGetActiveUniformBlockiv == NULL) return 1;
// 	ptrgoglGetActiveUniformBlockName = goglGetProcAddress("glGetActiveUniformBlockName");
// 	if(ptrgoglGetActiveUniformBlockName == NULL) return 1;
// 	ptrgoglUniformBlockBinding = goglGetProcAddress("glUniformBlockBinding");
// 	if(ptrgoglUniformBlockBinding == NULL) return 1;
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

// VERSION_2_1
const (
	PIXEL_PACK_BUFFER = 0x88EB
	SRGB8_ALPHA8 = 0x8C43
	PIXEL_PACK_BUFFER_BINDING = 0x88ED
	SRGB8 = 0x8C41
	PIXEL_UNPACK_BUFFER_BINDING = 0x88EF
	COMPRESSED_SRGB_ALPHA = 0x8C49
	FLOAT_MAT4x2 = 0x8B69
	FLOAT_MAT4x3 = 0x8B6A
	FLOAT_MAT3x4 = 0x8B68
	FLOAT_MAT3x2 = 0x8B67
	COMPRESSED_SRGB = 0x8C48
	SRGB = 0x8C40
	SRGB_ALPHA = 0x8C42
	FLOAT_MAT2x4 = 0x8B66
	FLOAT_MAT2x3 = 0x8B65
	PIXEL_UNPACK_BUFFER = 0x88EC
)
// VERSION_2_0
const (
	STENCIL_BACK_VALUE_MASK = 0x8CA4
	VERTEX_ATTRIB_ARRAY_SIZE = 0x8623
	STENCIL_BACK_PASS_DEPTH_PASS = 0x8803
	VERTEX_SHADER = 0x8B31
	VERTEX_ATTRIB_ARRAY_TYPE = 0x8625
	INT_VEC4 = 0x8B55
	INT_VEC3 = 0x8B54
	INT_VEC2 = 0x8B53
	MAX_FRAGMENT_UNIFORM_COMPONENTS = 0x8B49
	MAX_DRAW_BUFFERS = 0x8824
	VERTEX_ATTRIB_ARRAY_NORMALIZED = 0x886A
	INFO_LOG_LENGTH = 0x8B84
	DELETE_STATUS = 0x8B80
	FLOAT_VEC4 = 0x8B52
	FLOAT_VEC2 = 0x8B50
	FLOAT_VEC3 = 0x8B51
	STENCIL_BACK_REF = 0x8CA3
	ACTIVE_ATTRIBUTE_MAX_LENGTH = 0x8B8A
	SHADING_LANGUAGE_VERSION = 0x8B8C
	MAX_COMBINED_TEXTURE_IMAGE_UNITS = 0x8B4D
	ACTIVE_UNIFORMS = 0x8B86
	ATTACHED_SHADERS = 0x8B85
	VERTEX_ATTRIB_ARRAY_ENABLED = 0x8622
	MAX_VERTEX_UNIFORM_COMPONENTS = 0x8B4A
	MAX_TEXTURE_IMAGE_UNITS = 0x8872
	STENCIL_BACK_FAIL = 0x8801
	VERTEX_ATTRIB_ARRAY_POINTER = 0x8645
	MAX_VARYING_FLOATS = 0x8B4B
	CURRENT_VERTEX_ATTRIB = 0x8626
	SAMPLER_1D_SHADOW = 0x8B61
	DRAW_BUFFER15 = 0x8834
	DRAW_BUFFER14 = 0x8833
	DRAW_BUFFER11 = 0x8830
	DRAW_BUFFER10 = 0x882F
	DRAW_BUFFER13 = 0x8832
	DRAW_BUFFER12 = 0x8831
	LINK_STATUS = 0x8B82
	CURRENT_PROGRAM = 0x8B8D
	LOWER_LEFT = 0x8CA1
	MAX_VERTEX_TEXTURE_IMAGE_UNITS = 0x8B4C
	VALIDATE_STATUS = 0x8B83
	STENCIL_BACK_WRITEMASK = 0x8CA5
	FLOAT_MAT4 = 0x8B5C
	COMPILE_STATUS = 0x8B81
	FLOAT_MAT2 = 0x8B5A
	SAMPLER_3D = 0x8B5F
	FLOAT_MAT3 = 0x8B5B
	SHADER_TYPE = 0x8B4F
	MAX_VERTEX_ATTRIBS = 0x8869
	SAMPLER_1D = 0x8B5D
	SAMPLER_2D_SHADOW = 0x8B62
	SAMPLER_CUBE = 0x8B60
	STENCIL_BACK_PASS_DEPTH_FAIL = 0x8802
	BOOL_VEC2 = 0x8B57
	BOOL_VEC3 = 0x8B58
	BOOL_VEC4 = 0x8B59
	FRAGMENT_SHADER = 0x8B30
	STENCIL_BACK_FUNC = 0x8800
	FRAGMENT_SHADER_DERIVATIVE_HINT = 0x8B8B
	BOOL = 0x8B56
	VERTEX_PROGRAM_POINT_SIZE = 0x8642
	UPPER_LEFT = 0x8CA2
	VERTEX_ATTRIB_ARRAY_STRIDE = 0x8624
	SHADER_SOURCE_LENGTH = 0x8B88
	ACTIVE_UNIFORM_MAX_LENGTH = 0x8B87
	SAMPLER_2D = 0x8B5E
	POINT_SPRITE_COORD_ORIGIN = 0x8CA0
	ACTIVE_ATTRIBUTES = 0x8B89
	DRAW_BUFFER2 = 0x8827
	DRAW_BUFFER3 = 0x8828
	DRAW_BUFFER0 = 0x8825
	DRAW_BUFFER1 = 0x8826
	DRAW_BUFFER6 = 0x882B
	DRAW_BUFFER7 = 0x882C
	DRAW_BUFFER4 = 0x8829
	DRAW_BUFFER5 = 0x882A
	DRAW_BUFFER8 = 0x882D
	DRAW_BUFFER9 = 0x882E
	BLEND_EQUATION_ALPHA = 0x883D
	BLEND_EQUATION_RGB = 0x8009
)
// VERSION_1_4
const (
	DEPTH_COMPONENT16 = 0x81A5
	MIRRORED_REPEAT = 0x8370
	BLEND_DST_RGB = 0x80C8
	DEPTH_COMPONENT32 = 0x81A7
	DEPTH_COMPONENT24 = 0x81A6
	TEXTURE_COMPARE_FUNC = 0x884D
	DECR_WRAP = 0x8508
	BLEND_SRC_RGB = 0x80C9
	TEXTURE_COMPARE_MODE = 0x884C
	MAX_TEXTURE_LOD_BIAS = 0x84FD
	TEXTURE_DEPTH_SIZE = 0x884A
	BLEND_DST_ALPHA = 0x80CA
	POINT_FADE_THRESHOLD_SIZE = 0x8128
	INCR_WRAP = 0x8507
	TEXTURE_LOD_BIAS = 0x8501
	BLEND_SRC_ALPHA = 0x80CB
)
// VERSION_1_5
const (
	ARRAY_BUFFER_BINDING = 0x8894
	STATIC_COPY = 0x88E6
	BUFFER_ACCESS = 0x88BB
	READ_ONLY = 0x88B8
	CURRENT_QUERY = 0x8865
	DYNAMIC_DRAW = 0x88E8
	READ_WRITE = 0x88BA
	STATIC_DRAW = 0x88E4
	ARRAY_BUFFER = 0x8892
	ELEMENT_ARRAY_BUFFER_BINDING = 0x8895
	BUFFER_SIZE = 0x8764
	STREAM_COPY = 0x88E2
	DYNAMIC_READ = 0x88E9
	BUFFER_MAP_POINTER = 0x88BD
	DYNAMIC_COPY = 0x88EA
	QUERY_COUNTER_BITS = 0x8864
	ELEMENT_ARRAY_BUFFER = 0x8893
	QUERY_RESULT_AVAILABLE = 0x8867
	STREAM_DRAW = 0x88E0
	SAMPLES_PASSED = 0x8914
	WRITE_ONLY = 0x88B9
	QUERY_RESULT = 0x8866
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = 0x889F
	BUFFER_MAPPED = 0x88BC
	STREAM_READ = 0x88E1
	BUFFER_USAGE = 0x8765
	STATIC_READ = 0x88E5
)
// VERSION_1_1
const (
	NEAREST_MIPMAP_LINEAR = 0x2702
	POINTS = 0x0000
	FILL = 0x1B02
	AND_REVERSE = 0x1502
	CLEAR = 0x1500
	DEPTH_FUNC = 0x0B74
	TEXTURE_WIDTH = 0x1000
	CCW = 0x0901
	GEQUAL = 0x0206
	OR_REVERSE = 0x150B
	BLUE = 0x1905
	RGB5_A1 = 0x8057
	EXTENSIONS = 0x1F03
	DRAW_BUFFER = 0x0C01
	DONT_CARE = 0x1100
	LEFT = 0x0406
	PROXY_TEXTURE_1D = 0x8063
	DOUBLE = 0x140A
	TRIANGLE_FAN = 0x0006
	POLYGON_OFFSET_LINE = 0x2A02
	NONE = 0
	TEXTURE_GREEN_SIZE = 0x805D
	CW = 0x0900
	POINT_SIZE_RANGE = 0x0B12
	BLEND_DST = 0x0BE0
	FRONT_FACE = 0x0B46
	DOUBLEBUFFER = 0x0C32
	BACK_LEFT = 0x0402
	VERSION = 0x1F02
	INCR = 0x1E02
	ONE_MINUS_DST_COLOR = 0x0307
	COLOR_BUFFER_BIT = 0x00004000
	INVERT = 0x150A
	ONE_MINUS_SRC_COLOR = 0x0301
	POINT = 0x1B00
	LEQUAL = 0x0203
	PACK_LSB_FIRST = 0x0D01
	GREATER = 0x0204
	TEXTURE_BINDING_1D = 0x8068
	UNSIGNED_BYTE = 0x1401
	NEAREST = 0x2600
	TEXTURE_MAG_FILTER = 0x2800
	KEEP = 0x1E00
	POLYGON_OFFSET_FACTOR = 0x8038
	PROXY_TEXTURE_2D = 0x8064
	PACK_SWAP_BYTES = 0x0D00
	RGBA2 = 0x8055
	ONE_MINUS_DST_ALPHA = 0x0305
	CULL_FACE_MODE = 0x0B45
	RGBA4 = 0x8056
	UNSIGNED_INT = 0x1405
	NOOP = 0x1505
	RGBA8 = 0x8058
	TEXTURE_BORDER_COLOR = 0x1004
	AND = 0x1501
	RGBA16 = 0x805B
	RGBA12 = 0x805A
	COLOR_LOGIC_OP = 0x0BF2
	GREEN = 0x1904
	INVALID_ENUM = 0x0500
	TEXTURE_RED_SIZE = 0x805C
	COPY_INVERTED = 0x150C
	UNPACK_ALIGNMENT = 0x0CF5
	LINES = 0x0001
	SHORT = 0x1402
	SRC_ALPHA = 0x0302
	LINEAR_MIPMAP_NEAREST = 0x2701
	RGB10_A2 = 0x8059
	SCISSOR_TEST = 0x0C11
	LINEAR = 0x2601
	TEXTURE_ALPHA_SIZE = 0x805F
	TRIANGLE_STRIP = 0x0005
	DECR = 0x1E03
	REPEAT = 0x2901
	DEPTH = 0x1801
	STENCIL_FUNC = 0x0B92
	OUT_OF_MEMORY = 0x0505
	POINT_SIZE_GRANULARITY = 0x0B13
	INT = 0x1404
	BYTE = 0x1400
	OR_INVERTED = 0x150D
	TRIANGLES = 0x0004
	TEXTURE_MIN_FILTER = 0x2801
	FRONT_AND_BACK = 0x0408
	DEPTH_TEST = 0x0B71
	FRONT = 0x0404
	LOGIC_OP_MODE = 0x0BF0
	UNPACK_SWAP_BYTES = 0x0CF0
	TEXTURE_1D = 0x0DE0
	UNPACK_SKIP_PIXELS = 0x0CF4
	DEPTH_RANGE = 0x0B70
	TEXTURE_BINDING_2D = 0x8069
	COLOR_WRITEMASK = 0x0C23
	PACK_ROW_LENGTH = 0x0D02
	NOR = 0x1508
	TRUE = 1
	VIEWPORT = 0x0BA2
	TEXTURE_INTERNAL_FORMAT = 0x1003
	NAND = 0x150E
	COLOR = 0x1800
	STENCIL_VALUE_MASK = 0x0B93
	NEVER = 0x0200
	MAX_VIEWPORT_DIMS = 0x0D3A
	TEXTURE_2D = 0x0DE1
	MAX_TEXTURE_SIZE = 0x0D33
	POLYGON_OFFSET_UNITS = 0x2A00
	NICEST = 0x1102
	TEXTURE_BLUE_SIZE = 0x805E
	SCISSOR_BOX = 0x0C10
	FRONT_LEFT = 0x0400
	PACK_SKIP_PIXELS = 0x0D04
	STENCIL_INDEX = 0x1901
	FASTEST = 0x1101
	POLYGON_SMOOTH = 0x0B41
	DEPTH_CLEAR_VALUE = 0x0B73
	RENDERER = 0x1F01
	ALWAYS = 0x0207
	SRC_ALPHA_SATURATE = 0x0308
	FALSE = 0
	EQUAL = 0x0202
	STEREO = 0x0C33
	LINE_SMOOTH_HINT = 0x0C52
	BLEND = 0x0BE2
	DEPTH_WRITEMASK = 0x0B72
	TEXTURE_HEIGHT = 0x1001
	STENCIL_TEST = 0x0B90
	LINEAR_MIPMAP_LINEAR = 0x2703
	DST_COLOR = 0x0306
	NEAREST_MIPMAP_NEAREST = 0x2700
	SRC_COLOR = 0x0300
	XOR = 0x1506
	STENCIL_WRITEMASK = 0x0B98
	POLYGON_OFFSET_POINT = 0x2A01
	OR = 0x1507
	DST_ALPHA = 0x0304
	CULL_FACE = 0x0B44
	LINE_WIDTH_RANGE = 0x0B22
	FLOAT = 0x1406
	NOTEQUAL = 0x0205
	DEPTH_COMPONENT = 0x1902
	ZERO = 0
	COPY = 0x1503
	TEXTURE = 0x1702
	RGBA = 0x1908
	PACK_SKIP_ROWS = 0x0D03
	LINE_LOOP = 0x0002
	RGB = 0x1907
	SUBPIXEL_BITS = 0x0D50
	NO_ERROR = 0
	BLEND_SRC = 0x0BE1
	LINE_WIDTH = 0x0B21
	STENCIL_PASS_DEPTH_FAIL = 0x0B95
	LESS = 0x0201
	TEXTURE_WRAP_T = 0x2803
	PACK_ALIGNMENT = 0x0D05
	TEXTURE_WRAP_S = 0x2802
	UNPACK_ROW_LENGTH = 0x0CF2
	LINE_SMOOTH = 0x0B20
	POLYGON_SMOOTH_HINT = 0x0C53
	LINE_STRIP = 0x0003
	STENCIL_FAIL = 0x0B94
	R3_G3_B2 = 0x2A10
	REPLACE = 0x1E01
	ONE_MINUS_SRC_ALPHA = 0x0303
	COLOR_CLEAR_VALUE = 0x0C22
	RED = 0x1903
	LINE = 0x1B01
	RGB8 = 0x8051
	STENCIL_REF = 0x0B97
	UNPACK_SKIP_ROWS = 0x0CF3
	RGB4 = 0x804F
	VENDOR = 0x1F00
	RGB5 = 0x8050
	DITHER = 0x0BD0
	STENCIL_CLEAR_VALUE = 0x0B91
	INVALID_OPERATION = 0x0502
	EQUIV = 0x1509
	AND_INVERTED = 0x1504
	BACK_RIGHT = 0x0403
	RGB12 = 0x8053
	POINT_SIZE = 0x0B11
	RGB10 = 0x8052
	STENCIL_PASS_DEPTH_PASS = 0x0B96
	RGB16 = 0x8054
	ONE = 1
	STENCIL_BUFFER_BIT = 0x00000400
	STENCIL = 0x1802
	FRONT_RIGHT = 0x0401
	READ_BUFFER = 0x0C02
	BACK = 0x0405
	DEPTH_BUFFER_BIT = 0x00000100
	UNSIGNED_SHORT = 0x1403
	POLYGON_OFFSET_FILL = 0x8037
	UNPACK_LSB_FIRST = 0x0CF1
	ALPHA = 0x1906
	INVALID_VALUE = 0x0501
	SET = 0x150F
	RIGHT = 0x0407
	LINE_WIDTH_GRANULARITY = 0x0B23
)
// VERSION_1_2
const (
	TEXTURE_MAX_LOD = 0x813B
	MAX_ELEMENTS_INDICES = 0x80E9
	TEXTURE_BINDING_3D = 0x806A
	CLAMP_TO_EDGE = 0x812F
	TEXTURE_3D = 0x806F
	PACK_SKIP_IMAGES = 0x806B
	TEXTURE_MAX_LEVEL = 0x813D
	UNSIGNED_BYTE_3_3_2 = 0x8032
	ALIASED_LINE_WIDTH_RANGE = 0x846E
	UNSIGNED_INT_2_10_10_10_REV = 0x8368
	UNPACK_IMAGE_HEIGHT = 0x806E
	UNSIGNED_SHORT_5_6_5_REV = 0x8364
	UNSIGNED_SHORT_1_5_5_5_REV = 0x8366
	UNSIGNED_BYTE_2_3_3_REV = 0x8362
	UNSIGNED_SHORT_4_4_4_4_REV = 0x8365
	UNSIGNED_INT_8_8_8_8 = 0x8035
	UNSIGNED_SHORT_5_6_5 = 0x8363
	TEXTURE_BASE_LEVEL = 0x813C
	SMOOTH_LINE_WIDTH_GRANULARITY = 0x0B23
	TEXTURE_WRAP_R = 0x8072
	UNPACK_SKIP_IMAGES = 0x806D
	PACK_IMAGE_HEIGHT = 0x806C
	TEXTURE_MIN_LOD = 0x813A
	BGR = 0x80E0
	UNSIGNED_INT_10_10_10_2 = 0x8036
	SMOOTH_LINE_WIDTH_RANGE = 0x0B22
	UNSIGNED_INT_8_8_8_8_REV = 0x8367
	UNSIGNED_SHORT_5_5_5_1 = 0x8034
	BGRA = 0x80E1
	PROXY_TEXTURE_3D = 0x8070
	TEXTURE_DEPTH = 0x8071
	SMOOTH_POINT_SIZE_GRANULARITY = 0x0B13
	MAX_3D_TEXTURE_SIZE = 0x8073
	MAX_ELEMENTS_VERTICES = 0x80E8
	SMOOTH_POINT_SIZE_RANGE = 0x0B12
	UNSIGNED_SHORT_4_4_4_4 = 0x8033
)
// VERSION_1_3
const (
	TEXTURE_BINDING_CUBE_MAP = 0x8514
	SAMPLE_ALPHA_TO_ONE = 0x809F
	COMPRESSED_RGBA = 0x84EE
	MAX_CUBE_MAP_TEXTURE_SIZE = 0x851C
	SAMPLE_COVERAGE_INVERT = 0x80AB
	TEXTURE4 = 0x84C4
	TEXTURE5 = 0x84C5
	TEXTURE6 = 0x84C6
	TEXTURE7 = 0x84C7
	TEXTURE0 = 0x84C0
	TEXTURE1 = 0x84C1
	TEXTURE2 = 0x84C2
	TEXTURE3 = 0x84C3
	TEXTURE8 = 0x84C8
	TEXTURE9 = 0x84C9
	SAMPLE_COVERAGE_VALUE = 0x80AA
	TEXTURE_CUBE_MAP_POSITIVE_Z = 0x8519
	TEXTURE_CUBE_MAP_POSITIVE_X = 0x8515
	TEXTURE_CUBE_MAP_POSITIVE_Y = 0x8517
	ACTIVE_TEXTURE = 0x84E0
	NUM_COMPRESSED_TEXTURE_FORMATS = 0x86A2
	TEXTURE_COMPRESSED = 0x86A1
	PROXY_TEXTURE_CUBE_MAP = 0x851B
	SAMPLE_COVERAGE = 0x80A0
	MULTISAMPLE = 0x809D
	COMPRESSED_TEXTURE_FORMATS = 0x86A3
	TEXTURE_COMPRESSED_IMAGE_SIZE = 0x86A0
	TEXTURE31 = 0x84DF
	TEXTURE30 = 0x84DE
	SAMPLE_ALPHA_TO_COVERAGE = 0x809E
	TEXTURE28 = 0x84DC
	TEXTURE29 = 0x84DD
	TEXTURE22 = 0x84D6
	TEXTURE23 = 0x84D7
	TEXTURE20 = 0x84D4
	TEXTURE21 = 0x84D5
	TEXTURE26 = 0x84DA
	TEXTURE27 = 0x84DB
	TEXTURE24 = 0x84D8
	TEXTURE25 = 0x84D9
	SAMPLE_BUFFERS = 0x80A8
	TEXTURE19 = 0x84D3
	COMPRESSED_RGB = 0x84ED
	TEXTURE18 = 0x84D2
	TEXTURE13 = 0x84CD
	TEXTURE12 = 0x84CC
	TEXTURE11 = 0x84CB
	TEXTURE10 = 0x84CA
	TEXTURE17 = 0x84D1
	TEXTURE16 = 0x84D0
	TEXTURE15 = 0x84CF
	TEXTURE14 = 0x84CE
	SAMPLES = 0x80A9
	TEXTURE_CUBE_MAP_NEGATIVE_Z = 0x851A
	TEXTURE_CUBE_MAP_NEGATIVE_X = 0x8516
	TEXTURE_CUBE_MAP_NEGATIVE_Y = 0x8518
	TEXTURE_COMPRESSION_HINT = 0x84EF
	TEXTURE_CUBE_MAP = 0x8513
	CLAMP_TO_BORDER = 0x812D
)
// VERSION_3_2
const (
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS = 0x8C29
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION = 0x8E4C
	TIMEOUT_EXPIRED = 0x911B
	CONTEXT_PROFILE_MASK = 0x9126
	UNSIGNALED = 0x9118
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY = 0x910C
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS = 0x8DE1
	MAX_INTEGER_SAMPLES = 0x9110
	MAX_VERTEX_OUTPUT_COMPONENTS = 0x9122
	FIRST_VERTEX_CONVENTION = 0x8E4D
	SAMPLE_MASK_VALUE = 0x8E52
	SYNC_CONDITION = 0x9113
	MAX_GEOMETRY_INPUT_COMPONENTS = 0x9123
	LINES_ADJACENCY = 0x000A
	SYNC_FLAGS = 0x9115
	MAX_GEOMETRY_OUTPUT_COMPONENTS = 0x9124
	CONDITION_SATISFIED = 0x911C
	TEXTURE_BINDING_2D_MULTISAMPLE = 0x9104
	SIGNALED = 0x9119
	TRIANGLES_ADJACENCY = 0x000C
	MAX_GEOMETRY_UNIFORM_COMPONENTS = 0x8DDF
	GEOMETRY_INPUT_TYPE = 0x8917
	WAIT_FAILED = 0x911D
	SYNC_FENCE = 0x9116
	TRIANGLE_STRIP_ADJACENCY = 0x000D
	SYNC_STATUS = 0x9114
	MAX_SAMPLE_MASK_WORDS = 0x8E59
	TEXTURE_FIXED_SAMPLE_LOCATIONS = 0x9107
	TIMEOUT_IGNORED = 0xFFFFFFFFFFFFFFFF
	TEXTURE_2D_MULTISAMPLE = 0x9100
	MAX_DEPTH_TEXTURE_SAMPLES = 0x910F
	SAMPLE_MASK = 0x8E51
	TEXTURE_SAMPLES = 0x9106
	TEXTURE_CUBE_MAP_SEAMLESS = 0x884F
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY = 0x9103
	OBJECT_TYPE = 0x9112
	TEXTURE_2D_MULTISAMPLE_ARRAY = 0x9102
	SYNC_FLUSH_COMMANDS_BIT = 0x00000001
	SAMPLE_POSITION = 0x8E50
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE = 0x910A
	INT_SAMPLER_2D_MULTISAMPLE = 0x9109
	MAX_GEOMETRY_OUTPUT_VERTICES = 0x8DE0
	MAX_FRAGMENT_INPUT_COMPONENTS = 0x9125
	CONTEXT_CORE_PROFILE_BIT = 0x00000001
	FRAMEBUFFER_ATTACHMENT_LAYERED = 0x8DA7
	PROXY_TEXTURE_2D_MULTISAMPLE = 0x9101
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY = 0x910D
	PROVOKING_VERTEX = 0x8E4F
	DEPTH_CLAMP = 0x864F
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY = 0x9105
	ALREADY_SIGNALED = 0x911A
	GEOMETRY_SHADER = 0x8DD9
	MAX_SERVER_WAIT_TIMEOUT = 0x9111
	MAX_COLOR_TEXTURE_SAMPLES = 0x910E
	LAST_VERTEX_CONVENTION = 0x8E4E
	SAMPLER_2D_MULTISAMPLE = 0x9108
	SAMPLER_2D_MULTISAMPLE_ARRAY = 0x910B
	GEOMETRY_OUTPUT_TYPE = 0x8918
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS = 0x8DA8
	CONTEXT_COMPATIBILITY_PROFILE_BIT = 0x00000002
	PROGRAM_POINT_SIZE = 0x8642
	SYNC_GPU_COMMANDS_COMPLETE = 0x9117
	GEOMETRY_VERTICES_OUT = 0x8916
	LINE_STRIP_ADJACENCY = 0x000B
)
// VERSION_3_3
const (
	TIME_ELAPSED = 0x88BF
	SAMPLER_BINDING = 0x8919
	TIMESTAMP = 0x8E28
	TEXTURE_SWIZZLE_R = 0x8E42
	ANY_SAMPLES_PASSED = 0x8C2F
	TEXTURE_SWIZZLE_B = 0x8E44
	TEXTURE_SWIZZLE_A = 0x8E45
	TEXTURE_SWIZZLE_G = 0x8E43
	VERTEX_ATTRIB_ARRAY_DIVISOR = 0x88FE
	RGB10_A2UI = 0x906F
	INT_2_10_10_10_REV = 0x8D9F
	SRC1_COLOR = 0x88F9
	ONE_MINUS_SRC1_COLOR = 0x88FA
	ONE_MINUS_SRC1_ALPHA = 0x88FB
	TEXTURE_SWIZZLE_RGBA = 0x8E46
	MAX_DUAL_SOURCE_DRAW_BUFFERS = 0x88FC
)
// VERSION_3_0
const (
	MAX_SAMPLES = 0x8D57
	RENDERBUFFER = 0x8D41
	TRANSFORM_FEEDBACK_VARYINGS = 0x8C83
	MAP_FLUSH_EXPLICIT_BIT = 0x0010
	MAX_PROGRAM_TEXEL_OFFSET = 0x8905
	PRIMITIVES_GENERATED = 0x8C87
	VERTEX_ARRAY_BINDING = 0x85B5
	BGR_INTEGER = 0x8D9A
	COLOR_ATTACHMENT8 = 0x8CE8
	COLOR_ATTACHMENT9 = 0x8CE9
	COLOR_ATTACHMENT0 = 0x8CE0
	COLOR_ATTACHMENT1 = 0x8CE1
	COLOR_ATTACHMENT2 = 0x8CE2
	COLOR_ATTACHMENT3 = 0x8CE3
	COLOR_ATTACHMENT4 = 0x8CE4
	COLOR_ATTACHMENT5 = 0x8CE5
	COLOR_ATTACHMENT6 = 0x8CE6
	MAP_WRITE_BIT = 0x0002
	COLOR_ATTACHMENT7 = 0x8CE7
	TEXTURE_BLUE_TYPE = 0x8C12
	RG8I = 0x8237
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS = 0x8C8A
	UNSIGNED_INT_5_9_9_9_REV = 0x8C3E
	COMPRESSED_SIGNED_RED_RGTC1 = 0x8DBC
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT = 0x8CD6
	UNSIGNED_INT_SAMPLER_1D_ARRAY = 0x8DD6
	READ_FRAMEBUFFER = 0x8CA8
	TEXTURE_BINDING_2D_ARRAY = 0x8C1D
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE = 0x8216
	VERTEX_ATTRIB_ARRAY_INTEGER = 0x88FD
	DEPTH32F_STENCIL8 = 0x8CAD
	BGRA_INTEGER = 0x8D9B
	TEXTURE_RED_TYPE = 0x8C10
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = 0x8CD2
	R32UI = 0x8236
	R16I = 0x8233
	R16F = 0x822D
	RGBA8UI = 0x8D7C
	MAX_ARRAY_TEXTURE_LAYERS = 0x88FF
	RGBA_INTEGER = 0x8D99
	RGB32I = 0x8D83
	INT_SAMPLER_CUBE = 0x8DCC
	RGB32F = 0x8815
	RENDERBUFFER_RED_SIZE = 0x8D50
	DRAW_FRAMEBUFFER_BINDING = FRAMEBUFFER_BINDING
	RG16 = 0x822C
	TRANSFORM_FEEDBACK_BUFFER = 0x8C8E
	QUERY_BY_REGION_WAIT = 0x8E15
	SEPARATE_ATTRIBS = 0x8C8D
	RGBA32I = 0x8D82
	TEXTURE_DEPTH_TYPE = 0x8C16
	RGBA32F = 0x8814
	RG16UI = 0x823A
	RGBA8I = 0x8D8E
	SAMPLER_1D_ARRAY_SHADOW = 0x8DC3
	QUERY_NO_WAIT = 0x8E14
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER = 0x8CDB
	CONTEXT_FLAGS = 0x821E
	RENDERBUFFER_WIDTH = 0x8D42
	DEPTH_BUFFER = 0x8223
	TEXTURE_ALPHA_TYPE = 0x8C13
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE = 0x8215
	FRAMEBUFFER_SRGB = 0x8DB9
	RGBA32UI = 0x8D70
	TEXTURE_SHARED_SIZE = 0x8C3F
	COMPRESSED_RG = 0x8226
	COMPRESSED_SIGNED_RG_RGTC2 = 0x8DBE
	UNSIGNED_INT_10F_11F_11F_REV = 0x8C3B
	QUERY_WAIT = 0x8E13
	RENDERBUFFER_DEPTH_SIZE = 0x8D54
	UNSIGNED_INT_SAMPLER_2D_ARRAY = 0x8DD7
	UNSIGNED_NORMALIZED = 0x8C17
	TEXTURE_BINDING_1D_ARRAY = 0x8C1C
	TRANSFORM_FEEDBACK_BUFFER_SIZE = 0x8C85
	DEPTH_COMPONENT32F = 0x8CAC
	COMPRESSED_RED = 0x8225
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT = 0x0001
	FRAMEBUFFER = 0x8D40
	QUERY_BY_REGION_NO_WAIT = 0x8E16
	TEXTURE_2D_ARRAY = 0x8C1A
	MAX_COLOR_ATTACHMENTS = 0x8CDF
	COLOR_ATTACHMENT15 = 0x8CEF
	COLOR_ATTACHMENT14 = 0x8CEE
	COLOR_ATTACHMENT13 = 0x8CED
	COLOR_ATTACHMENT12 = 0x8CEC
	COLOR_ATTACHMENT11 = 0x8CEB
	COLOR_ATTACHMENT10 = 0x8CEA
	RGB32UI = 0x8D71
	RENDERBUFFER_HEIGHT = 0x8D43
	FIXED_ONLY = 0x891D
	RED_INTEGER = 0x8D94
	RENDERBUFFER_BINDING = 0x8CA7
	FRAMEBUFFER_DEFAULT = 0x8218
	MAP_INVALIDATE_RANGE_BIT = 0x0004
	NUM_EXTENSIONS = 0x821D
	READ_FRAMEBUFFER_BINDING = 0x8CAA
	RG32F = 0x8230
	MAP_INVALIDATE_BUFFER_BIT = 0x0008
	FRAMEBUFFER_UNDEFINED = 0x8219
	RG32I = 0x823B
	RGB8I = 0x8D8F
	FLOAT_32_UNSIGNED_INT_24_8_REV = 0x8DAD
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS = 0x8C80
	RG8UI = 0x8238
	INDEX = 0x8222
	RG16F = 0x822F
	RG16I = 0x8239
	UNSIGNED_INT_SAMPLER_CUBE = 0x8DD4
	RG = 0x8227
	R32F = 0x822E
	MAJOR_VERSION = 0x821B
	R32I = 0x8235
	CLAMP_READ_COLOR = 0x891C
	RENDERBUFFER_GREEN_SIZE = 0x8D51
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING = 0x8210
	BLUE_INTEGER = 0x8D96
	RG32UI = 0x823C
	R16 = 0x822A
	MAP_READ_BIT = 0x0001
	RGB16UI = 0x8D77
	TRANSFORM_FEEDBACK_BUFFER_START = 0x8C84
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 0x8CD3
	R11F_G11F_B10F = 0x8C3A
	RENDERBUFFER_STENCIL_SIZE = 0x8D55
	MAP_UNSYNCHRONIZED_BIT = 0x0020
	GREEN_INTEGER = 0x8D95
	DEPTH24_STENCIL8 = 0x88F0
	PROXY_TEXTURE_1D_ARRAY = 0x8C19
	MINOR_VERSION = 0x821C
	INT_SAMPLER_1D_ARRAY = 0x8DCE
	FRAMEBUFFER_COMPLETE = 0x8CD5
	RGB8UI = 0x8D7D
	SAMPLER_2D_ARRAY_SHADOW = 0x8DC4
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE = 0x8CD0
	INT_SAMPLER_2D = 0x8DCA
	UNSIGNED_INT_VEC3 = 0x8DC7
	UNSIGNED_INT_VEC2 = 0x8DC6
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE = 0x8214
	TEXTURE_1D_ARRAY = 0x8C18
	UNSIGNED_INT_VEC4 = 0x8DC8
	R8 = 0x8229
	FRAMEBUFFER_BINDING = 0x8CA6
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS = 0x8C8B
	INT_SAMPLER_1D = 0x8DC9
	FRAMEBUFFER_UNSUPPORTED = 0x8CDD
	RASTERIZER_DISCARD = 0x8C89
	RGBA16UI = 0x8D76
	HALF_FLOAT = 0x140B
	R8I = 0x8231
	R16UI = 0x8234
	RENDERBUFFER_ALPHA_SIZE = 0x8D53
	SAMPLER_CUBE_SHADOW = 0x8DC5
	COMPRESSED_RED_RGTC1 = 0x8DBB
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH = 0x8C76
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE = 0x8213
	MAX_RENDERBUFFER_SIZE = 0x84E8
	UNSIGNED_INT_24_8 = 0x84FA
	DEPTH_STENCIL_ATTACHMENT = 0x821A
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME = 0x8CD1
	COMPRESSED_RG_RGTC2 = 0x8DBD
	RENDERBUFFER_SAMPLES = 0x8CAB
	INT_SAMPLER_3D = 0x8DCB
	MIN_PROGRAM_TEXEL_OFFSET = 0x8904
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = 0x8CD7
	INVALID_FRAMEBUFFER_OPERATION = 0x0506
	UNSIGNED_INT_SAMPLER_2D = 0x8DD2
	STENCIL_BUFFER = 0x8224
	PROXY_TEXTURE_2D_ARRAY = 0x8C1B
	RGB16I = 0x8D89
	RGB16F = 0x881B
	RGB_INTEGER = 0x8D98
	BUFFER_MAP_OFFSET = 0x9121
	RGBA16I = 0x8D88
	STENCIL_INDEX8 = 0x8D48
	DRAW_FRAMEBUFFER = 0x8CA9
	STENCIL_INDEX1 = 0x8D46
	RGBA16F = 0x881A
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = 0x8C88
	STENCIL_INDEX4 = 0x8D47
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE = 0x8211
	BUFFER_ACCESS_FLAGS = 0x911F
	TRANSFORM_FEEDBACK_BUFFER_BINDING = 0x8C8F
	BUFFER_MAP_LENGTH = 0x9120
	CLIP_DISTANCE3 = 0x3003
	CLIP_DISTANCE2 = 0x3002
	MAX_CLIP_DISTANCES = 0x0D32
	CLIP_DISTANCE1 = 0x3001
	CLIP_DISTANCE0 = 0x3000
	CLIP_DISTANCE7 = 0x3007
	CLIP_DISTANCE6 = 0x3006
	CLIP_DISTANCE5 = 0x3005
	CLIP_DISTANCE4 = 0x3004
	INT_SAMPLER_2D_ARRAY = 0x8DCF
	TEXTURE_GREEN_TYPE = 0x8C11
	UNSIGNED_INT_SAMPLER_3D = 0x8DD3
	STENCIL_ATTACHMENT = 0x8D20
	INTERLEAVED_ATTRIBS = 0x8C8C
	TRANSFORM_FEEDBACK_BUFFER_MODE = 0x8C7F
	RENDERBUFFER_INTERNAL_FORMAT = 0x8D44
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE = 0x8D56
	STENCIL_INDEX16 = 0x8D49
	TEXTURE_STENCIL_SIZE = 0x88F1
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE = 0x8217
	DEPTH_STENCIL = 0x84F9
	FRAMEBUFFER_ATTACHMENT_RED_SIZE = 0x8212
	COMPARE_REF_TO_TEXTURE = 0x884E
	RGB9_E5 = 0x8C3D
	RG_INTEGER = 0x8228
	SAMPLER_2D_ARRAY = 0x8DC1
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER = 0x8CDC
	RG8 = 0x822B
	RENDERBUFFER_BLUE_SIZE = 0x8D52
	SAMPLER_1D_ARRAY = 0x8DC0
	UNSIGNED_INT_SAMPLER_1D = 0x8DD1
	DEPTH_ATTACHMENT = 0x8D00
	R8UI = 0x8232
)
// VERSION_3_1
const (
	UNIFORM_BUFFER_SIZE = 0x8A2A
	MAX_VERTEX_UNIFORM_BLOCKS = 0x8A2B
	RG_SNORM = 0x8F91
	UNSIGNED_INT_SAMPLER_2D_RECT = 0x8DD5
	COPY_READ_BUFFER = 0x8F36
	RED_SNORM = 0x8F90
	UNIFORM_BLOCK_INDEX = 0x8A3A
	SIGNED_NORMALIZED = 0x8F9C
	ACTIVE_UNIFORM_BLOCKS = 0x8A36
	TEXTURE_BUFFER = 0x8C2A
	UNIFORM_BLOCK_BINDING = 0x8A3F
	SAMPLER_2D_RECT = 0x8B63
	MAX_TEXTURE_BUFFER_SIZE = 0x8C2B
	RGBA8_SNORM = 0x8F97
	TEXTURE_BUFFER_FORMAT = 0x8C2E
	UNIFORM_OFFSET = 0x8A3B
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS = 0x8A33
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH = 0x8A35
	UNIFORM_SIZE = 0x8A38
	TEXTURE_RECTANGLE = 0x84F5
	MAX_RECTANGLE_TEXTURE_SIZE = 0x84F8
	UNSIGNED_INT_SAMPLER_BUFFER = 0x8DD8
	R16_SNORM = 0x8F98
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER = 0x8A44
	UNIFORM_NAME_LENGTH = 0x8A39
	UNIFORM_BLOCK_NAME_LENGTH = 0x8A41
	UNIFORM_BUFFER_OFFSET_ALIGNMENT = 0x8A34
	RGB_SNORM = 0x8F92
	PRIMITIVE_RESTART_INDEX = 0x8F9E
	RGB16_SNORM = 0x8F9A
	INVALID_INDEX = 0xFFFFFFFF
	UNIFORM_BLOCK_ACTIVE_UNIFORMS = 0x8A42
	RGB8_SNORM = 0x8F96
	UNIFORM_BUFFER_START = 0x8A29
	MAX_COMBINED_UNIFORM_BLOCKS = 0x8A2E
	SAMPLER_BUFFER = 0x8DC2
	SAMPLER_2D_RECT_SHADOW = 0x8B64
	RGBA_SNORM = 0x8F93
	UNIFORM_ARRAY_STRIDE = 0x8A3C
	MAX_UNIFORM_BUFFER_BINDINGS = 0x8A2F
	INT_SAMPLER_BUFFER = 0x8DD0
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES = 0x8A43
	TEXTURE_BINDING_RECTANGLE = 0x84F6
	UNIFORM_MATRIX_STRIDE = 0x8A3D
	INT_SAMPLER_2D_RECT = 0x8DCD
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER = 0x8A46
	MAX_UNIFORM_BLOCK_SIZE = 0x8A30
	UNIFORM_TYPE = 0x8A37
	UNIFORM_BUFFER = 0x8A11
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS = 0x8A31
	RGBA16_SNORM = 0x8F9B
	PRIMITIVE_RESTART = 0x8F9D
	UNIFORM_IS_ROW_MAJOR = 0x8A3E
	UNIFORM_BUFFER_BINDING = 0x8A28
	PROXY_TEXTURE_RECTANGLE = 0x84F7
	RG16_SNORM = 0x8F99
	COPY_WRITE_BUFFER = 0x8F37
	TEXTURE_BUFFER_DATA_STORE_BINDING = 0x8C2D
	UNIFORM_BLOCK_DATA_SIZE = 0x8A40
	MAX_FRAGMENT_UNIFORM_BLOCKS = 0x8A2D
	R8_SNORM = 0x8F94
	RG8_SNORM = 0x8F95
	TEXTURE_BINDING_BUFFER = 0x8C2C
)
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
// VERSION_1_2

// http://www.opengl.org/sdk/docs/man3/xhtml/glBlendColor.xml
func BlendColor(red Clampf, green Clampf, blue Clampf, alpha Clampf)  {
	C.goglBlendColor((C.GLclampf)(red), (C.GLclampf)(green), (C.GLclampf)(blue), (C.GLclampf)(alpha))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glBlendEquation.xml
func BlendEquation(mode Enum)  {
	C.goglBlendEquation((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawRangeElements.xml
func DrawRangeElements(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Pointer)  {
	C.goglDrawRangeElements((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glTexImage3D.xml
func TexImage3D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, depth Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
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
func SampleCoverage(value Clampf, invert Boolean)  {
	C.goglSampleCoverage((C.GLclampf)(value), (C.GLboolean)(invert))
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
func DrawElementsInstancedBaseVertex(mode Enum, count Sizei, type_ Enum, indices Pointer, primcount Sizei, basevertex Int)  {
	C.goglDrawElementsInstancedBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(primcount), (C.GLint)(basevertex))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiDrawElementsBaseVertex.xml
func MultiDrawElementsBaseVertex(mode Enum, count *Sizei, type_ Enum, indices *Pointer, primcount Sizei, basevertex *Int)  {
	C.goglMultiDrawElementsBaseVertex((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(primcount), (*C.GLint)(basevertex))
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
func DrawArraysInstanced(mode Enum, first Int, count Sizei, primcount Sizei)  {
	C.goglDrawArraysInstanced((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(primcount))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glDrawElementsInstanced.xml
func DrawElementsInstanced(mode Enum, count Sizei, type_ Enum, indices Pointer, primcount Sizei)  {
	C.goglDrawElementsInstanced((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(primcount))
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
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib1d.xml
func VertexAttrib1d(index Uint, x Double)  {
	C.goglVertexAttrib1d((C.GLuint)(index), (C.GLdouble)(x))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib1dv.xml
func VertexAttrib1dv(index Uint, v *Double)  {
	C.goglVertexAttrib1dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib1f.xml
func VertexAttrib1f(index Uint, x Float)  {
	C.goglVertexAttrib1f((C.GLuint)(index), (C.GLfloat)(x))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib1fv.xml
func VertexAttrib1fv(index Uint, v *Float)  {
	C.goglVertexAttrib1fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib1s.xml
func VertexAttrib1s(index Uint, x Short)  {
	C.goglVertexAttrib1s((C.GLuint)(index), (C.GLshort)(x))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib1sv.xml
func VertexAttrib1sv(index Uint, v *Short)  {
	C.goglVertexAttrib1sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib2d.xml
func VertexAttrib2d(index Uint, x Double, y Double)  {
	C.goglVertexAttrib2d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib2dv.xml
func VertexAttrib2dv(index Uint, v *Double)  {
	C.goglVertexAttrib2dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib2f.xml
func VertexAttrib2f(index Uint, x Float, y Float)  {
	C.goglVertexAttrib2f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib2fv.xml
func VertexAttrib2fv(index Uint, v *Float)  {
	C.goglVertexAttrib2fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib2s.xml
func VertexAttrib2s(index Uint, x Short, y Short)  {
	C.goglVertexAttrib2s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib2sv.xml
func VertexAttrib2sv(index Uint, v *Short)  {
	C.goglVertexAttrib2sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib3d.xml
func VertexAttrib3d(index Uint, x Double, y Double, z Double)  {
	C.goglVertexAttrib3d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib3dv.xml
func VertexAttrib3dv(index Uint, v *Double)  {
	C.goglVertexAttrib3dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib3f.xml
func VertexAttrib3f(index Uint, x Float, y Float, z Float)  {
	C.goglVertexAttrib3f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib3fv.xml
func VertexAttrib3fv(index Uint, v *Float)  {
	C.goglVertexAttrib3fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib3s.xml
func VertexAttrib3s(index Uint, x Short, y Short, z Short)  {
	C.goglVertexAttrib3s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib3sv.xml
func VertexAttrib3sv(index Uint, v *Short)  {
	C.goglVertexAttrib3sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4Nbv.xml
func VertexAttrib4Nbv(index Uint, v *Byte)  {
	C.goglVertexAttrib4Nbv((C.GLuint)(index), (*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4Niv.xml
func VertexAttrib4Niv(index Uint, v *Int)  {
	C.goglVertexAttrib4Niv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4Nsv.xml
func VertexAttrib4Nsv(index Uint, v *Short)  {
	C.goglVertexAttrib4Nsv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4Nub.xml
func VertexAttrib4Nub(index Uint, x Ubyte, y Ubyte, z Ubyte, w Ubyte)  {
	C.goglVertexAttrib4Nub((C.GLuint)(index), (C.GLubyte)(x), (C.GLubyte)(y), (C.GLubyte)(z), (C.GLubyte)(w))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4Nubv.xml
func VertexAttrib4Nubv(index Uint, v *Ubyte)  {
	C.goglVertexAttrib4Nubv((C.GLuint)(index), (*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4Nuiv.xml
func VertexAttrib4Nuiv(index Uint, v *Uint)  {
	C.goglVertexAttrib4Nuiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4Nusv.xml
func VertexAttrib4Nusv(index Uint, v *Ushort)  {
	C.goglVertexAttrib4Nusv((C.GLuint)(index), (*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4bv.xml
func VertexAttrib4bv(index Uint, v *Byte)  {
	C.goglVertexAttrib4bv((C.GLuint)(index), (*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4d.xml
func VertexAttrib4d(index Uint, x Double, y Double, z Double, w Double)  {
	C.goglVertexAttrib4d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4dv.xml
func VertexAttrib4dv(index Uint, v *Double)  {
	C.goglVertexAttrib4dv((C.GLuint)(index), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4f.xml
func VertexAttrib4f(index Uint, x Float, y Float, z Float, w Float)  {
	C.goglVertexAttrib4f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4fv.xml
func VertexAttrib4fv(index Uint, v *Float)  {
	C.goglVertexAttrib4fv((C.GLuint)(index), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4iv.xml
func VertexAttrib4iv(index Uint, v *Int)  {
	C.goglVertexAttrib4iv((C.GLuint)(index), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4s.xml
func VertexAttrib4s(index Uint, x Short, y Short, z Short, w Short)  {
	C.goglVertexAttrib4s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4sv.xml
func VertexAttrib4sv(index Uint, v *Short)  {
	C.goglVertexAttrib4sv((C.GLuint)(index), (*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4ubv.xml
func VertexAttrib4ubv(index Uint, v *Ubyte)  {
	C.goglVertexAttrib4ubv((C.GLuint)(index), (*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4uiv.xml
func VertexAttrib4uiv(index Uint, v *Uint)  {
	C.goglVertexAttrib4uiv((C.GLuint)(index), (*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttrib4usv.xml
func VertexAttrib4usv(index Uint, v *Ushort)  {
	C.goglVertexAttrib4usv((C.GLuint)(index), (*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribPointer.xml
func VertexAttribPointer(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, pointer Pointer)  {
	C.goglVertexAttribPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// VERSION_1_4

// http://www.opengl.org/sdk/docs/man3/xhtml/glBlendFuncSeparate.xml
func BlendFuncSeparate(sfactorRGB Enum, dfactorRGB Enum, sfactorAlpha Enum, dfactorAlpha Enum)  {
	C.goglBlendFuncSeparate((C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiDrawArrays.xml
func MultiDrawArrays(mode Enum, first *Int, count *Sizei, primcount Sizei)  {
	C.goglMultiDrawArrays((C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(primcount))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glMultiDrawElements.xml
func MultiDrawElements(mode Enum, count *Sizei, type_ Enum, indices *Pointer, primcount Sizei)  {
	C.goglMultiDrawElements((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(primcount))
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
func ClearColor(red Clampf, green Clampf, blue Clampf, alpha Clampf)  {
	C.goglClearColor((C.GLclampf)(red), (C.GLclampf)(green), (C.GLclampf)(blue), (C.GLclampf)(alpha))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClearStencil.xml
func ClearStencil(s Int)  {
	C.goglClearStencil((C.GLint)(s))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glClearDepth.xml
func ClearDepth(depth Clampd)  {
	C.goglClearDepth((C.GLclampd)(depth))
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
func DepthRange(near Clampd, far Clampd)  {
	C.goglDepthRange((C.GLclampd)(near), (C.GLclampd)(far))
}
// http://www.opengl.org/sdk/docs/man3/xhtml/glViewport.xml
func Viewport(x Int, y Int, width Sizei, height Sizei)  {
	C.goglViewport((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
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
func InitVersion21() error {
	var ret C.int
	if ret = C.init_VERSION_2_1(); ret != 0 {
		return errors.New("unable to initialize VERSION_2_1")
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
func Init() error {
	var err error
	if err = InitVersion14(); err != nil {
		return err
	}
	if err = InitVersion15(); err != nil {
		return err
	}
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
	if err = InitVersion32(); err != nil {
		return err
	}
	if err = InitVersion33(); err != nil {
		return err
	}
	if err = InitVersion30(); err != nil {
		return err
	}
	if err = InitVersion31(); err != nil {
		return err
	}
	if err = InitVersion21(); err != nil {
		return err
	}
	if err = InitVersion20(); err != nil {
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

// Allocates a GL string
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
// EOF