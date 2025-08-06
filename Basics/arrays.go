package main

// Arrays
// there are two ways to declare an array:

// 1. With the var keyword:
// func main() {
// 	var arr1 = [3]int{1, 2, 3}
// 	arr2 := [5]int{4, 5, 6, 7, 8}

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

// 2. With the := sign:
// func main() {
// 	var arr1 = [...]int{1, 2, 3}
// 	arr2 := [...]int{4, 5, 6, 7, 8}

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

// Access Elements of an Array
// func main() {
// 	prices := [3]int{10, 20, 30}

// 	fmt.Println(prices[0])
// 	fmt.Println(prices[2])
// }

// Change Elements of an Array
// func main() {
// 	prices := [3]int{10, 20, 30}

// 	prices[2] = 50
// 	fmt.Println(prices)
// }

// Array Initialization
// func main() {
// 	arr1 := [5]int{}              //not initialized
// 	arr2 := [5]int{1, 2}          //partially initialized
// 	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// 	fmt.Println(arr3)
// }

// Initialize Only Specific Elements
// func main() {
// 	arr1 := [5]int{1: 10, 2: 40, 4: 30}

// 	fmt.Println(arr1)
// }
