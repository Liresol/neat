package libthrees

import (
	//"os/exec"
	//"os"
	"math/rand"
	"math"
	"fmt"
	//"strings"
	"time"
)
var _ = time.Now().Nanosecond()

var (
	//NUM_RANGE_1 = 4
	//NUM_RANGE_2 = 4
	//NUM_RANGE_3 = 4
	//What
	NUM_RANGE_CARD_INITIALIZE = []int{6,8}
	FULL_DECK = []Three{GetThree(1),GetThree(2),GetThree(3),GetThree(1),GetThree(2),GetThree(3),GetThree(1),GetThree(2),GetThree(3),GetThree(1),GetThree(2),GetThree(3)}
)

const (
	//Equivalent to 1/21.
	BONUS_CHANCE = float64(0.04761904761904)
)

type Game struct {
	b Board
	nextCard Card
	deck []Three
}

func (g *Game) Initialize() {
	rand.Seed(int64(time.Now().Nanosecond()))
	g.ReloadDeck()
	g.InitializeBoard()
	g.UpdateNextCard()
}

func (g *Game) SeedInit(i int) {
	rand.Seed(int64(i))
	g.ReloadDeck()
	g.InitializeBoard()
	g.UpdateNextCard()
}

func shuffle(t []Three) ([]Three) {
	iter := 0
	for iter < 2 {
		iter += 1
	for i := 1;i<len(t);i++ {
		j := rand.Intn(i)
		t[i], t[j] = t[j], t[i]
	}
}
	return t
}

func (g *Game) ReloadDeck() {
	if(len(g.deck) > 0) {
		panic("Attempted to reload non-empty deck")
	}
	g.deck = FULL_DECK
	g.deck = shuffle(g.deck)
}

func (g *Game) DrawFromDeck() (Three) {
	if(len(g.deck) == 0) {
		g.ReloadDeck()
	}
	ret := g.deck[0]
	g.deck = g.deck[1:]
	return ret
}

func (g *Game) GetBonus() (Three) {
	buf := g.b.maxValue()
	buf /= 3
	if buf <= 8 {
		return empty()
	}
	buf = int(math.Log2(float64(buf))-4)
	if(buf == 0) {
		return GetThree(4)
	}
	ret := rand.Intn(buf)
	return GetThree(ret + 4)
}

func (g *Game) UpdateNextCard() {
	if rand.Float64() > BONUS_CHANCE {
		g.nextCard.target = g.DrawFromDeck()
	} else {
		buf := g.GetBonus()
		if buf == empty() {
			g.nextCard.target = g.DrawFromDeck()
		} else {
			g.nextCard.target = buf
		}
	}
}

func (g *Game) PeekNextCard() Three {
	return g.nextCard.target
}

func (g *Game) TestDeck() bool {
	count1 := 0
	count2 := 0
	count3 := 0
	for _, c := range g.deck {
		switch c {
		case GetThree(1):
			count1++
		case GetThree(2):
			count2++
		case GetThree(3):
			count3++
		default:
			return false
		}
	}
	if count1 > 4 || count2 > 4 || count3 > 4 {
		return false
	}
	/*
	fmt.Println(count1)
	fmt.Println(count2)
	fmt.Println(count3)
	fmt.Println("--------")
	*/
	return true
}

func createCards (num int, numOfThree int) []*Card {
	result := []*Card{}

	for i := 0;i < num;i++ {
		result = append(result, &Card{GetThree(numOfThree), Pos{}})
	}

	return result
}

func CreateCard (num int) *Card {
	result := &Card{GetThree(num), Pos{}}

	return result
}

func (g *Game) InitializeBoard() {
	cards := []*Card{}
	const size = BOARD_SIZE

	//Check carefully
	for i := 0; i < 8; i++ {
		cards = append(cards, CreateCard(g.DrawFromDeck().Value()))
	}

	permIndex := rand.Perm(BOARD_SIZE * BOARD_SIZE)
	positions := GetAllCombPos()

	for index, v := range cards {
		cards[index] = &Card{v.target, positions[permIndex[index]]}
	}

	param := make([]ValueDefiner, len(cards))
	for index, v := range cards {
		param[index] = v
	}
	board,_ := GetInitializedBoard(param)
	g.b = board
}

