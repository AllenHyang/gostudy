package handleMsg

import (
	"net"
	"strconv"
	"testing"
	"time"
)

func Test_msgCenter(t *testing.T) {
	msgCenter := GetMsgCenterIns()
	users := GenUsers()
	for _, u := range users {
		msgCenter.AddUser(u)
		msgCenter.Sub(u, "/Test")
	}

	msg := Msg{"/Test", []byte("hello world")}
	msgCenter.Pub(msg)
	time.Sleep(time.Second)
}

func Test_msgCenterSingleton(t *testing.T) {
	m1 := GetMsgCenterIns()
	m2 := GetMsgCenterIns()
	if m1 != m2 {
		t.Error("Not equal")
	}
}

func GenUsers() []*remoteUser {
	users := make([]*remoteUser, 10)
	for i := 0; i < 10; i++ {
		ip := net.IPNet{IP: []byte("192.168.0." + strconv.Itoa(i)), Mask: []byte("255.255.255.255")}
		name := "test " + strconv.Itoa(i)
		u := NewRemoteUser(name, &ip)
		users[i] = u
	}
	return users
}
