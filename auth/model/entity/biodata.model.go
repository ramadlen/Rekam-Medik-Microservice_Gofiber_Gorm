package model

import (
	"time"

	"gorm.io/gorm"
)

type DataIdentitas struct {
	ID                 uint64 `json:"id" gorm:"primaryKey;column:id"`
	NamaPanggilan      string `json:"nama_panggilan" gorm:"column:nama_panggilan"`
	GolonganDarah      string `json:"golongan_darah" gorm:"column:golongan_darah" `
	Email       	   string `gorm:"column:email" json:"email" validate:"required,email"`
	Password      	   string `json:"-" gorm:"column:password" `
	CreateAt           time.Time `json:"create_at" gorm:"column:create_at;autoCreateTime:true;"`
	UpdateAt           time.Time `json:"update_at" gorm:"column:update_at;autoCreateTime;autoUpdateTime:true;" `
	DeleteAt      gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at" `
}

func (d *DataIdentitas) TableName() string {
	return "data_identitas"
}
