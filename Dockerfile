FROM centos:7.6.180

WORKDIR /var

ADD gin-sample .
ADD templates/* templates/

CMD ["./gin-sample", "-logtostderr", "-v 2"]
