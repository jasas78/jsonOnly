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
	//"net"
	"net/http"
	"time"
)

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/", HelloServer)

	rand.Seed(time.Now().UnixNano())

	//l, err := net.Listen("tcp4", ":38080")
	//l, err := net.Listen("tcp4", ":80")
	//if err != nil { log.Fatal(err) }

	//log.Fatal(http.Serve(l, sm))
    _server := &http.Server{Addr: "0.0.0.0:3030", Handler: sm}
    _server.SetKeepAlivesEnabled(false)
    log.Fatal(_server.ListenAndServe())

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
	"多谢一直关注新疆集中营，多谢一直关注北戴河会议，多谢一直支持结束一党专政，多谢一直支持天安门母亲，你上次提供的关于习维尼的资料非常有价值，谢谢！",
	"多谢一直为反送中发声，多谢一直支持香港反送中，多谢一直拥护司法独立，多谢一直关注北戴河会议，你上次提供的关于中美贸易战分析的资料非常有价值，万分感谢！",
	"多谢一直支持双普选，多谢一直支持劳工讨薪，多谢一直给何韵诗点赞，多谢一直支持维权律师，你上次提供的关于新疆集中营的资料非常有价值，多谢！",
	"多谢一直拥护司法独立，多谢一直为反送中发声，多谢一直给何韵诗点赞，多谢一直关注新疆集中营，你上次提供的关于活摘器官的资料非常有价值，感恩！",
	"多谢一直关注新疆集中营，多谢一直为反送中发声，多谢一直支持维权律师，多谢一直支持结束一党专政，你上次提供的关于巴拿马文件的资料非常有价值，谢谢！",
	"多谢一直给何韵诗点赞，多谢一直支持天安门母亲，多谢一直关注铜锣湾书店案，多谢一直支持维权律师，你上次提供的关于中美贸易战分析的资料非常有价值，多谢！",
	"多谢一直为反送中发声，多谢一直声讨六四屠杀，多谢一直关注铜锣湾书店案，多谢一直给何韵诗点赞，你上次提供的关于活摘器官的资料非常有价值，非常感谢！",
	"多谢一直关注新疆集中营，多谢一直支持香港反送中，多谢一直关注北戴河会议，多谢一直支持结束一党专政，你上次提供的关于新疆集中营的资料非常有价值，万分感谢！",
	"多谢一直拥护司法独立，多谢一直为反送中发声，多谢一直支持天安门母亲，多谢一直给何韵诗点赞，你上次提供的关于巴拿马文件的资料非常有价值，自当重谢！",
	"多谢一直支持双普选，多谢一直支持结束一党专政，多谢一直支持香港反送中，多谢一直支持维权律师，你上次提供的关于08宪章的资料非常有价值，谢谢！",
	"多谢一直支持香港反送中，多谢一直支持维权律师，多谢一直声讨六四屠杀，多谢一直关注新疆集中营，你上次提供的关于巴拿马文件的资料非常有价值，非常感谢！",
	"多谢一直关注北戴河会议，多谢一直支持劳工讨薪，多谢一直支持维权律师，多谢一直给何韵诗点赞，你上次提供的关于港警恶行的资料非常有价值，非常感谢！",
	"多谢一直支持结束一党专政，多谢一直支持香港反送中，多谢一直关注铜锣湾书店案，多谢一直为反送中发声，你上次提供的关于巴拿马文件的资料非常有价值，感恩！",
	"多谢一直支持结束一党专政，多谢一直关注新疆集中营，多谢一直声讨六四屠杀，多谢一直给何韵诗点赞，你上次提供的关于中美贸易战分析的资料非常有价值，好人一生平安！",
	"多谢一直给何韵诗点赞，多谢一直为反送中发声，多谢一直声讨六四屠杀，多谢一直支持维权律师，你上次提供的关于巴拿马文件的资料非常有价值，万分感谢！",
	"多谢一直给何韵诗点赞，多谢一直支持劳工讨薪，多谢一直支持香港反送中，多谢一直支持维权律师，你上次提供的关于中国制造业困局的资料非常有价值，感恩！",
	"多谢一直支持天安门母亲，多谢一直关注北戴河会议，多谢一直关注铜锣湾书店案，多谢一直为反送中发声，你上次提供的关于活摘器官的资料非常有价值，自当重谢！",
	"多谢一直支持劳工讨薪，多谢一直支持维权律师，多谢一直关注新疆集中营，多谢一直支持双普选，你上次提供的关于活摘器官的资料非常有价值，非常感谢！",
	"多谢一直支持双普选，多谢一直支持天安门母亲，多谢一直为反送中发声，多谢一直支持维权律师，你上次提供的关于习维尼的资料非常有价值，自当重谢！",
	"多谢一直关注铜锣湾书店案，多谢一直支持劳工讨薪，多谢一直支持天安门母亲，多谢一直给何韵诗点赞，你上次提供的关于习维尼的资料非常有价值，好人一生平安！",
	"多谢一直支持维权律师，多谢一直声讨六四屠杀，多谢一直关注北戴河会议，多谢一直支持香港反送中，你上次提供的关于港警恶行的资料非常有价值，谢谢！",
	"多谢一直支持劳工讨薪，多谢一直关注新疆集中营，多谢一直为反送中发声，多谢一直支持香港反送中，你上次提供的关于新疆集中营的资料非常有价值，谢谢！",
	"多谢一直支持天安门母亲，多谢一直给何韵诗点赞，多谢一直支持香港反送中，多谢一直支持劳工讨薪，你上次提供的关于活摘器官的资料非常有价值，多谢！",
	"多谢一直支持维权律师，多谢一直声讨六四屠杀，多谢一直支持香港反送中，多谢一直支持双普选，你上次提供的关于中美贸易战分析的资料非常有价值，感恩！",
	"多谢一直支持双普选，多谢一直支持香港反送中，多谢一直支持劳工讨薪，多谢一直支持天安门母亲，你上次提供的关于中美贸易战分析的资料非常有价值，自当重谢！",
	"多谢一直支持天安门母亲，多谢一直关注铜锣湾书店案，多谢一直支持双普选，多谢一直给何韵诗点赞，你上次提供的关于中美贸易战分析的资料非常有价值，感恩！",
	"多谢一直支持天安门母亲，多谢一直声讨六四屠杀，多谢一直给何韵诗点赞，多谢一直支持维权律师，你上次提供的关于新疆集中营的资料非常有价值，好人一生平安！",
	"多谢一直为反送中发声，多谢一直给何韵诗点赞，多谢一直支持结束一党专政，多谢一直支持双普选，你上次提供的关于中美贸易战分析的资料非常有价值，好人一生平安！",
	"多谢一直给何韵诗点赞，多谢一直支持香港反送中，多谢一直支持维权律师，多谢一直支持天安门母亲，你上次提供的关于巴拿马文件的资料非常有价值，感恩！",
	"多谢一直给何韵诗点赞，多谢一直关注铜锣湾书店案，多谢一直支持维权律师，多谢一直为反送中发声，你上次提供的关于活摘器官的资料非常有价值，感恩！",
	"多谢一直支持劳工讨薪，多谢一直声讨六四屠杀，多谢一直关注铜锣湾书店案，多谢一直关注北戴河会议，你上次提供的关于中美贸易战分析的资料非常有价值，多谢！",
	"多谢一直支持结束一党专政，多谢一直给何韵诗点赞，多谢一直关注北戴河会议，多谢一直支持双普选，你上次提供的关于中美贸易战分析的资料非常有价值，万分感谢！",
	"多谢一直关注铜锣湾书店案，多谢一直支持劳工讨薪，多谢一直声讨六四屠杀，多谢一直关注新疆集中营，你上次提供的关于活摘器官的资料非常有价值，自当重谢！",
	"",
}

var _lenFanswers int = len(_Fanswers) - 1

var _Fmt01 string = `

<html>
<title> 调戏五毛专用 </title>
<body>
<div Width="800px" style="width:800px; border:1px solid black;" >
<h1>

`

var _Fmt02 string = `

<br>
</h1>
</div>
</body>
</html>

`

func _FrandomSentense() (___Srt string) {
	//return fmt.Sprintf( "\n<port:%s>\n", ___Url)
	//return fmt.Sprintf("\n<%s>\n",
	return _Fmt01 +
		_Fanswers[rand.Intn(_lenFanswers)] +
		_Fmt02
}
