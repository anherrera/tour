package main

import (
	"fmt"
	"github.com/anherrera/tour/imports"
	"github.com/anherrera/tour/packages"
	"time"
)

func main() {
	fmt.Println("Welcome!")
	fmt.Println("The time is", time.Now())
	packages.RandomNumberPrint()
	imports.MathIsNeat()

	var wife imports.Wife
	wife.HowHerDayWas = "Alright, I suppose. Busy."
	wife.IsSheCute = true
	wife.IsSheMad = false

	wife2 := &imports.Wife{"Super bad", true, true}
	fmt.Println("The wife is mad?", wife2.IsSheMad)
	fmt.Println("Will she cook then?", wife2.CookDinner())

	var hubby imports.Husband
	hubby.IsHeSad = true
	fmt.Println("The husband is sad?", hubby.IsHeSad)
	fmt.Println("Will he mow the lawn?", hubby.MowTheLawn())

	fmt.Println("DONE-ZO!");
}
