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
              nodeSelector:
              #  workload_type: spot
              containers:
              - name: golang
                image: golang:1.12
                command:
                - cat
                tty: true
              - name: kaniko
                image: registry.cn-beijing.aliyuncs.com/acs-sample/jenkins-slave-kaniko:0.6.0
                command:
                - cat
                tty: true
                volumeMounts:
                - name: ymian
                  mountPath: /root/.docker
              - name: kubectl
                image: roffe/kubectl:v1.13.2
                command:
                - cat
                tty: true
              - name: busybox
                image: ymian/busybox
                command:
                - cat
                tty: true
              volumes:
              - name: ymian
                secret:
                  # you can replace secret to yours
                  secretName: ymian
                  items:
                  - key: .dockerconfigjson
                    path: config.json
            """
        }
    }

    stages {
        stage('Build') {
            steps {
                container('golang') {
                    git url: 'https://github.com/HYmian/gin-sample.git'
                    sh """
                    go build
                    """
                }
            }
        }

        stage('Image Build And Publish') {
            steps {
                container("kaniko") {
                    // you can replace `--destination=ymian/gin-sample` to yours
                    sh "kaniko -f `pwd`/Dockerfile -c `pwd` -d ymian/gin-sample"
                }
            }
        }

        stage('Deploy') {
            steps {
                container("kubectl") {
                    withKubeConfig(
                        [
                            // you can replace `mo` to yours
                            credentialsId: 'm0', 
                            serverUrl: 'https://kubernetes.default.svc.cluster.local'
                        ]
                    ) {
                        sh 'kubectl apply -f `pwd`/deploy.yaml'
                        sh 'kubectl wait --for=condition=Ready pod -l app=gin-sample --timeout=60s'
                    }
                }
            }
        }

        stage('Test') {
            steps {
                container("busybox") {
                    sh """
                    curl -d '{"message":"this is my first webhook"}' -H "Content-Type: application/json" -X POST webhook-gateway-svc.argo-events.svc.cluster.local:12000/foo
                    """
                }
            }
        }
    }
}
