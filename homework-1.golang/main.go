package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("start...")
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		fmt.Println("error...")
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}
