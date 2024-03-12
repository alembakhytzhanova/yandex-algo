package main

import "fmt"

func main() {
	g1, g2 := 0, 0

	var a, b, c, d, mode int
	fmt.Scan(&a)
	g1 += a
	fmt.Scan(&b)
	g2 += b
	fmt.Scan(&c)
	g1 += c
	fmt.Scan(&d)
	g2 += d

	fmt.Scan(&mode)

	switch {
	case g1 > g2:
		fmt.Println(0)
	case mode == 2 && a > d:
		fmt.Println(g2 - g1)
	case mode == 2:
		fmt.Println(g2 - g1 + 1)
	case g2-g1+c > b:
		fmt.Println(g2 - g1)
	default:
		fmt.Println(g2 - g1 + 1)
	}
}
