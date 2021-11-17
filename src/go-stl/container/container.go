package container
import(
	Utils"go-stl/utils"
)
//实现container接口,这是容器的基础接口,其他的容器继承这个容器
type Container interface{
	Iterator()(i* Utils.InputIterator)//创建一个迭代器
	Empty()bool//判断容器是否为空
	Size()int//返回容器中元素的个数
	Clear()//清除所有元素
	Values() []interface{}//返回所有元素
}