package main

// for loop
// func main() {
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(i)
// 	}
// }

// continue keyword
// func main() {
// 	for i := 0; i < 5; i++ {
// 		if i == 3 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }

// break keyword
// func main() {
// 	for i := 0; i < 5; i++ {
// 		if i == 3 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// }

// nested loops
// func main() {
// 	adj := [2]string{"big", "tasty"}
// 	fruits := [3]string{"apple", "orange", "banana"}
// 	for i := 0; i < len(adj); i++ {
// 		for j := 0; j < len(fruits); j++ {
// 			fmt.Println(adj[i], fruits[j])
// 		}
// 	}
// }

// range keyword
// func main() {
// 	fruits := [3]string{"apple", "orange", "banana"}
// 	for idx, val := range fruits {
// 		fmt.Printf("%v\t%v\n", idx, val)
// 	}
// }

// omit indexes, only values
// func main() {
// 	fruits := [3]string{"apple", "orange", "banana"}
// 	for _, val := range fruits {
// 		fmt.Printf("%v\n", val)
// 	}
// }

// only indexes, no values
// func main() {
// 	fruits := [3]string{"apple", "orange", "banana"}

// 	for idx, _ := range fruits {
// 		fmt.Printf("%v\n", idx)
// 	}
// }
