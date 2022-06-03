package main


import (
    "fmt"
    "reflect"
)


func hello() {

    var name string;
    var age int;
    var price float64;

    var city = "Warsaw"
    var length = 3.5

    quantity := 4
    fruit := "apple"

    allApples := 3
    eatenApples := 0.5
    applesLeft := float64(allApples) - eatenApples

    fmt.Println("Name: ",name,"-", "Age: ",age, "_", "Price: ",price, "_")
    fmt.Println("City: ",city, "_", "Length: ",length, "_", "Qtd: ",quantity, "_", "Fruit: ",fruit)

    fmt.Println(applesLeft)
    fmt.Println("Hello, Go!")

    fmt.Println(reflect.TypeOf("Hello, Go!"))
    fmt.Printf("%T\n", "Hello, Go!")

    fmt.Println("------------------------------------")

    var test interface{}
	test = 1
	fmt.Printf("%T: %v\n", test, test)
	test = "don't do this"
	fmt.Printf("%T: %v\n", test, test)

    fmt.Println("------------------------------------")
    fmt.Println("------------------------------------")

    i := 1
	p1 := &i
	var p2 *int
	fmt.Printf("p1: %v. p2: %v\n", p1, p2)
	s1 := []int{}
	var s2 []int
	fmt.Printf("s1 nil? %v, s2 nil? %v\n", s1 == nil, s2 == nil)
	m1 := map[int]string{}
	var m2 map[int]string
	fmt.Printf("m1 nil? %v, m2 nil? %v\n", m1 == nil, m2 == nil)
	c1 := make(chan int)
	var c2 chan int
	fmt.Printf("c1 nil? %v, c2 nil? %v\n", c1 == nil, c2 == nil)
	var i1 interface{} = 1
	var i2 interface{}
	fmt.Printf("i1 nil? %v, i2 nil? %v\n", i1 == nil, i2 == nil)
}