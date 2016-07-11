/*
The proceeding file was copied from derui's go-threes program. All credit goes to derui for creating this.
*/
package libthrees

import (
	"fmt"
	"math"
	"strconv"
)

type Three struct {
	value int
	base  int
}

func empty() Three {
	return Three{0, 0}
}

func (t Three) IsEmpty() bool {
	return t.value == 0 && t.base == 0
}

func GetThree(num int) Three {
	if num <= 0 {
		return empty()
	}

	switch num {
	case 1:
		fallthrough
	case 2:
		return Three{num, 0}
	default:
		break
	}

	return Three{3 * int(math.Floor(math.Pow(2., float64(num-3)))), num - 3}
}

func (t Three) String() string {
	return strconv.Itoa(t.value)
}

func (t Three) Value() int {
	return t.value
}

func (t Three) Base() int {
	return t.base
}

func (t Three) Score() int {

	if t.IsEmpty() {
		return 0
	}

	switch t.value {
	case 1:
		fallthrough
	case 2:
		return 0
	default:
		break
	}

	pow := math.Floor(math.Pow(3., float64(t.base+1)))
	return int(pow)
}

func (t Three) CanMerge(other Three) bool {

	switch t.value {
	case 1:
		return other.value == 2
	case 2:
		return other.value == 1
	default:
		return t.value == other.value
	}
}

func (t Three) Merge(other Three) Three {

	if !t.CanMerge(other) {
		panic(fmt.Sprintf("Can not merge with %s", other))
	}

	switch t.value {
	case 1:
		fallthrough
	case 2:
		return GetThree(3)
	default:
		return GetThree(3 + t.base + 1)
	}
}


func GetAllCombPos() (ret []Pos) {

	const size = BOARD_SIZE
	for i := 0; i < size; i++ {
		for j := 0; j < size;j++ {
			ret = append(ret, GetPos(i, j))
		}
	}

	return
}

