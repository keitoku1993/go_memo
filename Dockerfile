# golangのイメージ取得
FROM golang:1.9.2

# ソースのコピー先
ENV SRC_DIR=/go/src/github.com/gouser

# コンパイルしたファイルの格納場所
ENV GOBIN=/go/bin

# RUN,やENTRYPOINTなどを実行するワークディレクトリを指定
WORKDIR $GOBIN

# ソースをコピーする
ADD . $SRC_DIR

# コンパイル
RUN cd /go/src;
RUN go install github.com/gouser/;

# プログラム実行
ENTRYPOINT [ "./gouser" ]

# 8080ポートを解放
EXPOSE 7070
