package main

import (
//"net"
)

type _Tconfig struct {
	//Name    string
	//MyId128 []byte
} //    _Tconfig
type _Tself struct {
	ProjName string
	//
	//     progPath string
	//     progMd5  _Tb128 // md5sum    : 16 byte : 727bf338cf523b90baccd24cca30b919
	//     progSha  _Tb256 // sha256sum : 32 byte : 2c6e3b458d5c482bc52a1d7d4f5a7d7766381c9f07d9b32ca605ae45b4e473f5
	//
	//     startTime    time.Time
	//     startTimEsha _Tb256
	//
	//     debugEnabled bool
} //    _Tself

type _TuExtMRead struct {
	//name string
} // _TuExtMRead

type _TuExtChanI struct {
} // _TuExtChanI

type _TuExtTimer struct {
} // _TuExtTimer

var (
	_VconfigFnWaitDn_json _TsrvInfo = _TsrvInfo{
		name:       "FnWaitDn",
		refreshUri: "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/FnWaitDn.json.rand",
		UriArrs:    []string{"127.0.0.1:32001", "127.0.0.1:32002"},
		//UdstAddr: []net.UDPAddr{
		//	net.UDPAddr{net.IPv4(127, 0, 0, 1), 32001, ""},
		//	net.UDPAddr{net.IPv4(127, 0, 0, 1), 32002, ""}},
		K256: [][]byte{_Vpasswd_udp_Fn_waitForCliens01, _Vpasswd_udp_Fn_waitForCliens02}}

	_VconfigDn2Cn_json _TsrvInfo = _TsrvInfo{
		name:       "Dn2Cn",
		refreshUri: "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/Dn2Cn.json.rand",
		UriArrs:    []string{"127.0.0.1:31001", "127.0.0.1:31002", "127.0.0.1:31003"},
		//UdstAddr: []net.UDPAddr{
		//	net.UDPAddr{net.IPv4(127, 0, 0, 1), 31001, ""},
		//	net.UDPAddr{net.IPv4(127, 0, 0, 1), 31002, ""},
		//	net.UDPAddr{net.IPv4(127, 0, 0, 1), 31003, ""}},
		K256: [][]byte{_Vpasswd_udp_Cn_waitForCliens01, _Vpasswd_udp_Cn_waitForCliens02}}

	_VconfigFnWaitDn_gob _TsrvInfo = _VconfigFnWaitDn_json
	_VconfigDn2Cn_gob    _TsrvInfo = _VconfigDn2Cn_json

	_Vconfig _Tconfig
	_Vself   _Tself
)

func init() {
	_VconfigFnWaitDn_gob.refreshUri = "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/FnWaitDn.gob.rand"
	_VconfigDn2Cn_gob.refreshUri = "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/Dn2Cn.gob.rand"
}

func _Ftry_download_rand_json_and_show_the_reult(___Vmsg string, ___KeyF *[]byte, ___VsrvInfo *_TsrvInfo, ___VrecoverObj *_TsrvInfo) {

	___VdownUri := ___VsrvInfo.refreshUri
	//___Vkey := ___VsrvInfo.K256

	//_, __Verr := _Ftry_download_rand_json01(___VdownUri, &___Vkey, ___VrecoverObj)
	_, __Verr := _Ftry_download_rand_json01(___VdownUri, ___KeyF, ___VrecoverObj)
	if nil == __Verr {
		_FpfN(___Vmsg+" 893481 01 : ok : %s , %v ", ___VdownUri, ___VrecoverObj)
	} else {
		_FpfN(___Vmsg+" 893481 02 : Error : %s , %v ", ___VdownUri, __Verr)
	}

} // _Ftry_download_rand_json_and_show_the_reult

func _Ftry_download_rand_gob_and_show_the_reult(___Vmsg string, ___KeyF *[]byte, ___VsrvInfo *_TsrvInfo, ___VrecoverObj *_TsrvInfo) {

	___VdownUri := ___VsrvInfo.refreshUri

	//_, __Verr := _Ftry_download_rand_json01(___VdownUri, ___KeyF, ___VrecoverObj)
	_, __Verr := _Ftry_download_rand_gob01(___VdownUri, ___KeyF, ___VrecoverObj)
	if nil == __Verr {
		_FpfN(___Vmsg+" 893481 03 : ok : %s , %v ", ___VdownUri, ___VrecoverObj)
	} else {
		_FpfN(___Vmsg+" 893481 04 : Error : %s , %v ", ___VdownUri, __Verr)
	}

} // _Ftry_download_rand_gob_and_show_the_reult

func main() {

	_FpfN("\n Start \n")

	//_FtestLen01()

	//_FtestER__write_json_and_rand_Exit
	_Ftry_gen_json01("FnWaitDn", &_Vpasswd_udp_FnWaitDn_download_config, &_VconfigFnWaitDn_json)
	_Ftry_gen_json01("Dn2Cn", &_Vpasswd_udp_Dn2Cn_download_config, &_VconfigDn2Cn_json)

	// _FtestER__write_gob_and_rand_Exit
	_Ftry_gen_gob01("FnWaitDn", &_Vpasswd_udp_FnWaitDn_download_config, &_VconfigFnWaitDn_gob)
	//_FpfNex(" 838191111111 ")
	_Ftry_gen_gob01("Dn2Cn", &_Vpasswd_udp_Dn2Cn_download_config, &_VconfigDn2Cn_gob)

	var (
		_Vconfig_tmp__Dn2Fn _TsrvInfo
		_Vconfig_tmp__Fn    _TsrvInfo
	)

	_FpfN("")
	_Ftry_download_rand_json_and_show_the_reult(" 839121 01 ", &_Vpasswd_udp_FnWaitDn_download_config, &_VconfigFnWaitDn_json, &_Vconfig_tmp__Fn)
	_Ftry_download_rand_json_and_show_the_reult(" 839121 02 ", &_Vpasswd_udp_Dn2Cn_download_config, &_VconfigDn2Cn_json, &_Vconfig_tmp__Dn2Fn)
	_Ftry_download_rand_json_and_show_the_reult(" 839121 03 ", &_Vpasswd_udp_FnWaitDn_download_config, &_VconfigDn2Cn_json, &_Vconfig_tmp__Dn2Fn)
	_FpfN("")

	_Ftry_download_rand_gob_and_show_the_reult(" 839121 11 ", &_Vpasswd_udp_FnWaitDn_download_config, &_VconfigFnWaitDn_gob, &_Vconfig_tmp__Fn)
	_Ftry_download_rand_gob_and_show_the_reult(" 839121 12 ", &_Vpasswd_udp_Dn2Cn_download_config, &_VconfigDn2Cn_gob, &_Vconfig_tmp__Dn2Fn)
	_Ftry_download_rand_gob_and_show_the_reult(" 839121 13 ", &_Vpasswd_udp_FnWaitDn_download_config, &_VconfigDn2Cn_gob, &_Vconfig_tmp__Dn2Fn)
	_FpfN("")

	// use a fake key to test the err

} // main
