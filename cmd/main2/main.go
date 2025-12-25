package main

import (
	"fmt"
	"math/rand"
	"strconv"

	ss "github.com/trng-tr/golab5/generic/generic2"
)

func main() {
	var s1 = ss.Stack[string]{} // empty slice of string
	for i := 1; i <= 10; i++ {
		s1.Push("Go00000000-" + strconv.Itoa(i))
	}
	fmt.Println("slice of string", s1)
	for i := 0; i < 4; i++ {
		removed, err := s1.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("removed", removed)
	}
	fmt.Println("remaining", s1)

	s2 := ss.Stack[ss.Customer]{} //empty slice of object
	for i := 1; i <= 5; i++ {
		s2.Push(ss.CustomerInput())
	}
	for _, c := range s2.Items {
		formatedC := c.String()
		fmt.Println(formatedC)
	}

	s3 := ss.Stack[float64]{}
	for i := 1; i <= 10; i++ {
		s3.Push(1 + rand.Float64()*(10-1))
	}
	fmt.Println("slice of float64", s3)
}
