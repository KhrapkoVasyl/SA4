package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventLoop(t *testing.T) {
	assert := assert.New(t)

	//Should create and start the event loop
	eventLoop := new(EventLoop)
	eventLoop.Start()
	assert.Equal(false, eventLoop.shouldStop, "The initial value of variable shouldStop should be equal false")
	assert.Equal(0, len(eventLoop.eq.commands), "The initial value of the commands slice length should be equal zero")

	err := eventLoop.Post(PrintCommand("hello"))
	assert.Equal(nil, err)

	err = eventLoop.Post(PrintCommand("wrong string"))
	assert.Equal(nil, err)

	err = eventLoop.Post(ReverseCommand("world"))
	assert.Equal(nil, err)

	err = eventLoop.Post(PrintCommand("hello2"))
	assert.Equal(nil, err)

	err = eventLoop.Post(ReverseCommand("revesedString"))
	assert.Equal(nil, err)

	assert.Equal(5, len(eventLoop.eq.commands))
	eventLoop.AwaitFinish()
	assert.Equal(true, eventLoop.shouldStop)
	assert.Equal(0, len(eventLoop.eq.commands))

	//calling post after loop finished

	postErr1 := eventLoop.Post(PrintCommand("print hello3"))
	assert.Error(postErr1)
	assert.EqualError(postErr1, "calling post after loop finished")
	postErr2 := eventLoop.Post(ReverseCommand("reverse testString"))
	assert.Error(postErr2)
	assert.EqualError(postErr2, "calling post after loop finished")
}
