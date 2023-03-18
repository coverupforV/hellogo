# day1

## 初见Go

```Go  
  package main


  import "fmt"


  func main() {
          /* 简单的程序 万能的hello world */
          fmt.Println("Hello Go")
  }
```

go run 表示 直接编译go语言并执行应用程序，一步完成

+ 第一行代码package main定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

+ 下一行import "fmt"告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。

+ 下一行func main()是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。

注意：这里面go语言的语法，定义函数的时候，‘{’ 必须和函数名在同一行，不能另起一行。

+ 下一行 /.../ 是注释，在程序执行时将被忽略。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 / 开头，并以 / 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。

+ 下一行fmt.Println(...)可以将字符串输出到控制台，并在最后自动增加换行字符 \n。 使用 fmt.Print("hello, world\n") 可以得到相同的结果。 Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。

## 变量声明

声明变量的一般形式是使用 var 关键字

第一种，指定变量类型，声明后若不赋值，使用默认值0。

```Go
var v_name v_type
v_name = value
package main


import "fmt"


func main() {
        var a int
        fmt.Printf(" = %d\n", a)
}

运行：
$go run test.go
a = 0
```

第二种，根据值自行判定变量类型。

```Go
var v_name = value
```

第三种，省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。

```Go
v_name := value


// 例如
var a int = 10
var b = 10
c : = 10
```

例如:

```Go
package main


import "fmt"


func main() {
        //第一种 使用默认值
        var a int
        fmt.Printf("a = %d\n", a)


        //第二种
        var b int = 10
        fmt.Printf("b = %d\n", b)


        //第三种 省略后面的数据类型,自动匹配类型
        var c = 20
        fmt.Printf("c = %d\n", c)


        //第四种 省略var关键字
        d := 3.14
        fmt.Printf("d = %f\n", d)
}
```

多变量声明

```Go
package main


import "fmt"


var x, y int
var ( //这种分解的写法,一般用于声明全局变量
        a int
        b bool
)


var c, d int = 1, 2
var e, f = 123, "liudanbing"


//这种不带声明格式的只能在函数体内声明
//g, h := 123, "需要在func函数体内实现"


func main() {
        g, h := 123, "需要在func函数体内实现"
        fmt.Println(x, y, a, b, c, d, e, f, g, h)


        //不能对g变量再次做初始化声明
        //g := 400


        _, value := 7, 5  //实际上7的赋值被废弃，变量 _  不具备读特性
        //fmt.Println(_) //_变量的是读不出来的
        fmt.Println(value) //5
}
```

## 常量

常量是一个简单值的标识符，在程序运行时，不会被修改的量。

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：

```Go
const identifier [type] = value
```

你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。

显式类型定义：

```Go
const b string = "abc"
```

隐式类型定义：

```Go
const b = "abc"
```
