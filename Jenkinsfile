node {
    stage 'fetch'
    git changelog: false, poll: false, url: 'https://github.com/HYmian/webDemo.git'

    stage 'build'
    def Go = tool name: 'Go1.6', type: 'org.jenkinsci.plugins.golang.GolangInstallation'
    withEnv(["GOROOT=${Go}", "GOPATH=${env.JENKINS_HOME}/Go"]) {
        withEnv(["PATH=${env.GOROOT}/bin:${env.GOPATH}/bin:${env.PATH}"]) {
        sh "go get -u github.com/astaxie/beego/orm"
        sh "go get -u github.com/go-martini/martini"
        sh "go get -u github.com/golang/glog"
        sh "go get -u github.com/martini-contrib/render"
        sh "go get -u github.com/go-sql-driver/mysql"
	sh "go get -u github.com/martini-contrib/binding"
        
        sh "go test"
        sh "go build"
        }
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
