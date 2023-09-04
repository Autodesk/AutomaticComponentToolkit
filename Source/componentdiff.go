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
// componentdiff.go
// contains the types and methods to diff componentdefinitions
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/xml"
)

// ComponentDiffBase is the base class for all component diff bases
type ComponentDiffBase struct {
	Path string `xml:"xpath,attr"`
}

// ComponentDiffElementRemove encodes the removal or an element
type ComponentDiffElementRemove struct {
	ComponentDiffBase
	XMLName xml.Name                 `xml:"removeelement"`
	Removal ComponentDiffableElement `xml:"diffable"`
}

// ComponentDiffElementAdd encodes the change of an element
type ComponentDiffElementAdd struct {
	ComponentDiffBase
	XMLName  xml.Name                 `xml:"addelement"`
	Addition ComponentDiffableElement `xml:"diffable"`
}

// ComponentDiffAttributeRemove encodes the removal or a scalar attribute
type ComponentDiffAttributeRemove struct {
	ComponentDiffBase
	XMLName xml.Name `xml:"removeattribute"`
}

// ComponentDiffAttributeAdd encodes the change of a scalar attribute
type ComponentDiffAttributeAdd struct {
	ComponentDiffBase
	XMLName xml.Name `xml:"addattribute"`
}

// ComponentDiffAttributeChange encodes the change of a scalar attribute
type ComponentDiffAttributeChange struct {
	ComponentDiffBase
	XMLName  xml.Name `xml:"changeattribute"`
	OldValue string   `xml:"oldvalue"`
	NewValue string   `xml:"newvalue"`
}

// ComponentDiff contains the difference between two component definitions
type ComponentDiff struct {
	XMLName            xml.Name                       `xml:"componentdiff"`
	AttributeRemovals  []ComponentDiffAttributeRemove `xml:"removeattribute"`
	AttributeAdditions []ComponentDiffAttributeAdd    `xml:"addattribute"`
	AttributeChanges   []ComponentDiffAttributeChange `xml:"changeattribute"`
	ElementRemovals    []ComponentDiffElementRemove   `xml:"removeelement"`
	ElementAdditions   []ComponentDiffElementAdd      `xml:"addelement"`
}

// ComponentDiffableElement is an interface for any element in a componentdefinition that can be diffed
type ComponentDiffableElement interface {
}

func diffParam(path string, paramA ComponentDefinitionParam, paramB ComponentDefinitionParam) ([]ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)

	pathA := path + "/param[@name='" + paramA.ParamName + "']"
	if paramA.ParamDescription != paramB.ParamDescription {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/description"
		change.OldValue = paramA.ParamDescription
		change.NewValue = paramB.ParamDescription
		changes = append(changes, change)
	}

	if paramA.ParamPass != paramB.ParamPass {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/pass"
		change.OldValue = paramA.ParamPass
		change.NewValue = paramB.ParamPass
		changes = append(changes, change)
	}

	if paramA.ParamType != paramB.ParamType {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/type"
		change.OldValue = paramA.ParamType
		change.NewValue = paramB.ParamType
		changes = append(changes, change)
	}

	if paramA.ParamClass != paramB.ParamClass {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/class"
		change.OldValue = paramA.ParamClass
		change.NewValue = paramB.ParamClass
		changes = append(changes, change)
	}

	return changes, nil
}

func diffMethod(path string, methodA ComponentDefinitionMethod, methodB ComponentDefinitionMethod) ([]ComponentDiffElementAdd, []ComponentDiffElementRemove, []ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)
	adds := make([]ComponentDiffElementAdd, 0)
	removes := make([]ComponentDiffElementRemove, 0)

	pathA := path + "/method[@name='" + methodA.MethodName + "']"
	pathB := path + "/method[@name='" + methodB.MethodName + "']"
	if methodA.MethodDescription != methodB.MethodDescription {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/description"
		change.OldValue = methodA.MethodDescription
		change.NewValue = methodB.MethodDescription
		changes = append(changes, change)
	}

	IFirstChangedParam := len(methodA.Params)
	for iA, paramA := range methodA.Params {
		BHasParamA := false
		if (iA < IFirstChangedParam) && (iA < len(methodB.Params)) {
			paramB := methodB.Params[iA]
			if paramA.ParamName == paramB.ParamName {
				Pchanges, err := diffParam(pathA, paramA, paramB)
				if err != nil {
					return adds, removes, changes, err
				}
				changes = append(changes, Pchanges...)
				BHasParamA = true
			}
		}
		if !BHasParamA {
			IFirstChangedParam = iA
			var remove ComponentDiffElementRemove
			remove.Path = pathA
			remove.Removal = paramA
			removes = append(removes, remove)
		}
	}

	for iB, paramB := range methodB.Params {
		AHasParamB := false
		if (iB < IFirstChangedParam) && (iB < len(methodA.Params)) {
			paramA := methodA.Params[iB]
			if paramA.ParamName == paramB.ParamName {
				AHasParamB = true
			}
		}
		if !AHasParamB {
			var add ComponentDiffElementAdd
			add.Path = pathB
			add.Addition = paramB
			adds = append(adds, add)
		}
	}
	return adds, removes, changes, nil
}

