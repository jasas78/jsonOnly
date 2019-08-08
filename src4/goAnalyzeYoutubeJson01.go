package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"

	//"io"
	//"log"
	//"strings"
)

type _STyt01 struct {
    //season_number string `json:"season_number"` // 
    //vcodec string `json:"vcodec"` //  "avc1.4d401f",
    //channel_id string // "UCeBfK3zzgqDQLhBovmguOuQ",
    //playlist_id string // "UUeBfK3zzgqDQLhBovmguOuQ",
    //channel_url string // "http://www.youtube.com/channel/UCeBfK3zzgqDQLhBovmguOuQ",
    //uploader_url string // "http://www.youtube.com/user/tanglung4",
    description string `json:"description"` // "6.16香港兩百萬人黑衣大遊行再度創造歷史。香港一週內舉行兩次百萬人大遊行，震撼整个世界。本次抗議人數佔香港
    id  string // "5VDDdJYDEdQ",
    thumbnail   string // "https://i.ytimg.com/vi/5VDDdJYDEdQ/maxresdefault.jpg",
    fulltitle   string // "唯有实行普选 才能保障人权",
    format_id   string // "136+251",
    uploader    string // "唐柏桥",
}


var 
(
    _jsonFile *os.File
    _filenameJson string = "1/1.json"
    _jsonByte []byte
    _err error
    _vstYT00 map[string]interface{}
    _vstYT10 []string = []string{ "description", "id", "webpage_url", "thumbnail", "fulltitle", "format_id", "upload_date", "uploader" }
)

// func Unmarshal(data []byte, v interface{}) error

func _readJsonFile() {
    // Open our _jsonFile
    _jsonFile, _err = os.Open( _filenameJson )

    // if we os.Open returns an error then handle it
    if _err != nil {
        fmt.Printf("Opened <%s> failed\n" , _filenameJson )
        fmt.Println(_err)
        os.Exit(110)
    }
    fmt.Printf("Successfully Opened <%s>.\n" , _filenameJson )

    // func ReadAll(r io.Reader) ([]byte, error)
    _jsonByte, _err = ioutil.ReadAll( _jsonFile )

    if _err != nil {
        fmt.Printf("ReadAll <%s> failed\n" , _filenameJson )
        fmt.Println(_err)
        os.Exit(111)
    }
    fmt.Printf("Successfully ReadAll json <%s>.\n" , _filenameJson )

    // defer the closing of our _jsonFile so that we can parse it later on
    defer _jsonFile.Close()
}

func _analyzeJsonObj() {
    ___valid := json.Valid( _jsonByte )
    fmt.Printf(" valid JSON <%v> \n" , ___valid )

    //fmt.Printf(" JSON obj <%d:%v> \n" , len(_jsonByte), _jsonByte )

    _err = json.Unmarshal(_jsonByte, &_vstYT00)

    if _err != nil {
        fmt.Printf("analyze JSON <%s> failed\n" , _filenameJson )
        fmt.Println(_err)
        os.Exit(121)
    }

    if 123 == 123 {
        //fmt.Printf(" JSON description 00 <%v> \n" , _vstYT00["description"] )
        for ___idx , ___key := range _vstYT10 {
            fmt.Printf("%d= %s:=<%v> \n" , 10000+___idx, ___key , _vstYT00[___key] )
        }
    }

}

func main() {
    /*
    fmt.Println("Hello World")
    */
    _readJsonFile()

    _analyzeJsonObj()

}
