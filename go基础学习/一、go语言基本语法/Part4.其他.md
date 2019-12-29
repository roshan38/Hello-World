# 第四部分
## 变量逃逸
直接上代码
---------------------------------
package main

import "fmt"

// 声明空结构体测试结构体逃逸情况

type Data struct {
}

func dummy() \*Data {
    var c Data
    return &c
}

func main() {
    fmt.Println(dummy())
}

-----------------------
### 执行逃逸分析：

$ go run -gcflags "-m -l" main.go
//使用 go run 运行程序时，-gcflags 参数是编译参数。其中 -m 表示进行内存分配分析，-l 表示避免程序内联，也就是避免进行程序优化。

\# command-line-arguments

./main.go:15:9: &c escapes to heap

./main.go:12:6: **moved to heap: c**

./main.go:20:19: dummy() escapes to heap

./main.go:20:13: main ... argument does not escape

&{}

注意上方加粗字，因为在函数运行完后，c的地址被返回，c不能被丢弃。因此需要保留c，c就移动到堆里，保证c的地址能被使用。
## 变量的生命周期
生命周期与作用域不同，作用域学C的时候了解过，生命周期可能会更长，就像上一节的例子中，c的生命周期超出了它的作用域。
## 常量和const关键字
go的常量用const定义，只能是bool、数字型（整数、浮点数、复数）和字符串//如果函数返回值是以上类型的也能const

const也可以批量声明：
const (
    a = 1
    b          //b==1
    c = 2
    d          //d==2
)

还有另一种批量声明：iota常量生产器，比较少用，在模拟枚举的时候有用，可以[展阅读一下]https://blog.csdn.net/thisinnocence/article/details/88046825)

**无类型常量**这里看不懂，不知道想表达什么

## 枚举
go语言现阶段没有枚举，可以通过const和iota来模拟枚举：
type Weapon int

const (
     Arrow Weapon = iota    // 开始生成枚举值, 默认为0
     Shuriken               // 1
     SniperRifle            // 2
     Rifle                  // 3
     Blower                 // 4
)
