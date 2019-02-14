package external_data

import (
	"github.com/garyburd/redigo/redis"
	"zCache/tool/logrus"
	"zCache/data"
)

func ImportFromRedis() (error){
	conn, err := redis.Dial("tcp", "192.168.228.143:6379")
	if err != nil {
		logrus.Warningf("redis connection fail [err: %+v]", err.Error())
		return err
	}
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("keys", "*"))
	if err != nil {
		logrus.Warningf("redis connection fail [err: %+v]", err.Error())
		return err
	}

	for _, key := range keys {
		value, err := redis.String(conn.Do("GET", key))
		if err != nil {
			logrus.Warningf("redis get fail [key: %+v, err: %+v]", key, err.Error())
			continue
		}
		// 导入数据
		_, err = zdata.CoreAdd(key,value)
		if err != nil {
			logrus.Warningf("CoreAdd fail [key: %+v, value: %+v, err: %+v]", key, value, err.Error())
			continue
		}
	}

	return nil
}

func ExportToRedis() (error){
	conn, err := redis.Dial("tcp", "192.168.228.143:6379")
	if err != nil {
		logrus.Warningf("redis connection fail [err: %+v]", err.Error())
		return err
	}
	defer conn.Close()

	node, err := zdata.CoreGetAll()
	if err != nil {
		logrus.Warningf("get all fail [err: %+v]", err.Error())
		return err
	}

	curNode := node
	for curNode != nil{
		_, err := redis.String(conn.Do("SET", curNode.Key, curNode.Value))
		if err != nil {
			logrus.Warningf("set fail [key: %+v, value: %+v, err: %+v]", curNode.Key, curNode.Value, err.Error())
			continue
		}
		curNode = curNode.Next
	}

	return nil
}