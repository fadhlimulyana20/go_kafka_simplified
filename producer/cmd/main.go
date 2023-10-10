package main

import (
	cfg "go_kafka/config"
	"go_kafka/producer"
)

func main() {
	topic := cfg.CONST_TOPIC
	producer.Produce(topic, 100)
}
