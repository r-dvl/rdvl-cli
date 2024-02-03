package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Order folder content.",
	Long:  `Order folder content.`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		i := 0
		for _, file := range files {
			if file.IsDir() {
				oldName := file.Name()
				newName := strconv.Itoa(i+1) + ". " + oldName
				err := os.Rename(oldName, newName)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				i++
			}
		}

		fmt.Println("Folders have been successfully renamed.")
	},
}

func init() {
	rootCmd.AddCommand(orderCmd)
}
