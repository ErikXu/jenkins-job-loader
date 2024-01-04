pipeline {
    agent any

    parameters {
        string(name: "DOMAIN", defaultValue: "http://jenkins:8080", description: "The domain of jenkins")
        string(name: "USER", defaultValue: "admin", description: "The user of jenkins")
        password(name: "PASSWORD", defaultValue: "", description: "The password of jenkins user")
        string(name: "FOLDER", defaultValue: "./jobs/example", description: "The folder to load jenkins job from")
        string(name: "CREDENTIALS_ID", defaultValue: "", description: "Your own credentials id if you want to replace with")
    }

    stages {
        stage("pre-build") {
            steps {
                script {
                    echo "------ Print deploy info ------"

                    env.GIT_COMMIT_MSG = sh (script: 'git log -1 --pretty=%B ${GIT_COMMIT}', returnStdout: true).trim()

                    echo "Current Commit: ${env.GIT_COMMIT}"
                    echo "Commit Description: ${env.GIT_COMMIT_MSG}"

                    env.DOMAIN =  params.DOMAIN
                    echo "Jenkins Domain: ${env.DOMAIN}"

                    env.USER =  params.USER
                    echo "Jenkins User: ${env.USER}"

                    env.FOLDER  = params.FOLDER
                    echo "Job Folder: ${env.FOLDER}"

                    env.CREDENTIALS_ID  = params.CREDENTIALS_ID
                    echo "Credentials Id: ${env.CREDENTIALS_ID}"
                }
            }
        }

        stage("build") {
            steps {
                script {
                    echo "------ Start building ------"                   
                    sh "cat build.sh | sed 's|\${PWD}|${env.WORKSPACE}|g' > build_in_jenkins.sh"
                    sh "sh build_in_jenkins.sh"
                }
            }
        }

        stage("run") {
            steps {
                script {
                        echo "------ Start running ------"
                        sh "sh run.sh"
                    }
            }
        }
    }
}