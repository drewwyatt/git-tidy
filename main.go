package main

import (
	"fmt"
	"os"
	"regexp"

	gUtils "github.com/drewwyatt/gitclean/git"
	flag "github.com/ogier/pflag"
)

// flags
var (
	force bool
)

var goneBranch = regexp.MustCompile(`(?m)^(?:\*| ) ([^\s]+)\s+[a-z0-9]+ \[[^:\n]+: gone\].*$`)

func checkForError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	git := gUtils.Git{}

	git.Fetch().Prune().ListRemoteBranches()
	submatches := goneBranch.FindAllStringSubmatch(git.Output, -1)
	for _, matches := range submatches {
		if len(matches) == 2 && matches[1] != "" {
			fmt.Printf("delete this branch: %s\n", matches[1])
		}
	}

	fmt.Println("Done.")
}

func init() {
	flag.BoolVarP(&force, "force", "f", false, "Force delete")
}
