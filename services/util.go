package services

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"ZCache/global"
)

func Md5Encode(msg string) []byte{
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(msg))
	result := Md5Inst.Sum([]byte(""))
	return result
}

func ByteToInt(msg []byte, bitSize int)(int64, error){
	encodedStr := "0x" + hex.EncodeToString(msg)
	data, err := strconv.ParseInt(encodedStr, 0, bitSize)
	if err != nil {
		return -1 , err
	}
	return data, nil
}


func GetHashIndex(msg string)(int64, error) {
	msgByte := Md5Encode(msg)
	data, err := ByteToInt(msgByte,global.Config.MaxLen)
	if err != nil {
		return -1 , err
	}
	return data, nil
}