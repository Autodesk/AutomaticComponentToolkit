/*++

Copyright (C) 2023 Autodesk Inc.

All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

--*/

//////////////////////////////////////////////////////////////////////////////////////////////////////
// languagerust.go
// functions to generate the Rust-layer of a library's API (can be used in bindings or implementation)
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"
)

func writeRustBaseTypeDefinitions(componentdefinition ComponentDefinition, w LanguageWriter, NameSpace string, BaseName string) error {
	w.Writeln("#[allow(unused_imports)]")
	w.Writeln("use std::ffi;")
	w.Writeln("")
	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Version definition for %s", NameSpace)
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("")
	w.Writeln("const %s_VERSION_MAJOR : usize = %d;", strings.ToUpper(NameSpace), majorVersion(componentdefinition.Version))
	w.Writeln("const %s_VERSION_MINOR : usize = %d;", strings.ToUpper(NameSpace), minorVersion(componentdefinition.Version))
	w.Writeln("const %s_VERSION_MICRO : usize= %d;", strings.ToUpper(NameSpace), microVersion(componentdefinition.Version))
	w.Writeln("const %s_VERSION_PRERELEASEINFO : &str = \"%s\";", strings.ToUpper(NameSpace), preReleaseInfo(componentdefinition.Version))
	w.Writeln("const %s_VERSION_BUILDINFO : &str = \"%s\";", strings.ToUpper(NameSpace), buildInfo(componentdefinition.Version))

	w.Writeln("")
	w.Writeln("")

	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Basic pointers definition for %s", NameSpace)
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("")
	w.Writeln("type Handle = std::ffi::c_void")

	if len(componentdefinition.Enums) > 0 {
		w.Writeln("/*************************************************************************************************************************")
		w.Writeln(" Enum definitions for %s", NameSpace)
		w.Writeln("**************************************************************************************************************************/")
		w.Writeln("")
		for i := 0; i < len(componentdefinition.Enums); i++ {
			enuminfo := componentdefinition.Enums[i]
			w.Writeln("#[repr(C, u16)]")
			w.Writeln("#[allow(non_snake_case)]")
			w.Writeln("pub enum %s {", enuminfo.Name)
			for j := 0; j < len(enuminfo.Options); j++ {
				option := enuminfo.Options[j]
				sep := ","
				if j == len(enuminfo.Options)-1 {
					sep = ""
				}
				w.Writeln("  pub %s = %d%s", option.Name, option.Value, sep)
			}
			w.Writeln("}")
			w.Writeln("")
		}
	}

	if len(componentdefinition.Structs) > 0 {
		w.Writeln("/*************************************************************************************************************************")
		w.Writeln(" Interface Struct definitions for %s", NameSpace)
		w.Writeln("**************************************************************************************************************************/")
		w.Writeln("")
		for i := 0; i < len(componentdefinition.Structs); i++ {
			structinfo := componentdefinition.Structs[i]
			w.Writeln("#[repr(C)]")
			w.Writeln("#[allow(non_snake_case)]")
			w.Writeln("pub struct %s {", structinfo.Name)
			for j := 0; j < len(structinfo.Members); j++ {
				member := structinfo.Members[j]
				last := j == len(structinfo.Members)-1
				err := writeRustMemberLine(member, w, structinfo.Name, last)
				if err != nil {
					return err
				}
			}
			w.Writeln("}")
			w.Writeln("")
		}
	}

	if len(componentdefinition.Functions) > 0 {
		w.Writeln("/*************************************************************************************************************************")
		w.Writeln(" Function type definitions for %s", NameSpace)
		w.Writeln("**************************************************************************************************************************/")
		w.Writeln("")
		for i := 0; i < len(componentdefinition.Functions); i++ {
			funcinfo := componentdefinition.Functions[i]
			w.Writeln("// %s", funcinfo.FunctionDescription)
			w.Writeln("//")
			parameterString := ""
			for j := 0; j < len(funcinfo.Params); j++ {
				RustParameters, err := generatePlainRustParameters(funcinfo.Params[j])
				RustParameter := RustParameters[0]
				if err != nil {
					return err
				}
				w.Writeln("// %s", RustParameter.ParamComment)
				if j == 0 {
					parameterString += fmt.Sprintf("%s : %s", RustParameter.ParamName, RustParameter.ParamType)
				} else {
					parameterString += fmt.Sprintf(", %s : %s", RustParameter.ParamName, RustParameter.ParamType)
				}
			}
			w.Writeln("type %s = unsafe extern \"C\" fn(%s);", funcinfo.FunctionName, parameterString)
		}
	}

	return nil
}

