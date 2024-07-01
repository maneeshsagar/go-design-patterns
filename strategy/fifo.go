package main

import "fmt"

type Fifo struct {
}

func NewFifo() *Fifo {
	return &Fifo{}
}

func (f *Fifo) Evict() {
	fmt.Println("evicting by Fifo")
}
