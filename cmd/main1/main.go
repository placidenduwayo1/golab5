package main

import (
	"fmt"
	"math/rand"

	"github.com/trng-tr/golab5/generic/generic1"
)

func main() {
	var s0 = make([]byte, 0)
	var s1 = make([]int16, 0)
	var s2 = make([]int32, 0)
	var s3 = make([]int64, 0)
	var s4 = make([]float32, 0)
	var s5 = make([]float64, 0)

	for i := 0; i < 6; i++ {
		s0 = append(s0, byte(rand.Intn(50-20+1)+20))
	}
	for i := 0; i < 5; i++ {
		s1 = append(s1, int16(rand.Intn(100-50+1)+50))
	}
	for i := 0; i < 7; i++ {
		s2 = append(s2, int32(rand.Intn(100-10+1)+10))
	}
	for i := 0; i < 4; i++ {
		s3 = append(s3, int64(rand.Intn(1_000_000-100_000+1)+100_000))
	}
	for i := 0; i < 5; i++ {
		s4 = append(s4, 10+rand.Float32()*(20-10))
	}
	for i := 0; i < 4; i++ {
		s5 = append(s5, 10+rand.Float64()*(100-10))
	}

	fmt.Println("s0: ", s0)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("s3: ", s3)
	fmt.Println("s4: ", s4)
	fmt.Println("s5: ", s5)
	fmt.Println("Somme slice 0: ", generic1.ComputeSum(s0))
	fmt.Println("Some slice 1: ", generic1.ComputeSum(s1))
	fmt.Println("Some slice 2: ", generic1.ComputeSum(s2))
	fmt.Println("Some slice 3: ", generic1.ComputeSum(s3))
	fmt.Println("Some slice 4: ", generic1.ComputeSum(s4))
	fmt.Println("Some slice 5: ", generic1.ComputeSum(s5))
}
