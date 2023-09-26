/*++

Copyright (C) 2023 Autodesk Inc. (Original Author)

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
// buildimplementationrust .go
// functions to generate Rust interface classes, implementation stubs and wrapper code that maps to
// the rust interfaces.
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strings"
)

// BuildImplementationPascal builds Pascal interface classes, implementation stubs and wrapper code that maps to the Pascal header
func BuildImplementationRust(component ComponentDefinition, outputFolder string, stubOutputFolder string, projectOutputFolder string, implementation ComponentDefinitionImplementation) error {
	forceRebuild := true
	LibraryName := component.LibraryName
	BaseName := component.BaseName
	modfiles := make([]string, 0)
	indentString := getIndentationString(implementation.Indentation)

	stubIdentifier := ""
	if len(implementation.StubIdentifier) > 0 {
		stubIdentifier = "_" + strings.ToLower(implementation.StubIdentifier)
	}

	InterfaceMod := BaseName + "_interfaces"
	IntfFileName := InterfaceMod + ".rs"
	IntfFilePath := path.Join(outputFolder, IntfFileName)
	modfiles = append(modfiles, IntfFilePath)
	log.Printf("Creating \"%s\"", IntfFilePath)
	IntfRSFile, err := CreateLanguageFile(IntfFilePath, indentString)
	if err != nil {
		return err
	}
	IntfRSFile.WriteCLicenseHeader(component,
		fmt.Sprintf("This is an autogenerated rust file in order to allow easy\ndevelopment of %s. The implementer of %s needs to\nderive concrete classes from the abstract classes in this header.", LibraryName, LibraryName),
		true)
	err = writeRustBaseTypeDefinitions(component, IntfRSFile, component.NameSpace, BaseName)
	if err != nil {
		return err
	}
	err = buildRustInterfaces(component, IntfRSFile)
	if err != nil {
		return err
	}

	IntfWrapperFileName := BaseName + "_interface_wrapper.rs"
	IntfWrapperFilePath := path.Join(outputFolder, IntfWrapperFileName)
	modfiles = append(modfiles, IntfWrapperFilePath)
	log.Printf("Creating \"%s\"", IntfWrapperFilePath)
	IntfWrapperRSFile, err := CreateLanguageFile(IntfWrapperFilePath, indentString)
	if err != nil {
		return err
	}
	IntfWrapperRSFile.WriteCLicenseHeader(component,
		fmt.Sprintf("This is an autogenerated Rust implementation file in order to allow easy\ndevelopment of %s. The functions in this file need to be implemented. It needs to be generated only once.", LibraryName),
		true)
	err = buildRustWrapper(component, IntfWrapperRSFile, InterfaceMod)
	if err != nil {
		return err
	}

	IntfHandleFileName := BaseName + "_interface_handle.rs"
	IntfHandleFilePath := path.Join(outputFolder, IntfHandleFileName)
	modfiles = append(modfiles, IntfHandleFilePath)
	log.Printf("Creating \"%s\"", IntfHandleFilePath)
	IntfHandleRSFile, err := CreateLanguageFile(IntfHandleFilePath, indentString)
	if err != nil {
		return err
	}
	IntfHandleRSFile.WriteCLicenseHeader(component,
		fmt.Sprintf("This is an autogenerated Rust implementation file in order to allow easy\ndevelopment of %s. The functions in this file need to be implemented. It needs to be generated only once.", LibraryName),
		true)
	err = buildRustHandle(component, IntfHandleRSFile, InterfaceMod)
	if err != nil {
		return err
	}

	IntfWrapperStubName := path.Join(stubOutputFolder, BaseName+stubIdentifier+".rs")
	modfiles = append(modfiles, IntfWrapperStubName)
	if forceRebuild || !FileExists(IntfWrapperStubName) {
		log.Printf("Creating \"%s\"", IntfWrapperStubName)
		stubfile, err := CreateLanguageFile(IntfWrapperStubName, indentString)
		if err != nil {
			return err
		}
		stubfile.WriteCLicenseHeader(component,
			fmt.Sprintf("This is an autogenerated Rust implementation file in order to allow easy\ndevelopment of %s. It needs to be generated only once.", LibraryName),
			true)
		if err != nil {
			return err
		}
		err = buildRustGlobalStubFile(component, stubfile, InterfaceMod)
		if err != nil {
			return err
		}
	} else {
		log.Printf("Omitting recreation of implementation stub \"%s\"", IntfWrapperStubName)
	}

	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		StubBase := BaseName + stubIdentifier
		StubClassName := path.Join(stubOutputFolder, StubBase+"_"+toSnakeCase(class.ClassName)+".rs")
		modfiles = append(modfiles, StubClassName)
		if forceRebuild || !FileExists(StubClassName) {
			log.Printf("Creating \"%s\"", StubClassName)
			stubfile, err := CreateLanguageFile(StubClassName, indentString)
			if err != nil {
				return err
			}
			stubfile.WriteCLicenseHeader(component,
				fmt.Sprintf("This is an autogenerated Rust implementation file in order to allow easy\ndevelopment of %s. It needs to be generated only once.", LibraryName),
				true)
			if err != nil {
				return err
			}
			err = buildRustStubFile(component, class, stubfile, InterfaceMod, StubBase)
			if err != nil {
				return err
			}
		} else {
			log.Printf("Omitting recreation of implementation stub \"%s\"", StubClassName)
		}
	}

	if len(projectOutputFolder) > 0 {
		IntfWrapperLibName := path.Join(projectOutputFolder, "lib.rs")
		if forceRebuild || !FileExists(IntfWrapperLibName) {
			log.Printf("Creating \"%s\"", IntfWrapperLibName)
			libfile, err := CreateLanguageFile(IntfWrapperLibName, indentString)
			if err != nil {
				return err
			}
			libfile.WriteCLicenseHeader(component,
				fmt.Sprintf("This is an autogenerated Rust implementation file in order to allow easy\ndevelopment of %s. It needs to be generated only once.", LibraryName),
				true)
			err = buildRustGlobalLibFile(component, libfile, projectOutputFolder, modfiles)
			if err != nil {
				return err
			}
		} else {
			log.Printf("Omitting recreation of lib \"%s\"", IntfWrapperLibName)
		}

		CargoFileName := path.Join(projectOutputFolder, "Cargo.toml")
		if forceRebuild || !FileExists(CargoFileName) {
			log.Printf("Creating Cargo file \"%s\" for Rust Implementation", CargoFileName)
			CargoFile, err := CreateLanguageFile(CargoFileName, indentString)
			if err != nil {
				return err
			}
			CargoFile.WriteTomlLicenseHeader(component,
				fmt.Sprintf("This is an autogenerated Cargo file for the development of %s.", LibraryName),
				true)
			LibPath, err := filepath.Rel(projectOutputFolder, IntfWrapperLibName)
			if err != nil {
				return err
			}
			buildCargoForRustImplementation(component, CargoFile, LibPath)
		} else {
			log.Printf("Omitting recreation of Cargo file \"%s\" for Rust Implementation", CargoFileName)
		}
	}

	return nil
}

func buildRustGlobalLibFile(component ComponentDefinition, w LanguageWriter, basedir string, modfiles []string) error {
	w.Writeln("")
	w.Writeln("#![feature(trait_upcasting)]")
	w.Writeln("#![allow(incomplete_features)]")
	w.Writeln("#![feature(vec_into_raw_parts)]")
	w.Writeln("")
	// Get all modules
	for i := 0; i < len(modfiles); i++ {
		modfile := modfiles[i]
		relfile, err := filepath.Rel(basedir, modfile)
		if err != nil {
			return err
		}
		w.Writeln("#[path = \"%s\"]", strings.ReplaceAll(relfile, "\\", "/"))
		IntfName := strings.TrimSuffix(filepath.Base(relfile), ".rs")
		w.Writeln("mod %s;", IntfName)
		w.Writeln("")
	}
	return nil
}

func buildRustInterfaces(component ComponentDefinition, w LanguageWriter) error {
	NameSpace := component.NameSpace
	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Traits defined for %s", NameSpace)
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("")
	for i := 0; i < len(component.Classes); i++ {
		classinfo := component.Classes[i]
		err := writeRustTrait(component, classinfo, w)
		if err != nil {
			return err
		}
	}
	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Trait defined for global methods of %s", NameSpace)
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("")
	err := writeRustGlobalTrait(component, w)
	if err != nil {
		return err
	}
	return nil
}

func buildCargoForRustImplementation(component ComponentDefinition, w LanguageWriter, path string) error {
	projectName := strings.ToLower(component.NameSpace)
	w.Writeln("[package]")
	w.Writeln("  name = \"%s\"", projectName)
	w.Writeln("  version = \"0.1.0\"")
	w.Writeln("[lib]")
	w.Writeln("  path = \"%s\"", strings.ReplaceAll(path, "\\", "/"))
	w.Writeln("  crate-type = [\"cdylib\"]")
	return nil
}

func writeRustTrait(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
	w.Writeln("// Trait for interface %s", class.ClassName)
	w.Writeln("//")
	if class.ClassDescription != "" {
		w.Writeln("// %s", class.ClassDescription)
		w.Writeln("//")
	}
	parentClassString := ""
	if !component.isBaseClass(class) {
		if class.ParentClass == "" {
			parentClassString = fmt.Sprintf(": %s ", component.Global.BaseClassName)
		} else {
			parentClassString = fmt.Sprintf(": %s ", class.ParentClass)
		}
	}
	w.Writeln("pub trait %s %s {", class.ClassName, parentClassString)
	w.AddIndentationLevel(1)
	methods := class.Methods
	if component.isBaseClass(class) {
		methods = append(
			methods,
			GetLastErrorMessageMethod(),
			ClearErrorMessageMethod(),
			RegisterErrorMessageMethod(),
			IncRefCountMethod(),
			DecRefCountMethod())
	}

	for j := 0; j < len(methods); j++ {
		method := methods[j]
		w.Writeln("")
		err := writeRustTraitFn(method, w, true, false, false)
		if err != nil {
			return err
		}
	}
	w.ResetIndentationLevel()
	w.Writeln("}")
	w.Writeln("")
	w.Writeln("")
	return nil
}

func writeRustTraitFn(method ComponentDefinitionMethod, w LanguageWriter, hasSelf bool, hasImpl bool, hasImplParent bool) error {
	methodName := toSnakeCase(method.MethodName)
	w.Writeln("// %s", methodName)
	w.Writeln("//")
	w.Writeln("// %s", method.MethodDescription)
	parameterString := ""
	parameterNames := ""
	if hasSelf {
		parameterString += "&mut self"
	}
	returnType := ""
	for k := 0; k < len(method.Params); k++ {
		param := method.Params[k]
		RustParams, err := generateRustParameters(param, false)
		if err != nil {
			return err
		}
		RustParam := RustParams[0]
		if param.ParamPass != "return" {
			if parameterString == "" {
				parameterString += fmt.Sprintf("%s : %s", RustParam.ParamName, RustParam.ParamType)
			} else {
				parameterString += fmt.Sprintf(", %s : %s", RustParam.ParamName, RustParam.ParamType)
			}
			if parameterNames == "" {
				parameterNames += RustParam.ParamName
			} else {
				parameterNames += fmt.Sprintf(", %s", RustParam.ParamName)
			}
		} else {
			returnType = RustParam.ParamType
		}
		w.Writeln("// %s", RustParam.ParamComment)
	}
	w.Writeln("//")
	if !hasImpl {
		if returnType == "" {
			w.Writeln("fn %s(%s);", methodName, parameterString)
		} else {
			w.Writeln("fn %s(%s) -> %s;", methodName, parameterString, returnType)
		}
	} else {
		if returnType == "" {
			w.Writeln("fn %s(%s) {", methodName, parameterString)
		} else {
			w.Writeln("fn %s(%s) -> %s {", methodName, parameterString, returnType)
		}
		w.AddIndentationLevel(1)
		if !hasImplParent {
			w.Writeln("unimplemented!();")
		} else {
			w.Writeln("self.parent.%s(%s)", methodName, parameterNames)
		}
		w.AddIndentationLevel(-1)
		w.Writeln("}")
	}
	return nil
}

func writeRustGlobalTrait(component ComponentDefinition, w LanguageWriter) error {
	w.Writeln("// Wrapper trait for global methods")
	w.Writeln("//")
	w.Writeln("pub trait Wrapper {")
	w.AddIndentationLevel(1)
	methods := component.Global.Methods
	for j := 0; j < len(methods); j++ {
		method := methods[j]
		w.Writeln("")
		err := writeRustTraitFn(method, w, false, false, false)
		if err != nil {
			return err
		}
	}
	w.ResetIndentationLevel()
	w.Writeln("}")
	return nil
}

func buildRustGlobalStubFile(component ComponentDefinition, w LanguageWriter, InterfaceMod string) error {
	w.Writeln("")
	w.Writeln("use %s::*;", InterfaceMod)
	w.Writeln("")
	w.Writeln("// Wrapper struct to implement the wrapper trait for global methods")
	w.Writeln("pub struct CWrapper;")
	w.Writeln("")
	w.Writeln("impl Wrapper for CWrapper {")
	w.Writeln("")
	w.AddIndentationLevel(1)
	methods := component.Global.Methods
	for j := 0; j < len(methods); j++ {
		method := methods[j]
		w.Writeln("")
		err := writeRustTraitFn(method, w, false, true, false)
		if err != nil {
			return err
		}
	}
	w.ResetIndentationLevel()
	w.Writeln("}")
	w.Writeln("")
	return nil
}

func getParentList(component ComponentDefinition, class ComponentDefinitionClass) ([]string, error) {
	parents := make([]string, 0)
	currClass := class
	for !component.isBaseClass(currClass) {
		parent := currClass.ParentClass
		if parent == "" {
			parent = component.baseClass().ClassName
		}
		parents = append(parents, parent)
		parClass, err := getClass(component, parent)
		if err != nil {
			return parents, err
		}
		currClass = parClass
	}
	return parents, nil
}

func getChildList(component ComponentDefinition, class ComponentDefinitionClass) []string {
	children := make([]string, 0)
	for i := 0; i < len(component.Classes); i++ {
		child := component.Classes[i]
		if child.ParentClass == class.ClassName {
			children = append(children, child.ClassName)
			children = append(children, getChildList(component, child)...)
		}
	}
	return children
}

func getClass(component ComponentDefinition, name string) (ComponentDefinitionClass, error) {
	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		if class.ClassName == name {
			return class, nil
		}
	}
	return component.baseClass(), fmt.Errorf("Cannot find class %s", name)
}

func buildRustStubFile(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter, InterfaceMod string, StubBase string) error {
	Name := class.ClassName
	parents, err := getParentList(component, class)
	if err != nil {
		return err
	}
	w.Writeln("")
	w.Writeln("use %s::*;", InterfaceMod)
	if len(parents) > 0 {
		parentName := parents[0]
		w.Writeln("use %s_%s::C%s;", StubBase, toSnakeCase(parentName), parentName)
	}
	w.Writeln("")
	w.Writeln("// Stub struct to implement the %s trait", Name)
	if len(parents) == 0 {
		w.Writeln("pub struct C%s;", Name)
	} else {
		w.Writeln("pub struct C%s {", Name)
		w.AddIndentationLevel(1)
		w.Writeln("parent : C%s", parents[0])
		w.ResetIndentationLevel()
		w.Writeln("}")
		w.Writeln("")
		w.Writeln("// Implementation of parent traits via parent")
		w.Writeln("")
		for i := 0; i < len(parents); i++ {
			parent := parents[i]
			parentClass, err := getClass(component, parent)
			if err != nil {
				return err
			}
			w.Writeln("impl %s for C%s {", parent, Name)
			w.AddIndentationLevel(1)
			methods := parentClass.Methods
			if component.isBaseClass(parentClass) {
				methods = append(
					methods,
					GetLastErrorMessageMethod(),
					ClearErrorMessageMethod(),
					RegisterErrorMessageMethod(),
					IncRefCountMethod(),
					DecRefCountMethod())
			}
			for j := 0; j < len(methods); j++ {
				method := methods[j]
				w.Writeln("")
				err := writeRustTraitFn(method, w, true, true, true)
				if err != nil {
					return err
				}
			}
			w.ResetIndentationLevel()
			w.Writeln("}")
		}
	}
	w.Writeln("")
	w.Writeln("impl %s for C%s {", Name, Name)
	w.Writeln("")
	w.AddIndentationLevel(1)
	methods := class.Methods
	if component.isBaseClass(class) {
		methods = append(
			methods,
			GetLastErrorMessageMethod(),
			ClearErrorMessageMethod(),
			RegisterErrorMessageMethod(),
			IncRefCountMethod(),
			DecRefCountMethod())
	}
	for j := 0; j < len(methods); j++ {
		method := methods[j]
		w.Writeln("")
		err := writeRustTraitFn(method, w, true, true, false)
		if err != nil {
			return err
		}
	}
	w.ResetIndentationLevel()
	w.Writeln("}")
	w.Writeln("")
	return nil
}

func buildRustWrapper(component ComponentDefinition, w LanguageWriter, InterfaceMod string) error {
	// Imports
	ModName := strings.ToLower(component.NameSpace)
	w.Writeln("")
	w.Writeln("// Calls from the C-Interface to the Rust traits via the CWrapper")
	w.Writeln("// These are the symbols exposed in the shared object interface")
	w.Writeln("")
	w.Writeln("use %s::*;", InterfaceMod)
	w.Writeln("use %s::CWrapper;", ModName)
	w.Writeln("use std::ffi::{c_char, CStr};")
	w.Writeln("")
	cprefix := ModName + "_"
	// Build the global methods
	err := writeGlobalRustWrapper(component, w, cprefix)
	if err != nil {
		return err
	}
	return nil
}

func buildRustHandle(component ComponentDefinition, w LanguageWriter, InterfaceMod string) error {
	w.Writeln("")
	w.Writeln("// Handle passed through interface define the casting maps needed to extract")
	w.Writeln("")
	w.Writeln("use %s::*;", InterfaceMod)
	w.Writeln("")
	w.Writeln("impl HandleImpl {")
	w.AddIndentationLevel(1)
	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		writeRustHandleAs(component, w, class, false)
		writeRustHandleAs(component, w, class, true)
		w.Writeln("")
	}
	w.AddIndentationLevel(-1)
	w.Writeln("}")
	return nil
}

func writeRustHandleAs(component ComponentDefinition, w LanguageWriter, class ComponentDefinitionClass, mut bool) error {
	children := getChildList(component, class)
	Name := class.ClassName
	if !mut {
		w.Writeln("pub fn as_%s(&self) -> Option<&dyn %s> {", toSnakeCase(Name), Name)
	} else {
		w.Writeln("pub fn as_mut_%s(&mut self) -> Option<&mut dyn %s> {", toSnakeCase(Name), Name)
	}
	w.AddIndentationLevel(1)
	w.Writeln("match self {")
	w.AddIndentationLevel(1)
	for i := 0; i < len(children); i++ {
		child := children[i]
		if !mut {
			w.Writeln("HandleImpl::T%s(ptr) => Some(ptr.as_ref()),", child)
		} else {
			w.Writeln("HandleImpl::T%s(ptr) => Some(ptr.as_mut()),", child)
		}
	}
	w.Writeln("_ => None")
	w.AddIndentationLevel(-1)
	w.Writeln("}")
	w.AddIndentationLevel(-1)
	w.Writeln("}")
	return nil
}

func writeGlobalRustWrapper(component ComponentDefinition, w LanguageWriter, cprefix string) error {
	errorprefix := strings.ToUpper(component.NameSpace)
	methods := component.Global.Methods
	for i := 0; i < len(methods); i++ {
		method := methods[i]
		err := writeRustMethodWrapper(method, w, cprefix, errorprefix)
		if err != nil {
			return err
		}
		w.Writeln("")
	}
	return nil
}

func writeRustMethodWrapper(method ComponentDefinitionMethod, w LanguageWriter, cprefix string, errorprefix string) error {
	// Build up the parameter strings
	parameterString := ""
	returnName := ""
	for k := 0; k < len(method.Params); k++ {
		param := method.Params[k]
		RustParams, err := generateRustParameters(param, true)
		if err != nil {
			return err
		}
		for i := 0; i < len(RustParams); i++ {
			RustParam := RustParams[i]
			if parameterString == "" {
				parameterString += fmt.Sprintf("%s : %s", RustParam.ParamName, RustParam.ParamType)
			} else {
				parameterString += fmt.Sprintf(", %s : %s", RustParam.ParamName, RustParam.ParamType)
			}
		}
	}
	w.Writeln("pub fn %s%s(%s) -> i32 {", cprefix, strings.ToLower(method.MethodName), parameterString)
	w.AddIndentationLevel(1)
	argsString := ""
	for k := 0; k < len(method.Params); k++ {
		param := method.Params[k]
		OName, err := writeRustParameterConversionArg(param, w, errorprefix)
		if err != nil {
			return err
		}
		if OName != "" {
			if argsString == "" {
				argsString = OName
			} else {
				argsString += fmt.Sprintf(", %s", OName)
			}
		}
	}
	if returnName != "" {
		w.Writeln("let %s = CWrapper::%s(%s);", returnName, toSnakeCase(method.MethodName), argsString)
	} else {
		w.Writeln("CWrapper::%s(%s);", toSnakeCase(method.MethodName), argsString)
	}
	for k := 0; k < len(method.Params); k++ {
		param := method.Params[k]
		err := writeRustParameterConversionOutPost(param, w, errorprefix)
		if err != nil {
			return err
		}
	}
	w.Writeln("// All ok")
	w.Writeln("%s_SUCCESS", errorprefix)
	w.AddIndentationLevel(-1)
	w.Writeln("}")
	return nil
}

func writeRustParameterConversionArg(param ComponentDefinitionParam, w LanguageWriter, errorprefix string) (string, error) {
	if param.ParamPass == "return" {
		return "", nil
	}
	IName := toSnakeCase(param.ParamName)
	OName := "_" + IName
	switch param.ParamType {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "single", "double":
		if param.ParamPass == "in" {
			w.Writeln("let %s = %s;", OName, IName)
		} else {
			w.Writeln("if %s.is_null() { return %s_ERROR_INVALIDPARAM; }", IName, errorprefix)
			w.Writeln("let %s = unsafe {&mut *%s};", OName, IName)
		}
	case "class", "optionalclass":
		if param.ParamPass == "in" {
			HName := "_handle_" + IName
			OpName := "_optional_" + IName
			w.Writeln("if %s.is_null() { return %s_ERROR_INVALIDPARAM; }", IName, errorprefix)
			w.Writeln("let %s = unsafe {&*%s};", HName, IName)
			w.Writeln("let %s = %s.as_%s();", OpName, HName, toSnakeCase(param.ParamClass))
			w.Writeln("if %s.is_none() { return %s_ERROR_INVALIDPARAM; }", OpName, errorprefix)
			w.Writeln("let %s = %s.unwrap();", OName, OpName)

		} else {
			HName := "_handle_" + IName
			OpName := "_optional_" + IName
			w.Writeln("if %s.is_null() { return %s_ERROR_INVALIDPARAM; }", IName, errorprefix)
			w.Writeln("let %s = unsafe {&mut *%s};", HName, IName)
			w.Writeln("let %s = %s.as_mut_%s();", OpName, HName, toSnakeCase(param.ParamClass))
			w.Writeln("if %s.is_none() { return %s_ERROR_INVALIDPARAM; }", OpName, errorprefix)
			w.Writeln("let %s = %s.unwrap();", OName, OpName)
		}
	case "string":
		if param.ParamPass == "in" {
			SName := "_str_" + IName
			OpName := "_optional_" + IName
			w.Writeln("if %s.is_null() { return %s_ERROR_INVALIDPARAM; }", IName, errorprefix)
			w.Writeln("let %s = unsafe{ CStr::from_ptr(%s) };", SName, IName)
			w.Writeln("let %s = %s.to_str();", OpName, SName)
			w.Writeln("if %s.is_err() { return %s_ERROR_INVALIDPARAM; }", OpName, errorprefix)
			w.Writeln("let %s = %s.unwrap();", OName, OpName)
		} else {
			SName := "_string_" + IName
			w.Writeln("let mut %s = String::new();", SName)
			w.Writeln("let %s = &mut %s;", OName, SName)
		}
	case "bool", "pointer", "struct", "basicarray", "structarray":
		//return fmt.Errorf("Conversion of type %s for parameter %s not supported", param.ParamType, IName)
	default:
		return "", fmt.Errorf("Conversion of type %s for parameter %s not supported as is unknown", param.ParamType, IName)
	}
	return OName, nil
}

func writeRustParameterConversionOutPost(param ComponentDefinitionParam, w LanguageWriter, errorprefix string) error {
	if param.ParamPass != "out" {
		return nil
	}
	// Any remaining bit needed to wire out variables
	IName := toSnakeCase(param.ParamName)
	switch param.ParamType {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "single", "double", "class", "optionalclass":
		return nil
	case "string":
		// Check the buffer size and if null
		BuffSizeName := IName + "_buffer_size"
		BuffName := IName + "_buffer"
		BuffSLName := "_buffer_slice_" + IName
		CNName := IName + "_needed_chars"
		CNRName := "_" + CNName
		SName := "_string_" + IName
		SLName := "_slice_" + IName
		w.Writeln("let %s = %s.as_bytes();", SLName, SName)
		w.Writeln("if %s > %s.len() { return %s_ERROR_BUFFERTOOSMALL; }", BuffSizeName, SLName, errorprefix)
		w.Writeln("if %s.is_null() { return %s_ERROR_INVALIDPARAM; }", BuffName, errorprefix)
		w.Writeln("if %s.is_null() { return %s_ERROR_INVALIDPARAM; }", CNName, errorprefix)
		w.Writeln("let mut %s = unsafe {  std::slice::from_raw_parts_mut(%s, %s.len()) };", BuffSLName, BuffName, SLName)
		w.Writeln("%s.clone_from_slice(%s);", BuffSLName, SLName)
		w.Writeln("let mut %s = unsafe { &mut *%s };", CNRName, CNName)
		w.Writeln("*%s = %s.len();", CNRName, SLName)
	case "bool", "pointer", "struct", "basicarray", "structarray":
		//return fmt.Errorf("Conversion of type %s for parameter %s not supported", param.ParamType, IName)
	default:
		return fmt.Errorf("Conversion of type %s for parameter %s not supported as is unknown", param.ParamType, IName)
	}
	return nil
}
