package cli

import (
	"flag"
	"fmt"
	"log"

	"github.com/ryanbabida/bingo/generator"
)

func Run() error {
	cardCount := flag.Int("count", 4, "number of cards to be generated")
	output := flag.String("output", "html", "print cards as html files")
	// clean := flag.Bool("clean", false, "cleans out bin folder")
	flag.Parse()

	cards := generator.GenerateCards(*cardCount)

	err := Write(cards, *output)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func Write(cards []generator.Card, output string) error {
	switch output {
	case "html":
		cardsHtml := generator.CardHtml{Cards: cards}
		err := cardsHtml.Write()

		if err != nil {
			return err
		}
	case "console":
		cardsConsole := generator.CardConsole{Cards: cards}
		err := cardsConsole.Write()

		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid output parameter: %v", output)
	}

	return nil
}
