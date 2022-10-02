package generator

import "testing"

func TestGenerateValuesByRange(t *testing.T) {
	lo := 1
	hi := 15
	vals := generateValuesByRange(lo, hi)

	for i := lo; i <= hi; i++ {
		found := false
		for _, val := range vals {
			if val == i {
				found = true
			}
		}

		if found == false {
			t.Fatalf(`%v is not in values from generatedValuesByRange(1, 15)`, i)
		}
	}
}

func TestGenerateCards(t *testing.T) {
	cards := generateCards()

	if len(cards) != 4 {
		t.Fatalf(`generateCards returns %v cards instead of 4`, len(cards))
	}
}

func TestGenerateCardsWithCount(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, exp := range expected {
		cards := GenerateCards(exp)

		if len(cards) != exp {
			t.Fatalf(`generateCards returns %v cards instead of %v`, len(cards), exp)
		}
	}
}
