package hyserial

import (
	"github.com/tarm/goserial"
	"io"
	"time"
)

type SerialPort io.ReadWriteCloser

type HYSerial struct {
	port SerialPort
}

func OpenSerial(com string, baud int, readTimeout time.Duration) (*HYSerial, error) {
	s := new(HYSerial)
	config := serial.Config{Name: com, Baud: baud, ReadTimeout: readTimeout}
	port, err := serial.OpenPort(&config)
	s.port = port
	return s, err
}

func (hys *HYSerial)Read(p []byte) (n int, err error){
	return hys.port.Read(p)

}

func (hys *HYSerial)Write(p []byte) (n int, err error){
	return hys.port.Write(p)

}
func (hys *HYSerial)Close() ( err error){
	return hys.port.Close()

}