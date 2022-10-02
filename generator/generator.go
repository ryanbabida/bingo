package generator

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	B []int
	I []int
	N []int
	G []int
	O []int
}

func GenerateCards(n int) []Card {
	var cards []Card

	for i := 0; i < n; i += 4 {
		set := generateCards()
		cards = append(cards, set...)
	}

	// cards are generated in sets of 4, cut the slice to just return n
	return cards[0:n]
}

func generateCards() []Card {
	var cards []Card

	bVals := generateValuesByRange(1, 15)
	iVals := generateValuesByRange(16, 30)
	nVals := generateValuesByRange(31, 45)
	gVals := generateValuesByRange(46, 60)
	oVals := generateValuesByRange(61, 75)

	cardCount := 3

	offset := 5

	for i := 0; i < cardCount; i++ {
		card := Card{
			B: bVals[offset*i : offset*i+5],
			I: iVals[offset*i : offset*i+5],
			N: nVals[offset*i : offset*i+5],
			G: gVals[offset*i : offset*i+5],
			O: oVals[offset*i : offset*i+5],
		}

		cards = append(cards, card)
	}

	bVals = generateValuesByRange(1, 15)
	iVals = generateValuesByRange(16, 30)
	// n values will be the Free values from the first 3 cards
	gVals = generateValuesByRange(46, 60)
	oVals = generateValuesByRange(61, 75)

	card := Card{
		B: bVals[0:5],
		I: iVals[0:5],
		G: gVals[0:5],
		O: oVals[0:5],
	}

	card.N = []int{cards[0].N[2], cards[1].N[2], -1, cards[2].N[2], cards[0].N[0]}

	cards = append(cards, card)

	// Free slot
	for _, card := range cards {
		card.N[2] = -1
	}

	return cards
}

func generateValuesByRange(lo, hi int) []int {
	var values []int

	for i := lo; i <= hi; i++ {
		values = append(values, i)
	}

	time.Sleep(1000)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })

	return values
}

type CardWriter interface {
	Write() error
}

type CardHtml struct {
	Cards []Card
}

func (c *CardHtml) Write() error {
	err := copyCSS()
	if err != nil {
		return err
	}

	for idx, card := range c.Cards {
		htmlBytes, err := os.ReadFile("static/template.html")
		if err != nil {
			return err
		}

		template := string(htmlBytes)
		for i := 0; i < 5; i++ {
			template = strings.ReplaceAll(template, "{B"+strconv.Itoa(i)+"}", strconv.Itoa(card.B[i]))
			template = strings.ReplaceAll(template, "{I"+strconv.Itoa(i)+"}", strconv.Itoa(card.I[i]))
			template = strings.ReplaceAll(template, "{N"+strconv.Itoa(i)+"}", strconv.Itoa(card.N[i]))
			template = strings.ReplaceAll(template, "{G"+strconv.Itoa(i)+"}", strconv.Itoa(card.G[i]))
			template = strings.ReplaceAll(template, "{O"+strconv.Itoa(i)+"}", strconv.Itoa(card.O[i]))
		}

		template = strings.ReplaceAll(template, "-1", "Free")

		f, err := os.Create("bin/cards-" + strconv.Itoa(idx+1) + ".html")
		if err != nil {
			return err
		}

		f.WriteString(template)

		defer f.Close()
	}

	return nil
}

type CardConsole struct {
	Cards []Card
}

func (c *CardConsole) Write() error {
	for _, card := range c.Cards {
		for i := 0; i < 5; i++ {
			fmt.Println(card.B[i], card.I[i], card.N[i], card.G[i], card.O[i])
		}
	}
	fmt.Println()

	return nil
}

func copyCSS() error {
	b, err := os.ReadFile("static/styles.css")
	if err != nil {
		return err
	}

	f, err := os.Create("bin/styles.css")
	if err != nil {
		return err
	}

	f.Write(b)
	defer f.Close()

	return nil
}
