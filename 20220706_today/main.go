package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("no args")
	}

	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}
