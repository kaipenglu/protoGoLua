package main

import (
	"./net/codec"
	"./net/tcplib"
	"./pb/pbcpl"
	"fmt"
)

func f(msgId uint32, msg interface{}) {
	fmt.Println(msgId, "#", msg)
}

func main() {
	cdc := codec.NewCodec()

	cdc.RegisterProto(1, &pbcpl.CSLogin{}, f)

	server := tcplib.NewTcpServer("127.0.0.1", "3456", cdc)
	server.Start()

	client := tcplib.NewTcpClient("127.0.0.1", "3456", cdc)
	client.Start()

	p := &pbcpl.CSLogin{}
	cmd := pbcpl.PACKET_ID_PACKET_CS_LOGIN_CSLogin
	p.Cmd = &cmd
	accountName := "jack"
	p.AccountName = &accountName
	password := "123456"
	p.Password = &password
	client.Send(p)

	endRunning := make(chan bool, 1)
	<-endRunning
}
