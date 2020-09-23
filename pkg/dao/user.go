package dao

import (
	"ego-demo/pkg/entity"
	"gorm.io/gorm"
)

type userDao struct {
	baseDao
}

func User(db *gorm.DB) *userDao {
	return &userDao{baseDao{db}}
}

// GetByUsername 根据用户名获取记录
func (dao *userDao) GetByEmail(email string) (*entity.UserEntity, error) {
	query := dao.db.Table(entity.TableUser).
		Where("email = ?", email).
		Where(entity.SoftDeleteCondition)

	user := new(entity.UserEntity)
	if err := query.First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
