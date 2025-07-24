package handlers

import (
	"strings"

	"github.com/KM3dd/http-from-scratch/internal/types"
)

func RootHnadler() types.Response {
	return types.Response{Code: 200, Message: "OK"}
}

func NotFoundHandler() types.Response {
	return types.Response{Code: 404, Message: "Not Found"}
}

func EchoHandler(route string) types.Response {
	text := strings.Split(route, "/")[2]
	return types.Response{Code: 200, Message: "OK", ContentType: "text/plain", ContentLength: len(text), Body: text}
}

func UserAgentHandler(request types.Request) types.Response {
	user_agent := strings.TrimSpace(strings.Split(request.Headers[1], ":")[1])
	return types.Response{Code: 200, Message: "OK", ContentType: "text/plain", ContentLength: len(user_agent), Body: user_agent}
}
