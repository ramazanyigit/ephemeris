package model

import "gorm.io/gorm"

type DiaryLog struct {
	gorm.Model
	Entry string `gorm:"type:text"`
}