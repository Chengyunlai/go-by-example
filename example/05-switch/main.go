package main

import (
	"fmt"
	"time"
)

func main() {

	a := 2
	// if 之后只能是bool类型， 而switch 可以作用其他类型
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default: // default 不是必须的，根据实际情况来写。
		fmt.Println("other")
		// 注意，没有break;在switch 语句中，默认每个case后自带一个break，表示到此结束 不向下执行了，跳出整个switch。
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
//语法二 省略变量 相当于作用在了bool 类型上
	switch{
    case true: fmt.Println("true") // 默认是true
    case false: fmt.Println("false")
	}

	num := 2

	//语法三 case 后可以跟随多个数值， 满足其中一个就执行
switch num{
case 1,2,3:
		fmt.Println("num符合其中某一个 执行代码")
		// fallthrough 表示强制执行后面的没有执行的case代码。
		fallthrough
case 4,5,6:
		fmt.Println("执行此代码")
}

//语法四 可以添加初始化变量 作用于switch内部
switch name:="huangrong"; name{
case "guojing":
		fmt.Println("shideguojing")
case "huangrong":
		fmt.Println("shidehuangrong")
} 
}
