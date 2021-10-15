// glog test
package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	fmt.Println("start...")
	glog.Info("This is info message")
}
