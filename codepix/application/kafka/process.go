package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
)

type KafkaProcessor struct {
	Database     *gorm.DB
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaProcessor(database *gorm.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database:     database,
		Producer:     producer,
		DeliveryChan: deliveryChan,
	}
}

func (k *KafkaProcessor) Consume() {
	port := 9092
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("kafka:%d", port),
		"group.id":          "consumergroup",
		"auto.offset.reset": "earLiest",
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{"test"}
	c.SubscribeTopics(topics, nil)

	fmt.Println(fmt.Sprintf("Kafka consumer has been started at %d port", port))
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(fmt.Sprintf("Message consumed with success: %d", string(msg.Value)))
		} else {
			fmt.Println(fmt.Sprintf("Message consumed with error: %d . Error: %d", string(msg.Value), err))
		}
	}
}
