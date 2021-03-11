package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)
func main(){
	li,err := net.Listen("tcp",":8080")
	if err != nil{
		log.Print(err)
	}
	defer li.Close()
    for {
		conn,err := li.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConnection(conn)
	}
	
}
func handleConnection(conn net.Conn){
	err := conn.SetDeadline(time.Now().Add(20*time.Second))
	if err != nil{
		log.Println("time OUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintln(conn,"read",ln)

	}
	defer conn.Close()
}