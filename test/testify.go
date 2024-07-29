package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSimpleRand 简单使用
func TestSimpleRand(t *testing.T) {
	t.Log("start ...")
	ass := assert.New(t)
	ass.Equal(1, 1)
	ass.NotEqual(1, 2)
	ass.NotNil("123")
	ass.IsType([]string{}, []string{""})

	ass.Contains("Hello World", "World")
	ass.Contains(map[string]string{"Hello": "World"}, "Hello")
	ass.Contains([]string{"Hello", "World"}, "Hello")
	ass.True(true)
	ass.True(false)

	t.Log("next ...")
	var s []string
	ass.Empty(s)
	ass.Nil(s)
	t.Log("end ...")
}

// TestCalculate 一般用的更多的是表驱动方式把同一个单元的测试用例都放在一起
func TestCalculate(t *testing.T) {
	ass := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		ass.Equal(test.input, test.expected)
	}
}
