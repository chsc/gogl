// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// ATI_draw_buffers: http://www.opengl.org/registry/specs/ATI/draw_buffers.txt
// 
// ATI_element_array: http://www.opengl.org/registry/specs/ATI/element_array.txt
// 
// ATI_envmap_bumpmap: http://www.opengl.org/registry/specs/ATI/envmap_bumpmap.txt
// 
// ATI_fragment_shader: http://www.opengl.org/registry/specs/ATI/fragment_shader.txt
// 
// ATI_map_object_buffer: http://www.opengl.org/registry/specs/ATI/map_object_buffer.txt
// 
// ATI_pn_triangles: http://www.opengl.org/registry/specs/ATI/pn_triangles.txt
// 
// ATI_separate_stencil: http://www.opengl.org/registry/specs/ATI/separate_stencil.txt
// 
// ATI_vertex_array_object: http://www.opengl.org/registry/specs/ATI/vertex_array_object.txt
// 
// ATI_vertex_attrib_array_object: http://www.opengl.org/registry/specs/ATI/vertex_attrib_array_object.txt
// 
// ATI_vertex_streams: http://www.opengl.org/registry/specs/ATI/vertex_streams.txt
// 
package ati

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
// //  ATI_draw_buffers
// void (APIENTRYP ptrglDrawBuffersATI)(GLsizei n, GLenum* bufs);
// //  ATI_element_array
// void (APIENTRYP ptrglElementPointerATI)(GLenum type, GLvoid* pointer);
// void (APIENTRYP ptrglDrawElementArrayATI)(GLenum mode, GLsizei count);
// void (APIENTRYP ptrglDrawRangeElementArrayATI)(GLenum mode, GLuint start, GLuint end, GLsizei count);
// //  ATI_envmap_bumpmap
// void (APIENTRYP ptrglTexBumpParameterivATI)(GLenum pname, GLint* param);
// void (APIENTRYP ptrglTexBumpParameterfvATI)(GLenum pname, GLfloat* param);
// void (APIENTRYP ptrglGetTexBumpParameterivATI)(GLenum pname, GLint* param);
// void (APIENTRYP ptrglGetTexBumpParameterfvATI)(GLenum pname, GLfloat* param);
// //  ATI_fragment_shader
// GLuint (APIENTRYP ptrglGenFragmentShadersATI)(GLuint range);
// void (APIENTRYP ptrglBindFragmentShaderATI)(GLuint id);
// void (APIENTRYP ptrglDeleteFragmentShaderATI)(GLuint id);
// void (APIENTRYP ptrglBeginFragmentShaderATI)();
// void (APIENTRYP ptrglEndFragmentShaderATI)();
// void (APIENTRYP ptrglPassTexCoordATI)(GLuint dst, GLuint coord, GLenum swizzle);
// void (APIENTRYP ptrglSampleMapATI)(GLuint dst, GLuint interp, GLenum swizzle);
// void (APIENTRYP ptrglColorFragmentOp1ATI)(GLenum op, GLuint dst, GLuint dstMask, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod);
// void (APIENTRYP ptrglColorFragmentOp2ATI)(GLenum op, GLuint dst, GLuint dstMask, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod, GLuint arg2, GLuint arg2Rep, GLuint arg2Mod);
// void (APIENTRYP ptrglColorFragmentOp3ATI)(GLenum op, GLuint dst, GLuint dstMask, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod, GLuint arg2, GLuint arg2Rep, GLuint arg2Mod, GLuint arg3, GLuint arg3Rep, GLuint arg3Mod);
// void (APIENTRYP ptrglAlphaFragmentOp1ATI)(GLenum op, GLuint dst, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod);
// void (APIENTRYP ptrglAlphaFragmentOp2ATI)(GLenum op, GLuint dst, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod, GLuint arg2, GLuint arg2Rep, GLuint arg2Mod);
// void (APIENTRYP ptrglAlphaFragmentOp3ATI)(GLenum op, GLuint dst, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod, GLuint arg2, GLuint arg2Rep, GLuint arg2Mod, GLuint arg3, GLuint arg3Rep, GLuint arg3Mod);
// void (APIENTRYP ptrglSetFragmentShaderConstantATI)(GLuint dst, GLfloat* value);
// //  ATI_map_object_buffer
// GLvoid* (APIENTRYP ptrglMapObjectBufferATI)(GLuint buffer);
// void (APIENTRYP ptrglUnmapObjectBufferATI)(GLuint buffer);
// //  ATI_pn_triangles
// void (APIENTRYP ptrglPNTrianglesiATI)(GLenum pname, GLint param);
// void (APIENTRYP ptrglPNTrianglesfATI)(GLenum pname, GLfloat param);
// //  ATI_separate_stencil
// void (APIENTRYP ptrglStencilOpSeparateATI)(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass);
// void (APIENTRYP ptrglStencilFuncSeparateATI)(GLenum frontfunc, GLenum backfunc, GLint ref, GLuint mask);
// //  ATI_vertex_array_object
// GLuint (APIENTRYP ptrglNewObjectBufferATI)(GLsizei size, GLvoid* pointer, GLenum usage);
// GLboolean (APIENTRYP ptrglIsObjectBufferATI)(GLuint buffer);
// void (APIENTRYP ptrglUpdateObjectBufferATI)(GLuint buffer, GLuint offset, GLsizei size, GLvoid* pointer, GLenum preserve);
// void (APIENTRYP ptrglGetObjectBufferfvATI)(GLuint buffer, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetObjectBufferivATI)(GLuint buffer, GLenum pname, GLint* params);
// void (APIENTRYP ptrglFreeObjectBufferATI)(GLuint buffer);
// void (APIENTRYP ptrglArrayObjectATI)(GLenum array, GLint size, GLenum type, GLsizei stride, GLuint buffer, GLuint offset);
// void (APIENTRYP ptrglGetArrayObjectfvATI)(GLenum array, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetArrayObjectivATI)(GLenum array, GLenum pname, GLint* params);
// void (APIENTRYP ptrglVariantArrayObjectATI)(GLuint id, GLenum type, GLsizei stride, GLuint buffer, GLuint offset);
// void (APIENTRYP ptrglGetVariantArrayObjectfvATI)(GLuint id, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetVariantArrayObjectivATI)(GLuint id, GLenum pname, GLint* params);
// //  ATI_vertex_attrib_array_object
// void (APIENTRYP ptrglVertexAttribArrayObjectATI)(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLuint buffer, GLuint offset);
// void (APIENTRYP ptrglGetVertexAttribArrayObjectfvATI)(GLuint index, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetVertexAttribArrayObjectivATI)(GLuint index, GLenum pname, GLint* params);
// //  ATI_vertex_streams
// void (APIENTRYP ptrglVertexStream1sATI)(GLenum stream, GLshort x);
// void (APIENTRYP ptrglVertexStream1svATI)(GLenum stream, GLshort* coords);
// void (APIENTRYP ptrglVertexStream1iATI)(GLenum stream, GLint x);
// void (APIENTRYP ptrglVertexStream1ivATI)(GLenum stream, GLint* coords);
// void (APIENTRYP ptrglVertexStream1fATI)(GLenum stream, GLfloat x);
// void (APIENTRYP ptrglVertexStream1fvATI)(GLenum stream, GLfloat* coords);
// void (APIENTRYP ptrglVertexStream1dATI)(GLenum stream, GLdouble x);
// void (APIENTRYP ptrglVertexStream1dvATI)(GLenum stream, GLdouble* coords);
// void (APIENTRYP ptrglVertexStream2sATI)(GLenum stream, GLshort x, GLshort y);
// void (APIENTRYP ptrglVertexStream2svATI)(GLenum stream, GLshort* coords);
// void (APIENTRYP ptrglVertexStream2iATI)(GLenum stream, GLint x, GLint y);
// void (APIENTRYP ptrglVertexStream2ivATI)(GLenum stream, GLint* coords);
// void (APIENTRYP ptrglVertexStream2fATI)(GLenum stream, GLfloat x, GLfloat y);
// void (APIENTRYP ptrglVertexStream2fvATI)(GLenum stream, GLfloat* coords);
// void (APIENTRYP ptrglVertexStream2dATI)(GLenum stream, GLdouble x, GLdouble y);
// void (APIENTRYP ptrglVertexStream2dvATI)(GLenum stream, GLdouble* coords);
// void (APIENTRYP ptrglVertexStream3sATI)(GLenum stream, GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrglVertexStream3svATI)(GLenum stream, GLshort* coords);
// void (APIENTRYP ptrglVertexStream3iATI)(GLenum stream, GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglVertexStream3ivATI)(GLenum stream, GLint* coords);
// void (APIENTRYP ptrglVertexStream3fATI)(GLenum stream, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglVertexStream3fvATI)(GLenum stream, GLfloat* coords);
// void (APIENTRYP ptrglVertexStream3dATI)(GLenum stream, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglVertexStream3dvATI)(GLenum stream, GLdouble* coords);
// void (APIENTRYP ptrglVertexStream4sATI)(GLenum stream, GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrglVertexStream4svATI)(GLenum stream, GLshort* coords);
// void (APIENTRYP ptrglVertexStream4iATI)(GLenum stream, GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglVertexStream4ivATI)(GLenum stream, GLint* coords);
// void (APIENTRYP ptrglVertexStream4fATI)(GLenum stream, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglVertexStream4fvATI)(GLenum stream, GLfloat* coords);
// void (APIENTRYP ptrglVertexStream4dATI)(GLenum stream, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglVertexStream4dvATI)(GLenum stream, GLdouble* coords);
// void (APIENTRYP ptrglNormalStream3bATI)(GLenum stream, GLbyte nx, GLbyte ny, GLbyte nz);
// void (APIENTRYP ptrglNormalStream3bvATI)(GLenum stream, GLbyte* coords);
// void (APIENTRYP ptrglNormalStream3sATI)(GLenum stream, GLshort nx, GLshort ny, GLshort nz);
// void (APIENTRYP ptrglNormalStream3svATI)(GLenum stream, GLshort* coords);
// void (APIENTRYP ptrglNormalStream3iATI)(GLenum stream, GLint nx, GLint ny, GLint nz);
// void (APIENTRYP ptrglNormalStream3ivATI)(GLenum stream, GLint* coords);
// void (APIENTRYP ptrglNormalStream3fATI)(GLenum stream, GLfloat nx, GLfloat ny, GLfloat nz);
// void (APIENTRYP ptrglNormalStream3fvATI)(GLenum stream, GLfloat* coords);
// void (APIENTRYP ptrglNormalStream3dATI)(GLenum stream, GLdouble nx, GLdouble ny, GLdouble nz);
// void (APIENTRYP ptrglNormalStream3dvATI)(GLenum stream, GLdouble* coords);
// void (APIENTRYP ptrglClientActiveVertexStreamATI)(GLenum stream);
// void (APIENTRYP ptrglVertexBlendEnviATI)(GLenum pname, GLint param);
// void (APIENTRYP ptrglVertexBlendEnvfATI)(GLenum pname, GLfloat param);
// 
// //  ATI_draw_buffers
// void goglDrawBuffersATI(GLsizei n, GLenum* bufs) {
// 	(*ptrglDrawBuffersATI)(n, bufs);
// }
// //  ATI_element_array
// void goglElementPointerATI(GLenum type_, GLvoid* pointer) {
// 	(*ptrglElementPointerATI)(type_, pointer);
// }
// void goglDrawElementArrayATI(GLenum mode, GLsizei count) {
// 	(*ptrglDrawElementArrayATI)(mode, count);
// }
// void goglDrawRangeElementArrayATI(GLenum mode, GLuint start, GLuint end, GLsizei count) {
// 	(*ptrglDrawRangeElementArrayATI)(mode, start, end, count);
// }
// //  ATI_envmap_bumpmap
// void goglTexBumpParameterivATI(GLenum pname, GLint* param) {
// 	(*ptrglTexBumpParameterivATI)(pname, param);
// }
// void goglTexBumpParameterfvATI(GLenum pname, GLfloat* param) {
// 	(*ptrglTexBumpParameterfvATI)(pname, param);
// }
// void goglGetTexBumpParameterivATI(GLenum pname, GLint* param) {
// 	(*ptrglGetTexBumpParameterivATI)(pname, param);
// }
// void goglGetTexBumpParameterfvATI(GLenum pname, GLfloat* param) {
// 	(*ptrglGetTexBumpParameterfvATI)(pname, param);
// }
// //  ATI_fragment_shader
// GLuint goglGenFragmentShadersATI(GLuint range_) {
// 	return (*ptrglGenFragmentShadersATI)(range_);
// }
// void goglBindFragmentShaderATI(GLuint id) {
// 	(*ptrglBindFragmentShaderATI)(id);
// }
// void goglDeleteFragmentShaderATI(GLuint id) {
// 	(*ptrglDeleteFragmentShaderATI)(id);
// }
// void goglBeginFragmentShaderATI() {
// 	(*ptrglBeginFragmentShaderATI)();
// }
// void goglEndFragmentShaderATI() {
// 	(*ptrglEndFragmentShaderATI)();
// }
// void goglPassTexCoordATI(GLuint dst, GLuint coord, GLenum swizzle) {
// 	(*ptrglPassTexCoordATI)(dst, coord, swizzle);
// }
// void goglSampleMapATI(GLuint dst, GLuint interp, GLenum swizzle) {
// 	(*ptrglSampleMapATI)(dst, interp, swizzle);
// }
// void goglColorFragmentOp1ATI(GLenum op, GLuint dst, GLuint dstMask, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod) {
// 	(*ptrglColorFragmentOp1ATI)(op, dst, dstMask, dstMod, arg1, arg1Rep, arg1Mod);
// }
// void goglColorFragmentOp2ATI(GLenum op, GLuint dst, GLuint dstMask, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod, GLuint arg2, GLuint arg2Rep, GLuint arg2Mod) {
// 	(*ptrglColorFragmentOp2ATI)(op, dst, dstMask, dstMod, arg1, arg1Rep, arg1Mod, arg2, arg2Rep, arg2Mod);
// }
// void goglColorFragmentOp3ATI(GLenum op, GLuint dst, GLuint dstMask, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod, GLuint arg2, GLuint arg2Rep, GLuint arg2Mod, GLuint arg3, GLuint arg3Rep, GLuint arg3Mod) {
// 	(*ptrglColorFragmentOp3ATI)(op, dst, dstMask, dstMod, arg1, arg1Rep, arg1Mod, arg2, arg2Rep, arg2Mod, arg3, arg3Rep, arg3Mod);
// }
// void goglAlphaFragmentOp1ATI(GLenum op, GLuint dst, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod) {
// 	(*ptrglAlphaFragmentOp1ATI)(op, dst, dstMod, arg1, arg1Rep, arg1Mod);
// }
// void goglAlphaFragmentOp2ATI(GLenum op, GLuint dst, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod, GLuint arg2, GLuint arg2Rep, GLuint arg2Mod) {
// 	(*ptrglAlphaFragmentOp2ATI)(op, dst, dstMod, arg1, arg1Rep, arg1Mod, arg2, arg2Rep, arg2Mod);
// }
// void goglAlphaFragmentOp3ATI(GLenum op, GLuint dst, GLuint dstMod, GLuint arg1, GLuint arg1Rep, GLuint arg1Mod, GLuint arg2, GLuint arg2Rep, GLuint arg2Mod, GLuint arg3, GLuint arg3Rep, GLuint arg3Mod) {
// 	(*ptrglAlphaFragmentOp3ATI)(op, dst, dstMod, arg1, arg1Rep, arg1Mod, arg2, arg2Rep, arg2Mod, arg3, arg3Rep, arg3Mod);
// }
// void goglSetFragmentShaderConstantATI(GLuint dst, GLfloat* value) {
// 	(*ptrglSetFragmentShaderConstantATI)(dst, value);
// }
// //  ATI_map_object_buffer
// GLvoid* goglMapObjectBufferATI(GLuint buffer) {
// 	return (*ptrglMapObjectBufferATI)(buffer);
// }
// void goglUnmapObjectBufferATI(GLuint buffer) {
// 	(*ptrglUnmapObjectBufferATI)(buffer);
// }
// //  ATI_pn_triangles
// void goglPNTrianglesiATI(GLenum pname, GLint param) {
// 	(*ptrglPNTrianglesiATI)(pname, param);
// }
// void goglPNTrianglesfATI(GLenum pname, GLfloat param) {
// 	(*ptrglPNTrianglesfATI)(pname, param);
// }
// //  ATI_separate_stencil
// void goglStencilOpSeparateATI(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
// 	(*ptrglStencilOpSeparateATI)(face, sfail, dpfail, dppass);
// }
// void goglStencilFuncSeparateATI(GLenum frontfunc, GLenum backfunc, GLint ref, GLuint mask) {
// 	(*ptrglStencilFuncSeparateATI)(frontfunc, backfunc, ref, mask);
// }
// //  ATI_vertex_array_object
// GLuint goglNewObjectBufferATI(GLsizei size, GLvoid* pointer, GLenum usage) {
// 	return (*ptrglNewObjectBufferATI)(size, pointer, usage);
// }
// GLboolean goglIsObjectBufferATI(GLuint buffer) {
// 	return (*ptrglIsObjectBufferATI)(buffer);
// }
// void goglUpdateObjectBufferATI(GLuint buffer, GLuint offset, GLsizei size, GLvoid* pointer, GLenum preserve) {
// 	(*ptrglUpdateObjectBufferATI)(buffer, offset, size, pointer, preserve);
// }
// void goglGetObjectBufferfvATI(GLuint buffer, GLenum pname, GLfloat* params) {
// 	(*ptrglGetObjectBufferfvATI)(buffer, pname, params);
// }
// void goglGetObjectBufferivATI(GLuint buffer, GLenum pname, GLint* params) {
// 	(*ptrglGetObjectBufferivATI)(buffer, pname, params);
// }
// void goglFreeObjectBufferATI(GLuint buffer) {
// 	(*ptrglFreeObjectBufferATI)(buffer);
// }
// void goglArrayObjectATI(GLenum array, GLint size, GLenum type_, GLsizei stride, GLuint buffer, GLuint offset) {
// 	(*ptrglArrayObjectATI)(array, size, type_, stride, buffer, offset);
// }
// void goglGetArrayObjectfvATI(GLenum array, GLenum pname, GLfloat* params) {
// 	(*ptrglGetArrayObjectfvATI)(array, pname, params);
// }
// void goglGetArrayObjectivATI(GLenum array, GLenum pname, GLint* params) {
// 	(*ptrglGetArrayObjectivATI)(array, pname, params);
// }
// void goglVariantArrayObjectATI(GLuint id, GLenum type_, GLsizei stride, GLuint buffer, GLuint offset) {
// 	(*ptrglVariantArrayObjectATI)(id, type_, stride, buffer, offset);
// }
// void goglGetVariantArrayObjectfvATI(GLuint id, GLenum pname, GLfloat* params) {
// 	(*ptrglGetVariantArrayObjectfvATI)(id, pname, params);
// }
// void goglGetVariantArrayObjectivATI(GLuint id, GLenum pname, GLint* params) {
// 	(*ptrglGetVariantArrayObjectivATI)(id, pname, params);
// }
// //  ATI_vertex_attrib_array_object
// void goglVertexAttribArrayObjectATI(GLuint index, GLint size, GLenum type_, GLboolean normalized, GLsizei stride, GLuint buffer, GLuint offset) {
// 	(*ptrglVertexAttribArrayObjectATI)(index, size, type_, normalized, stride, buffer, offset);
// }
// void goglGetVertexAttribArrayObjectfvATI(GLuint index, GLenum pname, GLfloat* params) {
// 	(*ptrglGetVertexAttribArrayObjectfvATI)(index, pname, params);
// }
// void goglGetVertexAttribArrayObjectivATI(GLuint index, GLenum pname, GLint* params) {
// 	(*ptrglGetVertexAttribArrayObjectivATI)(index, pname, params);
// }
// //  ATI_vertex_streams
// void goglVertexStream1sATI(GLenum stream, GLshort x) {
// 	(*ptrglVertexStream1sATI)(stream, x);
// }
// void goglVertexStream1svATI(GLenum stream, GLshort* coords) {
// 	(*ptrglVertexStream1svATI)(stream, coords);
// }
// void goglVertexStream1iATI(GLenum stream, GLint x) {
// 	(*ptrglVertexStream1iATI)(stream, x);
// }
// void goglVertexStream1ivATI(GLenum stream, GLint* coords) {
// 	(*ptrglVertexStream1ivATI)(stream, coords);
// }
// void goglVertexStream1fATI(GLenum stream, GLfloat x) {
// 	(*ptrglVertexStream1fATI)(stream, x);
// }
// void goglVertexStream1fvATI(GLenum stream, GLfloat* coords) {
// 	(*ptrglVertexStream1fvATI)(stream, coords);
// }
// void goglVertexStream1dATI(GLenum stream, GLdouble x) {
// 	(*ptrglVertexStream1dATI)(stream, x);
// }
// void goglVertexStream1dvATI(GLenum stream, GLdouble* coords) {
// 	(*ptrglVertexStream1dvATI)(stream, coords);
// }
// void goglVertexStream2sATI(GLenum stream, GLshort x, GLshort y) {
// 	(*ptrglVertexStream2sATI)(stream, x, y);
// }
// void goglVertexStream2svATI(GLenum stream, GLshort* coords) {
// 	(*ptrglVertexStream2svATI)(stream, coords);
// }
// void goglVertexStream2iATI(GLenum stream, GLint x, GLint y) {
// 	(*ptrglVertexStream2iATI)(stream, x, y);
// }
// void goglVertexStream2ivATI(GLenum stream, GLint* coords) {
// 	(*ptrglVertexStream2ivATI)(stream, coords);
// }
// void goglVertexStream2fATI(GLenum stream, GLfloat x, GLfloat y) {
// 	(*ptrglVertexStream2fATI)(stream, x, y);
// }
// void goglVertexStream2fvATI(GLenum stream, GLfloat* coords) {
// 	(*ptrglVertexStream2fvATI)(stream, coords);
// }
// void goglVertexStream2dATI(GLenum stream, GLdouble x, GLdouble y) {
// 	(*ptrglVertexStream2dATI)(stream, x, y);
// }
// void goglVertexStream2dvATI(GLenum stream, GLdouble* coords) {
// 	(*ptrglVertexStream2dvATI)(stream, coords);
// }
// void goglVertexStream3sATI(GLenum stream, GLshort x, GLshort y, GLshort z) {
// 	(*ptrglVertexStream3sATI)(stream, x, y, z);
// }
// void goglVertexStream3svATI(GLenum stream, GLshort* coords) {
// 	(*ptrglVertexStream3svATI)(stream, coords);
// }
// void goglVertexStream3iATI(GLenum stream, GLint x, GLint y, GLint z) {
// 	(*ptrglVertexStream3iATI)(stream, x, y, z);
// }
// void goglVertexStream3ivATI(GLenum stream, GLint* coords) {
// 	(*ptrglVertexStream3ivATI)(stream, coords);
// }
// void goglVertexStream3fATI(GLenum stream, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglVertexStream3fATI)(stream, x, y, z);
// }
// void goglVertexStream3fvATI(GLenum stream, GLfloat* coords) {
// 	(*ptrglVertexStream3fvATI)(stream, coords);
// }
// void goglVertexStream3dATI(GLenum stream, GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglVertexStream3dATI)(stream, x, y, z);
// }
// void goglVertexStream3dvATI(GLenum stream, GLdouble* coords) {
// 	(*ptrglVertexStream3dvATI)(stream, coords);
// }
// void goglVertexStream4sATI(GLenum stream, GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrglVertexStream4sATI)(stream, x, y, z, w);
// }
// void goglVertexStream4svATI(GLenum stream, GLshort* coords) {
// 	(*ptrglVertexStream4svATI)(stream, coords);
// }
// void goglVertexStream4iATI(GLenum stream, GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrglVertexStream4iATI)(stream, x, y, z, w);
// }
// void goglVertexStream4ivATI(GLenum stream, GLint* coords) {
// 	(*ptrglVertexStream4ivATI)(stream, coords);
// }
// void goglVertexStream4fATI(GLenum stream, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglVertexStream4fATI)(stream, x, y, z, w);
// }
// void goglVertexStream4fvATI(GLenum stream, GLfloat* coords) {
// 	(*ptrglVertexStream4fvATI)(stream, coords);
// }
// void goglVertexStream4dATI(GLenum stream, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglVertexStream4dATI)(stream, x, y, z, w);
// }
// void goglVertexStream4dvATI(GLenum stream, GLdouble* coords) {
// 	(*ptrglVertexStream4dvATI)(stream, coords);
// }
// void goglNormalStream3bATI(GLenum stream, GLbyte nx, GLbyte ny, GLbyte nz) {
// 	(*ptrglNormalStream3bATI)(stream, nx, ny, nz);
// }
// void goglNormalStream3bvATI(GLenum stream, GLbyte* coords) {
// 	(*ptrglNormalStream3bvATI)(stream, coords);
// }
// void goglNormalStream3sATI(GLenum stream, GLshort nx, GLshort ny, GLshort nz) {
// 	(*ptrglNormalStream3sATI)(stream, nx, ny, nz);
// }
// void goglNormalStream3svATI(GLenum stream, GLshort* coords) {
// 	(*ptrglNormalStream3svATI)(stream, coords);
// }
// void goglNormalStream3iATI(GLenum stream, GLint nx, GLint ny, GLint nz) {
// 	(*ptrglNormalStream3iATI)(stream, nx, ny, nz);
// }
// void goglNormalStream3ivATI(GLenum stream, GLint* coords) {
// 	(*ptrglNormalStream3ivATI)(stream, coords);
// }
// void goglNormalStream3fATI(GLenum stream, GLfloat nx, GLfloat ny, GLfloat nz) {
// 	(*ptrglNormalStream3fATI)(stream, nx, ny, nz);
// }
// void goglNormalStream3fvATI(GLenum stream, GLfloat* coords) {
// 	(*ptrglNormalStream3fvATI)(stream, coords);
// }
// void goglNormalStream3dATI(GLenum stream, GLdouble nx, GLdouble ny, GLdouble nz) {
// 	(*ptrglNormalStream3dATI)(stream, nx, ny, nz);
// }
// void goglNormalStream3dvATI(GLenum stream, GLdouble* coords) {
// 	(*ptrglNormalStream3dvATI)(stream, coords);
// }
// void goglClientActiveVertexStreamATI(GLenum stream) {
// 	(*ptrglClientActiveVertexStreamATI)(stream);
// }
// void goglVertexBlendEnviATI(GLenum pname, GLint param) {
// 	(*ptrglVertexBlendEnviATI)(pname, param);
// }
// void goglVertexBlendEnvfATI(GLenum pname, GLfloat param) {
// 	(*ptrglVertexBlendEnvfATI)(pname, param);
// }
// 
// int init_ATI_draw_buffers() {
// 	ptrglDrawBuffersATI = goglGetProcAddress("glDrawBuffersATI");
// 	if(ptrglDrawBuffersATI == NULL) return 1;
// 	return 0;
// }
// int init_ATI_element_array() {
// 	ptrglElementPointerATI = goglGetProcAddress("glElementPointerATI");
// 	if(ptrglElementPointerATI == NULL) return 1;
// 	ptrglDrawElementArrayATI = goglGetProcAddress("glDrawElementArrayATI");
// 	if(ptrglDrawElementArrayATI == NULL) return 1;
// 	ptrglDrawRangeElementArrayATI = goglGetProcAddress("glDrawRangeElementArrayATI");
// 	if(ptrglDrawRangeElementArrayATI == NULL) return 1;
// 	return 0;
// }
// int init_ATI_envmap_bumpmap() {
// 	ptrglTexBumpParameterivATI = goglGetProcAddress("glTexBumpParameterivATI");
// 	if(ptrglTexBumpParameterivATI == NULL) return 1;
// 	ptrglTexBumpParameterfvATI = goglGetProcAddress("glTexBumpParameterfvATI");
// 	if(ptrglTexBumpParameterfvATI == NULL) return 1;
// 	ptrglGetTexBumpParameterivATI = goglGetProcAddress("glGetTexBumpParameterivATI");
// 	if(ptrglGetTexBumpParameterivATI == NULL) return 1;
// 	ptrglGetTexBumpParameterfvATI = goglGetProcAddress("glGetTexBumpParameterfvATI");
// 	if(ptrglGetTexBumpParameterfvATI == NULL) return 1;
// 	return 0;
// }
// int init_ATI_fragment_shader() {
// 	ptrglGenFragmentShadersATI = goglGetProcAddress("glGenFragmentShadersATI");
// 	if(ptrglGenFragmentShadersATI == NULL) return 1;
// 	ptrglBindFragmentShaderATI = goglGetProcAddress("glBindFragmentShaderATI");
// 	if(ptrglBindFragmentShaderATI == NULL) return 1;
// 	ptrglDeleteFragmentShaderATI = goglGetProcAddress("glDeleteFragmentShaderATI");
// 	if(ptrglDeleteFragmentShaderATI == NULL) return 1;
// 	ptrglBeginFragmentShaderATI = goglGetProcAddress("glBeginFragmentShaderATI");
// 	if(ptrglBeginFragmentShaderATI == NULL) return 1;
// 	ptrglEndFragmentShaderATI = goglGetProcAddress("glEndFragmentShaderATI");
// 	if(ptrglEndFragmentShaderATI == NULL) return 1;
// 	ptrglPassTexCoordATI = goglGetProcAddress("glPassTexCoordATI");
// 	if(ptrglPassTexCoordATI == NULL) return 1;
// 	ptrglSampleMapATI = goglGetProcAddress("glSampleMapATI");
// 	if(ptrglSampleMapATI == NULL) return 1;
// 	ptrglColorFragmentOp1ATI = goglGetProcAddress("glColorFragmentOp1ATI");
// 	if(ptrglColorFragmentOp1ATI == NULL) return 1;
// 	ptrglColorFragmentOp2ATI = goglGetProcAddress("glColorFragmentOp2ATI");
// 	if(ptrglColorFragmentOp2ATI == NULL) return 1;
// 	ptrglColorFragmentOp3ATI = goglGetProcAddress("glColorFragmentOp3ATI");
// 	if(ptrglColorFragmentOp3ATI == NULL) return 1;
// 	ptrglAlphaFragmentOp1ATI = goglGetProcAddress("glAlphaFragmentOp1ATI");
// 	if(ptrglAlphaFragmentOp1ATI == NULL) return 1;
// 	ptrglAlphaFragmentOp2ATI = goglGetProcAddress("glAlphaFragmentOp2ATI");
// 	if(ptrglAlphaFragmentOp2ATI == NULL) return 1;
// 	ptrglAlphaFragmentOp3ATI = goglGetProcAddress("glAlphaFragmentOp3ATI");
// 	if(ptrglAlphaFragmentOp3ATI == NULL) return 1;
// 	ptrglSetFragmentShaderConstantATI = goglGetProcAddress("glSetFragmentShaderConstantATI");
// 	if(ptrglSetFragmentShaderConstantATI == NULL) return 1;
// 	return 0;
// }
// int init_ATI_map_object_buffer() {
// 	ptrglMapObjectBufferATI = goglGetProcAddress("glMapObjectBufferATI");
// 	if(ptrglMapObjectBufferATI == NULL) return 1;
// 	ptrglUnmapObjectBufferATI = goglGetProcAddress("glUnmapObjectBufferATI");
// 	if(ptrglUnmapObjectBufferATI == NULL) return 1;
// 	return 0;
// }
// int init_ATI_pn_triangles() {
// 	ptrglPNTrianglesiATI = goglGetProcAddress("glPNTrianglesiATI");
// 	if(ptrglPNTrianglesiATI == NULL) return 1;
// 	ptrglPNTrianglesfATI = goglGetProcAddress("glPNTrianglesfATI");
// 	if(ptrglPNTrianglesfATI == NULL) return 1;
// 	return 0;
// }
// int init_ATI_separate_stencil() {
// 	ptrglStencilOpSeparateATI = goglGetProcAddress("glStencilOpSeparateATI");
// 	if(ptrglStencilOpSeparateATI == NULL) return 1;
// 	ptrglStencilFuncSeparateATI = goglGetProcAddress("glStencilFuncSeparateATI");
// 	if(ptrglStencilFuncSeparateATI == NULL) return 1;
// 	return 0;
// }
// int init_ATI_vertex_array_object() {
// 	ptrglNewObjectBufferATI = goglGetProcAddress("glNewObjectBufferATI");
// 	if(ptrglNewObjectBufferATI == NULL) return 1;
// 	ptrglIsObjectBufferATI = goglGetProcAddress("glIsObjectBufferATI");
// 	if(ptrglIsObjectBufferATI == NULL) return 1;
// 	ptrglUpdateObjectBufferATI = goglGetProcAddress("glUpdateObjectBufferATI");
// 	if(ptrglUpdateObjectBufferATI == NULL) return 1;
// 	ptrglGetObjectBufferfvATI = goglGetProcAddress("glGetObjectBufferfvATI");
// 	if(ptrglGetObjectBufferfvATI == NULL) return 1;
// 	ptrglGetObjectBufferivATI = goglGetProcAddress("glGetObjectBufferivATI");
// 	if(ptrglGetObjectBufferivATI == NULL) return 1;
// 	ptrglFreeObjectBufferATI = goglGetProcAddress("glFreeObjectBufferATI");
// 	if(ptrglFreeObjectBufferATI == NULL) return 1;
// 	ptrglArrayObjectATI = goglGetProcAddress("glArrayObjectATI");
// 	if(ptrglArrayObjectATI == NULL) return 1;
// 	ptrglGetArrayObjectfvATI = goglGetProcAddress("glGetArrayObjectfvATI");
// 	if(ptrglGetArrayObjectfvATI == NULL) return 1;
// 	ptrglGetArrayObjectivATI = goglGetProcAddress("glGetArrayObjectivATI");
// 	if(ptrglGetArrayObjectivATI == NULL) return 1;
// 	ptrglVariantArrayObjectATI = goglGetProcAddress("glVariantArrayObjectATI");
// 	if(ptrglVariantArrayObjectATI == NULL) return 1;
// 	ptrglGetVariantArrayObjectfvATI = goglGetProcAddress("glGetVariantArrayObjectfvATI");
// 	if(ptrglGetVariantArrayObjectfvATI == NULL) return 1;
// 	ptrglGetVariantArrayObjectivATI = goglGetProcAddress("glGetVariantArrayObjectivATI");
// 	if(ptrglGetVariantArrayObjectivATI == NULL) return 1;
// 	return 0;
// }
// int init_ATI_vertex_attrib_array_object() {
// 	ptrglVertexAttribArrayObjectATI = goglGetProcAddress("glVertexAttribArrayObjectATI");
// 	if(ptrglVertexAttribArrayObjectATI == NULL) return 1;
// 	ptrglGetVertexAttribArrayObjectfvATI = goglGetProcAddress("glGetVertexAttribArrayObjectfvATI");
// 	if(ptrglGetVertexAttribArrayObjectfvATI == NULL) return 1;
// 	ptrglGetVertexAttribArrayObjectivATI = goglGetProcAddress("glGetVertexAttribArrayObjectivATI");
// 	if(ptrglGetVertexAttribArrayObjectivATI == NULL) return 1;
// 	return 0;
// }
// int init_ATI_vertex_streams() {
// 	ptrglVertexStream1sATI = goglGetProcAddress("glVertexStream1sATI");
// 	if(ptrglVertexStream1sATI == NULL) return 1;
// 	ptrglVertexStream1svATI = goglGetProcAddress("glVertexStream1svATI");
// 	if(ptrglVertexStream1svATI == NULL) return 1;
// 	ptrglVertexStream1iATI = goglGetProcAddress("glVertexStream1iATI");
// 	if(ptrglVertexStream1iATI == NULL) return 1;
// 	ptrglVertexStream1ivATI = goglGetProcAddress("glVertexStream1ivATI");
// 	if(ptrglVertexStream1ivATI == NULL) return 1;
// 	ptrglVertexStream1fATI = goglGetProcAddress("glVertexStream1fATI");
// 	if(ptrglVertexStream1fATI == NULL) return 1;
// 	ptrglVertexStream1fvATI = goglGetProcAddress("glVertexStream1fvATI");
// 	if(ptrglVertexStream1fvATI == NULL) return 1;
// 	ptrglVertexStream1dATI = goglGetProcAddress("glVertexStream1dATI");
// 	if(ptrglVertexStream1dATI == NULL) return 1;
// 	ptrglVertexStream1dvATI = goglGetProcAddress("glVertexStream1dvATI");
// 	if(ptrglVertexStream1dvATI == NULL) return 1;
// 	ptrglVertexStream2sATI = goglGetProcAddress("glVertexStream2sATI");
// 	if(ptrglVertexStream2sATI == NULL) return 1;
// 	ptrglVertexStream2svATI = goglGetProcAddress("glVertexStream2svATI");
// 	if(ptrglVertexStream2svATI == NULL) return 1;
// 	ptrglVertexStream2iATI = goglGetProcAddress("glVertexStream2iATI");
// 	if(ptrglVertexStream2iATI == NULL) return 1;
// 	ptrglVertexStream2ivATI = goglGetProcAddress("glVertexStream2ivATI");
// 	if(ptrglVertexStream2ivATI == NULL) return 1;
// 	ptrglVertexStream2fATI = goglGetProcAddress("glVertexStream2fATI");
// 	if(ptrglVertexStream2fATI == NULL) return 1;
// 	ptrglVertexStream2fvATI = goglGetProcAddress("glVertexStream2fvATI");
// 	if(ptrglVertexStream2fvATI == NULL) return 1;
// 	ptrglVertexStream2dATI = goglGetProcAddress("glVertexStream2dATI");
// 	if(ptrglVertexStream2dATI == NULL) return 1;
// 	ptrglVertexStream2dvATI = goglGetProcAddress("glVertexStream2dvATI");
// 	if(ptrglVertexStream2dvATI == NULL) return 1;
// 	ptrglVertexStream3sATI = goglGetProcAddress("glVertexStream3sATI");
// 	if(ptrglVertexStream3sATI == NULL) return 1;
// 	ptrglVertexStream3svATI = goglGetProcAddress("glVertexStream3svATI");
// 	if(ptrglVertexStream3svATI == NULL) return 1;
// 	ptrglVertexStream3iATI = goglGetProcAddress("glVertexStream3iATI");
// 	if(ptrglVertexStream3iATI == NULL) return 1;
// 	ptrglVertexStream3ivATI = goglGetProcAddress("glVertexStream3ivATI");
// 	if(ptrglVertexStream3ivATI == NULL) return 1;
// 	ptrglVertexStream3fATI = goglGetProcAddress("glVertexStream3fATI");
// 	if(ptrglVertexStream3fATI == NULL) return 1;
// 	ptrglVertexStream3fvATI = goglGetProcAddress("glVertexStream3fvATI");
// 	if(ptrglVertexStream3fvATI == NULL) return 1;
// 	ptrglVertexStream3dATI = goglGetProcAddress("glVertexStream3dATI");
// 	if(ptrglVertexStream3dATI == NULL) return 1;
// 	ptrglVertexStream3dvATI = goglGetProcAddress("glVertexStream3dvATI");
// 	if(ptrglVertexStream3dvATI == NULL) return 1;
// 	ptrglVertexStream4sATI = goglGetProcAddress("glVertexStream4sATI");
// 	if(ptrglVertexStream4sATI == NULL) return 1;
// 	ptrglVertexStream4svATI = goglGetProcAddress("glVertexStream4svATI");
// 	if(ptrglVertexStream4svATI == NULL) return 1;
// 	ptrglVertexStream4iATI = goglGetProcAddress("glVertexStream4iATI");
// 	if(ptrglVertexStream4iATI == NULL) return 1;
// 	ptrglVertexStream4ivATI = goglGetProcAddress("glVertexStream4ivATI");
// 	if(ptrglVertexStream4ivATI == NULL) return 1;
// 	ptrglVertexStream4fATI = goglGetProcAddress("glVertexStream4fATI");
// 	if(ptrglVertexStream4fATI == NULL) return 1;
// 	ptrglVertexStream4fvATI = goglGetProcAddress("glVertexStream4fvATI");
// 	if(ptrglVertexStream4fvATI == NULL) return 1;
// 	ptrglVertexStream4dATI = goglGetProcAddress("glVertexStream4dATI");
// 	if(ptrglVertexStream4dATI == NULL) return 1;
// 	ptrglVertexStream4dvATI = goglGetProcAddress("glVertexStream4dvATI");
// 	if(ptrglVertexStream4dvATI == NULL) return 1;
// 	ptrglNormalStream3bATI = goglGetProcAddress("glNormalStream3bATI");
// 	if(ptrglNormalStream3bATI == NULL) return 1;
// 	ptrglNormalStream3bvATI = goglGetProcAddress("glNormalStream3bvATI");
// 	if(ptrglNormalStream3bvATI == NULL) return 1;
// 	ptrglNormalStream3sATI = goglGetProcAddress("glNormalStream3sATI");
// 	if(ptrglNormalStream3sATI == NULL) return 1;
// 	ptrglNormalStream3svATI = goglGetProcAddress("glNormalStream3svATI");
// 	if(ptrglNormalStream3svATI == NULL) return 1;
// 	ptrglNormalStream3iATI = goglGetProcAddress("glNormalStream3iATI");
// 	if(ptrglNormalStream3iATI == NULL) return 1;
// 	ptrglNormalStream3ivATI = goglGetProcAddress("glNormalStream3ivATI");
// 	if(ptrglNormalStream3ivATI == NULL) return 1;
// 	ptrglNormalStream3fATI = goglGetProcAddress("glNormalStream3fATI");
// 	if(ptrglNormalStream3fATI == NULL) return 1;
// 	ptrglNormalStream3fvATI = goglGetProcAddress("glNormalStream3fvATI");
// 	if(ptrglNormalStream3fvATI == NULL) return 1;
// 	ptrglNormalStream3dATI = goglGetProcAddress("glNormalStream3dATI");
// 	if(ptrglNormalStream3dATI == NULL) return 1;
// 	ptrglNormalStream3dvATI = goglGetProcAddress("glNormalStream3dvATI");
// 	if(ptrglNormalStream3dvATI == NULL) return 1;
// 	ptrglClientActiveVertexStreamATI = goglGetProcAddress("glClientActiveVertexStreamATI");
// 	if(ptrglClientActiveVertexStreamATI == NULL) return 1;
// 	ptrglVertexBlendEnviATI = goglGetProcAddress("glVertexBlendEnviATI");
// 	if(ptrglVertexBlendEnviATI == NULL) return 1;
// 	ptrglVertexBlendEnvfATI = goglGetProcAddress("glVertexBlendEnvfATI");
// 	if(ptrglVertexBlendEnvfATI == NULL) return 1;
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

