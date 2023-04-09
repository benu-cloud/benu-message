package rabbitmq

import "time"

// message broker settings
type MessageBrokerSettings struct {
	Host           string
	Port           uint
	VHost          string
	Username       string
	Password       string
	PublishTimeout time.Duration
}
