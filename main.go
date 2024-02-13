package main

import (
	"fmt"
)

func main() {
	c := make(map[string]string)
	c["you"] = "alica, bob, charlie"
	fmt.Println(c["you"])
}
