package main

import (
	"fmt"
	"os"
	"regexp"

	"gopkg.in/AlecAivazis/survey.v1"

	gUtils "github.com/drewwyatt/gitclean/git"
	flag "github.com/ogier/pflag"
)

// build info
var version = "development snapshot"

// flags
var (
	force        bool
	interactive  bool
	printVersion bool
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

	if printVersion {
		fmt.Println(version)
		return
	}

	directory := flag.Arg(0) // first trailing argument after flags (if any)
	if directory == "" {
		directory = "."
	}

	git := gUtils.NewExecutor(directory)

	goneBranches := []string{}
	branchesToDelete := []string{}

	git.Fetch().Prune().ListRemoteBranches()
	submatches := goneBranch.FindAllStringSubmatch(git.Output, -1)
	for _, matches := range submatches {
		if len(matches) == 2 && matches[1] != "" {
			if interactive {
				goneBranches = append(goneBranches, matches[1])
			} else {
				branchesToDelete = append(branchesToDelete, matches[1])
			}
		}
	}

	prompt := &survey.MultiSelect{
		Message: "Use the spacebar to select the branches you would like to delete:",
		Options: goneBranches,
	}
	survey.AskOne(prompt, &branchesToDelete, nil)

	for _, branch := range branchesToDelete {
		git.Delete(branch, force)
	}

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
	flag.BoolVarP(&interactive, "interactive", "i", false, "Select branches you would like to delete from a list")
	flag.BoolVarP(&printVersion, "version", "v", false, "Print the version of gitclean")
}
