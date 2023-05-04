package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/aws/smithy-go/rand"
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

func sendFile(size int) error {
    file := make([]byte, size)

    _, err := io.ReadFull(rand.Reader, file)
    if err != nil {
        return err
    }

    conn, err := net.Dial("tcp", ":3000")
    if err != nil {
        return err
    }

    n, err := conn.Write(file)
    if err != nil {
        return err
    }

    fmt.Printf("written %d bytes\n", n)
    return nil
}

func main() {
    fs := &FileServer{}
    fs.start()
}
