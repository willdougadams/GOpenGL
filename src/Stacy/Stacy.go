package Stacy

import (
	"fmt"
)

type Stacy struct {
	prettiness int
	height string
	titties bool
}

func (stacy *Stacy) Init() *Stacy {
	stacy.prettiness = 100
	stacy.height = "small"
	stacy.titties = true
	return stacy
}

func (stacy *Stacy) IsGreat() {
	fmt.Printf("Stacy is Great!\n")
}
