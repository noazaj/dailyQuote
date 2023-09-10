package main

import (
	"fmt"

	"github.com/zajicekn/dailyQuote/quote"
)

func main() {
	quoteMessage := quote.Quote()
	fmt.Println(quoteMessage)
}
