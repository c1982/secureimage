package main

import (
	"fmt"
	"secureimage"
)

func main() {

	trusted, err := secureimage.Check("bad.jpg")

	if err != nil {
		panic(err)
	}

	if trusted {
		fmt.Println("yes.")
	} else {
		fmt.Println("bad image file")
	}
}
