pipeline {
    agent {
        kubernetes {
            label 'jenkins-pod'
            defaultContainer 'jnlp'
            yaml """
            apiVersion: v1
            kind: Pod
            metadata:
              labels:
                app: jenkins-slave-pod
            spec:
                containers:
                - name: golang
                  image: "go:1.11"
                  workingDir: /home/jenkins
                  command:
                  - cat
                  tty: true
                - name: kaniko
                  image: "registry.cn-beijing.aliyuncs.com/acs-sample/jenkins-slave-kaniko:0.6.0"
                  workingDir: /home/jenkins
                  command:
                  - cat
                  tty: true
            """
        }
    }

    stages {
        stage('Git'){
            steps{
                git branch: 'master', credentialsId: '', url: 'https://github.com/HYmian/webDemo.git'
            }
        }

        stage('Build') {
            steps {
                container('golang') {
                    sh 'go build'
                }
            }
        }
      }
}
