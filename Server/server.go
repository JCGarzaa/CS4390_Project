// server
package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
    "time"
)

const (
    SERVER_HOST = "localhost"
    SERVER_PORT = "8888"
    SERVER_TYPE = "tcp"
)

type ClientInfo struct {
    Name string
    InitialConectionTime time.Time
}

var connectionsInfo map[string]ClientInfo = make(map[string]ClientInfo)

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

        ip := connection.RemoteAddr().String()
        fmt.Println("Connection accepted from " + ip)
        connectionTime := time.Now()

        _, err = connection.Write([]byte("Welcome to the Basic Math Server. To exit, type 'exit'\n"))
        if err != nil {
            fmt.Println("Error Writing initial welcome to client:", err)
            os.Exit(1)
        }

        buffer := make([]byte, 1024)
        initialPayloadLength, err := connection.Read(buffer) 
        if err != nil {
            fmt.Println("Error reading from client:", err)
            os.Exit(1)
        }

        clientName := strings.TrimSpace(string(buffer[:initialPayloadLength])) // trim leading and trailing whitespaces
        // store the client name and the time the connection was accepted
        connectionsInfo[ip] = ClientInfo{Name: clientName, InitialConectionTime: connectionTime}
        go processClient(connection)
    }
}

func processClient(connection net.Conn) {
    ip := connection.RemoteAddr().String()
    clientName := connectionsInfo[ip].Name
    for {
        buffer := make([]byte, 1024)
        messageLength, err := connection.Read(buffer)
        if err != nil {
            fmt.Println("Error reading from client:", err)
            os.Exit(1)
        }


        message := strings.TrimSpace(string(buffer[:messageLength])) // trim the message of leading and trailing whitespaces
        fmt.Println("Received from " + clientName + ": " + message) // log the message received from the client
        if message == "exit" {
            fmt.Println("Client " + clientName + " requested to close the connection. Closing connection...")
            connection.Close()
            fmt.Println("Client " + clientName + " disconnected.")
            duration := time.Since(connectionsInfo[ip].InitialConectionTime)
            fmt.Println("Connection duration: ", duration, "\n")
            delete (connectionsInfo, ip) // remove the connection time from the map
            return
        }

        response := solveMathProblem(message) // Calculate the math problem
        fmt.Println("Sending response to " + clientName + ": " + response) // log the response being sent to the client
        _, err = connection.Write([]byte(response)) // send the response to the client
    }
}

func solveMathProblem(problem string) string {
    out, err := exec.Command("python", "Server/eval.py", problem).Output()
    if err != nil {
        fmt.Println("Error running eval.py:", err)
        return "Error solving math problem. Please try again"
    }
    return string(out)
}
