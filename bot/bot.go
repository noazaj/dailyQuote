package bot

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zajicekn/dailyQuote/quote"
)

func Bot() string {
	// This variable will hold the Quote and
	// Author from the quote package
	quoteMessage := quote.Quote()

	// Load the .env variables
	err := godotenv.Load(".env")
	if err != nil {
		cwd, _ := os.Getwd()
		log.Println("Current working directory:", cwd)
		log.Fatal("Error in loading .env file: ", err)
	}

	// Soon to implement API token

	return quoteMessage
}
