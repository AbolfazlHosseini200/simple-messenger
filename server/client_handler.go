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
		Send(splited[2],splited[1],splited[3])
	}
	if strings.HasPrefix(reader,"refresh#"){
		splited:=strings.Split(reader,"#")
		Update(splited[1])
	}
	if err!=nil {
		fmt.Println(err)
		continue
	}
}
}

func Update(sender string) {
	msg:="All clients:"
	for i:=0;i<len(clients);i++ {
		msg+=clients[i].name[:len(clients[i].name)-1]+"/"
	}
	Send(sender[:len(sender)-1],msg," ")
}
func Send(dest string,s string,name string)  {
	cc:=nameToCon(dest)
	if cc==nil {
		cc=nameToCon(name)
		fmt.Fprintf(cc,"print#thers no such user"+"\n")
		return
	}
	fmt.Fprintf(cc,"print#"+name[:len(name)-1]+":"+s+"\n")
}
func nameToCon(name string)net.Conn{
	for i:=0;i<len(clients);i++ {
		if clients[i].name[:len(clients[i].name)-1]==name{
			return clients[i].con
		}
	}
	return nil
}