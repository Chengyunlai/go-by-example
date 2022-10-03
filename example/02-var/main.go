package main

import (
	"fmt"
	"math"
)

func main() {
	// 语法：var 变量名 类型(可以省略) = 值
	var a = "initial" // 自动推导

	var b, c int = 1, 2 //声明类型

	var d = true // 自动推导

	var e float64 // 声明未赋值

	// 语法：变量名 := 值
	f := float32(e) // 简短定义

	g := a + "foo"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple

	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
