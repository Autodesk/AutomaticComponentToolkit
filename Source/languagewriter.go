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
// languagewriter.go
// A toolkit to automatically generate software components: abstract API, implementation stubs and language bindings
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// LanguageWriter is a wrapper around a io.Writer that handles indentation
type LanguageWriter struct {
	Indentation  int
	IndentString string
	Writer       io.Writer
	CurrentLine  string
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// AddIndentationLevel adds number of indentation the writers output
func (writer *LanguageWriter) AddIndentationLevel(levels int) error {
	writer.Indentation = max(writer.Indentation+levels, 0)
	return nil
}

// ResetIndentationLevel adds indentation to all output
func (writer *LanguageWriter) ResetIndentationLevel() error {
	writer.Indentation = 0
	return nil
}

// Writeln formats a string and writes it to a line. Pairs of leading spaces will be replaced by the indent IndentString.
func (writer *LanguageWriter) Writeln(format string, a ...interface{}) (int, error) {

	leadingSpaces := 0
	for _, rune := range format {
		if rune == ' ' {
			leadingSpaces = leadingSpaces + 1
		} else {
			break
		}
	}
	leadingIndents := leadingSpaces / 2

	indentedFormat := strings.Repeat(writer.IndentString, leadingIndents+writer.Indentation) + format[leadingIndents*2:]
	return fmt.Fprintf(writer.Writer, indentedFormat+"\n", a...)
}

// Writelns writes multiple lines and processes indentation
func (writer *LanguageWriter) Writelns(prefix string, lines []string) error {
	for idx := 0; idx < len(lines); idx++ {
		_, err := writer.Writeln(prefix + lines[idx])
		if err != nil {
			return err
		}
	}

	return nil
}

// BeginLine clears the CurrentLine buffer
func (writer *LanguageWriter) BeginLine() {
	writer.CurrentLine = ""
}

// Printf formats a string and appends it to the CurrentLine buffer
func (writer *LanguageWriter) Printf(format string, a ...interface{}) {
	writer.CurrentLine = writer.CurrentLine + fmt.Sprintf(format, a...)
}

// EndLine flushes the CurrentBuffer to the internal writer
func (writer *LanguageWriter) EndLine() (int, error) {
	return writer.Writeln(writer.CurrentLine)
}

// WriteCMakeLicenseHeader writes a license header into a writer with CMake-style comments
func (writer *LanguageWriter) WriteCMakeLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "#[[", "\n]]", "")
}

// WriteCLicenseHeader writes a license header into a writer with C-style comments
func (writer *LanguageWriter) WriteCLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "/*", "*/", "")
}

// WritePascalLicenseHeader writes a license header into a writer Pascal-style comments
func (writer *LanguageWriter) WritePascalLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "(*", "*)", "")
}

// WritePythonLicenseHeader writes a license header into a writer Python-style comments
func (writer *LanguageWriter) WritePythonLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "'''", "'''", "")
}

// WriteJavaLicenseHeader writes a license header into a writer Java-style comments
func (writer *LanguageWriter) WriteJavaLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "/*", "*/", "")
}

// WriteTomlLicenseHeader writes a license header into a writer for TOML-style line prefix comments
func (writer *LanguageWriter) WriteTomlLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "", "", "# ")
}

// WritePlainLicenseHeader writes a license header into a writer without comments
func (writer *LanguageWriter) WritePlainLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "", "", "")
}

// WriteLicenseHeader writes a license header into a writer with C-style comments
func WriteLicenseHeader(w io.Writer, component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(w, component, abstract, includeVersion, "/*", "*/", "")
}

// writeLicenseHeaderEx writes a license header into a writer.
func writeLicenseHeaderEx(w io.Writer, component ComponentDefinition, abstract string, includeVersion bool, CommandStart string, CommandEnd string, prefix string) {
	ACTVersion := component.ACTVersion
	version := component.Version
	copyright := component.Copyright
	year := component.Year

	if len(CommandStart) > 0 {
		fmt.Fprintf(w, "%s++\n", CommandStart)
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "%sCopyright (C) %d %s\n", prefix, year, copyright)
	fmt.Fprintf(w, "%s\n", prefix)
	for i := 0; i < len(component.License.Lines); i++ {
		line := component.License.Lines[i]
		fmt.Fprintf(w, "%s%s\n", prefix, line.Value)
	}
	fmt.Fprintf(w, "%s\n", prefix)
	if includeVersion {
		fmt.Fprintf(w, "%sThis file has been generated by the Automatic Component Toolkit (ACT) version %s.\n", prefix, ACTVersion)
		fmt.Fprintf(w, "%s\n", prefix)
	}
	if len(abstract) > 0 {
		fmt.Fprintf(w, "%sAbstract: %s\n", prefix, abstract)
		if includeVersion {
			fmt.Fprintf(w, "%s\n%sInterface version: %d.%d.%d\n", prefix, prefix, majorVersion(version), minorVersion(version), microVersion(version))
		}
	}
	fmt.Fprintf(w, "\n")
	if len(CommandEnd) > 0 {
		fmt.Fprintf(w, "%s\n", CommandEnd)
		fmt.Fprintf(w, "\n")
	}
}

// CreateLanguageFile creates a LanguageWriter and sets its indent string
func CreateLanguageFile(fileName string, indentString string) (LanguageWriter, error) {
	var result LanguageWriter
	var err error

	result.IndentString = indentString
	result.Indentation = 0
	result.Writer, err = os.Create(fileName)
	if err != nil {
		return result, err
	}

	return result, nil
}
