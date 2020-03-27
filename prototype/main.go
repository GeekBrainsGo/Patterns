package main

func main() {

	a := A{}
	n := 5
	a.N = &n

	b := a.Copy()
	b.N = 6

}

type A struct {
	N *int
	A *A
}

func(a *A) Copy() A {
	return A{
		N: &(*a.N),
	}
}
