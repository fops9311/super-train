package memorizeresult

type Memorized func(int) int

func Cachemem(f Memorized) Memorized {
	cache := make(map[int]int)
	return func(in int) int {
		if val, ok := cache[in]; ok {
			return val
		}
		res := f(in)
		cache[in] = res
		return res
	}
}
