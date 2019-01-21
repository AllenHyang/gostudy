package hyimu

import (
	"bytes"
	"encoding/binary"
	"hybot/Packet"
	"log"
	"math"
)

type IMU struct {
	rawData []byte  // id1 data id2 data
	info Info
}

func (d *IMU) parseBytes(b []byte) int {
	if len(b) != 2 {
		panic("Input b size should be 2")
	}
	value, err := Packet.ReadFromLenBytes(2, bytes.NewReader(b), binary.LittleEndian)
	if err != nil {
		panic(err)
	}
	return value
}

func (d *IMU) calAngular() float32 {
	value := d.parseBytes(d.rawData[5:6])
	return float32(value) * math.Pi / 1800
}

func (d *IMU) calXMagneticIntense() float32 {
	return 0
}

func (d *IMU) Apply(m []byte) {
	log.Printf("consuming,% X\n", m)

}

var SerialStruct = Packet.DataStruct{Header: []byte{0xaa, 0x00}, BytesOfLength: 1}
