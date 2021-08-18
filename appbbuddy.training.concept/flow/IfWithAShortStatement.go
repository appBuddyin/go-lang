// Like for, the if statement can start with a short
// statement to execute before the condition.
// Variables declared by the statement are
// only in scope until the end of the if

package flow

import (
	"fmt"
	"math"
)

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func ifWithShort() {
	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)
}
