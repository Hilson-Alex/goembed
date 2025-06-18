# Go Embbed Compiler

This is a change of the Go compiler created to be an embeddable part of the [Butterfly Compiler](https://github.com/Hilson-Alex/Butterfly).
The reason for this project came from the need to run it from code and without the need for it to be installed on the user's pc.

It is based on the 1.24 version of the original go compiler, with some files removed to reduce size. The target is not to be complete, but runnable and to have the minimum needed to compile a go program.
