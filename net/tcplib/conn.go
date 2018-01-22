package tcplib

import (
	"../codec"
	"net"
)

type TcpConn struct {
	tcpServer *TcpServer
	conn      *ComConn
	cdc       *codec.Codec
}

func NewTcpConn(tcpServer *TcpServer, nativeConn *net.Conn) *TcpConn {
	c := &TcpConn{}
	c.tcpServer = tcpServer
	c.cdc = tcpServer.cdc
	c.conn = &ComConn{nativeConn, c.cdc}
	return c
}

func (c *TcpConn) Read() {
	c.conn.Read()
}

func (c *TcpConn) Send(i interface{}) {
	c.conn.Send(i)
}
