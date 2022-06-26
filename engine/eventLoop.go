package engine

type Handler interface {
	Post(cmd Command)
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

func (loop *EventLoop) Post(cmd Command) {
	loop.eq.push(cmd)
}

func (loop *EventLoop) AwaitFinish() {
	loop.Post(StopCommand{})
	<-loop.stopSignal
}
