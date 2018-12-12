package zdata

import (
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
	"ZCache/types"
	"errors"
	"fmt"
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
		logrus.Warningf(fmt.Sprintf("CoreGet Failed[Key:%s,Err:%s]", key, err.Error()))
		return nil, err
	}
	return node, nil
}

func CoreGetAll()(*types.DataNode, error){
	if nil == global.GlobalVar.GRoot {
		return nil, errors.New("GRoot nil")
	}

	var index int64
	var rspRoot *types.DataNode = nil
	for index = 0; index < global.Config.MaxLen;index++{
		err := GetAll(global.GlobalVar.GRoot[index], index, &rspRoot,&rspRoot)
		if err != nil{
			return nil , err
		}
	}
	return rspRoot,nil
}