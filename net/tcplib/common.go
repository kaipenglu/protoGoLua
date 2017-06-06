package tcplib

import (
	"../codec"
	"bufio"
	"encoding/binary"
	"net"
)

type ComConn struct {
	conn *net.Conn
	cdc  *codec.Codec
}

func (c *ComConn) Read() {
	reader := bufio.NewReader(*(c.conn))
	scanner := bufio.NewScanner(reader)

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		la := uint32(len(data))
		if la < 4 {
			return 0, nil, nil
		}
		lc := binary.LittleEndian.Uint32(data[:4]) + 4
		if la >= lc {
			return int(lc), data[4:lc], nil
		}
		if atEOF {
			return 0, data, bufio.ErrFinalToken
		}
		return
	}

	scanner.Split(split)
	go func() {
		for scanner.Scan() {
			c.cdc.Decode(scanner.Bytes())
		}
	}()
}

func (c *ComConn) Write(b []byte) {
    if len(b) == 0 {
        return
    }
	a := make([]byte, 4)
	binary.LittleEndian.PutUint32(a, uint32(len(b)))
	(*(c.conn)).Write(a)
	(*(c.conn)).Write(b)
}
