package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "forge",
	Short: "Forge - A code generation tool for full-stack applications",
	Long: `Forge is a comprehensive code generation system that transforms YAML specifications 
into production-ready full-stack applications with MongoDB, Go, TypeScript, and Kotlin support.

It automatically generates models, API endpoints, validation logic, permission systems, 
HTTP handlers, database operations, and React integration.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(initCmd)
}
