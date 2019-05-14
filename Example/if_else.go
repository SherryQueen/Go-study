package Example

import "fmt"

func IfElse() {
	for i := 1; i < 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}

		if i%4 == 0 {
			fmt.Printf("%d is divisible by 4\n", i)
		}
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
