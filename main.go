package main

import (
	"bufio"
	"os"

	"github.com/KhrapkoVasyl/SA4/engine"
)

const inputFile = "input.txt"

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := engine.Parse(commandLine) // parse the line to get a Command
			eventLoop.Post(cmd)
		}
	}

	eventLoop.AwaitFinish()
}
