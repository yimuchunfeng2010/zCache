package services

import (
	"github.com/samuel/go-zookeeper/zk"
	"ZCache/tool/logrus"
	"time"
	"fmt"
	"strings"
	"strconv"
)


func main() {

	fmt.Printf("ZKOperateTest\n")

	var hosts = []string{"192.168.228.141:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*500)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	//var path = "/ZAB"
	//var data = []byte("1234")
	//var flags int32 = 1
	//// permission
	//var acls = zk.WorldACL(zk.PermAll)
	//
	//// create
	//p, err_create := conn.Create(path, data, flags, acls)
	//if err_create != nil {
	//	fmt.Println("AAA", err_create)
	//	return
	//}
	//fmt.Println("created:", p)

	msg, _, err := conn.Get("/ZAB")
	if err != nil {
		fmt.Println("BBB", err)
		return
	}
	fmt.Println("CCC", string(msg))
}
//package main
//
//import (
//"fmt"
//"github.com/samuel/go-zookeeper/zk"
//"time"
//)
//
//var hosts = []string{"192.168.228.141:2181"}
//
//var path1 = "/whatzk"
//
//var flags int32 = zk.FlagEphemeral
//var data1 = []byte("hello,this is a zk go test demo!!!")
//var acls = zk.WorldACL(zk.PermAll)
//
//func main() {
//	option := zk.WithEventCallback(callback)
//	conn, _, err := zk.Connect(hosts, time.Second*5, option)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer conn.Close()
//
//	_, _, ech, err := conn.ExistsW(path1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	create(conn, path1, data1)
//
//	_, _, ech, err = conn.ExistsW(path1)
//	if err != nil {
//		fmt.Println("EEEE", err)
//		return
//	}
//	conn.Delete(path1, 0)
//	fmt.Println("LLL")
//	go watchCreataNode(ech)
//}
//
//func callback(event zk.Event) {
//	fmt.Println("*******************")
//	fmt.Println("path:", event.Path)
//	fmt.Println("type:", event.Type.String())
//	fmt.Println("state:", event.State.String())
//	fmt.Println("-------------------")
//}
//
//func create(conn *zk.Conn, path string, data []byte) {
//	_, err_create := conn.Create(path, data, flags, acls)
//	if err_create != nil {
//		fmt.Println(err_create)
//		return
//	}
//
//}
//
//func watchCreataNode(ech <-chan zk.Event) {
//	event := <-ech
//	fmt.Println("*******************")
//	fmt.Println("path:", event.Path)
//	fmt.Println("type:", event.Type.String())
//	fmt.Println("state:", event.State.String())
//	fmt.Println("-------------------")
//}

// 获取子节点
//
//p, err_create = conn.Create("/ZAB/r", []byte("ass"), 3, acls)
//if nil != err_create {
//fmt.Println("EEEE", err_create)
//return
//}
//fmt.Println("TTTT", p)
//p, err_create := conn.Create("/ZAB/w", []byte("ass"), 3, acls)
//if nil != err_create {
//fmt.Println("EEEE", err_create)
//return
//}
//fmt.Println("OOO", p)
//
//p, err_create = conn.Create("/ZAB/w", []byte("ass"), 3, acls)
//if nil != err_create {
//fmt.Println("EEEE", err_create)
//return
//}
//fmt.Println("OOO", p)
//msg, _, err := conn.Children("/ZAB")
//if err != nil {
//fmt.Println("BBB", err)
//return
//}
//
//fmt.Println("CCC", msg)
////flags有4种取值：
////0:永久，除非手动删除
////zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
////zk.FlagSequence  = 2:会自动在节点后面添加序号
////3:Ephemeral和Sequence，即，短暂且自动添加序号
//// permission
//option := zk.WithEventCallback(callback)
//conn, _, err := zk.Connect(hosts, time.Second*5, option)
//if err != nil {
//fmt.Println(err)
//return
//}
//defer conn.Close()
//
//_, _, ech, err := conn.ExistsW(path1)
//if err != nil {
//fmt.Println(err)
//return
//}
//
//create(conn, path1, data1)
//
//_, _, ech, err = conn.ExistsW(path1)
//if err != nil {
//fmt.Println("EEEE", err)
//return
//}
//conn.Delete(path1, 0)
//fmt.Println("LLL")
//go watchCreataNode(ech)

func ZookeeperInit() error {

	var hosts = []string{"192.168.228.141:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}
	defer conn.Close()

	// 创建Lock节点(永久节点)
	var lockPath = "/Lock"
	var lockData []byte = []byte("Lock")
	var lockFlags int32 = 0
	_, err = conn.Create(lockPath, lockData, lockFlags, zk.WorldACL(zk.PermAll))
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}

	return nil
}
func Lock() (lockName string, err error) {

	// 创建临时写节点
	var hosts = []string{"192.168.228.141:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return "", err
	}
	defer conn.Close()

	// 获取当前子节点
	children, _, err := conn.Children("/Lock")
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return "", err
	}

	maxChild := GetMaxchild(children)

	// 创建当前节点
	var wLockPath = "/Lock/w"
	var wLockData []byte = []byte(strconv.FormatInt(time.Now().Unix(), 10))
	var wLockFlags int32 = 1 // 永久序列增长节点
	var acl = zk.WorldACL(zk.PermAll)

	lockPath, err := conn.Create(wLockPath, wLockData, wLockFlags, acl)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return "", err
	}

	if "" != maxChild {
		// 对最大子节点设置观察点
		_, _, ech, err := conn.ExistsW(maxChild)
		if err != nil {
			logrus.Errorf("%s", err.Error())
			return "", err
		}
		timeout := 10 // 超时时间10s
		for timeout > 0 {
			select {
			case _, ok := <-ech:
				if ok {
					return lockPath, nil
				}
			default:
				time.Sleep(time.Second)
				timeout--
			}
		}
	} else {
		return lockPath, nil
	}

	return "", nil
}

