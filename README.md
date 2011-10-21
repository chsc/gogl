GoGL 
====

GoGl is a OpenGL binding generator for Go.
It's in a very early stage of development and not functional yet.

Goals
-----

* Lightweight (use only extensions you really need)
* No external dependencies like GLEW

TODO
----

* Complete spec parser
* Linux, MacOS, Windows support
* Filter by Category, extension
* ...
	
Usage examples
--------------

Download latest spec files, generate OpenGL 3.1 interface, exclude all deprecated functions:

	gogl -download -version=3.1 -deprecation=true

Read spec files from current folder, generate OpenGL compatible interface:

	gogl -version=1.1
