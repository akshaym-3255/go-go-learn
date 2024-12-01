package subscriber

import (
	"fmt"

)

type CollectingSubscriber struct {
	ReadChan chan string
	messages []string
}

func NewCollectingSubscriber() (*CollectingSubscriber) {
	return &CollectingSubscriber{
		messages:  make([]string, 0),
	}
}

func (c *CollectingSubscriber) Process() {

	for {
		val , ok := <-c.ReadChan

		if !ok {
			fmt.Println("error occured")
			break
		}
		fmt.Println("colleting subecrber", val)
		c.messages = append(c.messages, val)
	}
}