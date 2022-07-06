package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("no args")
	}

	logger := log.New(os.Stdout, "Yuta says :", log.Llongfile)

	args := os.Args[1:]
	for _, arg := range args {
		res, err := http.Get(arg)
		if err != nil {
			logger.Println(err)
			continue
		}
		if res.StatusCode != http.StatusOK {
			logger.Println(res.StatusCode)
			continue
		}
		scanner := bufio.NewScanner(res.Body)
		for scanner.Scan() {
			logger.Println(scanner.Text())
		}
	}
}
