package main

import (
	aoc "aoc/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = aoc.ConfigureLogging()

type ConversionMap struct {
	sourceValue int
	targetValue int
	rangeValue  int
}

func solve(rawTxt []byte) string {

	paragraphs := strings.Split(string(rawTxt), "\n\n")
	conversionMapsList := make([][]ConversionMap, 0)
	seedsStr := strings.Split(paragraphs[0], ": ")[1]
	seedsStrSlice := strings.Split(seedsStr, " ")
	log.Info(seedsStrSlice)
	var startingSeeds []int
	var rangeSeeds []int
	startValue := 0
	for i, seedStr := range seedsStrSlice {
		if i%2 == 0 {
			startValue, _ = strconv.Atoi(seedStr)
			startingSeeds = append(startingSeeds, startValue)
		} else {
			rangeValue, _ := strconv.Atoi(seedStr)
			rangeSeeds = append(rangeSeeds, rangeValue)
		}
	}
	log.Info(startingSeeds)
	log.Info(rangeSeeds)

	for _, paragraph := range paragraphs[1:] {
		conversionMaps := make([]ConversionMap, 0)

		titleAndMapsStr := strings.Split(paragraph, ":\n")
		// title := titleAndMapsStr[0]
		conversionMapsStr := titleAndMapsStr[1]
		conversionMapsStrSlice := strings.Split(conversionMapsStr, "\n")
		for _, conversionMapStr := range conversionMapsStrSlice {

			values := strings.Split(conversionMapStr, " ")
			sourceValue, _ := strconv.Atoi(values[1])
			targetValue, _ := strconv.Atoi(values[0])
			rangeValue, _ := strconv.Atoi(values[2])

			conversionMap := ConversionMap{
				sourceValue: sourceValue,
				targetValue: targetValue,
				rangeValue:  rangeValue,
			}

			conversionMaps = append(conversionMaps, conversionMap)
		}
		conversionMapsList = append(conversionMapsList, conversionMaps)
	}

	log.Info(conversionMapsList)

	minLocation := conversionMapsList[len(conversionMapsList)-1][1].targetValue

	for i, seed := range startingSeeds {
		log.Info("Seed: ", seed)
		log.Info("Percent: ", float64(float64(i)/float64(len(startingSeeds))))
		for j := 0; j < rangeSeeds[i]; j++ {
			location := seed + j
			for _, conversionMaps := range conversionMapsList {
				for _, conversionMap := range conversionMaps {
					if location >= conversionMap.sourceValue && location <= conversionMap.sourceValue+conversionMap.rangeValue {
						offset := location - conversionMap.sourceValue
						// log.Info("Seed ", seed, " is converted to ", conversionMap.targetValue+offset)
						location = conversionMap.targetValue + offset
						break
					}
				}
			}
			if location < minLocation {
				log.Info("Min location updated: ", location)
				minLocation = location
				// minSeedIndex = i
			}
		}
	}

	return fmt.Sprintf("%d", minLocation)
}
