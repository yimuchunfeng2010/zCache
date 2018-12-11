package global

import (
	"ZCache/data"
)

var Config = struct {
	MaxLen int
}{
	MaxLen:1024,
}

var GlobalVar = struct {
	Root *zdata.Node
}{
	Root:nil,
}
