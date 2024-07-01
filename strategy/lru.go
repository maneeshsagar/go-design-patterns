package main

import "fmt"

type LRU struct {
}

func NewLRU() *LRU {
	return &LRU{}
}

func (f *LRU) Evict() {
	fmt.Println("evicting by LRU")
}
