package main

import (
	"errors"
	"github.com/boggo/neat"
	"github.com/boggo/neat/archiver"
	"github.com/boggo/neat/decoder"
	"github.com/boggo/neat/experiments/threes/libthrees"
	"github.com/boggo/neat/popeval"
	"github.com/boggo/neat/reporter"
	"github.com/boggo/neat/settings"
	//"math"
	"sort"
	//"math/rand"
	//"time"
	"fmt"
)

var _ = fmt.Print

var (
	D = libthrees.DOWN
	U = libthrees.UP
	R = libthrees.RIGHT
	L = libthrees.LEFT
)

type threesEval struct{}

//0 is Down
//1 is Up
//2 is Right
//3 is Left

func intDir(i int) libthrees.Direction {
	switch i {
	case 0:
		return D
	case 1:
		return U
	case 2:
		return R
	case 3:
		return L
	default:
		panic("Wat.")
	}
}

/*
func threes(action []float64, g libthrees.Game) [255]float64 {
	if(len(action) != 4) {
		fmt.Println("Wat.")
	}
	return g.GetFloatState()
}
*/

func (eval threesEval) Evaluate(org *neat.Organism) (err error) {
	if org.Phenome == nil {
		err = errors.New("Cannot evaluate an org without a Phenome")
		org.Fitness = []float64{0} // Minimal fitness
		return
	}
	var g libthrees.Game

	fitness := float64(0)
	numtests := 100

	for trials := 0; trials < numtests; trials++ {
		g = libthrees.Game{}
		//Causes the games to be deterministic.
		g.SeedInit(4761+trials)
		//g.Initialize()
		for !g.IsOver() {
			//255 inputs
			inputs := g.GetFloatState()
			action, err2 := org.Analyze(inputs)
			if err2 != nil {
				err = err2
				org.Fitness = []float64{0}
				return
			}
			//fmt.Println(action)
			var priority []float64
			for _, i := range action {
				priority = append(priority, i)
			}
			sort.Float64s(priority)
			//fmt.Println(priority)
			//fmt.Println(action)
			//fmt.Println(priority)
			for i := 0; i < 4; i++ {
				j := 0
				for j = 0; j < 4; j++ {
					if priority[3-i] == action[j] {
						break
					}
				}
				if g.CanMove(intDir(j)) {
					g.Move(intDir(j))
					break
				}
			}
			fitness += 1
		}
		//fmt.Println(fitness)
	}
	fitness /= float64(numtests)
	org.Fitness = []float64{fitness}
	return
}
func main() {

	// Load the settings
	ldr := settings.NewJSON("threes-settings.json")
	s, err := ldr.Load()
	if err != nil {
		panic(err)
	}

	// Create the archiver
	a := archiver.NewJSON("threes-pop.json")

	// Create the reporter
	r := reporter.NewConsole()

	// Create the evaluators
	o := &threesEval{}
	/*
		*The threes game is not concurrent-safe.
		p := popeval.NewConcurrent()
	*/
	p := popeval.NewSerial()

	// Create the decoder
	d := decoder.NewNEAT()

	// Iterate the experiment
	neat.Iterate(s, 100, d, p, o, a, r)
}
