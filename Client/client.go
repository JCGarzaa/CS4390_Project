package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "strings"
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

    fmt.Println("Connected to server on " + CLIENT_HOST + ":" + CLIENT_PORT)
    
    // send name to server
    var initialPayload []byte = []byte(name)
    _, err = connection.Write(initialPayload)
    if err != nil {
        fmt.Println("Error sending name to server:", err)
        return
    }

    // create a reader to read user input
    reader := bufio.NewReader(os.Stdin)

    // create a reader to read responses from the server
    defer connection.Close()
    for {
        // read response from server
        buffer := make([]byte, 1024)
        messageLength, err := connection.Read(buffer)
        if err != nil {
            if err.Error() == "EOF" {
                fmt.Println("Server closed connection. Exiting...")
                connection.Close()
            } else {
                fmt.Println("Error reading from server:", err)
            }
            return 
        }
        fmt.Println("SERVER RESPONSE:", string(buffer[:messageLength]))

        fmt.Print("Enter math expression to send to the server: ")
        expression, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading from user:", err)
            return
        }
        expression = strings.TrimSpace(expression)

        // send message to server
        _, err = connection.Write([]byte(expression))
        if err != nil {
            fmt.Println("Error sending message to server:", err)
            return
        }
    }
}
