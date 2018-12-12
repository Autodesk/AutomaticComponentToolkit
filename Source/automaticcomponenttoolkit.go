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
	"path"
	"log"
	"io/ioutil"
	"os"
	"encoding/xml"
)


func main () {

	log.Printf ("Automatic Component Toolkit v1.0\n");
	log.Printf ("---------------------------------------\n");
	if (len (os.Args) < 2) {
		log.Fatal ("Please run with the configuration XML as command line parameter.");
		log.Fatal ("To specify a path for the generated source code use the optional flag \"-o ABSOLUTE_PATH_TO_OUTPUT_FOLDER\"");
	}
	
	outfolderBase, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if (len (os.Args) >= 4) {
		if os.Args[2] == "-o" {
			outfolderBase = os.Args[3]
		}
	}
	log.Printf("Output directory: " + outfolderBase)

	log.Printf ("Loading %s", os.Args[1]);
	file, err := os.Open(os.Args[1]);
	if (err != nil) {
		log.Fatal (err);
	}

	bytes, err := ioutil.ReadAll (file);
	if (err != nil) {
		log.Fatal (err);
	}
	
	log.Printf ("Parsing Component Description File ");
	var component ComponentDefinition;
	err = xml.Unmarshal(bytes, &component)
	if (err != nil) {
		log.Fatal (err);
	}

	if component.BaseName == "" {
		log.Fatal ("Invalid export basename");
	}

	outputFolder := path.Join(outfolderBase, component.NameSpace + "_component");
	outputFolderBindings := path.Join(outputFolder, "Bindings")
	outputFolderImplementations := path.Join(outputFolder, "Implementations")
	
	err = CheckComponentDefinition (component);
	if (err != nil) {
		log.Fatal (err);
	}

	err  = os.MkdirAll(outputFolder, os.ModePerm);
	if (err != nil) {
		log.Fatal (err);
	}

	if (len(component.BindingList.Bindings) > 0) {
		err  = os.MkdirAll(outputFolderBindings, os.ModePerm);
		if (err != nil) {
			log.Fatal (err);
		}	
	}
	for bindingindex := 0; bindingindex < len(component.BindingList.Bindings); bindingindex++ {
		binding := component.BindingList.Bindings[bindingindex];
		indentString := getIndentationString(binding.Indentation)
		log.Printf ("Exporting Interface Binding for Languge \"%s\"", binding.Language);
		
		switch (binding.Language) {
			case "C": {
				outputFolderBindingC := outputFolderBindings + "/C";

				err  = os.MkdirAll(outputFolderBindingC, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}
				
				err = BuildBindingC(component, outputFolderBindingC)
				if (err != nil) {
					log.Fatal (err);
				}
			}

			case "CDynamic": {
				outputFolderBindingCDynamic := outputFolderBindings + "/CDynamic";

				err  = os.MkdirAll(outputFolderBindingCDynamic, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}
				
				CTypesHeaderName := path.Join(outputFolderBindingCDynamic, component.BaseName + "_types.h");
				err = CreateCTypesHeader (component, CTypesHeaderName);
				if (err != nil) {
					log.Fatal (err);
				}				
				
				err = BuildBindingCDynamic(component, outputFolderBindingCDynamic);
				if (err != nil) {
					log.Fatal (err);
				}
			}

			case "CppDynamic": {
				outputFolderBindingCppDynamic := outputFolderBindings + "/CppDynamic";

				err  = os.MkdirAll(outputFolderBindingCppDynamic, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}
				
				CTypesHeaderName := path.Join(outputFolderBindingCppDynamic, component.BaseName + "_types.h");
				err = CreateCTypesHeader (component, CTypesHeaderName);
				if (err != nil) {
					log.Fatal (err);
				}
				
				err = BuildBindingCppDynamic(component, outputFolderBindingCppDynamic);
				if (err != nil) {
					log.Fatal (err);
				}
			}

			case "Cpp": {
				outputFolderBindingCpp := outputFolderBindings + "/Cpp";

				err  = os.MkdirAll(outputFolderBindingCpp, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}

				CTypesHeaderName := path.Join(outputFolderBindingCpp, component.BaseName + "_types.h");
				err = CreateCTypesHeader (component, CTypesHeaderName);
				if (err != nil) {
					log.Fatal (err);
				}
				
				CHeaderName := path.Join(outputFolderBindingCpp, component.BaseName + ".h");
				err = CreateCHeader (component, CHeaderName);
				if (err != nil) {
					log.Fatal (err);
				}
				
				err = BuildBindingCPP(component, outputFolderBindingCpp);
				if (err != nil) {
					log.Fatal (err);
				}
			}

			case "Go": {
				outputFolderBindingGo := outputFolderBindings + "/Go";

				err  = os.MkdirAll(outputFolderBindingGo, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}

				err := BuildBindingGo(component, outputFolderBindingGo);
				if (err != nil) {
					log.Fatal (err);
				}
			}

			case "Node": {
				outputFolderBindingNode := outputFolderBindings + "/NodeJS";

				err  = os.MkdirAll(outputFolderBindingNode, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}
				
				CTypesHeaderName := path.Join(outputFolderBindingNode, component.BaseName + "_types.h");
				err = CreateCTypesHeader (component, CTypesHeaderName);
				if (err != nil) {
					log.Fatal (err);
				}
				
				err = BuildBindingCDynamic(component, outputFolderBindingNode);
				if (err != nil) {
					log.Fatal (err);
				}
				
				err := BuildBindingNode(component, outputFolderBindingNode, indentString);
				if (err != nil) {
					log.Fatal (err);
				}
			}
			
			case "Pascal": {
				outputFolderBindingPascal := outputFolderBindings + "/Pascal";
				
				err  = os.MkdirAll(outputFolderBindingPascal, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}
								
				err = BuildBindingPascalDynamic(component, outputFolderBindingPascal, indentString);
				if (err != nil) {
					log.Fatal (err);
				}
			}

			case "Python": {
				outputFolderBindingPython := outputFolderBindings + "/Python";
				
				err  = os.MkdirAll(outputFolderBindingPython, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}
								
				err = BuildBindingPythonDynamic(component, outputFolderBindingPython, indentString);
				if (err != nil) {
					log.Fatal (err);
				}
			}
			
			case"Fortran": {
				log.Printf ("Interface binding for language \"%s\" is not yet supported.", binding.Language);
			}

			default:
				log.Fatal ("Unknown binding export");
		}
	}

	if (len(component.ImplementationList.Implementations) > 0) {
		err  = os.MkdirAll(outputFolderImplementations, os.ModePerm);
		if (err != nil) {
			log.Fatal (err);
		}
	}
	for implementationindex := 0; implementationindex < len(component.ImplementationList.Implementations); implementationindex++ {
		implementation := component.ImplementationList.Implementations[implementationindex];
		log.Printf ("Exporting Implementation Interface for Language \"%s\"", implementation.Language);
		
		switch (implementation.Language) {
			case "Cpp": {
				outputFolderImplementationProject := outputFolderImplementations + "/Cpp";
				outputFolderImplementationCpp := outputFolderImplementations + "/Cpp/Interfaces";
				outputFolderImplementationCppStub := outputFolderImplementations + "/Cpp/Stub";

				err  = os.MkdirAll(outputFolderImplementationCpp, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}

				err  = os.MkdirAll(outputFolderImplementationCppStub, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}

				CTypesHeaderName := path.Join(outputFolderImplementationCpp, component.BaseName + "_types.h");
				err = CreateCTypesHeader (component, CTypesHeaderName);
				if (err != nil) {
					log.Fatal (err);
				}
				
				CHeaderName := path.Join(outputFolderImplementationCpp, component.BaseName + ".h");
				err = CreateCHeader (component, CHeaderName);
				if (err != nil) {
					log.Fatal (err);
				}
				
				err = BuildImplementationCPP(component, outputFolderImplementationCpp, outputFolderImplementationCppStub,
					outputFolderImplementationProject, implementation);
				if (err != nil) {
					log.Fatal (err);
				}
			}

			case "Pascal": {
				outputFolderImplementationProject := outputFolderImplementations + "/Pascal";
				outputFolderImplementationPascal := outputFolderImplementations + "/Pascal/Interfaces";
				outputFolderImplementationPascalStub := outputFolderImplementations + "/Pascal/Stub";

				err  = os.MkdirAll(outputFolderImplementationPascal, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}

				err  = os.MkdirAll(outputFolderImplementationPascalStub, os.ModePerm);
				if (err != nil) {
					log.Fatal (err);
				}
				
				err = BuildImplementationPascal(component, outputFolderImplementationPascal, outputFolderImplementationPascalStub,
					outputFolderImplementationProject, implementation);
				if (err != nil) {
					log.Fatal (err);
				}
			}
			
			case "Fortran": {
				log.Printf ("Implementation in language \"%s\" is not yet supported.", implementation.Language);
			}
			default:
				log.Fatal ("Unknown export");
		}
	}

}
