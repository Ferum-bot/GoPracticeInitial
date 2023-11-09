package main

import "fmt"

const (
	maxY = 100
	maxX
)

type Item struct {
	X int32
	Y int32
}

func NewItem(x, y int32) (*Item, error) {
	if x < 0 || x > maxX {
		return nil, fmt.Errorf("X parameter is out of bound! Passed %d, max bound %d", x, maxX)
	}
	if y < 0 || y > maxY {
		return nil, fmt.Errorf("Y parameter is out of bound! Passed %d, max bound %d", y, maxY)
	}

	createdItem := Item{
		X: x,
		Y: y,
	}

	// Go compiler does "escape analysis" and will allocate createdItem on the Heap
	return &createdItem, nil
}

func (item *Item) Move(x, y int32) {
	item.X = x
	item.Y = y
}

type Player struct {
	Name string
	Item Item
	Keys []Key
}

type mover interface {
	Move(x, y int32)
}

func moveAll(movers []mover, x, y int32) {
	for _, mover := range movers {
		mover.Move(x, y)
	}
}

func (player *Player) FoundKey(key Key) error {
	if key < Jade || key >= invalidKey {
		return fmt.Errorf("target key %d is invliad", key)
	}

	if !containsKey(player.Keys, key) {
		player.Keys = append(player.Keys, key)
	}
	return nil
}

func containsKey(keys []Key, targetKey Key) bool {
	for _, key := range keys {
		if key == targetKey {
			return true
		}
	}
	return false
}

type Key byte

// Go version of "enum"
const (
	Jade Key = iota + 1
	Cooper
	Crystal
	invalidKey
)

func (key Key) String() string {
	switch key {
	case Jade:
		return "jade"
	case Cooper:
		return "cooper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", key)
}

func main() {
	var firstItem Item
	fmt.Println(firstItem)
	fmt.Printf("First item: %#v \n", firstItem)

	secondItem := Item{1, 2}
	fmt.Printf("Second item: %#v \n", secondItem)

	thirdItem := Item{
		Y: 10,
		X: 12,
	}
	fmt.Printf("Third item: %#v \n", thirdItem)

	fmt.Println(NewItem(1, 2))
	fmt.Println(NewItem(10, -20))

	secondItem.Move(8, 9)
	fmt.Printf("Second item(moved): %#v", secondItem)

	firstPlayer := Player{
		Name: "Ferum-bot",
		Item: Item{112, 113},
	}
	fmt.Printf("First player: %#v \n", firstPlayer)
	fmt.Printf("First player: %#v \n", firstPlayer.Item.X)

	myMovers := []mover{
		&firstItem,
		&secondItem,
		&thirdItem,
	}

	moveAll(myMovers, 23, 23)
	for _, mover := range myMovers {
		fmt.Println(mover)
	}

	myKey := Jade
	fmt.Println("k:", myKey)
	fmt.Println("k:", Key(17))

	err := firstPlayer.FoundKey(Jade)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(firstPlayer)
	}
}
