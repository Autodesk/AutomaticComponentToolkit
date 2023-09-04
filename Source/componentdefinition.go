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
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	eSpecialMethodNone         = 0
	eSpecialMethodVersion      = 1
	eSpecialMethodRelease      = 2
	eSpecialMethodAcquire      = 3
	eSpecialMethodError        = 4
	eSpecialMethodInjection    = 5
	eSpecialMethodSymbolLookup = 6
	eSpecialMethodJournal      = 7
	eSpecialMethodPrerelease   = 8
	eSpecialMethodBuildinfo    = 9
)

// ComponentDefinitionParam definition of a method parameter used in the component's API
type ComponentDefinitionParam struct {
	ComponentDiffableElement
	XMLName          xml.Name `xml:"param"`
	ParamName        string   `xml:"name,attr"`
	ParamType        string   `xml:"type,attr"`
	ParamPass        string   `xml:"pass,attr"`
	ParamClass       string   `xml:"class,attr"`
	ParamDescription string   `xml:"description,attr"`
}

// ComponentDefinitionMethod definition of a method provided by the component's API
type ComponentDefinitionMethod struct {
	ComponentDiffableElement
	XMLName           xml.Name                   `xml:"method"`
	MethodName        string                     `xml:"name,attr"`
	MethodDescription string                     `xml:"description,attr"`
	Params            []ComponentDefinitionParam `xml:"param"`
}

// ComponentDefinitionClass definition of a class provided by the component's API
type ComponentDefinitionClass struct {
	ComponentDiffableElement
	XMLName          xml.Name                    `xml:"class"`
	ClassName        string                      `xml:"name,attr"`
	ClassDescription string                      `xml:"description,attr"`
	ParentClass      string                      `xml:"parent,attr"`
	Methods          []ComponentDefinitionMethod `xml:"method"`
}

// ComponentDefinitionFunctionType definition of a function interface provided by the component's API
type ComponentDefinitionFunctionType struct {
	ComponentDiffableElement
	XMLName             xml.Name                   `xml:"functiontype"`
	FunctionName        string                     `xml:"name,attr"`
	FunctionDescription string                     `xml:"description,attr"`
	Params              []ComponentDefinitionParam `xml:"param"`
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
	XMLName            xml.Name                    `xml:"global"`
	BaseClassName      string                      `xml:"baseclassname,attr"`
	ErrorMethod        string                      `xml:"errormethod,attr"`
	ReleaseMethod      string                      `xml:"releasemethod,attr"`
	AcquireMethod      string                      `xml:"acquiremethod,attr"`
	SymbolLookupMethod string                      `xml:"symbollookupmethod,attr"`
	InjectionMethod    string                      `xml:"injectionmethod,attr"`
	JournalMethod      string                      `xml:"journalmethod,attr"`
	VersionMethod      string                      `xml:"versionmethod,attr"`
	PrereleaseMethod   string                      `xml:"prereleasemethod,attr"`
	BuildinfoMethod    string                      `xml:"buildinfomethod,attr"`
	Methods            []ComponentDefinitionMethod `xml:"method"`
}

// ComponentDefinitionBinding definition of a specific languages for which bindings to the component's API will be generated
type ComponentDefinitionBinding struct {
	ComponentDiffableElement
	XMLName         xml.Name `xml:"binding"`
	Language        string   `xml:"language,attr"`
	Indentation     string   `xml:"indentation,attr"`
	ClassIdentifier string   `xml:"classidentifier,attr"`
}

// ComponentDefinitionImplementation definition of a specific languages for which bindings to the component's API will be generated
type ComponentDefinitionImplementation struct {
	ComponentDiffableElement
	XMLName         xml.Name `xml:"implementation"`
	Language        string   `xml:"language,attr"`
	Indentation     string   `xml:"indentation,attr"`
	ClassIdentifier string   `xml:"classidentifier,attr"`
	StubIdentifier  string   `xml:"stubidentifier,attr"`
}

