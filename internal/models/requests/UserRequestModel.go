package requests

type UserRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type LoginRequest struct {
	UserRequest
	Captcha   string `json:"captcha" form:"captcha" binding:"required"`
	CaptchaId string `json:"captcha_id" form:"captcha_id" binding:"required"`
}

type CreateUserRequest struct {
	UserRequest
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
}

type GetUserInfoRequest struct {
	Data string `json:"data" form:"data"`
	Type string `json:"type" form:"type" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
}
