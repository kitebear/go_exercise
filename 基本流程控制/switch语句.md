        与串联的if语句类似，switch语句提供了一个多分支条件执行的方法。不过在这里用一个专有名词来代表分支——case。每一个case可以携带一个表达式或一个类型说明符。前者又可被简称为case表达式。因此，Go语言的switch语句又分为表达式switch语句和类型switch语句。
    
        先说表达式switch语句。在此类switch语句中，每个case会携带一个表达式。与if语句中的条件表达式不同，这里的case表达式的结果类型并不一定是bool。不过，它们的结果类型需要与switch表达式的结果类型一致。所谓switch表达式是指switch语句中要被判定的那个表达式。switch语句会依据该表达式的结果与各个case表达式的结果是否相同来决定执行哪个分支。请看下面的示例：

    var name string
    // 省略若干条语句
    switch name {
    case "Golang":
        fmt.Println("A programming language from Google.")
    case "Rust":
        fmt.Println("A programming language from Mozilla.")
    default:
        fmt.Println("Unknown!")
    }  
        可以看到，在上述switch语句中，name充当了switch表达式，而"Go"和"Rust"充当了case表达式。它们的结果类型是一致的，都是string。顺便说一句，可以有只包含一个字面量或标识符的表达式。它们是最简单的表达式，属于基本表达式的一种。
    
        请大家注意switch语句的写法。switch表达式必须紧随switch关键字出现。在后面的花括号中，一个关键字case、case表达式、冒号以及后跟的若干条语句组成为一条case语句。在switch语句中可以有若干条case语句。Go语言会依照从上至下的顺序对每一条case语句中case表达式进行求值。只要被发现其表达式与switch表达式的结果相同，该case语句就会被选中。它包含的那些语句就会被执行。而其余的case语句则会被忽略。
    
        switch语句中还可以存在一个特殊的case——default case。顾名思义，当没有一个常规的case被选中的时候，default case就会被选中。上面示例中就存在一个default case。它由关键字default、冒号和后跟的一条语句组成。实际上，default case不一定被追加在最后。它可以是第一个case，或者出现在任意顺位上。
    
        另外，与if语句一样，switch语句还可以包含初始化子句，且其出现位置和写法也如出一辙。如：

    names := []string{"Golang", "Java", "Rust", "C"}
    switch name := names[0]; name {
    case "Golang":
        fmt.Println("A programming language from Google.")
    case "Rust":
        fmt.Println("A programming language from Mozilla.")
    default:
        fmt.Println("Unknown!")
    }
        好了，我们已经对switch语句的一般形式——表达式switch语句——有所了解了。下面我们来说说类型switch语句。它与一般形式有两点差别。第一点，紧随case关键字的不是表达式，而是类型说明符。类型说明符由若干个类型字面量组成，且多个类型字面量之间由英文逗号分隔。第二点，它的switch表达式是非常特殊的。这种特殊的表达式也起到了类型断言的作用，但其表现形式很特殊，如：v.(type)，其中v必须代表一个接口类型的值。注意，该类表达式只能出现在类型switch语句中，且只能充当switch表达式。一个类型switch语句的示例如下：

    v := 11
    switch i := interface{}(v).(type) {
    case int, int8, int16, int32, int64:
        fmt.Printf("A signed integer: %d. The type is %T. \n", i, i)
    case uint, uint8, uint16, uint32, uint64:
        fmt.Printf("A unsigned integer: %d. The type is %T. \n", i, i)
    default:
        fmt.Println("Unknown!")
    }
        请注意，我们在这里把switch表达式的结果赋给了一个变量。如此一来，我们就可以在该switch语句中使用这个结果了。这段代码被执行后，标准输出上会打印出A signed integer: 11. The type is int.。
    
        最后，我们来说一下fallthrough。它既是一个关键字，又可以代表一条语句。fallthrough语句可被包含在表达式switch语句中的case语句中。它的作用是使控制权流转到下一个case。不过要注意，fallthrough语句仅能作为case语句中的最后一条语句出现。并且，包含它的case语句不能是其所属switch语句的最后一条case语句。