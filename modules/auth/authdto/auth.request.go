package authdto

import "duongGoGo/modules/user/userdto"

type SignInDto struct {
	userdto.CreateUserRequestDto
}

type SignUpDto struct {
	userdto.CreateUserRequestDto
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
