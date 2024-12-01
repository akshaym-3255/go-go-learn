package main

import (
	"pub_sub/agent"
	"pub_sub/subscriber"
)

func main(){

	
	agent := agent.NewAgent()

	normalSubscriber := subscriber.NewNormalSubscriber()

	collecting := subscriber.NewCollectingSubscriber()


	topic := "normal"

	ch := agent.Subscribe(topic)
	normalSubscriber.ReadChan = ch

	go normalSubscriber.Process()

	collectingch := agent.Subscribe(topic)
	collecting.ReadChan = collectingch

	go collecting.Process()
	for {
		agent.Publish(topic, "hello ")
	}
	
}