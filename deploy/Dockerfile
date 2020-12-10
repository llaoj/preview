FROM registry-vpc.cn-beijing.aliyuncs.com/llaoj/go-unoconv:1.14-0.4

WORKDIR /go/src/app
ADD . .
ADD deploy/conf/oconv.conf /etc/supervisor/conf.d/

ENV GOPROXY https://goproxy.io

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["supervisord","-n"]
