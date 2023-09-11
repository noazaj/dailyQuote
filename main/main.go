package main

import (
	"fmt"

	"github.com/zajicekn/dailyQuote/bot"
)

func main() {
	resp := bot.Bot()
	fmt.Println(resp)
}
