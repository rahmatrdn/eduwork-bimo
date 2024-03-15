package main

import "fmt"

func main() {
	var radius float64 = 5.0
	var area float64

	const pi float64 = 3.14

	area = (radius * radius) * pi

	fmt.Println(area)
}
