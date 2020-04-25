package stringsvc

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// Message represents the message recieved from the queue
type Message struct {
	Message string `json:"message"`
}

func MakeUppercaseEndpoint(s Svc) Endpoint {
	return func(request interface{}) (interface{}, error) {
		m := request.(Message)
		u := s.toUppercase(m.Message)
		return &Message{u}, nil
	}
}

// DecodeUppercaseMessage turns the Amqp message into a Message struct
func DecodeUppercaseMessage(m amqp.Delivery) (interface{}, error) {
	var msg Message
	err := json.Unmarshal(m.Body, &msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}
