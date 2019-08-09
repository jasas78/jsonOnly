package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	id          string // "5VDDdJYDEdQ",
	thumbnail   string // "https://i.ytimg.com/vi/5VDDdJYDEdQ/maxresdefault.jpg",
	fulltitle   string // "唯有实行普选 才能保障人权",
	format_id   string // "136+251",
	uploader    string // "唐柏桥",
}

type _STrec struct {
	filesize                            int
	format_id, ext, vcodec, acodec, url string
}

var (
	_jsonFile     *os.File
	_filenameJson string = "1/1.json"
	_jsonByte     []byte
	_err          error
	_vstYT00      map[string]interface{}
	_vstYT10      []string = []string{"description", "id", "webpage_url", "thumbnail", "fulltitle", "format_id", "upload_date", "uploader"}
	_recLen       int
	_recArr       []_STrec
	_s300         string = ""
	_s400         string = ""
)

// func Unmarshal(data []byte, v interface{}) error

func _readJsonFile() {
	// Open our _jsonFile
	_jsonFile, _err = os.Open(_filenameJson)

	// if we os.Open returns an error then handle it
	if _err != nil {
		fmt.Printf("Opened <%s> failed\n", _filenameJson)
		fmt.Println(_err)
		os.Exit(110)
	}

	fmt.Printf("Successfully Opened <%s>.\n", _filenameJson)

	// func ReadAll(r io.Reader) ([]byte, error)
	_jsonByte, _err = ioutil.ReadAll(_jsonFile)

	if _err != nil {
		fmt.Printf("ReadAll <%s> failed\n", _filenameJson)
		fmt.Println(_err)
		os.Exit(111)
	}
	fmt.Printf("Successfully ReadAll json <%s>.\n", _filenameJson)

	// defer the closing of our _jsonFile so that we can parse it later on
	defer _jsonFile.Close()
}

func _analyzeJsonObj() {
	___valid := json.Valid(_jsonByte)
	fmt.Printf(" valid JSON <%v> \n", ___valid)

	//fmt.Printf(" JSON obj <%d:%v> \n" , len(_jsonByte), _jsonByte )

	_err = json.Unmarshal(_jsonByte, &_vstYT00)

	if _err != nil {
		fmt.Printf("analyze JSON <%s> failed\n", _filenameJson)
		fmt.Println(_err)
		os.Exit(121)
	}

	if 123 == 123 {
		//fmt.Printf(" JSON description 00 <%v> \n" , _vstYT00["description"] )
		for ___idx, ___key := range _vstYT10 {
			fmt.Printf("%d= %s:=<%v> \n", 10000+___idx, ___key, _vstYT00[___key])
		}
	}

	var _formatsNode interface{} = _vstYT00["formats"]
	if nil != _formatsNode {
		fmt.Printf(" formats not nil \n")
		//_formatArrs := _formatsNode.([]map[string]interface{})
		_formatArrs := _formatsNode.([]interface{})

		_recLen = len(_formatArrs)
		fmt.Printf(" formats len %d \n", _recLen)
		_recArr = make([]_STrec, _recLen)

		for ___idx, ___fNode := range _formatArrs {
			_s300 += fmt.Sprintf("%d= ", 30000+___idx)
			___fMap := ___fNode.(map[string]interface{})
			_recArr[___idx] = _STrec{}
			for ___key, ___value := range ___fMap {
				switch ___key {
				case "filesize":
					{
						switch vv := ___value.(type) {
						case float64:
							_s300 += fmt.Sprintf(" <%s:%v>", ___key, int(vv))
							_recArr[___idx].filesize = int(vv)

						}
					}
				case "vcodec", "acodec", "format_id", "ext":
					{
						_s300 += fmt.Sprintf(" <%s:%s>", ___key, ___value)
						switch ___key {
						case "vcodec":
							{
								_recArr[___idx].vcodec = fmt.Sprintf("%s", ___value)
							}
						case "acodec":
							{
								_recArr[___idx].acodec = fmt.Sprintf("%s", ___value)
							}
						case "format_id":
							{
								_recArr[___idx].format_id = fmt.Sprintf("%s", ___value)
							}
						case "ext":
							{
								_recArr[___idx].ext = fmt.Sprintf("%s", ___value)
							}
						}
					}
				case "url":
					{
						var urlV string = fmt.Sprintf("%s", ___value)
						_s300 += fmt.Sprintf(" <%s:%s>", ___key, urlV[:18])
						_recArr[___idx].url = urlV
					}
				}
			}
			_s300 += fmt.Sprintf("\n")
		}
		if 30000 == 30001 {
			fmt.Printf("%s", _s300)
		}

		for ___idx := 0; ___idx < _recLen; ___idx++ {
			_s400 += fmt.Sprintf("%d=", 40000+___idx)
			_s400 += fmt.Sprintf(" %3s", _recArr[___idx].format_id)
			_s400 += fmt.Sprintf(" %9d", _recArr[___idx].filesize)
			_s400 += fmt.Sprintf(" %4s", _recArr[___idx].ext)
			_s400 += fmt.Sprintf(" %13s", _recArr[___idx].vcodec)
			_s400 += fmt.Sprintf(" %13s", _recArr[___idx].acodec)
			_s400 += fmt.Sprintf(" %s", _recArr[___idx].url)
			_s400 += fmt.Sprintf("\n")
		}
		if 40000 == 40000 {
			fmt.Printf("%s", _s400)
		}
	}
}

// fmt.Printf("#!/bin/bash\n\n")

func main() {
	/*
	   fmt.Println("Hello World")
	*/
	_readJsonFile()

	_analyzeJsonObj()

}
