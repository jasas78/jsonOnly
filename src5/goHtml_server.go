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
	"math/rand"
	"net"
	"net/http"
	"time"
)

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/", HelloServer)

	rand.Seed(time.Now().UnixNano())

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
	switch r.Host {
	case "ttt.jjj123.com":
		{
			fmt.Fprintf(w, _FrandomSentense())
		}
	default:
		{
			fmt.Fprintf(w, r.RemoteAddr)
			fmt.Fprintf(w, "\n<%s>\n", r.Host)
		}
	}
}

var _Fanswers []string = []string{
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}
var _lenFanswers int = len(_Fanswers) - 1

func _FrandomSentense() (___Srt string) {
	//return fmt.Sprintf( "\n<port:%s>\n", ___Url)
	return fmt.Sprintf("\n<%s>\n", _Fanswers[rand.Intn(_lenFanswers)])
}
