# Learn-Java-By-Example

* [learn-go-with-tests](https://quii.gitbook.io/learn-go-with-tests/)
* [learn-go-with-tests-cn](https://studygolang.gitbook.io/learn-go-with-tests/)
* [Go-By-Example](https://gobyexample.com/)

## hello

* string
* if
* switch
* const
* testing

## integers

* ``integers``, `adder`

## iteration

* for

## arrays

* arrays
* slices
    * init slices
    * append

## Structs, methods & interfaces

* structs
* methods
* interfaces

## Pointers & errors

* Pointers 指针
    * 当你传值给函数或方法时，Go 会复制这些值。因此，如果你写的函数需要更改状态，你就需要用指针指向你想要更改的值
    * Go 取值的副本在大多数时候是有效的，但是有时候你不希望你的系统只使用副本，在这种情况下你需要传递一个引用。例如，非常庞大的数据或者你只想有一个实例
* nil
    * 指针可以是 nil
    * 当函数返回一个的指针，你需要确保检查过它是否为 nil，否则你可能会抛出一个执行异常，编译器在这里不能帮到你
    * nil 非常适合描述一个可能丢失的值
* 错误
    * 错误是在调用函数或方法时表示失败的
    * 通过测试我们得出结论，在错误中检查字符串会导致测试不稳定。因此，我们用一个有意义的值重构了，这样就更容易测试代码，同时对于我们
      API 的用户来说也更简单

## maps

* 初始化map，并且使用 `增删改查` 操作map

## di

* 依赖注入

## mocking

* mocking

## Concurrency

* goroutines 是 Go 的基本并发单元，它让我们可以同时检查多个网站
* anonymous functions（匿名函数），我们用它来启动每个检查网站的并发进程
* channels，用来组织和控制不同进程之间的交流，使我们能够避免 race condition（竞争条件） 的问题。
* the race detector（竞争探测器） 帮助我们调试并发代码的问题
