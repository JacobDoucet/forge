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
	Long: `Initialize a new forge project by creating a models directory 
with an example model YAML file to get you started.`,
	RunE: runInit,
}

var modelsDir string

func init() {
	initCmd.Flags().StringVarP(&modelsDir, "dir", "d", "./models", "directory to create models in")
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
		return fmt.Errorf("file %s already exists, refusing to overwrite", exampleModelPath)
	}

	if err := os.WriteFile(exampleModelPath, []byte(exampleModel), 0644); err != nil {
		return fmt.Errorf("failed to write example model: %w", err)
	}

	fmt.Printf("✓ Created models directory: %s\n", modelsDir)
	fmt.Printf("✓ Created example model: %s\n", exampleModelPath)
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Edit the example model in", exampleModelPath)
	fmt.Println("  2. Run 'forge build' to generate code")
	fmt.Println()
	fmt.Println("Example build command:")
	fmt.Printf("  forge build --specDir %s --goOutDir ./generated --goPkgRoot myapp/generated\n", modelsDir)

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
        type: "[]string"
      - name: assigneeId
        type: string
    
    # MongoDB collection configuration
    collection:
      - type: mongo
        name: tasks

    # Define indexes for efficient queries
    indexes:
      - fields:
          - status
      - fields:
          - assigneeId
          - status

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
      create:
        - admin
        - user
      read:
        - admin
        - user
      update:
        - admin
        - user
      delete:
        - admin

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
  roles:
    - admin
    - user
    - guest
`
