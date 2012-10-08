package main

import (
	"net"
	"fmt"
	"io"
)

const BUFF_LEN = 4096
func HandleConnection(c net.Conn){
	fmt.Println("Connection from: ", c.RemoteAddr())
	buff := make([]byte, BUFF_LEN)
	for{
		count, err := c.Read(buff);
		if err != nil{
			if err == io.EOF{
				fmt.Println("client disconnect")
				return
			}
			fmt.Println(err)
			return
		}
		if count != 0{
			fmt.Print(string(buff))
			c.Write(buff)
		}
	}
}

