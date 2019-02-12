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
// languagepascal.go
// functions to generate the Pascal-layer of a library's API (can be used in bindings or implementation)
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"
)



func writePascalBaseTypeDefinitions(componentdefinition ComponentDefinition, w LanguageWriter, NameSpace string, BaseName string) error {

	w.Writeln ("(*************************************************************************************************************************");
	w.Writeln (" Version definition for %s", NameSpace);
	w.Writeln ("**************************************************************************************************************************)");
	w.Writeln ("");
	w.Writeln ("const");
	w.Writeln ("  %s_VERSION_MAJOR = %d;", strings.ToUpper (NameSpace), majorVersion(componentdefinition.Version));
	w.Writeln ("  %s_VERSION_MINOR = %d;", strings.ToUpper (NameSpace), minorVersion(componentdefinition.Version));
	w.Writeln ("  %s_VERSION_MICRO = %d;", strings.ToUpper (NameSpace), microVersion(componentdefinition.Version));
	w.Writeln ("");
	w.Writeln ("");
	
	w.Writeln ("(*************************************************************************************************************************");
	w.Writeln (" General type definitions");
	w.Writeln ("**************************************************************************************************************************)");
	w.Writeln ("")
	
	w.Writeln ("type");
	w.Writeln ("  T%sResult = Cardinal;", NameSpace);
	w.Writeln ("  T%sHandle = Pointer;", NameSpace);
	w.Writeln ("")
	
	w.Writeln ("  P%sResult = ^T%sResult;", NameSpace, NameSpace);
	w.Writeln ("  P%sHandle = ^T%sHandle;", NameSpace, NameSpace);
	w.Writeln ("")

	w.Writeln ("(*************************************************************************************************************************");
	w.Writeln (" Error Constants for %s", NameSpace);
	w.Writeln ("**************************************************************************************************************************)");
	w.Writeln ("");
	w.Writeln ("const");
	w.Writeln ("  %s_SUCCESS = 0;", strings.ToUpper (NameSpace));
		
	for i := 0; i < len(componentdefinition.Errors.Errors); i++ {
		errorcode := componentdefinition.Errors.Errors[i];
		w.Writeln ("  %s_ERROR_%s = %d;", strings.ToUpper (NameSpace), errorcode.Name, errorcode.Code);
	}
	w.Writeln ("");
	
	if (len(componentdefinition.Enums) > 0) {
		w.Writeln ("(*************************************************************************************************************************");
		w.Writeln (" Declaration of enums");
		w.Writeln ("**************************************************************************************************************************)");
		w.Writeln ("");
		w.Writeln ("type");	
		w.Writeln ("");

		for i := 0; i < len(componentdefinition.Enums); i++ {
			enum := componentdefinition.Enums[i];
			w.Writeln ("  T%s%s = (", NameSpace, enum.Name);
			
			for j := 0; j < len(enum.Options); j++ {			
				comma := "";
				if (j < len(enum.Options) - 1) {
					comma = ",";
				}
				option := enum.Options[j];
				w.Writeln ("    e%s%s%s", enum.Name, option.Name, comma);
			}
			
			w.Writeln ( "  );");
			w.Writeln ( "");
		}
	}
	
	if len(componentdefinition.Structs) > 0 {
		w.Writeln ("(*************************************************************************************************************************");
		w.Writeln (" Declaration of structs");
		w.Writeln ("**************************************************************************************************************************)");
		w.Writeln ("");
		w.Writeln ("type");	
		w.Writeln ("");
			
		for i := 0; i < len(componentdefinition.Structs); i++ {
			structinfo := componentdefinition.Structs[i];
			w.Writeln ( "  P%s%s = ^T%s%s;", NameSpace, structinfo.Name, NameSpace, structinfo.Name);
			w.Writeln ( "  T%s%s = packed record", NameSpace, structinfo.Name);
			
			for j := 0; j < len(structinfo.Members); j++ {
				element := structinfo.Members[j];
				arrayprefix := "";
				if (element.Rows > 0) {
					if (element.Columns > 0) {
						arrayprefix = fmt.Sprintf ("array [0..%d, 0..%d] of ", element.Columns - 1, element.Rows - 1)
					} else {
						arrayprefix = fmt.Sprintf ("array [0..%d] of ",element.Rows - 1)
					}
				}
			
				switch (element.Type) {
					case "uint8":
						w.Writeln ( "    F%s: %sByte;", element.Name, arrayprefix);
					case "uint16":
						w.Writeln ( "    F%s: %sWord;", element.Name, arrayprefix);
					case "uint32":
						w.Writeln ( "    F%s: %sCardinal;", element.Name, arrayprefix);
					case "uint64":
						w.Writeln ( "    F%s: %sQWord;", element.Name, arrayprefix);
					case "int8":
						w.Writeln ( "    F%s: %sSmallInt;", element.Name, arrayprefix);
					case "int16":
						w.Writeln ( "    F%s: %sShortInt;", element.Name, arrayprefix);
					case "int32":
						w.Writeln ( "    F%s: %sInteger;", element.Name, arrayprefix);
					case "int64":
						w.Writeln ( "    F%s: %sInt64;", element.Name, arrayprefix);
					case "bool":
						w.Writeln ( "    F%s: %sByte;", element.Name, arrayprefix);
					case "single":
						w.Writeln ( "    F%s: %sSingle;", element.Name, arrayprefix);
					case "double":
						w.Writeln ( "    F%s: %sDouble;", element.Name, arrayprefix);
					case "string":
						return fmt.Errorf ("it is not possible for struct s%s%s to contain a string value", NameSpace, structinfo.Name);
					case "handle":
						return fmt.Errorf ("it is not possible for struct s%s%s to contain a handle value", NameSpace, structinfo.Name);
					case "enum":
						w.Writeln ( "    F%s: %sInteger;", element.Name, arrayprefix);
				}
			}
			
			w.Writeln ("  end;");
			w.Writeln ("");
		}
		
		w.Writeln ( "");

		w.Writeln ("(*************************************************************************************************************************");
		w.Writeln (" Declaration of struct arrays");
		w.Writeln ("**************************************************************************************************************************)");
		w.Writeln ("");
		
		for i := 0; i < len(componentdefinition.Structs); i++ {
			structinfo := componentdefinition.Structs[i];
			w.Writeln ("  ArrayOf%s%s = array of T%s%s;", NameSpace, structinfo.Name, NameSpace, structinfo.Name);
		}

		w.Writeln ("");
	}

	if len(componentdefinition.Functions) > 0 {
		w.Writeln ("(*************************************************************************************************************************");
		w.Writeln (" Declaration of function types");
		w.Writeln ("**************************************************************************************************************************)");
		w.Writeln ("");
		w.Writeln ("type");
		w.Writeln ("");
		for i := 0; i < len(componentdefinition.Functions); i++ {
			funcinfo := componentdefinition.Functions[i];
			arguments := ""
			for j := 0; j<len(funcinfo.Params); j++ {
				param := funcinfo.Params[j]
				if (arguments != "") {
					arguments = arguments + "; "
				}
				cParams, err := generatePlainPascalParameter(param, "", funcinfo.FunctionName, NameSpace)
				if (err != nil) {
					return err
				}
				arguments = arguments + cParams[0].ParamConvention + cParams[0].ParamName + ": " + cParams[0].ParamType
			}

			w.Writeln ("  P%s_%s = function(%s): Integer; cdecl;", NameSpace, funcinfo.FunctionName, arguments);
		}
	}
	
	w.Writeln ( "");

	return nil;
}

