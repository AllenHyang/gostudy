package server

import (
	"bytes"
	"hybot/Packet"
	"hybot/mqtt/handleMsg"
	"hybot/mqttData"
	"io"
	"log"
	"net"
	"os"
	"runtime/debug"
	"time"
)



func Run() {
	log.Println("Start")
	log.Println("start listening")
	listen, err := net.Listen("tcp", ":9999")

	defer listen.Close()
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		log.Println("connection established,time:", time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Println("Accept error:", err)
			os.Exit(1)
		}

		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	//buf := bytes.NewBuffer([]byte{})
	dChan := make(chan []byte, 10)
	ctlChan := make(chan bool)

	go WaitData(dChan, handleMsg.NewHandleMsg(conn), ctlChan)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("Connection lost %s \n", conn.RemoteAddr())
			} else {
				log.Println("read error, the client may not be proper closed.", err)
				//os.Exit(1)

			}
			ctlChan <- true

			_ = handleMsg.GetMsgCenterIns().RemoveUserByAddr(conn.RemoteAddr())

			return
		}
		dChan <- buf[:n]
	}
}

func WaitData(msgChan chan []byte, Consume Packet.Consume, quitChan chan bool) {
	buff := bytes.NewBuffer([]byte{})

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error,PANIC,==================,", err, buff.Bytes())
			debug.PrintStack()
			WaitData(msgChan, Consume, quitChan)
		}
	}()

	for {
		select {
		case <-quitChan:
			log.Println("quit WaitData,")
			return
		case b := <-msgChan:

			buff.Write(b)
			Packet.Scanner(buff, Consume, mqttData.MQTTDataStruct) //TODO: 大量数据发来的时候，会触发panic,json unmarshel错误
			continue
		}
	}
}
