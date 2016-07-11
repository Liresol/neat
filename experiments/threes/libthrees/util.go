/*
The proceeding file was copied from derui's go-threes program. All credit goes to derui for creating this.
*/
package libthrees

import (
	"fmt"
)

type Accessor interface {
	X() int
	Y() int
}

type ValueDefiner interface {
	Accessor
	Value() Three
}

type Pos struct {
	x int
	y int
}

func (t *Pos) X() int {
	return t.x
}

func (t *Pos) Y() int {
	return t.y
}

func (t *Pos) String() string {
	return fmt.Sprintf("(%d,%d)", t.x, t.y)
}

func GetPos(x, y int) Pos {
	return Pos{x, y}
}

type PosValue struct {
	x     int
	y     int
	value Three
}

func (t *PosValue) X() int {
	return t.x
}

func (t *PosValue) Y() int {
	return t.y
}

func (t *PosValue) Value() Three {
	return t.value
}

