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
	w.Writeln("use std::ffi::c_void;")
	w.Writeln("")
	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Version definition for %s", NameSpace)
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("")
	w.Writeln("#[allow(dead_code)]")
	w.Writeln("pub const %s_VERSION_MAJOR : usize = %d;", strings.ToUpper(NameSpace), majorVersion(componentdefinition.Version))
	w.Writeln("#[allow(dead_code)]")
	w.Writeln("pub const %s_VERSION_MINOR : usize = %d;", strings.ToUpper(NameSpace), minorVersion(componentdefinition.Version))
	w.Writeln("#[allow(dead_code)]")
	w.Writeln("pub const %s_VERSION_MICRO : usize= %d;", strings.ToUpper(NameSpace), microVersion(componentdefinition.Version))
	w.Writeln("#[allow(dead_code)]")
	w.Writeln("pub const %s_VERSION_PRERELEASEINFO : &str = \"%s\";", strings.ToUpper(NameSpace), preReleaseInfo(componentdefinition.Version))
	w.Writeln("#[allow(dead_code)]")
	w.Writeln("pub const %s_VERSION_BUILDINFO : &str = \"%s\";", strings.ToUpper(NameSpace), buildInfo(componentdefinition.Version))

	w.Writeln("")
	w.Writeln("")

	if len(componentdefinition.Errors.Errors) > 0 {
		w.Writeln("/*************************************************************************************************************************")
		w.Writeln(" Error constants for %s", NameSpace)
		w.Writeln("**************************************************************************************************************************/")
		w.Writeln("")
		w.Writeln("#[allow(dead_code)]")
		w.Writeln("pub const %s_SUCCESS : i32 = 0;", strings.ToUpper(NameSpace))
		for i := 0; i < len(componentdefinition.Errors.Errors); i++ {
			errorcode := componentdefinition.Errors.Errors[i]
			w.Writeln("#[allow(dead_code)]")
			if errorcode.Description != "" {
				w.Writeln("pub const %s_ERROR_%s : i32 = %d; /** %s */", strings.ToUpper(NameSpace), errorcode.Name, errorcode.Code, errorcode.Description)
			} else {
				w.Writeln("pub const %s_ERROR_%s : i32 = %d;", strings.ToUpper(NameSpace), errorcode.Name, errorcode.Code)
			}
		}
		w.Writeln("")
		w.Writeln("")
	}

	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Handle definiton for %s", NameSpace)
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("")
	w.Writeln("// Enum of all traits - this acts as a handle as we pass trait pointers through the interface")
	w.Writeln("")
	w.Writeln("#[allow(dead_code)]")
	w.Writeln("pub enum HandleImpl {")
	w.AddIndentationLevel(1)
	for i := 0; i < len(componentdefinition.Classes); i++ {
		class := componentdefinition.Classes[i]
		if i != len(componentdefinition.Classes)-1 {
			w.Writeln("T%s(u64, Box<dyn %s>),", class.ClassName, class.ClassName)
		} else {
			w.Writeln("T%s(u64, Box<dyn %s>)", class.ClassName, class.ClassName)
		}
	}
	w.AddIndentationLevel(-1)
	w.Writeln("}")
	w.Writeln("")
	w.Writeln("pub type Handle = *mut HandleImpl;")
	for i := 0; i < len(componentdefinition.Classes); i++ {
		class := componentdefinition.Classes[i]
		w.Writeln("pub type %sHandle =Handle;", class.ClassName)
	}

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
			w.Writeln("#[derive(Clone)]")
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
			w.Writeln("pub type %s = unsafe extern \"C\" fn(%s);", funcName, parameterString)
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
		if param.ParamType == "string" {
			if param.ParamPass == "out" {
				Params = make([]RustParameter, 3)
				Params[0].ParamType = "usize"
				Params[0].ParamName = toSnakeCase(param.ParamName) + "_buffer_size"
				Params[0].ParamComment = fmt.Sprintf("* @param[in] %s - size of the buffer (including trailing 0)", Params[0].ParamName)

				Params[1].ParamType = "*mut usize"
				Params[1].ParamName = toSnakeCase(param.ParamName) + "_needed_chars"
				Params[1].ParamComment = fmt.Sprintf("* @param[out] %s - will be filled with the count of the written bytes, or needed buffer size.", Params[1].ParamName)

				Params[2].ParamType = "*mut u8"
				Params[2].ParamName = toSnakeCase(param.ParamName) + "_buffer"
				Params[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s buffer of %s, may be NULL", Params[2].ParamName, param.ParamClass, param.ParamDescription)

				return Params, nil
			}
		}

		if param.ParamType == "basicarray" {
			basicParam := param
			basicParam.ParamType = param.ParamClass
			basicParam.ParamPass = "in"
			basicTypeName, err := generateRustParameterType(basicParam, isPlain)
			if err != nil {
				return nil, err
			}
			if param.ParamPass == "out" {
				Params = make([]RustParameter, 3)
				Params[0].ParamType = "usize"
				Params[0].ParamName = toSnakeCase(param.ParamName) + "_buffer_size"
				Params[0].ParamComment = fmt.Sprintf("* @param[in] %s - size of the buffer (including trailing 0)", Params[0].ParamName)

				Params[1].ParamType = "*mut usize"
				Params[1].ParamName = toSnakeCase(param.ParamName) + "_count"
				Params[1].ParamComment = fmt.Sprintf("* @param[out] %s - will be filled with the count of the written elements, or needed buffer size.", Params[1].ParamName)

				Params[2].ParamType = fmt.Sprintf("*mut %s", basicTypeName)
				Params[2].ParamName = toSnakeCase(param.ParamName) + "_buffer"
				Params[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s buffer of %s, may be NULL", Params[0].ParamName, param.ParamClass, param.ParamDescription)

				return Params, nil
			} else {
				Params = make([]RustParameter, 2)
				Params[0].ParamType = "usize"
				Params[0].ParamName = toSnakeCase(param.ParamName) + "_buffer_size"
				Params[0].ParamComment = fmt.Sprintf("* @param[in] %s - size of the buffer (including trailing 0)", Params[0].ParamName)

				Params[1].ParamType = fmt.Sprintf("*const %s", basicTypeName)
				Params[1].ParamName = toSnakeCase(param.ParamName) + "_buffer"
				Params[1].ParamComment = fmt.Sprintf("* @param[in] %s - %s buffer of %s, may be NULL", Params[0].ParamName, param.ParamClass, param.ParamDescription)

				return Params, nil
			}
		}

		if param.ParamType == "structarray" {
			if param.ParamPass == "out" {
				Params = make([]RustParameter, 3)
				Params[0].ParamType = "usize"
				Params[0].ParamName = toSnakeCase(param.ParamName) + "_buffer_size"
				Params[0].ParamComment = fmt.Sprintf("* @param[in] %s - size of the buffer (including trailing 0)", Params[0].ParamName)

				Params[1].ParamType = "*mut usize"
				Params[1].ParamName = toSnakeCase(param.ParamName) + "_count"
				Params[1].ParamComment = fmt.Sprintf("* @param[out] %s - will be filled with the count of the written elements, or needed buffer size.", Params[1].ParamName)

				Params[2].ParamType = fmt.Sprintf("*mut %s", param.ParamClass)
				Params[2].ParamName = toSnakeCase(param.ParamName) + "_buffer"
				Params[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s buffer of %s, may be NULL", Params[0].ParamName, param.ParamClass, param.ParamDescription)

				return Params, nil
			} else {
				Params = make([]RustParameter, 2)
				Params[0].ParamType = "usize"
				Params[0].ParamName = toSnakeCase(param.ParamName) + "_buffer_size"
				Params[0].ParamComment = fmt.Sprintf("* @param[in] %s - size of the buffer (including trailing 0)", Params[0].ParamName)

				Params[1].ParamType = fmt.Sprintf("*const %s", param.ParamClass)
				Params[1].ParamName = toSnakeCase(param.ParamName) + "_buffer"
				Params[1].ParamComment = fmt.Sprintf("* @param[in] %s - %s buffer of %s, may be NULL", Params[0].ParamName, param.ParamClass, param.ParamDescription)

				return Params, nil
			}
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
	BasicType := false
	switch ParamTypeName {
	case "uint8":
		RustParamTypeName = "u8"
		BasicType = true
	case "uint16":
		RustParamTypeName = "u16"
		BasicType = true
	case "uint32":
		RustParamTypeName = "u32"
		BasicType = true
	case "uint64":
		RustParamTypeName = "u64"
		BasicType = true
	case "int8":
		RustParamTypeName = "i8"
		BasicType = true
	case "int16":
		RustParamTypeName = "i16"
		BasicType = true
	case "int32":
		RustParamTypeName = "i32"
		BasicType = true
	case "int64":
		RustParamTypeName = "i64"
		BasicType = true
	case "bool":
		if isPlain {
			RustParamTypeName = "u8"
		} else {
			RustParamTypeName = "bool"
		}
		BasicType = true
	case "single":
		RustParamTypeName = "f32"
		BasicType = true
	case "double":
		RustParamTypeName = "f64"
		BasicType = true
	case "pointer":
		basicParam := param
		basicParam.ParamType = param.ParamClass
		basicParam.ParamPass = "return"
		basicTypeName, err := generateRustParameterType(basicParam, isPlain)
		if err != nil {
			basicTypeName = param.ParamClass
		}
		if isPlain {
			if param.ParamPass == "in" {
				RustParamTypeName = fmt.Sprintf("*const %s", basicTypeName)
			} else {
				RustParamTypeName = fmt.Sprintf("*mut %s", basicTypeName)
			}
		} else {
			switch param.ParamPass {
			case "out":
				RustParamTypeName = fmt.Sprintf("&mut %s", basicTypeName)
			case "in":
				RustParamTypeName = fmt.Sprintf("&%s", basicTypeName)
			case "return":
				RustParamTypeName = fmt.Sprintf("%s", basicTypeName)
			}
		}
	case "string":
		if isPlain {
			RustParamTypeName = "*const c_char"
		} else {
			switch param.ParamPass {
			case "out":
				RustParamTypeName = "&mut String"
			case "in":
				RustParamTypeName = "&str"
			case "return":
				RustParamTypeName = "String"
			}
		}

	case "enum":
		if isPlain {
			RustParamTypeName = fmt.Sprintf("u16")
		} else {
			RustParamTypeName = ParamClass
		}
		BasicType = true
	case "functiontype":
		RustParamTypeName = fmt.Sprintf("%s", ParamClass)
		BasicType = true
	case "struct":
		if isPlain {
			if param.ParamPass == "in" {
				RustParamTypeName = fmt.Sprintf("*const %s", ParamClass)
			} else {
				RustParamTypeName = fmt.Sprintf("*mut %s", ParamClass)
			}
		} else {
			switch param.ParamPass {
			case "out":
				RustParamTypeName = fmt.Sprintf("&mut %s", ParamClass)
			case "in":
				RustParamTypeName = fmt.Sprintf("&%s", ParamClass)
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
			RustParamTypeName = fmt.Sprintf("%sHandle", ParamClass)
			BasicType = true
		} else {
			switch param.ParamPass {
			case "out":
				RustParamTypeName = fmt.Sprintf("&mut dyn %s", ParamClass)
			case "in":
				RustParamTypeName = fmt.Sprintf("& dyn %s", ParamClass)
			case "return":
				RustParamTypeName = fmt.Sprintf("Box<dyn %s>", ParamClass)
			}
		}

	default:
		return "", fmt.Errorf("invalid parameter type \"%s\" for Rust parameter", ParamTypeName)
	}
	if BasicType {
		if param.ParamPass == "out" {
			if isPlain {
				RustParamOutTypeName := fmt.Sprintf("*mut %s", RustParamTypeName)
				return RustParamOutTypeName, nil
			} else {
				RustParamOutTypeName := fmt.Sprintf("&mut %s", RustParamTypeName)
				return RustParamOutTypeName, nil
			}
		}
		if param.ParamPass == "return" {
			if isPlain {
				RustParamOutTypeName := fmt.Sprintf("*mut %s", RustParamTypeName)
				return RustParamOutTypeName, nil
			}
		}
	}
	return RustParamTypeName, nil
}
