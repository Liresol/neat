package main
import (
	"github.com/boggo/neat/experiments/threes/libthrees"
	"fmt"
)

var D = libthrees.DOWN
var U = libthrees.UP
var R = libthrees.RIGHT
var L = libthrees.LEFT

func main() {
	test := libthrees.Game{}
	test.Initialize()
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
}