func calcScore(t Board) int {
	area := t.GetArea()
	score := 0

	for _, v := range area {
		score += v.Value().Score()
	}

	return score
}

func (g *Game) CalcScore() int {
	return calcScore(g.b)
}

func makeMover(generator func () Three) Mover {
	return func(indices []int) (int, Three) {
		if len(indices) == 1 {
			return indices[0], generator()
		}
		target := rand.Intn(len(indices))
		return indices[target], generator()
	}
}

func (g *Game) IsOver() bool {
	return !g.b.CanSomeMove()
}

func (g *Game) Move(dir Direction) bool {
	generator := func () Three {
		return g.nextCard.Value()
	}
	if g.IsOver() {
		return false
	}
	switch dir {
	case UP:
		if g.b.Move(UP, makeMover(generator)) {
			g.UpdateNextCard()
		}
		return true
	case DOWN:
		if g.b.Move(DOWN, makeMover(generator)) {
			g.UpdateNextCard()
		}
		return true
	case LEFT:
		if g.b.Move(LEFT, makeMover(generator)) {
			g.UpdateNextCard()
		}
		return true
	case RIGHT:
		if g.b.Move(RIGHT, makeMover(generator)) {
			g.UpdateNextCard()
		}
		return true
	}
	return false
}

func (g *Game) GetState() [255]bool {
	sliceFk := [16]bool{}
	arr := [255]bool{}
	table := [15]int{1,2,3,6,12,24,48,96,192,384,768,1536,3072,6144,12288}
	iter := 0
	for fkthis := 0; fkthis < 15; fkthis++ {
		sliceFk = g.b.printThree(table[fkthis])
		for i := 0; i < 16; i++ {
			arr[i+iter] = sliceFk[i]
		}
		sliceFk = [16]bool{}
		iter += 16
	}
	iter = 0
	for fkthis := 0; fkthis < 15; fkthis++ {
		if(g.nextCard.Value().Value() == table[fkthis]) {
			arr[240+iter] = true
		}
		iter += 1
	}

	return arr
}

func (g *Game) GetFloatState() []float64 {
	arr := g.GetState()
	var ret []float64
	for i:=0;i<255;i++ {
		if(arr[i] == true) {
			ret = append(ret, float64(1))
		} else {
			ret = append(ret, float64(0))
		}
	}
	return ret
}

func (g *Game) Debug (numOfThree int) [16]bool {
	return g.b.printThree(numOfThree)
}

func (g *Game) MaxValue() int {
	return g.b.maxValue()
}

func (g *Game) CanMove(dir Direction) bool {
	return g.b.CanMove(dir)
}

func (g *Game) PrintBoard() {
	for i:=0;i<9;i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	fmt.Print("|")
	kek, _ := g.b.At(Pos{0,0})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{1,0})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{2,0})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{3,0})
	fmt.Print(kek.Value())
	fmt.Print("|")
	fmt.Print("\n")
	fmt.Print("|")
	kek, _ = g.b.At(Pos{0,1})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{1,1})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{2,1})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{3,1})
	fmt.Print(kek.Value())
	fmt.Print("|")
	fmt.Print("\n")
	fmt.Print("|")
	kek, _ = g.b.At(Pos{0,2})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{1,2})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{2,2})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{3,2})
	fmt.Print(kek.Value())
	fmt.Print("|")
	fmt.Print("\n")
	fmt.Print("|")
	kek, _ = g.b.At(Pos{0,3})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{1,3})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{2,3})
	fmt.Print(kek.Value())
	fmt.Print("|")
	kek, _ = g.b.At(Pos{3,3})
	fmt.Print(kek.Value())
	fmt.Print("|")
	fmt.Print("\n")
	for i:=0;i<9;i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}
