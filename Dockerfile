FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/inquisitor /opt/bin/inquisitor

CMD ["/opt/bin/inquisitor","server","-b","0.0.0.0","-p","80"]
