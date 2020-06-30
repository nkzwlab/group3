FROM ubuntu:18.04

# Go用の環境変数を設定
ENV GOPATH $HOME/go
ENV PATH $PATH:/usr/lib/go-1.12/bin
ENV PATH $PATH:$GOPATH/bin

# 各種ツールをインストール
RUN apt-get update -y
RUN apt-get install -y software-properties-common
RUN add-apt-repository -y ppa:longsleep/golang-backports
RUN apt-get install -y git curl golang-1.12-go make mysql-client

RUN go get "github.com/julienschmidt/httprouter"\
           "github.com/go-sql-driver/mysql"


# working dir
RUN mkdir /api
WORKDIR /api

# コンテナとフォイルを共有
ADD . /api

#CMD make start
CMD make start
