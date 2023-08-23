package toolkit

// IndexStrArr return the first index of string s found in arr;
// If not found ,return -1
func IndexStrArr(s string, arr []string) int {
	for i, v := range arr {
		if s == v {
			return i
		}
	}
	return -1
}

// 字符串切片去重（用hashMap）
func DeduplicateStrArrWithMap(arr []string) []string {
	res := make([]string, 0, len(arr))
	m := make(map[string]bool, len(arr))
	for _, v := range arr {
		_, exist := m[v]
		if exist { //存在
			continue
		}
		//不存在相同,追加
		m[v] = true
		res = append(res, v)
	}
	return res
}

// 字符串切片值是否相等
func StrArrEqual(arr1 []string, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// 字符串切片求交集
func StrArraysIntersect(arr1 []string, arr2 []string) []string {
	res := make([]string, 0, len(arr1))
	for _, v := range arr1 {
		if IndexStrArr(v, arr2) != -1 {
			res = append(res, v)
		}
	}
	return res
}
