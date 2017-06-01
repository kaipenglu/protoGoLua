package tcplib

import (
    "net"
)

type TcpConn struct {
    tcpServer *TcpServer
    conn *ComConn
}

func NewTcpConn(tcpServer *TcpServer, nativeConn *net.Conn) *TcpConn {
    return &TcpConn{tcpServer, &ComConn{nativeConn}}
}

func (c *TcpConn) Read () {
    go (*(c.conn)).Read()
}

func (c *TcpConn) Send(str string) {
    (*(c.conn)).Write(str)
}
