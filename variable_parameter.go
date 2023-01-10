package variable_parameter

// TakeFirstParam 获取传递的可变参数，如果没有传递的话，则返回参数对应的默认值
func TakeFirstParam[T any](parameters []T) T {
	if len(parameters) > 0 {
		return parameters[0]
	} else {
		var zero T
		return zero
	}
}

// TakeFirstParamOrDefault 获取传递的可变参数，如果没有传的话则使用给定的默认值
func TakeFirstParamOrDefault[T any](parameters []T, defaultValue T) T {
	if len(parameters) > 0 {
		return parameters[0]
	} else {
		return defaultValue
	}
}

// TakeFirstParamOrDefaultFunc 获取传递的可变参数，如果没有传的话则执行给定的defaultValueFunc来获取默认值，如果传递的话defaultValueFunc函数不会被执行
func TakeFirstParamOrDefaultFunc[T any](parameters []T, defaultValueFunc func() T) T {
	if len(parameters) > 0 {
		return parameters[0]
	} else {
		return defaultValueFunc()
	}
}

// SetDefaultParam 如果传递了参数啥都不会做，如果没有传递的话则把给定的defaultValue放到parameters中，这样后面就可以直接使用parameters[0]的形式使用参数而不必担心越界
func SetDefaultParam[T any](parameters []T, defaultValue T) []T {
	if len(parameters) != 0 {
		return parameters
	}
	return append(parameters, defaultValue)
}

// SetDefaultParamByFunc 如果传递了参数啥都不会做，如果没有传递的话则执行defaultValueFunc来获取一个默认值，并把得到的默认值放到parameters中，这样后面就可以直接使用parameters[0]的形式使用参数而不必担心越界
func SetDefaultParamByFunc[T any](parameters []T, defaultValueFunc func() T) []T {
	if len(parameters) != 0 {
		return parameters
	}
	return append(parameters, defaultValueFunc())
}
