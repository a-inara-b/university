package Index

func FirstOccur(arr []float64, num float64) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
			break
		}
	}
}

func LastOccur(arr []float64, num float64) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == num {
			return i
			break
		}
	}
}

func Count(arr []float64, num float64) int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			cnt = cnt + 1
		}
	}
	return cnt
}

func Merge(arr1, arr2 []float64) []float64 {
	var res [len(arr1) + len(arr2)]float64
	for i := 0; i < len(arr1); i++ {
		res[i] = arr1[i]
	}
	for i := 0; i < len(arr2); i++ {
		res[i+len(arr1)] = arr2[i]
	}
	return res
}

func Sort(n []float64) []float64 {
	var i = 1
	for i < len(n) {
		var j = i
		for j >= 1 && n[j] < n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]

			j--
		}

		i++
	}
	return n
}
