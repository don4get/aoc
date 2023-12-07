package main

import (
	aoc "aoc/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = aoc.ConfigureLogging()

func part1(rawTxt []byte) string {

	handsMap := make(map[string]int, 0)

	handsMap["55555"] = 7
	handsMap["44441"] = 6
	handsMap["33322"] = 5
	handsMap["33311"] = 4
	handsMap["22221"] = 3
	handsMap["22111"] = 2
	handsMap["11111"] = 1

	lines := strings.Split(string(rawTxt), "\n")

	decks := make([]string, 0)
	bets := make([]int, 0)

	for _, line := range lines {
		decksBetsStr := strings.Split(line, " ")
		if len(decksBetsStr) < 2 {
			continue
		}
		decks = append(decks, decksBetsStr[0])
		bet, err := strconv.Atoi(decksBetsStr[1])

		if err != nil {
			log.Fatal(err)
		}

		bets = append(bets, bet)
	}

	cardsOrder := "AKQJT98765432"
	sortedDecks := make([]string, len(decks))
	for i, deck := range decks {
		deckSlice := strings.Split(deck, "")

		sort.Slice(deckSlice, func(i, j int) bool {
			return strings.IndexByte(cardsOrder, deckSlice[i][0]) < strings.IndexByte(cardsOrder, deckSlice[j][0])
		})
		// log.Info(deckSlice)
		sort.Slice(deckSlice, func(i, j int) bool {
			return strings.Count(deck, string(deckSlice[i][0])) > strings.Count(deck, string(deckSlice[j][0]))
		})
		sortedDeck := strings.Join(deckSlice, "")
		sortedDecks[i] = sortedDeck
	}

	decksBetsMap := make(map[string]int, 0)
	for i, deck := range sortedDecks {
		decksBetsMap[deck] = bets[i]
	}
	decksSortedMap := make(map[string]string, 0)
	for i, deck := range sortedDecks {
		decksSortedMap[deck] = decks[i]
	}

	sort.Slice(sortedDecks, func(i, j int) bool {
		countI := make([]int, 0)
		countJ := make([]int, 0)
		for k := range sortedDecks[i] {
			countI = append(countI, strings.Count(sortedDecks[i], string(sortedDecks[i][k])))
			countJ = append(countJ, strings.Count(sortedDecks[j], string(sortedDecks[j][k])))
		}

		countIStr := intSliceToString(countI)
		countJStr := intSliceToString(countJ)

		if handsMap[countIStr] != handsMap[countJStr] {
			return handsMap[countIStr] > handsMap[countJStr]
		}

		correspondingDeckI := decksSortedMap[sortedDecks[i]]
		correspondingDeckJ := decksSortedMap[sortedDecks[j]]

		for k := range correspondingDeckI {
			if strings.IndexByte(cardsOrder, correspondingDeckI[k]) == strings.IndexByte(cardsOrder, correspondingDeckJ[k]) {
				continue
			}
			return strings.IndexByte(cardsOrder, correspondingDeckI[k]) < strings.IndexByte(cardsOrder, correspondingDeckJ[k])
		}
		return false
	})

	result := 0
	slices.Reverse(sortedDecks)
	for _, deck := range sortedDecks {
		log.Info(deck)
	}
	for i, deck := range sortedDecks {
		bet, present := decksBetsMap[deck]
		if !present {
			log.Fatal(present)
		}
		result += bet * (i + 1)

	}

	return fmt.Sprintf("%d", result)
}

func part2(rawTxt []byte) string {

	handsMap := make(map[string]int, 0)

	handsMap["55555"] = 10
	handsMap["44441"] = 6
	handsMap["33322"] = 5
	handsMap["33311"] = 4
	handsMap["22221"] = 3
	handsMap["22111"] = 2
	handsMap["11111"] = 1

	lines := strings.Split(string(rawTxt), "\n")

	decks := make([]string, 0)
	bets := make([]int, 0)

	for _, line := range lines {
		decksBetsStr := strings.Split(line, " ")
		if len(decksBetsStr) < 2 {
			continue
		}
		decks = append(decks, decksBetsStr[0])
		bet, err := strconv.Atoi(decksBetsStr[1])

		if err != nil {
			log.Fatal(err)
		}

		bets = append(bets, bet)
	}

	cardsOrder := "AKQT98765432J"
	sortedDecks := make([]string, len(decks))
	for i, deck := range decks {
		deckSlice := strings.Split(deck, "")

		sort.Slice(deckSlice, func(i, j int) bool {
			return strings.IndexByte(cardsOrder, deckSlice[i][0]) < strings.IndexByte(cardsOrder, deckSlice[j][0])
		})
		sort.Slice(deckSlice, func(i, j int) bool {
			return strings.Count(deck, string(deckSlice[i][0])) > strings.Count(deck, string(deckSlice[j][0]))
		})
		sortedDeck := strings.Join(deckSlice, "")
		sortedDecks[i] = sortedDeck
	}

	// for i, deck := range sortedDecks {
	// 	countJokers := strings.Count(deck, "J")
	// 	if countJokers == 5 {
	// 		sortedDecks[i] = "AAAAA"
	// 		continue
	// 	}
	// 	deck = strings.Replace(deck, "J", "", -1)
	// 	firstStr := string(deck[0])
	// 	sortedDecks[i] = strings.Repeat(firstStr, countJokers) + deck
	// }

	decksBetsMap := make(map[string]int, 0)
	for i, deck := range sortedDecks {
		decksBetsMap[deck] = bets[i]
	}
	decksSortedMap := make(map[string]string, 0)
	for i, deck := range sortedDecks {
		decksSortedMap[deck] = decks[i]
	}

	sort.Slice(sortedDecks, func(i, j int) bool {
		deckI := sortedDecks[i]

		countJokers := strings.Count(deckI, "J")
		if countJokers == 5 {
			deckI = "AAAAA"
		} else {
			deckI = strings.Replace(deckI, "J", "", -1)
			firstStr := string(deckI[0])
			deckI = strings.Repeat(firstStr, countJokers) + deckI
		}

		deckJ := sortedDecks[j]

		countJokers = strings.Count(deckJ, "J")
		if countJokers == 5 {
			deckJ = "AAAAA"
		} else {
			deckJ = strings.Replace(deckJ, "J", "", -1)
			firstStr := string(deckJ[0])
			deckJ = strings.Repeat(firstStr, countJokers) + deckJ
		}

		countI := make([]int, 0)
		countJ := make([]int, 0)
		for k := range deckI {
			countI = append(countI, strings.Count(deckI, string(deckI[k])))
			countJ = append(countJ, strings.Count(deckJ, string(deckJ[k])))
		}

		countIStr := intSliceToString(countI)
		countJStr := intSliceToString(countJ)

		if handsMap[countIStr] != handsMap[countJStr] {
			return handsMap[countIStr] > handsMap[countJStr]
		}

		correspondingDeckI := decksSortedMap[sortedDecks[i]]
		correspondingDeckJ := decksSortedMap[sortedDecks[j]]

		for k := range correspondingDeckI {
			if strings.IndexByte(cardsOrder, correspondingDeckI[k]) == strings.IndexByte(cardsOrder, correspondingDeckJ[k]) {
				continue
			}
			return strings.IndexByte(cardsOrder, correspondingDeckI[k]) < strings.IndexByte(cardsOrder, correspondingDeckJ[k])
		}
		return false
	})

	result := 0
	slices.Reverse(sortedDecks)
	for _, deck := range sortedDecks {
		log.Info(deck)
	}
	for i, deck := range sortedDecks {
		bet, present := decksBetsMap[deck]
		if !present {
			log.Fatal(present)
		}
		result += bet * (i + 1)

	}

	return fmt.Sprintf("%d", result)
}

// sort.Slice(sortedDecks, func(i, j int) bool {
// 	CountJokersI := strings.Count(sortedDecks[i], "J")
// 	CountJokersJ := strings.Count(sortedDecks[j], "J")

// 	countI := []int{1, 1, 1, 1, 1}
// 	countJ := []int{1, 1, 1, 1, 1}
// 	for k := range sortedDecks[i] {
// 		ikStr := string(sortedDecks[i][k])
// 		jkStr := string(sortedDecks[j][k])
// 		countIK := 0
// 		countJK := 0
// 		if ikStr != "J" {
// 			countIK = strings.Count(sortedDecks[i], ikStr)
// 		}

// 		if jkStr != "J" {
// 			countJK = strings.Count(sortedDecks[j], jkStr)
// 		}

// 		if k == 0 {
// 			for l := 0; l < min(countIK+CountJokersI, 5); l++ {
// 				countI[l] = countIK + CountJokersI
// 			}

// 			for l := 0; l < min(countJK+CountJokersJ, 5); l++ {
// 				countJ[l] = countJK + CountJokersJ
// 			}
// 			// countIK += CountJokersI
// 			// countJK += CountJokersJ
// 		}
// 		countI[k] = max(countIK, countI[k])
// 		countJ[k] = max(countJK, countJ[k])
// 	}

// 	countIStr := intSliceToString(countI)
// 	countJStr := intSliceToString(countJ)
// 	log.Info(sortedDecks[i], " ", countIStr, " ", handsMap[countIStr])
// 	if handsMap[countIStr] != handsMap[countJStr] {
// 		return handsMap[countIStr] > handsMap[countJStr]
// 	}

// 	correspondingDeckI := decksSortedMap[sortedDecks[i]]
// 	correspondingDeckJ := decksSortedMap[sortedDecks[j]]

// 	for k := range correspondingDeckI {
// 		if strings.IndexByte(cardsOrder, correspondingDeckI[k]) == strings.IndexByte(cardsOrder, correspondingDeckJ[k]) {
// 			continue
// 		}
// 		return strings.IndexByte(cardsOrder, correspondingDeckI[k]) < strings.IndexByte(cardsOrder, correspondingDeckJ[k])
// 	}
// 	return false
// })

// result := 0
// slices.Reverse(sortedDecks)
// for _, deck := range sortedDecks {
// 	log.Info(deck)
// }
// for i, deck := range sortedDecks {
// 	bet, present := decksBetsMap[deck]
// 	if !present {
// 		log.Fatal(present)
// 	}
// 	result += bet * (i + 1)

// }

// return fmt.Sprintf("%d", result)
// deckOccurances := make([]map[string]int, 0)
// for i, deck := range decks {
// 	nbOccurances := make(map[string]int, 0)

// 	for _, c := range deck {
// 		cStr := string(c)

// 		_, present := nbOccurances[cStr]
// 		if present {
// 			continue
// 		}
// 		nbOccurances[cStr] = strings.Count(deck, string(cStr))
// 	}

// 	deckOccurances = append(deckOccurances, nbOccurances)
// }

// for _, nbOccurances := range deckOccurances {
// 	for k, v := range nbOccurances {

// 	}

func intSliceToString(slice []int) string {
	// Convert each integer to a string
	var stringSlice []string
	for _, num := range slice {
		stringSlice = append(stringSlice, strconv.Itoa(num))
	}

	// Join the string representations using ","
	result := strings.Join(stringSlice, "")

	return result
}
