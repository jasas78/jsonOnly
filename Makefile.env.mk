

CFGmakeRun:=Makefile.run.go01.mk


GoTOP:=go2019_0208_1414pm__udp_hold
GoTOP:= \
	goAnalyzeYoutubeJson01 \
	goHtml_server \
	goHtml_cgi01 \
	udp_7001__json_generator \
	jpg_gen_index \

#	udp_0001_test_udp \
#	udp_0997__server  udp_0998__client udp_0888__testOnly

GoTOP:=$(strip $(GoTOP))

udp_base__group01:=    \
	base_1101__gen_md5_sha								\
	base_1102__sleep									\
    base_1103__checkError                               \
    base_1104__checkInterface                           \
	base_1106__hex										\
	base_1111__http_get									\
    \
    base_1121__Sprintf									\
	base_1122__Fprintf									\
	base_1123__Cprintf									\
	base_1125__get_funcName								\
	base_1126__printf_hex								\
	base_1204__udp_listen								\
	base_1210__tcpNode__data_const						\
    \
	base_1246__rand_gen_append							\
	base_1431__data_pack_base_struct					\
    \
	base_1800__FnDnSnCn_const_struct					\
	base_1801__srvInfo_struct							\
	base_1802__loginReq_struct							\
	base_1803__dataTran_struct							\
	base_1804__cmdMap_const								\
    \
	base_1981__read_write_file							\
	base_1990__encode_decode__test_dataLength			\
	base_1991__encode_decode__aes_cbc					\
	base_1992__encode_decode__rand_only					\
	base_1993__encode_decode__rand_aes					\
	base_1996__json_read_write							\
    base_1997__encode_decode__json                      \
    base_1998__encode_decode__gob                       \
    base_1999__encode_decode__bin                       \


udp_0001_test_udp:= udp_0001_test_udp
udp_0997__server:= udp_0997__server  udp_0901__type
udp_0998__client:= udp_0998__client  udp_0901__type
udp_0888__testOnly:=   $(udp_base__group01)         \
    \
    udp_0888__testOnly
jpg_gen_index:=   jpg_gen_index         


goAnalyzeYoutubeJson01 := goAnalyzeYoutubeJson01 
goHtml_server:= 	goHtml_server
goHtml_cgi01:= 	goHtml_cgi01

udp_7001__json_generator:= 	$(udp_base__group01)		\
    \
	udp_1899__FnDnSnCn_const_passwd						\
    \
	udp_7001__json_generator

rj1:=vim_json_generator
$(rj1):= make vg DST=src7/udp_7001__json_generator.go 

rj2:=buil_json_generator_only
$(rj2):=make bl 

rj3:=run_json_generator_only
$(rj3):=./lnx/udp_7001__json_generator.lnx.LinuxX64.exe

rj9:=run_json_generator_all
$(rj9):=make rj1 rj2  rj3

rt:rj9

gr:=git_recover_rand
gr:
	rm -f        json/*.json.rand json/*.gob.rand
	git checkout json/*.json.rand json/*.gob.rand


rg1:=vim__jpg_gen_idx
$(rg1):= make vg DST=src8/jpg_gen_index.go 

rg2 : rj2

rg3:=run_jpg_gen_index
$(rg3):=./lnx/jpg_gen_index.lnx.LinuxX64.exe  ./jpg_  /jpg_

rg9:=run_jpg_gen_index__all
$(rg9):=make rg1 rj2  rg3

testList:= rj1 rj2 rj3 rj9  rg1 rg3 rg9
$(testList): 
	$($($@))



showRunHelpListLast += $(testList)

GoPreLinuxALL:= LinuxX64
GoPreDockerALL:=LinuxX64

bt: btgo


xx1OBJ1?=NA_1159363978952183809.info.json
xx1OBJ2?=11.json
xx1:
	@echo
	make vg125
	@echo
	make bl
	@echo
	cat        lnx/goAnalyzeYoutubeJson01.lnx.LinuxX64.exe   \
		> objSaved/goAnalyzeYoutubeJson01.lnx.LinuxX64.exe
	@echo
	cd 1/ && ../lnx/goAnalyzeYoutubeJson01.lnx.LinuxX64.exe   $(xx1OBJ1)
	cd 1/ && ../lnx/goAnalyzeYoutubeJson01.lnx.LinuxX64.exe   $(xx1OBJ2)
	@echo
	ls -l                                         1/$(xx1OBJ1)* 1/$(xx1OBJ2)* 
	@echo


