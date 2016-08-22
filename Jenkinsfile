node {
    stage 'fetch'
    git changelog: false, poll: false, url: 'https://github.com/HYmian/webDemo.git'

    stage 'build'
    docker.image("golang:1.7-alpine").inside {
        echo pwd()
        git 'https://github.com/HYmian/webDemo.git'
        sh 'mv $PWD/vendor/* /go/src/'
        sh 'go build'
    }
    input '构建完成，是否继续打包？'

    stage 'pack'
    def im = docker.build('ymian/webapp', '.')
    
    stage 'test'
    sh "docker-compose -f docker-compose-test.yml up -d"
    sh "./test-docker-compose.sh"
    sh "docker-compose -f docker-compose-test.yml stop"
    sh "docker-compose -f docker-compose-test.yml rm -f"

    stage 'deploy'
    sh "rancher-compose -p webDemo -e env.conf up -d"

    stage 'publish'
    docker.withRegistry('https://index.docker.io/v1/', '505') {
        im.push('latest')
    }
    echo "yeah!"
}
