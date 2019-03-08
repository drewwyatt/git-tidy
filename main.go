package main

import (
	"fmt"

	flag "github.com/ogier/pflag"
)

// flags
var (
	force bool
)

func main() {
	flag.Parse()
	if force {
		fmt.Println("you forced it")
	} else {
		fmt.Println("not forced")
	}
}

func init() {
	flag.BoolVarP(&force, "force", "f", false, "Force delete")
}
