FROM golang:1.10.4
LABEL maintainer="Prasad Kavinda <pp.kavinda@gmail.com>"

WORKDIR $GOPATH/src/github.com/ppkavinda/drive-torrent
COPY . .


RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD ["go","run","main.go"]
