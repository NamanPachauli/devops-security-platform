pipeline {
    agent {
        docker { 
            image 'golang:1.21'  // Official Go image
            args '-v /var/run/docker.sock:/var/run/docker.sock' // To run Docker inside container
        }
    }

    environment {
        APP_NAME = "devops-security-platform"
        IMAGE_NAME = "devops-security-platform-image"
        CONTAINER_NAME = "devops-security-platform-container"
        PORT = "9090"
    }

    stages {
        stage('Checkout Code') {
            steps {
                git branch: 'main', url: 'https://github.com/NamanPachauli/devops-security-platform.git'
            }
        }

        stage('Build Go App') {
            steps {
                echo "Building Go application..."
                sh 'go mod tidy'
                sh 'go build -o $APP_NAME ./...'
            }
        }

        stage('Build Docker Image') {
            steps {
                echo "Building Docker image..."
                sh 'docker build -t $IMAGE_NAME .'
            }
        }

        stage('Remove Old Container') {
            steps {
                echo "Removing old container if exists..."
                sh '''
                if [ $(docker ps -a -q -f name=$CONTAINER_NAME) ]; then
                    docker rm -f $CONTAINER_NAME
                fi
                '''
            }
        }

        stage('Run Container') {
            steps {
                echo "Running Docker container..."
                sh 'docker run -d -p $PORT:$PORT --name $CONTAINER_NAME $IMAGE_NAME'
            }
        }
    }

    post {
        success {
            echo "Pipeline completed successfully! ✅"
        }
        failure {
            echo "Pipeline failed. ❌"
        }
    }
}