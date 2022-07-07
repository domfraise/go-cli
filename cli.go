package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, world.")
	var x float64 = 10
	y := 20.6
	fmt.Println(x + y)
	//slices
	array := []int{1, 2, 3}
	fmt.Println(array)
	newArray := append(array, 4, 5, 6)

	fmt.Println(array)
	fmt.Println(newArray)

	//maps
	myMap := map[string]string{}
	newmap, _, that := addItemToMap(myMap, "key", "value")
	addItemToMap(myMap, "key2", "value2")
	fmt.Println(newmap, that)

	//functions
	fmt.Println(namedReturnValue(1))

	// memory pointers
	var list []string
	list = append(list, "this")
	p := passByReference(&list)
	fmt.Println(list, p)

	//loops
	thisMap := map[string]string{"key1": "value1", "key2": "value2"}
	for k, v := range thisMap {
		fmt.Printf("key=%s, value=%s\n", k, v)
	}

	for _, v := range thisMap {
		fmt.Printf("value=%s\n", v)
	}

	myNum := 20
	xBig := func() bool {
		return myNum > 1000
	}
	fmt.Println(xBig())
	myNum = 100000
	fmt.Println(xBig())

	//objects
	myPair := pair{3, 5}
	fmt.Println(myPair.MyMethod())
	var myInterface MyInterface
	myInterface = myPair
	fmt.Println(myInterface.MyMethod())

	//errors
	myMap2 := map[int]string{3: "three", 4: "four"}
	if x, ok := myMap2[1]; !ok {
		fmt.Println("Index not here!")
	} else {
		fmt.Println(x)
	}

	if _, err := strconv.Atoi("non-int"); err != nil {
		fmt.Println(err)
	}

	//concurrency
	c := make(chan int)

	go increment(0, c)
	go increment(10, c)
	go increment(-110, c)

	fmt.Println(<-c, <-c, <-c)

	cs := make(chan string)
	css := make(chan chan string)
	go func() { cs <- "wordy" }()
	go func() { c <- 84 }()

	select {
	case i := <-c:
		fmt.Printf("it is a %T", i)
	case <-cs:
		fmt.Println("its a string")
	case <-css:
		fmt.Println("didn't happen")
	}

}
func increment(i int, c chan int) {
	c <- i + 1
}

type MyInterface interface {
	MyMethod() string
}
type pair struct {
	x, y int
}

func (p pair) MyMethod() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func passByReference(incomingSlice *[]string) []string {
	*incomingSlice = append(*incomingSlice, "that")

	return *incomingSlice
}

func namedReturnValue(number int) (z string) {

	if number == 1 {
		z = "one"
	} else {
		z = "something else"
	}
	return
}

func addItemToMap(myMap map[string]string, key string, value string) (map[string]string, string, string) {
	myMap[key] = value
	return myMap, key, value
}
