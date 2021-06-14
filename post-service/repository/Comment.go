package repository

import "gorm.io/gorm"

type CommentRepo struct {
	Database *gorm.DB
}
