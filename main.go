package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("start up server")
	ln, err := bind_socket()
	if err != nil{
		return
	}
	network_manager := new(NetworkManager)
	network_manager.control_channels = make(map[string]chan string)
	//start_the_world()
	go func(){
		timer := time.NewTimer(time.Second * 10)
		<-timer.C
		fmt.Println("quit all channel")
		network_manager.QuitAllChannel()
	}()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept failed")
			continue
		}
		go HandleConnection(conn, network_manager)
	}
	//clean up the world
}
