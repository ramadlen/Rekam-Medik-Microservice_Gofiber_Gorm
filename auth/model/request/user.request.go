package request

type UserRequestCreateRequest struct {
	NamaPanggilan string `gorm:"column:nama_panggilan" json:"nama_panggilan" validate:"required"` //bindingbisadihapusapabilapenyebaberror
	GolonganDarah string `gorm:"column:golongan_darah" json:"golongan_darah" validate:"required"`
	Email         string `gorm:"column:email" json:"email" validate:"required,email"`
	Password      string `gorm:"column:password" json:"password" validate:"required"`
}
type UserUpdateRequest struct {
	NamaPanggilan string `gorm:"column:nama_panggilan" json:"nama_panggilan" validate:"required"`
	GolonganDarah string `gorm:"column:golongan_darah" json:"golongan_darah"`
}
type UserEmailRequest struct {
	Email string `gorm:"column:email" json:"email" validate:"required"`
}
