package service

import (
	"ego-demo/pkg/request"
	"ego-demo/pkg/response"
)

type UserService interface {
	Auth(req request.UserAuthRequest) (*response.UserAuthResponse, error)
	Register(req request.UserRegisterRequest) error
}

