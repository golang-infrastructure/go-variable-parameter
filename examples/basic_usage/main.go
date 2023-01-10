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
