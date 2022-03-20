package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://cisco.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://outlook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkRoutine(link, c) //spawn new go routine/thread
	}

	// for {
	// 	go checkRoutine(<-c, c)
	// }

	// for l := range c {
	// 	go checkRoutine(l, c)
	// }

	for l := range c {
		go func(link string) { //function literals (lamda/anonymour function like)
			time.Sleep(5 * time.Second)
			checkRoutine(link, c)
		}(l)
	}
}

func checkRoutine(link string, c chan string) {
	_, err := http.Get(link) //blocking code
	if err != nil {
		fmt.Println("looks like ", link, "is down :/")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
