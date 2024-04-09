package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"aa", "bb", "cc"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// message, err := greetings.Hello("Gladys")
	fmt.Println(message)

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	f := "apple"
	fmt.Println("apple:", f)

	var b1, c1 int = 1, 2
	fmt.Println("apple1:", b1, c1)

	var f1 string = "123"
	fmt.Println("apple2:", f1)

	f2 := "123"
	fmt.Println("apple3:", f2)

	x := 10
	ptr := &x
	fmt.Println(ptr)
	fmt.Println(*ptr)

	var p1 *Person
	p2 := &Person{Name: "jakson", Age: 10}
	fmt.Println(p1 == nil) // true
	fmt.Println(p2 == nil) // false

}

type Person struct {
	Name string
	Age  int
}
