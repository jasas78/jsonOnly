package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "fmt"
    "log"
)

func helloworld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Hello, World!")
}


func main() {

    router := httprouter.New()

    router.GET("/", helloworld)


    fmt.Println("Running :-)")

    //http.Server.SetKeepAlivesEnabled(false)
    server := &http.Server{Addr: ":3030", Handler: router}
    server.SetKeepAlivesEnabled(false)


    //log.Fatal(http.ListenAndServe(":80", router))
    log.Fatal(server.ListenAndServe())
}
