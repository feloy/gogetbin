FROM golang:1.13

COPY main.go build.sh ./

CMD go run main.go

