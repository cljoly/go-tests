package main

import (
	"fmt"

	finder "github.com/b4b4r07/go-finder"
)

// Testing https://github.com/b4b4r07/go-finder

func main() {
	cli, err := finder.New()
	if err != nil {
		panic(err)
	}
	args, _ := cli.Select([]string{"aiue", "tsrn"})
	fmt.Println("args:", args)
}