func getPascalParameterType(ParamTypeName string, NameSpace string, ParamClass string, isPlain bool, isImplementation bool)(string, error) {
	PascalParamTypeName := "";
	switch (ParamTypeName) {
		case "uint8":
			PascalParamTypeName = "Byte";

		case "uint16":
			PascalParamTypeName = "Word";

		case "uint32":
			PascalParamTypeName = "Cardinal";
			
		case "uint64":
			PascalParamTypeName = "QWord";

		case "int8":
			PascalParamTypeName = "ShortInt";

		case "int16":
			PascalParamTypeName = "SmallInt";

		case "int32":
			PascalParamTypeName = "Integer";
			
		case "int64":
			PascalParamTypeName = "Int64";
			
		case "bool":
			if isPlain {
				PascalParamTypeName = "Byte";
			} else {
				PascalParamTypeName = "Boolean";
			}
			
		case "single":
			PascalParamTypeName = "Single";

		case "double":
			PascalParamTypeName = "Double";
			
		case "string":
			if isPlain {
				PascalParamTypeName = "PAnsiChar";
			} else {
				PascalParamTypeName = "String";
			}

		case "enum":
			if isPlain {
				PascalParamTypeName = fmt.Sprintf ("Integer");
			} else {
				PascalParamTypeName = fmt.Sprintf ("T%s%s", NameSpace, ParamClass);
			}
		
		case "functiontype":
			if isPlain {				
				PascalParamTypeName = fmt.Sprintf ("P%s_%s", NameSpace, ParamClass);
			} else {
				PascalParamTypeName = fmt.Sprintf ("P%s_%s", NameSpace, ParamClass);
			}

		case "struct":
			if isPlain {				
				PascalParamTypeName = fmt.Sprintf ("P%s%s", NameSpace, ParamClass);
			} else {
				PascalParamTypeName = fmt.Sprintf ("T%s%s", NameSpace, ParamClass);
			}

		case "basicarray":
			basicTypeName, err := getPascalParameterType(ParamClass, NameSpace, "", isPlain, isImplementation);
			if (err != nil) {
				return "", err;
			}
			
			if isPlain {
				PascalParamTypeName = fmt.Sprintf ("P%s", basicTypeName);
			} else {
				if isImplementation {
					PascalParamTypeName = fmt.Sprintf ("P%s", basicTypeName);
				} else {
					PascalParamTypeName = fmt.Sprintf ("T%sDynArray", basicTypeName);
				}
			}

		case "structarray":
			if isPlain {
				PascalParamTypeName = fmt.Sprintf ("P%s%s", NameSpace, ParamClass)
			} else {
				if isImplementation {
					PascalParamTypeName = fmt.Sprintf ("P%s%s", NameSpace, ParamClass)
				} else {
					PascalParamTypeName = fmt.Sprintf ("ArrayOf%s%s", NameSpace, ParamClass);
				}
			}
			
		case "handle":
			if isPlain {
				PascalParamTypeName = fmt.Sprintf ("T%sHandle", NameSpace)
			} else {
				if isImplementation {
					PascalParamTypeName = "TObject";
				} else {
					PascalParamTypeName = fmt.Sprintf ("T%s%s", NameSpace, ParamClass);
				}
			}
		
		default:
			return "", fmt.Errorf ("invalid parameter type \"%s\" for Pascal parameter", ParamTypeName);
	}
	
	return PascalParamTypeName, nil;
}


