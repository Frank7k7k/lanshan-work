package main

import (
	"math"
)

func add(x int, y int) int {
	return x + y
}

//1.返回两个数相加的值 2.圆的面积 3.判断素数 4.随机数，二分法

func main() {
	result := add(3, 5)
	println(result)
	var r = 10.00
	var s = math.Pi * r * r
	println(s)

}
