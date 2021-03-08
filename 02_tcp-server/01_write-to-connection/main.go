package main
import (
	"fmt"
	"io"
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
		io.WriteString(conn,"Ohio mina\n")
		fmt.Fprintln(conn,"arigato gozaimasu")
		fmt.Fprintf(conn,"Inu")
		conn.Close()
	}
}