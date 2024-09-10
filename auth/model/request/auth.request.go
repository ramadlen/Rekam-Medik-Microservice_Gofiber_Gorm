package request

// type LoginRequest struct {
// 	Email    string `json:"email" gorm:"column:email" validate:"required,email"`
// 	Password string `json:"password" gorm:"column:password" validate:"required"`
// }
type LoginRequest struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}