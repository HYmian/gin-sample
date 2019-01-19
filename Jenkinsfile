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
                  image: golang:1.11
                  command:
                  - cat
                  tty: true
                - name: kaniko
                  image: registry.cn-beijing.aliyuncs.com/acs-sample/jenkins-slave-kaniko:0.6.0
                  workingDir: /home/jenkins
                  command:
                  - cat
                  tty: true
            """
        }
    }

    stages {
        stage('Build') {
            steps {
                container('golang') {
                    git url: 'https://github.com/HYmian/webDemo.git'
                    sh """
                    mkdir -p /go/src/github.com/HYmian
                    ln -s `pwd` /go/src/github.com/HYmain/webDemo
                    cd /go/src/github.com/hashicorp/webDemo && go build
                    """
                }
            }
        }
      }
}
