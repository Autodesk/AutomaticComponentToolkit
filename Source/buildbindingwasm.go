/*++

Copyright (C) 2024 Autodesk Inc. (Original Author)

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
// buildwasmbindingwasm.go
// functions to generate WASM bindings which rely on emscripten
// It produces a bindings.cpp file which needs to be compiled using emcc / em++ and needs cpp bindings
// in path
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Store the wasm binding path (could be useful later)
var wasmBindingFile = ""

// Type mapping (More things should move here)
func ResolveCppType(paramType string, componentdefinition ComponentDefinition) string {
	switch paramType {
	case "string":
		return "std::string"
	case "bool":
		return "bool"
	default:
		return fmt.Sprintf("%s_%s", componentdefinition.LibraryName, paramType)
	}
}

// shouldSkipMethod checks if a method should be skipped based on its parameter names.
func shouldSkipMethod(method ComponentDefinitionMethod) bool {
	for _, param := range method.Params {
		if param.ParamName == "SymbolLookupMethod" || param.ParamName == "SymbolAddressMethod" {
			return true // Found a parameter that triggers the skip, so we can return immediately.
		}
	}
	return false // No parameters matched, do not skip.
}

// Check if a static wrapper is needed
func NeedsStaticWrapper(m ComponentDefinitionMethod) bool {
	// 1) Skip methods with out-params or callbacks entirely
	for _, p := range m.Params {
		if p.ParamPass == "out" || p.ParamType == "functiontype" {
			return false
		}
	}
	// 2) If it returns a struct or structarray, wrap it
	for _, p := range m.Params {
		if p.ParamPass == "return" &&
			(p.ParamType == "struct" || p.ParamType == "structarray") {
			return true
		}
	}
	// 3) If it takes a struct/structarray as input, wrap it
	for _, p := range m.Params {
		if (p.ParamPass == "" || p.ParamPass == "in") &&
			(p.ParamType == "struct" || p.ParamType == "structarray" || p.ParamType == "class" || p.ParamType == "optionalclass") {
			return true
		}
	}
	// otherwise no wrapper needed
	return false
}

// Returns true if we need to emit an out-param wrapper
func NeedsOutParamWrapper(m ComponentDefinitionMethod) bool {
	for _, p := range m.Params {
		if p.ParamPass == "out" && p.ParamType != "functiontype" {
			return true
		}
	}
	return false
}

// Generate wrapper structs
func GenerateStructWrappers(componentdefinition ComponentDefinition) string {
	var builder strings.Builder

	for _, s := range componentdefinition.Structs {
		builder.WriteString(fmt.Sprintf("struct s%sWrapper {\n", s.Name))
		builder.WriteString(fmt.Sprintf("    s%s value;\n", s.Name))

		for _, m := range s.Members {
			typeString := ""
			if m.Type == "enum" {
				typeString = fmt.Sprintf("e%s", m.Class)
			} else {
				typeString = ResolveCppType(m.Type, componentdefinition)
			}

			if m.Rows > 0 || m.Columns > 0 {
				if m.Rows > 0 && m.Columns > 0 {
					// 2D Array
					for i := 0; i < m.Columns; i++ {
						for j := 0; j < m.Rows; j++ {
							builder.WriteString(fmt.Sprintf("    %s get_%s_%d_%d() const { return value.m_%s[%d][%d]; }\n",
								typeString, m.Name, i, j, m.Name, i, j))
							builder.WriteString(fmt.Sprintf("    void set_%s_%d_%d(%s v) { value.m_%s[%d][%d] = v; }\n",
								m.Name, i, j, typeString, m.Name, i, j))
						}
					}
				} else {
					// 1D Array
					n := m.Rows
					if n == 0 {
						n = m.Columns
					}
					for i := 0; i < n; i++ {
						builder.WriteString(fmt.Sprintf("    %s get_%s%d() const { return value.m_%s[%d]; }\n",
							typeString, m.Name, i, m.Name, i))
						builder.WriteString(fmt.Sprintf("    void set_%s%d(%s v) { value.m_%s[%d] = v; }\n",
							m.Name, i, typeString, m.Name, i))
					}
				}
			} else {
				// Scalar
				builder.WriteString(fmt.Sprintf("    %s get_%s() const { return value.m_%s; }\n",
					typeString, m.Name, m.Name))
				builder.WriteString(fmt.Sprintf("    void set_%s(%s v) { value.m_%s = v; }\n",
					m.Name, typeString, m.Name))
			}
		}

		// toStruct
		builder.WriteString(fmt.Sprintf("\n    s%s toStruct() const { return value; }\n", s.Name))

		// fromStruct
		builder.WriteString(fmt.Sprintf("\n    static s%sWrapper fromStruct(const emscripten::val &js) {\n", s.Name))
		builder.WriteString(fmt.Sprintf("        s%sWrapper wrapper;\n", s.Name))
		for _, m := range s.Members {
			typeString := ""
			if m.Type == "enum" {
				typeString = fmt.Sprintf("e%s", m.Class)
			} else {
				typeString = ResolveCppType(m.Type, componentdefinition)
			}

			if m.Rows > 0 || m.Columns > 0 {
				if m.Rows > 0 && m.Columns > 0 {
					// 2D Array
					for i := 0; i < m.Rows; i++ {
						for j := 0; j < m.Columns; j++ {
							builder.WriteString(fmt.Sprintf("        wrapper.value.m_%s[%d][%d] = js[\"%s_%d_%d\"].as<%s>();\n",
								m.Name, j, i, m.Name, i, j, typeString))
						}
					}
				} else {
					// 1D Array
					n := m.Rows
					if n == 0 {
						n = m.Columns
					}
					for i := 0; i < n; i++ {
						builder.WriteString(fmt.Sprintf("        wrapper.value.m_%s[%d] = js[\"%s%d\"].as<%s>();\n",
							m.Name, i, m.Name, i, typeString))
					}
				}
			} else {
				// Scalar
				builder.WriteString(fmt.Sprintf("        wrapper.value.m_%s = js[\"%s\"].as<%s>();\n",
					m.Name, m.Name, typeString))
			}
		}
		builder.WriteString("        return wrapper;\n")
		builder.WriteString("    }\n")

		builder.WriteString("};\n\n")
	}

	return builder.String()
}

// Generate necessary Cpp headers
func GenerateCppHeader(componentdefinition ComponentDefinition) string {
	var builder strings.Builder

	builder.WriteString("#include <iostream>\n")
	builder.WriteString("#include <vector>\n")
	builder.WriteString("#include <string>\n")
	builder.WriteString("#include <emscripten.h>\n")
	builder.WriteString("#include <emscripten/bind.h>\n")
	builder.WriteString(fmt.Sprintf("#include \"Cpp/%s_implicit.hpp\"\n", componentdefinition.BaseName))
	builder.WriteString("\n")
	builder.WriteString("using namespace emscripten;\n")
	builder.WriteString(fmt.Sprintf("using namespace %s;\n\n", componentdefinition.NameSpace))

	return builder.String()
}

// A common function to generate method wrappers (struct, struct arrays, class pointers and globals)
func generateMethodWrappers(
	component ComponentDefinition,
	isOutParam bool,
) string {
	var result strings.Builder

	emitOne := func(method ComponentDefinitionMethod, className string, isGlobal bool) {
		// 1. Predicate check
		if isOutParam && !NeedsOutParamWrapper(method) {
			return
		}
		if !isOutParam && !NeedsStaticWrapper(method) {
			return
		}

		// 2. Skip functiontype and callback
		for _, p := range method.Params {
			if p.ParamType == "functiontype" {
				return
			}
		}

		methodName := method.MethodName
		// Return type
		var returnType string
		if isOutParam {
			returnType = "emscripten::val"
		} else {
			returnType = "void"
			for _, p := range method.Params {
				if p.ParamPass == "return" {
					if p.ParamType == "struct" {
						returnType = fmt.Sprintf("s%sWrapper", p.ParamClass)
					} else if p.ParamType == "class" || p.ParamType == "optionalclass" {
						returnType = fmt.Sprintf("P%s", p.ParamClass)
					} else if p.ParamType == "basicarray" {
						returnType = fmt.Sprintf("std::vector<%s_%s>", component.LibraryName, p.ParamClass)
					} else {
						returnType = ResolveCppType(p.ParamType, component)
					}
					break
				}
			}
		}

		// Function name and signature
		result.WriteString("static ")
		result.WriteString(returnType)
		result.WriteString(" wrap_")
		result.WriteString(className)
		result.WriteString("_")
		result.WriteString(methodName)
		result.WriteString("(")
		// For class: CClass &self
		result.WriteString(fmt.Sprintf("C%s &self", className))
		// Params (excluding return and (if out) out-params)
		var paramList []string
		for _, p := range method.Params {
			if p.ParamPass == "return" || (isOutParam && p.ParamPass == "out") {
				continue
			}
			switch p.ParamType {
			case "struct":
				paramList = append(paramList, fmt.Sprintf("const s%sWrapper& %s", p.ParamClass, p.ParamName))
			case "structarray":
				paramList = append(paramList, fmt.Sprintf("const std::vector<s%sWrapper>& %s", p.ParamClass, p.ParamName))
			case "basicarray":
				paramList = append(paramList, fmt.Sprintf("std::vector<%s_%s>& %s", component.LibraryName, p.ParamClass, p.ParamName))
			case "enum":
				paramList = append(paramList, fmt.Sprintf("const e%s& %s", p.ParamClass, p.ParamName))
			case "class", "optionalclass":
				paramList = append(paramList, fmt.Sprintf("P%s& %s", p.ParamClass, p.ParamName))
			default:
				paramList = append(paramList, fmt.Sprintf("const %s& %s", ResolveCppType(p.ParamType, component), p.ParamName))
			}
		}
		// Append param list (comma logic)
		if len(paramList) > 0 {
			result.WriteString(", ")
			result.WriteString(strings.Join(paramList, ", "))
		}
		result.WriteString(") {\n")

		// structarray conversion
		if !isOutParam {
			for _, p := range method.Params {
				if p.ParamType == "structarray" {
					result.WriteString(fmt.Sprintf("    std::vector<s%s> converted_%s;\n", p.ParamClass, p.ParamName))
					result.WriteString(fmt.Sprintf("    converted_%s.reserve(%s.size());\n", p.ParamName, p.ParamName))
					result.WriteString(fmt.Sprintf("    for (const auto& w : %s) converted_%s.push_back(w.toStruct());\n", p.ParamName, p.ParamName))
				}
			}
		}

		if isOutParam {
			result.WriteString("    emscripten::val output = emscripten::val::object();\n")
			// Out param declarations
			for _, p := range method.Params {
				if p.ParamPass == "out" {
					paramType := ""
					if p.ParamType == "struct" {
						paramType = fmt.Sprintf("s%sWrapper", p.ParamClass)
					} else if p.ParamType == "structarray" {
						paramType = fmt.Sprintf("std::vector<s%s>", p.ParamClass)
					} else if p.ParamType == "basicarray" {
						paramType = fmt.Sprintf("std::vector<%s_%s>", component.LibraryName, p.ParamClass)
					} else if p.ParamType == "enum" {
						paramType = fmt.Sprintf("e%s", p.ParamClass)
					} else if p.ParamType == "class" {
						paramType = fmt.Sprintf("P%s", p.ParamClass)
					} else if p.ParamType == "optionalclass" {
						paramType = fmt.Sprintf("P%s", p.ParamClass)
					} else {
						paramType = ResolveCppType(p.ParamType, component)
					}
					result.WriteString(fmt.Sprintf("    %s %s;\n", paramType, p.ParamName))
				}
			}
		}

		// Compose call args
		var callArgs []string
		for _, p := range method.Params {
			if p.ParamPass == "return" {
				continue
			}
			if isOutParam && p.ParamPass == "out" {
				if p.ParamType == "struct" {
					callArgs = append(callArgs, fmt.Sprintf("%s.value", p.ParamName))
				} else {
					callArgs = append(callArgs, p.ParamName)
				}
			} else {
				switch p.ParamType {
				case "struct":
					callArgs = append(callArgs, fmt.Sprintf("%s.toStruct()", p.ParamName))
				case "structarray":
					callArgs = append(callArgs, fmt.Sprintf("converted_%s", p.ParamName))
				case "class", "optionalclass":
					callArgs = append(callArgs,
						fmt.Sprintf("classParam(%s)", p.ParamName),
					)
				default:
					callArgs = append(callArgs, p.ParamName)
				}
			}
		}

		// Call expr (class/global)
		var callExpr string
		callExpr = fmt.Sprintf("self.%s(%s)", methodName, strings.Join(callArgs, ", "))

		// Output assignment
		if isOutParam {
			var returnVar string
			for _, p := range method.Params {
				if p.ParamPass == "return" {
					returnVar = "return_value"
					break
				}
			}
			if returnVar != "" {
				returnType := "std::string"
				for _, p := range method.Params {
					if p.ParamPass == "return" && p.ParamType != "string" {
						returnType = ResolveCppType(p.ParamType, component)
						break
					}
				}
				result.WriteString(fmt.Sprintf("    %s %s = %s;\n", returnType, returnVar, callExpr))
				result.WriteString(fmt.Sprintf("    output.set(\"return\", %s);\n", returnVar))
			} else {
				result.WriteString(fmt.Sprintf("    %s;\n", callExpr))
			}
			for _, p := range method.Params {
				if p.ParamPass == "out" {
					result.WriteString(fmt.Sprintf("    output.set(\"%s\", %s);\n", p.ParamName, p.ParamName))
				}
			}
			result.WriteString("    return output;\n")
		} else {
			retType := "void"
			for _, p := range method.Params {
				if p.ParamPass == "return" {
					if p.ParamType == "struct" {
						retType = fmt.Sprintf("s%sWrapper", p.ParamClass)
					} else {
						retType = ResolveCppType(p.ParamType, component)
					}
					break
				}
			}
			if retType != "void" {
				result.WriteString(fmt.Sprintf("    auto result = %s;\n", callExpr))
				if strings.HasPrefix(retType, "s") {
					result.WriteString(fmt.Sprintf("    %s wrapper;\n", retType))
					result.WriteString("    wrapper.value = result;\n")
					result.WriteString("    return wrapper;\n")
				} else {
					result.WriteString("    return result;\n")
				}
			} else {
				result.WriteString(fmt.Sprintf("    %s;\n", callExpr))
			}
		}
		result.WriteString("}\n\n")
	}

	// Instance methods
	for _, class := range component.Classes {
		for _, method := range class.Methods {
			emitOne(method, class.ClassName, false)
		}
	}
	// Global methods
	for _, method := range component.Global.Methods {
		emitOne(method, "Wrapper", true)
	}
	return result.String()
}

// Call to generate static method wrappers
func GenerateStaticMethodWrappers(component ComponentDefinition) string {
	return generateMethodWrappers(component, false)
}

// Call to generate out param method wrappers
func GenerateOutParamMethodWrappers(component ComponentDefinition) string {
	return generateMethodWrappers(component, true)
}

// Actual binding block
func GenerateEmscriptenBindings(component ComponentDefinition) string {
	var result strings.Builder

	result.WriteString("// ================== Emscripten Bindings ==================\n")
	result.WriteString("EMSCRIPTEN_BINDINGS(" + component.LibraryName + ") {\n")

	if len(component.Enums) > 0 {
		result.WriteString("    // Enums\n")
		for _, e := range component.Enums {
			result.WriteString(fmt.Sprintf("    enum_<e%s>(\"e%s\")\n", e.Name, e.Name))
			for _, opt := range e.Options {
				result.WriteString(fmt.Sprintf("        .value(\"%s\", e%s::%s)\n", opt.Name, e.Name, opt.Name))
			}
			result.WriteString("    ;\n")
		}
		result.WriteString("\n")
	}

	if len(component.Structs) > 0 {
		result.WriteString("    // Register JS bindings for struct-array wrappers\n")
		for _, s := range component.Structs {
			result.WriteString(fmt.Sprintf("    register_vector<s%sWrapper>(\"std::vector<s%s>\");\n", s.Name, s.Name))
		}
		result.WriteString("\n")

		result.WriteString("    // Structs as exposed JS classes\n")
		for _, s := range component.Structs {
			result.WriteString(fmt.Sprintf("    class_<s%sWrapper>(\"s%s\")\n", s.Name, s.Name))
			result.WriteString("        .constructor<>()\n")
			result.WriteString(fmt.Sprintf("        .class_function(\"fromStruct\", &s%sWrapper::fromStruct)\n", s.Name))

			for _, m := range s.Members {
				if m.Rows > 0 || m.Columns > 0 {
					if m.Rows > 0 && m.Columns > 0 {
						for i := 0; i < m.Columns; i++ {
							for j := 0; j < m.Rows; j++ {
								result.WriteString(fmt.Sprintf("        .function(\"get_%s_%d_%d\", &s%sWrapper::get_%s_%d_%d)\n", m.Name, i, j, s.Name, m.Name, i, j))
								result.WriteString(fmt.Sprintf("        .function(\"set_%s_%d_%d\", &s%sWrapper::set_%s_%d_%d)\n", m.Name, i, j, s.Name, m.Name, i, j))
							}
						}
					} else {
						size := m.Rows
						if m.Rows == 0 {
							size = m.Columns
						}
						for i := 0; i < size; i++ {
							result.WriteString(fmt.Sprintf("        .function(\"get_%s%d\", &s%sWrapper::get_%s%d)\n", m.Name, i, s.Name, m.Name, i))
							result.WriteString(fmt.Sprintf("        .function(\"set_%s%d\", &s%sWrapper::set_%s%d)\n", m.Name, i, s.Name, m.Name, i))
						}
					}
				} else {
					result.WriteString(fmt.Sprintf("        .function(\"get_%s\", &s%sWrapper::get_%s)\n", m.Name, s.Name, m.Name))
					result.WriteString(fmt.Sprintf("        .function(\"set_%s\", &s%sWrapper::set_%s)\n", m.Name, s.Name, m.Name))
				}
			}
			result.WriteString("    ;\n")
		}
		result.WriteString("\n")
	}

	result.WriteString("    // Binding Methods\n")
	for _, cls := range component.Classes {
		base := fmt.Sprintf("C%s", cls.ClassName)
		if cls.ParentClass != "" {
			result.WriteString(fmt.Sprintf("    class_<%s, base<C%s>>(\"%s\")\n", base, cls.ParentClass, base))
		} else {
			result.WriteString(fmt.Sprintf("    class_<%s>(\"%s\")\n", base, base))
		}
		result.WriteString(fmt.Sprintf("        .smart_ptr<std::shared_ptr<%s>>(\"shared_ptr<%s>\")\n", base, base))

		for _, method := range cls.Methods {
			skip := false
			for _, p := range method.Params {
				if p.ParamType == "functiontype" {
					skip = true
					break
				}
			}
			if skip {
				result.WriteString(fmt.Sprintf("        // .function(\"%s\", &wrap_%s_%s) // Skipped due to callback\n", method.MethodName, cls.ClassName, method.MethodName))
				continue
			}

			wrapperName := fmt.Sprintf("wrap_%s_%s", cls.ClassName, method.MethodName)
			if NeedsStaticWrapper(method) || NeedsOutParamWrapper(method) {
				result.WriteString(fmt.Sprintf("        .function(\"%s\", &%s)\n", method.MethodName, wrapperName))
			} else {
				result.WriteString(fmt.Sprintf("        .function(\"%s\", &C%s::%s)\n", method.MethodName, cls.ClassName, method.MethodName))
			}
		}
		result.WriteString("    ;\n")
	}

	result.WriteString("    // CWrapper global bindings\n")
	result.WriteString("    class_<CWrapper>(\"CWrapper\")\n")
	result.WriteString("        .constructor<>()\n")
	for _, method := range component.Global.Methods {
		if shouldSkipMethod(method) {
			result.WriteString(fmt.Sprintf("        // .function(\"%s\", &CWrapper::%s) // Explicitly skipped (Returns a void pointer)\n", method.MethodName, method.MethodName))
			continue
		}

		skipCallback := false
		for _, p := range method.Params {
			if p.ParamType == "functiontype" {
				skipCallback = true
				break
			}
		}
		if skipCallback {
			result.WriteString(fmt.Sprintf("        // .function(\"%s\", &wrap_Wrapper_%s) // Skipped due to callback\n", method.MethodName, method.MethodName))
			continue
		}

		wrapperName := fmt.Sprintf("wrap_Wrapper_%s", method.MethodName)
		if NeedsStaticWrapper(method) || NeedsOutParamWrapper(method) {
			result.WriteString(fmt.Sprintf("        .function(\"%s\", &%s)\n", method.MethodName, wrapperName))
		} else {
			result.WriteString(fmt.Sprintf("        .function(\"%s\", &CWrapper::%s)\n", method.MethodName, method.MethodName))
		}
	}
	result.WriteString("    ;\n")

	result.WriteString("}\n")
	return result.String()
}

// Assembly the binding to a single Cpp
func BuildWASMBinding(componentdefinition ComponentDefinition, outputFolder string, outputFolderExample string, indentString string) error {
	// Step 1: Generate the header block
	header := GenerateCppHeader(componentdefinition)

	// Step 2: Generate the struct wrapper section
	structWrappers := GenerateStructWrappers(componentdefinition)

	// Step 3: Static Method Wrappers (without out-params)
	staticWrappers := GenerateStaticMethodWrappers(componentdefinition)

	// Step 4: Static Method Wrappers for out-params
	outParamWrappers := GenerateOutParamMethodWrappers(componentdefinition)

	// Step 5: Generate the Emscripten bindings block (currently commented out)
	bindingsBlock := GenerateEmscriptenBindings(componentdefinition)

	// Determine the output filename based on the library's base name
	outputFileName := filepath.Join(outputFolder, fmt.Sprintf("%s_bindings.cpp", componentdefinition.BaseName))

	// Create the file
	file, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outputFileName, err)
	}
	defer file.Close()

	// Write license header
	writeLicenseHeaderEx(file, componentdefinition, "C++ Emscripten wrapper for WebAssembly", true, "/*", "*/")

	// Write the header
	if _, err := file.WriteString(header); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write the struct wrappers
	if _, err := file.WriteString(structWrappers); err != nil {
		return fmt.Errorf("failed to write struct wrappers: %w", err)
	}

	// Write the static wrappers (normal)
	if _, err := file.WriteString(staticWrappers); err != nil {
		return fmt.Errorf("failed to write static wrappers: %w", err)
	}

	// Write the static wrappers (out-param methods)
	if _, err := file.WriteString(outParamWrappers); err != nil {
		return fmt.Errorf("failed to write out-param static wrappers: %w", err)
	}

	// Write the Emscripten bindings block (disabled for now)
	if _, err := file.WriteString(bindingsBlock); err != nil {
		return fmt.Errorf("failed to write Emscripten bindings block: %w", err)
	}

	// Done
	fmt.Printf("Successfully generated WASM bindings with name : %s\n", outputFileName)
	return nil
}
