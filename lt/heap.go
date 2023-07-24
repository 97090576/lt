package main

// 堆是一种特殊的完全二叉树，而完全二叉树指除了最后一层都塞满的二叉树，最后一层也是从左到右塞的，所以可以使用数据来存储完全二叉树，不需要使用树节点，可以省去大量的指针存储，只需要存储值就行，而每个节点的子节点的索引位置是确定的，故可以直接通过索引访问

type Heap interface {
	Push(interface{})
	Pop() interface{}
	Pick() interface{}
}
