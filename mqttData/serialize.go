package mqttData

import (
	"encoding/json"
	"fmt"
)

type Serialize interface {
	ToBytes(data interface{}) []byte
	Loads([]byte, interface{})
}

type JsonSerialize struct{}

func (m *JsonSerialize) ToBytes(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		panic("Json marshal to byte failed")
	}
	return b
}

func (m *JsonSerialize) Loads(b []byte, target interface{}) {

	err := json.Unmarshal(b, target)
	if err != nil {
		fmt.Println(err)
		panic("Json unmarshal to MqttData error ")
	}

}
