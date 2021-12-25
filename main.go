package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lamda"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	OK      bool   `json:"ok"`
}

func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Process Request ID %f", request.ID),
		OK:      true,
	}, nil
}

func main() {
	lamda.Start()
}
