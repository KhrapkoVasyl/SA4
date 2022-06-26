package engine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var result PrintCommand
	assert := assert.New(t)

	result = Parse("print hello")
	assert.Equal(
		PrintCommand("hello"),
		result,
		"Should return a print command with an argument")

	result = Parse("reverse hello")
	assert.Equal(
		ReverseCommand("hello"),
		result,
		"Should return a reverse command with an argument")

	result = Parse("print")
	assert.Equal(
		PrintCommand("SYNTAX ERROR: Not enough arguments"),
		result,
		"Should return a print command with syntax error, when no argument is passed")

	result = Parse("reverse")
	assert.Equal(
		PrintCommand("SYNTAX ERROR: Not enough arguments"),
		result,
		"Should return a print command with syntax error, when no argument is passed")

	result = Parse("")
	assert.Equal(
		PrintCommand("SYNTAX ERROR: Not enough arguments"),
		result,
		"Should return a print command with syntax error, when an empty string is passed")

	result = Parse("print hello1 hello2")
	assert.Equal(
		PrintCommand("SYNTAX ERROR: Too many arguments"),
		result,
		"Should return a print command with syntax error, when passed too many arguments")

	result = Parse("reverse hello1 hello2")
	assert.Equal(
		PrintCommand("SYNTAX ERROR: Too many arguments"),
		result,
		"Should return a print command with syntax error, when passed too many arguments")

	result = Parse("alert hello")
	assert.Equal(
		PrintCommand("SYNTAX ERROR: Unknown instruction"),
		result,
		"Should return a print command with syntax error, when passed unknown instruction")
}

func ExampleParse() {
	result = Parse("reverse hello")
	fmt.Println(result)
	// result: ReverseCommand("hello")
}
