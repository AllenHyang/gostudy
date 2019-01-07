package handleMsg

import (
	"hybot/mqttData"
	"log"
	"net"
)

type HandleMsg struct {
	conn      net.Conn
	msgCenter *MessageCenter
	user      *remoteUser
}

func NewHandleMsg(conn net.Conn) *HandleMsg {
	h := HandleMsg{conn: conn, msgCenter: GetMsgCenterIns(), user: nil}
	return &h
}

func (h *HandleMsg) isUserInitialed() bool {
	if h.user == nil {
		return false
	}
	return true
}
func (h *HandleMsg) Apply(m []byte) {
	mqttdata := mqttData.Loads(m)
	cmd := string(mqttdata.Data["cmd"])
	switch cmd {
	case "pub":
		msg := Msg{topic: mqttData.Topic(mqttdata.Data["topic"]), payload: []byte(mqttdata.Data["value"])}
		h.msgCenter.Pub(msg)
		//h.conn.Write([]byte("0,Pubbed"))

	case "sub":
		if !h.isUserInitialed() {
			log.Println("Node is not initialed,Please initNode first!")
			//h.conn.Write([]byte("-1,Node is not initialed,Please initNode first!"))

			return
		}
		h.msgCenter.Sub(h.user, mqttData.Topic(mqttdata.Data["topic"]))
		//h.conn.Write([]byte("0,Subbed"))

	case "initNode":
		name := mqttdata.Data["value"]
		u := NewRemoteUser(string(name), h.conn.RemoteAddr(), h.conn)
		h.user = u
		_ = h.msgCenter.AddUser(u)
		h.conn.Write([]byte("0,Connected"))
	default:
		log.Println("Func not found")
	}

}
