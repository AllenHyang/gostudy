package mqtt

import (
	"fmt"
	"net"
)

type Client struct {
	conn *net.Conn
}


func (c *Client) Connect()  error {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("connect error:", err)
		return  err
	}
	c.conn = &conn
	return nil
}

func (c *Client) Close()error {
	return (*c.conn).Close()
}
func (c *Client) Write(data []byte) (int, error) {
	return (*c.conn).Write(data)
}
