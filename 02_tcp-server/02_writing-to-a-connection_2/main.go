package main

import (
	"fmt"
	"log"
	"net"
)
func main(){
	li,err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Panic(err)
	}
	defer li.Close()
	for{
		conn,err := li.Accept()
		if err != nil{
			log.Panic(err)
		}
		b :=[]byte{'h','e','l','l','o'}
		n,err := conn.Write(b)
		s := fmt.Sprintln("\nno. of bytes written",n)
		fmt.Fprintln(conn,s)
		conn.Close()
	
	}
}