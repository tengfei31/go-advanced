package test

import (
	"math/rand"
	"testing"
	"time"
)

type Number []int

func TestRandom(t *testing.T) {
	var arr = generateArr()
	t.Log(arr)
	// 合并一个数组
	var newArr Number
	for i :=0; i < len(arr); i++ {
		newArr = append(newArr, arr[i]...)
	}
	t.Log(newArr)
	// 组合数字
	for i := 0; i < len(newArr); i++ {

	}

}

// generateArr 生成二维数组
func generateArr() []Number {
	rand.Seed(time.Now().UnixNano())
	var arr []Number
	arr = make([]Number, 6)
	for i := 0; i < 6; i++ {
		arr[i] = make(Number, 4)
		for j := 0; j < 4; j++ {
			arr[i][j] = rand.Int()
		}
	}
	return arr
}
