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
	"log"
	"path"
	"fmt"
	"strings"
)

// BuildBindingCSharp builds CSharp bindings of a library's API in form of dynamically loaded functions
// handles.
func BuildBindingCSharp(component ComponentDefinition, outputFolder string, outputFolderExample string, indentString string) error {

	namespace := component.NameSpace;
	baseName := component.BaseName;
	
	CSharpImpl := path.Join(outputFolder, namespace+".cs");
	log.Printf("Creating \"%s\"", CSharpImpl)
	CSharpImplFile, err := CreateLanguageFile (CSharpImpl, indentString)
	if err != nil {
		return err;
	}
	
	
	err = BuildBindingCSharpImplementation (component, CSharpImplFile, namespace, baseName);

	return err;
}



func getCSharpParameterType(ParamTypeName string, NameSpace string, ParamClass string, isPlain bool) (string, error) {
	CSharpParamTypeName := "";
	switch (ParamTypeName) {
		case "uint8":
			CSharpParamTypeName = "Byte";

		case "uint16":
			CSharpParamTypeName = "UInt16";

		case "uint32":
			CSharpParamTypeName = "UInt32";
			
		case "uint64":
			CSharpParamTypeName = "UInt64";

		case "int8":
			CSharpParamTypeName = "Int8";

		case "int16":
			CSharpParamTypeName = "Int16";

		case "int32":
			CSharpParamTypeName = "Int32";
			
		case "int64":
			CSharpParamTypeName = "Int64";
			
		case "bool":
			if isPlain {
				CSharpParamTypeName = "Int32";
			} else {
				CSharpParamTypeName = "bool";
			}
			
		case "single":
			CSharpParamTypeName = "Single";

		case "double":
			CSharpParamTypeName = "Double";

		case "pointer":
			CSharpParamTypeName = "IntPtr";
			
		case "string":
			if isPlain {
				CSharpParamTypeName = "String";
			} else {
				CSharpParamTypeName = "String";
			}

		case "enum":
			CSharpParamTypeName = fmt.Sprintf ("Int32");
		
		case "functiontype":
			CSharpParamTypeName = fmt.Sprintf ("IntPtr");

		case "struct":
			CSharpParamTypeName = fmt.Sprintf ("IntPtr");

		case "basicarray":
			CSharpParamTypeName = fmt.Sprintf ("IntPtr");

		case "structarray":
			CSharpParamTypeName = fmt.Sprintf ("IntPtr");
			
		case "class":
			if isPlain {
				CSharpParamTypeName = "IntPtr";
			} else {
				CSharpParamTypeName = "C" + ParamClass;
			}
		
		default:
	}
	
	return CSharpParamTypeName, nil;
}



func getCSharpPlainParameters(method ComponentDefinitionMethod, NameSpace string, ClassName string, isGlobal bool) (string, error) {
	parameters := "";
	
	for k := 0; k < len(method.Params); k++ {
		param := method.Params [k];
		ParamTypeName, err := getCSharpParameterType(param.ParamType, NameSpace, param.ParamClass, true);
		if err != nil {
			return "", err;
		}
		
		switch (param.ParamPass) {
			case "in":
				if (parameters != "") {
					parameters = parameters + ", ";
				}			
				parameters = parameters + ParamTypeName + " A" + param.ParamName;
				
			case "out", "return":
				if (parameters != "") {
					parameters = parameters + ", ";
				}			
				parameters = parameters + "out " + ParamTypeName + " A" + param.ParamName;

		}
	}
	
	return parameters, nil;
}


func getCSharpClassParameters(method ComponentDefinitionMethod, NameSpace string, ClassName string, isGlobal bool) (string, string, error) {
	parameters := "";
	returnType := "";
	
	for k := 0; k < len(method.Params); k++ {
		param := method.Params [k];
		ParamTypeName, err := getCSharpParameterType(param.ParamType, NameSpace, param.ParamClass, false);
		if err != nil {
			return "", "", err;
		}
		
		switch (param.ParamPass) {
			case "in":
				if (parameters != "") {
					parameters = parameters + ", ";
				}			
				parameters = parameters + ParamTypeName + " A" + param.ParamName;
				
			case "out":
				if (parameters != "") {
					parameters = parameters + ", ";
				}			
				parameters = parameters + "out " + ParamTypeName + " A" + param.ParamName;

			case "return":
				if (returnType != "") {
					return "", "", fmt.Errorf ("duplicate return value \"%s\" for Pascal method \"%s\"", param.ParamName, method.MethodName);
				}
				returnType = ParamTypeName;
		}
	}
	
	if (returnType == "") {
		returnType = "void";
	}

	return parameters, returnType, nil;
}



