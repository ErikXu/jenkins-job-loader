package cmd

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/bndr/gojenkins"
)

func init() {
	unloadCmd.Flags().StringVarP(&JenkinsDomain, "domain", "d", "http://jenkins:8080", "Domain, eg: http://jenkins:8080")
	unloadCmd.Flags().StringVarP(&JenkinsUser, "user", "u", "", "Username")
	unloadCmd.Flags().StringVarP(&JenkinsPassword, "password", "p", "", "Password")
	unloadCmd.Flags().StringVarP(&JobFolder, "folder", "f", "", "Folder")

	unloadCmd.MarkFlagRequired("domain")
	unloadCmd.MarkFlagRequired("user")
	unloadCmd.MarkFlagRequired("password")
	unloadCmd.MarkFlagRequired("folder")
}

var unloadCmd = &cobra.Command{
	Use:   "unload",
	Short: "Unload jenkins job",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		jenkins := gojenkins.CreateJenkins(nil, JenkinsDomain, JenkinsUser, JenkinsPassword)
		_, err := jenkins.Init(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}

		files, err := ioutil.ReadDir(JobFolder)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, file := range files {
			name := fileNameWithoutExtension(file.Name())
			_, err = jenkins.DeleteJob(ctx, name)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Job [%s] is unloaded\n", name)
		}
	},
}
