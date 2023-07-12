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
// buildimplementationts.go
// Builds typescript interface definitions.
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"log"
	"path"
	"strings"
	"errors"
	"unicode"
)

type TypeScriptOptions struct {
	Camelize bool
	JsArrays bool
	LineLength int
}

// BuildImplementationTS builds typescript interface definitions
func BuildImplementationTS(
	component ComponentDefinition, 
	outputFolder string, 
	implementation ComponentDefinitionImplementation,
) error {

	log.Printf("Creating TypeScript Implementation")

	options := TypeScriptOptions{
		Camelize: false,
		JsArrays: false,
		LineLength: 80,
	}

	filename := path.Join(outputFolder, "fusion-environment-act.d.ts")
	log.Printf("Creating \"%s\"", filename)
	file, err := CreateLanguageFile(filename, "	")
	if err != nil {
		return err
	}

	file.WriteCLicenseHeader(component, "", true)
	file.Writeln("")
	file.Writeln("export namespace %s {", component.NameSpace)

	err = writeTypescriptInterfaces(component, file, options)
	if err != nil {
		return err
	}	
	err = writeTypescriptEnums(component, file, options)
	if err != nil {
		return err
	}	
	file.Writeln("} // %s ", component.NameSpace)

	return nil
}

func writeTypescriptEnums(
	component ComponentDefinition,
	writer LanguageWriter,
	options TypeScriptOptions,
) error {
		
	for _, enum := range component.Enums {
		err := writeTypescriptEnum(enum, writer, options)
		if err != nil {
			return err
		}	
	}
	return nil
}

func writeTypescriptEnum(
	enum ComponentDefinitionEnum,
	writer LanguageWriter,
	options TypeScriptOptions,
) error {
	writer.Writeln("const enum %s {", getId(enum.Name, options))
	writer.Indentation++
	for _, option := range enum.Options {
		writeCommentEnumOption(option, writer, options)
		identifier := getId(option.Name, options)
		value := option.Value
		writer.Writeln("%s = %d,", identifier, value)
	}
	writer.Indentation--		
	writer.Writeln("}")
	writer.Writeln("")
	return nil
}

func writeTypescriptInterfaces(
	component ComponentDefinition, 
	writer LanguageWriter, 
	options TypeScriptOptions,
) error {
	for _, class := range component.Classes {
		err := writeTypescriptInterface(class, writer, options)
		if err != nil {
			return err
		}	
	}
	return nil
}

func writeTypescriptInterface(
	class ComponentDefinitionClass, 
	writer LanguageWriter, 
	options TypeScriptOptions,
) error {
	writeCommentClass(class, writer, options)
	identifier := getId(class.ClassName, options)
	extends := ""
	if (class.ParentClass != "") {
		extends = "extends " + class.ParentClass + " "
	}
	writer.Writeln("interface %s %s{", identifier, extends)

	for _, method := range class.Methods {
		writer.Indentation++
    var err error
    if (method.PropertyGet != "" || method.PropertySet != "" ) {
      err = writeTypescriptProperty(class, method, writer, options)
    } else {
		  err = writeTypescriptMethod(class, method, writer, options)
    }
		if err != nil {
			return err
		}	
		writer.Indentation--
	}
	writer.Writeln("}")
	writer.Writeln("")
	return nil
}

