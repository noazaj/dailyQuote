package main

import (
	"fmt"
	"log"

	"github.com/zajicekn/dailyQuote/bot"
)

func main() {
	message, err := bot.Bot()
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println(message)
}
