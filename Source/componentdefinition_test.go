package main

import (
	"testing"
)

func TestCheckMethodValidatesParamDescriptions(t *testing.T) {
	tests := []struct {
		name        string
		description string
		valid       bool
	}{
		{"regular description", "The new value of this Variable", true},
		{"empty description", "", true},
		{"C comment breakout", "x */ int injected = 1; /* y", false},
		{"Pascal comment breakout", "x *) begin end; (* y", false},
	}

	var component ComponentDefinition
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			method := ComponentDefinitionMethod{
				MethodName:        "SetValue",
				MethodDescription: "Sets the value",
				Params: []ComponentDefinitionParam{
					{ParamName: "Value", ParamType: "double", ParamPass: "in", ParamDescription: test.description},
				},
			}
			err := component.checkMethod(method, "Variable")
			if test.valid && err != nil {
				t.Errorf("checkMethod rejected param description %q: %v", test.description, err)
			}
			if !test.valid && err == nil {
				t.Errorf("checkMethod accepted param description %q", test.description)
			}
		})
	}
}

func TestDescriptionsRejectCommentBreakout(t *testing.T) {
	safe := []string{"Returns the value", "A pointer * to the buffer", "Path C:\\Data", "Size in 100% units"}
	unsafe := []string{
		"x */ int injected = 1; /* y",
		"x /* y",
		"x *) begin end; (* y",
		"x \\u002a/ class Evil {} /\\u002a",
		"x ends with a backslash \\",
		"x ends with a trigraph ??/",
	}
	for _, s := range safe {
		if !descriptionIsValid(s) {
			t.Errorf("descriptionIsValid(%q) = false, expected true", s)
		}
		if !commentTextIsSafe(s) {
			t.Errorf("commentTextIsSafe(%q) = false, expected true", s)
		}
	}
	for _, s := range unsafe {
		if descriptionIsValid(s) {
			t.Errorf("descriptionIsValid(%q) = true, expected false", s)
		}
		if commentTextIsSafe(s) {
			t.Errorf("commentTextIsSafe(%q) = true, expected false", s)
		}
	}
	if errorDescriptionIsValid("x */ int injected = 1; /* y") {
		t.Errorf("errorDescriptionIsValid accepted a comment breakout")
	}
	if !errorDescriptionIsValid("The value is out of range") {
		t.Errorf("errorDescriptionIsValid rejected a regular description")
	}
}

func TestCheckEnumsValidatesDescriptions(t *testing.T) {
	tests := []struct {
		name              string
		enumDescription   string
		optionDescription string
		valid             bool
	}{
		{"regular descriptions", "The color", "Red color", true},
		{"empty descriptions", "", "", true},
		{"enum comment breakout", "x */ int injected = 1; /* y", "", false},
		{"option comment breakout", "", "x */ int injected = 1; /* y", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var component ComponentDefinition
			component.NameMapsLookup.enumMap = make(map[string]bool)
			component.Enums = []ComponentDefinitionEnum{{
				Name:        "Color",
				Description: test.enumDescription,
				Options:     []ComponentDefinitionEnumOption{{Name: "Red", Value: 0, Description: test.optionDescription}},
			}}
			err := component.checkEnums()
			if test.valid && err != nil {
				t.Errorf("checkEnums rejected valid descriptions: %v", err)
			}
			if !test.valid && err == nil {
				t.Errorf("checkEnums accepted enum description %q and option description %q", test.enumDescription, test.optionDescription)
			}
		})
	}
}

func TestCheckFunctionTypesValidatesParamDescriptions(t *testing.T) {
	tests := []struct {
		name        string
		description string
		valid       bool
	}{
		{"regular description", "The progress", true},
		{"empty description", "", true},
		{"C comment breakout", "x */ int injected = 1; /* y", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var component ComponentDefinition
			component.NameMapsLookup.functionTypeMap = make(map[string]bool)
			component.Functions = []ComponentDefinitionFunctionType{{
				FunctionName:        "ProgressCallback",
				FunctionDescription: "Callback to report progress",
				Params: []ComponentDefinitionParam{
					{ParamName: "Progress", ParamType: "single", ParamPass: "in", ParamDescription: test.description},
				},
			}}
			err := component.checkFunctionTypes()
			if test.valid && err != nil {
				t.Errorf("checkFunctionTypes rejected param description %q: %v", test.description, err)
			}
			if !test.valid && err == nil {
				t.Errorf("checkFunctionTypes accepted param description %q", test.description)
			}
		})
	}
}
