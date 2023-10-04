package main

import (
	"flag"
	"fmt"
	"math"
)

func bounds(target float64, accuracy float64) (float64, float64) {
	return target - accuracy, target + accuracy
}

func main() {
	root := flag.Int("rt", 1, "a non-negative integer that is the root value (2 = square root, 3 = cube root, etc) ")
	accuracy := flag.Float64("acc", 1, "a non-negative number which describes how far away the found root may be from the actual root")
	flag.Parse()

	fmt.Printf("%f, %d, %f\n", *operand, *root, *accuracy)
	var operandFixed float64
	foundRoot := 1.0
	if *operand == 0 {
		fmt.Println(0)
		return
	}
	if *operand < 0 {
		operandFixed = *operand * -1
	} else {
		operandFixed = *operand
	}
	lowerBound, higherBound := bounds(operandFixed, *accuracy)
	for {
		exponentiation := math.Pow(foundRoot, float64(*root))
		fmt.Printf("%f ** %f == %f\n", foundRoot, float64(*root), exponentiation)
		if exponentiation < higherBound && exponentiation > lowerBound {
			fmt.Printf("found: %f, actualFound: %f", foundRoot, math.Pow(foundRoot, float64(*root)))
			return
		}

		foundRoot = foundRoot - ((math.Pow(foundRoot, float64(*root)) - operandFixed) / (float64(*root) * foundRoot))
		//fmt.Println(foundRoot)
	}
}
