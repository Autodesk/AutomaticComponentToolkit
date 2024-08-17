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
	"strconv"
	"strings"
)

const (
	eACTModeGenerate = 0
	eACTModeDiff     = 1
)

func createComponent(component ComponentDefinition, outfolderBase string, bindingsDirectoryOverride string, interfacesDirectoryOverride string, stubDirectoryOverride string, suppressBindings bool, suppressStub bool, suppressInterfaces bool, suppressSubcomponents bool, suppressLicense bool, suppressExamples bool) (error) {

	log.Printf("Creating Component \"%s\"", component.LibraryName)
	
	if (!suppressSubcomponents) {	
		for _, subComponent := range component.ImportedComponentDefinitions {
			err := createComponent(subComponent, outfolderBase, "", "", "", suppressBindings, suppressStub, suppressInterfaces, suppressSubcomponents, suppressLicense, suppressExamples)
			if (err != nil) {
				return err
			}
		}
	}

	outputFolder := path.Join(outfolderBase, component.NameSpace+"_component")
	outputFolderBindings := path.Join(outputFolder, "Bindings")
	outputFolderExamples := path.Join(outputFolder, "Examples")
	outputFolderDocumentation := path.Join(outputFolder, "Documentations")
	outputFolderImplementations := path.Join(outputFolder, "Implementations")
	
	if bindingsDirectoryOverride != "" {
		outputFolderBindings = bindingsDirectoryOverride;
	}

	err := os.MkdirAll(outputFolder, os.ModePerm)
	if err != nil {
		return err
	}

	if (!suppressLicense) {
		licenseFileName := path.Join(outputFolder, "license.txt")
		log.Printf("Creating \"%s\"", licenseFileName)
		licenseFile, err := CreateLanguageFile(licenseFileName, "")
		if err != nil {
			return err
		}
		licenseFile.WritePlainLicenseHeader(component, "", false)
	} else {
		log.Printf("Suppressing license...")
	}
	
	if (!suppressBindings) {
		if len(component.BindingList.Bindings) > 0 {
			err = os.MkdirAll(outputFolderBindings, os.ModePerm)
			if err != nil {
				return err
			}
		}
		for bindingindex := 0; bindingindex < len(component.BindingList.Bindings); bindingindex++ {
			binding := component.BindingList.Bindings[bindingindex]
			indentString := getIndentationString(binding.Indentation)
			log.Printf("Exporting Interface Binding for Language \"%s\"", binding.Language)

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
					
					outputFolderExampleCDynamic := "";
					if (!suppressExamples) {
						outputFolderExampleCDynamic = outputFolderExamples + "/CDynamic"
						err = os.MkdirAll(outputFolderExampleCDynamic, os.ModePerm)
						if err != nil {
							return err
						}
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
				
					outputFolderDocumentationCppExplicit := "";
					
					if (binding.Documentation != "") {
					
						outputFolderDocumentationCppExplicit = outputFolderDocumentation + "/Cpp"
						err = os.MkdirAll(outputFolderDocumentationCppExplicit, os.ModePerm)
						if err != nil {
							log.Fatal(err)
						}
					}
				
					outputFolderBindingCppDynamic := outputFolderBindings + "/CppDynamic"
					err = os.MkdirAll(outputFolderBindingCppDynamic, os.ModePerm)
					if err != nil {
						return err
					}
					
					outputFolderExampleCppDynamic := "";
					if (!suppressExamples) {
						outputFolderExampleCppDynamic = outputFolderExamples + "/CppDynamic"
						err = os.MkdirAll(outputFolderExampleCppDynamic, os.ModePerm)
						if err != nil {
							return err
						}
					}

					CPPTypesHeaderName := path.Join(outputFolderBindingCppDynamic, component.BaseName+"_types.hpp")
					err = CreateCPPTypesHeader(component, CPPTypesHeaderName, false)
					if err != nil {
						return err
					}

					err = BuildBindingCppExplicit(component, outputFolderBindingCppDynamic, outputFolderExampleCppDynamic,
						outputFolderDocumentationCppExplicit, indentString, binding.ClassIdentifier)
					if err != nil {
						return err
					}
				}

			case "Cpp":
				{
					outputFolderDocumentationCppImplicit := "";
					
					if (binding.Documentation != "") {
						outputFolderDocumentationCppImplicit = outputFolderDocumentation + "/Cpp";
						err = os.MkdirAll(outputFolderDocumentationCppImplicit, os.ModePerm)
						if err != nil {
							log.Fatal(err)
						}
					}

					outputFolderBindingCppImplicit := outputFolderBindings + "/Cpp"
					err = os.MkdirAll(outputFolderBindingCppImplicit, os.ModePerm)
					if err != nil {
						return err
					}
					
					outputFolderExampleCppImplicit := "";
					if (!suppressExamples) {
						outputFolderExampleCppImplicit = outputFolderExamples + "/Cpp"
						err = os.MkdirAll(outputFolderExampleCppImplicit, os.ModePerm)
						if err != nil {
							return err
						}
					}

					CPPTypesHeaderName := path.Join(outputFolderBindingCppImplicit, component.BaseName+"_types.hpp")
					err = CreateCPPTypesHeader(component, CPPTypesHeaderName, false)
					if err != nil {
						return err
					}

					CPPABIHeaderName := path.Join(outputFolderBindingCppImplicit, component.BaseName+"_abi.hpp")
					err = CreateCPPAbiHeader(component, CPPABIHeaderName)
					if err != nil {
						return err
					}

					err = BuildBindingCppImplicit(component, outputFolderBindingCppImplicit, outputFolderExampleCppImplicit,
						outputFolderDocumentationCppImplicit, indentString, binding.ClassIdentifier)
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

					outputFolderExampleGo := "";
					if (!suppressExamples) {
						outputFolderExampleGo = outputFolderExamples + "/Go"
						err = os.MkdirAll(outputFolderExampleGo, os.ModePerm)
						if err != nil {
							return err
						}
					}

					err := BuildBindingGo(component, outputFolderBindingGo, outputFolderExampleGo, indentString)
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

					
					outputFolderExamplePascal := "";
					if (!suppressExamples) {
						outputFolderExamplePascal = outputFolderExamples + "/Pascal"
						err = os.MkdirAll(outputFolderExamplePascal, os.ModePerm)
						if err != nil {
							return err
						}
					}

					err = BuildBindingPascalDynamic(component, outputFolderBindingPascal, outputFolderExamplePascal, indentString)
					if err != nil {
						return err
					}
				}

			case "CSharp":
				{
					outputFolderBindingCSharp := outputFolderBindings + "/CSharp";
					err  = os.MkdirAll(outputFolderBindingCSharp, os.ModePerm);
					if (err != nil) {
						log.Fatal (err);
					}

					
					outputFolderExampleCSharp := "";
					if (!suppressExamples) {
						outputFolderExampleCSharp = outputFolderExamples + "/CSharp";
						err  = os.MkdirAll(outputFolderExampleCSharp, os.ModePerm);
						if (err != nil) {
							log.Fatal (err);
						}
					}
					
					err = BuildBindingCSharp(component, outputFolderBindingCSharp, outputFolderExampleCSharp, indentString);
					if (err != nil) {
						log.Fatal (err);
					}
				}

			case "Python":
				{
					outputFolderBindingPython := outputFolderBindings + "/Python"
					err = os.MkdirAll(outputFolderBindingPython, os.ModePerm)
					if err != nil {
						return err
					}

					outputFolderExamplePython := "";
					if (!suppressExamples) {
						outputFolderExamplePython = outputFolderExamples + "/Python"
						err = os.MkdirAll(outputFolderExamplePython, os.ModePerm)
						if err != nil {
							return err
						}
					}

					err = BuildBindingPythonDynamic(component, outputFolderBindingPython, outputFolderExamplePython, indentString)
					if err != nil {
						return err
					}
				}

			case "Java":
				{
					version := 9
					if len(binding.Version) > 0 {
						if (binding.Version == "8") || (binding.Version == "1.8") {
							version = 8
						} else if (binding.Version == "9") || (binding.Version == "1.9") {
							version = 9
						} else {
							log.Fatal("Unknown/Unsupported java binding version: " + binding.Version)
							log.Fatal("Supported java versions are 8 and 9")
						}
					}
					versionStr := strconv.Itoa(version)
					outputFolderBindingJava := outputFolderBindings + "/Java" + versionStr
					err = os.MkdirAll(outputFolderBindingJava, os.ModePerm)
					if err != nil {
						return err
					}

					outputFolderExampleJava := outputFolderExamples + "/Java" + versionStr
					err = os.MkdirAll(outputFolderExampleJava, os.ModePerm)
					if err != nil {
						return err
					}

					err = BuildBindingJavaDynamic(component, outputFolderBindingJava, outputFolderExampleJava, indentString, version)
					if err != nil {
						return err
					}
				}

			case "Cppwasmtime":
				{
					outputFolderExampleCppwasmtime := "";
					if (!suppressExamples) {
						outputFolderExampleCppwasmtime = outputFolderExamples + "/Cppwasmtime"
						err = os.MkdirAll(outputFolderExampleCppwasmtime, os.ModePerm)
						if err != nil {
							return err
						}
					}

					// host files
					outputFolderBindingCppwasmtimeHost := outputFolderBindings + "/Cppwasmtime/host"
					err = os.MkdirAll(outputFolderBindingCppwasmtimeHost, os.ModePerm)
					if err != nil {
						return err
					}

					CTypesHeaderNameHost := path.Join(outputFolderBindingCppwasmtimeHost, component.BaseName+"_types.h")
					err = CreateCTypesHeader(component, CTypesHeaderNameHost)
					if err != nil {
						return err
					}

					err = BuildBindingCExplicit(component, outputFolderBindingCppwasmtimeHost, outputFolderExampleCppwasmtime, indentString)
					if err != nil {
						return err
					}

					outputFolderDocumentationCppwasmtime := "";
					err = BuildBindingCppwasmtimeHost(component, outputFolderBindingCppwasmtimeHost, outputFolderExampleCppwasmtime,
						outputFolderDocumentationCppwasmtime, indentString, binding.ClassIdentifier)
					if err != nil {
						return err
					}

					// guest files
					outputFolderBindingCppwasmtimeGuest := outputFolderBindings + "/Cppwasmtime/guest"
					err = os.MkdirAll(outputFolderBindingCppwasmtimeGuest, os.ModePerm)
					if err != nil {
						return err
					}

					CppTypesHeaderNameGuest := path.Join(outputFolderBindingCppwasmtimeGuest, component.BaseName+"_types.hpp")
					err = CreateCPPTypesHeader(component, CppTypesHeaderNameGuest, true)
					if err != nil {
						return err
					}

					err = BuildBindingCppwasmtimeGuest(component, outputFolderBindingCppwasmtimeGuest, outputFolderExampleCppwasmtime,
					 	outputFolderDocumentationCppwasmtime, indentString, binding.ClassIdentifier)

					// TODO: example + documentation
				}

			case "Fortran":
				{
					log.Printf("Interface binding for language \"%s\" is not yet supported.", binding.Language)
				}

			default:
				log.Fatal("Unknown binding export")
			}
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

				if (!suppressStub) {
				
					if (stubDirectoryOverride != "") {
						outputFolderImplementationCppStub = stubDirectoryOverride;
					}
				
					err = os.MkdirAll(outputFolderImplementationCppStub, os.ModePerm)
					if err != nil {
						return err
					}
				}

				if (!suppressInterfaces) {

					if (interfacesDirectoryOverride != "") {
						outputFolderImplementationCpp = interfacesDirectoryOverride;
					}

					err = os.MkdirAll(outputFolderImplementationCpp, os.ModePerm)
					if err != nil {
						return err
					}

					CTypesHeaderName := path.Join(outputFolderImplementationCpp, component.BaseName+"_types.hpp")
					err = CreateCPPTypesHeader(component, CTypesHeaderName, false)
					if err != nil {
						return err
					}

					CHeaderName := path.Join(outputFolderImplementationCpp, component.BaseName+"_abi.hpp")
					err = CreateCPPAbiHeader(component, CHeaderName)
					if err != nil {
						return err
					}
				}

				err = BuildImplementationCPP(component, outputFolderImplementationCpp, outputFolderImplementationCppStub,
					outputFolderImplementationProject, implementation, suppressStub, suppressInterfaces)
				if err != nil {
					return err
				}
			}

		case "Pascal":
			{
				outputFolderImplementationProject := outputFolderImplementations + "/Pascal"
				outputFolderImplementationPascal := outputFolderImplementations + "/Pascal/Interfaces"
				outputFolderImplementationPascalStub := outputFolderImplementations + "/Pascal/Stub"

				if (!suppressStub) {
					if (stubDirectoryOverride != "") {
						outputFolderImplementationPascalStub = stubDirectoryOverride;
					}
				
					err = os.MkdirAll(outputFolderImplementationPascalStub, os.ModePerm)
					if err != nil {
						return err
					}
				}


				if (!suppressInterfaces) {
					if (interfacesDirectoryOverride != "") {
						outputFolderImplementationPascal = interfacesDirectoryOverride;
					}
					err = os.MkdirAll(outputFolderImplementationPascal, os.ModePerm)
					if err != nil {
						return err
					}
					err = BuildImplementationPascal(component, outputFolderImplementationPascal, outputFolderImplementationPascalStub,
						outputFolderImplementationProject, implementation, suppressStub, suppressInterfaces)
					if err != nil {
						return err
					}
					
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


func printUsageInfo() {
	fmt.Fprintln(os.Stdout, "Run ACT with the Interface Description XML as command line parameter:")
	fmt.Fprintln(os.Stdout, "  $>act INTERFACEDESCRIPTION.xml [FLAGS]")
	fmt.Fprintln(os.Stdout, "  ")
	fmt.Fprintln(os.Stdout, "  run  act -h  to print this message.")
	fmt.Fprintln(os.Stdout, "  ")
	fmt.Fprintln(os.Stdout, "ACT has the following optional flags:")
	fmt.Fprintln(os.Stdout, "  -o: specify a path for the generated source code: \"-o ABSOLUTE_PATH_TO_OUTPUT_FOLDER\"")
	fmt.Fprintln(os.Stdout, "  -d: create a diff between two versions of an Interface Description XML: \"-d OTHER_IDL_FILE\"")
	fmt.Fprintln(os.Stdout, "  -bindings: specify a bindings override directory: \"-bindings TARGET_PATH_TO_BINDINGS\"")
	fmt.Fprintln(os.Stdout, "  -interfaces: specify a interfaces override directory: \"-bindings TARGET_PATH_TO_INTERFACES\"; Note that all implementations will use the same path.")
	fmt.Fprintln(os.Stdout, "  -stubs: specify a stubs override directory: \"-bindings TARGET_PATH_TO_STUBS\"; Note that all implementations will use the same path.")
	fmt.Fprintln(os.Stdout, "  -suppresslicense: do not generate a license-file")
	fmt.Fprintln(os.Stdout, "  -suppressbindings: do not generate bindings (even if the XML-file specifies them)")
	fmt.Fprintln(os.Stdout, "  -suppressstub: do not generate the content of the stubs-folder")
	fmt.Fprintln(os.Stdout, "  -suppressinterfaces: do not generate the contents of the interfaces-folder")
	fmt.Fprintln(os.Stdout, "  -suppresssubcomponents: do not generate any files for subcomponents")
	fmt.Fprintln(os.Stdout, "  -suppressexamples: do not generate any examples")
	fmt.Fprintln(os.Stdout, "  ")
	fmt.Fprintln(os.Stdout, "Tutorials, info and source-code on: https://github.com/Autodesk/AutomaticComponentToolkit/ .")
	fmt.Fprintln(os.Stdout, "ACT stops now.")
}

func main() {
	ACTVersion := "1.8.0-develop"
	fmt.Fprintln(os.Stdout, "Automatic Component Toolkit v"+ACTVersion)
	if len(os.Args) < 2 {
		printUsageInfo()
		return
	}
	if (strings.ToLower(os.Args[1]) == "-v") || (strings.ToLower(os.Args[1]) == "--version") {
		fmt.Fprintln(os.Stdout, "Version: "+ACTVersion)
		return
	}
	if (strings.ToLower(os.Args[1]) == "-h") || (strings.ToLower(os.Args[1]) == "--help") || (strings.ToLower(os.Args[1]) == "--usage") || (strings.ToLower(os.Args[1]) == "/h") {
		printUsageInfo()
		return
	}
	log.Printf("---------------------------------------\n")

	mode := eACTModeGenerate
	outfolderBase, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	diffFile := ""
	bindingsDirectoryOverride := ""
	interfacesDirectoryOverride := ""
	stubDirectoryOverride := ""
	
	suppressLicense := false;
	suppressBindings := false;
	suppressStub := false;
	suppressInterfaces := false;
	suppressSubcomponents := false;
	suppressExamples := false;
	
	if len(os.Args) >= 4 {
		for idx := 2; idx < len(os.Args); idx ++ {
			if os.Args[idx] == "-o" {
				outfolderBase = os.Args[idx + 1]
			}
	
			if os.Args[idx] == "-d" {
				diffFile = os.Args[idx + 1]
				mode = eACTModeDiff
			}
			
			if os.Args[idx] == "-bindings" {
				bindingsDirectoryOverride = os.Args[idx + 1]
				log.Printf("Bindings override directory: %s", bindingsDirectoryOverride)
			}

			if os.Args[idx] == "-interfaces" {
				interfacesDirectoryOverride = os.Args[idx + 1]
				log.Printf("Interfaces override directory: %s", interfacesDirectoryOverride)
			}

			if os.Args[idx] == "-stubs" {
				stubDirectoryOverride = os.Args[idx + 1]
				log.Printf("Stub override directory: %s", stubDirectoryOverride)
			}
				
			if os.Args[idx] == "-suppresslicense" {
				suppressLicense = true;
			}

			if os.Args[idx] == "-suppressbindings" {
				suppressBindings = true;
			}

			if os.Args[idx] == "-suppressstub" {
				suppressStub = true;
			}

			if os.Args[idx] == "-suppressinterfaces" {
				suppressInterfaces = true;
			}

			if os.Args[idx] == "-suppresssubcomponents" {
				suppressSubcomponents = true;
			}

			if os.Args[idx] == "-suppressexamples" {
				suppressExamples = true;
			}
			
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
	
	err = createComponent(component, outfolderBase, bindingsDirectoryOverride, interfacesDirectoryOverride, stubDirectoryOverride, suppressBindings, suppressStub, suppressInterfaces, suppressSubcomponents, suppressLicense, suppressExamples)
	if (err != nil) {
		if err == ErrPythonBuildFailed {
			log.Println("Python binding generation failed (Due to usage of reserved keywords)")
		} else {
			log.Println("Fatal error")
			log.Fatal(err)
		}
	} else {
		log.Println("Success")
	}
}
