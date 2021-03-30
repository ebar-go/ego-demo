package impl

import (
	"ego-demo/pkg/dao"
	"ego-demo/pkg/data"
	"ego-demo/pkg/entity"
	"github.com/ebar-go/ego/component/mysql"
)

type userRepo struct {
	db mysql.Database
}

func (repo userRepo) FindByEmail(email string) (*entity.UserEntity, error) {
	return dao.User(repo.db.GetInstance()).GetByEmail(email)
}

func (repo userRepo) BuildUserData(userEntity *entity.UserEntity) data.User {
	return data.User{
		Id:    userEntity.Id,
		Email: userEntity.Email,
	}
}

