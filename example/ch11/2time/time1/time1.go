package main

import (
	"fmt"
	"time"
)

func main() {
	//liststart1
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	d := 2*time.Hour + 30*time.Minute + 45*time.Second // d の型はtime.Duration
	fmt.Println(d)                                     // 2h30m45s
	//listend1
}
