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
// automaticcomponenttoolkit.go
// A toolkit to automatically generate software components: abstract API, implementation stubs and language bindings
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path"
)

const (
	eACTModeGenerate = 0
	eACTModeDiff     = 1
)

func createComponent(component ComponentDefinition, outfolderBase string) error {

	log.Printf("Creating Component \"%s\"", component.LibraryName)
	for _, subComponent := range component.ImportedComponentDefinitions {
		err := createComponent(subComponent, outfolderBase)
		if err != nil {
			return err
		}
	}

	outputFolder := path.Join(outfolderBase, component.NameSpace+"_component")
	outputFolderBindings := path.Join(outputFolder, "Bindings")
	outputFolderExamples := path.Join(outputFolder, "Examples")
	outputFolderImplementations := path.Join(outputFolder, "Implementations")

	err := os.MkdirAll(outputFolder, os.ModePerm)
	if err != nil {
		return err
	}

	licenseFileName := path.Join(outputFolder, "license.txt")
	log.Printf("Creating \"%s\"", licenseFileName)
	licenseFile, err := CreateLanguageFile(licenseFileName, "")
	if err != nil {
		return err
	}
	licenseFile.WritePlainLicenseHeader(component, "", false)

	if len(component.BindingList.Bindings) > 0 {
		err = os.MkdirAll(outputFolderBindings, os.ModePerm)
		if err != nil {
			return err
		}
	}
	for bindingindex := 0; bindingindex < len(component.BindingList.Bindings); bindingindex++ {
		binding := component.BindingList.Bindings[bindingindex]
		indentString := getIndentationString(binding.Indentation)
		log.Printf("Exporting Interface Binding for Languge \"%s\"", binding.Language)

		switch binding.Language {
		case "C":
			{
				outputFolderBindingC := outputFolderBindings + "/C"

				err = os.MkdirAll(outputFolderBindingC, os.ModePerm)
				if err != nil {
					return err
				}

				err = BuildBindingC(component, outputFolderBindingC)
				if err != nil {
					return err
				}
			}

		case "CDynamic":
			{
				outputFolderBindingCDynamic := outputFolderBindings + "/CDynamic"
				err = os.MkdirAll(outputFolderBindingCDynamic, os.ModePerm)
				if err != nil {
					return err
				}
				outputFolderExampleCDynamic := outputFolderExamples + "/CDynamic"
				err = os.MkdirAll(outputFolderExampleCDynamic, os.ModePerm)
				if err != nil {
					return err
				}

				CTypesHeaderName := path.Join(outputFolderBindingCDynamic, component.BaseName+"_types.h")
				err = CreateCTypesHeader(component, CTypesHeaderName)
				if err != nil {
					return err
				}

				err = BuildBindingCExplicit(component, outputFolderBindingCDynamic, outputFolderExampleCDynamic, indentString)
				if err != nil {
					return err
				}
			}

		case "CppDynamic":
			{
				outputFolderBindingCppDynamic := outputFolderBindings + "/CppDynamic"
				err = os.MkdirAll(outputFolderBindingCppDynamic, os.ModePerm)
				if err != nil {
					return err
				}
				outputFolderExampleCppDynamic := outputFolderExamples + "/CppDynamic"
				err = os.MkdirAll(outputFolderExampleCppDynamic, os.ModePerm)
				if err != nil {
					return err
				}

				CPPTypesHeaderName := path.Join(outputFolderBindingCppDynamic, component.BaseName+"_types.hpp")
				err = CreateCPPTypesHeader(component, CPPTypesHeaderName)
				if err != nil {
					return err
				}

				CPPABIHeaderName := path.Join(outputFolderBindingCppDynamic, component.BaseName+"_abi.hpp")
				err = CreateCPPAbiHeader(component, CPPABIHeaderName)
				if err != nil {
					return err
				}

				err = BuildBindingCppExplicit(component, outputFolderBindingCppDynamic, outputFolderExampleCppDynamic,
					indentString, binding.ClassIdentifier)
				if err != nil {
					return err
				}
			}

		case "Cpp":
			{
				outputFolderBindingCppImplicit := outputFolderBindings + "/Cpp"
				err = os.MkdirAll(outputFolderBindingCppImplicit, os.ModePerm)
				if err != nil {
					return err
				}
				outputFolderExampleCppImplicit := outputFolderExamples + "/Cpp"
				err = os.MkdirAll(outputFolderExampleCppImplicit, os.ModePerm)
				if err != nil {
					return err
				}

				CPPTypesHeaderName := path.Join(outputFolderBindingCppImplicit, component.BaseName+"_types.hpp")
				err = CreateCPPTypesHeader(component, CPPTypesHeaderName)
				if err != nil {
					return err
				}

				CPPABIHeaderName := path.Join(outputFolderBindingCppImplicit, component.BaseName+"_abi.hpp")
				err = CreateCPPAbiHeader(component, CPPABIHeaderName)
				if err != nil {
					return err
				}

				err = BuildBindingCppImplicit(component, outputFolderBindingCppImplicit, outputFolderExampleCppImplicit,
					indentString, binding.ClassIdentifier)
				if err != nil {
					return err
				}
			}

		case "Go":
			{
				outputFolderBindingGo := outputFolderBindings + "/Go"
				err = os.MkdirAll(outputFolderBindingGo, os.ModePerm)
				if err != nil {
					return err
				}

				outputFolderExampleGo := outputFolderExamples + "/Go"
				err = os.MkdirAll(outputFolderExampleGo, os.ModePerm)
				if err != nil {
					return err
				}

				err := BuildBindingGo(component, outputFolderBindingGo, outputFolderExampleGo)
				if err != nil {
					return err
				}
			}

		case "Node":
			{
				outputFolderBindingNode := outputFolderBindings + "/NodeJS"

				err = os.MkdirAll(outputFolderBindingNode, os.ModePerm)
				if err != nil {
					return err
				}

				CTypesHeaderName := path.Join(outputFolderBindingNode, component.BaseName+"_types.h")
				err = CreateCTypesHeader(component, CTypesHeaderName)
				if err != nil {
					return err
				}

				err = BuildBindingCExplicit(component, outputFolderBindingNode, "", indentString)
				if err != nil {
					return err
				}

				err := BuildBindingNode(component, outputFolderBindingNode, indentString)
				if err != nil {
					return err
				}
			}

		case "Pascal":
			{
				outputFolderBindingPascal := outputFolderBindings + "/Pascal"
				err = os.MkdirAll(outputFolderBindingPascal, os.ModePerm)
				if err != nil {
					return err
				}

				outputFolderExamplePascal := outputFolderExamples + "/Pascal"
				err = os.MkdirAll(outputFolderExamplePascal, os.ModePerm)
				if err != nil {
					return err
				}

				err = BuildBindingPascalDynamic(component, outputFolderBindingPascal, outputFolderExamplePascal, indentString)
				if err != nil {
					return err
				}
			}

		case "CSharp":
			{
				outputFolderBindingCSharp := outputFolderBindings + "/CSharp"
				err = os.MkdirAll(outputFolderBindingCSharp, os.ModePerm)
				if err != nil {
					log.Fatal(err)
				}

				outputFolderExampleCSharp := outputFolderExamples + "/CSharp"
				err = os.MkdirAll(outputFolderExampleCSharp, os.ModePerm)
				if err != nil {
					log.Fatal(err)
				}

				err = BuildBindingCSharp(component, outputFolderBindingCSharp, outputFolderExampleCSharp, indentString)
				if err != nil {
					log.Fatal(err)
				}
			}
		case "Python":
			{
				outputFolderBindingPython := outputFolderBindings + "/Python"
				err = os.MkdirAll(outputFolderBindingPython, os.ModePerm)
				if err != nil {
					return err
				}

				outputFolderExamplePython := outputFolderExamples + "/Python"
				err = os.MkdirAll(outputFolderExamplePython, os.ModePerm)
				if err != nil {
					return err
				}

				err = BuildBindingPythonDynamic(component, outputFolderBindingPython, outputFolderExamplePython, indentString)
				if err != nil {
					return err
				}
			}

		case "Fortran":
			{
				log.Printf("Interface binding for language \"%s\" is not yet supported.", binding.Language)
			}

		default:
			log.Fatal("Unknown binding export")
		}
	}

	if len(component.ImplementationList.Implementations) > 0 {
		err = os.MkdirAll(outputFolderImplementations, os.ModePerm)
		if err != nil {
			return err
		}
	}
	for implementationindex := 0; implementationindex < len(component.ImplementationList.Implementations); implementationindex++ {
		implementation := component.ImplementationList.Implementations[implementationindex]
		log.Printf("Exporting Implementation Interface for Language \"%s\"", implementation.Language)

		switch implementation.Language {
		case "Cpp":
			{
				outputFolderImplementationProject := outputFolderImplementations + "/Cpp"
				outputFolderImplementationCpp := outputFolderImplementations + "/Cpp/Interfaces"
				outputFolderImplementationCppStub := outputFolderImplementations + "/Cpp/Stub"

				err = os.MkdirAll(outputFolderImplementationCpp, os.ModePerm)
				if err != nil {
					return err
				}

				err = os.MkdirAll(outputFolderImplementationCppStub, os.ModePerm)
				if err != nil {
					return err
				}

				CTypesHeaderName := path.Join(outputFolderImplementationCpp, component.BaseName+"_types.hpp")
				err = CreateCPPTypesHeader(component, CTypesHeaderName)
				if err != nil {
					return err
				}

				CHeaderName := path.Join(outputFolderImplementationCpp, component.BaseName+"_abi.hpp")
				err = CreateCPPAbiHeader(component, CHeaderName)
				if err != nil {
					return err
				}

				err = BuildImplementationCPP(component, outputFolderImplementationCpp, outputFolderImplementationCppStub,
					outputFolderImplementationProject, implementation)
				if err != nil {
					return err
				}
			}

		case "Pascal":
			{
				outputFolderImplementationProject := outputFolderImplementations + "/Pascal"
				outputFolderImplementationPascal := outputFolderImplementations + "/Pascal/Interfaces"
				outputFolderImplementationPascalStub := outputFolderImplementations + "/Pascal/Stub"

				err = os.MkdirAll(outputFolderImplementationPascal, os.ModePerm)
				if err != nil {
					return err
				}

				err = os.MkdirAll(outputFolderImplementationPascalStub, os.ModePerm)
				if err != nil {
					return err
				}

				err = BuildImplementationPascal(component, outputFolderImplementationPascal, outputFolderImplementationPascalStub,
					outputFolderImplementationProject, implementation)
				if err != nil {
					return err
				}
			}

		case "Fortran":
			{
				log.Printf("Implementation in language \"%s\" is not yet supported.", implementation.Language)
			}
		default:
			log.Fatal("Unknown export")
		}
	}

	return nil
}

