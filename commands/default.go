package commands

import (
	"fmt"
	"regexp"

	"gopkg.in/AlecAivazis/survey.v1"

	gUtils "github.com/drewwyatt/git-tidy/git"
)

var goneBranch = regexp.MustCompile(`(?m)^(?:\*| ) ([^\s]+)\s+[a-z0-9]+ \[[^:\n]+: gone\].*$`)

func Default(directory string, interactive bool, force bool) {
	if directory == "" {
		directory = "."
	}

	git := gUtils.NewGit(directory, gUtils.NewExecutorWithExec())

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
		Message: "Branches to delete:",
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
