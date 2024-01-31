package image_service

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/module/image/image_repository"
	"amikom-pedia-api/utils/aws"
	"context"
	"database/sql"
	"mime/multipart"
)

type ImageServiceImpl struct {
	ImageRepository image_repository.ImageRepository
	DB              *sql.DB
	AWSS3           *aws.AwsS3
}

func NewImageService(imageRepository image_repository.ImageRepository, Db *sql.DB, awsS3 *aws.AwsS3) ImageService {
	return &ImageServiceImpl{ImageRepository: imageRepository, DB: Db, AWSS3: awsS3}
}

func (imageService *ImageServiceImpl) UploadToS3(ctx context.Context, userID string, imgAvatar, imgHeader *multipart.FileHeader) error {
	tx, err := imageService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	if imgAvatar != nil {
		avatar, err := imageService.AWSS3.UploadFile(imgAvatar, aws.ImgAvatar)
		helper.PanicIfError(err)

		avatarDomain := domain.Image{
			UserID:    sql.NullString{Valid: true, String: userID},
			FilePath:  avatar.FilePath,
			ImageType: avatar.ImageType,
			ImageUrl:  avatar.ImageUrl,
		}

		imageService.ImageRepository.Create(ctx, tx, avatarDomain)
	}

	if imgHeader != nil {
		header, err := imageService.AWSS3.UploadFile(imgHeader, aws.ImgHeader)
		helper.PanicIfError(err)

		headerDomain := domain.Image{
			UserID:    sql.NullString{Valid: true, String: userID},
			FilePath:  header.FilePath,
			ImageType: header.ImageType,
			ImageUrl:  header.ImageUrl,
		}

		imageService.ImageRepository.Create(ctx, tx, headerDomain)
	}

	return nil
}
