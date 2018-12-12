package global

import (
	"ZCache/types"
	"sync"
)

var Config = struct {
	MaxLen int64
	UserEmail string
	MailPort string
	MailAuthCode string
	MailSmtpHost string
	ToMail string
}{
	MaxLen: 1024,
	UserEmail:"123456789@qq.com",
	MailPort:"587",
	MailAuthCode:"mnkxcahklkrebfbb",
	MailSmtpHost: "smtp.qq.com",
	ToMail:"987654321@qq.com",
}

var GlobalVar = struct {
	Root    *types.Node
	GRoot   []*types.Node
	GRWLock *sync.RWMutex


}{
	Root:    nil,
	GRoot:   nil,
	GRWLock: nil,

}
