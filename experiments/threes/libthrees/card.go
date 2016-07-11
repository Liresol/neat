/*
The proceeding file was copied from derui's go-threes program. All credit goes to derui for creating this.
*/
package libthrees
//package main?


import (
	//"github.com/derui/go-threes/libthrees"
	//"fmt"
	//"strings"
)


type Card struct {
	target Three
	pos Pos
}

func (t *Card) X() int {
	return t.pos.X()
}

func (t *Card) Y() int {
	return t.pos.Y()
}

func (t *Card) Value() Three {
	return t.target
}
