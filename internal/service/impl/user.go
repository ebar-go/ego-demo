package impl

import (
	"ego-demo/internal/dto/request"
	"ego-demo/internal/dto/response"
	"ego-demo/internal/entity"
	"ego-demo/internal/enum"
	"ego-demo/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/ebar-go/ego/component/auth"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/egu"
)

type userService struct {
	jwt auth.Jwt
	userRepo repository.UserRepo
}

type userClaims struct {
	jwt.StandardClaims
	UserId int
}

func NewUserService(userRepo repository.UserRepo, jwt auth.Jwt) *userService {
	return &userService{jwt: jwt, userRepo: userRepo}
}

// Auth 校验用户登录
func (service *userService) Auth(req request.UserAuthRequest) (*response.UserAuthResponse, error) {
	// 调用dao的方法，根据邮箱获取用户信息
	userEntity, err :=service.userRepo.FindByEmail(req.Email)
	if err != nil {
		// 查询数据库失败时,返回自定义的error类型，用于panic拦截输出
		return nil, errors.Sprintf(enum.DataNotFound, "获取用户信息失败:%v", err)
	}

	// 校验密码
	if !userEntity.ValidatePass(req.Pass) {
		return nil, errors.New(enum.PasswordWrong, "密码错误")
	}

	// 组装结构体
	res := new(response.UserAuthResponse)
	// 生成token
	claims := new(userClaims)
	claims.ExpiresAt = egu.GetTimeStamp() + 3600
	claims.UserId = userEntity.ID

	res.Token, err = service.jwt.GenerateToken(claims)
	if err != nil {
		return nil, errors.New(enum.TokenGenerateFailed, err.Error())
	}

	return res, nil
}

// Register 注册
func (service *userService) Register(req request.UserRegisterRequest) error {
	user := new(entity.UserEntity)
	user.Email = req.Email
	user.Password = egu.Md5(req.Pass)

	if err := service.userRepo.CreateUser(user); err != nil {
		return errors.New(enum.DatabaseSaveFailed, err.Error())
	}

	return nil
}
