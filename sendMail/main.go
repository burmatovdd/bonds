package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/gomail.v2"
	"os"
)

var ctx = context.Background()

type Info struct {
	Message string `json:"message"`
}

func main() {
	fmt.Println("REDIS_CONN: ", os.Getenv("REDIS_CONN"))
	var redisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_CONN"),
	})

	subscriber := redisClient.Subscribe(ctx, "sendMail")

	info := Info{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println("err: ", err.Error())
		}
		if err = json.Unmarshal([]byte(msg.Payload), &info); err != nil {
			fmt.Println("err: ", err.Error())
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Println("message is ", msg.Payload)
		fmt.Printf("%+v\n", info)

		abc := gomail.NewMessage()

		abc.SetHeader("From", "burmatov.1969@mail.ru")
		abc.SetHeader("To", info.Message)
		abc.SetHeader("Subject", "Test email")
		abc.SetBody("text/plain", "This is the best body")

		a := gomail.NewDialer("smtp.mail.ru", 465, "burmatov.1969@mail.ru", "JBmMH31usL6sffSiUfzf")

		if err := a.DialAndSend(abc); err != nil {
			fmt.Println("err: ", err.Error())
		}
	}

}
