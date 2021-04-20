package dao

import "gorm.io/gorm"

type baseDao struct {
	db *gorm.DB
}

func (dao *baseDao) Create(entity interface{}) error {
	return dao.db.Create(entity).Error
}

func (dao *baseDao) Update(entity interface{}) error {
	return dao.db.Save(entity).Error
}
