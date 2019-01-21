package hymqtt

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"testing"
)

func TestInitNode(t *testing.T) {

	client := InitNode("test client")
	client.Publish("/test", "helloworld")

}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}
func TestSub(t *testing.T) {

	client := InitNode("test client")
	client.Subscribe("/test", onMessageReceived)

	c := make(chan os.Signal, 1)
	<-c
}
