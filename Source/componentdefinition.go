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
// componentdefinition.go
// contains the types used to define a component's API
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"strconv"
	"fmt"
	"errors"
	"encoding/xml"
	"regexp"
	"strings"
	"log"
	"math"
)

const (
	eSpecialMethodNone = 0
	eSpecialMethodRelease = 1
	eSpecialMethodVersion = 2
	eSpecialMethodJournal = 3
)

// ComponentDefinitionParam definition of a method parameter used in the component's API
type ComponentDefinitionParam struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"param"`
	ParamName string `xml:"name,attr"`
	ParamType string `xml:"type,attr"`
	ParamPass string `xml:"pass,attr"`
	ParamClass string `xml:"class,attr"`
	ParamDescription string `xml:"description,attr"`
}

// ComponentDefinitionMethod definition of a method provided by the component's API
type ComponentDefinitionMethod struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"method"`
	MethodName string `xml:"name,attr"`
	MethodDescription string `xml:"description,attr"`
	DLLSuffix string `xml:"dllsuffix,attr"`
	Params   []ComponentDefinitionParam `xml:"param"`
}

// ComponentDefinitionClass definition of a class provided by the component's API
type ComponentDefinitionClass struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"class"`
	ClassName string `xml:"name,attr"`
	ClassDescription string `xml:"description,attr"`
	ParentClass string `xml:"parent,attr"`
	Methods   []ComponentDefinitionMethod `xml:"method"`
}

// ComponentDefinitionFunctionType definition of a function interface provided by the component's API
type ComponentDefinitionFunctionType struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"functiontype"`
	FunctionName string `xml:"name,attr"`
	FunctionDescription string `xml:"description,attr"`
	Params   []ComponentDefinitionParam `xml:"param"`
}

// ComponentDefinitionBindingList definition of the language bindings to be generated for the component's API
type ComponentDefinitionBindingList struct {
	ComponentDiffableElement
	Bindings []ComponentDefinitionBinding `xml:"binding"`
}

// ComponentDefinitionImplementationList definition of the implementation interfaces or stubs to be generated for the component's API
type ComponentDefinitionImplementationList struct {
	ComponentDiffableElement
	Implementations []ComponentDefinitionImplementation `xml:"implementation"`
}

// ComponentDefinitionGlobal definition of global functions provided the component's API
type ComponentDefinitionGlobal struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"global"`
	ReleaseMethod string `xml:"releasemethod,attr"`
	JournalMethod string `xml:"journalmethod,attr"`
	VersionMethod string `xml:"versionmethod,attr"`
	Methods   []ComponentDefinitionMethod `xml:"method"`
}

// ComponentDefinitionBinding definition of a specific languages for which bindings to the component's API will be generated
type ComponentDefinitionBinding struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"binding"`
	Language string `xml:"language,attr"`
	Indentation string `xml:"indentation,attr"`
}

// ComponentDefinitionImplementation definition of a specific languages for which bindings to the component's API will be generated
type ComponentDefinitionImplementation struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"implementation"`
	Language string `xml:"language,attr"`
	Indentation string `xml:"indentation,attr"`
	ClassIdentifier string `xml:"classidentifier,attr"`
	StubIdentifier string `xml:"stubidentifier,attr"`
}

// ComponentDefinitionEnumOption definition of an enum used in the component's API
type ComponentDefinitionEnumOption struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"option"`
	Name string `xml:"name,attr"`
	Value int `xml:"value,attr"`
}

// ComponentDefinitionEnum definition of all enums used in the component's API
type ComponentDefinitionEnum struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"enum"`
	Name string `xml:"name,attr"`
	Options []ComponentDefinitionEnumOption `xml:"option"`
}

// ComponentDefinitionError definition of an error used in the component's API
type ComponentDefinitionError struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"error"`
	Name string `xml:"name,attr"`
	Code int `xml:"code,attr"`
	Description string `xml:"description,attr"`
}

// ComponentDefinitionErrors definition of errors in the component's API
type ComponentDefinitionErrors struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"errors"`
	Errors []ComponentDefinitionError `xml:"error"`
}

// ComponentDefinitionMember definition of a single struct provided by the component's API
type ComponentDefinitionMember struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"member"`
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Class string `xml:"class,attr"`
	Rows int `xml:"rows,attr"`
	Columns int `xml:"columns,attr"`
}

