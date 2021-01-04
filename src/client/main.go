package main

import (
	"cc/yangcy/stringutil"
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	fmt.Printf("1234567 ==> %v\n", stringutil.Reverse("1234567"))
	glog.Info("This is the first log from ")
	glog.Flush()
}
