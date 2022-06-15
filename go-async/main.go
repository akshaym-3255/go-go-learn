package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func TwoTable(two chan string) {
	for i := 1; i < 11; i++ {
		fmt.Println(2 * i)
	}
	two <- "done"

}

func ThreeTable(three chan string) {
	for i := 1; i < 11; i++ {
		fmt.Println(3 * i)
		time.Sleep(6000)
	}
	three <- "done"
}

func GetUser(user chan User) {
	fmt.Println("inside")
	resp, err := http.Get("https://randomuser.me/api/")

	if err != nil {
		log.Fatalf("error occured")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error occured")
	}
	//fmt.Println(string(body))
	var user1 User
	json.Unmarshal(body, &user1)
	user <- user1
	//fmt.Println(user.Info)

}

func main() {
	two := make(chan string)
	three := make(chan string)
	user := make(chan User)
	go TwoTable(two)
	go ThreeTable(three)
	go GetUser(user)
	IsDone := <-two
	IsDoneThree := <-three
	result := <-user
	fmt.Println(result.Info.Seed)
	fmt.Println(IsDone, IsDoneThree)
}
