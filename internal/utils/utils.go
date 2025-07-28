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

	first_line := fmt.Sprintf("HTTP/1.1 %d %s%s", response.Code, response.Message, CRLF)
	if response.ContentType != "" {
		content_type = fmt.Sprintf("Content-Type: %s%s", response.ContentType, CRLF)
	}
	if response.ContentLength != 0 {
		content_length = fmt.Sprintf("Content-Length: %d%s", response.ContentLength, CRLF)
	}

	resp = []byte(first_line + content_type + content_length + CRLF + response.Body)
	return resp
}

func ParseRequest(buf []byte, n int) types.Request {

	request := strings.Split(string(buf[:n]), CRLF)
	request_first_line := strings.Split(request[0], " ")

	operation := request_first_line[0] // GET, POST

	route := strings.Split(request_first_line[1], "/")[1:]

	headersMap := make(map[string]string)
	headers := request[1 : len(request)-2]
	for _, h := range headers {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headersMap[key] = value
		}
	}
	body := request[len(request)-1]

	return types.Request{Operation: operation, Route: route, Headers: headersMap, Body: body}

}
