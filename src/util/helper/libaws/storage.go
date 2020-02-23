package libaws

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetStorage ...
func (aw *Aws) GetStorage() *s3.S3 {
	cfg := aw.Sessions()
	result := s3.New(session.New(), cfg)
	return result
}

// UploadFile ...
func (aw *Aws) UploadFile(filePath string) error {
	file, _ := os.Open(filePath)
	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	file.Read(buffer)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:           aws.String(filePath),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	svc := aw.GetStorage()
	_, err := svc.PutObject(params)
	if err != nil {
		return err
	}

	return nil

}

// Readmetric ...
func (aw *Aws) Readmetric(key string) (interface{}, error) {
	bucket := os.Getenv("S3_BUCKET_NAME")
	s3Client := aw.GetStorage()
	results, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return nil, err
	}
	var metricData interface{}
	err = json.Unmarshal(buf.Bytes(), metricData)
	if err != nil {
		return nil, err
	}

	return metricData, nil
}
