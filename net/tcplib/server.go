package tcplib

import (
    "fmt"
    "net"
)

type TcpServer struct {
    Host string
    Port string
    connMap map[string] *TcpConn
    // protoMap map[in32] int32
}

func (srv *TcpServer) Start() {
    srv.connMap = make(map[string] *TcpConn)
    ln, _ := net.Listen("tcp", srv.Host + ":" + srv.Port)

    go func() {
        for {
            nativeConn, _ := ln.Accept()
            tcpConn := NewTcpConn(srv, &nativeConn)
            addr := (*(*(tcpConn.conn)).conn).RemoteAddr().String()
            srv.connMap[addr] = tcpConn
            go tcpConn.Read()
            fmt.Println("A client connected : " + addr)
        }
    } ()
}

func (srv *TcpServer) Broadcast(str string) {
    for _, v := range srv.connMap {
        v.Send(str)
    }
}
