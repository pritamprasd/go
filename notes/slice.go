package main

import "fmt"
// NOtice how address changes, when slice gets resized
func main() {
	s := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, slice %v, \t%p\n", cap(s), len(s),s, s)
	}
}