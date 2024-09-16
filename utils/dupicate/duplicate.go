package utils

func RemoveDuplicateElement[T string | int](arr []T) (ouco []T) {
	dupMap := make(map[T]bool)
	for _, dup := range arr {
		if !dupMap[dup] {
			dupMap[dup] = true
		}
	}
	for k, _ := range dupMap {
		ouco = append(ouco, k)
	}

	return
}