type pascalParameter struct {
	ParamType string
	ParamName string
	ParamComment string
	ParamConvention string
	ParamTypeNoConvention string
}

func generatePlainPascalParameter(param ComponentDefinitionParam, className string, methodName string, NameSpace string) ([]pascalParameter, error) {
	cParams := make([]pascalParameter,1)
	cParamTypeName, err := getPascalParameterType(param.ParamType, NameSpace, param.ParamClass, true, false);
	if (err != nil) {
		return nil, err;
	}

	switch (param.ParamPass) {
	case "in":
		switch (param.ParamType) {
			case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "n" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

			case "bool":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "b" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;
				
			case "single":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "f" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

			case "double":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "d" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;
				
			case "string":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

			case "enum":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "e" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

			case "struct":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

			case "basicarray", "structarray":
				cParams = make([]pascalParameter,2)
				cParams[0].ParamType = "QWord";
				cParams[0].ParamName = "n" + param.ParamName + "Count";
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - Number of elements in buffer", cParams[0].ParamName);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

				cParams[1].ParamType = cParamTypeName;
				cParams[1].ParamName = "p" + param.ParamName + "Buffer";
				cParams[1].ParamComment = fmt.Sprintf("* @param[in] %s - %s buffer of %s", cParams[1].ParamName, param.ParamClass, param.ParamDescription);
				cParams[1].ParamConvention = "const ";
				cParams[1].ParamTypeNoConvention = cParams[1].ParamType;

			case "functiontype":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

			case "handle":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

			default:
				return nil, fmt.Errorf ("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, className, methodName, param.ParamName);
		}
	
	case "out":
	
		switch (param.ParamType) {
		
			case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "bool", "single", "double", "enum":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "out ";
				cParams[0].ParamTypeNoConvention = "P" + cParamTypeName;

			case "struct":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "";
				cParams[0].ParamTypeNoConvention = "P" + cParamTypeName[1:];
				
			case "basicarray":
				cParams = make([]pascalParameter,3)
				cParams[0].ParamType = "QWord";
				cParams[0].ParamName = "n" + param.ParamName + "Count";
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - Number of elements in buffer", cParams[0].ParamName);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

				cParams[1].ParamType = "QWord";
				cParams[1].ParamName = "p" + param.ParamName + "NeededCount";
				cParams[1].ParamComment = fmt.Sprintf("* @param[out] %s - will be filled with the count of the written elements, or needed buffer size.", cParams[1].ParamName);
				cParams[1].ParamConvention = "out ";
				cParams[1].ParamTypeNoConvention = "PQWord";

				cParams[2].ParamType = cParamTypeName;
				cParams[2].ParamName = "p" + param.ParamName + "Buffer";
				cParams[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s buffer of %s", cParams[2].ParamName, param.ParamClass, param.ParamDescription);
				cParams[2].ParamConvention = "";
				cParams[2].ParamTypeNoConvention = cParams[2].ParamType;

			case "structarray":
				cParams = make([]pascalParameter,3)
				cParams[0].ParamType = "QWord";
				cParams[0].ParamName = "n" + param.ParamName + "Count";
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - Number of elements in buffer", cParams[0].ParamName);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

				cParams[1].ParamType = "QWord";
				cParams[1].ParamName = "p" + param.ParamName + "NeededCount";
				cParams[1].ParamComment = fmt.Sprintf("* @param[out] %s - will be filled with the count of the written elements, or needed buffer size.", cParams[1].ParamName);
				cParams[1].ParamConvention = "out ";
				cParams[1].ParamTypeNoConvention = "PQWord";

				cParams[2].ParamType = cParamTypeName;
				cParams[2].ParamName = "p" + param.ParamName + "Buffer";
				cParams[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s buffer of %s", cParams[2].ParamName, param.ParamClass, param.ParamDescription);
				cParams[2].ParamConvention = "";
				cParams[2].ParamTypeNoConvention = cParams[2].ParamType;
				
			case "string":
				cParams = make([]pascalParameter,3)
				cParams[0].ParamType = "Cardinal";
				cParams[0].ParamName = "n" + param.ParamName + "BufferSize";
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - size of the buffer (including trailing 0)", cParams[0].ParamName);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

				cParams[1].ParamType = "Cardinal";
				cParams[1].ParamName = "p" + param.ParamName + "NeededChars";
				cParams[1].ParamComment = fmt.Sprintf("* @param[out] %s - will be filled with the count of the written bytes, or needed buffer size.", cParams[1].ParamName);
				cParams[1].ParamConvention = "out ";
				cParams[1].ParamTypeNoConvention = "PCardinal";

				cParams[2].ParamType = "PAnsiChar";
				cParams[2].ParamName = "p" + param.ParamName + "Buffer";
				cParams[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s buffer of %s, may be NULL", cParams[2].ParamName, param.ParamClass, param.ParamDescription);
				cParams[2].ParamConvention = "";
				cParams[2].ParamTypeNoConvention = cParams[2].ParamType;

			case "handle":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "out ";
				cParams[0].ParamTypeNoConvention = "P" + cParamTypeName[1:];
	
			default:
				return nil, fmt.Errorf ("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, className, methodName, param.ParamName);
		}

	case "return":
	
		switch (param.ParamType) {
		
			case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "bool", "single", "double", "enum":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "out ";
				cParams[0].ParamTypeNoConvention = "P" + cParamTypeName;

			case "struct":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "";
				cParams[0].ParamTypeNoConvention = "P" + cParamTypeName[1:];
				
			case "basicarray":
				cParams = make([]pascalParameter,3)
				cParams[0].ParamType = "QWord";
				cParams[0].ParamName = "n" + param.ParamName + "Count";
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - Number of elements in buffer", cParams[0].ParamName);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

				cParams[1].ParamType = "QWord";
				cParams[1].ParamName = "p" + param.ParamName + "NeededCount";
				cParams[1].ParamComment = fmt.Sprintf("* @param[out] %s - will be filled with the count of the written elements, or needed buffer size.", cParams[1].ParamName);
				cParams[1].ParamConvention = "out ";
				cParams[1].ParamTypeNoConvention = "PQWord";

				cParams[2].ParamType = cParamTypeName;
				cParams[2].ParamName = "p" + param.ParamName + "Buffer";
				cParams[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s buffer of %s", cParams[2].ParamName, param.ParamClass, param.ParamDescription);
				cParams[2].ParamConvention = "";
				cParams[2].ParamTypeNoConvention = cParams[2].ParamType;

			case "structarray":
				cParams = make([]pascalParameter,3)
				cParams[0].ParamType = "QWord";
				cParams[0].ParamName = "n" + param.ParamName + "Count";
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - Number of elements in buffer", cParams[0].ParamName);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

				cParams[1].ParamType = "QWord";
				cParams[1].ParamName = "p" + param.ParamName + "NeededCount";
				cParams[1].ParamComment = fmt.Sprintf("* @param[out] %s - will be filled with the count of the written elements, or needed buffer size.", cParams[1].ParamName);
				cParams[1].ParamConvention = "out ";
				cParams[1].ParamTypeNoConvention = "PQWord";

				cParams[2].ParamType = cParamTypeName;
				cParams[2].ParamName = "p" + param.ParamName + "Buffer";
				cParams[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s buffer of %s", cParams[2].ParamName, param.ParamClass, param.ParamDescription);
				cParams[2].ParamConvention = "";
				cParams[2].ParamTypeNoConvention = cParams[2].ParamType;
				
			case "string":
				cParams = make([]pascalParameter,3)
				cParams[0].ParamType = "Cardinal";
				cParams[0].ParamName = "n" + param.ParamName + "BufferSize";
				cParams[0].ParamComment = fmt.Sprintf("* @param[in] %s - size of the buffer (including trailing 0)", cParams[0].ParamName);
				cParams[0].ParamConvention = "const ";
				cParams[0].ParamTypeNoConvention = cParams[0].ParamType;

				cParams[1].ParamType = "Cardinal";
				cParams[1].ParamName = "p" + param.ParamName + "NeededChars";
				cParams[1].ParamComment = fmt.Sprintf("* @param[out] %s - will be filled with the count of the written bytes, or needed buffer size.", cParams[1].ParamName);
				cParams[1].ParamConvention = "out ";
				cParams[1].ParamTypeNoConvention = "PCardinal";

				cParams[2].ParamType = "PAnsiChar";
				cParams[2].ParamName = "p" + param.ParamName + "Buffer";
				cParams[2].ParamComment = fmt.Sprintf("* @param[out] %s - %s buffer of %s, may be NULL", cParams[2].ParamName, param.ParamClass, param.ParamDescription);
				cParams[2].ParamConvention = "";
				cParams[2].ParamTypeNoConvention = cParams[2].ParamType;

			case "handle":
				cParams[0].ParamType = cParamTypeName;
				cParams[0].ParamName = "p" + param.ParamName;
				cParams[0].ParamComment = fmt.Sprintf("* @param[out] %s - %s", cParams[0].ParamName, param.ParamDescription);
				cParams[0].ParamConvention = "out ";
				cParams[0].ParamTypeNoConvention = "P" + cParamTypeName[1:];
	
			default:
				return nil, fmt.Errorf ("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, className, methodName, param.ParamName);
		}
		
	default:
		return nil, fmt.Errorf ("invalid method parameter passing \"%s\" for %s.%s (%s)", param.ParamPass, className, methodName, param.ParamName);
	}

	return cParams, nil;
}
