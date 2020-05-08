package main

import (
	"fmt"
)

func main() {
	if err := render(640, 480, 512, 4); err != nil {
		fmt.Println(err)
	}
}
