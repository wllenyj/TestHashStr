package hashstr

import (
	"github.com/dgryski/dgohash"
)

func JavaHash(key string) uint {
	hash := dgohash.NewJava32()
	hash.Write([]byte(key))
	return uint(hash.Sum32())
}

func DJBHash(key string) uint {
	hash := dgohash.NewDjb32()
	hash.Write([]byte(key))
	return uint(hash.Sum32())
}

func SuperFastHash(key string) uint {
	hash := dgohash.NewSuperFastHash()
	hash.Write([]byte(key))
	return uint(hash.Sum32())
}
