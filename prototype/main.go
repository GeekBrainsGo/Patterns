package main

func main() {

	a := A{}
	n := 5
	a.N = &n

	b := a.Copy()
	m := 6
	b.N = &m

}

type A struct {
	N *int
	A *A
}

func (a *A) Copy() A {
	return A{
		N: &(*a.N),
	}
}
