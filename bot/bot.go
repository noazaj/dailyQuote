package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/zajicekn/dailyQuote/quote"
)

type message struct {
	Chat_id string `json:"chat_id"`
	Text    string `json:"text"`
}

func Bot() (string, error) {
	textMessageQuote := quote.Quote()

	// Create the message entity from the message struct. This will include
	// the ID of the chatbot and the text that will be sent (quote & author)
	messageJson, err := json.Marshal(message{Chat_id: "@ZenQuoteBot", Text: textMessageQuote})
	if err != nil {
		return "", fmt.Errorf("error marshalling the JSON data: %v", err)
	}

	token := os.Getenv("API_KEY")
	if token == "" {
		return "", fmt.Errorf("API_KEY not found in environment variables")
	}

	methodName := "sendMessage"
	url := fmt.Sprintf("https://api.telegram.org/bot%v/%v", token, methodName)

	resp, err := http.Post(url, "application/json", bytes.NewReader(messageJson))
	if err != nil {
		return "", fmt.Errorf("error in posting message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("request failed with status code: %d, response: %s", resp.StatusCode, body)
	}

	return "Message sent successfully!", nil
}

func init() {
	// Load the .env variables once during package initialization
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
