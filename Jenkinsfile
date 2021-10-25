pipeline {
    agent any
    tools {
        go 'go1.17.2'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {        
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }

        stage('Release') {
            when {
                buildingTag()
            }
            steps {
                sh "echo 'COMPLETED!!'"   
            }
        }
    }   
}