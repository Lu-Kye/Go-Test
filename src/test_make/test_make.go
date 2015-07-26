package test_make

import (
	"fmt"
	"runtime"
	"unsafe"
)

func TestMake() {
	arr := make([]int, 0, 1000)
	fmt.Printf("arr: %T, %d\n", arr, unsafe.Sizeof(arr))

	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	fmt.Println(m.StackInuse)
	fmt.Println(m.HeapInuse)

	a := int(123)
	b := int64(123)
	c := "foo"
	d := struct {
		FieldA float32
		FieldB string
		FieldC int
	}{0, "bar", 0}

	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("d: %T, %d\n", d, unsafe.Sizeof(d))
	fmt.Printf("m: %T, %d\n", m, unsafe.Sizeof(m))
}
