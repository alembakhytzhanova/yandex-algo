package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, ans, a int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a)
		ans += a / 4
		if a%4 <= 2 {
			ans += a % 4
		} else {
			ans += 2
		}
	}

	fmt.Println(ans)
}
