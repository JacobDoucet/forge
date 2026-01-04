package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new forge project with example models",
	Long: `Initialize a new forge project by creating:
- A models directory with an example model YAML file
- A .forge.yml configuration file`,
	RunE: runInit,
}

var (
	modelsDir  string
	goPkg      string
	kotlinPkg  string
	skipConfig bool
)

func init() {
	initCmd.Flags().StringVarP(&modelsDir, "dir", "d", "./models", "directory to create models in")
	initCmd.Flags().StringVar(&goPkg, "goPkg", "myapp/generated", "Go package root for generated code")
	initCmd.Flags().StringVar(&kotlinPkg, "kotlinPkg", "com.myapp.generated", "Kotlin package root for generated code")
	initCmd.Flags().BoolVar(&skipConfig, "skip-config", false, "skip creating .forge.yml config file")
}

func runInit(cmd *cobra.Command, args []string) error {
	// Create the models directory
	if err := os.MkdirAll(modelsDir, 0755); err != nil {
		return fmt.Errorf("failed to create models directory: %w", err)
	}

	// Create an example model file
	exampleModelPath := filepath.Join(modelsDir, "example.yaml")

	// Check if file already exists
	if _, err := os.Stat(exampleModelPath); err == nil {
		fmt.Printf("⚠ Model file %s already exists, skipping\n", exampleModelPath)
	} else {
		if err := os.WriteFile(exampleModelPath, []byte(exampleModel), 0644); err != nil {
			return fmt.Errorf("failed to write example model: %w", err)
		}
		fmt.Printf("✓ Created example model: %s\n", exampleModelPath)
	}

	// Create .forge.yml config file
	if !skipConfig {
		configPath := ConfigFileName
		if _, err := os.Stat(configPath); err == nil {
			fmt.Printf("⚠ Config file %s already exists, skipping\n", configPath)
		} else {
			config := &Config{
				SpecDir: modelsDir,
				Go: &GoConfig{
					OutDir:  "./generated/go",
					PkgRoot: goPkg,
				},
				TypeScript: &TypeScriptConfig{
					OutDir: "./generated/typescript",
				},
				Kotlin: &KotlinConfig{
					OutDir:  "./generated/kotlin",
					PkgRoot: kotlinPkg,
				},
			}

			if err := WriteConfig(configPath, config); err != nil {
				return fmt.Errorf("failed to write config file: %w", err)
			}
			fmt.Printf("✓ Created config file: %s\n", configPath)
		}
	}

	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Edit your models in", modelsDir)
	fmt.Println("  2. Customize .forge.yml if needed")
	fmt.Println("  3. Run 'forge build' to generate code")

	return nil
}

const exampleModel = `# Example Forge Model Specification
# This file demonstrates the key features of Forge model definitions.
# See the README for complete documentation.

# Define enums for status values
enums:
  - name: TaskStatus
    type: string
    values:
      - pending
      - in_progress
      - completed
      - cancelled

  - name: TaskPriority
    type: string
    values:
      - low
      - medium
      - high
      - urgent

# Define your data objects
objects:
  # A simple Task object with MongoDB collection and HTTP endpoints
  - name: Task
    fields:
      - name: title
        type: string
        required: true
      - name: description
        type: string
      - name: status
        type: TaskStatus
        required: true
      - name: priority
        type: TaskPriority
      - name: dueDate
        type: timestamp
      - name: tags
        type: List<string>
      - name: assigneeId
        type: string
    
    # MongoDB collection configuration
    collection:
      - type: mongo
        name: tasks

    # Define indexes for efficient queries
    indexes:
      - name: status_idx
        fields:
          - name: status
      - name: assignee_status_idx
        fields:
          - name: assigneeId
          - name: status

    # HTTP endpoint configuration
    http:
      endpoint: tasks
      methods:
        - GET
        - POST
        - PUT
        - DELETE

    # Permission configuration (RBAC)
    permissions:
      read:
        - rbac: admin
        - rbac: user
      write:
        - rbac: admin
        - rbac: user

  # A nested object (no collection - used as an embedded type)
  - name: TaskComment
    fields:
      - name: text
        type: string
        required: true
      - name: authorId
        type: string
        required: true
      - name: createdAt
        type: timestamp

# Define custom errors
errors:
  - name: TaskNotFound
    code: TASK_NOT_FOUND
    message: "The requested task was not found"
  - name: InvalidTaskStatus
    code: INVALID_TASK_STATUS  
    message: "The provided task status is invalid"

# Define event types for real-time updates
events:
  - TaskCreated
  - TaskUpdated
  - TaskDeleted
  - TaskStatusChanged

# Global permissions definition
permissions:
  rbac:
    - name: admin
    - name: user
    - name: guest
`