func writeRustMemberLine(member ComponentDefinitionMember, w LanguageWriter, StructName string, last bool) error {
	suffix := ","
	if last {
		suffix = ""
	}
	arraysuffix := suffix
	if member.Rows > 0 {
		if member.Columns > 0 {
			arraysuffix = fmt.Sprintf("[%d][%d]%s", member.Columns, member.Rows, suffix)
		} else {
			arraysuffix = fmt.Sprintf("[%d]%s", member.Rows, suffix)
		}
	}
	switch member.Type {
	case "uint8":
		w.Writeln("  pub %s: u8%s", member.Name, arraysuffix)
	case "uint16":
		w.Writeln("  pub %s: u16%s", member.Name, arraysuffix)
	case "uint32":
		w.Writeln("  pub %s: u32%s", member.Name, arraysuffix)
	case "uint64":
		w.Writeln("  pub %s: u64%s", member.Name, arraysuffix)
	case "int8":
		w.Writeln("  pub %s: i8%s", member.Name, arraysuffix)
	case "int16":
		w.Writeln("  pub %s: i16%s", member.Name, arraysuffix)
	case "int32":
		w.Writeln("  pub %s: i32%s", member.Name, arraysuffix)
	case "int64":
		w.Writeln("  pub %s: i64%s", member.Name, arraysuffix)
	case "bool":
		w.Writeln("  pub %s: bool%s", member.Name, arraysuffix)
	case "single":
		w.Writeln("  pub %s: f32%s", member.Name, arraysuffix)
	case "double":
		w.Writeln("  pub %s: f64%s", member.Name, arraysuffix)
	case "pointer":
		w.Writeln("  pub %s: c_void%s", member.Name, arraysuffix)
	case "string":
		return fmt.Errorf("it is not possible for struct %s to contain a string value", StructName)
	case "class", "optionalclass":
		return fmt.Errorf("it is not possible for struct %s to contain a handle value", StructName)
	case "enum":
		w.Writeln("  pub %s: u16%s", member.Name, arraysuffix)
	}
	return nil
}

// CParameter is a handy representation of a function parameter in C
type RustParameter struct {
	ParamType    string
	ParamName    string
	ParamComment string
}

func generatePlainRustParameters(param ComponentDefinitionParam) ([]RustParameter, error) {

}

func generateRustParameterType(param ComponentDefinitionParam, isPlain bool) (string, error) {
	RustParamTypeName := ""
	ParamTypeName := param.ParamType
	ParamClass := param.ParamClass
	switch ParamTypeName {
	case "uint8":
		RustParamTypeName = "u8"

	case "uint16":
		RustParamTypeName = "u16"

	case "uint32":
		RustParamTypeName = "u32"

	case "uint64":
		RustParamTypeName = "u64"

	case "int8":
		RustParamTypeName = "i8"

	case "int16":
		RustParamTypeName = "i16"

	case "int32":
		RustParamTypeName = "i32"

	case "int64":
		RustParamTypeName = "i64"

	case "bool":
		if isPlain {
			RustParamTypeName = "u8"
		} else {
			RustParamTypeName = "bool"
		}

	case "single":
		RustParamTypeName = "f32"

	case "double":
		RustParamTypeName = "f64"

	case "pointer":
		RustParamTypeName = "c_void"

	case "string":
		if isPlain {
			RustParamTypeName = "*mut char"
		} else {
			// TODO
			return "", fmt.Errorf("Not yet handled")
		}

	case "enum":
		if isPlain {
			RustParamTypeName = fmt.Sprintf("u16")
		} else {
			// TODO
			return "", fmt.Errorf("Not yet handled")
		}

	case "functiontype":
		if isPlain {
			RustParamTypeName = fmt.Sprintf("%s", ParamClass)
		} else {
			// TODO
			return "", fmt.Errorf("Not yet handled")
		}

	case "struct":
		if isPlain {
			RustParamTypeName = fmt.Sprintf("*mut %s", ParamClass)
		} else {
			// TODO
			return "", fmt.Errorf("Not yet handled")
		}

	case "basicarray":
		basicTypeName, err := generateRustParameterType(param, isPlain)
		if err != nil {
			return "", err
		}

		if isPlain {
			RustParamTypeName = fmt.Sprintf("*mut %s", basicTypeName)
		} else {
			// TODO
			return "", fmt.Errorf("Not yet handled")
		}

	case "structarray":
		if isPlain {
			RustParamTypeName = fmt.Sprintf("*mut %s", ParamClass)
		} else {
			// TODO
			return "", fmt.Errorf("Not yet handled")
		}

	case "class", "optionalclass":
		if isPlain {
			RustParamTypeName = fmt.Sprintf("Handle")
		} else {
			// TODO
			return "", fmt.Errorf("Not yet handled")
		}

	default:
		return "", fmt.Errorf("invalid parameter type \"%s\" for Pascal parameter", ParamTypeName)
	}

	return RustParamTypeName, nil
}
