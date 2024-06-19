package main

import "fmt"

func divide(p, q float64) (float64, error) {
	if q == 0 {
		return 0, fmt.Errorf("denominator is 0")
	}
	return p / q, nil
}

func main() {
	ans, _ := divide(10, 0)
	fmt.Printf("%.2f\n", ans)
}
