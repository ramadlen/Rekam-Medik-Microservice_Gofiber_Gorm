package response

import "time"

type UserResponse struct {
	ID            uint64 `gorm:"primaryKey" json:"id"`
	NamaPanggilan string ` json:"nama_panggilan"`
	GolonganDarah string ` json:"golongan_darah"`

	Email    string    `json:"email" gorm:"column:email"`
	Password string    `json:"-" gorm:"column:password"`
	CreateAt time.Time ` json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}