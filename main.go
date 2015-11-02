package main

import (
	"fmt"
	"github.com/newmannh/go-euler/fetching"
)

func main() {
	problemToDisplay, err := fetching.FetchProblem(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The problem:\n\n%s\n", problemToDisplay)
}
