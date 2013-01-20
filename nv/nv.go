// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// NV_bindless_texture: http://www.opengl.org/registry/specs/NV/bindless_texture.txt
// 
// NV_conditional_render: http://www.opengl.org/registry/specs/NV/conditional_render.txt
// 
// NV_copy_image: http://www.opengl.org/registry/specs/NV/copy_image.txt
// 
// NV_depth_buffer_float: http://www.opengl.org/registry/specs/NV/depth_buffer_float.txt
// 
// NV_evaluators: http://www.opengl.org/registry/specs/NV/evaluators.txt
// 
// NV_explicit_multisample: http://www.opengl.org/registry/specs/NV/explicit_multisample.txt
// 
// NV_fence: http://www.opengl.org/registry/specs/NV/fence.txt
// 
// NV_fragment_program: http://www.opengl.org/registry/specs/NV/fragment_program.txt
// 
// NV_framebuffer_multisample_coverage: http://www.opengl.org/registry/specs/NV/framebuffer_multisample_coverage.txt
// 
// NV_geometry_program4: http://www.opengl.org/registry/specs/NV/geometry_program4.txt
// 
// NV_gpu_program4: http://www.opengl.org/registry/specs/NV/gpu_program4.txt
// 
// NV_gpu_program5: http://www.opengl.org/registry/specs/NV/gpu_program5.txt
// 
// NV_gpu_shader5: http://www.opengl.org/registry/specs/NV/gpu_shader5.txt
// 
// NV_half_float: http://www.opengl.org/registry/specs/NV/half_float.txt
// 
// NV_occlusion_query: http://www.opengl.org/registry/specs/NV/occlusion_query.txt
// 
// NV_parameter_buffer_object: http://www.opengl.org/registry/specs/NV/parameter_buffer_object.txt
// 
// NV_path_rendering: http://www.opengl.org/registry/specs/NV/path_rendering.txt
// 
// NV_pixel_data_range: http://www.opengl.org/registry/specs/NV/pixel_data_range.txt
// 
// NV_point_sprite: http://www.opengl.org/registry/specs/NV/point_sprite.txt
// 
// NV_present_video: http://www.opengl.org/registry/specs/NV/present_video.txt
// 
// NV_primitive_restart: http://www.opengl.org/registry/specs/NV/primitive_restart.txt
// 
// NV_register_combiners: http://www.opengl.org/registry/specs/NV/register_combiners.txt
// 
// NV_register_combiners2: http://www.opengl.org/registry/specs/NV/register_combiners2.txt
// 
// NV_shader_buffer_load: http://www.opengl.org/registry/specs/NV/shader_buffer_load.txt
// 
// NV_texture_barrier: http://www.opengl.org/registry/specs/NV/texture_barrier.txt
// 
// NV_texture_multisample: http://www.opengl.org/registry/specs/NV/texture_multisample.txt
// 
// NV_transform_feedback: http://www.opengl.org/registry/specs/NV/transform_feedback.txt
// 
// NV_transform_feedback2: http://www.opengl.org/registry/specs/NV/transform_feedback2.txt
// 
// NV_vdpau_interop: http://www.opengl.org/registry/specs/NV/vdpau_interop.txt
// 
// NV_vertex_array_range: http://www.opengl.org/registry/specs/NV/vertex_array_range.txt
// 
// NV_vertex_attrib_integer_64bit: http://www.opengl.org/registry/specs/NV/vertex_attrib_integer_64bit.txt
// 
// NV_vertex_buffer_unified_memory: http://www.opengl.org/registry/specs/NV/vertex_buffer_unified_memory.txt
// 
// NV_vertex_program: http://www.opengl.org/registry/specs/NV/vertex_program.txt
// 
// NV_vertex_program4: http://www.opengl.org/registry/specs/NV/vertex_program4.txt
// 
// NV_video_capture: http://www.opengl.org/registry/specs/NV/video_capture.txt
// 
package nv

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
// //  NV_bindless_texture
// GLuint64 (APIENTRYP ptrglGetTextureHandleNV)(GLuint texture);
// GLuint64 (APIENTRYP ptrglGetTextureSamplerHandleNV)(GLuint texture, GLuint sampler);
// void (APIENTRYP ptrglMakeTextureHandleResidentNV)(GLuint64 handle);
// void (APIENTRYP ptrglMakeTextureHandleNonResidentNV)(GLuint64 handle);
// GLuint64 (APIENTRYP ptrglGetImageHandleNV)(GLuint texture, GLint level, GLboolean layered, GLint layer, GLenum format);
// void (APIENTRYP ptrglMakeImageHandleResidentNV)(GLuint64 handle, GLenum access);
// void (APIENTRYP ptrglMakeImageHandleNonResidentNV)(GLuint64 handle);
// void (APIENTRYP ptrglUniformHandleui64NV)(GLint location, GLuint64 value);
// void (APIENTRYP ptrglUniformHandleui64vNV)(GLint location, GLsizei count, GLuint64* value);
// void (APIENTRYP ptrglProgramUniformHandleui64NV)(GLuint program, GLint location, GLuint64 value);
// void (APIENTRYP ptrglProgramUniformHandleui64vNV)(GLuint program, GLint location, GLsizei count, GLuint64* values);
// GLboolean (APIENTRYP ptrglIsTextureHandleResidentNV)(GLuint64 handle);
// GLboolean (APIENTRYP ptrglIsImageHandleResidentNV)(GLuint64 handle);
// //  NV_conditional_render
// void (APIENTRYP ptrglBeginConditionalRenderNV)(GLuint id, GLenum mode);
// void (APIENTRYP ptrglEndConditionalRenderNV)();
// //  NV_copy_image
// void (APIENTRYP ptrglCopyImageSubDataNV)(GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei width, GLsizei height, GLsizei depth);
// //  NV_depth_buffer_float
// void (APIENTRYP ptrglDepthRangedNV)(GLdouble zNear, GLdouble zFar);
// void (APIENTRYP ptrglClearDepthdNV)(GLdouble depth);
// void (APIENTRYP ptrglDepthBoundsdNV)(GLdouble zmin, GLdouble zmax);
// //  NV_evaluators
// void (APIENTRYP ptrglMapControlPointsNV)(GLenum target, GLuint index, GLenum type, GLsizei ustride, GLsizei vstride, GLint uorder, GLint vorder, GLboolean packed, GLvoid* points);
// void (APIENTRYP ptrglMapParameterivNV)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglMapParameterfvNV)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetMapControlPointsNV)(GLenum target, GLuint index, GLenum type, GLsizei ustride, GLsizei vstride, GLboolean packed, GLvoid* points);
// void (APIENTRYP ptrglGetMapParameterivNV)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetMapParameterfvNV)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetMapAttribParameterivNV)(GLenum target, GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetMapAttribParameterfvNV)(GLenum target, GLuint index, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglEvalMapsNV)(GLenum target, GLenum mode);
// //  NV_explicit_multisample
// void (APIENTRYP ptrglGetMultisamplefvNV)(GLenum pname, GLuint index, GLfloat* val);
// void (APIENTRYP ptrglSampleMaskIndexedNV)(GLuint index, GLbitfield mask);
// void (APIENTRYP ptrglTexRenderbufferNV)(GLenum target, GLuint renderbuffer);
// //  NV_fence
// void (APIENTRYP ptrglDeleteFencesNV)(GLsizei n, GLuint* fences);
// void (APIENTRYP ptrglGenFencesNV)(GLsizei n, GLuint* fences);
// GLboolean (APIENTRYP ptrglIsFenceNV)(GLuint fence);
// GLboolean (APIENTRYP ptrglTestFenceNV)(GLuint fence);
// void (APIENTRYP ptrglGetFenceivNV)(GLuint fence, GLenum pname, GLint* params);
// void (APIENTRYP ptrglFinishFenceNV)(GLuint fence);
// void (APIENTRYP ptrglSetFenceNV)(GLuint fence, GLenum condition);
// //  NV_fragment_program
// void (APIENTRYP ptrglProgramNamedParameter4fNV)(GLuint id, GLsizei len, GLubyte* name, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglProgramNamedParameter4dNV)(GLuint id, GLsizei len, GLubyte* name, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglProgramNamedParameter4fvNV)(GLuint id, GLsizei len, GLubyte* name, GLfloat* v);
// void (APIENTRYP ptrglProgramNamedParameter4dvNV)(GLuint id, GLsizei len, GLubyte* name, GLdouble* v);
// void (APIENTRYP ptrglGetProgramNamedParameterfvNV)(GLuint id, GLsizei len, GLubyte* name, GLfloat* params);
// void (APIENTRYP ptrglGetProgramNamedParameterdvNV)(GLuint id, GLsizei len, GLubyte* name, GLdouble* params);
// //  NV_framebuffer_multisample_coverage
// void (APIENTRYP ptrglRenderbufferStorageMultisampleCoverageNV)(GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLenum internalformat, GLsizei width, GLsizei height);
// //  NV_geometry_program4
// void (APIENTRYP ptrglProgramVertexLimitNV)(GLenum target, GLint limit);
// void (APIENTRYP ptrglFramebufferTextureEXT)(GLenum target, GLenum attachment, GLuint texture, GLint level);
// void (APIENTRYP ptrglFramebufferTextureLayerEXT)(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer);
// void (APIENTRYP ptrglFramebufferTextureFaceEXT)(GLenum target, GLenum attachment, GLuint texture, GLint level, GLenum face);
// //  NV_gpu_program4
// void (APIENTRYP ptrglProgramLocalParameterI4iNV)(GLenum target, GLuint index, GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglProgramLocalParameterI4ivNV)(GLenum target, GLuint index, GLint* params);
// void (APIENTRYP ptrglProgramLocalParametersI4ivNV)(GLenum target, GLuint index, GLsizei count, GLint* params);
// void (APIENTRYP ptrglProgramLocalParameterI4uiNV)(GLenum target, GLuint index, GLuint x, GLuint y, GLuint z, GLuint w);
// void (APIENTRYP ptrglProgramLocalParameterI4uivNV)(GLenum target, GLuint index, GLuint* params);
// void (APIENTRYP ptrglProgramLocalParametersI4uivNV)(GLenum target, GLuint index, GLsizei count, GLuint* params);
// void (APIENTRYP ptrglProgramEnvParameterI4iNV)(GLenum target, GLuint index, GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglProgramEnvParameterI4ivNV)(GLenum target, GLuint index, GLint* params);
// void (APIENTRYP ptrglProgramEnvParametersI4ivNV)(GLenum target, GLuint index, GLsizei count, GLint* params);
// void (APIENTRYP ptrglProgramEnvParameterI4uiNV)(GLenum target, GLuint index, GLuint x, GLuint y, GLuint z, GLuint w);
// void (APIENTRYP ptrglProgramEnvParameterI4uivNV)(GLenum target, GLuint index, GLuint* params);
// void (APIENTRYP ptrglProgramEnvParametersI4uivNV)(GLenum target, GLuint index, GLsizei count, GLuint* params);
// void (APIENTRYP ptrglGetProgramLocalParameterIivNV)(GLenum target, GLuint index, GLint* params);
// void (APIENTRYP ptrglGetProgramLocalParameterIuivNV)(GLenum target, GLuint index, GLuint* params);
// void (APIENTRYP ptrglGetProgramEnvParameterIivNV)(GLenum target, GLuint index, GLint* params);
// void (APIENTRYP ptrglGetProgramEnvParameterIuivNV)(GLenum target, GLuint index, GLuint* params);
// //  NV_gpu_program5
// void (APIENTRYP ptrglProgramSubroutineParametersuivNV)(GLenum target, GLsizei count, GLuint* params);
// void (APIENTRYP ptrglGetProgramSubroutineParameteruivNV)(GLenum target, GLuint index, GLuint* param);
// //  NV_gpu_shader5
// void (APIENTRYP ptrglUniform1i64NV)(GLint location, GLint64EXT x);
// void (APIENTRYP ptrglUniform2i64NV)(GLint location, GLint64EXT x, GLint64EXT y);
// void (APIENTRYP ptrglUniform3i64NV)(GLint location, GLint64EXT x, GLint64EXT y, GLint64EXT z);
// void (APIENTRYP ptrglUniform4i64NV)(GLint location, GLint64EXT x, GLint64EXT y, GLint64EXT z, GLint64EXT w);
// void (APIENTRYP ptrglUniform1i64vNV)(GLint location, GLsizei count, GLint64EXT* value);
// void (APIENTRYP ptrglUniform2i64vNV)(GLint location, GLsizei count, GLint64EXT* value);
// void (APIENTRYP ptrglUniform3i64vNV)(GLint location, GLsizei count, GLint64EXT* value);
// void (APIENTRYP ptrglUniform4i64vNV)(GLint location, GLsizei count, GLint64EXT* value);
// void (APIENTRYP ptrglUniform1ui64NV)(GLint location, GLuint64EXT x);
// void (APIENTRYP ptrglUniform2ui64NV)(GLint location, GLuint64EXT x, GLuint64EXT y);
// void (APIENTRYP ptrglUniform3ui64NV)(GLint location, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z);
// void (APIENTRYP ptrglUniform4ui64NV)(GLint location, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z, GLuint64EXT w);
// void (APIENTRYP ptrglUniform1ui64vNV)(GLint location, GLsizei count, GLuint64EXT* value);
// void (APIENTRYP ptrglUniform2ui64vNV)(GLint location, GLsizei count, GLuint64EXT* value);
// void (APIENTRYP ptrglUniform3ui64vNV)(GLint location, GLsizei count, GLuint64EXT* value);
// void (APIENTRYP ptrglUniform4ui64vNV)(GLint location, GLsizei count, GLuint64EXT* value);
// void (APIENTRYP ptrglGetUniformi64vNV)(GLuint program, GLint location, GLint64EXT* params);
// void (APIENTRYP ptrglProgramUniform1i64NV)(GLuint program, GLint location, GLint64EXT x);
// void (APIENTRYP ptrglProgramUniform2i64NV)(GLuint program, GLint location, GLint64EXT x, GLint64EXT y);
// void (APIENTRYP ptrglProgramUniform3i64NV)(GLuint program, GLint location, GLint64EXT x, GLint64EXT y, GLint64EXT z);
// void (APIENTRYP ptrglProgramUniform4i64NV)(GLuint program, GLint location, GLint64EXT x, GLint64EXT y, GLint64EXT z, GLint64EXT w);
// void (APIENTRYP ptrglProgramUniform1i64vNV)(GLuint program, GLint location, GLsizei count, GLint64EXT* value);
// void (APIENTRYP ptrglProgramUniform2i64vNV)(GLuint program, GLint location, GLsizei count, GLint64EXT* value);
// void (APIENTRYP ptrglProgramUniform3i64vNV)(GLuint program, GLint location, GLsizei count, GLint64EXT* value);
// void (APIENTRYP ptrglProgramUniform4i64vNV)(GLuint program, GLint location, GLsizei count, GLint64EXT* value);
// void (APIENTRYP ptrglProgramUniform1ui64NV)(GLuint program, GLint location, GLuint64EXT x);
// void (APIENTRYP ptrglProgramUniform2ui64NV)(GLuint program, GLint location, GLuint64EXT x, GLuint64EXT y);
// void (APIENTRYP ptrglProgramUniform3ui64NV)(GLuint program, GLint location, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z);
// void (APIENTRYP ptrglProgramUniform4ui64NV)(GLuint program, GLint location, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z, GLuint64EXT w);
// void (APIENTRYP ptrglProgramUniform1ui64vNV)(GLuint program, GLint location, GLsizei count, GLuint64EXT* value);
// void (APIENTRYP ptrglProgramUniform2ui64vNV)(GLuint program, GLint location, GLsizei count, GLuint64EXT* value);
// void (APIENTRYP ptrglProgramUniform3ui64vNV)(GLuint program, GLint location, GLsizei count, GLuint64EXT* value);
// void (APIENTRYP ptrglProgramUniform4ui64vNV)(GLuint program, GLint location, GLsizei count, GLuint64EXT* value);
// //  NV_half_float
// void (APIENTRYP ptrglVertex2hNV)(GLhalfNV x, GLhalfNV y);
// void (APIENTRYP ptrglVertex2hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglVertex3hNV)(GLhalfNV x, GLhalfNV y, GLhalfNV z);
// void (APIENTRYP ptrglVertex3hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglVertex4hNV)(GLhalfNV x, GLhalfNV y, GLhalfNV z, GLhalfNV w);
// void (APIENTRYP ptrglVertex4hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglNormal3hNV)(GLhalfNV nx, GLhalfNV ny, GLhalfNV nz);
// void (APIENTRYP ptrglNormal3hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglColor3hNV)(GLhalfNV red, GLhalfNV green, GLhalfNV blue);
// void (APIENTRYP ptrglColor3hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglColor4hNV)(GLhalfNV red, GLhalfNV green, GLhalfNV blue, GLhalfNV alpha);
// void (APIENTRYP ptrglColor4hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglTexCoord1hNV)(GLhalfNV s);
// void (APIENTRYP ptrglTexCoord1hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglTexCoord2hNV)(GLhalfNV s, GLhalfNV t);
// void (APIENTRYP ptrglTexCoord2hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglTexCoord3hNV)(GLhalfNV s, GLhalfNV t, GLhalfNV r);
// void (APIENTRYP ptrglTexCoord3hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglTexCoord4hNV)(GLhalfNV s, GLhalfNV t, GLhalfNV r, GLhalfNV q);
// void (APIENTRYP ptrglTexCoord4hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglMultiTexCoord1hNV)(GLenum target, GLhalfNV s);
// void (APIENTRYP ptrglMultiTexCoord1hvNV)(GLenum target, GLhalfNV* v);
// void (APIENTRYP ptrglMultiTexCoord2hNV)(GLenum target, GLhalfNV s, GLhalfNV t);
// void (APIENTRYP ptrglMultiTexCoord2hvNV)(GLenum target, GLhalfNV* v);
// void (APIENTRYP ptrglMultiTexCoord3hNV)(GLenum target, GLhalfNV s, GLhalfNV t, GLhalfNV r);
// void (APIENTRYP ptrglMultiTexCoord3hvNV)(GLenum target, GLhalfNV* v);
// void (APIENTRYP ptrglMultiTexCoord4hNV)(GLenum target, GLhalfNV s, GLhalfNV t, GLhalfNV r, GLhalfNV q);
// void (APIENTRYP ptrglMultiTexCoord4hvNV)(GLenum target, GLhalfNV* v);
// void (APIENTRYP ptrglFogCoordhNV)(GLhalfNV fog);
// void (APIENTRYP ptrglFogCoordhvNV)(GLhalfNV* fog);
// void (APIENTRYP ptrglSecondaryColor3hNV)(GLhalfNV red, GLhalfNV green, GLhalfNV blue);
// void (APIENTRYP ptrglSecondaryColor3hvNV)(GLhalfNV* v);
// void (APIENTRYP ptrglVertexWeighthNV)(GLhalfNV weight);
// void (APIENTRYP ptrglVertexWeighthvNV)(GLhalfNV* weight);
// void (APIENTRYP ptrglVertexAttrib1hNV)(GLuint index, GLhalfNV x);
// void (APIENTRYP ptrglVertexAttrib1hvNV)(GLuint index, GLhalfNV* v);
// void (APIENTRYP ptrglVertexAttrib2hNV)(GLuint index, GLhalfNV x, GLhalfNV y);
// void (APIENTRYP ptrglVertexAttrib2hvNV)(GLuint index, GLhalfNV* v);
// void (APIENTRYP ptrglVertexAttrib3hNV)(GLuint index, GLhalfNV x, GLhalfNV y, GLhalfNV z);
// void (APIENTRYP ptrglVertexAttrib3hvNV)(GLuint index, GLhalfNV* v);
// void (APIENTRYP ptrglVertexAttrib4hNV)(GLuint index, GLhalfNV x, GLhalfNV y, GLhalfNV z, GLhalfNV w);
// void (APIENTRYP ptrglVertexAttrib4hvNV)(GLuint index, GLhalfNV* v);
// void (APIENTRYP ptrglVertexAttribs1hvNV)(GLuint index, GLsizei n, GLhalfNV* v);
// void (APIENTRYP ptrglVertexAttribs2hvNV)(GLuint index, GLsizei n, GLhalfNV* v);
// void (APIENTRYP ptrglVertexAttribs3hvNV)(GLuint index, GLsizei n, GLhalfNV* v);
// void (APIENTRYP ptrglVertexAttribs4hvNV)(GLuint index, GLsizei n, GLhalfNV* v);
// //  NV_occlusion_query
// void (APIENTRYP ptrglGenOcclusionQueriesNV)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglDeleteOcclusionQueriesNV)(GLsizei n, GLuint* ids);
// GLboolean (APIENTRYP ptrglIsOcclusionQueryNV)(GLuint id);
// void (APIENTRYP ptrglBeginOcclusionQueryNV)(GLuint id);
// void (APIENTRYP ptrglEndOcclusionQueryNV)();
// void (APIENTRYP ptrglGetOcclusionQueryivNV)(GLuint id, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetOcclusionQueryuivNV)(GLuint id, GLenum pname, GLuint* params);
// //  NV_parameter_buffer_object
// void (APIENTRYP ptrglProgramBufferParametersfvNV)(GLenum target, GLuint buffer, GLuint index, GLsizei count, GLfloat* params);
// void (APIENTRYP ptrglProgramBufferParametersIivNV)(GLenum target, GLuint buffer, GLuint index, GLsizei count, GLint* params);
// void (APIENTRYP ptrglProgramBufferParametersIuivNV)(GLenum target, GLuint buffer, GLuint index, GLsizei count, GLuint* params);
// //  NV_path_rendering
// GLuint (APIENTRYP ptrglGenPathsNV)(GLsizei range);
// void (APIENTRYP ptrglDeletePathsNV)(GLuint path, GLsizei range);
// GLboolean (APIENTRYP ptrglIsPathNV)(GLuint path);
// void (APIENTRYP ptrglPathCommandsNV)(GLuint path, GLsizei numCommands, GLubyte* commands, GLsizei numCoords, GLenum coordType, GLvoid* coords);
// void (APIENTRYP ptrglPathCoordsNV)(GLuint path, GLsizei numCoords, GLenum coordType, GLvoid* coords);
// void (APIENTRYP ptrglPathSubCommandsNV)(GLuint path, GLsizei commandStart, GLsizei commandsToDelete, GLsizei numCommands, GLubyte* commands, GLsizei numCoords, GLenum coordType, GLvoid* coords);
// void (APIENTRYP ptrglPathSubCoordsNV)(GLuint path, GLsizei coordStart, GLsizei numCoords, GLenum coordType, GLvoid* coords);
// void (APIENTRYP ptrglPathStringNV)(GLuint path, GLenum format, GLsizei length, GLvoid* pathString);
// void (APIENTRYP ptrglPathGlyphsNV)(GLuint firstPathName, GLenum fontTarget, GLvoid* fontName, GLbitfield fontStyle, GLsizei numGlyphs, GLenum type, GLvoid* charcodes, GLenum handleMissingGlyphs, GLuint pathParameterTemplate, GLfloat emScale);
// void (APIENTRYP ptrglPathGlyphRangeNV)(GLuint firstPathName, GLenum fontTarget, GLvoid* fontName, GLbitfield fontStyle, GLuint firstGlyph, GLsizei numGlyphs, GLenum handleMissingGlyphs, GLuint pathParameterTemplate, GLfloat emScale);
// void (APIENTRYP ptrglWeightPathsNV)(GLuint resultPath, GLsizei numPaths, GLuint* paths, GLfloat* weights);
// void (APIENTRYP ptrglCopyPathNV)(GLuint resultPath, GLuint srcPath);
// void (APIENTRYP ptrglInterpolatePathsNV)(GLuint resultPath, GLuint pathA, GLuint pathB, GLfloat weight);
// void (APIENTRYP ptrglTransformPathNV)(GLuint resultPath, GLuint srcPath, GLenum transformType, GLfloat* transformValues);
// void (APIENTRYP ptrglPathParameterivNV)(GLuint path, GLenum pname, GLint* value);
// void (APIENTRYP ptrglPathParameteriNV)(GLuint path, GLenum pname, GLint value);
// void (APIENTRYP ptrglPathParameterfvNV)(GLuint path, GLenum pname, GLfloat* value);
// void (APIENTRYP ptrglPathParameterfNV)(GLuint path, GLenum pname, GLfloat value);
// void (APIENTRYP ptrglPathDashArrayNV)(GLuint path, GLsizei dashCount, GLfloat* dashArray);
// void (APIENTRYP ptrglPathStencilFuncNV)(GLenum func, GLint ref, GLuint mask);
// void (APIENTRYP ptrglPathStencilDepthOffsetNV)(GLfloat factor, GLfloat units);
// void (APIENTRYP ptrglStencilFillPathNV)(GLuint path, GLenum fillMode, GLuint mask);
// void (APIENTRYP ptrglStencilStrokePathNV)(GLuint path, GLint reference, GLuint mask);
// void (APIENTRYP ptrglStencilFillPathInstancedNV)(GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLenum fillMode, GLuint mask, GLenum transformType, GLfloat* transformValues);
// void (APIENTRYP ptrglStencilStrokePathInstancedNV)(GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLint reference, GLuint mask, GLenum transformType, GLfloat* transformValues);
// void (APIENTRYP ptrglPathCoverDepthFuncNV)(GLenum func);
// void (APIENTRYP ptrglPathColorGenNV)(GLenum color, GLenum genMode, GLenum colorFormat, GLfloat* coeffs);
// void (APIENTRYP ptrglPathTexGenNV)(GLenum texCoordSet, GLenum genMode, GLint components, GLfloat* coeffs);
// void (APIENTRYP ptrglPathFogGenNV)(GLenum genMode);
// void (APIENTRYP ptrglCoverFillPathNV)(GLuint path, GLenum coverMode);
// void (APIENTRYP ptrglCoverStrokePathNV)(GLuint path, GLenum coverMode);
// void (APIENTRYP ptrglCoverFillPathInstancedNV)(GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLenum coverMode, GLenum transformType, GLfloat* transformValues);
// void (APIENTRYP ptrglCoverStrokePathInstancedNV)(GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLenum coverMode, GLenum transformType, GLfloat* transformValues);
// void (APIENTRYP ptrglGetPathParameterivNV)(GLuint path, GLenum pname, GLint* value);
// void (APIENTRYP ptrglGetPathParameterfvNV)(GLuint path, GLenum pname, GLfloat* value);
// void (APIENTRYP ptrglGetPathCommandsNV)(GLuint path, GLubyte* commands);
// void (APIENTRYP ptrglGetPathCoordsNV)(GLuint path, GLfloat* coords);
// void (APIENTRYP ptrglGetPathDashArrayNV)(GLuint path, GLfloat* dashArray);
// void (APIENTRYP ptrglGetPathMetricsNV)(GLbitfield metricQueryMask, GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLsizei stride, GLfloat* metrics);
// void (APIENTRYP ptrglGetPathMetricRangeNV)(GLbitfield metricQueryMask, GLuint firstPathName, GLsizei numPaths, GLsizei stride, GLfloat* metrics);
// void (APIENTRYP ptrglGetPathSpacingNV)(GLenum pathListMode, GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLfloat advanceScale, GLfloat kerningScale, GLenum transformType, GLfloat* returnedSpacing);
// void (APIENTRYP ptrglGetPathColorGenivNV)(GLenum color, GLenum pname, GLint* value);
// void (APIENTRYP ptrglGetPathColorGenfvNV)(GLenum color, GLenum pname, GLfloat* value);
// void (APIENTRYP ptrglGetPathTexGenivNV)(GLenum texCoordSet, GLenum pname, GLint* value);
// void (APIENTRYP ptrglGetPathTexGenfvNV)(GLenum texCoordSet, GLenum pname, GLfloat* value);
// GLboolean (APIENTRYP ptrglIsPointInFillPathNV)(GLuint path, GLuint mask, GLfloat x, GLfloat y);
// GLboolean (APIENTRYP ptrglIsPointInStrokePathNV)(GLuint path, GLfloat x, GLfloat y);
// GLfloat (APIENTRYP ptrglGetPathLengthNV)(GLuint path, GLsizei startSegment, GLsizei numSegments);
// GLboolean (APIENTRYP ptrglPointAlongPathNV)(GLuint path, GLsizei startSegment, GLsizei numSegments, GLfloat distance, GLfloat* x, GLfloat* y, GLfloat* tangentX, GLfloat* tangentY);
// //  NV_pixel_data_range
// void (APIENTRYP ptrglPixelDataRangeNV)(GLenum target, GLsizei length, GLvoid* pointer);
// void (APIENTRYP ptrglFlushPixelDataRangeNV)(GLenum target);
// //  NV_point_sprite
// void (APIENTRYP ptrglPointParameteriNV)(GLenum pname, GLint param);
// void (APIENTRYP ptrglPointParameterivNV)(GLenum pname, GLint* params);
// //  NV_present_video
// void (APIENTRYP ptrglPresentFrameKeyedNV)(GLuint video_slot, GLuint64EXT minPresentTime, GLuint beginPresentTimeId, GLuint presentDurationId, GLenum type, GLenum target0, GLuint fill0, GLuint key0, GLenum target1, GLuint fill1, GLuint key1);
// void (APIENTRYP ptrglPresentFrameDualFillNV)(GLuint video_slot, GLuint64EXT minPresentTime, GLuint beginPresentTimeId, GLuint presentDurationId, GLenum type, GLenum target0, GLuint fill0, GLenum target1, GLuint fill1, GLenum target2, GLuint fill2, GLenum target3, GLuint fill3);
// void (APIENTRYP ptrglGetVideoivNV)(GLuint video_slot, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetVideouivNV)(GLuint video_slot, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglGetVideoi64vNV)(GLuint video_slot, GLenum pname, GLint64EXT* params);
// void (APIENTRYP ptrglGetVideoui64vNV)(GLuint video_slot, GLenum pname, GLuint64EXT* params);
// //  NV_primitive_restart
// void (APIENTRYP ptrglPrimitiveRestartNV)();
// void (APIENTRYP ptrglPrimitiveRestartIndexNV)(GLuint index);
// //  NV_register_combiners
// void (APIENTRYP ptrglCombinerParameterfvNV)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglCombinerParameterfNV)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglCombinerParameterivNV)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglCombinerParameteriNV)(GLenum pname, GLint param);
// void (APIENTRYP ptrglCombinerInputNV)(GLenum stage, GLenum portion, GLenum variable, GLenum input, GLenum mapping, GLenum componentUsage);
// void (APIENTRYP ptrglCombinerOutputNV)(GLenum stage, GLenum portion, GLenum abOutput, GLenum cdOutput, GLenum sumOutput, GLenum scale, GLenum bias, GLboolean abDotProduct, GLboolean cdDotProduct, GLboolean muxSum);
// void (APIENTRYP ptrglFinalCombinerInputNV)(GLenum variable, GLenum input, GLenum mapping, GLenum componentUsage);
// void (APIENTRYP ptrglGetCombinerInputParameterfvNV)(GLenum stage, GLenum portion, GLenum variable, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetCombinerInputParameterivNV)(GLenum stage, GLenum portion, GLenum variable, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetCombinerOutputParameterfvNV)(GLenum stage, GLenum portion, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetCombinerOutputParameterivNV)(GLenum stage, GLenum portion, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetFinalCombinerInputParameterfvNV)(GLenum variable, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetFinalCombinerInputParameterivNV)(GLenum variable, GLenum pname, GLint* params);
// //  NV_register_combiners2
// void (APIENTRYP ptrglCombinerStageParameterfvNV)(GLenum stage, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetCombinerStageParameterfvNV)(GLenum stage, GLenum pname, GLfloat* params);
// //  NV_shader_buffer_load
// void (APIENTRYP ptrglMakeBufferResidentNV)(GLenum target, GLenum access);
// void (APIENTRYP ptrglMakeBufferNonResidentNV)(GLenum target);
// GLboolean (APIENTRYP ptrglIsBufferResidentNV)(GLenum target);
// void (APIENTRYP ptrglMakeNamedBufferResidentNV)(GLuint buffer, GLenum access);
// void (APIENTRYP ptrglMakeNamedBufferNonResidentNV)(GLuint buffer);
// GLboolean (APIENTRYP ptrglIsNamedBufferResidentNV)(GLuint buffer);
// void (APIENTRYP ptrglGetBufferParameterui64vNV)(GLenum target, GLenum pname, GLuint64EXT* params);
// void (APIENTRYP ptrglGetNamedBufferParameterui64vNV)(GLuint buffer, GLenum pname, GLuint64EXT* params);
// void (APIENTRYP ptrglGetIntegerui64vNV)(GLenum value, GLuint64EXT* result);
// void (APIENTRYP ptrglUniformui64NV)(GLint location, GLuint64EXT value);
// void (APIENTRYP ptrglUniformui64vNV)(GLint location, GLsizei count, GLuint64EXT* value);
// void (APIENTRYP ptrglGetUniformui64vNV)(GLuint program, GLint location, GLuint64EXT* params);
// void (APIENTRYP ptrglProgramUniformui64NV)(GLuint program, GLint location, GLuint64EXT value);
// void (APIENTRYP ptrglProgramUniformui64vNV)(GLuint program, GLint location, GLsizei count, GLuint64EXT* value);
// //  NV_texture_barrier
// void (APIENTRYP ptrglTextureBarrierNV)();
// //  NV_texture_multisample
// void (APIENTRYP ptrglTexImage2DMultisampleCoverageNV)(GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLint internalFormat, GLsizei width, GLsizei height, GLboolean fixedSampleLocations);
// void (APIENTRYP ptrglTexImage3DMultisampleCoverageNV)(GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLint internalFormat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedSampleLocations);
// void (APIENTRYP ptrglTextureImage2DMultisampleNV)(GLuint texture, GLenum target, GLsizei samples, GLint internalFormat, GLsizei width, GLsizei height, GLboolean fixedSampleLocations);
// void (APIENTRYP ptrglTextureImage3DMultisampleNV)(GLuint texture, GLenum target, GLsizei samples, GLint internalFormat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedSampleLocations);
// void (APIENTRYP ptrglTextureImage2DMultisampleCoverageNV)(GLuint texture, GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLint internalFormat, GLsizei width, GLsizei height, GLboolean fixedSampleLocations);
// void (APIENTRYP ptrglTextureImage3DMultisampleCoverageNV)(GLuint texture, GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLint internalFormat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedSampleLocations);
// //  NV_transform_feedback
// void (APIENTRYP ptrglBeginTransformFeedbackNV)(GLenum primitiveMode);
// void (APIENTRYP ptrglEndTransformFeedbackNV)();
// void (APIENTRYP ptrglTransformFeedbackAttribsNV)(GLsizei count, GLint* attribs, GLenum bufferMode);
// void (APIENTRYP ptrglBindBufferRangeNV)(GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size);
// void (APIENTRYP ptrglBindBufferOffsetNV)(GLenum target, GLuint index, GLuint buffer, GLintptr offset);
// void (APIENTRYP ptrglBindBufferBaseNV)(GLenum target, GLuint index, GLuint buffer);
// void (APIENTRYP ptrglTransformFeedbackVaryingsNV)(GLuint program, GLsizei count, GLint* locations, GLenum bufferMode);
// void (APIENTRYP ptrglActiveVaryingNV)(GLuint program, GLchar* name);
// GLint (APIENTRYP ptrglGetVaryingLocationNV)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglGetActiveVaryingNV)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLsizei* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrglGetTransformFeedbackVaryingNV)(GLuint program, GLuint index, GLint* location);
// void (APIENTRYP ptrglTransformFeedbackStreamAttribsNV)(GLsizei count, GLint* attribs, GLsizei nbuffers, GLint* bufstreams, GLenum bufferMode);
// //  NV_transform_feedback2
// void (APIENTRYP ptrglBindTransformFeedbackNV)(GLenum target, GLuint id);
// void (APIENTRYP ptrglDeleteTransformFeedbacksNV)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglGenTransformFeedbacksNV)(GLsizei n, GLuint* ids);
// GLboolean (APIENTRYP ptrglIsTransformFeedbackNV)(GLuint id);
// void (APIENTRYP ptrglPauseTransformFeedbackNV)();
// void (APIENTRYP ptrglResumeTransformFeedbackNV)();
// void (APIENTRYP ptrglDrawTransformFeedbackNV)(GLenum mode, GLuint id);
// //  NV_vdpau_interop
// void (APIENTRYP ptrglVDPAUInitNV)(GLvoid* vdpDevice, GLvoid* getProcAddress);
// void (APIENTRYP ptrglVDPAUFiniNV)();
// GLvdpauSurfaceNV (APIENTRYP ptrglVDPAURegisterVideoSurfaceNV)(GLvoid* vdpSurface, GLenum target, GLsizei numTextureNames, GLuint* textureNames);
// GLvdpauSurfaceNV (APIENTRYP ptrglVDPAURegisterOutputSurfaceNV)(GLvoid* vdpSurface, GLenum target, GLsizei numTextureNames, GLuint* textureNames);
// void (APIENTRYP ptrglVDPAUIsSurfaceNV)(GLvdpauSurfaceNV surface);
// void (APIENTRYP ptrglVDPAUUnregisterSurfaceNV)(GLvdpauSurfaceNV surface);
// void (APIENTRYP ptrglVDPAUGetSurfaceivNV)(GLvdpauSurfaceNV surface, GLenum pname, GLsizei bufSize, GLsizei* length, GLint* values);
// void (APIENTRYP ptrglVDPAUSurfaceAccessNV)(GLvdpauSurfaceNV surface, GLenum access);
// void (APIENTRYP ptrglVDPAUMapSurfacesNV)(GLsizei numSurfaces, GLvdpauSurfaceNV* surfaces);
// void (APIENTRYP ptrglVDPAUUnmapSurfacesNV)(GLsizei numSurface, GLvdpauSurfaceNV* surfaces);
// //  NV_vertex_array_range
// void (APIENTRYP ptrglFlushVertexArrayRangeNV)();
// void (APIENTRYP ptrglVertexArrayRangeNV)(GLsizei length, GLvoid* pointer);
// //  NV_vertex_attrib_integer_64bit
// void (APIENTRYP ptrglVertexAttribL1i64NV)(GLuint index, GLint64EXT x);
// void (APIENTRYP ptrglVertexAttribL2i64NV)(GLuint index, GLint64EXT x, GLint64EXT y);
// void (APIENTRYP ptrglVertexAttribL3i64NV)(GLuint index, GLint64EXT x, GLint64EXT y, GLint64EXT z);
// void (APIENTRYP ptrglVertexAttribL4i64NV)(GLuint index, GLint64EXT x, GLint64EXT y, GLint64EXT z, GLint64EXT w);
// void (APIENTRYP ptrglVertexAttribL1i64vNV)(GLuint index, GLint64EXT* v);
// void (APIENTRYP ptrglVertexAttribL2i64vNV)(GLuint index, GLint64EXT* v);
// void (APIENTRYP ptrglVertexAttribL3i64vNV)(GLuint index, GLint64EXT* v);
// void (APIENTRYP ptrglVertexAttribL4i64vNV)(GLuint index, GLint64EXT* v);
// void (APIENTRYP ptrglVertexAttribL1ui64NV)(GLuint index, GLuint64EXT x);
// void (APIENTRYP ptrglVertexAttribL2ui64NV)(GLuint index, GLuint64EXT x, GLuint64EXT y);
// void (APIENTRYP ptrglVertexAttribL3ui64NV)(GLuint index, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z);
// void (APIENTRYP ptrglVertexAttribL4ui64NV)(GLuint index, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z, GLuint64EXT w);
// void (APIENTRYP ptrglVertexAttribL1ui64vNV)(GLuint index, GLuint64EXT* v);
// void (APIENTRYP ptrglVertexAttribL2ui64vNV)(GLuint index, GLuint64EXT* v);
// void (APIENTRYP ptrglVertexAttribL3ui64vNV)(GLuint index, GLuint64EXT* v);
// void (APIENTRYP ptrglVertexAttribL4ui64vNV)(GLuint index, GLuint64EXT* v);
// void (APIENTRYP ptrglGetVertexAttribLi64vNV)(GLuint index, GLenum pname, GLint64EXT* params);
// void (APIENTRYP ptrglGetVertexAttribLui64vNV)(GLuint index, GLenum pname, GLuint64EXT* params);
// void (APIENTRYP ptrglVertexAttribLFormatNV)(GLuint index, GLint size, GLenum type, GLsizei stride);
// //  NV_vertex_buffer_unified_memory
// void (APIENTRYP ptrglBufferAddressRangeNV)(GLenum pname, GLuint index, GLuint64EXT address, GLsizeiptr length);
// void (APIENTRYP ptrglVertexFormatNV)(GLint size, GLenum type, GLsizei stride);
// void (APIENTRYP ptrglNormalFormatNV)(GLenum type, GLsizei stride);
// void (APIENTRYP ptrglColorFormatNV)(GLint size, GLenum type, GLsizei stride);
// void (APIENTRYP ptrglIndexFormatNV)(GLenum type, GLsizei stride);
// void (APIENTRYP ptrglTexCoordFormatNV)(GLint size, GLenum type, GLsizei stride);
// void (APIENTRYP ptrglEdgeFlagFormatNV)(GLsizei stride);
// void (APIENTRYP ptrglSecondaryColorFormatNV)(GLint size, GLenum type, GLsizei stride);
// void (APIENTRYP ptrglFogCoordFormatNV)(GLenum type, GLsizei stride);
// void (APIENTRYP ptrglVertexAttribFormatNV)(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride);
// void (APIENTRYP ptrglVertexAttribIFormatNV)(GLuint index, GLint size, GLenum type, GLsizei stride);
// void (APIENTRYP ptrglGetIntegerui64i_vNV)(GLenum value, GLuint index, GLuint64EXT* result);
// //  NV_vertex_program
// GLboolean (APIENTRYP ptrglAreProgramsResidentNV)(GLsizei n, GLuint* programs, GLboolean* residences);
// void (APIENTRYP ptrglBindProgramNV)(GLenum target, GLuint id);
// void (APIENTRYP ptrglDeleteProgramsNV)(GLsizei n, GLuint* programs);
// void (APIENTRYP ptrglExecuteProgramNV)(GLenum target, GLuint id, GLfloat* params);
// void (APIENTRYP ptrglGenProgramsNV)(GLsizei n, GLuint* programs);
// void (APIENTRYP ptrglGetProgramParameterdvNV)(GLenum target, GLuint index, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglGetProgramParameterfvNV)(GLenum target, GLuint index, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetProgramivNV)(GLuint id, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetProgramStringNV)(GLuint id, GLenum pname, GLubyte* program);
// void (APIENTRYP ptrglGetTrackMatrixivNV)(GLenum target, GLuint address, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetVertexAttribdvNV)(GLuint index, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglGetVertexAttribfvNV)(GLuint index, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetVertexAttribivNV)(GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetVertexAttribPointervNV)(GLuint index, GLenum pname, GLvoid** pointer);
// GLboolean (APIENTRYP ptrglIsProgramNV)(GLuint id);
// void (APIENTRYP ptrglLoadProgramNV)(GLenum target, GLuint id, GLsizei len, GLubyte* program);
// void (APIENTRYP ptrglProgramParameter4dNV)(GLenum target, GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglProgramParameter4dvNV)(GLenum target, GLuint index, GLdouble* v);
// void (APIENTRYP ptrglProgramParameter4fNV)(GLenum target, GLuint index, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglProgramParameter4fvNV)(GLenum target, GLuint index, GLfloat* v);
// void (APIENTRYP ptrglProgramParameters4dvNV)(GLenum target, GLuint index, GLsizei count, GLdouble* v);
// void (APIENTRYP ptrglProgramParameters4fvNV)(GLenum target, GLuint index, GLsizei count, GLfloat* v);
// void (APIENTRYP ptrglRequestResidentProgramsNV)(GLsizei n, GLuint* programs);
// void (APIENTRYP ptrglTrackMatrixNV)(GLenum target, GLuint address, GLenum matrix, GLenum transform);
// void (APIENTRYP ptrglVertexAttribPointerNV)(GLuint index, GLint fsize, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglVertexAttrib1dNV)(GLuint index, GLdouble x);
// void (APIENTRYP ptrglVertexAttrib1dvNV)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttrib1fNV)(GLuint index, GLfloat x);
// void (APIENTRYP ptrglVertexAttrib1fvNV)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglVertexAttrib1sNV)(GLuint index, GLshort x);
// void (APIENTRYP ptrglVertexAttrib1svNV)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttrib2dNV)(GLuint index, GLdouble x, GLdouble y);
// void (APIENTRYP ptrglVertexAttrib2dvNV)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttrib2fNV)(GLuint index, GLfloat x, GLfloat y);
// void (APIENTRYP ptrglVertexAttrib2fvNV)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglVertexAttrib2sNV)(GLuint index, GLshort x, GLshort y);
// void (APIENTRYP ptrglVertexAttrib2svNV)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttrib3dNV)(GLuint index, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglVertexAttrib3dvNV)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttrib3fNV)(GLuint index, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglVertexAttrib3fvNV)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglVertexAttrib3sNV)(GLuint index, GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrglVertexAttrib3svNV)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttrib4dNV)(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglVertexAttrib4dvNV)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglVertexAttrib4fNV)(GLuint index, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglVertexAttrib4fvNV)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglVertexAttrib4sNV)(GLuint index, GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrglVertexAttrib4svNV)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttrib4ubNV)(GLuint index, GLubyte x, GLubyte y, GLubyte z, GLubyte w);
// void (APIENTRYP ptrglVertexAttrib4ubvNV)(GLuint index, GLubyte* v);
// void (APIENTRYP ptrglVertexAttribs1dvNV)(GLuint index, GLsizei count, GLdouble* v);
// void (APIENTRYP ptrglVertexAttribs1fvNV)(GLuint index, GLsizei count, GLfloat* v);
// void (APIENTRYP ptrglVertexAttribs1svNV)(GLuint index, GLsizei count, GLshort* v);
// void (APIENTRYP ptrglVertexAttribs2dvNV)(GLuint index, GLsizei count, GLdouble* v);
// void (APIENTRYP ptrglVertexAttribs2fvNV)(GLuint index, GLsizei count, GLfloat* v);
// void (APIENTRYP ptrglVertexAttribs2svNV)(GLuint index, GLsizei count, GLshort* v);
// void (APIENTRYP ptrglVertexAttribs3dvNV)(GLuint index, GLsizei count, GLdouble* v);
// void (APIENTRYP ptrglVertexAttribs3fvNV)(GLuint index, GLsizei count, GLfloat* v);
// void (APIENTRYP ptrglVertexAttribs3svNV)(GLuint index, GLsizei count, GLshort* v);
// void (APIENTRYP ptrglVertexAttribs4dvNV)(GLuint index, GLsizei count, GLdouble* v);
// void (APIENTRYP ptrglVertexAttribs4fvNV)(GLuint index, GLsizei count, GLfloat* v);
// void (APIENTRYP ptrglVertexAttribs4svNV)(GLuint index, GLsizei count, GLshort* v);
// void (APIENTRYP ptrglVertexAttribs4ubvNV)(GLuint index, GLsizei count, GLubyte* v);
// //  NV_vertex_program4
// void (APIENTRYP ptrglVertexAttribI1iEXT)(GLuint index, GLint x);
// void (APIENTRYP ptrglVertexAttribI2iEXT)(GLuint index, GLint x, GLint y);
// void (APIENTRYP ptrglVertexAttribI3iEXT)(GLuint index, GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglVertexAttribI4iEXT)(GLuint index, GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglVertexAttribI1uiEXT)(GLuint index, GLuint x);
// void (APIENTRYP ptrglVertexAttribI2uiEXT)(GLuint index, GLuint x, GLuint y);
// void (APIENTRYP ptrglVertexAttribI3uiEXT)(GLuint index, GLuint x, GLuint y, GLuint z);
// void (APIENTRYP ptrglVertexAttribI4uiEXT)(GLuint index, GLuint x, GLuint y, GLuint z, GLuint w);
// void (APIENTRYP ptrglVertexAttribI1ivEXT)(GLuint index, GLint* v);
// void (APIENTRYP ptrglVertexAttribI2ivEXT)(GLuint index, GLint* v);
// void (APIENTRYP ptrglVertexAttribI3ivEXT)(GLuint index, GLint* v);
// void (APIENTRYP ptrglVertexAttribI4ivEXT)(GLuint index, GLint* v);
// void (APIENTRYP ptrglVertexAttribI1uivEXT)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttribI2uivEXT)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttribI3uivEXT)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttribI4uivEXT)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttribI4bvEXT)(GLuint index, GLbyte* v);
// void (APIENTRYP ptrglVertexAttribI4svEXT)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttribI4ubvEXT)(GLuint index, GLubyte* v);
// void (APIENTRYP ptrglVertexAttribI4usvEXT)(GLuint index, GLushort* v);
// void (APIENTRYP ptrglVertexAttribIPointerEXT)(GLuint index, GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglGetVertexAttribIivEXT)(GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetVertexAttribIuivEXT)(GLuint index, GLenum pname, GLuint* params);
// //  NV_video_capture
// void (APIENTRYP ptrglBeginVideoCaptureNV)(GLuint video_capture_slot);
// void (APIENTRYP ptrglBindVideoCaptureStreamBufferNV)(GLuint video_capture_slot, GLuint stream, GLenum frame_region, GLintptrARB offset);
// void (APIENTRYP ptrglBindVideoCaptureStreamTextureNV)(GLuint video_capture_slot, GLuint stream, GLenum frame_region, GLenum target, GLuint texture);
// void (APIENTRYP ptrglEndVideoCaptureNV)(GLuint video_capture_slot);
// void (APIENTRYP ptrglGetVideoCaptureivNV)(GLuint video_capture_slot, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetVideoCaptureStreamivNV)(GLuint video_capture_slot, GLuint stream, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetVideoCaptureStreamfvNV)(GLuint video_capture_slot, GLuint stream, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetVideoCaptureStreamdvNV)(GLuint video_capture_slot, GLuint stream, GLenum pname, GLdouble* params);
// GLenum (APIENTRYP ptrglVideoCaptureNV)(GLuint video_capture_slot, GLuint* sequence_num, GLuint64EXT* capture_time);
// void (APIENTRYP ptrglVideoCaptureStreamParameterivNV)(GLuint video_capture_slot, GLuint stream, GLenum pname, GLint* params);
// void (APIENTRYP ptrglVideoCaptureStreamParameterfvNV)(GLuint video_capture_slot, GLuint stream, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglVideoCaptureStreamParameterdvNV)(GLuint video_capture_slot, GLuint stream, GLenum pname, GLdouble* params);
// 
// //  NV_bindless_texture
// GLuint64 goglGetTextureHandleNV(GLuint texture) {
// 	return (*ptrglGetTextureHandleNV)(texture);
// }
// GLuint64 goglGetTextureSamplerHandleNV(GLuint texture, GLuint sampler) {
// 	return (*ptrglGetTextureSamplerHandleNV)(texture, sampler);
// }
// void goglMakeTextureHandleResidentNV(GLuint64 handle) {
// 	(*ptrglMakeTextureHandleResidentNV)(handle);
// }
// void goglMakeTextureHandleNonResidentNV(GLuint64 handle) {
// 	(*ptrglMakeTextureHandleNonResidentNV)(handle);
// }
// GLuint64 goglGetImageHandleNV(GLuint texture, GLint level, GLboolean layered, GLint layer, GLenum format) {
// 	return (*ptrglGetImageHandleNV)(texture, level, layered, layer, format);
// }
// void goglMakeImageHandleResidentNV(GLuint64 handle, GLenum access) {
// 	(*ptrglMakeImageHandleResidentNV)(handle, access);
// }
// void goglMakeImageHandleNonResidentNV(GLuint64 handle) {
// 	(*ptrglMakeImageHandleNonResidentNV)(handle);
// }
// void goglUniformHandleui64NV(GLint location, GLuint64 value) {
// 	(*ptrglUniformHandleui64NV)(location, value);
// }
// void goglUniformHandleui64vNV(GLint location, GLsizei count, GLuint64* value) {
// 	(*ptrglUniformHandleui64vNV)(location, count, value);
// }
// void goglProgramUniformHandleui64NV(GLuint program, GLint location, GLuint64 value) {
// 	(*ptrglProgramUniformHandleui64NV)(program, location, value);
// }
// void goglProgramUniformHandleui64vNV(GLuint program, GLint location, GLsizei count, GLuint64* values) {
// 	(*ptrglProgramUniformHandleui64vNV)(program, location, count, values);
// }
// GLboolean goglIsTextureHandleResidentNV(GLuint64 handle) {
// 	return (*ptrglIsTextureHandleResidentNV)(handle);
// }
// GLboolean goglIsImageHandleResidentNV(GLuint64 handle) {
// 	return (*ptrglIsImageHandleResidentNV)(handle);
// }
// //  NV_conditional_render
// void goglBeginConditionalRenderNV(GLuint id, GLenum mode) {
// 	(*ptrglBeginConditionalRenderNV)(id, mode);
// }
// void goglEndConditionalRenderNV() {
// 	(*ptrglEndConditionalRenderNV)();
// }
// //  NV_copy_image
// void goglCopyImageSubDataNV(GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei width, GLsizei height, GLsizei depth) {
// 	(*ptrglCopyImageSubDataNV)(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, width, height, depth);
// }
// //  NV_depth_buffer_float
// void goglDepthRangedNV(GLdouble zNear, GLdouble zFar) {
// 	(*ptrglDepthRangedNV)(zNear, zFar);
// }
// void goglClearDepthdNV(GLdouble depth) {
// 	(*ptrglClearDepthdNV)(depth);
// }
// void goglDepthBoundsdNV(GLdouble zmin, GLdouble zmax) {
// 	(*ptrglDepthBoundsdNV)(zmin, zmax);
// }
// //  NV_evaluators
// void goglMapControlPointsNV(GLenum target, GLuint index, GLenum type_, GLsizei ustride, GLsizei vstride, GLint uorder, GLint vorder, GLboolean packed, GLvoid* points) {
// 	(*ptrglMapControlPointsNV)(target, index, type_, ustride, vstride, uorder, vorder, packed, points);
// }
// void goglMapParameterivNV(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglMapParameterivNV)(target, pname, params);
// }
// void goglMapParameterfvNV(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglMapParameterfvNV)(target, pname, params);
// }
// void goglGetMapControlPointsNV(GLenum target, GLuint index, GLenum type_, GLsizei ustride, GLsizei vstride, GLboolean packed, GLvoid* points) {
// 	(*ptrglGetMapControlPointsNV)(target, index, type_, ustride, vstride, packed, points);
// }
// void goglGetMapParameterivNV(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetMapParameterivNV)(target, pname, params);
// }
// void goglGetMapParameterfvNV(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetMapParameterfvNV)(target, pname, params);
// }
// void goglGetMapAttribParameterivNV(GLenum target, GLuint index, GLenum pname, GLint* params) {
// 	(*ptrglGetMapAttribParameterivNV)(target, index, pname, params);
// }
// void goglGetMapAttribParameterfvNV(GLenum target, GLuint index, GLenum pname, GLfloat* params) {
// 	(*ptrglGetMapAttribParameterfvNV)(target, index, pname, params);
// }
// void goglEvalMapsNV(GLenum target, GLenum mode) {
// 	(*ptrglEvalMapsNV)(target, mode);
// }
// //  NV_explicit_multisample
// void goglGetMultisamplefvNV(GLenum pname, GLuint index, GLfloat* val) {
// 	(*ptrglGetMultisamplefvNV)(pname, index, val);
// }
// void goglSampleMaskIndexedNV(GLuint index, GLbitfield mask) {
// 	(*ptrglSampleMaskIndexedNV)(index, mask);
// }
// void goglTexRenderbufferNV(GLenum target, GLuint renderbuffer) {
// 	(*ptrglTexRenderbufferNV)(target, renderbuffer);
// }
// //  NV_fence
// void goglDeleteFencesNV(GLsizei n, GLuint* fences) {
// 	(*ptrglDeleteFencesNV)(n, fences);
// }
// void goglGenFencesNV(GLsizei n, GLuint* fences) {
// 	(*ptrglGenFencesNV)(n, fences);
// }
// GLboolean goglIsFenceNV(GLuint fence) {
// 	return (*ptrglIsFenceNV)(fence);
// }
// GLboolean goglTestFenceNV(GLuint fence) {
// 	return (*ptrglTestFenceNV)(fence);
// }
// void goglGetFenceivNV(GLuint fence, GLenum pname, GLint* params) {
// 	(*ptrglGetFenceivNV)(fence, pname, params);
// }
// void goglFinishFenceNV(GLuint fence) {
// 	(*ptrglFinishFenceNV)(fence);
// }
// void goglSetFenceNV(GLuint fence, GLenum condition) {
// 	(*ptrglSetFenceNV)(fence, condition);
// }
// //  NV_fragment_program
// void goglProgramNamedParameter4fNV(GLuint id, GLsizei len, GLubyte* name, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglProgramNamedParameter4fNV)(id, len, name, x, y, z, w);
// }
// void goglProgramNamedParameter4dNV(GLuint id, GLsizei len, GLubyte* name, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglProgramNamedParameter4dNV)(id, len, name, x, y, z, w);
// }
// void goglProgramNamedParameter4fvNV(GLuint id, GLsizei len, GLubyte* name, GLfloat* v) {
// 	(*ptrglProgramNamedParameter4fvNV)(id, len, name, v);
// }
// void goglProgramNamedParameter4dvNV(GLuint id, GLsizei len, GLubyte* name, GLdouble* v) {
// 	(*ptrglProgramNamedParameter4dvNV)(id, len, name, v);
// }
// void goglGetProgramNamedParameterfvNV(GLuint id, GLsizei len, GLubyte* name, GLfloat* params) {
// 	(*ptrglGetProgramNamedParameterfvNV)(id, len, name, params);
// }
// void goglGetProgramNamedParameterdvNV(GLuint id, GLsizei len, GLubyte* name, GLdouble* params) {
// 	(*ptrglGetProgramNamedParameterdvNV)(id, len, name, params);
// }
// //  NV_framebuffer_multisample_coverage
// void goglRenderbufferStorageMultisampleCoverageNV(GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLenum internalformat, GLsizei width, GLsizei height) {
// 	(*ptrglRenderbufferStorageMultisampleCoverageNV)(target, coverageSamples, colorSamples, internalformat, width, height);
// }
// //  NV_geometry_program4
// void goglProgramVertexLimitNV(GLenum target, GLint limit) {
// 	(*ptrglProgramVertexLimitNV)(target, limit);
// }
// void goglFramebufferTextureEXT(GLenum target, GLenum attachment, GLuint texture, GLint level) {
// 	(*ptrglFramebufferTextureEXT)(target, attachment, texture, level);
// }
// void goglFramebufferTextureLayerEXT(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer) {
// 	(*ptrglFramebufferTextureLayerEXT)(target, attachment, texture, level, layer);
// }
// void goglFramebufferTextureFaceEXT(GLenum target, GLenum attachment, GLuint texture, GLint level, GLenum face) {
// 	(*ptrglFramebufferTextureFaceEXT)(target, attachment, texture, level, face);
// }
// //  NV_gpu_program4
// void goglProgramLocalParameterI4iNV(GLenum target, GLuint index, GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrglProgramLocalParameterI4iNV)(target, index, x, y, z, w);
// }
// void goglProgramLocalParameterI4ivNV(GLenum target, GLuint index, GLint* params) {
// 	(*ptrglProgramLocalParameterI4ivNV)(target, index, params);
// }
// void goglProgramLocalParametersI4ivNV(GLenum target, GLuint index, GLsizei count, GLint* params) {
// 	(*ptrglProgramLocalParametersI4ivNV)(target, index, count, params);
// }
// void goglProgramLocalParameterI4uiNV(GLenum target, GLuint index, GLuint x, GLuint y, GLuint z, GLuint w) {
// 	(*ptrglProgramLocalParameterI4uiNV)(target, index, x, y, z, w);
// }
// void goglProgramLocalParameterI4uivNV(GLenum target, GLuint index, GLuint* params) {
// 	(*ptrglProgramLocalParameterI4uivNV)(target, index, params);
// }
// void goglProgramLocalParametersI4uivNV(GLenum target, GLuint index, GLsizei count, GLuint* params) {
// 	(*ptrglProgramLocalParametersI4uivNV)(target, index, count, params);
// }
// void goglProgramEnvParameterI4iNV(GLenum target, GLuint index, GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrglProgramEnvParameterI4iNV)(target, index, x, y, z, w);
// }
// void goglProgramEnvParameterI4ivNV(GLenum target, GLuint index, GLint* params) {
// 	(*ptrglProgramEnvParameterI4ivNV)(target, index, params);
// }
// void goglProgramEnvParametersI4ivNV(GLenum target, GLuint index, GLsizei count, GLint* params) {
// 	(*ptrglProgramEnvParametersI4ivNV)(target, index, count, params);
// }
// void goglProgramEnvParameterI4uiNV(GLenum target, GLuint index, GLuint x, GLuint y, GLuint z, GLuint w) {
// 	(*ptrglProgramEnvParameterI4uiNV)(target, index, x, y, z, w);
// }
// void goglProgramEnvParameterI4uivNV(GLenum target, GLuint index, GLuint* params) {
// 	(*ptrglProgramEnvParameterI4uivNV)(target, index, params);
// }
// void goglProgramEnvParametersI4uivNV(GLenum target, GLuint index, GLsizei count, GLuint* params) {
// 	(*ptrglProgramEnvParametersI4uivNV)(target, index, count, params);
// }
// void goglGetProgramLocalParameterIivNV(GLenum target, GLuint index, GLint* params) {
// 	(*ptrglGetProgramLocalParameterIivNV)(target, index, params);
// }
// void goglGetProgramLocalParameterIuivNV(GLenum target, GLuint index, GLuint* params) {
// 	(*ptrglGetProgramLocalParameterIuivNV)(target, index, params);
// }
// void goglGetProgramEnvParameterIivNV(GLenum target, GLuint index, GLint* params) {
// 	(*ptrglGetProgramEnvParameterIivNV)(target, index, params);
// }
// void goglGetProgramEnvParameterIuivNV(GLenum target, GLuint index, GLuint* params) {
// 	(*ptrglGetProgramEnvParameterIuivNV)(target, index, params);
// }
// //  NV_gpu_program5
// void goglProgramSubroutineParametersuivNV(GLenum target, GLsizei count, GLuint* params) {
// 	(*ptrglProgramSubroutineParametersuivNV)(target, count, params);
// }
// void goglGetProgramSubroutineParameteruivNV(GLenum target, GLuint index, GLuint* param) {
// 	(*ptrglGetProgramSubroutineParameteruivNV)(target, index, param);
// }
// //  NV_gpu_shader5
// void goglUniform1i64NV(GLint location, GLint64EXT x) {
// 	(*ptrglUniform1i64NV)(location, x);
// }
// void goglUniform2i64NV(GLint location, GLint64EXT x, GLint64EXT y) {
// 	(*ptrglUniform2i64NV)(location, x, y);
// }
// void goglUniform3i64NV(GLint location, GLint64EXT x, GLint64EXT y, GLint64EXT z) {
// 	(*ptrglUniform3i64NV)(location, x, y, z);
// }
// void goglUniform4i64NV(GLint location, GLint64EXT x, GLint64EXT y, GLint64EXT z, GLint64EXT w) {
// 	(*ptrglUniform4i64NV)(location, x, y, z, w);
// }
// void goglUniform1i64vNV(GLint location, GLsizei count, GLint64EXT* value) {
// 	(*ptrglUniform1i64vNV)(location, count, value);
// }
// void goglUniform2i64vNV(GLint location, GLsizei count, GLint64EXT* value) {
// 	(*ptrglUniform2i64vNV)(location, count, value);
// }
// void goglUniform3i64vNV(GLint location, GLsizei count, GLint64EXT* value) {
// 	(*ptrglUniform3i64vNV)(location, count, value);
// }
// void goglUniform4i64vNV(GLint location, GLsizei count, GLint64EXT* value) {
// 	(*ptrglUniform4i64vNV)(location, count, value);
// }
// void goglUniform1ui64NV(GLint location, GLuint64EXT x) {
// 	(*ptrglUniform1ui64NV)(location, x);
// }
// void goglUniform2ui64NV(GLint location, GLuint64EXT x, GLuint64EXT y) {
// 	(*ptrglUniform2ui64NV)(location, x, y);
// }
// void goglUniform3ui64NV(GLint location, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z) {
// 	(*ptrglUniform3ui64NV)(location, x, y, z);
// }
// void goglUniform4ui64NV(GLint location, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z, GLuint64EXT w) {
// 	(*ptrglUniform4ui64NV)(location, x, y, z, w);
// }
// void goglUniform1ui64vNV(GLint location, GLsizei count, GLuint64EXT* value) {
// 	(*ptrglUniform1ui64vNV)(location, count, value);
// }
// void goglUniform2ui64vNV(GLint location, GLsizei count, GLuint64EXT* value) {
// 	(*ptrglUniform2ui64vNV)(location, count, value);
// }
// void goglUniform3ui64vNV(GLint location, GLsizei count, GLuint64EXT* value) {
// 	(*ptrglUniform3ui64vNV)(location, count, value);
// }
// void goglUniform4ui64vNV(GLint location, GLsizei count, GLuint64EXT* value) {
// 	(*ptrglUniform4ui64vNV)(location, count, value);
// }
// void goglGetUniformi64vNV(GLuint program, GLint location, GLint64EXT* params) {
// 	(*ptrglGetUniformi64vNV)(program, location, params);
// }
// void goglProgramUniform1i64NV(GLuint program, GLint location, GLint64EXT x) {
// 	(*ptrglProgramUniform1i64NV)(program, location, x);
// }
// void goglProgramUniform2i64NV(GLuint program, GLint location, GLint64EXT x, GLint64EXT y) {
// 	(*ptrglProgramUniform2i64NV)(program, location, x, y);
// }
// void goglProgramUniform3i64NV(GLuint program, GLint location, GLint64EXT x, GLint64EXT y, GLint64EXT z) {
// 	(*ptrglProgramUniform3i64NV)(program, location, x, y, z);
// }
// void goglProgramUniform4i64NV(GLuint program, GLint location, GLint64EXT x, GLint64EXT y, GLint64EXT z, GLint64EXT w) {
// 	(*ptrglProgramUniform4i64NV)(program, location, x, y, z, w);
// }
// void goglProgramUniform1i64vNV(GLuint program, GLint location, GLsizei count, GLint64EXT* value) {
// 	(*ptrglProgramUniform1i64vNV)(program, location, count, value);
// }
// void goglProgramUniform2i64vNV(GLuint program, GLint location, GLsizei count, GLint64EXT* value) {
// 	(*ptrglProgramUniform2i64vNV)(program, location, count, value);
// }
// void goglProgramUniform3i64vNV(GLuint program, GLint location, GLsizei count, GLint64EXT* value) {
// 	(*ptrglProgramUniform3i64vNV)(program, location, count, value);
// }
// void goglProgramUniform4i64vNV(GLuint program, GLint location, GLsizei count, GLint64EXT* value) {
// 	(*ptrglProgramUniform4i64vNV)(program, location, count, value);
// }
// void goglProgramUniform1ui64NV(GLuint program, GLint location, GLuint64EXT x) {
// 	(*ptrglProgramUniform1ui64NV)(program, location, x);
// }
// void goglProgramUniform2ui64NV(GLuint program, GLint location, GLuint64EXT x, GLuint64EXT y) {
// 	(*ptrglProgramUniform2ui64NV)(program, location, x, y);
// }
// void goglProgramUniform3ui64NV(GLuint program, GLint location, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z) {
// 	(*ptrglProgramUniform3ui64NV)(program, location, x, y, z);
// }
// void goglProgramUniform4ui64NV(GLuint program, GLint location, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z, GLuint64EXT w) {
// 	(*ptrglProgramUniform4ui64NV)(program, location, x, y, z, w);
// }
// void goglProgramUniform1ui64vNV(GLuint program, GLint location, GLsizei count, GLuint64EXT* value) {
// 	(*ptrglProgramUniform1ui64vNV)(program, location, count, value);
// }
// void goglProgramUniform2ui64vNV(GLuint program, GLint location, GLsizei count, GLuint64EXT* value) {
// 	(*ptrglProgramUniform2ui64vNV)(program, location, count, value);
// }
// void goglProgramUniform3ui64vNV(GLuint program, GLint location, GLsizei count, GLuint64EXT* value) {
// 	(*ptrglProgramUniform3ui64vNV)(program, location, count, value);
// }
// void goglProgramUniform4ui64vNV(GLuint program, GLint location, GLsizei count, GLuint64EXT* value) {
// 	(*ptrglProgramUniform4ui64vNV)(program, location, count, value);
// }
// //  NV_half_float
// void goglVertex2hNV(GLhalfNV x, GLhalfNV y) {
// 	(*ptrglVertex2hNV)(x, y);
// }
// void goglVertex2hvNV(GLhalfNV* v) {
// 	(*ptrglVertex2hvNV)(v);
// }
// void goglVertex3hNV(GLhalfNV x, GLhalfNV y, GLhalfNV z) {
// 	(*ptrglVertex3hNV)(x, y, z);
// }
// void goglVertex3hvNV(GLhalfNV* v) {
// 	(*ptrglVertex3hvNV)(v);
// }
// void goglVertex4hNV(GLhalfNV x, GLhalfNV y, GLhalfNV z, GLhalfNV w) {
// 	(*ptrglVertex4hNV)(x, y, z, w);
// }
// void goglVertex4hvNV(GLhalfNV* v) {
// 	(*ptrglVertex4hvNV)(v);
// }
// void goglNormal3hNV(GLhalfNV nx, GLhalfNV ny, GLhalfNV nz) {
// 	(*ptrglNormal3hNV)(nx, ny, nz);
// }
// void goglNormal3hvNV(GLhalfNV* v) {
// 	(*ptrglNormal3hvNV)(v);
// }
// void goglColor3hNV(GLhalfNV red, GLhalfNV green, GLhalfNV blue) {
// 	(*ptrglColor3hNV)(red, green, blue);
// }
// void goglColor3hvNV(GLhalfNV* v) {
// 	(*ptrglColor3hvNV)(v);
// }
// void goglColor4hNV(GLhalfNV red, GLhalfNV green, GLhalfNV blue, GLhalfNV alpha) {
// 	(*ptrglColor4hNV)(red, green, blue, alpha);
// }
// void goglColor4hvNV(GLhalfNV* v) {
// 	(*ptrglColor4hvNV)(v);
// }
// void goglTexCoord1hNV(GLhalfNV s) {
// 	(*ptrglTexCoord1hNV)(s);
// }
// void goglTexCoord1hvNV(GLhalfNV* v) {
// 	(*ptrglTexCoord1hvNV)(v);
// }
// void goglTexCoord2hNV(GLhalfNV s, GLhalfNV t) {
// 	(*ptrglTexCoord2hNV)(s, t);
// }
// void goglTexCoord2hvNV(GLhalfNV* v) {
// 	(*ptrglTexCoord2hvNV)(v);
// }
// void goglTexCoord3hNV(GLhalfNV s, GLhalfNV t, GLhalfNV r) {
// 	(*ptrglTexCoord3hNV)(s, t, r);
// }
// void goglTexCoord3hvNV(GLhalfNV* v) {
// 	(*ptrglTexCoord3hvNV)(v);
// }
// void goglTexCoord4hNV(GLhalfNV s, GLhalfNV t, GLhalfNV r, GLhalfNV q) {
// 	(*ptrglTexCoord4hNV)(s, t, r, q);
// }
// void goglTexCoord4hvNV(GLhalfNV* v) {
// 	(*ptrglTexCoord4hvNV)(v);
// }
// void goglMultiTexCoord1hNV(GLenum target, GLhalfNV s) {
// 	(*ptrglMultiTexCoord1hNV)(target, s);
// }
// void goglMultiTexCoord1hvNV(GLenum target, GLhalfNV* v) {
// 	(*ptrglMultiTexCoord1hvNV)(target, v);
// }
// void goglMultiTexCoord2hNV(GLenum target, GLhalfNV s, GLhalfNV t) {
// 	(*ptrglMultiTexCoord2hNV)(target, s, t);
// }
// void goglMultiTexCoord2hvNV(GLenum target, GLhalfNV* v) {
// 	(*ptrglMultiTexCoord2hvNV)(target, v);
// }
// void goglMultiTexCoord3hNV(GLenum target, GLhalfNV s, GLhalfNV t, GLhalfNV r) {
// 	(*ptrglMultiTexCoord3hNV)(target, s, t, r);
// }
// void goglMultiTexCoord3hvNV(GLenum target, GLhalfNV* v) {
// 	(*ptrglMultiTexCoord3hvNV)(target, v);
// }
// void goglMultiTexCoord4hNV(GLenum target, GLhalfNV s, GLhalfNV t, GLhalfNV r, GLhalfNV q) {
// 	(*ptrglMultiTexCoord4hNV)(target, s, t, r, q);
// }
// void goglMultiTexCoord4hvNV(GLenum target, GLhalfNV* v) {
// 	(*ptrglMultiTexCoord4hvNV)(target, v);
// }
// void goglFogCoordhNV(GLhalfNV fog) {
// 	(*ptrglFogCoordhNV)(fog);
// }
// void goglFogCoordhvNV(GLhalfNV* fog) {
// 	(*ptrglFogCoordhvNV)(fog);
// }
// void goglSecondaryColor3hNV(GLhalfNV red, GLhalfNV green, GLhalfNV blue) {
// 	(*ptrglSecondaryColor3hNV)(red, green, blue);
// }
// void goglSecondaryColor3hvNV(GLhalfNV* v) {
// 	(*ptrglSecondaryColor3hvNV)(v);
// }
// void goglVertexWeighthNV(GLhalfNV weight) {
// 	(*ptrglVertexWeighthNV)(weight);
// }
// void goglVertexWeighthvNV(GLhalfNV* weight) {
// 	(*ptrglVertexWeighthvNV)(weight);
// }
// void goglVertexAttrib1hNV(GLuint index, GLhalfNV x) {
// 	(*ptrglVertexAttrib1hNV)(index, x);
// }
// void goglVertexAttrib1hvNV(GLuint index, GLhalfNV* v) {
// 	(*ptrglVertexAttrib1hvNV)(index, v);
// }
// void goglVertexAttrib2hNV(GLuint index, GLhalfNV x, GLhalfNV y) {
// 	(*ptrglVertexAttrib2hNV)(index, x, y);
// }
// void goglVertexAttrib2hvNV(GLuint index, GLhalfNV* v) {
// 	(*ptrglVertexAttrib2hvNV)(index, v);
// }
// void goglVertexAttrib3hNV(GLuint index, GLhalfNV x, GLhalfNV y, GLhalfNV z) {
// 	(*ptrglVertexAttrib3hNV)(index, x, y, z);
// }
// void goglVertexAttrib3hvNV(GLuint index, GLhalfNV* v) {
// 	(*ptrglVertexAttrib3hvNV)(index, v);
// }
// void goglVertexAttrib4hNV(GLuint index, GLhalfNV x, GLhalfNV y, GLhalfNV z, GLhalfNV w) {
// 	(*ptrglVertexAttrib4hNV)(index, x, y, z, w);
// }
// void goglVertexAttrib4hvNV(GLuint index, GLhalfNV* v) {
// 	(*ptrglVertexAttrib4hvNV)(index, v);
// }
// void goglVertexAttribs1hvNV(GLuint index, GLsizei n, GLhalfNV* v) {
// 	(*ptrglVertexAttribs1hvNV)(index, n, v);
// }
// void goglVertexAttribs2hvNV(GLuint index, GLsizei n, GLhalfNV* v) {
// 	(*ptrglVertexAttribs2hvNV)(index, n, v);
// }
// void goglVertexAttribs3hvNV(GLuint index, GLsizei n, GLhalfNV* v) {
// 	(*ptrglVertexAttribs3hvNV)(index, n, v);
// }
// void goglVertexAttribs4hvNV(GLuint index, GLsizei n, GLhalfNV* v) {
// 	(*ptrglVertexAttribs4hvNV)(index, n, v);
// }
// //  NV_occlusion_query
// void goglGenOcclusionQueriesNV(GLsizei n, GLuint* ids) {
// 	(*ptrglGenOcclusionQueriesNV)(n, ids);
// }
// void goglDeleteOcclusionQueriesNV(GLsizei n, GLuint* ids) {
// 	(*ptrglDeleteOcclusionQueriesNV)(n, ids);
// }
// GLboolean goglIsOcclusionQueryNV(GLuint id) {
// 	return (*ptrglIsOcclusionQueryNV)(id);
// }
// void goglBeginOcclusionQueryNV(GLuint id) {
// 	(*ptrglBeginOcclusionQueryNV)(id);
// }
// void goglEndOcclusionQueryNV() {
// 	(*ptrglEndOcclusionQueryNV)();
// }
// void goglGetOcclusionQueryivNV(GLuint id, GLenum pname, GLint* params) {
// 	(*ptrglGetOcclusionQueryivNV)(id, pname, params);
// }
// void goglGetOcclusionQueryuivNV(GLuint id, GLenum pname, GLuint* params) {
// 	(*ptrglGetOcclusionQueryuivNV)(id, pname, params);
// }
// //  NV_parameter_buffer_object
// void goglProgramBufferParametersfvNV(GLenum target, GLuint buffer, GLuint index, GLsizei count, GLfloat* params) {
// 	(*ptrglProgramBufferParametersfvNV)(target, buffer, index, count, params);
// }
// void goglProgramBufferParametersIivNV(GLenum target, GLuint buffer, GLuint index, GLsizei count, GLint* params) {
// 	(*ptrglProgramBufferParametersIivNV)(target, buffer, index, count, params);
// }
// void goglProgramBufferParametersIuivNV(GLenum target, GLuint buffer, GLuint index, GLsizei count, GLuint* params) {
// 	(*ptrglProgramBufferParametersIuivNV)(target, buffer, index, count, params);
// }
// //  NV_path_rendering
// GLuint goglGenPathsNV(GLsizei range_) {
// 	return (*ptrglGenPathsNV)(range_);
// }
// void goglDeletePathsNV(GLuint path, GLsizei range_) {
// 	(*ptrglDeletePathsNV)(path, range_);
// }
// GLboolean goglIsPathNV(GLuint path) {
// 	return (*ptrglIsPathNV)(path);
// }
// void goglPathCommandsNV(GLuint path, GLsizei numCommands, GLubyte* commands, GLsizei numCoords, GLenum coordType, GLvoid* coords) {
// 	(*ptrglPathCommandsNV)(path, numCommands, commands, numCoords, coordType, coords);
// }
// void goglPathCoordsNV(GLuint path, GLsizei numCoords, GLenum coordType, GLvoid* coords) {
// 	(*ptrglPathCoordsNV)(path, numCoords, coordType, coords);
// }
// void goglPathSubCommandsNV(GLuint path, GLsizei commandStart, GLsizei commandsToDelete, GLsizei numCommands, GLubyte* commands, GLsizei numCoords, GLenum coordType, GLvoid* coords) {
// 	(*ptrglPathSubCommandsNV)(path, commandStart, commandsToDelete, numCommands, commands, numCoords, coordType, coords);
// }
// void goglPathSubCoordsNV(GLuint path, GLsizei coordStart, GLsizei numCoords, GLenum coordType, GLvoid* coords) {
// 	(*ptrglPathSubCoordsNV)(path, coordStart, numCoords, coordType, coords);
// }
// void goglPathStringNV(GLuint path, GLenum format, GLsizei length, GLvoid* pathString) {
// 	(*ptrglPathStringNV)(path, format, length, pathString);
// }
// void goglPathGlyphsNV(GLuint firstPathName, GLenum fontTarget, GLvoid* fontName, GLbitfield fontStyle, GLsizei numGlyphs, GLenum type_, GLvoid* charcodes, GLenum handleMissingGlyphs, GLuint pathParameterTemplate, GLfloat emScale) {
// 	(*ptrglPathGlyphsNV)(firstPathName, fontTarget, fontName, fontStyle, numGlyphs, type_, charcodes, handleMissingGlyphs, pathParameterTemplate, emScale);
// }
// void goglPathGlyphRangeNV(GLuint firstPathName, GLenum fontTarget, GLvoid* fontName, GLbitfield fontStyle, GLuint firstGlyph, GLsizei numGlyphs, GLenum handleMissingGlyphs, GLuint pathParameterTemplate, GLfloat emScale) {
// 	(*ptrglPathGlyphRangeNV)(firstPathName, fontTarget, fontName, fontStyle, firstGlyph, numGlyphs, handleMissingGlyphs, pathParameterTemplate, emScale);
// }
// void goglWeightPathsNV(GLuint resultPath, GLsizei numPaths, GLuint* paths, GLfloat* weights) {
// 	(*ptrglWeightPathsNV)(resultPath, numPaths, paths, weights);
// }
// void goglCopyPathNV(GLuint resultPath, GLuint srcPath) {
// 	(*ptrglCopyPathNV)(resultPath, srcPath);
// }
// void goglInterpolatePathsNV(GLuint resultPath, GLuint pathA, GLuint pathB, GLfloat weight) {
// 	(*ptrglInterpolatePathsNV)(resultPath, pathA, pathB, weight);
// }
// void goglTransformPathNV(GLuint resultPath, GLuint srcPath, GLenum transformType, GLfloat* transformValues) {
// 	(*ptrglTransformPathNV)(resultPath, srcPath, transformType, transformValues);
// }
// void goglPathParameterivNV(GLuint path, GLenum pname, GLint* value) {
// 	(*ptrglPathParameterivNV)(path, pname, value);
// }
// void goglPathParameteriNV(GLuint path, GLenum pname, GLint value) {
// 	(*ptrglPathParameteriNV)(path, pname, value);
// }
// void goglPathParameterfvNV(GLuint path, GLenum pname, GLfloat* value) {
// 	(*ptrglPathParameterfvNV)(path, pname, value);
// }
// void goglPathParameterfNV(GLuint path, GLenum pname, GLfloat value) {
// 	(*ptrglPathParameterfNV)(path, pname, value);
// }
// void goglPathDashArrayNV(GLuint path, GLsizei dashCount, GLfloat* dashArray) {
// 	(*ptrglPathDashArrayNV)(path, dashCount, dashArray);
// }
// void goglPathStencilFuncNV(GLenum func_, GLint ref, GLuint mask) {
// 	(*ptrglPathStencilFuncNV)(func_, ref, mask);
// }
// void goglPathStencilDepthOffsetNV(GLfloat factor, GLfloat units) {
// 	(*ptrglPathStencilDepthOffsetNV)(factor, units);
// }
// void goglStencilFillPathNV(GLuint path, GLenum fillMode, GLuint mask) {
// 	(*ptrglStencilFillPathNV)(path, fillMode, mask);
// }
// void goglStencilStrokePathNV(GLuint path, GLint reference, GLuint mask) {
// 	(*ptrglStencilStrokePathNV)(path, reference, mask);
// }
// void goglStencilFillPathInstancedNV(GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLenum fillMode, GLuint mask, GLenum transformType, GLfloat* transformValues) {
// 	(*ptrglStencilFillPathInstancedNV)(numPaths, pathNameType, paths, pathBase, fillMode, mask, transformType, transformValues);
// }
// void goglStencilStrokePathInstancedNV(GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLint reference, GLuint mask, GLenum transformType, GLfloat* transformValues) {
// 	(*ptrglStencilStrokePathInstancedNV)(numPaths, pathNameType, paths, pathBase, reference, mask, transformType, transformValues);
// }
// void goglPathCoverDepthFuncNV(GLenum func_) {
// 	(*ptrglPathCoverDepthFuncNV)(func_);
// }
// void goglPathColorGenNV(GLenum color, GLenum genMode, GLenum colorFormat, GLfloat* coeffs) {
// 	(*ptrglPathColorGenNV)(color, genMode, colorFormat, coeffs);
// }
// void goglPathTexGenNV(GLenum texCoordSet, GLenum genMode, GLint components, GLfloat* coeffs) {
// 	(*ptrglPathTexGenNV)(texCoordSet, genMode, components, coeffs);
// }
// void goglPathFogGenNV(GLenum genMode) {
// 	(*ptrglPathFogGenNV)(genMode);
// }
// void goglCoverFillPathNV(GLuint path, GLenum coverMode) {
// 	(*ptrglCoverFillPathNV)(path, coverMode);
// }
// void goglCoverStrokePathNV(GLuint path, GLenum coverMode) {
// 	(*ptrglCoverStrokePathNV)(path, coverMode);
// }
// void goglCoverFillPathInstancedNV(GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLenum coverMode, GLenum transformType, GLfloat* transformValues) {
// 	(*ptrglCoverFillPathInstancedNV)(numPaths, pathNameType, paths, pathBase, coverMode, transformType, transformValues);
// }
// void goglCoverStrokePathInstancedNV(GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLenum coverMode, GLenum transformType, GLfloat* transformValues) {
// 	(*ptrglCoverStrokePathInstancedNV)(numPaths, pathNameType, paths, pathBase, coverMode, transformType, transformValues);
// }
// void goglGetPathParameterivNV(GLuint path, GLenum pname, GLint* value) {
// 	(*ptrglGetPathParameterivNV)(path, pname, value);
// }
// void goglGetPathParameterfvNV(GLuint path, GLenum pname, GLfloat* value) {
// 	(*ptrglGetPathParameterfvNV)(path, pname, value);
// }
// void goglGetPathCommandsNV(GLuint path, GLubyte* commands) {
// 	(*ptrglGetPathCommandsNV)(path, commands);
// }
// void goglGetPathCoordsNV(GLuint path, GLfloat* coords) {
// 	(*ptrglGetPathCoordsNV)(path, coords);
// }
// void goglGetPathDashArrayNV(GLuint path, GLfloat* dashArray) {
// 	(*ptrglGetPathDashArrayNV)(path, dashArray);
// }
// void goglGetPathMetricsNV(GLbitfield metricQueryMask, GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLsizei stride, GLfloat* metrics) {
// 	(*ptrglGetPathMetricsNV)(metricQueryMask, numPaths, pathNameType, paths, pathBase, stride, metrics);
// }
// void goglGetPathMetricRangeNV(GLbitfield metricQueryMask, GLuint firstPathName, GLsizei numPaths, GLsizei stride, GLfloat* metrics) {
// 	(*ptrglGetPathMetricRangeNV)(metricQueryMask, firstPathName, numPaths, stride, metrics);
// }
// void goglGetPathSpacingNV(GLenum pathListMode, GLsizei numPaths, GLenum pathNameType, GLvoid* paths, GLuint pathBase, GLfloat advanceScale, GLfloat kerningScale, GLenum transformType, GLfloat* returnedSpacing) {
// 	(*ptrglGetPathSpacingNV)(pathListMode, numPaths, pathNameType, paths, pathBase, advanceScale, kerningScale, transformType, returnedSpacing);
// }
// void goglGetPathColorGenivNV(GLenum color, GLenum pname, GLint* value) {
// 	(*ptrglGetPathColorGenivNV)(color, pname, value);
// }
// void goglGetPathColorGenfvNV(GLenum color, GLenum pname, GLfloat* value) {
// 	(*ptrglGetPathColorGenfvNV)(color, pname, value);
// }
// void goglGetPathTexGenivNV(GLenum texCoordSet, GLenum pname, GLint* value) {
// 	(*ptrglGetPathTexGenivNV)(texCoordSet, pname, value);
// }
// void goglGetPathTexGenfvNV(GLenum texCoordSet, GLenum pname, GLfloat* value) {
// 	(*ptrglGetPathTexGenfvNV)(texCoordSet, pname, value);
// }
// GLboolean goglIsPointInFillPathNV(GLuint path, GLuint mask, GLfloat x, GLfloat y) {
// 	return (*ptrglIsPointInFillPathNV)(path, mask, x, y);
// }
// GLboolean goglIsPointInStrokePathNV(GLuint path, GLfloat x, GLfloat y) {
// 	return (*ptrglIsPointInStrokePathNV)(path, x, y);
// }
// GLfloat goglGetPathLengthNV(GLuint path, GLsizei startSegment, GLsizei numSegments) {
// 	return (*ptrglGetPathLengthNV)(path, startSegment, numSegments);
// }
// GLboolean goglPointAlongPathNV(GLuint path, GLsizei startSegment, GLsizei numSegments, GLfloat distance, GLfloat* x, GLfloat* y, GLfloat* tangentX, GLfloat* tangentY) {
// 	return (*ptrglPointAlongPathNV)(path, startSegment, numSegments, distance, x, y, tangentX, tangentY);
// }
// //  NV_pixel_data_range
// void goglPixelDataRangeNV(GLenum target, GLsizei length, GLvoid* pointer) {
// 	(*ptrglPixelDataRangeNV)(target, length, pointer);
// }
// void goglFlushPixelDataRangeNV(GLenum target) {
// 	(*ptrglFlushPixelDataRangeNV)(target);
// }
// //  NV_point_sprite
// void goglPointParameteriNV(GLenum pname, GLint param) {
// 	(*ptrglPointParameteriNV)(pname, param);
// }
// void goglPointParameterivNV(GLenum pname, GLint* params) {
// 	(*ptrglPointParameterivNV)(pname, params);
// }
// //  NV_present_video
// void goglPresentFrameKeyedNV(GLuint video_slot, GLuint64EXT minPresentTime, GLuint beginPresentTimeId, GLuint presentDurationId, GLenum type_, GLenum target0, GLuint fill0, GLuint key0, GLenum target1, GLuint fill1, GLuint key1) {
// 	(*ptrglPresentFrameKeyedNV)(video_slot, minPresentTime, beginPresentTimeId, presentDurationId, type_, target0, fill0, key0, target1, fill1, key1);
// }
// void goglPresentFrameDualFillNV(GLuint video_slot, GLuint64EXT minPresentTime, GLuint beginPresentTimeId, GLuint presentDurationId, GLenum type_, GLenum target0, GLuint fill0, GLenum target1, GLuint fill1, GLenum target2, GLuint fill2, GLenum target3, GLuint fill3) {
// 	(*ptrglPresentFrameDualFillNV)(video_slot, minPresentTime, beginPresentTimeId, presentDurationId, type_, target0, fill0, target1, fill1, target2, fill2, target3, fill3);
// }
// void goglGetVideoivNV(GLuint video_slot, GLenum pname, GLint* params) {
// 	(*ptrglGetVideoivNV)(video_slot, pname, params);
// }
// void goglGetVideouivNV(GLuint video_slot, GLenum pname, GLuint* params) {
// 	(*ptrglGetVideouivNV)(video_slot, pname, params);
// }
// void goglGetVideoi64vNV(GLuint video_slot, GLenum pname, GLint64EXT* params) {
// 	(*ptrglGetVideoi64vNV)(video_slot, pname, params);
// }
// void goglGetVideoui64vNV(GLuint video_slot, GLenum pname, GLuint64EXT* params) {
// 	(*ptrglGetVideoui64vNV)(video_slot, pname, params);
// }
// //  NV_primitive_restart
// void goglPrimitiveRestartNV() {
// 	(*ptrglPrimitiveRestartNV)();
// }
// void goglPrimitiveRestartIndexNV(GLuint index) {
// 	(*ptrglPrimitiveRestartIndexNV)(index);
// }
// //  NV_register_combiners
// void goglCombinerParameterfvNV(GLenum pname, GLfloat* params) {
// 	(*ptrglCombinerParameterfvNV)(pname, params);
// }
// void goglCombinerParameterfNV(GLenum pname, GLfloat param) {
// 	(*ptrglCombinerParameterfNV)(pname, param);
// }
// void goglCombinerParameterivNV(GLenum pname, GLint* params) {
// 	(*ptrglCombinerParameterivNV)(pname, params);
// }
// void goglCombinerParameteriNV(GLenum pname, GLint param) {
// 	(*ptrglCombinerParameteriNV)(pname, param);
// }
// void goglCombinerInputNV(GLenum stage, GLenum portion, GLenum variable, GLenum input, GLenum mapping, GLenum componentUsage) {
// 	(*ptrglCombinerInputNV)(stage, portion, variable, input, mapping, componentUsage);
// }
// void goglCombinerOutputNV(GLenum stage, GLenum portion, GLenum abOutput, GLenum cdOutput, GLenum sumOutput, GLenum scale, GLenum bias, GLboolean abDotProduct, GLboolean cdDotProduct, GLboolean muxSum) {
// 	(*ptrglCombinerOutputNV)(stage, portion, abOutput, cdOutput, sumOutput, scale, bias, abDotProduct, cdDotProduct, muxSum);
// }
// void goglFinalCombinerInputNV(GLenum variable, GLenum input, GLenum mapping, GLenum componentUsage) {
// 	(*ptrglFinalCombinerInputNV)(variable, input, mapping, componentUsage);
// }
// void goglGetCombinerInputParameterfvNV(GLenum stage, GLenum portion, GLenum variable, GLenum pname, GLfloat* params) {
// 	(*ptrglGetCombinerInputParameterfvNV)(stage, portion, variable, pname, params);
// }
// void goglGetCombinerInputParameterivNV(GLenum stage, GLenum portion, GLenum variable, GLenum pname, GLint* params) {
// 	(*ptrglGetCombinerInputParameterivNV)(stage, portion, variable, pname, params);
// }
// void goglGetCombinerOutputParameterfvNV(GLenum stage, GLenum portion, GLenum pname, GLfloat* params) {
// 	(*ptrglGetCombinerOutputParameterfvNV)(stage, portion, pname, params);
// }
// void goglGetCombinerOutputParameterivNV(GLenum stage, GLenum portion, GLenum pname, GLint* params) {
// 	(*ptrglGetCombinerOutputParameterivNV)(stage, portion, pname, params);
// }
// void goglGetFinalCombinerInputParameterfvNV(GLenum variable, GLenum pname, GLfloat* params) {
// 	(*ptrglGetFinalCombinerInputParameterfvNV)(variable, pname, params);
// }
// void goglGetFinalCombinerInputParameterivNV(GLenum variable, GLenum pname, GLint* params) {
// 	(*ptrglGetFinalCombinerInputParameterivNV)(variable, pname, params);
// }
// //  NV_register_combiners2
// void goglCombinerStageParameterfvNV(GLenum stage, GLenum pname, GLfloat* params) {
// 	(*ptrglCombinerStageParameterfvNV)(stage, pname, params);
// }
// void goglGetCombinerStageParameterfvNV(GLenum stage, GLenum pname, GLfloat* params) {
// 	(*ptrglGetCombinerStageParameterfvNV)(stage, pname, params);
// }
// //  NV_shader_buffer_load
// void goglMakeBufferResidentNV(GLenum target, GLenum access) {
// 	(*ptrglMakeBufferResidentNV)(target, access);
// }
// void goglMakeBufferNonResidentNV(GLenum target) {
// 	(*ptrglMakeBufferNonResidentNV)(target);
// }
// GLboolean goglIsBufferResidentNV(GLenum target) {
// 	return (*ptrglIsBufferResidentNV)(target);
// }
// void goglMakeNamedBufferResidentNV(GLuint buffer, GLenum access) {
// 	(*ptrglMakeNamedBufferResidentNV)(buffer, access);
// }
// void goglMakeNamedBufferNonResidentNV(GLuint buffer) {
// 	(*ptrglMakeNamedBufferNonResidentNV)(buffer);
// }
// GLboolean goglIsNamedBufferResidentNV(GLuint buffer) {
// 	return (*ptrglIsNamedBufferResidentNV)(buffer);
// }
// void goglGetBufferParameterui64vNV(GLenum target, GLenum pname, GLuint64EXT* params) {
// 	(*ptrglGetBufferParameterui64vNV)(target, pname, params);
// }
// void goglGetNamedBufferParameterui64vNV(GLuint buffer, GLenum pname, GLuint64EXT* params) {
// 	(*ptrglGetNamedBufferParameterui64vNV)(buffer, pname, params);
// }
// void goglGetIntegerui64vNV(GLenum value, GLuint64EXT* result) {
// 	(*ptrglGetIntegerui64vNV)(value, result);
// }
// void goglUniformui64NV(GLint location, GLuint64EXT value) {
// 	(*ptrglUniformui64NV)(location, value);
// }
// void goglUniformui64vNV(GLint location, GLsizei count, GLuint64EXT* value) {
// 	(*ptrglUniformui64vNV)(location, count, value);
// }
// void goglGetUniformui64vNV(GLuint program, GLint location, GLuint64EXT* params) {
// 	(*ptrglGetUniformui64vNV)(program, location, params);
// }
// void goglProgramUniformui64NV(GLuint program, GLint location, GLuint64EXT value) {
// 	(*ptrglProgramUniformui64NV)(program, location, value);
// }
// void goglProgramUniformui64vNV(GLuint program, GLint location, GLsizei count, GLuint64EXT* value) {
// 	(*ptrglProgramUniformui64vNV)(program, location, count, value);
// }
// //  NV_texture_barrier
// void goglTextureBarrierNV() {
// 	(*ptrglTextureBarrierNV)();
// }
// //  NV_texture_multisample
// void goglTexImage2DMultisampleCoverageNV(GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLint internalFormat, GLsizei width, GLsizei height, GLboolean fixedSampleLocations) {
// 	(*ptrglTexImage2DMultisampleCoverageNV)(target, coverageSamples, colorSamples, internalFormat, width, height, fixedSampleLocations);
// }
// void goglTexImage3DMultisampleCoverageNV(GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLint internalFormat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedSampleLocations) {
// 	(*ptrglTexImage3DMultisampleCoverageNV)(target, coverageSamples, colorSamples, internalFormat, width, height, depth, fixedSampleLocations);
// }
// void goglTextureImage2DMultisampleNV(GLuint texture, GLenum target, GLsizei samples, GLint internalFormat, GLsizei width, GLsizei height, GLboolean fixedSampleLocations) {
// 	(*ptrglTextureImage2DMultisampleNV)(texture, target, samples, internalFormat, width, height, fixedSampleLocations);
// }
// void goglTextureImage3DMultisampleNV(GLuint texture, GLenum target, GLsizei samples, GLint internalFormat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedSampleLocations) {
// 	(*ptrglTextureImage3DMultisampleNV)(texture, target, samples, internalFormat, width, height, depth, fixedSampleLocations);
// }
// void goglTextureImage2DMultisampleCoverageNV(GLuint texture, GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLint internalFormat, GLsizei width, GLsizei height, GLboolean fixedSampleLocations) {
// 	(*ptrglTextureImage2DMultisampleCoverageNV)(texture, target, coverageSamples, colorSamples, internalFormat, width, height, fixedSampleLocations);
// }
// void goglTextureImage3DMultisampleCoverageNV(GLuint texture, GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLint internalFormat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedSampleLocations) {
// 	(*ptrglTextureImage3DMultisampleCoverageNV)(texture, target, coverageSamples, colorSamples, internalFormat, width, height, depth, fixedSampleLocations);
// }
// //  NV_transform_feedback
// void goglBeginTransformFeedbackNV(GLenum primitiveMode) {
// 	(*ptrglBeginTransformFeedbackNV)(primitiveMode);
// }
// void goglEndTransformFeedbackNV() {
// 	(*ptrglEndTransformFeedbackNV)();
// }
// void goglTransformFeedbackAttribsNV(GLsizei count, GLint* attribs, GLenum bufferMode) {
// 	(*ptrglTransformFeedbackAttribsNV)(count, attribs, bufferMode);
// }
// void goglBindBufferRangeNV(GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size) {
// 	(*ptrglBindBufferRangeNV)(target, index, buffer, offset, size);
// }
// void goglBindBufferOffsetNV(GLenum target, GLuint index, GLuint buffer, GLintptr offset) {
// 	(*ptrglBindBufferOffsetNV)(target, index, buffer, offset);
// }
// void goglBindBufferBaseNV(GLenum target, GLuint index, GLuint buffer) {
// 	(*ptrglBindBufferBaseNV)(target, index, buffer);
// }
// void goglTransformFeedbackVaryingsNV(GLuint program, GLsizei count, GLint* locations, GLenum bufferMode) {
// 	(*ptrglTransformFeedbackVaryingsNV)(program, count, locations, bufferMode);
// }
// void goglActiveVaryingNV(GLuint program, GLchar* name) {
// 	(*ptrglActiveVaryingNV)(program, name);
// }
// GLint goglGetVaryingLocationNV(GLuint program, GLchar* name) {
// 	return (*ptrglGetVaryingLocationNV)(program, name);
// }
// void goglGetActiveVaryingNV(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLsizei* size, GLenum* type_, GLchar* name) {
// 	(*ptrglGetActiveVaryingNV)(program, index, bufSize, length, size, type_, name);
// }
// void goglGetTransformFeedbackVaryingNV(GLuint program, GLuint index, GLint* location) {
// 	(*ptrglGetTransformFeedbackVaryingNV)(program, index, location);
// }
// void goglTransformFeedbackStreamAttribsNV(GLsizei count, GLint* attribs, GLsizei nbuffers, GLint* bufstreams, GLenum bufferMode) {
// 	(*ptrglTransformFeedbackStreamAttribsNV)(count, attribs, nbuffers, bufstreams, bufferMode);
// }
// //  NV_transform_feedback2
// void goglBindTransformFeedbackNV(GLenum target, GLuint id) {
// 	(*ptrglBindTransformFeedbackNV)(target, id);
// }
// void goglDeleteTransformFeedbacksNV(GLsizei n, GLuint* ids) {
// 	(*ptrglDeleteTransformFeedbacksNV)(n, ids);
// }
// void goglGenTransformFeedbacksNV(GLsizei n, GLuint* ids) {
// 	(*ptrglGenTransformFeedbacksNV)(n, ids);
// }
// GLboolean goglIsTransformFeedbackNV(GLuint id) {
// 	return (*ptrglIsTransformFeedbackNV)(id);
// }
// void goglPauseTransformFeedbackNV() {
// 	(*ptrglPauseTransformFeedbackNV)();
// }
// void goglResumeTransformFeedbackNV() {
// 	(*ptrglResumeTransformFeedbackNV)();
// }
// void goglDrawTransformFeedbackNV(GLenum mode, GLuint id) {
// 	(*ptrglDrawTransformFeedbackNV)(mode, id);
// }
// //  NV_vdpau_interop
// void goglVDPAUInitNV(GLvoid* vdpDevice, GLvoid* getProcAddress) {
// 	(*ptrglVDPAUInitNV)(vdpDevice, getProcAddress);
// }
// void goglVDPAUFiniNV() {
// 	(*ptrglVDPAUFiniNV)();
// }
// GLvdpauSurfaceNV goglVDPAURegisterVideoSurfaceNV(GLvoid* vdpSurface, GLenum target, GLsizei numTextureNames, GLuint* textureNames) {
// 	return (*ptrglVDPAURegisterVideoSurfaceNV)(vdpSurface, target, numTextureNames, textureNames);
// }
// GLvdpauSurfaceNV goglVDPAURegisterOutputSurfaceNV(GLvoid* vdpSurface, GLenum target, GLsizei numTextureNames, GLuint* textureNames) {
// 	return (*ptrglVDPAURegisterOutputSurfaceNV)(vdpSurface, target, numTextureNames, textureNames);
// }
// void goglVDPAUIsSurfaceNV(GLvdpauSurfaceNV surface) {
// 	(*ptrglVDPAUIsSurfaceNV)(surface);
// }
// void goglVDPAUUnregisterSurfaceNV(GLvdpauSurfaceNV surface) {
// 	(*ptrglVDPAUUnregisterSurfaceNV)(surface);
// }
// void goglVDPAUGetSurfaceivNV(GLvdpauSurfaceNV surface, GLenum pname, GLsizei bufSize, GLsizei* length, GLint* values) {
// 	(*ptrglVDPAUGetSurfaceivNV)(surface, pname, bufSize, length, values);
// }
// void goglVDPAUSurfaceAccessNV(GLvdpauSurfaceNV surface, GLenum access) {
// 	(*ptrglVDPAUSurfaceAccessNV)(surface, access);
// }
// void goglVDPAUMapSurfacesNV(GLsizei numSurfaces, GLvdpauSurfaceNV* surfaces) {
// 	(*ptrglVDPAUMapSurfacesNV)(numSurfaces, surfaces);
// }
// void goglVDPAUUnmapSurfacesNV(GLsizei numSurface, GLvdpauSurfaceNV* surfaces) {
// 	(*ptrglVDPAUUnmapSurfacesNV)(numSurface, surfaces);
// }
// //  NV_vertex_array_range
// void goglFlushVertexArrayRangeNV() {
// 	(*ptrglFlushVertexArrayRangeNV)();
// }
// void goglVertexArrayRangeNV(GLsizei length, GLvoid* pointer) {
// 	(*ptrglVertexArrayRangeNV)(length, pointer);
// }
// //  NV_vertex_attrib_integer_64bit
// void goglVertexAttribL1i64NV(GLuint index, GLint64EXT x) {
// 	(*ptrglVertexAttribL1i64NV)(index, x);
// }
// void goglVertexAttribL2i64NV(GLuint index, GLint64EXT x, GLint64EXT y) {
// 	(*ptrglVertexAttribL2i64NV)(index, x, y);
// }
// void goglVertexAttribL3i64NV(GLuint index, GLint64EXT x, GLint64EXT y, GLint64EXT z) {
// 	(*ptrglVertexAttribL3i64NV)(index, x, y, z);
// }
// void goglVertexAttribL4i64NV(GLuint index, GLint64EXT x, GLint64EXT y, GLint64EXT z, GLint64EXT w) {
// 	(*ptrglVertexAttribL4i64NV)(index, x, y, z, w);
// }
// void goglVertexAttribL1i64vNV(GLuint index, GLint64EXT* v) {
// 	(*ptrglVertexAttribL1i64vNV)(index, v);
// }
// void goglVertexAttribL2i64vNV(GLuint index, GLint64EXT* v) {
// 	(*ptrglVertexAttribL2i64vNV)(index, v);
// }
// void goglVertexAttribL3i64vNV(GLuint index, GLint64EXT* v) {
// 	(*ptrglVertexAttribL3i64vNV)(index, v);
// }
// void goglVertexAttribL4i64vNV(GLuint index, GLint64EXT* v) {
// 	(*ptrglVertexAttribL4i64vNV)(index, v);
// }
// void goglVertexAttribL1ui64NV(GLuint index, GLuint64EXT x) {
// 	(*ptrglVertexAttribL1ui64NV)(index, x);
// }
// void goglVertexAttribL2ui64NV(GLuint index, GLuint64EXT x, GLuint64EXT y) {
// 	(*ptrglVertexAttribL2ui64NV)(index, x, y);
// }
// void goglVertexAttribL3ui64NV(GLuint index, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z) {
// 	(*ptrglVertexAttribL3ui64NV)(index, x, y, z);
// }
// void goglVertexAttribL4ui64NV(GLuint index, GLuint64EXT x, GLuint64EXT y, GLuint64EXT z, GLuint64EXT w) {
// 	(*ptrglVertexAttribL4ui64NV)(index, x, y, z, w);
// }
// void goglVertexAttribL1ui64vNV(GLuint index, GLuint64EXT* v) {
// 	(*ptrglVertexAttribL1ui64vNV)(index, v);
// }
// void goglVertexAttribL2ui64vNV(GLuint index, GLuint64EXT* v) {
// 	(*ptrglVertexAttribL2ui64vNV)(index, v);
// }
// void goglVertexAttribL3ui64vNV(GLuint index, GLuint64EXT* v) {
// 	(*ptrglVertexAttribL3ui64vNV)(index, v);
// }
// void goglVertexAttribL4ui64vNV(GLuint index, GLuint64EXT* v) {
// 	(*ptrglVertexAttribL4ui64vNV)(index, v);
// }
// void goglGetVertexAttribLi64vNV(GLuint index, GLenum pname, GLint64EXT* params) {
// 	(*ptrglGetVertexAttribLi64vNV)(index, pname, params);
// }
// void goglGetVertexAttribLui64vNV(GLuint index, GLenum pname, GLuint64EXT* params) {
// 	(*ptrglGetVertexAttribLui64vNV)(index, pname, params);
// }
// void goglVertexAttribLFormatNV(GLuint index, GLint size, GLenum type_, GLsizei stride) {
// 	(*ptrglVertexAttribLFormatNV)(index, size, type_, stride);
// }
// //  NV_vertex_buffer_unified_memory
// void goglBufferAddressRangeNV(GLenum pname, GLuint index, GLuint64EXT address, GLsizeiptr length) {
// 	(*ptrglBufferAddressRangeNV)(pname, index, address, length);
// }
// void goglVertexFormatNV(GLint size, GLenum type_, GLsizei stride) {
// 	(*ptrglVertexFormatNV)(size, type_, stride);
// }
// void goglNormalFormatNV(GLenum type_, GLsizei stride) {
// 	(*ptrglNormalFormatNV)(type_, stride);
// }
// void goglColorFormatNV(GLint size, GLenum type_, GLsizei stride) {
// 	(*ptrglColorFormatNV)(size, type_, stride);
// }
// void goglIndexFormatNV(GLenum type_, GLsizei stride) {
// 	(*ptrglIndexFormatNV)(type_, stride);
// }
// void goglTexCoordFormatNV(GLint size, GLenum type_, GLsizei stride) {
// 	(*ptrglTexCoordFormatNV)(size, type_, stride);
// }
// void goglEdgeFlagFormatNV(GLsizei stride) {
// 	(*ptrglEdgeFlagFormatNV)(stride);
// }
// void goglSecondaryColorFormatNV(GLint size, GLenum type_, GLsizei stride) {
// 	(*ptrglSecondaryColorFormatNV)(size, type_, stride);
// }
// void goglFogCoordFormatNV(GLenum type_, GLsizei stride) {
// 	(*ptrglFogCoordFormatNV)(type_, stride);
// }
// void goglVertexAttribFormatNV(GLuint index, GLint size, GLenum type_, GLboolean normalized, GLsizei stride) {
// 	(*ptrglVertexAttribFormatNV)(index, size, type_, normalized, stride);
// }
// void goglVertexAttribIFormatNV(GLuint index, GLint size, GLenum type_, GLsizei stride) {
// 	(*ptrglVertexAttribIFormatNV)(index, size, type_, stride);
// }
// void goglGetIntegerui64i_vNV(GLenum value, GLuint index, GLuint64EXT* result) {
// 	(*ptrglGetIntegerui64i_vNV)(value, index, result);
// }
// //  NV_vertex_program
// GLboolean goglAreProgramsResidentNV(GLsizei n, GLuint* programs, GLboolean* residences) {
// 	return (*ptrglAreProgramsResidentNV)(n, programs, residences);
// }
// void goglBindProgramNV(GLenum target, GLuint id) {
// 	(*ptrglBindProgramNV)(target, id);
// }
// void goglDeleteProgramsNV(GLsizei n, GLuint* programs) {
// 	(*ptrglDeleteProgramsNV)(n, programs);
// }
// void goglExecuteProgramNV(GLenum target, GLuint id, GLfloat* params) {
// 	(*ptrglExecuteProgramNV)(target, id, params);
// }
// void goglGenProgramsNV(GLsizei n, GLuint* programs) {
// 	(*ptrglGenProgramsNV)(n, programs);
// }
// void goglGetProgramParameterdvNV(GLenum target, GLuint index, GLenum pname, GLdouble* params) {
// 	(*ptrglGetProgramParameterdvNV)(target, index, pname, params);
// }
// void goglGetProgramParameterfvNV(GLenum target, GLuint index, GLenum pname, GLfloat* params) {
// 	(*ptrglGetProgramParameterfvNV)(target, index, pname, params);
// }
// void goglGetProgramivNV(GLuint id, GLenum pname, GLint* params) {
// 	(*ptrglGetProgramivNV)(id, pname, params);
// }
// void goglGetProgramStringNV(GLuint id, GLenum pname, GLubyte* program) {
// 	(*ptrglGetProgramStringNV)(id, pname, program);
// }
// void goglGetTrackMatrixivNV(GLenum target, GLuint address, GLenum pname, GLint* params) {
// 	(*ptrglGetTrackMatrixivNV)(target, address, pname, params);
// }
// void goglGetVertexAttribdvNV(GLuint index, GLenum pname, GLdouble* params) {
// 	(*ptrglGetVertexAttribdvNV)(index, pname, params);
// }
// void goglGetVertexAttribfvNV(GLuint index, GLenum pname, GLfloat* params) {
// 	(*ptrglGetVertexAttribfvNV)(index, pname, params);
// }
// void goglGetVertexAttribivNV(GLuint index, GLenum pname, GLint* params) {
// 	(*ptrglGetVertexAttribivNV)(index, pname, params);
// }
// void goglGetVertexAttribPointervNV(GLuint index, GLenum pname, GLvoid** pointer) {
// 	(*ptrglGetVertexAttribPointervNV)(index, pname, pointer);
// }
// GLboolean goglIsProgramNV(GLuint id) {
// 	return (*ptrglIsProgramNV)(id);
// }
// void goglLoadProgramNV(GLenum target, GLuint id, GLsizei len, GLubyte* program) {
// 	(*ptrglLoadProgramNV)(target, id, len, program);
// }
// void goglProgramParameter4dNV(GLenum target, GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglProgramParameter4dNV)(target, index, x, y, z, w);
// }
// void goglProgramParameter4dvNV(GLenum target, GLuint index, GLdouble* v) {
// 	(*ptrglProgramParameter4dvNV)(target, index, v);
// }
// void goglProgramParameter4fNV(GLenum target, GLuint index, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglProgramParameter4fNV)(target, index, x, y, z, w);
// }
// void goglProgramParameter4fvNV(GLenum target, GLuint index, GLfloat* v) {
// 	(*ptrglProgramParameter4fvNV)(target, index, v);
// }
// void goglProgramParameters4dvNV(GLenum target, GLuint index, GLsizei count, GLdouble* v) {
// 	(*ptrglProgramParameters4dvNV)(target, index, count, v);
// }
// void goglProgramParameters4fvNV(GLenum target, GLuint index, GLsizei count, GLfloat* v) {
// 	(*ptrglProgramParameters4fvNV)(target, index, count, v);
// }
// void goglRequestResidentProgramsNV(GLsizei n, GLuint* programs) {
// 	(*ptrglRequestResidentProgramsNV)(n, programs);
// }
// void goglTrackMatrixNV(GLenum target, GLuint address, GLenum matrix, GLenum transform) {
// 	(*ptrglTrackMatrixNV)(target, address, matrix, transform);
// }
// void goglVertexAttribPointerNV(GLuint index, GLint fsize, GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglVertexAttribPointerNV)(index, fsize, type_, stride, pointer);
// }
// void goglVertexAttrib1dNV(GLuint index, GLdouble x) {
// 	(*ptrglVertexAttrib1dNV)(index, x);
// }
// void goglVertexAttrib1dvNV(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttrib1dvNV)(index, v);
// }
// void goglVertexAttrib1fNV(GLuint index, GLfloat x) {
// 	(*ptrglVertexAttrib1fNV)(index, x);
// }
// void goglVertexAttrib1fvNV(GLuint index, GLfloat* v) {
// 	(*ptrglVertexAttrib1fvNV)(index, v);
// }
// void goglVertexAttrib1sNV(GLuint index, GLshort x) {
// 	(*ptrglVertexAttrib1sNV)(index, x);
// }
// void goglVertexAttrib1svNV(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttrib1svNV)(index, v);
// }
// void goglVertexAttrib2dNV(GLuint index, GLdouble x, GLdouble y) {
// 	(*ptrglVertexAttrib2dNV)(index, x, y);
// }
// void goglVertexAttrib2dvNV(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttrib2dvNV)(index, v);
// }
// void goglVertexAttrib2fNV(GLuint index, GLfloat x, GLfloat y) {
// 	(*ptrglVertexAttrib2fNV)(index, x, y);
// }
// void goglVertexAttrib2fvNV(GLuint index, GLfloat* v) {
// 	(*ptrglVertexAttrib2fvNV)(index, v);
// }
// void goglVertexAttrib2sNV(GLuint index, GLshort x, GLshort y) {
// 	(*ptrglVertexAttrib2sNV)(index, x, y);
// }
// void goglVertexAttrib2svNV(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttrib2svNV)(index, v);
// }
// void goglVertexAttrib3dNV(GLuint index, GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglVertexAttrib3dNV)(index, x, y, z);
// }
// void goglVertexAttrib3dvNV(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttrib3dvNV)(index, v);
// }
// void goglVertexAttrib3fNV(GLuint index, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglVertexAttrib3fNV)(index, x, y, z);
// }
// void goglVertexAttrib3fvNV(GLuint index, GLfloat* v) {
// 	(*ptrglVertexAttrib3fvNV)(index, v);
// }
// void goglVertexAttrib3sNV(GLuint index, GLshort x, GLshort y, GLshort z) {
// 	(*ptrglVertexAttrib3sNV)(index, x, y, z);
// }
// void goglVertexAttrib3svNV(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttrib3svNV)(index, v);
// }
// void goglVertexAttrib4dNV(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglVertexAttrib4dNV)(index, x, y, z, w);
// }
// void goglVertexAttrib4dvNV(GLuint index, GLdouble* v) {
// 	(*ptrglVertexAttrib4dvNV)(index, v);
// }
// void goglVertexAttrib4fNV(GLuint index, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglVertexAttrib4fNV)(index, x, y, z, w);
// }
// void goglVertexAttrib4fvNV(GLuint index, GLfloat* v) {
// 	(*ptrglVertexAttrib4fvNV)(index, v);
// }
// void goglVertexAttrib4sNV(GLuint index, GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrglVertexAttrib4sNV)(index, x, y, z, w);
// }
// void goglVertexAttrib4svNV(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttrib4svNV)(index, v);
// }
// void goglVertexAttrib4ubNV(GLuint index, GLubyte x, GLubyte y, GLubyte z, GLubyte w) {
// 	(*ptrglVertexAttrib4ubNV)(index, x, y, z, w);
// }
// void goglVertexAttrib4ubvNV(GLuint index, GLubyte* v) {
// 	(*ptrglVertexAttrib4ubvNV)(index, v);
// }
// void goglVertexAttribs1dvNV(GLuint index, GLsizei count, GLdouble* v) {
// 	(*ptrglVertexAttribs1dvNV)(index, count, v);
// }
// void goglVertexAttribs1fvNV(GLuint index, GLsizei count, GLfloat* v) {
// 	(*ptrglVertexAttribs1fvNV)(index, count, v);
// }
// void goglVertexAttribs1svNV(GLuint index, GLsizei count, GLshort* v) {
// 	(*ptrglVertexAttribs1svNV)(index, count, v);
// }
// void goglVertexAttribs2dvNV(GLuint index, GLsizei count, GLdouble* v) {
// 	(*ptrglVertexAttribs2dvNV)(index, count, v);
// }
// void goglVertexAttribs2fvNV(GLuint index, GLsizei count, GLfloat* v) {
// 	(*ptrglVertexAttribs2fvNV)(index, count, v);
// }
// void goglVertexAttribs2svNV(GLuint index, GLsizei count, GLshort* v) {
// 	(*ptrglVertexAttribs2svNV)(index, count, v);
// }
// void goglVertexAttribs3dvNV(GLuint index, GLsizei count, GLdouble* v) {
// 	(*ptrglVertexAttribs3dvNV)(index, count, v);
// }
// void goglVertexAttribs3fvNV(GLuint index, GLsizei count, GLfloat* v) {
// 	(*ptrglVertexAttribs3fvNV)(index, count, v);
// }
// void goglVertexAttribs3svNV(GLuint index, GLsizei count, GLshort* v) {
// 	(*ptrglVertexAttribs3svNV)(index, count, v);
// }
// void goglVertexAttribs4dvNV(GLuint index, GLsizei count, GLdouble* v) {
// 	(*ptrglVertexAttribs4dvNV)(index, count, v);
// }
// void goglVertexAttribs4fvNV(GLuint index, GLsizei count, GLfloat* v) {
// 	(*ptrglVertexAttribs4fvNV)(index, count, v);
// }
// void goglVertexAttribs4svNV(GLuint index, GLsizei count, GLshort* v) {
// 	(*ptrglVertexAttribs4svNV)(index, count, v);
// }
// void goglVertexAttribs4ubvNV(GLuint index, GLsizei count, GLubyte* v) {
// 	(*ptrglVertexAttribs4ubvNV)(index, count, v);
// }
// //  NV_vertex_program4
// void goglVertexAttribI1iEXT(GLuint index, GLint x) {
// 	(*ptrglVertexAttribI1iEXT)(index, x);
// }
// void goglVertexAttribI2iEXT(GLuint index, GLint x, GLint y) {
// 	(*ptrglVertexAttribI2iEXT)(index, x, y);
// }
// void goglVertexAttribI3iEXT(GLuint index, GLint x, GLint y, GLint z) {
// 	(*ptrglVertexAttribI3iEXT)(index, x, y, z);
// }
// void goglVertexAttribI4iEXT(GLuint index, GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrglVertexAttribI4iEXT)(index, x, y, z, w);
// }
// void goglVertexAttribI1uiEXT(GLuint index, GLuint x) {
// 	(*ptrglVertexAttribI1uiEXT)(index, x);
// }
// void goglVertexAttribI2uiEXT(GLuint index, GLuint x, GLuint y) {
// 	(*ptrglVertexAttribI2uiEXT)(index, x, y);
// }
// void goglVertexAttribI3uiEXT(GLuint index, GLuint x, GLuint y, GLuint z) {
// 	(*ptrglVertexAttribI3uiEXT)(index, x, y, z);
// }
// void goglVertexAttribI4uiEXT(GLuint index, GLuint x, GLuint y, GLuint z, GLuint w) {
// 	(*ptrglVertexAttribI4uiEXT)(index, x, y, z, w);
// }
// void goglVertexAttribI1ivEXT(GLuint index, GLint* v) {
// 	(*ptrglVertexAttribI1ivEXT)(index, v);
// }
// void goglVertexAttribI2ivEXT(GLuint index, GLint* v) {
// 	(*ptrglVertexAttribI2ivEXT)(index, v);
// }
// void goglVertexAttribI3ivEXT(GLuint index, GLint* v) {
// 	(*ptrglVertexAttribI3ivEXT)(index, v);
// }
// void goglVertexAttribI4ivEXT(GLuint index, GLint* v) {
// 	(*ptrglVertexAttribI4ivEXT)(index, v);
// }
// void goglVertexAttribI1uivEXT(GLuint index, GLuint* v) {
// 	(*ptrglVertexAttribI1uivEXT)(index, v);
// }
// void goglVertexAttribI2uivEXT(GLuint index, GLuint* v) {
// 	(*ptrglVertexAttribI2uivEXT)(index, v);
// }
// void goglVertexAttribI3uivEXT(GLuint index, GLuint* v) {
// 	(*ptrglVertexAttribI3uivEXT)(index, v);
// }
// void goglVertexAttribI4uivEXT(GLuint index, GLuint* v) {
// 	(*ptrglVertexAttribI4uivEXT)(index, v);
// }
// void goglVertexAttribI4bvEXT(GLuint index, GLbyte* v) {
// 	(*ptrglVertexAttribI4bvEXT)(index, v);
// }
// void goglVertexAttribI4svEXT(GLuint index, GLshort* v) {
// 	(*ptrglVertexAttribI4svEXT)(index, v);
// }
// void goglVertexAttribI4ubvEXT(GLuint index, GLubyte* v) {
// 	(*ptrglVertexAttribI4ubvEXT)(index, v);
// }
// void goglVertexAttribI4usvEXT(GLuint index, GLushort* v) {
// 	(*ptrglVertexAttribI4usvEXT)(index, v);
// }
// void goglVertexAttribIPointerEXT(GLuint index, GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {
// 	(*ptrglVertexAttribIPointerEXT)(index, size, type_, stride, pointer);
// }
// void goglGetVertexAttribIivEXT(GLuint index, GLenum pname, GLint* params) {
// 	(*ptrglGetVertexAttribIivEXT)(index, pname, params);
// }
// void goglGetVertexAttribIuivEXT(GLuint index, GLenum pname, GLuint* params) {
// 	(*ptrglGetVertexAttribIuivEXT)(index, pname, params);
// }
// //  NV_video_capture
// void goglBeginVideoCaptureNV(GLuint video_capture_slot) {
// 	(*ptrglBeginVideoCaptureNV)(video_capture_slot);
// }
// void goglBindVideoCaptureStreamBufferNV(GLuint video_capture_slot, GLuint stream, GLenum frame_region, GLintptrARB offset) {
// 	(*ptrglBindVideoCaptureStreamBufferNV)(video_capture_slot, stream, frame_region, offset);
// }
// void goglBindVideoCaptureStreamTextureNV(GLuint video_capture_slot, GLuint stream, GLenum frame_region, GLenum target, GLuint texture) {
// 	(*ptrglBindVideoCaptureStreamTextureNV)(video_capture_slot, stream, frame_region, target, texture);
// }
// void goglEndVideoCaptureNV(GLuint video_capture_slot) {
// 	(*ptrglEndVideoCaptureNV)(video_capture_slot);
// }
// void goglGetVideoCaptureivNV(GLuint video_capture_slot, GLenum pname, GLint* params) {
// 	(*ptrglGetVideoCaptureivNV)(video_capture_slot, pname, params);
// }
// void goglGetVideoCaptureStreamivNV(GLuint video_capture_slot, GLuint stream, GLenum pname, GLint* params) {
// 	(*ptrglGetVideoCaptureStreamivNV)(video_capture_slot, stream, pname, params);
// }
// void goglGetVideoCaptureStreamfvNV(GLuint video_capture_slot, GLuint stream, GLenum pname, GLfloat* params) {
// 	(*ptrglGetVideoCaptureStreamfvNV)(video_capture_slot, stream, pname, params);
// }
// void goglGetVideoCaptureStreamdvNV(GLuint video_capture_slot, GLuint stream, GLenum pname, GLdouble* params) {
// 	(*ptrglGetVideoCaptureStreamdvNV)(video_capture_slot, stream, pname, params);
// }
// GLenum goglVideoCaptureNV(GLuint video_capture_slot, GLuint* sequence_num, GLuint64EXT* capture_time) {
// 	return (*ptrglVideoCaptureNV)(video_capture_slot, sequence_num, capture_time);
// }
// void goglVideoCaptureStreamParameterivNV(GLuint video_capture_slot, GLuint stream, GLenum pname, GLint* params) {
// 	(*ptrglVideoCaptureStreamParameterivNV)(video_capture_slot, stream, pname, params);
// }
// void goglVideoCaptureStreamParameterfvNV(GLuint video_capture_slot, GLuint stream, GLenum pname, GLfloat* params) {
// 	(*ptrglVideoCaptureStreamParameterfvNV)(video_capture_slot, stream, pname, params);
// }
// void goglVideoCaptureStreamParameterdvNV(GLuint video_capture_slot, GLuint stream, GLenum pname, GLdouble* params) {
// 	(*ptrglVideoCaptureStreamParameterdvNV)(video_capture_slot, stream, pname, params);
// }
// 
// int init_NV_bindless_texture() {
// 	ptrglGetTextureHandleNV = goglGetProcAddress("glGetTextureHandleNV");
// 	if(ptrglGetTextureHandleNV == NULL) return 1;
// 	ptrglGetTextureSamplerHandleNV = goglGetProcAddress("glGetTextureSamplerHandleNV");
// 	if(ptrglGetTextureSamplerHandleNV == NULL) return 1;
// 	ptrglMakeTextureHandleResidentNV = goglGetProcAddress("glMakeTextureHandleResidentNV");
// 	if(ptrglMakeTextureHandleResidentNV == NULL) return 1;
// 	ptrglMakeTextureHandleNonResidentNV = goglGetProcAddress("glMakeTextureHandleNonResidentNV");
// 	if(ptrglMakeTextureHandleNonResidentNV == NULL) return 1;
// 	ptrglGetImageHandleNV = goglGetProcAddress("glGetImageHandleNV");
// 	if(ptrglGetImageHandleNV == NULL) return 1;
// 	ptrglMakeImageHandleResidentNV = goglGetProcAddress("glMakeImageHandleResidentNV");
// 	if(ptrglMakeImageHandleResidentNV == NULL) return 1;
// 	ptrglMakeImageHandleNonResidentNV = goglGetProcAddress("glMakeImageHandleNonResidentNV");
// 	if(ptrglMakeImageHandleNonResidentNV == NULL) return 1;
// 	ptrglUniformHandleui64NV = goglGetProcAddress("glUniformHandleui64NV");
// 	if(ptrglUniformHandleui64NV == NULL) return 1;
// 	ptrglUniformHandleui64vNV = goglGetProcAddress("glUniformHandleui64vNV");
// 	if(ptrglUniformHandleui64vNV == NULL) return 1;
// 	ptrglProgramUniformHandleui64NV = goglGetProcAddress("glProgramUniformHandleui64NV");
// 	if(ptrglProgramUniformHandleui64NV == NULL) return 1;
// 	ptrglProgramUniformHandleui64vNV = goglGetProcAddress("glProgramUniformHandleui64vNV");
// 	if(ptrglProgramUniformHandleui64vNV == NULL) return 1;
// 	ptrglIsTextureHandleResidentNV = goglGetProcAddress("glIsTextureHandleResidentNV");
// 	if(ptrglIsTextureHandleResidentNV == NULL) return 1;
// 	ptrglIsImageHandleResidentNV = goglGetProcAddress("glIsImageHandleResidentNV");
// 	if(ptrglIsImageHandleResidentNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_conditional_render() {
// 	ptrglBeginConditionalRenderNV = goglGetProcAddress("glBeginConditionalRenderNV");
// 	if(ptrglBeginConditionalRenderNV == NULL) return 1;
// 	ptrglEndConditionalRenderNV = goglGetProcAddress("glEndConditionalRenderNV");
// 	if(ptrglEndConditionalRenderNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_copy_image() {
// 	ptrglCopyImageSubDataNV = goglGetProcAddress("glCopyImageSubDataNV");
// 	if(ptrglCopyImageSubDataNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_depth_buffer_float() {
// 	ptrglDepthRangedNV = goglGetProcAddress("glDepthRangedNV");
// 	if(ptrglDepthRangedNV == NULL) return 1;
// 	ptrglClearDepthdNV = goglGetProcAddress("glClearDepthdNV");
// 	if(ptrglClearDepthdNV == NULL) return 1;
// 	ptrglDepthBoundsdNV = goglGetProcAddress("glDepthBoundsdNV");
// 	if(ptrglDepthBoundsdNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_evaluators() {
// 	ptrglMapControlPointsNV = goglGetProcAddress("glMapControlPointsNV");
// 	if(ptrglMapControlPointsNV == NULL) return 1;
// 	ptrglMapParameterivNV = goglGetProcAddress("glMapParameterivNV");
// 	if(ptrglMapParameterivNV == NULL) return 1;
// 	ptrglMapParameterfvNV = goglGetProcAddress("glMapParameterfvNV");
// 	if(ptrglMapParameterfvNV == NULL) return 1;
// 	ptrglGetMapControlPointsNV = goglGetProcAddress("glGetMapControlPointsNV");
// 	if(ptrglGetMapControlPointsNV == NULL) return 1;
// 	ptrglGetMapParameterivNV = goglGetProcAddress("glGetMapParameterivNV");
// 	if(ptrglGetMapParameterivNV == NULL) return 1;
// 	ptrglGetMapParameterfvNV = goglGetProcAddress("glGetMapParameterfvNV");
// 	if(ptrglGetMapParameterfvNV == NULL) return 1;
// 	ptrglGetMapAttribParameterivNV = goglGetProcAddress("glGetMapAttribParameterivNV");
// 	if(ptrglGetMapAttribParameterivNV == NULL) return 1;
// 	ptrglGetMapAttribParameterfvNV = goglGetProcAddress("glGetMapAttribParameterfvNV");
// 	if(ptrglGetMapAttribParameterfvNV == NULL) return 1;
// 	ptrglEvalMapsNV = goglGetProcAddress("glEvalMapsNV");
// 	if(ptrglEvalMapsNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_explicit_multisample() {
// 	ptrglGetMultisamplefvNV = goglGetProcAddress("glGetMultisamplefvNV");
// 	if(ptrglGetMultisamplefvNV == NULL) return 1;
// 	ptrglSampleMaskIndexedNV = goglGetProcAddress("glSampleMaskIndexedNV");
// 	if(ptrglSampleMaskIndexedNV == NULL) return 1;
// 	ptrglTexRenderbufferNV = goglGetProcAddress("glTexRenderbufferNV");
// 	if(ptrglTexRenderbufferNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_fence() {
// 	ptrglDeleteFencesNV = goglGetProcAddress("glDeleteFencesNV");
// 	if(ptrglDeleteFencesNV == NULL) return 1;
// 	ptrglGenFencesNV = goglGetProcAddress("glGenFencesNV");
// 	if(ptrglGenFencesNV == NULL) return 1;
// 	ptrglIsFenceNV = goglGetProcAddress("glIsFenceNV");
// 	if(ptrglIsFenceNV == NULL) return 1;
// 	ptrglTestFenceNV = goglGetProcAddress("glTestFenceNV");
// 	if(ptrglTestFenceNV == NULL) return 1;
// 	ptrglGetFenceivNV = goglGetProcAddress("glGetFenceivNV");
// 	if(ptrglGetFenceivNV == NULL) return 1;
// 	ptrglFinishFenceNV = goglGetProcAddress("glFinishFenceNV");
// 	if(ptrglFinishFenceNV == NULL) return 1;
// 	ptrglSetFenceNV = goglGetProcAddress("glSetFenceNV");
// 	if(ptrglSetFenceNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_fragment_program() {
// 	ptrglProgramNamedParameter4fNV = goglGetProcAddress("glProgramNamedParameter4fNV");
// 	if(ptrglProgramNamedParameter4fNV == NULL) return 1;
// 	ptrglProgramNamedParameter4dNV = goglGetProcAddress("glProgramNamedParameter4dNV");
// 	if(ptrglProgramNamedParameter4dNV == NULL) return 1;
// 	ptrglProgramNamedParameter4fvNV = goglGetProcAddress("glProgramNamedParameter4fvNV");
// 	if(ptrglProgramNamedParameter4fvNV == NULL) return 1;
// 	ptrglProgramNamedParameter4dvNV = goglGetProcAddress("glProgramNamedParameter4dvNV");
// 	if(ptrglProgramNamedParameter4dvNV == NULL) return 1;
// 	ptrglGetProgramNamedParameterfvNV = goglGetProcAddress("glGetProgramNamedParameterfvNV");
// 	if(ptrglGetProgramNamedParameterfvNV == NULL) return 1;
// 	ptrglGetProgramNamedParameterdvNV = goglGetProcAddress("glGetProgramNamedParameterdvNV");
// 	if(ptrglGetProgramNamedParameterdvNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_framebuffer_multisample_coverage() {
// 	ptrglRenderbufferStorageMultisampleCoverageNV = goglGetProcAddress("glRenderbufferStorageMultisampleCoverageNV");
// 	if(ptrglRenderbufferStorageMultisampleCoverageNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_geometry_program4() {
// 	ptrglProgramVertexLimitNV = goglGetProcAddress("glProgramVertexLimitNV");
// 	if(ptrglProgramVertexLimitNV == NULL) return 1;
// 	ptrglFramebufferTextureEXT = goglGetProcAddress("glFramebufferTextureEXT");
// 	if(ptrglFramebufferTextureEXT == NULL) return 1;
// 	ptrglFramebufferTextureLayerEXT = goglGetProcAddress("glFramebufferTextureLayerEXT");
// 	if(ptrglFramebufferTextureLayerEXT == NULL) return 1;
// 	ptrglFramebufferTextureFaceEXT = goglGetProcAddress("glFramebufferTextureFaceEXT");
// 	if(ptrglFramebufferTextureFaceEXT == NULL) return 1;
// 	return 0;
// }
// int init_NV_gpu_program4() {
// 	ptrglProgramLocalParameterI4iNV = goglGetProcAddress("glProgramLocalParameterI4iNV");
// 	if(ptrglProgramLocalParameterI4iNV == NULL) return 1;
// 	ptrglProgramLocalParameterI4ivNV = goglGetProcAddress("glProgramLocalParameterI4ivNV");
// 	if(ptrglProgramLocalParameterI4ivNV == NULL) return 1;
// 	ptrglProgramLocalParametersI4ivNV = goglGetProcAddress("glProgramLocalParametersI4ivNV");
// 	if(ptrglProgramLocalParametersI4ivNV == NULL) return 1;
// 	ptrglProgramLocalParameterI4uiNV = goglGetProcAddress("glProgramLocalParameterI4uiNV");
// 	if(ptrglProgramLocalParameterI4uiNV == NULL) return 1;
// 	ptrglProgramLocalParameterI4uivNV = goglGetProcAddress("glProgramLocalParameterI4uivNV");
// 	if(ptrglProgramLocalParameterI4uivNV == NULL) return 1;
// 	ptrglProgramLocalParametersI4uivNV = goglGetProcAddress("glProgramLocalParametersI4uivNV");
// 	if(ptrglProgramLocalParametersI4uivNV == NULL) return 1;
// 	ptrglProgramEnvParameterI4iNV = goglGetProcAddress("glProgramEnvParameterI4iNV");
// 	if(ptrglProgramEnvParameterI4iNV == NULL) return 1;
// 	ptrglProgramEnvParameterI4ivNV = goglGetProcAddress("glProgramEnvParameterI4ivNV");
// 	if(ptrglProgramEnvParameterI4ivNV == NULL) return 1;
// 	ptrglProgramEnvParametersI4ivNV = goglGetProcAddress("glProgramEnvParametersI4ivNV");
// 	if(ptrglProgramEnvParametersI4ivNV == NULL) return 1;
// 	ptrglProgramEnvParameterI4uiNV = goglGetProcAddress("glProgramEnvParameterI4uiNV");
// 	if(ptrglProgramEnvParameterI4uiNV == NULL) return 1;
// 	ptrglProgramEnvParameterI4uivNV = goglGetProcAddress("glProgramEnvParameterI4uivNV");
// 	if(ptrglProgramEnvParameterI4uivNV == NULL) return 1;
// 	ptrglProgramEnvParametersI4uivNV = goglGetProcAddress("glProgramEnvParametersI4uivNV");
// 	if(ptrglProgramEnvParametersI4uivNV == NULL) return 1;
// 	ptrglGetProgramLocalParameterIivNV = goglGetProcAddress("glGetProgramLocalParameterIivNV");
// 	if(ptrglGetProgramLocalParameterIivNV == NULL) return 1;
// 	ptrglGetProgramLocalParameterIuivNV = goglGetProcAddress("glGetProgramLocalParameterIuivNV");
// 	if(ptrglGetProgramLocalParameterIuivNV == NULL) return 1;
// 	ptrglGetProgramEnvParameterIivNV = goglGetProcAddress("glGetProgramEnvParameterIivNV");
// 	if(ptrglGetProgramEnvParameterIivNV == NULL) return 1;
// 	ptrglGetProgramEnvParameterIuivNV = goglGetProcAddress("glGetProgramEnvParameterIuivNV");
// 	if(ptrglGetProgramEnvParameterIuivNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_gpu_program5() {
// 	ptrglProgramSubroutineParametersuivNV = goglGetProcAddress("glProgramSubroutineParametersuivNV");
// 	if(ptrglProgramSubroutineParametersuivNV == NULL) return 1;
// 	ptrglGetProgramSubroutineParameteruivNV = goglGetProcAddress("glGetProgramSubroutineParameteruivNV");
// 	if(ptrglGetProgramSubroutineParameteruivNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_gpu_shader5() {
// 	ptrglUniform1i64NV = goglGetProcAddress("glUniform1i64NV");
// 	if(ptrglUniform1i64NV == NULL) return 1;
// 	ptrglUniform2i64NV = goglGetProcAddress("glUniform2i64NV");
// 	if(ptrglUniform2i64NV == NULL) return 1;
// 	ptrglUniform3i64NV = goglGetProcAddress("glUniform3i64NV");
// 	if(ptrglUniform3i64NV == NULL) return 1;
// 	ptrglUniform4i64NV = goglGetProcAddress("glUniform4i64NV");
// 	if(ptrglUniform4i64NV == NULL) return 1;
// 	ptrglUniform1i64vNV = goglGetProcAddress("glUniform1i64vNV");
// 	if(ptrglUniform1i64vNV == NULL) return 1;
// 	ptrglUniform2i64vNV = goglGetProcAddress("glUniform2i64vNV");
// 	if(ptrglUniform2i64vNV == NULL) return 1;
// 	ptrglUniform3i64vNV = goglGetProcAddress("glUniform3i64vNV");
// 	if(ptrglUniform3i64vNV == NULL) return 1;
// 	ptrglUniform4i64vNV = goglGetProcAddress("glUniform4i64vNV");
// 	if(ptrglUniform4i64vNV == NULL) return 1;
// 	ptrglUniform1ui64NV = goglGetProcAddress("glUniform1ui64NV");
// 	if(ptrglUniform1ui64NV == NULL) return 1;
// 	ptrglUniform2ui64NV = goglGetProcAddress("glUniform2ui64NV");
// 	if(ptrglUniform2ui64NV == NULL) return 1;
// 	ptrglUniform3ui64NV = goglGetProcAddress("glUniform3ui64NV");
// 	if(ptrglUniform3ui64NV == NULL) return 1;
// 	ptrglUniform4ui64NV = goglGetProcAddress("glUniform4ui64NV");
// 	if(ptrglUniform4ui64NV == NULL) return 1;
// 	ptrglUniform1ui64vNV = goglGetProcAddress("glUniform1ui64vNV");
// 	if(ptrglUniform1ui64vNV == NULL) return 1;
// 	ptrglUniform2ui64vNV = goglGetProcAddress("glUniform2ui64vNV");
// 	if(ptrglUniform2ui64vNV == NULL) return 1;
// 	ptrglUniform3ui64vNV = goglGetProcAddress("glUniform3ui64vNV");
// 	if(ptrglUniform3ui64vNV == NULL) return 1;
// 	ptrglUniform4ui64vNV = goglGetProcAddress("glUniform4ui64vNV");
// 	if(ptrglUniform4ui64vNV == NULL) return 1;
// 	ptrglGetUniformi64vNV = goglGetProcAddress("glGetUniformi64vNV");
// 	if(ptrglGetUniformi64vNV == NULL) return 1;
// 	ptrglProgramUniform1i64NV = goglGetProcAddress("glProgramUniform1i64NV");
// 	if(ptrglProgramUniform1i64NV == NULL) return 1;
// 	ptrglProgramUniform2i64NV = goglGetProcAddress("glProgramUniform2i64NV");
// 	if(ptrglProgramUniform2i64NV == NULL) return 1;
// 	ptrglProgramUniform3i64NV = goglGetProcAddress("glProgramUniform3i64NV");
// 	if(ptrglProgramUniform3i64NV == NULL) return 1;
// 	ptrglProgramUniform4i64NV = goglGetProcAddress("glProgramUniform4i64NV");
// 	if(ptrglProgramUniform4i64NV == NULL) return 1;
// 	ptrglProgramUniform1i64vNV = goglGetProcAddress("glProgramUniform1i64vNV");
// 	if(ptrglProgramUniform1i64vNV == NULL) return 1;
// 	ptrglProgramUniform2i64vNV = goglGetProcAddress("glProgramUniform2i64vNV");
// 	if(ptrglProgramUniform2i64vNV == NULL) return 1;
// 	ptrglProgramUniform3i64vNV = goglGetProcAddress("glProgramUniform3i64vNV");
// 	if(ptrglProgramUniform3i64vNV == NULL) return 1;
// 	ptrglProgramUniform4i64vNV = goglGetProcAddress("glProgramUniform4i64vNV");
// 	if(ptrglProgramUniform4i64vNV == NULL) return 1;
// 	ptrglProgramUniform1ui64NV = goglGetProcAddress("glProgramUniform1ui64NV");
// 	if(ptrglProgramUniform1ui64NV == NULL) return 1;
// 	ptrglProgramUniform2ui64NV = goglGetProcAddress("glProgramUniform2ui64NV");
// 	if(ptrglProgramUniform2ui64NV == NULL) return 1;
// 	ptrglProgramUniform3ui64NV = goglGetProcAddress("glProgramUniform3ui64NV");
// 	if(ptrglProgramUniform3ui64NV == NULL) return 1;
// 	ptrglProgramUniform4ui64NV = goglGetProcAddress("glProgramUniform4ui64NV");
// 	if(ptrglProgramUniform4ui64NV == NULL) return 1;
// 	ptrglProgramUniform1ui64vNV = goglGetProcAddress("glProgramUniform1ui64vNV");
// 	if(ptrglProgramUniform1ui64vNV == NULL) return 1;
// 	ptrglProgramUniform2ui64vNV = goglGetProcAddress("glProgramUniform2ui64vNV");
// 	if(ptrglProgramUniform2ui64vNV == NULL) return 1;
// 	ptrglProgramUniform3ui64vNV = goglGetProcAddress("glProgramUniform3ui64vNV");
// 	if(ptrglProgramUniform3ui64vNV == NULL) return 1;
// 	ptrglProgramUniform4ui64vNV = goglGetProcAddress("glProgramUniform4ui64vNV");
// 	if(ptrglProgramUniform4ui64vNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_half_float() {
// 	ptrglVertex2hNV = goglGetProcAddress("glVertex2hNV");
// 	if(ptrglVertex2hNV == NULL) return 1;
// 	ptrglVertex2hvNV = goglGetProcAddress("glVertex2hvNV");
// 	if(ptrglVertex2hvNV == NULL) return 1;
// 	ptrglVertex3hNV = goglGetProcAddress("glVertex3hNV");
// 	if(ptrglVertex3hNV == NULL) return 1;
// 	ptrglVertex3hvNV = goglGetProcAddress("glVertex3hvNV");
// 	if(ptrglVertex3hvNV == NULL) return 1;
// 	ptrglVertex4hNV = goglGetProcAddress("glVertex4hNV");
// 	if(ptrglVertex4hNV == NULL) return 1;
// 	ptrglVertex4hvNV = goglGetProcAddress("glVertex4hvNV");
// 	if(ptrglVertex4hvNV == NULL) return 1;
// 	ptrglNormal3hNV = goglGetProcAddress("glNormal3hNV");
// 	if(ptrglNormal3hNV == NULL) return 1;
// 	ptrglNormal3hvNV = goglGetProcAddress("glNormal3hvNV");
// 	if(ptrglNormal3hvNV == NULL) return 1;
// 	ptrglColor3hNV = goglGetProcAddress("glColor3hNV");
// 	if(ptrglColor3hNV == NULL) return 1;
// 	ptrglColor3hvNV = goglGetProcAddress("glColor3hvNV");
// 	if(ptrglColor3hvNV == NULL) return 1;
// 	ptrglColor4hNV = goglGetProcAddress("glColor4hNV");
// 	if(ptrglColor4hNV == NULL) return 1;
// 	ptrglColor4hvNV = goglGetProcAddress("glColor4hvNV");
// 	if(ptrglColor4hvNV == NULL) return 1;
// 	ptrglTexCoord1hNV = goglGetProcAddress("glTexCoord1hNV");
// 	if(ptrglTexCoord1hNV == NULL) return 1;
// 	ptrglTexCoord1hvNV = goglGetProcAddress("glTexCoord1hvNV");
// 	if(ptrglTexCoord1hvNV == NULL) return 1;
// 	ptrglTexCoord2hNV = goglGetProcAddress("glTexCoord2hNV");
// 	if(ptrglTexCoord2hNV == NULL) return 1;
// 	ptrglTexCoord2hvNV = goglGetProcAddress("glTexCoord2hvNV");
// 	if(ptrglTexCoord2hvNV == NULL) return 1;
// 	ptrglTexCoord3hNV = goglGetProcAddress("glTexCoord3hNV");
// 	if(ptrglTexCoord3hNV == NULL) return 1;
// 	ptrglTexCoord3hvNV = goglGetProcAddress("glTexCoord3hvNV");
// 	if(ptrglTexCoord3hvNV == NULL) return 1;
// 	ptrglTexCoord4hNV = goglGetProcAddress("glTexCoord4hNV");
// 	if(ptrglTexCoord4hNV == NULL) return 1;
// 	ptrglTexCoord4hvNV = goglGetProcAddress("glTexCoord4hvNV");
// 	if(ptrglTexCoord4hvNV == NULL) return 1;
// 	ptrglMultiTexCoord1hNV = goglGetProcAddress("glMultiTexCoord1hNV");
// 	if(ptrglMultiTexCoord1hNV == NULL) return 1;
// 	ptrglMultiTexCoord1hvNV = goglGetProcAddress("glMultiTexCoord1hvNV");
// 	if(ptrglMultiTexCoord1hvNV == NULL) return 1;
// 	ptrglMultiTexCoord2hNV = goglGetProcAddress("glMultiTexCoord2hNV");
// 	if(ptrglMultiTexCoord2hNV == NULL) return 1;
// 	ptrglMultiTexCoord2hvNV = goglGetProcAddress("glMultiTexCoord2hvNV");
// 	if(ptrglMultiTexCoord2hvNV == NULL) return 1;
// 	ptrglMultiTexCoord3hNV = goglGetProcAddress("glMultiTexCoord3hNV");
// 	if(ptrglMultiTexCoord3hNV == NULL) return 1;
// 	ptrglMultiTexCoord3hvNV = goglGetProcAddress("glMultiTexCoord3hvNV");
// 	if(ptrglMultiTexCoord3hvNV == NULL) return 1;
// 	ptrglMultiTexCoord4hNV = goglGetProcAddress("glMultiTexCoord4hNV");
// 	if(ptrglMultiTexCoord4hNV == NULL) return 1;
// 	ptrglMultiTexCoord4hvNV = goglGetProcAddress("glMultiTexCoord4hvNV");
// 	if(ptrglMultiTexCoord4hvNV == NULL) return 1;
// 	ptrglFogCoordhNV = goglGetProcAddress("glFogCoordhNV");
// 	if(ptrglFogCoordhNV == NULL) return 1;
// 	ptrglFogCoordhvNV = goglGetProcAddress("glFogCoordhvNV");
// 	if(ptrglFogCoordhvNV == NULL) return 1;
// 	ptrglSecondaryColor3hNV = goglGetProcAddress("glSecondaryColor3hNV");
// 	if(ptrglSecondaryColor3hNV == NULL) return 1;
// 	ptrglSecondaryColor3hvNV = goglGetProcAddress("glSecondaryColor3hvNV");
// 	if(ptrglSecondaryColor3hvNV == NULL) return 1;
// 	ptrglVertexWeighthNV = goglGetProcAddress("glVertexWeighthNV");
// 	if(ptrglVertexWeighthNV == NULL) return 1;
// 	ptrglVertexWeighthvNV = goglGetProcAddress("glVertexWeighthvNV");
// 	if(ptrglVertexWeighthvNV == NULL) return 1;
// 	ptrglVertexAttrib1hNV = goglGetProcAddress("glVertexAttrib1hNV");
// 	if(ptrglVertexAttrib1hNV == NULL) return 1;
// 	ptrglVertexAttrib1hvNV = goglGetProcAddress("glVertexAttrib1hvNV");
// 	if(ptrglVertexAttrib1hvNV == NULL) return 1;
// 	ptrglVertexAttrib2hNV = goglGetProcAddress("glVertexAttrib2hNV");
// 	if(ptrglVertexAttrib2hNV == NULL) return 1;
// 	ptrglVertexAttrib2hvNV = goglGetProcAddress("glVertexAttrib2hvNV");
// 	if(ptrglVertexAttrib2hvNV == NULL) return 1;
// 	ptrglVertexAttrib3hNV = goglGetProcAddress("glVertexAttrib3hNV");
// 	if(ptrglVertexAttrib3hNV == NULL) return 1;
// 	ptrglVertexAttrib3hvNV = goglGetProcAddress("glVertexAttrib3hvNV");
// 	if(ptrglVertexAttrib3hvNV == NULL) return 1;
// 	ptrglVertexAttrib4hNV = goglGetProcAddress("glVertexAttrib4hNV");
// 	if(ptrglVertexAttrib4hNV == NULL) return 1;
// 	ptrglVertexAttrib4hvNV = goglGetProcAddress("glVertexAttrib4hvNV");
// 	if(ptrglVertexAttrib4hvNV == NULL) return 1;
// 	ptrglVertexAttribs1hvNV = goglGetProcAddress("glVertexAttribs1hvNV");
// 	if(ptrglVertexAttribs1hvNV == NULL) return 1;
// 	ptrglVertexAttribs2hvNV = goglGetProcAddress("glVertexAttribs2hvNV");
// 	if(ptrglVertexAttribs2hvNV == NULL) return 1;
// 	ptrglVertexAttribs3hvNV = goglGetProcAddress("glVertexAttribs3hvNV");
// 	if(ptrglVertexAttribs3hvNV == NULL) return 1;
// 	ptrglVertexAttribs4hvNV = goglGetProcAddress("glVertexAttribs4hvNV");
// 	if(ptrglVertexAttribs4hvNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_occlusion_query() {
// 	ptrglGenOcclusionQueriesNV = goglGetProcAddress("glGenOcclusionQueriesNV");
// 	if(ptrglGenOcclusionQueriesNV == NULL) return 1;
// 	ptrglDeleteOcclusionQueriesNV = goglGetProcAddress("glDeleteOcclusionQueriesNV");
// 	if(ptrglDeleteOcclusionQueriesNV == NULL) return 1;
// 	ptrglIsOcclusionQueryNV = goglGetProcAddress("glIsOcclusionQueryNV");
// 	if(ptrglIsOcclusionQueryNV == NULL) return 1;
// 	ptrglBeginOcclusionQueryNV = goglGetProcAddress("glBeginOcclusionQueryNV");
// 	if(ptrglBeginOcclusionQueryNV == NULL) return 1;
// 	ptrglEndOcclusionQueryNV = goglGetProcAddress("glEndOcclusionQueryNV");
// 	if(ptrglEndOcclusionQueryNV == NULL) return 1;
// 	ptrglGetOcclusionQueryivNV = goglGetProcAddress("glGetOcclusionQueryivNV");
// 	if(ptrglGetOcclusionQueryivNV == NULL) return 1;
// 	ptrglGetOcclusionQueryuivNV = goglGetProcAddress("glGetOcclusionQueryuivNV");
// 	if(ptrglGetOcclusionQueryuivNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_parameter_buffer_object() {
// 	ptrglProgramBufferParametersfvNV = goglGetProcAddress("glProgramBufferParametersfvNV");
// 	if(ptrglProgramBufferParametersfvNV == NULL) return 1;
// 	ptrglProgramBufferParametersIivNV = goglGetProcAddress("glProgramBufferParametersIivNV");
// 	if(ptrglProgramBufferParametersIivNV == NULL) return 1;
// 	ptrglProgramBufferParametersIuivNV = goglGetProcAddress("glProgramBufferParametersIuivNV");
// 	if(ptrglProgramBufferParametersIuivNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_path_rendering() {
// 	ptrglGenPathsNV = goglGetProcAddress("glGenPathsNV");
// 	if(ptrglGenPathsNV == NULL) return 1;
// 	ptrglDeletePathsNV = goglGetProcAddress("glDeletePathsNV");
// 	if(ptrglDeletePathsNV == NULL) return 1;
// 	ptrglIsPathNV = goglGetProcAddress("glIsPathNV");
// 	if(ptrglIsPathNV == NULL) return 1;
// 	ptrglPathCommandsNV = goglGetProcAddress("glPathCommandsNV");
// 	if(ptrglPathCommandsNV == NULL) return 1;
// 	ptrglPathCoordsNV = goglGetProcAddress("glPathCoordsNV");
// 	if(ptrglPathCoordsNV == NULL) return 1;
// 	ptrglPathSubCommandsNV = goglGetProcAddress("glPathSubCommandsNV");
// 	if(ptrglPathSubCommandsNV == NULL) return 1;
// 	ptrglPathSubCoordsNV = goglGetProcAddress("glPathSubCoordsNV");
// 	if(ptrglPathSubCoordsNV == NULL) return 1;
// 	ptrglPathStringNV = goglGetProcAddress("glPathStringNV");
// 	if(ptrglPathStringNV == NULL) return 1;
// 	ptrglPathGlyphsNV = goglGetProcAddress("glPathGlyphsNV");
// 	if(ptrglPathGlyphsNV == NULL) return 1;
// 	ptrglPathGlyphRangeNV = goglGetProcAddress("glPathGlyphRangeNV");
// 	if(ptrglPathGlyphRangeNV == NULL) return 1;
// 	ptrglWeightPathsNV = goglGetProcAddress("glWeightPathsNV");
// 	if(ptrglWeightPathsNV == NULL) return 1;
// 	ptrglCopyPathNV = goglGetProcAddress("glCopyPathNV");
// 	if(ptrglCopyPathNV == NULL) return 1;
// 	ptrglInterpolatePathsNV = goglGetProcAddress("glInterpolatePathsNV");
// 	if(ptrglInterpolatePathsNV == NULL) return 1;
// 	ptrglTransformPathNV = goglGetProcAddress("glTransformPathNV");
// 	if(ptrglTransformPathNV == NULL) return 1;
// 	ptrglPathParameterivNV = goglGetProcAddress("glPathParameterivNV");
// 	if(ptrglPathParameterivNV == NULL) return 1;
// 	ptrglPathParameteriNV = goglGetProcAddress("glPathParameteriNV");
// 	if(ptrglPathParameteriNV == NULL) return 1;
// 	ptrglPathParameterfvNV = goglGetProcAddress("glPathParameterfvNV");
// 	if(ptrglPathParameterfvNV == NULL) return 1;
// 	ptrglPathParameterfNV = goglGetProcAddress("glPathParameterfNV");
// 	if(ptrglPathParameterfNV == NULL) return 1;
// 	ptrglPathDashArrayNV = goglGetProcAddress("glPathDashArrayNV");
// 	if(ptrglPathDashArrayNV == NULL) return 1;
// 	ptrglPathStencilFuncNV = goglGetProcAddress("glPathStencilFuncNV");
// 	if(ptrglPathStencilFuncNV == NULL) return 1;
// 	ptrglPathStencilDepthOffsetNV = goglGetProcAddress("glPathStencilDepthOffsetNV");
// 	if(ptrglPathStencilDepthOffsetNV == NULL) return 1;
// 	ptrglStencilFillPathNV = goglGetProcAddress("glStencilFillPathNV");
// 	if(ptrglStencilFillPathNV == NULL) return 1;
// 	ptrglStencilStrokePathNV = goglGetProcAddress("glStencilStrokePathNV");
// 	if(ptrglStencilStrokePathNV == NULL) return 1;
// 	ptrglStencilFillPathInstancedNV = goglGetProcAddress("glStencilFillPathInstancedNV");
// 	if(ptrglStencilFillPathInstancedNV == NULL) return 1;
// 	ptrglStencilStrokePathInstancedNV = goglGetProcAddress("glStencilStrokePathInstancedNV");
// 	if(ptrglStencilStrokePathInstancedNV == NULL) return 1;
// 	ptrglPathCoverDepthFuncNV = goglGetProcAddress("glPathCoverDepthFuncNV");
// 	if(ptrglPathCoverDepthFuncNV == NULL) return 1;
// 	ptrglPathColorGenNV = goglGetProcAddress("glPathColorGenNV");
// 	if(ptrglPathColorGenNV == NULL) return 1;
// 	ptrglPathTexGenNV = goglGetProcAddress("glPathTexGenNV");
// 	if(ptrglPathTexGenNV == NULL) return 1;
// 	ptrglPathFogGenNV = goglGetProcAddress("glPathFogGenNV");
// 	if(ptrglPathFogGenNV == NULL) return 1;
// 	ptrglCoverFillPathNV = goglGetProcAddress("glCoverFillPathNV");
// 	if(ptrglCoverFillPathNV == NULL) return 1;
// 	ptrglCoverStrokePathNV = goglGetProcAddress("glCoverStrokePathNV");
// 	if(ptrglCoverStrokePathNV == NULL) return 1;
// 	ptrglCoverFillPathInstancedNV = goglGetProcAddress("glCoverFillPathInstancedNV");
// 	if(ptrglCoverFillPathInstancedNV == NULL) return 1;
// 	ptrglCoverStrokePathInstancedNV = goglGetProcAddress("glCoverStrokePathInstancedNV");
// 	if(ptrglCoverStrokePathInstancedNV == NULL) return 1;
// 	ptrglGetPathParameterivNV = goglGetProcAddress("glGetPathParameterivNV");
// 	if(ptrglGetPathParameterivNV == NULL) return 1;
// 	ptrglGetPathParameterfvNV = goglGetProcAddress("glGetPathParameterfvNV");
// 	if(ptrglGetPathParameterfvNV == NULL) return 1;
// 	ptrglGetPathCommandsNV = goglGetProcAddress("glGetPathCommandsNV");
// 	if(ptrglGetPathCommandsNV == NULL) return 1;
// 	ptrglGetPathCoordsNV = goglGetProcAddress("glGetPathCoordsNV");
// 	if(ptrglGetPathCoordsNV == NULL) return 1;
// 	ptrglGetPathDashArrayNV = goglGetProcAddress("glGetPathDashArrayNV");
// 	if(ptrglGetPathDashArrayNV == NULL) return 1;
// 	ptrglGetPathMetricsNV = goglGetProcAddress("glGetPathMetricsNV");
// 	if(ptrglGetPathMetricsNV == NULL) return 1;
// 	ptrglGetPathMetricRangeNV = goglGetProcAddress("glGetPathMetricRangeNV");
// 	if(ptrglGetPathMetricRangeNV == NULL) return 1;
// 	ptrglGetPathSpacingNV = goglGetProcAddress("glGetPathSpacingNV");
// 	if(ptrglGetPathSpacingNV == NULL) return 1;
// 	ptrglGetPathColorGenivNV = goglGetProcAddress("glGetPathColorGenivNV");
// 	if(ptrglGetPathColorGenivNV == NULL) return 1;
// 	ptrglGetPathColorGenfvNV = goglGetProcAddress("glGetPathColorGenfvNV");
// 	if(ptrglGetPathColorGenfvNV == NULL) return 1;
// 	ptrglGetPathTexGenivNV = goglGetProcAddress("glGetPathTexGenivNV");
// 	if(ptrglGetPathTexGenivNV == NULL) return 1;
// 	ptrglGetPathTexGenfvNV = goglGetProcAddress("glGetPathTexGenfvNV");
// 	if(ptrglGetPathTexGenfvNV == NULL) return 1;
// 	ptrglIsPointInFillPathNV = goglGetProcAddress("glIsPointInFillPathNV");
// 	if(ptrglIsPointInFillPathNV == NULL) return 1;
// 	ptrglIsPointInStrokePathNV = goglGetProcAddress("glIsPointInStrokePathNV");
// 	if(ptrglIsPointInStrokePathNV == NULL) return 1;
// 	ptrglGetPathLengthNV = goglGetProcAddress("glGetPathLengthNV");
// 	if(ptrglGetPathLengthNV == NULL) return 1;
// 	ptrglPointAlongPathNV = goglGetProcAddress("glPointAlongPathNV");
// 	if(ptrglPointAlongPathNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_pixel_data_range() {
// 	ptrglPixelDataRangeNV = goglGetProcAddress("glPixelDataRangeNV");
// 	if(ptrglPixelDataRangeNV == NULL) return 1;
// 	ptrglFlushPixelDataRangeNV = goglGetProcAddress("glFlushPixelDataRangeNV");
// 	if(ptrglFlushPixelDataRangeNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_point_sprite() {
// 	ptrglPointParameteriNV = goglGetProcAddress("glPointParameteriNV");
// 	if(ptrglPointParameteriNV == NULL) return 1;
// 	ptrglPointParameterivNV = goglGetProcAddress("glPointParameterivNV");
// 	if(ptrglPointParameterivNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_present_video() {
// 	ptrglPresentFrameKeyedNV = goglGetProcAddress("glPresentFrameKeyedNV");
// 	if(ptrglPresentFrameKeyedNV == NULL) return 1;
// 	ptrglPresentFrameDualFillNV = goglGetProcAddress("glPresentFrameDualFillNV");
// 	if(ptrglPresentFrameDualFillNV == NULL) return 1;
// 	ptrglGetVideoivNV = goglGetProcAddress("glGetVideoivNV");
// 	if(ptrglGetVideoivNV == NULL) return 1;
// 	ptrglGetVideouivNV = goglGetProcAddress("glGetVideouivNV");
// 	if(ptrglGetVideouivNV == NULL) return 1;
// 	ptrglGetVideoi64vNV = goglGetProcAddress("glGetVideoi64vNV");
// 	if(ptrglGetVideoi64vNV == NULL) return 1;
// 	ptrglGetVideoui64vNV = goglGetProcAddress("glGetVideoui64vNV");
// 	if(ptrglGetVideoui64vNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_primitive_restart() {
// 	ptrglPrimitiveRestartNV = goglGetProcAddress("glPrimitiveRestartNV");
// 	if(ptrglPrimitiveRestartNV == NULL) return 1;
// 	ptrglPrimitiveRestartIndexNV = goglGetProcAddress("glPrimitiveRestartIndexNV");
// 	if(ptrglPrimitiveRestartIndexNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_register_combiners() {
// 	ptrglCombinerParameterfvNV = goglGetProcAddress("glCombinerParameterfvNV");
// 	if(ptrglCombinerParameterfvNV == NULL) return 1;
// 	ptrglCombinerParameterfNV = goglGetProcAddress("glCombinerParameterfNV");
// 	if(ptrglCombinerParameterfNV == NULL) return 1;
// 	ptrglCombinerParameterivNV = goglGetProcAddress("glCombinerParameterivNV");
// 	if(ptrglCombinerParameterivNV == NULL) return 1;
// 	ptrglCombinerParameteriNV = goglGetProcAddress("glCombinerParameteriNV");
// 	if(ptrglCombinerParameteriNV == NULL) return 1;
// 	ptrglCombinerInputNV = goglGetProcAddress("glCombinerInputNV");
// 	if(ptrglCombinerInputNV == NULL) return 1;
// 	ptrglCombinerOutputNV = goglGetProcAddress("glCombinerOutputNV");
// 	if(ptrglCombinerOutputNV == NULL) return 1;
// 	ptrglFinalCombinerInputNV = goglGetProcAddress("glFinalCombinerInputNV");
// 	if(ptrglFinalCombinerInputNV == NULL) return 1;
// 	ptrglGetCombinerInputParameterfvNV = goglGetProcAddress("glGetCombinerInputParameterfvNV");
// 	if(ptrglGetCombinerInputParameterfvNV == NULL) return 1;
// 	ptrglGetCombinerInputParameterivNV = goglGetProcAddress("glGetCombinerInputParameterivNV");
// 	if(ptrglGetCombinerInputParameterivNV == NULL) return 1;
// 	ptrglGetCombinerOutputParameterfvNV = goglGetProcAddress("glGetCombinerOutputParameterfvNV");
// 	if(ptrglGetCombinerOutputParameterfvNV == NULL) return 1;
// 	ptrglGetCombinerOutputParameterivNV = goglGetProcAddress("glGetCombinerOutputParameterivNV");
// 	if(ptrglGetCombinerOutputParameterivNV == NULL) return 1;
// 	ptrglGetFinalCombinerInputParameterfvNV = goglGetProcAddress("glGetFinalCombinerInputParameterfvNV");
// 	if(ptrglGetFinalCombinerInputParameterfvNV == NULL) return 1;
// 	ptrglGetFinalCombinerInputParameterivNV = goglGetProcAddress("glGetFinalCombinerInputParameterivNV");
// 	if(ptrglGetFinalCombinerInputParameterivNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_register_combiners2() {
// 	ptrglCombinerStageParameterfvNV = goglGetProcAddress("glCombinerStageParameterfvNV");
// 	if(ptrglCombinerStageParameterfvNV == NULL) return 1;
// 	ptrglGetCombinerStageParameterfvNV = goglGetProcAddress("glGetCombinerStageParameterfvNV");
// 	if(ptrglGetCombinerStageParameterfvNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_shader_buffer_load() {
// 	ptrglMakeBufferResidentNV = goglGetProcAddress("glMakeBufferResidentNV");
// 	if(ptrglMakeBufferResidentNV == NULL) return 1;
// 	ptrglMakeBufferNonResidentNV = goglGetProcAddress("glMakeBufferNonResidentNV");
// 	if(ptrglMakeBufferNonResidentNV == NULL) return 1;
// 	ptrglIsBufferResidentNV = goglGetProcAddress("glIsBufferResidentNV");
// 	if(ptrglIsBufferResidentNV == NULL) return 1;
// 	ptrglMakeNamedBufferResidentNV = goglGetProcAddress("glMakeNamedBufferResidentNV");
// 	if(ptrglMakeNamedBufferResidentNV == NULL) return 1;
// 	ptrglMakeNamedBufferNonResidentNV = goglGetProcAddress("glMakeNamedBufferNonResidentNV");
// 	if(ptrglMakeNamedBufferNonResidentNV == NULL) return 1;
// 	ptrglIsNamedBufferResidentNV = goglGetProcAddress("glIsNamedBufferResidentNV");
// 	if(ptrglIsNamedBufferResidentNV == NULL) return 1;
// 	ptrglGetBufferParameterui64vNV = goglGetProcAddress("glGetBufferParameterui64vNV");
// 	if(ptrglGetBufferParameterui64vNV == NULL) return 1;
// 	ptrglGetNamedBufferParameterui64vNV = goglGetProcAddress("glGetNamedBufferParameterui64vNV");
// 	if(ptrglGetNamedBufferParameterui64vNV == NULL) return 1;
// 	ptrglGetIntegerui64vNV = goglGetProcAddress("glGetIntegerui64vNV");
// 	if(ptrglGetIntegerui64vNV == NULL) return 1;
// 	ptrglUniformui64NV = goglGetProcAddress("glUniformui64NV");
// 	if(ptrglUniformui64NV == NULL) return 1;
// 	ptrglUniformui64vNV = goglGetProcAddress("glUniformui64vNV");
// 	if(ptrglUniformui64vNV == NULL) return 1;
// 	ptrglGetUniformui64vNV = goglGetProcAddress("glGetUniformui64vNV");
// 	if(ptrglGetUniformui64vNV == NULL) return 1;
// 	ptrglProgramUniformui64NV = goglGetProcAddress("glProgramUniformui64NV");
// 	if(ptrglProgramUniformui64NV == NULL) return 1;
// 	ptrglProgramUniformui64vNV = goglGetProcAddress("glProgramUniformui64vNV");
// 	if(ptrglProgramUniformui64vNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_texture_barrier() {
// 	ptrglTextureBarrierNV = goglGetProcAddress("glTextureBarrierNV");
// 	if(ptrglTextureBarrierNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_texture_multisample() {
// 	ptrglTexImage2DMultisampleCoverageNV = goglGetProcAddress("glTexImage2DMultisampleCoverageNV");
// 	if(ptrglTexImage2DMultisampleCoverageNV == NULL) return 1;
// 	ptrglTexImage3DMultisampleCoverageNV = goglGetProcAddress("glTexImage3DMultisampleCoverageNV");
// 	if(ptrglTexImage3DMultisampleCoverageNV == NULL) return 1;
// 	ptrglTextureImage2DMultisampleNV = goglGetProcAddress("glTextureImage2DMultisampleNV");
// 	if(ptrglTextureImage2DMultisampleNV == NULL) return 1;
// 	ptrglTextureImage3DMultisampleNV = goglGetProcAddress("glTextureImage3DMultisampleNV");
// 	if(ptrglTextureImage3DMultisampleNV == NULL) return 1;
// 	ptrglTextureImage2DMultisampleCoverageNV = goglGetProcAddress("glTextureImage2DMultisampleCoverageNV");
// 	if(ptrglTextureImage2DMultisampleCoverageNV == NULL) return 1;
// 	ptrglTextureImage3DMultisampleCoverageNV = goglGetProcAddress("glTextureImage3DMultisampleCoverageNV");
// 	if(ptrglTextureImage3DMultisampleCoverageNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_transform_feedback() {
// 	ptrglBeginTransformFeedbackNV = goglGetProcAddress("glBeginTransformFeedbackNV");
// 	if(ptrglBeginTransformFeedbackNV == NULL) return 1;
// 	ptrglEndTransformFeedbackNV = goglGetProcAddress("glEndTransformFeedbackNV");
// 	if(ptrglEndTransformFeedbackNV == NULL) return 1;
// 	ptrglTransformFeedbackAttribsNV = goglGetProcAddress("glTransformFeedbackAttribsNV");
// 	if(ptrglTransformFeedbackAttribsNV == NULL) return 1;
// 	ptrglBindBufferRangeNV = goglGetProcAddress("glBindBufferRangeNV");
// 	if(ptrglBindBufferRangeNV == NULL) return 1;
// 	ptrglBindBufferOffsetNV = goglGetProcAddress("glBindBufferOffsetNV");
// 	if(ptrglBindBufferOffsetNV == NULL) return 1;
// 	ptrglBindBufferBaseNV = goglGetProcAddress("glBindBufferBaseNV");
// 	if(ptrglBindBufferBaseNV == NULL) return 1;
// 	ptrglTransformFeedbackVaryingsNV = goglGetProcAddress("glTransformFeedbackVaryingsNV");
// 	if(ptrglTransformFeedbackVaryingsNV == NULL) return 1;
// 	ptrglActiveVaryingNV = goglGetProcAddress("glActiveVaryingNV");
// 	if(ptrglActiveVaryingNV == NULL) return 1;
// 	ptrglGetVaryingLocationNV = goglGetProcAddress("glGetVaryingLocationNV");
// 	if(ptrglGetVaryingLocationNV == NULL) return 1;
// 	ptrglGetActiveVaryingNV = goglGetProcAddress("glGetActiveVaryingNV");
// 	if(ptrglGetActiveVaryingNV == NULL) return 1;
// 	ptrglGetTransformFeedbackVaryingNV = goglGetProcAddress("glGetTransformFeedbackVaryingNV");
// 	if(ptrglGetTransformFeedbackVaryingNV == NULL) return 1;
// 	ptrglTransformFeedbackStreamAttribsNV = goglGetProcAddress("glTransformFeedbackStreamAttribsNV");
// 	if(ptrglTransformFeedbackStreamAttribsNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_transform_feedback2() {
// 	ptrglBindTransformFeedbackNV = goglGetProcAddress("glBindTransformFeedbackNV");
// 	if(ptrglBindTransformFeedbackNV == NULL) return 1;
// 	ptrglDeleteTransformFeedbacksNV = goglGetProcAddress("glDeleteTransformFeedbacksNV");
// 	if(ptrglDeleteTransformFeedbacksNV == NULL) return 1;
// 	ptrglGenTransformFeedbacksNV = goglGetProcAddress("glGenTransformFeedbacksNV");
// 	if(ptrglGenTransformFeedbacksNV == NULL) return 1;
// 	ptrglIsTransformFeedbackNV = goglGetProcAddress("glIsTransformFeedbackNV");
// 	if(ptrglIsTransformFeedbackNV == NULL) return 1;
// 	ptrglPauseTransformFeedbackNV = goglGetProcAddress("glPauseTransformFeedbackNV");
// 	if(ptrglPauseTransformFeedbackNV == NULL) return 1;
// 	ptrglResumeTransformFeedbackNV = goglGetProcAddress("glResumeTransformFeedbackNV");
// 	if(ptrglResumeTransformFeedbackNV == NULL) return 1;
// 	ptrglDrawTransformFeedbackNV = goglGetProcAddress("glDrawTransformFeedbackNV");
// 	if(ptrglDrawTransformFeedbackNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_vdpau_interop() {
// 	ptrglVDPAUInitNV = goglGetProcAddress("glVDPAUInitNV");
// 	if(ptrglVDPAUInitNV == NULL) return 1;
// 	ptrglVDPAUFiniNV = goglGetProcAddress("glVDPAUFiniNV");
// 	if(ptrglVDPAUFiniNV == NULL) return 1;
// 	ptrglVDPAURegisterVideoSurfaceNV = goglGetProcAddress("glVDPAURegisterVideoSurfaceNV");
// 	if(ptrglVDPAURegisterVideoSurfaceNV == NULL) return 1;
// 	ptrglVDPAURegisterOutputSurfaceNV = goglGetProcAddress("glVDPAURegisterOutputSurfaceNV");
// 	if(ptrglVDPAURegisterOutputSurfaceNV == NULL) return 1;
// 	ptrglVDPAUIsSurfaceNV = goglGetProcAddress("glVDPAUIsSurfaceNV");
// 	if(ptrglVDPAUIsSurfaceNV == NULL) return 1;
// 	ptrglVDPAUUnregisterSurfaceNV = goglGetProcAddress("glVDPAUUnregisterSurfaceNV");
// 	if(ptrglVDPAUUnregisterSurfaceNV == NULL) return 1;
// 	ptrglVDPAUGetSurfaceivNV = goglGetProcAddress("glVDPAUGetSurfaceivNV");
// 	if(ptrglVDPAUGetSurfaceivNV == NULL) return 1;
// 	ptrglVDPAUSurfaceAccessNV = goglGetProcAddress("glVDPAUSurfaceAccessNV");
// 	if(ptrglVDPAUSurfaceAccessNV == NULL) return 1;
// 	ptrglVDPAUMapSurfacesNV = goglGetProcAddress("glVDPAUMapSurfacesNV");
// 	if(ptrglVDPAUMapSurfacesNV == NULL) return 1;
// 	ptrglVDPAUUnmapSurfacesNV = goglGetProcAddress("glVDPAUUnmapSurfacesNV");
// 	if(ptrglVDPAUUnmapSurfacesNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_vertex_array_range() {
// 	ptrglFlushVertexArrayRangeNV = goglGetProcAddress("glFlushVertexArrayRangeNV");
// 	if(ptrglFlushVertexArrayRangeNV == NULL) return 1;
// 	ptrglVertexArrayRangeNV = goglGetProcAddress("glVertexArrayRangeNV");
// 	if(ptrglVertexArrayRangeNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_vertex_attrib_integer_64bit() {
// 	ptrglVertexAttribL1i64NV = goglGetProcAddress("glVertexAttribL1i64NV");
// 	if(ptrglVertexAttribL1i64NV == NULL) return 1;
// 	ptrglVertexAttribL2i64NV = goglGetProcAddress("glVertexAttribL2i64NV");
// 	if(ptrglVertexAttribL2i64NV == NULL) return 1;
// 	ptrglVertexAttribL3i64NV = goglGetProcAddress("glVertexAttribL3i64NV");
// 	if(ptrglVertexAttribL3i64NV == NULL) return 1;
// 	ptrglVertexAttribL4i64NV = goglGetProcAddress("glVertexAttribL4i64NV");
// 	if(ptrglVertexAttribL4i64NV == NULL) return 1;
// 	ptrglVertexAttribL1i64vNV = goglGetProcAddress("glVertexAttribL1i64vNV");
// 	if(ptrglVertexAttribL1i64vNV == NULL) return 1;
// 	ptrglVertexAttribL2i64vNV = goglGetProcAddress("glVertexAttribL2i64vNV");
// 	if(ptrglVertexAttribL2i64vNV == NULL) return 1;
// 	ptrglVertexAttribL3i64vNV = goglGetProcAddress("glVertexAttribL3i64vNV");
// 	if(ptrglVertexAttribL3i64vNV == NULL) return 1;
// 	ptrglVertexAttribL4i64vNV = goglGetProcAddress("glVertexAttribL4i64vNV");
// 	if(ptrglVertexAttribL4i64vNV == NULL) return 1;
// 	ptrglVertexAttribL1ui64NV = goglGetProcAddress("glVertexAttribL1ui64NV");
// 	if(ptrglVertexAttribL1ui64NV == NULL) return 1;
// 	ptrglVertexAttribL2ui64NV = goglGetProcAddress("glVertexAttribL2ui64NV");
// 	if(ptrglVertexAttribL2ui64NV == NULL) return 1;
// 	ptrglVertexAttribL3ui64NV = goglGetProcAddress("glVertexAttribL3ui64NV");
// 	if(ptrglVertexAttribL3ui64NV == NULL) return 1;
// 	ptrglVertexAttribL4ui64NV = goglGetProcAddress("glVertexAttribL4ui64NV");
// 	if(ptrglVertexAttribL4ui64NV == NULL) return 1;
// 	ptrglVertexAttribL1ui64vNV = goglGetProcAddress("glVertexAttribL1ui64vNV");
// 	if(ptrglVertexAttribL1ui64vNV == NULL) return 1;
// 	ptrglVertexAttribL2ui64vNV = goglGetProcAddress("glVertexAttribL2ui64vNV");
// 	if(ptrglVertexAttribL2ui64vNV == NULL) return 1;
// 	ptrglVertexAttribL3ui64vNV = goglGetProcAddress("glVertexAttribL3ui64vNV");
// 	if(ptrglVertexAttribL3ui64vNV == NULL) return 1;
// 	ptrglVertexAttribL4ui64vNV = goglGetProcAddress("glVertexAttribL4ui64vNV");
// 	if(ptrglVertexAttribL4ui64vNV == NULL) return 1;
// 	ptrglGetVertexAttribLi64vNV = goglGetProcAddress("glGetVertexAttribLi64vNV");
// 	if(ptrglGetVertexAttribLi64vNV == NULL) return 1;
// 	ptrglGetVertexAttribLui64vNV = goglGetProcAddress("glGetVertexAttribLui64vNV");
// 	if(ptrglGetVertexAttribLui64vNV == NULL) return 1;
// 	ptrglVertexAttribLFormatNV = goglGetProcAddress("glVertexAttribLFormatNV");
// 	if(ptrglVertexAttribLFormatNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_vertex_buffer_unified_memory() {
// 	ptrglBufferAddressRangeNV = goglGetProcAddress("glBufferAddressRangeNV");
// 	if(ptrglBufferAddressRangeNV == NULL) return 1;
// 	ptrglVertexFormatNV = goglGetProcAddress("glVertexFormatNV");
// 	if(ptrglVertexFormatNV == NULL) return 1;
// 	ptrglNormalFormatNV = goglGetProcAddress("glNormalFormatNV");
// 	if(ptrglNormalFormatNV == NULL) return 1;
// 	ptrglColorFormatNV = goglGetProcAddress("glColorFormatNV");
// 	if(ptrglColorFormatNV == NULL) return 1;
// 	ptrglIndexFormatNV = goglGetProcAddress("glIndexFormatNV");
// 	if(ptrglIndexFormatNV == NULL) return 1;
// 	ptrglTexCoordFormatNV = goglGetProcAddress("glTexCoordFormatNV");
// 	if(ptrglTexCoordFormatNV == NULL) return 1;
// 	ptrglEdgeFlagFormatNV = goglGetProcAddress("glEdgeFlagFormatNV");
// 	if(ptrglEdgeFlagFormatNV == NULL) return 1;
// 	ptrglSecondaryColorFormatNV = goglGetProcAddress("glSecondaryColorFormatNV");
// 	if(ptrglSecondaryColorFormatNV == NULL) return 1;
// 	ptrglFogCoordFormatNV = goglGetProcAddress("glFogCoordFormatNV");
// 	if(ptrglFogCoordFormatNV == NULL) return 1;
// 	ptrglVertexAttribFormatNV = goglGetProcAddress("glVertexAttribFormatNV");
// 	if(ptrglVertexAttribFormatNV == NULL) return 1;
// 	ptrglVertexAttribIFormatNV = goglGetProcAddress("glVertexAttribIFormatNV");
// 	if(ptrglVertexAttribIFormatNV == NULL) return 1;
// 	ptrglGetIntegerui64i_vNV = goglGetProcAddress("glGetIntegerui64i_vNV");
// 	if(ptrglGetIntegerui64i_vNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_vertex_program() {
// 	ptrglAreProgramsResidentNV = goglGetProcAddress("glAreProgramsResidentNV");
// 	if(ptrglAreProgramsResidentNV == NULL) return 1;
// 	ptrglBindProgramNV = goglGetProcAddress("glBindProgramNV");
// 	if(ptrglBindProgramNV == NULL) return 1;
// 	ptrglDeleteProgramsNV = goglGetProcAddress("glDeleteProgramsNV");
// 	if(ptrglDeleteProgramsNV == NULL) return 1;
// 	ptrglExecuteProgramNV = goglGetProcAddress("glExecuteProgramNV");
// 	if(ptrglExecuteProgramNV == NULL) return 1;
// 	ptrglGenProgramsNV = goglGetProcAddress("glGenProgramsNV");
// 	if(ptrglGenProgramsNV == NULL) return 1;
// 	ptrglGetProgramParameterdvNV = goglGetProcAddress("glGetProgramParameterdvNV");
// 	if(ptrglGetProgramParameterdvNV == NULL) return 1;
// 	ptrglGetProgramParameterfvNV = goglGetProcAddress("glGetProgramParameterfvNV");
// 	if(ptrglGetProgramParameterfvNV == NULL) return 1;
// 	ptrglGetProgramivNV = goglGetProcAddress("glGetProgramivNV");
// 	if(ptrglGetProgramivNV == NULL) return 1;
// 	ptrglGetProgramStringNV = goglGetProcAddress("glGetProgramStringNV");
// 	if(ptrglGetProgramStringNV == NULL) return 1;
// 	ptrglGetTrackMatrixivNV = goglGetProcAddress("glGetTrackMatrixivNV");
// 	if(ptrglGetTrackMatrixivNV == NULL) return 1;
// 	ptrglGetVertexAttribdvNV = goglGetProcAddress("glGetVertexAttribdvNV");
// 	if(ptrglGetVertexAttribdvNV == NULL) return 1;
// 	ptrglGetVertexAttribfvNV = goglGetProcAddress("glGetVertexAttribfvNV");
// 	if(ptrglGetVertexAttribfvNV == NULL) return 1;
// 	ptrglGetVertexAttribivNV = goglGetProcAddress("glGetVertexAttribivNV");
// 	if(ptrglGetVertexAttribivNV == NULL) return 1;
// 	ptrglGetVertexAttribPointervNV = goglGetProcAddress("glGetVertexAttribPointervNV");
// 	if(ptrglGetVertexAttribPointervNV == NULL) return 1;
// 	ptrglIsProgramNV = goglGetProcAddress("glIsProgramNV");
// 	if(ptrglIsProgramNV == NULL) return 1;
// 	ptrglLoadProgramNV = goglGetProcAddress("glLoadProgramNV");
// 	if(ptrglLoadProgramNV == NULL) return 1;
// 	ptrglProgramParameter4dNV = goglGetProcAddress("glProgramParameter4dNV");
// 	if(ptrglProgramParameter4dNV == NULL) return 1;
// 	ptrglProgramParameter4dvNV = goglGetProcAddress("glProgramParameter4dvNV");
// 	if(ptrglProgramParameter4dvNV == NULL) return 1;
// 	ptrglProgramParameter4fNV = goglGetProcAddress("glProgramParameter4fNV");
// 	if(ptrglProgramParameter4fNV == NULL) return 1;
// 	ptrglProgramParameter4fvNV = goglGetProcAddress("glProgramParameter4fvNV");
// 	if(ptrglProgramParameter4fvNV == NULL) return 1;
// 	ptrglProgramParameters4dvNV = goglGetProcAddress("glProgramParameters4dvNV");
// 	if(ptrglProgramParameters4dvNV == NULL) return 1;
// 	ptrglProgramParameters4fvNV = goglGetProcAddress("glProgramParameters4fvNV");
// 	if(ptrglProgramParameters4fvNV == NULL) return 1;
// 	ptrglRequestResidentProgramsNV = goglGetProcAddress("glRequestResidentProgramsNV");
// 	if(ptrglRequestResidentProgramsNV == NULL) return 1;
// 	ptrglTrackMatrixNV = goglGetProcAddress("glTrackMatrixNV");
// 	if(ptrglTrackMatrixNV == NULL) return 1;
// 	ptrglVertexAttribPointerNV = goglGetProcAddress("glVertexAttribPointerNV");
// 	if(ptrglVertexAttribPointerNV == NULL) return 1;
// 	ptrglVertexAttrib1dNV = goglGetProcAddress("glVertexAttrib1dNV");
// 	if(ptrglVertexAttrib1dNV == NULL) return 1;
// 	ptrglVertexAttrib1dvNV = goglGetProcAddress("glVertexAttrib1dvNV");
// 	if(ptrglVertexAttrib1dvNV == NULL) return 1;
// 	ptrglVertexAttrib1fNV = goglGetProcAddress("glVertexAttrib1fNV");
// 	if(ptrglVertexAttrib1fNV == NULL) return 1;
// 	ptrglVertexAttrib1fvNV = goglGetProcAddress("glVertexAttrib1fvNV");
// 	if(ptrglVertexAttrib1fvNV == NULL) return 1;
// 	ptrglVertexAttrib1sNV = goglGetProcAddress("glVertexAttrib1sNV");
// 	if(ptrglVertexAttrib1sNV == NULL) return 1;
// 	ptrglVertexAttrib1svNV = goglGetProcAddress("glVertexAttrib1svNV");
// 	if(ptrglVertexAttrib1svNV == NULL) return 1;
// 	ptrglVertexAttrib2dNV = goglGetProcAddress("glVertexAttrib2dNV");
// 	if(ptrglVertexAttrib2dNV == NULL) return 1;
// 	ptrglVertexAttrib2dvNV = goglGetProcAddress("glVertexAttrib2dvNV");
// 	if(ptrglVertexAttrib2dvNV == NULL) return 1;
// 	ptrglVertexAttrib2fNV = goglGetProcAddress("glVertexAttrib2fNV");
// 	if(ptrglVertexAttrib2fNV == NULL) return 1;
// 	ptrglVertexAttrib2fvNV = goglGetProcAddress("glVertexAttrib2fvNV");
// 	if(ptrglVertexAttrib2fvNV == NULL) return 1;
// 	ptrglVertexAttrib2sNV = goglGetProcAddress("glVertexAttrib2sNV");
// 	if(ptrglVertexAttrib2sNV == NULL) return 1;
// 	ptrglVertexAttrib2svNV = goglGetProcAddress("glVertexAttrib2svNV");
// 	if(ptrglVertexAttrib2svNV == NULL) return 1;
// 	ptrglVertexAttrib3dNV = goglGetProcAddress("glVertexAttrib3dNV");
// 	if(ptrglVertexAttrib3dNV == NULL) return 1;
// 	ptrglVertexAttrib3dvNV = goglGetProcAddress("glVertexAttrib3dvNV");
// 	if(ptrglVertexAttrib3dvNV == NULL) return 1;
// 	ptrglVertexAttrib3fNV = goglGetProcAddress("glVertexAttrib3fNV");
// 	if(ptrglVertexAttrib3fNV == NULL) return 1;
// 	ptrglVertexAttrib3fvNV = goglGetProcAddress("glVertexAttrib3fvNV");
// 	if(ptrglVertexAttrib3fvNV == NULL) return 1;
// 	ptrglVertexAttrib3sNV = goglGetProcAddress("glVertexAttrib3sNV");
// 	if(ptrglVertexAttrib3sNV == NULL) return 1;
// 	ptrglVertexAttrib3svNV = goglGetProcAddress("glVertexAttrib3svNV");
// 	if(ptrglVertexAttrib3svNV == NULL) return 1;
// 	ptrglVertexAttrib4dNV = goglGetProcAddress("glVertexAttrib4dNV");
// 	if(ptrglVertexAttrib4dNV == NULL) return 1;
// 	ptrglVertexAttrib4dvNV = goglGetProcAddress("glVertexAttrib4dvNV");
// 	if(ptrglVertexAttrib4dvNV == NULL) return 1;
// 	ptrglVertexAttrib4fNV = goglGetProcAddress("glVertexAttrib4fNV");
// 	if(ptrglVertexAttrib4fNV == NULL) return 1;
// 	ptrglVertexAttrib4fvNV = goglGetProcAddress("glVertexAttrib4fvNV");
// 	if(ptrglVertexAttrib4fvNV == NULL) return 1;
// 	ptrglVertexAttrib4sNV = goglGetProcAddress("glVertexAttrib4sNV");
// 	if(ptrglVertexAttrib4sNV == NULL) return 1;
// 	ptrglVertexAttrib4svNV = goglGetProcAddress("glVertexAttrib4svNV");
// 	if(ptrglVertexAttrib4svNV == NULL) return 1;
// 	ptrglVertexAttrib4ubNV = goglGetProcAddress("glVertexAttrib4ubNV");
// 	if(ptrglVertexAttrib4ubNV == NULL) return 1;
// 	ptrglVertexAttrib4ubvNV = goglGetProcAddress("glVertexAttrib4ubvNV");
// 	if(ptrglVertexAttrib4ubvNV == NULL) return 1;
// 	ptrglVertexAttribs1dvNV = goglGetProcAddress("glVertexAttribs1dvNV");
// 	if(ptrglVertexAttribs1dvNV == NULL) return 1;
// 	ptrglVertexAttribs1fvNV = goglGetProcAddress("glVertexAttribs1fvNV");
// 	if(ptrglVertexAttribs1fvNV == NULL) return 1;
// 	ptrglVertexAttribs1svNV = goglGetProcAddress("glVertexAttribs1svNV");
// 	if(ptrglVertexAttribs1svNV == NULL) return 1;
// 	ptrglVertexAttribs2dvNV = goglGetProcAddress("glVertexAttribs2dvNV");
// 	if(ptrglVertexAttribs2dvNV == NULL) return 1;
// 	ptrglVertexAttribs2fvNV = goglGetProcAddress("glVertexAttribs2fvNV");
// 	if(ptrglVertexAttribs2fvNV == NULL) return 1;
// 	ptrglVertexAttribs2svNV = goglGetProcAddress("glVertexAttribs2svNV");
// 	if(ptrglVertexAttribs2svNV == NULL) return 1;
// 	ptrglVertexAttribs3dvNV = goglGetProcAddress("glVertexAttribs3dvNV");
// 	if(ptrglVertexAttribs3dvNV == NULL) return 1;
// 	ptrglVertexAttribs3fvNV = goglGetProcAddress("glVertexAttribs3fvNV");
// 	if(ptrglVertexAttribs3fvNV == NULL) return 1;
// 	ptrglVertexAttribs3svNV = goglGetProcAddress("glVertexAttribs3svNV");
// 	if(ptrglVertexAttribs3svNV == NULL) return 1;
// 	ptrglVertexAttribs4dvNV = goglGetProcAddress("glVertexAttribs4dvNV");
// 	if(ptrglVertexAttribs4dvNV == NULL) return 1;
// 	ptrglVertexAttribs4fvNV = goglGetProcAddress("glVertexAttribs4fvNV");
// 	if(ptrglVertexAttribs4fvNV == NULL) return 1;
// 	ptrglVertexAttribs4svNV = goglGetProcAddress("glVertexAttribs4svNV");
// 	if(ptrglVertexAttribs4svNV == NULL) return 1;
// 	ptrglVertexAttribs4ubvNV = goglGetProcAddress("glVertexAttribs4ubvNV");
// 	if(ptrglVertexAttribs4ubvNV == NULL) return 1;
// 	return 0;
// }
// int init_NV_vertex_program4() {
// 	ptrglVertexAttribI1iEXT = goglGetProcAddress("glVertexAttribI1iEXT");
// 	if(ptrglVertexAttribI1iEXT == NULL) return 1;
// 	ptrglVertexAttribI2iEXT = goglGetProcAddress("glVertexAttribI2iEXT");
// 	if(ptrglVertexAttribI2iEXT == NULL) return 1;
// 	ptrglVertexAttribI3iEXT = goglGetProcAddress("glVertexAttribI3iEXT");
// 	if(ptrglVertexAttribI3iEXT == NULL) return 1;
// 	ptrglVertexAttribI4iEXT = goglGetProcAddress("glVertexAttribI4iEXT");
// 	if(ptrglVertexAttribI4iEXT == NULL) return 1;
// 	ptrglVertexAttribI1uiEXT = goglGetProcAddress("glVertexAttribI1uiEXT");
// 	if(ptrglVertexAttribI1uiEXT == NULL) return 1;
// 	ptrglVertexAttribI2uiEXT = goglGetProcAddress("glVertexAttribI2uiEXT");
// 	if(ptrglVertexAttribI2uiEXT == NULL) return 1;
// 	ptrglVertexAttribI3uiEXT = goglGetProcAddress("glVertexAttribI3uiEXT");
// 	if(ptrglVertexAttribI3uiEXT == NULL) return 1;
// 	ptrglVertexAttribI4uiEXT = goglGetProcAddress("glVertexAttribI4uiEXT");
// 	if(ptrglVertexAttribI4uiEXT == NULL) return 1;
// 	ptrglVertexAttribI1ivEXT = goglGetProcAddress("glVertexAttribI1ivEXT");
// 	if(ptrglVertexAttribI1ivEXT == NULL) return 1;
// 	ptrglVertexAttribI2ivEXT = goglGetProcAddress("glVertexAttribI2ivEXT");
// 	if(ptrglVertexAttribI2ivEXT == NULL) return 1;
// 	ptrglVertexAttribI3ivEXT = goglGetProcAddress("glVertexAttribI3ivEXT");
// 	if(ptrglVertexAttribI3ivEXT == NULL) return 1;
// 	ptrglVertexAttribI4ivEXT = goglGetProcAddress("glVertexAttribI4ivEXT");
// 	if(ptrglVertexAttribI4ivEXT == NULL) return 1;
// 	ptrglVertexAttribI1uivEXT = goglGetProcAddress("glVertexAttribI1uivEXT");
// 	if(ptrglVertexAttribI1uivEXT == NULL) return 1;
// 	ptrglVertexAttribI2uivEXT = goglGetProcAddress("glVertexAttribI2uivEXT");
// 	if(ptrglVertexAttribI2uivEXT == NULL) return 1;
// 	ptrglVertexAttribI3uivEXT = goglGetProcAddress("glVertexAttribI3uivEXT");
// 	if(ptrglVertexAttribI3uivEXT == NULL) return 1;
// 	ptrglVertexAttribI4uivEXT = goglGetProcAddress("glVertexAttribI4uivEXT");
// 	if(ptrglVertexAttribI4uivEXT == NULL) return 1;
// 	ptrglVertexAttribI4bvEXT = goglGetProcAddress("glVertexAttribI4bvEXT");
// 	if(ptrglVertexAttribI4bvEXT == NULL) return 1;
// 	ptrglVertexAttribI4svEXT = goglGetProcAddress("glVertexAttribI4svEXT");
// 	if(ptrglVertexAttribI4svEXT == NULL) return 1;
// 	ptrglVertexAttribI4ubvEXT = goglGetProcAddress("glVertexAttribI4ubvEXT");
// 	if(ptrglVertexAttribI4ubvEXT == NULL) return 1;
// 	ptrglVertexAttribI4usvEXT = goglGetProcAddress("glVertexAttribI4usvEXT");
// 	if(ptrglVertexAttribI4usvEXT == NULL) return 1;
// 	ptrglVertexAttribIPointerEXT = goglGetProcAddress("glVertexAttribIPointerEXT");
// 	if(ptrglVertexAttribIPointerEXT == NULL) return 1;
// 	ptrglGetVertexAttribIivEXT = goglGetProcAddress("glGetVertexAttribIivEXT");
// 	if(ptrglGetVertexAttribIivEXT == NULL) return 1;
// 	ptrglGetVertexAttribIuivEXT = goglGetProcAddress("glGetVertexAttribIuivEXT");
// 	if(ptrglGetVertexAttribIuivEXT == NULL) return 1;
// 	return 0;
// }
// int init_NV_video_capture() {
// 	ptrglBeginVideoCaptureNV = goglGetProcAddress("glBeginVideoCaptureNV");
// 	if(ptrglBeginVideoCaptureNV == NULL) return 1;
// 	ptrglBindVideoCaptureStreamBufferNV = goglGetProcAddress("glBindVideoCaptureStreamBufferNV");
// 	if(ptrglBindVideoCaptureStreamBufferNV == NULL) return 1;
// 	ptrglBindVideoCaptureStreamTextureNV = goglGetProcAddress("glBindVideoCaptureStreamTextureNV");
// 	if(ptrglBindVideoCaptureStreamTextureNV == NULL) return 1;
// 	ptrglEndVideoCaptureNV = goglGetProcAddress("glEndVideoCaptureNV");
// 	if(ptrglEndVideoCaptureNV == NULL) return 1;
// 	ptrglGetVideoCaptureivNV = goglGetProcAddress("glGetVideoCaptureivNV");
// 	if(ptrglGetVideoCaptureivNV == NULL) return 1;
// 	ptrglGetVideoCaptureStreamivNV = goglGetProcAddress("glGetVideoCaptureStreamivNV");
// 	if(ptrglGetVideoCaptureStreamivNV == NULL) return 1;
// 	ptrglGetVideoCaptureStreamfvNV = goglGetProcAddress("glGetVideoCaptureStreamfvNV");
// 	if(ptrglGetVideoCaptureStreamfvNV == NULL) return 1;
// 	ptrglGetVideoCaptureStreamdvNV = goglGetProcAddress("glGetVideoCaptureStreamdvNV");
// 	if(ptrglGetVideoCaptureStreamdvNV == NULL) return 1;
// 	ptrglVideoCaptureNV = goglGetProcAddress("glVideoCaptureNV");
// 	if(ptrglVideoCaptureNV == NULL) return 1;
// 	ptrglVideoCaptureStreamParameterivNV = goglGetProcAddress("glVideoCaptureStreamParameterivNV");
// 	if(ptrglVideoCaptureStreamParameterivNV == NULL) return 1;
// 	ptrglVideoCaptureStreamParameterfvNV = goglGetProcAddress("glVideoCaptureStreamParameterfvNV");
// 	if(ptrglVideoCaptureStreamParameterfvNV == NULL) return 1;
// 	ptrglVideoCaptureStreamParameterdvNV = goglGetProcAddress("glVideoCaptureStreamParameterdvNV");
// 	if(ptrglVideoCaptureStreamParameterdvNV == NULL) return 1;
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

// NV_bindless_texture
const (
)
// NV_blend_square
const (
)
// NV_conditional_render
const (
	QUERY_BY_REGION_NO_WAIT_NV = 0x8E16
	QUERY_BY_REGION_WAIT_NV = 0x8E15
	QUERY_NO_WAIT_NV = 0x8E14
	QUERY_WAIT_NV = 0x8E13
)
// NV_copy_depth_to_color
const (
	DEPTH_STENCIL_TO_BGRA_NV = 0x886F
	DEPTH_STENCIL_TO_RGBA_NV = 0x886E
)
// NV_copy_image
const (
)
// NV_depth_buffer_float
const (
	DEPTH32F_STENCIL8_NV = 0x8DAC
	DEPTH_BUFFER_FLOAT_MODE_NV = 0x8DAF
	DEPTH_COMPONENT32F_NV = 0x8DAB
	FLOAT_32_UNSIGNED_INT_24_8_REV_NV = 0x8DAD
)
// NV_depth_clamp
const (
	DEPTH_CLAMP_NV = 0x864F
)
// NV_evaluators
const (
	EVAL_2D_NV = 0x86C0
	EVAL_FRACTIONAL_TESSELLATION_NV = 0x86C5
	EVAL_TRIANGULAR_2D_NV = 0x86C1
	EVAL_VERTEX_ATTRIB0_NV = 0x86C6
	EVAL_VERTEX_ATTRIB10_NV = 0x86D0
	EVAL_VERTEX_ATTRIB11_NV = 0x86D1
	EVAL_VERTEX_ATTRIB12_NV = 0x86D2
	EVAL_VERTEX_ATTRIB13_NV = 0x86D3
	EVAL_VERTEX_ATTRIB14_NV = 0x86D4
	EVAL_VERTEX_ATTRIB15_NV = 0x86D5
	EVAL_VERTEX_ATTRIB1_NV = 0x86C7
	EVAL_VERTEX_ATTRIB2_NV = 0x86C8
	EVAL_VERTEX_ATTRIB3_NV = 0x86C9
	EVAL_VERTEX_ATTRIB4_NV = 0x86CA
	EVAL_VERTEX_ATTRIB5_NV = 0x86CB
	EVAL_VERTEX_ATTRIB6_NV = 0x86CC
	EVAL_VERTEX_ATTRIB7_NV = 0x86CD
	EVAL_VERTEX_ATTRIB8_NV = 0x86CE
	EVAL_VERTEX_ATTRIB9_NV = 0x86CF
	MAP_ATTRIB_U_ORDER_NV = 0x86C3
	MAP_ATTRIB_V_ORDER_NV = 0x86C4
	MAP_TESSELLATION_NV = 0x86C2
	MAX_MAP_TESSELLATION_NV = 0x86D6
	MAX_RATIONAL_EVAL_ORDER_NV = 0x86D7
)
// NV_explicit_multisample
const (
	INT_SAMPLER_RENDERBUFFER_NV = 0x8E57
	MAX_SAMPLE_MASK_WORDS_NV = 0x8E59
	SAMPLER_RENDERBUFFER_NV = 0x8E56
	SAMPLE_MASK_NV = 0x8E51
	SAMPLE_MASK_VALUE_NV = 0x8E52
	SAMPLE_POSITION_NV = 0x8E50
	TEXTURE_BINDING_RENDERBUFFER_NV = 0x8E53
	TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV = 0x8E54
	TEXTURE_RENDERBUFFER_NV = 0x8E55
	UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV = 0x8E58
)
// NV_fence
const (
	ALL_COMPLETED_NV = 0x84F2
	FENCE_CONDITION_NV = 0x84F4
	FENCE_STATUS_NV = 0x84F3
)
// NV_float_buffer
const (
	FLOAT_CLEAR_COLOR_VALUE_NV = 0x888D
	FLOAT_R16_NV = 0x8884
	FLOAT_R32_NV = 0x8885
	FLOAT_RG16_NV = 0x8886
	FLOAT_RG32_NV = 0x8887
	FLOAT_RGB16_NV = 0x8888
	FLOAT_RGB32_NV = 0x8889
	FLOAT_RGBA16_NV = 0x888A
	FLOAT_RGBA32_NV = 0x888B
	FLOAT_RGBA_MODE_NV = 0x888E
	FLOAT_RGBA_NV = 0x8883
	FLOAT_RGB_NV = 0x8882
	FLOAT_RG_NV = 0x8881
	FLOAT_R_NV = 0x8880
	TEXTURE_FLOAT_COMPONENTS_NV = 0x888C
)
// NV_fog_distance
const (
	EYE_PLANE_ABSOLUTE_NV = 0x855C
	EYE_RADIAL_NV = 0x855B
	FOG_DISTANCE_MODE_NV = 0x855A
)
// NV_fragment_program
const (
	FRAGMENT_PROGRAM_BINDING_NV = 0x8873
	FRAGMENT_PROGRAM_NV = 0x8870
	MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV = 0x8868
	MAX_TEXTURE_COORDS_NV = 0x8871
	MAX_TEXTURE_IMAGE_UNITS_NV = 0x8872
	PROGRAM_ERROR_STRING_NV = 0x8874
)
// NV_fragment_program2
const (
	MAX_PROGRAM_IF_DEPTH_NV = 0x88F6
	MAX_PROGRAM_LOOP_COUNT_NV = 0x88F8
	MAX_PROGRAM_LOOP_DEPTH_NV = 0x88F7
)
// NV_fragment_program4
const (
)
// NV_fragment_program_option
const (
)
// NV_framebuffer_multisample_coverage
const (
	MAX_MULTISAMPLE_COVERAGE_MODES_NV = 0x8E11
	MULTISAMPLE_COVERAGE_MODES_NV = 0x8E12
	RENDERBUFFER_COLOR_SAMPLES_NV = 0x8E10
	RENDERBUFFER_COVERAGE_SAMPLES_NV = 0x8CAB
)
// NV_geometry_program4
const (
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT = 0x8DA7
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT = 0x8CD4
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT = 0x8DA9
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT = 0x8DA8
	GEOMETRY_INPUT_TYPE_EXT = 0x8DDB
	GEOMETRY_OUTPUT_TYPE_EXT = 0x8DDC
	GEOMETRY_PROGRAM_NV = 0x8C26
	GEOMETRY_VERTICES_OUT_EXT = 0x8DDA
	LINES_ADJACENCY_EXT = 0x000A
	LINE_STRIP_ADJACENCY_EXT = 0x000B
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT = 0x8C29
	MAX_PROGRAM_OUTPUT_VERTICES_NV = 0x8C27
	MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV = 0x8C28
	PROGRAM_POINT_SIZE_EXT = 0x8642
	TRIANGLES_ADJACENCY_EXT = 0x000C
	TRIANGLE_STRIP_ADJACENCY_EXT = 0x000D
)
// NV_geometry_shader4
const (
)
// NV_gpu_program4
const (
	MAX_PROGRAM_ATTRIB_COMPONENTS_NV = 0x8908
	MAX_PROGRAM_GENERIC_ATTRIBS_NV = 0x8DA5
	MAX_PROGRAM_GENERIC_RESULTS_NV = 0x8DA6
	MAX_PROGRAM_RESULT_COMPONENTS_NV = 0x8909
	MAX_PROGRAM_TEXEL_OFFSET_NV = 0x8905
	MIN_PROGRAM_TEXEL_OFFSET_NV = 0x8904
	PROGRAM_ATTRIB_COMPONENTS_NV = 0x8906
	PROGRAM_RESULT_COMPONENTS_NV = 0x8907
)
// NV_gpu_program5
const (
	FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV = 0x8E5D
	MAX_FRAGMENT_INTERPOLATION_OFFSET_NV = 0x8E5C
	MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV = 0x8E5A
	MAX_PROGRAM_SUBROUTINE_NUM_NV = 0x8F45
	MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV = 0x8F44
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV = 0x8E5F
	MIN_FRAGMENT_INTERPOLATION_OFFSET_NV = 0x8E5B
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV = 0x8E5E
)
// NV_gpu_shader5
const (
	FLOAT16_NV = 0x8FF8
	FLOAT16_VEC2_NV = 0x8FF9
	FLOAT16_VEC3_NV = 0x8FFA
	FLOAT16_VEC4_NV = 0x8FFB
	INT16_NV = 0x8FE4
	INT16_VEC2_NV = 0x8FE5
	INT16_VEC3_NV = 0x8FE6
	INT16_VEC4_NV = 0x8FE7
	INT64_VEC2_NV = 0x8FE9
	INT64_VEC3_NV = 0x8FEA
	INT64_VEC4_NV = 0x8FEB
	INT8_NV = 0x8FE0
	INT8_VEC2_NV = 0x8FE1
	INT8_VEC3_NV = 0x8FE2
	INT8_VEC4_NV = 0x8FE3
	PATCHES = 0x000E
	UNSIGNED_INT16_NV = 0x8FF0
	UNSIGNED_INT16_VEC2_NV = 0x8FF1
	UNSIGNED_INT16_VEC3_NV = 0x8FF2
	UNSIGNED_INT16_VEC4_NV = 0x8FF3
	UNSIGNED_INT64_VEC2_NV = 0x8FF5
	UNSIGNED_INT64_VEC3_NV = 0x8FF6
	UNSIGNED_INT64_VEC4_NV = 0x8FF7
	UNSIGNED_INT8_NV = 0x8FEC
	UNSIGNED_INT8_VEC2_NV = 0x8FED
	UNSIGNED_INT8_VEC3_NV = 0x8FEE
	UNSIGNED_INT8_VEC4_NV = 0x8FEF
)
// NV_half_float
const (
	HALF_FLOAT_NV = 0x140B
)
// NV_light_max_exponent
const (
	MAX_SHININESS_NV = 0x8504
	MAX_SPOT_EXPONENT_NV = 0x8505
)
// NV_multisample_coverage
const (
	COLOR_SAMPLES_NV = 0x8E20
	COVERAGE_SAMPLES_NV = 0x80A9
)
// NV_multisample_filter_hint
const (
	MULTISAMPLE_FILTER_HINT_NV = 0x8534
)
// NV_occlusion_query
const (
	CURRENT_OCCLUSION_QUERY_ID_NV = 0x8865
	PIXEL_COUNTER_BITS_NV = 0x8864
	PIXEL_COUNT_AVAILABLE_NV = 0x8867
	PIXEL_COUNT_NV = 0x8866
)
// NV_packed_depth_stencil
const (
	DEPTH_STENCIL_NV = 0x84F9
	UNSIGNED_INT_24_8_NV = 0x84FA
)
// NV_parameter_buffer_object
const (
	FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV = 0x8DA4
	GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV = 0x8DA3
	MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV = 0x8DA0
	MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV = 0x8DA1
	VERTEX_PROGRAM_PARAMETER_BUFFER_NV = 0x8DA2
)
// NV_parameter_buffer_object2
const (
)
// NV_path_rendering
const (
	ACCUM_ADJACENT_PAIRS_NV = 0x90AD
	ADJACENT_PAIRS_NV = 0x90AE
	AFFINE_2D_NV = 0x9092
	AFFINE_3D_NV = 0x9094
	ARC_TO_NV = 0xFE
	BEVEL_NV = 0x90A6
	BOLD_BIT_NV = 0x01
	BOUNDING_BOX_NV = 0x908D
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV = 0x909C
	CIRCULAR_CCW_ARC_TO_NV = 0xF8
	CIRCULAR_CW_ARC_TO_NV = 0xFA
	CIRCULAR_TANGENT_ARC_TO_NV = 0xFC
	CLOSE_PATH_NV = 0x00
	CONVEX_HULL_NV = 0x908B
	COUNT_DOWN_NV = 0x9089
	COUNT_UP_NV = 0x9088
	CUBIC_CURVE_TO_NV = 0x0C
	DUP_FIRST_CUBIC_CURVE_TO_NV = 0xF2
	DUP_LAST_CUBIC_CURVE_TO_NV = 0xF4
	FILE_NAME_NV = 0x9074
	FIRST_TO_REST_NV = 0x90AF
	FONT_ASCENDER_NV = 0x00200000
	FONT_DESCENDER_NV = 0x00400000
	FONT_HAS_KERNING_NV = 0x10000000
	FONT_HEIGHT_NV = 0x00800000
	FONT_MAX_ADVANCE_HEIGHT_NV = 0x02000000
	FONT_MAX_ADVANCE_WIDTH_NV = 0x01000000
	FONT_UNDERLINE_POSITION_NV = 0x04000000
	FONT_UNDERLINE_THICKNESS_NV = 0x08000000
	FONT_UNITS_PER_EM_NV = 0x00100000
	FONT_X_MAX_BOUNDS_NV = 0x00040000
	FONT_X_MIN_BOUNDS_NV = 0x00010000
	FONT_Y_MAX_BOUNDS_NV = 0x00080000
	FONT_Y_MIN_BOUNDS_NV = 0x00020000
	GLYPH_HAS_KERNING_NV = 0x100
	GLYPH_HEIGHT_BIT_NV = 0x02
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV = 0x10
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV = 0x04
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV = 0x08
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV = 0x80
	GLYPH_VERTICAL_BEARING_X_BIT_NV = 0x20
	GLYPH_VERTICAL_BEARING_Y_BIT_NV = 0x40
	GLYPH_WIDTH_BIT_NV = 0x01
	HORIZONTAL_LINE_TO_NV = 0x06
	ITALIC_BIT_NV = 0x02
	LARGE_CCW_ARC_TO_NV = 0x16
	LARGE_CW_ARC_TO_NV = 0x18
	LINE_TO_NV = 0x04
	MITER_REVERT_NV = 0x90A7
	MITER_TRUNCATE_NV = 0x90A8
	MOVE_TO_CONTINUES_NV = 0x90B6
	MOVE_TO_NV = 0x02
	MOVE_TO_RESETS_NV = 0x90B5
	MULTI_HULLS_NV = 0x908C
	PATH_CLIENT_LENGTH_NV = 0x907F
	PATH_COMMAND_COUNT_NV = 0x909D
	PATH_COMPUTED_LENGTH_NV = 0x90A0
	PATH_COORD_COUNT_NV = 0x909E
	PATH_COVER_DEPTH_FUNC_NV = 0x90BF
	PATH_DASH_ARRAY_COUNT_NV = 0x909F
	PATH_DASH_CAPS_NV = 0x907B
	PATH_DASH_OFFSET_NV = 0x907E
	PATH_DASH_OFFSET_RESET_NV = 0x90B4
	PATH_END_CAPS_NV = 0x9076
	PATH_ERROR_POSITION_NV = 0x90AB
	PATH_FILL_BOUNDING_BOX_NV = 0x90A1
	PATH_FILL_COVER_MODE_NV = 0x9082
	PATH_FILL_MASK_NV = 0x9081
	PATH_FILL_MODE_NV = 0x9080
	PATH_FOG_GEN_MODE_NV = 0x90AC
	PATH_FORMAT_PS_NV = 0x9071
	PATH_FORMAT_SVG_NV = 0x9070
	PATH_GEN_COEFF_NV = 0x90B1
	PATH_GEN_COLOR_FORMAT_NV = 0x90B2
	PATH_GEN_COMPONENTS_NV = 0x90B3
	PATH_GEN_MODE_NV = 0x90B0
	PATH_INITIAL_DASH_CAP_NV = 0x907C
	PATH_INITIAL_END_CAP_NV = 0x9077
	PATH_JOIN_STYLE_NV = 0x9079
	PATH_MITER_LIMIT_NV = 0x907A
	PATH_OBJECT_BOUNDING_BOX_NV = 0x908A
	PATH_SAMPLE_QUALITY_NV = 0x9085
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV = 0x90BD
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV = 0x90BE
	PATH_STENCIL_FUNC_NV = 0x90B7
	PATH_STENCIL_REF_NV = 0x90B8
	PATH_STENCIL_VALUE_MASK_NV = 0x90B9
	PATH_STROKE_BOUNDING_BOX_NV = 0x90A2
	PATH_STROKE_BOUND_NV = 0x9086
	PATH_STROKE_COVER_MODE_NV = 0x9083
	PATH_STROKE_MASK_NV = 0x9084
	PATH_STROKE_OVERSAMPLE_COUNT_NV = 0x9087
	PATH_STROKE_WIDTH_NV = 0x9075
	PATH_TERMINAL_DASH_CAP_NV = 0x907D
	PATH_TERMINAL_END_CAP_NV = 0x9078
	PROJECTIVE_2D_NV = 0x9093
	PROJECTIVE_3D_NV = 0x9095
	QUADRATIC_CURVE_TO_NV = 0x0A
	RECT_NV = 0xF6
	RELATIVE_ARC_TO_NV = 0xFF
	RELATIVE_CUBIC_CURVE_TO_NV = 0x0D
	RELATIVE_HORIZONTAL_LINE_TO_NV = 0x07
	RELATIVE_LARGE_CCW_ARC_TO_NV = 0x17
	RELATIVE_LARGE_CW_ARC_TO_NV = 0x19
	RELATIVE_LINE_TO_NV = 0x05
	RELATIVE_MOVE_TO_NV = 0x03
	RELATIVE_QUADRATIC_CURVE_TO_NV = 0x0B
	RELATIVE_SMALL_CCW_ARC_TO_NV = 0x13
	RELATIVE_SMALL_CW_ARC_TO_NV = 0x15
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV = 0x11
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV = 0x0F
	RELATIVE_VERTICAL_LINE_TO_NV = 0x09
	RESTART_PATH_NV = 0xF0
	ROUND_NV = 0x90A4
	SKIP_MISSING_GLYPH_NV = 0x90A9
	SMALL_CCW_ARC_TO_NV = 0x12
	SMALL_CW_ARC_TO_NV = 0x14
	SMOOTH_CUBIC_CURVE_TO_NV = 0x10
	SMOOTH_QUADRATIC_CURVE_TO_NV = 0x0E
	SQUARE_NV = 0x90A3
	STANDARD_FONT_NAME_NV = 0x9072
	SYSTEM_FONT_NAME_NV = 0x9073
	TRANSLATE_2D_NV = 0x9090
	TRANSLATE_3D_NV = 0x9091
	TRANSLATE_X_NV = 0x908E
	TRANSLATE_Y_NV = 0x908F
	TRANSPOSE_AFFINE_2D_NV = 0x9096
	TRANSPOSE_AFFINE_3D_NV = 0x9098
	TRANSPOSE_PROJECTIVE_2D_NV = 0x9097
	TRANSPOSE_PROJECTIVE_3D_NV = 0x9099
	TRIANGULAR_NV = 0x90A5
	USE_MISSING_GLYPH_NV = 0x90AA
	UTF16_NV = 0x909B
	UTF8_NV = 0x909A
	VERTICAL_LINE_TO_NV = 0x08
)
// NV_pixel_data_range
const (
	READ_PIXEL_DATA_RANGE_LENGTH_NV = 0x887B
	READ_PIXEL_DATA_RANGE_NV = 0x8879
	READ_PIXEL_DATA_RANGE_POINTER_NV = 0x887D
	WRITE_PIXEL_DATA_RANGE_LENGTH_NV = 0x887A
	WRITE_PIXEL_DATA_RANGE_NV = 0x8878
	WRITE_PIXEL_DATA_RANGE_POINTER_NV = 0x887C
)
// NV_point_sprite
const (
	COORD_REPLACE_NV = 0x8862
	POINT_SPRITE_NV = 0x8861
	POINT_SPRITE_R_MODE_NV = 0x8863
)
// NV_present_video
const (
	CURRENT_TIME_NV = 0x8E28
	FIELDS_NV = 0x8E27
	FRAME_NV = 0x8E26
	NUM_FILL_STREAMS_NV = 0x8E29
	PRESENT_DURATION_NV = 0x8E2B
	PRESENT_TIME_NV = 0x8E2A
)
// NV_primitive_restart
const (
	PRIMITIVE_RESTART_INDEX_NV = 0x8559
	PRIMITIVE_RESTART_NV = 0x8558
)
// NV_register_combiners
const (
	BIAS_BY_NEGATIVE_ONE_HALF_NV = 0x8541
	COLOR_SUM_CLAMP_NV = 0x854F
	COMBINER0_NV = 0x8550
	COMBINER1_NV = 0x8551
	COMBINER2_NV = 0x8552
	COMBINER3_NV = 0x8553
	COMBINER4_NV = 0x8554
	COMBINER5_NV = 0x8555
	COMBINER6_NV = 0x8556
	COMBINER7_NV = 0x8557
	COMBINER_AB_DOT_PRODUCT_NV = 0x8545
	COMBINER_AB_OUTPUT_NV = 0x854A
	COMBINER_BIAS_NV = 0x8549
	COMBINER_CD_DOT_PRODUCT_NV = 0x8546
	COMBINER_CD_OUTPUT_NV = 0x854B
	COMBINER_COMPONENT_USAGE_NV = 0x8544
	COMBINER_INPUT_NV = 0x8542
	COMBINER_MAPPING_NV = 0x8543
	COMBINER_MUX_SUM_NV = 0x8547
	COMBINER_SCALE_NV = 0x8548
	COMBINER_SUM_OUTPUT_NV = 0x854C
	CONSTANT_COLOR0_NV = 0x852A
	CONSTANT_COLOR1_NV = 0x852B
	DISCARD_NV = 0x8530
	EXPAND_NEGATE_NV = 0x8539
	EXPAND_NORMAL_NV = 0x8538
	E_TIMES_F_NV = 0x8531
	HALF_BIAS_NEGATE_NV = 0x853B
	HALF_BIAS_NORMAL_NV = 0x853A
	MAX_GENERAL_COMBINERS_NV = 0x854D
	NUM_GENERAL_COMBINERS_NV = 0x854E
	PRIMARY_COLOR_NV = 0x852C
	REGISTER_COMBINERS_NV = 0x8522
	SCALE_BY_FOUR_NV = 0x853F
	SCALE_BY_ONE_HALF_NV = 0x8540
	SCALE_BY_TWO_NV = 0x853E
	SECONDARY_COLOR_NV = 0x852D
	SIGNED_IDENTITY_NV = 0x853C
	SIGNED_NEGATE_NV = 0x853D
	SPARE0_NV = 0x852E
	SPARE0_PLUS_SECONDARY_COLOR_NV = 0x8532
	SPARE1_NV = 0x852F
	TEXTURE0_ARB = 0x84C0
	TEXTURE1_ARB = 0x84C1
	UNSIGNED_IDENTITY_NV = 0x8536
	UNSIGNED_INVERT_NV = 0x8537
	VARIABLE_A_NV = 0x8523
	VARIABLE_B_NV = 0x8524
	VARIABLE_C_NV = 0x8525
	VARIABLE_D_NV = 0x8526
	VARIABLE_E_NV = 0x8527
	VARIABLE_F_NV = 0x8528
	VARIABLE_G_NV = 0x8529
)
// NV_register_combiners2
const (
	PER_STAGE_CONSTANTS_NV = 0x8535
)
// NV_shader_atomic_float
const (
)
// NV_shader_buffer_load
const (
	BUFFER_GPU_ADDRESS_NV = 0x8F1D
	GPU_ADDRESS_NV = 0x8F34
	MAX_SHADER_BUFFER_ADDRESS_NV = 0x8F35
)
// NV_shader_buffer_store
const (
	READ_WRITE = 0x88BA
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV = 0x00000010
	WRITE_ONLY = 0x88B9
)
// NV_tessellation_program5
const (
	MAX_PROGRAM_PATCH_ATTRIBS_NV = 0x86D8
	TESS_CONTROL_PROGRAM_NV = 0x891E
	TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV = 0x8C74
	TESS_EVALUATION_PROGRAM_NV = 0x891F
	TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV = 0x8C75
)
// NV_texgen_emboss
const (
	EMBOSS_CONSTANT_NV = 0x855E
	EMBOSS_LIGHT_NV = 0x855D
	EMBOSS_MAP_NV = 0x855F
)
// NV_texgen_reflection
const (
	NORMAL_MAP_NV = 0x8511
	REFLECTION_MAP_NV = 0x8512
)
// NV_texture_barrier
const (
)
// NV_texture_compression_vtc
const (
)
// NV_texture_env_combine4
const (
	COMBINE4_NV = 0x8503
	OPERAND3_ALPHA_NV = 0x859B
	OPERAND3_RGB_NV = 0x8593
	SOURCE3_ALPHA_NV = 0x858B
	SOURCE3_RGB_NV = 0x8583
)
// NV_texture_expand_normal
const (
	TEXTURE_UNSIGNED_REMAP_MODE_NV = 0x888F
)
// NV_texture_multisample
const (
	TEXTURE_COLOR_SAMPLES_NV = 0x9046
	TEXTURE_COVERAGE_SAMPLES_NV = 0x9045
)
// NV_texture_rectangle
const (
	MAX_RECTANGLE_TEXTURE_SIZE_NV = 0x84F8
	PROXY_TEXTURE_RECTANGLE_NV = 0x84F7
	TEXTURE_BINDING_RECTANGLE_NV = 0x84F6
	TEXTURE_RECTANGLE_NV = 0x84F5
)
// NV_texture_shader
const (
	CONST_EYE_NV = 0x86E5
	CULL_FRAGMENT_NV = 0x86E7
	CULL_MODES_NV = 0x86E0
	DEPENDENT_AR_TEXTURE_2D_NV = 0x86E9
	DEPENDENT_GB_TEXTURE_2D_NV = 0x86EA
	DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV = 0x86F3
	DOT_PRODUCT_DEPTH_REPLACE_NV = 0x86ED
	DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV = 0x86F1
	DOT_PRODUCT_NV = 0x86EC
	DOT_PRODUCT_REFLECT_CUBE_MAP_NV = 0x86F2
	DOT_PRODUCT_TEXTURE_2D_NV = 0x86EE
	DOT_PRODUCT_TEXTURE_CUBE_MAP_NV = 0x86F0
	DOT_PRODUCT_TEXTURE_RECTANGLE_NV = 0x864E
	DSDT8_MAG8_INTENSITY8_NV = 0x870B
	DSDT8_MAG8_NV = 0x870A
	DSDT8_NV = 0x8709
	DSDT_MAG_INTENSITY_NV = 0x86DC
	DSDT_MAG_NV = 0x86F6
	DSDT_MAG_VIB_NV = 0x86F7
	DSDT_NV = 0x86F5
	DS_BIAS_NV = 0x8716
	DS_SCALE_NV = 0x8710
	DT_BIAS_NV = 0x8717
	DT_SCALE_NV = 0x8711
	HILO16_NV = 0x86F8
	HILO_NV = 0x86F4
	HI_BIAS_NV = 0x8714
	HI_SCALE_NV = 0x870E
	LO_BIAS_NV = 0x8715
	LO_SCALE_NV = 0x870F
	MAGNITUDE_BIAS_NV = 0x8718
	MAGNITUDE_SCALE_NV = 0x8712
	OFFSET_TEXTURE_2D_BIAS_NV = 0x86E3
	OFFSET_TEXTURE_2D_MATRIX_NV = 0x86E1
	OFFSET_TEXTURE_2D_NV = 0x86E8
	OFFSET_TEXTURE_2D_SCALE_NV = 0x86E2
	OFFSET_TEXTURE_BIAS_NV = 0x86E3
	OFFSET_TEXTURE_MATRIX_NV = 0x86E1
	OFFSET_TEXTURE_RECTANGLE_NV = 0x864C
	OFFSET_TEXTURE_RECTANGLE_SCALE_NV = 0x864D
	OFFSET_TEXTURE_SCALE_NV = 0x86E2
	PASS_THROUGH_NV = 0x86E6
	PREVIOUS_TEXTURE_INPUT_NV = 0x86E4
	RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV = 0x86D9
	SHADER_CONSISTENT_NV = 0x86DD
	SHADER_OPERATION_NV = 0x86DF
	SIGNED_ALPHA8_NV = 0x8706
	SIGNED_ALPHA_NV = 0x8705
	SIGNED_HILO16_NV = 0x86FA
	SIGNED_HILO_NV = 0x86F9
	SIGNED_INTENSITY8_NV = 0x8708
	SIGNED_INTENSITY_NV = 0x8707
	SIGNED_LUMINANCE8_ALPHA8_NV = 0x8704
	SIGNED_LUMINANCE8_NV = 0x8702
	SIGNED_LUMINANCE_ALPHA_NV = 0x8703
	SIGNED_LUMINANCE_NV = 0x8701
	SIGNED_RGB8_NV = 0x86FF
	SIGNED_RGB8_UNSIGNED_ALPHA8_NV = 0x870D
	SIGNED_RGBA8_NV = 0x86FC
	SIGNED_RGBA_NV = 0x86FB
	SIGNED_RGB_NV = 0x86FE
	SIGNED_RGB_UNSIGNED_ALPHA_NV = 0x870C
	TEXTURE_BORDER_VALUES_NV = 0x871A
	TEXTURE_DS_SIZE_NV = 0x871D
	TEXTURE_DT_SIZE_NV = 0x871E
	TEXTURE_HI_SIZE_NV = 0x871B
	TEXTURE_LO_SIZE_NV = 0x871C
	TEXTURE_MAG_SIZE_NV = 0x871F
	TEXTURE_SHADER_NV = 0x86DE
	UNSIGNED_INT_8_8_S8_S8_REV_NV = 0x86DB
	UNSIGNED_INT_S8_S8_8_8_NV = 0x86DA
	VIBRANCE_BIAS_NV = 0x8719
	VIBRANCE_SCALE_NV = 0x8713
)
// NV_texture_shader2
const (
	DOT_PRODUCT_TEXTURE_3D_NV = 0x86EF
)
// NV_texture_shader3
const (
	DEPENDENT_HILO_TEXTURE_2D_NV = 0x8858
	DEPENDENT_RGB_TEXTURE_3D_NV = 0x8859
	DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV = 0x885A
	DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV = 0x885D
	DOT_PRODUCT_PASS_THROUGH_NV = 0x885B
	DOT_PRODUCT_TEXTURE_1D_NV = 0x885C
	FORCE_BLUE_TO_ONE_NV = 0x8860
	HILO8_NV = 0x885E
	OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV = 0x8856
	OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV = 0x8857
	OFFSET_HILO_TEXTURE_2D_NV = 0x8854
	OFFSET_HILO_TEXTURE_RECTANGLE_NV = 0x8855
	OFFSET_PROJECTIVE_TEXTURE_2D_NV = 0x8850
	OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV = 0x8851
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV = 0x8852
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV = 0x8853
	SIGNED_HILO8_NV = 0x885F
)
// NV_transform_feedback
const (
	ACTIVE_VARYINGS_NV = 0x8C81
	ACTIVE_VARYING_MAX_LENGTH_NV = 0x8C82
	BACK_PRIMARY_COLOR_NV = 0x8C77
	BACK_SECONDARY_COLOR_NV = 0x8C78
	CLIP_DISTANCE_NV = 0x8C7A
	GENERIC_ATTRIB_NV = 0x8C7D
	INTERLEAVED_ATTRIBS_NV = 0x8C8C
	LAYER_NV = 0x8DAA
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV = 0x8C8A
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV = 0x8C8B
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV = 0x8C80
	NEXT_BUFFER_NV = -2
	PRIMITIVES_GENERATED_NV = 0x8C87
	PRIMITIVE_ID_NV = 0x8C7C
	RASTERIZER_DISCARD_NV = 0x8C89
	SEPARATE_ATTRIBS_NV = 0x8C8D
	SKIP_COMPONENTS1_NV = -6
	SKIP_COMPONENTS2_NV = -5
	SKIP_COMPONENTS3_NV = -4
	SKIP_COMPONENTS4_NV = -3
	TEXTURE_COORD_NV = 0x8C79
	TRANSFORM_FEEDBACK_ATTRIBS_NV = 0x8C7E
	TRANSFORM_FEEDBACK_BUFFER_BINDING_NV = 0x8C8F
	TRANSFORM_FEEDBACK_BUFFER_MODE_NV = 0x8C7F
	TRANSFORM_FEEDBACK_BUFFER_NV = 0x8C8E
	TRANSFORM_FEEDBACK_BUFFER_SIZE_NV = 0x8C85
	TRANSFORM_FEEDBACK_BUFFER_START_NV = 0x8C84
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV = 0x8C88
	TRANSFORM_FEEDBACK_RECORD_NV = 0x8C86
	TRANSFORM_FEEDBACK_VARYINGS_NV = 0x8C83
	VERTEX_ID_NV = 0x8C7B
)
// NV_transform_feedback2
const (
	TRANSFORM_FEEDBACK_BINDING_NV = 0x8E25
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV = 0x8E24
	TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV = 0x8E23
	TRANSFORM_FEEDBACK_NV = 0x8E22
)
// NV_vdpau_interop
const (
	SURFACE_MAPPED_NV = 0x8700
	SURFACE_REGISTERED_NV = 0x86FD
	SURFACE_STATE_NV = 0x86EB
	WRITE_DISCARD_NV = 0x88BE
)
// NV_vertex_array_range
const (
	MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV = 0x8520
	VERTEX_ARRAY_RANGE_LENGTH_NV = 0x851E
	VERTEX_ARRAY_RANGE_NV = 0x851D
	VERTEX_ARRAY_RANGE_POINTER_NV = 0x8521
	VERTEX_ARRAY_RANGE_VALID_NV = 0x851F
)
// NV_vertex_array_range2
const (
	VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV = 0x8533
)
// NV_vertex_attrib_integer_64bit
const (
)
// NV_vertex_buffer_unified_memory
const (
	COLOR_ARRAY_ADDRESS_NV = 0x8F23
	COLOR_ARRAY_LENGTH_NV = 0x8F2D
	DRAW_INDIRECT_ADDRESS_NV = 0x8F41
	DRAW_INDIRECT_LENGTH_NV = 0x8F42
	DRAW_INDIRECT_UNIFIED_NV = 0x8F40
	EDGE_FLAG_ARRAY_ADDRESS_NV = 0x8F26
	EDGE_FLAG_ARRAY_LENGTH_NV = 0x8F30
	ELEMENT_ARRAY_ADDRESS_NV = 0x8F29
	ELEMENT_ARRAY_LENGTH_NV = 0x8F33
	ELEMENT_ARRAY_UNIFIED_NV = 0x8F1F
	FOG_COORD_ARRAY_ADDRESS_NV = 0x8F28
	FOG_COORD_ARRAY_LENGTH_NV = 0x8F32
	INDEX_ARRAY_ADDRESS_NV = 0x8F24
	INDEX_ARRAY_LENGTH_NV = 0x8F2E
	NORMAL_ARRAY_ADDRESS_NV = 0x8F22
	NORMAL_ARRAY_LENGTH_NV = 0x8F2C
	SECONDARY_COLOR_ARRAY_ADDRESS_NV = 0x8F27
	SECONDARY_COLOR_ARRAY_LENGTH_NV = 0x8F31
	TEXTURE_COORD_ARRAY_ADDRESS_NV = 0x8F25
	TEXTURE_COORD_ARRAY_LENGTH_NV = 0x8F2F
	VERTEX_ARRAY_ADDRESS_NV = 0x8F21
	VERTEX_ARRAY_LENGTH_NV = 0x8F2B
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV = 0x8F20
	VERTEX_ATTRIB_ARRAY_LENGTH_NV = 0x8F2A
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV = 0x8F1E
)
// NV_vertex_program
const (
	ATTRIB_ARRAY_POINTER_NV = 0x8645
	ATTRIB_ARRAY_SIZE_NV = 0x8623
	ATTRIB_ARRAY_STRIDE_NV = 0x8624
	ATTRIB_ARRAY_TYPE_NV = 0x8625
	CURRENT_ATTRIB_NV = 0x8626
	CURRENT_MATRIX_NV = 0x8641
	CURRENT_MATRIX_STACK_DEPTH_NV = 0x8640
	IDENTITY_NV = 0x862A
	INVERSE_NV = 0x862B
	INVERSE_TRANSPOSE_NV = 0x862D
	MAP1_VERTEX_ATTRIB0_4_NV = 0x8660
	MAP1_VERTEX_ATTRIB10_4_NV = 0x866A
	MAP1_VERTEX_ATTRIB11_4_NV = 0x866B
	MAP1_VERTEX_ATTRIB12_4_NV = 0x866C
	MAP1_VERTEX_ATTRIB13_4_NV = 0x866D
	MAP1_VERTEX_ATTRIB14_4_NV = 0x866E
	MAP1_VERTEX_ATTRIB15_4_NV = 0x866F
	MAP1_VERTEX_ATTRIB1_4_NV = 0x8661
	MAP1_VERTEX_ATTRIB2_4_NV = 0x8662
	MAP1_VERTEX_ATTRIB3_4_NV = 0x8663
	MAP1_VERTEX_ATTRIB4_4_NV = 0x8664
	MAP1_VERTEX_ATTRIB5_4_NV = 0x8665
	MAP1_VERTEX_ATTRIB6_4_NV = 0x8666
	MAP1_VERTEX_ATTRIB7_4_NV = 0x8667
	MAP1_VERTEX_ATTRIB8_4_NV = 0x8668
	MAP1_VERTEX_ATTRIB9_4_NV = 0x8669
	MAP2_VERTEX_ATTRIB0_4_NV = 0x8670
	MAP2_VERTEX_ATTRIB10_4_NV = 0x867A
	MAP2_VERTEX_ATTRIB11_4_NV = 0x867B
	MAP2_VERTEX_ATTRIB12_4_NV = 0x867C
	MAP2_VERTEX_ATTRIB13_4_NV = 0x867D
	MAP2_VERTEX_ATTRIB14_4_NV = 0x867E
	MAP2_VERTEX_ATTRIB15_4_NV = 0x867F
	MAP2_VERTEX_ATTRIB1_4_NV = 0x8671
	MAP2_VERTEX_ATTRIB2_4_NV = 0x8672
	MAP2_VERTEX_ATTRIB3_4_NV = 0x8673
	MAP2_VERTEX_ATTRIB4_4_NV = 0x8674
	MAP2_VERTEX_ATTRIB5_4_NV = 0x8675
	MAP2_VERTEX_ATTRIB6_4_NV = 0x8676
	MAP2_VERTEX_ATTRIB7_4_NV = 0x8677
	MAP2_VERTEX_ATTRIB8_4_NV = 0x8678
	MAP2_VERTEX_ATTRIB9_4_NV = 0x8679
	MATRIX0_NV = 0x8630
	MATRIX1_NV = 0x8631
	MATRIX2_NV = 0x8632
	MATRIX3_NV = 0x8633
	MATRIX4_NV = 0x8634
	MATRIX5_NV = 0x8635
	MATRIX6_NV = 0x8636
	MATRIX7_NV = 0x8637
	MAX_TRACK_MATRICES_NV = 0x862F
	MAX_TRACK_MATRIX_STACK_DEPTH_NV = 0x862E
	MODELVIEW_PROJECTION_NV = 0x8629
	PROGRAM_ERROR_POSITION_NV = 0x864B
	PROGRAM_LENGTH_NV = 0x8627
	PROGRAM_PARAMETER_NV = 0x8644
	PROGRAM_RESIDENT_NV = 0x8647
	PROGRAM_STRING_NV = 0x8628
	PROGRAM_TARGET_NV = 0x8646
	TRACK_MATRIX_NV = 0x8648
	TRACK_MATRIX_TRANSFORM_NV = 0x8649
	TRANSPOSE_NV = 0x862C
	VERTEX_ATTRIB_ARRAY0_NV = 0x8650
	VERTEX_ATTRIB_ARRAY10_NV = 0x865A
	VERTEX_ATTRIB_ARRAY11_NV = 0x865B
	VERTEX_ATTRIB_ARRAY12_NV = 0x865C
	VERTEX_ATTRIB_ARRAY13_NV = 0x865D
	VERTEX_ATTRIB_ARRAY14_NV = 0x865E
	VERTEX_ATTRIB_ARRAY15_NV = 0x865F
	VERTEX_ATTRIB_ARRAY1_NV = 0x8651
	VERTEX_ATTRIB_ARRAY2_NV = 0x8652
	VERTEX_ATTRIB_ARRAY3_NV = 0x8653
	VERTEX_ATTRIB_ARRAY4_NV = 0x8654
	VERTEX_ATTRIB_ARRAY5_NV = 0x8655
	VERTEX_ATTRIB_ARRAY6_NV = 0x8656
	VERTEX_ATTRIB_ARRAY7_NV = 0x8657
	VERTEX_ATTRIB_ARRAY8_NV = 0x8658
	VERTEX_ATTRIB_ARRAY9_NV = 0x8659
	VERTEX_PROGRAM_BINDING_NV = 0x864A
	VERTEX_PROGRAM_NV = 0x8620
	VERTEX_PROGRAM_POINT_SIZE_NV = 0x8642
	VERTEX_PROGRAM_TWO_SIDE_NV = 0x8643
	VERTEX_STATE_PROGRAM_NV = 0x8621
)
// NV_vertex_program1_1
const (
)
// NV_vertex_program2
const (
)
// NV_vertex_program2_option
const (
)
// NV_vertex_program3
const (
	MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB = 0x8B4C
)
// NV_vertex_program4
const (
	VERTEX_ATTRIB_ARRAY_INTEGER_NV = 0x88FD
)
// NV_video_capture
const (
	FAILURE_NV = 0x9030
	FIELD_LOWER_NV = 0x9023
	FIELD_UPPER_NV = 0x9022
	LAST_VIDEO_CAPTURE_STATUS_NV = 0x9027
	NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV = 0x9025
	NUM_VIDEO_CAPTURE_STREAMS_NV = 0x9024
	PARTIAL_SUCCESS_NV = 0x902E
	SUCCESS_NV = 0x902F
	VIDEO_BUFFER_BINDING_NV = 0x9021
	VIDEO_BUFFER_INTERNAL_FORMAT_NV = 0x902D
	VIDEO_BUFFER_NV = 0x9020
	VIDEO_BUFFER_PITCH_NV = 0x9028
	VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV = 0x903B
	VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV = 0x903A
	VIDEO_CAPTURE_FRAME_HEIGHT_NV = 0x9039
	VIDEO_CAPTURE_FRAME_WIDTH_NV = 0x9038
	VIDEO_CAPTURE_SURFACE_ORIGIN_NV = 0x903C
	VIDEO_CAPTURE_TO_422_SUPPORTED_NV = 0x9026
	VIDEO_COLOR_CONVERSION_MATRIX_NV = 0x9029
	VIDEO_COLOR_CONVERSION_MAX_NV = 0x902A
	VIDEO_COLOR_CONVERSION_MIN_NV = 0x902B
	VIDEO_COLOR_CONVERSION_OFFSET_NV = 0x902C
	YCBAYCR8A_4224_NV = 0x9032
	YCBYCR8_422_NV = 0x9031
	Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV = 0x9036
	Z4Y12Z4CB12Z4CR12_444_NV = 0x9037
	Z4Y12Z4CB12Z4Y12Z4CR12_422_NV = 0x9035
	Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV = 0x9034
	Z6Y10Z6CB10Z6Y10Z6CR10_422_NV = 0x9033
)
// NV_bindless_texture

func GetTextureHandleNV(texture Uint) Uint64 {
	return (Uint64)(C.goglGetTextureHandleNV((C.GLuint)(texture)))
}
func GetTextureSamplerHandleNV(texture Uint, sampler Uint) Uint64 {
	return (Uint64)(C.goglGetTextureSamplerHandleNV((C.GLuint)(texture), (C.GLuint)(sampler)))
}
func MakeTextureHandleResidentNV(handle Uint64)  {
	C.goglMakeTextureHandleResidentNV((C.GLuint64)(handle))
}
func MakeTextureHandleNonResidentNV(handle Uint64)  {
	C.goglMakeTextureHandleNonResidentNV((C.GLuint64)(handle))
}
func GetImageHandleNV(texture Uint, level Int, layered Boolean, layer Int, format Enum) Uint64 {
	return (Uint64)(C.goglGetImageHandleNV((C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(layered), (C.GLint)(layer), (C.GLenum)(format)))
}
func MakeImageHandleResidentNV(handle Uint64, access Enum)  {
	C.goglMakeImageHandleResidentNV((C.GLuint64)(handle), (C.GLenum)(access))
}
func MakeImageHandleNonResidentNV(handle Uint64)  {
	C.goglMakeImageHandleNonResidentNV((C.GLuint64)(handle))
}
func UniformHandleui64NV(location Int, value Uint64)  {
	C.goglUniformHandleui64NV((C.GLint)(location), (C.GLuint64)(value))
}
func UniformHandleui64vNV(location Int, count Sizei, value *Uint64)  {
	C.goglUniformHandleui64vNV((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(value))
}
func ProgramUniformHandleui64NV(program Uint, location Int, value Uint64)  {
	C.goglProgramUniformHandleui64NV((C.GLuint)(program), (C.GLint)(location), (C.GLuint64)(value))
}
func ProgramUniformHandleui64vNV(program Uint, location Int, count Sizei, values *Uint64)  {
	C.goglProgramUniformHandleui64vNV((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(values))
}
func IsTextureHandleResidentNV(handle Uint64) Boolean {
	return (Boolean)(C.goglIsTextureHandleResidentNV((C.GLuint64)(handle)))
}
func IsImageHandleResidentNV(handle Uint64) Boolean {
	return (Boolean)(C.goglIsImageHandleResidentNV((C.GLuint64)(handle)))
}
// NV_conditional_render

func BeginConditionalRenderNV(id Uint, mode Enum)  {
	C.goglBeginConditionalRenderNV((C.GLuint)(id), (C.GLenum)(mode))
}
func EndConditionalRenderNV()  {
	C.goglEndConditionalRenderNV()
}
// NV_copy_image

func CopyImageSubDataNV(srcName Uint, srcTarget Enum, srcLevel Int, srcX Int, srcY Int, srcZ Int, dstName Uint, dstTarget Enum, dstLevel Int, dstX Int, dstY Int, dstZ Int, width Sizei, height Sizei, depth Sizei)  {
	C.goglCopyImageSubDataNV((C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
// NV_depth_buffer_float

func DepthRangedNV(zNear Double, zFar Double)  {
	C.goglDepthRangedNV((C.GLdouble)(zNear), (C.GLdouble)(zFar))
}
func ClearDepthdNV(depth Double)  {
	C.goglClearDepthdNV((C.GLdouble)(depth))
}
func DepthBoundsdNV(zmin Double, zmax Double)  {
	C.goglDepthBoundsdNV((C.GLdouble)(zmin), (C.GLdouble)(zmax))
}
// NV_evaluators

func MapControlPointsNV(target Enum, index Uint, type_ Enum, ustride Sizei, vstride Sizei, uorder Int, vorder Int, packed Boolean, points Pointer)  {
	C.goglMapControlPointsNV((C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(type_), (C.GLsizei)(ustride), (C.GLsizei)(vstride), (C.GLint)(uorder), (C.GLint)(vorder), (C.GLboolean)(packed), (unsafe.Pointer)(points))
}
func MapParameterivNV(target Enum, pname Enum, params *Int)  {
	C.goglMapParameterivNV((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func MapParameterfvNV(target Enum, pname Enum, params *Float)  {
	C.goglMapParameterfvNV((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetMapControlPointsNV(target Enum, index Uint, type_ Enum, ustride Sizei, vstride Sizei, packed Boolean, points Pointer)  {
	C.goglGetMapControlPointsNV((C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(type_), (C.GLsizei)(ustride), (C.GLsizei)(vstride), (C.GLboolean)(packed), (unsafe.Pointer)(points))
}
func GetMapParameterivNV(target Enum, pname Enum, params *Int)  {
	C.goglGetMapParameterivNV((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetMapParameterfvNV(target Enum, pname Enum, params *Float)  {
	C.goglGetMapParameterfvNV((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetMapAttribParameterivNV(target Enum, index Uint, pname Enum, params *Int)  {
	C.goglGetMapAttribParameterivNV((C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetMapAttribParameterfvNV(target Enum, index Uint, pname Enum, params *Float)  {
	C.goglGetMapAttribParameterfvNV((C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func EvalMapsNV(target Enum, mode Enum)  {
	C.goglEvalMapsNV((C.GLenum)(target), (C.GLenum)(mode))
}
// NV_explicit_multisample

func GetMultisamplefvNV(pname Enum, index Uint, val *Float)  {
	C.goglGetMultisamplefvNV((C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(val))
}
func SampleMaskIndexedNV(index Uint, mask Bitfield)  {
	C.goglSampleMaskIndexedNV((C.GLuint)(index), (C.GLbitfield)(mask))
}
func TexRenderbufferNV(target Enum, renderbuffer Uint)  {
	C.goglTexRenderbufferNV((C.GLenum)(target), (C.GLuint)(renderbuffer))
}
// NV_fence

func DeleteFencesNV(n Sizei, fences *Uint)  {
	C.goglDeleteFencesNV((C.GLsizei)(n), (*C.GLuint)(fences))
}
func GenFencesNV(n Sizei, fences *Uint)  {
	C.goglGenFencesNV((C.GLsizei)(n), (*C.GLuint)(fences))
}
func IsFenceNV(fence Uint) Boolean {
	return (Boolean)(C.goglIsFenceNV((C.GLuint)(fence)))
}
func TestFenceNV(fence Uint) Boolean {
	return (Boolean)(C.goglTestFenceNV((C.GLuint)(fence)))
}
func GetFenceivNV(fence Uint, pname Enum, params *Int)  {
	C.goglGetFenceivNV((C.GLuint)(fence), (C.GLenum)(pname), (*C.GLint)(params))
}
func FinishFenceNV(fence Uint)  {
	C.goglFinishFenceNV((C.GLuint)(fence))
}
func SetFenceNV(fence Uint, condition Enum)  {
	C.goglSetFenceNV((C.GLuint)(fence), (C.GLenum)(condition))
}
// NV_fragment_program

func ProgramNamedParameter4fNV(id Uint, len Sizei, name *Ubyte, x Float, y Float, z Float, w Float)  {
	C.goglProgramNamedParameter4fNV((C.GLuint)(id), (C.GLsizei)(len), (*C.GLubyte)(name), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func ProgramNamedParameter4dNV(id Uint, len Sizei, name *Ubyte, x Double, y Double, z Double, w Double)  {
	C.goglProgramNamedParameter4dNV((C.GLuint)(id), (C.GLsizei)(len), (*C.GLubyte)(name), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func ProgramNamedParameter4fvNV(id Uint, len Sizei, name *Ubyte, v *Float)  {
	C.goglProgramNamedParameter4fvNV((C.GLuint)(id), (C.GLsizei)(len), (*C.GLubyte)(name), (*C.GLfloat)(v))
}
func ProgramNamedParameter4dvNV(id Uint, len Sizei, name *Ubyte, v *Double)  {
	C.goglProgramNamedParameter4dvNV((C.GLuint)(id), (C.GLsizei)(len), (*C.GLubyte)(name), (*C.GLdouble)(v))
}
func GetProgramNamedParameterfvNV(id Uint, len Sizei, name *Ubyte, params *Float)  {
	C.goglGetProgramNamedParameterfvNV((C.GLuint)(id), (C.GLsizei)(len), (*C.GLubyte)(name), (*C.GLfloat)(params))
}
func GetProgramNamedParameterdvNV(id Uint, len Sizei, name *Ubyte, params *Double)  {
	C.goglGetProgramNamedParameterdvNV((C.GLuint)(id), (C.GLsizei)(len), (*C.GLubyte)(name), (*C.GLdouble)(params))
}
// NV_framebuffer_multisample_coverage

func RenderbufferStorageMultisampleCoverageNV(target Enum, coverageSamples Sizei, colorSamples Sizei, internalformat Enum, width Sizei, height Sizei)  {
	C.goglRenderbufferStorageMultisampleCoverageNV((C.GLenum)(target), (C.GLsizei)(coverageSamples), (C.GLsizei)(colorSamples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
// NV_geometry_program4

func ProgramVertexLimitNV(target Enum, limit Int)  {
	C.goglProgramVertexLimitNV((C.GLenum)(target), (C.GLint)(limit))
}
func FramebufferTextureEXT(target Enum, attachment Enum, texture Uint, level Int)  {
	C.goglFramebufferTextureEXT((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTextureLayerEXT(target Enum, attachment Enum, texture Uint, level Int, layer Int)  {
	C.goglFramebufferTextureLayerEXT((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}
func FramebufferTextureFaceEXT(target Enum, attachment Enum, texture Uint, level Int, face Enum)  {
	C.goglFramebufferTextureFaceEXT((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(face))
}
// NV_gpu_program4

func ProgramLocalParameterI4iNV(target Enum, index Uint, x Int, y Int, z Int, w Int)  {
	C.goglProgramLocalParameterI4iNV((C.GLenum)(target), (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func ProgramLocalParameterI4ivNV(target Enum, index Uint, params *Int)  {
	C.goglProgramLocalParameterI4ivNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(params))
}
func ProgramLocalParametersI4ivNV(target Enum, index Uint, count Sizei, params *Int)  {
	C.goglProgramLocalParametersI4ivNV((C.GLenum)(target), (C.GLuint)(index), (C.GLsizei)(count), (*C.GLint)(params))
}
func ProgramLocalParameterI4uiNV(target Enum, index Uint, x Uint, y Uint, z Uint, w Uint)  {
	C.goglProgramLocalParameterI4uiNV((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}
func ProgramLocalParameterI4uivNV(target Enum, index Uint, params *Uint)  {
	C.goglProgramLocalParameterI4uivNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLuint)(params))
}
func ProgramLocalParametersI4uivNV(target Enum, index Uint, count Sizei, params *Uint)  {
	C.goglProgramLocalParametersI4uivNV((C.GLenum)(target), (C.GLuint)(index), (C.GLsizei)(count), (*C.GLuint)(params))
}
func ProgramEnvParameterI4iNV(target Enum, index Uint, x Int, y Int, z Int, w Int)  {
	C.goglProgramEnvParameterI4iNV((C.GLenum)(target), (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func ProgramEnvParameterI4ivNV(target Enum, index Uint, params *Int)  {
	C.goglProgramEnvParameterI4ivNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(params))
}
func ProgramEnvParametersI4ivNV(target Enum, index Uint, count Sizei, params *Int)  {
	C.goglProgramEnvParametersI4ivNV((C.GLenum)(target), (C.GLuint)(index), (C.GLsizei)(count), (*C.GLint)(params))
}
func ProgramEnvParameterI4uiNV(target Enum, index Uint, x Uint, y Uint, z Uint, w Uint)  {
	C.goglProgramEnvParameterI4uiNV((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}
func ProgramEnvParameterI4uivNV(target Enum, index Uint, params *Uint)  {
	C.goglProgramEnvParameterI4uivNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLuint)(params))
}
func ProgramEnvParametersI4uivNV(target Enum, index Uint, count Sizei, params *Uint)  {
	C.goglProgramEnvParametersI4uivNV((C.GLenum)(target), (C.GLuint)(index), (C.GLsizei)(count), (*C.GLuint)(params))
}
func GetProgramLocalParameterIivNV(target Enum, index Uint, params *Int)  {
	C.goglGetProgramLocalParameterIivNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(params))
}
func GetProgramLocalParameterIuivNV(target Enum, index Uint, params *Uint)  {
	C.goglGetProgramLocalParameterIuivNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLuint)(params))
}
func GetProgramEnvParameterIivNV(target Enum, index Uint, params *Int)  {
	C.goglGetProgramEnvParameterIivNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(params))
}
func GetProgramEnvParameterIuivNV(target Enum, index Uint, params *Uint)  {
	C.goglGetProgramEnvParameterIuivNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLuint)(params))
}
// NV_gpu_program5

func ProgramSubroutineParametersuivNV(target Enum, count Sizei, params *Uint)  {
	C.goglProgramSubroutineParametersuivNV((C.GLenum)(target), (C.GLsizei)(count), (*C.GLuint)(params))
}
func GetProgramSubroutineParameteruivNV(target Enum, index Uint, param *Uint)  {
	C.goglGetProgramSubroutineParameteruivNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLuint)(param))
}
// NV_gpu_shader5

func Uniform1i64NV(location Int, x Int64)  {
	C.goglUniform1i64NV((C.GLint)(location), (C.GLint64EXT)(x))
}
func Uniform2i64NV(location Int, x Int64, y Int64)  {
	C.goglUniform2i64NV((C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y))
}
func Uniform3i64NV(location Int, x Int64, y Int64, z Int64)  {
	C.goglUniform3i64NV((C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y), (C.GLint64EXT)(z))
}
func Uniform4i64NV(location Int, x Int64, y Int64, z Int64, w Int64)  {
	C.goglUniform4i64NV((C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y), (C.GLint64EXT)(z), (C.GLint64EXT)(w))
}
func Uniform1i64vNV(location Int, count Sizei, value *Int64)  {
	C.goglUniform1i64vNV((C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(value))
}
func Uniform2i64vNV(location Int, count Sizei, value *Int64)  {
	C.goglUniform2i64vNV((C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(value))
}
func Uniform3i64vNV(location Int, count Sizei, value *Int64)  {
	C.goglUniform3i64vNV((C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(value))
}
func Uniform4i64vNV(location Int, count Sizei, value *Int64)  {
	C.goglUniform4i64vNV((C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(value))
}
func Uniform1ui64NV(location Int, x Uint64)  {
	C.goglUniform1ui64NV((C.GLint)(location), (C.GLuint64EXT)(x))
}
func Uniform2ui64NV(location Int, x Uint64, y Uint64)  {
	C.goglUniform2ui64NV((C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y))
}
func Uniform3ui64NV(location Int, x Uint64, y Uint64, z Uint64)  {
	C.goglUniform3ui64NV((C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y), (C.GLuint64EXT)(z))
}
func Uniform4ui64NV(location Int, x Uint64, y Uint64, z Uint64, w Uint64)  {
	C.goglUniform4ui64NV((C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y), (C.GLuint64EXT)(z), (C.GLuint64EXT)(w))
}
func Uniform1ui64vNV(location Int, count Sizei, value *Uint64)  {
	C.goglUniform1ui64vNV((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(value))
}
func Uniform2ui64vNV(location Int, count Sizei, value *Uint64)  {
	C.goglUniform2ui64vNV((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(value))
}
func Uniform3ui64vNV(location Int, count Sizei, value *Uint64)  {
	C.goglUniform3ui64vNV((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(value))
}
func Uniform4ui64vNV(location Int, count Sizei, value *Uint64)  {
	C.goglUniform4ui64vNV((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(value))
}
func GetUniformi64vNV(program Uint, location Int, params *Int64)  {
	C.goglGetUniformi64vNV((C.GLuint)(program), (C.GLint)(location), (*C.GLint64EXT)(params))
}
func ProgramUniform1i64NV(program Uint, location Int, x Int64)  {
	C.goglProgramUniform1i64NV((C.GLuint)(program), (C.GLint)(location), (C.GLint64EXT)(x))
}
func ProgramUniform2i64NV(program Uint, location Int, x Int64, y Int64)  {
	C.goglProgramUniform2i64NV((C.GLuint)(program), (C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y))
}
func ProgramUniform3i64NV(program Uint, location Int, x Int64, y Int64, z Int64)  {
	C.goglProgramUniform3i64NV((C.GLuint)(program), (C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y), (C.GLint64EXT)(z))
}
func ProgramUniform4i64NV(program Uint, location Int, x Int64, y Int64, z Int64, w Int64)  {
	C.goglProgramUniform4i64NV((C.GLuint)(program), (C.GLint)(location), (C.GLint64EXT)(x), (C.GLint64EXT)(y), (C.GLint64EXT)(z), (C.GLint64EXT)(w))
}
func ProgramUniform1i64vNV(program Uint, location Int, count Sizei, value *Int64)  {
	C.goglProgramUniform1i64vNV((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(value))
}
func ProgramUniform2i64vNV(program Uint, location Int, count Sizei, value *Int64)  {
	C.goglProgramUniform2i64vNV((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(value))
}
func ProgramUniform3i64vNV(program Uint, location Int, count Sizei, value *Int64)  {
	C.goglProgramUniform3i64vNV((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(value))
}
func ProgramUniform4i64vNV(program Uint, location Int, count Sizei, value *Int64)  {
	C.goglProgramUniform4i64vNV((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint64EXT)(value))
}
func ProgramUniform1ui64NV(program Uint, location Int, x Uint64)  {
	C.goglProgramUniform1ui64NV((C.GLuint)(program), (C.GLint)(location), (C.GLuint64EXT)(x))
}
func ProgramUniform2ui64NV(program Uint, location Int, x Uint64, y Uint64)  {
	C.goglProgramUniform2ui64NV((C.GLuint)(program), (C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y))
}
func ProgramUniform3ui64NV(program Uint, location Int, x Uint64, y Uint64, z Uint64)  {
	C.goglProgramUniform3ui64NV((C.GLuint)(program), (C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y), (C.GLuint64EXT)(z))
}
func ProgramUniform4ui64NV(program Uint, location Int, x Uint64, y Uint64, z Uint64, w Uint64)  {
	C.goglProgramUniform4ui64NV((C.GLuint)(program), (C.GLint)(location), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y), (C.GLuint64EXT)(z), (C.GLuint64EXT)(w))
}
func ProgramUniform1ui64vNV(program Uint, location Int, count Sizei, value *Uint64)  {
	C.goglProgramUniform1ui64vNV((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(value))
}
func ProgramUniform2ui64vNV(program Uint, location Int, count Sizei, value *Uint64)  {
	C.goglProgramUniform2ui64vNV((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(value))
}
func ProgramUniform3ui64vNV(program Uint, location Int, count Sizei, value *Uint64)  {
	C.goglProgramUniform3ui64vNV((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(value))
}
func ProgramUniform4ui64vNV(program Uint, location Int, count Sizei, value *Uint64)  {
	C.goglProgramUniform4ui64vNV((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(value))
}
// NV_half_float

func Vertex2hNV(x Half, y Half)  {
	C.goglVertex2hNV((C.GLhalfNV)(x), (C.GLhalfNV)(y))
}
func Vertex2hvNV(v *Half)  {
	C.goglVertex2hvNV((*C.GLhalfNV)(v))
}
func Vertex3hNV(x Half, y Half, z Half)  {
	C.goglVertex3hNV((C.GLhalfNV)(x), (C.GLhalfNV)(y), (C.GLhalfNV)(z))
}
func Vertex3hvNV(v *Half)  {
	C.goglVertex3hvNV((*C.GLhalfNV)(v))
}
func Vertex4hNV(x Half, y Half, z Half, w Half)  {
	C.goglVertex4hNV((C.GLhalfNV)(x), (C.GLhalfNV)(y), (C.GLhalfNV)(z), (C.GLhalfNV)(w))
}
func Vertex4hvNV(v *Half)  {
	C.goglVertex4hvNV((*C.GLhalfNV)(v))
}
func Normal3hNV(nx Half, ny Half, nz Half)  {
	C.goglNormal3hNV((C.GLhalfNV)(nx), (C.GLhalfNV)(ny), (C.GLhalfNV)(nz))
}
func Normal3hvNV(v *Half)  {
	C.goglNormal3hvNV((*C.GLhalfNV)(v))
}
func Color3hNV(red Half, green Half, blue Half)  {
	C.goglColor3hNV((C.GLhalfNV)(red), (C.GLhalfNV)(green), (C.GLhalfNV)(blue))
}
func Color3hvNV(v *Half)  {
	C.goglColor3hvNV((*C.GLhalfNV)(v))
}
func Color4hNV(red Half, green Half, blue Half, alpha Half)  {
	C.goglColor4hNV((C.GLhalfNV)(red), (C.GLhalfNV)(green), (C.GLhalfNV)(blue), (C.GLhalfNV)(alpha))
}
func Color4hvNV(v *Half)  {
	C.goglColor4hvNV((*C.GLhalfNV)(v))
}
func TexCoord1hNV(s Half)  {
	C.goglTexCoord1hNV((C.GLhalfNV)(s))
}
func TexCoord1hvNV(v *Half)  {
	C.goglTexCoord1hvNV((*C.GLhalfNV)(v))
}
func TexCoord2hNV(s Half, t Half)  {
	C.goglTexCoord2hNV((C.GLhalfNV)(s), (C.GLhalfNV)(t))
}
func TexCoord2hvNV(v *Half)  {
	C.goglTexCoord2hvNV((*C.GLhalfNV)(v))
}
func TexCoord3hNV(s Half, t Half, r Half)  {
	C.goglTexCoord3hNV((C.GLhalfNV)(s), (C.GLhalfNV)(t), (C.GLhalfNV)(r))
}
func TexCoord3hvNV(v *Half)  {
	C.goglTexCoord3hvNV((*C.GLhalfNV)(v))
}
func TexCoord4hNV(s Half, t Half, r Half, q Half)  {
	C.goglTexCoord4hNV((C.GLhalfNV)(s), (C.GLhalfNV)(t), (C.GLhalfNV)(r), (C.GLhalfNV)(q))
}
func TexCoord4hvNV(v *Half)  {
	C.goglTexCoord4hvNV((*C.GLhalfNV)(v))
}
func MultiTexCoord1hNV(target Enum, s Half)  {
	C.goglMultiTexCoord1hNV((C.GLenum)(target), (C.GLhalfNV)(s))
}
func MultiTexCoord1hvNV(target Enum, v *Half)  {
	C.goglMultiTexCoord1hvNV((C.GLenum)(target), (*C.GLhalfNV)(v))
}
func MultiTexCoord2hNV(target Enum, s Half, t Half)  {
	C.goglMultiTexCoord2hNV((C.GLenum)(target), (C.GLhalfNV)(s), (C.GLhalfNV)(t))
}
func MultiTexCoord2hvNV(target Enum, v *Half)  {
	C.goglMultiTexCoord2hvNV((C.GLenum)(target), (*C.GLhalfNV)(v))
}
func MultiTexCoord3hNV(target Enum, s Half, t Half, r Half)  {
	C.goglMultiTexCoord3hNV((C.GLenum)(target), (C.GLhalfNV)(s), (C.GLhalfNV)(t), (C.GLhalfNV)(r))
}
func MultiTexCoord3hvNV(target Enum, v *Half)  {
	C.goglMultiTexCoord3hvNV((C.GLenum)(target), (*C.GLhalfNV)(v))
}
func MultiTexCoord4hNV(target Enum, s Half, t Half, r Half, q Half)  {
	C.goglMultiTexCoord4hNV((C.GLenum)(target), (C.GLhalfNV)(s), (C.GLhalfNV)(t), (C.GLhalfNV)(r), (C.GLhalfNV)(q))
}
func MultiTexCoord4hvNV(target Enum, v *Half)  {
	C.goglMultiTexCoord4hvNV((C.GLenum)(target), (*C.GLhalfNV)(v))
}
func FogCoordhNV(fog Half)  {
	C.goglFogCoordhNV((C.GLhalfNV)(fog))
}
func FogCoordhvNV(fog *Half)  {
	C.goglFogCoordhvNV((*C.GLhalfNV)(fog))
}
func SecondaryColor3hNV(red Half, green Half, blue Half)  {
	C.goglSecondaryColor3hNV((C.GLhalfNV)(red), (C.GLhalfNV)(green), (C.GLhalfNV)(blue))
}
func SecondaryColor3hvNV(v *Half)  {
	C.goglSecondaryColor3hvNV((*C.GLhalfNV)(v))
}
func VertexWeighthNV(weight Half)  {
	C.goglVertexWeighthNV((C.GLhalfNV)(weight))
}
func VertexWeighthvNV(weight *Half)  {
	C.goglVertexWeighthvNV((*C.GLhalfNV)(weight))
}
func VertexAttrib1hNV(index Uint, x Half)  {
	C.goglVertexAttrib1hNV((C.GLuint)(index), (C.GLhalfNV)(x))
}
func VertexAttrib1hvNV(index Uint, v *Half)  {
	C.goglVertexAttrib1hvNV((C.GLuint)(index), (*C.GLhalfNV)(v))
}
func VertexAttrib2hNV(index Uint, x Half, y Half)  {
	C.goglVertexAttrib2hNV((C.GLuint)(index), (C.GLhalfNV)(x), (C.GLhalfNV)(y))
}
func VertexAttrib2hvNV(index Uint, v *Half)  {
	C.goglVertexAttrib2hvNV((C.GLuint)(index), (*C.GLhalfNV)(v))
}
func VertexAttrib3hNV(index Uint, x Half, y Half, z Half)  {
	C.goglVertexAttrib3hNV((C.GLuint)(index), (C.GLhalfNV)(x), (C.GLhalfNV)(y), (C.GLhalfNV)(z))
}
func VertexAttrib3hvNV(index Uint, v *Half)  {
	C.goglVertexAttrib3hvNV((C.GLuint)(index), (*C.GLhalfNV)(v))
}
func VertexAttrib4hNV(index Uint, x Half, y Half, z Half, w Half)  {
	C.goglVertexAttrib4hNV((C.GLuint)(index), (C.GLhalfNV)(x), (C.GLhalfNV)(y), (C.GLhalfNV)(z), (C.GLhalfNV)(w))
}
func VertexAttrib4hvNV(index Uint, v *Half)  {
	C.goglVertexAttrib4hvNV((C.GLuint)(index), (*C.GLhalfNV)(v))
}
func VertexAttribs1hvNV(index Uint, n Sizei, v *Half)  {
	C.goglVertexAttribs1hvNV((C.GLuint)(index), (C.GLsizei)(n), (*C.GLhalfNV)(v))
}
func VertexAttribs2hvNV(index Uint, n Sizei, v *Half)  {
	C.goglVertexAttribs2hvNV((C.GLuint)(index), (C.GLsizei)(n), (*C.GLhalfNV)(v))
}
func VertexAttribs3hvNV(index Uint, n Sizei, v *Half)  {
	C.goglVertexAttribs3hvNV((C.GLuint)(index), (C.GLsizei)(n), (*C.GLhalfNV)(v))
}
func VertexAttribs4hvNV(index Uint, n Sizei, v *Half)  {
	C.goglVertexAttribs4hvNV((C.GLuint)(index), (C.GLsizei)(n), (*C.GLhalfNV)(v))
}
// NV_occlusion_query

func GenOcclusionQueriesNV(n Sizei, ids *Uint)  {
	C.goglGenOcclusionQueriesNV((C.GLsizei)(n), (*C.GLuint)(ids))
}
func DeleteOcclusionQueriesNV(n Sizei, ids *Uint)  {
	C.goglDeleteOcclusionQueriesNV((C.GLsizei)(n), (*C.GLuint)(ids))
}
func IsOcclusionQueryNV(id Uint) Boolean {
	return (Boolean)(C.goglIsOcclusionQueryNV((C.GLuint)(id)))
}
func BeginOcclusionQueryNV(id Uint)  {
	C.goglBeginOcclusionQueryNV((C.GLuint)(id))
}
func EndOcclusionQueryNV()  {
	C.goglEndOcclusionQueryNV()
}
func GetOcclusionQueryivNV(id Uint, pname Enum, params *Int)  {
	C.goglGetOcclusionQueryivNV((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetOcclusionQueryuivNV(id Uint, pname Enum, params *Uint)  {
	C.goglGetOcclusionQueryuivNV((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(params))
}
// NV_parameter_buffer_object

func ProgramBufferParametersfvNV(target Enum, buffer Uint, index Uint, count Sizei, params *Float)  {
	C.goglProgramBufferParametersfvNV((C.GLenum)(target), (C.GLuint)(buffer), (C.GLuint)(index), (C.GLsizei)(count), (*C.GLfloat)(params))
}
func ProgramBufferParametersIivNV(target Enum, buffer Uint, index Uint, count Sizei, params *Int)  {
	C.goglProgramBufferParametersIivNV((C.GLenum)(target), (C.GLuint)(buffer), (C.GLuint)(index), (C.GLsizei)(count), (*C.GLint)(params))
}
func ProgramBufferParametersIuivNV(target Enum, buffer Uint, index Uint, count Sizei, params *Uint)  {
	C.goglProgramBufferParametersIuivNV((C.GLenum)(target), (C.GLuint)(buffer), (C.GLuint)(index), (C.GLsizei)(count), (*C.GLuint)(params))
}
// NV_path_rendering

func GenPathsNV(range_ Sizei) Uint {
	return (Uint)(C.goglGenPathsNV((C.GLsizei)(range_)))
}
func DeletePathsNV(path Uint, range_ Sizei)  {
	C.goglDeletePathsNV((C.GLuint)(path), (C.GLsizei)(range_))
}
func IsPathNV(path Uint) Boolean {
	return (Boolean)(C.goglIsPathNV((C.GLuint)(path)))
}
func PathCommandsNV(path Uint, numCommands Sizei, commands *Ubyte, numCoords Sizei, coordType Enum, coords Pointer)  {
	C.goglPathCommandsNV((C.GLuint)(path), (C.GLsizei)(numCommands), (*C.GLubyte)(commands), (C.GLsizei)(numCoords), (C.GLenum)(coordType), (unsafe.Pointer)(coords))
}
func PathCoordsNV(path Uint, numCoords Sizei, coordType Enum, coords Pointer)  {
	C.goglPathCoordsNV((C.GLuint)(path), (C.GLsizei)(numCoords), (C.GLenum)(coordType), (unsafe.Pointer)(coords))
}
func PathSubCommandsNV(path Uint, commandStart Sizei, commandsToDelete Sizei, numCommands Sizei, commands *Ubyte, numCoords Sizei, coordType Enum, coords Pointer)  {
	C.goglPathSubCommandsNV((C.GLuint)(path), (C.GLsizei)(commandStart), (C.GLsizei)(commandsToDelete), (C.GLsizei)(numCommands), (*C.GLubyte)(commands), (C.GLsizei)(numCoords), (C.GLenum)(coordType), (unsafe.Pointer)(coords))
}
func PathSubCoordsNV(path Uint, coordStart Sizei, numCoords Sizei, coordType Enum, coords Pointer)  {
	C.goglPathSubCoordsNV((C.GLuint)(path), (C.GLsizei)(coordStart), (C.GLsizei)(numCoords), (C.GLenum)(coordType), (unsafe.Pointer)(coords))
}
func PathStringNV(path Uint, format Enum, length Sizei, pathString Pointer)  {
	C.goglPathStringNV((C.GLuint)(path), (C.GLenum)(format), (C.GLsizei)(length), (unsafe.Pointer)(pathString))
}
func PathGlyphsNV(firstPathName Uint, fontTarget Enum, fontName Pointer, fontStyle Bitfield, numGlyphs Sizei, type_ Enum, charcodes Pointer, handleMissingGlyphs Enum, pathParameterTemplate Uint, emScale Float)  {
	C.goglPathGlyphsNV((C.GLuint)(firstPathName), (C.GLenum)(fontTarget), (unsafe.Pointer)(fontName), (C.GLbitfield)(fontStyle), (C.GLsizei)(numGlyphs), (C.GLenum)(type_), (unsafe.Pointer)(charcodes), (C.GLenum)(handleMissingGlyphs), (C.GLuint)(pathParameterTemplate), (C.GLfloat)(emScale))
}
func PathGlyphRangeNV(firstPathName Uint, fontTarget Enum, fontName Pointer, fontStyle Bitfield, firstGlyph Uint, numGlyphs Sizei, handleMissingGlyphs Enum, pathParameterTemplate Uint, emScale Float)  {
	C.goglPathGlyphRangeNV((C.GLuint)(firstPathName), (C.GLenum)(fontTarget), (unsafe.Pointer)(fontName), (C.GLbitfield)(fontStyle), (C.GLuint)(firstGlyph), (C.GLsizei)(numGlyphs), (C.GLenum)(handleMissingGlyphs), (C.GLuint)(pathParameterTemplate), (C.GLfloat)(emScale))
}
func WeightPathsNV(resultPath Uint, numPaths Sizei, paths *Uint, weights *Float)  {
	C.goglWeightPathsNV((C.GLuint)(resultPath), (C.GLsizei)(numPaths), (*C.GLuint)(paths), (*C.GLfloat)(weights))
}
func CopyPathNV(resultPath Uint, srcPath Uint)  {
	C.goglCopyPathNV((C.GLuint)(resultPath), (C.GLuint)(srcPath))
}
func InterpolatePathsNV(resultPath Uint, pathA Uint, pathB Uint, weight Float)  {
	C.goglInterpolatePathsNV((C.GLuint)(resultPath), (C.GLuint)(pathA), (C.GLuint)(pathB), (C.GLfloat)(weight))
}
func TransformPathNV(resultPath Uint, srcPath Uint, transformType Enum, transformValues *Float)  {
	C.goglTransformPathNV((C.GLuint)(resultPath), (C.GLuint)(srcPath), (C.GLenum)(transformType), (*C.GLfloat)(transformValues))
}
func PathParameterivNV(path Uint, pname Enum, value *Int)  {
	C.goglPathParameterivNV((C.GLuint)(path), (C.GLenum)(pname), (*C.GLint)(value))
}
func PathParameteriNV(path Uint, pname Enum, value Int)  {
	C.goglPathParameteriNV((C.GLuint)(path), (C.GLenum)(pname), (C.GLint)(value))
}
func PathParameterfvNV(path Uint, pname Enum, value *Float)  {
	C.goglPathParameterfvNV((C.GLuint)(path), (C.GLenum)(pname), (*C.GLfloat)(value))
}
func PathParameterfNV(path Uint, pname Enum, value Float)  {
	C.goglPathParameterfNV((C.GLuint)(path), (C.GLenum)(pname), (C.GLfloat)(value))
}
func PathDashArrayNV(path Uint, dashCount Sizei, dashArray *Float)  {
	C.goglPathDashArrayNV((C.GLuint)(path), (C.GLsizei)(dashCount), (*C.GLfloat)(dashArray))
}
func PathStencilFuncNV(func_ Enum, ref Int, mask Uint)  {
	C.goglPathStencilFuncNV((C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
func PathStencilDepthOffsetNV(factor Float, units Float)  {
	C.goglPathStencilDepthOffsetNV((C.GLfloat)(factor), (C.GLfloat)(units))
}
func StencilFillPathNV(path Uint, fillMode Enum, mask Uint)  {
	C.goglStencilFillPathNV((C.GLuint)(path), (C.GLenum)(fillMode), (C.GLuint)(mask))
}
func StencilStrokePathNV(path Uint, reference Int, mask Uint)  {
	C.goglStencilStrokePathNV((C.GLuint)(path), (C.GLint)(reference), (C.GLuint)(mask))
}
func StencilFillPathInstancedNV(numPaths Sizei, pathNameType Enum, paths Pointer, pathBase Uint, fillMode Enum, mask Uint, transformType Enum, transformValues *Float)  {
	C.goglStencilFillPathInstancedNV((C.GLsizei)(numPaths), (C.GLenum)(pathNameType), (unsafe.Pointer)(paths), (C.GLuint)(pathBase), (C.GLenum)(fillMode), (C.GLuint)(mask), (C.GLenum)(transformType), (*C.GLfloat)(transformValues))
}
func StencilStrokePathInstancedNV(numPaths Sizei, pathNameType Enum, paths Pointer, pathBase Uint, reference Int, mask Uint, transformType Enum, transformValues *Float)  {
	C.goglStencilStrokePathInstancedNV((C.GLsizei)(numPaths), (C.GLenum)(pathNameType), (unsafe.Pointer)(paths), (C.GLuint)(pathBase), (C.GLint)(reference), (C.GLuint)(mask), (C.GLenum)(transformType), (*C.GLfloat)(transformValues))
}
func PathCoverDepthFuncNV(func_ Enum)  {
	C.goglPathCoverDepthFuncNV((C.GLenum)(func_))
}
func PathColorGenNV(color Enum, genMode Enum, colorFormat Enum, coeffs *Float)  {
	C.goglPathColorGenNV((C.GLenum)(color), (C.GLenum)(genMode), (C.GLenum)(colorFormat), (*C.GLfloat)(coeffs))
}
func PathTexGenNV(texCoordSet Enum, genMode Enum, components Int, coeffs *Float)  {
	C.goglPathTexGenNV((C.GLenum)(texCoordSet), (C.GLenum)(genMode), (C.GLint)(components), (*C.GLfloat)(coeffs))
}
func PathFogGenNV(genMode Enum)  {
	C.goglPathFogGenNV((C.GLenum)(genMode))
}
func CoverFillPathNV(path Uint, coverMode Enum)  {
	C.goglCoverFillPathNV((C.GLuint)(path), (C.GLenum)(coverMode))
}
func CoverStrokePathNV(path Uint, coverMode Enum)  {
	C.goglCoverStrokePathNV((C.GLuint)(path), (C.GLenum)(coverMode))
}
func CoverFillPathInstancedNV(numPaths Sizei, pathNameType Enum, paths Pointer, pathBase Uint, coverMode Enum, transformType Enum, transformValues *Float)  {
	C.goglCoverFillPathInstancedNV((C.GLsizei)(numPaths), (C.GLenum)(pathNameType), (unsafe.Pointer)(paths), (C.GLuint)(pathBase), (C.GLenum)(coverMode), (C.GLenum)(transformType), (*C.GLfloat)(transformValues))
}
func CoverStrokePathInstancedNV(numPaths Sizei, pathNameType Enum, paths Pointer, pathBase Uint, coverMode Enum, transformType Enum, transformValues *Float)  {
	C.goglCoverStrokePathInstancedNV((C.GLsizei)(numPaths), (C.GLenum)(pathNameType), (unsafe.Pointer)(paths), (C.GLuint)(pathBase), (C.GLenum)(coverMode), (C.GLenum)(transformType), (*C.GLfloat)(transformValues))
}
func GetPathParameterivNV(path Uint, pname Enum, value *Int)  {
	C.goglGetPathParameterivNV((C.GLuint)(path), (C.GLenum)(pname), (*C.GLint)(value))
}
func GetPathParameterfvNV(path Uint, pname Enum, value *Float)  {
	C.goglGetPathParameterfvNV((C.GLuint)(path), (C.GLenum)(pname), (*C.GLfloat)(value))
}
func GetPathCommandsNV(path Uint, commands *Ubyte)  {
	C.goglGetPathCommandsNV((C.GLuint)(path), (*C.GLubyte)(commands))
}
func GetPathCoordsNV(path Uint, coords *Float)  {
	C.goglGetPathCoordsNV((C.GLuint)(path), (*C.GLfloat)(coords))
}
func GetPathDashArrayNV(path Uint, dashArray *Float)  {
	C.goglGetPathDashArrayNV((C.GLuint)(path), (*C.GLfloat)(dashArray))
}
func GetPathMetricsNV(metricQueryMask Bitfield, numPaths Sizei, pathNameType Enum, paths Pointer, pathBase Uint, stride Sizei, metrics *Float)  {
	C.goglGetPathMetricsNV((C.GLbitfield)(metricQueryMask), (C.GLsizei)(numPaths), (C.GLenum)(pathNameType), (unsafe.Pointer)(paths), (C.GLuint)(pathBase), (C.GLsizei)(stride), (*C.GLfloat)(metrics))
}
func GetPathMetricRangeNV(metricQueryMask Bitfield, firstPathName Uint, numPaths Sizei, stride Sizei, metrics *Float)  {
	C.goglGetPathMetricRangeNV((C.GLbitfield)(metricQueryMask), (C.GLuint)(firstPathName), (C.GLsizei)(numPaths), (C.GLsizei)(stride), (*C.GLfloat)(metrics))
}
func GetPathSpacingNV(pathListMode Enum, numPaths Sizei, pathNameType Enum, paths Pointer, pathBase Uint, advanceScale Float, kerningScale Float, transformType Enum, returnedSpacing *Float)  {
	C.goglGetPathSpacingNV((C.GLenum)(pathListMode), (C.GLsizei)(numPaths), (C.GLenum)(pathNameType), (unsafe.Pointer)(paths), (C.GLuint)(pathBase), (C.GLfloat)(advanceScale), (C.GLfloat)(kerningScale), (C.GLenum)(transformType), (*C.GLfloat)(returnedSpacing))
}
func GetPathColorGenivNV(color Enum, pname Enum, value *Int)  {
	C.goglGetPathColorGenivNV((C.GLenum)(color), (C.GLenum)(pname), (*C.GLint)(value))
}
func GetPathColorGenfvNV(color Enum, pname Enum, value *Float)  {
	C.goglGetPathColorGenfvNV((C.GLenum)(color), (C.GLenum)(pname), (*C.GLfloat)(value))
}
func GetPathTexGenivNV(texCoordSet Enum, pname Enum, value *Int)  {
	C.goglGetPathTexGenivNV((C.GLenum)(texCoordSet), (C.GLenum)(pname), (*C.GLint)(value))
}
func GetPathTexGenfvNV(texCoordSet Enum, pname Enum, value *Float)  {
	C.goglGetPathTexGenfvNV((C.GLenum)(texCoordSet), (C.GLenum)(pname), (*C.GLfloat)(value))
}
func IsPointInFillPathNV(path Uint, mask Uint, x Float, y Float) Boolean {
	return (Boolean)(C.goglIsPointInFillPathNV((C.GLuint)(path), (C.GLuint)(mask), (C.GLfloat)(x), (C.GLfloat)(y)))
}
func IsPointInStrokePathNV(path Uint, x Float, y Float) Boolean {
	return (Boolean)(C.goglIsPointInStrokePathNV((C.GLuint)(path), (C.GLfloat)(x), (C.GLfloat)(y)))
}
func GetPathLengthNV(path Uint, startSegment Sizei, numSegments Sizei) Float {
	return (Float)(C.goglGetPathLengthNV((C.GLuint)(path), (C.GLsizei)(startSegment), (C.GLsizei)(numSegments)))
}
func PointAlongPathNV(path Uint, startSegment Sizei, numSegments Sizei, distance Float, x *Float, y *Float, tangentX *Float, tangentY *Float) Boolean {
	return (Boolean)(C.goglPointAlongPathNV((C.GLuint)(path), (C.GLsizei)(startSegment), (C.GLsizei)(numSegments), (C.GLfloat)(distance), (*C.GLfloat)(x), (*C.GLfloat)(y), (*C.GLfloat)(tangentX), (*C.GLfloat)(tangentY)))
}
// NV_pixel_data_range

func PixelDataRangeNV(target Enum, length Sizei, pointer Pointer)  {
	C.goglPixelDataRangeNV((C.GLenum)(target), (C.GLsizei)(length), (unsafe.Pointer)(pointer))
}
func FlushPixelDataRangeNV(target Enum)  {
	C.goglFlushPixelDataRangeNV((C.GLenum)(target))
}
// NV_point_sprite

func PointParameteriNV(pname Enum, param Int)  {
	C.goglPointParameteriNV((C.GLenum)(pname), (C.GLint)(param))
}
func PointParameterivNV(pname Enum, params *Int)  {
	C.goglPointParameterivNV((C.GLenum)(pname), (*C.GLint)(params))
}
// NV_present_video

func PresentFrameKeyedNV(video_slot Uint, minPresentTime Uint64, beginPresentTimeId Uint, presentDurationId Uint, type_ Enum, target0 Enum, fill0 Uint, key0 Uint, target1 Enum, fill1 Uint, key1 Uint)  {
	C.goglPresentFrameKeyedNV((C.GLuint)(video_slot), (C.GLuint64EXT)(minPresentTime), (C.GLuint)(beginPresentTimeId), (C.GLuint)(presentDurationId), (C.GLenum)(type_), (C.GLenum)(target0), (C.GLuint)(fill0), (C.GLuint)(key0), (C.GLenum)(target1), (C.GLuint)(fill1), (C.GLuint)(key1))
}
func PresentFrameDualFillNV(video_slot Uint, minPresentTime Uint64, beginPresentTimeId Uint, presentDurationId Uint, type_ Enum, target0 Enum, fill0 Uint, target1 Enum, fill1 Uint, target2 Enum, fill2 Uint, target3 Enum, fill3 Uint)  {
	C.goglPresentFrameDualFillNV((C.GLuint)(video_slot), (C.GLuint64EXT)(minPresentTime), (C.GLuint)(beginPresentTimeId), (C.GLuint)(presentDurationId), (C.GLenum)(type_), (C.GLenum)(target0), (C.GLuint)(fill0), (C.GLenum)(target1), (C.GLuint)(fill1), (C.GLenum)(target2), (C.GLuint)(fill2), (C.GLenum)(target3), (C.GLuint)(fill3))
}
func GetVideoivNV(video_slot Uint, pname Enum, params *Int)  {
	C.goglGetVideoivNV((C.GLuint)(video_slot), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetVideouivNV(video_slot Uint, pname Enum, params *Uint)  {
	C.goglGetVideouivNV((C.GLuint)(video_slot), (C.GLenum)(pname), (*C.GLuint)(params))
}
func GetVideoi64vNV(video_slot Uint, pname Enum, params *Int64)  {
	C.goglGetVideoi64vNV((C.GLuint)(video_slot), (C.GLenum)(pname), (*C.GLint64EXT)(params))
}
func GetVideoui64vNV(video_slot Uint, pname Enum, params *Uint64)  {
	C.goglGetVideoui64vNV((C.GLuint)(video_slot), (C.GLenum)(pname), (*C.GLuint64EXT)(params))
}
// NV_primitive_restart

func PrimitiveRestartNV()  {
	C.goglPrimitiveRestartNV()
}
func PrimitiveRestartIndexNV(index Uint)  {
	C.goglPrimitiveRestartIndexNV((C.GLuint)(index))
}
// NV_register_combiners

func CombinerParameterfvNV(pname Enum, params *Float)  {
	C.goglCombinerParameterfvNV((C.GLenum)(pname), (*C.GLfloat)(params))
}
func CombinerParameterfNV(pname Enum, param Float)  {
	C.goglCombinerParameterfNV((C.GLenum)(pname), (C.GLfloat)(param))
}
func CombinerParameterivNV(pname Enum, params *Int)  {
	C.goglCombinerParameterivNV((C.GLenum)(pname), (*C.GLint)(params))
}
func CombinerParameteriNV(pname Enum, param Int)  {
	C.goglCombinerParameteriNV((C.GLenum)(pname), (C.GLint)(param))
}
func CombinerInputNV(stage Enum, portion Enum, variable Enum, input Enum, mapping Enum, componentUsage Enum)  {
	C.goglCombinerInputNV((C.GLenum)(stage), (C.GLenum)(portion), (C.GLenum)(variable), (C.GLenum)(input), (C.GLenum)(mapping), (C.GLenum)(componentUsage))
}
func CombinerOutputNV(stage Enum, portion Enum, abOutput Enum, cdOutput Enum, sumOutput Enum, scale Enum, bias Enum, abDotProduct Boolean, cdDotProduct Boolean, muxSum Boolean)  {
	C.goglCombinerOutputNV((C.GLenum)(stage), (C.GLenum)(portion), (C.GLenum)(abOutput), (C.GLenum)(cdOutput), (C.GLenum)(sumOutput), (C.GLenum)(scale), (C.GLenum)(bias), (C.GLboolean)(abDotProduct), (C.GLboolean)(cdDotProduct), (C.GLboolean)(muxSum))
}
func FinalCombinerInputNV(variable Enum, input Enum, mapping Enum, componentUsage Enum)  {
	C.goglFinalCombinerInputNV((C.GLenum)(variable), (C.GLenum)(input), (C.GLenum)(mapping), (C.GLenum)(componentUsage))
}
func GetCombinerInputParameterfvNV(stage Enum, portion Enum, variable Enum, pname Enum, params *Float)  {
	C.goglGetCombinerInputParameterfvNV((C.GLenum)(stage), (C.GLenum)(portion), (C.GLenum)(variable), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetCombinerInputParameterivNV(stage Enum, portion Enum, variable Enum, pname Enum, params *Int)  {
	C.goglGetCombinerInputParameterivNV((C.GLenum)(stage), (C.GLenum)(portion), (C.GLenum)(variable), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetCombinerOutputParameterfvNV(stage Enum, portion Enum, pname Enum, params *Float)  {
	C.goglGetCombinerOutputParameterfvNV((C.GLenum)(stage), (C.GLenum)(portion), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetCombinerOutputParameterivNV(stage Enum, portion Enum, pname Enum, params *Int)  {
	C.goglGetCombinerOutputParameterivNV((C.GLenum)(stage), (C.GLenum)(portion), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetFinalCombinerInputParameterfvNV(variable Enum, pname Enum, params *Float)  {
	C.goglGetFinalCombinerInputParameterfvNV((C.GLenum)(variable), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetFinalCombinerInputParameterivNV(variable Enum, pname Enum, params *Int)  {
	C.goglGetFinalCombinerInputParameterivNV((C.GLenum)(variable), (C.GLenum)(pname), (*C.GLint)(params))
}
// NV_register_combiners2

func CombinerStageParameterfvNV(stage Enum, pname Enum, params *Float)  {
	C.goglCombinerStageParameterfvNV((C.GLenum)(stage), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetCombinerStageParameterfvNV(stage Enum, pname Enum, params *Float)  {
	C.goglGetCombinerStageParameterfvNV((C.GLenum)(stage), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// NV_shader_buffer_load

func MakeBufferResidentNV(target Enum, access Enum)  {
	C.goglMakeBufferResidentNV((C.GLenum)(target), (C.GLenum)(access))
}
func MakeBufferNonResidentNV(target Enum)  {
	C.goglMakeBufferNonResidentNV((C.GLenum)(target))
}
func IsBufferResidentNV(target Enum) Boolean {
	return (Boolean)(C.goglIsBufferResidentNV((C.GLenum)(target)))
}
func MakeNamedBufferResidentNV(buffer Uint, access Enum)  {
	C.goglMakeNamedBufferResidentNV((C.GLuint)(buffer), (C.GLenum)(access))
}
func MakeNamedBufferNonResidentNV(buffer Uint)  {
	C.goglMakeNamedBufferNonResidentNV((C.GLuint)(buffer))
}
func IsNamedBufferResidentNV(buffer Uint) Boolean {
	return (Boolean)(C.goglIsNamedBufferResidentNV((C.GLuint)(buffer)))
}
func GetBufferParameterui64vNV(target Enum, pname Enum, params *Uint64)  {
	C.goglGetBufferParameterui64vNV((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint64EXT)(params))
}
func GetNamedBufferParameterui64vNV(buffer Uint, pname Enum, params *Uint64)  {
	C.goglGetNamedBufferParameterui64vNV((C.GLuint)(buffer), (C.GLenum)(pname), (*C.GLuint64EXT)(params))
}
func GetIntegerui64vNV(value Enum, result *Uint64)  {
	C.goglGetIntegerui64vNV((C.GLenum)(value), (*C.GLuint64EXT)(result))
}
func Uniformui64NV(location Int, value Uint64)  {
	C.goglUniformui64NV((C.GLint)(location), (C.GLuint64EXT)(value))
}
func Uniformui64vNV(location Int, count Sizei, value *Uint64)  {
	C.goglUniformui64vNV((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(value))
}
func GetUniformui64vNV(program Uint, location Int, params *Uint64)  {
	C.goglGetUniformui64vNV((C.GLuint)(program), (C.GLint)(location), (*C.GLuint64EXT)(params))
}
func ProgramUniformui64NV(program Uint, location Int, value Uint64)  {
	C.goglProgramUniformui64NV((C.GLuint)(program), (C.GLint)(location), (C.GLuint64EXT)(value))
}
func ProgramUniformui64vNV(program Uint, location Int, count Sizei, value *Uint64)  {
	C.goglProgramUniformui64vNV((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64EXT)(value))
}
// NV_texture_barrier

func TextureBarrierNV()  {
	C.goglTextureBarrierNV()
}
// NV_texture_multisample

func TexImage2DMultisampleCoverageNV(target Enum, coverageSamples Sizei, colorSamples Sizei, internalFormat Int, width Sizei, height Sizei, fixedSampleLocations Boolean)  {
	C.goglTexImage2DMultisampleCoverageNV((C.GLenum)(target), (C.GLsizei)(coverageSamples), (C.GLsizei)(colorSamples), (C.GLint)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(fixedSampleLocations))
}
func TexImage3DMultisampleCoverageNV(target Enum, coverageSamples Sizei, colorSamples Sizei, internalFormat Int, width Sizei, height Sizei, depth Sizei, fixedSampleLocations Boolean)  {
	C.goglTexImage3DMultisampleCoverageNV((C.GLenum)(target), (C.GLsizei)(coverageSamples), (C.GLsizei)(colorSamples), (C.GLint)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(fixedSampleLocations))
}
func TextureImage2DMultisampleNV(texture Uint, target Enum, samples Sizei, internalFormat Int, width Sizei, height Sizei, fixedSampleLocations Boolean)  {
	C.goglTextureImage2DMultisampleNV((C.GLuint)(texture), (C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(fixedSampleLocations))
}
func TextureImage3DMultisampleNV(texture Uint, target Enum, samples Sizei, internalFormat Int, width Sizei, height Sizei, depth Sizei, fixedSampleLocations Boolean)  {
	C.goglTextureImage3DMultisampleNV((C.GLuint)(texture), (C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(fixedSampleLocations))
}
func TextureImage2DMultisampleCoverageNV(texture Uint, target Enum, coverageSamples Sizei, colorSamples Sizei, internalFormat Int, width Sizei, height Sizei, fixedSampleLocations Boolean)  {
	C.goglTextureImage2DMultisampleCoverageNV((C.GLuint)(texture), (C.GLenum)(target), (C.GLsizei)(coverageSamples), (C.GLsizei)(colorSamples), (C.GLint)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(fixedSampleLocations))
}
func TextureImage3DMultisampleCoverageNV(texture Uint, target Enum, coverageSamples Sizei, colorSamples Sizei, internalFormat Int, width Sizei, height Sizei, depth Sizei, fixedSampleLocations Boolean)  {
	C.goglTextureImage3DMultisampleCoverageNV((C.GLuint)(texture), (C.GLenum)(target), (C.GLsizei)(coverageSamples), (C.GLsizei)(colorSamples), (C.GLint)(internalFormat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(fixedSampleLocations))
}
// NV_transform_feedback

func BeginTransformFeedbackNV(primitiveMode Enum)  {
	C.goglBeginTransformFeedbackNV((C.GLenum)(primitiveMode))
}
func EndTransformFeedbackNV()  {
	C.goglEndTransformFeedbackNV()
}
func TransformFeedbackAttribsNV(count Sizei, attribs *Int, bufferMode Enum)  {
	C.goglTransformFeedbackAttribsNV((C.GLsizei)(count), (*C.GLint)(attribs), (C.GLenum)(bufferMode))
}
func BindBufferRangeNV(target Enum, index Uint, buffer Uint, offset Intptr, size Sizeiptr)  {
	C.goglBindBufferRangeNV((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
func BindBufferOffsetNV(target Enum, index Uint, buffer Uint, offset Intptr)  {
	C.goglBindBufferOffsetNV((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset))
}
func BindBufferBaseNV(target Enum, index Uint, buffer Uint)  {
	C.goglBindBufferBaseNV((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer))
}
func TransformFeedbackVaryingsNV(program Uint, count Sizei, locations *Int, bufferMode Enum)  {
	C.goglTransformFeedbackVaryingsNV((C.GLuint)(program), (C.GLsizei)(count), (*C.GLint)(locations), (C.GLenum)(bufferMode))
}
func ActiveVaryingNV(program Uint, name *Char)  {
	C.goglActiveVaryingNV((C.GLuint)(program), (*C.GLchar)(name))
}
func GetVaryingLocationNV(program Uint, name *Char) Int {
	return (Int)(C.goglGetVaryingLocationNV((C.GLuint)(program), (*C.GLchar)(name)))
}
func GetActiveVaryingNV(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Sizei, type_ *Enum, name *Char)  {
	C.goglGetActiveVaryingNV((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLsizei)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}
func GetTransformFeedbackVaryingNV(program Uint, index Uint, location *Int)  {
	C.goglGetTransformFeedbackVaryingNV((C.GLuint)(program), (C.GLuint)(index), (*C.GLint)(location))
}
func TransformFeedbackStreamAttribsNV(count Sizei, attribs *Int, nbuffers Sizei, bufstreams *Int, bufferMode Enum)  {
	C.goglTransformFeedbackStreamAttribsNV((C.GLsizei)(count), (*C.GLint)(attribs), (C.GLsizei)(nbuffers), (*C.GLint)(bufstreams), (C.GLenum)(bufferMode))
}
// NV_transform_feedback2

func BindTransformFeedbackNV(target Enum, id Uint)  {
	C.goglBindTransformFeedbackNV((C.GLenum)(target), (C.GLuint)(id))
}
func DeleteTransformFeedbacksNV(n Sizei, ids *Uint)  {
	C.goglDeleteTransformFeedbacksNV((C.GLsizei)(n), (*C.GLuint)(ids))
}
func GenTransformFeedbacksNV(n Sizei, ids *Uint)  {
	C.goglGenTransformFeedbacksNV((C.GLsizei)(n), (*C.GLuint)(ids))
}
func IsTransformFeedbackNV(id Uint) Boolean {
	return (Boolean)(C.goglIsTransformFeedbackNV((C.GLuint)(id)))
}
func PauseTransformFeedbackNV()  {
	C.goglPauseTransformFeedbackNV()
}
func ResumeTransformFeedbackNV()  {
	C.goglResumeTransformFeedbackNV()
}
func DrawTransformFeedbackNV(mode Enum, id Uint)  {
	C.goglDrawTransformFeedbackNV((C.GLenum)(mode), (C.GLuint)(id))
}
// NV_vdpau_interop

func VDPAUInitNV(vdpDevice Pointer, getProcAddress Pointer)  {
	C.goglVDPAUInitNV((unsafe.Pointer)(vdpDevice), (unsafe.Pointer)(getProcAddress))
}
func VDPAUFiniNV()  {
	C.goglVDPAUFiniNV()
}
func VDPAURegisterVideoSurfaceNV(vdpSurface Pointer, target Enum, numTextureNames Sizei, textureNames *Uint) Intptr {
	return (Intptr)(C.goglVDPAURegisterVideoSurfaceNV((unsafe.Pointer)(vdpSurface), (C.GLenum)(target), (C.GLsizei)(numTextureNames), (*C.GLuint)(textureNames)))
}
func VDPAURegisterOutputSurfaceNV(vdpSurface Pointer, target Enum, numTextureNames Sizei, textureNames *Uint) Intptr {
	return (Intptr)(C.goglVDPAURegisterOutputSurfaceNV((unsafe.Pointer)(vdpSurface), (C.GLenum)(target), (C.GLsizei)(numTextureNames), (*C.GLuint)(textureNames)))
}
func VDPAUIsSurfaceNV(surface Intptr)  {
	C.goglVDPAUIsSurfaceNV((C.GLvdpauSurfaceNV)(surface))
}
func VDPAUUnregisterSurfaceNV(surface Intptr)  {
	C.goglVDPAUUnregisterSurfaceNV((C.GLvdpauSurfaceNV)(surface))
}
func VDPAUGetSurfaceivNV(surface Intptr, pname Enum, bufSize Sizei, length *Sizei, values *Int)  {
	C.goglVDPAUGetSurfaceivNV((C.GLvdpauSurfaceNV)(surface), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(values))
}
func VDPAUSurfaceAccessNV(surface Intptr, access Enum)  {
	C.goglVDPAUSurfaceAccessNV((C.GLvdpauSurfaceNV)(surface), (C.GLenum)(access))
}
func VDPAUMapSurfacesNV(numSurfaces Sizei, surfaces *Intptr)  {
	C.goglVDPAUMapSurfacesNV((C.GLsizei)(numSurfaces), (*C.GLvdpauSurfaceNV)(surfaces))
}
func VDPAUUnmapSurfacesNV(numSurface Sizei, surfaces *Intptr)  {
	C.goglVDPAUUnmapSurfacesNV((C.GLsizei)(numSurface), (*C.GLvdpauSurfaceNV)(surfaces))
}
// NV_vertex_array_range

func FlushVertexArrayRangeNV()  {
	C.goglFlushVertexArrayRangeNV()
}
func VertexArrayRangeNV(length Sizei, pointer Pointer)  {
	C.goglVertexArrayRangeNV((C.GLsizei)(length), (unsafe.Pointer)(pointer))
}
// NV_vertex_attrib_integer_64bit

func VertexAttribL1i64NV(index Uint, x Int64)  {
	C.goglVertexAttribL1i64NV((C.GLuint)(index), (C.GLint64EXT)(x))
}
func VertexAttribL2i64NV(index Uint, x Int64, y Int64)  {
	C.goglVertexAttribL2i64NV((C.GLuint)(index), (C.GLint64EXT)(x), (C.GLint64EXT)(y))
}
func VertexAttribL3i64NV(index Uint, x Int64, y Int64, z Int64)  {
	C.goglVertexAttribL3i64NV((C.GLuint)(index), (C.GLint64EXT)(x), (C.GLint64EXT)(y), (C.GLint64EXT)(z))
}
func VertexAttribL4i64NV(index Uint, x Int64, y Int64, z Int64, w Int64)  {
	C.goglVertexAttribL4i64NV((C.GLuint)(index), (C.GLint64EXT)(x), (C.GLint64EXT)(y), (C.GLint64EXT)(z), (C.GLint64EXT)(w))
}
func VertexAttribL1i64vNV(index Uint, v *Int64)  {
	C.goglVertexAttribL1i64vNV((C.GLuint)(index), (*C.GLint64EXT)(v))
}
func VertexAttribL2i64vNV(index Uint, v *Int64)  {
	C.goglVertexAttribL2i64vNV((C.GLuint)(index), (*C.GLint64EXT)(v))
}
func VertexAttribL3i64vNV(index Uint, v *Int64)  {
	C.goglVertexAttribL3i64vNV((C.GLuint)(index), (*C.GLint64EXT)(v))
}
func VertexAttribL4i64vNV(index Uint, v *Int64)  {
	C.goglVertexAttribL4i64vNV((C.GLuint)(index), (*C.GLint64EXT)(v))
}
func VertexAttribL1ui64NV(index Uint, x Uint64)  {
	C.goglVertexAttribL1ui64NV((C.GLuint)(index), (C.GLuint64EXT)(x))
}
func VertexAttribL2ui64NV(index Uint, x Uint64, y Uint64)  {
	C.goglVertexAttribL2ui64NV((C.GLuint)(index), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y))
}
func VertexAttribL3ui64NV(index Uint, x Uint64, y Uint64, z Uint64)  {
	C.goglVertexAttribL3ui64NV((C.GLuint)(index), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y), (C.GLuint64EXT)(z))
}
func VertexAttribL4ui64NV(index Uint, x Uint64, y Uint64, z Uint64, w Uint64)  {
	C.goglVertexAttribL4ui64NV((C.GLuint)(index), (C.GLuint64EXT)(x), (C.GLuint64EXT)(y), (C.GLuint64EXT)(z), (C.GLuint64EXT)(w))
}
func VertexAttribL1ui64vNV(index Uint, v *Uint64)  {
	C.goglVertexAttribL1ui64vNV((C.GLuint)(index), (*C.GLuint64EXT)(v))
}
func VertexAttribL2ui64vNV(index Uint, v *Uint64)  {
	C.goglVertexAttribL2ui64vNV((C.GLuint)(index), (*C.GLuint64EXT)(v))
}
func VertexAttribL3ui64vNV(index Uint, v *Uint64)  {
	C.goglVertexAttribL3ui64vNV((C.GLuint)(index), (*C.GLuint64EXT)(v))
}
func VertexAttribL4ui64vNV(index Uint, v *Uint64)  {
	C.goglVertexAttribL4ui64vNV((C.GLuint)(index), (*C.GLuint64EXT)(v))
}
func GetVertexAttribLi64vNV(index Uint, pname Enum, params *Int64)  {
	C.goglGetVertexAttribLi64vNV((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint64EXT)(params))
}
func GetVertexAttribLui64vNV(index Uint, pname Enum, params *Uint64)  {
	C.goglGetVertexAttribLui64vNV((C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint64EXT)(params))
}
func VertexAttribLFormatNV(index Uint, size Int, type_ Enum, stride Sizei)  {
	C.goglVertexAttribLFormatNV((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride))
}
// NV_vertex_buffer_unified_memory

func BufferAddressRangeNV(pname Enum, index Uint, address Uint64, length Sizeiptr)  {
	C.goglBufferAddressRangeNV((C.GLenum)(pname), (C.GLuint)(index), (C.GLuint64EXT)(address), (C.GLsizeiptr)(length))
}
func VertexFormatNV(size Int, type_ Enum, stride Sizei)  {
	C.goglVertexFormatNV((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride))
}
func NormalFormatNV(type_ Enum, stride Sizei)  {
	C.goglNormalFormatNV((C.GLenum)(type_), (C.GLsizei)(stride))
}
func ColorFormatNV(size Int, type_ Enum, stride Sizei)  {
	C.goglColorFormatNV((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride))
}
func IndexFormatNV(type_ Enum, stride Sizei)  {
	C.goglIndexFormatNV((C.GLenum)(type_), (C.GLsizei)(stride))
}
func TexCoordFormatNV(size Int, type_ Enum, stride Sizei)  {
	C.goglTexCoordFormatNV((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride))
}
func EdgeFlagFormatNV(stride Sizei)  {
	C.goglEdgeFlagFormatNV((C.GLsizei)(stride))
}
func SecondaryColorFormatNV(size Int, type_ Enum, stride Sizei)  {
	C.goglSecondaryColorFormatNV((C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride))
}
func FogCoordFormatNV(type_ Enum, stride Sizei)  {
	C.goglFogCoordFormatNV((C.GLenum)(type_), (C.GLsizei)(stride))
}
func VertexAttribFormatNV(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei)  {
	C.goglVertexAttribFormatNV((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLsizei)(stride))
}
func VertexAttribIFormatNV(index Uint, size Int, type_ Enum, stride Sizei)  {
	C.goglVertexAttribIFormatNV((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride))
}
func GetIntegerui64i_vNV(value Enum, index Uint, result *Uint64)  {
	C.goglGetIntegerui64i_vNV((C.GLenum)(value), (C.GLuint)(index), (*C.GLuint64EXT)(result))
}
// NV_vertex_program

func AreProgramsResidentNV(n Sizei, programs *Uint, residences *Boolean) Boolean {
	return (Boolean)(C.goglAreProgramsResidentNV((C.GLsizei)(n), (*C.GLuint)(programs), (*C.GLboolean)(residences)))
}
func BindProgramNV(target Enum, id Uint)  {
	C.goglBindProgramNV((C.GLenum)(target), (C.GLuint)(id))
}
func DeleteProgramsNV(n Sizei, programs *Uint)  {
	C.goglDeleteProgramsNV((C.GLsizei)(n), (*C.GLuint)(programs))
}
func ExecuteProgramNV(target Enum, id Uint, params *Float)  {
	C.goglExecuteProgramNV((C.GLenum)(target), (C.GLuint)(id), (*C.GLfloat)(params))
}
func GenProgramsNV(n Sizei, programs *Uint)  {
	C.goglGenProgramsNV((C.GLsizei)(n), (*C.GLuint)(programs))
}
func GetProgramParameterdvNV(target Enum, index Uint, pname Enum, params *Double)  {
	C.goglGetProgramParameterdvNV((C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}
func GetProgramParameterfvNV(target Enum, index Uint, pname Enum, params *Float)  {
	C.goglGetProgramParameterfvNV((C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetProgramivNV(id Uint, pname Enum, params *Int)  {
	C.goglGetProgramivNV((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetProgramStringNV(id Uint, pname Enum, program *Ubyte)  {
	C.goglGetProgramStringNV((C.GLuint)(id), (C.GLenum)(pname), (*C.GLubyte)(program))
}
func GetTrackMatrixivNV(target Enum, address Uint, pname Enum, params *Int)  {
	C.goglGetTrackMatrixivNV((C.GLenum)(target), (C.GLuint)(address), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetVertexAttribdvNV(index Uint, pname Enum, params *Double)  {
	C.goglGetVertexAttribdvNV((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}
func GetVertexAttribfvNV(index Uint, pname Enum, params *Float)  {
	C.goglGetVertexAttribfvNV((C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetVertexAttribivNV(index Uint, pname Enum, params *Int)  {
	C.goglGetVertexAttribivNV((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetVertexAttribPointervNV(index Uint, pname Enum, pointer *Pointer)  {
	C.goglGetVertexAttribPointervNV((C.GLuint)(index), (C.GLenum)(pname), (*unsafe.Pointer)(pointer))
}
func IsProgramNV(id Uint) Boolean {
	return (Boolean)(C.goglIsProgramNV((C.GLuint)(id)))
}
func LoadProgramNV(target Enum, id Uint, len Sizei, program *Ubyte)  {
	C.goglLoadProgramNV((C.GLenum)(target), (C.GLuint)(id), (C.GLsizei)(len), (*C.GLubyte)(program))
}
func ProgramParameter4dNV(target Enum, index Uint, x Double, y Double, z Double, w Double)  {
	C.goglProgramParameter4dNV((C.GLenum)(target), (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func ProgramParameter4dvNV(target Enum, index Uint, v *Double)  {
	C.goglProgramParameter4dvNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLdouble)(v))
}
func ProgramParameter4fNV(target Enum, index Uint, x Float, y Float, z Float, w Float)  {
	C.goglProgramParameter4fNV((C.GLenum)(target), (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func ProgramParameter4fvNV(target Enum, index Uint, v *Float)  {
	C.goglProgramParameter4fvNV((C.GLenum)(target), (C.GLuint)(index), (*C.GLfloat)(v))
}
func ProgramParameters4dvNV(target Enum, index Uint, count Sizei, v *Double)  {
	C.goglProgramParameters4dvNV((C.GLenum)(target), (C.GLuint)(index), (C.GLsizei)(count), (*C.GLdouble)(v))
}
func ProgramParameters4fvNV(target Enum, index Uint, count Sizei, v *Float)  {
	C.goglProgramParameters4fvNV((C.GLenum)(target), (C.GLuint)(index), (C.GLsizei)(count), (*C.GLfloat)(v))
}
func RequestResidentProgramsNV(n Sizei, programs *Uint)  {
	C.goglRequestResidentProgramsNV((C.GLsizei)(n), (*C.GLuint)(programs))
}
func TrackMatrixNV(target Enum, address Uint, matrix Enum, transform Enum)  {
	C.goglTrackMatrixNV((C.GLenum)(target), (C.GLuint)(address), (C.GLenum)(matrix), (C.GLenum)(transform))
}
func VertexAttribPointerNV(index Uint, fsize Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglVertexAttribPointerNV((C.GLuint)(index), (C.GLint)(fsize), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func VertexAttrib1dNV(index Uint, x Double)  {
	C.goglVertexAttrib1dNV((C.GLuint)(index), (C.GLdouble)(x))
}
func VertexAttrib1dvNV(index Uint, v *Double)  {
	C.goglVertexAttrib1dvNV((C.GLuint)(index), (*C.GLdouble)(v))
}
func VertexAttrib1fNV(index Uint, x Float)  {
	C.goglVertexAttrib1fNV((C.GLuint)(index), (C.GLfloat)(x))
}
func VertexAttrib1fvNV(index Uint, v *Float)  {
	C.goglVertexAttrib1fvNV((C.GLuint)(index), (*C.GLfloat)(v))
}
func VertexAttrib1sNV(index Uint, x Short)  {
	C.goglVertexAttrib1sNV((C.GLuint)(index), (C.GLshort)(x))
}
func VertexAttrib1svNV(index Uint, v *Short)  {
	C.goglVertexAttrib1svNV((C.GLuint)(index), (*C.GLshort)(v))
}
func VertexAttrib2dNV(index Uint, x Double, y Double)  {
	C.goglVertexAttrib2dNV((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
func VertexAttrib2dvNV(index Uint, v *Double)  {
	C.goglVertexAttrib2dvNV((C.GLuint)(index), (*C.GLdouble)(v))
}
func VertexAttrib2fNV(index Uint, x Float, y Float)  {
	C.goglVertexAttrib2fNV((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y))
}
func VertexAttrib2fvNV(index Uint, v *Float)  {
	C.goglVertexAttrib2fvNV((C.GLuint)(index), (*C.GLfloat)(v))
}
func VertexAttrib2sNV(index Uint, x Short, y Short)  {
	C.goglVertexAttrib2sNV((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y))
}
func VertexAttrib2svNV(index Uint, v *Short)  {
	C.goglVertexAttrib2svNV((C.GLuint)(index), (*C.GLshort)(v))
}
func VertexAttrib3dNV(index Uint, x Double, y Double, z Double)  {
	C.goglVertexAttrib3dNV((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func VertexAttrib3dvNV(index Uint, v *Double)  {
	C.goglVertexAttrib3dvNV((C.GLuint)(index), (*C.GLdouble)(v))
}
func VertexAttrib3fNV(index Uint, x Float, y Float, z Float)  {
	C.goglVertexAttrib3fNV((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func VertexAttrib3fvNV(index Uint, v *Float)  {
	C.goglVertexAttrib3fvNV((C.GLuint)(index), (*C.GLfloat)(v))
}
func VertexAttrib3sNV(index Uint, x Short, y Short, z Short)  {
	C.goglVertexAttrib3sNV((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func VertexAttrib3svNV(index Uint, v *Short)  {
	C.goglVertexAttrib3svNV((C.GLuint)(index), (*C.GLshort)(v))
}
func VertexAttrib4dNV(index Uint, x Double, y Double, z Double, w Double)  {
	C.goglVertexAttrib4dNV((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func VertexAttrib4dvNV(index Uint, v *Double)  {
	C.goglVertexAttrib4dvNV((C.GLuint)(index), (*C.GLdouble)(v))
}
func VertexAttrib4fNV(index Uint, x Float, y Float, z Float, w Float)  {
	C.goglVertexAttrib4fNV((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func VertexAttrib4fvNV(index Uint, v *Float)  {
	C.goglVertexAttrib4fvNV((C.GLuint)(index), (*C.GLfloat)(v))
}
func VertexAttrib4sNV(index Uint, x Short, y Short, z Short, w Short)  {
	C.goglVertexAttrib4sNV((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func VertexAttrib4svNV(index Uint, v *Short)  {
	C.goglVertexAttrib4svNV((C.GLuint)(index), (*C.GLshort)(v))
}
func VertexAttrib4ubNV(index Uint, x Ubyte, y Ubyte, z Ubyte, w Ubyte)  {
	C.goglVertexAttrib4ubNV((C.GLuint)(index), (C.GLubyte)(x), (C.GLubyte)(y), (C.GLubyte)(z), (C.GLubyte)(w))
}
func VertexAttrib4ubvNV(index Uint, v *Ubyte)  {
	C.goglVertexAttrib4ubvNV((C.GLuint)(index), (*C.GLubyte)(v))
}
func VertexAttribs1dvNV(index Uint, count Sizei, v *Double)  {
	C.goglVertexAttribs1dvNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLdouble)(v))
}
func VertexAttribs1fvNV(index Uint, count Sizei, v *Float)  {
	C.goglVertexAttribs1fvNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLfloat)(v))
}
func VertexAttribs1svNV(index Uint, count Sizei, v *Short)  {
	C.goglVertexAttribs1svNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLshort)(v))
}
func VertexAttribs2dvNV(index Uint, count Sizei, v *Double)  {
	C.goglVertexAttribs2dvNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLdouble)(v))
}
func VertexAttribs2fvNV(index Uint, count Sizei, v *Float)  {
	C.goglVertexAttribs2fvNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLfloat)(v))
}
func VertexAttribs2svNV(index Uint, count Sizei, v *Short)  {
	C.goglVertexAttribs2svNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLshort)(v))
}
func VertexAttribs3dvNV(index Uint, count Sizei, v *Double)  {
	C.goglVertexAttribs3dvNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLdouble)(v))
}
func VertexAttribs3fvNV(index Uint, count Sizei, v *Float)  {
	C.goglVertexAttribs3fvNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLfloat)(v))
}
func VertexAttribs3svNV(index Uint, count Sizei, v *Short)  {
	C.goglVertexAttribs3svNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLshort)(v))
}
func VertexAttribs4dvNV(index Uint, count Sizei, v *Double)  {
	C.goglVertexAttribs4dvNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLdouble)(v))
}
func VertexAttribs4fvNV(index Uint, count Sizei, v *Float)  {
	C.goglVertexAttribs4fvNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLfloat)(v))
}
func VertexAttribs4svNV(index Uint, count Sizei, v *Short)  {
	C.goglVertexAttribs4svNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLshort)(v))
}
func VertexAttribs4ubvNV(index Uint, count Sizei, v *Ubyte)  {
	C.goglVertexAttribs4ubvNV((C.GLuint)(index), (C.GLsizei)(count), (*C.GLubyte)(v))
}
// NV_vertex_program4

func VertexAttribI1iEXT(index Uint, x Int)  {
	C.goglVertexAttribI1iEXT((C.GLuint)(index), (C.GLint)(x))
}
func VertexAttribI2iEXT(index Uint, x Int, y Int)  {
	C.goglVertexAttribI2iEXT((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y))
}
func VertexAttribI3iEXT(index Uint, x Int, y Int, z Int)  {
	C.goglVertexAttribI3iEXT((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func VertexAttribI4iEXT(index Uint, x Int, y Int, z Int, w Int)  {
	C.goglVertexAttribI4iEXT((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func VertexAttribI1uiEXT(index Uint, x Uint)  {
	C.goglVertexAttribI1uiEXT((C.GLuint)(index), (C.GLuint)(x))
}
func VertexAttribI2uiEXT(index Uint, x Uint, y Uint)  {
	C.goglVertexAttribI2uiEXT((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y))
}
func VertexAttribI3uiEXT(index Uint, x Uint, y Uint, z Uint)  {
	C.goglVertexAttribI3uiEXT((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z))
}
func VertexAttribI4uiEXT(index Uint, x Uint, y Uint, z Uint, w Uint)  {
	C.goglVertexAttribI4uiEXT((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}
func VertexAttribI1ivEXT(index Uint, v *Int)  {
	C.goglVertexAttribI1ivEXT((C.GLuint)(index), (*C.GLint)(v))
}
func VertexAttribI2ivEXT(index Uint, v *Int)  {
	C.goglVertexAttribI2ivEXT((C.GLuint)(index), (*C.GLint)(v))
}
func VertexAttribI3ivEXT(index Uint, v *Int)  {
	C.goglVertexAttribI3ivEXT((C.GLuint)(index), (*C.GLint)(v))
}
func VertexAttribI4ivEXT(index Uint, v *Int)  {
	C.goglVertexAttribI4ivEXT((C.GLuint)(index), (*C.GLint)(v))
}
func VertexAttribI1uivEXT(index Uint, v *Uint)  {
	C.goglVertexAttribI1uivEXT((C.GLuint)(index), (*C.GLuint)(v))
}
func VertexAttribI2uivEXT(index Uint, v *Uint)  {
	C.goglVertexAttribI2uivEXT((C.GLuint)(index), (*C.GLuint)(v))
}
func VertexAttribI3uivEXT(index Uint, v *Uint)  {
	C.goglVertexAttribI3uivEXT((C.GLuint)(index), (*C.GLuint)(v))
}
func VertexAttribI4uivEXT(index Uint, v *Uint)  {
	C.goglVertexAttribI4uivEXT((C.GLuint)(index), (*C.GLuint)(v))
}
func VertexAttribI4bvEXT(index Uint, v *Byte)  {
	C.goglVertexAttribI4bvEXT((C.GLuint)(index), (*C.GLbyte)(v))
}
func VertexAttribI4svEXT(index Uint, v *Short)  {
	C.goglVertexAttribI4svEXT((C.GLuint)(index), (*C.GLshort)(v))
}
func VertexAttribI4ubvEXT(index Uint, v *Ubyte)  {
	C.goglVertexAttribI4ubvEXT((C.GLuint)(index), (*C.GLubyte)(v))
}
func VertexAttribI4usvEXT(index Uint, v *Ushort)  {
	C.goglVertexAttribI4usvEXT((C.GLuint)(index), (*C.GLushort)(v))
}
func VertexAttribIPointerEXT(index Uint, size Int, type_ Enum, stride Sizei, pointer Pointer)  {
	C.goglVertexAttribIPointerEXT((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}
func GetVertexAttribIivEXT(index Uint, pname Enum, params *Int)  {
	C.goglGetVertexAttribIivEXT((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetVertexAttribIuivEXT(index Uint, pname Enum, params *Uint)  {
	C.goglGetVertexAttribIuivEXT((C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint)(params))
}
// NV_video_capture

func BeginVideoCaptureNV(video_capture_slot Uint)  {
	C.goglBeginVideoCaptureNV((C.GLuint)(video_capture_slot))
}
func BindVideoCaptureStreamBufferNV(video_capture_slot Uint, stream Uint, frame_region Enum, offset Intptr)  {
	C.goglBindVideoCaptureStreamBufferNV((C.GLuint)(video_capture_slot), (C.GLuint)(stream), (C.GLenum)(frame_region), (C.GLintptrARB)(offset))
}
func BindVideoCaptureStreamTextureNV(video_capture_slot Uint, stream Uint, frame_region Enum, target Enum, texture Uint)  {
	C.goglBindVideoCaptureStreamTextureNV((C.GLuint)(video_capture_slot), (C.GLuint)(stream), (C.GLenum)(frame_region), (C.GLenum)(target), (C.GLuint)(texture))
}
func EndVideoCaptureNV(video_capture_slot Uint)  {
	C.goglEndVideoCaptureNV((C.GLuint)(video_capture_slot))
}
func GetVideoCaptureivNV(video_capture_slot Uint, pname Enum, params *Int)  {
	C.goglGetVideoCaptureivNV((C.GLuint)(video_capture_slot), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetVideoCaptureStreamivNV(video_capture_slot Uint, stream Uint, pname Enum, params *Int)  {
	C.goglGetVideoCaptureStreamivNV((C.GLuint)(video_capture_slot), (C.GLuint)(stream), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetVideoCaptureStreamfvNV(video_capture_slot Uint, stream Uint, pname Enum, params *Float)  {
	C.goglGetVideoCaptureStreamfvNV((C.GLuint)(video_capture_slot), (C.GLuint)(stream), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetVideoCaptureStreamdvNV(video_capture_slot Uint, stream Uint, pname Enum, params *Double)  {
	C.goglGetVideoCaptureStreamdvNV((C.GLuint)(video_capture_slot), (C.GLuint)(stream), (C.GLenum)(pname), (*C.GLdouble)(params))
}
func VideoCaptureNV(video_capture_slot Uint, sequence_num *Uint, capture_time *Uint64) Enum {
	return (Enum)(C.goglVideoCaptureNV((C.GLuint)(video_capture_slot), (*C.GLuint)(sequence_num), (*C.GLuint64EXT)(capture_time)))
}
func VideoCaptureStreamParameterivNV(video_capture_slot Uint, stream Uint, pname Enum, params *Int)  {
	C.goglVideoCaptureStreamParameterivNV((C.GLuint)(video_capture_slot), (C.GLuint)(stream), (C.GLenum)(pname), (*C.GLint)(params))
}
func VideoCaptureStreamParameterfvNV(video_capture_slot Uint, stream Uint, pname Enum, params *Float)  {
	C.goglVideoCaptureStreamParameterfvNV((C.GLuint)(video_capture_slot), (C.GLuint)(stream), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func VideoCaptureStreamParameterdvNV(video_capture_slot Uint, stream Uint, pname Enum, params *Double)  {
	C.goglVideoCaptureStreamParameterdvNV((C.GLuint)(video_capture_slot), (C.GLuint)(stream), (C.GLenum)(pname), (*C.GLdouble)(params))
}
func InitNvBindlessTexture() error {
	var ret C.int
	if ret = C.init_NV_bindless_texture(); ret != 0 {
		return errors.New("unable to initialize NV_bindless_texture")
	}
	return nil
}
func InitNvConditionalRender() error {
	var ret C.int
	if ret = C.init_NV_conditional_render(); ret != 0 {
		return errors.New("unable to initialize NV_conditional_render")
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
func InitNvDepthBufferFloat() error {
	var ret C.int
	if ret = C.init_NV_depth_buffer_float(); ret != 0 {
		return errors.New("unable to initialize NV_depth_buffer_float")
	}
	return nil
}
func InitNvEvaluators() error {
	var ret C.int
	if ret = C.init_NV_evaluators(); ret != 0 {
		return errors.New("unable to initialize NV_evaluators")
	}
	return nil
}
func InitNvExplicitMultisample() error {
	var ret C.int
	if ret = C.init_NV_explicit_multisample(); ret != 0 {
		return errors.New("unable to initialize NV_explicit_multisample")
	}
	return nil
}
func InitNvFence() error {
	var ret C.int
	if ret = C.init_NV_fence(); ret != 0 {
		return errors.New("unable to initialize NV_fence")
	}
	return nil
}
func InitNvFragmentProgram() error {
	var ret C.int
	if ret = C.init_NV_fragment_program(); ret != 0 {
		return errors.New("unable to initialize NV_fragment_program")
	}
	return nil
}
func InitNvFramebufferMultisampleCoverage() error {
	var ret C.int
	if ret = C.init_NV_framebuffer_multisample_coverage(); ret != 0 {
		return errors.New("unable to initialize NV_framebuffer_multisample_coverage")
	}
	return nil
}
func InitNvGeometryProgram4() error {
	var ret C.int
	if ret = C.init_NV_geometry_program4(); ret != 0 {
		return errors.New("unable to initialize NV_geometry_program4")
	}
	return nil
}
func InitNvGpuProgram4() error {
	var ret C.int
	if ret = C.init_NV_gpu_program4(); ret != 0 {
		return errors.New("unable to initialize NV_gpu_program4")
	}
	return nil
}
func InitNvGpuProgram5() error {
	var ret C.int
	if ret = C.init_NV_gpu_program5(); ret != 0 {
		return errors.New("unable to initialize NV_gpu_program5")
	}
	return nil
}
func InitNvGpuShader5() error {
	var ret C.int
	if ret = C.init_NV_gpu_shader5(); ret != 0 {
		return errors.New("unable to initialize NV_gpu_shader5")
	}
	return nil
}
func InitNvHalfFloat() error {
	var ret C.int
	if ret = C.init_NV_half_float(); ret != 0 {
		return errors.New("unable to initialize NV_half_float")
	}
	return nil
}
func InitNvOcclusionQuery() error {
	var ret C.int
	if ret = C.init_NV_occlusion_query(); ret != 0 {
		return errors.New("unable to initialize NV_occlusion_query")
	}
	return nil
}
func InitNvParameterBufferObject() error {
	var ret C.int
	if ret = C.init_NV_parameter_buffer_object(); ret != 0 {
		return errors.New("unable to initialize NV_parameter_buffer_object")
	}
	return nil
}
func InitNvPathRendering() error {
	var ret C.int
	if ret = C.init_NV_path_rendering(); ret != 0 {
		return errors.New("unable to initialize NV_path_rendering")
	}
	return nil
}
func InitNvPixelDataRange() error {
	var ret C.int
	if ret = C.init_NV_pixel_data_range(); ret != 0 {
		return errors.New("unable to initialize NV_pixel_data_range")
	}
	return nil
}
func InitNvPointSprite() error {
	var ret C.int
	if ret = C.init_NV_point_sprite(); ret != 0 {
		return errors.New("unable to initialize NV_point_sprite")
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
func InitNvPrimitiveRestart() error {
	var ret C.int
	if ret = C.init_NV_primitive_restart(); ret != 0 {
		return errors.New("unable to initialize NV_primitive_restart")
	}
	return nil
}
func InitNvRegisterCombiners() error {
	var ret C.int
	if ret = C.init_NV_register_combiners(); ret != 0 {
		return errors.New("unable to initialize NV_register_combiners")
	}
	return nil
}
func InitNvRegisterCombiners2() error {
	var ret C.int
	if ret = C.init_NV_register_combiners2(); ret != 0 {
		return errors.New("unable to initialize NV_register_combiners2")
	}
	return nil
}
func InitNvShaderBufferLoad() error {
	var ret C.int
	if ret = C.init_NV_shader_buffer_load(); ret != 0 {
		return errors.New("unable to initialize NV_shader_buffer_load")
	}
	return nil
}
func InitNvTextureBarrier() error {
	var ret C.int
	if ret = C.init_NV_texture_barrier(); ret != 0 {
		return errors.New("unable to initialize NV_texture_barrier")
	}
	return nil
}
func InitNvTextureMultisample() error {
	var ret C.int
	if ret = C.init_NV_texture_multisample(); ret != 0 {
		return errors.New("unable to initialize NV_texture_multisample")
	}
	return nil
}
func InitNvTransformFeedback() error {
	var ret C.int
	if ret = C.init_NV_transform_feedback(); ret != 0 {
		return errors.New("unable to initialize NV_transform_feedback")
	}
	return nil
}
func InitNvTransformFeedback2() error {
	var ret C.int
	if ret = C.init_NV_transform_feedback2(); ret != 0 {
		return errors.New("unable to initialize NV_transform_feedback2")
	}
	return nil
}
func InitNvVdpauInterop() error {
	var ret C.int
	if ret = C.init_NV_vdpau_interop(); ret != 0 {
		return errors.New("unable to initialize NV_vdpau_interop")
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
func InitNvVertexAttribInteger64bit() error {
	var ret C.int
	if ret = C.init_NV_vertex_attrib_integer_64bit(); ret != 0 {
		return errors.New("unable to initialize NV_vertex_attrib_integer_64bit")
	}
	return nil
}
func InitNvVertexBufferUnifiedMemory() error {
	var ret C.int
	if ret = C.init_NV_vertex_buffer_unified_memory(); ret != 0 {
		return errors.New("unable to initialize NV_vertex_buffer_unified_memory")
	}
	return nil
}
func InitNvVertexProgram() error {
	var ret C.int
	if ret = C.init_NV_vertex_program(); ret != 0 {
		return errors.New("unable to initialize NV_vertex_program")
	}
	return nil
}
func InitNvVertexProgram4() error {
	var ret C.int
	if ret = C.init_NV_vertex_program4(); ret != 0 {
		return errors.New("unable to initialize NV_vertex_program4")
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
// EOF