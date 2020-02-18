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

type pan struct {
}

func (k knife) Cut() {
	fmt.Println("Chop Chop Chop...!")
}

func (p pan) Cook() {
	fmt.Println("bla bla bla...!")
}
