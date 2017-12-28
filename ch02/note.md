# 基本数据类型和操作符

- 文件名&关键字&标识符
> 大写变量可导出
- Go程序结构
- 常量和变量
- 数据类型和操作符
- 字符串类型

## 常量
const 常量使用const修饰，代表永远是只读的，不能修改。<br>
const 只能修饰boolean number string

const identifier [type] = value，其中type可以省略<br>
也可以使用表达式 但不可以通过函数赋值

>推荐写法

    const(
        a = 1
        b = 2
        c = 3
    )

    const(
        a = iota
        b
        c
    )