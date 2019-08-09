package main

import (
	"bufio"
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
	protocol    string // "https",  "m3u8_native"
}

type _STrec struct {
	filesize                                      int
	format_id, ext, vcodec, acodec, url, protocol string
}

type _STdst struct {
	vo1_ao2_both3 int
	size          int
	idx           int
	url           string
	protocol      string
	ext           string
}

var (
	_jsonFile         *os.File
	_filenameJson     string = "________________.json"
	_jsonByte         []byte
	_err              error
	_vstYT00          map[string]interface{}
	_vstYT10          []string = []string{"description", "id", "webpage_url", "thumbnail", "fulltitle", "format_id", "upload_date", "uploader", "protocol"}
	_recLen           int
	_recArr           []_STrec
	_s300             string = ""
	_s400             string = ""
	_vDstMaxAllowVo   _STdst
	_vDstMaxAllowAo   _STdst
	_vDstMaxAllowBoth _STdst
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
				case "protocol":
					{
						var protocolV string = fmt.Sprintf("%s", ___value)
						_s300 += fmt.Sprintf(" <%s:%s>", ___key, protocolV)
						_recArr[___idx].protocol = protocolV
					}
				}
			}
			_s300 += fmt.Sprintf("\n")
		}
		if 30000 == 30001 {
			fmt.Printf("%s", _s300)
		}

		_vDstMaxAllowVo = _STdst{}
		_vDstMaxAllowAo = _STdst{}
		_vDstMaxAllowBoth = _STdst{}

		for ___idx := 0; ___idx < _recLen; ___idx++ {
			_s400 += fmt.Sprintf("#%d=", 40000+___idx)
			_s400 += fmt.Sprintf(" %3s", _recArr[___idx].format_id)
			_s400 += fmt.Sprintf(" %9d", _recArr[___idx].filesize)
			_s400 += fmt.Sprintf(" %4s", _recArr[___idx].ext)
			_s400 += fmt.Sprintf(" %13s", _recArr[___idx].vcodec)
			_s400 += fmt.Sprintf(" %13s", _recArr[___idx].acodec)
			_s400 += fmt.Sprintf(" %s", _recArr[___idx].url)
			_s400 += fmt.Sprintf("\n")

			__fSize := _recArr[___idx].filesize
			//if __fSize != 0 {
			if __fSize >= 0 {
				var __vSign = 0
				if _recArr[___idx].vcodec == "" || _recArr[___idx].vcodec == "none" {
					__vSign = 1
				}
				if _recArr[___idx].acodec == "" || _recArr[___idx].acodec == "none" {
					__vSign += 2
				}
				switch __vSign {
				case 0:
					{ // both vo & ao
						if __fSize < 50000000 && __fSize >= _vDstMaxAllowBoth.size {
							_vDstMaxAllowBoth.idx = ___idx
							_vDstMaxAllowBoth.vo1_ao2_both3 = 3
							_vDstMaxAllowBoth.size = __fSize
							_vDstMaxAllowBoth.url = _recArr[___idx].url
							_vDstMaxAllowBoth.protocol = _recArr[___idx].protocol
							_vDstMaxAllowBoth.ext = _recArr[___idx].ext
						}
					}
				case 1:
					{ // vo null , ao only
						if __fSize < 45000000 && __fSize >= _vDstMaxAllowAo.size {
							_vDstMaxAllowAo.idx = ___idx
							_vDstMaxAllowAo.vo1_ao2_both3 = 2
							_vDstMaxAllowAo.size = __fSize
							_vDstMaxAllowAo.url = _recArr[___idx].url
							_vDstMaxAllowAo.protocol = _recArr[___idx].protocol
							_vDstMaxAllowAo.ext = _recArr[___idx].ext
						}
					}
				case 2:
					{ // ao null , vo only
						if __fSize < 35000000 && __fSize >= _vDstMaxAllowVo.size {
							_vDstMaxAllowVo.idx = ___idx
							_vDstMaxAllowVo.vo1_ao2_both3 = 1
							_vDstMaxAllowVo.size = __fSize
							_vDstMaxAllowVo.url = _recArr[___idx].url
							_vDstMaxAllowVo.protocol = _recArr[___idx].protocol
							_vDstMaxAllowVo.ext = _recArr[___idx].ext
						}
					}
				case 3:
					{ // vo null , ao null , shit.
						fmt.Printf(" 100053 : not-vo, not-ao , what is this ? \n")
						os.Exit(143)
					}
				default:
					{ // unknown
						fmt.Printf(" 100059 : unknow what happpens.\n")
						os.Exit(149)
					}
				}
			}
		}
		if 40000 == 40000 {
			fmt.Printf("%s", _s400)
		}
	}
}

