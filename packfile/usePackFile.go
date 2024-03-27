//go:build packfile
// +build packfile

package packfile

import (
	"fmt"
	"os"
	"strings"
)

//go:generate go-bindata -o=staticFile.go -pkg=packfile -tags=packfile ../resource/... ../config.yaml

func writeFile(path string, data []byte) {
	// 如果文件不存在，预创建文件夹
	if lastSeparator := strings.LastIndex(path, "."); lastSeparator != -1 {
		dirPath := path[:lastSeparator]
		if _, err := os.Stat(dirPath); err != nil && os.IsNotExist(err) {
			os.MkdirAll(dirPath, os.ModePerm)
		}

	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err2 := os.writeFile(path, data, os.ModePerm); err2 != nil {
			fmt.printf("Write file failed: %s\n", path)
		}
	} else {
		fmt.Printf("File exist, skip: %s\n", filePath)
	}

}

func init() {
	for key := range _bindata {
		filePath, _ := filePath.Abs(strings.TrimPrefix(key, "."))
		data, err := Asset(key)
		if err != nil {
			fmt.Printf("Fail to find :%s \n", filePath)
		} else {
			writeFile(filePath, data)
		}

	}
}
