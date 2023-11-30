package d01

import (
	aoc "aoc/utils"
	"fmt"
	"slices"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = aoc.ConfigureLogging()

func solve(rawTxt []byte) string {

	digitList := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	txt := string(rawTxt)
	lines := strings.Split(txt, "\n")

	sum := 0
	for _, line := range lines {
		log.Info(line)

		modifiedLine := ""
		for ind := 0; ind < len(line); ind++ {
			subLine := line[ind:]
			indices := []int{10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000}
			for i, digit := range digitList {
				if strings.Contains(subLine, digit) {
					// stillDigitsToBeFound = true
					indices[i] = strings.Index(subLine, digit)
					// log.Info("Found ", digit, " at ", indices[i])
				}
			}

			if slices.Contains(indices, 0) {
				modifiedLine += fmt.Sprintf("%d", slices.Index(indices, 0)+1)
			} else {
				modifiedLine += string(subLine[0])
			}

		}

		line = modifiedLine

		log.Info("Modified line: ", line)

		firstDigit := rune(-1)
		lastDigit := rune(-1)
		for _, c := range line {
			if unicode.IsDigit(c) {
				if firstDigit == rune(-1) {
					firstDigit = c
				}
				lastDigit = c
			}
		}
		log.Info("firstDigit: ", string(firstDigit))
		log.Info("lastDigit: ", string(lastDigit))
		digits := int(firstDigit-'0')*10 + int(lastDigit-'0')
		log.Info("Into int: ", digits)
		sum += digits
	}

	return fmt.Sprintf("%d", sum)
}

// stillDigitsToBeFound := true
// for stillDigitsToBeFound {
// 	// log.Info(line)
// 	stillDigitsToBeFound = false
//
// 	for i, digit := range digitList {
// 		if strings.Contains(line, digit) {
// 			stillDigitsToBeFound = true
// 			indices[i] = strings.Index(line, digit)
// 			// log.Info("Found ", digit, " at ", indices[i])
// 		}

// 		// line = strings.ReplaceAll(line, keyString, fmt.Sprintf("%d", valueInt))
// 	}
// 	if stillDigitsToBeFound {
// 		minValue := slices.Min(indices)
// 		minIndex := slices.Index(indices, minValue)
// 		line = strings.Replace(line, digitList[minIndex], fmt.Sprintf("%d", minIndex+1), 1)
// 		indices[minIndex] = 10000

// 	}
