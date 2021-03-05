package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const orginalStr = "{\"values\": [60,3241234,29,266122,72754,11394,0.16,0.869999999999999,58840,20,0,44,395511,34.7122169562927,0.878262153482706,170736,27091,947298,34.9672584991325,0,33,2,0,1,2,1,30,1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,1,0,60,3241234,29,266122,72754,11394,0.16,0.869999999999999,58840,20,0,44,395511,34.7122169562927,0.878262153482706,170736,27091,947298,34.9672584991325,0,33,2,0,1,2,1,30,1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,1,0],\"rows\": 2,\"cols\": 147,\"essayIds\":[\"1111\",\"2222\"]}"

// zip 压缩字符串
const nbody = "eJzsU0tKs0EQvMrPrIswVd3TPZ0b/GcIWQTNLhDwQ0XEu8uHCBHFx86FM8x0beoBRT+2u8Pp9ri07b9ddJicMocKiqCEVA4HaeXoGwb6ZkZdHow5vUMdHe6wGoOE+SYpMWqESrnycirEYT6VPcDsaQFlL6I8VXOlVaTG9CqaBjrMIHQQAmEreHmv//cuL+bHPP5I7ys3vvHluxyfJ7lU+avlN9ayR7s53697I7Sr82lF9EQ7Lsvh4f/10ra7RpINTZLa/uk5AAD//059qkg="

// gzip 压缩字符串
//const gZipStr = "H4sIAAAAAAAAAO1TzUqDMRB8Fcl5KNmf7GZ9A5+h9FC0t0LBDxWRvrsbREihpXrroRsyOznMzMKSr/K+3b/tlvL4sLYKYSUWBQfYjJjh7E1BJKGoK7KEbjEXWu9awRUVqpBojQiiK089WTTjYB8672xMTbSzVwN5dTEkD0Koc/QhC3NuXSNIuKWnCDgbJabtID/3F/92aOrndfQvv2tpc96ceMouTTK73Ndyi2vZoLwePsa/YZTnw34wUkfZLcv28+kl3+tCWQWFs8rm+A1ATm0cbAMAAA=="
const gZipStr = "H4sIAAAAAAAAAO1TQUoEQQy8+4w+F0unkk46/sA3yB4W3dvCgoOKiH83gwgjKOrNg2m6UpeqCoQ8t4fD6f64tMtr71CaUA1M0F1IBGMYRDQNfSdeMD23hTGndbCjwwyaY4hAbRelF8/hTMaqi0mnDLXJ6A6JHuoonoK0YM5Vlh4c0zJFOcpTFawmhWW7krf/jj97sumf6+RXft+lbfO2iR/ZV5NsXf7X8hfXske7Oz/W2RDt5nwqIhZox2U5PF3drufUpKqhsartXy5eAX/piaRqAwAA"

func main() {

	s, gzipErr := io.Copy(os.Stdout, GZipDecoder(strings.NewReader(gZipStr)))
	fmt.Println("Gzip解压后的字符串=======", s)
	if gzipErr != nil {
		log.Fatalf("Error copying decoded value to stdout: %s", gzipErr)
	}

	_, err := io.Copy(os.Stdout, ZipDecoder(strings.NewReader(nbody)))
	if err != nil {
		log.Fatalf("Error copying decoded value to stdout: %s", err)
	}

	z, _ := base64.StdEncoding.DecodeString(nbody)
	r, _ := zlib.NewReader(bytes.NewReader(z))
	result, _ := ioutil.ReadAll(r)
	fmt.Println("\n===================", string(result))

	f, _ := base64.StdEncoding.DecodeString(gZipStr)
	g, _ := gzip.NewReader(bytes.NewReader(f))
	gr, _ := ioutil.ReadAll(g)
	fmt.Println("\n===================", string(gr))
}

func GZipDecoder(r io.Reader) io.Reader {

	// We simply set up a custom chain of Decoders
	d, err := gzip.NewReader(
		base64.NewDecoder(base64.StdEncoding, r))

	// This should only occur if one of the Decoders can not reset
	// its internal buffer. Hence, it validates a panic.
	if err != nil {
		panic(fmt.Sprintf("Error setting up decoder chain: %s", err))
	}

	// We return an io.Reader which can be used as any other
	return d

}

func ZipDecoder(r io.Reader) io.Reader {

	// We simply set up a custom chain of Decoders
	d, err := zlib.NewReader(
		base64.NewDecoder(base64.StdEncoding, r))

	// This should only occur if one of the Decoders can not reset
	// its internal buffer. Hence, it validates a panic.
	if err != nil {
		panic(fmt.Sprintf("Error setting up decoder chain: %s", err))
	}

	// We return an io.Reader which can be used as any other
	return d

}
