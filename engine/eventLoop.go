package engine

import "errors"

type Handler interface {
	Post(cmd Command) error
}

type EventLoop struct {
	eq         *EventQueue
	shouldStop bool
	stopSignal chan struct{}
}

func (loop *EventLoop) Start() {
	loop.eq = &EventQueue{
		notEmpty: make(chan struct{}),
	}

	loop.stopSignal = make(chan struct{})

	go func() {
		for !loop.shouldStop || !loop.eq.empty() {
			cmd := loop.eq.pull()
			cmd.Execute(loop)
		}

		loop.stopSignal <- struct{}{}
	}()
}

func (loop *EventLoop) Post(cmd Command) error {

	if loop.shouldStop {
		// if we send a Command after calling the AwaitFinish() method,
		// then it will not pass this check in the Post method, and
		// will not get into our EventQueue (an appropriate error will be displayed).
		// This check does not prevent us from using the ReverseCommand,
		// because all commands, that entered the EventQueue before the StopCommand, will
		// be executed and will call Post() of PrintCommand before the value of
		// shouldStop changes to true

		return errors.New("calling post after loop finished")
	}

	loop.eq.push(cmd)
	return nil
}

func (loop *EventLoop) AwaitFinish() {
	loop.Post(StopCommand{})
	<-loop.stopSignal
}
