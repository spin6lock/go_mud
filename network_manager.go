package main

import (
	"fmt"
	"net"
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

func (this *NetworkManager) QuitAllChannel(){
	for k, v := range this.control_channels {
		fmt.Println(k)
		v <- "quit" //block here
	}
}

