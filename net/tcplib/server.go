package tcplib

import (
	"../codec"
	"fmt"
	"net"
)

type TcpServer struct {
	host    string
	port    string
	connMap map[string]*TcpConn
	cdc     *codec.Codec
}

func NewTcpServer(host, port string, cdc *codec.Codec) *TcpServer {
	srv := &TcpServer{}
	srv.host = host
	srv.port = port
	srv.connMap = make(map[string]*TcpConn)
	srv.cdc = cdc
	return srv
}

func (srv *TcpServer) Start() {
	ln, _ := net.Listen("tcp", srv.host+":"+srv.port)

	go func() {
		for {
			nativeConn, _ := ln.Accept()
			tcpConn := NewTcpConn(srv, &nativeConn)
			addr := (*(*(tcpConn.conn)).conn).RemoteAddr().String()
			srv.connMap[addr] = tcpConn
			tcpConn.Read()
			fmt.Println("A client connected : " + addr)
		}
	}()
}

func (srv *TcpServer) Broadcast(i interface{}) {
	for _, v := range srv.connMap {
		v.Send(i)
	}
}
