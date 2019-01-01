package mqtt

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_packet(t *testing.T) {
	buff := bytes.NewBuffer([]byte{})
	msg := []byte("123abcdefg")

	packet := NewPacket(msg)
	packet.Header = []byte("hyMQTT")
	packet.Pack(buff)
	packet.Header = []byte("HYMQTT")

	packet.Pack(buff)
	//
	packet.Pack(buff)
	msg = []byte("123abcdefg")
	packet = NewPacket(msg)
	packet.Pack(buff)
	packet.payload = []byte("123")
	packet.Pack(buff)

	fmt.Println(buff.Bytes())
	Scanner(bytes.NewReader(buff.Bytes()))
	a := []byte("abc")
	b := []byte("abc")
	println(bytes.Equal(a, b))
}

func S() func() bool {
	a := 1
	f := func() bool {
		a += 1
		if a < 5 {
			return true
		} else {
			return false
		}
	}
	return f

}
