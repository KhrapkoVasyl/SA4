package main

import (
	"bufio"
	"log"
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
			postErr := eventLoop.Post(cmd)
			if postErr != nil {
				log.Fatalf("Post error: %s", postErr)
			}
		}
	}

	eventLoop.AwaitFinish()

	// Throws an error when trying to call a Post() method after the finish of the EventLoop
	err := eventLoop.Post(engine.PrintCommand("error command"))
	if err != nil {
		log.Fatalf("Post error: %s", err)
	}

}
