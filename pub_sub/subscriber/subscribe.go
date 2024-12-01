package subscriber

import "fmt"

type NormalSubscriber struct {
	ReadChan chan string
}

func NewNormalSubscriber() (*NormalSubscriber) {
	return &NormalSubscriber{
	}
}

func (n *NormalSubscriber) Process() {

	for {
		val , ok := <-n.ReadChan

		if !ok {
			fmt.Println("error occured")
			break
		}
		fmt.Println("normal subecrber", val)
		fmt.Println(val)
	}
}