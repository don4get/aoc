package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithDummy(t *testing.T) {
	inputText, err := os.ReadFile("dummy.txt")

	if err != nil {
		t.Errorf("Error reading input.txt: %v", err)
	}

	result := solve(inputText)

	expected := "4361"

	assert.Equal(t, expected, result, "Result should be %v, but was %v", expected, result)
}

func TestWithDummy2(t *testing.T) {
	inputText, err := os.ReadFile("dummy2.txt")

	if err != nil {
		t.Errorf("Error reading input.txt: %v", err)
	}

	result := solve(inputText)

	expected := "467835"

	assert.Equal(t, expected, result, "Result should be %v, but was %v", expected, result)
}

func TestWithInput(t *testing.T) {

	inputText, err := os.ReadFile("input2.txt")

	if err != nil {
		t.Errorf("Error reading input.txt: %v", err)
	}

	// logrus.Info(string(inputText))

	result := solve(inputText)

	log.Info("Result: ", result)
}
