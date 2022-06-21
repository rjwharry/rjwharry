package main

import "fmt"

func a() {
	fmt.Println("Inside  a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited")
	fmt.Println("Exitting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()")
	fmt.Println("Exitting b()")
}

func main() {
	a()
	fmt.Println("main() ended")
}
