package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(zipFile string, destDir string) ([]string, error) {
	zipReader, err := zip.OpenReader(zipFile)
	var paths []string
	if err != nil {
		return []string{}, err
	}
	defer zipReader.Close()
	for _, v := range zipReader.File {
		if strings.Index(v.Name, "..") > -1 {
			return []string{}, fmt.Errorf("%s 文件名不合法", v.Name)
		}

		fpath := filepath.Join(destDir, v.Name)
		paths = append(paths, fpath)
		if v.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)

		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return []string{}, err
			}
			inFile, err := v.Open()
			if err != nil {
				return []string{}, err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, v.Mode())
			if err != nil {
				return []string{}, err
			}
			defer outFile.Close()
			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return []string{}, err
			}
		}
	}
	return paths, nil

}

func ZipFiles(filename string, files []string, oldForm, newForm string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = newZipFile.Close()
	}()
	zipWriter := zip.NewWriter(newZipFile)
	defer func() {
		_ = zipWriter.Close()
	}()
	for _, file := range files {
		err = func(file string) error {
			zipFile, err := os.Open(file)
			if err != nil {
				return err
			}
			defer zipFile.Close()
			info, err := zipFile.Stat()
			if err != nil {
				return err
			}
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			header.Name = strings.Replace(file, oldForm, newForm, -1)

			header.Method = zip.Deflate

			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			if _, err = io.Copy(writer, zipFile); err != nil {
				return err
			}
			return nil

		}(file)
		if err != nil {
			return err
		}
	}
	return nil
}
