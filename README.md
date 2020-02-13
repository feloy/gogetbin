# gogetbin

A container that listens for go path and returns built binary

## Run locally

```sh
$ docker build . -t gogetbin
$ docker run -d -p 8080:8080 gogetbin
$ curl http://localhost:8080/github.com/github/hub > hub
$ chmod +x hub
$ ./hub
```

## Run gloabally from Google Cloud Run

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)
