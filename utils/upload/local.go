package upload

import (
	"errors"
	"ginserver/global"
	"ginserver/utils"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Local struct {
}

func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {

	ext := path.Ext(file.Filename)
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	filename := name + "_" + time.Now().Format(time.DateOnly) + ext
	mkdirErr := os.MkdirAll(global.GVA_CONFIG.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		global.GVA_LOG.Error("function os.MkdirAll() failed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() failed,err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.GVA_CONFIG.Local.StorePath + "/" + filename
	filepath := global.GVA_CONFIG.Local.Path + "/" + filename
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close()
	out, createErr := os.Create(p)
	if createErr != nil {
		global.GVA_LOG.Error("function os.Create() failed", zap.Any("err", createErr.Error()))
	}
	defer out.Close()

	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		global.GVA_LOG.Error("function io.Copy() failed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.copy() failed,err:" + copyErr.Error())
	}
	return filepath, filename, nil

}

func (*Local) DeleteFile(key string) error {
	p := global.GVA_CONFIG.Local.StorePath + "/" + key
	if strings.Contains(p, global.GVA_CONFIG.Local.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败，err:" + err.Error())
		}

	}
	return nil
}
