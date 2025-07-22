package utils

import (
	"fmt"

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
		content_type = fmt.Sprintf("Content-Length: %d\r\n", response.ContentLength)
	}

	resp = []byte(first_line + content_type + content_length + CRLF + response.Body)

	return resp
}
