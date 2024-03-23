package utils

import "os"

const (
	breakpointDir = "./breakpointDir/"
	finishDir     = "./fileDir/"
)

func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(breakpointDir + FileMd5)
	return err
}
