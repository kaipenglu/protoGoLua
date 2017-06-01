package tcplib

import (
    "net"
)

type TcpClient struct {
    Host string
    Port string
    conn *ComConn
}

func (c *TcpClient) Start () {
    nativeConn, _ := net.Dial("tcp", c.Host + ":" + c.Port)
    c.conn = &ComConn{conn : &nativeConn}
    go (*(c.conn)).Read()
}

func (c *TcpClient) Send(str string) {
    (*(c.conn)).Write(str)
}