func writeTypescriptMethod(
	class ComponentDefinitionClass, 
	method ComponentDefinitionMethod, 
	writer LanguageWriter, 
	options TypeScriptOptions,
) error {
	inParams := filterPass(method.Params, "in")
	outParams := filterPass(method.Params, "out")
	returnParams := filterPass(method.Params, "return")

	writer.Writeln("");
	writeCommentMethod(class, method, writer, options)
	writer.BeginLine()
	writer.Printf("%s: (", getId(method.MethodName, options))
	for i, param := range inParams {
		writer.Printf(
			"%s: %s", 
			getId(param.ParamName, options), 
			getType(param, options),
		)
		if (i + 1 < len(inParams)) {
			writer.Printf(", ")
		}
	}
	writer.Printf(") => ")

	if (len(outParams) > 0) {
		writer.Printf("[")
		for i, param := range outParams {
			writer.Printf(
				"%s: %s", 
				getId(param.ParamName, options), 
				getType(param, options),
			)
			if (i + 1 < len(outParams)) {
				writer.Printf(", ")
			}
		}
		writer.Printf("]")
	} else {
		if (len(returnParams) > 1) {
			return errors.New("More than one return value.")
		} else if (len(returnParams) == 1) {
			writer.Printf(getType(returnParams[0], options))
		} else {
			writer.Printf("void")
		}
	}

	writer.Printf(";")
	writer.EndLine()
	return nil
} 

func writeTypescriptProperty(
	class ComponentDefinitionClass, 
	method ComponentDefinitionMethod, 
	writer LanguageWriter, 
	options TypeScriptOptions,
) error {
  if (method.PropertySet != "") {
    // Ignore setters
    return nil
  }
  getter := &method
  var setter *ComponentDefinitionMethod
  for _, method := range class.Methods {
    if method.PropertySet == getter.PropertyGet {
      setter = &method
      continue
    }
  }
  returnParams := filterPass(getter.Params, "return")
  if (len(returnParams) != 1) {
    return errors.New("Property getters should have a single return value.")
  }
  inParams := filterPass(setter.Params, "in")
  if (len(inParams) != 1) {
    return errors.New("Property setters should have a single input parameter")
  }
  readOnly := "readonly "
  if (setter != nil) {
    readOnly = ""
  }
  writer.Writeln("")
  writeCommentProperty(class, *getter, writer, options)
  writer.Writeln(
    "%s%s: %s;", 
    readOnly,
    getId(getter.PropertyGet, options), 
    getType(returnParams[0], options),
  )
  return nil
}

func filterPass(
	params []ComponentDefinitionParam, 
	pass string,
) []ComponentDefinitionParam {
	var result []ComponentDefinitionParam;
	for _, param := range params {
		if (param.ParamPass == pass) {
			result = append(result, param)
		}
	}	
	return result;
}

func getId(identifier string, options TypeScriptOptions) string {
	if (options.Camelize) {
		return camelize(identifier)
	}
	return identifier
}

func getType(
	param ComponentDefinitionParam, 
	options TypeScriptOptions,
) string {
	return getTypeString(param.ParamType, param.ParamClass, options);
}

func getTypeString(
	paramType string, 
	paramClass string,
	options TypeScriptOptions,
) string {
	if (paramType == "class" || paramType == "enum") {
		if (options.JsArrays && strings.HasSuffix(paramClass, "Vector")) {
			return strings.TrimSuffix(paramClass, "Vector") + "[]"
		}
		return paramClass
	} else if (paramType == "basicarray") {
		return getTypeString(paramClass, "", options) + "[]"
	} else if (
		paramType == "double" ||
		paramType == "int16"	||
		paramType == "int32"	||
		paramType == "int64"	||
		paramType == "uint16" ||
		paramType == "uint32" ||
		paramType == "uint64") {
		return "number"
	} else if (paramType == "bool") {
		return "boolean"
	}
	return paramType
}

func camelize(identifier string) string {
  if len(identifier) == 0 {
    return identifier
  }
  result := []rune(identifier)
  result[0] = unicode.ToLower(result[0])
  return string(result)
}

func writeCommentEnumOption(
	option ComponentDefinitionEnumOption, 
	writer LanguageWriter,
	options TypeScriptOptions,
) {
	writer.Writeln("/**")
	lines := getCommentLines(" * ", option.Description, writer, options)
	for _, line := range lines {
		writer.Writeln(" * " + line)
	}
	writer.Writeln(" */")
}

func writeCommentClass(
	class ComponentDefinitionClass, 
	writer LanguageWriter,
	options TypeScriptOptions,
) {
	writer.Writeln("/**")
	lines := getCommentLines(" * ", class.ClassDescription, writer, options)
	for _, line := range lines {
		writer.Writeln(" * " + line)
	}
	writer.Writeln(" */")
}

