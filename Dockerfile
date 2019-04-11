FROM golang:1.10.4
LABEL maintainer="Prasad Kavinda <pp.kavinda@gmail.com>"
WORKDIR $GOPATH/src/github.com/ppkavinda/drive-torrent
COPY . .
RUN go get github.com/anacrolix/dht
RUN go get github.com/anacrolix/torrent
RUN go get github.com/ppkavinda/drive-torrent/db
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/websocket
RUN go get golang.org/x/net/context
RUN go get golang.org/x/oauth2
RUN go get google.golang.org/api/drive/v3