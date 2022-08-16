package helpers

func UniqueString(arr []string) []string {
	result := make([]string, 0, len(arr))
	encountered := map[string]bool{}
	for v := range arr {
		encountered[arr[v]] = true
	}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

/*
func uniqueInt(arr []int) []int {
	result := make([]int, 0, len(arr))
	encountered := map[int]bool{}
	for v := range arr {
		encountered[arr[v]] = true
	}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}*/
