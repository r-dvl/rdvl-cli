package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// Variable to store the key to be passed with the --hide option
var hideKey string

// secretsCmd represents the `secrets` command
var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Command to manage secrets in YAML files",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the --hide option was provided
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

	// Define the --hide option with a short flag -k to avoid conflicts with -h (help)
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
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v", file, err)
	}

	// Unmarshal YAML content into a map of string to interface{}
	var content map[string]interface{}
	err = yaml.Unmarshal(data, &content)
	if err != nil {
		return fmt.Errorf("error parsing YAML file %s: %v", file, err)
	}

	// Find and replace the key in the map
	changed := replaceValue(content, key)

	// If a change was made, save the file
	if changed {
		// Convert the map to JSON to preserve the order
		jsonData, err := json.Marshal(content)
		if err != nil {
			return fmt.Errorf("error generating JSON for %s: %v", file, err)
		}

		// Convert JSON to YAML
		var newData interface{}
		err = json.Unmarshal(jsonData, &newData)
		if err != nil {
			return fmt.Errorf("error parsing JSON for %s: %v", file, err)
		}

		newYAMLData, err := yaml.Marshal(newData)
		if err != nil {
			return fmt.Errorf("error generating YAML for %s: %v", file, err)
		}

		err = os.WriteFile(file, newYAMLData, 0644)
		if err != nil {
			return fmt.Errorf("error writing file %s: %v", file, err)
		}
		fmt.Printf("Updated: %s\n", file)
	}

	return nil
}

// replaceValue replaces the value of the specified key in the map and its submaps
func replaceValue(content map[string]interface{}, key string) bool {
	changed := false
	for k, v := range content {
		if k == key {
			content[k] = "secret"
			changed = true
		} else if nestedMap, ok := v.(map[string]interface{}); ok {
			if replaceValue(nestedMap, key) {
				changed = true
			}
		}
	}
	return changed
}
