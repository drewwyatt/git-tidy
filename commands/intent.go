package commands

const (
	// DefaultIntent Default Intent Type
	DefaultIntent = iota
	// VersionIntent Intent to print version
	VersionIntent = iota
)

// Intent Used for deriving user intent based on predefined business rules
type Intent struct {
	PrintVersion bool
}

// NewIntent Intent Constructor
func NewIntent(printVersion bool) Intent {
	return Intent{
		PrintVersion: printVersion,
	}
}

// Is Get derived intent type
func (i *Intent) Is() int {
	if i.PrintVersion {
		return VersionIntent
	}

	return DefaultIntent
}
