package main

import "./net/tcplib"

func main() {
    server := &tcplib.TcpServer{Host : "127.0.0.1", Port : "3456"};
    server.Start();

    client1 := &tcplib.TcpClient{Host : "127.0.0.1", Port : "3456"};
    client1.Start();

    client2 := &tcplib.TcpClient{Host : "127.0.0.1", Port : "3456"};
    client2.Start();

    server.Broadcast("Hello, clients!")
    client1.Send("Hello, Server!")
    client2.Send("Hello, Server!")

    endRunning := make(chan bool, 1)
    <-endRunning
}
