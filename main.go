package main

import (
	"bytes"
	"html"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		path := strings.TrimLeft(html.EscapeString(r.URL.Path), "/")

		cmd := exec.Command("./build.sh", path)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(stdout)

		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
		io.Copy(w, buf)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