func _FgenVoAoLine(___w *bufio.Writer, ___srcVo, ___srcAo string) {
	__vFnameMp4 := _filenameJson + ".mp4"
	fmt.Fprintf(___w, "rm -f %s \n", __vFnameMp4)
	fmt.Fprintf(___w, "/usr/bin/ffmpeg -i %s  -i %s        -ac 1 -ab 25000   -map 0:v:0 -map 1:a:0      %s \n", ___srcVo, ___srcAo, __vFnameMp4)
	// https://superuser.com/questions/1137612/ffmpeg-replace-audio-in-video
}

func _FgenWavMp3Line(___w *bufio.Writer, ___src string) {
	__vFnameWav := _filenameJson + ".wav"
	__vFnameMp3 := _filenameJson + ".mp3"
	fmt.Fprintf(___w, "rm -f %s \n", __vFnameWav)
	fmt.Fprintf(___w, "/usr/bin/ffmpeg -i %s         -vn -ac 1 -ab 25000        %s \n", ___src, __vFnameWav)
	fmt.Fprintf(___w, "rm -f %s \n", __vFnameMp3)
	fmt.Fprintf(___w, "lame       %s       %s\n", __vFnameWav, __vFnameMp3)
}

func _FgetDownloadLine(___w *bufio.Writer, ___dst, ___src, ___protocol string) {
	switch ___protocol {
	case "https":
		{
			fmt.Fprintf(___w, "rm -f %s \n", ___dst)
			fmt.Fprintf(___w, "wget -c -O %s  '%s'\n", ___dst, ___src)
		}
	case "m3u8_native":
		{
			fmt.Fprintf(___w, "rm -f %s \n", ___dst)
			fmt.Fprintf(___w, "/usr/bin/ffmpeg -i '%s' %s \n", ___src, ___dst)
		}
	default:
		{
			fmt.Printf("\n\n unknown protocol :<%s>\n\n", ___protocol)
			os.Exit(179)
		}

	}
}

