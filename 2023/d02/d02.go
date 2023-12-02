package d02

import (
	aoc "aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = aoc.ConfigureLogging()

func solve(rawTxt []byte) string {

	colors := []string{"red", "green", "blue"}
	sum := 0

	lines := strings.Split(string(rawTxt), "\n")

	for _, line := range lines {
		minNumbers := []int{0, 0, 0}

		idStr := strings.Split(line, ":")[0]
		idStr = strings.ReplaceAll(idStr, "Game ", "")
		_, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error(err)
		}
		// log.Info("ID: ", id)
		result := strings.Split(line, ":")[1]
		subsets := strings.Split(result, ";")

		for _, subset := range subsets {
			// log.Info("Subset: ", subset)
			diceInfo := strings.Split(subset, ",")

			for _, dice := range diceInfo {
				for i, color := range colors {
					if strings.Contains(dice, color) {
						// log.Info("Found ", color)
						numberSpace := strings.Replace(dice, color, "", 1)
						number := strings.ReplaceAll(numberSpace, " ", "")
						numberInt, err := strconv.Atoi(number)
						if err != nil {
							log.Error(err)
						}

						minNumbers[i] = int(math.Max(float64(minNumbers[i]), float64(numberInt)))
					}
				}
			}

		}

		// if minNumbers[0] > 12 || minNumbers[1] > 13 || minNumbers[2] > 14 {
		// 	sum += 0
		// } else {
		// 	sum += id
		// 	log.Info("ID: ", id)
		// }
		// log.Info("MinNumbers: ", minNumbers)
		// log.Info("Sum: ", sum)

		minimalSet := minNumbers[0] * minNumbers[1] * minNumbers[2]
		sum += minimalSet

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
