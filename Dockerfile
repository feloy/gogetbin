FROM golang

COPY main.go build.sh ./

CMD go run main.go

