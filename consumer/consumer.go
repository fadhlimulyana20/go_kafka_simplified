package consumer

import (
	cfg "go_kafka/config"
	"log"

	"github.com/IBM/sarama"
)

func Consume(topic string) {
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{cfg.CONST_HOST}, config)
	if err != nil {
		log.Fatalf("[ERROR][Consumer] %v", err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("[ERROR][Consumer] %v", err)
	}

	defer partitionConsumer.Close()
	for msg := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", msg.Partition, msg.Offset, string(msg.Value))
	}
}
