package simulation

import "fmt"

// Results holds the data obtained from the simulation
type Results struct {
	Susceptible, Infected, Resistant int
}

// Simulate runs the actual simulation for the program to display
func Simulate(pop, initI int, S2I, I2R, S2R float64) []Results {
	fmt.Println("hi")
	return []Results{{0, 0, 0}, {1, 1, 1}}
}
