package api

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kyokomi/emoji"
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/quote"
)

// Stonk is a struct for your stonk
type Stonk struct {
	SymbolList []string
}

// queryStonk makes an API call to gather quote
func queryStonk(s string) *finance.Quote {
	q, err := quote.Get(s)

	if q == nil {
		fmt.Println("Can't find this stonk!")
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}

	return q
}

// Print outputs stonk data
func (s Stonk) Print() {
	var outputEmoji string

	for _, quote := range s.SymbolList {
		var output strings.Builder
		data := queryStonk(quote)
		if data.RegularMarketChangePercent < 0 {
			outputEmoji = ":chart_with_downwards_trend:"
		} else if data.RegularMarketChangePercent > 10 {
			outputEmoji = ":rocket:"
		} else if data.RegularMarketChangePercent > 0 {
			outputEmoji = ":chart_with_upwards_trend:"
		} else {
			outputEmoji = ":face_with_rolling_eyes:"
		}
		output.WriteString(data.Symbol)
		output.WriteString(" ")
		output.WriteString(fmt.Sprintf("$%.2f", data.RegularMarketPrice))
		output.WriteString(" ")
		output.WriteString(fmt.Sprintf("%.2f%%", data.RegularMarketChangePercent))
		output.WriteString(" ")
		output.WriteString(emoji.Sprint(outputEmoji))
		fmt.Println(output.String())
	}
}