func diffClass(path string, classA ComponentDefinitionClass, classB ComponentDefinitionClass) ([]ComponentDiffElementAdd, []ComponentDiffElementRemove, []ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)
	adds := make([]ComponentDiffElementAdd, 0)
	removes := make([]ComponentDiffElementRemove, 0)

	pathA := path + "/class[@name='" + classA.ClassName + "']"
	pathB := path + "/class[@name='" + classB.ClassName + "']"
	if classA.ClassDescription != classB.ClassDescription {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/description"
		change.OldValue = classA.ClassDescription
		change.NewValue = classB.ClassDescription
		changes = append(changes, change)
	}

	if classA.ParentClass != classB.ParentClass {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/parent"
		change.OldValue = classA.ParentClass
		change.NewValue = classB.ParentClass
		changes = append(changes, change)
	}

	for _, methodA := range classA.Methods {
		BHasMethodA := false
		for _, methodB := range classB.Methods {
			if methodA.MethodName == methodB.MethodName {
				BHasMethodA = true
				Madds, Mremoves, Mchanges, err := diffMethod(pathA, methodA, methodB)
				if err != nil {
					return adds, removes, changes, err
				}
				adds = append(adds, Madds...)
				removes = append(removes, Mremoves...)
				changes = append(changes, Mchanges...)
				break
			}
		}
		if !BHasMethodA {
			var remove ComponentDiffElementRemove
			remove.Path = path
			remove.Removal = methodA
			removes = append(removes, remove)
		}
	}

	for _, methodB := range classB.Methods {
		AHasMethodB := false
		for _, methodA := range classA.Methods {
			if methodA.MethodName == methodB.MethodName {
				AHasMethodB = true
				break
			}
		}
		if !AHasMethodB {
			var add ComponentDiffElementAdd
			add.Path = pathB
			add.Addition = methodB
			adds = append(adds, add)
		}
	}

	return adds, removes, changes, nil
}

func diffClasses(path string, classesA []ComponentDefinitionClass, classesB []ComponentDefinitionClass) ([]ComponentDiffElementAdd, []ComponentDiffElementRemove, []ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)
	adds := make([]ComponentDiffElementAdd, 0)
	removes := make([]ComponentDiffElementRemove, 0)

	for _, classA := range classesA {
		BHasClassA := false
		for _, classB := range classesB {
			if classA.ClassName == classB.ClassName {
				BHasClassA = true
				Cadds, Cremoves, Cchanges, err := diffClass(path, classA, classB)
				if err != nil {
					return adds, removes, changes, err
				}
				adds = append(adds, Cadds...)
				removes = append(removes, Cremoves...)
				changes = append(changes, Cchanges...)
				break
			}
		}
		if !BHasClassA {
			var remove ComponentDiffElementRemove
			remove.Path = path
			remove.Removal = classA
			removes = append(removes, remove)
		}
	}

	for _, classB := range classesB {
		AHasClassB := false
		for _, classA := range classesA {
			if classB.ClassName == classA.ClassName {
				AHasClassB = true
				break
			}
		}
		if !AHasClassB {
			var add ComponentDiffElementAdd
			add.Path = path
			add.Addition = classB
			adds = append(adds, add)
		}
	}

	return adds, removes, changes, nil
}

