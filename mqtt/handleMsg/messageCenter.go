package handleMsg

import (
	"errors"
	"fmt"
	"hybot/Packet"
	"hybot/mqttData"
	"log"
	"net"
	"sync"
)

var ins *MessageCenter
var once sync.Once

var UserNotExistError = errors.New("UserNotExistError")
var UserExistError = errors.New("UserExistError")

type MessageCenter struct {
	user map[string]*remoteUser
}

type Msg struct {
	topic   mqttData.Topic
	payload []byte
}

func GetMsgCenterIns() *MessageCenter {
	once.Do(func() {
		ins = &MessageCenter{}
		ins.user = make(map[string]*remoteUser)
	})

	return ins
}

func (m *MessageCenter) AddUser(u *remoteUser) error {
	log.Printf("Add user,%s \n", u.GetName())
	if _, ok := m.user[u.GetName()]; ok {
		return UserExistError
	}
	m.user[u.GetName()] = u
	return nil
}

func (m *MessageCenter) RemoveUser(u *remoteUser) {
	log.Printf(" Remove user,%s \n", u.GetName())

	delete(m.user, u.GetName())
}

func (m *MessageCenter) RemoveUserByAddr(addr net.Addr) error {
	u, err := m.GetUserByAddr(addr)
	if err != nil {
		return err
	}
	m.RemoveUser(u)
	return nil
}

func (m *MessageCenter) GetUserByAddr(addr net.Addr) (*remoteUser, error) {
	for _, u := range m.user {
		if u.addr == addr {
			return u, nil
		}
	}
	return nil, UserNotExistError

}
func (m *MessageCenter) Sub(u *remoteUser, topic mqttData.Topic) {
	m.user[u.GetName()] = u
	u.AddSub(topic)
}

func send(u remoteUser, msg Msg) {
	log.Printf("Sending msg: %s, To addr: %s \n", msg.payload, u.addr.String())
	m := mqttData.NewResponseData(msg.topic, msg.payload)
	b := Packet.PackToBytes(m, mqttData.MQTTHeader)
	_, _ = u.conn.Write(b)
}

func (m *MessageCenter) Pub(msg Msg) {
	for _, u := range m.user {
		fmt.Println(u.addr, u.subs, u.conn.RemoteAddr())

		if u.IsSubbed(msg.topic) {
			//转发消息
			go send(*u, msg)
		}
	}
}
