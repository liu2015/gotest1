package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	breakpointDir = "./breakpointDir/"
	finishDir     = "./fileDir/"
)

// 断点续传
func BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (string, error) {
	path := breakpointDir + fileMd5 + "/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	pathC, err := makeFileContent(content, fileName, path, contentNumber)
	return pathC, err
}

// 检查Md5
func CheckMd5(context []byte, chunkMd5 string) (CanUpload bool) {

	fileMd5 := MD5V(context)
	if fileMd5 == chunkMd5 {
		return true
	} else {
		return false
	}
}

// 创建切片内容
func makeFileContent(context []byte, fileName string, FileDir string, contentNumber int) (string, error) {

	if strings.Index(fileName, "..") > -1 || strings.Index(FileDir, "..") > -1 {
		return "", errors.New("文件名或者路径不合法")

	}
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber)
	f, err := os.Create(path)
	if err != nil {
		return path, err
	} else {
		_, err = f.Write(context)
		if err != nil {
			return path, err
		}
	}
	defer f.Close()
	return path, nil

}

// 创建切片
func MakeFile(fileName string, FileMd5 string) (string, error) {

	rd, err := os.ReadDir(breakpointDir + FileMd5)
	if err != nil {
		return finishDir + fileName, err
	}
	_ = os.MkdirAll(finishDir, os.ModePerm)
	fd, err := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return finishDir + fileName, err
	}

	defer fd.Close()
	for v := range rd {
		content, _ := os.ReadFile(breakpointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(v))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
			return finishDir + fileName, err
		}
	}
	return finishDir + fileName, nil
}

// 移除切片
func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(breakpointDir + FileMd5)
	return err
}
