package main

func main() {

	a := A{}.Build(0) //< Способ 1
	b := BuildA(0) //< Способ 2

}

type A struct {
	N int
}

// Отличие билдер от конструктора - билдер сложнее
// Билдер может сходить в БД или обратиться к другим объектам
func(a *A) Build(n int) A {
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
