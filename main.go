package main

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		path := strings.TrimLeft(html.EscapeString(r.URL.Path), "/")

		redirectRoot := os.Getenv("REDIRECT_ROOT")
		if redirectRoot != "" && path == "" {
			w.Header().Add("Location", redirectRoot)
			w.WriteHeader(http.StatusMovedPermanently)
			return
		}

		cmd := exec.Command("./build.sh", path)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := cmd.Start(); err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		outbuf := new(bytes.Buffer)
		outbuf.ReadFrom(stdout)

		errbuf := new(bytes.Buffer)
		errbuf.ReadFrom(stderr)

		if err := cmd.Wait(); err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				exitCode := exitError.ExitCode()
				switch exitCode {
				case 4:
					w.WriteHeader(http.StatusNotFound)
				case 5:
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			fmt.Println(err)
			io.Copy(os.Stderr, errbuf)
			return
		}

		w.Header().Add("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, outbuf)
		io.Copy(os.Stderr, errbuf)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
