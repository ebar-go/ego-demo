package entity

type BaseEntity struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	IsDeleted int `json:"is_deleted" gorm:"column:is_deleted"`
	// 对time进行格式化
	CreatedAt int `gorm:"column:created_at" json:"created_at"`
	UpdatedAt int `gorm:"column:updated_at" json:"updated_at"`
}

const (
	TableUser = "users"
)

const (
	SoftDeleteCondition = "is_deleted = 0"
)
