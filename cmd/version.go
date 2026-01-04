package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version information - set via ldflags at build time
var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("forge %s\n", Version)
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
