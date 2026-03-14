pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                script {
                    docker.image('golang:1.21').inside("-v /c/ProgramData/Jenkins/.jenkins/workspace/devops-security-platform:/workspace") {
                        bat 'go build'
                    }
                }
            }
        }
    }
}