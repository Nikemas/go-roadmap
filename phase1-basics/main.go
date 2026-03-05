package main

import "fmt" 

func main () {
	var name string = "Shadow"
	var age int = 20
	var height float64 = 1.80
	var isStudent bool = false

	city := "Bishkek"
	score := 95

	var x, y int = 10, 20

	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Возраст: %d\n", age)
	fmt.Printf("Рост: %.2f\n", height)
	fmt.Printf("Студент: %v\n", isStudent)
	fmt.Printf("Город: %s\n", city)
	fmt.Printf("Очки: %d\n", score)
	fmt.Printf("X=%d, Y=%d\n", x, y)
	
	city = "Osh"
	fmt.Printf("Новый город: %s\n", city)
	constants()
}

func constants(){
	const Pi = 3.14159
	const AppName = "ShadowApp"
	const MaxUser = 100

	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
	)

	fmt.Printf("Pi: %f\n", Pi)
	fmt.Printf("App: %s\n", AppName)
	fmt.Printf("Max Use: %d\n", MaxUser)
	fmt.Printf("Monday=%d, Fiday=%d\n", Monday, Friday)
}