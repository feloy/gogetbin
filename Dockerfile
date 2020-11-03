FROM golang:1.15
RUN apt update -y && apt install -y git  upx

COPY main.go build.sh ./

CMD go run main.go

