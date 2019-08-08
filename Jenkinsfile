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
              annotations:
                k8s.aliyun.com/eci-cpu: 2
                k8s.aliyun.com/eci-memory: 4Gi
            spec:
#              nodeSelector:
#                type: virtual-kubelet
#              tolerations:
#              - key: virtual-kubelet.io/provider
#                operator: Exists
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
                  - key: config.json
                    path: config.json
            """
        }
    }

    stages {
        stage('Build') {
            steps {
                container('golang') {
                    sh """
                    go build -mod vendor -v
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

        stage('Deploy to pro') {
            when {
              branch "master"
            }
            steps {
                container("kubectl") {
                    withKubeConfig(
                        [
                            // you can replace `mo` to yours
                            credentialsId: 'pro-env',
                            serverUrl: 'https://kubernetes.default.svc.cluster.local'
                        ]
                    ) {
                        sh 'kubectl apply -f `pwd`/deploy.yaml -n pro'
                        sh 'kubectl wait --for=condition=Ready pod -l app=gin-sample --timeout=60s'
                    }
                }
            }
        }

        stage('Deploy other') {
            when {
              not { branch "master" }
            }
            steps {
                container("kubectl") {
                    withKubeConfig(
                        [
                            // you can replace `mo` to yours
                            credentialsId: 'test-env',
                            serverUrl: 'https://kubernetes.default.svc.cluster.local'
                        ]
                    ) {
                        sh 'kubectl apply -f `pwd`/deploy.yaml -n test'
                        sh 'kubectl wait --for=condition=Ready pod -l app=gin-sample --timeout=60s'
                    }
                }
            }
        }

        stage('Test') {
            when {
              not { branch "master" }
            }
            steps {
                container("busybox") {
                    sh """
                    x=`curl http://webdemo.default.svc.cluster.local:3000/stress/3 -w '%{size_download}' -so /dev/null`;
                    if [ $x -eq '3072' ]; then
                        exit 0;
                    else
                        echo x=$x;
                        exit 1;
                    fi
                    """
                }
            }
        }
    }
}
