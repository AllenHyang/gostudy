package hymqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
)

var server = "tcp://127.0.0.1:1883"

type HYClient interface {
	Publish(topic string, payload interface{})
	Subscribe(topic string, callback MQTT.MessageHandler)
	Disconnect()
}

type hyclient struct {
	mqtttClient MQTT.Client
}

func InitNode(clientID string) *hyclient {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(server)
	opts.SetClientID(clientID)
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return &hyclient{client}
}

func (c *hyclient) Publish(topic string, payload interface{}) {
	//func (c *client) Publish(topic string, qos byte, retained bool, payload interface{}) Token {
	//token:= c.mqtttClient.Publish(topic,byte(0),false,payload)
	if token := c.mqtttClient.Publish(topic, byte(0), false, payload); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

//func (c *client) Subscribe(topic string, qos byte, callback MessageHandler) Token {
func (c *hyclient) Subscribe(topic string, callback MQTT.MessageHandler) {

	if token := c.mqtttClient.Subscribe(topic, byte(0), callback); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func (c *hyclient) Disconnect() {
	//func (c *client) Disconnect(quiesce uint) {
	c.mqtttClient.Disconnect(250)

}
