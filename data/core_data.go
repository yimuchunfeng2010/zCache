package zdata

import ("ZCache/services"
	"ZCache/global"
	"errors"
	"ZCache/types"
)


func CoreAdd(key string, data types.CacheData) (*types.Node, error) {
	hashIndex , err := services.GetHashIndex(key)
	if err != nil {
		return nil , err
	}
	if nil == global.GlobalVar.GRoot{
		return nil, errors.New("GRoot nil")
	}
	global.GlobalVar.GRoot[hashIndex] , err = Add(global.GlobalVar.GRoot[hashIndex], key, data)
	if err != nil {
		return nil, err
	}

	return global.GlobalVar.GRoot[hashIndex],nil
}

func CoreDelete(key string) (*types.Node, error) {
	hashIndex , err := services.GetHashIndex(key)
	if err != nil {
		return nil , err
	}
	if nil == global.GlobalVar.GRoot{
		return nil, errors.New("GRoot nil")
	}

	global.GlobalVar.GRoot[hashIndex] , err = Delete(global.GlobalVar.GRoot[hashIndex], key)
	if err != nil {
		return nil, err
	}

	return global.GlobalVar.GRoot[hashIndex],nil
}

//查找并返回节点
func CoreUpdate(key string, data types.CacheData) (*types.Node, error) {
	hashIndex , err := services.GetHashIndex(key)
	if err != nil {
		return nil , err
	}

	if nil == global.GlobalVar.GRoot{
		return nil, errors.New("GRoot nil")
	}

	global.GlobalVar.GRoot[hashIndex] , err = Update(global.GlobalVar.GRoot[hashIndex], key, data)
	if err != nil {
		return nil, err
	}

	return global.GlobalVar.GRoot[hashIndex],nil
}

//查找并返回节点
func CoreGet(key string) (*types.Node, error) {
	hashIndex , err := services.GetHashIndex(key)
	if err != nil {
		return nil , err
	}
	if nil == global.GlobalVar.GRoot{
		return nil, errors.New("GRoot nil")
	}

	node , err := Get(global.GlobalVar.GRoot[hashIndex], key)
	if err != nil {
		return nil, err
	}
	return node,nil
}