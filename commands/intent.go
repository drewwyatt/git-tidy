package commands

type Intent int

const (
	// DefaultIntent Default Intent Type
	DefaultIntent Intent = iota
	// VersionIntent Intent to print version
	VersionIntent Intent = iota
)

// IntentInfo Used for deriving user intent based on predefined business rules
type IntentInfo struct {
	PrintVersion bool
}

// NewIntentInfo IntentInfo Constructor
func NewIntentInfo(printVersion bool) IntentInfo {
	return IntentInfo{
		PrintVersion: printVersion,
	}
}

// Intent Get derived Intent
func (i *IntentInfo) Intent() Intent {
	if i.PrintVersion {
		return VersionIntent
	}

	return DefaultIntent
}
