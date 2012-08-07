// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// AMD_blend_minmax_factor: http://www.opengl.org/registry/specs/AMD/blend_minmax_factor.txt
// 
// AMD_conservative_depth: http://www.opengl.org/registry/specs/AMD/conservative_depth.txt
// 
// AMD_debug_output: http://www.opengl.org/registry/specs/AMD/debug_output.txt
// 
// AMD_depth_clamp_separate: http://www.opengl.org/registry/specs/AMD/depth_clamp_separate.txt
// 
// AMD_draw_buffers_blend: http://www.opengl.org/registry/specs/AMD/draw_buffers_blend.txt
// 
// AMD_multi_draw_indirect: http://www.opengl.org/registry/specs/AMD/multi_draw_indirect.txt
// 
// AMD_name_gen_delete: http://www.opengl.org/registry/specs/AMD/name_gen_delete.txt
// 
// AMD_performance_monitor: http://www.opengl.org/registry/specs/AMD/performance_monitor.txt
// 
// AMD_pinned_memory: http://www.opengl.org/registry/specs/AMD/pinned_memory.txt
// 
// AMD_sample_positions: http://www.opengl.org/registry/specs/AMD/sample_positions.txt
// 
// AMD_seamless_cubemap_per_texture: http://www.opengl.org/registry/specs/AMD/seamless_cubemap_per_texture.txt
// 
// AMD_shader_stencil_export: http://www.opengl.org/registry/specs/AMD/shader_stencil_export.txt
// 
// AMD_stencil_operation_extended: http://www.opengl.org/registry/specs/AMD/stencil_operation_extended.txt
// 
// AMD_texture_texture4: http://www.opengl.org/registry/specs/AMD/texture_texture4.txt
// 
// AMD_transform_feedback3_lines_triangles: http://www.opengl.org/registry/specs/AMD/transform_feedback3_lines_triangles.txt
// 
// AMD_vertex_shader_layer: http://www.opengl.org/registry/specs/AMD/vertex_shader_layer.txt
// 
// AMD_vertex_shader_tessellator: http://www.opengl.org/registry/specs/AMD/vertex_shader_tessellator.txt
// 
// AMD_vertex_shader_viewport_index: http://www.opengl.org/registry/specs/AMD/vertex_shader_viewport_index.txt
// 
package amd

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
// //  AMD_blend_minmax_factor
// //  AMD_conservative_depth
// //  AMD_debug_output
// void (APIENTRYP ptrglDebugMessageEnableAMD)(GLenum category, GLenum severity, GLsizei count, GLuint* ids, GLboolean enabled);
// void (APIENTRYP ptrglDebugMessageInsertAMD)(GLenum category, GLenum severity, GLuint id, GLsizei length, GLchar* buf);
// void (APIENTRYP ptrglDebugMessageCallbackAMD)(GLDEBUGPROCAMD callback, GLvoid* userParam);
// GLuint (APIENTRYP ptrglGetDebugMessageLogAMD)(GLuint count, GLsizei bufsize, GLenum* categories, GLuint* severities, GLuint* ids, GLsizei* lengths, GLchar* message);
// //  AMD_depth_clamp_separate
// //  AMD_draw_buffers_blend
// void (APIENTRYP ptrglBlendFuncIndexedAMD)(GLuint buf, GLenum src, GLenum dst);
// void (APIENTRYP ptrglBlendFuncSeparateIndexedAMD)(GLuint buf, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha);
// void (APIENTRYP ptrglBlendEquationIndexedAMD)(GLuint buf, GLenum mode);
// void (APIENTRYP ptrglBlendEquationSeparateIndexedAMD)(GLuint buf, GLenum modeRGB, GLenum modeAlpha);
// //  AMD_multi_draw_indirect
// void (APIENTRYP ptrglMultiDrawArraysIndirectAMD)(GLenum mode, GLvoid* indirect, GLsizei primcount, GLsizei stride);
// void (APIENTRYP ptrglMultiDrawElementsIndirectAMD)(GLenum mode, GLenum type, GLvoid* indirect, GLsizei primcount, GLsizei stride);
// //  AMD_name_gen_delete
// void (APIENTRYP ptrglGenNamesAMD)(GLenum identifier, GLuint num, GLuint* names);
// void (APIENTRYP ptrglDeleteNamesAMD)(GLenum identifier, GLuint num, GLuint* names);
// GLboolean (APIENTRYP ptrglIsNameAMD)(GLenum identifier, GLuint name);
// //  AMD_performance_monitor
// void (APIENTRYP ptrglGetPerfMonitorGroupsAMD)(GLint* numGroups, GLsizei groupsSize, GLuint* groups);
// void (APIENTRYP ptrglGetPerfMonitorCountersAMD)(GLuint group, GLint* numCounters, GLint* maxActiveCounters, GLsizei counterSize, GLuint* counters);
// void (APIENTRYP ptrglGetPerfMonitorGroupStringAMD)(GLuint group, GLsizei bufSize, GLsizei* length, GLchar* groupString);
// void (APIENTRYP ptrglGetPerfMonitorCounterStringAMD)(GLuint group, GLuint counter, GLsizei bufSize, GLsizei* length, GLchar* counterString);
// void (APIENTRYP ptrglGetPerfMonitorCounterInfoAMD)(GLuint group, GLuint counter, GLenum pname, GLvoid* data);
// void (APIENTRYP ptrglGenPerfMonitorsAMD)(GLsizei n, GLuint* monitors);
// void (APIENTRYP ptrglDeletePerfMonitorsAMD)(GLsizei n, GLuint* monitors);
// void (APIENTRYP ptrglSelectPerfMonitorCountersAMD)(GLuint monitor, GLboolean enable, GLuint group, GLint numCounters, GLuint* counterList);
// void (APIENTRYP ptrglBeginPerfMonitorAMD)(GLuint monitor);
// void (APIENTRYP ptrglEndPerfMonitorAMD)(GLuint monitor);
// void (APIENTRYP ptrglGetPerfMonitorCounterDataAMD)(GLuint monitor, GLenum pname, GLsizei dataSize, GLuint* data, GLint* bytesWritten);
// //  AMD_pinned_memory
// //  AMD_sample_positions
// void (APIENTRYP ptrglSetMultisamplefvAMD)(GLenum pname, GLuint index, GLfloat* val);
// //  AMD_seamless_cubemap_per_texture
// //  AMD_shader_stencil_export
// //  AMD_stencil_operation_extended
// void (APIENTRYP ptrglStencilOpValueAMD)(GLenum face, GLuint value);
// //  AMD_texture_texture4
// //  AMD_transform_feedback3_lines_triangles
// //  AMD_vertex_shader_layer
// //  AMD_vertex_shader_tessellator
// void (APIENTRYP ptrglTessellationFactorAMD)(GLfloat factor);
// void (APIENTRYP ptrglTessellationModeAMD)(GLenum mode);
// //  AMD_vertex_shader_viewport_index
// 
// //  AMD_blend_minmax_factor
// //  AMD_conservative_depth
// //  AMD_debug_output
// void goglDebugMessageEnableAMD(GLenum category, GLenum severity, GLsizei count, GLuint* ids, GLboolean enabled) {
// 	(*ptrglDebugMessageEnableAMD)(category, severity, count, ids, enabled);
// }
// void goglDebugMessageInsertAMD(GLenum category, GLenum severity, GLuint id, GLsizei length, GLchar* buf) {
// 	(*ptrglDebugMessageInsertAMD)(category, severity, id, length, buf);
// }
// void goglDebugMessageCallbackAMD(GLDEBUGPROCAMD callback, GLvoid* userParam) {
// 	(*ptrglDebugMessageCallbackAMD)(callback, userParam);
// }
// GLuint goglGetDebugMessageLogAMD(GLuint count, GLsizei bufsize, GLenum* categories, GLuint* severities, GLuint* ids, GLsizei* lengths, GLchar* message) {
// 	return (*ptrglGetDebugMessageLogAMD)(count, bufsize, categories, severities, ids, lengths, message);
// }
// //  AMD_depth_clamp_separate
// //  AMD_draw_buffers_blend
// void goglBlendFuncIndexedAMD(GLuint buf, GLenum src, GLenum dst) {
// 	(*ptrglBlendFuncIndexedAMD)(buf, src, dst);
// }
// void goglBlendFuncSeparateIndexedAMD(GLuint buf, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
// 	(*ptrglBlendFuncSeparateIndexedAMD)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
// }
// void goglBlendEquationIndexedAMD(GLuint buf, GLenum mode) {
// 	(*ptrglBlendEquationIndexedAMD)(buf, mode);
// }
// void goglBlendEquationSeparateIndexedAMD(GLuint buf, GLenum modeRGB, GLenum modeAlpha) {
// 	(*ptrglBlendEquationSeparateIndexedAMD)(buf, modeRGB, modeAlpha);
// }
// //  AMD_multi_draw_indirect
// void goglMultiDrawArraysIndirectAMD(GLenum mode, GLvoid* indirect, GLsizei primcount, GLsizei stride) {
// 	(*ptrglMultiDrawArraysIndirectAMD)(mode, indirect, primcount, stride);
// }
// void goglMultiDrawElementsIndirectAMD(GLenum mode, GLenum type_, GLvoid* indirect, GLsizei primcount, GLsizei stride) {
// 	(*ptrglMultiDrawElementsIndirectAMD)(mode, type_, indirect, primcount, stride);
// }
// //  AMD_name_gen_delete
// void goglGenNamesAMD(GLenum identifier, GLuint num, GLuint* names) {
// 	(*ptrglGenNamesAMD)(identifier, num, names);
// }
// void goglDeleteNamesAMD(GLenum identifier, GLuint num, GLuint* names) {
// 	(*ptrglDeleteNamesAMD)(identifier, num, names);
// }
// GLboolean goglIsNameAMD(GLenum identifier, GLuint name) {
// 	return (*ptrglIsNameAMD)(identifier, name);
// }
// //  AMD_performance_monitor
// void goglGetPerfMonitorGroupsAMD(GLint* numGroups, GLsizei groupsSize, GLuint* groups) {
// 	(*ptrglGetPerfMonitorGroupsAMD)(numGroups, groupsSize, groups);
// }
// void goglGetPerfMonitorCountersAMD(GLuint group, GLint* numCounters, GLint* maxActiveCounters, GLsizei counterSize, GLuint* counters) {
// 	(*ptrglGetPerfMonitorCountersAMD)(group, numCounters, maxActiveCounters, counterSize, counters);
// }
// void goglGetPerfMonitorGroupStringAMD(GLuint group, GLsizei bufSize, GLsizei* length, GLchar* groupString) {
// 	(*ptrglGetPerfMonitorGroupStringAMD)(group, bufSize, length, groupString);
// }
// void goglGetPerfMonitorCounterStringAMD(GLuint group, GLuint counter, GLsizei bufSize, GLsizei* length, GLchar* counterString) {
// 	(*ptrglGetPerfMonitorCounterStringAMD)(group, counter, bufSize, length, counterString);
// }
// void goglGetPerfMonitorCounterInfoAMD(GLuint group, GLuint counter, GLenum pname, GLvoid* data) {
// 	(*ptrglGetPerfMonitorCounterInfoAMD)(group, counter, pname, data);
// }
// void goglGenPerfMonitorsAMD(GLsizei n, GLuint* monitors) {
// 	(*ptrglGenPerfMonitorsAMD)(n, monitors);
// }
// void goglDeletePerfMonitorsAMD(GLsizei n, GLuint* monitors) {
// 	(*ptrglDeletePerfMonitorsAMD)(n, monitors);
// }
// void goglSelectPerfMonitorCountersAMD(GLuint monitor, GLboolean enable, GLuint group, GLint numCounters, GLuint* counterList) {
// 	(*ptrglSelectPerfMonitorCountersAMD)(monitor, enable, group, numCounters, counterList);
// }
// void goglBeginPerfMonitorAMD(GLuint monitor) {
// 	(*ptrglBeginPerfMonitorAMD)(monitor);
// }
// void goglEndPerfMonitorAMD(GLuint monitor) {
// 	(*ptrglEndPerfMonitorAMD)(monitor);
// }
// void goglGetPerfMonitorCounterDataAMD(GLuint monitor, GLenum pname, GLsizei dataSize, GLuint* data, GLint* bytesWritten) {
// 	(*ptrglGetPerfMonitorCounterDataAMD)(monitor, pname, dataSize, data, bytesWritten);
// }
// //  AMD_pinned_memory
// //  AMD_sample_positions
// void goglSetMultisamplefvAMD(GLenum pname, GLuint index, GLfloat* val) {
// 	(*ptrglSetMultisamplefvAMD)(pname, index, val);
// }
// //  AMD_seamless_cubemap_per_texture
// //  AMD_shader_stencil_export
// //  AMD_stencil_operation_extended
// void goglStencilOpValueAMD(GLenum face, GLuint value) {
// 	(*ptrglStencilOpValueAMD)(face, value);
// }
// //  AMD_texture_texture4
// //  AMD_transform_feedback3_lines_triangles
// //  AMD_vertex_shader_layer
// //  AMD_vertex_shader_tessellator
// void goglTessellationFactorAMD(GLfloat factor) {
// 	(*ptrglTessellationFactorAMD)(factor);
// }
// void goglTessellationModeAMD(GLenum mode) {
// 	(*ptrglTessellationModeAMD)(mode);
// }
// //  AMD_vertex_shader_viewport_index
// 
// int init_AMD_blend_minmax_factor() {
// 	return 0;
// }
// int init_AMD_conservative_depth() {
// 	return 0;
// }
// int init_AMD_debug_output() {
// 	ptrglDebugMessageEnableAMD = goglGetProcAddress("glDebugMessageEnableAMD");
// 	if(ptrglDebugMessageEnableAMD == NULL) return 1;
// 	ptrglDebugMessageInsertAMD = goglGetProcAddress("glDebugMessageInsertAMD");
// 	if(ptrglDebugMessageInsertAMD == NULL) return 1;
// 	ptrglDebugMessageCallbackAMD = goglGetProcAddress("glDebugMessageCallbackAMD");
// 	if(ptrglDebugMessageCallbackAMD == NULL) return 1;
// 	ptrglGetDebugMessageLogAMD = goglGetProcAddress("glGetDebugMessageLogAMD");
// 	if(ptrglGetDebugMessageLogAMD == NULL) return 1;
// 	return 0;
// }
// int init_AMD_depth_clamp_separate() {
// 	return 0;
// }
// int init_AMD_draw_buffers_blend() {
// 	ptrglBlendFuncIndexedAMD = goglGetProcAddress("glBlendFuncIndexedAMD");
// 	if(ptrglBlendFuncIndexedAMD == NULL) return 1;
// 	ptrglBlendFuncSeparateIndexedAMD = goglGetProcAddress("glBlendFuncSeparateIndexedAMD");
// 	if(ptrglBlendFuncSeparateIndexedAMD == NULL) return 1;
// 	ptrglBlendEquationIndexedAMD = goglGetProcAddress("glBlendEquationIndexedAMD");
// 	if(ptrglBlendEquationIndexedAMD == NULL) return 1;
// 	ptrglBlendEquationSeparateIndexedAMD = goglGetProcAddress("glBlendEquationSeparateIndexedAMD");
// 	if(ptrglBlendEquationSeparateIndexedAMD == NULL) return 1;
// 	return 0;
// }
// int init_AMD_multi_draw_indirect() {
// 	ptrglMultiDrawArraysIndirectAMD = goglGetProcAddress("glMultiDrawArraysIndirectAMD");
// 	if(ptrglMultiDrawArraysIndirectAMD == NULL) return 1;
// 	ptrglMultiDrawElementsIndirectAMD = goglGetProcAddress("glMultiDrawElementsIndirectAMD");
// 	if(ptrglMultiDrawElementsIndirectAMD == NULL) return 1;
// 	return 0;
// }
// int init_AMD_name_gen_delete() {
// 	ptrglGenNamesAMD = goglGetProcAddress("glGenNamesAMD");
// 	if(ptrglGenNamesAMD == NULL) return 1;
// 	ptrglDeleteNamesAMD = goglGetProcAddress("glDeleteNamesAMD");
// 	if(ptrglDeleteNamesAMD == NULL) return 1;
// 	ptrglIsNameAMD = goglGetProcAddress("glIsNameAMD");
// 	if(ptrglIsNameAMD == NULL) return 1;
// 	return 0;
// }
// int init_AMD_performance_monitor() {
// 	ptrglGetPerfMonitorGroupsAMD = goglGetProcAddress("glGetPerfMonitorGroupsAMD");
// 	if(ptrglGetPerfMonitorGroupsAMD == NULL) return 1;
// 	ptrglGetPerfMonitorCountersAMD = goglGetProcAddress("glGetPerfMonitorCountersAMD");
// 	if(ptrglGetPerfMonitorCountersAMD == NULL) return 1;
// 	ptrglGetPerfMonitorGroupStringAMD = goglGetProcAddress("glGetPerfMonitorGroupStringAMD");
// 	if(ptrglGetPerfMonitorGroupStringAMD == NULL) return 1;
// 	ptrglGetPerfMonitorCounterStringAMD = goglGetProcAddress("glGetPerfMonitorCounterStringAMD");
// 	if(ptrglGetPerfMonitorCounterStringAMD == NULL) return 1;
// 	ptrglGetPerfMonitorCounterInfoAMD = goglGetProcAddress("glGetPerfMonitorCounterInfoAMD");
// 	if(ptrglGetPerfMonitorCounterInfoAMD == NULL) return 1;
// 	ptrglGenPerfMonitorsAMD = goglGetProcAddress("glGenPerfMonitorsAMD");
// 	if(ptrglGenPerfMonitorsAMD == NULL) return 1;
// 	ptrglDeletePerfMonitorsAMD = goglGetProcAddress("glDeletePerfMonitorsAMD");
// 	if(ptrglDeletePerfMonitorsAMD == NULL) return 1;
// 	ptrglSelectPerfMonitorCountersAMD = goglGetProcAddress("glSelectPerfMonitorCountersAMD");
// 	if(ptrglSelectPerfMonitorCountersAMD == NULL) return 1;
// 	ptrglBeginPerfMonitorAMD = goglGetProcAddress("glBeginPerfMonitorAMD");
// 	if(ptrglBeginPerfMonitorAMD == NULL) return 1;
// 	ptrglEndPerfMonitorAMD = goglGetProcAddress("glEndPerfMonitorAMD");
// 	if(ptrglEndPerfMonitorAMD == NULL) return 1;
// 	ptrglGetPerfMonitorCounterDataAMD = goglGetProcAddress("glGetPerfMonitorCounterDataAMD");
// 	if(ptrglGetPerfMonitorCounterDataAMD == NULL) return 1;
// 	return 0;
// }
// int init_AMD_pinned_memory() {
// 	return 0;
// }
// int init_AMD_sample_positions() {
// 	ptrglSetMultisamplefvAMD = goglGetProcAddress("glSetMultisamplefvAMD");
// 	if(ptrglSetMultisamplefvAMD == NULL) return 1;
// 	return 0;
// }
// int init_AMD_seamless_cubemap_per_texture() {
// 	return 0;
// }
// int init_AMD_shader_stencil_export() {
// 	return 0;
// }
// int init_AMD_stencil_operation_extended() {
// 	ptrglStencilOpValueAMD = goglGetProcAddress("glStencilOpValueAMD");
// 	if(ptrglStencilOpValueAMD == NULL) return 1;
// 	return 0;
// }
// int init_AMD_texture_texture4() {
// 	return 0;
// }
// int init_AMD_transform_feedback3_lines_triangles() {
// 	return 0;
// }
// int init_AMD_vertex_shader_layer() {
// 	return 0;
// }
// int init_AMD_vertex_shader_tessellator() {
// 	ptrglTessellationFactorAMD = goglGetProcAddress("glTessellationFactorAMD");
// 	if(ptrglTessellationFactorAMD == NULL) return 1;
// 	ptrglTessellationModeAMD = goglGetProcAddress("glTessellationModeAMD");
// 	if(ptrglTessellationModeAMD == NULL) return 1;
// 	return 0;
// }
// int init_AMD_vertex_shader_viewport_index() {
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

