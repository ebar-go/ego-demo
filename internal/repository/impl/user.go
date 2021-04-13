package impl

import (
	"ego-demo/internal/dao"
	"ego-demo/internal/entity"
	"fmt"
	"github.com/ebar-go/ego/component/mysql"
	"github.com/ebar-go/egu"
	"gorm.io/gorm"
)

type userRepo struct {
	db mysql.Database
}

func NewUserRepo(db mysql.Database) *userRepo  {
	return &userRepo{db : db}
}

func (repo userRepo) FindByEmail(email string) (*entity.UserEntity, error) {
	return dao.User(repo.db.GetInstance()).GetByEmail(email)
}

func (repo userRepo) CreateUser(userEntity *entity.UserEntity) error {
	return repo.db.GetInstance().Transaction(func(tx *gorm.DB) error {
		userDao := dao.User(tx)

		// 根据邮箱获取用户信息
		_, err := userDao.GetByEmail(userEntity.Email)
		if err == nil {
			return fmt.Errorf("email exist")
		}
		if err != gorm.ErrRecordNotFound {
			return fmt.Errorf("unable to get user: %v", err)
		}

		now := egu.GetTimeStamp()

		userEntity.CreatedAt = now
		userEntity.UpdatedAt = now

		if err := userDao.Create(userEntity); err != nil {
			return fmt.Errorf( "unable to create user: %v", err)
		}

		return nil

	})
}

