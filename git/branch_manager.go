package git

import (
	"fmt"
	"regexp"

	"github.com/drewwyatt/git-tidy/reporting"
)

type branchDeletionError struct {
	Branch string
	Msg    string
}

var reporter = reporting.NewReporter()

// BranchManager higher level abstraction of Git
type BranchManager struct {
	git Git

	GoneBranches         []string
	DeletedBranches      []string
	BranchDeletionErrors []branchDeletionError
}

// NewBranchManager Constructor for branch manager
func NewBranchManager(directory string, exec Executor) BranchManager {
	git := Git{directory: directory, exec: exec}
	return BranchManager{git: git}
}

// Init run setup git commands and set gone branches
func (bm *BranchManager) Init() []error {
	var possibleErrors []error
	possibleErrors = append(possibleErrors, bm.fetch())
	possibleErrors = append(possibleErrors, bm.prune())
	possibleErrors = append(possibleErrors, bm.listRemoteBranches())

	var errors []error
	for idx, error := range possibleErrors {
		if error != nil {
			errors = append(errors, error)
			reporter.FailProcess(idx, error.Error())
		} else {
			reporter.PassProcess(idx)
		}
	}

	if len(errors) == 0 {
		bm.GoneBranches = bm.findGoneBranches()
		reporter.Complete("Initialization complete")
	} else {
		reporter.Fail("Failed to initialize")
	}

	return errors
}

func (bm *BranchManager) fetch() error {
	reporter.Report("git fetch")
	return bm.git.Fetch()
}

func (bm *BranchManager) prune() error {
	reporter.Report("git remote prune origin")
	return bm.git.Prune()
}

func (bm *BranchManager) listRemoteBranches() error {
	reporter.Report("git branch -vv")
	return bm.git.ListRemoteBranches()
}

func (bm *BranchManager) Delete(branch string, force bool) {
	reporter.Report(fmt.Sprintf("git branch %s %s", getDeleteArg(force), branch))
	error := bm.git.Delete(branch, force)
	if error != nil {
		bm.BranchDeletionErrors = append(bm.BranchDeletionErrors, branchDeletionError{Branch: branch, Msg: bm.git.ErrorMsg})
	} else {
		bm.DeletedBranches = append(bm.DeletedBranches, branch)
	}
}

var goneBranch = regexp.MustCompile(`(?m)^(?:\*| ) ([^\s]+)\s+[a-z0-9]+ \[[^:\n]+: gone\].*$`)

func (bm *BranchManager) findGoneBranches() []string {
	var goneBranches []string
	submatches := goneBranch.FindAllStringSubmatch(bm.git.Output, -1)
	for _, matches := range submatches {
		if len(matches) == 2 && matches[1] != "" {
			goneBranches = append(goneBranches, matches[1])
		}
	}

	return goneBranches
}
