package utils
//实现了一个迭代器，它用于遍历容器中的部分元素和全部元素

//输入迭代器
type InputIterator interface{
	Next() bool
	Value() interface{}
	Index() int
	Begin()
	First()
}
//输出迭代器
type OutputIterator interface{}
//正向迭代器
type ForwardIterator interface{
	InputIterator
}
//双向迭代器：支持双向移动和多次读写操作
type BidIterator interface{
	ForwardIterator
}
//随机访问迭代器，支持双向移动，并且支持多次读写操作
type RandomIterator interface{
	BidIterator
}


