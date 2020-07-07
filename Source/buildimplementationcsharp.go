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






func buildCSharpInterface(component ComponentDefinition, class ComponentDefinitionClass, NameSpace string, NameSpaceImplementation string, ClassIdentifier string, BaseName string, outputFolder string, indentString string) error {
		outIntfName := "I" + ClassIdentifier + class.ClassName;
		CSharpBaseClassName := "I" + component.Global.BaseClassName;

		StubIntfFileName := path.Join(outputFolder, BaseName + "_I" + class.ClassName +".cs");

		log.Printf("Creating \"%s\"", StubIntfFileName)
		stubintfw, err := CreateLanguageFile(StubIntfFileName, indentString)
		if err != nil {
			return err
		}
		stubintfw.WriteCLicenseHeader(component,
			fmt.Sprintf("This is the class declaration of %s", outIntfName),
			false)
		

		stubintfw.Writeln("")
		stubintfw.Writeln("using System;")
		stubintfw.Writeln("using System.Runtime.InteropServices;")
		stubintfw.Writeln("")
		stubintfw.Writeln("namespace %s {", NameSpace)

		if (!component.isBaseClass(class)) {
			if (class.ParentClass == "") {
				class.ParentClass = component.Global.BaseClassName
			}
		}

		stubintfw.Writeln("")
		stubintfw.Writeln("/*************************************************************************************************************************")
		stubintfw.Writeln(" Class interface of %s ", outIntfName)
		stubintfw.Writeln("**************************************************************************************************************************/")
		stubintfw.Writeln("")

		stubintfw.Writeln("  [ComVisible(true)]")
		stubintfw.Writeln("  [Guid(ContractGuids.IID_%s)]", outIntfName)
		stubintfw.Writeln("  [InterfaceType(ComInterfaceType.InterfaceIsIUnknown)]")
		stubintfw.Writeln("  ")
		
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
		
		stubintfw.Writeln("}")
		stubintfw.Writeln("")

	return nil
}




func buildCSharpInterfaces(component ComponentDefinition, NameSpace string, NameSpaceImplementation string, ClassIdentifier string, BaseName string, outputFolder string, indentString string) error {

	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		err :=  buildCSharpInterface(component, class, NameSpace, NameSpaceImplementation, ClassIdentifier, BaseName, outputFolder, indentString)
		if err != nil {
			return err
		}
	}

	return nil
}



func buildCSharpStubClass(component ComponentDefinition, class ComponentDefinitionClass, NameSpace string, NameSpaceImplementation string, ClassIdentifier string, BaseName string, outputFolder string, indentString string, stubIdentifier string, forceRecreation bool) error {
		outClassName := "C" + ClassIdentifier + class.ClassName

		StubIntfFileName := path.Join(outputFolder, BaseName + stubIdentifier + "_I" + class.ClassName +".cs");
		StubImplFileName := path.Join(outputFolder, BaseName + stubIdentifier + "_" + class.ClassName +".cs");
		if !forceRecreation && ( FileExists(StubIntfFileName) || FileExists(StubImplFileName) ) {
			log.Printf("Omitting recreation of Stub implementation for \"%s\"", outClassName)
			return nil
		}

		log.Printf("Creating \"%s\"", StubIntfFileName)
		stubintfw, err := CreateLanguageFile(StubIntfFileName, indentString)
		if err != nil {
			return err
		}
		stubintfw.WriteCLicenseHeader(component,
			fmt.Sprintf("This is the class declaration of %s", outClassName),
			false)
		
		log.Printf("Creating \"%s\"", StubImplFileName)
		stubimplw, err := CreateLanguageFile(StubImplFileName, indentString)
		if err != nil {
			return err
		}
		stubimplw.WriteCLicenseHeader(component,
			fmt.Sprintf("This is a stub class definition of %s", outClassName),
			false)

		stubintfw.Writeln("")
		stubintfw.Writeln("using System;")
		stubintfw.Writeln("using System.Runtime.InteropServices;")
		stubintfw.Writeln("")
		stubintfw.Writeln("namespace %s {", NameSpace)

		if (!component.isBaseClass(class)) {
			if (class.ParentClass == "") {
				class.ParentClass = component.Global.BaseClassName
			}
		}

		stubintfw.Writeln("")
		stubintfw.Writeln("/*************************************************************************************************************************")
		stubintfw.Writeln(" Class interface of %s ", outClassName)
		stubintfw.Writeln("**************************************************************************************************************************/")
		stubintfw.Writeln("")

		stubintfw.Writeln("")

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
