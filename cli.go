package main

import "fmt"

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
