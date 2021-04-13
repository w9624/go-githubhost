package main

import (
	"fmt"
	"log"
	"net/http"

	"go-githubhost/runner"
)

func main() {
	_ = Run()
}

func Run() (err error) {
	// http server
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		text, err := runner.Run()
		if err != nil {
			_, _ = fmt.Fprintf(writer, fmt.Sprintf("异常: %v", err))
		}
		_, _ = fmt.Fprintln(writer, text)
	})

	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}

	return
}