// ATI_draw_buffers
const (
	DRAW_BUFFER0_ATI = 0x8825
	DRAW_BUFFER10_ATI = 0x882F
	DRAW_BUFFER11_ATI = 0x8830
	DRAW_BUFFER12_ATI = 0x8831
	DRAW_BUFFER13_ATI = 0x8832
	DRAW_BUFFER14_ATI = 0x8833
	DRAW_BUFFER15_ATI = 0x8834
	DRAW_BUFFER1_ATI = 0x8826
	DRAW_BUFFER2_ATI = 0x8827
	DRAW_BUFFER3_ATI = 0x8828
	DRAW_BUFFER4_ATI = 0x8829
	DRAW_BUFFER5_ATI = 0x882A
	DRAW_BUFFER6_ATI = 0x882B
	DRAW_BUFFER7_ATI = 0x882C
	DRAW_BUFFER8_ATI = 0x882D
	DRAW_BUFFER9_ATI = 0x882E
	MAX_DRAW_BUFFERS_ATI = 0x8824
)
// ATI_element_array
const (
	ELEMENT_ARRAY_ATI = 0x8768
	ELEMENT_ARRAY_POINTER_ATI = 0x876A
	ELEMENT_ARRAY_TYPE_ATI = 0x8769
)
// ATI_envmap_bumpmap
const (
	BUMP_ENVMAP_ATI = 0x877B
	BUMP_NUM_TEX_UNITS_ATI = 0x8777
	BUMP_ROT_MATRIX_ATI = 0x8775
	BUMP_ROT_MATRIX_SIZE_ATI = 0x8776
	BUMP_TARGET_ATI = 0x877C
	BUMP_TEX_UNITS_ATI = 0x8778
	DU8DV8_ATI = 0x877A
	DUDV_ATI = 0x8779
)
// ATI_fragment_shader
const (
	X2X_BIT_ATI = 0x00000001
	X4X_BIT_ATI = 0x00000002
	X8X_BIT_ATI = 0x00000004
	ADD_ATI = 0x8963
	BIAS_BIT_ATI = 0x00000008
	BLUE_BIT_ATI = 0x00000004
	CND0_ATI = 0x896B
	CND_ATI = 0x896A
	COLOR_ALPHA_PAIRING_ATI = 0x8975
	COMP_BIT_ATI = 0x00000002
	CON_0_ATI = 0x8941
	CON_10_ATI = 0x894B
	CON_11_ATI = 0x894C
	CON_12_ATI = 0x894D
	CON_13_ATI = 0x894E
	CON_14_ATI = 0x894F
	CON_15_ATI = 0x8950
	CON_16_ATI = 0x8951
	CON_17_ATI = 0x8952
	CON_18_ATI = 0x8953
	CON_19_ATI = 0x8954
	CON_1_ATI = 0x8942
	CON_20_ATI = 0x8955
	CON_21_ATI = 0x8956
	CON_22_ATI = 0x8957
	CON_23_ATI = 0x8958
	CON_24_ATI = 0x8959
	CON_25_ATI = 0x895A
	CON_26_ATI = 0x895B
	CON_27_ATI = 0x895C
	CON_28_ATI = 0x895D
	CON_29_ATI = 0x895E
	CON_2_ATI = 0x8943
	CON_30_ATI = 0x895F
	CON_31_ATI = 0x8960
	CON_3_ATI = 0x8944
	CON_4_ATI = 0x8945
	CON_5_ATI = 0x8946
	CON_6_ATI = 0x8947
	CON_7_ATI = 0x8948
	CON_8_ATI = 0x8949
	CON_9_ATI = 0x894A
	DOT2_ADD_ATI = 0x896C
	DOT3_ATI = 0x8966
	DOT4_ATI = 0x8967
	EIGHTH_BIT_ATI = 0x00000020
	FRAGMENT_SHADER_ATI = 0x8920
	GREEN_BIT_ATI = 0x00000002
	HALF_BIT_ATI = 0x00000008
	LERP_ATI = 0x8969
	MAD_ATI = 0x8968
	MOV_ATI = 0x8961
	MUL_ATI = 0x8964
	NEGATE_BIT_ATI = 0x00000004
	NUM_FRAGMENT_CONSTANTS_ATI = 0x896F
	NUM_FRAGMENT_REGISTERS_ATI = 0x896E
	NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI = 0x8973
	NUM_INSTRUCTIONS_PER_PASS_ATI = 0x8971
	NUM_INSTRUCTIONS_TOTAL_ATI = 0x8972
	NUM_LOOPBACK_COMPONENTS_ATI = 0x8974
	NUM_PASSES_ATI = 0x8970
	QUARTER_BIT_ATI = 0x00000010
	RED_BIT_ATI = 0x00000001
	REG_0_ATI = 0x8921
	REG_10_ATI = 0x892B
	REG_11_ATI = 0x892C
	REG_12_ATI = 0x892D
	REG_13_ATI = 0x892E
	REG_14_ATI = 0x892F
	REG_15_ATI = 0x8930
	REG_16_ATI = 0x8931
	REG_17_ATI = 0x8932
	REG_18_ATI = 0x8933
	REG_19_ATI = 0x8934
	REG_1_ATI = 0x8922
	REG_20_ATI = 0x8935
	REG_21_ATI = 0x8936
	REG_22_ATI = 0x8937
	REG_23_ATI = 0x8938
	REG_24_ATI = 0x8939
	REG_25_ATI = 0x893A
	REG_26_ATI = 0x893B
	REG_27_ATI = 0x893C
	REG_28_ATI = 0x893D
	REG_29_ATI = 0x893E
	REG_2_ATI = 0x8923
	REG_30_ATI = 0x893F
	REG_31_ATI = 0x8940
	REG_3_ATI = 0x8924
	REG_4_ATI = 0x8925
	REG_5_ATI = 0x8926
	REG_6_ATI = 0x8927
	REG_7_ATI = 0x8928
	REG_8_ATI = 0x8929
	REG_9_ATI = 0x892A
	SATURATE_BIT_ATI = 0x00000040
	SECONDARY_INTERPOLATOR_ATI = 0x896D
	SUB_ATI = 0x8965
	SWIZZLE_STQ_ATI = 0x8977
	SWIZZLE_STQ_DQ_ATI = 0x8979
	SWIZZLE_STRQ_ATI = 0x897A
	SWIZZLE_STRQ_DQ_ATI = 0x897B
	SWIZZLE_STR_ATI = 0x8976
	SWIZZLE_STR_DR_ATI = 0x8978
)
// ATI_map_object_buffer
const (
)
// ATI_meminfo
const (
	RENDERBUFFER_FREE_MEMORY_ATI = 0x87FD
	TEXTURE_FREE_MEMORY_ATI = 0x87FC
	VBO_FREE_MEMORY_ATI = 0x87FB
)
// ATI_pixel_format_float
const (
	COLOR_CLEAR_UNCLAMPED_VALUE_ATI = 0x8835
	RGBA_FLOAT_MODE_ATI = 0x8820
)
// ATI_pn_triangles
const (
	MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI = 0x87F1
	PN_TRIANGLES_ATI = 0x87F0
	PN_TRIANGLES_NORMAL_MODE_ATI = 0x87F3
	PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI = 0x87F7
	PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI = 0x87F8
	PN_TRIANGLES_POINT_MODE_ATI = 0x87F2
	PN_TRIANGLES_POINT_MODE_CUBIC_ATI = 0x87F6
	PN_TRIANGLES_POINT_MODE_LINEAR_ATI = 0x87F5
	PN_TRIANGLES_TESSELATION_LEVEL_ATI = 0x87F4
)
// ATI_separate_stencil
const (
	STENCIL_BACK_FAIL_ATI = 0x8801
	STENCIL_BACK_FUNC_ATI = 0x8800
	STENCIL_BACK_PASS_DEPTH_FAIL_ATI = 0x8802
	STENCIL_BACK_PASS_DEPTH_PASS_ATI = 0x8803
)
// ATI_text_fragment_shader
const (
	TEXT_FRAGMENT_SHADER_ATI = 0x8200
)
// ATI_texture_env_combine3
const (
	MODULATE_ADD_ATI = 0x8744
	MODULATE_SIGNED_ADD_ATI = 0x8745
	MODULATE_SUBTRACT_ATI = 0x8746
)
// ATI_texture_float
const (
	ALPHA_FLOAT16_ATI = 0x881C
	ALPHA_FLOAT32_ATI = 0x8816
	INTENSITY_FLOAT16_ATI = 0x881D
	INTENSITY_FLOAT32_ATI = 0x8817
	LUMINANCE_ALPHA_FLOAT16_ATI = 0x881F
	LUMINANCE_ALPHA_FLOAT32_ATI = 0x8819
	LUMINANCE_FLOAT16_ATI = 0x881E
	LUMINANCE_FLOAT32_ATI = 0x8818
	RGBA_FLOAT16_ATI = 0x881A
	RGBA_FLOAT32_ATI = 0x8814
	RGB_FLOAT16_ATI = 0x881B
	RGB_FLOAT32_ATI = 0x8815
)
// ATI_texture_mirror_once
const (
	MIRROR_CLAMP_ATI = 0x8742
	MIRROR_CLAMP_TO_EDGE_ATI = 0x8743
)
// ATI_vertex_array_object
const (
	ARRAY_OBJECT_BUFFER_ATI = 0x8766
	ARRAY_OBJECT_OFFSET_ATI = 0x8767
	DISCARD_ATI = 0x8763
	DYNAMIC_ATI = 0x8761
	OBJECT_BUFFER_SIZE_ATI = 0x8764
	OBJECT_BUFFER_USAGE_ATI = 0x8765
	PRESERVE_ATI = 0x8762
	STATIC_ATI = 0x8760
)
// ATI_vertex_attrib_array_object
const (
)
// ATI_vertex_streams
const (
	MAX_VERTEX_STREAMS_ATI = 0x876B
	VERTEX_SOURCE_ATI = 0x8774
	VERTEX_STREAM0_ATI = 0x876C
	VERTEX_STREAM1_ATI = 0x876D
	VERTEX_STREAM2_ATI = 0x876E
	VERTEX_STREAM3_ATI = 0x876F
	VERTEX_STREAM4_ATI = 0x8770
	VERTEX_STREAM5_ATI = 0x8771
	VERTEX_STREAM6_ATI = 0x8772
	VERTEX_STREAM7_ATI = 0x8773
)
// ATI_draw_buffers

