package engine

import "fmt"

// Command represents actions that can be performed
// in a single event loop iteration.
type Command interface {
	Execute(handler Handler)
}

type PrintCommand string

func (pc PrintCommand) Execute(h Handler) {
	fmt.Println(string(pc))
}

type ReverseCommand string

func (rc ReverseCommand) Execute(h Handler) {
	reversedStr := reverseStr(string(rc))
	h.Post(PrintCommand(reversedStr))
}

func reverseStr(str string) string {
	runeSlice := []rune(str)
	for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}

	return string(runeSlice)
}

type StopCommand struct{}

func (sc StopCommand) Execute(h Handler) {
	h.(*EventLoop).stop = true
}