// AMD_blend_minmax_factor
const (
	FACTOR_MAX_AMD = 0x901D
	FACTOR_MIN_AMD = 0x901C
)
// AMD_conservative_depth
const (
)
// AMD_debug_output
const (
	DEBUG_CATEGORY_API_ERROR_AMD = 0x9149
	DEBUG_CATEGORY_APPLICATION_AMD = 0x914F
	DEBUG_CATEGORY_DEPRECATION_AMD = 0x914B
	DEBUG_CATEGORY_OTHER_AMD = 0x9150
	DEBUG_CATEGORY_PERFORMANCE_AMD = 0x914D
	DEBUG_CATEGORY_SHADER_COMPILER_AMD = 0x914E
	DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD = 0x914C
	DEBUG_CATEGORY_WINDOW_SYSTEM_AMD = 0x914A
	DEBUG_LOGGED_MESSAGES_AMD = 0x9145
	DEBUG_SEVERITY_HIGH_AMD = 0x9146
	DEBUG_SEVERITY_LOW_AMD = 0x9148
	DEBUG_SEVERITY_MEDIUM_AMD = 0x9147
	MAX_DEBUG_LOGGED_MESSAGES_AMD = 0x9144
	MAX_DEBUG_MESSAGE_LENGTH_AMD = 0x9143
)
// AMD_depth_clamp_separate
const (
	DEPTH_CLAMP_FAR_AMD = 0x901F
	DEPTH_CLAMP_NEAR_AMD = 0x901E
)
// AMD_draw_buffers_blend
const (
)
// AMD_multi_draw_indirect
const (
)
// AMD_name_gen_delete
const (
	DATA_BUFFER_AMD = 0x9151
	PERFORMANCE_MONITOR_AMD = 0x9152
	QUERY_OBJECT_AMD = 0x9153
	SAMPLER_OBJECT_AMD = 0x9155
	VERTEX_ARRAY_OBJECT_AMD = 0x9154
)
// AMD_performance_monitor
const (
	COUNTER_RANGE_AMD = 0x8BC1
	COUNTER_TYPE_AMD = 0x8BC0
	PERCENTAGE_AMD = 0x8BC3
	PERFMON_RESULT_AMD = 0x8BC6
	PERFMON_RESULT_AVAILABLE_AMD = 0x8BC4
	PERFMON_RESULT_SIZE_AMD = 0x8BC5
	UNSIGNED_INT64_AMD = 0x8BC2
)
// AMD_pinned_memory
const (
	EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD = 0x9160
)
// AMD_sample_positions
const (
	SUBSAMPLE_DISTANCE_AMD = 0x883F
)
// AMD_seamless_cubemap_per_texture
const (
)
// AMD_shader_stencil_export
const (
)
// AMD_stencil_operation_extended
const (
	REPLACE_VALUE_AMD = 0x874B
	SET_AMD = 0x874A
	STENCIL_BACK_OP_VALUE_AMD = 0x874D
	STENCIL_OP_VALUE_AMD = 0x874C
)
// AMD_texture_texture4
const (
)
// AMD_transform_feedback3_lines_triangles
const (
)
// AMD_vertex_shader_layer
const (
)
// AMD_vertex_shader_tessellator
const (
	CONTINUOUS_AMD = 0x9007
	DISCRETE_AMD = 0x9006
	INT_SAMPLER_BUFFER_AMD = 0x9002
	SAMPLER_BUFFER_AMD = 0x9001
	TESSELLATION_FACTOR_AMD = 0x9005
	TESSELLATION_MODE_AMD = 0x9004
	UNSIGNED_INT_SAMPLER_BUFFER_AMD = 0x9003
)
// AMD_vertex_shader_viewport_index
const (
)
// AMD_blend_minmax_factor

