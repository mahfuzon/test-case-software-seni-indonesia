package main

import "fmt"

type Robot interface {
	info() int
}

type Autobot struct {
	power int
}

type Decepticon struct {
	power int
}

func (a *Autobot) hit(d Decepticon) string {
	if a.power > d.power {
		return "win"
	}

	return "lose"
}

func (a *Autobot) info() int {
	return a.power
}

func (d *Decepticon) hit(a Autobot) string {
	if d.power > a.power {
		return "win"
	}
	return "lose"
}

func (d *Decepticon) info() int {
	return d.power
}

func printInfo(s Robot) int {
	return s.info()
}

func main() {
	optimusPrime := Autobot{power: 11}
	bumbleBee := Autobot{power: 10}
	megatron := Decepticon{power: 12}

	// Below are test cases.
	// If you change codes below, it does not impact what we run in the background
	// --
	fmt.Println(optimusPrime.hit(megatron)) // win
	fmt.Println(megatron.hit(optimusPrime)) // win
	fmt.Println(bumbleBee.hit(megatron))    // lose
	// --
	fmt.Println(printInfo(&optimusPrime)) // 11
	fmt.Println(printInfo(&bumbleBee))    // 10
	fmt.Println(printInfo(&megatron))     // 12
}
