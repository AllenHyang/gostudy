package Packet

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Packet struct {
	Header    []byte
	lengthMsg uint32
	payload   []byte
}

func (p *Packet) Pack(buff io.Writer) error {
	var err error
	err = binary.Write(buff, binary.BigEndian, &p.Header)
	err = binary.Write(buff, binary.BigEndian, &p.lengthMsg)
	err = binary.Write(buff, binary.BigEndian, &p.payload)

	return err
}
func NewPacket(msg []byte, header []byte) *Packet {

	p := Packet{[]byte(header), uint32(len(msg)), msg}

	return &p
}

func PackToBytes(msg []byte, header []byte) []byte {
	p := NewPacket(msg, header)
	buff := bytes.NewBuffer([]byte{})
	_ = p.Pack(buff)
	return buff.Bytes()
}

type Consume interface {
	Apply(m []byte)
}

func Scanner(buffer *bytes.Buffer, Consume Consume, header []byte) {
	//log.Printf("buffer,%s \n", buffer.Bytes())

	if buffer.Len() < len(header)+4 { //不够header长度，返回
		return
	}
	for buffer.Len() > 0 {
		b, err := buffer.ReadByte()
		if err != io.EOF && b == header[0] { //找到Header开始
			headerLefts := buffer.Next(len(header) - 1)       //header 读出来
			if bytes.Equal(headerLefts, []byte(header[1:])) { //是否为header
				if buffer.Len() >= 4 { //够length的长度
					lengthBytes := buffer.Next(4) //读长度
					var lengthMsg uint32
					_ = binary.Read(bytes.NewReader(lengthBytes), binary.BigEndian, &lengthMsg)
					if buffer.Len() >= int(lengthMsg) { //长度够，读出来，不够返回
						msg := buffer.Next(int(lengthMsg))
						go Consume.Apply(msg)
					} else {
						//长度不够，需要把header length重新写入
						leftBytes := buffer.Bytes()
						buffer.Reset()
						packet := Packet{header, lengthMsg, leftBytes}
						_ = packet.Pack(buffer)
						return
					}
				} else { //不够length的长度，把header 写回去
					leftBytes := buffer.Bytes()
					buffer.Reset()
					buffer.Write(header)
					buffer.Write(leftBytes)
					return

				}

			}
		}
	}
}
