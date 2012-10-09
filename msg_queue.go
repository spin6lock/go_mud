package main

import (
	"net"
	"fmt"
	"io"
)

type NetworkManager struct{
	online_count int;
	//ref to every conn queue, in order to kill by master
}

const BUFF_LEN = 4096
//TODO add control channel
func HandleConnection(c net.Conn){
	fmt.Println("Connection from: ", c.RemoteAddr())
	//join connection manager
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

func bind_socket() (net.Listener, error){
	ln, err := net.Listen("tcp", ":19840")
	if err != nil {
		fmt.Println("bind failed")
	}
	return ln, err
}

