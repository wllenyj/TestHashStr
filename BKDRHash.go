package hashstr

const primeRK = 7
func Hash1(sep string) uint {
	hash := uint(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint(sep[i])
	}
	return hash
}

const primeRK2 = 16777619
func Hash2(sep string) uint {
	hash := uint(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK2 + uint(sep[i])
	}
	return hash
}

const primeRK3 = 131 
func Hash3(sep string) uint {
	hash := uint(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK3 + uint(sep[i])
	}
	return hash
}
