package main

import (
	"bufio"
	"fmt"
	"os"
)

func WriteLines(filePath string, lines []string) {
	//创建一个新文件，写入内容 5 句 “http://c.biancheng.net/golang/”
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for _, line := range lines {
		write.WriteString(line)
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func main() {
	//lines := make([]string, 0)
	//for i := 0; i < 19*166000; i++ {
	//	lines = append(lines, "12345\n")
	//}
	//filePath := "/Users/mac/test_uid_19m.txt"
	//WriteLines(filePath, lines)

	ch := make(chan int)
	select {
	case i := <-ch:
		println(i)

	default:
		println("default")
	}
	fmt.Println("finish")
}
