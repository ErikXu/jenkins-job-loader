package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/bndr/gojenkins"
)

func init() {
	loadCmd.Flags().StringVarP(&JenkinsDomain, "domain", "d", "http://jenkins:8080", "Domain, eg: http://jenkins:8080")
	loadCmd.Flags().StringVarP(&JenkinsUser, "user", "u", "", "Username")
	loadCmd.Flags().StringVarP(&JenkinsPassword, "password", "p", "", "Password")
	loadCmd.Flags().StringVarP(&JobFolder, "folder", "f", "", "Folder")
	loadCmd.Flags().StringVarP(&CredentialsId, "credentials", "c", "", "Credentials Id")

	loadCmd.MarkFlagRequired("domain")
	loadCmd.MarkFlagRequired("user")
	loadCmd.MarkFlagRequired("password")
	loadCmd.MarkFlagRequired("folder")
}

var CredentialsId string

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load jenkins job",
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
			content, err := os.ReadFile(filepath.Join(JobFolder, file.Name()))
			if err != nil {
				fmt.Println(err)
				return
			}

			config := string(content)

			if CredentialsId != "" {
				regexCredentials := regexp.MustCompile(`<credentialsId>.*</credentialsId>`)
				config = regexCredentials.ReplaceAllString(config, fmt.Sprintf("<credentialsId>%s</credentialsId>", CredentialsId))
			}

			name := fileNameWithoutExtension(file.Name())
			_, err = jenkins.CreateJob(ctx, config, name)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Job [%s] is loaded\n", name)
		}
	},
}
