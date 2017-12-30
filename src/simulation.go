package solarsim

import "fmt"

type Simulation struct {
	planets []Planet
}

func (sim *Simulation) step() {
	for i, p := range sim.planets {
		fmt.Println("Test Planet %d : %s", i, p.name)
		p.update()
	}
}