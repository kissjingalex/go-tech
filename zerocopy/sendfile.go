package main

import (
	"fmt"
	"os"
	"syscall"
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

	// 获取源文件的文件描述符
	srcFd := int(src.Fd())
	fmt.Printf("srcFd=%d\n", srcFd)
	// 获取目标文件的文件描述符
	targetFd := int(target.Fd())
	fmt.Printf("targetFd=%d\n", targetFd)

	// 使用 Sendfile 实现零拷贝 (拷贝 10 个字节)
	// 如果因为字符编码导致的字符截断问题 (如中文乱码问题), 结果自动保留到截断前的最后完整字节
	// 例如文件内容为 “星期三四五六七”，count 参数为 4, 那么只会拷贝第一个字 (一个汉字 3 个字节)
	// 但是需要注意的是，方法的返回值 written 不受影响 (和 count 参数保持一致)
	// 所以实际开发中，第三个参数 offset 必须设置正确，否则就可能引起乱码或数据丢失问题
	// 注：offset参数不能为nil
	offset := int64(0)
	n, err := syscall.Sendfile(targetFd, srcFd, &offset, 4)
	if err != nil {
		fmt.Printf("error=%v\n", err) // socket operation on non-socket
		return
	}
	fmt.Printf("写入字节数: %d", n)
}
