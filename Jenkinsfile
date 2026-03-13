pipeline {
    agent any

    stages {

        stage('Build Go App') {
            steps {
                bat 'go build -o auth-service.exe ./auth-service'
            }
        }

        stage('Build Docker Image') {
            steps {
                bat 'docker build -t devops-auth-service .'
            }
        }

        stage('Remove Old Container') {
            steps {
                bat 'docker rm -f auth-container || exit 0'
            }
        }

        stage('Run Container') {
            steps {
                bat 'docker run -d -p 9095:9092 --name auth-container devops-auth-service'
            }
        }

    }
}