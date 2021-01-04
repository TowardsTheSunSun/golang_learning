package main

import (
	"fmt"
	"math"
)

/*
课程2 https://tour.go-zh.org/flowcontrol/1
*/
func main() {
	//for循环
	fmt.Printf("sum from %v to %v step %v is %v\n", 1, 10, 1, sum(1, 10, 1))
	fmt.Printf("sum2 from %v to %v step %v is %v\n", 1, 10, 1, sum2(1, 10, 1))
	fmt.Printf("sum3 from %v to %v step %v is %v\n", 1, 10, 1, sum3(1, 10, 1))
	//fmt.Printf("start noEndingLoop()")
	//noEndingLoop()
	fmt.Printf("max(3,4) by ifMatchBetween %v\n", max(3, 4))
	fmt.Printf("if pow(2, 3) exceed 5 then return 5: %v \n", ifPowerExceedReturn(2, 3, 5))
	fmt.Printf("if pow(2, 3) exceed 9 then return 9: %v \n", ifPowerExceedReturn(2, 3, 9))

}

//正常的循环
func sum(from, to, step int) int {
	sum := 0
	for i := from; i <= to; i += step {
		sum += i
	}
	return sum
}

//初始化语句和后置语句是可选的
func sum2(from, to, step int) int {
	sum := 0
	for ; from <= to; from += step {
		sum += from
	}
	return sum
}

//for就是go语言中的while
func sum3(from, to, step int) int {
	sum := 0
	for from <= to {
		sum += from
		from += step
	}
	return sum
}

//死循环
func noEndingLoop() {
	for true {
		fmt.Print(".")
	}
}

func max(x, y int) int {
	return ifConditionReturn(x > y, x, y)
}

//如果条件满足返回a，不满足返回b
func ifConditionReturn(condition bool, matchReturn, notMathReturn int) int {
	if condition {
		return matchReturn
	} else {
		return notMathReturn
	}
}

//如果幂值超过一个数字，则返回这个数字
//条件中可以初始化变量， 这个变量只能用在if else中
func ifPowerExceedReturn(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v > lim {
		return lim
	} else {
		return v
	}
}
