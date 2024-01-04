package cmd

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/bndr/gojenkins"
)

func init() {
	exportCmd.Flags().StringVarP(&JenkinsDomain, "domain", "d", "http://jenkins:8080", "Domain, eg: http://jenkins:8080")
	exportCmd.Flags().StringVarP(&JenkinsUser, "user", "u", "", "Username")
	exportCmd.Flags().StringVarP(&JenkinsPassword, "password", "p", "", "Password")
	exportCmd.Flags().StringVarP(&JobFolder, "folder", "f", "", "Folder")

	exportCmd.MarkFlagRequired("domain")
	exportCmd.MarkFlagRequired("user")
	exportCmd.MarkFlagRequired("password")
	exportCmd.MarkFlagRequired("folder")
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export jenkins jobs to files",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		jenkins := gojenkins.CreateJenkins(nil, JenkinsDomain, JenkinsUser, JenkinsPassword)
		_, err := jenkins.Init(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}

		err = os.RemoveAll(JobFolder)
		if err != nil {
			fmt.Println(err)
			return
		}

		os.MkdirAll(JobFolder, os.ModePerm)

		jobs, err := jenkins.GetAllJobs(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, job := range jobs {
			config, err := job.GetConfig(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}

			name := job.Raw.Name
			filepath := path.Join(JobFolder, fmt.Sprintf("%s.xml", name))
			f, err := os.Create(filepath)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.WriteString(config)
			fmt.Printf("Job [%s] is exported\n", name)
		}
	},
}
