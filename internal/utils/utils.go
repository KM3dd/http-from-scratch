package utils

import (
	"fmt"
	"strings"

	"github.com/KM3dd/http-from-scratch/internal/types"
)

const CRLF = "\r\n"

func BuildResponse(response types.Response) []byte {

	var resp []byte
	var content_type = ""
	var content_length = ""
	first_line := fmt.Sprintf("HTTP/1.1 %d %s\r\n", response.Code, response.Message)
	if response.ContentType != "" {
		content_type = fmt.Sprintf("Content-Type: %s\r\n", response.ContentType)
	}
	if response.ContentLength != 0 {
		content_length = fmt.Sprintf("Content-Length: %d\r\n", response.ContentLength)
	}

	resp = []byte(first_line + content_type + content_length + CRLF + response.Body)
	return resp
}

func ParseRequest(buf []byte, n int) types.Request {

	request := strings.Split(string(buf[:n]), "\r\n")

	request_first_line := strings.Split(request[0], " ")

	operation := request_first_line[0]
	route := strings.Split(request_first_line[1], "/")[1:]
	fmt.Println("Here is route :", route)
	headers := request[1 : len(request)-2]
	body := request[len(request)-1]

	return types.Request{Operation: operation, Route: route, Headers: headers, Body: body}

}
