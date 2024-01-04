package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(unloadCmd)
	rootCmd.AddCommand(exportCmd)
}

var rootCmd = &cobra.Command{
	Use:   "jenkins-job-loader",
	Short: "jenkins-job-loader is used to load jenkins jobs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jenkins-job-loader is used to load jenkins jobs")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
