package auth

import (
	"duongGoGo/modules/user"
)

type Module struct {
}

var authService *Service
var userRepository *user.Repository

func (m Module) Init() {
	authService = new(Service)
	userRepository = new(user.Repository)
}
