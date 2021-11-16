//arraylist数据结构的实现
package arraylist

import (
	container "go-stl/container"
)

//定义arraylist数据结构
type Arraylist struct{
	elements []interface{}
	size int
}
type Arraylister interface{
	container.Container
	Add()
}
// func(arraylist *Arraylist)Remove(index int){
// 	if !arraylist.withinRange(index){
// 		return
// 	}
// 	arraylist.elements[index]=nil
// 	copy(arraylist.elements[index:],arraylist.elements[index+1:arraylist.size-1])
// 	arraylist.size--

// }
func(arraylist *Arraylist)Add(values ...interface{}){
	//为即将插入的元素分配内存
	arraylist.Allocate(len(values))
	for _,value :=range values{
		arraylist.elements[arraylist.size]=value
		arraylist.size++
	}
}
func(arraylist *Arraylist)IndexOf(value interface{})int{
	if arraylist.size==0{
		return -1
	}
	//独特的for循环，可以记住这种语法
	for index,element:=range arraylist.elements{
		if element ==value{
			return index
		}
	}
	return -1//表示没找到
}
//交换两个元素的值
func(arraylist *Arraylist)Swap(i int ,j int ){
	if arraylist.withinRange(i)&&arraylist.withinRange(j){
		arraylist.elements[i],arraylist.elements[j]=arraylist.elements[j],arraylist.elements[i]
	}
}
//插入值
func(arraylist *Arraylist)Insert(index int, values ...interface{}){
	if !arraylist.withinRange(index){
		if index==arraylist.size{
			arraylist.Add(values...)
		}
		return
	}
	l:=len(values)
	arraylist.Allocate(l)
	copy(arraylist.elements[index+l:],arraylist.elements[index:arraylist.size-1])
	copy(arraylist.elements[index:index+l],values)
}
//set的方式设定初值
func(arraylist *Arraylist)Set(index int ,value interface{}){
	if !arraylist.withinRange(index){
		if index==arraylist.size{
			arraylist.Add(value)
		}
	}
	arraylist.elements[index]=value
}
//判断索引位置是否正确
func(arraylist *Arraylist)withinRange(i int)bool{
	return i>=0&&i<arraylist.size
}
//返回所有值
func(arraylist *Arraylist)Values()[]interface{}{
	newElements:=make([]interface{},arraylist.size)
	copy(newElements,arraylist.elements[:arraylist.size])
	return newElements
}
func(arraylist *Arraylist)Clear(){
	arraylist.size=0
	arraylist.elements=[]interface{}{}
}
func(arraylist *Arraylist)Empty()bool{
	return arraylist.size==0
}
func(arraylist*Arraylist)Size()int{
	return arraylist.size
}
//扩容数组
func(arraylist *Arraylist)Allocate(n int){
	currentCapacity:=cap(arraylist.elements)
	if arraylist.size+n>=currentCapacity{
		newCapacity:=int(2.0*float32(currentCapacity+n))
		arraylist.resize(newCapacity)
	}
}
//重新填充数组
func(arraylist *Arraylist)resize(cap int){
	//重新分配内存空间
	newElements:=make([]interface{},cap,cap)
	//将原来的数组元素复制其中
	copy(newElements,arraylist.elements)
	arraylist.elements=newElements
}
