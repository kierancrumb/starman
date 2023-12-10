package main

import (
	"fmt"
	"starman/gemini"
	"os"
)

func main() {
	args := os.Args
	res, meta := gemini.Request(args[1])
	if meta != "" {
		fmt.Println(res)	
	}
}