pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build') {
            steps {
                script {
                    docker.image('golang:1.21').inside("-v /c/ProgramData/Jenkins/.jenkins/workspace/devops-security-platform:/workspace -w /workspace") {
                        bat 'go build'
                    }
                }
            }
        }
    }
}