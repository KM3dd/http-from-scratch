package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/KM3dd/http-from-scratch/internal/handlers"
	"github.com/KM3dd/http-from-scratch/internal/types"
	"github.com/KM3dd/http-from-scratch/internal/utils"
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
	var resp types.Response

	request := utils.ParseRequest(buf, n)

	if request.Route == "/" {
		resp = handlers.RootHnadler()

	} else if strings.Split(request.Route, "/")[1] == "echo" {
		resp = handlers.EchoHandler(request.Route)

	} else if strings.Split(request.Route, "/")[1] == "user-agent" {
		resp = handlers.UserAgentHandler(request)

	} else {
		resp = handlers.NotFoundHandler()

	}
	response = utils.BuildResponse(resp)
	return response
}
