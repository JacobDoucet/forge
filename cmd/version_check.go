package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CheckAndUpdateVersion compares the current version with the required version
// and prompts the user to update if they don't match.
// Returns true if the program should continue, false if it should exit.
func CheckAndUpdateVersion(requiredVersion string) (bool, error) {
	if requiredVersion == "" {
		return true, nil
	}

	currentVersion := getVersion()

	// Normalize versions for comparison (remove 'v' prefix if present)
	normalizedRequired := normalizeVersion(requiredVersion)
	normalizedCurrent := normalizeVersion(currentVersion)

	if versionsCompatible(normalizedCurrent, normalizedRequired) {
		return true, nil
	}

	// Special case: if current is "dev", we're running a development build
	if currentVersion == "dev" {
		fmt.Printf("⚠️  Running development build. Config specifies version %s\n", requiredVersion)
		fmt.Print("Continue anyway? [y/N]: ")
		if !promptYesNo() {
			return false, nil
		}
		return true, nil
	}

	fmt.Printf("\n⚠️  Version mismatch detected!\n")
	fmt.Printf("   Current version:  %s\n", currentVersion)
	fmt.Printf("   Required version: %s\n", requiredVersion)
	fmt.Println()

	// Ask user if they want to update
	fmt.Printf("Would you like to update forge to version %s? [y/N]: ", requiredVersion)
	if !promptYesNo() {
		fmt.Println("Aborting. Please update manually or change forgeVersion in your config.")
		return false, nil
	}

	// Perform the update
	if err := updateForge(requiredVersion); err != nil {
		return false, fmt.Errorf("failed to update forge: %w", err)
	}

	fmt.Println()
	fmt.Printf("✅ Successfully updated to version %s\n", requiredVersion)
	fmt.Println("Please re-run your command.")
	return false, nil
}

// normalizeVersion removes the 'v' prefix from a version string if present
func normalizeVersion(version string) string {
	return strings.TrimPrefix(version, "v")
}

// versionsCompatible checks if the current version is compatible with the required version.
// It allows commits after a tag (e.g., v1.0.3-1-gff58aa8 is compatible with v1.0.3)
// and versions with -dirty suffix during development.
func versionsCompatible(current, required string) bool {
	// Exact match
	if current == required {
		return true
	}

	// Allow if current version starts with required version followed by - (for git describe suffixes)
	// e.g., "1.0.3-1-gff58aa8" matches "1.0.3", "1.0.3-dirty" matches "1.0.3"
	if strings.HasPrefix(current, required+"-") {
		return true
	}

	return false
}

// promptYesNo reads a yes/no response from stdin
// Returns true for 'y' or 'Y', false otherwise
func promptYesNo() bool {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}

// updateForge runs go install to update to the specified version
func updateForge(version string) error {
	// Ensure version has 'v' prefix for go install
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}

	installPath := fmt.Sprintf("github.com/JacobDoucet/forge@%s", version)

	fmt.Printf("Running: go install %s\n", installPath)

	cmd := exec.Command("go", "install", installPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go install failed: %w", err)
	}

	return nil
}
