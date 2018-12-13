package tool

import (
	"ZCache/global"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"time"
	"ZCache/types"
)

func Md5Encode(msg string) []byte {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(msg))
	result := Md5Inst.Sum([]byte(""))
	return result
}

func ByteToInt(msg []byte) (int64, error) {
	encodedStr := "0x" + hex.EncodeToString(msg)
	data, err := strconv.ParseInt(encodedStr, 0, 64)
	if err != nil {
		return -1, err
	}
	return data, nil
}

func GetHashIndex(msg string) (int64, error) {
	msgByte := Md5Encode(msg)
	msgByte = msgByte[0 : len(msgByte)/2-1]
	data, err := ByteToInt(msgByte)
	if err != nil {
		return -1, err
	}
	data = data % global.Config.MaxLen
	return data, nil
}

func GetDataLogFileName() string {
	now := time.Now()
	x := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)
	var fileName string
	if "windows" == runtime.GOOS {
		fileName = "data_log/" + x.Format("2006-01-02_15-04-05_112") + ".txt"
	} else {
		fileName = "data_log\\" + x.Format("2006-01-02_15-04-05_112") + ".txt"
	}

	return fileName
}

func GetAllFile(pathname string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return nil, err
	}
	files := make([]string, 0)
	for _, fi := range rd {
		if false == fi.IsDir() {
			files = append(files, fi.Name())
		}

	}

	return files, err
}

func GetNewestFile(pathname string) (string, error) {
	files, err := GetAllFile(pathname)
	if err != nil {
		return "", err
	}
	fmt.Println(files)
	if len(files) == 0 {
		return "", errors.New("File Not Found")
	}
	newFile := ""
	var timeIndex int64 = 0
	for _, file := range files {
		curIndex := GetFileModTime(file)
		if curIndex >= timeIndex {
			newFile = file
		}
	}
	return newFile, nil
}

//获取文件修改时间 返回unix时间戳
func GetFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

func GetDataLogDir() string {
	dir := ""
	if "windows" == runtime.GOOS {
		dir = "./data_log/"
	} else {
		dir = ".\\data_log\\"
	}
	return dir
}

func CurrentFile() (string, int) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "", -1
	}
	return file, line

}
func GetFileNameLine() string {
	currentFile, line := CurrentFile()
	fileInfo := currentFile + " " + strconv.Itoa(line)

	return fileInfo
}

func ClusterHealthCheck(operation types.OperationType)(bool , error){
	switch operation {
	case types.OPERATION_TYPE_SET,types.OPERATION_TYPE_DELETE,types.OPERATION_TYPE_POST:
		if global.GlobalVar.GClusterHealthState != types.CLUSTER_HEALTH_TYPE_HEALTH{
			return false, nil
		}
	case types.OPERATION_TYPE_GET:
		if global.GlobalVar.GClusterHealthState == types.CLUSTER_HEALTH_TYPE_UNHEALTH{
			return false, nil
		}
	default:
		return false, errors.New("Invalid Operation")
	}

	return true,nil
}
