/*++

Copyright (C) 2018 Autodesk Inc. (Original Author)

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
// buildbindingcsharp.go
// functions to generate C#-bindings of a library's API in form of dynamically loaded functions
// handles.
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"path"
	"strings"
)

// BuildBindingCSharp builds CSharp bindings of a library's API in form of dynamically loaded functions
// handles.
func BuildBindingCSharp(component ComponentDefinition, outputFolder string, outputFolderExample string, indentString string) error {
	namespace := component.NameSpace
	baseName := component.BaseName

	CSharpImpl := path.Join(outputFolder, namespace+".cs")
	log.Printf("Creating \"%s\"", CSharpImpl)
	CSharpImplFile, err := CreateLanguageFile(CSharpImpl, indentString)
	if err != nil {
		return err
	}

	err = buildBindingCSharpImplementation(component, CSharpImplFile, namespace, baseName)
	return err
}

func getCSharpParameterType(ParamTypeName string, NameSpace string, ParamClass string, isPlain bool) (string, error) {
	CSharpParamTypeName := ""
	switch ParamTypeName {
	case "uint8":
		CSharpParamTypeName = "Byte"

	case "uint16":
		CSharpParamTypeName = "UInt16"

	case "uint32":
		CSharpParamTypeName = "UInt32"

	case "uint64":
		CSharpParamTypeName = "UInt64"

	case "int8":
		CSharpParamTypeName = "Int8"

	case "int16":
		CSharpParamTypeName = "Int16"

	case "int32":
		CSharpParamTypeName = "Int32"

	case "int64":
		CSharpParamTypeName = "Int64"

	case "bool":
		if isPlain {
			CSharpParamTypeName = "Int32"
		} else {
			CSharpParamTypeName = "bool"
		}

	case "single":
		CSharpParamTypeName = "Single"

	case "double":
		CSharpParamTypeName = "Double"

	case "pointer":
		CSharpParamTypeName = "UInt64"

	case "string":
		if isPlain {
			CSharpParamTypeName = "byte[]"
		} else {
			CSharpParamTypeName = "String"
		}

	case "enum":
		if isPlain {
			CSharpParamTypeName = "Int32"
		} else {
			CSharpParamTypeName = "e" + ParamClass
		}

	case "functiontype":
		CSharpParamTypeName = fmt.Sprintf("IntPtr")

	case "struct":
		if isPlain {
			CSharpParamTypeName = "internal" + ParamClass
		} else {
			CSharpParamTypeName = "s" + ParamClass
		}

	case "basicarray":
		CSharpParamTypeName = fmt.Sprintf("IntPtr")

	case "structarray":
		CSharpParamTypeName = fmt.Sprintf("IntPtr")

	case "class":
		if isPlain {
			CSharpParamTypeName = "IntPtr"
		} else {
			CSharpParamTypeName = "C" + ParamClass
		}

	default:
	}

	return CSharpParamTypeName, nil
}

func getCSharpPlainParameters(method ComponentDefinitionMethod, NameSpace string, ClassName string, isGlobal bool) (string, error) {
	parameters := ""

	for k := 0; k < len(method.Params); k++ {
		param := method.Params[k]
		ParamTypeName, err := getCSharpParameterType(param.ParamType, NameSpace, param.ParamClass, true)
		if err != nil {
			return "", err
		}

		switch param.ParamPass {
		case "in":
			if parameters != "" {
				parameters = parameters + ", "
			}
			parameters = parameters + ParamTypeName + " A" + param.ParamName

		case "out", "return":
			if parameters != "" {
				parameters = parameters + ", "
			}

			switch param.ParamType {
			case "string":
				parameters = parameters + fmt.Sprintf("UInt32 size%s, out UInt32 needed%s, IntPtr data%s", param.ParamName, param.ParamName, param.ParamName)

			default:
				parameters = parameters + "out " + ParamTypeName + " A" + param.ParamName
			}

		}
	}

	return parameters, nil
}

func getCSharpClassParameters(method ComponentDefinitionMethod, NameSpace string, ClassName string, isGlobal bool) (string, string, error) {
	parameters := ""
	returnType := ""

	for k := 0; k < len(method.Params); k++ {
		param := method.Params[k]
		ParamTypeName, err := getCSharpParameterType(param.ParamType, NameSpace, param.ParamClass, false)
		if err != nil {
			return "", "", err
		}

		switch param.ParamPass {
		case "in":
			if parameters != "" {
				parameters = parameters + ", "
			}
			parameters = parameters + ParamTypeName + " A" + param.ParamName

		case "out":
			if parameters != "" {
				parameters = parameters + ", "
			}
			parameters = parameters + "out " + ParamTypeName + " A" + param.ParamName

		case "return":
			if returnType != "" {
				return "", "", fmt.Errorf("duplicate return value \"%s\" for Pascal method \"%s\"", param.ParamName, method.MethodName)
			}
			returnType = ParamTypeName
		}
	}

	if returnType == "" {
		returnType = "void"
	}

	return parameters, returnType, nil
}

func writeCSharpClassMethodImplementation(method ComponentDefinitionMethod, w LanguageWriter, NameSpace string, ClassName string, isGlobal bool, spacing string) error {

	defineCommands := make([]string, 0)
	initCommands := make([]string, 0)
	resultCommands := make([]string, 0)
	postInitCommands := make([]string, 0)

	doInitCall := false

	callFunctionName := ""
	callFunctionParameters := ""
	initCallParameters := ""

	if isGlobal {
		callFunctionName = fmt.Sprintf("%s", method.MethodName)
	} else {
		callFunctionName = fmt.Sprintf("%s_%s", ClassName, method.MethodName)
		callFunctionParameters = "Handle"
	}

	initCallParameters = callFunctionParameters

	for k := 0; k < len(method.Params); k++ {
		param := method.Params[k]
		ParamTypeName, err := getCSharpParameterType(param.ParamType, NameSpace, param.ParamClass, false)
		if err != nil {
			return err
		}

		callFunctionParameter := ""
		initCallParameter := ""

		switch param.ParamPass {
		case "in":

			switch param.ParamType {
			case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
				callFunctionParameter = "A" + param.ParamName
				initCallParameter = callFunctionParameter

			case "single":
				callFunctionParameter = "A" + param.ParamName
				initCallParameter = callFunctionParameter

			case "double":
				callFunctionParameter = "A" + param.ParamName
				initCallParameter = callFunctionParameter

			case "pointer":
				callFunctionParameter = "A" + param.ParamName
				initCallParameter = callFunctionParameter

			case "string":
				defineCommands = append(defineCommands, fmt.Sprintf("  byte[] byte%s = Encoding.UTF8.GetBytes(A%s + char.MinValue);", param.ParamName, param.ParamName))
				callFunctionParameter = "byte" + param.ParamName
				initCallParameter = callFunctionParameter

			case "enum":
				defineCommands = append(defineCommands, fmt.Sprintf("  Int32 enum%s = (Int32) A%s;", param.ParamName, param.ParamName))
				callFunctionParameter = "enum" + param.ParamName
				initCallParameter = callFunctionParameter

			case "bool":
				callFunctionParameter = "( A" + param.ParamName + " ? 1 : 0 )"
				initCallParameter = callFunctionParameter

			case "struct":
				defineCommands = append(defineCommands, fmt.Sprintf("  Internal.internal%s int%s = Internal.%sWrapper.convertStructToInternal_%s (A%s);", param.ParamClass, param.ParamName, NameSpace, param.ParamClass, param.ParamName))
				callFunctionParameter = "int" + param.ParamName
				initCallParameter = callFunctionParameter

			case "basicarray":
				callFunctionParameter = "IntPtr.Zero"
				initCallParameter = callFunctionParameter

			case "structarray":
				callFunctionParameter = "IntPtr.Zero"
				initCallParameter = callFunctionParameter

			case "functiontype":
				callFunctionParameter = "IntPtr.Zero"
				initCallParameter = callFunctionParameter

			case "class":
				callFunctionParameter = "A" + param.ParamName + ".GetHandle()"
				initCallParameter = callFunctionParameter

			default:
				return fmt.Errorf("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, ClassName, method.MethodName, param.ParamName)
			}

		case "out":

			switch param.ParamType {
			case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":

				callFunctionParameter = "out A" + param.ParamName
				initCallParameter = callFunctionParameter

			case "single":
				callFunctionParameter = "out A" + param.ParamName
				initCallParameter = callFunctionParameter

			case "double":
				callFunctionParameter = "out A" + param.ParamName
				initCallParameter = callFunctionParameter

			case "pointer":
				defineCommands = append(defineCommands, fmt.Sprintf("  %s result%s = 0;", ParamTypeName, param.ParamName))
				callFunctionParameter = "out result" + param.ParamName
				resultCommands = append(resultCommands, fmt.Sprintf("  A%s = result%s;", param.ParamName, param.ParamName))
				initCallParameter = callFunctionParameter

			case "string":

				initCommands = append(initCommands, fmt.Sprintf("  UInt32 size%s = 0;", param.ParamName))
				initCommands = append(initCommands, fmt.Sprintf("  UInt32 needed%s = 0;", param.ParamName))

				initCallParameter = fmt.Sprintf("size%s, out needed%s, IntPtr.Zero", param.ParamName, param.ParamName)

				postInitCommands = append(postInitCommands, fmt.Sprintf("  size%s = needed%s + 1;", param.ParamName, param.ParamName))
				postInitCommands = append(postInitCommands, fmt.Sprintf("  byte[] bytes%s = new byte[size%s];", param.ParamName, param.ParamName))
				postInitCommands = append(postInitCommands, fmt.Sprintf("  GCHandle data%s = GCHandle.Alloc(bytes%s, GCHandleType.Pinned);", param.ParamName, param.ParamName))

				callFunctionParameter = fmt.Sprintf("size%s, out needed%s, data%s.AddrOfPinnedObject()", param.ParamName, param.ParamName, param.ParamName)

				resultCommands = append(resultCommands, fmt.Sprintf("  data%s.Free();", param.ParamName))
				resultCommands = append(resultCommands, fmt.Sprintf("  A%s = Encoding.UTF8.GetString(bytes%s).TrimEnd(char.MinValue);", param.ParamName, param.ParamName))

				doInitCall = true

			case "enum":
				defineCommands = append(defineCommands, fmt.Sprintf("  Int32 result%s = 0;", param.ParamName))
				callFunctionParameter = "out result" + param.ParamName
				initCallParameter = callFunctionParameter
				resultCommands = append(resultCommands, fmt.Sprintf("  A%s = (e%s) (result%s);", param.ParamName, param.ParamClass, param.ParamName))

			case "bool":
				defineCommands = append(defineCommands, fmt.Sprintf("  Int32 result%s = 0;", param.ParamName))
				callFunctionParameter = "out result" + param.ParamName
				initCallParameter = callFunctionParameter
				resultCommands = append(resultCommands, fmt.Sprintf("  A%s = (result%s != 0);", param.ParamName, param.ParamName))

			case "struct":
				defineCommands = append(defineCommands, fmt.Sprintf("  Internal.internal%s intresult%s;", param.ParamClass, param.ParamName))
				callFunctionParameter = "out intresult" + param.ParamName
				initCallParameter = callFunctionParameter
				resultCommands = append(resultCommands, fmt.Sprintf("  A%s = Internal.%sWrapper.convertInternalToStruct_%s (intresult%s);", param.ParamName, NameSpace, param.ParamClass, param.ParamName))

			case "basicarray", "structarray":

				defineCommands = append(defineCommands, fmt.Sprintf("  IntPtr result%s = IntPtr.Zero;", param.ParamName))
				callFunctionParameter = "out result" + param.ParamName
				initCallParameter = callFunctionParameter
				resultCommands = append(resultCommands, fmt.Sprintf("  A%s = result%s;", param.ParamName, param.ParamName))

				/*defineCommands = append (defineCommands, "  countNeeded" + param.ParamName + ": QWord;");
				defineCommands = append (defineCommands, "  countWritten" + param.ParamName + ": QWord;");
				initCommands = append (initCommands, "  countNeeded" + param.ParamName + ":= 0;");
				initCommands = append (initCommands, "  countWritten" + param.ParamName + ":= 0;");

				initCallParameters = initCallParameters + fmt.Sprintf("0, countNeeded%s, nil", param.ParamName)

				postInitCommands = append (postInitCommands, fmt.Sprintf("  SetLength (A%s, countNeeded%s);", param.ParamName, param.ParamName));

				callFunctionParameters = callFunctionParameters + fmt.Sprintf("countNeeded%s, countWritten%s, @A%s[0]", param.ParamName, param.ParamName, param.ParamName)

				doInitCall = true; */

			case "class":
				defineCommands = append(defineCommands, fmt.Sprintf("  IntPtr new%s = IntPtr.Zero;", param.ParamName))
				callFunctionParameter = "out new" + param.ParamName
				initCallParameter = callFunctionParameter
				resultCommands = append(resultCommands, fmt.Sprintf("  A%s = new C%s (new%s );", param.ParamName, param.ParamClass, param.ParamName))

			default:
				return fmt.Errorf("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, ClassName, method.MethodName, param.ParamName)
			}

		case "return":

			switch param.ParamType {
			case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "single", "double", "pointer":

				defineCommands = append(defineCommands, fmt.Sprintf("  %s result%s = 0;", ParamTypeName, param.ParamName))
				callFunctionParameter = "out result" + param.ParamName
				initCallParameter = callFunctionParameter
				resultCommands = append(resultCommands, fmt.Sprintf("  return result%s;", param.ParamName))

			case "string":

				initCommands = append(initCommands, fmt.Sprintf("  UInt32 size%s = 0;", param.ParamName))
				initCommands = append(initCommands, fmt.Sprintf("  UInt32 needed%s = 0;", param.ParamName))

				initCallParameter = fmt.Sprintf("size%s, out needed%s, IntPtr.Zero", param.ParamName, param.ParamName)

				postInitCommands = append(postInitCommands, fmt.Sprintf("  size%s = needed%s + 1;", param.ParamName, param.ParamName))
				postInitCommands = append(postInitCommands, fmt.Sprintf("  byte[] bytes%s = new byte[size%s];", param.ParamName, param.ParamName))
				postInitCommands = append(postInitCommands, fmt.Sprintf("  GCHandle data%s = GCHandle.Alloc(bytes%s, GCHandleType.Pinned);", param.ParamName, param.ParamName))

				callFunctionParameter = fmt.Sprintf("size%s, out needed%s, data%s.AddrOfPinnedObject()", param.ParamName, param.ParamName, param.ParamName)

				resultCommands = append(resultCommands, fmt.Sprintf("  data%s.Free();", param.ParamName))
				resultCommands = append(resultCommands, fmt.Sprintf("  return Encoding.UTF8.GetString(bytes%s).TrimEnd(char.MinValue);", param.ParamName))

				doInitCall = true

			case "enum":
				defineCommands = append(defineCommands, fmt.Sprintf("  Int32 result%s = 0;", param.ParamName))
				callFunctionParameter = "out result" + param.ParamName
				initCallParameter = callFunctionParameter
				resultCommands = append(resultCommands, fmt.Sprintf("  return (e%s) (result%s);", param.ParamClass, param.ParamName))

			case "bool":
				defineCommands = append(defineCommands, fmt.Sprintf("  Int32 result%s = 0;", param.ParamName))
				callFunctionParameter = "out result" + param.ParamName
				initCallParameter = callFunctionParameter
				resultCommands = append(resultCommands, fmt.Sprintf("  return (result%s != 0);", param.ParamName))

			case "struct":
				defineCommands = append(defineCommands, fmt.Sprintf("  Internal.internal%s intresult%s;", param.ParamClass, param.ParamName))
				callFunctionParameter = "out intresult" + param.ParamName
				initCallParameter = callFunctionParameter
				resultCommands = append(resultCommands, fmt.Sprintf("  return Internal.%sWrapper.convertInternalToStruct_%s (intresult%s);", NameSpace, param.ParamClass, param.ParamName))

			case "basicarray", "structarray":
				/*defineCommands = append (defineCommands, "  countNeeded" + param.ParamName + ": QWord;");
				defineCommands = append (defineCommands, "  countWritten" + param.ParamName + ": QWord;");
				initCommands = append (initCommands, "  countNeeded" + param.ParamName + ":= 0;");
				initCommands = append (initCommands, "  countWritten" + param.ParamName + ":= 0;");

				initCallParameters = initCallParameters + fmt.Sprintf("0, countNeeded%s, nil", param.ParamName)

				postInitCommands = append (postInitCommands, fmt.Sprintf("  SetLength (Result, countNeeded%s);", param.ParamName));

				callFunctionParameters = callFunctionParameters + fmt.Sprintf("countNeeded%s, countWritten%s, @Result[0]", param.ParamName, param.ParamName)

				doInitCall = true; */

			case "class":

				defineCommands = append(defineCommands, fmt.Sprintf("  IntPtr new%s = IntPtr.Zero;", param.ParamName))
				callFunctionParameter = "out new" + param.ParamName
				initCallParameter = callFunctionParameter
				resultCommands = append(resultCommands, fmt.Sprintf("  return new C%s (new%s );", param.ParamClass, param.ParamName))

			default:
				return fmt.Errorf("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, ClassName, method.MethodName, param.ParamName)
			}

		}

		if callFunctionParameters != "" {
			callFunctionParameters = callFunctionParameters + ", "
		}

		if initCallParameters != "" {
			initCallParameters = initCallParameters + ", "
		}

		callFunctionParameters = callFunctionParameters + callFunctionParameter
		initCallParameters = initCallParameters + initCallParameter

	}

	if len(defineCommands) > 0 {
		w.Writelns(spacing, defineCommands)
	}

	if len(initCommands) > 0 {
		w.Writelns(spacing, initCommands)
	}

	if doInitCall {
		w.Writeln(spacing+"  CheckError (Internal.%sWrapper.%s (%s));", NameSpace, callFunctionName, initCallParameters)
	}

	w.Writelns(spacing, postInitCommands)

	w.Writeln("")

	w.Writeln(spacing+"  CheckError (Internal.%sWrapper.%s (%s));", NameSpace, callFunctionName, callFunctionParameters)

	w.Writelns(spacing, resultCommands)

	return nil
}

