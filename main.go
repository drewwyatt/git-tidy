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

	git.Delete("hey", force)
	git.Delete("ho", force)

	if len(git.DeletedBranches) > 0 {
		fmt.Println("Deleted branches:")
		for _, branch := range git.DeletedBranches {
			fmt.Println(branch)
		}
	}

	if len(git.BranchDeletionErrors) > 0 {
		fmt.Println("Errors:")
		for _, err := range git.BranchDeletionErrors {
			fmt.Printf("[%s]: %s", err.Branch, err.Msg)
		}
	}

	fmt.Println("Done.")
}

func init() {
	flag.BoolVarP(&force, "force", "f", false, "Force delete")
}
