# Go可变参数

# 一、这是什么
这个库为Go中的可变参数提供了一些辅助方法，以便更爽的使用可变参数。

# 二、安装
```bash
go get -u github.com/golang-infrastructure/go-variable-parameter
```

# 三、Example 

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

var DefaultFooOptions = FooOptions{
	Foo: "default foo",
	Bar: 10,
}

func Foo(optionsVariableParams ...FooOptions) {
	// 如果传递了options则使用传递的，如果没传递则使用默认的
	options := variable_parameter.TakeFirstParamOrDefault(optionsVariableParams, DefaultFooOptions)
	fmt.Println(options.Foo)
}

func main() {

	// 不传递参数
	Foo() // Output: default foo

	// 传递参数
	Foo(FooOptions{Foo: "custom foo"}) // Output: custom foo

}
```


