package array

func Contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func Uniq(arr []string) []string {
	m := make(map[string]struct{})
	for _, ele := range arr {
		m[ele] = struct{}{}
	}

	uniq := [] string{}
	for i := range m {
		uniq = append(uniq, i)
	}
	return uniq
}