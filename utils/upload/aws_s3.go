package upload

import (
	"errors"
	"fmt"
	"ginserver/global"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

type AwsS3 struct {
}

func (*AwsS3) UploadFile(file *multipart.FileHeader) (string, string, error) {

	session := newSession()
	uploader := s3manager.NewUploader(session)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	filename := global.GVA_CONFIG.AwsS3.PathPrefix + "/" + fileKey
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close()

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(global.GVA_CONFIG.AwsS3.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		global.GVA_LOG.Error("function uploader.Upload() failed", zap.Any("err", err.Error()))
		return "", "", err
	}
	return global.GVA_CONFIG.AwsS3.BaseURL + "/" + filename, fileKey, nil

}

// func (*AwsS3) DeleteFile(key string) {
// 	session := newSession()
// 	svc := s3.New(session)
// 	filename := global.GVA_CONFIG.AwsS3.PathPrefix + "/" + key
// 	bucket := global.GVA_CONFIG.AwsS3.Bucket
// 	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
// 		Bucket: aws.String(bucket),
// 		Key:    aws.String(filename),
// 	})
// 	if err != nil {
// 		global.GVA_LOG.Error("function svc.DeleteObject() failed", zap.Any("err", err.Error))
// 	}
// }

func (*AwsS3) DeleteFile(key string) error {
	session := newSession()
	svc := s3.New(session)
	filename := global.GVA_CONFIG.AwsS3.PathPrefix + "/" + key
	bucket := global.GVA_CONFIG.AwsS3.Bucket
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		global.GVA_LOG.Error("function svc.DeleteObject() failed", zap.Any("err", err.Error()))
		return errors.New("function svc.DeleteObject() failed, err:" + err.Error())
	}
	_ = svc.WaitUntilObjectExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	return nil

}

// 创建new 一个session
func newSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String(global.GVA_CONFIG.AwsS3.Region),
		Endpoint:         aws.String(global.GVA_CONFIG.AwsS3.Endpoint),
		S3ForcePathStyle: aws.Bool(global.GVA_CONFIG.AwsS3.S3ForcePathStyle),
		DisableSSL:       aws.Bool(global.GVA_CONFIG.AwsS3.DisableSSL),
		Credentials: credentials.NewStaticCredentials(
			global.GVA_CONFIG.AwsS3.SecretID,
			global.GVA_CONFIG.AwsS3.SecretKey,
			"",
		),
	})
	return sess
}
