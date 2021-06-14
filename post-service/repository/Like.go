package repository

import "gorm.io/gorm"

type LikeRepo struct {
	Database *gorm.DB
}