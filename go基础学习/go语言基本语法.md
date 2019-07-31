# go语言基本语法
Go语言的基本类型有：

bool

string

int、int8、int16、int32、int64

uint、uint8、uint16、uint32、uint64、uintptr //u即无符号

byte // uint8 的别名

rune // int32 的别名 代表一个 Unicode 码

float32、float64

complex64、complex128//复数类型，复数类型的值一般由浮点数表示的实数部分、加号“+”、浮点数表示的虚数部分，以及小写字母“i”组成。比如：*3.7E+1 + 5.98E-2i*
## 声明
**在一个变量被声明后，系统自动赋予它该类型的默认值，int为0,float为0.0，bool为false，string为空字符串，指针为nil

例1，通过var声明：

var a,b \*[]int64,float64//a是一个int64类型切片指针，现在它是nil。b是一个float64类型变量现在它是0.0

例2，通过‘:=’声明：

a,b:=1,1.1

var声明还有另一种形式：
var (
    a int
    b string
    c []float32
    d func() bool
    e struct {
        x int
    }
)