func diffEnum(path string, enumA ComponentDefinitionEnum, enumB ComponentDefinitionEnum) ([]ComponentDiffElementAdd, []ComponentDiffElementRemove, []ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)
	adds := make([]ComponentDiffElementAdd, 0)
	removes := make([]ComponentDiffElementRemove, 0)

	pathA := path + "/enum[@name='" + enumA.Name + "']"
	pathB := path + "/enum[@name='" + enumB.Name + "']"

	for _, optionA := range enumA.Options {
		BHasOptionA := false
		for _, optionB := range enumB.Options {
			if optionA.Name == optionB.Name {
				BHasOptionA = true
				if optionA.Value != optionB.Value {
					var change ComponentDiffAttributeChange
					change.Path = pathA + "/value"
					change.OldValue = string(optionA.Value)
					change.NewValue = string(optionB.Value)
					changes = append(changes, change)
				}
				break
			}
		}
		if !BHasOptionA {
			var remove ComponentDiffElementRemove
			remove.Path = path
			remove.Removal = optionA
			removes = append(removes, remove)
		}
	}

	for _, optionB := range enumB.Options {
		AHasOptionB := false
		for _, optionA := range enumA.Options {
			if optionA.Name == optionB.Name {
				AHasOptionB = true
				break
			}
		}
		if !AHasOptionB {
			var add ComponentDiffElementAdd
			add.Path = pathB
			add.Addition = enumB
			adds = append(adds, add)
		}
	}

	return adds, removes, changes, nil
}

func diffEnums(path string, enumsA []ComponentDefinitionEnum, enumsB []ComponentDefinitionEnum) ([]ComponentDiffElementAdd, []ComponentDiffElementRemove, []ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)
	adds := make([]ComponentDiffElementAdd, 0)
	removes := make([]ComponentDiffElementRemove, 0)

	for _, enumA := range enumsA {
		BHasEnumA := false
		for _, enumB := range enumsB {
			if enumA.Name == enumB.Name {
				BHasEnumA = true
				Eadds, Eremoves, Echanges, err := diffEnum(path, enumA, enumB)
				if err != nil {
					return adds, removes, changes, err
				}
				adds = append(adds, Eadds...)
				removes = append(removes, Eremoves...)
				changes = append(changes, Echanges...)
				break
			}
		}
		if !BHasEnumA {
			var remove ComponentDiffElementRemove
			remove.Path = path
			remove.Removal = enumA
			removes = append(removes, remove)
		}
	}

	for _, enumB := range enumsB {
		AHasEnumB := false
		for _, enumA := range enumsA {
			if enumB.Name == enumA.Name {
				AHasEnumB = true
				break
			}
		}
		if !AHasEnumB {
			var add ComponentDiffElementAdd
			add.Path = path
			add.Addition = enumB
			adds = append(adds, add)
		}
	}

	return adds, removes, changes, nil
}

func diffError(path string, errorA ComponentDefinitionError, errorB ComponentDefinitionError) ([]ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)

	pathA := path + "/error[@name='" + errorA.Name + "']"
	if errorA.Code != errorB.Code {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/code"
		change.OldValue = string(errorA.Code)
		change.NewValue = string(errorB.Code)
		changes = append(changes, change)
	}

	if errorA.Description != errorB.Description {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/description"
		change.OldValue = errorA.Description
		change.NewValue = errorB.Description
		changes = append(changes, change)
	}

	return changes, nil
}

func diffErrors(path string, errorsA []ComponentDefinitionError, errorsB []ComponentDefinitionError) ([]ComponentDiffElementAdd, []ComponentDiffElementRemove, []ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)
	adds := make([]ComponentDiffElementAdd, 0)
	removes := make([]ComponentDiffElementRemove, 0)

	for _, errorA := range errorsA {
		BHasErrorA := false
		for _, errorB := range errorsB {
			if errorA.Name == errorB.Name {
				BHasErrorA = true
				Echanges, err := diffError(path, errorA, errorB)
				if err != nil {
					return adds, removes, changes, err
				}
				changes = append(changes, Echanges...)
				break
			}
		}
		if !BHasErrorA {
			var remove ComponentDiffElementRemove
			remove.Path = path
			remove.Removal = errorA
			removes = append(removes, remove)
		}
	}

	for _, errorB := range errorsB {
		AHasErrorB := false
		for _, errorA := range errorsA {
			if errorB.Name == errorA.Name {
				AHasErrorB = true
				break
			}
		}
		if !AHasErrorB {
			var add ComponentDiffElementAdd
			add.Path = path
			add.Addition = errorB
			adds = append(adds, add)
		}
	}

	return adds, removes, changes, nil
}

