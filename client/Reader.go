package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Read(conn net.Conn){
for{
	message,err:=bufio.NewReader(conn).ReadString('\n')
	if err!=nil{
		fmt.Println(err)
		continue
	}
	if strings.HasPrefix(message,"print#"){
		fmt.Print(message[6:])
	}
}
}