package main

import (
	"cc/yangcy/stringutil"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"math/rand"
)

const pi = 3.141592653

/*
课程1
@see https://tour.go-zh.org/basics/1
*/
func main() {
	//glog使用
	flag.Parse()
	glog.Info("This is the first log from ")
	glog.Flush()
	//自己的module使用
	fmt.Printf("1234567 ==> %v\n", stringutil.Reverse("1234567"))
	//使用math/rand
	fmt.Printf("rand.Int31n(): %v\n", rand.Int31n(100))
	//使用本类中定义的函数
	fmt.Printf("1+2=%v", add(1, 2))
	//返回两个结果
	a, b := swap(1, 2)
	fmt.Printf("1 , 2 swap %v , %v\n", a, b)
	fmt.Println(swap(1, 2))
	//使用常量
	fmt.Printf("Pi constant is %v\n", pi)
	//cannot refer to unexported name stringutil.blankString
	//fmt.Printf("stringutil.blankString constant is \"%v\"", stringutil.anotherBlankString)
	//大些变量默认导出
	fmt.Printf("stringutil.AnExportBlankString constant is \"%v\"\n", stringutil.AnExportBlankString)
}

func add(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}
