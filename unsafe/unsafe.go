package main

import (
	"fmt"
	"unsafe"
)

// unsafe的主要作用：指针类型转换
// golang 不支持显示的类型指针转换，如不能强制把 *int 转成 *float64，而通过 unsafe.Pointer 可以进行不同指针类型转换

// unsafe的4个规则
// 1. 任何指针类型都可以转换为unsafe.Pointer，如：unsafe.Pointer(&foo)
// 2. unsafe.Pointer都可以转换为任何指针类型，如：(*float64)(unsafe.Pointer(&foo))
// 3. uintptr可以转换为unsafe.Pointer
// 4. unsafe.Pointer可以转换uintptr

func ChangeType()  {
	var (
		foo int
		bar *int // int指针类型地址的值
	)
	foo = 10
	bar = &foo // bar 存储的是foo的指针地址

	// 将int强改成float64
	var baz *float64
	baz = (*float64)(unsafe.Pointer(bar)) // unsafe.Pointer 可以转换为任何类型
	fmt.Printf("%T\n", baz)

	// 任何指针类型可以转换为unsafe.Pointer
	v := unsafe.Pointer(bar)
	fmt.Printf("%T\n",v)
	fmt.Printf("%v",v)
}
type SizeOf struct {
	Size int //64位操作系统8位
	Size1 int //64位操作系统8位
	Raw string
}
func Offset()  {
	var foo int32 = 10
	var s string = `10`
	var s2 string
	//var bar = 20
	//var fooPoint = uintptr(&foo)
	// (unsafe.Offsetof(&bar) + )

	// sizeof 只返回数据类型的大小，不管实际引用的数据类型的大小
	fmt.Println(unsafe.Sizeof(&foo))
	fmt.Println(unsafe.Sizeof(&s))
	fmt.Println(unsafe.Sizeof(&s2))
	fmt.Println(unsafe.Sizeof(&SizeOf{}))
}

func main()  {
	//ChangeType()
	Offset()
}