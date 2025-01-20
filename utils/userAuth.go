package utils

import "myapp/rabbitmq"

func UserAuth(email string)bool{
	if rabbitmq.UserEmailSender(email){
		
	}
	return true
}