func main() {
	ACTVersion := "1.6.0"
	fmt.Fprintln(os.Stdout, "Automatic Component Toolkit v"+ACTVersion)
	if len(os.Args) < 2 {
		log.Fatal("Please run with the Interface Description XML as command line parameter.")
		log.Fatal("To specify a path for the generated source code use the optional flag \"-o ABSOLUTE_PATH_TO_OUTPUT_FOLDER\"")
		log.Fatal("To create a diff between two versions of an Interface Description XML use the optional flag \"-d OTHER_IDL_FILE\"")
	}
	if os.Args[1] == "-v" {
		fmt.Fprintln(os.Stdout, "Version: "+ACTVersion)
		return
	}
	log.Printf("---------------------------------------\n")

	mode := eACTModeGenerate
	outfolderBase, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	diffFile := ""
	if len(os.Args) >= 4 {
		if os.Args[2] == "-o" {
			outfolderBase = os.Args[3]
		}

		if os.Args[2] == "-d" {
			diffFile = os.Args[3]
			mode = eACTModeDiff
		}
	}
	if mode == eACTModeGenerate {
		log.Printf("Output directory: " + outfolderBase)
	}

	log.Printf("Loading Component Description File")
	component, err := ReadComponentDefinition(os.Args[1], ACTVersion)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Checking Component Description")
	err = component.CheckComponentDefinition()
	if err != nil {
		log.Fatal(err)
	}

	if mode == eACTModeDiff {
		log.Printf("Loading Component Description File to compare to")
		componentB, err := ReadComponentDefinition(diffFile, ACTVersion)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Checking Component Description B")
		err = componentB.CheckComponentDefinition()
		if err != nil {
			log.Fatal(err)
		}
		diff, err := DiffComponentDefinitions(component, componentB)
		if err != nil {
			log.Fatal(err)
		}

		output, err := xml.MarshalIndent(diff, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		writer, err := os.Create("diff.xml")
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(output)
		writer.Write(output)

		return
	}

	// This needs to go into a "preparation function"
	// baseClass, err := setupBaseClassDefinition(true)
	// if (err != nil) {
	// 	log.Fatal (err);
	// }
	// component.Classes = append([]ComponentDefinitionClass{baseClass}, component.Classes...)
	// for i := 0; i < len(component.Classes); i++ {
	// 	if (!component.Classes[i].isBaseClass()) {
	// 		if (component.Classes[i].ParentClass == "") {
	// 			component.Classes[i].ParentClass = "BaseClass";
	// 		}
	// 	}
	// }

	err = createComponent(component, outfolderBase)
	if err != nil {
		log.Println("Fatal error")
		log.Fatal(err)
	} else {
		log.Println("Success")
	}

}
