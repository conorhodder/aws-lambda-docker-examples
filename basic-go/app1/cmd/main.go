package main

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handle)
}

func handle(_ context.Context, _ any) (string, error) {
	slog.Info("received new event")
	return "success from app 1", nil
}
