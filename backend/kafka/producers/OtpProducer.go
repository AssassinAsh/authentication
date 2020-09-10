package producers

import (
	"authentication/config"
	"authentication/connections"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

//OtpProducer - Kafka Otp Producer
func OtpProducer(event interface{}) {
	log.Println("Otp Producer Running...")
	var cfg config.KafkaConfig

	err := cleanenv.ReadConfig("application.yaml", &cfg)

	if err != nil {
		fmt.Println("Error in getting kafka configs :", err)
	}

	log.Println("Sending Message")
	err = connections.SendMessage(cfg.KafkaConfig.Producer.Otp.Server, cfg.KafkaConfig.Producer.Otp.Topic, event)

	if err != nil {
		log.Panic("Error in Marshalling", err)
	}

}
