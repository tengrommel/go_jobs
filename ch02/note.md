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

    var (
        a = 1
        b = 3
    )

## 值类型和引用类型

1.值类型：变量存储值，内存通常在栈中分配。 int float bool string struct
2.引用类型：变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配。通过GC回收。 指针 slice map chan

## 变量的作用域
1.在函数内部
2.在函数外部作用于整个包，大写为整个程序

## 类型转换

type(variable)

    var a int = 8;
    var b int32=int32(a)