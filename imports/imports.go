package imports

import (
	"fmt"
	"math"
	"github.com/darkhelmet/env"
)

func MathIsNeat() {
	fmt.Println("Pi is like, around", math.Pi)
	line := fmt.Sprintf("Now you have %.5f problems", math.Sqrt(7))
	fmt.Println(line)
	fmt.Println("3 + 3 is", add(3, 3))
	fmt.Println("Path is...", env.String("PATH"))
}

func add(x, y int) int {
	return x + y
}

type Wife struct {
	HowHerDayWas string
	IsSheCute bool
	IsSheMad bool
}

func (w *Wife) CookDinner() bool {
	willSheCook := ! w.IsSheMad
	return willSheCook
}

type Husband struct {
	HowHisDayWas string
	IsHeHandsome bool
	IsHeSad bool
}

func (h *Husband) MowTheLawn() bool {
	willHeDoIt := ! h.IsHeSad
	return willHeDoIt
}
