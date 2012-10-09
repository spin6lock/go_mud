package main

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("start up server")
	ln, err := bind_socket()
	if err != nil{
		return
	}
	//network_manager()
	//start_the_world()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept failed")
			continue
		}
		go HandleConnection(conn)
	}
	//clean up the world
}
