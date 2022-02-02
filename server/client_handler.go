package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)
var clientInfo Client
func Handle(c net.Conn,client Client)  {
clientInfo =client
for {
	reader,err:=bufio.NewReader(c).ReadString('\n')
	if strings.HasPrefix(reader,"send#"){
	splited:=strings.Split(reader,"#")
	Send(splited[2],splited[1],clientInfo.name)
	}
	if err!=nil {
		fmt.Println(err)
		continue
	}
}
}
func Send(dest string,s string,name string)  {
	cc:=nameToCon(dest)
	if cc==nil {
		fmt.Fprintf(clientInfo.con,"print#"+name[:len(name)-1]+":"+s+"\n")
		return
	}
	fmt.Fprintf(cc,"print#"+name[:len(name)-1]+":"+s+"\n")
}
func nameToCon(name string)net.Conn{
	for i:=0;i<len(clients);i++ {
		if clients[i].name==name{
			return clients[i].con
		}
	}
	return nil
}