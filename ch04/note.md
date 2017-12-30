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