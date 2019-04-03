# ![ACT logo](images/ACT_logo_50px.png) Automatic Component Toolkit

## Specification of the Interface Description Language of the Automatic Component Toolkit (ACT-IDL)



| **Version** | 1.5.0 |
| --- | --- |

## Disclaimer

THESE MATERIALS ARE PROVIDED "AS IS." The contributors expressly disclaim any warranties (express, implied, or otherwise), including implied warranties of merchantability, non-infringement, fitness for a particular purpose, or title, related to the materials. The entire risk as to implementing or otherwise using the materials is assumed by the implementer and user. IN NO EVENT WILL ANY MEMBER BE LIABLE TO ANY OTHER PARTY FOR LOST PROFITS OR ANY FORM OF INDIRECT, SPECIAL, INCIDENTAL, OR CONSEQUENTIAL DAMAGES OF ANY CHARACTER FROM ANY CAUSES OF ACTION OF ANY KIND WITH RESPECT TO THIS DELIVERABLE OR ITS GOVERNING AGREEMENT, WHETHER BASED ON BREACH OF CONTRACT, TORT (INCLUDING NEGLIGENCE), OR OTHERWISE, AND WHETHER OR NOT THE OTHER MEMBER HAS BEEN ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

## Table of Contents

- [Preface](#preface)
   * [Document Conventions](#document-conventions)
   * [Language Notes](#language-notes)
 - [Elements and types in the ACT-IDL](#elements-and-types-in-the-act-idl)
   * [1. Component](#1-component)
   * [2. License](#2-license)
   * [3. License Line](#3-license-line)
   * [4. Bindings](#4-bindings)
   * [5. Implementations](#5-implementations)
   * [6. Export](#6-export)
   * [7. Global](#7-global)
   * [8. Class](#8-class)
   * [9. Function Type](#9-function-type)
   * [10. Param](#10-param)
   * [11. Enum](#11-enum)
   * [12. Option](#12-option)
   * [13. Struct](#13-struct)
   * [14. Member](#14-member)
   * [15. Errors](#15-errors)
   * [16. Error](#16-error)
   * [17. Simple Types](#17-simple-types)
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
| year | **ST\_Year** | optional | the current year | The year associcated with the copyright. |
| @anyAttribute | | | | |

It is RECOMMENDED that components generated with ACT follow the [semantic versioning scheme](https://semver.org/).
The "version" attribute encodes the semantic version of this component. Major, Minor and Micro-version info MUST be included. Pre-release information and build information MAY be included.

The \<component> element is the root element of a ACT-IDL file.
There MUST be exactly one \<component> element in a ACT-IDL file.
A component MUST have exactly one child [license](#2-license) element, 
one child [bindings](#4-bindings) element, 
one child [implementations](#5-implementations) element, 
one child [errors](#15-errors) element and 
one child [global](#7-global) element.

The names of the \<struct>-, \<enum>-, \<functiontype>- and \<class>-elements MUST be unique within the \<component>.

>**Note:** Regarding the \"uniqueness\" of attributes of type **xs:string**.
>Within this specification strings are considered equal regardless of the case of the individual letters.

## 2. License
Element **\<license>** of type **CT\_License**

![element license](images/element_license.png)

The \<license> element contains a list of at least one child [line](#3-line) element.
The license lines will be included as comments at the start of all generated source code files.

## 3. Line
Element **\<line>** of type **CT\_LicenseLine**

![element line](images/element_line.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| value | **xs:string** | required | | A line of the license. |

## 4. Bindings
Element **\<bindings>** of type **CT\_BindingList**

![element bindings](images/element_bindings.png)

The CT\_BindingList type contains a list of [binding](#6-export) elements.
The \<binding> elements in the \<bindings> element determine the language bindings that will be generated.

## 5. Implementations
Element **\<implementations>** of type **CT\_ImplementationsList**

![element implementation](images/element_implementations.png)

The CT\_ImplementationsList type contains a list of [implementation](#6-export) elements.
The \<implementation> elements in the \<implementations> element determine the languages for which implementation stubs will be generated.

## 6. Export
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
| language | **ST\_Language** | required | | The programming langugage of this export. |
| indentation | **ST\_Indentation** | optional | 4spaces | Which string should be used to denote a single level of indentation in the generated source code files. |
| stubidentifier | **ST\_StubIdentifier** | optional | "" | Generated sources files of this export will follow the naming schme "...${BaseName}_${stubidentifier}...". Only used in \<implementation> right now. |
| classidentifier | **ST\_ClassIdentifier** | optional | "" | Generated classes of this export will follow the naming schme "...${ClassIdentifier}_${ClassName}...". The only binding that supports this are the C++-bindings.|

## 7. Global
Element **\<global>** of type **CT\_Global**

![element global](images/element_global.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| baseclassname | **ST\_Name** | required | | Specifies the name of a class that is the base class for all classes of the generated component. |
| releasemethod | **ST\_Name** | required | | Specifies the name of the method used to release a class instance owned by the generated component. |
| versionmethod | **ST\_Name** | required | | Specifies the name of the method used to obtain the major, minor and micro version of the component. |
| prereleasemethod | **ST\_Name** | required | | Specifies the name of the method used to obtain the prerelease information of the component. |
| buildinfomethod | **ST\_Name** | required | | Specifies the name of the method used to obtain the build information of the component. |
| errormethod | **ST\_Name** | required | | Specifies the name of the method used to query the last error that occured during the call of class's method. |
| journalmethod | **ST\_Name** | optional | | Specifies the name of the method used to set the journal file. If ommitted, journalling will not be built into the component. |

The \<global> element contains a list of [method](#9-function-type) elements that define the exported global functions of the component.
The names of the \<method> elements MUST be unique within the \<global> element.

The `baseclassname`-attribute must be the name of a \<class> element within the components list of classes.
This class will be the base class for all classes of the generated component.

The `releasemethod`-attribute must be the name of a \<method> within the \<global> element of a component that has exactly one parameter with `type="class"`, `class="BaseClass"` and `pass="in"`.
The `versionmethod`-attribute must be the name of a \<method> within the \<global> element of a component that has exactly three parameters. The three parameters MUST be of type `type="uint32"` and `pass="out"`.
The `prereleasemethod`-attribute is optional an can be the name of a \<method> within the \<global> element of a component that has two parameters.
The first parameter MUST be of type `type="bool"` and `pass="return"`, the second parameter MUST be of type `type="string"` and `pass="out"`.
The `buildinfomethod`-attribute is optional an can be the name of a \<method> within the \<global> element of a component that has two parameters.
The first parameter MUST be of type `type="bool"` and `pass="return"`, the second parameter MUST be of type `type="string"` and `pass="out"`.


The `errormethod`-attribute must be the name of a \<method> within the \<global> element of a method that has exactly three parameters:
1. `type="class"`, `class="$BASECLASSNAME"` and `pass="in"`, where `"$BASECLASSNAME"` is the value of the `baseclassname` attribute of the \<global> element.
2. `type="string"` and `pass="out"`: outputs the last error message
3. `type="bool"` and `pass="return"`: returns the instance of the baseclass has an error.

If the `journalmethod` attribute is given, it must be the name of a \<method> within the \<global> element of a method that has exactly one parameter with `type="string"` and `pass="in"`.

 **Note**
 `type="handle"` is equivalent to `type="class"` for backwards compatibility. It will be removed in a later version.

## 8. Class
Element **\<class>** of type **CT\_Class**

![element class](images/element_class.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this class. |
| parent | **ST\_Name** | optional | | The name of the parent class of this class. |
| description | **ST\_Description** | optional | | A description of this class. |

The \<class> element contains a list of [method](#9-function-type) elements that define the exported member functions of this class.
The names of the \<method> elements MUST be unique in this list.

If the `parent`-attribute is empty, and the name of this class differs from the `baseclassname`-attribute of the \<global> element, `baseclassname` will be considered as the parent class of this class.

A class MUST be defined in the list of \<class> elements before it is used as parent-class of another class. This restiction rules out circular inheritance. Moreover, the default `baseclassname` MUST be defined as the first \<class> within the IDL-file.

## 9. Function Type
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
| description | **ST\_Description** | required | | A description of this function type. |

The CT\_FunctionType-type describes the signature of a function in the interface.
Each element of type CT\_FunctionType contains a list of [param](#10-param) elements.
The names of the param in this list MUST be unique.
This list MUST contain zero or one param-elements with the pass-value \"return\".

The \<functiontype>-element can be used to define callback functions into the consumer's code.

## 10. Param
Element **\<param>** of type **CT\_Param**

![element param](images/element_param.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this parameter. |
| description | **ST\_Description** | required | | A description of this parameter. |
| pass | **ST\_Pass** | required | | Specifies whether the parameter is passed "in", "out" or as "return"-value of the enclosing functiontype. |
| type | **ST\_Type** | required | | The type of this parameter. |
| class | **ST\_Name** | optional | | Required if the type is an [**ST\_ComposedType**](#173-composedtype) |


## 11. Enum
Element **\<enum>** of type **CT\_Enum**

![element enum](images/element_enum.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this Enumeration. |

The \<enum> element defines an enumerated type (see https://en.wikipedia.org/wiki/Enumerated_type), i.e. a set of named values.<br/>
It contains a list of at least one [option](#12-option) element.
The names as well as the values of the options in this list MUST be unique within a \<enum> element.


## 12. Option
Element **\<option>** of type **CT\_Option**

![element option](images/element_option.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this option. |
| value | **xs:nonNegativeInteger** | required | | The numerical value of this option. |


## 13. Struct
Element **\<struct>** of type **CT\_Struct**

![element struct](images/element_struct.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this struct. |

The \<struct> element defines a composite data type (see https://en.wikipedia.org/wiki/Composite_data_type). <br/>
It contains a list of at least one [member](#14-member) element.
The names of the member elements MUST be unique within a struct element.


## 14. Member
Element **\<member>** of type **CT\_Member**

![element member](images/element_member.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_Name** | required | | The name of this member. |
| type | **ST\_ScalarType** | required | | The scalar type of this member. |
| rows | **xs:positiveInteger** | optional | 1 | The number of rows of this member. |
| columns | **xs:positiveInteger** | optional | 1 | The number of columns of this member. |

The \<member> element defines a member (or "field") within a struct. Only [**ST\_ScalarType**](#172-scalartype) is allowed within structs.
By default, the member defines a single value of its type within the enclusing struct. One- or two-dimensional arrays of fixed size can be
defined by setting the rows and colums attributes to the desired size of the array.


## 15. Errors
Element **\<errors>** of type **CT\_ErrorList**

![element errors](images/element_errors.png)

The \<errors> element contains a list of [\<error>](#16-error) elements.
The names and codes of the \<error> elements in this list MUST be unique within the \<errors> element.

Each ACT-component MUST contain at least the following 8 error codes:

`NOTIMPLEMENTED`, `INVALIDPARAM`, `INVALIDCAST`, `BUFFERTOOSMALL`, `GENERICEXCEPTION`, `COULDNOTLOADLIBRARY`, `COULDNOTFINDLIBRARYEXPORT`, `INCOMPATIBLEBINARYVERSION`

## 16. Error
Element **\<error>** of type **CT\_Error**

![element error](images/element_error.png)

##### Attributes
| Name | Type | Use | Default | Annotation |
| --- | --- | --- | --- | --- |
| name | **ST\_ErrorName** | required | | The name of this error. |
| code | **xs:positiveInteger** | required | | The numerical error code of this error. |
| description | **ST\_ErrorDescription** | otpional | | A short description of this error. |


## 17. Simple Types
The simple types of this specification encode features, concepts, data types,
and naming rules used in or required by programming languages.

For now, please look the up in the [ACT.xsd](../Source/ACT.xsd).

### 17.1 Type
Supported types are:
- `bool`: denotes a boolean value (`true` or `false`).
Although this can be encoded in a single bit, the thin C89-layer APIs generated by ACT will use a unsigned 8 bit value (a `uint8` in ACT terms) to encode a boolean value.
A numerical value of `0` encodes `false`, all oher values encode `true`.
Implementations and bindings should use the definition of a boolean value native to the respective language of the implementation or binding.
- `uint8`, `uint16`, `uint32`, `uint64`:
An _unsigned_ integer vaules ranging from 0 - 2<sup>8</sup>-1, 0 - 2<sup>16</sup>-1, 0 - 2<sup>32</sup>-1, 0 - 2<sup>64</sup>-1, respectively.
- `int8`, `int16`, `int32`, `int64`:
A _signed_ integer vaules ranging from -2<sup>7</sup> - 2<sup>7</sup>-1, -2<sup>15</sup> - 2<sup>15</sup>-1,
-2<sup>31</sup> - 2<sup>31</sup>-1,
-2<sup>63</sup> - 2<sup>63</sup>-1, respectively.
- `pointer`: An address memory without knowledge of the kind of data that resides there. In C++, this corresponds to a `void*`.
- `string` denotes a null-terminated string. If a component requires arbitrary strings that can contain null-characters, on should use the type `basicarray` of class `uint8`.
- `single`: Single precision floating point number.
- `double`: Double precision floating point number.
- `struct`: see [13. Struct](#13-struct)
- `enum`: see [11. Enum](#11-enum)
- `basicarray`: an array of [ST\_ScalarTypes](#17-2-scalartype)
- `enumarray`: an array of [enums](#11-enum)
- `structarray`: an array of [structs](#13-struct)
- `handle`: the identifier (address, unique identifier, hash, ...) of a class instance [class instance](#8-class)
- `functiontype`: see [9. Function Type](#9-function-type)

### 17.2 ScalarType
A subset of scalar or integral of ST\_Type:

`bool`, `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32`, `int64`, `single`, `double`, `pointer`.

### 17.3 ComposedType
A subset of more complex types, or types composed of other ST\_Types:

`string`, `enum`, `basicarray`, `enumarray`, `structarray`, `handle`, `functiontype`

### 17.4 Name
### 17.5 Description
### 17.6 ErrorName
### 17.7 ErrorDescription
### 17.8 Pass
### 17.9 Language
### 17.10 Indentation
### 17.11 Year
### 17.12 Version
### 17.13 Stub Identifier
### 17.14 Class Identifier
### 17.16 NameSpace
### 17.15 Library Name
### 17.16 Base Name


# Appendix A. XSD Schema of ACT-IDL
See [ACT.xsd](../Source/ACT.xsd)

# Appendix B. Example of ACT-IDL
See [libPrimes.xml](../Examples/Primes/libPrimes.xml)
