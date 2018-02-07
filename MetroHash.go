package hashstr

import (
	"github.com/dgryski/go-metro"
)

func MetroHash(key string) uint {
	return uint(metro.Hash64Str(key, 0))
}
