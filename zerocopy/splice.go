package main

import (
	//syscall "golang.org/x/sys/unix"
	"os"
)

func main() {
	// 设置源文件
	src, err := os.Open("/tmp/source.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()
	// 设置目标文件
	target, err := os.Create("/tmp/target.txt")
	if err != nil {
		panic(err)
	}
	defer target.Close()

	// 创建管道文件
	// 作为两个文件传输数据的中介
	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	defer pipeReader.Close()
	defer pipeWriter.Close()

	// 设置文件读写模式
	// 笔者在标准库中没有找到对应的常量说明
	// 读者可以参考这个文档:
	//   https://pkg.go.dev/golang.org/x/sys/unix#pkg-constants
	//   SPLICE_F_NONBLOCK = 0x2
	//spliceNonBlock := 0x02
	//// 使用 Splice 将数据从源文件描述符移动到管道 writer
	//_, err = syscall.Splice(int(src.Fd()), nil, int(pipeWriter.Fd()), nil, 1024, spliceNonBlock)
	//if err != nil {
	//	panic(err)
	//}
	//// 使用 Splice 将数据从管道 reader 移动到目标文件描述符
	//n, err := syscall.Splice(int(pipeReader.Fd()), nil, int(target.Fd()), nil, 1024, spliceNonBlock)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("写入字节数: %d", n)
}
