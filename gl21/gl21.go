// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// VERSION_2_1
// 
// VERSION_2_0
// 
// VERSION_1_3_DEPRECATED
// 
// VERSION_1_4
// 
// VERSION_1_5
// 
// VERSION_1_0
// 
// VERSION_1_1
// 
// VERSION_1_4_DEPRECATED
// 
// VERSION_1_2
// 
// VERSION_1_3
// 
// VERSION_1_1_DEPRECATED
// 
// VERSION_1_0_DEPRECATED
// 
// VERSION_1_2_DEPRECATED
// 
package gl21

// #cgo darwin  LDFLAGS: -framework OpenGL
// #cgo linux   LDFLAGS: -lGL
// #cgo windows LDFLAGS: -lopengl32
// 
// #ifdef __APPLE__
// // TODO: add context header
// #elif defined(_WIN32) && !defined(APIENTRY) && !defined(__CYGWIN__) && !defined(__SCITECH_SNAP__)
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
// /* OpenGL 3.0 also reuses entry points from these extensions: */
// /* ARB_framebuffer_object */
// /* ARB_map_buffer_range */
// /* ARB_vertex_array_object */
// /* OpenGL 3.1 also reuses entry points from these extensions: */
// /* ARB_copy_buffer */
// /* ARB_uniform_buffer_object */
// /* OpenGL 3.2 also reuses entry points from these extensions: */
// /* ARB_draw_elements_base_vertex */
// /* ARB_provoking_vertex */
// /* ARB_sync */
// /* ARB_texture_multisample */
// /* OpenGL 3.3 also reuses entry points from these extensions: */
// /* ARB_blend_func_extended */
// /* ARB_sampler_objects */
// /* ARB_explicit_attrib_location, but it has none */
// /* ARB_occlusion_query2 (no entry points) */
// /* ARB_shader_bit_encoding (no entry points) */
// /* ARB_texture_rgb10_a2ui (no entry points) */
// /* ARB_texture_swizzle (no entry points) */
// /* ARB_timer_query */
// /* ARB_vertex_type_2_10_10_10_rev */
// /* OpenGL 4.0 also reuses entry points from these extensions: */
// /* ARB_texture_query_lod (no entry points) */
// /* ARB_draw_indirect */
// /* ARB_gpu_shader5 (no entry points) */
// /* ARB_gpu_shader_fp64 */
// /* ARB_shader_subroutine */
// /* ARB_tessellation_shader */
// /* ARB_texture_buffer_object_rgb32 (no entry points) */
// /* ARB_texture_cube_map_array (no entry points) */
// /* ARB_texture_gather (no entry points) */
// /* ARB_transform_feedback2 */
// /* ARB_transform_feedback3 */
// /* OpenGL 4.1 reuses entry points from these extensions: */
// /* ARB_ES2_compatibility */
// /* ARB_get_program_binary */
// /* ARB_separate_shader_objects */
// /* ARB_shader_precision (no entry points) */
// /* ARB_vertex_attrib_64bit */
// /* ARB_viewport_array */
// /* OpenGL 4.2 reuses entry points from these extensions: */
// /* ARB_base_instance */
// /* ARB_shading_language_420pack (no entry points) */
// /* ARB_transform_feedback_instanced */
// /* ARB_compressed_texture_pixel_storage (no entry points) */
// /* ARB_conservative_depth (no entry points) */
// /* ARB_internalformat_query */
// /* ARB_map_buffer_alignment (no entry points) */
// /* ARB_shader_atomic_counters */
// /* ARB_shader_image_load_store */
// /* ARB_shading_language_packing (no entry points) */
// /* ARB_texture_storage */
// /* All ARB_fragment_program entry points are shared with ARB_vertex_program. */
// /* This is really a WGL extension, but defines some associated GL enums.
// * ATI does not export "GL_ATI_pixel_format_float" in the GL_EXTENSIONS string.
// */
// /* Some NV_fragment_program entry points are shared with ARB_vertex_program. */
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
// //  VERSION_1_3_DEPRECATED
// void (APIENTRYP ptrgoglClientActiveTexture)(GLenum texture);
// void (APIENTRYP ptrgoglMultiTexCoord1d)(GLenum target, GLdouble s);
// void (APIENTRYP ptrgoglMultiTexCoord1dv)(GLenum target, GLdouble* v);
// void (APIENTRYP ptrgoglMultiTexCoord1f)(GLenum target, GLfloat s);
// void (APIENTRYP ptrgoglMultiTexCoord1fv)(GLenum target, GLfloat* v);
// void (APIENTRYP ptrgoglMultiTexCoord1i)(GLenum target, GLint s);
// void (APIENTRYP ptrgoglMultiTexCoord1iv)(GLenum target, GLint* v);
// void (APIENTRYP ptrgoglMultiTexCoord1s)(GLenum target, GLshort s);
// void (APIENTRYP ptrgoglMultiTexCoord1sv)(GLenum target, GLshort* v);
// void (APIENTRYP ptrgoglMultiTexCoord2d)(GLenum target, GLdouble s, GLdouble t);
// void (APIENTRYP ptrgoglMultiTexCoord2dv)(GLenum target, GLdouble* v);
// void (APIENTRYP ptrgoglMultiTexCoord2f)(GLenum target, GLfloat s, GLfloat t);
// void (APIENTRYP ptrgoglMultiTexCoord2fv)(GLenum target, GLfloat* v);
// void (APIENTRYP ptrgoglMultiTexCoord2i)(GLenum target, GLint s, GLint t);
// void (APIENTRYP ptrgoglMultiTexCoord2iv)(GLenum target, GLint* v);
// void (APIENTRYP ptrgoglMultiTexCoord2s)(GLenum target, GLshort s, GLshort t);
// void (APIENTRYP ptrgoglMultiTexCoord2sv)(GLenum target, GLshort* v);
// void (APIENTRYP ptrgoglMultiTexCoord3d)(GLenum target, GLdouble s, GLdouble t, GLdouble r);
// void (APIENTRYP ptrgoglMultiTexCoord3dv)(GLenum target, GLdouble* v);
// void (APIENTRYP ptrgoglMultiTexCoord3f)(GLenum target, GLfloat s, GLfloat t, GLfloat r);
// void (APIENTRYP ptrgoglMultiTexCoord3fv)(GLenum target, GLfloat* v);
// void (APIENTRYP ptrgoglMultiTexCoord3i)(GLenum target, GLint s, GLint t, GLint r);
// void (APIENTRYP ptrgoglMultiTexCoord3iv)(GLenum target, GLint* v);
// void (APIENTRYP ptrgoglMultiTexCoord3s)(GLenum target, GLshort s, GLshort t, GLshort r);
// void (APIENTRYP ptrgoglMultiTexCoord3sv)(GLenum target, GLshort* v);
// void (APIENTRYP ptrgoglMultiTexCoord4d)(GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q);
// void (APIENTRYP ptrgoglMultiTexCoord4dv)(GLenum target, GLdouble* v);
// void (APIENTRYP ptrgoglMultiTexCoord4f)(GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q);
// void (APIENTRYP ptrgoglMultiTexCoord4fv)(GLenum target, GLfloat* v);
// void (APIENTRYP ptrgoglMultiTexCoord4i)(GLenum target, GLint s, GLint t, GLint r, GLint q);
// void (APIENTRYP ptrgoglMultiTexCoord4iv)(GLenum target, GLint* v);
// void (APIENTRYP ptrgoglMultiTexCoord4s)(GLenum target, GLshort s, GLshort t, GLshort r, GLshort q);
// void (APIENTRYP ptrgoglMultiTexCoord4sv)(GLenum target, GLshort* v);
// void (APIENTRYP ptrgoglLoadTransposeMatrixf)(GLfloat* m);
// void (APIENTRYP ptrgoglLoadTransposeMatrixd)(GLdouble* m);
// void (APIENTRYP ptrgoglMultTransposeMatrixf)(GLfloat* m);
// void (APIENTRYP ptrgoglMultTransposeMatrixd)(GLdouble* m);
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
// //  VERSION_1_4_DEPRECATED
// void (APIENTRYP ptrgoglFogCoordf)(GLfloat coord);
// void (APIENTRYP ptrgoglFogCoordfv)(GLfloat* coord);
// void (APIENTRYP ptrgoglFogCoordd)(GLdouble coord);
// void (APIENTRYP ptrgoglFogCoorddv)(GLdouble* coord);
// void (APIENTRYP ptrgoglFogCoordPointer)(GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrgoglSecondaryColor3b)(GLbyte red, GLbyte green, GLbyte blue);
// void (APIENTRYP ptrgoglSecondaryColor3bv)(GLbyte* v);
// void (APIENTRYP ptrgoglSecondaryColor3d)(GLdouble red, GLdouble green, GLdouble blue);
// void (APIENTRYP ptrgoglSecondaryColor3dv)(GLdouble* v);
// void (APIENTRYP ptrgoglSecondaryColor3f)(GLfloat red, GLfloat green, GLfloat blue);
// void (APIENTRYP ptrgoglSecondaryColor3fv)(GLfloat* v);
// void (APIENTRYP ptrgoglSecondaryColor3i)(GLint red, GLint green, GLint blue);
// void (APIENTRYP ptrgoglSecondaryColor3iv)(GLint* v);
// void (APIENTRYP ptrgoglSecondaryColor3s)(GLshort red, GLshort green, GLshort blue);
// void (APIENTRYP ptrgoglSecondaryColor3sv)(GLshort* v);
// void (APIENTRYP ptrgoglSecondaryColor3ub)(GLubyte red, GLubyte green, GLubyte blue);
// void (APIENTRYP ptrgoglSecondaryColor3ubv)(GLubyte* v);
// void (APIENTRYP ptrgoglSecondaryColor3ui)(GLuint red, GLuint green, GLuint blue);
// void (APIENTRYP ptrgoglSecondaryColor3uiv)(GLuint* v);
// void (APIENTRYP ptrgoglSecondaryColor3us)(GLushort red, GLushort green, GLushort blue);
// void (APIENTRYP ptrgoglSecondaryColor3usv)(GLushort* v);
// void (APIENTRYP ptrgoglSecondaryColorPointer)(GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrgoglWindowPos2d)(GLdouble x, GLdouble y);
// void (APIENTRYP ptrgoglWindowPos2dv)(GLdouble* v);
// void (APIENTRYP ptrgoglWindowPos2f)(GLfloat x, GLfloat y);
// void (APIENTRYP ptrgoglWindowPos2fv)(GLfloat* v);
// void (APIENTRYP ptrgoglWindowPos2i)(GLint x, GLint y);
// void (APIENTRYP ptrgoglWindowPos2iv)(GLint* v);
// void (APIENTRYP ptrgoglWindowPos2s)(GLshort x, GLshort y);
// void (APIENTRYP ptrgoglWindowPos2sv)(GLshort* v);
// void (APIENTRYP ptrgoglWindowPos3d)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrgoglWindowPos3dv)(GLdouble* v);
// void (APIENTRYP ptrgoglWindowPos3f)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrgoglWindowPos3fv)(GLfloat* v);
// void (APIENTRYP ptrgoglWindowPos3i)(GLint x, GLint y, GLint z);
// void (APIENTRYP ptrgoglWindowPos3iv)(GLint* v);
// void (APIENTRYP ptrgoglWindowPos3s)(GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrgoglWindowPos3sv)(GLshort* v);
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
// //  VERSION_1_1_DEPRECATED
// void (APIENTRYP ptrgoglArrayElement)(GLint i);
// void (APIENTRYP ptrgoglColorPointer)(GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrgoglDisableClientState)(GLenum array);
// void (APIENTRYP ptrgoglEdgeFlagPointer)(GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrgoglEnableClientState)(GLenum array);
// void (APIENTRYP ptrgoglIndexPointer)(GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrgoglInterleavedArrays)(GLenum format, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrgoglNormalPointer)(GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrgoglTexCoordPointer)(GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrgoglVertexPointer)(GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// GLboolean (APIENTRYP ptrgoglAreTexturesResident)(GLsizei n, GLuint* textures, GLboolean* residences);
// void (APIENTRYP ptrgoglPrioritizeTextures)(GLsizei n, GLuint* textures, GLclampf* priorities);
// void (APIENTRYP ptrgoglIndexub)(GLubyte c);
// void (APIENTRYP ptrgoglIndexubv)(GLubyte* c);
// void (APIENTRYP ptrgoglPopClientAttrib)();
// void (APIENTRYP ptrgoglPushClientAttrib)(GLbitfield mask);
// //  VERSION_1_0_DEPRECATED
// void (APIENTRYP ptrgoglNewList)(GLuint list, GLenum mode);
// void (APIENTRYP ptrgoglEndList)();
// void (APIENTRYP ptrgoglCallList)(GLuint list);
// void (APIENTRYP ptrgoglCallLists)(GLsizei n, GLenum type, GLvoid* lists);
// void (APIENTRYP ptrgoglDeleteLists)(GLuint list, GLsizei range);
// GLuint (APIENTRYP ptrgoglGenLists)(GLsizei range);
// void (APIENTRYP ptrgoglListBase)(GLuint base);
// void (APIENTRYP ptrgoglBegin)(GLenum mode);
// void (APIENTRYP ptrgoglBitmap)(GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap);
// void (APIENTRYP ptrgoglColor3b)(GLbyte red, GLbyte green, GLbyte blue);
// void (APIENTRYP ptrgoglColor3bv)(GLbyte* v);
// void (APIENTRYP ptrgoglColor3d)(GLdouble red, GLdouble green, GLdouble blue);
// void (APIENTRYP ptrgoglColor3dv)(GLdouble* v);
// void (APIENTRYP ptrgoglColor3f)(GLfloat red, GLfloat green, GLfloat blue);
// void (APIENTRYP ptrgoglColor3fv)(GLfloat* v);
// void (APIENTRYP ptrgoglColor3i)(GLint red, GLint green, GLint blue);
// void (APIENTRYP ptrgoglColor3iv)(GLint* v);
// void (APIENTRYP ptrgoglColor3s)(GLshort red, GLshort green, GLshort blue);
// void (APIENTRYP ptrgoglColor3sv)(GLshort* v);
// void (APIENTRYP ptrgoglColor3ub)(GLubyte red, GLubyte green, GLubyte blue);
// void (APIENTRYP ptrgoglColor3ubv)(GLubyte* v);
// void (APIENTRYP ptrgoglColor3ui)(GLuint red, GLuint green, GLuint blue);
// void (APIENTRYP ptrgoglColor3uiv)(GLuint* v);
// void (APIENTRYP ptrgoglColor3us)(GLushort red, GLushort green, GLushort blue);
// void (APIENTRYP ptrgoglColor3usv)(GLushort* v);
// void (APIENTRYP ptrgoglColor4b)(GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha);
// void (APIENTRYP ptrgoglColor4bv)(GLbyte* v);
// void (APIENTRYP ptrgoglColor4d)(GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha);
// void (APIENTRYP ptrgoglColor4dv)(GLdouble* v);
// void (APIENTRYP ptrgoglColor4f)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrgoglColor4fv)(GLfloat* v);
// void (APIENTRYP ptrgoglColor4i)(GLint red, GLint green, GLint blue, GLint alpha);
// void (APIENTRYP ptrgoglColor4iv)(GLint* v);
// void (APIENTRYP ptrgoglColor4s)(GLshort red, GLshort green, GLshort blue, GLshort alpha);
// void (APIENTRYP ptrgoglColor4sv)(GLshort* v);
// void (APIENTRYP ptrgoglColor4ub)(GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha);
// void (APIENTRYP ptrgoglColor4ubv)(GLubyte* v);
// void (APIENTRYP ptrgoglColor4ui)(GLuint red, GLuint green, GLuint blue, GLuint alpha);
// void (APIENTRYP ptrgoglColor4uiv)(GLuint* v);
// void (APIENTRYP ptrgoglColor4us)(GLushort red, GLushort green, GLushort blue, GLushort alpha);
// void (APIENTRYP ptrgoglColor4usv)(GLushort* v);
// void (APIENTRYP ptrgoglEdgeFlag)(GLboolean flag);
// void (APIENTRYP ptrgoglEdgeFlagv)(GLboolean* flag);
// void (APIENTRYP ptrgoglEnd)();
// void (APIENTRYP ptrgoglIndexd)(GLdouble c);
// void (APIENTRYP ptrgoglIndexdv)(GLdouble* c);
// void (APIENTRYP ptrgoglIndexf)(GLfloat c);
// void (APIENTRYP ptrgoglIndexfv)(GLfloat* c);
// void (APIENTRYP ptrgoglIndexi)(GLint c);
// void (APIENTRYP ptrgoglIndexiv)(GLint* c);
// void (APIENTRYP ptrgoglIndexs)(GLshort c);
// void (APIENTRYP ptrgoglIndexsv)(GLshort* c);
// void (APIENTRYP ptrgoglNormal3b)(GLbyte nx, GLbyte ny, GLbyte nz);
// void (APIENTRYP ptrgoglNormal3bv)(GLbyte* v);
// void (APIENTRYP ptrgoglNormal3d)(GLdouble nx, GLdouble ny, GLdouble nz);
// void (APIENTRYP ptrgoglNormal3dv)(GLdouble* v);
// void (APIENTRYP ptrgoglNormal3f)(GLfloat nx, GLfloat ny, GLfloat nz);
// void (APIENTRYP ptrgoglNormal3fv)(GLfloat* v);
// void (APIENTRYP ptrgoglNormal3i)(GLint nx, GLint ny, GLint nz);
// void (APIENTRYP ptrgoglNormal3iv)(GLint* v);
// void (APIENTRYP ptrgoglNormal3s)(GLshort nx, GLshort ny, GLshort nz);
// void (APIENTRYP ptrgoglNormal3sv)(GLshort* v);
// void (APIENTRYP ptrgoglRasterPos2d)(GLdouble x, GLdouble y);
// void (APIENTRYP ptrgoglRasterPos2dv)(GLdouble* v);
// void (APIENTRYP ptrgoglRasterPos2f)(GLfloat x, GLfloat y);
// void (APIENTRYP ptrgoglRasterPos2fv)(GLfloat* v);
// void (APIENTRYP ptrgoglRasterPos2i)(GLint x, GLint y);
// void (APIENTRYP ptrgoglRasterPos2iv)(GLint* v);
// void (APIENTRYP ptrgoglRasterPos2s)(GLshort x, GLshort y);
// void (APIENTRYP ptrgoglRasterPos2sv)(GLshort* v);
// void (APIENTRYP ptrgoglRasterPos3d)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrgoglRasterPos3dv)(GLdouble* v);
// void (APIENTRYP ptrgoglRasterPos3f)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrgoglRasterPos3fv)(GLfloat* v);
// void (APIENTRYP ptrgoglRasterPos3i)(GLint x, GLint y, GLint z);
// void (APIENTRYP ptrgoglRasterPos3iv)(GLint* v);
// void (APIENTRYP ptrgoglRasterPos3s)(GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrgoglRasterPos3sv)(GLshort* v);
// void (APIENTRYP ptrgoglRasterPos4d)(GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrgoglRasterPos4dv)(GLdouble* v);
// void (APIENTRYP ptrgoglRasterPos4f)(GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrgoglRasterPos4fv)(GLfloat* v);
// void (APIENTRYP ptrgoglRasterPos4i)(GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrgoglRasterPos4iv)(GLint* v);
// void (APIENTRYP ptrgoglRasterPos4s)(GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrgoglRasterPos4sv)(GLshort* v);
// void (APIENTRYP ptrgoglRectd)(GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2);
// void (APIENTRYP ptrgoglRectdv)(GLdouble* v1, GLdouble* v2);
// void (APIENTRYP ptrgoglRectf)(GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2);
// void (APIENTRYP ptrgoglRectfv)(GLfloat* v1, GLfloat* v2);
// void (APIENTRYP ptrgoglRecti)(GLint x1, GLint y1, GLint x2, GLint y2);
// void (APIENTRYP ptrgoglRectiv)(GLint* v1, GLint* v2);
// void (APIENTRYP ptrgoglRects)(GLshort x1, GLshort y1, GLshort x2, GLshort y2);
// void (APIENTRYP ptrgoglRectsv)(GLshort* v1, GLshort* v2);
// void (APIENTRYP ptrgoglTexCoord1d)(GLdouble s);
// void (APIENTRYP ptrgoglTexCoord1dv)(GLdouble* v);
// void (APIENTRYP ptrgoglTexCoord1f)(GLfloat s);
// void (APIENTRYP ptrgoglTexCoord1fv)(GLfloat* v);
// void (APIENTRYP ptrgoglTexCoord1i)(GLint s);
// void (APIENTRYP ptrgoglTexCoord1iv)(GLint* v);
// void (APIENTRYP ptrgoglTexCoord1s)(GLshort s);
// void (APIENTRYP ptrgoglTexCoord1sv)(GLshort* v);
// void (APIENTRYP ptrgoglTexCoord2d)(GLdouble s, GLdouble t);
// void (APIENTRYP ptrgoglTexCoord2dv)(GLdouble* v);
// void (APIENTRYP ptrgoglTexCoord2f)(GLfloat s, GLfloat t);
// void (APIENTRYP ptrgoglTexCoord2fv)(GLfloat* v);
// void (APIENTRYP ptrgoglTexCoord2i)(GLint s, GLint t);
// void (APIENTRYP ptrgoglTexCoord2iv)(GLint* v);
// void (APIENTRYP ptrgoglTexCoord2s)(GLshort s, GLshort t);
// void (APIENTRYP ptrgoglTexCoord2sv)(GLshort* v);
// void (APIENTRYP ptrgoglTexCoord3d)(GLdouble s, GLdouble t, GLdouble r);
// void (APIENTRYP ptrgoglTexCoord3dv)(GLdouble* v);
// void (APIENTRYP ptrgoglTexCoord3f)(GLfloat s, GLfloat t, GLfloat r);
// void (APIENTRYP ptrgoglTexCoord3fv)(GLfloat* v);
// void (APIENTRYP ptrgoglTexCoord3i)(GLint s, GLint t, GLint r);
// void (APIENTRYP ptrgoglTexCoord3iv)(GLint* v);
// void (APIENTRYP ptrgoglTexCoord3s)(GLshort s, GLshort t, GLshort r);
// void (APIENTRYP ptrgoglTexCoord3sv)(GLshort* v);
// void (APIENTRYP ptrgoglTexCoord4d)(GLdouble s, GLdouble t, GLdouble r, GLdouble q);
// void (APIENTRYP ptrgoglTexCoord4dv)(GLdouble* v);
// void (APIENTRYP ptrgoglTexCoord4f)(GLfloat s, GLfloat t, GLfloat r, GLfloat q);
// void (APIENTRYP ptrgoglTexCoord4fv)(GLfloat* v);
// void (APIENTRYP ptrgoglTexCoord4i)(GLint s, GLint t, GLint r, GLint q);
// void (APIENTRYP ptrgoglTexCoord4iv)(GLint* v);
// void (APIENTRYP ptrgoglTexCoord4s)(GLshort s, GLshort t, GLshort r, GLshort q);
// void (APIENTRYP ptrgoglTexCoord4sv)(GLshort* v);
// void (APIENTRYP ptrgoglVertex2d)(GLdouble x, GLdouble y);
// void (APIENTRYP ptrgoglVertex2dv)(GLdouble* v);
// void (APIENTRYP ptrgoglVertex2f)(GLfloat x, GLfloat y);
// void (APIENTRYP ptrgoglVertex2fv)(GLfloat* v);
// void (APIENTRYP ptrgoglVertex2i)(GLint x, GLint y);
// void (APIENTRYP ptrgoglVertex2iv)(GLint* v);
// void (APIENTRYP ptrgoglVertex2s)(GLshort x, GLshort y);
// void (APIENTRYP ptrgoglVertex2sv)(GLshort* v);
// void (APIENTRYP ptrgoglVertex3d)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrgoglVertex3dv)(GLdouble* v);
// void (APIENTRYP ptrgoglVertex3f)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrgoglVertex3fv)(GLfloat* v);
// void (APIENTRYP ptrgoglVertex3i)(GLint x, GLint y, GLint z);
// void (APIENTRYP ptrgoglVertex3iv)(GLint* v);
// void (APIENTRYP ptrgoglVertex3s)(GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrgoglVertex3sv)(GLshort* v);
// void (APIENTRYP ptrgoglVertex4d)(GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrgoglVertex4dv)(GLdouble* v);
// void (APIENTRYP ptrgoglVertex4f)(GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrgoglVertex4fv)(GLfloat* v);
// void (APIENTRYP ptrgoglVertex4i)(GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrgoglVertex4iv)(GLint* v);
// void (APIENTRYP ptrgoglVertex4s)(GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrgoglVertex4sv)(GLshort* v);
// void (APIENTRYP ptrgoglClipPlane)(GLenum plane, GLdouble* equation);
// void (APIENTRYP ptrgoglColorMaterial)(GLenum face, GLenum mode);
// void (APIENTRYP ptrgoglFogf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglFogfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglFogi)(GLenum pname, GLint param);
// void (APIENTRYP ptrgoglFogiv)(GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglLightf)(GLenum light, GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglLightfv)(GLenum light, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglLighti)(GLenum light, GLenum pname, GLint param);
// void (APIENTRYP ptrgoglLightiv)(GLenum light, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglLightModelf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglLightModelfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglLightModeli)(GLenum pname, GLint param);
// void (APIENTRYP ptrgoglLightModeliv)(GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglLineStipple)(GLint factor, GLushort pattern);
// void (APIENTRYP ptrgoglMaterialf)(GLenum face, GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglMaterialfv)(GLenum face, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglMateriali)(GLenum face, GLenum pname, GLint param);
// void (APIENTRYP ptrgoglMaterialiv)(GLenum face, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglPolygonStipple)(GLubyte* mask);
// void (APIENTRYP ptrgoglShadeModel)(GLenum mode);
// void (APIENTRYP ptrgoglTexEnvf)(GLenum target, GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglTexEnvfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglTexEnvi)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrgoglTexEnviv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglTexGend)(GLenum coord, GLenum pname, GLdouble param);
// void (APIENTRYP ptrgoglTexGendv)(GLenum coord, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrgoglTexGenf)(GLenum coord, GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglTexGenfv)(GLenum coord, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglTexGeni)(GLenum coord, GLenum pname, GLint param);
// void (APIENTRYP ptrgoglTexGeniv)(GLenum coord, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglFeedbackBuffer)(GLsizei size, GLenum type, GLfloat* buffer);
// void (APIENTRYP ptrgoglSelectBuffer)(GLsizei size, GLuint* buffer);
// GLint (APIENTRYP ptrgoglRenderMode)(GLenum mode);
// void (APIENTRYP ptrgoglInitNames)();
// void (APIENTRYP ptrgoglLoadName)(GLuint name);
// void (APIENTRYP ptrgoglPassThrough)(GLfloat token);
// void (APIENTRYP ptrgoglPopName)();
// void (APIENTRYP ptrgoglPushName)(GLuint name);
// void (APIENTRYP ptrgoglClearAccum)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrgoglClearIndex)(GLfloat c);
// void (APIENTRYP ptrgoglIndexMask)(GLuint mask);
// void (APIENTRYP ptrgoglAccum)(GLenum op, GLfloat value);
// void (APIENTRYP ptrgoglPopAttrib)();
// void (APIENTRYP ptrgoglPushAttrib)(GLbitfield mask);
// void (APIENTRYP ptrgoglMap1d)(GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points);
// void (APIENTRYP ptrgoglMap1f)(GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points);
// void (APIENTRYP ptrgoglMap2d)(GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points);
// void (APIENTRYP ptrgoglMap2f)(GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points);
// void (APIENTRYP ptrgoglMapGrid1d)(GLint un, GLdouble u1, GLdouble u2);
// void (APIENTRYP ptrgoglMapGrid1f)(GLint un, GLfloat u1, GLfloat u2);
// void (APIENTRYP ptrgoglMapGrid2d)(GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2);
// void (APIENTRYP ptrgoglMapGrid2f)(GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2);
// void (APIENTRYP ptrgoglEvalCoord1d)(GLdouble u);
// void (APIENTRYP ptrgoglEvalCoord1dv)(GLdouble* u);
// void (APIENTRYP ptrgoglEvalCoord1f)(GLfloat u);
// void (APIENTRYP ptrgoglEvalCoord1fv)(GLfloat* u);
// void (APIENTRYP ptrgoglEvalCoord2d)(GLdouble u, GLdouble v);
// void (APIENTRYP ptrgoglEvalCoord2dv)(GLdouble* u);
// void (APIENTRYP ptrgoglEvalCoord2f)(GLfloat u, GLfloat v);
// void (APIENTRYP ptrgoglEvalCoord2fv)(GLfloat* u);
// void (APIENTRYP ptrgoglEvalMesh1)(GLenum mode, GLint i1, GLint i2);
// void (APIENTRYP ptrgoglEvalPoint1)(GLint i);
// void (APIENTRYP ptrgoglEvalMesh2)(GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2);
// void (APIENTRYP ptrgoglEvalPoint2)(GLint i, GLint j);
// void (APIENTRYP ptrgoglAlphaFunc)(GLenum func, GLclampf ref);
// void (APIENTRYP ptrgoglPixelZoom)(GLfloat xfactor, GLfloat yfactor);
// void (APIENTRYP ptrgoglPixelTransferf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrgoglPixelTransferi)(GLenum pname, GLint param);
// void (APIENTRYP ptrgoglPixelMapfv)(GLenum map, GLsizei mapsize, GLfloat* values);
// void (APIENTRYP ptrgoglPixelMapuiv)(GLenum map, GLsizei mapsize, GLuint* values);
// void (APIENTRYP ptrgoglPixelMapusv)(GLenum map, GLsizei mapsize, GLushort* values);
// void (APIENTRYP ptrgoglCopyPixels)(GLint x, GLint y, GLsizei width, GLsizei height, GLenum type);
// void (APIENTRYP ptrgoglDrawPixels)(GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrgoglGetClipPlane)(GLenum plane, GLdouble* equation);
// void (APIENTRYP ptrgoglGetLightfv)(GLenum light, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetLightiv)(GLenum light, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetMapdv)(GLenum target, GLenum query, GLdouble* v);
// void (APIENTRYP ptrgoglGetMapfv)(GLenum target, GLenum query, GLfloat* v);
// void (APIENTRYP ptrgoglGetMapiv)(GLenum target, GLenum query, GLint* v);
// void (APIENTRYP ptrgoglGetMaterialfv)(GLenum face, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetMaterialiv)(GLenum face, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetPixelMapfv)(GLenum map, GLfloat* values);
// void (APIENTRYP ptrgoglGetPixelMapuiv)(GLenum map, GLuint* values);
// void (APIENTRYP ptrgoglGetPixelMapusv)(GLenum map, GLushort* values);
// void (APIENTRYP ptrgoglGetPolygonStipple)(GLubyte* mask);
// void (APIENTRYP ptrgoglGetTexEnvfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetTexEnviv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetTexGendv)(GLenum coord, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrgoglGetTexGenfv)(GLenum coord, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetTexGeniv)(GLenum coord, GLenum pname, GLint* params);
// GLboolean (APIENTRYP ptrgoglIsList)(GLuint list);
// void (APIENTRYP ptrgoglFrustum)(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar);
// void (APIENTRYP ptrgoglLoadIdentity)();
// void (APIENTRYP ptrgoglLoadMatrixf)(GLfloat* m);
// void (APIENTRYP ptrgoglLoadMatrixd)(GLdouble* m);
// void (APIENTRYP ptrgoglMatrixMode)(GLenum mode);
// void (APIENTRYP ptrgoglMultMatrixf)(GLfloat* m);
// void (APIENTRYP ptrgoglMultMatrixd)(GLdouble* m);
// void (APIENTRYP ptrgoglOrtho)(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar);
// void (APIENTRYP ptrgoglPopMatrix)();
// void (APIENTRYP ptrgoglPushMatrix)();
// void (APIENTRYP ptrgoglRotated)(GLdouble angle, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrgoglRotatef)(GLfloat angle, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrgoglScaled)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrgoglScalef)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrgoglTranslated)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrgoglTranslatef)(GLfloat x, GLfloat y, GLfloat z);
// //  VERSION_1_2_DEPRECATED
// void (APIENTRYP ptrgoglColorTable)(GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* table);
// void (APIENTRYP ptrgoglColorTableParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglColorTableParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglCopyColorTable)(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrgoglGetColorTable)(GLenum target, GLenum format, GLenum type, GLvoid* table);
// void (APIENTRYP ptrgoglGetColorTableParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetColorTableParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglColorSubTable)(GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data);
// void (APIENTRYP ptrgoglCopyColorSubTable)(GLenum target, GLsizei start, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrgoglConvolutionFilter1D)(GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* image);
// void (APIENTRYP ptrgoglConvolutionFilter2D)(GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* image);
// void (APIENTRYP ptrgoglConvolutionParameterf)(GLenum target, GLenum pname, GLfloat params);
// void (APIENTRYP ptrgoglConvolutionParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglConvolutionParameteri)(GLenum target, GLenum pname, GLint params);
// void (APIENTRYP ptrgoglConvolutionParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglCopyConvolutionFilter1D)(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrgoglCopyConvolutionFilter2D)(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrgoglGetConvolutionFilter)(GLenum target, GLenum format, GLenum type, GLvoid* image);
// void (APIENTRYP ptrgoglGetConvolutionParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetConvolutionParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetSeparableFilter)(GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span);
// void (APIENTRYP ptrgoglSeparableFilter2D)(GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column);
// void (APIENTRYP ptrgoglGetHistogram)(GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values);
// void (APIENTRYP ptrgoglGetHistogramParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetHistogramParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglGetMinmax)(GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values);
// void (APIENTRYP ptrgoglGetMinmaxParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrgoglGetMinmaxParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrgoglHistogram)(GLenum target, GLsizei width, GLenum internalformat, GLboolean sink);
// void (APIENTRYP ptrgoglMinmax)(GLenum target, GLenum internalformat, GLboolean sink);
// void (APIENTRYP ptrgoglResetHistogram)(GLenum target);
// void (APIENTRYP ptrgoglResetMinmax)(GLenum target);
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
// //  VERSION_1_3_DEPRECATED
// void goglClientActiveTexture(GLenum texture) {
// 	(*ptrgoglClientActiveTexture)(texture);
// }
// void goglMultiTexCoord1d(GLenum target, GLdouble s) {
// 	(*ptrgoglMultiTexCoord1d)(target, s);
// }
// void goglMultiTexCoord1dv(GLenum target, GLdouble* v) {
// 	(*ptrgoglMultiTexCoord1dv)(target, v);
// }
// void goglMultiTexCoord1f(GLenum target, GLfloat s) {
// 	(*ptrgoglMultiTexCoord1f)(target, s);
// }
// void goglMultiTexCoord1fv(GLenum target, GLfloat* v) {
// 	(*ptrgoglMultiTexCoord1fv)(target, v);
// }
// void goglMultiTexCoord1i(GLenum target, GLint s) {
// 	(*ptrgoglMultiTexCoord1i)(target, s);
// }
// void goglMultiTexCoord1iv(GLenum target, GLint* v) {
// 	(*ptrgoglMultiTexCoord1iv)(target, v);
// }
// void goglMultiTexCoord1s(GLenum target, GLshort s) {
// 	(*ptrgoglMultiTexCoord1s)(target, s);
// }
// void goglMultiTexCoord1sv(GLenum target, GLshort* v) {
// 	(*ptrgoglMultiTexCoord1sv)(target, v);
// }
// void goglMultiTexCoord2d(GLenum target, GLdouble s, GLdouble t) {
// 	(*ptrgoglMultiTexCoord2d)(target, s, t);
// }
// void goglMultiTexCoord2dv(GLenum target, GLdouble* v) {
// 	(*ptrgoglMultiTexCoord2dv)(target, v);
// }
// void goglMultiTexCoord2f(GLenum target, GLfloat s, GLfloat t) {
// 	(*ptrgoglMultiTexCoord2f)(target, s, t);
// }
// void goglMultiTexCoord2fv(GLenum target, GLfloat* v) {
// 	(*ptrgoglMultiTexCoord2fv)(target, v);
// }
// void goglMultiTexCoord2i(GLenum target, GLint s, GLint t) {
// 	(*ptrgoglMultiTexCoord2i)(target, s, t);
// }
// void goglMultiTexCoord2iv(GLenum target, GLint* v) {
// 	(*ptrgoglMultiTexCoord2iv)(target, v);
// }
// void goglMultiTexCoord2s(GLenum target, GLshort s, GLshort t) {
// 	(*ptrgoglMultiTexCoord2s)(target, s, t);
// }
// void goglMultiTexCoord2sv(GLenum target, GLshort* v) {
// 	(*ptrgoglMultiTexCoord2sv)(target, v);
// }
// void goglMultiTexCoord3d(GLenum target, GLdouble s, GLdouble t, GLdouble r) {
// 	(*ptrgoglMultiTexCoord3d)(target, s, t, r);
// }
// void goglMultiTexCoord3dv(GLenum target, GLdouble* v) {
// 	(*ptrgoglMultiTexCoord3dv)(target, v);
// }
// void goglMultiTexCoord3f(GLenum target, GLfloat s, GLfloat t, GLfloat r) {
// 	(*ptrgoglMultiTexCoord3f)(target, s, t, r);
// }
// void goglMultiTexCoord3fv(GLenum target, GLfloat* v) {
// 	(*ptrgoglMultiTexCoord3fv)(target, v);
// }
// void goglMultiTexCoord3i(GLenum target, GLint s, GLint t, GLint r) {
// 	(*ptrgoglMultiTexCoord3i)(target, s, t, r);
// }
// void goglMultiTexCoord3iv(GLenum target, GLint* v) {
// 	(*ptrgoglMultiTexCoord3iv)(target, v);
// }
// void goglMultiTexCoord3s(GLenum target, GLshort s, GLshort t, GLshort r) {
// 	(*ptrgoglMultiTexCoord3s)(target, s, t, r);
// }
// void goglMultiTexCoord3sv(GLenum target, GLshort* v) {
// 	(*ptrgoglMultiTexCoord3sv)(target, v);
// }
// void goglMultiTexCoord4d(GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
// 	(*ptrgoglMultiTexCoord4d)(target, s, t, r, q);
// }
// void goglMultiTexCoord4dv(GLenum target, GLdouble* v) {
// 	(*ptrgoglMultiTexCoord4dv)(target, v);
// }
// void goglMultiTexCoord4f(GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
// 	(*ptrgoglMultiTexCoord4f)(target, s, t, r, q);
// }
// void goglMultiTexCoord4fv(GLenum target, GLfloat* v) {
// 	(*ptrgoglMultiTexCoord4fv)(target, v);
// }
// void goglMultiTexCoord4i(GLenum target, GLint s, GLint t, GLint r, GLint q) {
// 	(*ptrgoglMultiTexCoord4i)(target, s, t, r, q);
// }
// void goglMultiTexCoord4iv(GLenum target, GLint* v) {
// 	(*ptrgoglMultiTexCoord4iv)(target, v);
// }
// void goglMultiTexCoord4s(GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
// 	(*ptrgoglMultiTexCoord4s)(target, s, t, r, q);
// }
// void goglMultiTexCoord4sv(GLenum target, GLshort* v) {
// 	(*ptrgoglMultiTexCoord4sv)(target, v);
// }
// void goglLoadTransposeMatrixf(GLfloat* m) {
// 	(*ptrgoglLoadTransposeMatrixf)(m);
// }
// void goglLoadTransposeMatrixd(GLdouble* m) {
// 	(*ptrgoglLoadTransposeMatrixd)(m);
// }
// void goglMultTransposeMatrixf(GLfloat* m) {
// 	(*ptrgoglMultTransposeMatrixf)(m);
// }
// void goglMultTransposeMatrixd(GLdouble* m) {
// 	(*ptrgoglMultTransposeMatrixd)(m);
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
// //  VERSION_1_4_DEPRECATED
// void goglFogCoordf(GLfloat coord) {
// 	(*ptrgoglFogCoordf)(coord);
// }
// void goglFogCoordfv(GLfloat* coord) {
// 	(*ptrgoglFogCoordfv)(coord);
// }
// void goglFogCoordd(GLdouble coord) {
// 	(*ptrgoglFogCoordd)(coord);
// }
// void goglFogCoorddv(GLdouble* coord) {
// 	(*ptrgoglFogCoorddv)(coord);
// }
// void goglFogCoordPointer(GLenum type, GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglFogCoordPointer)(type, stride, pointer);
// }
// void goglSecondaryColor3b(GLbyte red, GLbyte green, GLbyte blue) {
// 	(*ptrgoglSecondaryColor3b)(red, green, blue);
// }
// void goglSecondaryColor3bv(GLbyte* v) {
// 	(*ptrgoglSecondaryColor3bv)(v);
// }
// void goglSecondaryColor3d(GLdouble red, GLdouble green, GLdouble blue) {
// 	(*ptrgoglSecondaryColor3d)(red, green, blue);
// }
// void goglSecondaryColor3dv(GLdouble* v) {
// 	(*ptrgoglSecondaryColor3dv)(v);
// }
// void goglSecondaryColor3f(GLfloat red, GLfloat green, GLfloat blue) {
// 	(*ptrgoglSecondaryColor3f)(red, green, blue);
// }
// void goglSecondaryColor3fv(GLfloat* v) {
// 	(*ptrgoglSecondaryColor3fv)(v);
// }
// void goglSecondaryColor3i(GLint red, GLint green, GLint blue) {
// 	(*ptrgoglSecondaryColor3i)(red, green, blue);
// }
// void goglSecondaryColor3iv(GLint* v) {
// 	(*ptrgoglSecondaryColor3iv)(v);
// }
// void goglSecondaryColor3s(GLshort red, GLshort green, GLshort blue) {
// 	(*ptrgoglSecondaryColor3s)(red, green, blue);
// }
// void goglSecondaryColor3sv(GLshort* v) {
// 	(*ptrgoglSecondaryColor3sv)(v);
// }
// void goglSecondaryColor3ub(GLubyte red, GLubyte green, GLubyte blue) {
// 	(*ptrgoglSecondaryColor3ub)(red, green, blue);
// }
// void goglSecondaryColor3ubv(GLubyte* v) {
// 	(*ptrgoglSecondaryColor3ubv)(v);
// }
// void goglSecondaryColor3ui(GLuint red, GLuint green, GLuint blue) {
// 	(*ptrgoglSecondaryColor3ui)(red, green, blue);
// }
// void goglSecondaryColor3uiv(GLuint* v) {
// 	(*ptrgoglSecondaryColor3uiv)(v);
// }
// void goglSecondaryColor3us(GLushort red, GLushort green, GLushort blue) {
// 	(*ptrgoglSecondaryColor3us)(red, green, blue);
// }
// void goglSecondaryColor3usv(GLushort* v) {
// 	(*ptrgoglSecondaryColor3usv)(v);
// }
// void goglSecondaryColorPointer(GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglSecondaryColorPointer)(size, type, stride, pointer);
// }
// void goglWindowPos2d(GLdouble x, GLdouble y) {
// 	(*ptrgoglWindowPos2d)(x, y);
// }
// void goglWindowPos2dv(GLdouble* v) {
// 	(*ptrgoglWindowPos2dv)(v);
// }
// void goglWindowPos2f(GLfloat x, GLfloat y) {
// 	(*ptrgoglWindowPos2f)(x, y);
// }
// void goglWindowPos2fv(GLfloat* v) {
// 	(*ptrgoglWindowPos2fv)(v);
// }
// void goglWindowPos2i(GLint x, GLint y) {
// 	(*ptrgoglWindowPos2i)(x, y);
// }
// void goglWindowPos2iv(GLint* v) {
// 	(*ptrgoglWindowPos2iv)(v);
// }
// void goglWindowPos2s(GLshort x, GLshort y) {
// 	(*ptrgoglWindowPos2s)(x, y);
// }
// void goglWindowPos2sv(GLshort* v) {
// 	(*ptrgoglWindowPos2sv)(v);
// }
// void goglWindowPos3d(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrgoglWindowPos3d)(x, y, z);
// }
// void goglWindowPos3dv(GLdouble* v) {
// 	(*ptrgoglWindowPos3dv)(v);
// }
// void goglWindowPos3f(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrgoglWindowPos3f)(x, y, z);
// }
// void goglWindowPos3fv(GLfloat* v) {
// 	(*ptrgoglWindowPos3fv)(v);
// }
// void goglWindowPos3i(GLint x, GLint y, GLint z) {
// 	(*ptrgoglWindowPos3i)(x, y, z);
// }
// void goglWindowPos3iv(GLint* v) {
// 	(*ptrgoglWindowPos3iv)(v);
// }
// void goglWindowPos3s(GLshort x, GLshort y, GLshort z) {
// 	(*ptrgoglWindowPos3s)(x, y, z);
// }
// void goglWindowPos3sv(GLshort* v) {
// 	(*ptrgoglWindowPos3sv)(v);
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
// //  VERSION_1_1_DEPRECATED
// void goglArrayElement(GLint i) {
// 	(*ptrgoglArrayElement)(i);
// }
// void goglColorPointer(GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglColorPointer)(size, type, stride, pointer);
// }
// void goglDisableClientState(GLenum array) {
// 	(*ptrgoglDisableClientState)(array);
// }
// void goglEdgeFlagPointer(GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglEdgeFlagPointer)(stride, pointer);
// }
// void goglEnableClientState(GLenum array) {
// 	(*ptrgoglEnableClientState)(array);
// }
// void goglIndexPointer(GLenum type, GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglIndexPointer)(type, stride, pointer);
// }
// void goglInterleavedArrays(GLenum format, GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglInterleavedArrays)(format, stride, pointer);
// }
// void goglNormalPointer(GLenum type, GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglNormalPointer)(type, stride, pointer);
// }
// void goglTexCoordPointer(GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglTexCoordPointer)(size, type, stride, pointer);
// }
// void goglVertexPointer(GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
// 	(*ptrgoglVertexPointer)(size, type, stride, pointer);
// }
// GLboolean goglAreTexturesResident(GLsizei n, GLuint* textures, GLboolean* residences) {
// 	return (*ptrgoglAreTexturesResident)(n, textures, residences);
// }
// void goglPrioritizeTextures(GLsizei n, GLuint* textures, GLclampf* priorities) {
// 	(*ptrgoglPrioritizeTextures)(n, textures, priorities);
// }
// void goglIndexub(GLubyte c) {
// 	(*ptrgoglIndexub)(c);
// }
// void goglIndexubv(GLubyte* c) {
// 	(*ptrgoglIndexubv)(c);
// }
// void goglPopClientAttrib() {
// 	(*ptrgoglPopClientAttrib)();
// }
// void goglPushClientAttrib(GLbitfield mask) {
// 	(*ptrgoglPushClientAttrib)(mask);
// }
// //  VERSION_1_0_DEPRECATED
// void goglNewList(GLuint list, GLenum mode) {
// 	(*ptrgoglNewList)(list, mode);
// }
// void goglEndList() {
// 	(*ptrgoglEndList)();
// }
// void goglCallList(GLuint list) {
// 	(*ptrgoglCallList)(list);
// }
// void goglCallLists(GLsizei n, GLenum type, GLvoid* lists) {
// 	(*ptrgoglCallLists)(n, type, lists);
// }
// void goglDeleteLists(GLuint list, GLsizei range) {
// 	(*ptrgoglDeleteLists)(list, range);
// }
// GLuint goglGenLists(GLsizei range) {
// 	return (*ptrgoglGenLists)(range);
// }
// void goglListBase(GLuint base) {
// 	(*ptrgoglListBase)(base);
// }
// void goglBegin(GLenum mode) {
// 	(*ptrgoglBegin)(mode);
// }
// void goglBitmap(GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
// 	(*ptrgoglBitmap)(width, height, xorig, yorig, xmove, ymove, bitmap);
// }
// void goglColor3b(GLbyte red, GLbyte green, GLbyte blue) {
// 	(*ptrgoglColor3b)(red, green, blue);
// }
// void goglColor3bv(GLbyte* v) {
// 	(*ptrgoglColor3bv)(v);
// }
// void goglColor3d(GLdouble red, GLdouble green, GLdouble blue) {
// 	(*ptrgoglColor3d)(red, green, blue);
// }
// void goglColor3dv(GLdouble* v) {
// 	(*ptrgoglColor3dv)(v);
// }
// void goglColor3f(GLfloat red, GLfloat green, GLfloat blue) {
// 	(*ptrgoglColor3f)(red, green, blue);
// }
// void goglColor3fv(GLfloat* v) {
// 	(*ptrgoglColor3fv)(v);
// }
// void goglColor3i(GLint red, GLint green, GLint blue) {
// 	(*ptrgoglColor3i)(red, green, blue);
// }
// void goglColor3iv(GLint* v) {
// 	(*ptrgoglColor3iv)(v);
// }
// void goglColor3s(GLshort red, GLshort green, GLshort blue) {
// 	(*ptrgoglColor3s)(red, green, blue);
// }
// void goglColor3sv(GLshort* v) {
// 	(*ptrgoglColor3sv)(v);
// }
// void goglColor3ub(GLubyte red, GLubyte green, GLubyte blue) {
// 	(*ptrgoglColor3ub)(red, green, blue);
// }
// void goglColor3ubv(GLubyte* v) {
// 	(*ptrgoglColor3ubv)(v);
// }
// void goglColor3ui(GLuint red, GLuint green, GLuint blue) {
// 	(*ptrgoglColor3ui)(red, green, blue);
// }
// void goglColor3uiv(GLuint* v) {
// 	(*ptrgoglColor3uiv)(v);
// }
// void goglColor3us(GLushort red, GLushort green, GLushort blue) {
// 	(*ptrgoglColor3us)(red, green, blue);
// }
// void goglColor3usv(GLushort* v) {
// 	(*ptrgoglColor3usv)(v);
// }
// void goglColor4b(GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
// 	(*ptrgoglColor4b)(red, green, blue, alpha);
// }
// void goglColor4bv(GLbyte* v) {
// 	(*ptrgoglColor4bv)(v);
// }
// void goglColor4d(GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
// 	(*ptrgoglColor4d)(red, green, blue, alpha);
// }
// void goglColor4dv(GLdouble* v) {
// 	(*ptrgoglColor4dv)(v);
// }
// void goglColor4f(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
// 	(*ptrgoglColor4f)(red, green, blue, alpha);
// }
// void goglColor4fv(GLfloat* v) {
// 	(*ptrgoglColor4fv)(v);
// }
// void goglColor4i(GLint red, GLint green, GLint blue, GLint alpha) {
// 	(*ptrgoglColor4i)(red, green, blue, alpha);
// }
// void goglColor4iv(GLint* v) {
// 	(*ptrgoglColor4iv)(v);
// }
// void goglColor4s(GLshort red, GLshort green, GLshort blue, GLshort alpha) {
// 	(*ptrgoglColor4s)(red, green, blue, alpha);
// }
// void goglColor4sv(GLshort* v) {
// 	(*ptrgoglColor4sv)(v);
// }
// void goglColor4ub(GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
// 	(*ptrgoglColor4ub)(red, green, blue, alpha);
// }
// void goglColor4ubv(GLubyte* v) {
// 	(*ptrgoglColor4ubv)(v);
// }
// void goglColor4ui(GLuint red, GLuint green, GLuint blue, GLuint alpha) {
// 	(*ptrgoglColor4ui)(red, green, blue, alpha);
// }
// void goglColor4uiv(GLuint* v) {
// 	(*ptrgoglColor4uiv)(v);
// }
// void goglColor4us(GLushort red, GLushort green, GLushort blue, GLushort alpha) {
// 	(*ptrgoglColor4us)(red, green, blue, alpha);
// }
// void goglColor4usv(GLushort* v) {
// 	(*ptrgoglColor4usv)(v);
// }
// void goglEdgeFlag(GLboolean flag) {
// 	(*ptrgoglEdgeFlag)(flag);
// }
// void goglEdgeFlagv(GLboolean* flag) {
// 	(*ptrgoglEdgeFlagv)(flag);
// }
// void goglEnd() {
// 	(*ptrgoglEnd)();
// }
// void goglIndexd(GLdouble c) {
// 	(*ptrgoglIndexd)(c);
// }
// void goglIndexdv(GLdouble* c) {
// 	(*ptrgoglIndexdv)(c);
// }
// void goglIndexf(GLfloat c) {
// 	(*ptrgoglIndexf)(c);
// }
// void goglIndexfv(GLfloat* c) {
// 	(*ptrgoglIndexfv)(c);
// }
// void goglIndexi(GLint c) {
// 	(*ptrgoglIndexi)(c);
// }
// void goglIndexiv(GLint* c) {
// 	(*ptrgoglIndexiv)(c);
// }
// void goglIndexs(GLshort c) {
// 	(*ptrgoglIndexs)(c);
// }
// void goglIndexsv(GLshort* c) {
// 	(*ptrgoglIndexsv)(c);
// }
// void goglNormal3b(GLbyte nx, GLbyte ny, GLbyte nz) {
// 	(*ptrgoglNormal3b)(nx, ny, nz);
// }
// void goglNormal3bv(GLbyte* v) {
// 	(*ptrgoglNormal3bv)(v);
// }
// void goglNormal3d(GLdouble nx, GLdouble ny, GLdouble nz) {
// 	(*ptrgoglNormal3d)(nx, ny, nz);
// }
// void goglNormal3dv(GLdouble* v) {
// 	(*ptrgoglNormal3dv)(v);
// }
// void goglNormal3f(GLfloat nx, GLfloat ny, GLfloat nz) {
// 	(*ptrgoglNormal3f)(nx, ny, nz);
// }
// void goglNormal3fv(GLfloat* v) {
// 	(*ptrgoglNormal3fv)(v);
// }
// void goglNormal3i(GLint nx, GLint ny, GLint nz) {
// 	(*ptrgoglNormal3i)(nx, ny, nz);
// }
// void goglNormal3iv(GLint* v) {
// 	(*ptrgoglNormal3iv)(v);
// }
// void goglNormal3s(GLshort nx, GLshort ny, GLshort nz) {
// 	(*ptrgoglNormal3s)(nx, ny, nz);
// }
// void goglNormal3sv(GLshort* v) {
// 	(*ptrgoglNormal3sv)(v);
// }
// void goglRasterPos2d(GLdouble x, GLdouble y) {
// 	(*ptrgoglRasterPos2d)(x, y);
// }
// void goglRasterPos2dv(GLdouble* v) {
// 	(*ptrgoglRasterPos2dv)(v);
// }
// void goglRasterPos2f(GLfloat x, GLfloat y) {
// 	(*ptrgoglRasterPos2f)(x, y);
// }
// void goglRasterPos2fv(GLfloat* v) {
// 	(*ptrgoglRasterPos2fv)(v);
// }
// void goglRasterPos2i(GLint x, GLint y) {
// 	(*ptrgoglRasterPos2i)(x, y);
// }
// void goglRasterPos2iv(GLint* v) {
// 	(*ptrgoglRasterPos2iv)(v);
// }
// void goglRasterPos2s(GLshort x, GLshort y) {
// 	(*ptrgoglRasterPos2s)(x, y);
// }
// void goglRasterPos2sv(GLshort* v) {
// 	(*ptrgoglRasterPos2sv)(v);
// }
// void goglRasterPos3d(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrgoglRasterPos3d)(x, y, z);
// }
// void goglRasterPos3dv(GLdouble* v) {
// 	(*ptrgoglRasterPos3dv)(v);
// }
// void goglRasterPos3f(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrgoglRasterPos3f)(x, y, z);
// }
// void goglRasterPos3fv(GLfloat* v) {
// 	(*ptrgoglRasterPos3fv)(v);
// }
// void goglRasterPos3i(GLint x, GLint y, GLint z) {
// 	(*ptrgoglRasterPos3i)(x, y, z);
// }
// void goglRasterPos3iv(GLint* v) {
// 	(*ptrgoglRasterPos3iv)(v);
// }
// void goglRasterPos3s(GLshort x, GLshort y, GLshort z) {
// 	(*ptrgoglRasterPos3s)(x, y, z);
// }
// void goglRasterPos3sv(GLshort* v) {
// 	(*ptrgoglRasterPos3sv)(v);
// }
// void goglRasterPos4d(GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrgoglRasterPos4d)(x, y, z, w);
// }
// void goglRasterPos4dv(GLdouble* v) {
// 	(*ptrgoglRasterPos4dv)(v);
// }
// void goglRasterPos4f(GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrgoglRasterPos4f)(x, y, z, w);
// }
// void goglRasterPos4fv(GLfloat* v) {
// 	(*ptrgoglRasterPos4fv)(v);
// }
// void goglRasterPos4i(GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrgoglRasterPos4i)(x, y, z, w);
// }
// void goglRasterPos4iv(GLint* v) {
// 	(*ptrgoglRasterPos4iv)(v);
// }
// void goglRasterPos4s(GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrgoglRasterPos4s)(x, y, z, w);
// }
// void goglRasterPos4sv(GLshort* v) {
// 	(*ptrgoglRasterPos4sv)(v);
// }
// void goglRectd(GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
// 	(*ptrgoglRectd)(x1, y1, x2, y2);
// }
// void goglRectdv(GLdouble* v1, GLdouble* v2) {
// 	(*ptrgoglRectdv)(v1, v2);
// }
// void goglRectf(GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
// 	(*ptrgoglRectf)(x1, y1, x2, y2);
// }
// void goglRectfv(GLfloat* v1, GLfloat* v2) {
// 	(*ptrgoglRectfv)(v1, v2);
// }
// void goglRecti(GLint x1, GLint y1, GLint x2, GLint y2) {
// 	(*ptrgoglRecti)(x1, y1, x2, y2);
// }
// void goglRectiv(GLint* v1, GLint* v2) {
// 	(*ptrgoglRectiv)(v1, v2);
// }
// void goglRects(GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
// 	(*ptrgoglRects)(x1, y1, x2, y2);
// }
// void goglRectsv(GLshort* v1, GLshort* v2) {
// 	(*ptrgoglRectsv)(v1, v2);
// }
// void goglTexCoord1d(GLdouble s) {
// 	(*ptrgoglTexCoord1d)(s);
// }
// void goglTexCoord1dv(GLdouble* v) {
// 	(*ptrgoglTexCoord1dv)(v);
// }
// void goglTexCoord1f(GLfloat s) {
// 	(*ptrgoglTexCoord1f)(s);
// }
// void goglTexCoord1fv(GLfloat* v) {
// 	(*ptrgoglTexCoord1fv)(v);
// }
// void goglTexCoord1i(GLint s) {
// 	(*ptrgoglTexCoord1i)(s);
// }
// void goglTexCoord1iv(GLint* v) {
// 	(*ptrgoglTexCoord1iv)(v);
// }
// void goglTexCoord1s(GLshort s) {
// 	(*ptrgoglTexCoord1s)(s);
// }
// void goglTexCoord1sv(GLshort* v) {
// 	(*ptrgoglTexCoord1sv)(v);
// }
// void goglTexCoord2d(GLdouble s, GLdouble t) {
// 	(*ptrgoglTexCoord2d)(s, t);
// }
// void goglTexCoord2dv(GLdouble* v) {
// 	(*ptrgoglTexCoord2dv)(v);
// }
// void goglTexCoord2f(GLfloat s, GLfloat t) {
// 	(*ptrgoglTexCoord2f)(s, t);
// }
// void goglTexCoord2fv(GLfloat* v) {
// 	(*ptrgoglTexCoord2fv)(v);
// }
// void goglTexCoord2i(GLint s, GLint t) {
// 	(*ptrgoglTexCoord2i)(s, t);
// }
// void goglTexCoord2iv(GLint* v) {
// 	(*ptrgoglTexCoord2iv)(v);
// }
// void goglTexCoord2s(GLshort s, GLshort t) {
// 	(*ptrgoglTexCoord2s)(s, t);
// }
// void goglTexCoord2sv(GLshort* v) {
// 	(*ptrgoglTexCoord2sv)(v);
// }
// void goglTexCoord3d(GLdouble s, GLdouble t, GLdouble r) {
// 	(*ptrgoglTexCoord3d)(s, t, r);
// }
// void goglTexCoord3dv(GLdouble* v) {
// 	(*ptrgoglTexCoord3dv)(v);
// }
// void goglTexCoord3f(GLfloat s, GLfloat t, GLfloat r) {
// 	(*ptrgoglTexCoord3f)(s, t, r);
// }
// void goglTexCoord3fv(GLfloat* v) {
// 	(*ptrgoglTexCoord3fv)(v);
// }
// void goglTexCoord3i(GLint s, GLint t, GLint r) {
// 	(*ptrgoglTexCoord3i)(s, t, r);
// }
// void goglTexCoord3iv(GLint* v) {
// 	(*ptrgoglTexCoord3iv)(v);
// }
// void goglTexCoord3s(GLshort s, GLshort t, GLshort r) {
// 	(*ptrgoglTexCoord3s)(s, t, r);
// }
// void goglTexCoord3sv(GLshort* v) {
// 	(*ptrgoglTexCoord3sv)(v);
// }
// void goglTexCoord4d(GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
// 	(*ptrgoglTexCoord4d)(s, t, r, q);
// }
// void goglTexCoord4dv(GLdouble* v) {
// 	(*ptrgoglTexCoord4dv)(v);
// }
// void goglTexCoord4f(GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
// 	(*ptrgoglTexCoord4f)(s, t, r, q);
// }
// void goglTexCoord4fv(GLfloat* v) {
// 	(*ptrgoglTexCoord4fv)(v);
// }
// void goglTexCoord4i(GLint s, GLint t, GLint r, GLint q) {
// 	(*ptrgoglTexCoord4i)(s, t, r, q);
// }
// void goglTexCoord4iv(GLint* v) {
// 	(*ptrgoglTexCoord4iv)(v);
// }
// void goglTexCoord4s(GLshort s, GLshort t, GLshort r, GLshort q) {
// 	(*ptrgoglTexCoord4s)(s, t, r, q);
// }
// void goglTexCoord4sv(GLshort* v) {
// 	(*ptrgoglTexCoord4sv)(v);
// }
// void goglVertex2d(GLdouble x, GLdouble y) {
// 	(*ptrgoglVertex2d)(x, y);
// }
// void goglVertex2dv(GLdouble* v) {
// 	(*ptrgoglVertex2dv)(v);
// }
// void goglVertex2f(GLfloat x, GLfloat y) {
// 	(*ptrgoglVertex2f)(x, y);
// }
// void goglVertex2fv(GLfloat* v) {
// 	(*ptrgoglVertex2fv)(v);
// }
// void goglVertex2i(GLint x, GLint y) {
// 	(*ptrgoglVertex2i)(x, y);
// }
// void goglVertex2iv(GLint* v) {
// 	(*ptrgoglVertex2iv)(v);
// }
// void goglVertex2s(GLshort x, GLshort y) {
// 	(*ptrgoglVertex2s)(x, y);
// }
// void goglVertex2sv(GLshort* v) {
// 	(*ptrgoglVertex2sv)(v);
// }
// void goglVertex3d(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrgoglVertex3d)(x, y, z);
// }
// void goglVertex3dv(GLdouble* v) {
// 	(*ptrgoglVertex3dv)(v);
// }
// void goglVertex3f(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrgoglVertex3f)(x, y, z);
// }
// void goglVertex3fv(GLfloat* v) {
// 	(*ptrgoglVertex3fv)(v);
// }
// void goglVertex3i(GLint x, GLint y, GLint z) {
// 	(*ptrgoglVertex3i)(x, y, z);
// }
// void goglVertex3iv(GLint* v) {
// 	(*ptrgoglVertex3iv)(v);
// }
// void goglVertex3s(GLshort x, GLshort y, GLshort z) {
// 	(*ptrgoglVertex3s)(x, y, z);
// }
// void goglVertex3sv(GLshort* v) {
// 	(*ptrgoglVertex3sv)(v);
// }
// void goglVertex4d(GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrgoglVertex4d)(x, y, z, w);
// }
// void goglVertex4dv(GLdouble* v) {
// 	(*ptrgoglVertex4dv)(v);
// }
// void goglVertex4f(GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrgoglVertex4f)(x, y, z, w);
// }
// void goglVertex4fv(GLfloat* v) {
// 	(*ptrgoglVertex4fv)(v);
// }
// void goglVertex4i(GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrgoglVertex4i)(x, y, z, w);
// }
// void goglVertex4iv(GLint* v) {
// 	(*ptrgoglVertex4iv)(v);
// }
// void goglVertex4s(GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrgoglVertex4s)(x, y, z, w);
// }
// void goglVertex4sv(GLshort* v) {
// 	(*ptrgoglVertex4sv)(v);
// }
// void goglClipPlane(GLenum plane, GLdouble* equation) {
// 	(*ptrgoglClipPlane)(plane, equation);
// }
// void goglColorMaterial(GLenum face, GLenum mode) {
// 	(*ptrgoglColorMaterial)(face, mode);
// }
// void goglFogf(GLenum pname, GLfloat param) {
// 	(*ptrgoglFogf)(pname, param);
// }
// void goglFogfv(GLenum pname, GLfloat* params) {
// 	(*ptrgoglFogfv)(pname, params);
// }
// void goglFogi(GLenum pname, GLint param) {
// 	(*ptrgoglFogi)(pname, param);
// }
// void goglFogiv(GLenum pname, GLint* params) {
// 	(*ptrgoglFogiv)(pname, params);
// }
// void goglLightf(GLenum light, GLenum pname, GLfloat param) {
// 	(*ptrgoglLightf)(light, pname, param);
// }
// void goglLightfv(GLenum light, GLenum pname, GLfloat* params) {
// 	(*ptrgoglLightfv)(light, pname, params);
// }
// void goglLighti(GLenum light, GLenum pname, GLint param) {
// 	(*ptrgoglLighti)(light, pname, param);
// }
// void goglLightiv(GLenum light, GLenum pname, GLint* params) {
// 	(*ptrgoglLightiv)(light, pname, params);
// }
// void goglLightModelf(GLenum pname, GLfloat param) {
// 	(*ptrgoglLightModelf)(pname, param);
// }
// void goglLightModelfv(GLenum pname, GLfloat* params) {
// 	(*ptrgoglLightModelfv)(pname, params);
// }
// void goglLightModeli(GLenum pname, GLint param) {
// 	(*ptrgoglLightModeli)(pname, param);
// }
// void goglLightModeliv(GLenum pname, GLint* params) {
// 	(*ptrgoglLightModeliv)(pname, params);
// }
// void goglLineStipple(GLint factor, GLushort pattern) {
// 	(*ptrgoglLineStipple)(factor, pattern);
// }
// void goglMaterialf(GLenum face, GLenum pname, GLfloat param) {
// 	(*ptrgoglMaterialf)(face, pname, param);
// }
// void goglMaterialfv(GLenum face, GLenum pname, GLfloat* params) {
// 	(*ptrgoglMaterialfv)(face, pname, params);
// }
// void goglMateriali(GLenum face, GLenum pname, GLint param) {
// 	(*ptrgoglMateriali)(face, pname, param);
// }
// void goglMaterialiv(GLenum face, GLenum pname, GLint* params) {
// 	(*ptrgoglMaterialiv)(face, pname, params);
// }
// void goglPolygonStipple(GLubyte* mask) {
// 	(*ptrgoglPolygonStipple)(mask);
// }
// void goglShadeModel(GLenum mode) {
// 	(*ptrgoglShadeModel)(mode);
// }
// void goglTexEnvf(GLenum target, GLenum pname, GLfloat param) {
// 	(*ptrgoglTexEnvf)(target, pname, param);
// }
// void goglTexEnvfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrgoglTexEnvfv)(target, pname, params);
// }
// void goglTexEnvi(GLenum target, GLenum pname, GLint param) {
// 	(*ptrgoglTexEnvi)(target, pname, param);
// }
// void goglTexEnviv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglTexEnviv)(target, pname, params);
// }
// void goglTexGend(GLenum coord, GLenum pname, GLdouble param) {
// 	(*ptrgoglTexGend)(coord, pname, param);
// }
// void goglTexGendv(GLenum coord, GLenum pname, GLdouble* params) {
// 	(*ptrgoglTexGendv)(coord, pname, params);
// }
// void goglTexGenf(GLenum coord, GLenum pname, GLfloat param) {
// 	(*ptrgoglTexGenf)(coord, pname, param);
// }
// void goglTexGenfv(GLenum coord, GLenum pname, GLfloat* params) {
// 	(*ptrgoglTexGenfv)(coord, pname, params);
// }
// void goglTexGeni(GLenum coord, GLenum pname, GLint param) {
// 	(*ptrgoglTexGeni)(coord, pname, param);
// }
// void goglTexGeniv(GLenum coord, GLenum pname, GLint* params) {
// 	(*ptrgoglTexGeniv)(coord, pname, params);
// }
// void goglFeedbackBuffer(GLsizei size, GLenum type, GLfloat* buffer) {
// 	(*ptrgoglFeedbackBuffer)(size, type, buffer);
// }
// void goglSelectBuffer(GLsizei size, GLuint* buffer) {
// 	(*ptrgoglSelectBuffer)(size, buffer);
// }
// GLint goglRenderMode(GLenum mode) {
// 	return (*ptrgoglRenderMode)(mode);
// }
// void goglInitNames() {
// 	(*ptrgoglInitNames)();
// }
// void goglLoadName(GLuint name) {
// 	(*ptrgoglLoadName)(name);
// }
// void goglPassThrough(GLfloat token) {
// 	(*ptrgoglPassThrough)(token);
// }
// void goglPopName() {
// 	(*ptrgoglPopName)();
// }
// void goglPushName(GLuint name) {
// 	(*ptrgoglPushName)(name);
// }
// void goglClearAccum(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
// 	(*ptrgoglClearAccum)(red, green, blue, alpha);
// }
// void goglClearIndex(GLfloat c) {
// 	(*ptrgoglClearIndex)(c);
// }
// void goglIndexMask(GLuint mask) {
// 	(*ptrgoglIndexMask)(mask);
// }
// void goglAccum(GLenum op, GLfloat value) {
// 	(*ptrgoglAccum)(op, value);
// }
// void goglPopAttrib() {
// 	(*ptrgoglPopAttrib)();
// }
// void goglPushAttrib(GLbitfield mask) {
// 	(*ptrgoglPushAttrib)(mask);
// }
// void goglMap1d(GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
// 	(*ptrgoglMap1d)(target, u1, u2, stride, order, points);
// }
// void goglMap1f(GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
// 	(*ptrgoglMap1f)(target, u1, u2, stride, order, points);
// }
// void goglMap2d(GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
// 	(*ptrgoglMap2d)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// void goglMap2f(GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
// 	(*ptrgoglMap2f)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// void goglMapGrid1d(GLint un, GLdouble u1, GLdouble u2) {
// 	(*ptrgoglMapGrid1d)(un, u1, u2);
// }
// void goglMapGrid1f(GLint un, GLfloat u1, GLfloat u2) {
// 	(*ptrgoglMapGrid1f)(un, u1, u2);
// }
// void goglMapGrid2d(GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
// 	(*ptrgoglMapGrid2d)(un, u1, u2, vn, v1, v2);
// }
// void goglMapGrid2f(GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
// 	(*ptrgoglMapGrid2f)(un, u1, u2, vn, v1, v2);
// }
// void goglEvalCoord1d(GLdouble u) {
// 	(*ptrgoglEvalCoord1d)(u);
// }
// void goglEvalCoord1dv(GLdouble* u) {
// 	(*ptrgoglEvalCoord1dv)(u);
// }
// void goglEvalCoord1f(GLfloat u) {
// 	(*ptrgoglEvalCoord1f)(u);
// }
// void goglEvalCoord1fv(GLfloat* u) {
// 	(*ptrgoglEvalCoord1fv)(u);
// }
// void goglEvalCoord2d(GLdouble u, GLdouble v) {
// 	(*ptrgoglEvalCoord2d)(u, v);
// }
// void goglEvalCoord2dv(GLdouble* u) {
// 	(*ptrgoglEvalCoord2dv)(u);
// }
// void goglEvalCoord2f(GLfloat u, GLfloat v) {
// 	(*ptrgoglEvalCoord2f)(u, v);
// }
// void goglEvalCoord2fv(GLfloat* u) {
// 	(*ptrgoglEvalCoord2fv)(u);
// }
// void goglEvalMesh1(GLenum mode, GLint i1, GLint i2) {
// 	(*ptrgoglEvalMesh1)(mode, i1, i2);
// }
// void goglEvalPoint1(GLint i) {
// 	(*ptrgoglEvalPoint1)(i);
// }
// void goglEvalMesh2(GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
// 	(*ptrgoglEvalMesh2)(mode, i1, i2, j1, j2);
// }
// void goglEvalPoint2(GLint i, GLint j) {
// 	(*ptrgoglEvalPoint2)(i, j);
// }
// void goglAlphaFunc(GLenum func, GLclampf ref) {
// 	(*ptrgoglAlphaFunc)(func, ref);
// }
// void goglPixelZoom(GLfloat xfactor, GLfloat yfactor) {
// 	(*ptrgoglPixelZoom)(xfactor, yfactor);
// }
// void goglPixelTransferf(GLenum pname, GLfloat param) {
// 	(*ptrgoglPixelTransferf)(pname, param);
// }
// void goglPixelTransferi(GLenum pname, GLint param) {
// 	(*ptrgoglPixelTransferi)(pname, param);
// }
// void goglPixelMapfv(GLenum map, GLsizei mapsize, GLfloat* values) {
// 	(*ptrgoglPixelMapfv)(map, mapsize, values);
// }
// void goglPixelMapuiv(GLenum map, GLsizei mapsize, GLuint* values) {
// 	(*ptrgoglPixelMapuiv)(map, mapsize, values);
// }
// void goglPixelMapusv(GLenum map, GLsizei mapsize, GLushort* values) {
// 	(*ptrgoglPixelMapusv)(map, mapsize, values);
// }
// void goglCopyPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
// 	(*ptrgoglCopyPixels)(x, y, width, height, type);
// }
// void goglDrawPixels(GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
// 	(*ptrgoglDrawPixels)(width, height, format, type, pixels);
// }
// void goglGetClipPlane(GLenum plane, GLdouble* equation) {
// 	(*ptrgoglGetClipPlane)(plane, equation);
// }
// void goglGetLightfv(GLenum light, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetLightfv)(light, pname, params);
// }
// void goglGetLightiv(GLenum light, GLenum pname, GLint* params) {
// 	(*ptrgoglGetLightiv)(light, pname, params);
// }
// void goglGetMapdv(GLenum target, GLenum query, GLdouble* v) {
// 	(*ptrgoglGetMapdv)(target, query, v);
// }
// void goglGetMapfv(GLenum target, GLenum query, GLfloat* v) {
// 	(*ptrgoglGetMapfv)(target, query, v);
// }
// void goglGetMapiv(GLenum target, GLenum query, GLint* v) {
// 	(*ptrgoglGetMapiv)(target, query, v);
// }
// void goglGetMaterialfv(GLenum face, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetMaterialfv)(face, pname, params);
// }
// void goglGetMaterialiv(GLenum face, GLenum pname, GLint* params) {
// 	(*ptrgoglGetMaterialiv)(face, pname, params);
// }
// void goglGetPixelMapfv(GLenum map, GLfloat* values) {
// 	(*ptrgoglGetPixelMapfv)(map, values);
// }
// void goglGetPixelMapuiv(GLenum map, GLuint* values) {
// 	(*ptrgoglGetPixelMapuiv)(map, values);
// }
// void goglGetPixelMapusv(GLenum map, GLushort* values) {
// 	(*ptrgoglGetPixelMapusv)(map, values);
// }
// void goglGetPolygonStipple(GLubyte* mask) {
// 	(*ptrgoglGetPolygonStipple)(mask);
// }
// void goglGetTexEnvfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetTexEnvfv)(target, pname, params);
// }
// void goglGetTexEnviv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglGetTexEnviv)(target, pname, params);
// }
// void goglGetTexGendv(GLenum coord, GLenum pname, GLdouble* params) {
// 	(*ptrgoglGetTexGendv)(coord, pname, params);
// }
// void goglGetTexGenfv(GLenum coord, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetTexGenfv)(coord, pname, params);
// }
// void goglGetTexGeniv(GLenum coord, GLenum pname, GLint* params) {
// 	(*ptrgoglGetTexGeniv)(coord, pname, params);
// }
// GLboolean goglIsList(GLuint list) {
// 	return (*ptrgoglIsList)(list);
// }
// void goglFrustum(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
// 	(*ptrgoglFrustum)(left, right, bottom, top, zNear, zFar);
// }
// void goglLoadIdentity() {
// 	(*ptrgoglLoadIdentity)();
// }
// void goglLoadMatrixf(GLfloat* m) {
// 	(*ptrgoglLoadMatrixf)(m);
// }
// void goglLoadMatrixd(GLdouble* m) {
// 	(*ptrgoglLoadMatrixd)(m);
// }
// void goglMatrixMode(GLenum mode) {
// 	(*ptrgoglMatrixMode)(mode);
// }
// void goglMultMatrixf(GLfloat* m) {
// 	(*ptrgoglMultMatrixf)(m);
// }
// void goglMultMatrixd(GLdouble* m) {
// 	(*ptrgoglMultMatrixd)(m);
// }
// void goglOrtho(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
// 	(*ptrgoglOrtho)(left, right, bottom, top, zNear, zFar);
// }
// void goglPopMatrix() {
// 	(*ptrgoglPopMatrix)();
// }
// void goglPushMatrix() {
// 	(*ptrgoglPushMatrix)();
// }
// void goglRotated(GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrgoglRotated)(angle, x, y, z);
// }
// void goglRotatef(GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrgoglRotatef)(angle, x, y, z);
// }
// void goglScaled(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrgoglScaled)(x, y, z);
// }
// void goglScalef(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrgoglScalef)(x, y, z);
// }
// void goglTranslated(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrgoglTranslated)(x, y, z);
// }
// void goglTranslatef(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrgoglTranslatef)(x, y, z);
// }
// //  VERSION_1_2_DEPRECATED
// void goglColorTable(GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* table) {
// 	(*ptrgoglColorTable)(target, internalformat, width, format, type, table);
// }
// void goglColorTableParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrgoglColorTableParameterfv)(target, pname, params);
// }
// void goglColorTableParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglColorTableParameteriv)(target, pname, params);
// }
// void goglCopyColorTable(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
// 	(*ptrgoglCopyColorTable)(target, internalformat, x, y, width);
// }
// void goglGetColorTable(GLenum target, GLenum format, GLenum type, GLvoid* table) {
// 	(*ptrgoglGetColorTable)(target, format, type, table);
// }
// void goglGetColorTableParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetColorTableParameterfv)(target, pname, params);
// }
// void goglGetColorTableParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglGetColorTableParameteriv)(target, pname, params);
// }
// void goglColorSubTable(GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
// 	(*ptrgoglColorSubTable)(target, start, count, format, type, data);
// }
// void goglCopyColorSubTable(GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
// 	(*ptrgoglCopyColorSubTable)(target, start, x, y, width);
// }
// void goglConvolutionFilter1D(GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* image) {
// 	(*ptrgoglConvolutionFilter1D)(target, internalformat, width, format, type, image);
// }
// void goglConvolutionFilter2D(GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* image) {
// 	(*ptrgoglConvolutionFilter2D)(target, internalformat, width, height, format, type, image);
// }
// void goglConvolutionParameterf(GLenum target, GLenum pname, GLfloat params) {
// 	(*ptrgoglConvolutionParameterf)(target, pname, params);
// }
// void goglConvolutionParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrgoglConvolutionParameterfv)(target, pname, params);
// }
// void goglConvolutionParameteri(GLenum target, GLenum pname, GLint params) {
// 	(*ptrgoglConvolutionParameteri)(target, pname, params);
// }
// void goglConvolutionParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglConvolutionParameteriv)(target, pname, params);
// }
// void goglCopyConvolutionFilter1D(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
// 	(*ptrgoglCopyConvolutionFilter1D)(target, internalformat, x, y, width);
// }
// void goglCopyConvolutionFilter2D(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrgoglCopyConvolutionFilter2D)(target, internalformat, x, y, width, height);
// }
// void goglGetConvolutionFilter(GLenum target, GLenum format, GLenum type, GLvoid* image) {
// 	(*ptrgoglGetConvolutionFilter)(target, format, type, image);
// }
// void goglGetConvolutionParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetConvolutionParameterfv)(target, pname, params);
// }
// void goglGetConvolutionParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglGetConvolutionParameteriv)(target, pname, params);
// }
// void goglGetSeparableFilter(GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
// 	(*ptrgoglGetSeparableFilter)(target, format, type, row, column, span);
// }
// void goglSeparableFilter2D(GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
// 	(*ptrgoglSeparableFilter2D)(target, internalformat, width, height, format, type, row, column);
// }
// void goglGetHistogram(GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
// 	(*ptrgoglGetHistogram)(target, reset, format, type, values);
// }
// void goglGetHistogramParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetHistogramParameterfv)(target, pname, params);
// }
// void goglGetHistogramParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglGetHistogramParameteriv)(target, pname, params);
// }
// void goglGetMinmax(GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
// 	(*ptrgoglGetMinmax)(target, reset, format, type, values);
// }
// void goglGetMinmaxParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrgoglGetMinmaxParameterfv)(target, pname, params);
// }
// void goglGetMinmaxParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrgoglGetMinmaxParameteriv)(target, pname, params);
// }
// void goglHistogram(GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
// 	(*ptrgoglHistogram)(target, width, internalformat, sink);
// }
// void goglMinmax(GLenum target, GLenum internalformat, GLboolean sink) {
// 	(*ptrgoglMinmax)(target, internalformat, sink);
// }
// void goglResetHistogram(GLenum target) {
// 	(*ptrgoglResetHistogram)(target);
// }
// void goglResetMinmax(GLenum target) {
// 	(*ptrgoglResetMinmax)(target);
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
// int init_VERSION_1_3_DEPRECATED() {
// 	ptrgoglClientActiveTexture = goglGetProcAddress("glClientActiveTexture");
// 	if(ptrgoglClientActiveTexture == NULL) return 1;
// 	ptrgoglMultiTexCoord1d = goglGetProcAddress("glMultiTexCoord1d");
// 	if(ptrgoglMultiTexCoord1d == NULL) return 1;
// 	ptrgoglMultiTexCoord1dv = goglGetProcAddress("glMultiTexCoord1dv");
// 	if(ptrgoglMultiTexCoord1dv == NULL) return 1;
// 	ptrgoglMultiTexCoord1f = goglGetProcAddress("glMultiTexCoord1f");
// 	if(ptrgoglMultiTexCoord1f == NULL) return 1;
// 	ptrgoglMultiTexCoord1fv = goglGetProcAddress("glMultiTexCoord1fv");
// 	if(ptrgoglMultiTexCoord1fv == NULL) return 1;
// 	ptrgoglMultiTexCoord1i = goglGetProcAddress("glMultiTexCoord1i");
// 	if(ptrgoglMultiTexCoord1i == NULL) return 1;
// 	ptrgoglMultiTexCoord1iv = goglGetProcAddress("glMultiTexCoord1iv");
// 	if(ptrgoglMultiTexCoord1iv == NULL) return 1;
// 	ptrgoglMultiTexCoord1s = goglGetProcAddress("glMultiTexCoord1s");
// 	if(ptrgoglMultiTexCoord1s == NULL) return 1;
// 	ptrgoglMultiTexCoord1sv = goglGetProcAddress("glMultiTexCoord1sv");
// 	if(ptrgoglMultiTexCoord1sv == NULL) return 1;
// 	ptrgoglMultiTexCoord2d = goglGetProcAddress("glMultiTexCoord2d");
// 	if(ptrgoglMultiTexCoord2d == NULL) return 1;
// 	ptrgoglMultiTexCoord2dv = goglGetProcAddress("glMultiTexCoord2dv");
// 	if(ptrgoglMultiTexCoord2dv == NULL) return 1;
// 	ptrgoglMultiTexCoord2f = goglGetProcAddress("glMultiTexCoord2f");
// 	if(ptrgoglMultiTexCoord2f == NULL) return 1;
// 	ptrgoglMultiTexCoord2fv = goglGetProcAddress("glMultiTexCoord2fv");
// 	if(ptrgoglMultiTexCoord2fv == NULL) return 1;
// 	ptrgoglMultiTexCoord2i = goglGetProcAddress("glMultiTexCoord2i");
// 	if(ptrgoglMultiTexCoord2i == NULL) return 1;
// 	ptrgoglMultiTexCoord2iv = goglGetProcAddress("glMultiTexCoord2iv");
// 	if(ptrgoglMultiTexCoord2iv == NULL) return 1;
// 	ptrgoglMultiTexCoord2s = goglGetProcAddress("glMultiTexCoord2s");
// 	if(ptrgoglMultiTexCoord2s == NULL) return 1;
// 	ptrgoglMultiTexCoord2sv = goglGetProcAddress("glMultiTexCoord2sv");
// 	if(ptrgoglMultiTexCoord2sv == NULL) return 1;
// 	ptrgoglMultiTexCoord3d = goglGetProcAddress("glMultiTexCoord3d");
// 	if(ptrgoglMultiTexCoord3d == NULL) return 1;
// 	ptrgoglMultiTexCoord3dv = goglGetProcAddress("glMultiTexCoord3dv");
// 	if(ptrgoglMultiTexCoord3dv == NULL) return 1;
// 	ptrgoglMultiTexCoord3f = goglGetProcAddress("glMultiTexCoord3f");
// 	if(ptrgoglMultiTexCoord3f == NULL) return 1;
// 	ptrgoglMultiTexCoord3fv = goglGetProcAddress("glMultiTexCoord3fv");
// 	if(ptrgoglMultiTexCoord3fv == NULL) return 1;
// 	ptrgoglMultiTexCoord3i = goglGetProcAddress("glMultiTexCoord3i");
// 	if(ptrgoglMultiTexCoord3i == NULL) return 1;
// 	ptrgoglMultiTexCoord3iv = goglGetProcAddress("glMultiTexCoord3iv");
// 	if(ptrgoglMultiTexCoord3iv == NULL) return 1;
// 	ptrgoglMultiTexCoord3s = goglGetProcAddress("glMultiTexCoord3s");
// 	if(ptrgoglMultiTexCoord3s == NULL) return 1;
// 	ptrgoglMultiTexCoord3sv = goglGetProcAddress("glMultiTexCoord3sv");
// 	if(ptrgoglMultiTexCoord3sv == NULL) return 1;
// 	ptrgoglMultiTexCoord4d = goglGetProcAddress("glMultiTexCoord4d");
// 	if(ptrgoglMultiTexCoord4d == NULL) return 1;
// 	ptrgoglMultiTexCoord4dv = goglGetProcAddress("glMultiTexCoord4dv");
// 	if(ptrgoglMultiTexCoord4dv == NULL) return 1;
// 	ptrgoglMultiTexCoord4f = goglGetProcAddress("glMultiTexCoord4f");
// 	if(ptrgoglMultiTexCoord4f == NULL) return 1;
// 	ptrgoglMultiTexCoord4fv = goglGetProcAddress("glMultiTexCoord4fv");
// 	if(ptrgoglMultiTexCoord4fv == NULL) return 1;
// 	ptrgoglMultiTexCoord4i = goglGetProcAddress("glMultiTexCoord4i");
// 	if(ptrgoglMultiTexCoord4i == NULL) return 1;
// 	ptrgoglMultiTexCoord4iv = goglGetProcAddress("glMultiTexCoord4iv");
// 	if(ptrgoglMultiTexCoord4iv == NULL) return 1;
// 	ptrgoglMultiTexCoord4s = goglGetProcAddress("glMultiTexCoord4s");
// 	if(ptrgoglMultiTexCoord4s == NULL) return 1;
// 	ptrgoglMultiTexCoord4sv = goglGetProcAddress("glMultiTexCoord4sv");
// 	if(ptrgoglMultiTexCoord4sv == NULL) return 1;
// 	ptrgoglLoadTransposeMatrixf = goglGetProcAddress("glLoadTransposeMatrixf");
// 	if(ptrgoglLoadTransposeMatrixf == NULL) return 1;
// 	ptrgoglLoadTransposeMatrixd = goglGetProcAddress("glLoadTransposeMatrixd");
// 	if(ptrgoglLoadTransposeMatrixd == NULL) return 1;
// 	ptrgoglMultTransposeMatrixf = goglGetProcAddress("glMultTransposeMatrixf");
// 	if(ptrgoglMultTransposeMatrixf == NULL) return 1;
// 	ptrgoglMultTransposeMatrixd = goglGetProcAddress("glMultTransposeMatrixd");
// 	if(ptrgoglMultTransposeMatrixd == NULL) return 1;
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
// int init_VERSION_1_4_DEPRECATED() {
// 	ptrgoglFogCoordf = goglGetProcAddress("glFogCoordf");
// 	if(ptrgoglFogCoordf == NULL) return 1;
// 	ptrgoglFogCoordfv = goglGetProcAddress("glFogCoordfv");
// 	if(ptrgoglFogCoordfv == NULL) return 1;
// 	ptrgoglFogCoordd = goglGetProcAddress("glFogCoordd");
// 	if(ptrgoglFogCoordd == NULL) return 1;
// 	ptrgoglFogCoorddv = goglGetProcAddress("glFogCoorddv");
// 	if(ptrgoglFogCoorddv == NULL) return 1;
// 	ptrgoglFogCoordPointer = goglGetProcAddress("glFogCoordPointer");
// 	if(ptrgoglFogCoordPointer == NULL) return 1;
// 	ptrgoglSecondaryColor3b = goglGetProcAddress("glSecondaryColor3b");
// 	if(ptrgoglSecondaryColor3b == NULL) return 1;
// 	ptrgoglSecondaryColor3bv = goglGetProcAddress("glSecondaryColor3bv");
// 	if(ptrgoglSecondaryColor3bv == NULL) return 1;
// 	ptrgoglSecondaryColor3d = goglGetProcAddress("glSecondaryColor3d");
// 	if(ptrgoglSecondaryColor3d == NULL) return 1;
// 	ptrgoglSecondaryColor3dv = goglGetProcAddress("glSecondaryColor3dv");
// 	if(ptrgoglSecondaryColor3dv == NULL) return 1;
// 	ptrgoglSecondaryColor3f = goglGetProcAddress("glSecondaryColor3f");
// 	if(ptrgoglSecondaryColor3f == NULL) return 1;
// 	ptrgoglSecondaryColor3fv = goglGetProcAddress("glSecondaryColor3fv");
// 	if(ptrgoglSecondaryColor3fv == NULL) return 1;
// 	ptrgoglSecondaryColor3i = goglGetProcAddress("glSecondaryColor3i");
// 	if(ptrgoglSecondaryColor3i == NULL) return 1;
// 	ptrgoglSecondaryColor3iv = goglGetProcAddress("glSecondaryColor3iv");
// 	if(ptrgoglSecondaryColor3iv == NULL) return 1;
// 	ptrgoglSecondaryColor3s = goglGetProcAddress("glSecondaryColor3s");
// 	if(ptrgoglSecondaryColor3s == NULL) return 1;
// 	ptrgoglSecondaryColor3sv = goglGetProcAddress("glSecondaryColor3sv");
// 	if(ptrgoglSecondaryColor3sv == NULL) return 1;
// 	ptrgoglSecondaryColor3ub = goglGetProcAddress("glSecondaryColor3ub");
// 	if(ptrgoglSecondaryColor3ub == NULL) return 1;
// 	ptrgoglSecondaryColor3ubv = goglGetProcAddress("glSecondaryColor3ubv");
// 	if(ptrgoglSecondaryColor3ubv == NULL) return 1;
// 	ptrgoglSecondaryColor3ui = goglGetProcAddress("glSecondaryColor3ui");
// 	if(ptrgoglSecondaryColor3ui == NULL) return 1;
// 	ptrgoglSecondaryColor3uiv = goglGetProcAddress("glSecondaryColor3uiv");
// 	if(ptrgoglSecondaryColor3uiv == NULL) return 1;
// 	ptrgoglSecondaryColor3us = goglGetProcAddress("glSecondaryColor3us");
// 	if(ptrgoglSecondaryColor3us == NULL) return 1;
// 	ptrgoglSecondaryColor3usv = goglGetProcAddress("glSecondaryColor3usv");
// 	if(ptrgoglSecondaryColor3usv == NULL) return 1;
// 	ptrgoglSecondaryColorPointer = goglGetProcAddress("glSecondaryColorPointer");
// 	if(ptrgoglSecondaryColorPointer == NULL) return 1;
// 	ptrgoglWindowPos2d = goglGetProcAddress("glWindowPos2d");
// 	if(ptrgoglWindowPos2d == NULL) return 1;
// 	ptrgoglWindowPos2dv = goglGetProcAddress("glWindowPos2dv");
// 	if(ptrgoglWindowPos2dv == NULL) return 1;
// 	ptrgoglWindowPos2f = goglGetProcAddress("glWindowPos2f");
// 	if(ptrgoglWindowPos2f == NULL) return 1;
// 	ptrgoglWindowPos2fv = goglGetProcAddress("glWindowPos2fv");
// 	if(ptrgoglWindowPos2fv == NULL) return 1;
// 	ptrgoglWindowPos2i = goglGetProcAddress("glWindowPos2i");
// 	if(ptrgoglWindowPos2i == NULL) return 1;
// 	ptrgoglWindowPos2iv = goglGetProcAddress("glWindowPos2iv");
// 	if(ptrgoglWindowPos2iv == NULL) return 1;
// 	ptrgoglWindowPos2s = goglGetProcAddress("glWindowPos2s");
// 	if(ptrgoglWindowPos2s == NULL) return 1;
// 	ptrgoglWindowPos2sv = goglGetProcAddress("glWindowPos2sv");
// 	if(ptrgoglWindowPos2sv == NULL) return 1;
// 	ptrgoglWindowPos3d = goglGetProcAddress("glWindowPos3d");
// 	if(ptrgoglWindowPos3d == NULL) return 1;
// 	ptrgoglWindowPos3dv = goglGetProcAddress("glWindowPos3dv");
// 	if(ptrgoglWindowPos3dv == NULL) return 1;
// 	ptrgoglWindowPos3f = goglGetProcAddress("glWindowPos3f");
// 	if(ptrgoglWindowPos3f == NULL) return 1;
// 	ptrgoglWindowPos3fv = goglGetProcAddress("glWindowPos3fv");
// 	if(ptrgoglWindowPos3fv == NULL) return 1;
// 	ptrgoglWindowPos3i = goglGetProcAddress("glWindowPos3i");
// 	if(ptrgoglWindowPos3i == NULL) return 1;
// 	ptrgoglWindowPos3iv = goglGetProcAddress("glWindowPos3iv");
// 	if(ptrgoglWindowPos3iv == NULL) return 1;
// 	ptrgoglWindowPos3s = goglGetProcAddress("glWindowPos3s");
// 	if(ptrgoglWindowPos3s == NULL) return 1;
// 	ptrgoglWindowPos3sv = goglGetProcAddress("glWindowPos3sv");
// 	if(ptrgoglWindowPos3sv == NULL) return 1;
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
// int init_VERSION_1_1_DEPRECATED() {
// 	ptrgoglArrayElement = goglGetProcAddress("glArrayElement");
// 	if(ptrgoglArrayElement == NULL) return 1;
// 	ptrgoglColorPointer = goglGetProcAddress("glColorPointer");
// 	if(ptrgoglColorPointer == NULL) return 1;
// 	ptrgoglDisableClientState = goglGetProcAddress("glDisableClientState");
// 	if(ptrgoglDisableClientState == NULL) return 1;
// 	ptrgoglEdgeFlagPointer = goglGetProcAddress("glEdgeFlagPointer");
// 	if(ptrgoglEdgeFlagPointer == NULL) return 1;
// 	ptrgoglEnableClientState = goglGetProcAddress("glEnableClientState");
// 	if(ptrgoglEnableClientState == NULL) return 1;
// 	ptrgoglIndexPointer = goglGetProcAddress("glIndexPointer");
// 	if(ptrgoglIndexPointer == NULL) return 1;
// 	ptrgoglInterleavedArrays = goglGetProcAddress("glInterleavedArrays");
// 	if(ptrgoglInterleavedArrays == NULL) return 1;
// 	ptrgoglNormalPointer = goglGetProcAddress("glNormalPointer");
// 	if(ptrgoglNormalPointer == NULL) return 1;
// 	ptrgoglTexCoordPointer = goglGetProcAddress("glTexCoordPointer");
// 	if(ptrgoglTexCoordPointer == NULL) return 1;
// 	ptrgoglVertexPointer = goglGetProcAddress("glVertexPointer");
// 	if(ptrgoglVertexPointer == NULL) return 1;
// 	ptrgoglAreTexturesResident = goglGetProcAddress("glAreTexturesResident");
// 	if(ptrgoglAreTexturesResident == NULL) return 1;
// 	ptrgoglPrioritizeTextures = goglGetProcAddress("glPrioritizeTextures");
// 	if(ptrgoglPrioritizeTextures == NULL) return 1;
// 	ptrgoglIndexub = goglGetProcAddress("glIndexub");
// 	if(ptrgoglIndexub == NULL) return 1;
// 	ptrgoglIndexubv = goglGetProcAddress("glIndexubv");
// 	if(ptrgoglIndexubv == NULL) return 1;
// 	ptrgoglPopClientAttrib = goglGetProcAddress("glPopClientAttrib");
// 	if(ptrgoglPopClientAttrib == NULL) return 1;
// 	ptrgoglPushClientAttrib = goglGetProcAddress("glPushClientAttrib");
// 	if(ptrgoglPushClientAttrib == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_0_DEPRECATED() {
// 	ptrgoglNewList = goglGetProcAddress("glNewList");
// 	if(ptrgoglNewList == NULL) return 1;
// 	ptrgoglEndList = goglGetProcAddress("glEndList");
// 	if(ptrgoglEndList == NULL) return 1;
// 	ptrgoglCallList = goglGetProcAddress("glCallList");
// 	if(ptrgoglCallList == NULL) return 1;
// 	ptrgoglCallLists = goglGetProcAddress("glCallLists");
// 	if(ptrgoglCallLists == NULL) return 1;
// 	ptrgoglDeleteLists = goglGetProcAddress("glDeleteLists");
// 	if(ptrgoglDeleteLists == NULL) return 1;
// 	ptrgoglGenLists = goglGetProcAddress("glGenLists");
// 	if(ptrgoglGenLists == NULL) return 1;
// 	ptrgoglListBase = goglGetProcAddress("glListBase");
// 	if(ptrgoglListBase == NULL) return 1;
// 	ptrgoglBegin = goglGetProcAddress("glBegin");
// 	if(ptrgoglBegin == NULL) return 1;
// 	ptrgoglBitmap = goglGetProcAddress("glBitmap");
// 	if(ptrgoglBitmap == NULL) return 1;
// 	ptrgoglColor3b = goglGetProcAddress("glColor3b");
// 	if(ptrgoglColor3b == NULL) return 1;
// 	ptrgoglColor3bv = goglGetProcAddress("glColor3bv");
// 	if(ptrgoglColor3bv == NULL) return 1;
// 	ptrgoglColor3d = goglGetProcAddress("glColor3d");
// 	if(ptrgoglColor3d == NULL) return 1;
// 	ptrgoglColor3dv = goglGetProcAddress("glColor3dv");
// 	if(ptrgoglColor3dv == NULL) return 1;
// 	ptrgoglColor3f = goglGetProcAddress("glColor3f");
// 	if(ptrgoglColor3f == NULL) return 1;
// 	ptrgoglColor3fv = goglGetProcAddress("glColor3fv");
// 	if(ptrgoglColor3fv == NULL) return 1;
// 	ptrgoglColor3i = goglGetProcAddress("glColor3i");
// 	if(ptrgoglColor3i == NULL) return 1;
// 	ptrgoglColor3iv = goglGetProcAddress("glColor3iv");
// 	if(ptrgoglColor3iv == NULL) return 1;
// 	ptrgoglColor3s = goglGetProcAddress("glColor3s");
// 	if(ptrgoglColor3s == NULL) return 1;
// 	ptrgoglColor3sv = goglGetProcAddress("glColor3sv");
// 	if(ptrgoglColor3sv == NULL) return 1;
// 	ptrgoglColor3ub = goglGetProcAddress("glColor3ub");
// 	if(ptrgoglColor3ub == NULL) return 1;
// 	ptrgoglColor3ubv = goglGetProcAddress("glColor3ubv");
// 	if(ptrgoglColor3ubv == NULL) return 1;
// 	ptrgoglColor3ui = goglGetProcAddress("glColor3ui");
// 	if(ptrgoglColor3ui == NULL) return 1;
// 	ptrgoglColor3uiv = goglGetProcAddress("glColor3uiv");
// 	if(ptrgoglColor3uiv == NULL) return 1;
// 	ptrgoglColor3us = goglGetProcAddress("glColor3us");
// 	if(ptrgoglColor3us == NULL) return 1;
// 	ptrgoglColor3usv = goglGetProcAddress("glColor3usv");
// 	if(ptrgoglColor3usv == NULL) return 1;
// 	ptrgoglColor4b = goglGetProcAddress("glColor4b");
// 	if(ptrgoglColor4b == NULL) return 1;
// 	ptrgoglColor4bv = goglGetProcAddress("glColor4bv");
// 	if(ptrgoglColor4bv == NULL) return 1;
// 	ptrgoglColor4d = goglGetProcAddress("glColor4d");
// 	if(ptrgoglColor4d == NULL) return 1;
// 	ptrgoglColor4dv = goglGetProcAddress("glColor4dv");
// 	if(ptrgoglColor4dv == NULL) return 1;
// 	ptrgoglColor4f = goglGetProcAddress("glColor4f");
// 	if(ptrgoglColor4f == NULL) return 1;
// 	ptrgoglColor4fv = goglGetProcAddress("glColor4fv");
// 	if(ptrgoglColor4fv == NULL) return 1;
// 	ptrgoglColor4i = goglGetProcAddress("glColor4i");
// 	if(ptrgoglColor4i == NULL) return 1;
// 	ptrgoglColor4iv = goglGetProcAddress("glColor4iv");
// 	if(ptrgoglColor4iv == NULL) return 1;
// 	ptrgoglColor4s = goglGetProcAddress("glColor4s");
// 	if(ptrgoglColor4s == NULL) return 1;
// 	ptrgoglColor4sv = goglGetProcAddress("glColor4sv");
// 	if(ptrgoglColor4sv == NULL) return 1;
// 	ptrgoglColor4ub = goglGetProcAddress("glColor4ub");
// 	if(ptrgoglColor4ub == NULL) return 1;
// 	ptrgoglColor4ubv = goglGetProcAddress("glColor4ubv");
// 	if(ptrgoglColor4ubv == NULL) return 1;
// 	ptrgoglColor4ui = goglGetProcAddress("glColor4ui");
// 	if(ptrgoglColor4ui == NULL) return 1;
// 	ptrgoglColor4uiv = goglGetProcAddress("glColor4uiv");
// 	if(ptrgoglColor4uiv == NULL) return 1;
// 	ptrgoglColor4us = goglGetProcAddress("glColor4us");
// 	if(ptrgoglColor4us == NULL) return 1;
// 	ptrgoglColor4usv = goglGetProcAddress("glColor4usv");
// 	if(ptrgoglColor4usv == NULL) return 1;
// 	ptrgoglEdgeFlag = goglGetProcAddress("glEdgeFlag");
// 	if(ptrgoglEdgeFlag == NULL) return 1;
// 	ptrgoglEdgeFlagv = goglGetProcAddress("glEdgeFlagv");
// 	if(ptrgoglEdgeFlagv == NULL) return 1;
// 	ptrgoglEnd = goglGetProcAddress("glEnd");
// 	if(ptrgoglEnd == NULL) return 1;
// 	ptrgoglIndexd = goglGetProcAddress("glIndexd");
// 	if(ptrgoglIndexd == NULL) return 1;
// 	ptrgoglIndexdv = goglGetProcAddress("glIndexdv");
// 	if(ptrgoglIndexdv == NULL) return 1;
// 	ptrgoglIndexf = goglGetProcAddress("glIndexf");
// 	if(ptrgoglIndexf == NULL) return 1;
// 	ptrgoglIndexfv = goglGetProcAddress("glIndexfv");
// 	if(ptrgoglIndexfv == NULL) return 1;
// 	ptrgoglIndexi = goglGetProcAddress("glIndexi");
// 	if(ptrgoglIndexi == NULL) return 1;
// 	ptrgoglIndexiv = goglGetProcAddress("glIndexiv");
// 	if(ptrgoglIndexiv == NULL) return 1;
// 	ptrgoglIndexs = goglGetProcAddress("glIndexs");
// 	if(ptrgoglIndexs == NULL) return 1;
// 	ptrgoglIndexsv = goglGetProcAddress("glIndexsv");
// 	if(ptrgoglIndexsv == NULL) return 1;
// 	ptrgoglNormal3b = goglGetProcAddress("glNormal3b");
// 	if(ptrgoglNormal3b == NULL) return 1;
// 	ptrgoglNormal3bv = goglGetProcAddress("glNormal3bv");
// 	if(ptrgoglNormal3bv == NULL) return 1;
// 	ptrgoglNormal3d = goglGetProcAddress("glNormal3d");
// 	if(ptrgoglNormal3d == NULL) return 1;
// 	ptrgoglNormal3dv = goglGetProcAddress("glNormal3dv");
// 	if(ptrgoglNormal3dv == NULL) return 1;
// 	ptrgoglNormal3f = goglGetProcAddress("glNormal3f");
// 	if(ptrgoglNormal3f == NULL) return 1;
// 	ptrgoglNormal3fv = goglGetProcAddress("glNormal3fv");
// 	if(ptrgoglNormal3fv == NULL) return 1;
// 	ptrgoglNormal3i = goglGetProcAddress("glNormal3i");
// 	if(ptrgoglNormal3i == NULL) return 1;
// 	ptrgoglNormal3iv = goglGetProcAddress("glNormal3iv");
// 	if(ptrgoglNormal3iv == NULL) return 1;
// 	ptrgoglNormal3s = goglGetProcAddress("glNormal3s");
// 	if(ptrgoglNormal3s == NULL) return 1;
// 	ptrgoglNormal3sv = goglGetProcAddress("glNormal3sv");
// 	if(ptrgoglNormal3sv == NULL) return 1;
// 	ptrgoglRasterPos2d = goglGetProcAddress("glRasterPos2d");
// 	if(ptrgoglRasterPos2d == NULL) return 1;
// 	ptrgoglRasterPos2dv = goglGetProcAddress("glRasterPos2dv");
// 	if(ptrgoglRasterPos2dv == NULL) return 1;
// 	ptrgoglRasterPos2f = goglGetProcAddress("glRasterPos2f");
// 	if(ptrgoglRasterPos2f == NULL) return 1;
// 	ptrgoglRasterPos2fv = goglGetProcAddress("glRasterPos2fv");
// 	if(ptrgoglRasterPos2fv == NULL) return 1;
// 	ptrgoglRasterPos2i = goglGetProcAddress("glRasterPos2i");
// 	if(ptrgoglRasterPos2i == NULL) return 1;
// 	ptrgoglRasterPos2iv = goglGetProcAddress("glRasterPos2iv");
// 	if(ptrgoglRasterPos2iv == NULL) return 1;
// 	ptrgoglRasterPos2s = goglGetProcAddress("glRasterPos2s");
// 	if(ptrgoglRasterPos2s == NULL) return 1;
// 	ptrgoglRasterPos2sv = goglGetProcAddress("glRasterPos2sv");
// 	if(ptrgoglRasterPos2sv == NULL) return 1;
// 	ptrgoglRasterPos3d = goglGetProcAddress("glRasterPos3d");
// 	if(ptrgoglRasterPos3d == NULL) return 1;
// 	ptrgoglRasterPos3dv = goglGetProcAddress("glRasterPos3dv");
// 	if(ptrgoglRasterPos3dv == NULL) return 1;
// 	ptrgoglRasterPos3f = goglGetProcAddress("glRasterPos3f");
// 	if(ptrgoglRasterPos3f == NULL) return 1;
// 	ptrgoglRasterPos3fv = goglGetProcAddress("glRasterPos3fv");
// 	if(ptrgoglRasterPos3fv == NULL) return 1;
// 	ptrgoglRasterPos3i = goglGetProcAddress("glRasterPos3i");
// 	if(ptrgoglRasterPos3i == NULL) return 1;
// 	ptrgoglRasterPos3iv = goglGetProcAddress("glRasterPos3iv");
// 	if(ptrgoglRasterPos3iv == NULL) return 1;
// 	ptrgoglRasterPos3s = goglGetProcAddress("glRasterPos3s");
// 	if(ptrgoglRasterPos3s == NULL) return 1;
// 	ptrgoglRasterPos3sv = goglGetProcAddress("glRasterPos3sv");
// 	if(ptrgoglRasterPos3sv == NULL) return 1;
// 	ptrgoglRasterPos4d = goglGetProcAddress("glRasterPos4d");
// 	if(ptrgoglRasterPos4d == NULL) return 1;
// 	ptrgoglRasterPos4dv = goglGetProcAddress("glRasterPos4dv");
// 	if(ptrgoglRasterPos4dv == NULL) return 1;
// 	ptrgoglRasterPos4f = goglGetProcAddress("glRasterPos4f");
// 	if(ptrgoglRasterPos4f == NULL) return 1;
// 	ptrgoglRasterPos4fv = goglGetProcAddress("glRasterPos4fv");
// 	if(ptrgoglRasterPos4fv == NULL) return 1;
// 	ptrgoglRasterPos4i = goglGetProcAddress("glRasterPos4i");
// 	if(ptrgoglRasterPos4i == NULL) return 1;
// 	ptrgoglRasterPos4iv = goglGetProcAddress("glRasterPos4iv");
// 	if(ptrgoglRasterPos4iv == NULL) return 1;
// 	ptrgoglRasterPos4s = goglGetProcAddress("glRasterPos4s");
// 	if(ptrgoglRasterPos4s == NULL) return 1;
// 	ptrgoglRasterPos4sv = goglGetProcAddress("glRasterPos4sv");
// 	if(ptrgoglRasterPos4sv == NULL) return 1;
// 	ptrgoglRectd = goglGetProcAddress("glRectd");
// 	if(ptrgoglRectd == NULL) return 1;
// 	ptrgoglRectdv = goglGetProcAddress("glRectdv");
// 	if(ptrgoglRectdv == NULL) return 1;
// 	ptrgoglRectf = goglGetProcAddress("glRectf");
// 	if(ptrgoglRectf == NULL) return 1;
// 	ptrgoglRectfv = goglGetProcAddress("glRectfv");
// 	if(ptrgoglRectfv == NULL) return 1;
// 	ptrgoglRecti = goglGetProcAddress("glRecti");
// 	if(ptrgoglRecti == NULL) return 1;
// 	ptrgoglRectiv = goglGetProcAddress("glRectiv");
// 	if(ptrgoglRectiv == NULL) return 1;
// 	ptrgoglRects = goglGetProcAddress("glRects");
// 	if(ptrgoglRects == NULL) return 1;
// 	ptrgoglRectsv = goglGetProcAddress("glRectsv");
// 	if(ptrgoglRectsv == NULL) return 1;
// 	ptrgoglTexCoord1d = goglGetProcAddress("glTexCoord1d");
// 	if(ptrgoglTexCoord1d == NULL) return 1;
// 	ptrgoglTexCoord1dv = goglGetProcAddress("glTexCoord1dv");
// 	if(ptrgoglTexCoord1dv == NULL) return 1;
// 	ptrgoglTexCoord1f = goglGetProcAddress("glTexCoord1f");
// 	if(ptrgoglTexCoord1f == NULL) return 1;
// 	ptrgoglTexCoord1fv = goglGetProcAddress("glTexCoord1fv");
// 	if(ptrgoglTexCoord1fv == NULL) return 1;
// 	ptrgoglTexCoord1i = goglGetProcAddress("glTexCoord1i");
// 	if(ptrgoglTexCoord1i == NULL) return 1;
// 	ptrgoglTexCoord1iv = goglGetProcAddress("glTexCoord1iv");
// 	if(ptrgoglTexCoord1iv == NULL) return 1;
// 	ptrgoglTexCoord1s = goglGetProcAddress("glTexCoord1s");
// 	if(ptrgoglTexCoord1s == NULL) return 1;
// 	ptrgoglTexCoord1sv = goglGetProcAddress("glTexCoord1sv");
// 	if(ptrgoglTexCoord1sv == NULL) return 1;
// 	ptrgoglTexCoord2d = goglGetProcAddress("glTexCoord2d");
// 	if(ptrgoglTexCoord2d == NULL) return 1;
// 	ptrgoglTexCoord2dv = goglGetProcAddress("glTexCoord2dv");
// 	if(ptrgoglTexCoord2dv == NULL) return 1;
// 	ptrgoglTexCoord2f = goglGetProcAddress("glTexCoord2f");
// 	if(ptrgoglTexCoord2f == NULL) return 1;
// 	ptrgoglTexCoord2fv = goglGetProcAddress("glTexCoord2fv");
// 	if(ptrgoglTexCoord2fv == NULL) return 1;
// 	ptrgoglTexCoord2i = goglGetProcAddress("glTexCoord2i");
// 	if(ptrgoglTexCoord2i == NULL) return 1;
// 	ptrgoglTexCoord2iv = goglGetProcAddress("glTexCoord2iv");
// 	if(ptrgoglTexCoord2iv == NULL) return 1;
// 	ptrgoglTexCoord2s = goglGetProcAddress("glTexCoord2s");
// 	if(ptrgoglTexCoord2s == NULL) return 1;
// 	ptrgoglTexCoord2sv = goglGetProcAddress("glTexCoord2sv");
// 	if(ptrgoglTexCoord2sv == NULL) return 1;
// 	ptrgoglTexCoord3d = goglGetProcAddress("glTexCoord3d");
// 	if(ptrgoglTexCoord3d == NULL) return 1;
// 	ptrgoglTexCoord3dv = goglGetProcAddress("glTexCoord3dv");
// 	if(ptrgoglTexCoord3dv == NULL) return 1;
// 	ptrgoglTexCoord3f = goglGetProcAddress("glTexCoord3f");
// 	if(ptrgoglTexCoord3f == NULL) return 1;
// 	ptrgoglTexCoord3fv = goglGetProcAddress("glTexCoord3fv");
// 	if(ptrgoglTexCoord3fv == NULL) return 1;
// 	ptrgoglTexCoord3i = goglGetProcAddress("glTexCoord3i");
// 	if(ptrgoglTexCoord3i == NULL) return 1;
// 	ptrgoglTexCoord3iv = goglGetProcAddress("glTexCoord3iv");
// 	if(ptrgoglTexCoord3iv == NULL) return 1;
// 	ptrgoglTexCoord3s = goglGetProcAddress("glTexCoord3s");
// 	if(ptrgoglTexCoord3s == NULL) return 1;
// 	ptrgoglTexCoord3sv = goglGetProcAddress("glTexCoord3sv");
// 	if(ptrgoglTexCoord3sv == NULL) return 1;
// 	ptrgoglTexCoord4d = goglGetProcAddress("glTexCoord4d");
// 	if(ptrgoglTexCoord4d == NULL) return 1;
// 	ptrgoglTexCoord4dv = goglGetProcAddress("glTexCoord4dv");
// 	if(ptrgoglTexCoord4dv == NULL) return 1;
// 	ptrgoglTexCoord4f = goglGetProcAddress("glTexCoord4f");
// 	if(ptrgoglTexCoord4f == NULL) return 1;
// 	ptrgoglTexCoord4fv = goglGetProcAddress("glTexCoord4fv");
// 	if(ptrgoglTexCoord4fv == NULL) return 1;
// 	ptrgoglTexCoord4i = goglGetProcAddress("glTexCoord4i");
// 	if(ptrgoglTexCoord4i == NULL) return 1;
// 	ptrgoglTexCoord4iv = goglGetProcAddress("glTexCoord4iv");
// 	if(ptrgoglTexCoord4iv == NULL) return 1;
// 	ptrgoglTexCoord4s = goglGetProcAddress("glTexCoord4s");
// 	if(ptrgoglTexCoord4s == NULL) return 1;
// 	ptrgoglTexCoord4sv = goglGetProcAddress("glTexCoord4sv");
// 	if(ptrgoglTexCoord4sv == NULL) return 1;
// 	ptrgoglVertex2d = goglGetProcAddress("glVertex2d");
// 	if(ptrgoglVertex2d == NULL) return 1;
// 	ptrgoglVertex2dv = goglGetProcAddress("glVertex2dv");
// 	if(ptrgoglVertex2dv == NULL) return 1;
// 	ptrgoglVertex2f = goglGetProcAddress("glVertex2f");
// 	if(ptrgoglVertex2f == NULL) return 1;
// 	ptrgoglVertex2fv = goglGetProcAddress("glVertex2fv");
// 	if(ptrgoglVertex2fv == NULL) return 1;
// 	ptrgoglVertex2i = goglGetProcAddress("glVertex2i");
// 	if(ptrgoglVertex2i == NULL) return 1;
// 	ptrgoglVertex2iv = goglGetProcAddress("glVertex2iv");
// 	if(ptrgoglVertex2iv == NULL) return 1;
// 	ptrgoglVertex2s = goglGetProcAddress("glVertex2s");
// 	if(ptrgoglVertex2s == NULL) return 1;
// 	ptrgoglVertex2sv = goglGetProcAddress("glVertex2sv");
// 	if(ptrgoglVertex2sv == NULL) return 1;
// 	ptrgoglVertex3d = goglGetProcAddress("glVertex3d");
// 	if(ptrgoglVertex3d == NULL) return 1;
// 	ptrgoglVertex3dv = goglGetProcAddress("glVertex3dv");
// 	if(ptrgoglVertex3dv == NULL) return 1;
// 	ptrgoglVertex3f = goglGetProcAddress("glVertex3f");
// 	if(ptrgoglVertex3f == NULL) return 1;
// 	ptrgoglVertex3fv = goglGetProcAddress("glVertex3fv");
// 	if(ptrgoglVertex3fv == NULL) return 1;
// 	ptrgoglVertex3i = goglGetProcAddress("glVertex3i");
// 	if(ptrgoglVertex3i == NULL) return 1;
// 	ptrgoglVertex3iv = goglGetProcAddress("glVertex3iv");
// 	if(ptrgoglVertex3iv == NULL) return 1;
// 	ptrgoglVertex3s = goglGetProcAddress("glVertex3s");
// 	if(ptrgoglVertex3s == NULL) return 1;
// 	ptrgoglVertex3sv = goglGetProcAddress("glVertex3sv");
// 	if(ptrgoglVertex3sv == NULL) return 1;
// 	ptrgoglVertex4d = goglGetProcAddress("glVertex4d");
// 	if(ptrgoglVertex4d == NULL) return 1;
// 	ptrgoglVertex4dv = goglGetProcAddress("glVertex4dv");
// 	if(ptrgoglVertex4dv == NULL) return 1;
// 	ptrgoglVertex4f = goglGetProcAddress("glVertex4f");
// 	if(ptrgoglVertex4f == NULL) return 1;
// 	ptrgoglVertex4fv = goglGetProcAddress("glVertex4fv");
// 	if(ptrgoglVertex4fv == NULL) return 1;
// 	ptrgoglVertex4i = goglGetProcAddress("glVertex4i");
// 	if(ptrgoglVertex4i == NULL) return 1;
// 	ptrgoglVertex4iv = goglGetProcAddress("glVertex4iv");
// 	if(ptrgoglVertex4iv == NULL) return 1;
// 	ptrgoglVertex4s = goglGetProcAddress("glVertex4s");
// 	if(ptrgoglVertex4s == NULL) return 1;
// 	ptrgoglVertex4sv = goglGetProcAddress("glVertex4sv");
// 	if(ptrgoglVertex4sv == NULL) return 1;
// 	ptrgoglClipPlane = goglGetProcAddress("glClipPlane");
// 	if(ptrgoglClipPlane == NULL) return 1;
// 	ptrgoglColorMaterial = goglGetProcAddress("glColorMaterial");
// 	if(ptrgoglColorMaterial == NULL) return 1;
// 	ptrgoglFogf = goglGetProcAddress("glFogf");
// 	if(ptrgoglFogf == NULL) return 1;
// 	ptrgoglFogfv = goglGetProcAddress("glFogfv");
// 	if(ptrgoglFogfv == NULL) return 1;
// 	ptrgoglFogi = goglGetProcAddress("glFogi");
// 	if(ptrgoglFogi == NULL) return 1;
// 	ptrgoglFogiv = goglGetProcAddress("glFogiv");
// 	if(ptrgoglFogiv == NULL) return 1;
// 	ptrgoglLightf = goglGetProcAddress("glLightf");
// 	if(ptrgoglLightf == NULL) return 1;
// 	ptrgoglLightfv = goglGetProcAddress("glLightfv");
// 	if(ptrgoglLightfv == NULL) return 1;
// 	ptrgoglLighti = goglGetProcAddress("glLighti");
// 	if(ptrgoglLighti == NULL) return 1;
// 	ptrgoglLightiv = goglGetProcAddress("glLightiv");
// 	if(ptrgoglLightiv == NULL) return 1;
// 	ptrgoglLightModelf = goglGetProcAddress("glLightModelf");
// 	if(ptrgoglLightModelf == NULL) return 1;
// 	ptrgoglLightModelfv = goglGetProcAddress("glLightModelfv");
// 	if(ptrgoglLightModelfv == NULL) return 1;
// 	ptrgoglLightModeli = goglGetProcAddress("glLightModeli");
// 	if(ptrgoglLightModeli == NULL) return 1;
// 	ptrgoglLightModeliv = goglGetProcAddress("glLightModeliv");
// 	if(ptrgoglLightModeliv == NULL) return 1;
// 	ptrgoglLineStipple = goglGetProcAddress("glLineStipple");
// 	if(ptrgoglLineStipple == NULL) return 1;
// 	ptrgoglMaterialf = goglGetProcAddress("glMaterialf");
// 	if(ptrgoglMaterialf == NULL) return 1;
// 	ptrgoglMaterialfv = goglGetProcAddress("glMaterialfv");
// 	if(ptrgoglMaterialfv == NULL) return 1;
// 	ptrgoglMateriali = goglGetProcAddress("glMateriali");
// 	if(ptrgoglMateriali == NULL) return 1;
// 	ptrgoglMaterialiv = goglGetProcAddress("glMaterialiv");
// 	if(ptrgoglMaterialiv == NULL) return 1;
// 	ptrgoglPolygonStipple = goglGetProcAddress("glPolygonStipple");
// 	if(ptrgoglPolygonStipple == NULL) return 1;
// 	ptrgoglShadeModel = goglGetProcAddress("glShadeModel");
// 	if(ptrgoglShadeModel == NULL) return 1;
// 	ptrgoglTexEnvf = goglGetProcAddress("glTexEnvf");
// 	if(ptrgoglTexEnvf == NULL) return 1;
// 	ptrgoglTexEnvfv = goglGetProcAddress("glTexEnvfv");
// 	if(ptrgoglTexEnvfv == NULL) return 1;
// 	ptrgoglTexEnvi = goglGetProcAddress("glTexEnvi");
// 	if(ptrgoglTexEnvi == NULL) return 1;
// 	ptrgoglTexEnviv = goglGetProcAddress("glTexEnviv");
// 	if(ptrgoglTexEnviv == NULL) return 1;
// 	ptrgoglTexGend = goglGetProcAddress("glTexGend");
// 	if(ptrgoglTexGend == NULL) return 1;
// 	ptrgoglTexGendv = goglGetProcAddress("glTexGendv");
// 	if(ptrgoglTexGendv == NULL) return 1;
// 	ptrgoglTexGenf = goglGetProcAddress("glTexGenf");
// 	if(ptrgoglTexGenf == NULL) return 1;
// 	ptrgoglTexGenfv = goglGetProcAddress("glTexGenfv");
// 	if(ptrgoglTexGenfv == NULL) return 1;
// 	ptrgoglTexGeni = goglGetProcAddress("glTexGeni");
// 	if(ptrgoglTexGeni == NULL) return 1;
// 	ptrgoglTexGeniv = goglGetProcAddress("glTexGeniv");
// 	if(ptrgoglTexGeniv == NULL) return 1;
// 	ptrgoglFeedbackBuffer = goglGetProcAddress("glFeedbackBuffer");
// 	if(ptrgoglFeedbackBuffer == NULL) return 1;
// 	ptrgoglSelectBuffer = goglGetProcAddress("glSelectBuffer");
// 	if(ptrgoglSelectBuffer == NULL) return 1;
// 	ptrgoglRenderMode = goglGetProcAddress("glRenderMode");
// 	if(ptrgoglRenderMode == NULL) return 1;
// 	ptrgoglInitNames = goglGetProcAddress("glInitNames");
// 	if(ptrgoglInitNames == NULL) return 1;
// 	ptrgoglLoadName = goglGetProcAddress("glLoadName");
// 	if(ptrgoglLoadName == NULL) return 1;
// 	ptrgoglPassThrough = goglGetProcAddress("glPassThrough");
// 	if(ptrgoglPassThrough == NULL) return 1;
// 	ptrgoglPopName = goglGetProcAddress("glPopName");
// 	if(ptrgoglPopName == NULL) return 1;
// 	ptrgoglPushName = goglGetProcAddress("glPushName");
// 	if(ptrgoglPushName == NULL) return 1;
// 	ptrgoglClearAccum = goglGetProcAddress("glClearAccum");
// 	if(ptrgoglClearAccum == NULL) return 1;
// 	ptrgoglClearIndex = goglGetProcAddress("glClearIndex");
// 	if(ptrgoglClearIndex == NULL) return 1;
// 	ptrgoglIndexMask = goglGetProcAddress("glIndexMask");
// 	if(ptrgoglIndexMask == NULL) return 1;
// 	ptrgoglAccum = goglGetProcAddress("glAccum");
// 	if(ptrgoglAccum == NULL) return 1;
// 	ptrgoglPopAttrib = goglGetProcAddress("glPopAttrib");
// 	if(ptrgoglPopAttrib == NULL) return 1;
// 	ptrgoglPushAttrib = goglGetProcAddress("glPushAttrib");
// 	if(ptrgoglPushAttrib == NULL) return 1;
// 	ptrgoglMap1d = goglGetProcAddress("glMap1d");
// 	if(ptrgoglMap1d == NULL) return 1;
// 	ptrgoglMap1f = goglGetProcAddress("glMap1f");
// 	if(ptrgoglMap1f == NULL) return 1;
// 	ptrgoglMap2d = goglGetProcAddress("glMap2d");
// 	if(ptrgoglMap2d == NULL) return 1;
// 	ptrgoglMap2f = goglGetProcAddress("glMap2f");
// 	if(ptrgoglMap2f == NULL) return 1;
// 	ptrgoglMapGrid1d = goglGetProcAddress("glMapGrid1d");
// 	if(ptrgoglMapGrid1d == NULL) return 1;
// 	ptrgoglMapGrid1f = goglGetProcAddress("glMapGrid1f");
// 	if(ptrgoglMapGrid1f == NULL) return 1;
// 	ptrgoglMapGrid2d = goglGetProcAddress("glMapGrid2d");
// 	if(ptrgoglMapGrid2d == NULL) return 1;
// 	ptrgoglMapGrid2f = goglGetProcAddress("glMapGrid2f");
// 	if(ptrgoglMapGrid2f == NULL) return 1;
// 	ptrgoglEvalCoord1d = goglGetProcAddress("glEvalCoord1d");
// 	if(ptrgoglEvalCoord1d == NULL) return 1;
// 	ptrgoglEvalCoord1dv = goglGetProcAddress("glEvalCoord1dv");
// 	if(ptrgoglEvalCoord1dv == NULL) return 1;
// 	ptrgoglEvalCoord1f = goglGetProcAddress("glEvalCoord1f");
// 	if(ptrgoglEvalCoord1f == NULL) return 1;
// 	ptrgoglEvalCoord1fv = goglGetProcAddress("glEvalCoord1fv");
// 	if(ptrgoglEvalCoord1fv == NULL) return 1;
// 	ptrgoglEvalCoord2d = goglGetProcAddress("glEvalCoord2d");
// 	if(ptrgoglEvalCoord2d == NULL) return 1;
// 	ptrgoglEvalCoord2dv = goglGetProcAddress("glEvalCoord2dv");
// 	if(ptrgoglEvalCoord2dv == NULL) return 1;
// 	ptrgoglEvalCoord2f = goglGetProcAddress("glEvalCoord2f");
// 	if(ptrgoglEvalCoord2f == NULL) return 1;
// 	ptrgoglEvalCoord2fv = goglGetProcAddress("glEvalCoord2fv");
// 	if(ptrgoglEvalCoord2fv == NULL) return 1;
// 	ptrgoglEvalMesh1 = goglGetProcAddress("glEvalMesh1");
// 	if(ptrgoglEvalMesh1 == NULL) return 1;
// 	ptrgoglEvalPoint1 = goglGetProcAddress("glEvalPoint1");
// 	if(ptrgoglEvalPoint1 == NULL) return 1;
// 	ptrgoglEvalMesh2 = goglGetProcAddress("glEvalMesh2");
// 	if(ptrgoglEvalMesh2 == NULL) return 1;
// 	ptrgoglEvalPoint2 = goglGetProcAddress("glEvalPoint2");
// 	if(ptrgoglEvalPoint2 == NULL) return 1;
// 	ptrgoglAlphaFunc = goglGetProcAddress("glAlphaFunc");
// 	if(ptrgoglAlphaFunc == NULL) return 1;
// 	ptrgoglPixelZoom = goglGetProcAddress("glPixelZoom");
// 	if(ptrgoglPixelZoom == NULL) return 1;
// 	ptrgoglPixelTransferf = goglGetProcAddress("glPixelTransferf");
// 	if(ptrgoglPixelTransferf == NULL) return 1;
// 	ptrgoglPixelTransferi = goglGetProcAddress("glPixelTransferi");
// 	if(ptrgoglPixelTransferi == NULL) return 1;
// 	ptrgoglPixelMapfv = goglGetProcAddress("glPixelMapfv");
// 	if(ptrgoglPixelMapfv == NULL) return 1;
// 	ptrgoglPixelMapuiv = goglGetProcAddress("glPixelMapuiv");
// 	if(ptrgoglPixelMapuiv == NULL) return 1;
// 	ptrgoglPixelMapusv = goglGetProcAddress("glPixelMapusv");
// 	if(ptrgoglPixelMapusv == NULL) return 1;
// 	ptrgoglCopyPixels = goglGetProcAddress("glCopyPixels");
// 	if(ptrgoglCopyPixels == NULL) return 1;
// 	ptrgoglDrawPixels = goglGetProcAddress("glDrawPixels");
// 	if(ptrgoglDrawPixels == NULL) return 1;
// 	ptrgoglGetClipPlane = goglGetProcAddress("glGetClipPlane");
// 	if(ptrgoglGetClipPlane == NULL) return 1;
// 	ptrgoglGetLightfv = goglGetProcAddress("glGetLightfv");
// 	if(ptrgoglGetLightfv == NULL) return 1;
// 	ptrgoglGetLightiv = goglGetProcAddress("glGetLightiv");
// 	if(ptrgoglGetLightiv == NULL) return 1;
// 	ptrgoglGetMapdv = goglGetProcAddress("glGetMapdv");
// 	if(ptrgoglGetMapdv == NULL) return 1;
// 	ptrgoglGetMapfv = goglGetProcAddress("glGetMapfv");
// 	if(ptrgoglGetMapfv == NULL) return 1;
// 	ptrgoglGetMapiv = goglGetProcAddress("glGetMapiv");
// 	if(ptrgoglGetMapiv == NULL) return 1;
// 	ptrgoglGetMaterialfv = goglGetProcAddress("glGetMaterialfv");
// 	if(ptrgoglGetMaterialfv == NULL) return 1;
// 	ptrgoglGetMaterialiv = goglGetProcAddress("glGetMaterialiv");
// 	if(ptrgoglGetMaterialiv == NULL) return 1;
// 	ptrgoglGetPixelMapfv = goglGetProcAddress("glGetPixelMapfv");
// 	if(ptrgoglGetPixelMapfv == NULL) return 1;
// 	ptrgoglGetPixelMapuiv = goglGetProcAddress("glGetPixelMapuiv");
// 	if(ptrgoglGetPixelMapuiv == NULL) return 1;
// 	ptrgoglGetPixelMapusv = goglGetProcAddress("glGetPixelMapusv");
// 	if(ptrgoglGetPixelMapusv == NULL) return 1;
// 	ptrgoglGetPolygonStipple = goglGetProcAddress("glGetPolygonStipple");
// 	if(ptrgoglGetPolygonStipple == NULL) return 1;
// 	ptrgoglGetTexEnvfv = goglGetProcAddress("glGetTexEnvfv");
// 	if(ptrgoglGetTexEnvfv == NULL) return 1;
// 	ptrgoglGetTexEnviv = goglGetProcAddress("glGetTexEnviv");
// 	if(ptrgoglGetTexEnviv == NULL) return 1;
// 	ptrgoglGetTexGendv = goglGetProcAddress("glGetTexGendv");
// 	if(ptrgoglGetTexGendv == NULL) return 1;
// 	ptrgoglGetTexGenfv = goglGetProcAddress("glGetTexGenfv");
// 	if(ptrgoglGetTexGenfv == NULL) return 1;
// 	ptrgoglGetTexGeniv = goglGetProcAddress("glGetTexGeniv");
// 	if(ptrgoglGetTexGeniv == NULL) return 1;
// 	ptrgoglIsList = goglGetProcAddress("glIsList");
// 	if(ptrgoglIsList == NULL) return 1;
// 	ptrgoglFrustum = goglGetProcAddress("glFrustum");
// 	if(ptrgoglFrustum == NULL) return 1;
// 	ptrgoglLoadIdentity = goglGetProcAddress("glLoadIdentity");
// 	if(ptrgoglLoadIdentity == NULL) return 1;
// 	ptrgoglLoadMatrixf = goglGetProcAddress("glLoadMatrixf");
// 	if(ptrgoglLoadMatrixf == NULL) return 1;
// 	ptrgoglLoadMatrixd = goglGetProcAddress("glLoadMatrixd");
// 	if(ptrgoglLoadMatrixd == NULL) return 1;
// 	ptrgoglMatrixMode = goglGetProcAddress("glMatrixMode");
// 	if(ptrgoglMatrixMode == NULL) return 1;
// 	ptrgoglMultMatrixf = goglGetProcAddress("glMultMatrixf");
// 	if(ptrgoglMultMatrixf == NULL) return 1;
// 	ptrgoglMultMatrixd = goglGetProcAddress("glMultMatrixd");
// 	if(ptrgoglMultMatrixd == NULL) return 1;
// 	ptrgoglOrtho = goglGetProcAddress("glOrtho");
// 	if(ptrgoglOrtho == NULL) return 1;
// 	ptrgoglPopMatrix = goglGetProcAddress("glPopMatrix");
// 	if(ptrgoglPopMatrix == NULL) return 1;
// 	ptrgoglPushMatrix = goglGetProcAddress("glPushMatrix");
// 	if(ptrgoglPushMatrix == NULL) return 1;
// 	ptrgoglRotated = goglGetProcAddress("glRotated");
// 	if(ptrgoglRotated == NULL) return 1;
// 	ptrgoglRotatef = goglGetProcAddress("glRotatef");
// 	if(ptrgoglRotatef == NULL) return 1;
// 	ptrgoglScaled = goglGetProcAddress("glScaled");
// 	if(ptrgoglScaled == NULL) return 1;
// 	ptrgoglScalef = goglGetProcAddress("glScalef");
// 	if(ptrgoglScalef == NULL) return 1;
// 	ptrgoglTranslated = goglGetProcAddress("glTranslated");
// 	if(ptrgoglTranslated == NULL) return 1;
// 	ptrgoglTranslatef = goglGetProcAddress("glTranslatef");
// 	if(ptrgoglTranslatef == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_2_DEPRECATED() {
// 	ptrgoglColorTable = goglGetProcAddress("glColorTable");
// 	if(ptrgoglColorTable == NULL) return 1;
// 	ptrgoglColorTableParameterfv = goglGetProcAddress("glColorTableParameterfv");
// 	if(ptrgoglColorTableParameterfv == NULL) return 1;
// 	ptrgoglColorTableParameteriv = goglGetProcAddress("glColorTableParameteriv");
// 	if(ptrgoglColorTableParameteriv == NULL) return 1;
// 	ptrgoglCopyColorTable = goglGetProcAddress("glCopyColorTable");
// 	if(ptrgoglCopyColorTable == NULL) return 1;
// 	ptrgoglGetColorTable = goglGetProcAddress("glGetColorTable");
// 	if(ptrgoglGetColorTable == NULL) return 1;
// 	ptrgoglGetColorTableParameterfv = goglGetProcAddress("glGetColorTableParameterfv");
// 	if(ptrgoglGetColorTableParameterfv == NULL) return 1;
// 	ptrgoglGetColorTableParameteriv = goglGetProcAddress("glGetColorTableParameteriv");
// 	if(ptrgoglGetColorTableParameteriv == NULL) return 1;
// 	ptrgoglColorSubTable = goglGetProcAddress("glColorSubTable");
// 	if(ptrgoglColorSubTable == NULL) return 1;
// 	ptrgoglCopyColorSubTable = goglGetProcAddress("glCopyColorSubTable");
// 	if(ptrgoglCopyColorSubTable == NULL) return 1;
// 	ptrgoglConvolutionFilter1D = goglGetProcAddress("glConvolutionFilter1D");
// 	if(ptrgoglConvolutionFilter1D == NULL) return 1;
// 	ptrgoglConvolutionFilter2D = goglGetProcAddress("glConvolutionFilter2D");
// 	if(ptrgoglConvolutionFilter2D == NULL) return 1;
// 	ptrgoglConvolutionParameterf = goglGetProcAddress("glConvolutionParameterf");
// 	if(ptrgoglConvolutionParameterf == NULL) return 1;
// 	ptrgoglConvolutionParameterfv = goglGetProcAddress("glConvolutionParameterfv");
// 	if(ptrgoglConvolutionParameterfv == NULL) return 1;
// 	ptrgoglConvolutionParameteri = goglGetProcAddress("glConvolutionParameteri");
// 	if(ptrgoglConvolutionParameteri == NULL) return 1;
// 	ptrgoglConvolutionParameteriv = goglGetProcAddress("glConvolutionParameteriv");
// 	if(ptrgoglConvolutionParameteriv == NULL) return 1;
// 	ptrgoglCopyConvolutionFilter1D = goglGetProcAddress("glCopyConvolutionFilter1D");
// 	if(ptrgoglCopyConvolutionFilter1D == NULL) return 1;
// 	ptrgoglCopyConvolutionFilter2D = goglGetProcAddress("glCopyConvolutionFilter2D");
// 	if(ptrgoglCopyConvolutionFilter2D == NULL) return 1;
// 	ptrgoglGetConvolutionFilter = goglGetProcAddress("glGetConvolutionFilter");
// 	if(ptrgoglGetConvolutionFilter == NULL) return 1;
// 	ptrgoglGetConvolutionParameterfv = goglGetProcAddress("glGetConvolutionParameterfv");
// 	if(ptrgoglGetConvolutionParameterfv == NULL) return 1;
// 	ptrgoglGetConvolutionParameteriv = goglGetProcAddress("glGetConvolutionParameteriv");
// 	if(ptrgoglGetConvolutionParameteriv == NULL) return 1;
// 	ptrgoglGetSeparableFilter = goglGetProcAddress("glGetSeparableFilter");
// 	if(ptrgoglGetSeparableFilter == NULL) return 1;
// 	ptrgoglSeparableFilter2D = goglGetProcAddress("glSeparableFilter2D");
// 	if(ptrgoglSeparableFilter2D == NULL) return 1;
// 	ptrgoglGetHistogram = goglGetProcAddress("glGetHistogram");
// 	if(ptrgoglGetHistogram == NULL) return 1;
// 	ptrgoglGetHistogramParameterfv = goglGetProcAddress("glGetHistogramParameterfv");
// 	if(ptrgoglGetHistogramParameterfv == NULL) return 1;
// 	ptrgoglGetHistogramParameteriv = goglGetProcAddress("glGetHistogramParameteriv");
// 	if(ptrgoglGetHistogramParameteriv == NULL) return 1;
// 	ptrgoglGetMinmax = goglGetProcAddress("glGetMinmax");
// 	if(ptrgoglGetMinmax == NULL) return 1;
// 	ptrgoglGetMinmaxParameterfv = goglGetProcAddress("glGetMinmaxParameterfv");
// 	if(ptrgoglGetMinmaxParameterfv == NULL) return 1;
// 	ptrgoglGetMinmaxParameteriv = goglGetProcAddress("glGetMinmaxParameteriv");
// 	if(ptrgoglGetMinmaxParameteriv == NULL) return 1;
// 	ptrgoglHistogram = goglGetProcAddress("glHistogram");
// 	if(ptrgoglHistogram == NULL) return 1;
// 	ptrgoglMinmax = goglGetProcAddress("glMinmax");
// 	if(ptrgoglMinmax == NULL) return 1;
// 	ptrgoglResetHistogram = goglGetProcAddress("glResetHistogram");
// 	if(ptrgoglResetHistogram == NULL) return 1;
// 	ptrgoglResetMinmax = goglGetProcAddress("glResetMinmax");
// 	if(ptrgoglResetMinmax == NULL) return 1;
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
	Int64    C.GLint64
	Uint64   C.GLuint64
	Intptr   C.GLintptr
	Sizeiptr C.GLsizeiptr
)

