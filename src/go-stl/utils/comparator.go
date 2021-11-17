package utils

//实现了一个比较器,类型为函数类型，本代码可以用于练习流程控制相关的知识点，这部分也需要掌握GO语言所有的数据类型
//具体用法如下所示
/*
if a>b return 1
a=b return 0
a<b return -1
*/
type Comparator func(a, b interface{}) int

func Cmp(a, b interface{}) int {
	switch a.(type) {
	case int:
		return CmpInt(a, b)
	case int8:
		return CmpInt8(a, b)
	case int16:
		return CmpInt16(a, b)
	case int32:
		return CmpInt32(a, b)
	case int64:
		return CmpInt64(a, b)
	case uint:
		return CmpUint(a, b)
	case uint8:
		return CmpUint8(a, b)
	case uint16:
		return CmpUint16(a, b)
	case uint32:
		return CmpUint32(a, b)
	case string:
		return CmpString(a, b)
	case bool:
		return CmpBool(a, b)
	case float32:
		return CmpFloat32(a, b)
	case float64:
		return CmpFloat64(a, b)
	case complex64:
		return CmpComplex64(a, b)
	case complex128:
		return CmpComplex128(a, b)
	case uintptr:
		return CmpUintptr(a, b)
	}

	return 1
}

//系统自带类型的比较器
func CmpBool(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(bool) == true && b.(bool) == false {
		return 1
	}
	return -1
}

//Uint类型比较器
func CmpUint8(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint8) > b.(uint8) {
		return 1
	}
	return -1

}

func CmpUint16(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint16) > b.(uint16) {
		return 1
	}
	return -1

}

func CmpUint32(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint32) > b.(uint32) {
		return 1
	}
	return -1

}

func CmpUint64(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint64) > b.(uint64) {
		return 1
	}
	return -1

}

//Int类型比较器
func CmpInt64(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int64) > b.(int64) {
		return 1
	}
	return -1

}

func CmpInt32(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int32) > b.(int32) {
		return 1
	}
	return -1

}

func CmpInt16(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int16) > b.(int16) {
		return 1
	}
	return -1

}

func CmpInt8(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int8) > b.(int8) {
		return 1
	}
	return -1

}

//float类型比较器
func CmpFloat64(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(float64) > b.(float64) {
		return 1
	}
	return -1
}

func CmpFloat32(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(float32) > b.(float32) {
		return 1
	}
	return -1
}

//字符串类型比较器
func CmpString(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(string) > b.(string) {
		return 1
	}
	return -1
}

//Complex类型比较器
func CmpComplex64(a, b interface{}) int {
	if a == b {
		return 0
	}
	//获取实数部分
	comA := a.(complex64)
	comB := b.(complex64)
	//比较实数部分
	if real(comA) > real(comB) {
		return 1
	}
	//如果实数部分相同，只比较虚数部分
	if real(comA) == real(comB) && imag(comA) > imag(comB) {
		return 1
	}
	return -1
}
func CmpComplex128(a, b interface{}) int {
	if a == b {
		return 0
	}
	//获取实数部分
	comA := a.(complex128)
	comB := b.(complex128)
	//比较实数部分
	if real(comA) > real(comB) {
		return 1
	}
	//如果实数部分相同，只比较虚数部分
	if real(comA) == real(comB) && imag(comA) > imag(comB) {
		return 1
	}
	return -1
}
func CmpUintptr(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uintptr) > b.(uintptr) {
		return 1
	}
	return -1
}

func CmpInt(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(int) > b.(int) {
		return 1
	}
	return -1

}
func CmpUint(a, b interface{}) int {
	if a == b {
		return 0
	}
	if a.(uint) > b.(uint) {
		return 1
	}
	return -1
}
