package main

import (
	"encoding/base64"
	"os"
)

func main() {
	str := encode()
	
	file,_:= os.Create("encoded_str.txt")
	defer file.Close()
	file.Write([]byte(str))


	decode(str)
}

//エンコード
func encode() string {

	file, _ := os.Open("kt_s2.jpg")
	defer file.Close()

	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

//デコード
func decode(str string) {
	data, _ := base64.StdEncoding.DecodeString(str) //[]byte

	file, _ := os.Create("encode_and_decord.jpg")
	defer file.Close()

	file.Write(data)
}
