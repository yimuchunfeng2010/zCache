package services

import (
	"github.com/samuel/go-zookeeper/zk"
	"zCache/tool/logrus"
	"zCache/global"
	"time"
	//"fmt"
	"strings"
	"strconv"
)

func ZookeeperInit() error {

	var hosts = []string{global.Config.ZkIPaddr}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		logrus.Errorf(err.Error())
		return err
	}
	defer conn.Close()

	// 创建Lock节点(永久节点)
	var lockPath = "/Lock"
	var lockData []byte = []byte("Lock")
	var lockFlags int32 = 0
	var acls = zk.WorldACL(zk.PermAll)
	_, err = conn.Create(lockPath, lockData, lockFlags, acls)
	if err != nil {
		if 0 == strings.Compare("zk: node already exists", err.Error()) {
			logrus.Infof("Create Node %s Success", lockPath)
			return nil
		}
		logrus.Errorf("%s", err.Error())
		return err
	}

	// 创建Lock节点(永久节点)
	var ClusterPath = "/Cluster"
	var ClusterData []byte = []byte("Cluster")
	var ClusterFlags int32 = 0
	_, err = conn.Create(ClusterPath, ClusterData, ClusterFlags, acls)
	if err != nil {
		if 0 == strings.Compare("zk: node already exists", err.Error()) {
			logrus.Infof("Create Node %s Success", lockPath)
			return nil
		}
		logrus.Errorf("%s", err.Error())
		return err
	}

	return nil
}
func Lock() (lockName string, err error) {

	//// 创建临时写节点
	//var hosts = []string{global.Config.ZkIPaddr}
	//conn, _, err := zk.Connect(hosts, time.Second*5)
	//if err != nil {
	//	logrus.Errorf("%s", err.Error())
	//	return "", err
	//}
	//defer conn.Close()
	//
	//// 获取当前子节点
	//children, _, err := conn.Children("/Lock")
	//if err != nil {
	//	logrus.Errorf("%s", err.Error())
	//	return "", err
	//}
	//
	//fmt.Println("children", children)
	//maxChild := GetMaxchild(children)
	//
	//// 创建当前节点
	//var wLockPath = "/Lock/w"
	//var wLockData []byte = []byte(strconv.FormatInt(time.Now().Unix(), 10))
	//var wLockFlags int32 = 2 // 永久序列增长节点
	//var acl = zk.WorldACL(zk.PermAll)
	//
	//lockPath, err := conn.Create(wLockPath, wLockData, wLockFlags, acl)
	//if err != nil {
	//	logrus.Errorf("%s", err.Error())
	//	return "", err
	//}
	//
	//if "" != maxChild {
	//	// 对最大子节点设置观察点
	//	_, _, ech, err := conn.ExistsW(maxChild)
	//	if err != nil {
	//		logrus.Errorf("%s", err.Error())
	//		return "", err
	//	}
	//	timeout := 60 // 超时时间10s
	//	for timeout > 0 {
	//		select {
	//		case _, ok := <-ech:
	//			if ok {
	//				return lockPath, nil
	//			}
	//		default:
	//			time.Sleep(time.Second)
	//			timeout--
	//		}
	//	}
	//	return "", nil
	//} else {
	//	return lockPath, nil
	//}
	//
	//return "", nil
	return "lock",nil
}

func Unlock(lockName string) error {

	//var hosts = []string{global.Config.ZkIPaddr}
	//conn, _, err := zk.Connect(hosts, time.Second*5)
	//if err != nil {
	//	logrus.Errorf("%s", err.Error())
	//	return err
	//}
	//defer conn.Close()
	//
	//// 删除节点
	//err = conn.Delete(lockName, 0)
	//if err != nil {
	//	logrus.Errorf("%s", err.Error())
	//	return err
	//}
	//
	//return nil
	return nil
}

