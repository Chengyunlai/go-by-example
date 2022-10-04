package main

import "fmt"

func main() {

	i := 1
	// 条件判断省略，死循环
	for {
		fmt.Println("loop")
		break // 跳出死循环；break 用来终止执行for循环语句，终止整个循环。
	}
	// 标准语法
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue // continue 终止当前循环的迭代，重新进入下一条件，进入循环。
		}
		fmt.Println(n)
	}

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("---")
	for i := 1; i < 10; i++ {
    if i == 5 {
        break // break
    }
    fmt.Println(i)
}

}
