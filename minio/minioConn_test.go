package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"testing"
)

func Test_minioUse(t *testing.T) {
	endpoint := "192.168.5.3:9000"
	accessKeyID := "R05t0i1BvQ7ILaMfNzrw"
	secretAccessKey := "pQ55LHorgsYkzvXzq7FyYO727s8rfvEUl5iFuAZP"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up
}
