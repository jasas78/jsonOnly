package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

var (
	_P            func(___Vfmt string, ___Vpara ...interface{}) (int, error) = fmt.Printf
	_S            func(___Vfmt string, ___Vpara ...interface{}) string       = fmt.Sprintf
	_Fpf          func(___Vfmt string, ___Vpara ...interface{}) string       = fmt.Sprintf
	_VcmdRemove01 string                                                     = "rm -fr 1/CCC_*"
	_VcmdMkdir01  string                                                     = "mkdir -p 1/"
	__Dir01       string
	__Uri01       string
)

func _Usage() {
	_P("\n\n\n Usage : %s <search_dir> <jpgURIpath> \n \n # will gen the output in the currect dir's child dir ./1/ \n\n", os.Args[0])
	os.Exit(3)
}

func main() {
	__VaLen := len(os.Args)
	_P(" 838192 01 : arg len %d \n", __VaLen)
	if 3 != __VaLen {
		_Usage()
	}

	__Dir01 = os.Args[1]
	__Uri01 = os.Args[2]
	__cmdL01 := "ls " + __Dir01 + "/"
	__Vcmd01 := exec.Command("sh", "-c", __cmdL01)
	// func (c *Cmd) CombinedOutput() ([]byte, error)
	__VstdoutStderr, __Verr2 := __Vcmd01.CombinedOutput()
	if __Verr2 != nil {
		_P(" 838192 03 : error met : <%v>", __Verr2)
		os.Exit(4)
	}

	// func MustCompile(str string) *Regexp
	__Vrep := regexp.MustCompile("\r\n|\r|\n| ")
	//func (re *Regexp) Split(s string, n int) []string
	__VfArr1 := __Vrep.Split(string(__VstdoutStderr), -1)

	__VfLen1 := len(__VfArr1)
	_P(" 838192 05 : get filename Amount : <%d> \n", __VfLen1)

	__VfArr2 := []string{}
	for __Vi1 := 0; __Vi1 < __VfLen1; __Vi1++ {
		//_P(" %d:[%s] ", __Vi1+1, __VfArr1[__Vi1])
		if 0 != len(__VfArr1[__Vi1]) {
			__Vcmd04 := exec.Command("sh", "-c", "basename "+__VfArr1[__Vi1]+" | awk '{printf $1}'")
			__VstdoutStderr5, __Verr5 := __Vcmd04.CombinedOutput()
			if __Verr5 != nil {
				_P(" 838192 06 : error met : <%v>", __Verr5)
				os.Exit(5)
			}

			__VfArr2 = append(__VfArr2, string(__VstdoutStderr5))
		}
	}

	//_P(" 838192 06 : why len : <%d> : %x \n", len(__VfArr1[__VfLen1-1]), __VfArr1[__VfLen1-1])

	__VfLen2 := len(__VfArr2)
	_P("\n 838192 07 : get real Amount %d\n", __VfLen2)

	__Vcmd02 := exec.Command("sh", "-c", _VcmdRemove01)
	__Vcmd02.CombinedOutput()

	__Vcmd03 := exec.Command("sh", "-c", _VcmdMkdir01)
	__Vcmd03.CombinedOutput()

	_Ffprintf("1/_index.md", _S(_Vpg01, "all_jpg", ""))

	_Fjpg_gen__index(&__VfArr2)

	_P("\n\n")
}

func _Fjpg_gen__index(___VfArr *[]string) {
	for __Vidx1, __Vbasename1 := range *___VfArr {
		//for __Vidx1, _ := range *___VfArr {
		//_P(" %s", __Vstr1)
		__Vj2 := __Vidx1 + 1

		__Vj3 := __Vj2 % 10
		__Vj4 := __Vj2 % 100
		__Vj5 := __Vj2 % 1000

		__Vk2 := (__Vj5 - __Vj4) / 100
		__Vk3 := (__Vj4 - __Vj3) / 10
		__Vk4 := __Vj3

		__Vstr2 := _S("%d , %d , %d , %d : %d %d %d", __Vj2, __Vj3, __Vj4, __Vj5, __Vk2, __Vk3, __Vk4)
		__Vstr3 := _S("1/CCC_%d00_%d99", __Vk2, __Vk2)
		__Vstr4 := _S("%d0_%d9", __Vk3, __Vk3)
		__Vstr5 := _S("%s/%s", __Vstr3, __Vstr4)
		__Vstr6 := _S("%s/%03d", __Vstr5, __Vj2)

		__Vmkdir1 := false
		__Vmkdir2 := false
		if __Vj2 == 1 {
			__Vmkdir1 = true
			__Vmkdir2 = true
		} else {
			if 0 == __Vk4 {
				if 0 == __Vk3 {
					__Vmkdir1 = true // __Vstr3
				}
				__Vmkdir2 = true // __Vstr5
			}
		}

		__Vstr7 := ""
		__Vstr8 := ""
		if __Vmkdir1 { // __Vstr3
			__Vstr7 = __Vstr3
		}
		if __Vmkdir2 { // __Vstr5
			__Vstr8 = __Vstr5
		}

		_P("vj2:%d : str2:%s : str3:%s str4:%s str5:%s : str7:%s , str8:%s : str6:%s\n",
			__Vj2, __Vstr2, __Vstr3, __Vstr4, __Vstr5, __Vstr7, __Vstr8, __Vstr6)

		if __Vmkdir1 { // __Vstr3
			__Vcmd21 := exec.Command("sh", "-c", "mkdir -p "+__Vstr7)
			__Vcmd21.CombinedOutput()
			_Ffprintf(__Vstr7+"/_index.md", _S(_Vpg01, __Vstr7, ""))
		}

		if __Vmkdir2 { // __Vstr5
			__Vcmd21 := exec.Command("sh", "-c", "mkdir -p "+__Vstr8)
			__Vcmd21.CombinedOutput()
			_Ffprintf(__Vstr8+"/_index.md", _S(_Vpg01, __Vstr8, ""))
		}

		_Ffprintf(_S(__Vstr5+"/CCD_%05d.md", __Vj2),
			_S(_Vpg02, __Vstr6, "", __Uri01, __Vbasename1))

	}
}

var _Vpg01 string = "+++\ntitle = \"%s\"\ndescription = \"\"\nweight = 20\n+++\n%s\n"
var _Vpg02 string = _Vpg01 + "<img class=\"center-fit-jpg\" src=\"%s/%s\" ></img>\n\n"

func _Ffprintf(___Vfname string, ___VoutStr string) {
	__Vf, __Verr := os.OpenFile(___Vfname, os.O_RDWR|os.O_CREATE, 0755)
	if __Verr != nil {
		_P(" 838191 01 : error <%v>\n\n", __Verr)
		os.Exit(5)
	}

	// func (f *File) WriteString(s string) (n int, err error)
	_, __Verr = __Vf.WriteString(___VoutStr)
	if __Verr != nil {
		_P(" 838191 03 : error <%v>\n\n", __Verr)
		os.Exit(6)
	}

	if __Verr := __Vf.Close(); __Verr != nil {
		_P(" 838191 05 : error <%v>\n\n", __Verr)
		os.Exit(7)
	}
}
