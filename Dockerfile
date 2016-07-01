FROM golang

WORKDIR /var

ADD boot.sh .
ADD webDemo .
ADD templates/* templates/

CMD ["./boot.sh"]
