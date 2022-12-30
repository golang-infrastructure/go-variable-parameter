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
	Foo()

	// 传递参数
	Foo(FooOptions{Foo: "custom foo"})

}