// ComponentDefinitionEnumOption definition of an enum used in the component's API
type ComponentDefinitionEnumOption struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"option"`
	Name    string   `xml:"name,attr"`
	Value   int      `xml:"value,attr"`
}

// ComponentDefinitionEnum definition of all enums used in the component's API
type ComponentDefinitionEnum struct {
	ComponentDiffableElement
	XMLName xml.Name                        `xml:"enum"`
	Name    string                          `xml:"name,attr"`
	Options []ComponentDefinitionEnumOption `xml:"option"`
}

// ComponentDefinitionError definition of an error used in the component's API
type ComponentDefinitionError struct {
	ComponentDiffableElement
	XMLName     xml.Name `xml:"error"`
	Name        string   `xml:"name,attr"`
	Code        int      `xml:"code,attr"`
	Description string   `xml:"description,attr"`
}

// ComponentDefinitionErrors definition of errors in the component's API
type ComponentDefinitionErrors struct {
	ComponentDiffableElement
	XMLName xml.Name                   `xml:"errors"`
	Errors  []ComponentDefinitionError `xml:"error"`
}

// ComponentDefinitionImportComponent definition of errors in the component's API
type ComponentDefinitionImportComponent struct {
	ComponentDiffableElement
	XMLName   xml.Name `xml:"importcomponent"`
	URI       string   `xml:"uri,attr"`
	Namespace string   `xml:"namespace,attr"`
}

// ComponentDefinitionMember definition of a single struct provided by the component's API
type ComponentDefinitionMember struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"member"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
	Rows    int      `xml:"rows,attr"`
	Columns int      `xml:"columns,attr"`
}

// ComponentDefinitionStruct definition of all structs provided by the component's API
type ComponentDefinitionStruct struct {
	ComponentDiffableElement
	XMLName xml.Name                    `xml:"struct"`
	Name    string                      `xml:"name,attr"`
	Members []ComponentDefinitionMember `xml:"member"`
}

// ComponentDefinitionLicenseLine a single line of the component's license
type ComponentDefinitionLicenseLine struct {
	ComponentDiffableElement
	XMLName xml.Name `xml:"line"`
	Value   string   `xml:"value,attr"`
}

// ComponentDefinitionLicense the component's license
type ComponentDefinitionLicense struct {
	ComponentDiffableElement
	XMLName xml.Name                         `xml:"license"`
	Lines   []ComponentDefinitionLicenseLine `xml:"line"`
}

// ComponentDefinition the complete definition of the component's API
type ComponentDefinition struct {
	ACTVersion         string
	XMLName            xml.Name                              `xml:"component"`
	Version            string                                `xml:"version,attr"`
	Copyright          string                                `xml:"copyright,attr"`
	Year               int                                   `xml:"year,attr"`
	NameSpace          string                                `xml:"namespace,attr"`
	LibraryName        string                                `xml:"libraryname,attr"`
	BaseName           string                                `xml:"basename,attr"`
	License            ComponentDefinitionLicense            `xml:"license"`
	Classes            []ComponentDefinitionClass            `xml:"class"`
	Functions          []ComponentDefinitionFunctionType     `xml:"functiontype"`
	BindingList        ComponentDefinitionBindingList        `xml:"bindings"`
	ImplementationList ComponentDefinitionImplementationList `xml:"implementations"`
	Enums              []ComponentDefinitionEnum             `xml:"enum"`
	Structs            []ComponentDefinitionStruct           `xml:"struct"`
	Global             ComponentDefinitionGlobal             `xml:"global"`
	Errors             ComponentDefinitionErrors             `xml:"errors"`
	ImportComponents   []ComponentDefinitionImportComponent  `xml:"importcomponent"`

	ImportedComponentDefinitions map[string]ComponentDefinition
	NameMapsLookup               NameMaps
}

// Normalize adds default values, changes deprecated constants to their later versions
func (component *ComponentDefinition) Normalize() {
	for i := 0; i < len(component.Classes); i++ {
		component.Classes[i].Normalize()
	}
	component.Global.Normalize()

	for _, importedComponent := range component.ImportedComponentDefinitions {
		importedComponent.Normalize()
	}
}

// Normalize adds default values, changes deprecated constants to their later versions
func (global *ComponentDefinitionGlobal) Normalize() {
	for i := 0; i < len(global.Methods); i++ {
		global.Methods[i].Normalize()
	}
}

// Normalize adds default values, changes deprecated constants to their later versions
func (class *ComponentDefinitionClass) Normalize() {
	for i := 0; i < len(class.Methods); i++ {
		class.Methods[i].Normalize()
	}
}

// Normalize adds default values, changes deprecated constants to their later versions
func (method *ComponentDefinitionMethod) Normalize() {
	for i := 0; i < len(method.Params); i++ {
		method.Params[i].Normalize()
	}
}

// Normalize adds default values, changes deprecated constants to their later versions
func (param *ComponentDefinitionParam) Normalize() {
	if param.ParamType == "handle" {
		param.ParamType = "class"
	}
}

// ReadComponentDefinition reads a ComponentDefinition from a file
func ReadComponentDefinition(FileName string, ACTVersion string) (ComponentDefinition, error) {
	var component ComponentDefinition
	component.ImportedComponentDefinitions = make(map[string]ComponentDefinition, 0)
	component.NameMapsLookup = NameMaps{
		enumMap:         make(map[string]bool, 0),
		structMap:       make(map[string]bool, 0),
		classMap:        make(map[string]bool, 0),
		functionTypeMap: make(map[string]bool, 0),
	}

	absFileName, err := filepath.Abs(FileName)
	if err != nil {
		return component, err
	}
	directory := filepath.Dir(absFileName)

	file, err := os.Open(FileName)
	if err != nil {
		return component, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return component, err
	}

	component.ACTVersion = ACTVersion
	err = xml.Unmarshal(bytes, &component)
	if err != nil {
		return component, err
	}

	for i := 0; i < len(component.ImportComponents); i++ {
		importComponent := component.ImportComponents[i]
		subFileName := filepath.Join(directory, importComponent.URI)

		subComponent, err := ReadComponentDefinition(subFileName, ACTVersion)
		if err != nil {
			return component, err
		}
		if subComponent.NameSpace != importComponent.Namespace {
			return component, fmt.Errorf("Namespace of imported component \"%s\" does not match declared namespace \"%s\"", importComponent.Namespace, subComponent.NameSpace)
		}
		component.ImportedComponentDefinitions[importComponent.Namespace] = subComponent
	}
	component.Normalize()

	return component, nil
}

func getIndentationString(str string) string {
	if str == "tabs" {
		return "\t"
	}
	index := strings.Index(str, "spaces")
	if index < 1 {
		log.Printf("invalid indentation: \"%s\". Using \"tabs\" instead\n", str)
		return "\t"
	}
	numSpaces, err := strconv.ParseUint(str[0:index], 10, 64)
	if err != nil {
		log.Printf("invalid indentation: \"%s\". Using \"4spaces\" instead\n", str)
		return "    "
	}
	indentString := ""
	var i uint64
	for i < numSpaces {
		indentString = indentString + " "
		i++
	}
	return indentString
}

func (component *ComponentDefinition) checkImplementations() error {
	implementations := component.ImplementationList.Implementations
	for i := 0; i < len(implementations); i++ {
		implementation := implementations[i]

		if len(implementation.ClassIdentifier) > 0 {
			if !nameSpaceIsValid(implementation.ClassIdentifier) {
				return fmt.Errorf("Invalid ClassIdentifier in implementation \"%s\"", implementation.Language)
			}
		}
		if len(implementation.StubIdentifier) > 0 {
			if !stubIdentifierIsValid(implementation.StubIdentifier) {
				return fmt.Errorf("Invalid StubIdentifier in implementation \"%s\"", implementation.Language)
			}
		}
	}
	return nil
}

func (component *ComponentDefinition) checkErrors() error {
	errorNameList := make(map[string]bool, 0)
	errorCodeList := make(map[int]bool, 0)

	errors := component.Errors
	for i := 0; i < len(errors.Errors); i++ {
		merror := errors.Errors[i]
		if !nameIsValidIdentifier(merror.Name) {
			return fmt.Errorf("invalid error name \"%s\"", merror.Name)
		}
		if errorNameList[strings.ToLower(merror.Name)] {
			return fmt.Errorf("duplicate error name \"%s\"", merror.Name)
		}
		errorNameList[strings.ToLower(merror.Name)] = true

		if errorCodeList[merror.Code] {
			return fmt.Errorf("duplicate error code \"%d\" for error \"%s\"", merror.Code, merror.Name)
		}
		errorCodeList[merror.Code] = true

		if !errorDescriptionIsValid(merror.Description) {
			return fmt.Errorf("invalid error description \"%s\" for error \"%s\"", merror.Description, merror.Name)
		}
	}

	requiredErrors := []string{"NOTIMPLEMENTED", "INVALIDPARAM",
		"INVALIDCAST", "BUFFERTOOSMALL", "GENERICEXCEPTION", "COULDNOTLOADLIBRARY", "COULDNOTFINDLIBRARYEXPORT", "INCOMPATIBLEBINARYVERSION"}
	for _, req := range requiredErrors {
		if !errorNameList[strings.ToLower(req)] {
			return fmt.Errorf("component is missing the required error \"%s\"", req)
		}
	}

	return nil
}

func errorDescriptionIsValid(name string) bool {
	var IsValidIdentifier = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_+\\-:,.=!/ ]*$").MatchString

	if name != "" {
		return IsValidIdentifier(name)
	}

	return false
}

func checkOptions(options []ComponentDefinitionEnumOption) error {
	optionLowerNameList := make(map[string]bool, 0)
	optionValueList := make(map[int]bool, 0)

	for j := 0; j < len(options); j++ {
		option := options[j]
		if !nameIsValidIdentifier(option.Name) {
			return fmt.Errorf("invalid option name \"%s\"", option.Name)
		}
		if math.Abs(float64(option.Value)) > math.Exp2(31)-1 {
			return fmt.Errorf("option value out of range \"%d\" in \"%s\"", option.Value, option.Name)
		}
		if optionValueList[option.Value] {
			return fmt.Errorf("duplicate option value \"%d\" in \"%s\"", option.Value, option.Name)
		}
		if optionLowerNameList[strings.ToLower(option.Name)] {
			return fmt.Errorf("duplicate option name \"%s\"", option.Name)
		}
		optionValueList[option.Value] = true
		optionLowerNameList[strings.ToLower(option.Name)] = true
	}
	return nil
}

func (component *ComponentDefinition) checkEnums() error {
	enums := component.Enums
	var enumNameList = &component.NameMapsLookup.enumMap
	enumLowerNameList := make(map[string]bool, 0)

	for i := 0; i < len(enums); i++ {
		enum := enums[i]
		if !nameIsValidIdentifier(enum.Name) {
			return fmt.Errorf("invalid enum name \"%s\"", enum.Name)
		}

		if enumLowerNameList[strings.ToLower(enum.Name)] {
			return fmt.Errorf("duplicate enum name \"%s\"", enum.Name)
		}

		err := checkOptions(enum.Options)
		if err != nil {
			return fmt.Errorf(err.Error()+" in enum = \"%s\"", enum.Name)
		}

		enumLowerNameList[strings.ToLower(enum.Name)] = true
		(*enumNameList)[enum.Name] = true
	}

	return nil
}

func (component *ComponentDefinition) checkStructs() error {
	structs := component.Structs
	var structNameList = &component.NameMapsLookup.structMap
	structLowerNameList := make(map[string]bool, 0)

	for i := 0; i < len(structs); i++ {
		mstruct := structs[i]
		if !nameIsValidIdentifier(mstruct.Name) {
			return fmt.Errorf("invalid struct name \"%s\"", mstruct.Name)
		}
		if structLowerNameList[mstruct.Name] == true {
			return fmt.Errorf("duplicate struct name \"%s\"", mstruct.Name)
		}
		(*structNameList)[mstruct.Name] = true
		structLowerNameList[strings.ToLower(mstruct.Name)] = true

		for j := 0; j < len(mstruct.Members); j++ {
			member := mstruct.Members[j]
			if !nameIsValidIdentifier(member.Name) {
				return fmt.Errorf("invalid member name \"%s\"", member.Name)
			}
		}
	}
	return nil
}

func (component *ComponentDefinition) checkClasses() error {
	classes := component.Classes
	baseClassName := component.Global.BaseClassName
	var classNameList = &component.NameMapsLookup.classMap

	classLowerNameList := make(map[string]bool, 0)
	classNameIndex := make(map[string]int, 0)
	for i := 0; i < len(classes); i++ {
		class := classes[i]
		if !nameIsValidIdentifier(class.ClassName) {
			return fmt.Errorf("invalid class name \"%s\"", class.ClassName)
		}
		if classLowerNameList[strings.ToLower(class.ClassName)] == true {
			return fmt.Errorf("duplicate class name \"%s\"", class.ClassName)
		}
		if len(class.ClassDescription) > 0 && !descriptionIsValid(class.ClassDescription) {
			return fmt.Errorf("invalid class description \"%s\" in class \"%s\"", class.ClassDescription, class.ClassName)
		}

		classLowerNameList[strings.ToLower(class.ClassName)] = true
		(*classNameList)[class.ClassName] = true
		classNameIndex[class.ClassName] = i
	}

	// Check parent class definitions
	for i := 0; i < len(classes); i++ {
		class := classes[i]
		parentClass := class.ParentClass
		if (baseClassName != class.ClassName) && (len(parentClass) == 0) {
			parentClass = baseClassName
		}
		if len(parentClass) > 0 {
			if !nameIsValidIdentifier(parentClass) {
				return fmt.Errorf("invalid parent class name \"%s\"", parentClass)
			}
			if (*classNameList)[parentClass] == false {
				return fmt.Errorf("unknown parent class \"%s\" for class \"%s\"", parentClass, class.ClassName)
			}
			if classNameIndex[parentClass] >= i {
				return fmt.Errorf("parent class \"%s\" for class \"%s\" is defined after its child class", parentClass, class.ClassName)
			}
			if strings.ToLower(class.ClassName) == strings.ToLower(parentClass) {
				return fmt.Errorf("class \"%s\" cannot be its own parent class \"%s\"", class.ClassName, parentClass)
			}
		}
	}

	return nil
}

func (component *ComponentDefinition) checkFunctionTypes() error {
	functions := component.Functions
	var functionNameList = &component.NameMapsLookup.functionTypeMap

	functionLowerNameList := make(map[string]bool, 0)
	for i := 0; i < len(functions); i++ {
		function := functions[i]
		if !nameIsValidIdentifier(function.FunctionName) {
			return fmt.Errorf("invalid functiontype name \"%s\"", function.FunctionName)
		}
		if functionLowerNameList[strings.ToLower(function.FunctionName)] == true {
			return fmt.Errorf("duplicate functiontype name \"%s\"", function.FunctionName)
		}
		if len(function.FunctionDescription) > 0 && !descriptionIsValid(function.FunctionDescription) {
			return fmt.Errorf("invalid function description \"%s\" in functiontype \"%s\"", function.FunctionDescription, function.FunctionName)
		}

		functionLowerNameList[strings.ToLower(function.FunctionName)] = true
		(*functionNameList)[function.FunctionName] = true
	}
	return nil
}

func checkDuplicateNames(nameMaps NameMaps) error {
	enumList := nameMaps.enumMap
	structList := nameMaps.structMap
	classList := nameMaps.classMap
	functionTypeList := nameMaps.functionTypeMap

	allLowerList := make(map[string]string, 0)

	for k := range structList {
		if val, ok := allLowerList[strings.ToLower(k)]; ok {
			if val == "struct" {
				return fmt.Errorf("duplicate struct name \"%s\"", k)
			}
			return fmt.Errorf("struct with name \"%s\" conflicts with %s of same name", k, val)
		}
		allLowerList[strings.ToLower(k)] = "struct"
	}

	for k := range enumList {
		if val, ok := allLowerList[strings.ToLower(k)]; ok {
			if val == "enum" {
				return fmt.Errorf("duplicate class name \"%s\"", k)
			}
			return fmt.Errorf("enum with name \"%s\" conflicts with %s of same name", k, val)
		}
		allLowerList[strings.ToLower(k)] = "enum"
	}

	for k := range classList {
		if val, ok := allLowerList[strings.ToLower(k)]; ok {
			if val == "class" {
				return fmt.Errorf("duplicate class name \"%s\"", k)
			}
			return fmt.Errorf("class with name \"%s\" conflicts with %s of same name", k, val)
		}
		allLowerList[strings.ToLower(k)] = "class"
	}

	for k := range functionTypeList {
		if val, ok := allLowerList[strings.ToLower(k)]; ok {
			if val == "functiontype" {
				return fmt.Errorf("duplicate functiontype name \"%s\"", k)
			}
			return fmt.Errorf("functiontype with name \"%s\" conflicts with %s of same name", k, val)
		}
		allLowerList[strings.ToLower(k)] = "functiontype"
	}
	return nil
}

func (component *ComponentDefinition) checkMethod(method ComponentDefinitionMethod, className string) error {
	if !nameIsValidIdentifier(method.MethodName) {
		return fmt.Errorf("invalid name for method \"%s.%s\"", className, method.MethodName)
	}
	if !descriptionIsValid(method.MethodDescription) {
		return fmt.Errorf("invalid description for method \"%s.%s\"", className, method.MethodName)
	}

	paramNameList := make(map[string]bool, 0)
	for k := 0; k < len(method.Params); k++ {
		param := method.Params[k]
		if !nameIsValidIdentifier(param.ParamName) {
			return fmt.Errorf("invalid param name \"%s\" in method \"%s.%s\"", param.ParamName, className, method.MethodName)
		}
		if !descriptionIsValid(method.MethodDescription) {
			return fmt.Errorf("invalid description for parameter \"%s.%s(... %s ...)\"", className, method.MethodName, param.ParamName)
		}
		if paramNameList[strings.ToLower(param.ParamName)] {
			return fmt.Errorf("duplicate name \"%s\" for parameter in method \"%s.%s\"", param.ParamName, className, method.MethodName)
		}
		paramNameList[strings.ToLower(param.ParamName)] = true

		if isScalarType(param.ParamType) || param.ParamType == "string" {
			// okay
		} else if param.ParamType == "basicarray" {
			if !isScalarType(param.ParamClass) {
				return fmt.Errorf("parameter \"%s\" for method \"%s.%s\" is an unknown basic type \"%s\"", param.ParamName, className, method.MethodName, param.ParamClass)
			}
		} else {
			currentNameMaps := component.NameMapsLookup

			namespace, paramClassName, err := decomposeParamClassName(param.ParamClass)
			if err != nil {
				return err
			}
			if len(namespace) > 0 {
				if subComponent, ok := component.ImportedComponentDefinitions[namespace]; ok {
					currentNameMaps = subComponent.NameMapsLookup
				} else {
					return fmt.Errorf("parameter \"%s\" of method \"%s.%s\" is of unknown class \"%s\": unknown namespace \"%s\"", param.ParamName, className, method.MethodName, param.ParamClass, namespace)
				}
			}

			if param.ParamType == "class" || param.ParamType == "optionalclass" {
				if currentNameMaps.classMap[paramClassName] != true {
					return fmt.Errorf("parameter \"%s\" of method \"%s.%s\" is of unknown class \"%s\"", param.ParamName, className, method.MethodName, param.ParamClass)
				}
			} else if (param.ParamType == "enum") || (param.ParamType == "enumarray") {
				if currentNameMaps.enumMap[paramClassName] != true {
					return fmt.Errorf("parameter \"%s\" for method \"%s.%s\" is an unknown enum \"%s\"", param.ParamName, className, method.MethodName, param.ParamClass)
				}
			} else if (param.ParamType == "structarray") || (param.ParamType == "struct") {
				if currentNameMaps.structMap[paramClassName] != true {
					return fmt.Errorf("parameter \"%s\" for method \"%s.%s\" is an unknown struct \"%s\"", param.ParamName, className, method.MethodName, param.ParamClass)
				}
			} else if param.ParamType == "functiontype" {
				if currentNameMaps.functionTypeMap[paramClassName] != true {
					return fmt.Errorf("parameter \"%s\" for method \"%s.%s\" is an unknown function type \"%s\"", param.ParamName, className, method.MethodName, param.ParamClass)
				}
			} else {
				return fmt.Errorf("parameter \"%s\" of method \"%s.%s\" is of unknown type \"%s\"", param.ParamName, className, method.MethodName, param.ParamType)
			}
		}

	}

	return nil
}

func (component *ComponentDefinition) checkClassMethods() error {
	classes := component.Classes

	for i := 0; i < len(classes); i++ {
		class := classes[i]

		methodNameList := make(map[string]bool, 0)
		for j := 0; j < len(class.Methods); j++ {
			method := class.Methods[j]

			if methodNameList[strings.ToLower(method.MethodName)] {
				return fmt.Errorf("duplicate name for method \"%s.%s\"", class.ClassName, method.MethodName)
			}
			methodNameList[strings.ToLower(method.MethodName)] = true

			err := component.checkMethod(method, class.ClassName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// decomposeParamClassName decomposes a classname into a namespace and the actual classname within this namespace
func decomposeParamClassName(paramClassName string) (string, string, error) {
	if len(paramClassName) == 0 {
		return "", paramClassName, nil
	}

	namespaceRegexp := "[A-Z][a-zA-Z0-9_]{0,63}"
	var IsValidParamClassName = regexp.MustCompile(fmt.Sprintf("^((%s):){0,1}([a-zA-Z0-9_]{0,64})$", namespaceRegexp))

	if !(IsValidParamClassName.MatchString(paramClassName)) {
		return "", "", fmt.Errorf("param class name \"%s\" is ill formatted", paramClassName)
	}
	slices := IsValidParamClassName.FindStringSubmatch(paramClassName)
	if len(slices) != 4 {
		return "", "", fmt.Errorf("param class name \"%s\" is ill formatted", paramClassName)
	}

	namespace := slices[2]
	className := slices[3]

	return namespace, className, nil
}

func nameIsValidIdentifier(name string) bool {
	var IsValidIdentifier = regexp.MustCompile("^[A-Z][a-zA-Z0-9_]{0,63}$").MatchString
	if name != "" {
		return IsValidIdentifier(name)
	}
	return false
}

func descriptionIsValid(description string) bool {
	var IsValidMethodDescription = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_\\\\/+\\-:,.=!?()'; |]*$").MatchString
	if description != "" {
		return IsValidMethodDescription(description)
	}
	return false
}

func isScalarType(typeStr string) bool {
	switch typeStr {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "bool", "single", "double", "pointer":
		return true
	}
	return false
}

func majorVersion(version string) int {
	isValid, versions, _ := decomposeVersionString(version)
	if !isValid {
		log.Fatal("invalid version")
	}
	return versions[0]
}
func minorVersion(version string) int {
	isValid, versions, _ := decomposeVersionString(version)
	if !isValid {
		log.Fatal("invalid version")
	}
	return versions[1]
}
func microVersion(version string) int {
	isValid, versions, _ := decomposeVersionString(version)
	if !isValid {
		log.Fatal("invalid version")
	}
	return versions[2]
}
func preReleaseInfo(version string) string {
	isValid, _, additionalData := decomposeVersionString(version)
	if !isValid {
		log.Fatal("invalid version")
	}
	return additionalData[0]
}
func buildInfo(version string) string {
	isValid, _, additionalData := decomposeVersionString(version)
	if !isValid {
		log.Fatal("invalid version")
	}
	return additionalData[1]
}

func decomposeVersionString(version string) (bool, [3]int, [2]string) {
	var IsValidVersion = regexp.MustCompile("^([0-9]+)\\.([0-9]+)\\.([0-9]+)(\\-[a-zA-Z0-9.\\-]+)?(\\+[a-zA-Z0-9.\\-]+)?$")

	var vers [3]int
	var data [2]string

	if !(IsValidVersion.MatchString(version)) {
		return false, vers, data
	}
	slices := IsValidVersion.FindStringSubmatch(version)
	if len(slices) != 6 {
		return false, vers, data
	}
	for i := 0; i < 3; i++ {
		ver, err := strconv.Atoi(slices[i+1])
		if err != nil {
			return false, vers, data
		}
		vers[i] = ver
	}
	for i := 0; i < 2; i++ {
		slice := slices[i+4]
		if len(slice) > 0 {
			data[i] = slice[1:]
		}
	}
	return true, vers, data
}

func nameSpaceIsValid(namespace string) bool {
	var IsValidNamespace = regexp.MustCompile("^[A-Z][a-zA-Z0-9_]{0,63}$").MatchString
	if namespace != "" {
		return IsValidNamespace(namespace)
	}
	return false
}

func stubIdentifierIsValid(stubIdentifier string) bool {
	var IsValidStubIdentifier = regexp.MustCompile("[a-zA-Z0-9_]{0,63}$").MatchString
	if stubIdentifier != "" {
		return IsValidStubIdentifier(stubIdentifier)
	}
	return false
}

func libraryNameIsValid(libraryname string) bool {
	var IsLibraryNameValid = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_+\\-:,.=!/ ]*$").MatchString
	if libraryname != "" {
		return IsLibraryNameValid(libraryname)
	}
	return false
}

func baseNameIsValid(baseName string) bool {
	var IsBaseNameValid = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_\\-.]*$").MatchString
	if baseName != "" {
		return IsBaseNameValid(baseName)
	}
	return false
}

func (component *ComponentDefinition) checkComponentHeader() error {
	versionIsValid, _, _ := decomposeVersionString(component.Version)
	if !versionIsValid {
		return fmt.Errorf("Version \"%s\" is invalid", component.Version)
	}
	if component.Copyright == "" {
		return errors.New("no Copyright information given")
	}
	if (component.Year < 2000) || (component.Year > 2100) {
		return errors.New("invalid year given")
	}
	if !nameSpaceIsValid(component.NameSpace) {
		return errors.New("Invalid Namespace")
	}
	if !libraryNameIsValid(component.LibraryName) {
		return errors.New("Invalid LilbraryName")
	}
	if component.BaseName == "" {
		log.Fatal("Invalid export basename")
	}
	if !baseNameIsValid(component.BaseName) {
		return errors.New("Invalid BaseName")
	}
	return nil
}

// NameMaps contains maps of names of elements in a component
type NameMaps struct {
	enumMap         map[string]bool
	structMap       map[string]bool
	classMap        map[string]bool
	functionTypeMap map[string]bool
}

// CheckComponentDefinition checks a component and returns an error, if it fails
func (component *ComponentDefinition) CheckComponentDefinition() error {
	err := component.checkComponentHeader()
	if err != nil {
		return err
	}

	for _, subComponent := range component.ImportedComponentDefinitions {
		err := subComponent.CheckComponentDefinition()
		if err != nil {
			return err
		}
	}

	err = component.checkErrors()
	if err != nil {
		return err
	}

	err = component.checkImplementations()
	if err != nil {
		return err
	}

	err = component.checkEnums()
	if err != nil {
		return err
	}

	err = component.checkStructs()
	if err != nil {
		return err
	}

	err = component.checkClasses()
	if err != nil {
		return err
	}

	err = component.checkFunctionTypes()
	if err != nil {
		return err
	}

	err = checkDuplicateNames(component.NameMapsLookup)
	if err != nil {
		return err
	}

	err = component.checkClassMethods()
	if err != nil {
		return err
	}

	globalMethodNameList := make(map[string]bool, 0)
	for i := 0; i < len(component.Global.Methods); i++ {
		method := component.Global.Methods[i]

		if globalMethodNameList[strings.ToLower(method.MethodName)] {
			return fmt.Errorf("duplicate name for method \"%s.%s\"", "global", method.MethodName)
		}
		globalMethodNameList[strings.ToLower(method.MethodName)] = true

		_, err := CheckHeaderSpecialFunction(method, component.Global)
		if err != nil {
			return err
		}

		err = component.checkMethod(method, "global")
		if err != nil {
			return err
		}
	}

	if component.Global.BaseClassName == "" {
		return errors.New("No base class name specified")
	}
	found := 0
	for i := 0; i < len(component.Classes); i++ {
		if component.Classes[i].ClassName == component.Global.BaseClassName {
			found++
		}
	}
	if found == 0 {
		return errors.New("Specified base class not found")
	} else if found > 1 {
		return errors.New("Base clase defined more than once")
	}
	return nil
}

// CheckHeaderSpecialFunction checks a special function of the header against their required definitions
func CheckHeaderSpecialFunction(method ComponentDefinitionMethod, global ComponentDefinitionGlobal) (int, error) {

	if global.ReleaseMethod == "" {
		return eSpecialMethodNone, errors.New("No release method specified")
	}

	if global.AcquireMethod == "" {
		return eSpecialMethodNone, errors.New("No acquire method specified")
	}

	if global.VersionMethod == "" {
		return eSpecialMethodNone, errors.New("No version method specified")
	}

	if global.ErrorMethod == "" {
		return eSpecialMethodNone, errors.New("No error method specified")
	}

	if global.ReleaseMethod == global.JournalMethod {
		return eSpecialMethodNone, errors.New("Release method can not be the same as the Journal method")
	}

	if global.ReleaseMethod == global.VersionMethod {
		return eSpecialMethodNone, errors.New("Release method can not be the same as the Version method")
	}

	if global.ReleaseMethod == global.AcquireMethod {
		return eSpecialMethodNone, errors.New("Release method can not be the same as the Acquire method")
	}

	if global.JournalMethod == global.VersionMethod {
		return eSpecialMethodNone, errors.New("Journal method can not be the same as the Version method")
	}

	if global.JournalMethod == global.AcquireMethod {
		return eSpecialMethodNone, errors.New("Journal method can not be the same as the Acquire method")
	}

	if method.MethodName == global.ReleaseMethod {
		if len(method.Params) != 1 {
			return eSpecialMethodNone, errors.New("Release method does not match the expected function template")
		}

		if (method.Params[0].ParamType != "class") || (method.Params[0].ParamClass != global.BaseClassName) || (method.Params[0].ParamPass != "in") {
			return eSpecialMethodNone, errors.New("Release method does not match the expected function template")
		}

		return eSpecialMethodRelease, nil
	}

	if method.MethodName == global.AcquireMethod {
		if len(method.Params) != 1 {
			return eSpecialMethodNone, errors.New("Acquire method does not match the expected function template")
		}

		if (method.Params[0].ParamType != "class") || (method.Params[0].ParamClass != global.BaseClassName) || (method.Params[0].ParamPass != "in") {
			return eSpecialMethodNone, errors.New("Acquire method does not match the expected function template")
		}

		return eSpecialMethodAcquire, nil
	}

	if method.MethodName == global.SymbolLookupMethod {
		if len(method.Params) != 1 {
			return eSpecialMethodNone, errors.New("SymbolLookup method does not match the expected function template")
		}
		if (method.Params[0].ParamType != "pointer") || (method.Params[0].ParamPass != "return") {
			return eSpecialMethodNone, errors.New("SymbolLookup method does not match the expected function template")
		}

		return eSpecialMethodSymbolLookup, nil
	}

	if method.MethodName == global.InjectionMethod {
		if len(method.Params) != 2 {
			return eSpecialMethodNone, errors.New("Injection method does not match the expected function template")
		}
		if (method.Params[0].ParamType != "string") || (method.Params[0].ParamPass != "in") ||
			(method.Params[1].ParamType != "pointer") || (method.Params[1].ParamPass != "in") {
			return eSpecialMethodNone, errors.New("Injection method does not match the expected function template")
		}

		return eSpecialMethodInjection, nil
	}

	if method.MethodName == global.JournalMethod {
		if len(method.Params) != 1 {
			return eSpecialMethodNone, errors.New("Journal method does not match the expected function template")
		}

		if (method.Params[0].ParamType != "string") || (method.Params[0].ParamPass != "in") {
			return eSpecialMethodNone, errors.New("Journal method does not match the expected function template")
		}

		return eSpecialMethodJournal, nil
	}

	if method.MethodName == global.VersionMethod {
		if len(method.Params) != 3 {
			return eSpecialMethodNone, errors.New("Version method does not match the expected function template")
		}

		if (method.Params[0].ParamType != "uint32") || (method.Params[0].ParamPass != "out") ||
			(method.Params[1].ParamType != "uint32") || (method.Params[1].ParamPass != "out") ||
			(method.Params[2].ParamType != "uint32") || (method.Params[2].ParamPass != "out") {
			return eSpecialMethodNone, errors.New("Version method does not match the expected function template")
		}

		return eSpecialMethodVersion, nil
	}

	if method.MethodName == global.ErrorMethod {
		if len(method.Params) != 3 {
			return eSpecialMethodNone, errors.New("Error method does not match the expected function template")
		}

		if (method.Params[0].ParamType != "class") || (method.Params[0].ParamPass != "in") ||
			(method.Params[1].ParamType != "string") || (method.Params[1].ParamPass != "out") ||
			(method.Params[2].ParamType != "bool") || (method.Params[2].ParamPass != "return") ||
			(method.Params[0].ParamClass != global.BaseClassName) {
			return eSpecialMethodNone, errors.New("Error method does not match the expected function template")
		}

		return eSpecialMethodError, nil
	}

	if len(global.PrereleaseMethod) > 0 && (global.PrereleaseMethod == global.BuildinfoMethod) {
		return eSpecialMethodNone, errors.New("Prerelease method can not be the same as the buildinfo method")
	}

	if method.MethodName == global.PrereleaseMethod {
		if len(method.Params) != 2 {
			return eSpecialMethodNone, errors.New("Prerelease method does not match the expected function template")
		}

		if (method.Params[0].ParamType != "bool") || (method.Params[0].ParamPass != "return") ||
			(method.Params[1].ParamType != "string") || (method.Params[1].ParamPass != "out") {
			return eSpecialMethodNone, errors.New("Prerelease method does not match the expected function template")
		}

		return eSpecialMethodPrerelease, nil
	}

	if method.MethodName == global.BuildinfoMethod {
		if len(method.Params) != 2 {
			return eSpecialMethodNone, errors.New("Buildinfo method does not match the expected function template")
		}

		if (method.Params[0].ParamType != "bool") || (method.Params[0].ParamPass != "return") ||
			(method.Params[1].ParamType != "string") || (method.Params[1].ParamPass != "out") {
			return eSpecialMethodNone, errors.New("Buildinfo method does not match the expected function template")
		}

		return eSpecialMethodBuildinfo, nil
	}

	return eSpecialMethodNone, nil
}

// GetLastErrorMessageMethod returns the xml definition of the GetLastErrorMessage-method
func GetLastErrorMessageMethod() ComponentDefinitionMethod {
	var method ComponentDefinitionMethod
	source := `<method name="GetLastErrorMessage" description = "Returns the last error registered of this class instance">
		<param name="ErrorMessage" type="string" pass="out" description="Message of the last error registered" />
		<param name="HasLastError" type="bool" pass="return" description="Has an error been registered already" />
	</method>`
	xml.Unmarshal([]byte(source), &method)
	return method
}

// RegisterErrorMessageMethod returns the xml definition of the RegisterErrorMessage-method
func RegisterErrorMessageMethod() ComponentDefinitionMethod {
	var method ComponentDefinitionMethod
	source := `<method name="RegisterErrorMessage" description = "Registers an error message with this class instance">
		<param name="ErrorMessage" type="string" pass="in" description="Error message to register" />
	</method>`
	xml.Unmarshal([]byte(source), &method)
	return method
}

// ClearErrorMessageMethod returns the xml definition of the ClearErrorMessage-method
func ClearErrorMessageMethod() ComponentDefinitionMethod {
	var method ComponentDefinitionMethod
	source := `	<method name="ClearErrorMessages" description = "Clears all registered messages of this class instance">
	</method>`
	xml.Unmarshal([]byte(source), &method)
	return method
}

// IncRefCountMethod returns the xml definition of the IncRefCount-method
func IncRefCountMethod() ComponentDefinitionMethod {
	var method ComponentDefinitionMethod
	source := `<method name="IncRefCount" description = "Increases the reference count of a class instance">
	</method>`
	xml.Unmarshal([]byte(source), &method)
	return method
}

// DecRefCountMethod returns the xml definition of the DecRefCount-method
func DecRefCountMethod() ComponentDefinitionMethod {
	var method ComponentDefinitionMethod
	source := `<method name="DecRefCount" description = "Decreases the reference count of a class instance and free releases it, if the last reference has been removed">
	<param name="HasBeenReleased" type="bool" pass="return" description="Has the object been released" />
	</method>`
	xml.Unmarshal([]byte(source), &method)
	return method
}

// ReleaseBaseClassInterfaceMethod returns the xml definition of a method that should decrease the reference count of a BaseClass interface.
func ReleaseBaseClassInterfaceMethod(baseClassName string) ComponentDefinitionMethod {
	var method ComponentDefinitionMethod
	source := `<method name="ReleaseBaseClassInterface" description = "Releases ownership of a base class interface. Deletes the reference, if necessary.">
		<param name="IBase" type="class" class="` + baseClassName + `" pass="in" description="The base class instance to release" />
	</method>`
	xml.Unmarshal([]byte(source), &method)
	return method
}

// AcquireBaseClassInterfaceMethod returns the xml definition of a method that should increase the reference count of a BaseClass interface.
func AcquireBaseClassInterfaceMethod(baseClassName string) ComponentDefinitionMethod {
	var method ComponentDefinitionMethod
	source := `<method name="AcquireBaseClassInterface" description = "Acquires shared ownership of a base class interface.">
		<param name="IBase" type="class" class="` + baseClassName + `" pass="in" description="The base class instance to acquire" />
	</method>`
	xml.Unmarshal([]byte(source), &method)
	return method
}

func (component *ComponentDefinition) isBaseClass(class ComponentDefinitionClass) bool {
	return class.ClassName == component.Global.BaseClassName
}

func (component *ComponentDefinition) baseClass() ComponentDefinitionClass {
	for i := 0; i < len(component.Classes); i++ {
		if component.isBaseClass(component.Classes[i]) {
			return component.Classes[i]
		}
	}
	var out ComponentDefinitionClass
	log.Fatal("No base class available")
	return out
}
