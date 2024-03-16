package ast

import (
	"fmt"
	"go/ast"
	"go/token"
)

func AddImport(astNode ast.Node, imp string) {

	impStr := fmt.Sprintf("\"%s\"", imp)
	ast.Inspect(astNode, func(n ast.Node) bool {
		if genDecl, ok := n.(*ast.GenDecl); ok {
			if genDecl.Tok == token.IMPORT {
				for v := range genDecl.Specs {
					if impNode, ok := genDecl.Specs[v].(*ast.ImportSpec); ok {
						if impNode.Path.Value == impStr {
							return false
						}
					}
				}
				genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: impStr,
					},
				})
			}
		}
		return true
	})
}

// 查询特定function方法

func FindFunction(astNode ast.Node, FunctionName string) *ast.FuncDecl {
	var funcDeclP *ast.FuncDecl
	ast.Inspect(astNode, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			if funcDecl.Name.String() == FunctionName {
				funcDeclP = funcDecl
				return false
			}
		}
		return true
	})

	return funcDeclP

}
