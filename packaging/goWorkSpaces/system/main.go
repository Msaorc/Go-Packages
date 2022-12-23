package main

import (
	"fmt"

	"github.com/Msaorc/Go-Packages/packaging/packageLocal/math"
)

func main() {
	r := math.Math{5, 35}
	fmt.Println(r.Add())
}
