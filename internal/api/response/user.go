package response

type UserLoginResDto struct {
	// State string `json:"state"`
	Token string `json:"token"`
}

type UserInfoResDto struct {
	UserName string `json:"userName" `
	Password string `json:"passWord" `
	Phone    string `json:"phone" `
	Email    string `json:"email"  `
}
