package handleMsg

import (
	"hybot/mqttData"
	"net"
)

type remoteUser struct {
	name string
	subs map[mqttData.Topic]bool
	addr net.Addr
	conn net.Conn
}

func NewRemoteUser(name string, ip net.Addr, conn net.Conn) *remoteUser {
	u := remoteUser{name: name, addr: ip, conn: conn}
	u.subs = make(map[mqttData.Topic]bool)
	return &u
}

func (u *remoteUser) GetName() string {
	return string(u.name)
}

func (u *remoteUser) AddSub(topic mqttData.Topic) {
	u.subs[topic] = true
}

func (u *remoteUser) RemoveSub(topic mqttData.Topic) {
	delete(u.subs, topic)
}

func (u *remoteUser) IsSubbed(topic mqttData.Topic) bool {
	_, ok := u.subs[topic]
	return ok
}
