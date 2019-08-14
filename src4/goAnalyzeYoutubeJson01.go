package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	//"io"
	//"log"
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
	filesize, width                               int
	format_id, ext, vcodec, acodec, url, protocol string
}

type _STdst struct {
	vo1_ao2_both3 int
	size          int
	width         int
	idx           int
	url           string
	protocol      string
	ext           string
}

var (
	_jsonFile         *os.File
	_filenameJson     string = "________________.json"
	_filenameBase     string = "________________"
	_filenameDir      string = "________________.dir"
	_filenamePure     string
	_jsonByte         []byte
	_err              error
	_vstYT00          map[string]interface{}
	_vstYT10          []string = []string{"description", "id", "webpage_url", "thumbnail", "fulltitle", "format_id", "upload_date", "uploader", "protocol"}
	_recLen           int
	_recArr           []_STrec
	_s200             string = ""
	_s300             string = ""
	_s400             string = ""
	_vDstMaxAllowVo   _STdst
	_vDstMaxAllowAo   _STdst
	_vDstMaxAllowBoth _STdst
	_vFnameMpX        string
	_vFnameVoX        string
	_vDescription01   string
	_vForceLargeMulti int = 1
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

	//fmt.Printf(" JSON description 00 <%v> \n" , _vstYT00["description"] )
	_vDescription01 =
		strings.Replace(strings.Replace(
			strings.Replace(strings.Replace(
				fmt.Sprintf("%s", _vstYT00["description"]),
				"\n", "_", -1),
				"\r", "_", -1),
			"'", "_", -1),
			"\"", "_", -1)
	for ___idx, ___key := range _vstYT10 {
		if ___key == "description" {
			_s200 += fmt.Sprintf("#%d= %s:=<%v> \n", 10000+___idx, ___key, _vDescription01)
		} else {
			_s200 += fmt.Sprintf("#%d= %s:=<%v> \n", 10000+___idx, ___key, _vstYT00[___key])
		}
	}
	if 123 == 123 {
		fmt.Printf("%s", _s200)
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
			_s300 += fmt.Sprintf("#%d= ", 30000+___idx)
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
				case "width":
					{
						switch vv := ___value.(type) {
						case float64:
							_s300 += fmt.Sprintf(" <%s:%v>", ___key, int(vv))
							_recArr[___idx].width = int(vv)

						}
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
			_s400 += fmt.Sprintf("#%d/%d=", 40000+___idx, _recLen)
			_s400 += fmt.Sprintf(" %3s", _recArr[___idx].format_id)
			_s400 += fmt.Sprintf(" %9d", _recArr[___idx].filesize)
			_s400 += fmt.Sprintf(" %5d", _recArr[___idx].width)
			_s400 += fmt.Sprintf(" %4s", _recArr[___idx].ext)
			_s400 += fmt.Sprintf(" %13s", _recArr[___idx].vcodec)
			_s400 += fmt.Sprintf(" %13s", _recArr[___idx].acodec)
			_s400 += fmt.Sprintf(" '%19s'", _recArr[___idx].protocol)
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
						if __fSize < (50000000*_vForceLargeMulti) &&
							(__fSize > _vDstMaxAllowBoth.size || _vDstMaxAllowBoth.vo1_ao2_both3 == 0) {
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
						if __fSize < (45000000*_vForceLargeMulti) &&
							(__fSize > _vDstMaxAllowAo.size || _vDstMaxAllowAo.vo1_ao2_both3 == 0) {
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
						if __fSize < (35000000*_vForceLargeMulti) &&
							(__fSize > _vDstMaxAllowVo.size || _vDstMaxAllowVo.vo1_ao2_both3 == 0) {
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
						fmt.Printf(" 100053 : not-vo, not-ao , what is this ? idx %d \n", ___idx)
						fmt.Printf(" 100054 : \n %s\n\n", _s400)
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

func _FgenDstCombineLine(___w *bufio.Writer, ___srcVo, ___srcAo, ___dstCo string) {
	__tmpExt1 := strings.TrimSuffix(___dstCo, ".webm")
	if __tmpExt1 == ___dstCo {
		_vFnameVoX = ___dstCo
	} else {
		_vFnameVoX = ___dstCo + ".mp4"
	}

	fmt.Fprintf(___w, "rm -f %s \n", _vFnameVoX)
	fmt.Fprintf(___w, "/usr/bin/ffmpeg "+
		"\\\n    -i %s  "+
		"\\\n    -i %s        "+
		"\\\n    -map 0:v:0 -map 1:a:0      "+
		"\\\n    -ac 1 -ar 22050 -b:a 25000   "+
		"\\\n    -c:v copy   "+
		"\\\n    %s "+
		"\\\n    && (echo 'ok     :gen %s' >> result.txt )"+
		"\\\n    || (echo 'failed :gen %s' >> result.txt )"+
		"\n",
		___srcVo, ___srcAo, _vFnameVoX, _vFnameVoX, _vFnameVoX)
	// https://superuser.com/questions/1137612/ffmpeg-replace-audio-in-video
}

func _FgenWavMp3Line(___w *bufio.Writer, ___src string) {
	__vFnameWav := _filenameJson + ".wav"
	__vFnameMp3 := _filenameJson + ".mp3"
	_vFnameMpX = _filenameJson + ".25k.mp4"
	fmt.Fprintf(___w, "rm -f %s \n", __vFnameWav)
	fmt.Fprintf(___w,
		"/usr/bin/ffmpeg \\\n    -i %s         \\\n    -vn -ac 1 -ar 22050 -b:a 25000        "+
			"\\\n    %s \n",
		___src, __vFnameWav)
	fmt.Fprintf(___w, "rm -f %s \n", __vFnameMp3)
	fmt.Fprintf(___w, "rm -f %s \n", _vFnameMpX)
	fmt.Fprintf(___w, "echo echo lame       \\\n    %s       \\\n    %s\n", __vFnameWav, __vFnameMp3)
	fmt.Fprintf(___w,
		"/usr/bin/ffmpeg \\\n    -i %s         \\\n    -vn -ac 1 -ar 22050 -b:a 25000        "+
			"\\\n    %s \n",
		___src, _vFnameMpX)
}

var _vYTcmd string = "${HOME}/bin/youtube-dl --ignore-errors --restrict-filenames  "

// -o '%%(upload_date)s_%(id)s.%%(ext)s'"

//func _FgetDownloadLine(___w *bufio.Writer, ___dst, ___src, ___protocol string) {
func _FgetDownloadLine(___w *bufio.Writer, ___dst string, ___vRec _STrec) {
	__protocol := ___vRec.protocol
	__src1 := ___vRec.url
	__src2 := _vstYT00["webpage_url"]
	__fmt := ___vRec.format_id
	switch __protocol {
	case "https", "http_dash_segments":
		{
			fmt.Fprintf(___w, "rm -f %s \n", ___dst)
			fmt.Fprintf(___w, "echo echo wget -c \\\n    -O %s  \\\n    '%s'\n", ___dst, __src1)
			fmt.Fprintf(___w, _vYTcmd+" \\\n    -f %s \\\n    -o %s  \\\n    '%s'\n",
				__fmt,
				___dst, __src2)
		}
	case "m3u8_native":
		{
			fmt.Fprintf(___w, "rm -f %s \n", ___dst)
			fmt.Fprintf(___w, "echo echo /usr/bin/ffmpeg -i '%s' \\\n    %s \n", __src1, ___dst)
			fmt.Fprintf(___w, _vYTcmd+" \\\n    -f %s \\\n    -o %s  \\\n    '%s'\n",
				__fmt,
				___dst, __src2)
		}
	default:
		{
			fmt.Printf("\n\n unknown protocol :<%s>\n\n", __protocol)
			os.Exit(179)
		}

	}
}

var _vShugoIndex01 string = `

mv  %s/%s %s/%s

cat > %s/_index.md << EOF3
+++
title = " %s %s "
description = " %s "
weight = 20
+++

{{< mymp4 mp4="%s" 
text="len $(cat %s/%s|wc -c)"
>}}

{{< mymp4x  mp4x="%s"
text="len $(cat %s/%s|wc -c)"
>}}


{{< mydiv text="%s" >}}
<br>

{{< mydiv text="%s" >}}


<br>

请大家传播时，不需要传播文件本身，<br>
原因是：一旦传播过大东西（例如，图片，文件），<br>
就会触发检查机制。<br>
我不知道检查机制的触发条件。<br>
但是我知道，不会说你传一个没有敏感词的网络地址都检查，<br>
否则，检查员得累死。<br><br>
直接转发网址就可以了：<br>
原因是，这是程序员网站，<br>
共匪不敢封锁，墙内可以直接下载。


EOF3

`

func _genIndexScripForHugo2() {
	__vFshName := _filenameJson + ".sh2"
	__vFileSh, __vErr := os.Create(__vFshName)
	if __vErr != nil {
		fmt.Printf("Create <%s> failed\n", __vFshName)
		fmt.Println(__vErr)
		os.Exit(140)
	}
	defer __vFileSh.Close()

	__vBfIoWriter := bufio.NewWriter(__vFileSh)
	fmt.Fprintf(__vBfIoWriter, "#!/bin/bash\n\n")

	fmt.Fprintf(__vBfIoWriter, "rm -f \\\n    %s.vo.* \\\n    %s.ao.* \\\n    %s.bo.* \\\n    %s.wav \n\n",
		_filenameJson,
		_filenameJson,
		_filenameJson,
		_filenameJson)

	fmt.Fprintf(__vBfIoWriter, "rm -fr               %s\nmkdir -p             %s\nmv %s %s.*         %s/\n\n",
		_filenameDir,
		_filenameDir,
		_filenameJson,
		_filenameJson,
		_filenameDir)

	fmt.Fprintf(__vBfIoWriter, _vShugoIndex01,
		_filenameDir,
		_filenameJson+".jpg",
		_filenameDir,
		_vFnameVoX+".jpg",

		_filenameDir,
		_filenamePure,
		_vstYT00["fulltitle"],
		_vDescription01,

		_vFnameVoX,
		_filenameDir,
		_vFnameVoX,

		_vFnameMpX,
		_filenameDir,
		_vFnameMpX,

		_vDescription01,

		_vstYT00["webpage_url"])

	fmt.Fprintf(__vBfIoWriter, "\n")
	__vBfIoWriter.Flush()
} // _genIndexScripForHugo2

// func fmt.Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// func os.Create(name string) (*File, error)
// func bufio.NewWriter(w io.Writer) *Writer
func _genYoutubeDownloadScript1() {
	__vFshName := _filenameJson + ".sh1"
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

	fmt.Fprintf(__vBfIoWriter, "wget -c -O %s.jpg \\\n    '%s' \n\n", _filenameJson, _vstYT00["thumbnail"])

	//fmt.Fprintf(__vBfIoWriter , "wget -c -O %s.jpg\n\n", _filenameJson )

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

	__ffVO := fmt.Sprintf("%s.vo.%s", _filenameJson, _vDstMaxAllowVo.ext)
	__ffAO := fmt.Sprintf("%s.ao.%s", _filenameJson, _vDstMaxAllowAo.ext)
	__ffBO := fmt.Sprintf("%s.bo.%s", _filenameJson, _vDstMaxAllowBoth.ext)
	__ffDstVo := fmt.Sprintf("%s.%s", _filenameJson, _vDstMaxAllowVo.ext)
	__ffDstBo := fmt.Sprintf("%s.%s", _filenameJson, _vDstMaxAllowBoth.ext)

	if _vDstMaxAllowAo.vo1_ao2_both3 != 0 &&
		_vDstMaxAllowVo.vo1_ao2_both3 != 0 {
		fmt.Fprintf(__vBfIoWriter, "# combine183838 01 : ideal : use vo + ao : %s + %s \n", __ffVO, __ffAO)
		_FgetDownloadLine(__vBfIoWriter, __ffVO, _recArr[_vDstMaxAllowVo.idx])
		_FgetDownloadLine(__vBfIoWriter, __ffAO, _recArr[_vDstMaxAllowAo.idx])

		fmt.Fprintf(__vBfIoWriter, "\n")
		_FgenWavMp3Line(__vBfIoWriter, __ffAO)
		_FgenDstCombineLine(__vBfIoWriter, __ffVO, _filenameJson+".wav", __ffDstVo)
	} else { // no-Vo or/and no-Ao
		if _vDstMaxAllowBoth.vo1_ao2_both3 == 0 { // no-both + ( no-Vo or no-Ao)
			fmt.Printf("\n\n# no-both , AND (no-Ao || no-Vo), what happens ? 1838111. 181 \n\n")
			os.Exit(181)
		}
		if _vDstMaxAllowVo.vo1_ao2_both3 != 0 { // use vo && both(as ao)
			fmt.Fprintf(__vBfIoWriter, "# combine183838 11 : use vo + both(ao) : %s + %s \n", __ffVO, __ffBO)
			_FgetDownloadLine(__vBfIoWriter, __ffVO, _recArr[_vDstMaxAllowVo.idx])
			_FgetDownloadLine(__vBfIoWriter, __ffBO, _recArr[_vDstMaxAllowBoth.idx])

			fmt.Fprintf(__vBfIoWriter, "\n")
			_FgenWavMp3Line(__vBfIoWriter, __ffBO)
			_FgenDstCombineLine(__vBfIoWriter, __ffVO, _filenameJson+".wav", __ffDstVo)
		} else {
			if _vDstMaxAllowAo.vo1_ao2_both3 != 0 { // use both(as vo) + ao
				fmt.Fprintf(__vBfIoWriter, "# combine183838 21 : use both(vo) + ao : %s + %s \n", __ffBO, __ffAO)
				_FgetDownloadLine(__vBfIoWriter, __ffBO, _recArr[_vDstMaxAllowBoth.idx])
				_FgetDownloadLine(__vBfIoWriter, __ffAO, _recArr[_vDstMaxAllowAo.idx])

				fmt.Fprintf(__vBfIoWriter, "\n")
				_FgenWavMp3Line(__vBfIoWriter, __ffAO)
				_FgenDstCombineLine(__vBfIoWriter, __ffBO, _filenameJson+".wav", __ffDstBo)
			} else { // use both only
				fmt.Fprintf(__vBfIoWriter, "# combine183838 31 : no vo , but both only, so , use both only : %s \n", __ffBO)
				_FgetDownloadLine(__vBfIoWriter, __ffBO, _recArr[_vDstMaxAllowBoth.idx])

				fmt.Fprintf(__vBfIoWriter, "\n")
				_FgenWavMp3Line(__vBfIoWriter, __ffBO)
				_FgenDstCombineLine(__vBfIoWriter, __ffBO, _filenameJson+".wav", __ffDstBo)
			}
		}
	}

	if 20004 == 20004 {
		fmt.Fprintf(__vBfIoWriter, "\n")
		fmt.Fprintf(__vBfIoWriter, "%s", _s200)
	}

	if 30004 == 30004 {
		fmt.Fprintf(__vBfIoWriter, "\n")
		fmt.Fprintf(__vBfIoWriter, "%s", _s300)
	}

	fmt.Fprintf(__vBfIoWriter, "\n")
	__vBfIoWriter.Flush()
} // _genYoutubeDownloadScript1

func main() {
	/*
	   fmt.Println("Hello World")
	   // lnx/goAnalyzeYoutubeJson01.lnx.LinuxX64.exe 11 --> 2
	   // lnx/goAnalyzeYoutubeJson01.lnx.LinuxX64.exe 11 22 --> 3
	   fmt.Printf(" args len %d \n" , len(os.Args))
	   os.Exit(100)
	*/
	if 2 == len(os.Args) {
		_filenameJson = os.Args[1]
	} else {
		if 3 == len(os.Args) && os.Args[1] == "-f" {
			_filenameJson = os.Args[2]
			_vForceLargeMulti = 2
		} else {
			fmt.Printf("\n\n  args len %d \n Usage : %s <filename.json>\n\n"+
				"    for aa1 in *.json ; do %s ${aa1} ;"+
				"echo . ./${aa1}.sh1 >> z1.txt;echo . ./${aa1}.sh2 >> z2.txt; done\n\n",
				len(os.Args), os.Args[0], os.Args[0])
			os.Exit(100)
		}
	}
	_filenameBase = strings.TrimSuffix(_filenameJson, ".json")
	_filenameDir = _filenameBase + "_dir"
	_filenamePure = strings.TrimSuffix(_filenameBase, ".info")

	_readJsonFile()

	_analyzeJsonObj()

	_genYoutubeDownloadScript1()

	_genIndexScripForHugo2()
}
