package consumers

import (
	"authentication/components"
	"authentication/config"
	"authentication/models"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/ilyakaznacheev/cleanenv"
)

//OtpConsumer -
func OtpConsumer() {

	var cfg config.KafkaConfig

	err := cleanenv.ReadConfig("application.yaml", &cfg)

	if err != nil {
		fmt.Println("Error in getting kafka configs :", err)
	}

	config := sarama.NewConfig()

	config.Consumer.Return.Errors = true

	// Create new consumer
	master, err := sarama.NewConsumer([]string{cfg.KafkaConfig.Consumer.Otp.Server}, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	// How to decide partition, is it fixed value...?
	consumer, err := master.ConsumePartition(cfg.KafkaConfig.Consumer.Otp.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
				otpProcessor(&msg.Value)
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")
}

func otpProcessor(otp *[]byte) {
	otpKafka := models.OtpKafka{}

	json.Unmarshal(*otp, &otpKafka)

	components.UpdateOtp(&otpKafka)
}
