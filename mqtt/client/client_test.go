package client

import (
	"bytes"
	"fmt"
	"hybot/Packet"
	"hybot/mqttData"
	"strconv"
	"testing"
	"time"
)

func Test_client(t *testing.T) {
	client := Client{}
	err := client.Connect()
	if err != nil {
		return
	}
	data := []byte("abcdefghigklmn")
	MqttHeader := []byte("HYMQTT")

	packet := Packet.NewPacket(data, MqttHeader)
	buff := bytes.NewBuffer([]byte{})
	packet.Pack(buff)

	for i := 0; i < 5000; i++ {
		go client.Write(buff.Bytes())
		//time.Sleep(time.Nanosecond * 10)
	}
	time.Sleep(time.Second * 5)
	client.Close()
	<-make(chan struct{})
}

func Test_sub(t *testing.T) {
	for i := 0; i < 10; i++ {
		client := NewClient()
		client.InitNode("Test sub," + strconv.Itoa(i))
		//time.Sleep(time.Second)
		client.Sub("/AAA")
	}
	//pub test
	client := NewClient()
	client.InitNode("Pub")
	client.Pub("/AAA", []byte("测试发布"))

	<-make(chan struct{})

}

func Test_addUser(t *testing.T) {
	client := Client{}
	err := client.Connect()
	if err != nil {
		t.Error("Connect failed")
		return
	}
	//defer client.Close()
	mdata := mqttData.NewInitData("Test one")

	packet := Packet.NewPacket(mdata, mqttData.MQTTHeader)
	buff := bytes.NewBuffer([]byte{})
	packet.Pack(buff)
	client.Write(buff.Bytes())

	fmt.Println("finished")
	time.Sleep(time.Second * 5)

}
