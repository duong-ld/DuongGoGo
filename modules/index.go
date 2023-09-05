package modules

import (
	"duongGoGo/modules/auth"
	"duongGoGo/modules/user"
)

var UserModule = new(user.Module)
var AuthModule = new(auth.Module)

func Init() {
	UserModule.Init()
	AuthModule.Init()
}
