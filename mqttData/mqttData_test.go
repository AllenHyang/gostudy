package mqttData

import (
	"fmt"
	"testing"
)

func Test_data(t *testing.T) {

	a := NewPubData("123", []byte("ffff"))
	fmt.Println(a)
	c := Loads(a)
	fmt.Println(c.Data)

}
