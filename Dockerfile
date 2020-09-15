FROM golang:1.15
RUN apt update && apt install git build-base upx

COPY main.go build.sh ./

CMD go run main.go

