package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// Version information - set via ldflags at build time
var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func getVersion() string {
	// If version was set via ldflags, use it
	if Version != "dev" {
		return Version
	}
	// Otherwise try to get it from build info (works with go install)
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return Version
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("forge %s\n", getVersion())
		if Commit != "none" {
			fmt.Printf("  commit: %s\n", Commit)
		}
		if BuildDate != "unknown" {
			fmt.Printf("  built:  %s\n", BuildDate)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
