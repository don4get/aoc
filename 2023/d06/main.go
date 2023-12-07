package main

import (
	aoc "aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = aoc.ConfigureLogging()

func part1(rawTxt []byte) string {
	lines := strings.Split(string(rawTxt), "\n")

	durationsStr := strings.Split(lines[0], ":")[1]
	durationsStrSlice := strings.Split(durationsStr, " ")
	durations := make([]int, 0)
	for _, durationStr := range durationsStrSlice {
		if durationStr == "" {
			continue
		}
		duration, err := strconv.Atoi(durationStr)

		if err != nil {
			log.Fatal(err)
		}

		durations = append(durations, duration)
	}

	recordsStr := strings.Split(lines[1], ":")[1]
	recordsStrSlice := strings.Split(recordsStr, " ")

	records := make([]int, 0)
	for _, recordStr := range recordsStrSlice {
		if recordStr == "" {
			continue
		}
		record, err := strconv.Atoi(recordStr)

		if err != nil {
			log.Fatal(err)
		}

		records = append(records, record)
	}

	nbWays := make([]int, 0)

	for j, duration := range durations {
		root1, root2 := solveQuadratic(float64(duration), float64(records[j]))

		if -root1*root1+float64(duration)*root1-float64(records[j]) == 0 {
			root1++

		}
		if -root2*root2+float64(duration)*root2-float64(records[j]) == 0 {
			root2--
		}

		nbWay := int(math.Ceil(root2) - math.Ceil(root1) + 1)
		nbWays = append(nbWays, nbWay)
	}

	result := 1

	for _, nbWay := range nbWays {
		result *= nbWay
	}

	return fmt.Sprintf("%d", result)
}

func part2(rawTxt []byte) string {
	lines := strings.Split(string(rawTxt), "\n")
	durationStr := strings.Split(lines[0], ":")[1]
	durationStr = strings.Replace(durationStr, " ", "", -1)

	recordStr := strings.Split(lines[1], ":")[1]
	recordStr = strings.Replace(recordStr, " ", "", -1)

	duration, err := strconv.Atoi(durationStr)

	if err != nil {
		log.Fatal(err)
	}

	record, err := strconv.Atoi(recordStr)

	if err != nil {
		log.Fatal(err)
	}

	root1, root2 := solveQuadratic(float64(duration), float64(record))
	if -root1*root1+float64(duration)*root1-float64(record) == 0 {
		root1++

	}
	if -root2*root2+float64(duration)*root2-float64(record) == 0 {
		root2--
	}

	nbWay := int(math.Ceil(root2) - math.Ceil(root1) + 1)

	return fmt.Sprintf("%d", nbWay)
}

func solveQuadratic(duration, record float64) (float64, float64) {
	// Coefficients of the quadratic equation
	a := -1.0
	b := duration
	c := -record

	// Calculate the discriminant
	discriminant := b*b - 4*a*c

	// Check if roots are real
	if discriminant < 0 {
		return math.NaN(), math.NaN()
	}

	// Calculate the roots
	root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	root2 := (-b - math.Sqrt(discriminant)) / (2 * a)

	return root1, root2
}
