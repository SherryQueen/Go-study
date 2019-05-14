package Example

import (
	"fmt"
	"math"
)

const s string = "hello world"

func Constant() {
	fmt.Println(s)

	const n = 50000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
