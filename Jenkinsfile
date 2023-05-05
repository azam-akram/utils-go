pipeline {
    agent any
    
    tools {
        go '1.18'
    }


    stages {
        stage('Go Version') {
            steps {
                sh 'go version'
            }
        }

        stage('Go Vet') {
            steps {
                sh 'go vet ./...'
            }
        }
        
        stage('Go Test') {
            steps {
                sh 'go test ./...'
            }
        }

    }
}