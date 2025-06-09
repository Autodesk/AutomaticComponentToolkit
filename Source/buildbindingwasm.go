//////////////////////////////////////////////////////////////////////////////////////////////////////
// buildbindingwasm.go
// functions to generate dynamic Python3-bindings of a library's API in form of explicitly loaded
// function handles.
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Store the python file path
var wasmBindingFile = ""

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
				typeString = fmt.Sprintf("Lib3MF_%s", m.Type)
			}

			// Determine if the member is an array
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
				typeString = fmt.Sprintf("Lib3MF_%s", m.Type)
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

// GenerateCppHeader generates the header part of the C++ file as a string
func GenerateCppHeader(componentdefinition ComponentDefinition) string {
	var builder strings.Builder

	// Write the standard includes
	builder.WriteString("#include <iostream>\n")
	builder.WriteString("#include <vector>\n")
	builder.WriteString("#include <string>\n")
	builder.WriteString("#include <emscripten.h>\n")
	builder.WriteString("#include <emscripten/bind.h>\n")

	// Use the BaseName to generate the #include for the implicit header
	builder.WriteString(fmt.Sprintf("#include \"Cpp/%s_implicit.hpp\"\n", componentdefinition.BaseName))

	// Namespace usage
	builder.WriteString("\n")
	builder.WriteString("using namespace emscripten;\n")
	builder.WriteString(fmt.Sprintf("using namespace %s;\n\n", componentdefinition.NameSpace))

	return builder.String()
}

func GenerateStaticMethodWrappers(component ComponentDefinition) string {
	var result strings.Builder

	for _, class := range component.Classes {
		className := class.ClassName
		for _, method := range class.Methods {
			// Skip methods with out-params
			hasOutParam := false
			for _, param := range method.Params {
				if param.ParamPass == "out" {
					hasOutParam = true
					break
				}
			}
			if hasOutParam {
				continue
			}

			methodName := method.MethodName
			result.WriteString("static ")

			// Determine return type
			returnType := "void"
			for _, p := range method.Params {
				if p.ParamPass == "return" {
					if p.ParamType == "struct" {
						returnType = fmt.Sprintf("s%sWrapper", p.ParamClass)
					} else {
						returnType = fmt.Sprintf("Lib3MF_%s", p.ParamType)
					}
					break
				}
			}

			// Function signature
			result.WriteString(fmt.Sprintf("%s wrap_%s_%s(C%s &self", returnType, className, methodName, className))

			// Parameter list
			var paramList []string
			for _, p := range method.Params {
				if p.ParamPass == "return" {
					continue
				}
				var param string
				switch p.ParamType {
				case "struct":
					param = fmt.Sprintf("const s%sWrapper& %s", p.ParamClass, p.ParamName)
				case "structarray":
					param = fmt.Sprintf("const std::vector<s%sWrapper>& %s", p.ParamClass, p.ParamName)
				default:
					param = fmt.Sprintf("Lib3MF_%s %s", p.ParamType, p.ParamName)
				}
				paramList = append(paramList, param)
			}
			if len(paramList) > 0 {
				result.WriteString(", " + strings.Join(paramList, ", "))
			}
			result.WriteString(") {\n")

			// structarray conversion
			for _, p := range method.Params {
				if p.ParamType == "structarray" {
					result.WriteString(fmt.Sprintf("    std::vector<s%s> converted_%s;\n", p.ParamClass, p.ParamName))
					result.WriteString(fmt.Sprintf("    converted_%s.reserve(%s.size());\n", p.ParamName, p.ParamName))
					result.WriteString(fmt.Sprintf("    for (const auto& w : %s) converted_%s.push_back(w.toStruct());\n", p.ParamName, p.ParamName))
				}
			}

			// Compose call
			var callArgs []string
			for _, p := range method.Params {
				if p.ParamPass == "return" {
					continue
				}
				switch p.ParamType {
				case "struct":
					callArgs = append(callArgs, fmt.Sprintf("%s.toStruct()", p.ParamName))
				case "structarray":
					callArgs = append(callArgs, fmt.Sprintf("converted_%s", p.ParamName))
				default:
					callArgs = append(callArgs, p.ParamName)
				}
			}
			callExpr := fmt.Sprintf("self.%s(%s)", methodName, strings.Join(callArgs, ", "))

			// Return or void
			if returnType != "void" {
				result.WriteString(fmt.Sprintf("    auto result = %s;\n", callExpr))
				if strings.HasPrefix(returnType, "s") { // wrapper
					result.WriteString(fmt.Sprintf("    %s wrapper;\n", returnType))
					result.WriteString("    wrapper.value = result;\n")
					result.WriteString("    return wrapper;\n")
				} else {
					result.WriteString("    return result;\n")
				}
			} else {
				result.WriteString(fmt.Sprintf("    %s;\n", callExpr))
			}
			result.WriteString("}\n\n")
		}
	}

	return result.String()
}

