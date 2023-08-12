pipeline {
    // install golang 1.20 on Jenkins node
    agent any
    tools {
        go 'go1.20'
    }
    environment {
        GO120MODULE = 'on'
        CGO_ENABLED = 0 
    }
    stages {

        
        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
                sh 'go get ./go.mod'
                sh 'make unit-tests'
            }
        }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'make run'
            }
        }
    }
}