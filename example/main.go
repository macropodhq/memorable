package main

import (
	"fmt"

	"github.com/BugHerd/memorable/heroku"
)

func main() {
	g := heroku.New(0)

	fmt.Printf("%s\n", g.Get())
}
