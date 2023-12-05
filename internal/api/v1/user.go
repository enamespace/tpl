package v1

type User struct {
	NickName string `json:"nickname" gorm:"column:nickname" validate:"required"`
	Password string `json:"password,omitempty" gorm:"column:password" validate:"required"`
	Email    string `json:"email" gorm:"column:email" validate:"required,email" `
	Gender   string `json:"gender" gorm:"column:gender" validate:"oneof=male female prefer_not_to"`
	Status   int    `json:"status" gorm:"column:status" validate:"omitempty"`
}
