package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"text/template"
)

type tmplTypeInfo struct {
	TypeName   string
	GoTypeName string
	TypeKind   string
	RefKind    string
	ZeroValue  string
	BitSize    int
	Sizeof     int
}

type tmplConfTypeInfoList struct {
	TL []tmplTypeInfo
}

var (
	typeInfoListForceconv = []tmplTypeInfo{
		{TypeName: "Bool", GoTypeName: "bool", TypeKind: "Bool", ZeroValue: "false", Sizeof: 1},
		{TypeName: "Int8", GoTypeName: "int8", TypeKind: "Int", ZeroValue: "0", Sizeof: 1},
		{TypeName: "Int16", GoTypeName: "int16", TypeKind: "Int", ZeroValue: "0", Sizeof: 2},
		{TypeName: "Int32", GoTypeName: "int32", TypeKind: "Int", ZeroValue: "0", Sizeof: 4},
		{TypeName: "Int64", GoTypeName: "int64", TypeKind: "Int", ZeroValue: "0", Sizeof: 8},
		{TypeName: "Uint8", GoTypeName: "uint8", TypeKind: "Uint", ZeroValue: "0", Sizeof: 1},
		{TypeName: "Uint16", GoTypeName: "uint16", TypeKind: "Uint", ZeroValue: "0", Sizeof: 2},
		{TypeName: "Uint32", GoTypeName: "uint32", TypeKind: "Uint", ZeroValue: "0", Sizeof: 4},
		{TypeName: "Uint64", GoTypeName: "uint64", TypeKind: "Uint", ZeroValue: "0", Sizeof: 8},
		{TypeName: "Float32", GoTypeName: "float32", TypeKind: "Float", ZeroValue: "0.0", Sizeof: 4},
		{TypeName: "Float64", GoTypeName: "float64", TypeKind: "Float", ZeroValue: "0.0", Sizeof: 8},
		{TypeName: "Complex64", GoTypeName: "complex64", TypeKind: "Complex", ZeroValue: "0", Sizeof: 4},
		{TypeName: "Complex128", GoTypeName: "complex128", TypeKind: "Complex", ZeroValue: "0", Sizeof: 8},
		{TypeName: "String", GoTypeName: "string", TypeKind: "String", ZeroValue: "\"\"", Sizeof: 16},
	}

	tmplDataForceconv = tmplConfTypeInfoList{
		TL: typeInfoListForceconv,
	}
)

type tmplSet struct {
	Name   string
	Input  string
	Output string
	Config interface{}
}

func main() {
	set := []tmplSet{
		{Name: "Forceconv", Input: "tmpl_forceconv.go", Output: "forceconv.gen.go", Config: tmplDataForceconv},
	}

	PROJECT_PATH := os.Getenv("GOPATH") + "/src/github.com/shunsukuda/forceconv/"
	TEMPLATE_DIR := PROJECT_PATH + "_internal/template/"

	for i := range set {
		fmt.Printf("generate code: %s\n", set[i].Name)
		inputPath := TEMPLATE_DIR + set[i].Input
		outputPath := PROJECT_PATH + set[i].Output
		tmplFile, _ := os.Open(inputPath)
		defer tmplFile.Close()
		tmplBuf := bytes.NewBuffer(nil)
		io.Copy(tmplBuf, tmplFile)
		fmt.Printf("input %s %d bytes\n", inputPath, tmplBuf.Len())
		tmpl, err := template.New(set[i].Name).Parse(tmplBuf.String())
		if err != nil {
			log.Fatal(err)
		}
		outFile, _ := os.Create(outputPath)
		defer outFile.Close()
		if err = tmpl.Execute(outFile, set[i].Config); err != nil {
			log.Fatal(err)
		}
		if err = exec.Command("go", "fmt", outputPath).Start(); err != nil {
			log.Fatal(err)
		}
	}

}
