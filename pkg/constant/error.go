package constant

import "fmt"

var (
	PasswordWrong = fmt.Errorf("密码错误")
	EmailExist    = fmt.Errorf("该邮箱已存在")
)
