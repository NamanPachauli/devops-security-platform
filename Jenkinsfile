pipeline {
    agent any

    stages {

        stage('Build') {
            steps {

                bat '''
                docker run --rm -v C:/ProgramData/Jenkins/.jenkins/workspace/devops-security-platform:/workspace -w /workspace golang:1.25 go version
                docker run --rm -v C:/ProgramData/Jenkins/.jenkins/workspace/devops-security-platform:/workspace -w /workspace golang:1.25 go build -o app.exe
                '''

            }
        }

    }
}
