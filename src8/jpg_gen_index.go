package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

var (
	_P func(___Vfmt string, ___Vpara ...interface{}) (int, error) = fmt.Printf
	_S func(___Vfmt string, ___Vpara ...interface{}) string       = fmt.Sprintf
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

}
