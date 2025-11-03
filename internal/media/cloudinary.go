package media

import (
    "context"
    "mime/multipart"
    "os"

    "github.com/cloudinary/cloudinary-go/v2"
    "github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Uploader interface {
    UploadFromFileHeader(ctx context.Context, fh *multipart.FileHeader, folder string) (string, error)
}

type CloudinaryUploader struct {
    cld *cloudinary.Cloudinary
}

func NewCloudinaryUploader() (*CloudinaryUploader, error) {
    cld, err := cloudinary.NewFromParams(
        os.Getenv("CLOUDINARY_CLOUD_NAME"),
        os.Getenv("CLOUDINARY_API_KEY"),
        os.Getenv("CLOUDINARY_API_SECRET"),
    )
    if err != nil {
        return nil, err
    }
    return &CloudinaryUploader{cld: cld}, nil
}

func (u *CloudinaryUploader) UploadFromFileHeader(ctx context.Context, fh *multipart.FileHeader, folder string) (string, error) {
    file, err := fh.Open()
    if err != nil {
        return "", err
    }
    defer file.Close()

    res, err := u.cld.Upload.Upload(ctx, file, uploader.UploadParams{
        ResourceType: "image",
        Folder:       folder,
    })
    if err != nil {
        return "", err
    }
    return res.SecureURL, nil
}