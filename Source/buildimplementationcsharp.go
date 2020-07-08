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
// buildimplementationcpp.go
// functions to generate C++ interface classes, implementation stubs and wrapper code that maps to
// the C-header.
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"errors"
	"path"
	"strings"
)

// BuildImplementationCPP builds CSharp interface classes, implementation stubs and wrapper code that maps to the C-header
func BuildImplementationCSharp(component ComponentDefinition, outputFolder string, stubOutputFolder string, projectOutputFolder string, implementation ComponentDefinitionImplementation, suppressStub bool, suppressInterfaces bool) error {
	forceRecreation := false

	NameSpace := component.NameSpace;
	ImplementationSubNameSpace := "Impl"
	
	BaseName := component.BaseName;

	indentString := getIndentationString(implementation.Indentation)
	stubIdentifier := ""
	if len(implementation.StubIdentifier) > 0 {
		stubIdentifier = "_" + strings.ToLower(implementation.StubIdentifier)
	}


	if (!suppressInterfaces) {

		err := buildCSharpInterfaces(component, NameSpace, ImplementationSubNameSpace, implementation.ClassIdentifier, BaseName, outputFolder, indentString)
		if err != nil {
			return err
		}
	
	}


	if (!suppressStub) {

		err := buildCSharpStub(component, NameSpace, ImplementationSubNameSpace, implementation.ClassIdentifier, BaseName, stubOutputFolder, indentString, stubIdentifier, forceRecreation)
		if err != nil {
			return err
		}
	
	}

	return nil
}






func buildCSharpInterface(component ComponentDefinition, class ComponentDefinitionClass, NameSpace string, NameSpaceImplementation string, ClassIdentifier string, stubintfw LanguageWriter) error {

		outIntfName := "I" + ClassIdentifier + class.ClassName;
		CSharpBaseClassName := "I" + component.Global.BaseClassName;

		if (!component.isBaseClass(class)) {
			if (class.ParentClass == "") {
				class.ParentClass = component.Global.BaseClassName
			}
		}

		stubintfw.Writeln("")
		stubintfw.Writeln("  /*************************************************************************************************************************")
		stubintfw.Writeln("   COM interface of %s ", outIntfName)
		stubintfw.Writeln("  **************************************************************************************************************************/")
		stubintfw.Writeln("")

		stubintfw.Writeln("  [ComVisible(true)]")
		stubintfw.Writeln("  [Guid(ContractGuids.IID_%s)]", outIntfName)
		stubintfw.Writeln("  [InterfaceType(ComInterfaceType.InterfaceIsIUnknown)]")
		
		CSharpParentIntfName := ""
		if !component.isBaseClass(class) {
			if class.ParentClass == "" {
				CSharpParentIntfName = ": " + CSharpBaseClassName
			} else {
				CSharpParentIntfName = ": I" + class.ParentClass
			}
		}		
		
		stubintfw.Writeln("  public interface %s %s", outIntfName, CSharpParentIntfName)
		stubintfw.Writeln("  {")
 
		for j := 0; j < len(class.Methods); j++ {
			method := class.Methods[j]

			parameters, returnType, err := getCSharpClassParameters(method, NameSpace, class.ClassName, false)
			if err != nil {
				return err
			}

			stubintfw.Writeln("    public %s %s (%s);", returnType, method.MethodName, parameters)
			stubintfw.Writeln("")
		} 
 
		stubintfw.Writeln("  }")

		stubintfw.Writeln("")
		


	return nil
}




