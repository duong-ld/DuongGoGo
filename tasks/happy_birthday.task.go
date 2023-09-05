package tasks

import (
	"duongGoGo/infra/rabbitmq"
	"duongGoGo/modules/user"
	"fmt"
)

func happyBirthday() {
	userRepo := new(user.Repository)
	users, err := userRepo.GetUserBirthdayToday()

	if err != nil {
		panic("Something went wrong when get user have birthday today")
		return
	}

	for _, u := range users {
		message := fmt.Sprintf("%s:%d", u.Email, 0)
		rabbitmq.PublishMessageToEmailQueue(message)
	}
}
