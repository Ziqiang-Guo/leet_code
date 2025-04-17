package main

import (
	"bufio"
	"fmt"
	"os"
)

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func (doublyLinkedList *DoublyLinkedList) insertNodeIntoDoublyLinkedList(nodeData int32) {
	node := &DoublyLinkedListNode{
		next: nil,
		prev: nil,
		data: nodeData,
	}

	if doublyLinkedList.head == nil {
		doublyLinkedList.head = node
	} else {
		doublyLinkedList.tail.next = node
		node.prev = doublyLinkedList.tail
	}

	doublyLinkedList.tail = node
}

func printDoublyLinkedList(node *DoublyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprintf(writer, sep)
		}
	}
}

/*
 * Complete the 'sortedInsert' function below.
 *
 * The function is expected to return an INTEGER_DOUBLY_LINKED_LIST.
 * The function accepts following parameters:
 *  1. INTEGER_DOUBLY_LINKED_LIST llist
 *  2. INTEGER data
 */

/*
 * For your reference:
 *
 * DoublyLinkedListNode {
 *     data int32
 *     next *DoublyLinkedListNode
 *     prev *DoublyLinkedListNode
 * }
 *
 */

func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	// 创建新节点
	newNode := &DoublyLinkedListNode{
		data: data,
		next: nil,
		prev: nil,
	}

	// 如果链表为空，返回新节点
	if llist == nil {
		return newNode
	}

	// 如果新数据小于头节点数据，插入到头部
	if data < llist.data {
		newNode.next = llist
		llist.prev = newNode
		return newNode
	}

	// 找到合适的插入位置
	current := llist
	for current.next != nil && current.next.data < data {
		current = current.next
	}

	// 插入到链表尾部
	if current.next == nil {
		current.next = newNode
		newNode.prev = current
	} else {
		// 插入到链表中间
		newNode.next = current.next
		newNode.prev = current
		current.next.prev = newNode
		current.next = newNode
	}

	return llist
}

func main() {
	// 创建一个有序的测试链表
	test_list := &DoublyLinkedList{}

	// 插入一些有序数据
	test_list.insertNodeIntoDoublyLinkedList(1)
	test_list.insertNodeIntoDoublyLinkedList(3)
	test_list.insertNodeIntoDoublyLinkedList(5)
	test_list.insertNodeIntoDoublyLinkedList(7)

	// 打印原始链表
	fmt.Println("原始链表:")
	writer := bufio.NewWriter(os.Stdout)
	printDoublyLinkedList(test_list.head, " -> ", writer)
	writer.Flush()
	fmt.Println()

	// 测试在不同位置插入数据
	// 1. 插入到链表头部
	fmt.Println("\n插入 0 到链表头部:")
	newHead := sortedInsert(test_list.head, 0)
	printDoublyLinkedList(newHead, " -> ", writer)
	writer.Flush()
	fmt.Println()

	// 2. 插入到链表中间
	fmt.Println("\n插入 4 到链表中间:")
	newHead = sortedInsert(newHead, 4)
	printDoublyLinkedList(newHead, " -> ", writer)
	writer.Flush()
	fmt.Println()

	// 3. 插入到链表尾部
	fmt.Println("\n插入 9 到链表尾部:")
	newHead = sortedInsert(newHead, 9)
	printDoublyLinkedList(newHead, " -> ", writer)
	writer.Flush()
	fmt.Println()
}
