package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	port := flag.String("port", "8000", "serve port=8000")
	flag.Parse()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", LogReq(http.FileServer(http.Dir(dir))))
	fmt.Println("Listening on port:", *port)
	http.ListenAndServe(":"+*port, nil)
}

func LogReq(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		handler.ServeHTTP(w, r)
	}

}
