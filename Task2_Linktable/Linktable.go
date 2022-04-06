package main

import (
	"errors"
	"fmt"
	"sync"
)

//结构体定义
//结点：LinkNode
type LinkNode struct {
	data interface{}
	next *LinkNode
}
//链表：LinkTable
type LinkTable struct {
	pHead *LinkNode
	length int
	mutex sync.Mutex
}

//函数
//创建带头结点的空链表
func CreateLinkTable() *LinkTable {
	pLinkTable := new(LinkTable)
	pLinkTable.pHead = new(LinkNode)
	return pLinkTable
}

//删除链表
func (p *LinkTable) clear()  {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.pHead.next = nil
	p.length = 0
}

//插入结点至链表头部
func (p *LinkTable) AddNodeAtHead(data ...interface{}) {
	for _, val := range data{
		//准备好结点
		node := LinkNode{val, nil}
		node.next = p.pHead.next
		//修改链表前加锁
		p.mutex.Lock()
		p.pHead.next = &node
		p.length++
		//完成修改后释放锁
		p.mutex.Unlock()
	}
}

//判断链表是否为空链表
func (p *LinkTable) IsEmpty() bool {
	if p.length == 0{
		return true
	} else {
		return false
	}
}

//删除第一个值为data的结点
func (p *LinkTable) DelNode(delVal interface{}) (err error){
	if p.IsEmpty() {
		return errors.New("链表为空")
	}
	for first,second:=p.pHead,p.pHead.next; second!=nil; first,second=second,second.next{
		if second.data == delVal {
			p.mutex.Lock()
			first.next = second.next
			p.length--
			p.mutex.Unlock()
			return nil
		}
	}
	return errors.New("链表中不存在该值对应的结点")
}

//获取头结点
func (p *LinkTable) GetLinkTableHead() (firstNode *LinkNode, err error){
	if p.IsEmpty() {
		firstNode = nil
		err = errors.New("链表为空")
		return
	}
	return p.pHead.next, nil
}

//获取下一个结点
func (p *LinkTable) GetNextLinkTableNode(pNode *LinkNode) (*LinkNode, error) {
	if pNode.next == nil {
		return nil, errors.New("不存在下一个结点")
	}
	return pNode.next, nil
}

//查找是否存在某值对应的结点
func (p *LinkTable) FindNode(data interface{}) (*LinkNode, error) {
	if p.IsEmpty() {
		return nil, errors.New("链表为空")
	}
	for q:=p.pHead.next; q!=nil; q=q.next{
		if q.data == data {
			return q, nil
		}
	}
	return nil, errors.New("链表中不存在该值对应结点")
}

//输出链表中全部结点
func (p *LinkTable) PrintLinkTable() error {
	if p.IsEmpty() {
		return errors.New("链表为空")
	}
	for q:=p.pHead.next; q!=nil; q=q.next{
		fmt.Printf("%v\t", q.data)
	}
	fmt.Println()
	return nil
}

func main() {
	list := CreateLinkTable()
	list.AddNodeAtHead(1,3,5,7,9)
	if list.PrintLinkTable() != nil {
		fmt.Println(errors.New("链表为空").Error())
	}

	list.DelNode(5)
	if list.PrintLinkTable() != nil {
		fmt.Println(errors.New("链表为空").Error())
	}

	_, ok := list.FindNode(3)
	if ok == nil {
		fmt.Println("查找成功")
	}

	list.clear()
	if list.PrintLinkTable() != nil {
		fmt.Println(errors.New("链表为空").Error())
	}
}














