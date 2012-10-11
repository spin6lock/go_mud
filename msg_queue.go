package main

import (
	"net"
	"fmt"
	"io"
)

type NetworkManager struct{
	online_count int //do i need mutex to protect online count?
	control_channels map[string]chan string
}

func (this *NetworkManager) RegisterControlChannel(c net.Conn) chan string{
	this.online_count++
	key := c.RemoteAddr().String()
	this.control_channels[key] = make(chan string)
	return this.control_channels[key]
}

func (this *NetworkManager) UnregisterControlChannel(c net.Conn){
	this.online_count--
	key := c.RemoteAddr().String()
	delete(this.control_channels, key)
}

const BUFF_LEN = 4096
func HandleConnection(c net.Conn, manager *NetworkManager){
	fmt.Println("Connection from: ", c.RemoteAddr())
	quit := manager.RegisterControlChannel(c)
	buff := make([]byte, BUFF_LEN)
	for{
		select{
			case <-quit:
				break
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

