package main

import (
	"fmt"
	"net"
)

const BUFF_LEN = 4096
func handleConnection(c net.Conn){
	buff := make([]byte, BUFF_LEN)
	for{
		count, err := c.Read(buff);
		if err != nil{
			fmt.Println(err)
			return
		}
		if count != 0{
			fmt.Println(string(buff))
			c.Write(buff)
		}
	}
}

func main(){
	fmt.Println("start up server")
	ln, err := net.Listen("tcp", ":19840")
	if err != nil {
		fmt.Println("bind failed")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept failed")
			continue
		}
		go handleConnection(conn)
	}
}
