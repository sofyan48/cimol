package libaws

import (
	"bytes"
	"net/http"
	"os"
	"time"

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
func (aw *Aws) UploadFile(data interface{}, ID, types, status string) {
	now := time.Now()
	path := "./metric/" + types + "/" + now.Format("01-02-2006") + "/" + ID

	aw.Storage.CreateFolderTree(path)
	aw.Storage.CreateJSONFile(data, path, status)

	// Open the file for use
	filePath := path + "/" + status + ".json"
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
	svc.PutObject(params)

}

// QueueReport ...
func (aw *Aws) QueueReport() {

}
