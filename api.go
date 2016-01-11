package main

import (
	"log"

	"github.com/ironbay/transporter-go"
	"github.com/ironbay/transporter-go/amqp"
	"github.com/ironbay/transporter-go/http"
	"github.com/uplevel-security/go/services/rabbit"
)

func main() {
	hl := http.NewListener()
	ae, err := amqp.NewEmitter(rabbit.Conn)
	if err != nil {
		log.Fatal(err)
	}

	hl.On("*", func(msg *transporter.Message, ctx transporter.Context) (interface{}, error) {
		var result interface{}
		err := ae.Send(msg.Channel, msg.Payload, msg.Meta, &result)
		return result, err
	})

	hl.Start()

}
