/*++

Copyright (C) 2019 Autodesk Inc. (Original Author)

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
// buildbindingccppdocumentation.go.go
// functions to generate the Sphinx documentation of a library's C++-bindings
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"path"
)


// BuildCCPPDocumentation builds the Sphinx documentation of a library's C++-bindings
func BuildCCPPDocumentation(component ComponentDefinition, outputFolder string, ClassIdentifier string) (error) {
	BaseName := component.BaseName

	globalFileName := path.Join(outputFolder, BaseName + ".rst")
	log.Printf("Creating \"%s\"", globalFileName)
	globalDocFile, err := CreateLanguageFile(globalFileName, "\t")
	if err != nil {
		return err
	}
	err = buildCCPPDocumentationGlobal(component, globalDocFile, ClassIdentifier)
	if err != nil {
		return err
	}

	typesFileName := path.Join(outputFolder, BaseName + "-types.rst")
	log.Printf("Creating \"%s\"", typesFileName)
	typesDocFile, err := CreateLanguageFile(typesFileName, "\t")
	if err != nil {
		return err
	}
	err = buildCCPPDocumentationTypes(component, typesDocFile, ClassIdentifier)
	if err != nil {
		return err
	}
	
	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		classFileName := path.Join(outputFolder, BaseName + "_" + class.ClassName + ".rst")
		log.Printf("Creating \"%s\"", classFileName)
		classDocFile, err := CreateLanguageFile(classFileName, "\t")
		if err != nil {
			return err
		}
		err = buildCCPPDocumentationClass(component, classDocFile, class, ClassIdentifier)
		if err != nil {
			return err
		}

	}


	err = buildCCPPDocumentationExample(component, outputFolder, ClassIdentifier, true, "_dynamic")
	if err != nil {
		return err
	}
	err = buildCCPPDocumentationExample(component, outputFolder, ClassIdentifier, false, "_implicit")
	if err != nil {
		return err
	}
	
	return nil
}

func buildCCPPDocumentationExample(component ComponentDefinition, outputFolder string, ClassIdentifier string, ExplicitLinking bool, suffix string) error {
	NameSpace := component.NameSpace

	DynamicCPPExample := path.Join(outputFolder, NameSpace +"_example"+suffix+".cpp")
	log.Printf("Creating \"%s\"", DynamicCPPExample)
	dyncppexamplefile, err := CreateLanguageFile(DynamicCPPExample, "  ")
	if err != nil {
		return err
	}
	buildDynamicCppExample(component, dyncppexamplefile, outputFolder, ClassIdentifier, ExplicitLinking)

	DynamicCPPCMake := path.Join(outputFolder, "CMakeLists"+suffix+".txt")
	log.Printf("Creating \"%s\"", DynamicCPPCMake)
	dyncppcmake, err := CreateLanguageFile(DynamicCPPCMake, "	")
	if err != nil {
		return err
	}
	buildCppDynamicExampleCMake(component, dyncppcmake, outputFolder, ExplicitLinking)
	return nil
}

func writeCPPDocumentationFunctionPointer(component ComponentDefinition, w LanguageWriter,
	functiontype ComponentDefinitionFunctionType) (error) {
	
	NameSpace := component.NameSpace
	returnType := "void"
	parameters := ""

	for j := 0; j < len(functiontype.Params); j++ {
		param := functiontype.Params[j]

		cParamTypeName, err := getCPPParameterTypeName(param.ParamType, NameSpace, param.ParamClass);
		if (err != nil) {
			return err;
		}
		if (parameters != "") {
			parameters = parameters + ", "
		}
		if (param.ParamPass == "in") {
			parameters = parameters + cParamTypeName
		} else {
			parameters = parameters + cParamTypeName + "*"
		}
	}
	w.Writeln("  .. cpp:type:: %s = %s(*)(%s)", functiontype.FunctionName, returnType, parameters)
	w.Writeln("    ")
	w.Writeln("    %s", functiontype.FunctionDescription)
	w.Writeln("    ")

	for j := 0; j < len(functiontype.Params); j++ {
		param := functiontype.Params[j]

		cParams, err := generateCCPPParameter(param, "", functiontype.FunctionName, NameSpace, true)
		if (err != nil) {
			return err;
		}
		for _, cParam := range cParams {
			w.Writeln("    %s", cParam.ParamDocumentationLine);
		}
	}
	w.Writeln("    ")

	return nil
}


func buildCCPPDocumentationGlobal(component ComponentDefinition, w LanguageWriter, ClassIdentifier string) (error) {

	NameSpace := component.NameSpace
	LibraryName := component.LibraryName
	global := component.Global

	wrapperName := "C"+ClassIdentifier+"Wrapper"

	w.Writeln("")
	w.Writeln("The wrapper class %s", wrapperName)
	w.Writeln("===================================================================================")
	w.Writeln("")
	w.Writeln("")
	w.Writeln(".. cpp:class:: %s::%s", NameSpace, wrapperName)

	w.Writeln("")
	w.Writeln("  All types of %s reside in the namespace %s and all", LibraryName, NameSpace)
	w.Writeln("  functionality of %s resides in %s::%s.", LibraryName, NameSpace, wrapperName)
	w.Writeln("")
	w.Writeln("  A suitable way to use %s::%s is as a singleton.", NameSpace, wrapperName)
	w.Writeln("")

	
	for j := 0; j < len(global.Methods); j++ {
		method := global.Methods[j]

		parameters, returntype, err := getDynamicCPPMethodParameters(method, NameSpace, ClassIdentifier, "Wrapper")
		if (err != nil) {
			return err
		}
		w.Writeln("  .. cpp:function:: %s %s(%s)", returntype, method.MethodName, parameters)
		w.Writeln("  ")
		w.Writeln("    %s", method.MethodDescription)
		w.Writeln("  ")
		writeCPPDocumentationParameters(method, w, NameSpace)
		w.Writeln("  ")
	}

	w.Writeln(".. cpp:type:: std::shared_ptr<%s> %s::P%s%s", wrapperName, NameSpace, ClassIdentifier, "Wrapper")
	w.Writeln("  ")
	
	// Load library functions
	// Check error functions of the base class

	return nil
}


func writeCPPDocumentationParameters(method ComponentDefinitionMethod, w LanguageWriter, NameSpace string) {
	for k := 0; k < len(method.Params); k++ {
		param := method.Params[k]
		variableName := getBindingCppVariableName(param)
		if (param.ParamPass == "return") {
			w.Writeln("    :returns: %s", param.ParamDescription )
		} else {
			w.Writeln("    :param %s: %s ", variableName, param.ParamDescription)
		}
	}
	w.Writeln("")
}

func buildCCPPDocumentationClass(component ComponentDefinition, w LanguageWriter, class ComponentDefinitionClass, ClassIdentifier string) (error) {
	
	NameSpace := component.NameSpace
	className := "C"+ClassIdentifier+class.ClassName

	w.Writeln("")
	w.Writeln("%s", className)
	w.Writeln("====================================================================================================")
	w.Writeln("")
	w.Writeln("")
	
	_, inheritanceSpecifier := getCPPInheritanceSpecifier(component, class, "C", ClassIdentifier)

	w.Writeln(".. cpp:class:: %s::%s %s", NameSpace, className, inheritanceSpecifier)
	w.Writeln("")
	w.Writeln("  %s", class.ClassDescription)
	w.Writeln("")
	w.Writeln("")

	w.Writeln("")
	w.Writeln("")
	for j := 0; j < len(class.Methods); j++ {
		method := class.Methods[j]

		parameters, returntype, err := getDynamicCPPMethodParameters(method, NameSpace, ClassIdentifier, class.ClassName)
		if (err != nil) {
			return err
		}
		w.Writeln("  .. cpp:function:: %s %s(%s)", returntype, method.MethodName, parameters)
		w.Writeln("")
		w.Writeln("    %s", method.MethodDescription)
		w.Writeln("")
		writeCPPDocumentationParameters(method, w, NameSpace)
		w.Writeln("")
	}

	w.Writeln(".. cpp:type:: std::shared_ptr<%s> %s::P%s%s", className, NameSpace, ClassIdentifier, class.ClassName)
	w.Writeln("")
	w.Writeln("  Shared pointer to %s to easily allow reference counting.", className)
	w.Writeln("")

	return nil
}

func buildCCPPDocumentationException(component ComponentDefinition, w LanguageWriter) {
	LibraryName := component.LibraryName
	NameSpace := component.NameSpace 

	ExceptionName := "E" + NameSpace + "Exception"
	w.Writeln("  ")
	w.Writeln("%s: The standard exception class of %s", ExceptionName, LibraryName)
	w.Writeln("============================================================================================================================================================================================================")
	w.Writeln("  ")
	w.Writeln("  Errors in %s are reported as Exceptions. It is recommended to not throw these exceptions in your client code.", LibraryName)
	w.Writeln("  ")
	w.Writeln("  ")
	w.Writeln("  .. cpp:class:: %s::%s", NameSpace, ExceptionName)
	w.Writeln("  ")
	w.Writeln("    .. cpp:function:: void %s::what() const noexcept", ExceptionName)
	w.Writeln("    ")
	w.Writeln("       Returns error message")
	w.Writeln("    ")
	w.Writeln("       :return: the error message of this exception")
	w.Writeln("    ")


	w.Writeln("  ")
	w.Writeln("    .. cpp:function:: %sResult %s::getErrorCode() const noexcept", NameSpace, ExceptionName)
	w.Writeln("    ")
	w.Writeln("       Returns error code")
	w.Writeln("    ")
	w.Writeln("       :return: the error code of this exception")
	w.Writeln("    ")
}


func buildCCPPDocumentationInputVector(component ComponentDefinition, w LanguageWriter, ClassIdentifier string) {
	LibraryName := component.LibraryName
	NameSpace := component.NameSpace 

	InputVector := "C" + ClassIdentifier + "InputVector"
	w.Writeln("  ")
	w.Writeln("%s: Adapter for passing arrays as input for functions", InputVector)
	w.Writeln("===============================================================================================================================================================")
	w.Writeln("  ")
	w.Writeln("  Several functions of %s expect arrays of integral types or structs as input parameters.", LibraryName)
	w.Writeln("  To not restrict the interface to, say, std::vector<type>,")
	w.Writeln("  and to have a more abstract interface than a location in memory and the number of elements to input to a function")
	w.Writeln("  %s provides a templated adapter class to pass arrays as input for functions.", LibraryName)
	w.Writeln("  ")
	w.Writeln("  Usually, instances of %s are generated anonymously (or even implicitly) in the call to a function that expects an input array.", InputVector)
	w.Writeln("  ")
	w.Writeln("  ")
	
	
	w.Writeln("  .. cpp:class:: template<typename T> %s::%s", NameSpace, InputVector)
	w.Writeln("  ")
	w.Writeln("    .. cpp:function:: %s(const std::vector<T>& vec)", InputVector)
	w.Writeln("  ")
	w.Writeln("      Constructs of a %s from a std::vector<T>", InputVector)
	w.Writeln("  ")
	w.Writeln("    .. cpp:function:: %s(const T* in_data, size_t in_size)", InputVector)
	w.Writeln("  ")
	w.Writeln("      Constructs of a %s from a memory address and a given number of elements", InputVector)
	w.Writeln("  ")

	w.Writeln("    .. cpp:function:: const T* %s::data() const", InputVector)
	w.Writeln("  ")
	w.Writeln("      returns the start address of the data captured by this %s", InputVector)
	w.Writeln("  ")

	w.Writeln("    .. cpp:function:: size_t %s::size() const", InputVector)
	w.Writeln("  ")
	w.Writeln("      returns the number of elements captured by this %s", InputVector)
	w.Writeln("  ")
	w.Writeln(" ")
}


func buildCCPPDocumentationStructs(component ComponentDefinition, w LanguageWriter) (error) {
	if len(component.Structs) == 0 {
		return nil
	}

	NameSpace := component.NameSpace

	w.Writeln("")
	w.Writeln("Structs")
	w.Writeln("--------------")
	w.Writeln("")
	w.Writeln("  All structs are defined as `packed`, i.e. with the")
	w.Writeln("  ")
	w.Writeln("  .. code-block:: c")
	w.Writeln("    ");
	w.Writeln("    #pragma pack (1)");
	w.Writeln("");

	for i := 0; i < len(component.Structs); i++ {
		structinfo := component.Structs[i];
		w.Writeln("  .. cpp:struct:: s%s", structinfo.Name);
		w.Writeln("  ");
		// w.Writeln("    %s", structinfo.Description);
		// w.Writeln("  ");
		for j := 0; j < len(structinfo.Members); j++ {
			member := structinfo.Members[j];
			arraysuffix := "";
			if (member.Rows > 0) {
				if (member.Columns > 0) {
					arraysuffix = fmt.Sprintf ("[%d][%d]", member.Columns, member.Rows)
				} else {
					arraysuffix = fmt.Sprintf ("[%d]",member.Rows)
				}
			}
			memberLine, err:= getCPPMemberLine(member, NameSpace, arraysuffix, structinfo.Name, "")
			if (err!=nil) {
				return err
			}
			w.Writeln("    .. cpp:member:: %s", memberLine)
			w.Writeln("  ");
		}
		w.Writeln("");
	}

	return nil
}

func buildCCPPDocumentationSimpleTypes(component ComponentDefinition, w LanguageWriter) {
	NameSpace := component.NameSpace

	w.Writeln("Simple types")
	w.Writeln("--------------")
	w.Writeln("")
	types := []string{"uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64"} 
	for _, _type := range types {
		w.Writeln("  .. cpp:type:: %s_t %s_%s", _type, NameSpace, _type)
		w.Writeln("  ")
	}
	w.Writeln("  .. cpp:type:: float %s_single", NameSpace)
	w.Writeln("  ")
	w.Writeln("  .. cpp:type:: double %s_double", NameSpace)
	w.Writeln("  ")
	w.Writeln("  .. cpp:type:: %s_pvoid = void*", NameSpace)
	w.Writeln("  ")
	w.Writeln("  .. cpp:type:: %sResult = %s_int32", NameSpace, NameSpace)
	w.Writeln("  ")
	w.Writeln("  ")
}

func buildCCPPDocumentationEnums(component ComponentDefinition, w LanguageWriter) {
	if len(component.Enums) == 0 {
		return
	}

	NameSpace := component.NameSpace

	w.Writeln("")
	w.Writeln("Enumerations")
	w.Writeln("--------------")
	w.Writeln("")
	for i := 0; i < len(component.Enums); i++ {
		enum := component.Enums[i]
		w.Writeln("  .. cpp:enum-class:: e%s : %s_int32", enum.Name, NameSpace);
		w.Writeln("  ")
		// w.Writeln("  %s", enum.Description)
		// w.Writeln("  ")
		for j := 0; j < len(enum.Options); j++ {
			option := enum.Options[j];
			w.Writeln("    .. cpp:enumerator:: %s = %d", option.Name, option.Value);
		}
		w.Writeln("  ");
	}
}

func buildCCPPDocumentationFunctionTypes(component ComponentDefinition, w LanguageWriter) (error) {
	if len(component.Functions) == 0 {
		return nil
	}
	w.Writeln("")
	w.Writeln("Function types")
	w.Writeln("---------------")
	w.Writeln("")
	w.Writeln("")
	for i := 0; i < len(component.Functions); i++ {
		functiontype := component.Functions[i]
		err := writeCPPDocumentationFunctionPointer(component, w, functiontype)
		if (err!=nil) {
			return err
		}
	}
	w.Writeln("")

	return nil
}

func buildCCPPDocumentationTypes(component ComponentDefinition, w LanguageWriter, ClassIdentifier string) (error) {
	LibraryName := component.LibraryName

	w.Writeln("")
	w.Writeln("Types used in %s", LibraryName)
	w.Writeln("==========================================================================================================")
	w.Writeln("")
	w.Writeln("")

	buildCCPPDocumentationSimpleTypes(component, w)
	buildCCPPDocumentationEnums(component, w)
	
	err := buildCCPPDocumentationStructs(component, w)
	if (err!=nil) {
		return err
	}
	err = buildCCPPDocumentationFunctionTypes(component, w)
	if (err!=nil) {
		return err
	}
	buildCCPPDocumentationException(component, w)
	buildCCPPDocumentationInputVector(component, w, ClassIdentifier)

	return nil
}
