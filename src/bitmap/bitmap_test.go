package bitmap

import (
	"fmt"
	"testing"
)

// TestBitMap
//
//	@Description: 测试 bitmap 功能
//	@Param t:
func TestBitMap(t *testing.T) {
	bitMap := new(BitMap)
	bitMap.Init(100)
	bitMap.Set(10)
	bitMap.Set(88)
	bitMap.Set(99)
	fmt.Println(bitMap.IsHave(888))
	fmt.Println(bitMap.IsHave(88))
}

// TestCreateNumFile
//
//	@Description: 生成 1000 万个整数的文件
//	@Param t:
func TestCreateNumFile(t *testing.T) {
	randomNum()
}

// TestSplitFile
//
//	@Description: 将大文件拆分为小文件
//	@Param t:
func TestSplitFile(t *testing.T) {
	splitFile()
}