func DrawBuffersATI(n Sizei, bufs *Enum)  {
	C.goglDrawBuffersATI((C.GLsizei)(n), (*C.GLenum)(bufs))
}
// ATI_element_array

func ElementPointerATI(type_ Enum, pointer Pointer)  {
	C.goglElementPointerATI((C.GLenum)(type_), (unsafe.Pointer)(pointer))
}
func DrawElementArrayATI(mode Enum, count Sizei)  {
	C.goglDrawElementArrayATI((C.GLenum)(mode), (C.GLsizei)(count))
}
func DrawRangeElementArrayATI(mode Enum, start Uint, end Uint, count Sizei)  {
	C.goglDrawRangeElementArrayATI((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count))
}
// ATI_envmap_bumpmap

func TexBumpParameterivATI(pname Enum, param *Int)  {
	C.goglTexBumpParameterivATI((C.GLenum)(pname), (*C.GLint)(param))
}
func TexBumpParameterfvATI(pname Enum, param *Float)  {
	C.goglTexBumpParameterfvATI((C.GLenum)(pname), (*C.GLfloat)(param))
}
func GetTexBumpParameterivATI(pname Enum, param *Int)  {
	C.goglGetTexBumpParameterivATI((C.GLenum)(pname), (*C.GLint)(param))
}
func GetTexBumpParameterfvATI(pname Enum, param *Float)  {
	C.goglGetTexBumpParameterfvATI((C.GLenum)(pname), (*C.GLfloat)(param))
}
// ATI_fragment_shader

