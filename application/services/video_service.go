package services

import (
	"context"
	"encoder/application/repositories"
	"encoder/domain"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

type VideoService struct {
	Video *domain.Video
	VideoRepository *repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(v.Video.FilePath)

	reader, err := obj.NewReader(ctx)

	if err != nil {
		return err
	}

	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	file, err := os.Create(os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4")
	if err != nil {
		return err
	}

	_, err = file.Write(body)

	if err != nil {
		return err
	}

	defer file.Close()

	log.Printf("Video %v has been stored", v.Video.ID)

	return nil
}