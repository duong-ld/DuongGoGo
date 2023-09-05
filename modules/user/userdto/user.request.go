package userdto

type GetUserRequestDto struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type CreateUserRequestDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type UpdateUserRequestDto struct {
	BirthDay string `json:"birthday"`
}
