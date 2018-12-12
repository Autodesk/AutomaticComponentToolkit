# Automatic Component Toolkit
The Automatic Component Toolkit (ACT) is a code generator that takes an instance of an [Interface Description Language](#interface-description-language-idl) file and generates 
a [thin C89 API](#thin-c89-api), [implementation stubs](#implementation-stubs) and [language bindings](#language-bindings) of your desired software component.

### Interface Description Language (IDL)
The IDL file defines the types and functions of your API and serves as the source for the automatically generated Code.
The exact schema of the IDL and explanation of each element is described in [Documentation/IDL.md](Documentation/IDL.md).

### Thin C89-API
A thin C89-API is a C header file that declares all functions, structs, enums and constants exported by your software component.

### Implementation Stubs
An implementation stub is a collection of source files in a certain programming language *L*
that implements
- the classes, structs, methods, ... defined by the IDL
- as well as the mapping between the thin C89 API and these native classes.
Such an implementation stub fulfills the interface and compiles by itself, however, it does not contain any domain logic.

### Language Bindings
A language binding of the component for programming language *C* implements the classes, enums, methods, ... defined by the IDL by
calling the functions exported by your component via the thin C89 API.
A consumer of your component only needs to include the language binding relevant for them and not worry about the C89 interface or the underlying implementation.

## How to use ACT:
1) Get ACT:
     1) From Source:
        1. Install go https://golang.org/doc/install
        2. Build automaticcomponenttoolkit.go:
        <br/>`.\build.bat`
    <br/>OR
     2) Download the precompiled binaries from one of the [releases](https://github.com/Autodesk/AutomaticComponentToolkit/releases) 
2) Write an interface description file for your desired API
3) Generate implementation stubs and language bindings for your API:
<br/>`act.exe Examples/Numbers/libnumbers.xml`
4) Integrate the generated code in your project

You are probably best of starting of with an [Example](#example).

## Language Support
ACT supports generation of bindings and implementation stubs for C++, C, Pascal, Golang, NodeJS and Python. However, not all features of the IDL are yet supported by the individual export language:

#### Feature Matrix: Bindings
| Binding     |         Status                                             | Operating Systems |   class   |  scalar type  |     struct    |  enumeration  |     string    | basicarray | structarray | Callbacks |
|:-----------:|:----------------------------------------------------------:|:-----------------:|:---------:|:-------------:|:-------------:|:-------------:|:-------------:|:----------:|:-----------:|:---------:|
| C++         | ![](Documentation/images/Tick.png) mature                  | Win, Linux, MacOS | in,return | in,out,return | in,out,return | in,out,return | in,out,return |   in,out   |    in,out   |    in     |
| C++ Dynamic | ![](Documentation/images/Tick.png) mature                  | Win               | in,return | in,out,return | in,out,return | in,out,return | in,out,return |   in,out   |    in,out   |    in     |
| C           | ![](Documentation/images/Tick.png) mature                  | Win, Linux, MacOS | in,return | in,out,return | in,out,return | in,out,return | in,out,return |   in,out   |    in,out   |    in     |
| C   Dynamic | ![](Documentation/images/Tick.png) mature                  | Win               | in,return | in,out,return | in,out,return | in,out,return | in,out,return |   in,out   |    in,out   |    in     |
| Pascal      | ![](Documentation/images/Tick.png) mature                  | Win, Linux, MacOS | in,return | in,out,return | in,out,return | in,out,return | in,out,return |   in,out   |    in,out   |    in     |
| Python      | ![](Documentation/images/Tick.png) complete (but unstable) | Win, Linux, MacOS | in,return | in,out,return | in,out,return | in,out,return | in,out,return |   in,out   |    in,out   |    in     |
| Golang      | ![](Documentation/images/O.png) partial support            | Win, Linux, MacOS | in,return | in,out,return |       ?       |       ?       |      ?        |       ?    |      ?      |     -     |
| NodeJS      | ![](Documentation/images/O.png) partial support            | Win, Linux, MacOS | in,return | in,out,return |       ?       |       ?       |      ?        |       ?    |      ?      |     -     |

#### Feature Matrix: Implementation Stubs
| Implementation |         Status                                        | Operating Systems |   class   |  scalar type  |     struct    |  enumeration  |     string    | basicarray | structarray | Callbacks | Journaling |
|:--------------:|:-----------------------------------------------------:|:-----------------:|:---------:|:-------------:|:-------------:|:-------------:|:-------------:|:----------:|:-----------:|:---------:|:----------:|
| C++            | ![](Documentation/images/Tick.png) mature             | Win, Linux, MacOS | in,return | in,out,return | in,out,return | in,out,return | in,out,return |   in,out   |    in,out   | in        | +          |
| Pascal         | ![](Documentation/images/X.png) in development        |                   |           |               |               |               |               |            |             |           |            |


## Example
TODO: Annotation of an example project

## Background: the hourglass pattern-API for shared software components
A very clean approach to creating software components is the hourglass pattern for APIs.
The rationale is to pipe any domain code with a thick API through a C89-interface, thereby catching all exceptions.

            Domain-Code with thick API
      \ (templates, complex classes, custom  /
         \ exceptions, custom control flows/
           \         ...                /
             \                        /
                Narrow C89-interface (thin API)
                  (return values,
                   error codes, 
                   error strings)
               /                     \
             /                         \
        Language bindings in any other language;
                      Thick API

This enables producers of libraries to use their own programming paradigms and styles, without affecting their consumer and results in a great isolation of code (and responsibility) between library-producer and -consumer.
Due to the very clear interface, such libraries are very easy to integrate in existing projects.

A much more detailed introduction the topic is this presentation: https://www.youtube.com/watch?v=PVYdHDm0q6Y

### Difficulty of this approach:
Generating (and maintaining!) the required layers of interfaces (language bindings, thin API and domain code-API) and their consistency is error prone _if_ it is not automated.

