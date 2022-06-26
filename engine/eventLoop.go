package engine

type Handler interface {
	Post(cmd Command)
}

type EventLoop struct {
	eq         *EventsQueue
	shoudlStop bool
	stopSignal chan struct{}
}

func (loop *EventLoop) Start() {
	loop.eq = &EventsQueue{
		notEmpty: make(chan struct{}),
	}

	loop.stopSignal = make(chan struct{})

	go func() {
		for !loop.shoudlStop || !loop.eq.isEmpty() {
			cmd := loop.eq.pull()
			cmd.Execute(loop)
		}

		loop.stopSignal <- struct{}{}
	}()
}

func (loop *EventLoop) Post(cmd Command) {
	// l.eq.push(cmd)
}

type stopCommand struct{}

func (scmd stopCommand) Execute(h Handler) {
	h.(*EventLoop).shoudlStop = true
}

func (loop *EventLoop) AwaitFinish() {
	loop.Post(stopCommand{})
	<-loop.stopSignal
}
