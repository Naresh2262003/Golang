package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	var data float64 = float64(num)

	data += 1.2
	fmt.Println(data)

	str := strconv.Itoa(num)
	// str += 2
	fmt.Println(str)
	fmt.Printf("Type of str is %T\n", str)

	strToNum, _ := strconv.Atoi(str)

	fmt.Println(strToNum)

	strToNum += 10
	fmt.Printf("Type of str is %T and value is %d\n", strToNum, strToNum)

	floatStr := "3.14"
	strFloat, _ := strconv.ParseFloat(floatStr, 64)

	strFloat += 2.06
	fmt.Println(strFloat)

}
