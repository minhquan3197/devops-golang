package auth

// RequestLogin login user
type RequestLogin struct {
	Username string `json:"username" valid:"required~Tên không được để trống"`
	Password string `json:"password" valid:"required~Mật khẩu không được để trống"`
}
