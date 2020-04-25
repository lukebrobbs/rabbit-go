package stringsvc

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// Message represents the message recieved from the queue
type Message struct {
	Message string `json:"message"`
}

type CountResponse struct {
	Count int `json: "count"`
}

func MakeUppercaseEndpoint(s Svc) Endpoint {
	return func(request interface{}) (interface{}, error) {
		m := request.(Message)
		u := s.toUppercase(m.Message)
		return &Message{u}, nil
	}
}

func MakeCountEndpoint(s Svc) Endpoint {
	return func(request interface{}) (interface{}, error) {
		m := request.(Message)
		u := s.count(m.Message)
		return &CountResponse{u}, nil
	}
}

// DecodeMessage turns the Amqp message into a Message struct
func DecodeMessage(m amqp.Delivery) (interface{}, error) {
	var msg Message
	err := json.Unmarshal(m.Body, &msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}
