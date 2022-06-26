package engine

import (
	"strings"
)

func Parse(commandLine string) Command {
	parts := strings.Fields(commandLine)

	if len(parts) < 2 {
		return PrintCommand("SYNTAX ERROR: Not enough arguments")
	}
	if len(parts) > 2 {
		return PrintCommand("SYNTAX ERROR: Too many arguments")
	}

	switch parts[0] {
	case "print":
		return PrintCommand(parts[1])
	case "reverse":
		return ReverseCommand(parts[1])
	default:
		return PrintCommand("SYNTAX ERROR: Unknown instruction")
	}
}
