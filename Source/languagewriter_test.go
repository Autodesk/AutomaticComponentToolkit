package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSanitizeHeaderText(t *testing.T) {
	tests := []struct {
		name       string
		commentEnd string
		input      string
		expected   string
	}{
		{"safe C text", "*/", "All rights reserved.", "All rights reserved."},
		{"safe Pascal text", "*)", "(C) 2018 Autodesk Inc.", "(C) 2018 Autodesk Inc."},
		{"safe CMake text", "\n]]", "list [a] of [b]", "list [a] of [b]"},
		{"C comment end", "*/", "x */ int evil(); /*", "x * / int evil(); / *"},
		{"C repeated comment end", "*/", "***//", "*** //"},
		{"C repeated comment start", "*/", "//**", "// **"},
		{"C comment end inside start", "*/", "/*/", "/ * /"},
		{"C trailing backslash", "*/", "a *\\", "a *\\."},
		{"C trailing backslash and whitespace", "*/", "a *\\ \t", "a *\\."},
		{"C trailing trigraph", "*/", "a *??/", "a *??/."},
		{"C inner backslash", "*/", "C:\\Path", "C:\\Path"},
		{"Java unicode comment end", "*/", "x \\u002a/ class Evil {} /\\u002a", "x \\ u002a/ class Evil {} /\\ u002a"},
		{"Java unicode escape with repeated u", "*/", "\\uuu002A/", "\\ uuu002A/"},
		{"Java non-escape backslash u", "*/", "C:\\users", "C:\\users"},
		{"Pascal comment end", "*)", "x *) begin halt; end; (*", "x * ) begin halt; end; ( *"},
		{"Pascal nested comment start", "*)", "((**))", "(( ** ))"},
		{"Python quotes", "'''", "x ''' ; import os ; '''", "x \\'\\'\\' ; import os ; \\'\\'\\'"},
		{"Python trailing backslash", "'''", "x \\", "x \\\\"},
		{"Python unicode escape", "'''", "\\N{", "\\\\N{"},
		{"CMake bracket end", "\n]]", "x ]] message(FATAL_ERROR pwned) #[[", "x ] ] message(FATAL_ERROR pwned) #[["},
		{"CMake repeated bracket end", "\n]]", "]]]]", "] ] ] ]"},
		{"control characters", "*/", "a\nb\rc\x00d", "a b c d"},
		{"plain keeps comment sequences", "", "*/ *) ''' ]]", "*/ *) ''' ]]"},
		{"plain removes control characters", "", "a\nb", "a b"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := sanitizeCommentText(test.input, test.commentEnd)
			if actual != test.expected {
				t.Errorf("sanitizeCommentText(%q, %q) = %q, expected %q", test.input, test.commentEnd, actual, test.expected)
			}
		})
	}
}

func TestWriteLicenseHeaderLineSplicing(t *testing.T) {
	var component ComponentDefinition
	component.Copyright = "Autodesk Inc."
	component.Year = 2026
	component.Version = "1.0.0"
	component.License.Lines = []ComponentDefinitionLicenseLine{
		{Value: "a *\\"},
		{Value: "/ int injected = 1; /\\"},
		{Value: "* comment resumes here"},
	}

	var buffer bytes.Buffer
	writeLicenseHeaderEx(&buffer, component, "", false, "/*", "*/")
	output := buffer.String()

	for _, line := range strings.Split(output, "\n") {
		trimmed := strings.TrimRight(line, " \t")
		if strings.HasSuffix(trimmed, "\\") || strings.HasSuffix(trimmed, "??/") {
			t.Errorf("header line %q would be spliced with the next line:\n%s", line, output)
		}
	}
}

func TestHeaderTextNeedsSanitizing(t *testing.T) {
	safe := []string{
		"Autodesk Inc.",
		"Redistribution and use in source and binary forms, with or without modification,",
		"(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS",
	}
	for _, s := range safe {
		if headerTextNeedsSanitizing(s) {
			t.Errorf("headerTextNeedsSanitizing(%q) = true, expected false", s)
		}
	}

	unsafe := []string{"*/", "*)", "(*", "'", "\\", "]]", "a\nb"}
	for _, s := range unsafe {
		if !headerTextNeedsSanitizing(s) {
			t.Errorf("headerTextNeedsSanitizing(%q) = false, expected true", s)
		}
	}
}

func TestWriteLicenseHeaderCannotBeEscaped(t *testing.T) {
	payload := "x */ *) (* ''' ]] \\N{ \n evil(); "
	var component ComponentDefinition
	component.Copyright = payload
	component.Year = 2026
	component.Version = "1.0.0"
	component.License.Lines = []ComponentDefinitionLicenseLine{{Value: payload}, {Value: "All rights reserved."}}

	styles := []struct {
		name         string
		commentStart string
		commentEnd   string
		counts       map[string]int
	}{
		{"C", "/*", "*/", map[string]int{"*/": 1, "/*": 1}},
		{"Pascal", "(*", "*)", map[string]int{"*)": 1, "(*": 1}},
		{"Python", "'''", "'''", map[string]int{"'''": 2}},
		{"CMake", "#[[", "\n]]", map[string]int{"]]": 1}},
	}

	for _, style := range styles {
		t.Run(style.name, func(t *testing.T) {
			var buffer bytes.Buffer
			writeLicenseHeaderEx(&buffer, component, "Abstract", true, style.commentStart, style.commentEnd)
			output := buffer.String()

			for sequence, expected := range style.counts {
				if actual := strings.Count(output, sequence); actual != expected {
					t.Errorf("%q occurs %d times in header, expected %d:\n%s", sequence, actual, expected, output)
				}
			}
			if !strings.HasSuffix(strings.TrimRight(output, "\n"), strings.TrimLeft(style.commentEnd, "\n")) {
				t.Errorf("header does not end with its comment terminator:\n%s", output)
			}
			if style.name == "Python" && strings.Contains(strings.ReplaceAll(output, "\\\\", ""), "\\N") {
				t.Errorf("header contains an unescaped backslash sequence:\n%s", output)
			}
		})
	}
}
