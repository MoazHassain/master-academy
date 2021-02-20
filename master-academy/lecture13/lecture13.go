package main

import "fmt"

// method-1
// func add(x int, y int) int {
// 	r := x + y
// 	return r
// }

// method-2
// func add(x, y int) int {
// 	r := x + y
// 	return r
// }

// method-3
// func add(x, y int) (r int) {
// 	r = x + y
// 	return r
// }

// func main() {
// 	a := add(20, 30)
// 	fmt.Println(a)
// }

func main() {
	var result *float32
	var mark float32 = 59.90
	result = &mark
	fmt.Println(result, mark, *result)
}
