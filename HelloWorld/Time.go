package main

import (
	"fmt"
	"time"
)

func main() {
	birthday, _ := time.Parse("Jan 2 2006", "Nov 10 2009") // time.Time
	age := time.Since(birthday)
	fmt.Printf("Go is %d days old\n", age/(time.Hour*24))

	t := time.Now()
	fmt.Println(t.In(time.UTC))
	home, _ := time.LoadLocation("Australia/Sydney")
	fmt.Println(t.In(home))
}