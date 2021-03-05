package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"unicode/utf8"
	"unsafe"
)

func main() {

	str := "{\"values\": [60,3241234,29,266122,72754,11394,0.16,0.869999999999999,58840,20,0,44,395511,34.7122169562927,0.878262153482706,170736,27091,947298,34.9672584991325,0,33,2,0,1,2,1,30,1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,1,0,60,3241234,29,266122,72754,11394,0.16,0.869999999999999,58840,20,0,44,395511,34.7122169562927,0.878262153482706,170736,27091,947298,34.9672584991325,0,33,2,0,1,2,1,30,1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,1,0],\"rows\": 2,\"cols\": 147,\"essayIds\":[\"1111\",\"2222\"]}"
	fmt.Println("原始字符串 ======= ", str)

	arr := []byte(str)
	fmt.Println("解压前字节数组 ======= ", arr)

	// zip压缩
	zipBytes := ZipBytes(arr)
	fmt.Println("zip压缩后的字节数组 ======= ", zipBytes)
	fmt.Println("zip压缩后的字符串长度 ======= ", utf8.RuneCountInString(string(zipBytes)))
	fmt.Println("zip压缩后的字符串666 ======= ", base64.StdEncoding.EncodeToString(zipBytes))

	// zip解压
	uZipBytes := UZipBytes(zipBytes)
	fmt.Println("zip解压字节数组 ======= ", uZipBytes)
	fmt.Println("zip解压后字符串 ======= ", string(uZipBytes))

	// Gzip压缩
	gzipBytes := GZipBytes(arr)
	fmt.Println("Gzip压缩字节数组 ======= ", gzipBytes)
	fmt.Println("Gzip压缩后的字符串长度 ======= ", utf8.RuneCountInString(string(gzipBytes)))
	fmt.Println("Gzip压缩后的字符串666 ======= ", base64.StdEncoding.EncodeToString(gzipBytes))

	// Gzip解压
	uGZipBytes := UGZipBytes(gzipBytes)
	fmt.Println("Gzip解压字节数组 ======= ", uGZipBytes)
	fmt.Println("Gzip解压后字符串 ======= ", string(uGZipBytes))

}

//GZip压缩
func GZipBytes(data []byte) []byte {
	var input bytes.Buffer
	g := gzip.NewWriter(&input) //面向api编程调用压缩算法的一个api
	//参数就是指向某个数据缓冲区默认压缩等级是DefaultCompression 在这里还有另一个api可以调用调整压缩级别
	//gzip.NewWirterLevel(&in,gzip.BestCompression) NoCompression（对应的int 0）、
	//BestSpeed（1）、DefaultCompression（-1）、HuffmanOnly（-2）BestCompression（9）这几个级别也可以
	//这样写gzip.NewWirterLevel(&in,0)
	//这里的异常返回最好还是处理下，我这里纯属省事
	_, _ = g.Write(data)
	_ = g.Close()
	return input.Bytes()
}

//GZip解压
func UGZipBytes(data []byte) []byte {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(data)
	r, _ := gzip.NewReader(&in)
	_ = r.Close() //这句放在后面也没有问题，不写也没有任何报错
	//机翻注释：关闭关闭读者。它不会关闭底层的io.Reader。为了验证GZIP校验和，读取器必须完全使用，直到io.EOF。

	_, _ = io.Copy(&out, r) //这里我看了下源码不是太明白，
	//我个人想法是这样的，Reader本身就是go中表示一个压缩文件的形式，r转化为[]byte就是一个符合压缩文件协议的压缩文件
	return out.Bytes()
}

//zip压缩
func ZipBytes(data []byte) []byte {

	var in bytes.Buffer
	z := zlib.NewWriter(&in)
	_, _ = z.Write(data)
	_ = z.Close()
	return in.Bytes()
}

//zip解压
func UZipBytes(data []byte) []byte {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(data)
	r, _ := zlib.NewReader(&in)
	_ = r.Close()
	_, _ = io.Copy(&out, r)
	return out.Bytes()
}

// string 转 []byte
func Str2sbyte1(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) +
		2*unsafe.Sizeof(&b))) = len(s)
	return
}

// []byte 转 string

func Sbyte2str1(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
