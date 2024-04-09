package us

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafe_1(t *testing.T) {
	i := 32
	t.Logf("size of i : %d\n", unsafe.Sizeof(i))
}

func TestUnsafe_2(t *testing.T) {
	n := struct {
		i string
		j int64
	}{
		i: "EDDYCJY",
		j: 1,
	}

	nPointer := unsafe.Pointer(&n)

	niPointer := (*string)(nPointer)
	*niPointer = "haha"

	// 注：uintptr不能存储在临时变量
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 2

	fmt.Printf("n.i: %s, n.j: %d", n.i, n.j)
}

func TestUnsafe_3(t *testing.T) {
	var f float64 = 3.5
	n := *(*uint64)(unsafe.Pointer(&f))

	fmt.Printf("n = %d\n", n)
}