func Unlock(lockName string) error {

	var hosts = []string{"192.168.228.141:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}
	defer conn.Close()

	// 删除节点
	err = conn.Delete(lockName, 0)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}

	return nil
}

func RLock() (lockName string, err error) {
	// 创建临时写节点
	var hosts = []string{"192.168.228.141:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return "", err
	}
	defer conn.Close()

	// 获取当前子节点
	children, _, err := conn.Children("/Lock")
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return "", err
	}

	maxChild := GetMaxWritechild(children)

	// 创建子节点
	var wLockPath = "/Lock/r"
	var wLockData []byte = []byte(strconv.FormatInt(time.Now().Unix(), 10))
	var wLockFlags int32 = 1 // 永久序列增长节点
	var acl = zk.WorldACL(zk.PermAll)

	lockPath, err := conn.Create(wLockPath, wLockData, wLockFlags, acl)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return "", err
	}

	if "" != maxChild {
		// 对最大子节点设置观察点
		_, _, ech, err := conn.ExistsW(maxChild)
		if err != nil {
			logrus.Errorf("%s", err.Error())
			return "", err
		}
		timeout := 10 // 超时时间10s
		for timeout > 0 {
			select {
			case _, ok := <-ech:
				if ok {
					return lockPath, nil
				}
			default:
				time.Sleep(time.Second)
				timeout--
			}
		}
	} else {
		return lockPath, nil
	}

	return "", nil
}

func RUnlock(lockName string) error {
	var hosts = []string{"192.168.228.141:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}
	defer conn.Close()

	// 删除节点
	err = conn.Delete(lockName, 0)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}

	return nil
}

func GetMaxchild(children []string) (child string) {
	if 1 == len(children) {
		return children[0]
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
func Callback(event zk.Event) {
	fmt.Println("*******************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("*******************")
}