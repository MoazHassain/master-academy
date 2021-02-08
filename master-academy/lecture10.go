package main

import "fmt"

func main() {
	var student [3]string

	student[0] = "moaz"
	student[1] = "miraz"
	student[2] = "mitul"

	fmt.Println(student)
	fmt.Println(len(student))
	fmt.Println(student[2])
}
