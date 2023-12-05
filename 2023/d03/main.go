package main

import (
	aoc "aoc/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = aoc.ConfigureLogging()

type Position struct {
	x int
	y int
}

func solve(rawTxt []byte) string {
	sum := 0

	lines := strings.Split(string(rawTxt), "\n")

	partsList := []string{}
	partPositions := []Position{}
	symbolsPositions := []Position{}
	gearPositions := []Position{}

	for j, line := range lines {
		symbolsAndParts := strings.Split(line, ".")

		for _, symbolAndPart := range symbolsAndParts {
			symbolAndPartIndex := strings.Index(line, symbolAndPart)

			pattern := regexp.MustCompile("[0-9]+")

			parts := pattern.FindAllString(symbolAndPart, -1)
			indices := pattern.FindAllStringIndex(symbolAndPart, -1)
			for k, part := range parts {
				partsList = append(partsList, part)
				i := symbolAndPartIndex + indices[k][0]
				partPositions = append(partPositions, Position{i, j})
			}

			for k, c := range symbolAndPart {
				if !unicode.IsDigit(c) {
					symbolsPositions = append(symbolsPositions, Position{k + symbolAndPartIndex, j})
				}
				if c == '*' {
					gearPositions = append(gearPositions, Position{k + symbolAndPartIndex, j})
				}
			}

			rep := strings.Repeat(".", len(symbolAndPart))
			line = strings.Replace(line, symbolAndPart, rep, 1)
		}
	}

	// for i, partPosition := range partPositions {
	// 	adjacent := false
	// 	for j, _ := range partsList[i] {
	// 		for _, symbolPosition := range symbolsPositions {
	// 			distanceX := int(math.Abs(float64(partPosition.x + j - symbolPosition.x)))
	// 			distanceY := int(math.Abs(float64(partPosition.y - symbolPosition.y)))

	// 			if distanceX < 2 && distanceY < 2 {
	// 				adjacent = true
	// 				continue
	// 			}
	// 		}
	// 	}
	// 	if adjacent {
	// 		partInt, _ := strconv.Atoi(partsList[i])
	// 		sum += partInt
	// 	}
	// }

	for _, gearPosition := range gearPositions {
		adjacentParts := 0
		valid := false
		for i, partPosition := range partPositions {
			for j, _ := range partsList[i] {
				distanceX := int(math.Abs(float64(gearPosition.x - partPosition.x - j)))
				distanceY := int(math.Abs(float64(gearPosition.y - partPosition.y)))

				if distanceX < 2 && distanceY < 2 {
					val, _ := strconv.Atoi(partsList[i])
					if adjacentParts == 0 {
						adjacentParts += val
					} else {
						valid = true
						adjacentParts *= val
					}
					break
				}
			}
		}
		if valid {
			sum += adjacentParts
		}
	}

	return fmt.Sprintf("%d", sum)
}
