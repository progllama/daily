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
		url := arg
		if arg[:8] != "https://" && arg[:7] != "http://" {
			url = "https://" + url
		}

		res, err := http.Get(url)
		if err != nil {
			logger.Println(err)
			continue
		}

		if res.StatusCode != http.StatusOK {
			logger.Println(res.StatusCode)
			continue
		}

		scanner := bufio.NewScanner(res.Body)
		var body string
		for scanner.Scan() {
			body += scanner.Text()
		}

		currentNode := Node{}
		tag := false
		for i, c := range body {
			if string(c) == "<" {
				tag = true
			}

			if string(c) == ">" {
				tag = false
			}

			if tag {

			} else {

			}
		}
	}
}

type Node struct {
	ParentNode *Node
	ChildNodes []Node
}
