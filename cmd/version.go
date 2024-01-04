package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.1.0")
	},
}