func writeCommentMethod(
	class ComponentDefinitionClass,
	method ComponentDefinitionMethod, 
	writer LanguageWriter,
	options TypeScriptOptions,
) {
	writer.Writeln("/**")
	lines := getCommentLines(" * ", method.MethodDescription, writer, options)
	for _, line := range lines {
		writer.Writeln(" * " + line)
	}
	inParams := filterPass(method.Params, "in")
	outParams := filterPass(method.Params, "out")
	returnParams := filterPass(method.Params, "return")

	writeCommentInParams(inParams, writer, options) 
	if (len(outParams) > 0) {
		writeCommentOutParams(outParams, writer, options)
	} else {
		writeCommentReturnParams(returnParams, writer, options)
	}
	writer.Writeln(" */")
}

func writeCommentProperty(
	class ComponentDefinitionClass,
	method ComponentDefinitionMethod, 
	writer LanguageWriter,
	options TypeScriptOptions,
) {
	writer.Writeln("/**")
	lines := getCommentLines(" * ", method.MethodDescription, writer, options)
	for _, line := range lines {
		writer.Writeln(" * " + line)
	}
	writer.Writeln(" */")
}

func writeCommentInParams(
	params []ComponentDefinitionParam, 
	writer LanguageWriter,
	options TypeScriptOptions,
) {
	for _, param := range params {
		prefix := " * @param {" + getType(param, options) + "} " + 
							getId(param.ParamName, options) + " "
		lines := getCommentLines(prefix, param.ParamDescription, writer, options)
		if (len(lines) > 0) {
			writer.Writeln(prefix + lines[0])
			prefix = " * " + strings.Repeat(" ", len(prefix) - len(" * "))
			for i := 1; i < len(lines); i++	{
				line := lines[i]
				writer.Writeln(prefix + line)
			}	
		} 
	}	
}

func writeCommentOutParams(
	params []ComponentDefinitionParam, 
	writer LanguageWriter,
	options TypeScriptOptions,
) {
	for _, param := range params {
		prefix := " * @returns {" + getType(param, options) + "} "
		prefix2 := prefix + getId(param.ParamName, options) + " "
		lines := getCommentLines(prefix2, param.ParamDescription, writer, options)
		if (len(lines) > 0) {
			writer.Writeln(prefix2 + lines[0])
			prefix2 = " * " + strings.Repeat(" ", len(prefix2) - len(" * "))
			for i := 1; i < len(lines); i++	{
				line := lines[i]
				writer.Writeln(prefix2 + line)
			}	
		} 
	}
}

func writeCommentReturnParams(
	params []ComponentDefinitionParam, 
	writer LanguageWriter,
	options TypeScriptOptions,
) {
	for _, param := range params {
		prefix := " * @returns {" + getType(param, options) + "} "
		lines := getCommentLines(prefix, param.ParamDescription, writer, options)
		if (len(lines) > 0) {
			writer.Writeln(prefix + lines[0])
			prefix = " * " + strings.Repeat(" ", len(prefix) - len(" * "))
			for i := 1; i < len(lines); i++	{
				line := lines[i]
				writer.Writeln(prefix + line)
			}	
		} 
	}
}

func getCommentLines(
	prefix string,
	comment string,
	writer LanguageWriter,	 
	options TypeScriptOptions,
) []string {
	indent := strings.Repeat(writer.IndentString, writer.Indentation)
	lineLength := options.LineLength - len(indent) - len(prefix)
	return getLines(comment, lineLength)
}

func getLines(input string, width int) []string {
	words := strings.Fields(input)
	var lines []string
	line := ""
	for _, word := range words {
		if len(line)+len(word) > width {
			lines = append(lines, line)
			line = word + " " 
		} else {
			line += word + " "
		}
	}
	lines = append(lines, line)
	return lines
}
