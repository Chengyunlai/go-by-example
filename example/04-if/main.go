package main

import "fmt"

func main() {
/**
 * if else 语句
 */
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
//if语句
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
//先执行得到num,再将num的作为判断
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
