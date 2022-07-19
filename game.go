package main

import (
	"fmt"
	"strconv"
)

type Game struct {
	id   int
	next *Game

	m int
	n int
}

func New(m, n int) *Game {
	if n < 1 {
		panic("太短没得玩")
	}
	first := &Game{id: 0, m: m, n: n}
	current := first
	current.next = first
	for i := 1; i < n; i++ {
		g := &Game{id: i, m: m, n: n}
		current.next = g
		current = g
		current.next = first
	}
	return first
}

func (g *Game) Play() (int, string) {
	order := ""
	first, last := g, g
	for {
		if last.next == first {
			break
		}
		last = last.next
	}

	for {
		for i := 1; i <= g.m-1; i++ {
			first = first.next
			last = last.next
		}
		order = order + strconv.Itoa(first.id) + " "
		first = first.next
		last.next = first
		if last == first {
			break
		}
	}
	return first.id, order
}
func main() {
	//0,1,2,3,4,5,6,7,8,9
	//3,7,1,6,2,9,8,0,5,4
	m, n := 4, 10
	win, order := New(m, n).Play()
	fmt.Printf("胜利者: %d, 淘汰顺序: %v", win, order)
}
