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
	"regexp"
	"strings"
)

func toSnakeCase(BaseType string) string {
	var matchCapital = regexp.MustCompile("(.)([A-Z])")
	underscored := matchCapital.ReplaceAllString(BaseType, "${1}_${2}")
	return strings.ToLower(underscored)
}

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
	w.Writeln("type Handle = std::ffi::c_void;")

	if len(componentdefinition.Enums) > 0 {
		w.Writeln("/*************************************************************************************************************************")
		w.Writeln(" Enum definitions for %s", NameSpace)
		w.Writeln("**************************************************************************************************************************/")
		w.Writeln("")
		for i := 0; i < len(componentdefinition.Enums); i++ {
			enuminfo := componentdefinition.Enums[i]
			w.Writeln("#[repr(C, u16)]")
			enumName := enuminfo.Name
			w.Writeln("pub enum %s {", enumName)
			for j := 0; j < len(enuminfo.Options); j++ {
				option := enuminfo.Options[j]
				sep := ","
				if j == len(enuminfo.Options)-1 {
					sep = ""
				}
				optionName := option.Name
				w.Writeln("  pub %s = %d%s", optionName, option.Value, sep)
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
			structName := structinfo.Name
			w.Writeln("#[repr(C)]")
			w.Writeln("pub struct %s {", structName)
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
				RustParameters, err := generateRustParameters(funcinfo.Params[j], true)
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
			w.Writeln("//")
			funcName := funcinfo.FunctionName
			w.Writeln("type %s = unsafe extern \"C\" fn(%s);", funcName, parameterString)
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
	memberName := toSnakeCase(member.Name)
	switch member.Type {
	case "uint8":
		w.Writeln("  pub %s: u8%s", memberName, arraysuffix)
	case "uint16":
		w.Writeln("  pub %s: u16%s", memberName, arraysuffix)
	case "uint32":
		w.Writeln("  pub %s: u32%s", memberName, arraysuffix)
	case "uint64":
		w.Writeln("  pub %s: u64%s", memberName, arraysuffix)
	case "int8":
		w.Writeln("  pub %s: i8%s", memberName, arraysuffix)
	case "int16":
		w.Writeln("  pub %s: i16%s", memberName, arraysuffix)
	case "int32":
		w.Writeln("  pub %s: i32%s", memberName, arraysuffix)
	case "int64":
		w.Writeln("  pub %s: i64%s", memberName, arraysuffix)
	case "bool":
		w.Writeln("  pub %s: bool%s", memberName, arraysuffix)
	case "single":
		w.Writeln("  pub %s: f32%s", memberName, arraysuffix)
	case "double":
		w.Writeln("  pub %s: f64%s", memberName, arraysuffix)
	case "pointer":
		w.Writeln("  pub %s: c_void%s", memberName, arraysuffix)
	case "string":
		return fmt.Errorf("it is not possible for struct %s to contain a string value", StructName)
	case "class", "optionalclass":
		return fmt.Errorf("it is not possible for struct %s to contain a handle value", StructName)
	case "enum":
		w.Writeln("  pub %s: u16%s", memberName, arraysuffix)
	}
	return nil
}

// CParameter is a handy representation of a function parameter in C
type RustParameter struct {
	ParamType    string
	ParamName    string
	ParamComment string
}

func generateRustParameters(param ComponentDefinitionParam, isPlain bool) ([]RustParameter, error) {
	Params := make([]RustParameter, 1)
	ParamTypeName, err := generateRustParameterType(param, isPlain)
	if err != nil {
		return nil, err
	}

	if isPlain {
		if param.ParamType == "basicarray" {
			return nil, fmt.Errorf("Not yet handled")
		}

		if param.ParamType == "structarray" {
			return nil, fmt.Errorf("Not yet handled")
		}
	}

	Params[0].ParamType = ParamTypeName
	Params[0].ParamName = toSnakeCase(param.ParamName)
	Params[0].ParamComment = fmt.Sprintf("* @param[%s] %s - %s", param.ParamPass, Params[0].ParamName, param.ParamDescription)

	return Params, nil
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
			return "", fmt.Errorf("%s Not yet handled", param.ParamType)
		}

	case "enum":
		if isPlain {
			RustParamTypeName = fmt.Sprintf("u16")
		} else {
			switch param.ParamPass {
			case "out":
				RustParamTypeName = fmt.Sprintf("&mut %s", ParamClass)
			case "in", "return":
				RustParamTypeName = fmt.Sprintf("%s", ParamClass)
			}
		}

	case "functiontype":
		RustParamTypeName = fmt.Sprintf("%s", ParamClass)

	case "struct":
		if isPlain {
			RustParamTypeName = fmt.Sprintf("*mut %s", ParamClass)
		} else {
			switch param.ParamPass {
			case "out":
				RustParamTypeName = fmt.Sprintf("&mut %s", ParamClass)
			case "in":
				RustParamTypeName = fmt.Sprintf("& %s", ParamClass)
			case "return":
				RustParamTypeName = fmt.Sprintf("%s", ParamClass)
			}
		}

	case "basicarray":
		basicParam := param
		basicParam.ParamType = param.ParamClass
		basicParam.ParamPass = "return"
		basicTypeName, err := generateRustParameterType(basicParam, isPlain)
		if err != nil {
			return "", err
		}

		if isPlain {
			RustParamTypeName = fmt.Sprintf("*mut %s", basicTypeName)
		} else {
			switch param.ParamPass {
			case "out":
				RustParamTypeName = fmt.Sprintf("&mut Vec<%s>", basicTypeName)
			case "in":
				RustParamTypeName = fmt.Sprintf("&[%s]", basicTypeName)
			case "return":
				RustParamTypeName = fmt.Sprintf("Vec<%s>", basicTypeName)
			}
		}

	case "structarray":
		if isPlain {
			RustParamTypeName = fmt.Sprintf("*mut %s", ParamClass)
		} else {
			switch param.ParamPass {
			case "out":
				RustParamTypeName = fmt.Sprintf("&mut Vec<%s>", ParamClass)
			case "in":
				RustParamTypeName = fmt.Sprintf("&[%s]", ParamClass)
			case "return":
				RustParamTypeName = fmt.Sprintf("Vec<%s>", ParamClass)
			}
		}

	case "class", "optionalclass":
		if isPlain {
			RustParamTypeName = fmt.Sprintf("Handle")
		} else {
			// TODO
			return "", fmt.Errorf("%s Not yet handled", param.ParamType)
		}

	default:
		return "", fmt.Errorf("invalid parameter type \"%s\" for Rust parameter", ParamTypeName)
	}

	return RustParamTypeName, nil
}
