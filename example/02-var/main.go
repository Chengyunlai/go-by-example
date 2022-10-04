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

	// 常量的声明
	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))

	print()
	change()
}

func print(){
	//var方式声明多变量
	var a,b,c int
	a=1 
	b=2 
	c=3
 //也可以写在一行
 	// var a1,a2,a3 int =10,20,30
 //也可以省略类型 根据数据进行类型推导
	// var a1,a2,a3,a4 =10,20,"ago",30
 //如果是多种类型 也可以使用集合
 	var(
		 a1 =""
		 a2 =10
 )
 fmt.Println(a,b,c)
//  fmt.Println(a1,a2,a3)
//  fmt.Println(a1,a2,a3,a4)
 fmt.Println(a1,a2)
}

/**
 * 交换元素
 */
func change(){
	var a,b = 10,20
	a,b = b,a
	fmt.Println(a,b)
}
