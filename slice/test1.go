package main

import "fmt"

/*
*
问题点：
1）slice的扩容
2）逃逸分析

新版本sdk已经解决了逃逸的问题了
https://www.cnblogs.com/mushroom/p/8998538.html
*/
func main() {
	s := []byte("")
	s1 := append(s, 'a')
	s2 := append(s, 'b')
	fmt.Println(s1, "==========", s2)
	fmt.Println(string(s1), "==========", string(s2))
}
