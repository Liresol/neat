/*
The proceeding file was copied from derui's go-threes program. All credit goes to derui for creating this.
*/
package libthrees
//package main?

import (
	"fmt"
)

const (
	BOARD_SIZE = 4
)

type Direction int

const (
	UP	Direction = iota
	DOWN	Direction = iota
	LEFT	Direction = iota
	RIGHT	Direction = iota
)

func (t Direction) String() string {
	switch t {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	default:
		return ""
	}
}

//?
type Mover func (indices []int) (int, Three)

type Line struct {
	data []Three

	dir Direction

	//shrinked
	moved bool
}

func (t Line) IsMoved() bool {
	return t.moved
}

type Board struct {
	area [BOARD_SIZE][BOARD_SIZE]Three
}

func GetBoard() Board {
	return Board{[BOARD_SIZE][BOARD_SIZE]Three{
		{empty(), empty(), empty(), empty()},
		{empty(), empty(), empty(), empty()},
		{empty(), empty(), empty(), empty()},
		{empty(), empty(), empty(), empty()}}}
}

func validPos(v int) bool {
	return 0 <= v && v < BOARD_SIZE
}

func (b *Board) GetArea() []ValueDefiner {
	ret := []ValueDefiner{}

	for yi, line := range b.area {
		for xi, v := range line {
			ret = append(ret, &PosValue{xi, yi, v})
		}
	}
	return ret
}

func GetInitializedBoard(accessors []ValueDefiner) (Board, error) {

	board := GetBoard()

	for _, v := range accessors {
		if !(validPos(v.X()) && validPos(v.Y())) {
			return GetBoard(), fmt.Errorf("the position of x and y must be from 0 to 3")
		}
		board.area[v.Y()][v.X()] = v.Value()
	}

	return board, nil
}

func (t Board) At(pos Pos) (Three, error) {
		if !(validPos(pos.X()) && validPos(pos.Y())) {
			return empty(), fmt.Errorf("(x,y) = (%d, %d) are not containers from the board", pos.X(), pos.Y())
		}
		return t.area[pos.Y()][pos.X()], nil
}

func getPositions(dir Direction) [][]Pos {
	switch dir {
	case UP:
		ret := make([][]Pos, 0)
		for xi := 0; xi < BOARD_SIZE; xi++ {
			line := make([]Pos, 0)
			for yi := 0; yi < BOARD_SIZE; yi++ {
				line = append(line, Pos{xi, yi})
			}
			ret = append(ret, line)
		}
		return ret
	case DOWN:
		ret := make([][]Pos, 0)
		for xi := 0; xi < BOARD_SIZE; xi++ {
			line := make([]Pos, 0)
			for yi := BOARD_SIZE - 1; yi >= 0; yi-- {
				line = append(line, Pos{xi, yi})
			}
			ret = append(ret, line)
		}

		return ret

	case LEFT:
		ret := make([][]Pos, 0)
		for yi := 0; yi < BOARD_SIZE; yi++ {
			line := make([]Pos, 0)
			for xi := 0; xi < BOARD_SIZE; xi++ {
				line = append(line, Pos{xi, yi})
			}
			ret = append(ret, line)
		}

		return ret
	case RIGHT:
		ret := make([][]Pos, 0)
		for yi := 0; yi < BOARD_SIZE; yi++ {
			line := make([]Pos, 0)
			for xi := BOARD_SIZE - 1; xi >= 0; xi-- {
				line = append(line, Pos{xi, yi})
			}
			ret = append(ret, line)
		}

		return ret
	default:
		panic("Unknown Direction value!")
	}
}

func (t Board) separateWithDirection(dir Direction) []Line {

	positions := getPositions(dir)

	ret := []Line{}
	for _, linePositions := range positions {

		line := Line{}

		for _, pos := range linePositions {
			v, _ := t.At(pos)
			line.data = append(line.data, v)
		}
		ret = append(ret, line)
	}

	return ret
}

func (b Board) maxValue() int {
	buf := 0
	for _, line := range b.area {
		for _, v := range line {
			if (v.Value() > buf) {
				buf = v.Value()
			}
		}
	}
	return buf
}

func makeNewLine() [BOARD_SIZE]Three {
	return [BOARD_SIZE]Three{empty(), empty(), empty(), empty()}
}

func (t *Board) CanMove(dir Direction) bool {
	lines := t.separateWithDirection(dir)
	allError := true
	newThrees := makeNewLine()

	for index, line := range lines {
		line.data = append(line.data, newThrees[index])
		merged, _ := moveLine(line)

		if merged {
			allError = false
		}
	}

	if allError {
		return false
	}
	return true
}

func (t *Board) CanSomeMove() bool {
	directions := []Direction{UP, LEFT, DOWN, RIGHT}

	for _, dir := range directions {
		if t.CanMove(dir) {
			return true;
		}
	}
	return false;
}

func (t *Board) Move(dir Direction, mover Mover) bool {
	lines := t.separateWithDirection(dir)
	positions := getPositions(dir)

	allError := true
	newThrees := makeNewLine()

	for index, line := range lines {
		line.data = append(line.data, newThrees[index])
		merged, mergedLine := moveLine(line)

		position := positions[index]

		for mergedLineIndex, v := range mergedLine.data {
			pos := position[mergedLineIndex]
			t.area[pos.Y()][pos.X()] = v
		}

		if merged {
			allError = false
		}

		lines[index] = mergedLine
	}
	if allError {
		return false
	}

	solvedLineIndices := []int{}
	for index, line := range lines {
		if line.IsMoved() {
			solvedLineIndices = append(solvedLineIndices, index)
		}
	}

	targetIndex, v := mover(solvedLineIndices)
	pos := positions[targetIndex][BOARD_SIZE - 1]
	t.area[pos.Y()][pos.X()] = v
	return true
}

func moveLine (line Line) (moved bool, ret Line) {
	solveSets := [][]int{{0,1},{1,2},{2,3}}
	moved = false
	ret = line
	ret.moved = false
	if line.data[0].IsEmpty() {
		ret.data = make([]Three, len(line.data)-1)
		copy(ret.data, line.data[1:])
		moved = true;
		ret.moved = true
		return
	}

	ret.data = make([]Three, len(line.data)-1)
	copy(ret.data, line.data[:BOARD_SIZE])

	for _, set := range solveSets {
		first := set[0]
		second := set[1]

		if line.data[first].IsEmpty() {
			first := ret.data[:first]
			second := line.data[second:]
			ret.data = append(first, second...)
			moved = true
			ret.moved = true
			break
		}
		if line.data[first].CanMerge(line.data[second]) {
			ret.data[first] = line.data[first].Merge(line.data[second])
			first := ret.data[:first+1]
			second := line.data[second+1:]
			ret.data = append(first, second...)
			moved = true
			ret.moved = true
			break
		}
	}

	return
}

func (b *Board) printThree(numOfThree int) [16]bool {
	iter := 0
	ret := [16]bool{}
	for _, line := range b.area {
		for _, v := range line {
			if (v.Value() == numOfThree) {
				ret[iter] = true
			}
			iter += 1
		}
	}
	return ret
}

