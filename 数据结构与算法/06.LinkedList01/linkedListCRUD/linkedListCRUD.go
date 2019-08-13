package main

import (
	"fmt"
	"testing"
)

type linkedList struct {
	Data int
	Next *linkedList
}

func main {
	var p *linkedList
	p.Data = 1
	h := p
	for a := 1; a < 20; a++ {
		p.Data = a
		p = p.Next
	}
	p = nil
	fmt.Println(*h)
}

func getData(Data int) i int{
     
}
