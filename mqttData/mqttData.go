package mqttData

type MqttData struct {
	serialize Serialize
	Data      map[string]string
}

var MQTTHeader = []byte("HYMQTT")

func NewMqttData() *MqttData {
	m := new(MqttData)
	m.serialize = new(JsonSerialize)
	m.Data = make(map[string]string)
	return m
}

func NewResponseData(topic Topic, msg []byte) []byte {
	m := NewMqttData()
	m.Data["topic"] = string(topic)
	m.Data["value"] = string(msg)
	return m.serialize.ToBytes(m.Data)
}

func NewPubData(topic Topic, msg []byte) []byte {
	m := NewMqttData()
	m.Data["cmd"] = "pub"
	m.Data["topic"] = string(topic)
	m.Data["value"] = string(msg)
	return m.serialize.ToBytes(m.Data)
}
func NewSubData(topic Topic) []byte {
	m := NewMqttData()
	m.Data["cmd"] = "sub"
	m.Data["topic"] = string(topic)
	return m.serialize.ToBytes(m.Data)
}

func NewInitData(name string) []byte {
	m := NewMqttData()
	m.Data["cmd"] = string("initNode")
	m.Data["value"] = string(name)
	return m.serialize.ToBytes(m.Data)
}

func Loads(b []byte) *MqttData {
	m := NewMqttData()
	m.serialize.Loads(b, &m.Data)
	return m
}
