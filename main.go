package main

import "fmt"

type butter interface {
	but()
}

type booker interface {
	book()
}

type X struct {

}

func (x X) Cut() {
	fmt.Println("Chop Chop Chop...!")
}