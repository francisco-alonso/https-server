package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type AdapterImpl struct {
	client mqtt.Client
}

// NewAdapter returns a new instance of AdapterImpl.
func NewAdapter() *AdapterImpl {
	return &AdapterImpl{}
}

// Connect initializes the MQTT client and connects to the broker.
func (a *AdapterImpl) Connect(brokerURL, clientId string) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID(clientId)
	opts.SetCleanSession(true)
	opts.SetConnectTimeout(30 * time.Second)
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		fmt.Printf("MQTT connection lost: %v\n", err)
	}
	a.client = mqtt.NewClient(opts)
	token := a.client.Connect()
	token.Wait()
	
	return token.Error()
}

// Publish sends a message to the specified topic.
func (a *AdapterImpl) Publish(topic string, payload interface{}) error {
	token := a.client.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}

// Subscribe registers a callback function to handle messages from the specified topic.
func (a *AdapterImpl) Subscribe(topic string, callback mqtt.MessageHandler) error {
	token := a.client.Subscribe(topic, 0, callback)
	token.Wait()
	return token.Error()
}

// Disconnect cleanly disconnects the MQTT client.
func (a *AdapterImpl) Disconnect(quiesce uint) {
	a.client.Disconnect(quiesce)
}