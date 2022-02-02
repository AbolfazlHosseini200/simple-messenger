package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
fmt.Println("trying to connect to server in a second...")
time.Sleep(time.Second)
connection, err:=net.Dial("tcp","2.tcp.ngrok.io:15934")
if err!=nil{
	fmt.Println(err)
	return
}
fmt.Println("connected to server successfully")
fmt.Println("enter your name:")
name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Fprintf(connection,name + "\n")
go Read(connection)
	for  {
		fmt.Println("enter destination:")
		dst,_:=bufio.NewReader(os.Stdin).ReadString('\n')
		dst = dst[:len(dst)-1]
		fmt.Println("enter message:")
		for  {
			msg,_:=bufio.NewReader(os.Stdin).ReadString('\n')
			if msg[:len(msg)-1] == "exit" {
				break
			}else if msg[:len(msg)-1] == "refresh"{
				fmt.Fprintf(connection,"refresh#"+name)
				break
			}
			fmt.Fprintf(connection,"send#"+msg[:len(msg)-1]+"#"+dst+"#" +name)
		}
	}


}
