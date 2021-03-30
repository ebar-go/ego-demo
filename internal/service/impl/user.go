package impl

import (
	"ego-demo/internal/enum/statusCode"
	"ego-demo/internal/repository"
	"ego-demo/internal/service"
	"ego-demo/pkg/dao"
	"ego-demo/pkg/entity"
	"ego-demo/pkg/request"
	"ego-demo/pkg/response"
	"fmt"
	"github.com/ebar-go/ego/component/auth"
	"github.com/ebar-go/ego/component/mysql"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/egu"
	"gorm.io/gorm"
)

type userService struct {
	db  mysql.Database
	jwt auth.Jwt
	userRepo repository.UserRepo
	tokenRepo repository.TokenRepo
}

func newUserService(db mysql.Database) service.UserService {
	return &userService{db: db}
}

// Auth 校验用户登录
func (service *userService) Auth(req request.UserAuthRequest) (*response.UserAuthResponse, error) {
	// 调用dao的方法，根据邮箱获取用户信息
	userEntity, err :=service.userRepo.FindByEmail(req.Email)
	if err != nil {
		// 查询数据库失败时,返回自定义的error类型，用于panic拦截输出
		return nil, errors.Sprintf(statusCode.DataNotFound, "获取用户信息失败:%v", err)
	}

	// 校验密码
	if !userEntity.ValidatePass(req.Pass) {
		return nil, errors.New(statusCode.PasswordWrong, "密码错误")
	}

	// 组装结构体
	res := new(response.UserAuthResponse)

	userData := service.userRepo.BuildUserData(userEntity)
	res.Token, err = service.tokenRepo.CreateToken(userData)
	if err != nil {
		return nil, errors.New(statusCode.TokenGenerateFailed, err.Error())
	}

	return res, nil
}

// Register 注册
func (service *userService) Register(req request.UserRegisterRequest) error {
	err := service.db.GetInstance().Transaction(func(tx *gorm.DB) error {
		userDao := dao.User(service.db.GetInstance())
		// 根据邮箱获取用户信息
		user, err := userDao.GetByEmail(req.Email)
		if err != nil && err != gorm.ErrRecordNotFound {
			return fmt.Errorf("获取用户信息失败:%v", err)
		}

		// 用户已存在
		if user != nil {
			return fmt.Errorf("该邮箱已被注册")
		}

		now := int(egu.GetTimeStamp())

		user = new(entity.UserEntity)
		user.Email = req.Email
		user.Password = egu.Md5(req.Pass)
		user.CreatedAt = now
		user.UpdatedAt = now

		if err := userDao.Create(user); err != nil {
			return fmt.Errorf( "保存数据失败:%v", err)
		}

		return nil
	})

	if err != nil {
		return errors.New(statusCode.DatabaseSaveFailed, err.Error())
	}


	return nil
}
