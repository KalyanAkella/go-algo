package main

import (
	"fmt"
	"sort"
)

type Person struct {
	height, weight uint
}

func (this *Person) CanBeAbove(that *Person) bool {
	return this.height <= that.height && this.weight <= that.weight
}

type Persons []*Person

func (this Persons) Len() int {
	return len(this)
}

func (this Persons) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByHeight struct{ Persons }

func (this ByHeight) Less(i, j int) bool {
	return this.Persons[j].height <= this.Persons[i].height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func computeMaxHeightWithPersonAtBottom(persons Persons, bottomIndex int) int {
	maxHeight := 0
	bottomPerson := persons[bottomIndex]
	for personIndex := bottomIndex + 1; personIndex < persons.Len(); personIndex++ {
		currentPerson := persons[personIndex]
		if currentPerson.CanBeAbove(bottomPerson) {
			maxHeight = max(maxHeight, computeMaxHeightWithPersonAtBottom(persons, personIndex))
		}
	}
	maxHeight += 1
	return maxHeight
}

func computeHeight(persons Persons) int {
	maxHeight := 0
	for i := range persons {
		maxHeight = max(maxHeight, computeMaxHeightWithPersonAtBottom(persons, i))
	}
	return maxHeight
}

//TODO: Memoise
func main() {
	persons := []*Person{
		&Person{65, 100},
		&Person{70, 150},
		&Person{56, 90},
		&Person{75, 190},
		&Person{60, 95},
		&Person{68, 110},
	}
	sort.Sort(ByHeight{persons})
	result := computeHeight(persons)
	fmt.Println(result)
}
