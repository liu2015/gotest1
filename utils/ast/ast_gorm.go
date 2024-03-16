package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func AddRegisterTablesAst(path, funcName, pk, varName, dbName, model string) {

	modelPk := fmt.Sprintf("github.com/flipped-aurora/gin-vue-admin/server/model/%s", pk)
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	if err != nil {
		fmt.Println(err)
	}
	AddImport(astFile, modelPk)
	funcNode := FindFunction(astFile, funcName)
	if funcNode != nil {
		ast.Print(fileSet, funcNode)
	}

}
