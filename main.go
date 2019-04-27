package main

import (
	"github.com/drewwyatt/gitclean/commands"
	flag "github.com/ogier/pflag"
)

// build info injected by gitclean
var version = "development snapshot"

ksjdhfkshfk

// flags
var (
	force        bool
	interactive  bool
	printVersion bool
)

func main() {
	flag.Parse()
	directory := flag.Arg(0) // first trailing argument after flags (if any)

	info := commands.NewIntentInfo(printVersion)

	switch i := info.Intent(); i {
	case commands.VersionIntent:
		commands.Version(version)
		return
	default:
		commands.Default(directory, interactive, force)
		return
	}
}

func init() {
	flag.BoolVarP(&force, "force", "f", false, "Force delete")
	flag.BoolVarP(&interactive, "interactive", "i", false, "Select branches you would like to delete from a list")
	flag.BoolVarP(&printVersion, "version", "v", false, "Print the version of gitclean")
}