// VERSION_2_1
const (
	COMPRESSED_SRGB = 0x8C48
	SRGB = 0x8C40
	SRGB_ALPHA = 0x8C42
	FLOAT_MAT2x4 = 0x8B66
	FLOAT_MAT2x3 = 0x8B65
	PIXEL_UNPACK_BUFFER = 0x88EC
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
)
// VERSION_2_0
const (
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
)
// VERSION_1_3_DEPRECATED
const (
	COMPRESSED_LUMINANCE = 0x84EA
	DOT3_RGBA = 0x86AF
	OPERAND0_RGB = 0x8590
	DOT3_RGB = 0x86AE
	CLIENT_ACTIVE_TEXTURE = 0x84E1
	COMBINE = 0x8570
	TRANSPOSE_MODELVIEW_MATRIX = 0x84E3
	SOURCE1_RGB = 0x8581
	INTERPOLATE = 0x8575
	RGB_SCALE = 0x8573
	COMBINE_ALPHA = 0x8572
	TRANSPOSE_TEXTURE_MATRIX = 0x84E5
	OPERAND1_ALPHA = 0x8599
	TRANSPOSE_COLOR_MATRIX = 0x84E6
	OPERAND1_RGB = 0x8591
	PREVIOUS = 0x8578
	OPERAND2_RGB = 0x8592
	CONSTANT = 0x8576
	SOURCE2_RGB = 0x8582
	MAX_TEXTURE_UNITS = 0x84E2
	ADD_SIGNED = 0x8574
	NORMAL_MAP = 0x8511
	SOURCE1_ALPHA = 0x8589
	COMPRESSED_INTENSITY = 0x84EC
	OPERAND2_ALPHA = 0x859A
	OPERAND0_ALPHA = 0x8598
	SOURCE0_ALPHA = 0x8588
	COMPRESSED_LUMINANCE_ALPHA = 0x84EB
	MULTISAMPLE_BIT = 0x20000000
	SUBTRACT = 0x84E7
	TRANSPOSE_PROJECTION_MATRIX = 0x84E4
	SOURCE2_ALPHA = 0x858A
	REFLECTION_MAP = 0x8512
	SOURCE0_RGB = 0x8580
	COMPRESSED_ALPHA = 0x84E9
	PRIMARY_COLOR = 0x8577
	COMBINE_RGB = 0x8571
)
// VERSION_2_1_DEPRECATED
const (
	SLUMINANCE8 = 0x8C47
	COMPRESSED_SLUMINANCE = 0x8C4A
	COMPRESSED_SLUMINANCE_ALPHA = 0x8C4B
	CURRENT_RASTER_SECONDARY_COLOR = 0x845F
	SLUMINANCE_ALPHA = 0x8C44
	SLUMINANCE = 0x8C46
	SLUMINANCE8_ALPHA8 = 0x8C45
)
// VERSION_1_5_DEPRECATED
const (
	FOG_COORDINATE_ARRAY_BUFFER_BINDING = 0x889D
	SRC2_RGB = 0x8582
	FOG_COORD_SRC = 0x8450
	VERTEX_ARRAY_BUFFER_BINDING = 0x8896
	SRC1_ALPHA = 0x8589
	FOG_COORD_ARRAY_TYPE = 0x8454
	SRC0_RGB = 0x8580
	WEIGHT_ARRAY_BUFFER_BINDING = 0x889E
	EDGE_FLAG_ARRAY_BUFFER_BINDING = 0x889B
	COLOR_ARRAY_BUFFER_BINDING = 0x8898
	FOG_COORD_ARRAY = 0x8457
	FOG_COORD_ARRAY_BUFFER_BINDING = 0x889D
	FOG_COORD = 0x8451
	TEXTURE_COORD_ARRAY_BUFFER_BINDING = 0x889A
	FOG_COORD_ARRAY_POINTER = 0x8456
	FOG_COORD_ARRAY_STRIDE = 0x8455
	SRC1_RGB = 0x8581
	SRC0_ALPHA = 0x8588
	CURRENT_FOG_COORD = 0x8453
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING = 0x889C
	NORMAL_ARRAY_BUFFER_BINDING = 0x8897
	SRC2_ALPHA = 0x858A
	INDEX_ARRAY_BUFFER_BINDING = 0x8899
)
// VERSION_1_4
const (
	BLEND_DST_ALPHA = 0x80CA
	POINT_FADE_THRESHOLD_SIZE = 0x8128
	INCR_WRAP = 0x8507
	TEXTURE_LOD_BIAS = 0x8501
	BLEND_SRC_ALPHA = 0x80CB
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
)
// VERSION_1_4_DEPRECATED
const (
	FOG_COORDINATE = 0x8451
	POINT_SIZE_MAX = 0x8127
	FOG_COORDINATE_ARRAY_STRIDE = 0x8455
	SECONDARY_COLOR_ARRAY_STRIDE = 0x845C
	TEXTURE_FILTER_CONTROL = 0x8500
	DEPTH_TEXTURE_MODE = 0x884B
	COLOR_SUM = 0x8458
	SECONDARY_COLOR_ARRAY_SIZE = 0x845A
	CURRENT_SECONDARY_COLOR = 0x8459
	FOG_COORDINATE_ARRAY_POINTER = 0x8456
	FOG_COORDINATE_SOURCE = 0x8450
	GENERATE_MIPMAP_HINT = 0x8192
	SECONDARY_COLOR_ARRAY_POINTER = 0x845D
	SECONDARY_COLOR_ARRAY_TYPE = 0x845B
	FOG_COORDINATE_ARRAY = 0x8457
	POINT_SIZE_MIN = 0x8126
	FOG_COORDINATE_ARRAY_TYPE = 0x8454
	CURRENT_FOG_COORDINATE = 0x8453
	FRAGMENT_DEPTH = 0x8452
	SECONDARY_COLOR_ARRAY = 0x845E
	GENERATE_MIPMAP = 0x8191
	COMPARE_R_TO_TEXTURE = 0x884E
	POINT_DISTANCE_ATTENUATION = 0x8129
)
// VERSION_1_2
const (
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
)
// VERSION_1_3
const (
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
)
// VERSION_1_1_DEPRECATED
const (
	QUADRATIC_ATTENUATION = 0x1209
	ACCUM_BUFFER_BIT = 0x00000200
	TEXTURE_GEN_MODE = 0x2500
	ALPHA_SCALE = 0x0D1C
	CLIENT_ALL_ATTRIB_BITS = 0xFFFFFFFF
	T4F_C4F_N3F_V4F = 0x2A2D
	TEXTURE_STACK_DEPTH = 0x0BA5
	INTENSITY12 = 0x804C
	FEEDBACK_BUFFER_SIZE = 0x0DF1
	FOG_BIT = 0x00000080
	Q = 0x2003
	PIXEL_MAP_G_TO_G = 0x0C77
	TEXTURE_MATRIX = 0x0BA8
	INTENSITY16 = 0x804D
	S = 0x2000
	TEXTURE_ENV_COLOR = 0x2201
	R = 0x2002
	T2F_C3F_V3F = 0x2A2A
	T = 0x2001
	PIXEL_MAP_I_TO_B_SIZE = 0x0CB4
	MAP_COLOR = 0x0D10
	PROJECTION = 0x1701
	EDGE_FLAG_ARRAY_POINTER = 0x8093
	VERTEX_ARRAY_POINTER = 0x808E
	CURRENT_INDEX = 0x0B01
	LIST_MODE = 0x0B30
	LINE_STIPPLE_PATTERN = 0x0B25
	NORMALIZE = 0x0BA1
	DEPTH_BITS = 0x0D56
	ATTRIB_STACK_DEPTH = 0x0BB0
	SPOT_EXPONENT = 0x1205
	PROJECTION_MATRIX = 0x0BA7
	TEXTURE_GEN_T = 0x0C61
	EYE_LINEAR = 0x2400
	NORMAL_ARRAY_STRIDE = 0x807F
	ACCUM_CLEAR_VALUE = 0x0B80
	TEXTURE_GEN_R = 0x0C62
	ORDER = 0x0A01
	TEXTURE_GEN_S = 0x0C60
	TEXTURE_GEN_Q = 0x0C63
	LIGHTING_BIT = 0x00000040
	BLUE_BIAS = 0x0D1B
	COLOR_ARRAY_STRIDE = 0x8083
	PIXEL_MAP_A_TO_A = 0x0C79
	BITMAP = 0x1A00
	V2F = 0x2A20
	FOG_END = 0x0B64
	LUMINANCE8 = 0x8040
	C4UB_V2F = 0x2A22
	TEXTURE_COORD_ARRAY_TYPE = 0x8089
	MAP2_GRID_DOMAIN = 0x0DD2
	EXP2 = 0x0801
	COLOR_MATERIAL_PARAMETER = 0x0B56
	INDEX_ARRAY_STRIDE = 0x8086
	VERTEX_ARRAY = 0x8074
	EYE_PLANE = 0x2502
	LUMINANCE4 = 0x803F
	MAX_LIST_NESTING = 0x0B31
	MAP_STENCIL = 0x0D11
	LIGHTING = 0x0B50
	ALPHA_BIAS = 0x0D1D
	PERSPECTIVE_CORRECTION_HINT = 0x0C50
	RED_SCALE = 0x0D14
	INDEX_LOGIC_OP = 0x0BF1
	DOMAIN = 0x0A02
	PIXEL_MAP_G_TO_G_SIZE = 0x0CB7
	X4D_COLOR_TEXTURE = 0x0604
	MAX_MODELVIEW_STACK_DEPTH = 0x0D36
	POINT_BIT = 0x00000002
	COEFF = 0x0A00
	ACCUM_BLUE_BITS = 0x0D5A
	INDEX_BITS = 0x0D51
	LUMINANCE12_ALPHA4 = 0x8046
	SMOOTH = 0x1D01
	RENDER_MODE = 0x0C40
	ACCUM_GREEN_BITS = 0x0D59
	MAP1_INDEX = 0x0D91
	X3_BYTES = 0x1408
	OBJECT_PLANE = 0x2501
	MAP1_GRID_SEGMENTS = 0x0DD1
	MAX_LIGHTS = 0x0D31
	C3F_V3F = 0x2A24
	POLYGON_BIT = 0x00000008
	TEXTURE_INTENSITY_SIZE = 0x8061
	PIXEL_MAP_R_TO_R_SIZE = 0x0CB6
	FLAT = 0x1D00
	INDEX_WRITEMASK = 0x0C21
	PIXEL_MAP_B_TO_B_SIZE = 0x0CB8
	LUMINANCE_ALPHA = 0x190A
	LIGHT_MODEL_AMBIENT = 0x0B53
	MAP1_GRID_DOMAIN = 0x0DD0
	LIST_BASE = 0x0B32
	COLOR_INDEXES = 0x1603
	V3F = 0x2A21
	BLUE_BITS = 0x0D54
	SPECULAR = 0x1202
	TEXTURE_BIT = 0x00040000
	MATRIX_MODE = 0x0BA0
	COPY_PIXEL_TOKEN = 0x0706
	MAX_CLIP_PLANES = 0x0D32
	PROJECTION_STACK_DEPTH = 0x0BA4
	FOG_START = 0x0B63
	FEEDBACK_BUFFER_TYPE = 0x0DF2
	COMPILE = 0x1300
	LUMINANCE8_ALPHA8 = 0x8045
	POLYGON_MODE = 0x0B40
	TRANSFORM_BIT = 0x00001000
	COLOR_ARRAY_SIZE = 0x8081
	INTENSITY8 = 0x804B
	COLOR_INDEX = 0x1900
	INTENSITY4 = 0x804A
	GREEN_SCALE = 0x0D18
	FEEDBACK = 0x1C01
	COLOR_ARRAY_POINTER = 0x8090
	RENDER = 0x1C00
	TEXTURE_LUMINANCE_SIZE = 0x8060
	AMBIENT = 0x1200
	T2F_V3F = 0x2A27
	MAX_PIXEL_MAP_TABLE = 0x0D34
	VERTEX_ARRAY_STRIDE = 0x807C
	MODELVIEW_STACK_DEPTH = 0x0BA3
	POINT_SMOOTH = 0x0B10
	DIFFUSE = 0x1201
	LUMINANCE12_ALPHA12 = 0x8047
	LINE_TOKEN = 0x0702
	TEXTURE_PRIORITY = 0x8066
	TEXTURE_COMPONENTS = 0x1003
	EMISSION = 0x1600
	POSITION = 0x1203
	MAP1_TEXTURE_COORD_4 = 0x0D96
	MAP1_TEXTURE_COORD_3 = 0x0D95
	MAP1_TEXTURE_COORD_2 = 0x0D94
	MAP1_COLOR_4 = 0x0D90
	MAP1_TEXTURE_COORD_1 = 0x0D93
	ALPHA_TEST = 0x0BC0
	X3D = 0x0601
	PIXEL_MAP_R_TO_R = 0x0C76
	LINEAR_ATTENUATION = 0x1208
	CURRENT_RASTER_DISTANCE = 0x0B09
	PIXEL_MAP_S_TO_S_SIZE = 0x0CB1
	N3F_V3F = 0x2A25
	PIXEL_MAP_I_TO_R = 0x0C72
	MAX_TEXTURE_STACK_DEPTH = 0x0D39
	COLOR_ARRAY = 0x8076
	PIXEL_MAP_I_TO_G = 0x0C73
	TEXTURE_BORDER = 0x1005
	PIXEL_MAP_I_TO_B = 0x0C74
	PIXEL_MAP_I_TO_A = 0x0C75
	LUMINANCE4_ALPHA4 = 0x8043
	SELECT = 0x1C02
	LINE_BIT = 0x00000004
	PIXEL_MAP_I_TO_I = 0x0C70
	PIXEL_MAP_A_TO_A_SIZE = 0x0CB9
	NORMAL_ARRAY_POINTER = 0x808F
	POLYGON_STIPPLE_BIT = 0x00000010
	LUMINANCE = 0x1909
	POINT_SMOOTH_HINT = 0x0C51
	TEXTURE_COORD_ARRAY_POINTER = 0x8092
	CURRENT_NORMAL = 0x0B02
	PIXEL_MAP_S_TO_S = 0x0C71
	POLYGON = 0x0009
	X3D_COLOR_TEXTURE = 0x0603
	TEXTURE_COORD_ARRAY_STRIDE = 0x808A
	X2_BYTES = 0x1407
	COLOR_ARRAY_TYPE = 0x8082
	MODELVIEW = 0x1700
	CURRENT_RASTER_TEXTURE_COORDS = 0x0B06
	INDEX_OFFSET = 0x0D13
	STACK_OVERFLOW = 0x0503
	T4F_V4F = 0x2A28
	ENABLE_BIT = 0x00002000
	MAP2_VERTEX_4 = 0x0DB8
	MAP2_COLOR_4 = 0x0DB0
	SPOT_DIRECTION = 0x1204
	CLAMP = 0x2900
	AUTO_NORMAL = 0x0D80
	INDEX_ARRAY_TYPE = 0x8085
	MAP2_VERTEX_3 = 0x0DB7
	QUADS = 0x0007
	X2D = 0x0600
	SPOT_CUTOFF = 0x1206
	MODELVIEW_MATRIX = 0x0BA6
	LIGHT_MODEL_LOCAL_VIEWER = 0x0B51
	MULT = 0x0103
	INTENSITY = 0x8049
	COLOR_MATERIAL_FACE = 0x0B55
	NAME_STACK_DEPTH = 0x0D70
	VERTEX_ARRAY_TYPE = 0x807B
	FEEDBACK_BUFFER_POINTER = 0x0DF0
	INDEX_SHIFT = 0x0D12
	RED_BIAS = 0x0D15
	EDGE_FLAG_ARRAY = 0x8079
	INDEX_CLEAR_VALUE = 0x0C20
	TEXTURE_COORD_ARRAY = 0x8078
	QUAD_STRIP = 0x0008
	CURRENT_RASTER_POSITION = 0x0B07
	INDEX_ARRAY_POINTER = 0x8091
	LIGHT_MODEL_TWO_SIDE = 0x0B52
	TEXTURE_COORD_ARRAY_SIZE = 0x8088
	MAX_PROJECTION_STACK_DEPTH = 0x0D38
	PIXEL_MAP_I_TO_G_SIZE = 0x0CB3
	SELECTION_BUFFER_SIZE = 0x0DF4
	DECAL = 0x2101
	MAP1_NORMAL = 0x0D92
	SHADE_MODEL = 0x0B54
	T2F_N3F_V3F = 0x2A2B
	CLIENT_VERTEX_ARRAY_BIT = 0x00000002
	ALPHA_TEST_REF = 0x0BC2
	CURRENT_RASTER_COLOR = 0x0B04
	AMBIENT_AND_DIFFUSE = 0x1602
	CLIP_PLANE2 = 0x3002
	ALPHA_TEST_FUNC = 0x0BC1
	CLIP_PLANE3 = 0x3003
	T2F_C4F_N3F_V3F = 0x2A2C
	CLIP_PLANE0 = 0x3000
	CLIP_PLANE1 = 0x3001
	CLIP_PLANE4 = 0x3004
	CLIP_PLANE5 = 0x3005
	MAP1_VERTEX_4 = 0x0D98
	POLYGON_STIPPLE = 0x0B42
	MAP1_VERTEX_3 = 0x0D97
	FOG_HINT = 0x0C54
	FOG_COLOR = 0x0B66
	DRAW_PIXEL_TOKEN = 0x0705
	LOAD = 0x0101
	TEXTURE_ENV_MODE = 0x2200
	NORMAL_ARRAY_TYPE = 0x807E
	BITMAP_TOKEN = 0x0704
	ALPHA_BITS = 0x0D55
	ALPHA16 = 0x803E
	CURRENT_RASTER_INDEX = 0x0B05
	ACCUM = 0x0100
	STACK_UNDERFLOW = 0x0504
	ALPHA12 = 0x803D
	PIXEL_MAP_B_TO_B = 0x0C78
	SCISSOR_BIT = 0x00080000
	EXP = 0x0800
	CLIENT_PIXEL_STORE_BIT = 0x00000001
	PIXEL_MAP_I_TO_R_SIZE = 0x0CB2
	X3D_COLOR = 0x0602
	DEPTH_BIAS = 0x0D1F
	TEXTURE_ENV = 0x2300
	MODULATE = 0x2100
	ALPHA4 = 0x803B
	VIEWPORT_BIT = 0x00000800
	EDGE_FLAG = 0x0B43
	CURRENT_TEXTURE_COORDS = 0x0B03
	ALPHA8 = 0x803C
	INDEX_MODE = 0x0C30
	STENCIL_BITS = 0x0D57
	AUX_BUFFERS = 0x0C00
	ACCUM_RED_BITS = 0x0D58
	PASS_THROUGH_TOKEN = 0x0700
	ZOOM_Y = 0x0D17
	CLIENT_ATTRIB_STACK_DEPTH = 0x0BB1
	MAX_EVAL_ORDER = 0x0D30
	ZOOM_X = 0x0D16
	LUMINANCE12 = 0x8041
	LUMINANCE16 = 0x8042
	COMPILE_AND_EXECUTE = 0x1301
	TEXTURE_RESIDENT = 0x8067
	CONSTANT_ATTENUATION = 0x1207
	FOG_MODE = 0x0B65
	COLOR_MATERIAL = 0x0B57
	X4_BYTES = 0x1409
	PIXEL_MODE_BIT = 0x00000020
	SHININESS = 0x1601
	NORMAL_ARRAY = 0x8075
	CURRENT_RASTER_POSITION_VALID = 0x0B08
	RETURN = 0x0102
	GREEN_BIAS = 0x0D19
	AUX1 = 0x040A
	LIGHT3 = 0x4003
	LIST_INDEX = 0x0B33
	AUX0 = 0x0409
	LIGHT2 = 0x4002
	MAP2_INDEX = 0x0DB1
	AUX3 = 0x040C
	LIGHT1 = 0x4001
	AUX2 = 0x040B
	LIGHT0 = 0x4000
	INDEX_ARRAY = 0x8077
	LIGHT7 = 0x4007
	LIGHT6 = 0x4006
	LIGHT5 = 0x4005
	LIGHT4 = 0x4004
	POLYGON_TOKEN = 0x0703
	PIXEL_MAP_I_TO_I_SIZE = 0x0CB0
	CURRENT_COLOR = 0x0B00
	LINE_STIPPLE = 0x0B24
	FOG = 0x0B60
	EVAL_BIT = 0x00010000
	FOG_DENSITY = 0x0B62
	LINE_STIPPLE_REPEAT = 0x0B26
	ALL_ATTRIB_BITS = 0xFFFFFFFF
	MAP2_TEXTURE_COORD_4 = 0x0DB6
	VERTEX_ARRAY_SIZE = 0x807A
	MAX_ATTRIB_STACK_DEPTH = 0x0D35
	MAP2_TEXTURE_COORD_1 = 0x0DB3
	MAP2_TEXTURE_COORD_2 = 0x0DB4
	HINT_BIT = 0x00008000
	MAP2_TEXTURE_COORD_3 = 0x0DB5
	MAP2_GRID_SEGMENTS = 0x0DD3
	RGBA_MODE = 0x0C31
	LUMINANCE16_ALPHA16 = 0x8048
	LUMINANCE6_ALPHA2 = 0x8044
	DEPTH_SCALE = 0x0D1E
	ADD = 0x0104
	LIST_BIT = 0x00020000
	SELECTION_BUFFER_POINTER = 0x0DF3
	RED_BITS = 0x0D52
	MAP2_NORMAL = 0x0DB2
	ACCUM_ALPHA_BITS = 0x0D5B
	T2F_C4UB_V3F = 0x2A29
	BLUE_SCALE = 0x0D1A
	POINT_TOKEN = 0x0701
	MAX_NAME_STACK_DEPTH = 0x0D37
	SPHERE_MAP = 0x2402
	C4F_N3F_V3F = 0x2A26
	MAX_CLIENT_ATTRIB_STACK_DEPTH = 0x0D3B
	C4UB_V3F = 0x2A23
	PIXEL_MAP_I_TO_A_SIZE = 0x0CB5
	LOGIC_OP = 0x0BF1
	OBJECT_LINEAR = 0x2401
	LINE_RESET_TOKEN = 0x0707
	GREEN_BITS = 0x0D53
	EDGE_FLAG_ARRAY_STRIDE = 0x808C
	CURRENT_BIT = 0x00000001
	FOG_INDEX = 0x0B61
)
// VERSION_1_2_DEPRECATED
const (
	SEPARATE_SPECULAR_COLOR = 0x81FA
	ALIASED_POINT_SIZE_RANGE = 0x846D
	RESCALE_NORMAL = 0x803A
	LIGHT_MODEL_COLOR_CONTROL = 0x81F8
	SINGLE_COLOR = 0x81F9
)
// VERSION_2_0_DEPRECATED
const (
	MAX_TEXTURE_COORDS = 0x8871
	POINT_SPRITE = 0x8861
	VERTEX_PROGRAM_TWO_SIDE = 0x8643
	COORD_REPLACE = 0x8862
)
// VERSION_1_2_DEPRECATED
func ColorTable(target Enum, internalformat Enum, width Sizei, format Enum, type_ Enum, table Pointer)  {
	C.goglColorTable((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(table))
}
func ColorTableParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglColorTableParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func ColorTableParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglColorTableParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func CopyColorTable(target Enum, internalformat Enum, x Int, y Int, width Sizei)  {
	C.goglCopyColorTable((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
func GetColorTable(target Enum, format Enum, type_ Enum, table Pointer)  {
	C.goglGetColorTable((C.GLenum)(target), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(table))
}
func GetColorTableParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetColorTableParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetColorTableParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetColorTableParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func ColorSubTable(target Enum, start Sizei, count Sizei, format Enum, type_ Enum, data Pointer)  {
	C.goglColorSubTable((C.GLenum)(target), (C.GLsizei)(start), (C.GLsizei)(count), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(data))
}
func CopyColorSubTable(target Enum, start Sizei, x Int, y Int, width Sizei)  {
	C.goglCopyColorSubTable((C.GLenum)(target), (C.GLsizei)(start), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
func ConvolutionFilter1D(target Enum, internalformat Enum, width Sizei, format Enum, type_ Enum, image Pointer)  {
	C.goglConvolutionFilter1D((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(image))
}
func ConvolutionFilter2D(target Enum, internalformat Enum, width Sizei, height Sizei, format Enum, type_ Enum, image Pointer)  {
	C.goglConvolutionFilter2D((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(image))
}
func ConvolutionParameterf(target Enum, pname Enum, params Float)  {
	C.goglConvolutionParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(params))
}
func ConvolutionParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglConvolutionParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func ConvolutionParameteri(target Enum, pname Enum, params Int)  {
	C.goglConvolutionParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(params))
}
func ConvolutionParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglConvolutionParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func CopyConvolutionFilter1D(target Enum, internalformat Enum, x Int, y Int, width Sizei)  {
	C.goglCopyConvolutionFilter1D((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
func CopyConvolutionFilter2D(target Enum, internalformat Enum, x Int, y Int, width Sizei, height Sizei)  {
	C.goglCopyConvolutionFilter2D((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func GetConvolutionFilter(target Enum, format Enum, type_ Enum, image Pointer)  {
	C.goglGetConvolutionFilter((C.GLenum)(target), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(image))
}
func GetConvolutionParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetConvolutionParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetConvolutionParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetConvolutionParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetSeparableFilter(target Enum, format Enum, type_ Enum, row Pointer, column Pointer, span Pointer)  {
	C.goglGetSeparableFilter((C.GLenum)(target), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(row), (unsafe.Pointer)(column), (unsafe.Pointer)(span))
}
func SeparableFilter2D(target Enum, internalformat Enum, width Sizei, height Sizei, format Enum, type_ Enum, row Pointer, column Pointer)  {
	C.goglSeparableFilter2D((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(row), (unsafe.Pointer)(column))
}
func GetHistogram(target Enum, reset Boolean, format Enum, type_ Enum, values Pointer)  {
	C.goglGetHistogram((C.GLenum)(target), (C.GLboolean)(reset), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(values))
}
func GetHistogramParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetHistogramParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetHistogramParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetHistogramParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetMinmax(target Enum, reset Boolean, format Enum, type_ Enum, values Pointer)  {
	C.goglGetMinmax((C.GLenum)(target), (C.GLboolean)(reset), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(values))
}
func GetMinmaxParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetMinmaxParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetMinmaxParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetMinmaxParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func Histogram(target Enum, width Sizei, internalformat Enum, sink Boolean)  {
	C.goglHistogram((C.GLenum)(target), (C.GLsizei)(width), (C.GLenum)(internalformat), (C.GLboolean)(sink))
}
func Minmax(target Enum, internalformat Enum, sink Boolean)  {
	C.goglMinmax((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLboolean)(sink))
}
func ResetHistogram(target Enum)  {
	C.goglResetHistogram((C.GLenum)(target))
}
func ResetMinmax(target Enum)  {
	C.goglResetMinmax((C.GLenum)(target))
}
// VERSION_2_1
func UniformMatrix2x3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
func UniformMatrix3x2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
func UniformMatrix2x4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
func UniformMatrix4x2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
func UniformMatrix3x4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
func UniformMatrix4x3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
// VERSION_2_0
func BlendEquationSeparate(modeRGB Enum, modeAlpha Enum)  {
	C.goglBlendEquationSeparate((C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func DrawBuffers(n Sizei, bufs *Enum)  {
	C.goglDrawBuffers((C.GLsizei)(n), (*C.GLenum)(bufs))
}
func StencilOpSeparate(face Enum, sfail Enum, dpfail Enum, dppass Enum)  {
	C.goglStencilOpSeparate((C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}
func StencilFuncSeparate(face Enum, func_ Enum, ref Int, mask Uint)  {
	C.goglStencilFuncSeparate((C.GLenum)(face), (C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
func StencilMaskSeparate(face Enum, mask Uint)  {
	C.goglStencilMaskSeparate((C.GLenum)(face), (C.GLuint)(mask))
}
func AttachShader(program Uint, shader Uint)  {
	C.goglAttachShader((C.GLuint)(program), (C.GLuint)(shader))
}
func BindAttribLocation(program Uint, index Uint, name *Char)  {
	C.goglBindAttribLocation((C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(name))
}
func CompileShader(shader Uint)  {
	C.goglCompileShader((C.GLuint)(shader))
}
func CreateProgram() Uint {
	return (Uint)(C.goglCreateProgram())
}
func CreateShader(type_ Enum) Uint {
	return (Uint)(C.goglCreateShader((C.GLenum)(type_)))
}
func DeleteProgram(program Uint)  {
	C.goglDeleteProgram((C.GLuint)(program))
}
func DeleteShader(shader Uint)  {
	C.goglDeleteShader((C.GLuint)(shader))
}
func DetachShader(program Uint, shader Uint)  {
	C.goglDetachShader((C.GLuint)(program), (C.GLuint)(shader))
}
func DisableVertexAttribArray(index Uint)  {
	C.goglDisableVertexAttribArray((C.GLuint)(index))
}
func EnableVertexAttribArray(index Uint)  {
	C.goglEnableVertexAttribArray((C.GLuint)(index))
}
func GetActiveAttrib(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)  {
	C.goglGetActiveAttrib((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
func GetActiveUniform(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)  {
	C.goglGetActiveUniform((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
func GetAttachedShaders(program Uint, maxCount Sizei, count *Sizei, obj *Uint)  {
	C.goglGetAttachedShaders((C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(count), (*C.GLuint)(obj))
}
func GetAttribLocation(program Uint, name *Char) Int {
	return (Int)(C.goglGetAttribLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
func GetProgramiv(program Uint, pname Enum, params *Int)  {
	C.goglGetProgramiv((C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetProgramInfoLog(program Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.goglGetProgramInfoLog((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
func GetShaderiv(shader Uint, pname Enum, params *Int)  {
	C.goglGetShaderiv((C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetShaderInfoLog(shader Uint, bufSize Sizei, length *Sizei, infoLog *Char)  {
	C.goglGetShaderInfoLog((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}
func GetShaderSource(shader Uint, bufSize Sizei, length *Sizei, source *Char)  {
	C.goglGetShaderSource((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(source))
}
func GetUniformLocation(program Uint, name *Char) Int {
	return (Int)(C.goglGetUniformLocation((C.GLuint)(program), (*C.GLchar)(name)))
}
func GetUniformfv(program Uint, location Int, params *Float)  {
	C.goglGetUniformfv((C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(params))
}
func GetUniformiv(program Uint, location Int, params *Int)  {
	C.goglGetUniformiv((C.GLuint)(program), (C.GLint)(location), (*C.GLint)(params))
}
func GetVertexAttribdv(index Uint, pname Enum, params *Double)  {
	C.goglGetVertexAttribdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}
func GetVertexAttribfv(index Uint, pname Enum, params *Float)  {
	C.goglGetVertexAttribfv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetVertexAttribiv(index Uint, pname Enum, params *Int)  {
	C.goglGetVertexAttribiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetVertexAttribPointerv(index Uint, pname Enum, pointer *Pointer)  {
	C.goglGetVertexAttribPointerv((C.GLuint)(index), (C.GLenum)(pname), (*unsafe.Pointer)(pointer))
}
func IsProgram(program Uint) Boolean {
	return (Boolean)(C.goglIsProgram((C.GLuint)(program)))
}
func IsShader(shader Uint) Boolean {
	return (Boolean)(C.goglIsShader((C.GLuint)(shader)))
}
func LinkProgram(program Uint)  {
	C.goglLinkProgram((C.GLuint)(program))
}
func ShaderSource(shader Uint, count Sizei, string_ **Char, length *Int)  {
	C.goglShaderSource((C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(string_)), (*C.GLint)(length))
}
func UseProgram(program Uint)  {
	C.goglUseProgram((C.GLuint)(program))
}
func Uniform1f(location Int, v0 Float)  {
	C.goglUniform1f((C.GLint)(location), (C.GLfloat)(v0))
}
func Uniform2f(location Int, v0 Float, v1 Float)  {
	C.goglUniform2f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}
func Uniform3f(location Int, v0 Float, v1 Float, v2 Float)  {
	C.goglUniform3f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
func Uniform4f(location Int, v0 Float, v1 Float, v2 Float, v3 Float)  {
	C.goglUniform4f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}
func Uniform1i(location Int, v0 Int)  {
	C.goglUniform1i((C.GLint)(location), (C.GLint)(v0))
}
func Uniform2i(location Int, v0 Int, v1 Int)  {
	C.goglUniform2i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
func Uniform3i(location Int, v0 Int, v1 Int, v2 Int)  {
	C.goglUniform3i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
func Uniform4i(location Int, v0 Int, v1 Int, v2 Int, v3 Int)  {
	C.goglUniform4i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
func Uniform1fv(location Int, count Sizei, value *Float)  {
	C.goglUniform1fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
func Uniform2fv(location Int, count Sizei, value *Float)  {
	C.goglUniform2fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
func Uniform3fv(location Int, count Sizei, value *Float)  {
	C.goglUniform3fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
func Uniform4fv(location Int, count Sizei, value *Float)  {
	C.goglUniform4fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}
func Uniform1iv(location Int, count Sizei, value *Int)  {
	C.goglUniform1iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
func Uniform2iv(location Int, count Sizei, value *Int)  {
	C.goglUniform2iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
func Uniform3iv(location Int, count Sizei, value *Int)  {
	C.goglUniform3iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
func Uniform4iv(location Int, count Sizei, value *Int)  {
	C.goglUniform4iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}
func UniformMatrix2fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
func UniformMatrix3fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
func UniformMatrix4fv(location Int, count Sizei, transpose Boolean, value *Float)  {
	C.goglUniformMatrix4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}
func ValidateProgram(program Uint)  {
	C.goglValidateProgram((C.GLuint)(program))
}
func VertexAttrib1d(index Uint, x Double)  {
	C.goglVertexAttrib1d((C.GLuint)(index), (C.GLdouble)(x))
}
func VertexAttrib1dv(index Uint, v *Double)  {
	C.goglVertexAttrib1dv((C.GLuint)(index), (*C.GLdouble)(v))
}
func VertexAttrib1f(index Uint, x Float)  {
	C.goglVertexAttrib1f((C.GLuint)(index), (C.GLfloat)(x))
}
func VertexAttrib1fv(index Uint, v *Float)  {
	C.goglVertexAttrib1fv((C.GLuint)(index), (*C.GLfloat)(v))
}
func VertexAttrib1s(index Uint, x Short)  {
	C.goglVertexAttrib1s((C.GLuint)(index), (C.GLshort)(x))
}
func VertexAttrib1sv(index Uint, v *Short)  {
	C.goglVertexAttrib1sv((C.GLuint)(index), (*C.GLshort)(v))
}
func VertexAttrib2d(index Uint, x Double, y Double)  {
	C.goglVertexAttrib2d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
func VertexAttrib2dv(index Uint, v *Double)  {
	C.goglVertexAttrib2dv((C.GLuint)(index), (*C.GLdouble)(v))
}
func VertexAttrib2f(index Uint, x Float, y Float)  {
	C.goglVertexAttrib2f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y))
}
func VertexAttrib2fv(index Uint, v *Float)  {
	C.goglVertexAttrib2fv((C.GLuint)(index), (*C.GLfloat)(v))
}
func VertexAttrib2s(index Uint, x Short, y Short)  {
	C.goglVertexAttrib2s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y))
}
func VertexAttrib2sv(index Uint, v *Short)  {
	C.goglVertexAttrib2sv((C.GLuint)(index), (*C.GLshort)(v))
}
func VertexAttrib3d(index Uint, x Double, y Double, z Double)  {
	C.goglVertexAttrib3d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func VertexAttrib3dv(index Uint, v *Double)  {
	C.goglVertexAttrib3dv((C.GLuint)(index), (*C.GLdouble)(v))
}
func VertexAttrib3f(index Uint, x Float, y Float, z Float)  {
	C.goglVertexAttrib3f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func VertexAttrib3fv(index Uint, v *Float)  {
	C.goglVertexAttrib3fv((C.GLuint)(index), (*C.GLfloat)(v))
}
func VertexAttrib3s(index Uint, x Short, y Short, z Short)  {
	C.goglVertexAttrib3s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func VertexAttrib3sv(index Uint, v *Short)  {
	C.goglVertexAttrib3sv((C.GLuint)(index), (*C.GLshort)(v))
}
func VertexAttrib4Nbv(index Uint, v *Byte)  {
	C.goglVertexAttrib4Nbv((C.GLuint)(index), (*C.GLbyte)(v))
}
func VertexAttrib4Niv(index Uint, v *Int)  {
	C.goglVertexAttrib4Niv((C.GLuint)(index), (*C.GLint)(v))
}
func VertexAttrib4Nsv(index Uint, v *Short)  {
	C.goglVertexAttrib4Nsv((C.GLuint)(index), (*C.GLshort)(v))
}
func VertexAttrib4Nub(index Uint, x Ubyte, y Ubyte, z Ubyte, w Ubyte)  {
	C.goglVertexAttrib4Nub((C.GLuint)(index), (C.GLubyte)(x), (C.GLubyte)(y), (C.GLubyte)(z), (C.GLubyte)(w))
}
func VertexAttrib4Nubv(index Uint, v *Ubyte)  {
	C.goglVertexAttrib4Nubv((C.GLuint)(index), (*C.GLubyte)(v))
}
func VertexAttrib4Nuiv(index Uint, v *Uint)  {
	C.goglVertexAttrib4Nuiv((C.GLuint)(index), (*C.GLuint)(v))
}
func VertexAttrib4Nusv(index Uint, v *Ushort)  {
	C.goglVertexAttrib4Nusv((C.GLuint)(index), (*C.GLushort)(v))
}
func VertexAttrib4bv(index Uint, v *Byte)  {
	C.goglVertexAttrib4bv((C.GLuint)(index), (*C.GLbyte)(v))
}
func VertexAttrib4d(index Uint, x Double, y Double, z Double, w Double)  {
	C.goglVertexAttrib4d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func VertexAttrib4dv(index Uint, v *Double)  {
	C.goglVertexAttrib4dv((C.GLuint)(index), (*C.GLdouble)(v))
}
func VertexAttrib4f(index Uint, x Float, y Float, z Float, w Float)  {
	C.goglVertexAttrib4f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func VertexAttrib4fv(index Uint, v *Float)  {
	C.goglVertexAttrib4fv((C.GLuint)(index), (*C.GLfloat)(v))
}
func VertexAttrib4iv(index Uint, v *Int)  {
	C.goglVertexAttrib4iv((C.GLuint)(index), (*C.GLint)(v))
}
func VertexAttrib4s(index Uint, x Short, y Short, z Short, w Short)  {
	C.goglVertexAttrib4s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func VertexAttrib4sv(index Uint, v *Short)  {
	C.goglVertexAttrib4sv((C.GLuint)(index), (*C.GLshort)(v))
}
func VertexAttrib4ubv(index Uint, v *Ubyte)  {
	C.goglVertexAttrib4ubv((C.GLuint)(index), (*C.GLubyte)(v))
}
func VertexAttrib4uiv(index Uint, v *Uint)  {
	C.goglVertexAttrib4uiv((C.GLuint)(index), (*C.GLuint)(v))
}
func VertexAttrib4usv(index Uint, v *Ushort)  {
	C.goglVertexAttrib4usv((C.GLuint)(index), (*C.GLushort)(v))
}
func VertexAttribPointer(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, pointer Pointer)  {
	C.goglVertexAttribPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
// VERSION_1_3_DEPRECATED
func ClientActiveTexture(texture Enum)  {
	C.goglClientActiveTexture((C.GLenum)(texture))
}
func MultiTexCoord1d(target Enum, s Double)  {
	C.goglMultiTexCoord1d((C.GLenum)(target), (C.GLdouble)(s))
}
func MultiTexCoord1dv(target Enum, v *Double)  {
	C.goglMultiTexCoord1dv((C.GLenum)(target), (*C.GLdouble)(v))
}
func MultiTexCoord1f(target Enum, s Float)  {
	C.goglMultiTexCoord1f((C.GLenum)(target), (C.GLfloat)(s))
}
func MultiTexCoord1fv(target Enum, v *Float)  {
	C.goglMultiTexCoord1fv((C.GLenum)(target), (*C.GLfloat)(v))
}
func MultiTexCoord1i(target Enum, s Int)  {
	C.goglMultiTexCoord1i((C.GLenum)(target), (C.GLint)(s))
}
func MultiTexCoord1iv(target Enum, v *Int)  {
	C.goglMultiTexCoord1iv((C.GLenum)(target), (*C.GLint)(v))
}
func MultiTexCoord1s(target Enum, s Short)  {
	C.goglMultiTexCoord1s((C.GLenum)(target), (C.GLshort)(s))
}
func MultiTexCoord1sv(target Enum, v *Short)  {
	C.goglMultiTexCoord1sv((C.GLenum)(target), (*C.GLshort)(v))
}
func MultiTexCoord2d(target Enum, s Double, t Double)  {
	C.goglMultiTexCoord2d((C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t))
}
func MultiTexCoord2dv(target Enum, v *Double)  {
	C.goglMultiTexCoord2dv((C.GLenum)(target), (*C.GLdouble)(v))
}
func MultiTexCoord2f(target Enum, s Float, t Float)  {
	C.goglMultiTexCoord2f((C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t))
}
func MultiTexCoord2fv(target Enum, v *Float)  {
	C.goglMultiTexCoord2fv((C.GLenum)(target), (*C.GLfloat)(v))
}
func MultiTexCoord2i(target Enum, s Int, t Int)  {
	C.goglMultiTexCoord2i((C.GLenum)(target), (C.GLint)(s), (C.GLint)(t))
}
func MultiTexCoord2iv(target Enum, v *Int)  {
	C.goglMultiTexCoord2iv((C.GLenum)(target), (*C.GLint)(v))
}
func MultiTexCoord2s(target Enum, s Short, t Short)  {
	C.goglMultiTexCoord2s((C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t))
}
func MultiTexCoord2sv(target Enum, v *Short)  {
	C.goglMultiTexCoord2sv((C.GLenum)(target), (*C.GLshort)(v))
}
func MultiTexCoord3d(target Enum, s Double, t Double, r Double)  {
	C.goglMultiTexCoord3d((C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r))
}
func MultiTexCoord3dv(target Enum, v *Double)  {
	C.goglMultiTexCoord3dv((C.GLenum)(target), (*C.GLdouble)(v))
}
func MultiTexCoord3f(target Enum, s Float, t Float, r Float)  {
	C.goglMultiTexCoord3f((C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r))
}
func MultiTexCoord3fv(target Enum, v *Float)  {
	C.goglMultiTexCoord3fv((C.GLenum)(target), (*C.GLfloat)(v))
}
func MultiTexCoord3i(target Enum, s Int, t Int, r Int)  {
	C.goglMultiTexCoord3i((C.GLenum)(target), (C.GLint)(s), (C.GLint)(t), (C.GLint)(r))
}
func MultiTexCoord3iv(target Enum, v *Int)  {
	C.goglMultiTexCoord3iv((C.GLenum)(target), (*C.GLint)(v))
}
func MultiTexCoord3s(target Enum, s Short, t Short, r Short)  {
	C.goglMultiTexCoord3s((C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r))
}
func MultiTexCoord3sv(target Enum, v *Short)  {
	C.goglMultiTexCoord3sv((C.GLenum)(target), (*C.GLshort)(v))
}
func MultiTexCoord4d(target Enum, s Double, t Double, r Double, q Double)  {
	C.goglMultiTexCoord4d((C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r), (C.GLdouble)(q))
}
func MultiTexCoord4dv(target Enum, v *Double)  {
	C.goglMultiTexCoord4dv((C.GLenum)(target), (*C.GLdouble)(v))
}
func MultiTexCoord4f(target Enum, s Float, t Float, r Float, q Float)  {
	C.goglMultiTexCoord4f((C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(q))
}
func MultiTexCoord4fv(target Enum, v *Float)  {
	C.goglMultiTexCoord4fv((C.GLenum)(target), (*C.GLfloat)(v))
}
func MultiTexCoord4i(target Enum, s Int, t Int, r Int, q Int)  {
	C.goglMultiTexCoord4i((C.GLenum)(target), (C.GLint)(s), (C.GLint)(t), (C.GLint)(r), (C.GLint)(q))
}
func MultiTexCoord4iv(target Enum, v *Int)  {
	C.goglMultiTexCoord4iv((C.GLenum)(target), (*C.GLint)(v))
}
func MultiTexCoord4s(target Enum, s Short, t Short, r Short, q Short)  {
	C.goglMultiTexCoord4s((C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r), (C.GLshort)(q))
}
func MultiTexCoord4sv(target Enum, v *Short)  {
	C.goglMultiTexCoord4sv((C.GLenum)(target), (*C.GLshort)(v))
}
func LoadTransposeMatrixf(m *Float)  {
	C.goglLoadTransposeMatrixf((*C.GLfloat)(m))
}
func LoadTransposeMatrixd(m *Double)  {
	C.goglLoadTransposeMatrixd((*C.GLdouble)(m))
}
func MultTransposeMatrixf(m *Float)  {
	C.goglMultTransposeMatrixf((*C.GLfloat)(m))
}
func MultTransposeMatrixd(m *Double)  {
	C.goglMultTransposeMatrixd((*C.GLdouble)(m))
}
// VERSION_1_4
func BlendFuncSeparate(sfactorRGB Enum, dfactorRGB Enum, sfactorAlpha Enum, dfactorAlpha Enum)  {
	C.goglBlendFuncSeparate((C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}
func MultiDrawArrays(mode Enum, first *Int, count *Sizei, primcount Sizei)  {
	C.goglMultiDrawArrays((C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(primcount))
}
func MultiDrawElements(mode Enum, count *Sizei, type_ Enum, indices *Pointer, primcount Sizei)  {
	C.goglMultiDrawElements((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(primcount))
}
func PointParameterf(pname Enum, param Float)  {
	C.goglPointParameterf((C.GLenum)(pname), (C.GLfloat)(param))
}
func PointParameterfv(pname Enum, params *Float)  {
	C.goglPointParameterfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
func PointParameteri(pname Enum, param Int)  {
	C.goglPointParameteri((C.GLenum)(pname), (C.GLint)(param))
}
func PointParameteriv(pname Enum, params *Int)  {
	C.goglPointParameteriv((C.GLenum)(pname), (*C.GLint)(params))
}
// VERSION_1_5
func GenQueries(n Sizei, ids *Uint)  {
	C.goglGenQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}
func DeleteQueries(n Sizei, ids *Uint)  {
	C.goglDeleteQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}
func IsQuery(id Uint) Boolean {
	return (Boolean)(C.goglIsQuery((C.GLuint)(id)))
}
func BeginQuery(target Enum, id Uint)  {
	C.goglBeginQuery((C.GLenum)(target), (C.GLuint)(id))
}
func EndQuery(target Enum)  {
	C.goglEndQuery((C.GLenum)(target))
}
func GetQueryiv(target Enum, pname Enum, params *Int)  {
	C.goglGetQueryiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetQueryObjectiv(id Uint, pname Enum, params *Int)  {
	C.goglGetQueryObjectiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetQueryObjectuiv(id Uint, pname Enum, params *Uint)  {
	C.goglGetQueryObjectuiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(params))
}
func BindBuffer(target Enum, buffer Uint)  {
	C.goglBindBuffer((C.GLenum)(target), (C.GLuint)(buffer))
}
func DeleteBuffers(n Sizei, buffers *Uint)  {
	C.goglDeleteBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}
func GenBuffers(n Sizei, buffers *Uint)  {
	C.goglGenBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}
func IsBuffer(buffer Uint) Boolean {
	return (Boolean)(C.goglIsBuffer((C.GLuint)(buffer)))
}
func BufferData(target Enum, size Sizeiptr, data Pointer, usage Enum)  {
	C.goglBufferData((C.GLenum)(target), (C.GLsizeiptr)(size), (unsafe.Pointer)(data), (C.GLenum)(usage))
}
func BufferSubData(target Enum, offset Intptr, size Sizeiptr, data Pointer)  {
	C.goglBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}
func GetBufferSubData(target Enum, offset Intptr, size Sizeiptr, data Pointer)  {
	C.goglGetBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}
func MapBuffer(target Enum, access Enum) Pointer {
	return (Pointer)(C.goglMapBuffer((C.GLenum)(target), (C.GLenum)(access)))
}
func UnmapBuffer(target Enum) Boolean {
	return (Boolean)(C.goglUnmapBuffer((C.GLenum)(target)))
}
func GetBufferParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetBufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetBufferPointerv(target Enum, pname Enum, params *Pointer)  {
	C.goglGetBufferPointerv((C.GLenum)(target), (C.GLenum)(pname), (*unsafe.Pointer)(params))
}
// VERSION_1_0
func CullFace(mode Enum)  {
	C.goglCullFace((C.GLenum)(mode))
}
func FrontFace(mode Enum)  {
	C.goglFrontFace((C.GLenum)(mode))
}
func Hint(target Enum, mode Enum)  {
	C.goglHint((C.GLenum)(target), (C.GLenum)(mode))
}
func LineWidth(width Float)  {
	C.goglLineWidth((C.GLfloat)(width))
}
func PointSize(size Float)  {
	C.goglPointSize((C.GLfloat)(size))
}
func PolygonMode(face Enum, mode Enum)  {
	C.goglPolygonMode((C.GLenum)(face), (C.GLenum)(mode))
}
func Scissor(x Int, y Int, width Sizei, height Sizei)  {
	C.goglScissor((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func TexParameterf(target Enum, pname Enum, param Float)  {
	C.goglTexParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TexParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func TexParameteri(target Enum, pname Enum, param Int)  {
	C.goglTexParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func TexParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func TexImage1D(target Enum, level Int, internalformat Int, width Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
func TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
func DrawBuffer(mode Enum)  {
	C.goglDrawBuffer((C.GLenum)(mode))
}
func Clear(mask Bitfield)  {
	C.goglClear((C.GLbitfield)(mask))
}
func ClearColor(red Clampf, green Clampf, blue Clampf, alpha Clampf)  {
	C.goglClearColor((C.GLclampf)(red), (C.GLclampf)(green), (C.GLclampf)(blue), (C.GLclampf)(alpha))
}
func ClearStencil(s Int)  {
	C.goglClearStencil((C.GLint)(s))
}
func ClearDepth(depth Clampd)  {
	C.goglClearDepth((C.GLclampd)(depth))
}
func StencilMask(mask Uint)  {
	C.goglStencilMask((C.GLuint)(mask))
}
func ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean)  {
	C.goglColorMask((C.GLboolean)(red), (C.GLboolean)(green), (C.GLboolean)(blue), (C.GLboolean)(alpha))
}
func DepthMask(flag Boolean)  {
	C.goglDepthMask((C.GLboolean)(flag))
}
func Disable(cap Enum)  {
	C.goglDisable((C.GLenum)(cap))
}
func Enable(cap Enum)  {
	C.goglEnable((C.GLenum)(cap))
}
func Finish()  {
	C.goglFinish()
}
func Flush()  {
	C.goglFlush()
}
func BlendFunc(sfactor Enum, dfactor Enum)  {
	C.goglBlendFunc((C.GLenum)(sfactor), (C.GLenum)(dfactor))
}
func LogicOp(opcode Enum)  {
	C.goglLogicOp((C.GLenum)(opcode))
}
func StencilFunc(func_ Enum, ref Int, mask Uint)  {
	C.goglStencilFunc((C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
func StencilOp(fail Enum, zfail Enum, zpass Enum)  {
	C.goglStencilOp((C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}
func DepthFunc(func_ Enum)  {
	C.goglDepthFunc((C.GLenum)(func_))
}
func PixelStoref(pname Enum, param Float)  {
	C.goglPixelStoref((C.GLenum)(pname), (C.GLfloat)(param))
}
func PixelStorei(pname Enum, param Int)  {
	C.goglPixelStorei((C.GLenum)(pname), (C.GLint)(param))
}
func ReadBuffer(mode Enum)  {
	C.goglReadBuffer((C.GLenum)(mode))
}
func ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglReadPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
func GetBooleanv(pname Enum, params *Boolean)  {
	C.goglGetBooleanv((C.GLenum)(pname), (*C.GLboolean)(params))
}
func GetDoublev(pname Enum, params *Double)  {
	C.goglGetDoublev((C.GLenum)(pname), (*C.GLdouble)(params))
}
func GetError() Enum {
	return (Enum)(C.goglGetError())
}
func GetFloatv(pname Enum, params *Float)  {
	C.goglGetFloatv((C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetIntegerv(pname Enum, params *Int)  {
	C.goglGetIntegerv((C.GLenum)(pname), (*C.GLint)(params))
}
func GetString(name Enum) *Ubyte {
	return (*Ubyte)(C.goglGetString((C.GLenum)(name)))
}
func GetTexImage(target Enum, level Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglGetTexImage((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
func GetTexParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetTexParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetTexLevelParameterfv(target Enum, level Int, pname Enum, params *Float)  {
	C.goglGetTexLevelParameterfv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetTexLevelParameteriv(target Enum, level Int, pname Enum, params *Int)  {
	C.goglGetTexLevelParameteriv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(params))
}
func IsEnabled(cap Enum) Boolean {
	return (Boolean)(C.goglIsEnabled((C.GLenum)(cap)))
}
func DepthRange(near Clampd, far Clampd)  {
	C.goglDepthRange((C.GLclampd)(near), (C.GLclampd)(far))
}
func Viewport(x Int, y Int, width Sizei, height Sizei)  {
	C.goglViewport((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_1
func DrawArrays(mode Enum, first Int, count Sizei)  {
	C.goglDrawArrays((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}
func DrawElements(mode Enum, count Sizei, type_ Enum, indices Pointer)  {
	C.goglDrawElements((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
func GetPointerv(pname Enum, params *Pointer)  {
	C.goglGetPointerv((C.GLenum)(pname), (*unsafe.Pointer)(params))
}
func PolygonOffset(factor Float, units Float)  {
	C.goglPolygonOffset((C.GLfloat)(factor), (C.GLfloat)(units))
}
func CopyTexImage1D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, border Int)  {
	C.goglCopyTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
}
func CopyTexImage2D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, height Sizei, border Int)  {
	C.goglCopyTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}
func CopyTexSubImage1D(target Enum, level Int, xoffset Int, x Int, y Int, width Sizei)  {
	C.goglCopyTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
func CopyTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, x Int, y Int, width Sizei, height Sizei)  {
	C.goglCopyTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func TexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
func TexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
func BindTexture(target Enum, texture Uint)  {
	C.goglBindTexture((C.GLenum)(target), (C.GLuint)(texture))
}
func DeleteTextures(n Sizei, textures *Uint)  {
	C.goglDeleteTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}
func GenTextures(n Sizei, textures *Uint)  {
	C.goglGenTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}
func IsTexture(texture Uint) Boolean {
	return (Boolean)(C.goglIsTexture((C.GLuint)(texture)))
}
// VERSION_1_4_DEPRECATED
func FogCoordf(coord Float)  {
	C.goglFogCoordf((C.GLfloat)(coord))
}
func FogCoordfv(coord *Float)  {
	C.goglFogCoordfv((*C.GLfloat)(coord))
}
func FogCoordd(coord Double)  {
	C.goglFogCoordd((C.GLdouble)(coord))
}
func FogCoorddv(coord *Double)  {
	C.goglFogCoorddv((*C.GLdouble)(coord))
}
func FogCoordPointer(type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglFogCoordPointer((C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func SecondaryColor3b(red Byte, green Byte, blue Byte)  {
	C.goglSecondaryColor3b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue))
}
func SecondaryColor3bv(v *Byte)  {
	C.goglSecondaryColor3bv((*C.GLbyte)(v))
}
func SecondaryColor3d(red Double, green Double, blue Double)  {
	C.goglSecondaryColor3d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue))
}
func SecondaryColor3dv(v *Double)  {
	C.goglSecondaryColor3dv((*C.GLdouble)(v))
}
func SecondaryColor3f(red Float, green Float, blue Float)  {
	C.goglSecondaryColor3f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue))
}
func SecondaryColor3fv(v *Float)  {
	C.goglSecondaryColor3fv((*C.GLfloat)(v))
}
func SecondaryColor3i(red Int, green Int, blue Int)  {
	C.goglSecondaryColor3i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue))
}
func SecondaryColor3iv(v *Int)  {
	C.goglSecondaryColor3iv((*C.GLint)(v))
}
func SecondaryColor3s(red Short, green Short, blue Short)  {
	C.goglSecondaryColor3s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue))
}
func SecondaryColor3sv(v *Short)  {
	C.goglSecondaryColor3sv((*C.GLshort)(v))
}
func SecondaryColor3ub(red Ubyte, green Ubyte, blue Ubyte)  {
	C.goglSecondaryColor3ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue))
}
func SecondaryColor3ubv(v *Ubyte)  {
	C.goglSecondaryColor3ubv((*C.GLubyte)(v))
}
func SecondaryColor3ui(red Uint, green Uint, blue Uint)  {
	C.goglSecondaryColor3ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue))
}
func SecondaryColor3uiv(v *Uint)  {
	C.goglSecondaryColor3uiv((*C.GLuint)(v))
}
func SecondaryColor3us(red Ushort, green Ushort, blue Ushort)  {
	C.goglSecondaryColor3us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue))
}
func SecondaryColor3usv(v *Ushort)  {
	C.goglSecondaryColor3usv((*C.GLushort)(v))
}
func SecondaryColorPointer(size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglSecondaryColorPointer((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func WindowPos2d(x Double, y Double)  {
	C.goglWindowPos2d((C.GLdouble)(x), (C.GLdouble)(y))
}
func WindowPos2dv(v *Double)  {
	C.goglWindowPos2dv((*C.GLdouble)(v))
}
func WindowPos2f(x Float, y Float)  {
	C.goglWindowPos2f((C.GLfloat)(x), (C.GLfloat)(y))
}
func WindowPos2fv(v *Float)  {
	C.goglWindowPos2fv((*C.GLfloat)(v))
}
func WindowPos2i(x Int, y Int)  {
	C.goglWindowPos2i((C.GLint)(x), (C.GLint)(y))
}
func WindowPos2iv(v *Int)  {
	C.goglWindowPos2iv((*C.GLint)(v))
}
func WindowPos2s(x Short, y Short)  {
	C.goglWindowPos2s((C.GLshort)(x), (C.GLshort)(y))
}
func WindowPos2sv(v *Short)  {
	C.goglWindowPos2sv((*C.GLshort)(v))
}
func WindowPos3d(x Double, y Double, z Double)  {
	C.goglWindowPos3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func WindowPos3dv(v *Double)  {
	C.goglWindowPos3dv((*C.GLdouble)(v))
}
func WindowPos3f(x Float, y Float, z Float)  {
	C.goglWindowPos3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func WindowPos3fv(v *Float)  {
	C.goglWindowPos3fv((*C.GLfloat)(v))
}
func WindowPos3i(x Int, y Int, z Int)  {
	C.goglWindowPos3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func WindowPos3iv(v *Int)  {
	C.goglWindowPos3iv((*C.GLint)(v))
}
func WindowPos3s(x Short, y Short, z Short)  {
	C.goglWindowPos3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func WindowPos3sv(v *Short)  {
	C.goglWindowPos3sv((*C.GLshort)(v))
}
// VERSION_1_2
func BlendColor(red Clampf, green Clampf, blue Clampf, alpha Clampf)  {
	C.goglBlendColor((C.GLclampf)(red), (C.GLclampf)(green), (C.GLclampf)(blue), (C.GLclampf)(alpha))
}
func BlendEquation(mode Enum)  {
	C.goglBlendEquation((C.GLenum)(mode))
}
func DrawRangeElements(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Pointer)  {
	C.goglDrawRangeElements((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}
func TexImage3D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, depth Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
func TexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
func CopyTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, x Int, y Int, width Sizei, height Sizei)  {
	C.goglCopyTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_3
func ActiveTexture(texture Enum)  {
	C.goglActiveTexture((C.GLenum)(texture))
}
func SampleCoverage(value Clampf, invert Boolean)  {
	C.goglSampleCoverage((C.GLclampf)(value), (C.GLboolean)(invert))
}
func CompressedTexImage3D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, depth Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
func CompressedTexImage2D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
func CompressedTexImage1D(target Enum, level Int, internalformat Enum, width Sizei, border Int, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
func CompressedTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
func CompressedTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
func CompressedTexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, imageSize Sizei, data Pointer)  {
	C.goglCompressedTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}
func GetCompressedTexImage(target Enum, level Int, img Pointer)  {
	C.goglGetCompressedTexImage((C.GLenum)(target), (C.GLint)(level), (unsafe.Pointer)(img))
}
// VERSION_1_1_DEPRECATED
func ArrayElement(i Int)  {
	C.goglArrayElement((C.GLint)(i))
}
func ColorPointer(size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglColorPointer((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func DisableClientState(array Enum)  {
	C.goglDisableClientState((C.GLenum)(array))
}
func EdgeFlagPointer(stride Sizei, pointer Pointer)  {
	C.goglEdgeFlagPointer((C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func EnableClientState(array Enum)  {
	C.goglEnableClientState((C.GLenum)(array))
}
func IndexPointer(type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglIndexPointer((C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func InterleavedArrays(format Enum, stride Sizei, pointer Pointer)  {
	C.goglInterleavedArrays((C.GLenum)(format), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func NormalPointer(type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglNormalPointer((C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func TexCoordPointer(size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglTexCoordPointer((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func VertexPointer(size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglVertexPointer((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func AreTexturesResident(n Sizei, textures *Uint, residences *Boolean) Boolean {
	return (Boolean)(C.goglAreTexturesResident((C.GLsizei)(n), (*C.GLuint)(textures), (*C.GLboolean)(residences)))
}
func PrioritizeTextures(n Sizei, textures *Uint, priorities *Clampf)  {
	C.goglPrioritizeTextures((C.GLsizei)(n), (*C.GLuint)(textures), (*C.GLclampf)(priorities))
}
func Indexub(c Ubyte)  {
	C.goglIndexub((C.GLubyte)(c))
}
func Indexubv(c *Ubyte)  {
	C.goglIndexubv((*C.GLubyte)(c))
}
func PopClientAttrib()  {
	C.goglPopClientAttrib()
}
func PushClientAttrib(mask Bitfield)  {
	C.goglPushClientAttrib((C.GLbitfield)(mask))
}
// VERSION_1_0_DEPRECATED
func NewList(list Uint, mode Enum)  {
	C.goglNewList((C.GLuint)(list), (C.GLenum)(mode))
}
func EndList()  {
	C.goglEndList()
}
func CallList(list Uint)  {
	C.goglCallList((C.GLuint)(list))
}
func CallLists(n Sizei, type_ Enum, lists Pointer)  {
	C.goglCallLists((C.GLsizei)(n), (C.GLenum)(type_), (unsafe.Pointer)(lists))
}
func DeleteLists(list Uint, range_ Sizei)  {
	C.goglDeleteLists((C.GLuint)(list), (C.GLsizei)(range_))
}
func GenLists(range_ Sizei) Uint {
	return (Uint)(C.goglGenLists((C.GLsizei)(range_)))
}
func ListBase(base Uint)  {
	C.goglListBase((C.GLuint)(base))
}
func Begin(mode Enum)  {
	C.goglBegin((C.GLenum)(mode))
}
func Bitmap(width Sizei, height Sizei, xorig Float, yorig Float, xmove Float, ymove Float, bitmap *Ubyte)  {
	C.goglBitmap((C.GLsizei)(width), (C.GLsizei)(height), (C.GLfloat)(xorig), (C.GLfloat)(yorig), (C.GLfloat)(xmove), (C.GLfloat)(ymove), (*C.GLubyte)(bitmap))
}
func Color3b(red Byte, green Byte, blue Byte)  {
	C.goglColor3b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue))
}
func Color3bv(v *Byte)  {
	C.goglColor3bv((*C.GLbyte)(v))
}
func Color3d(red Double, green Double, blue Double)  {
	C.goglColor3d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue))
}
func Color3dv(v *Double)  {
	C.goglColor3dv((*C.GLdouble)(v))
}
func Color3f(red Float, green Float, blue Float)  {
	C.goglColor3f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue))
}
func Color3fv(v *Float)  {
	C.goglColor3fv((*C.GLfloat)(v))
}
func Color3i(red Int, green Int, blue Int)  {
	C.goglColor3i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue))
}
func Color3iv(v *Int)  {
	C.goglColor3iv((*C.GLint)(v))
}
func Color3s(red Short, green Short, blue Short)  {
	C.goglColor3s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue))
}
func Color3sv(v *Short)  {
	C.goglColor3sv((*C.GLshort)(v))
}
func Color3ub(red Ubyte, green Ubyte, blue Ubyte)  {
	C.goglColor3ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue))
}
func Color3ubv(v *Ubyte)  {
	C.goglColor3ubv((*C.GLubyte)(v))
}
func Color3ui(red Uint, green Uint, blue Uint)  {
	C.goglColor3ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue))
}
func Color3uiv(v *Uint)  {
	C.goglColor3uiv((*C.GLuint)(v))
}
func Color3us(red Ushort, green Ushort, blue Ushort)  {
	C.goglColor3us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue))
}
func Color3usv(v *Ushort)  {
	C.goglColor3usv((*C.GLushort)(v))
}
func Color4b(red Byte, green Byte, blue Byte, alpha Byte)  {
	C.goglColor4b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue), (C.GLbyte)(alpha))
}
func Color4bv(v *Byte)  {
	C.goglColor4bv((*C.GLbyte)(v))
}
func Color4d(red Double, green Double, blue Double, alpha Double)  {
	C.goglColor4d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue), (C.GLdouble)(alpha))
}
func Color4dv(v *Double)  {
	C.goglColor4dv((*C.GLdouble)(v))
}
func Color4f(red Float, green Float, blue Float, alpha Float)  {
	C.goglColor4f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
func Color4fv(v *Float)  {
	C.goglColor4fv((*C.GLfloat)(v))
}
func Color4i(red Int, green Int, blue Int, alpha Int)  {
	C.goglColor4i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue), (C.GLint)(alpha))
}
func Color4iv(v *Int)  {
	C.goglColor4iv((*C.GLint)(v))
}
func Color4s(red Short, green Short, blue Short, alpha Short)  {
	C.goglColor4s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue), (C.GLshort)(alpha))
}
func Color4sv(v *Short)  {
	C.goglColor4sv((*C.GLshort)(v))
}
func Color4ub(red Ubyte, green Ubyte, blue Ubyte, alpha Ubyte)  {
	C.goglColor4ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue), (C.GLubyte)(alpha))
}
func Color4ubv(v *Ubyte)  {
	C.goglColor4ubv((*C.GLubyte)(v))
}
func Color4ui(red Uint, green Uint, blue Uint, alpha Uint)  {
	C.goglColor4ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue), (C.GLuint)(alpha))
}
func Color4uiv(v *Uint)  {
	C.goglColor4uiv((*C.GLuint)(v))
}
func Color4us(red Ushort, green Ushort, blue Ushort, alpha Ushort)  {
	C.goglColor4us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue), (C.GLushort)(alpha))
}
func Color4usv(v *Ushort)  {
	C.goglColor4usv((*C.GLushort)(v))
}
func EdgeFlag(flag Boolean)  {
	C.goglEdgeFlag((C.GLboolean)(flag))
}
func EdgeFlagv(flag *Boolean)  {
	C.goglEdgeFlagv((*C.GLboolean)(flag))
}
func End()  {
	C.goglEnd()
}
func Indexd(c Double)  {
	C.goglIndexd((C.GLdouble)(c))
}
func Indexdv(c *Double)  {
	C.goglIndexdv((*C.GLdouble)(c))
}
func Indexf(c Float)  {
	C.goglIndexf((C.GLfloat)(c))
}
func Indexfv(c *Float)  {
	C.goglIndexfv((*C.GLfloat)(c))
}
func Indexi(c Int)  {
	C.goglIndexi((C.GLint)(c))
}
func Indexiv(c *Int)  {
	C.goglIndexiv((*C.GLint)(c))
}
func Indexs(c Short)  {
	C.goglIndexs((C.GLshort)(c))
}
func Indexsv(c *Short)  {
	C.goglIndexsv((*C.GLshort)(c))
}
func Normal3b(nx Byte, ny Byte, nz Byte)  {
	C.goglNormal3b((C.GLbyte)(nx), (C.GLbyte)(ny), (C.GLbyte)(nz))
}
func Normal3bv(v *Byte)  {
	C.goglNormal3bv((*C.GLbyte)(v))
}
func Normal3d(nx Double, ny Double, nz Double)  {
	C.goglNormal3d((C.GLdouble)(nx), (C.GLdouble)(ny), (C.GLdouble)(nz))
}
func Normal3dv(v *Double)  {
	C.goglNormal3dv((*C.GLdouble)(v))
}
func Normal3f(nx Float, ny Float, nz Float)  {
	C.goglNormal3f((C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz))
}
func Normal3fv(v *Float)  {
	C.goglNormal3fv((*C.GLfloat)(v))
}
func Normal3i(nx Int, ny Int, nz Int)  {
	C.goglNormal3i((C.GLint)(nx), (C.GLint)(ny), (C.GLint)(nz))
}
func Normal3iv(v *Int)  {
	C.goglNormal3iv((*C.GLint)(v))
}
func Normal3s(nx Short, ny Short, nz Short)  {
	C.goglNormal3s((C.GLshort)(nx), (C.GLshort)(ny), (C.GLshort)(nz))
}
func Normal3sv(v *Short)  {
	C.goglNormal3sv((*C.GLshort)(v))
}
func RasterPos2d(x Double, y Double)  {
	C.goglRasterPos2d((C.GLdouble)(x), (C.GLdouble)(y))
}
func RasterPos2dv(v *Double)  {
	C.goglRasterPos2dv((*C.GLdouble)(v))
}
func RasterPos2f(x Float, y Float)  {
	C.goglRasterPos2f((C.GLfloat)(x), (C.GLfloat)(y))
}
func RasterPos2fv(v *Float)  {
	C.goglRasterPos2fv((*C.GLfloat)(v))
}
func RasterPos2i(x Int, y Int)  {
	C.goglRasterPos2i((C.GLint)(x), (C.GLint)(y))
}
func RasterPos2iv(v *Int)  {
	C.goglRasterPos2iv((*C.GLint)(v))
}
func RasterPos2s(x Short, y Short)  {
	C.goglRasterPos2s((C.GLshort)(x), (C.GLshort)(y))
}
func RasterPos2sv(v *Short)  {
	C.goglRasterPos2sv((*C.GLshort)(v))
}
func RasterPos3d(x Double, y Double, z Double)  {
	C.goglRasterPos3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func RasterPos3dv(v *Double)  {
	C.goglRasterPos3dv((*C.GLdouble)(v))
}
func RasterPos3f(x Float, y Float, z Float)  {
	C.goglRasterPos3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func RasterPos3fv(v *Float)  {
	C.goglRasterPos3fv((*C.GLfloat)(v))
}
func RasterPos3i(x Int, y Int, z Int)  {
	C.goglRasterPos3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func RasterPos3iv(v *Int)  {
	C.goglRasterPos3iv((*C.GLint)(v))
}
func RasterPos3s(x Short, y Short, z Short)  {
	C.goglRasterPos3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func RasterPos3sv(v *Short)  {
	C.goglRasterPos3sv((*C.GLshort)(v))
}
func RasterPos4d(x Double, y Double, z Double, w Double)  {
	C.goglRasterPos4d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func RasterPos4dv(v *Double)  {
	C.goglRasterPos4dv((*C.GLdouble)(v))
}
func RasterPos4f(x Float, y Float, z Float, w Float)  {
	C.goglRasterPos4f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func RasterPos4fv(v *Float)  {
	C.goglRasterPos4fv((*C.GLfloat)(v))
}
func RasterPos4i(x Int, y Int, z Int, w Int)  {
	C.goglRasterPos4i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func RasterPos4iv(v *Int)  {
	C.goglRasterPos4iv((*C.GLint)(v))
}
func RasterPos4s(x Short, y Short, z Short, w Short)  {
	C.goglRasterPos4s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func RasterPos4sv(v *Short)  {
	C.goglRasterPos4sv((*C.GLshort)(v))
}
func Rectd(x1 Double, y1 Double, x2 Double, y2 Double)  {
	C.goglRectd((C.GLdouble)(x1), (C.GLdouble)(y1), (C.GLdouble)(x2), (C.GLdouble)(y2))
}
func Rectdv(v1 *Double, v2 *Double)  {
	C.goglRectdv((*C.GLdouble)(v1), (*C.GLdouble)(v2))
}
func Rectf(x1 Float, y1 Float, x2 Float, y2 Float)  {
	C.goglRectf((C.GLfloat)(x1), (C.GLfloat)(y1), (C.GLfloat)(x2), (C.GLfloat)(y2))
}
func Rectfv(v1 *Float, v2 *Float)  {
	C.goglRectfv((*C.GLfloat)(v1), (*C.GLfloat)(v2))
}
func Recti(x1 Int, y1 Int, x2 Int, y2 Int)  {
	C.goglRecti((C.GLint)(x1), (C.GLint)(y1), (C.GLint)(x2), (C.GLint)(y2))
}
func Rectiv(v1 *Int, v2 *Int)  {
	C.goglRectiv((*C.GLint)(v1), (*C.GLint)(v2))
}
func Rects(x1 Short, y1 Short, x2 Short, y2 Short)  {
	C.goglRects((C.GLshort)(x1), (C.GLshort)(y1), (C.GLshort)(x2), (C.GLshort)(y2))
}
func Rectsv(v1 *Short, v2 *Short)  {
	C.goglRectsv((*C.GLshort)(v1), (*C.GLshort)(v2))
}
func TexCoord1d(s Double)  {
	C.goglTexCoord1d((C.GLdouble)(s))
}
func TexCoord1dv(v *Double)  {
	C.goglTexCoord1dv((*C.GLdouble)(v))
}
func TexCoord1f(s Float)  {
	C.goglTexCoord1f((C.GLfloat)(s))
}
func TexCoord1fv(v *Float)  {
	C.goglTexCoord1fv((*C.GLfloat)(v))
}
func TexCoord1i(s Int)  {
	C.goglTexCoord1i((C.GLint)(s))
}
func TexCoord1iv(v *Int)  {
	C.goglTexCoord1iv((*C.GLint)(v))
}
func TexCoord1s(s Short)  {
	C.goglTexCoord1s((C.GLshort)(s))
}
func TexCoord1sv(v *Short)  {
	C.goglTexCoord1sv((*C.GLshort)(v))
}
func TexCoord2d(s Double, t Double)  {
	C.goglTexCoord2d((C.GLdouble)(s), (C.GLdouble)(t))
}
func TexCoord2dv(v *Double)  {
	C.goglTexCoord2dv((*C.GLdouble)(v))
}
func TexCoord2f(s Float, t Float)  {
	C.goglTexCoord2f((C.GLfloat)(s), (C.GLfloat)(t))
}
func TexCoord2fv(v *Float)  {
	C.goglTexCoord2fv((*C.GLfloat)(v))
}
func TexCoord2i(s Int, t Int)  {
	C.goglTexCoord2i((C.GLint)(s), (C.GLint)(t))
}
func TexCoord2iv(v *Int)  {
	C.goglTexCoord2iv((*C.GLint)(v))
}
func TexCoord2s(s Short, t Short)  {
	C.goglTexCoord2s((C.GLshort)(s), (C.GLshort)(t))
}
func TexCoord2sv(v *Short)  {
	C.goglTexCoord2sv((*C.GLshort)(v))
}
func TexCoord3d(s Double, t Double, r Double)  {
	C.goglTexCoord3d((C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r))
}
func TexCoord3dv(v *Double)  {
	C.goglTexCoord3dv((*C.GLdouble)(v))
}
func TexCoord3f(s Float, t Float, r Float)  {
	C.goglTexCoord3f((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r))
}
func TexCoord3fv(v *Float)  {
	C.goglTexCoord3fv((*C.GLfloat)(v))
}
func TexCoord3i(s Int, t Int, r Int)  {
	C.goglTexCoord3i((C.GLint)(s), (C.GLint)(t), (C.GLint)(r))
}
func TexCoord3iv(v *Int)  {
	C.goglTexCoord3iv((*C.GLint)(v))
}
func TexCoord3s(s Short, t Short, r Short)  {
	C.goglTexCoord3s((C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r))
}
func TexCoord3sv(v *Short)  {
	C.goglTexCoord3sv((*C.GLshort)(v))
}
func TexCoord4d(s Double, t Double, r Double, q Double)  {
	C.goglTexCoord4d((C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r), (C.GLdouble)(q))
}
func TexCoord4dv(v *Double)  {
	C.goglTexCoord4dv((*C.GLdouble)(v))
}
func TexCoord4f(s Float, t Float, r Float, q Float)  {
	C.goglTexCoord4f((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(q))
}
func TexCoord4fv(v *Float)  {
	C.goglTexCoord4fv((*C.GLfloat)(v))
}
func TexCoord4i(s Int, t Int, r Int, q Int)  {
	C.goglTexCoord4i((C.GLint)(s), (C.GLint)(t), (C.GLint)(r), (C.GLint)(q))
}
func TexCoord4iv(v *Int)  {
	C.goglTexCoord4iv((*C.GLint)(v))
}
func TexCoord4s(s Short, t Short, r Short, q Short)  {
	C.goglTexCoord4s((C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r), (C.GLshort)(q))
}
func TexCoord4sv(v *Short)  {
	C.goglTexCoord4sv((*C.GLshort)(v))
}
func Vertex2d(x Double, y Double)  {
	C.goglVertex2d((C.GLdouble)(x), (C.GLdouble)(y))
}
func Vertex2dv(v *Double)  {
	C.goglVertex2dv((*C.GLdouble)(v))
}
func Vertex2f(x Float, y Float)  {
	C.goglVertex2f((C.GLfloat)(x), (C.GLfloat)(y))
}
func Vertex2fv(v *Float)  {
	C.goglVertex2fv((*C.GLfloat)(v))
}
func Vertex2i(x Int, y Int)  {
	C.goglVertex2i((C.GLint)(x), (C.GLint)(y))
}
func Vertex2iv(v *Int)  {
	C.goglVertex2iv((*C.GLint)(v))
}
func Vertex2s(x Short, y Short)  {
	C.goglVertex2s((C.GLshort)(x), (C.GLshort)(y))
}
func Vertex2sv(v *Short)  {
	C.goglVertex2sv((*C.GLshort)(v))
}
func Vertex3d(x Double, y Double, z Double)  {
	C.goglVertex3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Vertex3dv(v *Double)  {
	C.goglVertex3dv((*C.GLdouble)(v))
}
func Vertex3f(x Float, y Float, z Float)  {
	C.goglVertex3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func Vertex3fv(v *Float)  {
	C.goglVertex3fv((*C.GLfloat)(v))
}
func Vertex3i(x Int, y Int, z Int)  {
	C.goglVertex3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func Vertex3iv(v *Int)  {
	C.goglVertex3iv((*C.GLint)(v))
}
func Vertex3s(x Short, y Short, z Short)  {
	C.goglVertex3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func Vertex3sv(v *Short)  {
	C.goglVertex3sv((*C.GLshort)(v))
}
func Vertex4d(x Double, y Double, z Double, w Double)  {
	C.goglVertex4d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func Vertex4dv(v *Double)  {
	C.goglVertex4dv((*C.GLdouble)(v))
}
func Vertex4f(x Float, y Float, z Float, w Float)  {
	C.goglVertex4f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func Vertex4fv(v *Float)  {
	C.goglVertex4fv((*C.GLfloat)(v))
}
func Vertex4i(x Int, y Int, z Int, w Int)  {
	C.goglVertex4i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func Vertex4iv(v *Int)  {
	C.goglVertex4iv((*C.GLint)(v))
}
func Vertex4s(x Short, y Short, z Short, w Short)  {
	C.goglVertex4s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func Vertex4sv(v *Short)  {
	C.goglVertex4sv((*C.GLshort)(v))
}
func ClipPlane(plane Enum, equation *Double)  {
	C.goglClipPlane((C.GLenum)(plane), (*C.GLdouble)(equation))
}
func ColorMaterial(face Enum, mode Enum)  {
	C.goglColorMaterial((C.GLenum)(face), (C.GLenum)(mode))
}
func Fogf(pname Enum, param Float)  {
	C.goglFogf((C.GLenum)(pname), (C.GLfloat)(param))
}
func Fogfv(pname Enum, params *Float)  {
	C.goglFogfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
func Fogi(pname Enum, param Int)  {
	C.goglFogi((C.GLenum)(pname), (C.GLint)(param))
}
func Fogiv(pname Enum, params *Int)  {
	C.goglFogiv((C.GLenum)(pname), (*C.GLint)(params))
}
func Lightf(light Enum, pname Enum, param Float)  {
	C.goglLightf((C.GLenum)(light), (C.GLenum)(pname), (C.GLfloat)(param))
}
func Lightfv(light Enum, pname Enum, params *Float)  {
	C.goglLightfv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func Lighti(light Enum, pname Enum, param Int)  {
	C.goglLighti((C.GLenum)(light), (C.GLenum)(pname), (C.GLint)(param))
}
func Lightiv(light Enum, pname Enum, params *Int)  {
	C.goglLightiv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(params))
}
func LightModelf(pname Enum, param Float)  {
	C.goglLightModelf((C.GLenum)(pname), (C.GLfloat)(param))
}
func LightModelfv(pname Enum, params *Float)  {
	C.goglLightModelfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
func LightModeli(pname Enum, param Int)  {
	C.goglLightModeli((C.GLenum)(pname), (C.GLint)(param))
}
func LightModeliv(pname Enum, params *Int)  {
	C.goglLightModeliv((C.GLenum)(pname), (*C.GLint)(params))
}
func LineStipple(factor Int, pattern Ushort)  {
	C.goglLineStipple((C.GLint)(factor), (C.GLushort)(pattern))
}
func Materialf(face Enum, pname Enum, param Float)  {
	C.goglMaterialf((C.GLenum)(face), (C.GLenum)(pname), (C.GLfloat)(param))
}
func Materialfv(face Enum, pname Enum, params *Float)  {
	C.goglMaterialfv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func Materiali(face Enum, pname Enum, param Int)  {
	C.goglMateriali((C.GLenum)(face), (C.GLenum)(pname), (C.GLint)(param))
}
func Materialiv(face Enum, pname Enum, params *Int)  {
	C.goglMaterialiv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(params))
}
func PolygonStipple(mask *Ubyte)  {
	C.goglPolygonStipple((*C.GLubyte)(mask))
}
func ShadeModel(mode Enum)  {
	C.goglShadeModel((C.GLenum)(mode))
}
func TexEnvf(target Enum, pname Enum, param Float)  {
	C.goglTexEnvf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TexEnvfv(target Enum, pname Enum, params *Float)  {
	C.goglTexEnvfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func TexEnvi(target Enum, pname Enum, param Int)  {
	C.goglTexEnvi((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func TexEnviv(target Enum, pname Enum, params *Int)  {
	C.goglTexEnviv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func TexGend(coord Enum, pname Enum, param Double)  {
	C.goglTexGend((C.GLenum)(coord), (C.GLenum)(pname), (C.GLdouble)(param))
}
func TexGendv(coord Enum, pname Enum, params *Double)  {
	C.goglTexGendv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLdouble)(params))
}
func TexGenf(coord Enum, pname Enum, param Float)  {
	C.goglTexGenf((C.GLenum)(coord), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TexGenfv(coord Enum, pname Enum, params *Float)  {
	C.goglTexGenfv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func TexGeni(coord Enum, pname Enum, param Int)  {
	C.goglTexGeni((C.GLenum)(coord), (C.GLenum)(pname), (C.GLint)(param))
}
func TexGeniv(coord Enum, pname Enum, params *Int)  {
	C.goglTexGeniv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLint)(params))
}
func FeedbackBuffer(size Sizei, type_ Enum, buffer *Float)  {
	C.goglFeedbackBuffer((C.GLsizei)(size), (C.GLenum)(type_), (*C.GLfloat)(buffer))
}
func SelectBuffer(size Sizei, buffer *Uint)  {
	C.goglSelectBuffer((C.GLsizei)(size), (*C.GLuint)(buffer))
}
func RenderMode(mode Enum) Int {
	return (Int)(C.goglRenderMode((C.GLenum)(mode)))
}
func InitNames()  {
	C.goglInitNames()
}
func LoadName(name Uint)  {
	C.goglLoadName((C.GLuint)(name))
}
func PassThrough(token Float)  {
	C.goglPassThrough((C.GLfloat)(token))
}
func PopName()  {
	C.goglPopName()
}
func PushName(name Uint)  {
	C.goglPushName((C.GLuint)(name))
}
func ClearAccum(red Float, green Float, blue Float, alpha Float)  {
	C.goglClearAccum((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
func ClearIndex(c Float)  {
	C.goglClearIndex((C.GLfloat)(c))
}
func IndexMask(mask Uint)  {
	C.goglIndexMask((C.GLuint)(mask))
}
func Accum(op Enum, value Float)  {
	C.goglAccum((C.GLenum)(op), (C.GLfloat)(value))
}
func PopAttrib()  {
	C.goglPopAttrib()
}
func PushAttrib(mask Bitfield)  {
	C.goglPushAttrib((C.GLbitfield)(mask))
}
func Map1d(target Enum, u1 Double, u2 Double, stride Int, order Int, points *Double)  {
	C.goglMap1d((C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLdouble)(points))
}
func Map1f(target Enum, u1 Float, u2 Float, stride Int, order Int, points *Float)  {
	C.goglMap1f((C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLfloat)(points))
}
func Map2d(target Enum, u1 Double, u2 Double, ustride Int, uorder Int, v1 Double, v2 Double, vstride Int, vorder Int, points *Double)  {
	C.goglMap2d((C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLdouble)(points))
}
func Map2f(target Enum, u1 Float, u2 Float, ustride Int, uorder Int, v1 Float, v2 Float, vstride Int, vorder Int, points *Float)  {
	C.goglMap2f((C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLfloat)(points))
}
func MapGrid1d(un Int, u1 Double, u2 Double)  {
	C.goglMapGrid1d((C.GLint)(un), (C.GLdouble)(u1), (C.GLdouble)(u2))
}
func MapGrid1f(un Int, u1 Float, u2 Float)  {
	C.goglMapGrid1f((C.GLint)(un), (C.GLfloat)(u1), (C.GLfloat)(u2))
}
func MapGrid2d(un Int, u1 Double, u2 Double, vn Int, v1 Double, v2 Double)  {
	C.goglMapGrid2d((C.GLint)(un), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(vn), (C.GLdouble)(v1), (C.GLdouble)(v2))
}
func MapGrid2f(un Int, u1 Float, u2 Float, vn Int, v1 Float, v2 Float)  {
	C.goglMapGrid2f((C.GLint)(un), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(vn), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
func EvalCoord1d(u Double)  {
	C.goglEvalCoord1d((C.GLdouble)(u))
}
func EvalCoord1dv(u *Double)  {
	C.goglEvalCoord1dv((*C.GLdouble)(u))
}
func EvalCoord1f(u Float)  {
	C.goglEvalCoord1f((C.GLfloat)(u))
}
func EvalCoord1fv(u *Float)  {
	C.goglEvalCoord1fv((*C.GLfloat)(u))
}
func EvalCoord2d(u Double, v Double)  {
	C.goglEvalCoord2d((C.GLdouble)(u), (C.GLdouble)(v))
}
func EvalCoord2dv(u *Double)  {
	C.goglEvalCoord2dv((*C.GLdouble)(u))
}
func EvalCoord2f(u Float, v Float)  {
	C.goglEvalCoord2f((C.GLfloat)(u), (C.GLfloat)(v))
}
func EvalCoord2fv(u *Float)  {
	C.goglEvalCoord2fv((*C.GLfloat)(u))
}
func EvalMesh1(mode Enum, i1 Int, i2 Int)  {
	C.goglEvalMesh1((C.GLenum)(mode), (C.GLint)(i1), (C.GLint)(i2))
}
func EvalPoint1(i Int)  {
	C.goglEvalPoint1((C.GLint)(i))
}
func EvalMesh2(mode Enum, i1 Int, i2 Int, j1 Int, j2 Int)  {
	C.goglEvalMesh2((C.GLenum)(mode), (C.GLint)(i1), (C.GLint)(i2), (C.GLint)(j1), (C.GLint)(j2))
}
func EvalPoint2(i Int, j Int)  {
	C.goglEvalPoint2((C.GLint)(i), (C.GLint)(j))
}
func AlphaFunc(func_ Enum, ref Clampf)  {
	C.goglAlphaFunc((C.GLenum)(func_), (C.GLclampf)(ref))
}
func PixelZoom(xfactor Float, yfactor Float)  {
	C.goglPixelZoom((C.GLfloat)(xfactor), (C.GLfloat)(yfactor))
}
func PixelTransferf(pname Enum, param Float)  {
	C.goglPixelTransferf((C.GLenum)(pname), (C.GLfloat)(param))
}
func PixelTransferi(pname Enum, param Int)  {
	C.goglPixelTransferi((C.GLenum)(pname), (C.GLint)(param))
}
func PixelMapfv(map_ Enum, mapsize Sizei, values *Float)  {
	C.goglPixelMapfv((C.GLenum)(map_), (C.GLsizei)(mapsize), (*C.GLfloat)(values))
}
func PixelMapuiv(map_ Enum, mapsize Sizei, values *Uint)  {
	C.goglPixelMapuiv((C.GLenum)(map_), (C.GLsizei)(mapsize), (*C.GLuint)(values))
}
func PixelMapusv(map_ Enum, mapsize Sizei, values *Ushort)  {
	C.goglPixelMapusv((C.GLenum)(map_), (C.GLsizei)(mapsize), (*C.GLushort)(values))
}
func CopyPixels(x Int, y Int, width Sizei, height Sizei, type_ Enum)  {
	C.goglCopyPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(type_))
}
func DrawPixels(width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglDrawPixels((C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
func GetClipPlane(plane Enum, equation *Double)  {
	C.goglGetClipPlane((C.GLenum)(plane), (*C.GLdouble)(equation))
}
func GetLightfv(light Enum, pname Enum, params *Float)  {
	C.goglGetLightfv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetLightiv(light Enum, pname Enum, params *Int)  {
	C.goglGetLightiv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetMapdv(target Enum, query Enum, v *Double)  {
	C.goglGetMapdv((C.GLenum)(target), (C.GLenum)(query), (*C.GLdouble)(v))
}
func GetMapfv(target Enum, query Enum, v *Float)  {
	C.goglGetMapfv((C.GLenum)(target), (C.GLenum)(query), (*C.GLfloat)(v))
}
func GetMapiv(target Enum, query Enum, v *Int)  {
	C.goglGetMapiv((C.GLenum)(target), (C.GLenum)(query), (*C.GLint)(v))
}
func GetMaterialfv(face Enum, pname Enum, params *Float)  {
	C.goglGetMaterialfv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetMaterialiv(face Enum, pname Enum, params *Int)  {
	C.goglGetMaterialiv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetPixelMapfv(map_ Enum, values *Float)  {
	C.goglGetPixelMapfv((C.GLenum)(map_), (*C.GLfloat)(values))
}
func GetPixelMapuiv(map_ Enum, values *Uint)  {
	C.goglGetPixelMapuiv((C.GLenum)(map_), (*C.GLuint)(values))
}
func GetPixelMapusv(map_ Enum, values *Ushort)  {
	C.goglGetPixelMapusv((C.GLenum)(map_), (*C.GLushort)(values))
}
func GetPolygonStipple(mask *Ubyte)  {
	C.goglGetPolygonStipple((*C.GLubyte)(mask))
}
func GetTexEnvfv(target Enum, pname Enum, params *Float)  {
	C.goglGetTexEnvfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetTexEnviv(target Enum, pname Enum, params *Int)  {
	C.goglGetTexEnviv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetTexGendv(coord Enum, pname Enum, params *Double)  {
	C.goglGetTexGendv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLdouble)(params))
}
func GetTexGenfv(coord Enum, pname Enum, params *Float)  {
	C.goglGetTexGenfv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetTexGeniv(coord Enum, pname Enum, params *Int)  {
	C.goglGetTexGeniv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLint)(params))
}
func IsList(list Uint) Boolean {
	return (Boolean)(C.goglIsList((C.GLuint)(list)))
}
func Frustum(left Double, right Double, bottom Double, top Double, zNear Double, zFar Double)  {
	C.goglFrustum((C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
}
func LoadIdentity()  {
	C.goglLoadIdentity()
}
func LoadMatrixf(m *Float)  {
	C.goglLoadMatrixf((*C.GLfloat)(m))
}
func LoadMatrixd(m *Double)  {
	C.goglLoadMatrixd((*C.GLdouble)(m))
}
func MatrixMode(mode Enum)  {
	C.goglMatrixMode((C.GLenum)(mode))
}
func MultMatrixf(m *Float)  {
	C.goglMultMatrixf((*C.GLfloat)(m))
}
func MultMatrixd(m *Double)  {
	C.goglMultMatrixd((*C.GLdouble)(m))
}
func Ortho(left Double, right Double, bottom Double, top Double, zNear Double, zFar Double)  {
	C.goglOrtho((C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
}
func PopMatrix()  {
	C.goglPopMatrix()
}
func PushMatrix()  {
	C.goglPushMatrix()
}
func Rotated(angle Double, x Double, y Double, z Double)  {
	C.goglRotated((C.GLdouble)(angle), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Rotatef(angle Float, x Float, y Float, z Float)  {
	C.goglRotatef((C.GLfloat)(angle), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func Scaled(x Double, y Double, z Double)  {
	C.goglScaled((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Scalef(x Float, y Float, z Float)  {
	C.goglScalef((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func Translated(x Double, y Double, z Double)  {
	C.goglTranslated((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Translatef(x Float, y Float, z Float)  {
	C.goglTranslatef((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func InitVERSION21() error {
	var ret C.int
	if ret = C.init_VERSION_2_1(); ret != 0 {
		return errors.New("unable to initialize VERSION_2_1")
	}
	return nil
}
func InitVERSION20() error {
	var ret C.int
	if ret = C.init_VERSION_2_0(); ret != 0 {
		return errors.New("unable to initialize VERSION_2_0")
	}
	return nil
}
func InitVERSION13DEPRECATED() error {
	var ret C.int
	if ret = C.init_VERSION_1_3_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_3_DEPRECATED")
	}
	return nil
}
func InitVERSION14() error {
	var ret C.int
	if ret = C.init_VERSION_1_4(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_4")
	}
	return nil
}
func InitVERSION15() error {
	var ret C.int
	if ret = C.init_VERSION_1_5(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_5")
	}
	return nil
}
func InitVERSION10() error {
	var ret C.int
	if ret = C.init_VERSION_1_0(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_0")
	}
	return nil
}
func InitVERSION11() error {
	var ret C.int
	if ret = C.init_VERSION_1_1(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_1")
	}
	return nil
}
func InitVERSION14DEPRECATED() error {
	var ret C.int
	if ret = C.init_VERSION_1_4_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_4_DEPRECATED")
	}
	return nil
}
func InitVERSION12() error {
	var ret C.int
	if ret = C.init_VERSION_1_2(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_2")
	}
	return nil
}
func InitVERSION13() error {
	var ret C.int
	if ret = C.init_VERSION_1_3(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_3")
	}
	return nil
}
func InitVERSION11DEPRECATED() error {
	var ret C.int
	if ret = C.init_VERSION_1_1_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_1_DEPRECATED")
	}
	return nil
}
func InitVERSION10DEPRECATED() error {
	var ret C.int
	if ret = C.init_VERSION_1_0_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_0_DEPRECATED")
	}
	return nil
}
func InitVERSION12DEPRECATED() error {
	var ret C.int
	if ret = C.init_VERSION_1_2_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_2_DEPRECATED")
	}
	return nil
}
func Init() error {
	var err error
	if err = InitVERSION13DEPRECATED(); err != nil {
		return err
	}
	if err = InitVERSION14(); err != nil {
		return err
	}
	if err = InitVERSION15(); err != nil {
		return err
	}
	if err = InitVERSION10(); err != nil {
		return err
	}
	if err = InitVERSION11(); err != nil {
		return err
	}
	if err = InitVERSION14DEPRECATED(); err != nil {
		return err
	}
	if err = InitVERSION12(); err != nil {
		return err
	}
	if err = InitVERSION13(); err != nil {
		return err
	}
	if err = InitVERSION11DEPRECATED(); err != nil {
		return err
	}
	if err = InitVERSION10DEPRECATED(); err != nil {
		return err
	}
	if err = InitVERSION12DEPRECATED(); err != nil {
		return err
	}
	if err = InitVERSION21(); err != nil {
		return err
	}
	if err = InitVERSION20(); err != nil {
		return err
	}
	return nil
}
// EOF