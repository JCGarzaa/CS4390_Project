// client
package main

import (
    "fmt"
    "net"
    "os"
)

const (
    CLIENT_HOST = "localhost"
    CLIENT_PORT = "8888"
    CLIENT_TYPE = "tcp"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run client.go <name>")
        os.Exit(1)
    }

    var name string = os.Args[1]

    // establish connection
    connection, err := net.Dial(CLIENT_TYPE, CLIENT_HOST+":"+CLIENT_PORT)

    if err != nil {
        fmt.Println("Error connecting to server:", err)
        os.Exit(1)
    }

    // send message
    _, err = connection.Write([]byte("Hello from " + name))
    buffer := make([]byte, 1024)
    messageLength, err := connection.Read(buffer)
    if err != nil {
        fmt.Println("Error reading from server:", err)
        os.Exit(1)
    }

    fmt.Println("Received: ", string(buffer[:messageLength]))
    defer connection.Close()
}