// AMD_conservative_depth

// AMD_debug_output

func DebugMessageEnableAMD(category Enum, severity Enum, count Sizei, ids *Uint, enabled Boolean)  {
	C.goglDebugMessageEnableAMD((C.GLenum)(category), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(ids), (C.GLboolean)(enabled))
}
func DebugMessageInsertAMD(category Enum, severity Enum, id Uint, length Sizei, buf *Char)  {
	C.goglDebugMessageInsertAMD((C.GLenum)(category), (C.GLenum)(severity), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(buf))
}
func DebugMessageCallbackAMD(callback Pointer, userParam Pointer)  {
	C.goglDebugMessageCallbackAMD((*[0]byte)(callback), (unsafe.Pointer)(userParam))
}
func GetDebugMessageLogAMD(count Uint, bufsize Sizei, categories *Enum, severities *Uint, ids *Uint, lengths *Sizei, message *Char) Uint {
	return (Uint)(C.goglGetDebugMessageLogAMD((C.GLuint)(count), (C.GLsizei)(bufsize), (*C.GLenum)(categories), (*C.GLuint)(severities), (*C.GLuint)(ids), (*C.GLsizei)(lengths), (*C.GLchar)(message)))
}
// AMD_depth_clamp_separate

// AMD_draw_buffers_blend

