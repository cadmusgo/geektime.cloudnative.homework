// read env test
// ref: https://gobyexample.com/environment-variables
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("start...")
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
}
