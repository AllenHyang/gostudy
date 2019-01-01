package mqtt

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func Run() {
	fmt.Println("start listening")

	listen, err := net.Listen("tcp", ":8888")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		fmt.Println("connection established,time:",time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			fmt.Println("Accept error:", err)
			os.Exit(1)

		}
		go HandleConn(conn)
	}

}
func HandleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Connection lost %s \n",conn.RemoteAddr())
			} else {
				fmt.Println("read error:", err)
				//
				panic(1)
			}
			return
		}
		fmt.Printf("read reuslt ,recieve %d bytes, is:%s,from %s ,%s,\n", n, buf[:n], conn.RemoteAddr(),time.Now().Format("2006-01-02 15:04:05"))
	}
}
