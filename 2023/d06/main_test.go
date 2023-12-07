package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithDummy(t *testing.T) {
	inputText, err := os.ReadFile("dummy.txt")

	if err != nil {
		t.Errorf("Error reading input.txt: %v", err)
	}

	result := part1(inputText)

	expected := "288"

	assert.Equal(t, expected, result, "Result should be %v, but was %v", expected, result)
}

func TestPart2WithDummy(t *testing.T) {
	inputText, err := os.ReadFile("dummy2.txt")

	if err != nil {
		t.Errorf("Error reading input.txt: %v", err)
	}

	result := part2(inputText)

	expected := "71503"

	assert.Equal(t, expected, result, "Result should be %v, but was %v", expected, result)
}

func TestPart1WithInput(t *testing.T) {

	inputText, err := os.ReadFile("input2.txt")

	if err != nil {
		t.Errorf("Error reading input.txt: %v", err)
	}

	// logrus.Info(string(inputText))

	result := part1(inputText)

	log.Info("Result: ", result)
}

func TestPart2WithInput(t *testing.T) {

	inputText, err := os.ReadFile("input2.txt")

	if err != nil {
		t.Errorf("Error reading input.txt: %v", err)
	}

	// logrus.Info(string(inputText))

	result := part2(inputText)

	log.Info("Result: ", result)
}

func TestPart3WithInput(t *testing.T) {

	inputText, err := os.ReadFile("input2.txt")

	if err != nil {
		t.Errorf("Error reading input.txt: %v", err)
	}

	// logrus.Info(string(inputText))

	result := part2(inputText)

	log.Info("Result: ", result)
}
