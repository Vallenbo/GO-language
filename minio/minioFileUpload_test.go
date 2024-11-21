package minio

// FileUploader.go MinIO example

import (
	"context"
	"log"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Test_fileUpload(*testing.T) {
	endpoint := "192.168.5.3:9000"
	accessKeyID := "R05t0i1BvQ7ILaMfNzrw"
	secretAccessKey := "pQ55LHorgsYkzvXzq7FyYO727s8rfvEUl5iFuAZP"
	ctx := context.Background()

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Upload the test file
	// Change the value of filePath if the file is in another location
	objectName := "test.txt" // 上传文件名
	filePath := "./test.txt"
	contentType := "application/octet-stream"

	// Upload the test file with FPutObject
	info, err := minioClient.FPutObject(ctx, "ms-project", objectName, filePath,
		minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}
