package entity

type UserEntity struct {
	BaseEntity
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

// TableName 指定模型的表名称
func (UserEntity) TableName() string {
	return TableUser
}
