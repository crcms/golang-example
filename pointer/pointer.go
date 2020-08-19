package pointer

import (
	"fmt"
)

func Point()  {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ",a) // 1
	fmt.Println("&a = ",&a) // point指针地址
	fmt.Println("*&a = ",*&a) // 1  &a => a的指针地址 *&a => 取得a指针地址的值
	fmt.Println("b = ",b) // a 指针地址的值
	fmt.Println("&b = ",&b) // b 值的引用值
	fmt.Println("*&b = ",*&b) // b 值的引用值
	fmt.Println("*b = ",*b) //
	fmt.Println("c = ",c)
	fmt.Println("*c = ",*c)
	fmt.Println("&c = ",&c)
	fmt.Println("*&c = ",*&c)
	fmt.Println("**c = ",**c)
	fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c)
	fmt.Println("x = ",x)
}
