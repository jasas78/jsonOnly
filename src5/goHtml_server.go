package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":38080", nil)
	//http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	/*
		fmt.Fprintf(w, "Hello, path:[ /%s ] ! [%s]",
			r.URL.Path[1:],
			r.RemoteAddr)
	*/
	fmt.Fprintf(w, r.RemoteAddr)
}
