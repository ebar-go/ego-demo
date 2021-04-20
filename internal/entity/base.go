package entity


type BaseEntity struct {
	ID        int `gorm:"primarykey"`
	IsDeleted int `json:"is_deleted" gorm:"column:is_deleted"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

const (
	TableUser = "users"
)

const (
	SoftDeleteCondition = "is_deleted = 0"
)
