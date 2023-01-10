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
