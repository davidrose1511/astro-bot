package main

import (
	"astro-bot/internal/ai"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Load Env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	ctx := context.Background()

	// 2. Connect to AI (OpenRouter)
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENROUTER_API_KEY is required")
	}
	
	brain := ai.NewService(apiKey)

	// 3. Start Chat Loop
	fmt.Println("-------------------------------------------")
	fmt.Println("🔮 ASTRO BOT (OpenRouter Edition)")
	fmt.Println("-------------------------------------------")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nYOU: ")
		if !scanner.Scan() {
			break
		}
		userText := scanner.Text()

		if strings.TrimSpace(userText) == "exit" {
			break
		}

		// Process Message
		response, err := brain.ProcessUserMessage(ctx, userText)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		// Handle Response
		if strings.Contains(response, ">>> ACTION:") {
			fmt.Println("-------------------------------------------")
			fmt.Println(response) // Just print the raw extraction
			fmt.Println("-------------------------------------------")
		} else {
			fmt.Printf("BOT: %s\n", response)
		}
	}
}