package config

import "time"

type PortNumber uint

// message broker settings
type MessageBrokerSettings struct {
	Host           string
	Port           PortNumber
	VHost          string
	Username       string
	Password       string
	PublishTimeout time.Duration
}
