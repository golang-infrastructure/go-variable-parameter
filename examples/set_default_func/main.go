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
