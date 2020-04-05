package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

/*
* @CreateTime: 2019/12/14 18:20
* @Author: hujiaming
* @Description: 对string类型进行加解密
 */

type Data struct {
	NewData string
}

func EncryptString(source string) []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(Data{NewData: source})
	if err != nil {
		log.Panic(err)
	}
	return encoded.Bytes()
}
func DecryptString(source []byte) string {
	var data Data
	decoder := gob.NewDecoder(bytes.NewReader(source))
	err := decoder.Decode(&data)
	if err != nil {
		log.Panic(err)
	}
	return data.NewData
}