func BlendFuncIndexedAMD(buf Uint, src Enum, dst Enum)  {
	C.goglBlendFuncIndexedAMD((C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}
func BlendFuncSeparateIndexedAMD(buf Uint, srcRGB Enum, dstRGB Enum, srcAlpha Enum, dstAlpha Enum)  {
	C.goglBlendFuncSeparateIndexedAMD((C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}
func BlendEquationIndexedAMD(buf Uint, mode Enum)  {
	C.goglBlendEquationIndexedAMD((C.GLuint)(buf), (C.GLenum)(mode))
}
func BlendEquationSeparateIndexedAMD(buf Uint, modeRGB Enum, modeAlpha Enum)  {
	C.goglBlendEquationSeparateIndexedAMD((C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
// AMD_multi_draw_indirect

func MultiDrawArraysIndirectAMD(mode Enum, indirect Pointer, primcount Sizei, stride Sizei)  {
	C.goglMultiDrawArraysIndirectAMD((C.GLenum)(mode), (unsafe.Pointer)(indirect), (C.GLsizei)(primcount), (C.GLsizei)(stride))
}
func MultiDrawElementsIndirectAMD(mode Enum, type_ Enum, indirect Pointer, primcount Sizei, stride Sizei)  {
	C.goglMultiDrawElementsIndirectAMD((C.GLenum)(mode), (C.GLenum)(type_), (unsafe.Pointer)(indirect), (C.GLsizei)(primcount), (C.GLsizei)(stride))
}
// AMD_name_gen_delete

func GenNamesAMD(identifier Enum, num Uint, names *Uint)  {
	C.goglGenNamesAMD((C.GLenum)(identifier), (C.GLuint)(num), (*C.GLuint)(names))
}
func DeleteNamesAMD(identifier Enum, num Uint, names *Uint)  {
	C.goglDeleteNamesAMD((C.GLenum)(identifier), (C.GLuint)(num), (*C.GLuint)(names))
}
func IsNameAMD(identifier Enum, name Uint) Boolean {
	return (Boolean)(C.goglIsNameAMD((C.GLenum)(identifier), (C.GLuint)(name)))
}
// AMD_performance_monitor

func GetPerfMonitorGroupsAMD(numGroups *Int, groupsSize Sizei, groups *Uint)  {
	C.goglGetPerfMonitorGroupsAMD((*C.GLint)(numGroups), (C.GLsizei)(groupsSize), (*C.GLuint)(groups))
}
func GetPerfMonitorCountersAMD(group Uint, numCounters *Int, maxActiveCounters *Int, counterSize Sizei, counters *Uint)  {
	C.goglGetPerfMonitorCountersAMD((C.GLuint)(group), (*C.GLint)(numCounters), (*C.GLint)(maxActiveCounters), (C.GLsizei)(counterSize), (*C.GLuint)(counters))
}
func GetPerfMonitorGroupStringAMD(group Uint, bufSize Sizei, length *Sizei, groupString *Char)  {
	C.goglGetPerfMonitorGroupStringAMD((C.GLuint)(group), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(groupString))
}
func GetPerfMonitorCounterStringAMD(group Uint, counter Uint, bufSize Sizei, length *Sizei, counterString *Char)  {
	C.goglGetPerfMonitorCounterStringAMD((C.GLuint)(group), (C.GLuint)(counter), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(counterString))
}
func GetPerfMonitorCounterInfoAMD(group Uint, counter Uint, pname Enum, data Pointer)  {
	C.goglGetPerfMonitorCounterInfoAMD((C.GLuint)(group), (C.GLuint)(counter), (C.GLenum)(pname), (unsafe.Pointer)(data))
}
func GenPerfMonitorsAMD(n Sizei, monitors *Uint)  {
	C.goglGenPerfMonitorsAMD((C.GLsizei)(n), (*C.GLuint)(monitors))
}
func DeletePerfMonitorsAMD(n Sizei, monitors *Uint)  {
	C.goglDeletePerfMonitorsAMD((C.GLsizei)(n), (*C.GLuint)(monitors))
}
func SelectPerfMonitorCountersAMD(monitor Uint, enable Boolean, group Uint, numCounters Int, counterList *Uint)  {
	C.goglSelectPerfMonitorCountersAMD((C.GLuint)(monitor), (C.GLboolean)(enable), (C.GLuint)(group), (C.GLint)(numCounters), (*C.GLuint)(counterList))
}
func BeginPerfMonitorAMD(monitor Uint)  {
	C.goglBeginPerfMonitorAMD((C.GLuint)(monitor))
}
func EndPerfMonitorAMD(monitor Uint)  {
	C.goglEndPerfMonitorAMD((C.GLuint)(monitor))
}
func GetPerfMonitorCounterDataAMD(monitor Uint, pname Enum, dataSize Sizei, data *Uint, bytesWritten *Int)  {
	C.goglGetPerfMonitorCounterDataAMD((C.GLuint)(monitor), (C.GLenum)(pname), (C.GLsizei)(dataSize), (*C.GLuint)(data), (*C.GLint)(bytesWritten))
}
// AMD_pinned_memory

// AMD_sample_positions

func SetMultisamplefvAMD(pname Enum, index Uint, val *Float)  {
	C.goglSetMultisamplefvAMD((C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(val))
}
// AMD_seamless_cubemap_per_texture

// AMD_shader_stencil_export

// AMD_stencil_operation_extended

func StencilOpValueAMD(face Enum, value Uint)  {
	C.goglStencilOpValueAMD((C.GLenum)(face), (C.GLuint)(value))
}
// AMD_texture_texture4

// AMD_transform_feedback3_lines_triangles

// AMD_vertex_shader_layer

// AMD_vertex_shader_tessellator

func TessellationFactorAMD(factor Float)  {
	C.goglTessellationFactorAMD((C.GLfloat)(factor))
}
func TessellationModeAMD(mode Enum)  {
	C.goglTessellationModeAMD((C.GLenum)(mode))
}
// AMD_vertex_shader_viewport_index

func InitAmdBlendMinmaxFactor() error {
	var ret C.int
	if ret = C.init_AMD_blend_minmax_factor(); ret != 0 {
		return errors.New("unable to initialize AMD_blend_minmax_factor")
	}
	return nil
}
func InitAmdConservativeDepth() error {
	var ret C.int
	if ret = C.init_AMD_conservative_depth(); ret != 0 {
		return errors.New("unable to initialize AMD_conservative_depth")
	}
	return nil
}
func InitAmdDebugOutput() error {
	var ret C.int
	if ret = C.init_AMD_debug_output(); ret != 0 {
		return errors.New("unable to initialize AMD_debug_output")
	}
	return nil
}
func InitAmdDepthClampSeparate() error {
	var ret C.int
	if ret = C.init_AMD_depth_clamp_separate(); ret != 0 {
		return errors.New("unable to initialize AMD_depth_clamp_separate")
	}
	return nil
}
func InitAmdDrawBuffersBlend() error {
	var ret C.int
	if ret = C.init_AMD_draw_buffers_blend(); ret != 0 {
		return errors.New("unable to initialize AMD_draw_buffers_blend")
	}
	return nil
}
func InitAmdMultiDrawIndirect() error {
	var ret C.int
	if ret = C.init_AMD_multi_draw_indirect(); ret != 0 {
		return errors.New("unable to initialize AMD_multi_draw_indirect")
	}
	return nil
}
func InitAmdNameGenDelete() error {
	var ret C.int
	if ret = C.init_AMD_name_gen_delete(); ret != 0 {
		return errors.New("unable to initialize AMD_name_gen_delete")
	}
	return nil
}
func InitAmdPerformanceMonitor() error {
	var ret C.int
	if ret = C.init_AMD_performance_monitor(); ret != 0 {
		return errors.New("unable to initialize AMD_performance_monitor")
	}
	return nil
}
func InitAmdPinnedMemory() error {
	var ret C.int
	if ret = C.init_AMD_pinned_memory(); ret != 0 {
		return errors.New("unable to initialize AMD_pinned_memory")
	}
	return nil
}
func InitAmdSamplePositions() error {
	var ret C.int
	if ret = C.init_AMD_sample_positions(); ret != 0 {
		return errors.New("unable to initialize AMD_sample_positions")
	}
	return nil
}
func InitAmdSeamlessCubemapPerTexture() error {
	var ret C.int
	if ret = C.init_AMD_seamless_cubemap_per_texture(); ret != 0 {
		return errors.New("unable to initialize AMD_seamless_cubemap_per_texture")
	}
	return nil
}
func InitAmdShaderStencilExport() error {
	var ret C.int
	if ret = C.init_AMD_shader_stencil_export(); ret != 0 {
		return errors.New("unable to initialize AMD_shader_stencil_export")
	}
	return nil
}
func InitAmdStencilOperationExtended() error {
	var ret C.int
	if ret = C.init_AMD_stencil_operation_extended(); ret != 0 {
		return errors.New("unable to initialize AMD_stencil_operation_extended")
	}
	return nil
}
func InitAmdTextureTexture4() error {
	var ret C.int
	if ret = C.init_AMD_texture_texture4(); ret != 0 {
		return errors.New("unable to initialize AMD_texture_texture4")
	}
	return nil
}
func InitAmdTransformFeedback3LinesTriangles() error {
	var ret C.int
	if ret = C.init_AMD_transform_feedback3_lines_triangles(); ret != 0 {
		return errors.New("unable to initialize AMD_transform_feedback3_lines_triangles")
	}
	return nil
}
func InitAmdVertexShaderLayer() error {
	var ret C.int
	if ret = C.init_AMD_vertex_shader_layer(); ret != 0 {
		return errors.New("unable to initialize AMD_vertex_shader_layer")
	}
	return nil
}
func InitAmdVertexShaderTessellator() error {
	var ret C.int
	if ret = C.init_AMD_vertex_shader_tessellator(); ret != 0 {
		return errors.New("unable to initialize AMD_vertex_shader_tessellator")
	}
	return nil
}
func InitAmdVertexShaderViewportIndex() error {
	var ret C.int
	if ret = C.init_AMD_vertex_shader_viewport_index(); ret != 0 {
		return errors.New("unable to initialize AMD_vertex_shader_viewport_index")
	}
	return nil
}
// EOF