func GenerateOutParamMethodWrappers(component ComponentDefinition) string {
	var result strings.Builder

	for _, class := range component.Classes {
		className := class.ClassName
		for _, method := range class.Methods {
			// Check if method has at least one out-param
			hasOutParam := false
			for _, param := range method.Params {
				if param.ParamPass == "out" {
					hasOutParam = true
					break
				}
			}
			if !hasOutParam {
				continue
			}

			methodName := method.MethodName
			result.WriteString("static emscripten::val ")
			result.WriteString(fmt.Sprintf("wrap_%s_%s(C%s &self", className, methodName, className))

			// Parameter list (no out params)
			var paramList []string
			for _, p := range method.Params {
				if p.ParamPass == "return" || p.ParamPass == "out" {
					continue
				}
				var param string
				switch p.ParamType {
				case "struct":
					param = fmt.Sprintf("const s%sWrapper& %s", p.ParamClass, p.ParamName)
				default:
					param = fmt.Sprintf("Lib3MF_%s %s", p.ParamType, p.ParamName)
				}
				paramList = append(paramList, param)
			}
			if len(paramList) > 0 {
				result.WriteString(", " + strings.Join(paramList, ", "))
			}
			result.WriteString(") {\n")

			// Create output object
			result.WriteString("    emscripten::val output = emscripten::val::object();\n")

			// Create local variables for out-params
			for _, p := range method.Params {
				if p.ParamPass == "out" {
					var paramType string
					if p.ParamType == "struct" {
						paramType = fmt.Sprintf("s%s", p.ParamClass)
					} else {
						paramType = fmt.Sprintf("Lib3MF_%s", p.ParamType)
					}
					result.WriteString(fmt.Sprintf("    %s %s;\n", paramType, p.ParamName))
				}
			}

			// Compose call arguments
			var callArgs []string
			for _, p := range method.Params {
				switch {
				case p.ParamPass == "return":
					continue
				case p.ParamPass == "out":
					if p.ParamType == "struct" {
						callArgs = append(callArgs, fmt.Sprintf("%s", p.ParamName+".value"))
					} else {
						callArgs = append(callArgs, p.ParamName)
					}
				case p.ParamType == "struct":
					callArgs = append(callArgs, fmt.Sprintf("%s.toStruct()", p.ParamName))
				default:
					callArgs = append(callArgs, p.ParamName)
				}
			}

			// Compose call statement
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
						returnType = fmt.Sprintf("Lib3MF_%s", p.ParamType)
						break
					}
				}
				result.WriteString(fmt.Sprintf("    %s %s = self.%s(%s);\n", returnType, returnVar, methodName, strings.Join(callArgs, ", ")))
				result.WriteString(fmt.Sprintf("    output.set(\"return\", %s);\n", returnVar))
			} else {
				result.WriteString(fmt.Sprintf("    self.%s(%s);\n", methodName, strings.Join(callArgs, ", ")))
			}

			// Populate output object
			for _, p := range method.Params {
				if p.ParamPass == "out" {
					result.WriteString(fmt.Sprintf("    output.set(\"%s\", %s);\n", p.ParamName, p.ParamName))
				}
			}

			result.WriteString("    return output;\n")
			result.WriteString("}\n\n")
		}
	}

	return result.String()
}

