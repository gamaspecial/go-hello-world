package main

import (
	"fmt"

	"github.com/gamaspecial/go-hello-world/subpkg"
)

func main() {
	fmt.Println(subpkg.Hello())
	fmt.Println(subpkg.Golang())
	fmt.Println(Goodbye())
}
