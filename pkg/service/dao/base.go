package dao

import "github.com/jinzhu/gorm"

type BaseDao struct {
	db *gorm.DB
}

func (dao *BaseDao) Create(entity interface{}) error {
	return dao.db.Create(entity).Error
}

func (dao *BaseDao) Update(entity interface{}) error {
	return dao.db.Save(entity).Error
}
