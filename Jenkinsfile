pipeline {
    agent any
    
    stages {
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