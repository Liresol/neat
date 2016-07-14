package main

import (
	"fmt"
	"github.com/boggo/neat/experiments/threes/libthrees"
	"math"
	"math/rand"
)

var D = libthrees.DOWN
var U = libthrees.UP
var R = libthrees.RIGHT
var L = libthrees.LEFT
var tMoves = []libthrees.Direction{D, U, R, L}

func randMove(g libthrees.Game) (libthrees.Direction, bool) {
	if g.IsOver() {
		return U, false
	}
	var pMoves []libthrees.Direction
	for _, d := range tMoves {
		if g.CanMove(d) {
			pMoves = append(pMoves, d)
		}
	}
	CMove := rand.Intn(len(pMoves))
	return pMoves[CMove], true

}

func avgMoves(i int) float64 {
	g := libthrees.Game{}
	moves := 0.0
	for I := 0; I < i; I++ {
		for !g.IsOver() {
			m, _ := randMove(g)
			g.Move(m)
			moves += 1
		}
		g = libthrees.Game{}
		g.Initialize()
	}
	moves /= float64(i)
	return moves
}

/*
func avgTotal(iter int, i int, c chan float64) {
	for I := 0; I < iter; I++ {
		go avgMoves(i,c)
	}
}
*/

func stdDev(f []float64) float64 {
	if len(f) == 0 {
		fmt.Println("Bruh")
		return -1.0
	}
	m := 0.0
	for _, v := range f {
		m += v
	}
	m /= float64(len(f))
	s := 0.0
	for _, v := range f {
		s += (v - m) * (v - m)
	}
	s /= float64(len(f))
	return math.Pow(s, 0.5)
}

func main() {
	sum := 0.0
	max := 0.0
	min := 9999999.0
	Ssum := []float64{}
	test := libthrees.Game{}
	test.Initialize()
	for i := 0; i < 200000; i++ {
		v := avgMoves(100)
		Ssum = append(Ssum,v)
		if i % 1000 == 0 {
			fmt.Println(i)
		}
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	for _,i := range Ssum {
		sum += i
	}
	sum /= float64(len(Ssum))
	fmt.Println(sum)
	fmt.Println(stdDev(Ssum))
	fmt.Println("The max is", max)
	fmt.Println("The min is", min)
	fmt.Println("Lel")
	/*
		fmt.Println(test)
		test.PrintBoard();
		test.Move(D)
		test.PrintBoard();
		test.Move(L)
		test.PrintBoard();
		test.Move(L)
		test.PrintBoard();
		test.Move(L)
		test.PrintBoard();
		test.Move(L)
		test.PrintBoard();
		fmt.Println(test)
		fmt.Println(test.TestDeck())
	*/
}