func writeCSharpClassMethodImplementation (method ComponentDefinitionMethod, w LanguageWriter, NameSpace string, ClassName string, isGlobal bool, spacing string) (error) {

	
	
	defineCommands := make([]string, 0);
	//initCommands := make([]string, 0);
	resultCommands := make([]string, 0);
	//postInitCommands := make([]string, 0);
	//wrapperCallPrefix := "";
	
	//doInitCall := false;
	
	
	callFunctionName := "";	
	callFunctionParameters := "";
	initCallParameters := "";
	//errorInstanceHandle := "";
	
	if isGlobal {
		callFunctionName = fmt.Sprintf ("%s", method.MethodName);
		//errorInstanceHandle = "nil";
	} else {
		callFunctionName = fmt.Sprintf ("%s_%s", ClassName, method.MethodName);		
		callFunctionParameters = "Handle";
		//errorInstanceHandle = "Self";
	}
	
	initCallParameters = callFunctionParameters;
	
	
	for k := 0; k < len(method.Params); k++ {
		param := method.Params [k];
		ParamTypeName, err := getCSharpParameterType(param.ParamType, NameSpace, param.ParamClass, false);
		if err != nil {
			return err;
		}
				
		if (callFunctionParameters != "") {
			callFunctionParameters = callFunctionParameters + ", ";
		}			

		if (initCallParameters != "") {
			initCallParameters = initCallParameters + ", ";
		}			
		
		switch (param.ParamPass) {
			case "in":

				switch (param.ParamType) {
					case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
						callFunctionParameters = callFunctionParameters + "A" + param.ParamName;

					case "single":
						callFunctionParameters = callFunctionParameters + "A" + param.ParamName;

					case "double":
						callFunctionParameters = callFunctionParameters + "A" + param.ParamName;

					case "pointer":						
						callFunctionParameters = callFunctionParameters + "A" + param.ParamName;
						
					case "string":
						callFunctionParameters = callFunctionParameters + "A" + param.ParamName;
						
					case "enum":

					case "bool":
						callFunctionParameters = callFunctionParameters + "( A" + param.ParamName + " ? 1 : 0 )";
						
					case "struct":

					case "basicarray":

					case "structarray":

					case "functiontype":

					case "class":
						callFunctionParameters = callFunctionParameters + "A" + param.ParamName + ".GetHandle()";

					default:
						return fmt.Errorf ("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, ClassName, method.MethodName, param.ParamName);
				}

			
				
			case "out":
			
				switch (param.ParamType) {
					case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
					
						defineCommands = append (defineCommands, fmt.Sprintf ("  %s result%s = 0;", ParamTypeName, param.ParamName));
						callFunctionParameters = callFunctionParameters + "out result" + param.ParamName;

					case "single":					
						defineCommands = append (defineCommands, fmt.Sprintf ("  %s result%s = 0.0;", ParamTypeName, param.ParamName));
						callFunctionParameters = callFunctionParameters + "out result" + param.ParamName;

					case "double":					
						defineCommands = append (defineCommands, fmt.Sprintf ("  %s result%s = 0.0;", ParamTypeName, param.ParamName));
						callFunctionParameters = callFunctionParameters + "out result" + param.ParamName;

					case "pointer":					
						defineCommands = append (defineCommands, fmt.Sprintf ("  %s result%s = (IntPtr) 0;", ParamTypeName, param.ParamName));
						callFunctionParameters = callFunctionParameters + "out result" + param.ParamName;
						
					case "string":
						/*defineCommands = append (defineCommands, "  bytesNeeded" + param.ParamName + ": Cardinal;");
						defineCommands = append (defineCommands, "  bytesWritten" + param.ParamName + ": Cardinal;");
						defineCommands = append (defineCommands, "  buffer" + param.ParamName + ": array of Char;");
						initCommands = append (initCommands, "  bytesNeeded" + param.ParamName + ":= 0;");
						initCommands = append (initCommands, "  bytesWritten" + param.ParamName + ":= 0;");
						
						initCallParameters = initCallParameters + fmt.Sprintf("0, bytesNeeded%s, nil", param.ParamName)
						
						postInitCommands = append (postInitCommands, fmt.Sprintf("  SetLength (buffer%s, bytesNeeded%s + 2);", param.ParamName, param.ParamName));
						
						callFunctionParameters = callFunctionParameters + fmt.Sprintf("bytesNeeded%s + 1, bytesWritten%s, @buffer%s[0]", param.ParamName, param.ParamName, param.ParamName)

						resultCommands = append (resultCommands, fmt.Sprintf ("  buffer%s[bytesNeeded%s + 1] := #0;", param.ParamName, param.ParamName));
						resultCommands = append (resultCommands, fmt.Sprintf ("  A%s := StrPas (@buffer%s[0]);", param.ParamName, param.ParamName));

						doInitCall = true; */
						
					case "enum":
						/*defineCommands = append (defineCommands, "  Result" + param.ParamName + ": Integer;");
						initCommands = append (initCommands, "  Result" + param.ParamName + " := 0;");
			
						callFunctionParameters = callFunctionParameters + "Result" + param.ParamName;
						initCallParameters = initCallParameters + "Result" + param.ParamName;
						resultCommands = append (resultCommands, fmt.Sprintf ("  A%s := convertConstTo%s (Result%s);", param.ParamName, param.ParamClass, param.ParamName)); */

					case "bool":
						defineCommands = append (defineCommands, fmt.Sprintf ("  Int32 result%s = 0;", param.ParamName));
						callFunctionParameters = callFunctionParameters + "out result" + param.ParamName;
						
					case "struct":
						/*callFunctionParameters = callFunctionParameters + "@A" + param.ParamName;
						initCallParameters = initCallParameters + "@A" + param.ParamName; */
						
					case "basicarray", "structarray":
						
						/*defineCommands = append (defineCommands, "  countNeeded" + param.ParamName + ": QWord;");
						defineCommands = append (defineCommands, "  countWritten" + param.ParamName + ": QWord;");
						initCommands = append (initCommands, "  countNeeded" + param.ParamName + ":= 0;");
						initCommands = append (initCommands, "  countWritten" + param.ParamName + ":= 0;");
						
						initCallParameters = initCallParameters + fmt.Sprintf("0, countNeeded%s, nil", param.ParamName)
						
						postInitCommands = append (postInitCommands, fmt.Sprintf("  SetLength (A%s, countNeeded%s);", param.ParamName, param.ParamName));
						
						callFunctionParameters = callFunctionParameters + fmt.Sprintf("countNeeded%s, countWritten%s, @A%s[0]", param.ParamName, param.ParamName, param.ParamName)

						doInitCall = true; */
					
					case "class":
						/*defineCommands = append (defineCommands, "  H" + param.ParamName + ": " + PlainParamTypeName + ";");
						initCommands = append (initCommands, "  Result := nil;");
						initCommands = append (initCommands, "  A%s := nil;", param.ParamName);
						initCommands = append (initCommands, "  H" + param.ParamName + " := nil;");
						callFunctionParameters = callFunctionParameters + "H" + param.ParamName;
						initCallParameters = initCallParameters + "nil";
						
						resultCommands = append (resultCommands, fmt.Sprintf ("  if Assigned (H%s) then", param.ParamName));
						resultCommands = append (resultCommands, fmt.Sprintf ("    A%s := T%s%s.Create (%s, H%s);", param.ParamName, NameSpace, param.ParamClass, wrapperInstanceName, param.ParamName)); */

					default:
						return fmt.Errorf ("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, ClassName, method.MethodName, param.ParamName);
				}


			case "return":

				switch (param.ParamType) {
					case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "single", "double", "pointer":
					
						defineCommands = append (defineCommands, fmt.Sprintf ("  %s result%s = 0;", ParamTypeName, param.ParamName));
						callFunctionParameters = callFunctionParameters + "out result" + param.ParamName;
						resultCommands = append (resultCommands, fmt.Sprintf ("  return result%s;", param.ParamName));

					case "string":
						/*defineCommands = append (defineCommands, "  bytesNeeded" + param.ParamName + ": Cardinal;");
						defineCommands = append (defineCommands, "  bytesWritten" + param.ParamName + ": Cardinal;");
						defineCommands = append (defineCommands, "  buffer" + param.ParamName + ": array of Char;");
						initCommands = append (initCommands, "  bytesNeeded" + param.ParamName + ":= 0;");
						initCommands = append (initCommands, "  bytesWritten" + param.ParamName + ":= 0;");
						
						initCallParameters = initCallParameters + fmt.Sprintf("0, bytesNeeded%s, nil", param.ParamName)

						postInitCommands = append (postInitCommands, fmt.Sprintf("  SetLength (buffer%s, bytesNeeded%s + 2);", param.ParamName, param.ParamName));
						
						callFunctionParameters = callFunctionParameters + fmt.Sprintf("bytesNeeded%s + 2, bytesWritten%s, @buffer%s[0]", param.ParamName, param.ParamName, param.ParamName)

						resultCommands = append (resultCommands, fmt.Sprintf ("  buffer%s[bytesNeeded%s + 1] := #0;", param.ParamName, param.ParamName));
						resultCommands = append (resultCommands, fmt.Sprintf ("  Result := StrPas (@buffer%s[0]);", param.ParamName));

						doInitCall = true; */

						
					case "enum":
						/*defineCommands = append (defineCommands, "  Result" + param.ParamName + ": Integer;");
						initCommands = append (initCommands, "  Result" + param.ParamName + " := 0;");
			
						callFunctionParameters = callFunctionParameters + "Result" + param.ParamName;
						initCallParameters = initCallParameters + "Result" + param.ParamName;
						resultCommands = append (resultCommands, fmt.Sprintf ("  Result := convertConstTo%s (Result%s);", param.ParamClass, param.ParamName)); */

					case "bool":
						defineCommands = append (defineCommands, fmt.Sprintf ("  Int32 result%s = 0;", param.ParamName));
						callFunctionParameters = callFunctionParameters + "out result" + param.ParamName;
						resultCommands = append (resultCommands, fmt.Sprintf ("  return (result%s != 0);", param.ParamName));
						
						
					case "struct":
						//callFunctionParameters = callFunctionParameters + "@Result";

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

						defineCommands = append (defineCommands, fmt.Sprintf ("  IntPtr new%s = (IntPtr) 0;", param.ParamName));
						callFunctionParameters = callFunctionParameters + "out new" + param.ParamName;
						resultCommands = append (resultCommands, fmt.Sprintf ("  return new C%s (new%s );", param.ParamClass, param.ParamName));

						/*defineCommands = append (defineCommands, "  H" + param.ParamName + ": " + PlainParamTypeName + ";");
						initCommands = append (initCommands, "  Result := nil;");
						initCommands = append (initCommands, "  H" + param.ParamName + " := nil;");
						callFunctionParameters = callFunctionParameters + "H" + param.ParamName;
						resultCommands = append (resultCommands, fmt.Sprintf ("  if Assigned (H%s) then", param.ParamName));
						resultCommands = append (resultCommands, fmt.Sprintf ("    Result := T%s%s.Create (%s, H%s);", NameSpace, param.ParamClass, wrapperInstanceName, param.ParamName)); */

					default:
						return fmt.Errorf ("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, ClassName, method.MethodName, param.ParamName);
				}
			
				
				
		}
	}
	
	
	
	if len (defineCommands) > 0 {
		w.Writelns (spacing, defineCommands);	
	}
		
/*	if (doInitCall) {
		w.Writeln (spacing + "  %sCheckError (%s, %s%s (%s));", wrapperCallPrefix, errorInstanceHandle, wrapperCallPrefix, callFunctionName, initCallParameters);
	}
	
	w.Writelns (spacing, postInitCommands);	
	
	
	
	w.Writeln (spacing + "end;");
	w.Writeln (""); */

	w.Writeln (spacing + "  Internal.%sWrapper.%s (%s);", NameSpace, callFunctionName, callFunctionParameters);
	
	w.Writelns (spacing, resultCommands);	
	
	return nil;
}



func BuildBindingCSharpImplementation (component ComponentDefinition, w LanguageWriter, NameSpace string, BaseName string) error {

	baseName := component.BaseName;
	global := component.Global;
	
	
	CSharpBaseClassName := "C" + component.Global.BaseClassName;
	w.Writeln ("using System;")
	w.Writeln ("using System.Runtime.InteropServices;")
	w.Writeln ("")

	w.Writeln ("namespace %s {", NameSpace)
	w.Writeln ("")

	w.Writeln ("  namespace Internal {")
	w.Writeln ("")
	w.Writeln ("    public class %sWrapper", NameSpace);
	w.Writeln ("    {")
	
	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		
		for j := 0; j < len(class.Methods); j++ {
			method := class.Methods[j]
			
			parameters, err := getCSharpPlainParameters (method, NameSpace, class.ClassName, false);
			if (err != nil) {
				return err;
			}
			
			w.Writeln ("      [DllImport(\"%s.dll\", EntryPoint = \"%s_%s_%s\", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]", baseName, strings.ToLower (NameSpace), strings.ToLower (class.ClassName), strings.ToLower (method.MethodName));
			
			if (parameters == "") {
				parameters = "IntPtr Handle";
			} else {
				parameters = "IntPtr Handle, " + parameters;
			}
			
			w.Writeln ("      public extern static Int32 %s_%s (%s);", class.ClassName, method.MethodName, parameters);
			w.Writeln ("")
			
		}

	}

	for j := 0; j < len(global.Methods); j++ {
		method := global.Methods[j]
		
		parameters, err := getCSharpPlainParameters (method, NameSpace, "", true);
		if (err != nil) {
			return err;
		}
		
			
		w.Writeln ("      [DllImport(\"%s.dll\", EntryPoint = \"%s_%s\", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]", baseName, strings.ToLower (NameSpace), strings.ToLower (method.MethodName));
		w.Writeln ("      public extern static Int32 %s (%s);", method.MethodName, parameters);
		w.Writeln ("")
	}
	
	
	w.Writeln ("    }")
	w.Writeln ("  }")

	w.Writeln ("")
	w.Writeln ("")

	
	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]

		CSharpParentClassName := ""
		if (!component.isBaseClass(class)) {
			if class.ParentClass == "" {
				CSharpParentClassName = ": " + CSharpBaseClassName
			} else {
				CSharpParentClassName = ": C" + class.ParentClass
			}
		}

		w.Writeln ("  class C%s %s", class.ClassName, CSharpParentClassName);
		w.Writeln ("  {")

		if (component.isBaseClass(class)) {
			w.Writeln ("    protected IntPtr Handle;")
			w.Writeln ("")
			w.Writeln ("    public C%s (IntPtr NewHandle)", class.ClassName)
			w.Writeln ("    {")
			w.Writeln ("      Handle = NewHandle;")
			w.Writeln ("    }")
			w.Writeln ("")
			w.Writeln ("    ~C%s ()", class.ClassName)
			w.Writeln ("    {")
			w.Writeln ("      if (Handle != (IntPtr) 0) {")
			w.Writeln ("        Internal.%sWrapper.%s (Handle);", NameSpace, component.Global.ReleaseMethod)
			w.Writeln ("        Handle = (IntPtr) 0;")
			w.Writeln ("      }")
			w.Writeln ("    }")
			w.Writeln ("")
			
			w.Writeln ("    public IntPtr GetHandle ()")
			w.Writeln ("    {")
			w.Writeln ("      return Handle;")
			w.Writeln ("    }")
			w.Writeln ("")
		} else {
			w.Writeln ("    public C%s (IntPtr NewHandle) : base (NewHandle)", class.ClassName)
			w.Writeln ("    {")
			w.Writeln ("    }")
			w.Writeln ("")
		}
		
		for j := 0; j < len(class.Methods); j++ {
			method := class.Methods[j]
			
			parameters, returnType, err := getCSharpClassParameters (method, NameSpace, class.ClassName, false);
			if (err != nil) {
				return err;
			}
			
			w.Writeln ("    %s %s (%s)", returnType, method.MethodName, parameters);
			w.Writeln ("    {");
			
			writeCSharpClassMethodImplementation (method, w, NameSpace, class.ClassName, false, "    ");
			
			w.Writeln ("    }");
			w.Writeln ("");
		}
		
		w.Writeln ("  }")
		w.Writeln ("")
	}

	
	w.Writeln ("  class CWrapper");
	w.Writeln ("  {")
		
	for j := 0; j < len(global.Methods); j++ {
		method := global.Methods[j]

		parameters, returnType, err := getCSharpClassParameters (method, NameSpace, "", true);
		if (err != nil) {
			return err;
		}

		w.Writeln ("    %s %s (%s)", returnType, method.MethodName, parameters);
		w.Writeln ("    {");
		
		writeCSharpClassMethodImplementation (method, w, NameSpace, "Wrapper", true, "    ");
		
		w.Writeln ("    }");
		w.Writeln ("");
	}
		
	w.Writeln ("  }")
	w.Writeln ("")

	w.Writeln ("}")

	
	return nil
}


