//go:build packfile
// +build packfile

package packfile

import (
	"os"
	"strings"
)

//go:generate go-bindata -o=staticFile.go -pkg=packfile -tags=packfile ../resource/... ../config.yaml

func writeFile(path string, data []byte) {
	if lastSeparator := strings.LastIndex(path, "."); lastSeparator != -1 {
		dirPath := path[:lastSeparator]
		if _, err := os.Stat(dirPath); err != nil && os.IsNotExist(err) {
			os.MkdirAll(dirPath, os.ModePerm)
		}

	}
}
