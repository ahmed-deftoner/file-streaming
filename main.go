package main

import (
	"fmt"
	"log"
	"net"
)

type FileServer struct {}

func (fs *FileServer) start() {
    ln, err := net.Listen("tcp", ":3000")
    if err != nil {
        log.Fatal(err)
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go fs.readloop(conn)
    }
}

func (fs *FileServer) readloop(conn net.Conn)  {
    buf := make([]byte, 2048)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            log.Fatal(err)
        }
        file := buf[:n]
        fmt.Println(file)
        fmt.Printf("Recieved %d bytes\n", n)
    }
}

func main() {

}
