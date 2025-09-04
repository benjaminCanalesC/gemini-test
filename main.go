package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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

	model := "gemini-2.5-flash-lite"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bienvenido al chat con Gemini AI. Escribe 'salir' o 'exit' para terminar.")

	for {
		fmt.Print("Usuario: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "salir" || strings.ToLower(input) == "exit" {
			fmt.Println("Chat terminado.")
			break
		}

		result, err := client.Models.GenerateContent(
			ctx,
			model,
			genai.Text(input),
			nil,
		)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Gemini:", result.Text())
	}
}
