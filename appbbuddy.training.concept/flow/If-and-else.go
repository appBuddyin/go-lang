// Variables declared inside
// an if short statement are also
// available inside any of the else blocks.

package flow


import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, only in if else block
	return lim
}

func ifAndElse() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
