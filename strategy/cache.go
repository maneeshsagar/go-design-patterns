package main

type EvictionPolicy interface {
	Evict()
}

type Cache struct {
	Capacity       int
	Storage        map[string]string
	EvictionPolicy EvictionPolicy
}

func NewCache(e EvictionPolicy) *Cache {
	return &Cache{
		Capacity:       10,
		Storage:        make(map[string]string),
		EvictionPolicy: e,
	}
}

func (c *Cache) Evict() {
	c.EvictionPolicy.Evict()
}
