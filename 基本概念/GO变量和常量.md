
我们在这里需要优先说明的是用于声明变量的关键字var，以及用于声明常量的关键字const。要知道，绝大多数的数据类型的值都可以被赋给一个变量，包括函数。而常量则不同，它只能被赋予基本数据类型的值本身。

变量和常量在声明方式方面也有所不同。我们可以在声明一个变量的时候直接为它赋值，也可以只声明不赋值。变量的声明并赋值方式如下：

```
    // 注释：普通赋值，由关键字var、变量名称、变量类型、特殊标记=，以及相应的值组成。
    // 若只声明不赋值，则去除最后两个组成部分即可。
    var num1 int = 1


    // 注释：平行赋值
    var num2, num3 int = 2, 3


    // 注释：多行赋值
    var (
        num4 int = 4
        num5 int = 5
    )
```

 上述这三种变量声明的方式，也适用于常量。但是要注意，对于常量不能出现只声明不赋值的情况。