func RLock() (lockName string, err error) {
	//// 创建临时写节点
	//var hosts = []string{global.Config.ZkIPaddr}
	//conn, _, err := zk.Connect(hosts, time.Second*5)
	//if err != nil {
	//	logrus.Errorf("%s", err.Error())
	//	return "", err
	//}
	//defer conn.Close()
	//
	//// 获取当前子节点
	//children, _, err := conn.Children("/Lock")
	//if err != nil {
	//	logrus.Errorf("%s", err.Error())
	//	return "", err
	//}
	//
	//maxChild := GetMaxWritechild(children)
	//
	//// 创建子节点
	//var wLockPath = "/Lock/r"
	//var wLockData []byte = []byte(strconv.FormatInt(time.Now().Unix(), 10))
	//var wLockFlags int32 = 2 // 永久序列增长节点
	//var acl = zk.WorldACL(zk.PermAll)
	//
	//lockPath, err := conn.Create(wLockPath, wLockData, wLockFlags, acl)
	//if err != nil {
	//	logrus.Errorf("%s", err.Error())
	//	return "", err
	//}
	//
	//if "" != maxChild {
	//	// 对最大子节点设置观察点
	//	_, _, ech, err := conn.ExistsW(maxChild)
	//	if err != nil {
	//		logrus.Errorf("%s", err.Error())
	//		return "", err
	//	}
	//	timeout := 60 // 超时时间10s
	//	for timeout > 0 {
	//		select {
	//		case _, ok := <-ech:
	//			if ok {
	//				return lockPath, nil
	//			}
	//		default:
	//			time.Sleep(time.Second)
	//			timeout--
	//		}
	//	}
	//	return "", nil
	//} else {
	//	return lockPath, nil
	//}
	//
	//return "", nil
	return "lock",nil
}

func RUnlock(lockName string) error {
	//var hosts = []string{global.Config.ZkIPaddr}
	//conn, _, err := zk.Connect(hosts, time.Second*5)
	//if err != nil {
	//	logrus.Errorf("%s", err.Error())
	//	return err
	//}
	//defer conn.Close()
	//
	//// 删除节点
	//err = conn.Delete(lockName, 0)
	//if err != nil {
	//	logrus.Errorf("%s", err.Error())
	//	return err
	//}
	//
	//return nil
	return nil
}

func GetMaxchild(children []string) (child string) {
	if 0 == len(children) {
		return ""
	}

	var maxChild = children[0]
	maxIndex := maxChild[1:]
	for _, value := range children {
		curIndex := value[1:]
		if curIndex > maxIndex {
			maxIndex = curIndex
		}
		maxChild = value
	}

	return maxChild
}

func GetMaxWritechild(children []string) (child string) {
	//过滤所有写节点
	writeChildren := make([]string, 0)
	for _, value := range children {
		if strings.HasPrefix(value, "w") {
			writeChildren = append(writeChildren, value)
		}
	}

	if 0 == len(writeChildren) {
		return ""
	}

	var maxChild = children[0]
	maxIndex := maxChild[1:]
	for _, value := range children {
		curIndex := value[1:]
		if curIndex > maxIndex {
			maxIndex = curIndex
		}
		maxChild = value
	}

	return maxChild
}

func RegisterNode()(error){
	var hosts = []string{global.Config.ZkIPaddr}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}
	defer conn.Close()

	var nodePath = "/Cluster/"
	var nodeData []byte = []byte(strconv.FormatInt(time.Now().Unix(), 10))
	var nodeFlags int32 = 4 // 临时增长序列节点
	var acl = zk.WorldACL(zk.PermAll)

	_, err = conn.Create(nodePath, nodeData, nodeFlags, acl)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return  err
	}

	return nil
}


func GetWorkingNode()(int,error){
	var hosts = []string{global.Config.ZkIPaddr}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return -1, err
	}
	defer conn.Close()

	children, _, err := conn.Children("/Lock")
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return -1, err
	}

	return len(children), nil
}