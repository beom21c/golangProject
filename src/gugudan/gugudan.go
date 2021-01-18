package main

import "fmt"

func main() {
	for dan := 1; dan <= 9; dan++ {
		fmt.Printf("%d단\n", dan)
		// continue를 사용하면 마지막 } 으로 이동한다.
		// 5단은 건너뜀
		if dan == 5 {
			continue
		}

		for j := 1; j <= 9; j++ {
			fmt.Printf("%d * %d = %d\n", dan, j, dan*j)
		}

		fmt.Println()
	}
}
