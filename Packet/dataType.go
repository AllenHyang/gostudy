package Packet

import (
	"encoding/binary"
	"io"
)

func readToUint8(r io.Reader, order binary.ByteOrder) (int, error) {
	var data uint8
	err := binary.Read(r, order, &data)
	return int(data), err

}

func readToUint32(r io.Reader, order binary.ByteOrder) (int, error) {
	var data uint32
	err := binary.Read(r, order, &data)
	return int(data), err

}

func ReadFromLenBytes(length int, r io.Reader, order binary.ByteOrder) (int, error) {
	switch length {
	case 1:
		return readToUint8(r, order)
	case 2:
		return readToUint16(r,order)
	case 4:
		return readToUint32(r, order)
	}
	panic("length of bytes not defined")

}

func readToUint16(r io.Reader, order binary.ByteOrder) (int, error) {
	var data uint16
	err := binary.Read(r, order, &data)
	return int(data), err
}
