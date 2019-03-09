package main

import (
	"fmt"
	"os"
	"regexp"

	flag "github.com/ogier/pflag"
)

// flags
var (
	force bool
)

var goneBranch = regexp.MustCompile(`(?m)^(?:\*| ) ([^\s]+)\s+[a-z0-9]{9} \[[^:]+: gone\].*$`)

func checkForError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func main() {
	flag.Parse()
	git := Git{}

	git.Fetch().Prune().ListRemoteBranches()
	submatches := goneBranch.FindAllStringSubmatch(git.output, -1)
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
