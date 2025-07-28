package handlers

import (
	"fmt"
	"os"

	"github.com/KM3dd/http-from-scratch/internal/types"
)

func RootHnadler() types.Response {
	return types.Response{Code: 200, Message: "OK"}
}

func NotFoundHandler() types.Response {
	return types.Response{Code: 404, Message: "Not Found"}
}

func EchoHandler(request types.Request) types.Response {
	text := request.Route[1]
	return types.Response{Code: 200, Message: "OK", ContentType: "text/plain", ContentLength: len(text), Body: text}
}

func UserAgentHandler(request types.Request) types.Response {
	user_agent := request.Headers["user-agent"]
	return types.Response{Code: 200, Message: "OK", ContentType: "text/plain", ContentLength: len(user_agent), Body: user_agent}
}

func FilesHandler(request types.Request) types.Response {
	file_name := request.Route[1]
	path := fmt.Sprintf("/tmp/%s", file_name)

	if request.Operation == "GET" {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return types.Response{Code: 404, Message: "Not Found"}
		}

		// Read file content
		data, err := os.ReadFile(path)
		if err != nil {
			return types.Response{Code: 500, Message: "Something went wrong"}
		}
		return types.Response{Code: 200, Message: "OK", ContentType: "application/octet-stream", ContentLength: len(string(data)), Body: string(data)}

	} else if request.Operation == "POST" {
		data := request.Body
		err := os.WriteFile(path, []byte(data), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return types.Response{Code: 500, Message: "Something went wrong"}
		}

		fmt.Println("File written successfully.")

		return types.Response{Code: 201, Message: "Created"}
	}
	return types.Response{Code: 500, Message: "Something went wrong"}
}
