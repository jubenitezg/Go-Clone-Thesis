package func_extractor

import (
	"go-func-extractor/common"
	"testing"
)

const (
	TestFile = "test_assets/file.go"
)

func TestFuncExtractorExtractFunctions(t *testing.T) {
	expected := []common.CodeFragment{
		{
			Code: "func (p *Person) SayHello() string {\n\treturn \"Hello \" + p.Name\n}",
			Id:   "some-id",
			Line: 10,
			Path: "test_assets/file.go",
		},
		{
			Code: "func (p *Person) SayAge() string {\n\treturn \"I am \" + strconv.Itoa(p.Age)\n}",
			Id:   "some-id",
			Line: 14,
			Path: "test_assets/file.go",
		},
		{
			Code: "func NewPerson(name string, age int) *Person {\n\treturn &Person{Name: name, Age: age}\n}",
			Id:   "some-id",
			Line: 18,
			Path: "test_assets/file.go",
		},
	}
	extractor := NewFuncExtractor(TestFile)
	functions, err := extractor.ExtractFunctions()
	if err != nil {
		t.Error("Failed to extract functions:", err)
	}
	if len(functions) != len(expected) {
		t.Errorf("Expected %d functions, got %d", len(expected), len(functions))
	}
	for i, expectedFragment := range expected {
		if functions[i].Code != expectedFragment.Code {
			t.Errorf("Expected code %s, got %s", expectedFragment.Code, functions[i].Code)
		}
		if functions[i].Line != expectedFragment.Line {
			t.Errorf("Expected line %d, got %d", expectedFragment.Line, functions[i].Line)
		}
		if functions[i].Path != expectedFragment.Path {
			t.Errorf("Expected path %s, got %s", expectedFragment.Path, functions[i].Path)
		}
	}
}
