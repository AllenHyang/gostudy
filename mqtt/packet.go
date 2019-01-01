package mqtt

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	MsgHeader = "HYMQTT"
)

type Packet struct {
	Header    []byte
	lengthMsg int32
	payload   []byte
}

func (p *Packet) Pack(buff io.Writer) error {
	var err error
	err = binary.Write(buff, binary.BigEndian, &p.Header)
	err = binary.Write(buff, binary.BigEndian, &p.lengthMsg)
	err = binary.Write(buff, binary.BigEndian, &p.payload)

	return err
}
func NewPacket(msg []byte) *Packet {

	p := Packet{[]byte(MsgHeader), int32(len(msg)), msg}

	return &p

}

func Unpack(reader io.Reader) (*Packet, error) {
	var err error
	p := &Packet{Header: make([]byte, 6), lengthMsg: int32(0)}
	err = binary.Read(reader, binary.BigEndian, &p.Header)
	err = binary.Read(reader, binary.BigEndian, &p.lengthMsg)

	p.payload = make([]byte, p.lengthMsg)
	err = binary.Read(reader, binary.BigEndian, &p.payload)
	return p, err

}
func Scanner(reader io.Reader) {
	s := bufio.NewScanner(reader)
	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		header := []byte(MsgHeader)
		if !atEOF && data[0] == MsgHeader[0] && bytes.Equal(data[:6], header) {
			if len(data) > 6+4 {
				length := int32(0)
				binary.Read(bytes.NewReader(data[6:6+4]), binary.BigEndian, &length)
				if int(length)+6+4 <= len(data) {
					return int(length) + 6 + 4, data[:int(length)+6+4], nil
				}
			}
		}
		return 1, []byte{}, nil
	})
	for s.Scan() {
		if len(s.Bytes()) > 0 {
			msg, err := Unpack(bytes.NewReader(s.Bytes()))
			if err != nil {
				fmt.Println("error,", err)
				return
			}
			fmt.Println(msg)
		}
	}
}
