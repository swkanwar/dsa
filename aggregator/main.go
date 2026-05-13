package main

import (
	"fmt"
	"time"
)

type Activity struct {
	userID string
	Action string
}

func main() {
	ch := make(chan Activity, 1000)

	go startAggregator(ch)
	ticker := time.NewTicker(500 * time.Millisecond)
	for i := 0; i < 1000; i++ {
		select {
		case <-ticker.C:
			ch <- Activity{
				userID: fmt.Sprintf("User ID %d", i),
				Action: fmt.Sprintf("Action %d", i),
			}
		}

	}
	select {}
}

func startAggregator(ch chan Activity) {
	var buffer []Activity

	// Flush interval 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case event, ok := <-ch:
			if !ok {
				return
			}
			buffer = append(buffer, event)

			if len(buffer) >= 20 {
				fmt.Println("buffer:", buffer)
				buffer = nil
			}

		case <-ticker.C:
			if len(buffer) > 0 {
				fmt.Println("Ticket flush buffer:", buffer)
				buffer = nil
			}
		}
	}
}
