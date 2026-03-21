package config

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var Cloudinary *cloudinary.Cloudinary

func ConnectCloudinary() {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		log.Fatal("Gagal konek ke Cloudinary:", err)
	}

	Cloudinary = cld
	log.Println("Cloudinary berhasil terhubung!")
}

func UploadImage(file interface{}) (string, error) {
	ctx := context.Background()

	uploadResult, err := Cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "medigo/obat",
	})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
