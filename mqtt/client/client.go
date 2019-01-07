package client

import (
	"fmt"
	"hybot/Packet"
	"hybot/mqttData"
	"net"
	"time"
)

type Client struct {
	conn *net.Conn
}

func NewClient() *Client {
	client := Client{}
	_ = client.Connect()
	return &client
}

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("connect error:", err)
		return err
	}
	c.conn = &conn
	return nil
}

func (c *Client) InitNode(name string) {
	m := mqttData.NewInitData(name)
	b := Packet.PackToBytes(m, mqttData.MQTTHeader)
	_, err := c.Write(b)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	time.Sleep(time.Millisecond * 100)
}

func (c *Client) Sub(topic mqttData.Topic) {
	mdata := mqttData.NewSubData(topic)
	b := Packet.PackToBytes(mdata, mqttData.MQTTHeader)
	_, err := c.Write(b)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func (c *Client) Pub(topic mqttData.Topic, payload []byte) {
	mdata := mqttData.NewPubData(topic, payload)
	b := Packet.PackToBytes(mdata, mqttData.MQTTHeader)
	_, err := c.Write(b)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func (c *Client) Close() error {
	return (*c.conn).Close()
}
func (c *Client) Write(data []byte) (int, error) {
	return (*c.conn).Write(data)
}
