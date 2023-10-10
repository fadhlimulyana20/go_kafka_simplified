package producer

import (
	cfg "go_kafka/config"
	"log"
	"strconv"
	"time"

	"github.com/IBM/sarama"
)

func Produce(topic string, limit int) {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{cfg.CONST_HOST}, config)
	if err != nil {
		log.Fatalf("[ERROR][Producer] %v", err)
		return
	}
	defer producer.Close()

	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   nil,
			Value: sarama.StringEncoder(str),
		}

		partition, offset, err := producer.SendMessage(msg)

		if err != nil {
			log.Fatalf("[ERROR][Producer] %v", err)
			return
		}
		log.Printf("[Producer] partition id: %d; offset:%d, value: %s\n", partition, offset, str)
	}
}
