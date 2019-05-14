package Queue

import (
	"fmt"
	"time"
)

func start(c chan string) {
	for i := 0; i < 3; i++ {
		if i%2 == 0 {
			go producer(c, i)
		} else {
			go consumer(c, i)
		}
	}
}

func producer(c chan string, id int) {
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("from: %d, content: %d\n", id, i)
		c <- msg
		fmt.Printf("producer %d send message %d\n", id, i)
	}
}

func consumer(ch chan string, id int) {
	select {
	case <-ch:
		msg := <-ch
		fmt.Printf("consumer %d consume message %s\n", id, msg)
	}
}

func main() {
	var c = make(chan string, 10)
	start(c)

	time.Sleep(5 * time.Second)
	close(c)
	time.Sleep(5 * time.Second)
}
