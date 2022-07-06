package producer

import (
	"fmt"
	"github.com/streadway/amqp"
)

func SendMsg(jsonData []byte) {
	fmt.Println("Processing contact form message...")
	conn, err := amqp.Dial("amqp://test:test@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()


	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}


    // attempt to publish a message to the queue!
        err = ch.Publish(
                "",
                "TestQueue",
                false,
                false,
                amqp.Publishing{
                        ContentType: "text/plain",
                        Body:        []byte(jsonData),
                },
        )

        if err != nil {
                fmt.Println(err)
        }

	fmt.Printf("%#v\n", q)
	conn.Close()
	fmt.Println("[Closed connection] Successfully Published Message to Queue")
}
