package main

import "fmt"

type Observable interface {
	Subscribe(o Observer)
	UnSubscrive(o Observer)
	NotifyAll()
}

type Item struct {
	Observers []Observer
}

func (i *Item) Subscribe(o Observer) {
	i.Observers = append(i.Observers, o)
}

func (i *Item) UnSubscrive(o Observer) {
	i.Observers = removeFromSlice(i.Observers, o)
}

func (i *Item) NotifyAll() {
	for _, oberver := range i.Observers {
		oberver.Update()
	}
}

// this is not from the function
func (i *Item) UpdateAvailibility() {
	fmt.Println("item Added")
	i.NotifyAll()
}

func NewItem() *Item {
	return &Item{}
}

func removeFromSlice(observers []Observer, observer Observer) (finalObserverSlice []Observer) {
	for index, obs := range observers {
		if obs.GetId() == observer.GetId() {
			finalObserverSlice = append(finalObserverSlice, observers[:index]...)
			finalObserverSlice = append(finalObserverSlice, observers[index+1:]...)
			return
		}
	}
	return
}
