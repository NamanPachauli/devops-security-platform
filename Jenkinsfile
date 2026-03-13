pipeline {
    agent any

    stages {

        stage('Clone Repository') {
            steps {
                git 'https://github.com/YOUR_GITHUB_USERNAME/devops-security-platform.git'
            }
        }

        stage('Build Go App') {
            steps {
                bat 'go build -o auth-service.exe'
            }
        }

        stage('Build Docker Image') {
            steps {
                bat 'docker build -t devops-auth-service .'
            }
        }

        stage('Run Container') {
            steps {
                bat 'docker run -d -p 9095:9092 --name auth-container devops-auth-service'
            }
        }

    }
}