package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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
	//牛顿法开根号
	sqrt(4)
	//switch case
	printSystemName(runtime.GOOS)
	//switch case float
	fmt.Printf("switchCaseFloat %v: %v\n", 1.0, switchCaseFloat(1.0))
	fmt.Printf("switchCaseFloat %v: %v\n", 2.0, switchCaseFloat(2.0))
	fmt.Printf("switchCaseFloat %v: %v\n", 11.0, switchCaseFloat(11.0))
	//switch no condition
	switchCaseNoCondition(time.Now())
	//defer
	deferOuter()
	//defer loop
	deferLoop()
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

//使用题目中的公式 z -= (z * z - x)/(2*z)，使得z逐渐接近于跟好根号x
func sqrt(x float64) {
	i := 0
	z := x / 2
	for z < x && i < 10 {
		fmt.Printf("sqrtForNumberGreatThanOne temp %v - %v\n", i, z)
		i++
		z -= (z*z - x) / (2 * z)
	}
	fmt.Printf("sqrtForNumberGreatThanOne result %v - %v\n", i, z)
}

//switch case string
func printSystemName(os string) {
	switch os {
	case "darwin":
		fmt.Printf("Your system is: %v\n", "OS X")
	case "linux":
		fmt.Printf("Your system is: %v\n", "Linux")
	default:
		fmt.Printf("Your system is: %v\n", os)
	}
}

//switch case float64
func switchCaseFloat(x float64) int {
	switch x {
	case 1.0:
		return 1
	case 2.0:
		return 2
	case 3.0:
		return 3
	default:
		return int(math.Floor(x))
	}
}

//switch no condition
func switchCaseNoCondition(t time.Time) {
	switch {
	case t.Hour() < 12:
		fmt.Printf("Good morning.\n")
	case t.Hour() < 17:
		fmt.Printf("Good afternoon.\n")
	default:
		fmt.Printf("Good evening.\n")
	}
}
func deferOuter() {
	defer deferInner()
	fmt.Printf("deferOuter run\n")
}
func deferInner() {
	fmt.Printf("deferInner run\n")
}
func deferLoop() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("defer loop: %v\n", i)
	}
	fmt.Printf("done.\n")
}
