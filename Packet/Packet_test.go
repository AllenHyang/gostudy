package Packet

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func Test_packet(t *testing.T) {
	buff := bytes.NewBuffer([]byte{})
	MqttHeader := []byte("HYMQTT")
	msg := []byte("abcdefghigklmn")
	packet := NewPacket(msg, MqttHeader)
	buff.Write([]byte{1, 2, 3})
	var i int
	for i < 1000 {
		packet.Pack(buff)
		i++
	}
	fmt.Println(buff.Bytes())
	f := func(msg []byte) { fmt.Printf("in consuming,%s \n", msg) }
	bf := bytes.NewBuffer(buff.Bytes())
	Scanner(bf, f, MqttHeader)

	Scanner(bf, f, MqttHeader)
	time.Sleep(time.Second * 5)
	fmt.Printf("%s", bf.Bytes())

}
