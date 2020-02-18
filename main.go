package main

import (
	"fmt"
)

type Cutter interface {
	Cut()
}

type Cooker interface {
	Cook()
}

//The chef has the ability to cut and cook
type chef struct {
	Cutter
	Cooker
}

type knife struct {
}

type chainsaw struct {
}

type oven struct {
}

type pan struct {
}

//Knife becomes an implementer of Cutter
func (k knife) Cut() {
	fmt.Println("Chop Chop Chop...!")
}

//Chainsaw becomes an implementer of Cutter
func (c chainsaw) Cut() {
	fmt.Println("ZRRRRRRRRRRR!!!!")
}

//Pan becomes an implementer of Cooker
func (p pan) Cook() {
	fmt.Println("sweeeeesh!")
}

//Oven becomes an implementer of Cooker
func (o oven) Cook() {
	fmt.Println("BHHHHHHHMMMMMMMMM!")
}

//Chef does his work that is cut and cook irrespective of the implementers of the interfaces
func (c chef) Do() {
	c.Cut()
	c.Cook()
}

