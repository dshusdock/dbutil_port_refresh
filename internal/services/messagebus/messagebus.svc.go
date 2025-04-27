package messagebus

import (
	"log"

	"github.com/asaskevich/EventBus"
)

// Service is the message bus service

type MessageBusSvc struct {
	Bus EventBus.Bus
}

var MBus* MessageBusSvc

func init() {
	// init the message bus service
	log.Println("Initializing message bus service")
	MBus = &MessageBusSvc{}
	MBus.Bus = EventBus.New()
}

func GetBus() EventBus.Bus {
	return MBus.Bus
}	
