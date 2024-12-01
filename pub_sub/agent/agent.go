package agent

import (
	"sync"
	"fmt"
)

type Agent struct {
	mu sync.Mutex
	subscribers map[string][]chan string
	quit chan bool
}

func NewAgent() (*Agent) {
	return &Agent{
		mu : sync.Mutex{},
		subscribers: make(map[string][]chan string),
		quit:  make(chan bool),
	}
}


func (a *Agent) Publish(topic string, message string) {
	fmt.Println(message)
	for _, sub := range a.subscribers[topic]{
		sub <- message
	}
}


func (a *Agent) Subscribe(topic string) (chan string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	readChan := make(chan string)
	a.subscribers[topic] = append(a.subscribers[topic], readChan)
	return readChan
}