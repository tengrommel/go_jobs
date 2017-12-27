# 第一章

## 编程来自生活

> 调试程序

- 垃圾回收
>内存自动回收，再也不需要开发人员管理内存<br>
开发人员专注业务实现，降低了心智负担<br>
只需要new分配内存，不需要释放。

- 天然并发
> 从语言层面支持并发，非常简单
<br>goroutine，轻量级线程，创建成千上万个goroutine成为可能
<br>基于CSP(Communicating Sequential Process)模型实现

    func main() {
        go fmt.Println("hello")
    }

运行命令

    go run hello.go test.go
    