package tcplib

import (
	"../codec"
	"net"
)

type TcpClient struct {
	host string
	port string
	conn *ComConn
	cdc  *codec.Codec
}

func NewTcpClient(host, port string, cdc *codec.Codec) (c *TcpClient) {
	c = &TcpClient{}
	c.host = host
	c.port = port
	c.cdc = cdc
	return
}

func (c *TcpClient) Start() {
	nativeConn, _ := net.Dial("tcp", c.host+":"+c.port)
	c.conn = &ComConn{conn: &nativeConn}
	c.conn.Read()
}

func (c *TcpClient) Send(i interface{}) {
	b := c.cdc.Encode(i)
    c.conn.Write(b)
}
