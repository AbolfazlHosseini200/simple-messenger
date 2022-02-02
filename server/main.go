package main

import (
	"bufio"
	"fmt"
	"net"
)
var clients []Client
type Client struct {
	name string
	con net.Conn
}
func main() {
	connection,err:=net.Listen("tcp","localhost:8080")
	if err!=nil{
		fmt.Println(err)
	}

	defer connection.Close()
	fmt.Println("Listening on port 8080")
	for{
		c,err:=connection.Accept()
		if err!=nil{
			fmt.Println(err)
		}
		var test Client
		test.con = c
		test.name,err=bufio.NewReader(c).ReadString('\n')
		if err!=nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Client ",c.RemoteAddr()," Connected")
		clients=append(clients,test)
		go Handle(c,test)
	}
}