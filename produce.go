package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type Producer struct {
 	ChatProducer sarama.SyncProducer
}

// Generate Producer with configured broker
func NewProducer() *Producer {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	c, err := sarama.NewSyncProducer([]string{
		"kafka-01:9092",
		"kafka-02:9092",
		"kafka-03:9092"}, config)
	if err != nil{
		panic(err)
	}

	return &Producer{ChatProducer: c}
}

func (p *Producer) Close()  error {
	err := p.ChatProducer.Close()
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) SendStringData(message string) error {
	partition, offset, err := p.ChatProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "chatting",
		Value: sarama.StringEncoder(message),
	})
	if err != nil {
		return err
	}

	fmt.Println("%d/%d", partition, offset)
	return nil
}