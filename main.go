package main

import (
	"exercism/go/gigasecond"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(gigasecond.AddGigasecond(t))
}