func buildCSharpInterfaces(component ComponentDefinition, NameSpace string, NameSpaceImplementation string, ClassIdentifier string, BaseName string, outputFolder string, indentString string) error {

	StubIntfFileName := path.Join(outputFolder, BaseName + "_Interfaces.cs");

	log.Printf("Creating \"%s\"", StubIntfFileName)
	stubintfw, err := CreateLanguageFile(StubIntfFileName, indentString)
	if err != nil {
		return err
	}
	stubintfw.WriteCLicenseHeader(component,
		fmt.Sprintf("This is the interface declaration file of %s", NameSpace),
		false)
	

	stubintfw.Writeln("")
	stubintfw.Writeln("using System;")
	stubintfw.Writeln("using System.Runtime.InteropServices;")
	stubintfw.Writeln("")
	
	stubintfw.Writeln("namespace %s {", NameSpace)

	stubintfw.Writeln("")
	stubintfw.Writeln("  /*************************************************************************************************************************")
	stubintfw.Writeln("   Error Codes and Exception definition of %s ", NameSpace)
	stubintfw.Writeln("  **************************************************************************************************************************/")
	stubintfw.Writeln("");

	stubintfw.Writeln("  public class ErrorCodes {")

	for i := 0; i < len(component.Errors.Errors); i++ {
		errorcode := component.Errors.Errors[i]
		stubintfw.Writeln("    public const UInt32 %s = %d;", errorcode.Name, errorcode.Code)
	}

	stubintfw.Writeln("  }")


	stubintfw.Writeln("");
	stubintfw.Writeln("  [Serializable()]")
	stubintfw.Writeln("  public class %sException : System.Exception", NameSpace)
	stubintfw.Writeln("  {")
	stubintfw.Writeln("    public static string ErrorCodeToString (UInt32 errorCode)")
	stubintfw.Writeln("    {")
	stubintfw.Writeln("      switch (errorCode)")
	stubintfw.Writeln("      {")
	for i := 0; i < len(component.Errors.Errors); i++ {
		errorcode := component.Errors.Errors[i]
		stubintfw.Writeln("        case ErrorCodes.%s: return \"%s\";", errorcode.Name, errorcode.Description)
	}

	stubintfw.Writeln("        default: return String.Format(\"Unknown Error #{0}\", errorCode);")
	stubintfw.Writeln("      }")
	stubintfw.Writeln("    }")
	stubintfw.Writeln("    ")
	
	stubintfw.Writeln("    public %sException() : base() { }", NameSpace)	
	stubintfw.Writeln("    public %sException(UInt32 errorCode) : base(ErrorCodeToString (errorCode)) { }", NameSpace)
	stubintfw.Writeln("    public %sException(string message) : base(message) { }", NameSpace)
	stubintfw.Writeln("    public %sException(string message, System.Exception inner) : base(message, inner) { }", NameSpace)
	stubintfw.Writeln("    protected %sException(System.Runtime.Serialization.SerializationInfo info,", NameSpace)
	stubintfw.Writeln("       System.Runtime.Serialization.StreamingContext context) : base(info, context) { }")
	stubintfw.Writeln("  }")

	
	

	stubintfw.Writeln("")
	stubintfw.Writeln("  /*************************************************************************************************************************")
	stubintfw.Writeln("   COM GUID definitions of %s ", NameSpace)
	stubintfw.Writeln("  **************************************************************************************************************************/")
	stubintfw.Writeln("")


	stubintfw.Writeln("  internal sealed class ContractGuids")
	stubintfw.Writeln("  {")
	
	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		
		if (strings.TrimSpace (class.CLSID) == "") {
			return errors.New ("Invalid class CLSID for " + class.ClassName);
		}

		if (strings.TrimSpace (class.IID) == "") {
			return errors.New ("Invalid class IID for " + class.ClassName);
		}
		
		stubintfw.Writeln("    public const string CLSID_I%s = \"%s\";", class.ClassName, strings.ToUpper (strings.TrimSpace (class.CLSID)));
		stubintfw.Writeln("    public const string IID_I%s = \"%s\";", class.ClassName, strings.ToUpper (strings.TrimSpace (class.IID)));
	}

	stubintfw.Writeln("    public const string CLSID_I%s = \"%s\";", "Wrapper", strings.ToUpper (strings.TrimSpace (component.Global.CLSID)));
	stubintfw.Writeln("    public const string IID_I%s = \"%s\";", "Wrapper", strings.ToUpper (strings.TrimSpace (component.Global.IID)));

	if (strings.TrimSpace (component.Global.CLSID) == "") {
		return errors.New ("Invalid wrapper CLSID");
	}

	if (strings.TrimSpace (component.Global.IID) == "") {
		return errors.New ("Invalid wrapper IID");
	}


	stubintfw.Writeln("  }")
	stubintfw.Writeln("")


	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		err :=  buildCSharpInterface(component, class, NameSpace, NameSpaceImplementation, ClassIdentifier, stubintfw)
		if err != nil {
			return err
		}
	}

	stubintfw.Writeln("}")
	stubintfw.Writeln("")
	
	return nil
}



