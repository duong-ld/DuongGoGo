package tasks

import "duongGoGo/infra/rabbitmq"

func sayHello() {
	rabbitmq.PublishMessageToEmailQueue("0:0")
}