func GenFragmentShadersATI(range_ Uint) Uint {
	return (Uint)(C.goglGenFragmentShadersATI((C.GLuint)(range_)))
}
func BindFragmentShaderATI(id Uint)  {
	C.goglBindFragmentShaderATI((C.GLuint)(id))
}
func DeleteFragmentShaderATI(id Uint)  {
	C.goglDeleteFragmentShaderATI((C.GLuint)(id))
}
func BeginFragmentShaderATI()  {
	C.goglBeginFragmentShaderATI()
}
func EndFragmentShaderATI()  {
	C.goglEndFragmentShaderATI()
}
func PassTexCoordATI(dst Uint, coord Uint, swizzle Enum)  {
	C.goglPassTexCoordATI((C.GLuint)(dst), (C.GLuint)(coord), (C.GLenum)(swizzle))
}
func SampleMapATI(dst Uint, interp Uint, swizzle Enum)  {
	C.goglSampleMapATI((C.GLuint)(dst), (C.GLuint)(interp), (C.GLenum)(swizzle))
}
func ColorFragmentOp1ATI(op Enum, dst Uint, dstMask Uint, dstMod Uint, arg1 Uint, arg1Rep Uint, arg1Mod Uint)  {
	C.goglColorFragmentOp1ATI((C.GLenum)(op), (C.GLuint)(dst), (C.GLuint)(dstMask), (C.GLuint)(dstMod), (C.GLuint)(arg1), (C.GLuint)(arg1Rep), (C.GLuint)(arg1Mod))
}
func ColorFragmentOp2ATI(op Enum, dst Uint, dstMask Uint, dstMod Uint, arg1 Uint, arg1Rep Uint, arg1Mod Uint, arg2 Uint, arg2Rep Uint, arg2Mod Uint)  {
	C.goglColorFragmentOp2ATI((C.GLenum)(op), (C.GLuint)(dst), (C.GLuint)(dstMask), (C.GLuint)(dstMod), (C.GLuint)(arg1), (C.GLuint)(arg1Rep), (C.GLuint)(arg1Mod), (C.GLuint)(arg2), (C.GLuint)(arg2Rep), (C.GLuint)(arg2Mod))
}
func ColorFragmentOp3ATI(op Enum, dst Uint, dstMask Uint, dstMod Uint, arg1 Uint, arg1Rep Uint, arg1Mod Uint, arg2 Uint, arg2Rep Uint, arg2Mod Uint, arg3 Uint, arg3Rep Uint, arg3Mod Uint)  {
	C.goglColorFragmentOp3ATI((C.GLenum)(op), (C.GLuint)(dst), (C.GLuint)(dstMask), (C.GLuint)(dstMod), (C.GLuint)(arg1), (C.GLuint)(arg1Rep), (C.GLuint)(arg1Mod), (C.GLuint)(arg2), (C.GLuint)(arg2Rep), (C.GLuint)(arg2Mod), (C.GLuint)(arg3), (C.GLuint)(arg3Rep), (C.GLuint)(arg3Mod))
}
func AlphaFragmentOp1ATI(op Enum, dst Uint, dstMod Uint, arg1 Uint, arg1Rep Uint, arg1Mod Uint)  {
	C.goglAlphaFragmentOp1ATI((C.GLenum)(op), (C.GLuint)(dst), (C.GLuint)(dstMod), (C.GLuint)(arg1), (C.GLuint)(arg1Rep), (C.GLuint)(arg1Mod))
}
func AlphaFragmentOp2ATI(op Enum, dst Uint, dstMod Uint, arg1 Uint, arg1Rep Uint, arg1Mod Uint, arg2 Uint, arg2Rep Uint, arg2Mod Uint)  {
	C.goglAlphaFragmentOp2ATI((C.GLenum)(op), (C.GLuint)(dst), (C.GLuint)(dstMod), (C.GLuint)(arg1), (C.GLuint)(arg1Rep), (C.GLuint)(arg1Mod), (C.GLuint)(arg2), (C.GLuint)(arg2Rep), (C.GLuint)(arg2Mod))
}
func AlphaFragmentOp3ATI(op Enum, dst Uint, dstMod Uint, arg1 Uint, arg1Rep Uint, arg1Mod Uint, arg2 Uint, arg2Rep Uint, arg2Mod Uint, arg3 Uint, arg3Rep Uint, arg3Mod Uint)  {
	C.goglAlphaFragmentOp3ATI((C.GLenum)(op), (C.GLuint)(dst), (C.GLuint)(dstMod), (C.GLuint)(arg1), (C.GLuint)(arg1Rep), (C.GLuint)(arg1Mod), (C.GLuint)(arg2), (C.GLuint)(arg2Rep), (C.GLuint)(arg2Mod), (C.GLuint)(arg3), (C.GLuint)(arg3Rep), (C.GLuint)(arg3Mod))
}
func SetFragmentShaderConstantATI(dst Uint, value *Float)  {
	C.goglSetFragmentShaderConstantATI((C.GLuint)(dst), (*C.GLfloat)(value))
}
// ATI_map_object_buffer

