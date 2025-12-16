# ![ACT logo](images/ACT_logo_50px.png) Automatic Component Toolkit

## Specification of the Interface Description Language of the Automatic Component Toolkit (ACT-IDL)



| **Version** | 1.8.1 |
| --- | --- |

## Disclaimer

THESE MATERIALS ARE PROVIDED "AS IS." The contributors expressly disclaim any warranties (express, implied, or otherwise), including implied warranties of merchantability, non-infringement, fitness for a particular purpose, or title, related to the materials. The entire risk as to implementing or otherwise using the materials is assumed by the implementer and user. IN NO EVENT WILL ANY MEMBER BE LIABLE TO ANY OTHER PARTY FOR LOST PROFITS OR ANY FORM OF INDIRECT, SPECIAL, INCIDENTAL, OR CONSEQUENTIAL DAMAGES OF ANY CHARACTER FROM ANY CAUSES OF ACTION OF ANY KIND WITH RESPECT TO THIS DELIVERABLE OR ITS GOVERNING AGREEMENT, WHETHER BASED ON BREACH OF CONTRACT, TORT (INCLUDING NEGLIGENCE), OR OTHERWISE, AND WHETHER OR NOT THE OTHER MEMBER HAS BEEN ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

## Table of Contents

- [Preface](#preface)
   * [Document Conventions](#document-conventions)
   * [Language Notes](#language-notes)
 - [Elements and types in the ACT-IDL](#elements-and-types-in-the-act-idl)
   * [1. Component](#1-component)
   * [2. Import Component](#2-import-component)
   * [3. License](#3-license)
   * [4. License Line](#4-license-line)
   * [5. Bindings](#5-bindings)
   * [6. Implementations](#6-implementations)
   * [7. Export](#7-export)
   * [8. Global](#8-global)
   * [9. Class](#9-class)
   * [10. Function Type](#10-function-type)
   * [11. Param](#11-param)
   * [12. Enum](#12-enum)
   * [13. Option](#13-option)
   * [14. Struct](#14-struct)
   * [15. Member](#15-member)
   * [16. Errors](#16-errors)
   * [17. Error](#17-error)
   * [18. Simple Types](#18-simple-types)
 - [Appendix A. XSD Schema of ACT-IDL](#appendix-a-xsd-schema-of-act-idl)
 - [Appendix B. Example of ACT-IDL](#appendix-b-example-of-act-idl)

# Preface

## Document Conventions

Except where otherwise noted, syntax descriptions are expressed in the ABNF format as defined in RFC 4234.

Glossary terms are formatted like _this_.

Syntax descriptions and code are formatted as `Markdown code blocks.`

Replaceable items, that is, an item intended to be replaced by a value, are formatted in _`monospace cursive`_ type.

Notes are formatted as follows:

>**Note:** This is a note.

## Language Notes

In this specification, the words that are used to define the significance of each requirement are written in uppercase. These words are used in accordance with their definitions in RFC 2119, and their respective meanings are reproduced below:

- _MUST._ This word, or the adjective "REQUIRED," means that the item is an absolute requirement of the specification.
- _SHOULD._ This word, or the adjective "RECOMMENDED," means that there may exist valid reasons in particular circumstances to ignore this item, but the full implications should be understood and the case carefully weighed before choosing a different course.
- _MAY._ This word, or the adjective "OPTIONAL," means that this item is truly optional.


# Elements and types in the ACT-IDL

## 1. Component
Element **\<component>** of type **CT\_Component**

![element component](images/element_component.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| libraryname | **ST\_LibraryName** | required | | Specifies the name of the component. Can contain spaces. |
| namespace | **ST\_NameSpace** | required | | Specifies the namespace for the components's functionality. |
| copyright | **xs:string** | required | | The legal copyright holder. |
| basename | **ST\_BaseName** | required | | The basename will be used as prefix for generated filenames and all sorts of identifiers in the generated source code. |
| version | **ST\_Version** | required | | The semantic version of this component. |
| year | **ST\_Year** | optional | the current year | The year associated with the copyright. |
| @anyAttribute | | | | |

It is RECOMMENDED that components generated with ACT follow the [semantic versioning scheme](https://semver.org/).
The "version" attribute encodes the semantic version of this component. Major, Minor and Micro-version info MUST be included. Pre-release information and build information MAY be included.

The \<component> element is the root element of a ACT-IDL file.
There MUST be exactly one \<component> element in a ACT-IDL file.
A component MUST have exactly one child [license](#3-license) element, 
one child [bindings](#5-bindings) element, 
one child [implementations](#6-implementations) element, 
one child [errors](#16-errors) element and 
one child [global](#8-global) element.

The names of the \<struct>-, \<enum>-, \<functiontype>- and \<class>-elements MUST be unique within the \<component>.

>**Note:** Regarding the \"uniqueness\" of attributes of type **xs:string**.
>Within this specification strings are considered equal regardless of the case of the individual letters.

## 2. Import Component
Element **\<importcomponent>** of type **CT\_ImportComponent**
![element importcomponent](images/element_importcomponent.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| uri | **xs:string** | optional | | The relative location of another ACT-IDL-file. |
| namespace | **ST\_NameSpace** | optional | | The namespace of the imported component. |

The \<importcomponent> element defines the namespace and the relative location of another ACT-IDL-file and is the mechanism that allows injecting another ACT-component into the ACT-component at hand.

The `namespace` attribute of the \<importcomponent> element MUST match the `namespace` of the \<component> element within the file at location of the `uri` attribute.

The `class`es, `functiontype`s, `struct`s and `enum`s of the imported component will be available as `param`s of methods in this ACT-component. 
To use an entity with name `Y` from another ACT component (with namespace `X`) as `class` of a `param` in this ACT component set the `class`-attribute to `class="X:Y"`.

To be able to inject a component `Inner` into a component `Outer`, component `Inner` must define the `symbollookupmethod` in its global section and component `Outer` must define the `injectionmethod`.

>**Note:** Component injection is an advanced feature. Not all bindings support it.
> See [the Injection example](../Examples/Injection) for a minimal working example.

## 3. License
Element **\<license>** of type **CT\_License**

![element license](images/element_license.png)

The \<license> element contains a list of at least one child [line](#3-line) element.
The license lines will be included as comments at the start of all generated source code files.

## 4. Line
Element **\<line>** of type **CT\_LicenseLine**

![element line](images/element_line.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| value | **xs:string** | required | | A line of the license. |

## 5. Bindings
Element **\<bindings>** of type **CT\_BindingList**

![element bindings](images/element_bindings.png)

The CT\_BindingList type contains a list of [binding](#7-export) elements.
The \<binding> elements in the \<bindings> element determine the language bindings that will be generated.

## 6. Implementations
Element **\<implementations>** of type **CT\_ImplementationsList**

![element implementation](images/element_implementations.png)

The CT\_ImplementationsList type contains a list of [implementation](#7-export) elements.
The \<implementation> elements in the \<implementations> element determine the languages for which implementation stubs will be generated.

>**Note:** Currently, only `Cpp` (C++) and `Pascal` are supported for implementations. Other languages listed in [ST\_Language](#189-language) are not yet supported for generating implementation stubs.

## 7. Export
Element **\<binding>**
<br/>
![element binding](images/element_binding.png)
<br/>
Element **\<implmentation>**
<br/>
![element implmentation](images/element_implementation.png)
<br/>
of type **CT\_Export**
<br/>
![type export](images/type_export.png)


##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| language | **ST\_Language** | required | | The programming language of this export. |
| indentation | **ST\_Indentation** | optional | 4spaces | Which string should be used to denote a single level of indentation in the generated source code files. |
| stubidentifier | **ST\_StubIdentifier** | optional | "" | Generated sources files of this export will follow the naming scheme "...${BaseName}_${stubidentifier}...". Only used in \<implementation> right now. |
| classidentifier | **ST\_ClassIdentifier** | optional | "" | Generated classes of this export will follow the naming scheme "...${ClassIdentifier}_${ClassName}...". The only binding that supports this are the C++-bindings. |
| documentation | **xs:string** | optional | | Specifies the documentation format to generate. Currently only "sphinx" is supported for C++ bindings. |
| version | **xs:string** | optional | | Specifies a version string for this binding. For Java bindings, this attribute is used to specify the Java version: "8" or "1.8" for Java 8, "9" or "1.9" for Java 9. If not specified for Java bindings, the default is Java 9. |

## 8. Global
Element **\<global>** of type **CT\_Global**

![element global](images/element_global.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| baseclassname | **ST\_Name** | required | | Specifies the name of a class that is the base class for all classes of the generated component. |
| stringoutclassname | **ST\_Name** | optional | | Specifies the name of a class (or its parent hierarchy) that enables string output parameter caching. When specified, methods in classes derived from this class that return strings can use caching to improve performance. If omitted, string output caching is disabled. |
| acquiremethod | **ST\_Name** | required | | Specifies the name of the method used to acquire ownership of a class instance owned by the generated component. |
| releasemethod | **ST\_Name** | required | | Specifies the name of the method used to release ownership of a class instance owned by the generated component. |
| errormethod | **ST\_Name** | required | | Specifies the name of the method used to query the last error that occured during the call of class's method. |
| versionmethod | **ST\_Name** | required | | Specifies the name of the method used to obtain the major, minor and micro version of the component. |
| prereleasemethod | **ST\_Name** | optional | | Specifies the name of the method used to obtain the prerelease information of the component. |
| classtypeidmethod | **ST\_Name** | required | | Specifies the name of the method in base class used to get class type id of an object. |
| buildinfomethod | **ST\_Name** | optional | | Specifies the name of the method used to obtain the build information of the component. |
| injectionmethod | **ST\_Name** | optional | | Specifies the name of the method used to inject the symbollookupmethod another ACT component into this component at runtime. |
| symbollookupmethod | **ST\_Name** | optional | | Specifies the name of the method that returns the address of a given symbol exported by this component. |
| journalmethod | **ST\_Name** | optional | | Specifies the name of the method used to set the journal file. If omitted, journalling will not be built into the component. |

The \<global> element contains a list of [method](#10-function-type) elements that define the exported global functions of the component and defines special methods of the component.
The names of the \<method> elements MUST be unique within the \<global> element.

The `baseclassname`-attribute must be the name of a \<class> element within the components list of classes.
This class will be the base class for all classes of the generated component.

The `acquiremethod`- and `releasemethod`-attributes must each be the name of a \<method> within the \<global> element of a component that has exactly one parameter with `type="class"`, `class="BaseClass"` and `pass="in"`.
The `versionmethod`-attribute must be the name of a \<method> within the \<global> element of a component that has exactly three parameters. The three parameters MUST be of type `type="uint32"` and `pass="out"`.
The `prereleasemethod`-attribute is optional and can be the name of a \<method> within the \<global> element of a component that has two parameters.
The first parameter MUST be of type `type="bool"` and `pass="return"`, the second parameter MUST be of type `type="string"` and `pass="out"`.
The `classtypeidmethod`-attribute must be the name of a \<method> within the baseclassname \<class> element of a component that has exactly one parameter with `type="uint64"` and `pass="return"`.
The `buildinfomethod`-attribute is optional and can be the name of a \<method> within the \<global> element of a component that has two parameters.
The first parameter MUST be of type `type="bool"` and `pass="return"`, the second parameter MUST be of type `type="string"` and `pass="out"`.

The `errormethod`-attribute must be the name of a \<method> within the \<global> element of a method that has exactly three parameters:
1. `type="class"`, `class="$BASECLASSNAME"` and `pass="in"`, where `"$BASECLASSNAME"` is the value of the `baseclassname` attribute of the \<global> element.
2. `type="string"` and `pass="out"`: outputs the last error message
3. `type="bool"` and `pass="return"`: returns the instance of the baseclass has an error.

If the `injectionmethod` attribute is given, it must be the name of a \<method> within the \<global> element of a method that has exactly two parameters with `type="string"` and `pass="in"` and `type="pointer"` and `pass="in"`.

If the `symbollookupmethod` attribute is given, it must be the name of a \<method> within the \<global> element of a method that has exactly one parameter with `type="pointer"` and `pass="return"`. The implementation of this method is fully autogenerated and returns the address of another internal lookup method. This internal lookup method in turn is similar to a `GetProcAddress`- or `dlsym`-method: given the name of a method in this component, it provides the address of a method in this component with this name. The return value of the `symbollookupmethod` is usually passed into the `injectionmethod` of another component.

If the `journalmethod` attribute is given, it must be the name of a \<method> within the \<global> element of a method that has exactly one parameter with `type="string"` and `pass="in"`.

## 9. Class
Element **\<class>** of type **CT\_Class**

![element class](images/element_class.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this class. |
| parent | **ST\_Name** | optional | | The name of the parent class of this class. |
| description | **ST\_Description** | optional | | A description of this class. |
| threadsafetyoption | **xs:string** | optional | "none" | Specifies the thread safety level for this class. Valid values are: `none` (no thread safety), `soft` (locks only when strings are returned), `strict` (locks on every method call). Child classes cannot have a less strict option than their parent. If not specified, inherits from parent or defaults to `none`. |

The \<class> element contains a list of [method](#10-function-type) elements that define the exported member functions of this class.
The names of the \<method> elements MUST be unique in this list.

If the `parent`-attribute is empty, and the name of this class differs from the `baseclassname`-attribute of the \<global> element, `baseclassname` will be considered as the parent class of this class.

A class MUST be defined in the list of \<class> elements before it is used as parent-class of another class. This restiction rules out circular inheritance. Moreover, the default `baseclassname` MUST be defined as the first \<class> within the IDL-file.

## 10. Function Type
Element **\<functiontype>**
<br/>
![element functiontype](images/element_functiontype.png)

Element **\<method>**
<br/>
![element method](images/element_method.png)

of Complex type **CT\_FunctionType**
<br/>
![type functiontype](images/type_functiontype.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this function type. |
| description | **ST\_ErrorDescription** | optional | | A description of this function type. |

>**Note:** The \<method> element (used within \<class> and \<global>) also supports the following additional attribute:
>- `disablestringoutcache`: **xs:boolean** (optional, default: `false`): When set to `true`, disables string output parameter caching for this method, even if the class is part of the string output class hierarchy defined by `stringoutclassname` in the \<global> element.

The CT\_FunctionType-type describes the signature of a function in the interface.
Each element of type CT\_FunctionType contains a list of [param](#11-param) elements.
The names of the param in this list MUST be unique.
This list MUST contain zero or one param-elements with the pass-value \"return\".

The \<functiontype>-element can be used to define callback functions into the consumer's code.

## 11. Param
Element **\<param>** of type **CT\_Param**

![element param](images/element_param.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this parameter. |
| pass | **ST\_Pass** | required | | Specifies whether the parameter is passed "in", "out" or as "return"-value of the enclosing functiontype. |
| type | **ST\_Type** | required | | The type of this parameter. |
| class | **ST\_NameSpacedClassName** | optional | | Required if the type is an [**ST\_ComposedType**](#183-composedtype). For imported components, use the format "Namespace:ClassName" to reference a class from another component. |
| description | **ST\_Description** | optional | | A description of this parameter. |


## 12. Enum
Element **\<enum>** of type **CT\_Enum**

![element enum](images/element_enum.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this enumerated type. |
| description | **ST\_Description** | optional | | A description of this enumerated type. |

The \<enum> element defines an enumerated type (see https://en.wikipedia.org/wiki/Enumerated_type), i.e. a set of named values.<br/>
It contains a list of at least one [option](#13-option) element.
The names as well as the values of the options in this list MUST be unique within a \<enum> element.


## 13. Option
Element **\<option>** of type **CT\_Option**

![element option](images/element_option.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this option. |
| value | **xs:nonNegativeInteger** | required | | The numerical value of this option. |
| description | **ST\_Description** | optional | | A description of this option. |


## 14. Struct
Element **\<struct>** of type **CT\_Struct**

![element struct](images/element_struct.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this struct. |
| description | **ST\_Description** | optional | | A description of this struct. |

The \<struct> element defines a composite data type (see https://en.wikipedia.org/wiki/Composite_data_type). <br/>
It contains a list of at least one [member](#15-member) element.
The names of the member elements MUST be unique within a struct element.


## 15. Member
Element **\<member>** of type **CT\_Member**

![element member](images/element_member.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this member. |
| type | **ST\_ScalarType** | required | | The scalar type of this member. |
| rows | **xs:positiveInteger** | optional | 1 | The number of rows of this member. |
| columns | **xs:positiveInteger** | optional | 1 | The number of columns of this member. |
| description | **ST\_Description** | optional | | A description of this member. |

The \<member> element defines a member (or "field") within a struct. Only [**ST\_ScalarType**](#182-scalartype) is allowed within structs.
By default, the member defines a single value of its type within the enclosing struct. One- or two-dimensional arrays of fixed size can be
defined by setting the rows and columns attributes to the desired size of the array.


## 16. Errors
Element **\<errors>** of type **CT\_ErrorList**

![element errors](images/element_errors.png)

The \<errors> element contains a list of [\<error>](#17-error) elements.
The names and codes of the \<error> elements in this list MUST be unique within the \<errors> element.

Each ACT-component MUST contain at least the following 8 error codes:

`NOTIMPLEMENTED`, `INVALIDPARAM`, `INVALIDCAST`, `BUFFERTOOSMALL`, `GENERICEXCEPTION`, `COULDNOTLOADLIBRARY`, `COULDNOTFINDLIBRARYEXPORT`, `INCOMPATIBLEBINARYVERSION`

## 17. Error
Element **\<error>** of type **CT\_Error**

![element error](images/element_error.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_ErrorName** | required | | The name of this error. |
| code | **xs:positiveInteger** | required | | The numerical error code of this error. |
| description | **ST\_ErrorDescription** | optional | | A short description of this error. |


## 18. Simple Types
The simple types of this specification encode features, concepts, data types,
and naming rules used in or required by programming languages.

For now, please look the up in the [ACT.xsd](../Source/ACT.xsd).

### 18.1 Type
Supported types are:
- `bool`: denotes a boolean value (`true` or `false`).
Although this can be encoded in a single bit, the thin C89-layer APIs generated by ACT will use an unsigned 8 bit value (a `uint8` in ACT terms) to encode a boolean value.
A numerical value of `0` encodes `false`, all other values encode `true`.
Implementations and bindings should use the definition of a boolean value that is native to the respective language of the implementation or binding.
- `uint8`, `uint16`, `uint32`, `uint64`:
An _unsigned_ integer values ranging from 0 - 2<sup>8</sup>-1, 0 - 2<sup>16</sup>-1, 0 - 2<sup>32</sup>-1, 0 - 2<sup>64</sup>-1, respectively.
- `int8`, `int16`, `int32`, `int64`:
A _signed_ integer values ranging from -2<sup>7</sup> - 2<sup>7</sup>-1, -2<sup>15</sup> - 2<sup>15</sup>-1,
-2<sup>31</sup> - 2<sup>31</sup>-1,
-2<sup>63</sup> - 2<sup>63</sup>-1, respectively.
- `pointer`: An address in memory without knowledge of the kind of data that resides there. In C++, this corresponds to a `void*`.
- `string` denotes a null-terminated string. If a component requires arbitrary strings that can contain null-characters, one should use the type `basicarray` of class `uint8`. When using `string` as `out`- or `return`-parameter, the size of the buffer that is passed through the ABI includes the terminating null-character.

- `single`: Single precision floating point number.
- `double`: Double precision floating point number.
- `struct`: see [13. Struct](#14-struct)
- `enum`: see [11. Enum](#12-enum)
- `basicarray`: an array of [ST\_ScalarTypes](#18-2-scalartype)
- `enumarray`: an array of [enums](#12-enum)
- `structarray`: an array of [structs](#14-struct)
- `functiontype`: see [9. Function Type](#10-function-type)
- `class`: the identifier (address, unique identifier, hash, ...) of a class instance [class instance](#9-class)
- `optionalclass`: behaves just like `class`, however, this instance may be empty, or null. A use case for `optionalclass` is e.g. a `findElementByName`-method of a list, which might or might not return a class instance.

**Note**
 `type="handle"` is equivalent to `type="class"` for backwards compatibility. It will be removed in a future version.

### 18.2 ScalarType
A subset of scalar or integral of ST\_Type:

`bool`, `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32`, `int64`, `single`, `double`, `pointer`.

### 18.3 ComposedType
A subset of more complex types, or types composed of other ST\_Types:

`string`, `enum`, `basicarray`, `enumarray`, `structarray`, `class`, `optionalclass`, `functiontype`

**Note:** `handle` is also included in ST\_ComposedType for backwards compatibility, but it is equivalent to `class` and will be removed in a future version.

### 18.4 Name
Type **ST\_Name** is a string that MUST match the pattern `[A-Z][a-zA-Z0-9_]{0,63}`. It is used for names of classes, structs, enums, methods, parameters, members, and options. Names MUST start with an uppercase letter and may contain letters, digits, and underscores.

### 18.5 Description
Type **ST\_Description** is a string that MUST match the pattern `[a-zA-Z][a-zA-Z0-9_\\/+\-:,.=!?()'; |]*`. It is used to provide human-readable descriptions for various elements. Descriptions MUST start with a letter and may contain letters, digits, and various punctuation characters.

### 18.6 ErrorName
Type **ST\_ErrorName** is a string that MUST match the pattern `[A-Z][A-Z0-9_]*`. It is used for error names. Error names MUST be in uppercase and may contain uppercase letters, digits, and underscores.

### 18.7 ErrorDescription
Type **ST\_ErrorDescription** is a string that MUST match the pattern `[a-zA-Z][a-zA-Z0-9_+\-:,.=!/ ]*`. It is used to provide short descriptions for errors. Error descriptions MUST start with a letter and may contain letters, digits, and various punctuation characters.

### 18.8 Pass
Type **ST\_Pass** is an enumeration with the following values:
- `in`: The parameter is passed into the function.
- `out`: The parameter is passed out of the function (the function writes to it).
- `return`: The parameter is the return value of the function.

### 18.9 Language
Type **ST\_Language** is an enumeration that specifies the programming language for bindings or implementations. Supported values are:
- `C`: C language binding
- `Cpp`: C++ language binding
- `CDynamic`: C language binding with dynamic linking
- `CppDynamic`: C++ language binding with dynamic linking
- `Python`: Python language binding
- `Pascal`: Pascal language binding
- `Fortran`: Fortran language binding (not yet supported for implementations)
- `Node`: Node.js language binding
- `Go`: Go language binding
- `CSharp`: C# language binding
- `Java`: Java language binding (supports Java 8 and Java 9; requires `version` attribute to specify "8", "1.8", "9", or "1.9"; defaults to Java 9 if not specified)
- `WASM`: JavaScript bindings through WebAssembly (generates WebAssembly bindings using Emscripten that can be used from JavaScript/TypeScript)
- `Cppwasmtime`: C++ bindings for WebAssembly using wasmtime (generates both host and guest bindings for wasmtime-based WebAssembly modules)

>**Note:** The XSD schema file (ACT.xsd) currently only enumerates `C`, `Cpp`, `CDynamic`, `CppDynamic`, `Python`, `Pascal`, `Fortran`, `Node`, `Go`, and `CSharp` in the ST_Language type definition. However, the codebase supports `Java`, `WASM`, and `Cppwasmtime` through the `@anyAttribute` mechanism, which allows additional language values beyond those explicitly defined in the schema. These additional languages are fully supported and documented here.

### 18.10 Indentation
Type **ST\_Indentation** is an enumeration that specifies the indentation style for generated source code files. Supported values are:
- `1spaces`, `2spaces`, `3spaces`, `4spaces`, `5spaces`, `6spaces`, `7spaces`, `8spaces`: Specifies the number of spaces for a single indentation level.
- `tabs`: Uses tab characters for indentation.

The default value is `4spaces`.

### 18.11 Year
Type **ST\_Year** is a positive integer that MUST be greater than 1900 and less than 2147483648. It represents the year associated with the copyright.

### 18.12 Version
Type **ST\_Version** is a string that MUST match the pattern `([0-9]+\.[0-9]+\.[0-9])(\-[a-zA-Z0-9.\-]+)?(\+[a-zA-Z0-9.\-]+)?`. It follows the semantic versioning scheme (see https://semver.org/):
- Major, minor, and micro version numbers are required (e.g., `1.2.3`).
- Pre-release information MAY be included (e.g., `1.2.3-alpha`).
- Build information MAY be included (e.g., `1.2.3+20130313144700`).

### 18.13 Stub Identifier
Type **ST\_StubIdentifier** is a string that MUST match the pattern `[A-Za-z0-9_]{0,63}`. It is used to customize the naming scheme for generated source files in implementations. When specified, generated source files will follow the naming scheme "...${BaseName}_${stubidentifier}...".

### 18.14 Class Identifier
Type **ST\_ClassIdentifier** is a string that MUST either be empty (`""`) or match the pattern `[A-Z][A-Za-z0-9_]{0,63}`. It is used to customize the naming scheme for generated classes. When specified, generated classes will follow the naming scheme "...${ClassIdentifier}_${ClassName}...". Currently, only C++ bindings support this attribute.

### 18.15 NameSpace
Type **ST\_NameSpace** is a string that MUST match the pattern `[A-Z][a-zA-Z0-9_]{0,63}`. It specifies the namespace for the component's functionality. Namespaces MUST start with an uppercase letter and may contain letters, digits, and underscores.

### 18.16 Library Name
Type **ST\_LibraryName** is a string that MUST match the pattern `[a-zA-Z][a-zA-Z0-9_+\-:,.=!/ ]*`. It specifies the name of the component and MAY contain spaces and various punctuation characters.

### 18.17 Base Name
Type **ST\_BaseName** is a string that MUST match the pattern `[a-zA-Z][a-zA-Z0-9_\-.]*`. It is used as a prefix for generated filenames and all sorts of identifiers in the generated source code. Base names MUST start with a letter and may contain letters, digits, underscores, hyphens, and dots.

### 18.18 NameSpaced Class Name
Type **ST\_NameSpacedClassName** is a string that MUST match the pattern `^(([A-Z][a-zA-Z0-9_]{0,63}):){0,1}([a-zA-Z0-9_]{0,64})`. It is used for the `class` attribute of parameters to reference classes, structs, enums, or function types. The format allows:
- A simple class name: `ClassName`
- A namespaced class name from an imported component: `Namespace:ClassName`

When referencing entities from imported components, use the format `Namespace:EntityName` where `Namespace` is the namespace of the imported component and `EntityName` is the name of the class, struct, enum, or function type.


# Appendix A. XSD Schema of ACT-IDL
See [ACT.xsd](../Source/ACT.xsd)

# Appendix B. Example of ACT-IDL
See [libPrimes.xml](../Examples/Primes/libPrimes.xml)
