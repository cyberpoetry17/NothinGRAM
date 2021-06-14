package repository

import "gorm.io/gorm"

type DislikeRepo struct {
	Database *gorm.DB
}
