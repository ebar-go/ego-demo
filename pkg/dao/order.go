package dao

import (
	"ego-demo/pkg/entity"
	"gorm.io/gorm"
)

type orderDao struct {
	baseDao
}

func Order(db *gorm.DB) *orderDao  {
	return &orderDao{baseDao{db}}
}

func (dao *orderDao) GetById(id int) (*entity.OrderEntity, error) {
	item := new(entity.OrderEntity)
	if err := dao.db.First(item, id).Error; err != nil {
		return nil, err
	}

	return item, nil
}
