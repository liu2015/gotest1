package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
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
	FuncNode := FindFunction(astFile, funcName)
	if FuncNode != nil {
		ast.Print(fileSet, FuncNode)
	}
	// 增加一个db库变量
	addDBVar(FuncNode.Body, varName, dbName)
	addAutoMigrate(FuncNode.Body, varName, pk, model)
	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)
	os.WriteFile(path, bf.Bytes(), 0666)

}

func addDBVar(astBody *ast.BlockStmt, varName, dbName string) {

	if dbName == "" {
		return
	}
	dbStr := fmt.Sprintf("\"%s\"", dbName)

	for v := range astBody.List {
		if assignStmt, ok := astBody.List[v].(*ast.AssignStmt); ok {

			if ident, ok := assignStmt.Lhs[0].(*ast.Ident); ok {
				if ident.Name == varName {
					return
				}
			}
		}
	}
	assignNode := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: varName}},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "global",
					},
					Sel: &ast.Ident{
						Name: "GetGlobalDBByDBName",
					},
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: dbStr,
					},
				},
			},
		},
	}
	astBody.List = append([]ast.Stmt{assignNode}, astBody.List...)

}

// db库变量增加
func addAutoMigrate(astBody *ast.BlockStmt, dbname string, pk string, model string) {
	if dbname == "" {
		dbname = "db"
	}
	flag := true
	ast.Inspect(astBody, func(n ast.Node) bool {
		switch n1 := n.(type) {
		case *ast.CallExpr:
			if s, ok := n1.Fun.(*ast.SelectorExpr); ok {
				if x, ok := s.X.(*ast.Ident); ok {
					if s.Sel.Name == "AutoMigrate" && x.Name == dbname {
						flag = false
						if !needAppendInit(n, pk, model) {
							return false
						}
						n1.Args = append(n1.Args, &ast.CompositeLit{
							Type: &ast.SelectorExpr{
								X: &ast.Ident{
									Name: pk,
								},
								Sel: &ast.Ident{
									Name: model,
								},
							}})
						return false
					}

				}
			}
		}
		return true
	})

	if flag {
		exprStmt := &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: dbname,
					},
					Sel: &ast.Ident{
						Name: "AutoMigrate",
					},
				},
				Args: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: pk,
							},
							Sel: &ast.Ident{
								Name: model,
							},
						},
					},
				},
			},
		}
		astBody.List = append(astBody.List, exprStmt)
	}

}

func NeedAppendModel(callNode ast.Node, pk string, model string) bool {
	flag := true
	ast.Inspect(callNode, func(n ast.Node) bool {
		switch n1 := n.(type) {
		case *ast.SelectorExpr:
			if x, ok := n1.X.(*ast.Ident); ok {
				if n1.Sel.Name == model && x.Name == pk {
					flag = false
					return false
				}
			}
		}
		return true
	})
	return flag
}
