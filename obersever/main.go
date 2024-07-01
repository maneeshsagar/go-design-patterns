package main

func main() {
	item := NewItem()
	person1 := NewPerson(1, "Maneesh")
	person2 := NewPerson(2, "Himanshu")

	item.Subscribe(person1)
	item.Subscribe(person2)
	item.UpdateAvailibility()
	item.UnSubscrive(person1)
	item.UpdateAvailibility()
}
