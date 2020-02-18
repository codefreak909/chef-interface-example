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

type knife struct {
}

type chainsaw struct {
}

type oven struct {
}

type pan struct {
}


func (k knife) Cut() {
	fmt.Println("Chop Chop Chop...!")
}

func (c chainsaw) Cut() {
	fmt.Println("ZRRRRRRRRRRR!!!!")
}

func (p pan) Cook() {
	fmt.Println("sweeeeesh!")
}


func (o oven) Cook() {
	fmt.Println("BHHHHHHHMMMMMMMMM!")
}