package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

// Variable to store the key to be passed with the --hide option
var hideKey string

// secretsCmd represents the `secrets` command
var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Command to manage secrets in YAML files",
	Run: func(cmd *cobra.Command, args []string) {
		if hideKey != "" {
			err := hideSecrets(hideKey)
			if err != nil {
				fmt.Printf("Error hiding secrets: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Secrets were successfully hidden.")
		} else {
			fmt.Println("You must provide the --hide option with the key you want to hide.")
		}
	},
}

func init() {
	// Add the command to the root
	rootCmd.AddCommand(secretsCmd)

	// Define the --hide option with a short flag -k
	secretsCmd.Flags().StringVarP(&hideKey, "hide", "k", "", "Key whose value will be replaced with 'secret' in YAML files")
}

func hideSecrets(key string) error {
	// Get all .yaml and .yml files in the current directory
	files, err := filepath.Glob("*.y*ml")
	if err != nil {
		return fmt.Errorf("could not find YAML files: %v", err)
	}

	// Iterate over each file found
	for _, file := range files {
		err := processFile(file, key)
		if err != nil {
			return err
		}
	}

	return nil
}

func processFile(file, key string) error {
	// Read the file
	input, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v", file, err)
	}

	// Create a regex pattern to match the key and its value, preserving indentation
	// Example: if `key` is "password", the pattern searches for lines like "  password: value"
	pattern := fmt.Sprintf(`(?m)^(\s*%s:\s*)(.+)$`, regexp.QuoteMeta(key))
	re := regexp.MustCompile(pattern)

	// Replace only the value of the key with "secret" while preserving the original indentation
	output := re.ReplaceAllString(string(input), "${1}secret")

	// Write the changes back to the file only if changes occurred
	if string(input) != output {
		err = os.WriteFile(file, []byte(output), 0644)
		if err != nil {
			return fmt.Errorf("error writing file %s: %v", file, err)
		}
		fmt.Printf("Updated: %s\n", file)
	}

	return nil
}
