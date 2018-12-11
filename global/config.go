package global

import (
	"ZCache/types"
	"sync"
)

var Config = struct {
	MaxLen int64
	}{
	MaxLen:1024,
}

var GlobalVar = struct {
	Root *types.Node
	GRoot []*types.Node
	GRWLock *sync.RWMutex
}{
	Root:nil,
	GRoot:nil,
	GRWLock:nil,
}
