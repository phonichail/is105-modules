package main

import (
	"fmt"
	"net"
)

func handler(c net.Conn) {
	c.Write([]byte("Ole Bjørnar Granås"))
	fmt.Println(c.RemoteAddr())
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
