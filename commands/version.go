package commands

import (
	"fmt"
)

// Version print the current version
func Version(version string) {
	fmt.Println(version)
	return
}
