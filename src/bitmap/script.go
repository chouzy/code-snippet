package bitmap

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const fileName = "random.txt"
const numIntegers = 10000000

// randomNum
//
//	@Description: 生成 1000 万个整数存入文件
func randomNum() {
	minValue := 1
	maxValue := 1000000

	// 设置随机数种子，确保每次运行生成的随机数不同
	rand.Seed(time.Now().UnixNano())

	// 创建文件并打开以供写入
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()

	// 创建写入缓冲区
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// 生成并写入随机整数到文件
	for i := 0; i < numIntegers; i++ {
		randomInt := rand.Intn(maxValue-minValue+1) + minValue
		_, err := fmt.Fprintf(writer, "%d\n", randomInt)
		if err != nil {
			fmt.Println("写入文件失败:", err)
			return
		}
	}

	fmt.Println("成功生成和写入", numIntegers, "个随机整数到文件:", fileName)
}

// splitFile
//
//	@Description: 将大文件拆分为小文件
func splitFile() {
	numIntegersPerFile := 100000
	numFiles := numIntegers / numIntegersPerFile
	sortDir := "sort"
	if err := os.Mkdir(sortDir, 0755); err != nil && !os.IsExist(err) {
		fmt.Println("无法创建目录:", err)
		return
	}
	// 打开原始文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建读取缓冲区
	scanner := bufio.NewScanner(file)
	// 拆分原始文件为小文件
	for i := 0; i <= numFiles; i++ {
		// 创建小文件并打开以供写入
		name := fmt.Sprintf("%s/%d.txt", sortDir, i)
		smallFile, err := os.Create(name)
		if err != nil {
			fmt.Println("无法创建文件:", err)
			return
		}
		defer smallFile.Close()

		// 创建写入缓冲区
		writer := bufio.NewWriter(smallFile)
		defer writer.Flush()

		// 写入指定数量的文件到小文件
		for j := 0; j < numIntegersPerFile && scanner.Scan(); j++ {
			integer := scanner.Text()
			_, err := fmt.Fprintln(writer, integer)
			if err != nil {
				fmt.Println("写入文件失败:", err)
				return
			}
		}
	}
	fmt.Println("拆分完成")
}
