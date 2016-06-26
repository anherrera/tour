package main

import (
	"fmt"
	"time"
	"github.com/anherrera/tour/packages"
)

func main() {
	fmt.Println("Welcome!")
	fmt.Println("The time is", time.Now())
	packages.RandomNumberPrint()
}