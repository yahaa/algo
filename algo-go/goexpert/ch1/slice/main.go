package main

import "fmt"

func q1() {
	var array [10]int
	slice := array[5:6]

	fmt.Printf("len of slice: %d\n", len(slice))
	fmt.Printf("cap of slice: %d\n", cap(slice))
	fmt.Printf("&array[5]==&slice[5] is %v\n", &array[5] == &slice[0])

}

func q2() {
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)

	newSlice := addElement(slice, 6)

	fmt.Printf("&slice[5]==&newSlice[5] is %v\n", &slice[0] == &newSlice[0])
}

func addElement(slice []int, e int) []int {
	return append(slice, e)
}

func q3() {
	orderLen := 5

	order := make([]uint16, 2*orderLen)

	pollorder := order[:orderLen:orderLen]

	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
}

func main() {
	q3()
}
