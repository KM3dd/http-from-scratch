package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		// goroutine for the new connection
		go HandleRequest(conn)
	}

}

func HandleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		os.Exit(1)
	}

	response := MakeResponse(buf, n)
	conn.Write(response)
	conn.Close()
}

func MakeResponse(buf []byte, n int) []byte {
	var response []byte
	request := strings.Split(string(buf[:n]), "\r\n")
	request_line := request[0]

	route := strings.Split(request_line, " ")[1]
	fmt.Println("Received request for route: ", route)
	if route == "/" {
		response = []byte("HTTP/1.1 200 OK\r\n\r\n")

	} else if strings.Split(route, "/")[1] == "echo" {
		text := strings.Split(route, "/")[2]
		response = []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(text), text))
	} else if strings.Split(route, "/")[1] == "user-agent" {
		user_agent := strings.TrimSpace(strings.Split(request[2], ":")[1])
		response = []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(user_agent), user_agent))
	} else {
		response = []byte("HTTP/1.1 404 Not Found\r\n\r\n")
	}
	return response
}
