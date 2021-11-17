package quene
//实现了队列数据结构,继承自容器
import(
	container"go-stl/container"
	arraylist"go-stl/arraylist"
)
type quene interface{
	container.Container
	Push()(e interface{})//入队列
	Pop()(e interface{})//出队列
	Front()(e interface{})//获取队列的头元素
	End()(e interface{})//获取队列的尾元素
}