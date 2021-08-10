package main

func All(s []string, f func(string) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func Any(s []string, f func(string) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}
func Filter(s []string, f func(string) bool) (sOut []string) {
	sOut = make([]string, 0)
	for _, v := range s {
		if f(v) {
			sOut = append(sOut, v)
		}
	}
	return sOut
}