func diffGlobal(path string, globalA ComponentDefinitionGlobal, globalB ComponentDefinitionGlobal) ([]ComponentDiffElementAdd, []ComponentDiffElementRemove, []ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)
	adds := make([]ComponentDiffElementAdd, 0)
	removes := make([]ComponentDiffElementRemove, 0)

	pathA := path + "/global"
	pathB := path + "/global"
	if globalA.JournalMethod != globalB.JournalMethod {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/journalmethod"
		change.OldValue = globalA.JournalMethod
		change.NewValue = globalB.JournalMethod
		changes = append(changes, change)
	}
	if globalA.ReleaseMethod != globalB.ReleaseMethod {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/releasemethod"
		change.OldValue = globalA.ReleaseMethod
		change.NewValue = globalB.ReleaseMethod
		changes = append(changes, change)
	}
	if globalA.VersionMethod != globalB.VersionMethod {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/versionmethod"
		change.OldValue = globalA.VersionMethod
		change.NewValue = globalB.VersionMethod
		changes = append(changes, change)
	}

	for _, methodA := range globalA.Methods {
		BHasMethodA := false
		for _, methodB := range globalB.Methods {
			if methodA.MethodName == methodB.MethodName {
				BHasMethodA = true
				Madds, Mremoves, Mchanges, err := diffMethod(pathA, methodA, methodB)
				if err != nil {
					return adds, removes, changes, err
				}
				adds = append(adds, Madds...)
				removes = append(removes, Mremoves...)
				changes = append(changes, Mchanges...)
				break
			}
		}
		if !BHasMethodA {
			var remove ComponentDiffElementRemove
			remove.Path = path
			remove.Removal = methodA
			removes = append(removes, remove)
		}
	}

	for _, methodB := range globalB.Methods {
		AHasMethodB := false
		for _, methodA := range globalA.Methods {
			if methodA.MethodName == methodB.MethodName {
				AHasMethodB = true
				break
			}
		}
		if !AHasMethodB {
			var add ComponentDiffElementAdd
			add.Path = pathB
			add.Addition = methodB
			adds = append(adds, add)
		}
	}

	return adds, removes, changes, nil
}

func diffMember(path string, memberA ComponentDefinitionMember, memberB ComponentDefinitionMember) ([]ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)

	pathA := path + "/member[@name='" + memberA.Name + "']"
	if memberA.Type != memberB.Type {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/type"
		change.OldValue = memberA.Type
		change.NewValue = memberB.Type
		changes = append(changes, change)
	}

	if memberA.Class != memberB.Class {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/class"
		change.OldValue = memberA.Class
		change.NewValue = memberB.Class
		changes = append(changes, change)
	}

	if memberA.Columns != memberB.Columns {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/columns"
		change.OldValue = string(memberA.Columns)
		change.NewValue = string(memberB.Columns)
		changes = append(changes, change)
	}

	if memberA.Rows != memberB.Rows {
		var change ComponentDiffAttributeChange
		change.Path = pathA + "/rows"
		change.OldValue = string(memberA.Rows)
		change.NewValue = string(memberB.Rows)
		changes = append(changes, change)
	}

	return changes, nil
}

func diffStruct(path string, structA ComponentDefinitionStruct, structB ComponentDefinitionStruct) ([]ComponentDiffElementAdd, []ComponentDiffElementRemove, []ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)
	adds := make([]ComponentDiffElementAdd, 0)
	removes := make([]ComponentDiffElementRemove, 0)

	pathA := path + "/structA[@name='" + structA.Name + "']"
	pathB := path + "/structB[@name='" + structB.Name + "']"

	IFirstChangedMember := len(structA.Members)
	for iA, memberA := range structA.Members {
		BHasMemberA := false
		if (iA < IFirstChangedMember) && (iA < len(structB.Members)) {
			memberB := structB.Members[iA]
			if memberA.Name == memberB.Name {
				Pchanges, err := diffMember(pathA, memberA, memberB)
				if err != nil {
					return adds, removes, changes, err
				}
				changes = append(changes, Pchanges...)
				BHasMemberA = true
			}
		}
		if !BHasMemberA {
			var remove ComponentDiffElementRemove
			remove.Path = path
			remove.Removal = memberA
			removes = append(removes, remove)
		}
	}

	for iB, memberB := range structB.Members {
		AHasMemberB := false
		if (iB < IFirstChangedMember) && (iB < len(structA.Members)) {
			memberA := structA.Members[iB]
			if memberB.Name == memberA.Name {
				AHasMemberB = true
			}
		}
		if !AHasMemberB {
			var add ComponentDiffElementAdd
			add.Path = pathB
			add.Addition = memberB
			adds = append(adds, add)
		}
	}

	return adds, removes, changes, nil
}

