package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestGenerateJournalFunctionCodeStructsAndArrays(t *testing.T) {
	method := ComponentDefinitionMethod{
		MethodName: "Update",
		Params: []ComponentDefinitionParam{
			{ParamName: "Point", ParamType: "struct", ParamClass: "Coordinate", ParamPass: "in"},
			{ParamName: "Points", ParamType: "structarray", ParamClass: "Coordinate", ParamPass: "in"},
			{ParamName: "Weights", ParamType: "basicarray", ParamClass: "double", ParamPass: "in"},
			{ParamName: "Foreign", ParamType: "struct", ParamClass: "Other:Thing", ParamPass: "in"},
			{ParamName: "Box", ParamType: "struct", ParamClass: "Box", ParamPass: "out"},
			{ParamName: "Corners", ParamType: "structarray", ParamClass: "Coordinate", ParamPass: "out"},
			{ParamName: "Values", ParamType: "basicarray", ParamClass: "uint32", ParamPass: "out"},
			{ParamName: "Center", ParamType: "struct", ParamClass: "Coordinate", ParamPass: "return"},
		},
	}

	initCode, successCode, err := generateJournalFunctionCode(method, "Lib", "Shape", false)
	if err != nil {
		t.Fatalf("generateJournalFunctionCode failed: %v", err)
	}
	initText := strings.Join(initCode, "\n")
	successText := strings.Join(successCode, "\n")

	expectedInit := []string{
		`pJournalEntry->addStructParameter("Point", "Coordinate", pPoint);`,
		`pJournalEntry->addStructArrayParameter("Points", "Coordinate", pPointsBuffer, nPointsBufferSize);`,
		`pJournalEntry->addBasicArrayParameter("Weights", "double", pWeightsBuffer, nWeightsBufferSize);`,
	}
	for _, expected := range expectedInit {
		if !strings.Contains(initText, expected) {
			t.Errorf("journal init code is missing %q:\n%s", expected, initText)
		}
	}
	if strings.Contains(initText, "Foreign") {
		t.Errorf("structs of imported components must not be journaled:\n%s", initText)
	}

	expectedSuccess := []string{
		`pJournalEntry->addStructResult("Box", "Box", pBox);`,
		`pJournalEntry->addStructArrayResult("Corners", "Coordinate", pCornersBuffer, nCornersBufferSize, pCornersNeededCount);`,
		`pJournalEntry->addBasicArrayResult("Values", "uint32", pValuesBuffer, nValuesBufferSize, pValuesNeededCount);`,
		`pJournalEntry->addStructResult("Center", "Coordinate", pCenter);`,
	}
	for _, expected := range expectedSuccess {
		if !strings.Contains(successText, expected) {
			t.Errorf("journal success code is missing %q:\n%s", expected, successText)
		}
	}
}

func TestBuildJournalingCPPSerializesStructsAndEscapesXML(t *testing.T) {
	var component ComponentDefinition
	component.NameSpace = "Lib"
	component.BaseName = "lib"
	component.Version = "1.0.0"
	component.Structs = []ComponentDefinitionStruct{
		{
			Name: "Sample",
			Members: []ComponentDefinitionMember{
				{Name: "Value", Type: "double"},
				{Name: "Mode", Type: "enum", Class: "Mode"},
				{Name: "Vector", Type: "single", Rows: 3},
				{Name: "Matrix", Type: "double", Rows: 4, Columns: 4},
			},
		},
	}

	var header, impl bytes.Buffer
	headerw := LanguageWriter{IndentString: "  ", Writer: &header}
	implw := LanguageWriter{IndentString: "  ", Writer: &impl}
	if err := buildJournalingCPP(component, headerw, implw); err != nil {
		t.Fatalf("buildJournalingCPP failed: %v", err)
	}
	headerText := header.String()
	implText := impl.String()

	expectedHeader := []string{
		"std::string LibJournalStructToXML (const sLibSample & Value, const std::string & sIndent);",
		"template <typename T> void addStructParameter(",
		"template <typename T> void addStructArrayResult(",
		"template <typename T> void addBasicArrayResult(",
	}
	for _, expected := range expectedHeader {
		if !strings.Contains(headerText, expected) {
			t.Errorf("journal header is missing %q", expected)
		}
	}

	expectedImpl := []string{
		`#include "lib_interfacejournal.hpp"`,
		`<member name=\"Value\" type=\"double\" value=\"" << LibJournalValueToString(Value.m_Value)`,
		`<member name=\"Mode\" type=\"enumMode\" value=\"" << LibJournalValueToString((Lib_int32) Value.m_Mode)`,
		`LibJournalValueToString(Value.m_Vector[nRow])`,
		`LibJournalValueToString(Value.m_Matrix[nColumn][nRow])`,
		`case '"': sResult += "&quot;"; break;`,
		`case '<': sResult += "&lt;"; break;`,
		`LibJournalEscapeXML(Value.m_sValue)`,
		`addParameter(sName, "int32", std::to_string(nValue));`,
	}
	for _, expected := range expectedImpl {
		if !strings.Contains(implText, expected) {
			t.Errorf("journal implementation is missing %q", expected)
		}
	}
}