// ComponentDefinitionStruct definition of all structs provided by the component's API
type ComponentDefinitionStruct struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"struct"`
	Name string `xml:"name,attr"`
	Members []ComponentDefinitionMember `xml:"member"`
}

// ComponentDefinitionLicenseLine a single line of the component's license
type ComponentDefinitionLicenseLine struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"line"`
	Value string `xml:"value,attr"`
}

// ComponentDefinitionLicense the component's license
type ComponentDefinitionLicense struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"license"`
	Lines   []ComponentDefinitionLicenseLine `xml:"line"`
}

// ComponentDefinition the complete definition of the component's API
type ComponentDefinition struct {
	ACTVersion string
	XMLName xml.Name `xml:"component"`
	Version string `xml:"version,attr"`
	Copyright string `xml:"copyright,attr"`
	Year int `xml:"year,attr"`
	NameSpace string `xml:"namespace,attr"`
	LibraryName string `xml:"libraryname,attr"`
	BaseName string `xml:"basename,attr"`
	License ComponentDefinitionLicense `xml:"license"`
	Classes []ComponentDefinitionClass `xml:"class"`
	Functions []ComponentDefinitionFunctionType `xml:"functiontype"`
	BindingList ComponentDefinitionBindingList `xml:"bindings"`
	ImplementationList ComponentDefinitionImplementationList `xml:"implementations"`
	Enums []ComponentDefinitionEnum `xml:"enum"`
	Structs []ComponentDefinitionStruct `xml:"struct"`
	Global ComponentDefinitionGlobal `xml:"global"`
	Errors ComponentDefinitionErrors `xml:"errors"`
}

func getIndentationString (str string) string {
	if str == "tabs" {
		return "\t";
	}
	index := strings.Index(str, "spaces");
	if (index < 1) {
		log.Printf ("invalid indentation: \"%s\". Using \"tabs\" instead\n", str);
		return "\t";
	}
	numSpaces, err := strconv.ParseUint(str[0:index], 10, 64);
	if err!=nil {
		log.Printf ("invalid indentation: \"%s\". Using \"4spaces\" instead\n", str);
		return "    ";
	}
	indentString := "";
	var i uint64;
	for i < numSpaces {
		indentString = indentString + " ";
		i++;
	}
	return indentString;
}

func checkImplementations(implementations[] ComponentDefinitionImplementation) error {
	for i := 0; i < len(implementations); i++ {
		implementation := implementations[i]

		if len(implementation.ClassIdentifier) > 0 {
			if !nameSpaceIsValid(implementation.ClassIdentifier) {
				return fmt.Errorf ("Invalid ClassIdentifier in implementation \"%s\"", implementation.Language);
			}
		}
		if len(implementation.StubIdentifier) > 0 {
			if !stubIdentifierIsValid(implementation.StubIdentifier) {
				return fmt.Errorf ("Invalid StubIdentifier in implementation \"%s\"", implementation.Language);
			}
		}
	}
	return nil
}

func checkErrors(errors ComponentDefinitionErrors) error {
	errorNameList := make(map[string]bool, 0);
	errorCodeList := make(map[int]bool, 0);
	for i := 0; i < len(errors.Errors); i++ {
		merror := errors.Errors[i];
		if !nameIsValidIdentifier(merror.Name) {
			return fmt.Errorf( "invalid error name \"%s\"", merror.Name);
		}
		if (errorNameList[strings.ToLower(merror.Name)]) {
			return fmt.Errorf( "duplicate error name \"%s\"", merror.Name);
		}
		errorNameList[strings.ToLower(merror.Name)] = true;

		if (errorCodeList[merror.Code]) {
			return fmt.Errorf( "duplicate error code \"%d\" for error \"%s\"", merror.Code, merror.Name);
		}
		errorCodeList[merror.Code] = true

		if !errorDescriptionIsValid(merror.Description) {
			return fmt.Errorf( "invalid error description \"%s\" for error \"%s\"", merror.Description, merror.Name);
		}
	}
	return nil
}

func errorDescriptionIsValid (name string) bool {
	var IsValidIdentifier = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_+\\-:,.=!/ ]*$").MatchString

	if (name != "") {
		return IsValidIdentifier (name);
	}
	
	return false;
}

func checkOptions(options[] ComponentDefinitionEnumOption) (error) {
	optionLowerNameList := make(map[string]bool, 0);
	optionValueList := make(map[int]bool, 0);

	for j := 0; j < len(options); j++ {
		option := options[j]
		if !nameIsValidIdentifier(option.Name) {
			return fmt.Errorf("invalid option name \"%s\"", option.Name)
		}
		if (math.Abs( float64(option.Value)) > math.Exp2(31) - 1) {
			return fmt.Errorf("option value out of range \"%d\" in \"%s\"", option.Value, option.Name)
		}
		if optionValueList[option.Value] {
			return fmt.Errorf("duplicate option value \"%d\" in \"%s\"", option.Value, option.Name);
		}
		if optionLowerNameList[strings.ToLower(option.Name)] {
			return fmt.Errorf("duplicate option name \"%s\"", option.Name);
		}
		optionValueList[option.Value] = true
		optionLowerNameList[strings.ToLower(option.Name)] = true
	}
	return nil
}

func checkEnums(enums[] ComponentDefinitionEnum) (map[string]bool, error) {
	enumLowerNameList := make(map[string]bool, 0);
	enumNameList := make(map[string]bool, 0);

	for i := 0; i < len(enums); i++ {
		enum := enums[i];
		if !nameIsValidIdentifier(enum.Name) {
			return nil, fmt.Errorf( "invalid enum name \"%s\"", enum.Name);
		}
		
		if (enumLowerNameList[strings.ToLower(enum.Name)]) {
			return nil, fmt.Errorf("duplicate enum name \"%s\"", enum.Name);
		}

		err := checkOptions(enum.Options)
		if err != nil {
			return nil, fmt.Errorf(err.Error() + " in enum = \"%s\"", enum.Name);
		}

		enumLowerNameList[strings.ToLower(enum.Name)] = true
		enumNameList[enum.Name] = true
	}

	return enumNameList, nil
}
	
func checkStructs(structs[] ComponentDefinitionStruct) (map[string]bool, error) {
	structLowerNameList := make(map[string]bool, 0)
	structNameList := make(map[string]bool, 0)

	for i := 0; i < len(structs); i++ {
		mstruct := structs[i];
		if !nameIsValidIdentifier(mstruct.Name) {
			return nil, fmt.Errorf ("invalid struct name \"%s\"", mstruct.Name)
		}
		if structLowerNameList[mstruct.Name] == true {
			return nil, fmt.Errorf ("duplicate struct name \"%s\"", mstruct.Name)
		}
		
		structNameList[mstruct.Name] = true
		structLowerNameList[strings.ToLower(mstruct.Name)] = true
	}
	return structNameList, nil
}

func checkClasses(classes[] ComponentDefinitionClass) (map[string]bool, error) {
	classLowerNameList := make(map[string]bool, 0)
	classNameList := make(map[string]bool, 0)
	for i := 0; i < len(classes); i++ {
		class := classes[i];
		if !nameIsValidIdentifier(class.ClassName) {
			return nil, fmt.Errorf ("invalid class name \"%s\"", class.ClassName);
		}
		if classLowerNameList[strings.ToLower(class.ClassName)] == true {
			return nil, fmt.Errorf ("duplicate class name \"%s\"", class.ClassName);
		}
		if len(class.ClassDescription) > 0 && !descriptionIsValid(class.ClassDescription) {
			return nil, fmt.Errorf ("invalid class description \"%s\" in class \"%s\"", class.ClassDescription, class.ClassName);
		}
		
		classLowerNameList[strings.ToLower(class.ClassName)] = true
		classNameList[class.ClassName] = true
	}

	for i := 0; i < len(classes); i++ {
		class := classes[i];
		parentClass := class.ParentClass;
		if (len(parentClass) > 0) {
			if !nameIsValidIdentifier(parentClass) {
				return nil, fmt.Errorf ("invalid class parent name \"%s\"", parentClass);
			}
			if (classNameList[parentClass] == false) {
				return nil, fmt.Errorf ("unknown parent class \"%s\" for class \"%s\"", parentClass, class.ClassName);
			}
			if (strings.ToLower(class.ClassName) == strings.ToLower(parentClass)) {
				return nil, fmt.Errorf ("class \"%s\" cannot be its own parent class \"%s\"", class.ClassName, parentClass);
			}

		}
	}

	return classNameList, nil
}

func checkFunctionTypes(functions[] ComponentDefinitionFunctionType) (map[string]bool, error) {
	functionLowerNameList := make(map[string]bool, 0)
	functionNameList := make(map[string]bool, 0)
	for i := 0; i < len(functions); i++ {
		function := functions[i];
		if !nameIsValidIdentifier(function.FunctionName) {
			return nil, fmt.Errorf ("invalid functiontype name \"%s\"", function.FunctionName);
		}
		if functionLowerNameList[strings.ToLower(function.FunctionName)] == true {
			return nil, fmt.Errorf ("duplicate functiontype name \"%s\"", function.FunctionName);
		}
		if len(function.FunctionDescription) > 0 && !descriptionIsValid(function.FunctionDescription) {
			return nil, fmt.Errorf ("invalid function description \"%s\" in functiontype \"%s\"", function.FunctionDescription, function.FunctionName);
		}
		
		functionLowerNameList[strings.ToLower(function.FunctionName)] = true
		functionNameList[function.FunctionName] = true
	}
	return functionNameList, nil
}

func checkDuplicateNames(enumList map[string]bool, structList map[string]bool, classList map[string]bool) (error) {
	allLowerList := make(map[string]string, 0)
    for k := range structList {
		if allLowerList[strings.ToLower(k)] == "struct" {
			return fmt.Errorf ("duplicate struct name \"%s\"", k)
		}
		allLowerList[strings.ToLower(k)] = "struct"
	}
	
	for k := range classList {
		if allLowerList[strings.ToLower(k)] == "struct" {
			return fmt.Errorf ("Class with name \"%s\" conflicts with struct of same name", k)
		}
		if allLowerList[strings.ToLower(k)] == "class" {
			return fmt.Errorf ("duplicate class name \"%s\"", k)
		}
		allLowerList[strings.ToLower(k)] = "class"
	}
	
	for k := range enumList {
		if allLowerList[strings.ToLower(k)] == "struct" {
			return fmt.Errorf ("Class with name \"%s\" conflicts with struct of same name", k)
		}
		if allLowerList[strings.ToLower(k)] == "class" {
			return fmt.Errorf ("enum with name \"%s\" conflicts with class of same name", k)
		}
		if allLowerList[strings.ToLower(k)] == "enum" {
			return fmt.Errorf ("duplicate enum name \"%s\"", k)
		}
		allLowerList[strings.ToLower(k)] = "enum"
    }

	return nil
}

func checkClassMethods(classes[] ComponentDefinitionClass, enumList map[string]bool, structList map[string]bool, classList map[string]bool, functionTypeList map[string]bool,) (error) {
	for i := 0; i < len(classes); i++ {
		class := classes[i];				
		methodNameList := make(map[string]bool, 0)
		for j := 0; j < len(class.Methods); j++ {
			method := class.Methods[j]
			if !nameIsValidIdentifier(method.MethodName) {
				return fmt.Errorf ("invalid name for method \"%s.%s\"", class.ClassName, method.MethodName);
			}
			if !descriptionIsValid(method.MethodDescription) {
				return fmt.Errorf ("invalid description for method \"%s.%s\"", class.ClassName, method.MethodName);
			}
			if (methodNameList[strings.ToLower(method.MethodName)]) {
				return fmt.Errorf ("duplicate name for method \"%s.%s\"", class.ClassName, method.MethodName)
			}
			methodNameList[strings.ToLower(method.MethodName)] = true
			
			paramNameList := make(map[string]bool, 0)
			for k := 0; k < len(method.Params); k++ {
				param := method.Params[k]
				if !nameIsValidIdentifier(param.ParamName) {
					return fmt.Errorf ("invalid param name \"%s\" in method \"%s.%s\"", param.ParamName, class.ClassName, method.MethodName);
				}
				if !descriptionIsValid(method.MethodDescription) {
					return fmt.Errorf ("invalid description for parameter \"%s.%s(... %s ...)\"", class.ClassName, method.MethodName, param.ParamName);
				}
				if (paramNameList[strings.ToLower(param.ParamName)]) {
					return fmt.Errorf ("duplicate name \"%s\" for parameter in method \"%s.%s\"", param.ParamName, class.ClassName, method.MethodName)
				}
				paramNameList[strings.ToLower(param.ParamName)] = true

				if (isScalarType(param.ParamType) || param.ParamType == "string") {
					// okay
				} else if (param.ParamType == "handle") {
					if (classList[param.ParamClass] != true) {
						return fmt.Errorf ("parameter \"%s\" of method \"%s.%s\" is of unknown class \"%s\"", param.ParamName, class.ClassName, method.MethodName, param.ParamClass);
					}
				} else if (param.ParamType == "enum") || (param.ParamType == "enumarray") {
					if (enumList[param.ParamClass] != true) {
						return fmt.Errorf ("parameter \"%s\" for method \"%s.%s\" is an unknown enum \"%s\"", param.ParamName, class.ClassName, method.MethodName, param.ParamClass);
					}
				} else if (param.ParamType == "structarray") || (param.ParamType == "struct") {
					if (structList[param.ParamClass] != true) {
						return fmt.Errorf ("parameter \"%s\" for method \"%s.%s\" is an unknown struct \"%s\"", param.ParamName, class.ClassName, method.MethodName, param.ParamClass);
					}
				} else if (param.ParamType == "basicarray") {
					if !isScalarType(param.ParamClass) {
						return fmt.Errorf ("parameter \"%s\" for method \"%s.%s\" is an unknown basic type \"%s\"", param.ParamName, class.ClassName, method.MethodName, param.ParamClass);
					}
				} else if (param.ParamType == "functiontype") {
					if (functionTypeList[param.ParamClass] != true) {
						return fmt.Errorf ("parameter \"%s\" for method \"%s.%s\" is an unknown function type \"%s\"", param.ParamName, class.ClassName, method.MethodName, param.ParamClass);
					}
				} else {
					return fmt.Errorf ("parameter \"%s\" of method \"%s.%s\" is of unknown type \"%s\"", param.ParamName, class.ClassName, method.MethodName, param.ParamType);
				}

			}
		}
	}
	return nil
}

func nameIsValidIdentifier (name string) bool {
	var IsValidIdentifier = regexp.MustCompile("^[A-Z][a-zA-Z0-9_]{0,63}$").MatchString
	if (name != "") {
		return IsValidIdentifier (name);
	}
	return false;
}

func descriptionIsValid (description string) bool {
	var IsValidMethodDescription = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_\\\\/+\\-:,.=!?()'; ]*$").MatchString
	if (description != "") {
		return IsValidMethodDescription (description);
	}
	return false;
}

func isScalarType(typeStr string) bool {
	switch (typeStr) {
		case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "bool", "single", "double":
			return true
	}
	return false
}

func majorVersion (version string) int {
	return versionTriple(version)[0]
}
func minorVersion (version string) int {
	return versionTriple(version)[1]
}
func microVersion (version string) int {
	return versionTriple(version)[2]
}

func versionTriple (version string) [3]int {
	if !versionIsValidVersion(version) {
		log.Fatal("invalid version")
	}
	
	versionTripleR, _ := regexp.Compile("([0-9]*)")
	trip := versionTripleR.FindAllString(version, -1)
	if len(trip) != 3 {
		log.Fatal("invalid version")
	}

	var vers [3]int;
	for i := 0; i < 3; i++ {
		ver, err := strconv.Atoi(trip[i])
		if err != nil {
			log.Fatal("invalid version")
		}
		vers[i] = ver
	}
	return vers
}

func versionIsValidVersion (version string) bool {
	var IsValidVersion = regexp.MustCompile("^([0-9]*)\\.([0-9]*)\\.([0-9]*)$").MatchString
	if (version != "") {
		return IsValidVersion (version);
	}
	return false;
}

func nameSpaceIsValid (namespace string) bool {
	var IsValidNamespace = regexp.MustCompile("^[A-Z][a-zA-Z0-9_]{0,63}$").MatchString
	if (namespace != "") {
		return IsValidNamespace (namespace);
	}
	return false;
}

func stubIdentifierIsValid (stubIdentifier string) bool {
	var IsValidStubIdentifier = regexp.MustCompile("[a-zA-Z0-9_]{0,63}$").MatchString
	if (stubIdentifier != "") {
		return IsValidStubIdentifier (stubIdentifier);
	}
	return false;
}

func libraryNameIsValid (libraryname string) bool {
	var IsLibraryNameValid = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_+\\-:,.=!/ ]*$").MatchString
	if (libraryname != "") {
		return IsLibraryNameValid (libraryname);
	}
	return false;
}

func baseNameIsValid (baseName string) bool {
	var IsBaseNameValid = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_\\-.]*$").MatchString
	if (baseName != "") {
		return IsBaseNameValid (baseName);
	}
	return false;
}

func checkComponentHeader(component ComponentDefinition) (error) {
	if !versionIsValidVersion(component.Version) {
		return fmt.Errorf("Version \"%s\" is invalid", component.Version)
	}
	if component.Copyright == "" {
		return errors.New ("no Copyright information given");
	}
	if (component.Year < 2000) || (component.Year > 2100) {
		return errors.New ("invalid year given");
	}
	if !nameSpaceIsValid(component.NameSpace) {
		return errors.New ("Invalid Namespace");
	}
	if !libraryNameIsValid(component.LibraryName) {
		return errors.New ("Invalid LilbraryName");
	}
	if component.BaseName == "" {
		log.Fatal ("Invalid export basename");
	}
	if !baseNameIsValid(component.BaseName) {
		return errors.New ("Invalid BaseName");
	}
	return nil
}


// CheckComponentDefinition checks a component and returns an error, if it fails
func CheckComponentDefinition (component ComponentDefinition) (error) {
	err := checkComponentHeader(component)
	if err != nil {
		return err
	}
	
	err = checkErrors(component.Errors)
	if err != nil {
		return err
	}

	err = checkImplementations(component.ImplementationList.Implementations)
	if err != nil {
		return err
	}

	var enumList = make(map[string]bool, 0)
	enumList, err = checkEnums(component.Enums)
	if err != nil {
		return err
	}
	
	var structList = make(map[string]bool, 0)
	structList, err = checkStructs(component.Structs)
	if err != nil {
		return err
	}

	var classList = make(map[string]bool, 0)
	classList, err = checkClasses(component.Classes)
	if err != nil {
		return err
	}

	var functionTypeList = make(map[string]bool, 0)
	functionTypeList, err = checkFunctionTypes(component.Functions)
	if err != nil {
		return err
	}

	err = checkDuplicateNames(enumList, structList, classList)
	if err != nil {
		return err
	}

	err = checkClassMethods(component.Classes, enumList, structList, classList, functionTypeList)
	if err != nil {
		return err
	}

	return nil
}


// CheckHeaderSpecialFunction checks a special function of the header against their required definitions
func CheckHeaderSpecialFunction (method ComponentDefinitionMethod, global ComponentDefinitionGlobal) (int, error) {

	if (global.ReleaseMethod == "") {
		return eSpecialMethodNone, errors.New ("No release method specified");
	}

	if (global.VersionMethod == "") {
		return eSpecialMethodNone, errors.New ("No version method specified");
	}

	if (global.ReleaseMethod == global.JournalMethod) {
		return eSpecialMethodNone, errors.New ("Release method can not be the same as the Journal method");
	}

	if (global.ReleaseMethod == global.VersionMethod) {
		return eSpecialMethodNone, errors.New ("Release method can not be the same as the Version method");
	}

	if (global.JournalMethod == global.VersionMethod) {
		return eSpecialMethodNone, errors.New ("Journal method can not be the same as the Version method");
	}
	
	if (method.MethodName == global.ReleaseMethod) {
		if (len (method.Params) != 1) {
			return eSpecialMethodNone, errors.New ("Release method does not match the expected function template");
		}
		
		if (method.Params[0].ParamType != "handle") || (method.Params[0].ParamClass != "BaseClass") || (method.Params[0].ParamPass != "in") {
			return eSpecialMethodNone, errors.New ("Release method does not match the expected function template");
		}

		return eSpecialMethodRelease, nil;
	}

	if (method.MethodName == global.JournalMethod) {
		if (len (method.Params) != 1) {
			return eSpecialMethodNone, errors.New ("Journal method does not match the expected function template");
		}
		
		if (method.Params[0].ParamType != "string") || (method.Params[0].ParamPass != "in") {
			return eSpecialMethodNone, errors.New ("Journal method does not match the expected function template");
		}
		
		return eSpecialMethodJournal, nil;
	}

	if (method.MethodName == global.VersionMethod) {
		if (len (method.Params) != 3) {
			return eSpecialMethodNone, errors.New ("Version method does not match the expected function template");
		}
		
		if (method.Params[0].ParamType != "uint32") || (method.Params[0].ParamPass != "out") || 
			(method.Params[1].ParamType != "uint32") || (method.Params[1].ParamPass != "out") || 
			(method.Params[2].ParamType != "uint32") || (method.Params[2].ParamPass != "out")  {
			return eSpecialMethodNone, errors.New ("Version method does not match the expected function template");
		}
		
		return eSpecialMethodVersion, nil;
	}
	
	return eSpecialMethodNone, nil;
}


