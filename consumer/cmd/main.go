package main

import (
	cfg "go_kafka/config"
	"go_kafka/consumer"
)

func main() {
	topic := cfg.CONST_TOPIC
	consumer.Consume(topic)
}
