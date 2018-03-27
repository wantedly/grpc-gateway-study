FROM golang:1.10.0-alpine3.7

RUN apk add --no-cache --update \
  ca-certificates \
  openssl \
  && update-ca-certificates


#  protobuf
#-----------------------------------------------
RUN apk add --no-cache --update \
  protobuf \
  protobuf-dev


#  glibc (https://github.com/sgerrand/alpine-pkg-glibc)
#-----------------------------------------------
# @see http://qiita.com/tkusumi/items/1668df2fc4cf61a68282
RUN apk add --no-cache --update --virtual .build-deps \
  wget \
  && wget -q -O /etc/apk/keys/sgerrand.rsa.pub \
  https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub \
  && wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.25-r0/glibc-2.25-r0.apk \
  && apk del .build-deps \
  && apk add glibc-2.25-r0.apk


#  Build dependencies
#-----------------------------------------------
RUN apk add --no-cache --update \
  build-base \
  curl


#  dep
#-----------------------------------------------
ENV DEP_VERSION 0.4.1

RUN curl -fL "https://github.com/golang/dep/releases/download/v$DEP_VERSION/dep-linux-amd64" -o dep \
  && chmod u+x dep \
  && mv ./dep /usr/local/bin/dep


#  grapi
#-----------------------------------------------
ENV GRAPI_VERSION 0.2.0

RUN curl -fL "https://github.com/izumin5210/grapi/releases/download/v$GRAPI_VERSION/grapi_linux_amd64" -o grapi \
  && chmod u+x grapi \
  && mv ./grapi /usr/local/bin/grapi


#  migrate
#-----------------------------------------------
ENV MIGRATE_VERSION 3.2.0

RUN curl -fL "https://github.com/golang-migrate/migrate/releases/download/v$MIGRATE_VERSION/migrate.linux-amd64.tar.gz" -o migrate.tar.gz \
  && tar xf migrate.tar.gz \
  && mv migrate.linux-amd64 migrate \
  && chmod u+x migrate \
  && mv ./migrate /usr/local/bin/migrate


#  Library
#-----------------------------------------------
RUN apk add --no-cache --update \
  bash \
  coreutils \
  git \
  openssh-client \
  && go get -u github.com/golang/lint/golint


#  App
#-----------------------------------------------
ENV APP_DIR /go/src/github.com/wantedly/grpc-gateway-study
ENV PATH $APP_DIR/bin:$PATH
WORKDIR $APP_DIR
RUN ln -sf $APP_DIR /app

COPY ./entrypoint.sh /usr/bin/entrypoint.sh
ENTRYPOINT ["/usr/bin/entrypoint.sh"]