// func fmt.Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// func os.Create(name string) (*File, error)
// func bufio.NewWriter(w io.Writer) *Writer
func _genYoutubeDownloadScript() {
	__vFshName := os.Args[1] + ".sh"
	__vFileSh, __vErr := os.Create(__vFshName)
	if __vErr != nil {
		fmt.Printf("Create <%s> failed\n", __vFshName)
		fmt.Println(__vErr)
		os.Exit(140)
	}
	defer __vFileSh.Close()

	__vBfIoWriter := bufio.NewWriter(__vFileSh)
	fmt.Fprintf(__vBfIoWriter, "#!/bin/bash\n\n")

	fmt.Fprintf(__vBfIoWriter, "%s\n", _s400)

	fmt.Fprintf(__vBfIoWriter, "wget -c -O %s.jpg\n\n", os.Args[1])

	//fmt.Fprintf(__vBfIoWriter , "wget -c -O %s.jpg\n\n", os.Args[1] )

	fmt.Fprintf(__vBfIoWriter, "# Name : idx      size vo1_ao2_both3   ext   protocol\n")

	fmt.Fprintf(__vBfIoWriter, "# Vo   : %3d %9d %1d              %5s  '%s' \n",
		_vDstMaxAllowVo.idx,
		_vDstMaxAllowVo.size,
		_vDstMaxAllowVo.vo1_ao2_both3,
		_vDstMaxAllowVo.ext,
		_recArr[_vDstMaxAllowVo.idx].protocol)

	fmt.Fprintf(__vBfIoWriter, "# Ao   : %3d %9d %1d              %5s  '%s' \n",
		_vDstMaxAllowAo.idx,
		_vDstMaxAllowAo.size,
		_vDstMaxAllowAo.vo1_ao2_both3,
		_vDstMaxAllowAo.ext,
		_recArr[_vDstMaxAllowAo.idx].protocol)

	fmt.Fprintf(__vBfIoWriter, "# Both : %3d %9d %1d              %5s  '%s' \n",
		_vDstMaxAllowBoth.idx,
		_vDstMaxAllowBoth.size,
		_vDstMaxAllowBoth.vo1_ao2_both3,
		_vDstMaxAllowBoth.ext,
		_recArr[_vDstMaxAllowBoth.idx].protocol)

	if _vDstMaxAllowBoth.vo1_ao2_both3 == 0 {
		if _vDstMaxAllowAo.vo1_ao2_both3 != 0 &&
			_vDstMaxAllowVo.vo1_ao2_both3 != 0 {
			__ff1 := fmt.Sprintf("%s.vo.%s", _filenameJson, _vDstMaxAllowVo.ext)
			__ff2 := fmt.Sprintf("%s.ao.%s", _filenameJson, _vDstMaxAllowAo.ext)
			fmt.Fprintf(__vBfIoWriter, "# no-both, so , use vo + ao : %s + %s \n", __ff1, __ff2)
			_FgetDownloadLine(__vBfIoWriter, __ff1, _vDstMaxAllowVo.url, _vDstMaxAllowVo.protocol)
			_FgetDownloadLine(__vBfIoWriter, __ff2, _vDstMaxAllowAo.url, _vDstMaxAllowAo.protocol)

			_FgenWavMp3Line(__vBfIoWriter, __ff2)
			_FgenVoAoLine(__vBfIoWriter, __ff1, _filenameJson+".wav")
		} else {
			fmt.Printf("\n\n  no-both , AND no(ao + vo) , what happens ? 1838111. \n\n")
			os.Exit(180)
		}
	} else { // both exist.
		if _vDstMaxAllowVo.vo1_ao2_both3 != 0 { // use vo && both
			__ff1 := fmt.Sprintf("%s.vo.%s", _filenameJson, _vDstMaxAllowVo.ext)
			__ff2 := fmt.Sprintf("%s.bo.%s", _filenameJson, _vDstMaxAllowBoth.ext)
			fmt.Fprintf(__vBfIoWriter, "# both exist, vo exist , use vo + both : %s + %s \n", __ff1, __ff2)
			_FgetDownloadLine(__vBfIoWriter, __ff1, _vDstMaxAllowVo.url, _vDstMaxAllowVo.protocol)
			_FgetDownloadLine(__vBfIoWriter, __ff2, _vDstMaxAllowBoth.url, _vDstMaxAllowBoth.protocol)

			_FgenWavMp3Line(__vBfIoWriter, __ff2)
			_FgenVoAoLine(__vBfIoWriter, __ff1, _filenameJson+".wav")
		} else { // use both only
			__ff1 := fmt.Sprintf("%s.bo.%s", _filenameJson, _vDstMaxAllowBoth.ext)
			fmt.Fprintf(__vBfIoWriter, "# no vo , but both only, so , use both only : %s \n", __ff1)
			_FgetDownloadLine(__vBfIoWriter, __ff1, _vDstMaxAllowBoth.url, _vDstMaxAllowBoth.protocol)

			_FgenWavMp3Line(__vBfIoWriter, __ff1)
			_FgenVoAoLine(__vBfIoWriter, __ff1, _filenameJson+".wav")
		}
	}

	fmt.Fprintf(__vBfIoWriter, "\n")
	__vBfIoWriter.Flush()
}

func main() {
	/*
	   fmt.Println("Hello World")
	   // lnx/goAnalyzeYoutubeJson01.lnx.LinuxX64.exe 11 --> 2
	   // lnx/goAnalyzeYoutubeJson01.lnx.LinuxX64.exe 11 22 --> 3
	   fmt.Printf(" args len %d \n" , len(os.Args))
	   os.Exit(100)
	*/
	if 2 != len(os.Args) {
		fmt.Printf("\n\n  args len %d \n Usage : %s <filename.json>", len(os.Args), os.Args[0])
		os.Exit(100)
	}
	_filenameJson = os.Args[1]

	_readJsonFile()

	_analyzeJsonObj()

	_genYoutubeDownloadScript()
}
