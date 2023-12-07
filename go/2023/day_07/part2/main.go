package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	bid   int
	cards string
}

func main() {
	hands := map[string][]Hand{}

	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_07/input.txt")
	rows := strings.Split(string(f), "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}
		rowContent := strings.Split(row, " ")
		bid, _ := strconv.Atoi(rowContent[1])
		hand := Hand{cards: rowContent[0], bid: bid}
		classifiedHand := classifyHand(hand)
		hands[classifiedHand] = append(hands[classifiedHand], hand)
	}

	for _, h := range hands {
		sort.Slice(h, func(i, j int) bool {
			return getFirstDifferentValue(h[i].cards, h[j].cards) < getFirstDifferentValue(h[j].cards, h[i].cards)
		})
	}

	hankedHands := []Hand{}
	hankedHands = append(hankedHands, hands["high_card"]...)
	hankedHands = append(hankedHands, hands["one_pair"]...)
	hankedHands = append(hankedHands, hands["two_pair"]...)
	hankedHands = append(hankedHands, hands["three_of_a_kind"]...)
	hankedHands = append(hankedHands, hands["full_house"]...)
	hankedHands = append(hankedHands, hands["four_of_a_kind"]...)
	hankedHands = append(hankedHands, hands["five_of_a_kind"]...)

	total := 0

	for i, h := range hankedHands {
		total += h.bid * (i + 1)
	}

	fmt.Println(total)
}

func classifyHand(hand Hand) string {
	cards := map[rune]int{}

	for _, card := range hand.cards {
		cards[card]++
	}

	criptedCard := criptCards(cards)

	switch criptedCard {
	case "5":
		return "five_of_a_kind"
	case "41":
		return "four_of_a_kind"
	case "32":
		return "full_house"
	case "311":
		return "three_of_a_kind"
	case "221":
		return "two_pair"
	case "2111":
		return "one_pair"
	default:
		return "high_card"
	}
}

func criptCards(cards map[rune]int) string {
	cardNumbers := []int{}

	for _, card := range cards {
		cardNumbers = append(cardNumbers, card)
	}

	sort.Slice(cardNumbers, func(i, j int) bool { return cardNumbers[i] > cardNumbers[j] })
	criptedCards := ""

	for _, card := range cardNumbers {
		criptedCards += fmt.Sprint(card)
	}

	if _, ok := cards['J']; ok {
		switch criptedCards {
		case "41":
			if cards['J'] == 1 || cards['J'] == 4 {
				criptedCards = "5"
			}
		case "32":
			if cards['J'] == 2 || cards['J'] == 3 {
				criptedCards = "5"
			}
		case "311":
			if cards['J'] == 1 || cards['J'] == 3 {
				criptedCards = "41"
			} else if cards['J'] == 2 {
				criptedCards = "5"
			}
		case "221":
			if cards['J'] == 2 {
				criptedCards = "41"
			} else if cards['J'] == 1 {
				criptedCards = "32"
			}
		case "2111":
			if cards['J'] == 1 || cards['J'] == 2 {
				criptedCards = "311"
			}
		default:
			if cards['J'] == 1 {
				criptedCards = "2111"
			}
		}
	}

	return criptedCards
}

func getFirstDifferentValue(cards1, cards2 string) int {
	cardsStrength := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 1,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	i := 0

	for cards1[i] == cards2[i] {
		i++
	}

	return cardsStrength[rune(cards1[i])]
}
