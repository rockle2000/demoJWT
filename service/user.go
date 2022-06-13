package service

type LoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserRes struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
}
