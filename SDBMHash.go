package hashstr

func SDBMHash(key string) uint {
	hash := uint(0)
	for i:=0; i<len(key); i++ {
		hash = uint(key[i]) + (hash<<6) + (hash<<16) - hash
	}
	return hash
}
