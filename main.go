package main

import (
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
    }
}


func main() {

}
