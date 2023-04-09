package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

func ParseArgs() (m MessageBrokerSettings) {
	// try to load env variables if they exist
	godotenv.Load()

	var rmqhost string
	var rmqport PortNumber = PortNumber(5672)
	var rmqvHost string
	var rmqusername string
	var rmqpassword string
	var rmqpublishTimeoutSeconds uint

	flag.StringVar(&rmqhost, "rmqhost", "localhost", "RabbitMQ message broker host.")
	flag.Var(&rmqport, "rmqport", "RabbitMQ message broker port. Should be in the range 0-65535.")
	flag.StringVar(&rmqvHost, "rmqvhost", "", "RabbitMQ virtual host.")
	flag.StringVar(&rmqusername, "rmqusername", "", "RabbitMQ username (required).")
	flag.StringVar(&rmqpassword, "rmqpassword", "", "RabbitMQ password (required).")
	flag.UintVar(&rmqpublishTimeoutSeconds, "rmqtimeout", 5, "RabbitMQ publish timeout in seconds")

	flag.Parse()

	if rmqusername == "" {
		fmt.Println("Error: the flag -rmqusername is required.")
		flag.Usage()
		os.Exit(1)
	}
	if rmqpassword == "" {
		fmt.Println("Error: the flag -rmqpassword is required.")
		flag.Usage()
		os.Exit(1)
	}

	m.Host = rmqhost
	m.Password = rmqpassword
	m.Port = rmqport
	m.Username = rmqusername
	m.VHost = rmqvHost
	m.PublishTimeout = time.Second * time.Duration(rmqpublishTimeoutSeconds)

	return
}
