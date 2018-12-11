package global

import (
	"ZCache/types"
)

var Config = struct {
	MaxLen int64
	}{
	MaxLen:1024,
}

var GlobalVar = struct {
	Root *types.Node
	GRoot []*types.Node
}{
	Root:nil,
	GRoot:nil,
}
