        go语句和通道类型是Go语言的并发编程理念的最终体现。在第五章，我已经详细介绍过了通道类型。相比之下，go语句在用法上要比通道简单很多。与defer语句相同，go语句也可以携带一条表达式语句。注意，go语句的执行会很快结束，并不会对当前流程的进行造成阻塞或明显的延迟。一个简单的示例如下：

    go fmt.Println("Go!")
        可以看到，go语句仅由一个关键字go和一条表达式语句构成。同样的，go语句的执行与其携带的表达式语句的执行在时间上没有必然联系。这里能够确定的仅仅是后者会在前者完成之后发生。在go语句被执行时，其携带的函数（也被称为go函数）以及要传给它的若干参数（如果有的话）会被封装成一个实体（即Goroutine），并被放入到相应的待运行队列中。Go语言的运行时系统会适时的从队列中取出待运行的Goroutine并执行相应的函数调用操作。注意，对传递给这里的函数的那些参数的求值会在go语句被执行时进行。这一点也是与defer语句类似的。
    
        正是由于go函数的执行时间的不确定性，所以Go语言提供了很多方法来帮助我们协调它们的执行。其中最简单粗暴的方法就是调用time.Sleep函数。请看下面的示例：

    package main

    import (
        "fmt"
    )

    func main() {
        go fmt.Println("Go!")
    }  
        这样一个命令源码文件被运行时，标准输出上不会有任何内容出现。因为还没等Go语言运行时系统调度那个go函数执行，主函数main就已经执行完毕了。函数main的执行完毕意味着整个程序的执行的结束。因此，这个go函数根本就没有执行的机会。
    
    但是，当我们在上述go语句的后面添加一条对time.Sleep函数的调用语句之后情况就会不同了：

    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        go fmt.Println("Go!")
        time.Sleep(100 * time.Millisecond)
    }
        语句time.Sleep(100 * time.Millisecond)会把main函数的执行结束时间向后延迟100毫秒。100毫秒虽短暂，但足够go函数被调度执行的了。上述命令源码文件在被运行时会如我们所愿地在标准输出上打印出Go!。

    
        另一个比较绅士的做法是在main函数的最后调用runtime.Gosched函数。相应的程序版本如下：

    package main

    import (
        "fmt"
        "runtime"
    )

    func main() {
        go fmt.Println("Go!")
        runtime.Gosched()
    }
    runtime.Gosched函数的作用是让当前正在运行的Goroutine（这里是运行main函数的那个Goroutine）暂时“休息”一下，而让Go运行时系统转去运行其它的Goroutine（这里是与go fmt.Println("Go!")对应并会封装fmt.Println("Go!")的那个Goroutine）。如此一来，我们就更加精细地控制了对几个Goroutine的运行的调度。
    
        当然，我们还有其它方法可以满足上述需求。并且，如果我们需要去左右更多的Goroutine的运行时机的话，下面这种方法也许更合适一些。请看代码：

    package main

    import (
        "fmt"
        "sync"
    )

    func main() {
        var wg sync.WaitGroup
        wg.Add(3)
        go func() {
            fmt.Println("Go!")
            wg.Done()
        }()
        go func() {
            fmt.Println("Go!")
            wg.Done()
        }()
        go func() {
            fmt.Println("Go!")
            wg.Done()
        }()
        wg.Wait()
    }
        sync.WaitGroup类型有三个方法可用——Add、Done和Wait。Add会使其所属值的一个内置整数得到相应增加，Done会使那个整数减1，而Wait方法会使当前Goroutine（这里是运行main函数的那个Goroutine）阻塞直到那个整数为0。这下你应该明白上面这个示例所采用的方法了。我们在main函数中启用了三个Goroutine来封装三个go函数。每个匿名函数的最后都调用了wg.Done方法，并以此表达当前的go函数会立即执行结束的情况。当这三个go函数都调用过wg.Done函数之后，处于main函数最后的那条wg.Wait()语句的阻塞作用将会消失，main函数的执行将立即结束。