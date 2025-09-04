package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash-lite",
		genai.Text("Write a short poem about Chile, in spanish."),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())
}
