package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("The Number is", rand.Intn(10), math.Nextafter(20, 30))
}
