package main

import (
	"fmt"
	"math/rand"
	"time"
)

//实现一个函数，对输入的扑克牌执行洗牌，保证其是均匀分布的，也就是说列表中的每一张扑克牌出现在列表的每一个位置上的概率必须相同。
func Shuffle(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	val := []int{5, 8, 10, 1, 3}
	Shuffle(val)
	fmt.Printf("%+v\n", val)
}
