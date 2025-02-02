package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func UserEmailSender(email string)bool{
	conn , err:=amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err,"Failed to connect to RabbitMQ")
	defer conn.Close()
	 
	ch,err:=conn.Channel()

	failOnError(err,"Failed to open a channel")
	defer ch.Close()

	q,err:=ch.QueueDeclare(
		"task_queue",	//name		
		false,	//durable
		false,	//delete when unused
		false,	//exclusive
		false,	//no-wait
		nil,	//arguments
	)

	failOnError(err, "Failed to declare a queue")

	body := "kumarpriydarshi6@gmail.com"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
	return true
}


