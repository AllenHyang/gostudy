package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hybot/Packet"
	"testing"
)

func Test_binary_read(t *testing.T) {

	lengthBytes := []byte{9,1}
	lengthMsg, _:= Packet.ReadFromLenBytes(len(lengthBytes), bytes.NewReader(lengthBytes), binary.LittleEndian)
	fmt.Println(lengthMsg)
}

