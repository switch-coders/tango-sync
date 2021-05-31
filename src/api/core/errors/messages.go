package errors

import (
	"fmt"
	"sort"
)

type Message string
type Parameters map[string]interface{}

const (
	// Error.
	ErrorCreatingUser           Message = "Error creating user."
	ErrorCreatingCategory       Message = "Error creating category."
	ErrorGettingResource        Message = "Error getting resource."
	ErrorUpsertResource         Message = "Error upsert resource."
	ErrorTooManyRequests        Message = "Too many requests error."
	ErrorUpdatingResource       Message = "Error Updating resource."
	ErrorUserNotFound           Message = "Error user doesn't exists."
	ErrorDataBaseConnection     Message = "Error trying to connect to database."
	ErrorBindingRequest         Message = "Error binding request."
	ErrorForbidden              Message = "Error forbidden action without authentication."
	ErrorUnmarshallingResponse  Message = "Error unmarshalling response."
	ErrorInvalidTangoToken      Message = "Error invalid tango access token."
	ErrorConnectingAMQP         Message = "Error to connecting amqp server."
	ErrorOpeningChannel         Message = "Failed to open a channel."
	ErrorFailsToPublishRabbitMQ Message = "Failed to publish message."
)

func (message Message) GetMessage() string {
	return string(message)
}

func (message Message) GetMessageWithParams(params Parameters) string {
	msg := message.GetMessage()
	keys := make([]string, 0)

	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		m := fmt.Sprintf(" %v:%v", k, params[k])
		msg = msg + m
	}

	return msg
}
