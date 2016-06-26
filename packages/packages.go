package packages

import (
	"fmt"
	"math/rand"
)

func RandomNumberPrint() {
	rand.Seed(54638943)
	fmt.Println("My fave number is", rand.Intn(10))
}