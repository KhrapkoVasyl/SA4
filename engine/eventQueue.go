package engine

import "sync"

type EventQueue struct {
	mu        sync.Mutex
	commands  []Command
	isWaiting bool

	notEmpty chan struct{}
}

func (eq *EventQueue) push(c Command) {
	eq.mu.Lock()
	defer eq.mu.Unlock()
	eq.commands = append(eq.commands, c)

	if eq.isWaiting {
		eq.notEmpty <- struct{}{}
	}
}

func (eq *EventQueue) pull() Command {
	eq.mu.Lock()
	defer eq.mu.Unlock()

	if len(eq.commands) == 0 {
		eq.isWaiting = true
		eq.mu.Unlock()
		<-eq.notEmpty
		eq.mu.Lock()
		// eq.isWaiting = false
	}

	res := eq.commands[0]
	eq.commands[0] = nil
	eq.commands = eq.commands[1:]
	return res
}

func (eq *EventQueue) empty() bool {
	eq.mu.Lock()
	defer eq.mu.Unlock()
	return len(eq.commands) == 0
}
