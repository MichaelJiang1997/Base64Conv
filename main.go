/*
Base64 Converter
Authour: Michael Jiang
Last-Modified: Thu May  9 23:43:25 DST 2019
*/

package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
)

//F2B filr to base64 code
func F2B(fName string) []byte {
	pCode, err := ioutil.ReadFile(fName)
	if err != nil {
		fmt.Println("error:", err)
	}
	var encoded bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &encoded)
	encoder.Write(pCode)
	encoder.Close()
	return encoded.Bytes()
}

//B2F base64 code to file
func B2F(fName string, bCode []byte) {
	var buffer bytes.Buffer
	var decoded bytes.Buffer
	temp := make([]byte, 4096)
	buffer.Write(bCode)
	decoder := base64.NewDecoder(base64.StdEncoding, &buffer)

	//this fucking loop make me sick
	for {
		n, _ := decoder.Read(temp)
		for i := 0; i < n; i++ {
			decoded.WriteByte(temp[i])
		}
		if n == 0 {
			break
		}
	}
	ioutil.WriteFile(fName, decoded.Bytes(), 0644)
}

//SaveBase64Code  save the base64 code to file
func SaveBase64Code(bName string, bCode []byte) {
	err := ioutil.WriteFile(bName, bCode, 0644)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

//ReadBase64Code  read the base64 code from file
func ReadBase64Code(fName string) (bCode []byte) {
	bCode, err := ioutil.ReadFile(fName)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	return bCode
}

//Helper 打印帮助信息
func Helper() {
	fmt.Println("Base64 Converter v0.0.1<Michael Jiang:sencom.top>")
	fmt.Println("file to Base64 code:")
	fmt.Println("Base64Conv -mode=F2B -input=[filename] -output=[filename]")
	fmt.Println("Base64 code to file:")
	fmt.Println("Base64Conv -mode=B2F -input=[filename] -output=[filename]")
}
func main() {
	//命令行参数：第一个参数，为参数名称，第二个参数为默认值，第三个参数是说明
	input := flag.String("input", "", "input filename")
	output := flag.String("output", "", "output filename")
	mode := flag.String("mode", "", "Conv mode")
	flag.Parse()
	var bCode []byte
	switch *mode {
	case "F2B":
		bCode = F2B(*input)
		SaveBase64Code(*output, bCode)
	case "B2F":
		bCode = ReadBase64Code(*input)
		B2F(*output, bCode)
	default:
		Helper()
	}

}
