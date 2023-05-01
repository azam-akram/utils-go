pipeline {
    agent any
    
    tools {
        go '1.16.13'
    }

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