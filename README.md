# jenkins-job-loader

A tool used to export, load and unload jenkins jobs.

## Language 

[ÖÐÎÄÎÄµµ](README_CN.md)

## Basic Knowledge

You can get the job configuration details by adding `/config.xml` after the job link, eg:

The job link is <http://jenkins.xxx.com/job/your-job>, 

and its config link is <http://jenkins.xxx.com/job/your-job/config.xml>

## Build

Use the following command to build the tool:

``` bash
go mod tidy

GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build --ldflags="-s" -v
```

## Usage

- Export jobs to folder

``` bash
# Template
./jenkins-job-loader export -d "{domain}" -u "{user}" -p "{password}" -f "{folder_export_to}"

# Example
./jenkins-job-loader export -d "http://jenkins.xxx.com" -u "admin" -p "123456" -f "./jobs/example"
```

- Load jobs from folder to jenkins

``` bash
# Template
./jenkins-job-loader load -d "{domain}" -u "{user}" -p "{password}" -f "{job_config_folder}" -c "{your_credentials_id}"

# Example 1: Load job with the raw credentials
./jenkins-job-loader load -d "http://jenkins.xxx.com" -u "admin" -p "123456" -f "./jobs/example"

# Example 2: Load job with the your credentials
./jenkins-job-loader load -d "http://jenkins.xxx.com" -u "admin" -p "123456" -f "./jobs/example" -c "1b7574ed-56e7-4af9-b851-052df8e53b87"
```

- Unload jobs in folder from jenkins

``` bash
# Template
./jenkins-job-loader unload -d "{domain}" -u "{user}" -p "{password}" -f "{job_config_folder}"

# Example
./jenkins-job-loader unload -d "http://jenkins.xxx.com" -u "admin" -p "123456" -f "./jobs/example"
```

## Init Jenkins Jobs

This repo consists a [Jenkinsfile](Jenkinsfile), you can use it to create your first jenkins multibranch pipeline job, and use this job to load jobs from folder.
