package basic

import (
	"reflect"
	"strconv"
)

// 定义抽象类跟类
type Object interface{}

type Numbers interface {
	~byte | ~int8 | ~int16 | ~uint16 | ~int | ~uint | ~int32 | ~uint32 | ~int64 | ~uint64
}

type Enum interface {
	~int8
}

type ComObj interface {
	comparable
}

// 执行
type Runnable interface {
	Run() int32
}

type HashCodeAble interface {
	HashCode() int32
}

func NullT[T Object]() T {
	return *new(T)
}

/**
 * 比较两个objec是否相同
 */
func ObjectEqual(a, b Object) bool {
	return reflect.DeepEqual(a, b)
}

// 转换num为 string
func NumString[T Numbers](t T) string {
	return strconv.FormatInt(int64(t), 10)
}

// 针对基础类的HashCode，返回int32的code值
func HashCode(a ...Object) int32 {
	if a == nil {
		return 0
	}
	result := int32(1)
	for _, element := range a {
		if element == nil {
			result = 31*result + 0
		} else {
			result = 31*result + ObjectHashCode(element)
		}
	}
	return result
}
