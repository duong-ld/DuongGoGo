package tasks

import (
	"duongGoGo/infra/rabbitmq"
	"duongGoGo/modules/user"
	"fmt"
)

func luckyUser() {
	userRepo := new(user.Repository)
	user, err := userRepo.GetLuckyUser()

	if err != nil {
		panic("Something went wrong when get lucky user")
		return
	}

	message := fmt.Sprintf("%s:%d", user.Email, 1)
	rabbitmq.PublishMessageToEmailQueue(message)
}
