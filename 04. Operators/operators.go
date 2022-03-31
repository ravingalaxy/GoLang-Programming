package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Maths operators:")
	fmt.Println("10+5 = ", 10+5)
	fmt.Println("10-5 = ", 10-5)
	fmt.Println("10*5 = ", 10*5)
	fmt.Println("10/5 = ", 10/5)
	fmt.Println("10%5 = ", 10%5)

	var base, exponent float64
	base = 2.0
	exponent = 3.0
	fmt.Println("2 to the power 3 = ", math.Pow(base, exponent))

	num := 10
	num++
	fmt.Println("10+1 = ", num)
	num2 := 10
	num2--
	fmt.Println("10-1 = ", num2)
	fmt.Println("Assignment Operator:")
	mynum := 10
	mynum += 5
	fmt.Println("10+5 = ", mynum)

	mynum1 := 10
	mynum1 -= 5
	fmt.Println("10-5 = ", mynum1)

	mynum2 := 10
	mynum2 *= 5
	fmt.Println("10*5 = ", mynum2)

	mynum3 := 10
	mynum3 /= 5
	fmt.Println("10/5 = ", mynum3)

	mynum4 := 10
	mynum4 %= 5
	fmt.Println("10%5 = ", mynum4)

	fmt.Println("Comparision operator")
	fmt.Println("5==5", 5 == 5)
	fmt.Println("5!=5", 5 != 5)
	fmt.Println("5>5", 5 > 5)
	fmt.Println("5<5", 5 < 5)
	fmt.Println("5<=5", 5 <= 5)
	fmt.Println("5>=5", 5 >= 5)
}
