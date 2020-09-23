package entity

type OrderEntity struct {
	BaseEntity
	UserId int `json:"user_id"`
	Amount float64 `json:"amount"`
}

func (OrderEntity) TableName() string  {
	return "orders"
}
