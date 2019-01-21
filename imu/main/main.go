package main

import (
	"bytes"
	"hybot/Packet"
	"hybot/imu/hyimu"
	"hybot/serial"
	"log"
)

func main() {


	s, err := hyserial.OpenSerial("com1", 115200, 0)
	if err != nil {
		return
	}
	buff := bytes.NewBuffer([]byte{})

	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		buff.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("% X\n", buff.Bytes())
		imu := new(hyimu.IMU)
		Packet.Scanner(1, 1, 0)

	}
}