func diffStructs(path string, structsA []ComponentDefinitionStruct, structsB []ComponentDefinitionStruct) ([]ComponentDiffElementAdd, []ComponentDiffElementRemove, []ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)
	adds := make([]ComponentDiffElementAdd, 0)
	removes := make([]ComponentDiffElementRemove, 0)

	for _, structA := range structsA {
		BHasStructA := false
		for _, structB := range structsB {
			if structA.Name == structB.Name {
				BHasStructA = true
				EAdds, ERemoves, Echanges, err := diffStruct(path, structA, structB)
				if err != nil {
					return adds, removes, changes, err
				}
				adds = append(adds, EAdds...)
				removes = append(removes, ERemoves...)
				changes = append(changes, Echanges...)
				break
			}
		}
		if !BHasStructA {
			var remove ComponentDiffElementRemove
			remove.Path = path
			remove.Removal = structA
			removes = append(removes, remove)
		}
	}

	for _, structB := range structsB {
		AHasStructB := false
		for _, structA := range structsA {
			if structB.Name == structA.Name {
				AHasStructB = true
				break
			}
		}
		if !AHasStructB {
			var add ComponentDiffElementAdd
			add.Path = path
			add.Addition = structB
			adds = append(adds, add)
		}
	}

	return adds, removes, changes, nil
}

func diffComponentAttributes(path string, componentA ComponentDefinition, componentB ComponentDefinition) ([]ComponentDiffAttributeChange, error) {
	changes := make([]ComponentDiffAttributeChange, 0)

	if componentA.Year != componentB.Year {
		var change ComponentDiffAttributeChange
		change.Path = path + "/year"
		change.OldValue = string(componentA.Year)
		change.NewValue = string(componentB.Year)
		changes = append(changes, change)
	}
	if componentA.NameSpace != componentB.NameSpace {
		var change ComponentDiffAttributeChange
		change.Path = path + "/namespace"
		change.OldValue = componentA.NameSpace
		change.NewValue = componentB.NameSpace
		changes = append(changes, change)
	}
	if componentA.LibraryName != componentB.LibraryName {
		var change ComponentDiffAttributeChange
		change.Path = path + "/libraryname"
		change.OldValue = componentA.LibraryName
		change.NewValue = componentB.LibraryName
		changes = append(changes, change)
	}
	if componentA.BaseName != componentB.BaseName {
		var change ComponentDiffAttributeChange
		change.Path = path + "/basename"
		change.OldValue = componentA.BaseName
		change.NewValue = componentB.BaseName
		changes = append(changes, change)
	}
	return changes, nil
}

// DiffComponentDefinitions generates a diff D = B - A between component definitions A and B such that A + D = B
func DiffComponentDefinitions(A ComponentDefinition, B ComponentDefinition) (ComponentDiff, error) {
	var diff ComponentDiff

	path := "/component"

	changes, err := diffComponentAttributes(path, A, B)
	if err != nil {
		return diff, err
	}

	// TODO: check license
	// TODO: check bindings(!?)

	adds, removes, changes, err := diffGlobal(path, A.Global, B.Global)
	if err != nil {
		return diff, err
	}
	diff.ElementAdditions = append(diff.ElementAdditions, adds...)
	diff.ElementRemovals = append(diff.ElementRemovals, removes...)
	diff.AttributeChanges = append(diff.AttributeChanges, changes...)

	adds, removes, changes, err = diffClasses(path, A.Classes, B.Classes)
	if err != nil {
		return diff, err
	}
	diff.ElementAdditions = append(diff.ElementAdditions, adds...)
	diff.ElementRemovals = append(diff.ElementRemovals, removes...)
	diff.AttributeChanges = append(diff.AttributeChanges, changes...)

	adds, removes, changes, err = diffEnums(path, A.Enums, B.Enums)
	if err != nil {
		return diff, err
	}
	diff.ElementAdditions = append(diff.ElementAdditions, adds...)
	diff.ElementRemovals = append(diff.ElementRemovals, removes...)
	diff.AttributeChanges = append(diff.AttributeChanges, changes...)

	adds, removes, changes, err = diffErrors(path+"/Errors", A.Errors.Errors, B.Errors.Errors)
	if err != nil {
		return diff, err
	}
	diff.ElementAdditions = append(diff.ElementAdditions, adds...)
	diff.ElementRemovals = append(diff.ElementRemovals, removes...)
	diff.AttributeChanges = append(diff.AttributeChanges, changes...)

	adds, removes, changes, err = diffStructs(path, A.Structs, B.Structs)
	if err != nil {
		return diff, err
	}
	diff.ElementAdditions = append(diff.ElementAdditions, adds...)
	diff.ElementRemovals = append(diff.ElementRemovals, removes...)
	diff.AttributeChanges = append(diff.AttributeChanges, changes...)
	return diff, nil
}
