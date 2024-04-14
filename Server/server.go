// server
package main

import (
    "fmt"
    "net"
    "os"
)

const (
    SERVER_HOST = "localhost"
    SERVER_PORT = "8888"
    SERVER_TYPE = "tcp"
)

func main() {
    fmt.Println("Starting server on " + SERVER_HOST + ":" + SERVER_PORT)
    server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
    if err != nil {
        fmt.Println("Error starting server:", err)
        os.Exit(1)
    }

    defer server.Close()
    fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
    fmt.Println("Waiting for connection...")
    for {
        connection, err := server.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            os.Exit(1)
        }

        fmt.Println("Connection accepted from " + connection.RemoteAddr().String())
        go processClient(connection)
    }
}

func processClient(connection net.Conn) {
    buffer := make([]byte, 1024)
    messageLength, err := connection.Read(buffer)
    if err != nil {
        fmt.Println("Error reading from client:", err)
        os.Exit(1)
    }
    fmt.Println("Received: ", string(buffer[:messageLength]))
    _, err = connection.Write([]byte("Thanks! Received message: " + string(buffer[:messageLength])))
    connection.Close()
}