func buildBindingCSharpImplementation(component ComponentDefinition, w LanguageWriter, NameSpace string, BaseName string) error {

	baseName := component.BaseName
	global := component.Global

	CSharpBaseClassName := "C" + component.Global.BaseClassName
	w.Writeln("using System;")
	w.Writeln("using System.Text;")
	w.Writeln("using System.Runtime.InteropServices;")
	w.Writeln("")

	w.Writeln("namespace %s {", NameSpace)
	w.Writeln("")

	for i := 0; i < len(component.Enums); i++ {
		enum := component.Enums[i]
		w.Writeln("  public enum e%s {", enum.Name)

		for j := 0; j < len(enum.Options); j++ {
			option := enum.Options[j]
			commavalue := ""
			if j < (len(enum.Options) - 1) {
				commavalue = ","
			}

			w.Writeln("    %s = %d%s", option.Name, option.Value, commavalue)
		}

		w.Writeln("  };")
		w.Writeln("")

	}

	for i := 0; i < len(component.Structs); i++ {
		structinfo := component.Structs[i]

		w.Writeln("  public struct s%s", structinfo.Name)
		w.Writeln("  {")

		for j := 0; j < len(structinfo.Members); j++ {
			element := structinfo.Members[j]

			arraysuffix := ""
			if element.Rows > 0 {
				if element.Columns > 0 {
					arraysuffix = fmt.Sprintf("[][]")
				} else {
					arraysuffix = fmt.Sprintf("[]")
				}
			}

			switch element.Type {
			case "uint8":
				w.Writeln("    public Byte%s %s;", arraysuffix, element.Name)
			case "uint16":
				w.Writeln("    public UInt16%s %s;", arraysuffix, element.Name)
			case "uint32":
				w.Writeln("    public UInt32%s %s;", arraysuffix, element.Name)
			case "uint64":
				w.Writeln("    public UInt64%s %s;", arraysuffix, element.Name)
			case "int8":
				w.Writeln("    public Int8%s %s;", arraysuffix, element.Name)
			case "int16":
				w.Writeln("    public Int16%s %s;", arraysuffix, element.Name)
			case "int32":
				w.Writeln("    public Int32%s %s;", arraysuffix, element.Name)
			case "int64":
				w.Writeln("    public Int64%s %s;", arraysuffix, element.Name)
			case "bool":
				w.Writeln("    public bool%s %s;", arraysuffix, element.Name)
			case "single":
				w.Writeln("    public Single%s %s;", arraysuffix, element.Name)
			case "double":
				w.Writeln("    public Double%s %s;", arraysuffix, element.Name)
			case "pointer":
				w.Writeln("    public UInt64%s %s;", arraysuffix, element.Name)
			case "string":
				return fmt.Errorf("it is not possible for struct s%s%s to contain a string value", NameSpace, structinfo.Name)
			case "class":
				return fmt.Errorf("it is not possible for struct s%s%s to contain a handle value", NameSpace, structinfo.Name)
			case "enum":
				w.Writeln("    public e%s%s %s;", element.Class, arraysuffix, element.Name)
			}
		}

		w.Writeln("  }")
		w.Writeln("")
	}

	w.Writeln("")

	w.Writeln("  namespace Internal {")
	w.Writeln("")

	for i := 0; i < len(component.Structs); i++ {
		structinfo := component.Structs[i]

		w.Writeln("    [StructLayout(LayoutKind.Explicit)]")
		w.Writeln("    public unsafe struct internal%s", structinfo.Name)
		w.Writeln("    {")

		fieldOffset := 0

		for j := 0; j < len(structinfo.Members); j++ {
			element := structinfo.Members[j]

			arraysuffix := ""
			fixedtag := ""
			multiplier := 1
			if element.Rows > 0 {
				if element.Columns > 0 {
					multiplier = element.Rows * element.Columns
					arraysuffix = fmt.Sprintf("[%d]", multiplier)
				} else {
					multiplier = element.Rows
					arraysuffix = fmt.Sprintf("[%d]", multiplier)
				}

				fixedtag = "fixed "
			}

			switch element.Type {
			case "uint8":
				w.Writeln("      [FieldOffset(%d)] public %sByte %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 1*multiplier
			case "uint16":
				w.Writeln("      [FieldOffset(%d)] public %sUInt16 %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 2*multiplier
			case "uint32":
				w.Writeln("      [FieldOffset(%d)] public %sUInt32 %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 4*multiplier
			case "uint64":
				w.Writeln("      [FieldOffset(%d)] public %sUInt64 %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 8*multiplier
			case "int8":
				w.Writeln("      [FieldOffset(%d)] public %sInt8 %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 1*multiplier
			case "int16":
				w.Writeln("      [FieldOffset(%d)] public %sInt16 %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 2*multiplier
			case "int32":
				w.Writeln("      [FieldOffset(%d)] public %sInt32 %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 4*multiplier
			case "int64":
				w.Writeln("      [FieldOffset(%d)] public %sInt64 %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 8*multiplier
			case "bool":
				w.Writeln("      [FieldOffset(%d)] public %sByte %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 1*multiplier
			case "single":
				w.Writeln("      [FieldOffset(%d)] public %sSingle %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 4*multiplier
			case "double":
				w.Writeln("      [FieldOffset(%d)] public %sDouble %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 8*multiplier
			case "pointer":
				w.Writeln("      [FieldOffset(%d)] public %sUInt64 %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 8*multiplier
			case "string":
				return fmt.Errorf("it is not possible for struct s%s%s to contain a string value", NameSpace, structinfo.Name)
			case "class":
				return fmt.Errorf("it is not possible for struct s%s%s to contain a handle value", NameSpace, structinfo.Name)
			case "enum":
				w.Writeln("      [FieldOffset(%d)] public %sInt32 %s%s;", fieldOffset, fixedtag, element.Name, arraysuffix)
				fieldOffset = fieldOffset + 4*multiplier
			}
		}

		w.Writeln("    }")
		w.Writeln("")
	}

	w.Writeln("")

	w.Writeln("    public class %sWrapper", NameSpace)
	w.Writeln("    {")

	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]

		for j := 0; j < len(class.Methods); j++ {
			method := class.Methods[j]

			parameters, err := getCSharpPlainParameters(method, NameSpace, class.ClassName, false)
			if err != nil {
				return err
			}

			w.Writeln("      [DllImport(\"%s.dll\", EntryPoint = \"%s_%s_%s\", CallingConvention=CallingConvention.Cdecl)]", baseName, strings.ToLower(NameSpace), strings.ToLower(class.ClassName), strings.ToLower(method.MethodName))

			if parameters == "" {
				parameters = "IntPtr Handle"
			} else {
				parameters = "IntPtr Handle, " + parameters
			}

			w.Writeln("      public unsafe extern static Int32 %s_%s (%s);", class.ClassName, method.MethodName, parameters)
			w.Writeln("")

		}

	}

	for j := 0; j < len(global.Methods); j++ {
		method := global.Methods[j]

		parameters, err := getCSharpPlainParameters(method, NameSpace, "", true)
		if err != nil {
			return err
		}

		w.Writeln("      [DllImport(\"%s.dll\", EntryPoint = \"%s_%s\", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]", baseName, strings.ToLower(NameSpace), strings.ToLower(method.MethodName))
		w.Writeln("      public extern static Int32 %s (%s);", method.MethodName, parameters)
		w.Writeln("")
	}

	for i := 0; i < len(component.Structs); i++ {
		structinfo := component.Structs[i]

		w.Writeln("      public unsafe static s%s convertInternalToStruct_%s (internal%s int%s)", structinfo.Name, structinfo.Name, structinfo.Name, structinfo.Name)
		w.Writeln("      {")
		w.Writeln("        s%s %s;", structinfo.Name, structinfo.Name)

		for j := 0; j < len(structinfo.Members); j++ {
			element := structinfo.Members[j]

			paramType, err := getCSharpParameterType(element.Type, NameSpace, element.Class, false)
			if err != nil {
				return err
			}

			castPrefix := ""
			castSuffix := ""
			switch element.Type {
			case "bool":
				castSuffix = " != 0"
			case "enum":
				castPrefix = fmt.Sprintf("(e%s) ", element.Class)
			}

			if element.Rows > 0 {
				if element.Columns > 0 {
					w.Writeln("        %s.%s = new %s[%d][];", structinfo.Name, element.Name, paramType, element.Columns)
					w.Writeln("        for (int colIndex = 0; colIndex < %d; colIndex++) {", element.Columns)
					w.Writeln("          %s.%s[colIndex] = new %s[%d];", structinfo.Name, element.Name, paramType, element.Rows)
					w.Writeln("          for (int rowIndex = 0; rowIndex < %d; rowIndex++) {", element.Rows)
					w.Writeln("            %s.%s[colIndex][rowIndex] = %sint%s.%s%s[colIndex * %d + rowIndex];", structinfo.Name, element.Name, castPrefix, structinfo.Name, element.Name, castSuffix, element.Rows)
					w.Writeln("          }")
					w.Writeln("        }")
					w.Writeln("")
				} else {
					w.Writeln("        %s.%s = new %s[%d];", structinfo.Name, element.Name, paramType, element.Rows)
					w.Writeln("        for (int rowIndex = 0; rowIndex < %d; rowIndex++) {", element.Rows)
					w.Writeln("          %s.%s[rowIndex] = %sint%s.%s%s[rowIndex];", structinfo.Name, element.Name, castPrefix, structinfo.Name, element.Name, castSuffix)
					w.Writeln("        }")
					w.Writeln("")
				}
			} else {
				w.Writeln("        %s.%s = %sint%s.%s%s;", structinfo.Name, element.Name, castPrefix, structinfo.Name, element.Name, castSuffix)
			}
		}

		w.Writeln("        return %s;", structinfo.Name)
		w.Writeln("      }")
		w.Writeln("")

		w.Writeln("      public unsafe static internal%s convertStructToInternal_%s (s%s %s)", structinfo.Name, structinfo.Name, structinfo.Name, structinfo.Name)
		w.Writeln("      {")
		w.Writeln("        internal%s int%s;", structinfo.Name, structinfo.Name)

		for j := 0; j < len(structinfo.Members); j++ {
			element := structinfo.Members[j]

			castPrefix := ""
			castSuffix := ""
			switch element.Type {
			case "bool":
				castSuffix = " (int)"
			case "enum":
				castPrefix = fmt.Sprintf("(Int32) ")
			}

			if element.Rows > 0 {
				if element.Columns > 0 {
					w.Writeln("        for (int colIndex = 0; colIndex < %d; colIndex++) {", element.Columns)
					w.Writeln("          for (int rowIndex = 0; rowIndex < %d; rowIndex++) {", element.Rows)
					w.Writeln("            int%s.%s[colIndex * %d + rowIndex] = %s%s.%s[colIndex][rowIndex]%s;", structinfo.Name, element.Name, element.Rows, castPrefix, structinfo.Name, element.Name, castSuffix)
					w.Writeln("          }")
					w.Writeln("        }")
					w.Writeln("")
				} else {
					w.Writeln("        for (int rowIndex = 0; rowIndex < %d; rowIndex++) {", element.Rows)
					w.Writeln("          int%s.%s[rowIndex] = %s%s.%s%s[rowIndex];", structinfo.Name, element.Name, castPrefix, structinfo.Name, element.Name, castSuffix)
					w.Writeln("        }")
					w.Writeln("")
				}
			} else {
				w.Writeln("        int%s.%s = %s%s.%s%s;", structinfo.Name, element.Name, castPrefix, structinfo.Name, element.Name, castSuffix)
			}

		}

		w.Writeln("        return int%s;", structinfo.Name)
		w.Writeln("      }")
		w.Writeln("")
	}

	w.Writeln("      public static void ThrowError(IntPtr Handle, Int32 errorCode)")
	w.Writeln("      {")
	w.Writeln("        String sMessage = \"%s Error\";", NameSpace)

	if len(component.Global.ErrorMethod) > 0 {
		w.Writeln("        if (Handle != IntPtr.Zero) {")
		w.Writeln("          UInt32 sizeMessage = 0;")
		w.Writeln("          UInt32 neededMessage = 0;")
		w.Writeln("          Int32 hasLastError = 0;")
		w.Writeln("          Int32 resultCode1 = %s (Handle, sizeMessage, out neededMessage, IntPtr.Zero, out hasLastError);", component.Global.ErrorMethod)
		w.Writeln("          if ((resultCode1 == 0) && (hasLastError != 0)) {")
		w.Writeln("            sizeMessage = neededMessage + 1;")
		w.Writeln("            byte[] bytesMessage = new byte[sizeMessage];")
		w.Writeln("")
		w.Writeln("            GCHandle dataMessage = GCHandle.Alloc(bytesMessage, GCHandleType.Pinned);")
		w.Writeln("            Int32 resultCode2 = %s(Handle, sizeMessage, out neededMessage, dataMessage.AddrOfPinnedObject(), out hasLastError);", component.Global.ErrorMethod)
		w.Writeln("            dataMessage.Free();")
		w.Writeln("")
		w.Writeln("            if ((resultCode2 == 0) && (hasLastError != 0)) {")
		w.Writeln("              sMessage = sMessage + \": \" + Encoding.UTF8.GetString(bytesMessage).TrimEnd(char.MinValue);")
		w.Writeln("            }")
		w.Writeln("          }")
		w.Writeln("        }")
		w.Writeln("")
	}
	w.Writeln("        throw new Exception(sMessage + \"(# \" + errorCode + \")\");")
	w.Writeln("      }")
	w.Writeln("")

	w.Writeln("    }")
	w.Writeln("  }")

	w.Writeln("")
	w.Writeln("")

	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]

		CSharpParentClassName := ""
		if !component.isBaseClass(class) {
			if class.ParentClass == "" {
				CSharpParentClassName = ": " + CSharpBaseClassName
			} else {
				CSharpParentClassName = ": C" + class.ParentClass
			}
		}

		w.Writeln("  class C%s %s", class.ClassName, CSharpParentClassName)
		w.Writeln("  {")

		if component.isBaseClass(class) {
			w.Writeln("    protected IntPtr Handle;")
			w.Writeln("")
			w.Writeln("    public C%s (IntPtr NewHandle)", class.ClassName)
			w.Writeln("    {")
			w.Writeln("      Handle = NewHandle;")
			w.Writeln("    }")
			w.Writeln("")
			w.Writeln("    ~C%s ()", class.ClassName)
			w.Writeln("    {")
			w.Writeln("      if (Handle != IntPtr.Zero) {")
			w.Writeln("        Internal.%sWrapper.%s (Handle);", NameSpace, component.Global.ReleaseMethod)
			w.Writeln("        Handle = IntPtr.Zero;")
			w.Writeln("      }")
			w.Writeln("    }")
			w.Writeln("")

			w.Writeln("    protected void CheckError (Int32 errorCode)")
			w.Writeln("    {")
			w.Writeln("      if (errorCode != 0) {")
			w.Writeln("        Internal.%sWrapper.ThrowError (Handle, errorCode);", NameSpace)
			w.Writeln("      }")
			w.Writeln("    }")
			w.Writeln("")

			w.Writeln("    public IntPtr GetHandle ()")
			w.Writeln("    {")
			w.Writeln("      return Handle;")
			w.Writeln("    }")
			w.Writeln("")

		} else {
			w.Writeln("    public C%s (IntPtr NewHandle) : base (NewHandle)", class.ClassName)
			w.Writeln("    {")
			w.Writeln("    }")
			w.Writeln("")
		}

		for j := 0; j < len(class.Methods); j++ {
			method := class.Methods[j]

			parameters, returnType, err := getCSharpClassParameters(method, NameSpace, class.ClassName, false)
			if err != nil {
				return err
			}

			w.Writeln("    public %s %s (%s)", returnType, method.MethodName, parameters)
			w.Writeln("    {")

			writeCSharpClassMethodImplementation(method, w, NameSpace, class.ClassName, false, "    ")

			w.Writeln("    }")
			w.Writeln("")
		}

		w.Writeln("  }")
		w.Writeln("")
	}

	w.Writeln("  class Wrapper")
	w.Writeln("  {")

	w.Writeln("    private static void CheckError (Int32 errorCode)")
	w.Writeln("    {")
	w.Writeln("      if (errorCode != 0) {")
	w.Writeln("        Internal.%sWrapper.ThrowError (IntPtr.Zero, errorCode);", NameSpace)
	w.Writeln("      }")
	w.Writeln("    }")
	w.Writeln("")

	for j := 0; j < len(global.Methods); j++ {
		method := global.Methods[j]

		parameters, returnType, err := getCSharpClassParameters(method, NameSpace, "", true)
		if err != nil {
			return err
		}

		w.Writeln("    public static %s %s (%s)", returnType, method.MethodName, parameters)
		w.Writeln("    {")

		writeCSharpClassMethodImplementation(method, w, NameSpace, "Wrapper", true, "    ")

		w.Writeln("    }")
		w.Writeln("")
	}

	w.Writeln("  }")
	w.Writeln("")

	w.Writeln("}")

	return nil
}
