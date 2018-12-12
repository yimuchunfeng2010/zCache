package zdata

import (
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
	"ZCache/types"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func CoreAdd(key string, value string) (*types.Node, error) {
	hashIndex, err := services.GetHashIndex(key)
	if err != nil {
		return nil, err
	}
	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}
	tmpNode, err := Add(global.GlobalVar.GRoot[hashIndex], key, value)
	if err != nil {
		return nil, err
	}

	global.GlobalVar.GRoot[hashIndex] = tmpNode

	return global.GlobalVar.GRoot[hashIndex], nil
}

func CoreDelete(key string) (*types.Node, error) {
	hashIndex, err := services.GetHashIndex(key)
	if err != nil {
		return nil, err
	}
	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}

	tmpNode, err := Delete(global.GlobalVar.GRoot[hashIndex], key)
	if err != nil {
		return nil, err
	}
	global.GlobalVar.GRoot[hashIndex] = tmpNode
	return global.GlobalVar.GRoot[hashIndex], nil
}

//查找并返回节点
func CoreUpdate(key string, Value string) (*types.Node, error) {
	hashIndex, err := services.GetHashIndex(key)
	if err != nil {
		return nil, err
	}

	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}

	tmpNode, err := Update(global.GlobalVar.GRoot[hashIndex], key, Value)
	if err != nil {
		return nil, err
	}
	global.GlobalVar.GRoot[hashIndex] = tmpNode

	return global.GlobalVar.GRoot[hashIndex], nil
}

//查找并返回节点
func CoreGet(key string) (*types.Node, error) {
	hashIndex, err := services.GetHashIndex(key)
	if err != nil {
		return nil, err
	}
	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}

	node, err := Get(global.GlobalVar.GRoot[hashIndex], key)
	if err != nil {
		logrus.Warningf("%s  CoreGet Failed[Key:%s,Err:%s]", services.GetFileNameLine(), key, err.Error())
		return nil, err
	}
	return node, nil
}

func CoreGetAll() (*types.DataNode, error) {
	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}

	var index int64
	var rspRoot *types.DataNode = nil
	for index = 0; index < global.Config.MaxLen; index++ {
		err := GetAll(global.GlobalVar.GRoot[index], index, &rspRoot, &rspRoot)
		if err != nil {
			return nil, err
		}
	}
	return rspRoot, nil
}

func CoreFlush() error {
	if nil == global.GlobalVar.GRoot {
		return errors.New("GRoot nil")
	}

	var index int64
	var rspRoot *types.DataNode = nil
	for index = 0; index < global.Config.MaxLen; index++ {
		err := GetAll(global.GlobalVar.GRoot[index], index, &rspRoot, &rspRoot)
		if err != nil {
			return err
		}
	}
	// 写文件
	file, err := os.OpenFile(services.GetDataLogFileName(), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileWrite := bufio.NewWriter(file)
	curNode := rspRoot
	for nil != curNode {
		msg := fmt.Sprintf("%s  %s\n", curNode.Key, curNode.Value)
		fileWrite.WriteString(msg)
		curNode = curNode.Next

	}
	fileWrite.Flush()

	return nil
}

func CoreImport() error {
	file, err := services.GetNewestFile(services.GetDataLogDir())
	if err != nil {
		return err
	}

	fi, err := os.Open(fmt.Sprintf("%s%s", services.GetDataLogDir(), file))
	if err != nil {
		logrus.Warningf("%s  Open File Failed! [Err:%s]\n", services.GetFileNameLine(), err.Error())
		return err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)

	for {
		data, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		array := strings.Split(string(data), "  ")
		if len(array) != 2 {
			logrus.Warningf("%s  Invaild Data! [Data: %s]\n", services.GetFileNameLine(), string(data))
			continue
		}
		key := array[0]
		value := array[1]
		_, err := CoreAdd(key, value)
		if err != nil {
			logrus.Warningf("%s  CoreAdd Data[Key: %s, Value: %s]\n", services.GetFileNameLine(), key, value)
			continue
		}

	}
	return nil
}
