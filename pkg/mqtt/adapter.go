package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client interface {
	Connect(brokerURL, clientID string) error
	Publish(topic string, payload interface{}) error
	Subscribe(topic string, callback mqtt.MessageHandler) error
	Disconnect(quiesce int)
}
