package services

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
)


func main() {

	fmt.Printf("ZKOperateTest\n")

	var hosts = []string{"192.168.228.143:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*500)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	var path = "/ZAB"
	var data = []byte("1234")
	var flags int32 = 1
	// permission
	var acls = zk.WorldACL(zk.PermAll)

	// create
	p, err_create := conn.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println("AAA", err_create)
		return
	}
	fmt.Println("created:", p)

	msg, _, err := conn.Get("/ZAB")
	if err != nil {
		fmt.Println("BBB", err)
		return
	}
	fmt.Println("CCC", string(msg))
}
