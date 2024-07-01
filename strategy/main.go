package main

func main() {
	lru := NewLFU()
	cache := NewCache(lru)
	cache.Evict()
}
