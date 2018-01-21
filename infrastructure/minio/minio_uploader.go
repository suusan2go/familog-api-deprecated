package minio

import (
	"mime/multipart"

	"github.com/suusan2go/familog-api/domain/model"
	"github.com/suusan2go/familog-api/lib/uploader"
)

// ImageUploader imagege file upload
type ImageUploader struct {
	Client uploader.Uploader
}

// UploadFile upload file
func (upl *ImageUploader) UploadFile(image *model.DiaryEntryImage, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	if _, err := upl.Client.UploadImage(src, image.FilePath); err != nil {
		return err
	}
	return nil
}
