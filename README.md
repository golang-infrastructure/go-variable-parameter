# Go可变参数

# 一、这是什么
这个库为Go中的可变参数提供了一些辅助方法，以便更爽的使用可变参数。

# 二、安装
```bash
go get -u github.com/golang-infrastructure/go-variable-parameter
```

# 三、Example 
## 3.1 TakeFirstParamOrDefault
```go
package main

import (
	"fmt"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
)

type FooOptions struct {
	Foo string
	Bar int
}

var DefaultFooOptions = &FooOptions{
	Foo: "default foo",
	Bar: 10,
}

func Foo(optionsVariableParams ...*FooOptions) {
	// 如果传递了options则使用传递的取出数组中的第一个元素返回，如果没传递则使用给出的默认值，适合默认值是一个固定的值的时候用
	options := variable_parameter.TakeFirstParamOrDefault(optionsVariableParams, DefaultFooOptions)

	// 后面的代码就可以直接使用options来操作啦
	fmt.Println(options.Foo)

}

func main() {

	// 不传递参数
	Foo() // Output: default foo

	// 传递参数
	Foo(&FooOptions{Foo: "custom foo"}) // Output: custom foo

}
```


## 3.2 TakeFirstParamOrDefaultFunc
```go
package main

import (
	"fmt"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"time"
)

type FooOptions struct {
	Foo string
	Bar int
}

func NewFooOptions() *FooOptions {

	// 假装有耗时操作
	time.Sleep(time.Second * 1)

	return &FooOptions{
		Foo: "default foo",
		Bar: 10,
	}
}

func Foo(optionsVariableParams ...*FooOptions) {
	// 如果后面可能会涉及到对options的修改之类的，则options无法使用单例，可能得每次都重新创建一个新的，则可以使用一个默认值的函数，仅在需要默认值的时候才运行
	options := variable_parameter.TakeFirstParamOrDefaultFunc(optionsVariableParams, func() *FooOptions {
		// 如果初始化比较耗时或者觉得仅在必要的时候才创建比较好可以用这种默认函数的方式，此函数仅在未传递参数时运行
		return NewFooOptions()
	})

	// 后面的代码就可以直接使用options来操作啦
	fmt.Println(options.Foo)
}

func main() {

	// 传递参数
	start := time.Now()
	Foo(&FooOptions{Foo: "custom foo"})
	cost := time.Now().Sub(start)
	fmt.Println("传递参数时耗时：" + cost.String())
	// Output:
	// custom foo
	// 传递参数时耗时：0s

	// 不传递参数
	start = time.Now()
	Foo()
	cost = time.Now().Sub(start)
	fmt.Println("不传递参数时耗时：" + cost.String())
	// Output:
	// default foo
	// 不传递参数时耗时：1.0136631s

}
```


## 3.3 SetDefaultParam 
```go
package main

import (
	"fmt"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
)

type FooOptions struct {
	Foo string
	Bar int
}

var DefaultFooOptions = &FooOptions{
	Foo: "default foo",
	Bar: 10,
}

func Foo(options ...*FooOptions) {
	// 如果没有传递参数的话，则设置一个默认的参数
	options = variable_parameter.SetDefaultParam(options, DefaultFooOptions)

	// 为什么不使用这种方式来操作呢？也许这样更简单？
	// 在递归互相调用的时候切片中可能会被重复放入默认值，尤其是在有很多的重载之类的或者高内聚的代码中会互相调用问题尤其明显
	//options = append(options, DefaultFooOptions)

	// 后面的代码就可以直接使用options[0]来操作而不必担心越界
	fmt.Println(options[0].Foo)
}

func main() {

	// 传递参数
	Foo(&FooOptions{Foo: "custom foo"}) // Output: custom foo

	// 不传递参数
	Foo() // Output: default foo

}
```


## 3.4 SetDefaultParamByFunc
```go
package main

import (
	"fmt"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"time"
)

type FooOptions struct {
	Foo string
	Bar int
}

func NewFooOptions() *FooOptions {

	// 假装有耗时操作
	time.Sleep(time.Second * 1)

	return &FooOptions{
		Foo: "default foo",
		Bar: 10,
	}
}

func Foo(options ...*FooOptions) {
	// 如果后面可能会涉及到对options[0]的修改之类的，则options[0]无法使用单例，可能得每次都重新创建一个新的，则可以使用一个默认值函数，仅在未传递参数的时候才会执行
	options = variable_parameter.SetDefaultParamByFunc(options, func() *FooOptions {
		// 如果初始化比较耗时或者觉得仅在必要的时候才创建比较好可以用这种默认函数的方式，此函数仅在未传递参数时运行
		return NewFooOptions()
	})

	// 后面的代码就可以直接使用options[0]来操作啦
	fmt.Println(options[0].Foo)
}

func main() {

	// 传递参数
	start := time.Now()
	Foo(&FooOptions{Foo: "custom foo"})
	cost := time.Now().Sub(start)
	fmt.Println("传递参数时耗时：" + cost.String())
	// Output:
	// custom foo
	// 传递参数时耗时：0s

	// 不传递参数
	start = time.Now()
	Foo()
	cost = time.Now().Sub(start)
	fmt.Println("不传递参数时耗时：" + cost.String())
	// Output:
	// default foo
	// 不传递参数时耗时：1.0136631s

}
```



