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
connection, err:=net.Dial("tcp","localhost:8080")
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
		fmt.Println("enter message:")
		for  {
			msg,_:=bufio.NewReader(os.Stdin).ReadString('\n')
			if msg == "exit" {
				break
			}
			fmt.Fprintf(connection,"send#"+msg[:len(msg)-1]+"#"+dst )
		}
	}


}
