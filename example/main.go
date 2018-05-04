package main

import (
	"fmt"
	"os"
	"secureimage"
)

func main() {
	trusted, err := secureimage.Check(os.Args[1])

	if err != nil {
		panic(err)
	}

	if trusted {
		fmt.Println("yes.")
	} else {
		fmt.Println("bad image file")
	}
}