func buildCSharpStubClass(component ComponentDefinition, class ComponentDefinitionClass, NameSpace string, NameSpaceImplementation string, ClassIdentifier string, BaseName string, outputFolder string, indentString string, stubIdentifier string, forceRecreation bool) error {

		CSharpBaseClassName := "C" + component.Global.BaseClassName;
		
		outClassName := "C" + ClassIdentifier + class.ClassName

		StubImplFileName := path.Join(outputFolder, BaseName + stubIdentifier + "_C" + class.ClassName +".cs");
		if !forceRecreation && ( FileExists(StubImplFileName) ) {
			log.Printf("Omitting recreation of Stub implementation for \"%s\"", outClassName)
			return nil
		}

		
		log.Printf("Creating \"%s\"", StubImplFileName)
		stubimplw, err := CreateLanguageFile(StubImplFileName, indentString)
		if err != nil {
			return err
		}
		stubimplw.WriteCLicenseHeader(component,
			fmt.Sprintf("This is a stub class definition of %s", outClassName),
			false)

		stubimplw.Writeln("")
		stubimplw.Writeln("using System;")
		stubimplw.Writeln("using System.Runtime.InteropServices;")
		stubimplw.Writeln("")
		stubimplw.Writeln("namespace %s {", NameSpace)

		if (!component.isBaseClass(class)) {
			if (class.ParentClass == "") {
				class.ParentClass = component.Global.BaseClassName
			}
		}

		stubimplw.Writeln("")
		stubimplw.Writeln("  /*************************************************************************************************************************")
		stubimplw.Writeln("   Class implementation of %s ", outClassName)
		stubimplw.Writeln("  **************************************************************************************************************************/")
		stubimplw.Writeln("")

		stubimplw.Writeln("  [ComVisible(true)]")
		stubimplw.Writeln("  [Guid(ContractGuids.CLSID_I%s)]", class.ClassName)
		
		CSharpParentClassName := ""
		if !component.isBaseClass(class) {
			if class.ParentClass == "" {
				CSharpParentClassName = ": " + CSharpBaseClassName + ", I" + class.ClassName
			} else {
				CSharpParentClassName = ": C" + class.ParentClass + ", I" + class.ClassName
			}
		} else {
			
			CSharpParentClassName = ": I" + class.ClassName
		
		}
		
		stubimplw.Writeln("  public interface %s %s", outClassName, CSharpParentClassName)
		stubimplw.Writeln("  {")
 
		for j := 0; j < len(class.Methods); j++ {
			method := class.Methods[j]

			parameters, returnType, err := getCSharpClassParameters(method, NameSpace, class.ClassName, false)
			if err != nil {
				return err
			}

			stubimplw.Writeln("    %s I%s.%s (%s)", returnType, class.ClassName, method.MethodName, parameters)
			stubimplw.Writeln("    {")
			stubimplw.Writeln("      throw new %sException (ErrorCodes.NOTIMPLEMENTED);", NameSpace)
			stubimplw.Writeln("    }")
			stubimplw.Writeln("")
		} 
 
		stubimplw.Writeln("  }")

		stubimplw.Writeln("")
		
		stubimplw.Writeln("}")
		stubimplw.Writeln("")

	return nil
}

func buildCSharpStub(component ComponentDefinition, NameSpace string, NameSpaceImplementation string, ClassIdentifier string, BaseName string, outputFolder string, indentString string, stubIdentifier string, forceRecreation bool) error {

	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		err :=  buildCSharpStubClass(component, class, NameSpace, NameSpaceImplementation, ClassIdentifier, BaseName, outputFolder, indentString, stubIdentifier, forceRecreation)
		if err != nil {
			return err
		}
	}

	return nil
}
