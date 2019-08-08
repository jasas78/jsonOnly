//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//func main() {
//	http.HandleFunc("/", HelloServer)
//    http.ServeType = "tcp4"
//	//http.ListenAndServe("0.0.0.0:38080", nil)
//	http.ListenAndServe("0.0.0.0:80", nil)
//}

// https://play.golang.org/p/EOZkK1UUpe
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/", HelloServer)

	//l, err := net.Listen("tcp4", ":38080")
	l, err := net.Listen("tcp4", ":80")

	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, sm))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	/*
		fmt.Fprintf(w, "Hello, path:[ /%s ] ! [%s]",
			r.URL.Path[1:],
			r.RemoteAddr)
	*/
	fmt.Fprintf(w, r.RemoteAddr)
}
