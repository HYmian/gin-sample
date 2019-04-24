FROM centos:7

WORKDIR /var

ADD boot.sh .
ADD gin-sample .
ADD templates/* templates/

CMD ["./boot.sh"]
