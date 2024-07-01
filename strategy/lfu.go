package main

import "fmt"

type LFU struct {
}

func NewLFU() *LFU {
	return &LFU{}
}

func (f *LFU) Evict() {
	fmt.Println("evicting by LRU")
}
