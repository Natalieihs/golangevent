package main

import "fmt"

type Event int

const (
	EventA Event = iota
	EventB
)

type EventEmitter struct {
	events map[Event]chan struct{}
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		events: make(map[Event]chan struct{}),
	}
}

func (e *EventEmitter) On(event Event, lister chan struct{}) {
	if _, ok := e.events[event]; !ok {
		e.events[event] = make(chan struct{})
	}
	e.events[event] = lister
}

// Emit 触发事件
func (e *EventEmitter) Emit(event Event) {
	if _, ok := e.events[event]; ok {
		close(e.events[event])
	}
}
func main() {
	emitter := NewEventEmitter()

	listenerA := make(chan struct{})
	emitter.On(EventA, listenerA)

	listenerB := make(chan struct{})
	emitter.On(EventB, listenerB)
	go func() {
		emitter.Emit(EventA)
	}()

	select {
	case <-listenerA:
		fmt.Println("EventA 被触发")
	case <-listenerB:
		fmt.Println("EventB 被触发")
	}
}
