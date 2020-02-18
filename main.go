package main

import "fmt"

type Cutter interface {
	Cut()
}

type Cooker interface {
	Cook()
}

type X struct {

}

func (x X) Cut() {
	fmt.Println("Chop Chop Chop...!")
}