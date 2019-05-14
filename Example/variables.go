package Example

import "fmt"

func Variables() {
	var a = "initial"
	fmt.Println("a: ", a)

	var b, c int = 1, 2
	fmt.Println("b:", b, " c:", c)

	var d = true
	fmt.Println("d:", d)

	var e int
	fmt.Println("e:", e)

	f := "short"
	fmt.Println("f", f)
}
