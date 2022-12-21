package main

import (
	"fmt"

	"github.com/Msaorc/Go-Packages/packaging/math"
)

func main() {
	result := math.Math{35, 7}
	fmt.Println(result.Add())
}
