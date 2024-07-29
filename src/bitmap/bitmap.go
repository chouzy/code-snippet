package bitmap

// BitMap
//
//	@Description: BitMap 结构体
type BitMap struct {
	Size int    // 数组的大小
	Data []byte // 存放数值的 byte 数组
}

// Init
//
//	@Description: 初始化 BitMap
//	@Param size: bitmap 的大小
func (b *BitMap) Init(size int) {
	b.Size = size
	b.Data = make([]byte, (size+7)/8)
}

// Set
//
//	@Description: 向 bitmap 中添加值
//	@Param num: 要添加的值
func (b *BitMap) Set(num int) {
	if num >= b.Size {
		return
	}
	b.Data[num/8] |= 1 << (7 - (num % 8))
}

// IsHave
//
//	@Description: 判断值是否存在于 bitmap
//	@Param num: 要判断的值
//	@Return bool: true-存在; false-不存在
func (b *BitMap) IsHave(num int) bool {
	if num >= b.Size {
		return false
	}
	bit := b.Data[num/8]
	return (bit & (1 << (7 - num%8))) != 0
}
