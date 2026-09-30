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
	"regexp"
	"strings"
	"unicode"
)

// javaUnicodeEscape matches Java unicode escapes, which Java decodes everywhere in a source file, including comments
var javaUnicodeEscape = regexp.MustCompile(`\\(u+[0-9a-fA-F]{4})`)

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
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "#[[", "\n]]")
}

// WriteCLicenseHeader writes a license header into a writer with C-style comments
func (writer *LanguageWriter) WriteCLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "/*", "*/")
}

// WritePascalLicenseHeader writes a license header into a writer Pascal-style comments
func (writer *LanguageWriter) WritePascalLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "(*", "*)")
}

// WritePythonLicenseHeader writes a license header into a writer Python-style comments
func (writer *LanguageWriter) WritePythonLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "'''", "'''")
}

// WriteJavaScriptLicenseHeader writes a license header with JS-style (C-style) comments.
func (writer *LanguageWriter) WriteJavaScriptLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "/*", "*/")
}

// WriteJavaLicenseHeader writes a license header into a writer Java-style comments
func (writer *LanguageWriter) WriteJavaLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "/*", "*/")
}

// WritePlainLicenseHeader writes a license header into a writer without comments
func (writer *LanguageWriter) WritePlainLicenseHeader(component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(writer.Writer, component, abstract, includeVersion, "", "")
}

// WriteLicenseHeader writes a license header into a writer with C-style comments
func WriteLicenseHeader(w io.Writer, component ComponentDefinition, abstract string, includeVersion bool) {
	writeLicenseHeaderEx(w, component, abstract, includeVersion, "/*", "*/")
}

// replaceUntilStable replaces old by new until old no longer occurs in s
func replaceUntilStable(s string, old string, new string) string {
	for strings.Contains(s, old) {
		s = strings.ReplaceAll(s, old, new)
	}
	return s
}

// sanitizeCommentText neutralizes character sequences in s that could terminate a comment or string literal
// which is closed by commentEnd, or otherwise alter the code that follows it. s is written as a single line.
func sanitizeCommentText(s string, commentEnd string) string {
	s = strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return ' '
		}
		return r
	}, s)

	switch commentEnd {
	case "*/":
		// An escape such as \u002a/ would otherwise close the comment in Java.
		s = javaUnicodeEscape.ReplaceAllString(s, "\\ $1")
		// A nested opening sequence triggers -Wcomment, which fails builds that use -Werror.
		s = replaceUntilStable(s, "*/", "* /")
		s = replaceUntilStable(s, "/*", "/ *")
		// C and C++ splice a line ending in a backslash (or the ??/ trigraph), optionally followed by
		// whitespace, with the next line before comments are recognized. This could assemble */ across lines.
		trimmed := strings.TrimRight(s, " ")
		if strings.HasSuffix(trimmed, "\\") || strings.HasSuffix(trimmed, "??/") {
			s = trimmed + "."
		}
	case "*)":
		// Free Pascal nests (* *) comments, so an opening sequence is as dangerous as a closing one.
		s = replaceUntilStable(s, "*)", "* )")
		s = replaceUntilStable(s, "(*", "( *")
	case "'''":
		// Python headers are string literals, so backslash escapes are interpreted as well.
		s = strings.ReplaceAll(s, "\\", "\\\\")
		s = strings.ReplaceAll(s, "'", "\\'")
	case "\n]]":
		s = replaceUntilStable(s, "]]", "] ]")
	}
	return s
}

// headerTextNeedsSanitizing returns true if s would be altered in the license header of any generated file
func headerTextNeedsSanitizing(s string) bool {
	for _, commentEnd := range []string{"", "*/", "*)", "'''", "\n]]"} {
		if sanitizeCommentText(s, commentEnd) != s {
			return true
		}
	}
	return false
}

// writeLicenseHeaderEx writes a license header into a writer.
func writeLicenseHeaderEx(w io.Writer, component ComponentDefinition, abstract string, includeVersion bool, CommandStart string, CommandEnd string) {
	ACTVersion := component.ACTVersion
	version := component.Version
	copyright := component.Copyright
	year := component.Year

	if len(CommandStart) > 0 {
		fmt.Fprintf(w, "%s++\n", CommandStart)
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "Copyright (C) %d %s\n", year, sanitizeCommentText(copyright, CommandEnd))
	fmt.Fprintf(w, "\n")
	for i := 0; i < len(component.License.Lines); i++ {
		line := component.License.Lines[i]
		fmt.Fprintf(w, "%s\n", sanitizeCommentText(line.Value, CommandEnd))
	}
	fmt.Fprintf(w, "\n")
	if includeVersion {
		fmt.Fprintf(w, "This file has been generated by the Automatic Component Toolkit (ACT) version %s.\n", ACTVersion)
		fmt.Fprintf(w, "\n")
	}
	if len(abstract) > 0 {
		fmt.Fprintf(w, "Abstract: %s\n", abstract)
		if includeVersion {
			fmt.Fprintf(w, "\nInterface version: %d.%d.%d\n", majorVersion(version), minorVersion(version), microVersion(version))
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
