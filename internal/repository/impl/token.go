package impl

import (
	"ego-demo/pkg/data"
	"github.com/ebar-go/ego/component/auth"
	"github.com/ebar-go/egu"
)

type tokenRepo struct {
	jwt auth.Jwt
}

func (repo tokenRepo) CreateToken(user data.User) (string, error) {
	// 生成token
	userClaims := new(data.UserClaims)
	userClaims.ExpiresAt = egu.GetTimeStamp() + 3600
	userClaims.User = user
	return repo.jwt.GenerateToken(userClaims)
}

func (repo tokenRepo) ValidateToken(token string) (data.User, error) {
	panic("implement me")
}

