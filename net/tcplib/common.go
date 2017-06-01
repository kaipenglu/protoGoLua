package tcplib

import (
    "net"
    "bufio"
    "fmt"
)

type ComConn struct {
    conn *net.Conn
}

func (c *ComConn) Read() {
    reader := bufio.NewReader(*(c.conn))
    for {
        message, err := reader.ReadString('\n')

        if err != nil {
            (*(c.conn)).Close()
            return
        }

        fmt.Println("A new Message : " + message)
    }
}

func (c *ComConn) Write(str string) {
    (*(c.conn)).Write([]byte(str + "\n"))
}
