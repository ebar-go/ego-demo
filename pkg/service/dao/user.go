package dao

import (
	"ego-demo/pkg/service/entity"
	"github.com/ebar-go/ego/component/mysql"
	"github.com/jinzhu/gorm"
)

type userDao struct {
	mysql.Dao
}

func User(db *gorm.DB) *userDao {
	return &userDao{mysql.Dao{DB: db}}
}

// GetByUsername 根据用户名获取记录
func (dao *userDao) GetByEmail(email string) (*entity.UserEntity, error) {
	query := dao.DB.Table(entity.TableUser).
		Where("email = ?", email).
		Where(entity.SoftDeleteCondition)

	user := new(entity.UserEntity)
	if err := query.First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
