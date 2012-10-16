package main

import (
	"net"
	"fmt"
	"io"
)


const BUFF_LEN = 4096
func HandleConnection(c net.Conn, manager *NetworkManager){
	defer c.Close()
	fmt.Println("Connection from: ", c.RemoteAddr())
	quit := manager.RegisterControlChannel(c)
	buff := make([]byte, BUFF_LEN)
	loop:
	for{
		select{
			case <-quit:
				break loop
			default:
				count, err := c.Read(buff)
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
	manager.UnregisterControlChannel(c)
}

func bind_socket() (net.Listener, error){
	ln, err := net.Listen("tcp", ":19840")
	if err != nil {
		fmt.Println("bind failed")
	}
	return ln, err
}

