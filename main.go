package main

import (
	"fmt"

	"github.com/hoge/hello-world/subpkg"
)

func main() {
	fmt.Println(subpkg.Hello())
	fmt.Println(subpkg.Golang())
	fmt.Println(Goodbye())
}
