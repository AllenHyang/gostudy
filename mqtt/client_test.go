package mqtt

import (
	"testing"
)

func Test_client(t *testing.T) {
	client := Client{}
	err := client.Connect()
	if err != nil {
		return
	}
	data := []byte("abcdefghigklmn")
	for i := 1; i < 5000; i++ {
		go client.Write(data)
		//time.Sleep(time.Nanosecond * 10)
	}
	<-make(chan struct{})
}
