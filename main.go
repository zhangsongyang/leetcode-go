package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(cmp.Equal("he", "he"))
	a := 100
	fmt.Println(&a)
	var pr *int = &a
	fmt.Println(pr)
	fmt.Println(*pr)
}
