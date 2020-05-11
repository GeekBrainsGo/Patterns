package main

import "fmt"

func main() {
	var q Queue
	q.Add("1")
	fmt.Println(q.Get())
	fmt.Println(q.Get())
	fmt.Println(q.Get())
	fmt.Println(q.Get())
	q.Add("1")
	q.Add("2")
	q.Add("3")
	fmt.Println(q.Get())
	fmt.Println(q.Get())
	fmt.Println(q.Get())
	fmt.Println(q.Get())

	var s Stack
	s.Add("1")
	fmt.Println(s.Get())
	fmt.Println(s.Get())
	fmt.Println(s.Get())
	fmt.Println(s.Get())
	s.Add("1")
	s.Add("2")
	s.Add("3")
	fmt.Println(s.Get())
	fmt.Println(s.Get())
	fmt.Println(s.Get())
	fmt.Println(s.Get())

}
