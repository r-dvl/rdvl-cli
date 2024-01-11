/*
Copyright © 2024 Raúl Del Valle Lima <rauldel.valle.lima@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

// orderCmd represents the order command
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Order folder content.",
	Long:  `Order folder content.`,
	Run: func(cmd *cobra.Command, args []string) {
		i := 0
		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() && path != "." {
				oldName := info.Name()
				newName := strconv.Itoa(i+1) + ". " + oldName
				err := os.Rename(oldName, newName)
				if err != nil {
					return err
				}
				i++
			}

			return nil
		})

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Folders have been successfully renamed.")
	},
}

func init() {
	rootCmd.AddCommand(orderCmd)
}
