package main

import (
	"fmt"
)

func main{
	//声明一个含10个byte类型变量的数组，并为前五个变量赋值
	var arr [10]byte
	arr = [10]byte{'a', 'b', 'c', 'd', 'e'}
	for i := 0; i < 5; i++ {
		fmt.Printf("%c", arr[i])
	}
    //把第3位移动至第6位，并把x替换第3位，完成插入操作
	arr[5] = arr[2]
	arr[2] = 'x'
	for i := 0; i < 6; i++ {
		fmt.Printf("%c", arr[i])
	}
	//把第6位移动至第2位，删除第6位，完成删除操作
	arr[1] = arr[5]
	arr[5] = ''
}
//结果有一个问题，22行的arr[5]无法改变为空字符
