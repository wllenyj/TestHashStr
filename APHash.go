package hashstr

func APHash(key string) uint {
	hash := uint(0)
	for i := 0; i < len(key); i++ {
		if (i & 1) == 0 {
			hash ^= ((hash << 7) ^ uint(key[i]) ^ (hash >> 3))
		} else {
			hash ^= (^((hash << 11) ^ uint(key[i]) ^ (hash >> 5)))
		}
	}
	return hash
}
