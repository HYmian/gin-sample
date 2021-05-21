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
                        sh '''
                        kubectl apply -f `pwd`/deploy/deploy.yaml -n pro
                        kubectl wait --for=condition=Ready pod -l app=gin-sample --timeout=60s -n pro
                        '''
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
                        sh '''
                        kubectl apply -f `pwd`/deploy/deploy.yaml -n test
                        kubectl wait --for=condition=Ready pod -l app=gin-sample --timeout=60s -n test
                        '''
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
                    sh "./validate.sh"
                }
            }
        }
    }

    post {
        always {
            notifyBuild()
        }
    }
}


def notifyBuild() {
  def subject = "${currentBuild.result}: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'"
  def details = """<p>STARTED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':</p>
    <p>Check console output at &QUOT;<a href='${env.BUILD_URL}'>${env.JOB_NAME} [${env.BUILD_NUMBER}]</a>&QUOT;</p>"""

  emailext (
      subject: subject,
      body: details,
      recipientProviders: [developers(), buildUser(), requestor(), upstreamDevelopers()]
    )
}
