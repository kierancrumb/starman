package main

import (
	"fmt"
	"starman/gemini"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("invalid arguments")
		return
	}

	var prefix = strings.Split(args[1], "://")[0]

	switch prefix {
	case "gemini":
		res, meta := gemini.Request(args[1])
		if meta == "text/gemini" {
			var out Document
			out.links, out.body = gemini.Parse(res)
			fmt.Println(out.body)
		} else {
			fmt.Println(res)
		}
	
	case "gopher":
		fmt.Println("Gopher protocol isn't supported yet.")

	case "http", "https":
		fmt.Println("go use google or smnthing idk")
	}
}

type Document struct {
	links []string
	body string
}