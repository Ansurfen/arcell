package main

/*
typedef struct {
	char** data;
	int len;
} StringSlice;

typedef struct {
	void* ptr;
} _Conn;

typedef struct {
	char* out;
	char* eout;
	unsigned char err;
} GSysRes;
*/
import "C"

import (
	"encoding/json"
	"go_utils/utils"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"unsafe"
)

type Packet struct {
	ACTION string `json:"_ACTION_"`
	PARAM  string `json:"_PARAM_"`
	DATA   string `json:"_DATA_"`
}

func main() {}

//export DialConn
func DialConn(network, addr *C.char) *C._Conn {
	conn, err := net.Dial(C.GoString(network), C.GoString(addr))
	if err != nil {
		log.Fatalf("Dial: %s\n", err)
	}
	ret := (*C._Conn)(C.malloc(C.size_t(unsafe.Sizeof(C._Conn{}))))
	ret.ptr = unsafe.Pointer(&conn)
	return ret
}

//export CloseConn
func CloseConn(conn *C._Conn) C.uchar {
	err := (*(*net.Conn)(conn.ptr)).Close()
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}

//export WriteConn
func WriteConn(conn *C._Conn, data *C.char) C.uchar {
	_, err := (*(*net.Conn)(conn.ptr)).Write([]byte(C.GoString(data)))
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}

//export ReadConn
func ReadConn(conn *C._Conn) *C.char {
	bytes, err := ioutil.ReadAll((*(*net.Conn)(conn.ptr)))
	if err != nil {
		panic(err)
	}
	return C.CString(string(bytes))
}

//export GSystem
func GSystem(name *C.char, args []*C.char) *C.GSysRes {
	var arg []string
	for _, v := range args {
		if C.GoString(v) != "" {
			arg = append(arg, C.GoString(v))
		}
	}
	cmd := exec.Command(C.GoString(name), arg...)
	out, err := cmd.CombinedOutput()
	ret := (*C.GSysRes)(C.malloc(C.size_t(unsafe.Sizeof(C.GSysRes{}))))
	if err != nil {
		ret.err = 1
		ret.eout = C.CString(err.Error())
	} else {
		ret.err = 0
		ret.eout = C.CString("")
	}
	ret.out = C.CString(string(out))
	return ret
}

//export _Replace
func _Replace(s, _old, _new *C.char, n C.int) *C.char {
	return C.CString(strings.Replace(C.GoString(s), C.GoString(_old), C.GoString(_new), int(n)))
}

//export _Split
func _Split(s, sep *C.char) *C.StringSlice {
	res := strings.Split(C.GoString(s), C.GoString(sep))
	return ToCStringSlice(res)
}

func ToCStringSlice(list []string) *C.StringSlice {
	buf := C.malloc(C.size_t(len(list)) * C.size_t(unsafe.Sizeof(uintptr(0))))
	s := unsafe.Slice((**C.char)(buf), len(list))
	for i, item := range list {
		s[i] = C.CString(item)
	}
	ret := (*C.StringSlice)(C.malloc(C.size_t(unsafe.Sizeof(C.StringSlice{}))))
	ret.data = (**C.char)(buf)
	ret.len = (C.int)(len(list))
	return ret
}

//export Len
func Len(v *C.char) C.int {
	return C.int(len(C.GoString(v)))
}

//export _ReplaceWithRegexp
func _ReplaceWithRegexp(expr, src, repl *C.char) *C.char {
	re, err := regexp.Compile(C.GoString(expr))
	if err != nil {
		panic(err)
	}
	return C.CString(re.ReplaceAllString(C.GoString(src), C.GoString(repl)))
}

//export _FindStr
func _FindStr(s, substr *C.char) C.int {
	return C.int(strings.Index(C.GoString(s), C.GoString(substr)))
}

//export _log
func _log(_type *C.char, data *C.char) {
	v := C.GoString(data)
	switch C.GoString(_type) {
	case "INFO":
		log.Println(v)
	case "PANIC":
		log.Panicln(v)
	case "FATAL":
		log.Fatalln(v)
	}
}

//export gout
func gout(str *C.char) {
	println(C.GoString(str))
}

//export GSystemWithStdout
func GSystemWithStdout(name *C.char, args []*C.char, g, l C.uchar, workdir *C.char) {
	var arg []string
	for _, v := range args {
		if C.GoString(v) != "" {
			arg = append(arg, C.GoString(v))
		}
	}
	cmd := exec.Command(C.GoString(name), arg...)
	out, err := cmd.CombinedOutput()
	out = []byte(strings.ReplaceAll(string(out), "\r\n", ""))
	if err != nil {
		println(err)
	}
	param := ""
	if g != 0 {
		param += "global"
	}
	if l != 0 {
		param += " " + C.GoString(workdir)
	}
	p := Packet{
		ACTION: "MSG",
		PARAM:  param,
		DATA:   string(out),
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.Dial("tcp", "127.0.0.1:1314")
	if err != nil {
		log.Fatalf("Dial: %s\n", err)
	}
	_, err = conn.Write(data)
	data, err = ioutil.ReadAll(conn)
	// 直接拿回link,我们自己去读
	f := 1.0
	println(string(data))
	// conf := utils.GetConf("tell", "json", "E:\\arcell\\daemon\\src")
	// println(conf.GetStringSlice("comment")[0])
	utils.SimilarText("", "", &f)
	// 去找守护进程拿数据，然后json操作匹配先全局在局部
	println(string(out))
}
