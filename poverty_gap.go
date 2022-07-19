package main


//实现一个函数，输入为任意长度的int64数组，返回元素最大差值，例如输入arr=[5,8,10,1,3]，返回9。
func povertyGap(arr []int64) int64 {
	max, min := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}
	return max - min
}

func main() {
	val := []int64{5, 8, 10, 1, 3}
	res := povertyGap(val)
	println(res)
}
