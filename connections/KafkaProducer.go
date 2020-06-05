package connections

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

func newKafkaConfigurations() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ChannelBufferSize = 1
	config.Version = sarama.V0_10_0_0
	return config
}

func newKafkaSyncProducer(server string) sarama.SyncProducer {
	kafka, err := sarama.NewSyncProducer([]string{server}, newKafkaConfigurations())

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	return kafka
}

//SendMessage - For sending message over kafka
func SendMessage(server string, topic string, event interface{}) error {

	fmt.Println(event)

	json, err := json.Marshal(event)

	if err != nil {
		return err
	}

	msgLog := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(json)),
	}

	partition, offset, err := newKafkaSyncProducer(server).SendMessage(msgLog)

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
	}

	fmt.Printf("Message: %+v\n", event)
	fmt.Printf("Message is stored in partition %d, offset %d\n", partition, offset)

	return nil
}
