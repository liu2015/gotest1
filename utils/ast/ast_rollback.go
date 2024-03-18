package ast

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bsm/ginkgo/v2/extensions/globals"
)

func RollBackAst(pk, model string) {

	RollGermBack(pk, model)
	RollRouterBack(pk, model)
}

func RollGermBack(pk, model string) {
	path := filepath.Join(global.GVA_CONFIG.AutoCode.Root, globals.GVA_CONFIG.AutoCode.Server, "initialize", "gorm.go")
	src, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
}
