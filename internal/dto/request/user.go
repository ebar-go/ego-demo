package request

// UserAuthRequest 登录校验请求
type UserAuthRequest struct {
	// 如果是表单提交，使用form,否则获取不到数据
	Email string `json:"email" binding:"required,email" comment:"邮箱"`       // 验证邮箱格式
	Pass  string `json:"pass" binding:"required,min=6,max=10" comment:"密码"` // 验证密码，长度为6~10
}

// UserRegisterRequest 注册请求
type UserRegisterRequest struct {
	// 邮箱
	Email string `json:"email" binding:"required,email" comment:"邮箱"`       // 验证邮箱格式
	// 密码
	Pass  string `json:"pass" binding:"required,min=6,max=10" comment:"密码"` // 验证密码，长度为6~10
}
