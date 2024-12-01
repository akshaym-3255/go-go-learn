package publisher

import (
	"fmt"
	"pub_sub/agent"
)

type Publisher struct {
	agent agent.Agent
}


func (p *Publisher)Publish(topic string, message string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("error occured handling panic")
		}
	}()

	p.agent.Publish(topic, message)
}