func MapObjectBufferATI(buffer Uint) Pointer {
	return (Pointer)(C.goglMapObjectBufferATI((C.GLuint)(buffer)))
}
func UnmapObjectBufferATI(buffer Uint)  {
	C.goglUnmapObjectBufferATI((C.GLuint)(buffer))
}
// ATI_pn_triangles

func PNTrianglesiATI(pname Enum, param Int)  {
	C.goglPNTrianglesiATI((C.GLenum)(pname), (C.GLint)(param))
}
func PNTrianglesfATI(pname Enum, param Float)  {
	C.goglPNTrianglesfATI((C.GLenum)(pname), (C.GLfloat)(param))
}
// ATI_separate_stencil

func StencilOpSeparateATI(face Enum, sfail Enum, dpfail Enum, dppass Enum)  {
	C.goglStencilOpSeparateATI((C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}
func StencilFuncSeparateATI(frontfunc Enum, backfunc Enum, ref Int, mask Uint)  {
	C.goglStencilFuncSeparateATI((C.GLenum)(frontfunc), (C.GLenum)(backfunc), (C.GLint)(ref), (C.GLuint)(mask))
}
// ATI_vertex_array_object

func NewObjectBufferATI(size Sizei, pointer Pointer, usage Enum) Uint {
	return (Uint)(C.goglNewObjectBufferATI((C.GLsizei)(size), (unsafe.Pointer)(pointer), (C.GLenum)(usage)))
}
func IsObjectBufferATI(buffer Uint) Boolean {
	return (Boolean)(C.goglIsObjectBufferATI((C.GLuint)(buffer)))
}
func UpdateObjectBufferATI(buffer Uint, offset Uint, size Sizei, pointer Pointer, preserve Enum)  {
	C.goglUpdateObjectBufferATI((C.GLuint)(buffer), (C.GLuint)(offset), (C.GLsizei)(size), (unsafe.Pointer)(pointer), (C.GLenum)(preserve))
}
func GetObjectBufferfvATI(buffer Uint, pname Enum, params *Float)  {
	C.goglGetObjectBufferfvATI((C.GLuint)(buffer), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetObjectBufferivATI(buffer Uint, pname Enum, params *Int)  {
	C.goglGetObjectBufferivATI((C.GLuint)(buffer), (C.GLenum)(pname), (*C.GLint)(params))
}
func FreeObjectBufferATI(buffer Uint)  {
	C.goglFreeObjectBufferATI((C.GLuint)(buffer))
}
func ArrayObjectATI(array Enum, size Int, type_ Enum, stride Sizei, buffer Uint, offset Uint)  {
	C.goglArrayObjectATI((C.GLenum)(array), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (C.GLuint)(buffer), (C.GLuint)(offset))
}
func GetArrayObjectfvATI(array Enum, pname Enum, params *Float)  {
	C.goglGetArrayObjectfvATI((C.GLenum)(array), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetArrayObjectivATI(array Enum, pname Enum, params *Int)  {
	C.goglGetArrayObjectivATI((C.GLenum)(array), (C.GLenum)(pname), (*C.GLint)(params))
}
func VariantArrayObjectATI(id Uint, type_ Enum, stride Sizei, buffer Uint, offset Uint)  {
	C.goglVariantArrayObjectATI((C.GLuint)(id), (C.GLenum)(type_), (C.GLsizei)(stride), (C.GLuint)(buffer), (C.GLuint)(offset))
}
func GetVariantArrayObjectfvATI(id Uint, pname Enum, params *Float)  {
	C.goglGetVariantArrayObjectfvATI((C.GLuint)(id), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetVariantArrayObjectivATI(id Uint, pname Enum, params *Int)  {
	C.goglGetVariantArrayObjectivATI((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
}
// ATI_vertex_attrib_array_object

func VertexAttribArrayObjectATI(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, buffer Uint, offset Uint)  {
	C.goglVertexAttribArrayObjectATI((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLsizei)(stride), (C.GLuint)(buffer), (C.GLuint)(offset))
}
func GetVertexAttribArrayObjectfvATI(index Uint, pname Enum, params *Float)  {
	C.goglGetVertexAttribArrayObjectfvATI((C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetVertexAttribArrayObjectivATI(index Uint, pname Enum, params *Int)  {
	C.goglGetVertexAttribArrayObjectivATI((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}
// ATI_vertex_streams

func VertexStream1sATI(stream Enum, x Short)  {
	C.goglVertexStream1sATI((C.GLenum)(stream), (C.GLshort)(x))
}
func VertexStream1svATI(stream Enum, coords *Short)  {
	C.goglVertexStream1svATI((C.GLenum)(stream), (*C.GLshort)(coords))
}
func VertexStream1iATI(stream Enum, x Int)  {
	C.goglVertexStream1iATI((C.GLenum)(stream), (C.GLint)(x))
}
func VertexStream1ivATI(stream Enum, coords *Int)  {
	C.goglVertexStream1ivATI((C.GLenum)(stream), (*C.GLint)(coords))
}
func VertexStream1fATI(stream Enum, x Float)  {
	C.goglVertexStream1fATI((C.GLenum)(stream), (C.GLfloat)(x))
}
func VertexStream1fvATI(stream Enum, coords *Float)  {
	C.goglVertexStream1fvATI((C.GLenum)(stream), (*C.GLfloat)(coords))
}
func VertexStream1dATI(stream Enum, x Double)  {
	C.goglVertexStream1dATI((C.GLenum)(stream), (C.GLdouble)(x))
}
func VertexStream1dvATI(stream Enum, coords *Double)  {
	C.goglVertexStream1dvATI((C.GLenum)(stream), (*C.GLdouble)(coords))
}
func VertexStream2sATI(stream Enum, x Short, y Short)  {
	C.goglVertexStream2sATI((C.GLenum)(stream), (C.GLshort)(x), (C.GLshort)(y))
}
func VertexStream2svATI(stream Enum, coords *Short)  {
	C.goglVertexStream2svATI((C.GLenum)(stream), (*C.GLshort)(coords))
}
func VertexStream2iATI(stream Enum, x Int, y Int)  {
	C.goglVertexStream2iATI((C.GLenum)(stream), (C.GLint)(x), (C.GLint)(y))
}
func VertexStream2ivATI(stream Enum, coords *Int)  {
	C.goglVertexStream2ivATI((C.GLenum)(stream), (*C.GLint)(coords))
}
func VertexStream2fATI(stream Enum, x Float, y Float)  {
	C.goglVertexStream2fATI((C.GLenum)(stream), (C.GLfloat)(x), (C.GLfloat)(y))
}
func VertexStream2fvATI(stream Enum, coords *Float)  {
	C.goglVertexStream2fvATI((C.GLenum)(stream), (*C.GLfloat)(coords))
}
func VertexStream2dATI(stream Enum, x Double, y Double)  {
	C.goglVertexStream2dATI((C.GLenum)(stream), (C.GLdouble)(x), (C.GLdouble)(y))
}
func VertexStream2dvATI(stream Enum, coords *Double)  {
	C.goglVertexStream2dvATI((C.GLenum)(stream), (*C.GLdouble)(coords))
}
func VertexStream3sATI(stream Enum, x Short, y Short, z Short)  {
	C.goglVertexStream3sATI((C.GLenum)(stream), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func VertexStream3svATI(stream Enum, coords *Short)  {
	C.goglVertexStream3svATI((C.GLenum)(stream), (*C.GLshort)(coords))
}
func VertexStream3iATI(stream Enum, x Int, y Int, z Int)  {
	C.goglVertexStream3iATI((C.GLenum)(stream), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func VertexStream3ivATI(stream Enum, coords *Int)  {
	C.goglVertexStream3ivATI((C.GLenum)(stream), (*C.GLint)(coords))
}
func VertexStream3fATI(stream Enum, x Float, y Float, z Float)  {
	C.goglVertexStream3fATI((C.GLenum)(stream), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func VertexStream3fvATI(stream Enum, coords *Float)  {
	C.goglVertexStream3fvATI((C.GLenum)(stream), (*C.GLfloat)(coords))
}
func VertexStream3dATI(stream Enum, x Double, y Double, z Double)  {
	C.goglVertexStream3dATI((C.GLenum)(stream), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func VertexStream3dvATI(stream Enum, coords *Double)  {
	C.goglVertexStream3dvATI((C.GLenum)(stream), (*C.GLdouble)(coords))
}
func VertexStream4sATI(stream Enum, x Short, y Short, z Short, w Short)  {
	C.goglVertexStream4sATI((C.GLenum)(stream), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func VertexStream4svATI(stream Enum, coords *Short)  {
	C.goglVertexStream4svATI((C.GLenum)(stream), (*C.GLshort)(coords))
}
func VertexStream4iATI(stream Enum, x Int, y Int, z Int, w Int)  {
	C.goglVertexStream4iATI((C.GLenum)(stream), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func VertexStream4ivATI(stream Enum, coords *Int)  {
	C.goglVertexStream4ivATI((C.GLenum)(stream), (*C.GLint)(coords))
}
func VertexStream4fATI(stream Enum, x Float, y Float, z Float, w Float)  {
	C.goglVertexStream4fATI((C.GLenum)(stream), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func VertexStream4fvATI(stream Enum, coords *Float)  {
	C.goglVertexStream4fvATI((C.GLenum)(stream), (*C.GLfloat)(coords))
}
func VertexStream4dATI(stream Enum, x Double, y Double, z Double, w Double)  {
	C.goglVertexStream4dATI((C.GLenum)(stream), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func VertexStream4dvATI(stream Enum, coords *Double)  {
	C.goglVertexStream4dvATI((C.GLenum)(stream), (*C.GLdouble)(coords))
}
func NormalStream3bATI(stream Enum, nx Byte, ny Byte, nz Byte)  {
	C.goglNormalStream3bATI((C.GLenum)(stream), (C.GLbyte)(nx), (C.GLbyte)(ny), (C.GLbyte)(nz))
}
func NormalStream3bvATI(stream Enum, coords *Byte)  {
	C.goglNormalStream3bvATI((C.GLenum)(stream), (*C.GLbyte)(coords))
}
func NormalStream3sATI(stream Enum, nx Short, ny Short, nz Short)  {
	C.goglNormalStream3sATI((C.GLenum)(stream), (C.GLshort)(nx), (C.GLshort)(ny), (C.GLshort)(nz))
}
func NormalStream3svATI(stream Enum, coords *Short)  {
	C.goglNormalStream3svATI((C.GLenum)(stream), (*C.GLshort)(coords))
}
func NormalStream3iATI(stream Enum, nx Int, ny Int, nz Int)  {
	C.goglNormalStream3iATI((C.GLenum)(stream), (C.GLint)(nx), (C.GLint)(ny), (C.GLint)(nz))
}
func NormalStream3ivATI(stream Enum, coords *Int)  {
	C.goglNormalStream3ivATI((C.GLenum)(stream), (*C.GLint)(coords))
}
func NormalStream3fATI(stream Enum, nx Float, ny Float, nz Float)  {
	C.goglNormalStream3fATI((C.GLenum)(stream), (C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz))
}
func NormalStream3fvATI(stream Enum, coords *Float)  {
	C.goglNormalStream3fvATI((C.GLenum)(stream), (*C.GLfloat)(coords))
}
func NormalStream3dATI(stream Enum, nx Double, ny Double, nz Double)  {
	C.goglNormalStream3dATI((C.GLenum)(stream), (C.GLdouble)(nx), (C.GLdouble)(ny), (C.GLdouble)(nz))
}
func NormalStream3dvATI(stream Enum, coords *Double)  {
	C.goglNormalStream3dvATI((C.GLenum)(stream), (*C.GLdouble)(coords))
}
func ClientActiveVertexStreamATI(stream Enum)  {
	C.goglClientActiveVertexStreamATI((C.GLenum)(stream))
}
func VertexBlendEnviATI(pname Enum, param Int)  {
	C.goglVertexBlendEnviATI((C.GLenum)(pname), (C.GLint)(param))
}
func VertexBlendEnvfATI(pname Enum, param Float)  {
	C.goglVertexBlendEnvfATI((C.GLenum)(pname), (C.GLfloat)(param))
}
func InitAtiDrawBuffers() error {
	var ret C.int
	if ret = C.init_ATI_draw_buffers(); ret != 0 {
		return errors.New("unable to initialize ATI_draw_buffers")
	}
	return nil
}
func InitAtiElementArray() error {
	var ret C.int
	if ret = C.init_ATI_element_array(); ret != 0 {
		return errors.New("unable to initialize ATI_element_array")
	}
	return nil
}
func InitAtiEnvmapBumpmap() error {
	var ret C.int
	if ret = C.init_ATI_envmap_bumpmap(); ret != 0 {
		return errors.New("unable to initialize ATI_envmap_bumpmap")
	}
	return nil
}
func InitAtiFragmentShader() error {
	var ret C.int
	if ret = C.init_ATI_fragment_shader(); ret != 0 {
		return errors.New("unable to initialize ATI_fragment_shader")
	}
	return nil
}
func InitAtiMapObjectBuffer() error {
	var ret C.int
	if ret = C.init_ATI_map_object_buffer(); ret != 0 {
		return errors.New("unable to initialize ATI_map_object_buffer")
	}
	return nil
}
func InitAtiPnTriangles() error {
	var ret C.int
	if ret = C.init_ATI_pn_triangles(); ret != 0 {
		return errors.New("unable to initialize ATI_pn_triangles")
	}
	return nil
}
func InitAtiSeparateStencil() error {
	var ret C.int
	if ret = C.init_ATI_separate_stencil(); ret != 0 {
		return errors.New("unable to initialize ATI_separate_stencil")
	}
	return nil
}
func InitAtiVertexArrayObject() error {
	var ret C.int
	if ret = C.init_ATI_vertex_array_object(); ret != 0 {
		return errors.New("unable to initialize ATI_vertex_array_object")
	}
	return nil
}
func InitAtiVertexAttribArrayObject() error {
	var ret C.int
	if ret = C.init_ATI_vertex_attrib_array_object(); ret != 0 {
		return errors.New("unable to initialize ATI_vertex_attrib_array_object")
	}
	return nil
}
func InitAtiVertexStreams() error {
	var ret C.int
	if ret = C.init_ATI_vertex_streams(); ret != 0 {
		return errors.New("unable to initialize ATI_vertex_streams")
	}
	return nil
}
// EOF