package variable_parameter

// TakeFirstParam 获取可变参数的第一个参数
func TakeFirstParam[T any](parameters []T) T {
	if len(parameters) > 0 {
		return parameters[0]
	} else {
		var zero T
		return zero
	}
}

// TakeFirstParamOrDefault 获取可变参数的第一个参数，如果没有传的话则使用默认值
func TakeFirstParamOrDefault[T any](parameters []T, defaultValue T) T {
	if len(parameters) > 0 {
		return parameters[0]
	} else {
		return defaultValue
	}
}
