package main

import "fmt"

func main() {

	a := A{}
	a1 := a.Build(0) //< Способ 1
	a2 := BuildA(0)  //< Способ 2
	fmt.Printf("%v\n%v\n", a1, a2)
}

type A struct {
	N int
}

// Отличие билдер от конструктора - билдер сложнее
// Билдер может сходить в БД или обратиться к другим объектам
func (a *A) Build(n int) A {
	return A{
		N: n,
	}
}

func BuildA(n int) A {
	// Ходим в БД
	return A{
		N: n,
	}
}
