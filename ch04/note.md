## 内置函数

- close:主要用来关闭channel
- len:用来求长度，比如string、array、slice、map、channel
- new:用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针。
- make:用来分配内存，主要用来分配引用类型，比如chan、map、slice。
- append:用来追加元素到数组、slice中
- panic和recover:用来做错误处理

> new和make的区别：new返回一个地址 make返回一个实例

## 递归函数

## 闭包

## 数组
> 同一种数据类型的固定长度的序列<br>var a [len]int <br>var a [5]int<br>长度是数组类型的一部分
*访问越界触发panic*
*数组是值类型，因此改变副本不会改变原始的值*

## 切片

- 切片：数组的一个引用，因此切片是引用类型
- 切片的长度可以改变，因此，切片是一个可变的数组
- 切片遍历方式和数组一样，可以用len()求长度
- cap可以求出slice最大的容量，0<=cap(slice)<=len(array)，其中array是slice引用的数组
- 切片的定义：var变量名 []类型，比如var str []string var arr []int

*string底层就是一个byte数组，但不可变<br>要改变需要转换为byte(中文)*