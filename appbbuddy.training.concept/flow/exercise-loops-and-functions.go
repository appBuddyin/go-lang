
package flow


import (
	"fmt"
)

func Sqrt(x float64) float64 {
	sNumber := 1.0
	for i := 0; i < 10; i++ {
		if sNumber*sNumber < x {
			sNumber = sNumber + (x-sNumber)/2
		} else {
			sNumber = sNumber - (x-sNumber)/2
		}
	}
	return sNumber
}

func exercise() {
	fmt.Println(Sqrt(16))
}
