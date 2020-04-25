package stringsvc

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Endpoint func(req interface{}) (interface{}, error)
type DecodeRequest func(m amqp.Delivery) (interface{}, error)
type DeliveryHandleFunc func(m amqp.Delivery)

// Subscriber can handle messages
type Subscriber interface {
	HandleDelivery(m amqp.Delivery)
}

type subscriber struct {
	e   Endpoint
	dec DecodeRequest
}

func (s subscriber) HandleDelivery(m amqp.Delivery) {
	d, err := s.dec(m)

	if err != nil {
		fmt.Println("add error handler to subsicriber here")
	}

	response, err := s.e(d)
	if err != nil {
		fmt.Println("add error handler to subsicriber here")
	}

	fmt.Println(response)
}

// NewAmqp returns a Subscriber
func NewAmqp(e Endpoint, dec DecodeRequest) Subscriber {
	return &subscriber{e, dec}
}
