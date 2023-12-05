package main

import (
	aoc "aoc/utils"
	"fmt"
	"slices"
	"strings"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = aoc.ConfigureLogging()

func solve(rawTxt []byte) string {
	sum := 0

	lines := strings.Split(string(rawTxt), "\n")
	nbLines := len(lines)
	copiesCards := make([]int, nbLines)
	for c := range copiesCards {
		copiesCards[c] = 1
	}

	for j, line := range lines {
		cards := strings.Split(line, ":")[1]
		nbCardsWinning := 0

		cardDecks := strings.Split(cards, "|")
		winningCards := strings.Split(cardDecks[0], " ")
		myCards := strings.Split(cardDecks[1], " ")

		for _, card := range winningCards {
			if card == "" {
				continue
			}
			if slices.Contains(myCards, card) {
				log.Info("Card ", card, " is winning")
				nbCardsWinning++
				log.Info("nbCardsWinning: ", nbCardsWinning)
			}
		}
		if nbCardsWinning >= 1 {
			// var toBeAdded float64 = math.Pow(2., float64(nbCardsWinning-1))
			// sum += int(toBeAdded)
			for i := 0; i < nbCardsWinning; i++ {
				copiesCards[j+i+1] += copiesCards[j]
			}

		}
	}

	for c := range copiesCards {
		sum += copiesCards[c]
	}

	return fmt.Sprintf("%d", sum)
}
