package util

import "unsafe"

// []byte 与 string 高效转化
func main() {

}

// string 转 []byte
func Str2sbyte(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) +
		2*unsafe.Sizeof(&b))) = len(s)
	return
}

// []byte 转 string

func Sbyte2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
