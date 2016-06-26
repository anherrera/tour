package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(54638943)
	fmt.Println("My fave number is", rand.Intn(10))
}