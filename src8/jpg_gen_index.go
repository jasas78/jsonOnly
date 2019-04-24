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
	_VcmdRemove01 string                                                     = "rm -fr 1/CCC_*"
	_VcmdMkdir01  string                                                     = "mkdir -p 1/"
)

func _Usage() {
	_P("\n\n\n Usage : %s <search_dir> \n \n # will gen the output in the currect dir's child dir ./1/ \n\n", os.Args[0])
	os.Exit(3)
}

func main() {
	__VaLen := len(os.Args)
	_P(" 838192 01 : arg len %d \n", __VaLen)
	if 2 != __VaLen {
		_Usage()
	}

	__Dir01 := os.Args[1]
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
			__VfArr2 = append(__VfArr2, __VfArr1[__Vi1])
		}
	}

	//_P(" 838192 06 : why len : <%d> : %x \n", len(__VfArr1[__VfLen1-1]), __VfArr1[__VfLen1-1])

	__VfLen2 := len(__VfArr2)
	_P("\n 838192 07 : get real Amount %d\n", __VfLen2)

	__Vcmd02 := exec.Command("sh", "-c", _VcmdRemove01)
	__Vcmd02.CombinedOutput()

	__Vcmd03 := exec.Command("sh", "-c", _VcmdMkdir01)
	__Vcmd03.CombinedOutput()

	_Fjpg_gen__index(&__VfArr2)

	_P("\n\n")
}

func _Fjpg_gen__index(___VfArr *[]string) {
	//for __Vidx1, __Vstr1 := range *___VfArr {
	for __Vidx1, _ := range *___VfArr {
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
		if __Vmkdir2 { // __Vstr3
			__Vstr8 = __Vstr5
		}

		_P("%d : %s : %s %s : %s %s \n", __Vj2, __Vstr2, __Vstr3, __Vstr4, __Vstr7, __Vstr8)

	}
}
