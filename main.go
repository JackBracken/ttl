package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	y, m, d := now.Date()

	lunchTime := time.Date(y, m, d, 12, 0, 0, 0, time.Local)

	if now.Before(lunchTime) {
		r := int(time.Until(lunchTime).Seconds())
		hrs := r / 3600
		mins := r % 3600 / 60
		secs := r % 3600 % 60

		fmt.Printf("Lunch is in %d hours, %d minutes and %d seconds.", hrs, mins, secs)
	} else {
		fmt.Println("Lunch time is over.")
	}
}
