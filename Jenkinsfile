pipeline {
    agent {
        docker { image 'golang:1.21' }
    }

    stages {
        stage('Clone') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o app .'
            }
        }

        stage('Run Docker') {
            steps {
                sh 'docker build -t do-host-network .'
                sh 'docker run -d -p 8000:8000 do-host-network'
            }
        }
    }
}
