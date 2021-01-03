FROM golang:1.14-buster

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y -q \
    fontconfig \
    xfonts-utils \
    supervisor \
    imagemagick \
    unoconv \
    && apt-get clean -y \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir -p /usr/share/fonts/win

ADD simsun.ttc /usr/share/fonts/win/

RUN mkfontscale \
    && mkfontdir \
    && fc-cache

RUN sed -i 's/rights="none" pattern="PDF"/rights="read|write" pattern="PDF"/' /etc/ImageMagick-6/policy.xml
RUN sed -i 's/name="height" value="16KP"/name="height" value="48KP"/' /etc/ImageMagick-6/policy.xml

ADD conf/supervisord.conf /etc/supervisor/
ADD conf/libreoffice.conf /etc/supervisor/conf.d/
ADD conf/preview.conf /etc/supervisor/conf.d/

WORKDIR /go/src/app
ADD . .
ENV GOPROXY https://goproxy.io
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["supervisord","-n"]