func GenerateEmscriptenBindings(component ComponentDefinition) string {
	var result strings.Builder

	result.WriteString("// ================== Emscripten Bindings ==================\n")
	result.WriteString("EMSCRIPTEN_BINDINGS(" + component.BaseName + ") {\n")

	// Register std::vector for struct wrappers
	result.WriteString("    // Register JS bindings for struct-array wrappers\n")
	for _, s := range component.Structs {
		result.WriteString(fmt.Sprintf("    register_vector<s%sWrapper>(\"std::vector<s%s>\");\n", s.Name, s.Name))
	}
	result.WriteString("\n")

	// Enums
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

	// Struct classes
	result.WriteString("    // Structs as exposed JS classes\n")
	for _, s := range component.Structs {
		result.WriteString(fmt.Sprintf("    class_<s%sWrapper>(\"s%s\")\n", s.Name, s.Name))
		result.WriteString("        .constructor<>()\n")
		result.WriteString(fmt.Sprintf("        .class_function(\"fromStruct\", &s%sWrapper::fromStruct)\n", s.Name))

		for _, m := range s.Members {
			if m.Rows > 0 || m.Columns > 0 { // array case
				if m.Rows > 0 && m.Columns > 0 { // 2D
					for i := 0; i < m.Columns; i++ {
						for j := 0; j < m.Rows; j++ {
							result.WriteString(fmt.Sprintf("        .function(\"get_%s_%d_%d\", &s%sWrapper::get_%s_%d_%d)\n", m.Name, i, j, s.Name, m.Name, i, j))
							result.WriteString(fmt.Sprintf("        .function(\"set_%s_%d_%d\", &s%sWrapper::set_%s_%d_%d)\n", m.Name, i, j, s.Name, m.Name, i, j))
						}
					}
				} else { // 1D
					size := m.Rows
					if m.Rows == 0 {
						size = m.Columns
					}
					for i := 0; i < size; i++ {
						result.WriteString(fmt.Sprintf("        .function(\"get_%s%d\", &s%sWrapper::get_%s%d)\n", m.Name, i, s.Name, m.Name, i))
						result.WriteString(fmt.Sprintf("        .function(\"set_%s%d\", &s%sWrapper::set_%s%d)\n", m.Name, i, s.Name, m.Name, i))
					}
				}
			} else { // single value
				result.WriteString(fmt.Sprintf("        .function(\"get_%s\", &s%sWrapper::get_%s)\n", m.Name, s.Name, m.Name))
				result.WriteString(fmt.Sprintf("        .function(\"set_%s\", &s%sWrapper::set_%s)\n", m.Name, s.Name, m.Name))
			}
		}
		result.WriteString("    ;\n")
	}
	result.WriteString("\n")

	// Class bindings
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
			// Here we assume you'll use the generated static wrappers if out-params, or direct bindings
			// For now, we’ll default to direct C++ methods; refine further if needed
			result.WriteString(fmt.Sprintf("        .function(\"%s\", &C%s::%s)\n", method.MethodName, cls.ClassName, method.MethodName))
		}
		result.WriteString("    ;\n")
	}

	// Global methods in CWrapper
	result.WriteString("    // CWrapper global bindings\n")
	result.WriteString("    class_<CWrapper>(\"CWrapper\")\n")
	result.WriteString("        .constructor<>()\n")
	for _, method := range component.Global.Methods {
		result.WriteString(fmt.Sprintf("        .function(\"%s\", &CWrapper::%s)\n", method.MethodName, method.MethodName))
	}
	result.WriteString("    ;\n")

	result.WriteString("}\n")
	return result.String()
}

// BuildWASMBinding generates the WASM C++ binding file.
func BuildWASMBinding(componentdefinition ComponentDefinition, outputFolder string, outputFolderExample string, indentString string) error {
	// Step 1: Generate the header block
	header := GenerateCppHeader(componentdefinition)

	// Step 2: Generate the struct wrapper section
	structWrappers := GenerateStructWrappers(componentdefinition)

	// Step 3: Static Method Wrappers (without out-params)
	staticWrappers := GenerateStaticMethodWrappers(componentdefinition)

	// Step 4: Static Method Wrappers for out-params
	outParamWrappers := GenerateOutParamMethodWrappers(componentdefinition)

	// Step 5: Generate the Emscripten bindings block
	bindingsBlock := GenerateEmscriptenBindings(componentdefinition)

	// Determine the output filename based on the library's base name
	outputFileName := filepath.Join(outputFolder, fmt.Sprintf("%s_bindings.cpp", componentdefinition.BaseName))

	// Create the file
	file, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outputFileName, err)
	}
	defer file.Close()

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

	// Write the Emscripten bindings block
	if _, err := file.WriteString(bindingsBlock); err != nil {
		return fmt.Errorf("failed to write Emscripten bindings block: %w", err)
	}

	// Done
	fmt.Printf("Successfully generated %s\n", outputFileName)
	return nil
}
