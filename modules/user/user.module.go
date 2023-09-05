package user

type Module struct {
}

var userService *Service
var userRepository *Repository

func (m Module) Init() {
	userService = new(Service)
	userRepository = new(Repository)
}
