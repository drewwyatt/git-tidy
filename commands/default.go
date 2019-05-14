package commands

import (
	"fmt"

	"github.com/drewwyatt/git-tidy/git"
	"gopkg.in/AlecAivazis/survey.v1"
)

func Default(directory string, interactive bool, force bool) {
	if directory == "" {
		directory = "."
	}

	branchManager := git.NewBranchManager(directory, git.NewExecutorWithExec())
	branchManager.Init()

	goneBranches := []string{}
	branchesToDelete := []string{}

	if interactive {
		goneBranches = branchManager.GoneBranches
	} else {
		branchesToDelete = branchManager.GoneBranches
	}

	prompt := &survey.MultiSelect{
		Message: "Branches to delete:",
		Options: goneBranches,
	}
	survey.AskOne(prompt, &branchesToDelete, nil)

	for _, branch := range branchesToDelete {
		branchManager.Delete(branch, force)
	}

	if len(branchManager.DeletedBranches) > 0 {
		fmt.Println("Deleted branches:")
		for _, branch := range branchManager.DeletedBranches {
			fmt.Println(branch)
		}
	}

	if len(branchManager.BranchDeletionErrors) > 0 {
		fmt.Println("Errors:")
		for _, err := range branchManager.BranchDeletionErrors {
			fmt.Printf("[%s]: %s", err.Branch, err.Msg)
		}
	}

	fmt.Println("Done.")
}
